// Copyright 2025 EURECOM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/giuliocarot0/gitc"
	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
)

type Amf struct {
	PlmnId        models.PlmnId
	AmfId         string
	Subscriptions map[models.AmfEventTypeAnyOf][]string
	SubMutex      sync.RWMutex
}

func NewAmf(plmnId models.PlmnId) *Amf {
	return &Amf{
		PlmnId:        plmnId,
		AmfId:         fmt.Sprintf("AMF-%s%s", plmnId.Mcc, plmnId.Mnc),
		Subscriptions: make(map[models.AmfEventTypeAnyOf][]string),
		SubMutex:      sync.RWMutex{},
	}
}

func (amf *Amf) InitAmf() {
	log.Printf("[%s] started", amf.AmfId)
	err := gitc.StartTask("AMF", func(msg gitc.Message) {
		switch msg.Type {
		case models.UeToAmfType:
			//			log.Printf("[%s] Received message UeToAmfMsg from %s", amf.AmfId, msg.From)
			amf.handleUeToAmfEvent(msg.Payload.(*models.UeToAmfMsg))
		}
	}, 1024)
	if err != nil {
		log.Fatalf("[%s] could not start AMF task: %s", amf.AmfId, err.Error())
	}
}

func (amf *Amf) handleUeToAmfEvent(msg *models.UeToAmfMsg) {
	//log.Printf("[%s] UeToAmfMsg: %+v", amf.AmfId, msg)

	// Process the message and notify subscribers
	amf.SubMutex.RLock()
	defer amf.SubMutex.RUnlock()

	//prepare the basic report
	amfReport := models.AmfEventReport{
		Type:      msg.EventType,
		TimeStamp: msg.TimeStamp,
		Supi:      models.PtrString(msg.Supi),
		Gpsi:      models.PtrString(msg.Gpsi), // MSISDN in E.164 format
		State: models.AmfEventState{
			Active: true},
		AccessTypeList: []models.AccessType{msg.AccessType},
		Location: &models.UserLocation{
			NrLocation: &models.NrLocation{
				UeLocationTimestamp:      &msg.TimeStamp,
				AgeOfLocationInformation: models.PtrInt32(0),
				Tai: models.Tai{
					PlmnId: msg.PlmnId,
					Tac:    "001010",
				},
				Ncgi: models.Ncgi{
					PlmnId:   msg.PlmnId,
					NrCellId: msg.CurrentCellId,
				},
			},
		},
	}

	switch msg.EventType {
	case models.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT:
		amfReport.CmInfoList = []models.CmInfo{{
			CmState:    msg.CmState,
			AccessType: msg.AccessType}}

	case models.AMFEVENTTYPEANYOF_REGISTRATION_STATE_REPORT:
		amfReport.RmInfoList = []models.RmInfo{{
			RmState:    msg.RmState,
			AccessType: msg.AccessType}}

	case models.AMFEVENTTYPEANYOF_LOCATION_REPORT:
		// no need to add anything specific for location report
	case models.AMFEVENTTYPEANYOF_LOSS_OF_CONNECTIVITY:
		amfReport.LossOfConnectReason = models.LOSSOFCONNECTIVITYREASONANYOF_DEREGISTERED

	case models.AMFEVENTTYPEANYOF_UES_IN_AREA_REPORT:
	}

	// buildup the notification structure
	amfNotification := &models.AmfEventNotification{
		ReportList: []models.AmfEventReport{amfReport},
	}

	//	log.Printf("[%s] generating notification : %+v", amf.AmfId, amfNotification)

	callbackBody, err := json.Marshal(amfNotification)
	if err != nil {
		log.Printf("[%s] error while marshalling notification: %s", amf.AmfId, err.Error())
		return
	}

	for _, callbackUrl := range amf.Subscriptions[msg.EventType] {
		go func(url string, data []byte) {
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(data)) // Replace nil with actual payload if needed
			if err != nil {
				log.Printf("Error notifying subscriber %s: %v", url, err)
				return
			}
			defer func() {
				_ = resp.Body.Close()
			}()
			//log.Printf("Notified subscriber %s with response status: %s", url, resp.Status)
		}(callbackUrl, callbackBody)
	}
}

// NORTHBOUND Definitions

func (amf *Amf) HandleNewSubscription(w http.ResponseWriter, r *http.Request) {

	subData := &models.AmfCreateEventSubscription{}

	if err := json.NewDecoder(r.Body).Decode(subData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sub, ok := subData.GetSubscriptionOk()
	if !ok {
		http.Error(w, "could not find subscription information", http.StatusBadRequest)
		return
	}

	callbackUrl, ok := sub.GetEventNotifyUriOk()
	if !ok {
		http.Error(w, "could not find callbackUri information", http.StatusBadRequest)
		return
	}

	amf.SubMutex.Lock()
	defer amf.SubMutex.Unlock()

	for _, event := range sub.GetEventList() {
		if listSub, exist := amf.Subscriptions[event.Type]; exist {
			amf.Subscriptions[event.Type] = append(listSub, *callbackUrl)
		} else {
			amf.Subscriptions[event.Type] = []string{*callbackUrl}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(subData); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

	log.Printf("[%s] created new subscription for: %s", amf.AmfId, *callbackUrl)
}

func (amf *Amf) RegisterNorthboundAPIs(r *mux.Router) {
	r.HandleFunc("/namf-evts/v1/subscriptions", amf.HandleNewSubscription)
	log.Printf("[%s] namf-evts has been registered", amf.AmfId)
}
