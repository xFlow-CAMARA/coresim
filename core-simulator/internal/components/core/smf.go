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
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/utils"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
)

type Smf struct {
	PlmnId        models.PlmnId
	SmfId         string
	Subscriptions map[models.SmfEventAnyOf][]string
	SubMutex      sync.RWMutex
	ipamInstance  *utils.IPAllocator
}

func NewSmf(plmnId models.PlmnId, ipamInstance *utils.IPAllocator) *Smf {
	return &Smf{
		PlmnId:        plmnId,
		SmfId:         fmt.Sprintf("SMF-%s%s", plmnId.Mcc, plmnId.Mnc),
		Subscriptions: make(map[models.SmfEventAnyOf][]string),
		SubMutex:      sync.RWMutex{},
		ipamInstance:  ipamInstance,
	}
}

func (smf *Smf) InitSmf() {
	log.Printf("[%s] started", smf.SmfId)

	err := gitc.StartTask("SMF", func(msg gitc.Message) {
		switch msg.Type {
		case models.UeToSmfType:
			//log.Printf("[%s] Received message UeToSmfMsg from %s", smf.SmfId, msg.From)
			smf.handleUeToSmfEvent(msg.Payload.(*models.UeToSmfMsg))
		}
	}, 1024)
	if err != nil {
		log.Fatalf("[%s] could not start SMF task: %s", smf.SmfId, err.Error())
	}
}

func (smf *Smf) handleUeToSmfEvent(msg *models.UeToSmfMsg) {
	//log.Printf("[%s] UeToSmfMsg: %+v", smf.SmfId, msg)

	// Process the message and notify subscribers
	smf.SubMutex.RLock()
	defer smf.SubMutex.RUnlock()

	//prepare the basic report
	smfEvent := models.EventNotification{
		Event:     msg.EventType,
		TimeStamp: msg.TimeStamp,
		Supi:      models.PtrString(msg.Supi),
		Gpsi:      models.PtrString(msg.Gpsi), // MSISDN in E.164 format
		Dnn:       &msg.Dnn,
		Snssai:    &msg.Snssai,
		AccType:   &msg.AccessType,
		PduSeId:   &msg.PduSessId,
		PlmnId:    &msg.PlmnId,
	}

	switch msg.EventType {
	case models.SMFEVENTANYOF_PDU_SES_EST:
		smfEvent.PduSessType = &models.PduSessionType{
			PduSessionTypeAnyOf: &msg.PduSessType,
		}
		smfEvent.Ipv4Addr = &msg.UeAddress

	case models.SMFEVENTANYOF_PDU_SES_REL:
		smfEvent.PduSessType = &models.PduSessionType{
			PduSessionTypeAnyOf: &msg.PduSessType,
		}
		smfEvent.Ipv4Addr = &msg.UeAddress

	case models.SMFEVENTANYOF_DDDS:
		smfEvent.DddStatus = &models.DlDataDeliveryStatus{
			DlDataDeliveryStatusAnyOf: &msg.DddsState,
		}
	case models.SMFEVENTANYOF_QOS_MON:
		smfEvent.CustomizedData = &models.CustomizedData{
			UsageReport: models.CustomUsageReport{
				Volume: models.Volume{
					Downlink: msg.UpReport.TotalDlBytes,
					Uplink:   msg.UpReport.TotalUlBytes,
					Total:    msg.UpReport.TotalBytes,
				},
				NoP: models.Volume{
					Downlink: msg.UpReport.NumDlPackets,
					Uplink:   msg.UpReport.NumUlPackets,
					Total:    msg.UpReport.NumOfPackets,
				},
				Trigger: "PERIODIC",
				SeId:    msg.PduSessId,
			},
		}
	case models.SMFEVENTANYOF_COMM_FAIL:
	}

	// buildup the notification structure
	amfNotification := &models.NsmfEventExposureNotification{
		NotifId:     "test",
		EventNotifs: []models.EventNotification{smfEvent},
	}
	//log.Printf("[%s] generating notification : %+v", smf.SmfId, amfNotification)

	callbackBody, err := json.Marshal(amfNotification)
	if err != nil {
		log.Printf("[%s] error while marshalling notification: %s", smf.SmfId, err.Error())
		return
	}

	for _, callbackUrl := range smf.Subscriptions[msg.EventType] {
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

func (smf *Smf) HandleNewSubscription(w http.ResponseWriter, r *http.Request) {

	subData := &models.NsmfEventExposure{}

	if err := json.NewDecoder(r.Body).Decode(subData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	subList, ok := subData.GetEventSubsOk()
	if !ok {
		http.Error(w, "could not find event list", http.StatusBadRequest)
		return
	}

	callbackUrl, ok := subData.GetNotifUriOk()
	if !ok {
		http.Error(w, "could not find callbackUri information", http.StatusBadRequest)
		return
	}

	smf.SubMutex.Lock()
	defer smf.SubMutex.Unlock()

	for _, event := range subList {
		if listSub, exist := smf.Subscriptions[event.Event]; exist {
			smf.Subscriptions[event.Event] = append(listSub, *callbackUrl)
		} else {
			smf.Subscriptions[event.Event] = []string{*callbackUrl}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(subData); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

	log.Printf("[%s] created new subscription for: %s", smf.SmfId, *callbackUrl)
}

func (smf *Smf) RegisterNorthboundAPIs(r *mux.Router) {
	r.HandleFunc("/nsmf-event-exposure/v1/subscriptions", smf.HandleNewSubscription)
	log.Printf("[%s] nsmf-event-exposure has been registered", smf.SmfId)
}
