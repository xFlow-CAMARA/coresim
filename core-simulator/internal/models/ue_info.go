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

package models

type UeInfoPatch struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Imsi string                 `json:"imsi"`
}

type UeInfo struct {
	Imsi               *string
	PduSessEst         map[string]*pduSesEst `json:"PDU_SES_EST"`
	PduSessRel         map[string]*pduSesRel `json:"PDU_SES_REL"`
	Ddds               map[string]*ddds      `json:"DDDS"`
	RegistrationInfo   *rmInfo               `json:"REGISTRATION_STATE_REPORT"`
	ConnectivityInfo   *cmInfo               `json:"CONNECTIVITY_STATE_REPORT"`
	LossOfConnectivity *lossOfConnectReason  `json:"LOSS_OF_CONNECTIVITY"`
	Location           *location             `json:"LOCATION_REPORT"`
}

type Snssai struct {
	Sst int32   `yaml:"sst" json:"sst"`
	Sd  *string `yaml:"sd" json:"sd"`
}

type pduSesEst struct {
	AdIpv4Addr  *string         `json:"AdIpv4Addr"`
	Dnn         *string         `json:"Dnn"`
	PduSeId     *int32          `json:"PduSeId"`
	PduSessType *PduSessionType `json:"PduSessType"`
	Snssai      *Snssai         `json:"Snssai"`
	TimeStamp   int64           `json:"TimeStamp"`
}
type pduSesRel struct {
	Ipv4Addr    *string         `json:"Ipv4Addr"`
	Dnn         *string         `json:"Dnn"`
	PduSeId     *int32          `json:"PduSeId"`
	PduSessType *PduSessionType `json:"PduSessType"`
	Snssai      *Snssai         `json:"Snssai"`
	TimeStamp   int64           `json:"TimeStamp"`
}

type ddds struct {
	DddStatus  *DlDataDeliveryStatus `json:"DddStatus"`
	PduSeId    *int32                `json:"PduSeId"`
	Dnn        *string               `json:"Dnn"`
	Snssai     *Snssai               `json:"Snssai"`
	UeIpv4Addr *string               `json:"UeIpv4Addr"`
	TimeStamp  int64                 `json:"TimeStamp"`
}

type rmInfo struct {
	RmInfo    RmInfo `json:"RmInfo"`
	TimeStamp int64  `json:"TimeStamp"`
}

type cmInfo struct {
	CmInfo    CmInfo `json:"CmInfo"`
	TimeStamp int64  `json:"TimeStamp"`
}

type location struct {
	UserLocation UserLocation `json:"UserLocation"`
	TimeStamp    int64        `json:"TimeStamp"`
}

type lossOfConnectReason struct {
	LossOfConnectReason LossOfConnectivityReasonAnyOf `json:"LossOfConnectReason"`
	TimeStamp           int64                         `json:"TimeStamp"`
}
