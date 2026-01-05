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

import (
	"time"

	"github.com/giuliocarot0/gitc"
)

const (
	UeToAmfType gitc.MessageType = iota
	AmfToUeType
	UeToSmfType
	SmfToUeType
	PcfToUeType
	UeToPcfType
)

type UeToAmfMsg struct {
	EventType     AmfEventTypeAnyOf
	TimeStamp     time.Time
	RmState       RmState
	CmState       CmState
	Supi          string
	Gpsi          string // MSISDN in E.164 format (e.g., "+33612345678")
	PlmnId        PlmnId
	CurrentCellId string
	AccessType    AccessType
}

type UeToSmfMsg struct {
	EventType   SmfEventAnyOf
	TimeStamp   time.Time
	Supi        string
	Gpsi        string // MSISDN in E.164 format
	PlmnId      PlmnId
	AccessType  AccessType
	Dnn         string
	Snssai      Snssai
	UeAddress   string
	PduSessType PduSessionTypeAnyOf
	PduSessId   int32
	DddsState   DlDataDeliveryStatusAnyOf
	UpReport    *UpStatsReport
}

type AmfToUeMsg struct {
}

type SmfToUeMsg struct {
}
type PcfToUeMsg struct {
}
type UeToPcfMsg struct {
}
