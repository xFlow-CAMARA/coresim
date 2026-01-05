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

type UeState int

const (
	Deregistered UeState = iota
	Registered
	Attached  //connected without PDU yet
	Idle      // PDU established, no traffic
	Connected // PDU session established, has recent UL/DL activity
	Handover  // Can occour at any time, it is transient leads to connected with pdu, or deregistered
)

type UeProcedure string

const (
	NoProcedure              UeProcedure = "NONE"
	Registration             UeProcedure = "REGISTRATION"
	Attach                   UeProcedure = "ATTACH"
	PduSessionEstablishement UeProcedure = "PDU_SES_EST"
	PduSessionFailure        UeProcedure = "PDU_SES_FAIL"
	PduSessionRelease        UeProcedure = "PDU_SES_REL"
	LossOfConnection         UeProcedure = "LOSS_OF_CONNECTION"
	Sleep                    UeProcedure = "IDLE_MODE"
	Paging                   UeProcedure = "PAGING"
	HandoverSuccessful       UeProcedure = "HO_SUCCESSFUL"
	HandoverFailure          UeProcedure = "HO_FAILED"
	HandoverInitiated        UeProcedure = "HO_INITIATED"
)

type Transition struct {
	To          UeState
	Probability float64
	Procedure   UeProcedure
}
