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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/giuliocarot0/gitc"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/utils"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
)

type Pcf struct {
	PlmnId        models.PlmnId
	PcfId         string
	Subscriptions map[string]*models.AppSessionContext
	SubMutex      sync.RWMutex
	ipamInstance  *utils.IPAllocator
}

func NewPcf(plmnId models.PlmnId, ipamInstance *utils.IPAllocator) *Pcf {
	return &Pcf{
		PlmnId:        plmnId,
		PcfId:         fmt.Sprintf("PCF-%s%s", plmnId.Mcc, plmnId.Mnc),
		Subscriptions: make(map[string]*models.AppSessionContext),
		SubMutex:      sync.RWMutex{},
		ipamInstance:  ipamInstance,
	}
}

func (pcf *Pcf) InitPcf() {
	log.Printf("[%s] started", pcf.PcfId)
	err := gitc.StartTask("PCF", func(msg gitc.Message) {
		switch msg.Type {
		default:
		}
	}, 1024)
	if err != nil {
		log.Fatalf("[%s] could not start PCF task: %s", pcf.PcfId, err.Error())
	}
}

// NORTHBOUND Definitions

func (pcf *Pcf) HandleNewSubscription(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		subData := &models.AppSessionContext{}

		if err := json.NewDecoder(r.Body).Decode(subData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		rData, ok := subData.GetAscReqDataOk()
		if !ok || rData == nil {
			http.Error(w, "Missing ascReqData", http.StatusBadRequest)
			return
		}

		ueAddr, ok := rData.GetUeIpv4Ok()
		if !ok || ueAddr == nil {
			http.Error(w, "Missing Ue Ipv4 Address", http.StatusBadRequest)
			return
		}

		// do lookup of the UE in the IP Management system
		supi, pduSessId, ok := pcf.ipamInstance.GetUserStringOk(*ueAddr)

		if !ok {
			http.Error(w, "requested UE is not connected to the network", http.StatusNotFound)
			return
		}
		// inform the network about the new policy decision
		log.Printf("received new policy decision for UE %s, pduSessId %d", supi, pduSessId)

		// now verify if it is a qos session or a routing decision
		if _, ok := rData.GetMedComponentsOk(); ok {
			// this is a qos session
			// this normally results in creating a new qos flow for the target pdu session
		} else if _, ok := rData.GetAfRoutReqOk(); ok {
			// this is a routing decision
			// this normally results in reconfiguring the UP path for the target PDU session
		} else {
			// this is not supported yet
			http.Error(w, "unsupported policy request", http.StatusBadRequest)
			return
		}

		subId := uuid.New().String()
		location := "/npcf-policyauthorization/v1/app-sessions/" + subId
		w.Header().Add("Location", location)
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(subData); err != nil {
			http.Error(w, "could not serialize response body", http.StatusInternalServerError)
			return
		}

		pcf.SubMutex.Lock()
		defer pcf.SubMutex.Unlock()
		pcf.Subscriptions[subId] = subData

		log.Printf("[%s] created new subscription", pcf.PcfId)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (pcf *Pcf) HandleUpdateSubscription(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "policy update is not supported yet", http.StatusNotImplemented)
}

func (pcf *Pcf) HandleDeleteSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed on this endpoint", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	appSessId := vars["appSessId"]

	if len(appSessId) == 0 {
		http.Error(w, "the request url is malformed", http.StatusBadRequest)
	}

	pcf.SubMutex.Lock()
	defer pcf.SubMutex.Unlock()

	if sub := pcf.Subscriptions[appSessId]; sub != nil {
		// do operations by notifying the network about the policy change
		delete(pcf.Subscriptions, appSessId)
		log.Printf("[%s] deleted subscription ", pcf.PcfId)

	} else {
		http.Error(w, "app-session context is not found", http.StatusNotFound)
		return
	}

}

func (pcf *Pcf) RegisterNorthboundAPIs(r *mux.Router) {
	r.HandleFunc("/npcf-policyauthorization/v1/app-sessions", pcf.HandleNewSubscription)
	r.HandleFunc("/npcf-policyauthorization/v1/app-sessions/{appSessId}/delete", pcf.HandleDeleteSubscription)
	r.HandleFunc("/npcf-policyauthorization/v1/app-sessions/{appSessId}", pcf.HandleUpdateSubscription)
	log.Printf("[%s] npcf-policyauthorization has been registered", pcf.PcfId)
}
