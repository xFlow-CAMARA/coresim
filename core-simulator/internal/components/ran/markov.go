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

package ran

import (
	"math/rand/v2"

	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
)

var transitions = map[models.UeState][]models.Transition{

	models.Deregistered: {
		{To: models.Registered, Probability: 0.90, Procedure: models.Registration},  // ue turn on and registered
		{To: models.Deregistered, Probability: 0.10, Procedure: models.NoProcedure}, // ue off or failed to register
	},
	models.Registered: {
		{To: models.Attached, Probability: 0.99, Procedure: models.Attach},                // initial Ue Setup completed, signalling is active
		{To: models.Registered, Probability: 0.005},                                       // attach pendind
		{To: models.Deregistered, Probability: 0.005, Procedure: models.LossOfConnection}, // attach failed - too many attemps/loss of connectivity
	},
	models.Attached: {
		{To: models.Connected, Probability: 0.90, Procedure: models.PduSessionEstablishement}, // pdu sess established
		{To: models.Attached, Probability: 0.05, Procedure: models.NoProcedure},               // pdu sess establishment pending
		{To: models.Deregistered, Probability: 0.05, Procedure: models.LossOfConnection},      // loss of connectivity, too many failed attempts
	},
	models.Idle: {
		{To: models.Idle, Probability: 0.94, Procedure: models.NoProcedure},               // this is only for paging (network initiated), Service request is handled by UE traffic generator routine
		{To: models.Connected, Probability: 0.05, Procedure: models.Paging},               // this is only for paging (network initiated), Service request is handled by UE traffic generator routine
		{To: models.Handover, Probability: 0.007, Procedure: models.HandoverInitiated},    // handover
		{To: models.Deregistered, Probability: 0.003, Procedure: models.LossOfConnection}, //loss of connectivity/ UL failure
	},
	models.Connected: {
		{To: models.Connected, Probability: 0.997, Procedure: models.NoProcedure},
		{To: models.Handover, Probability: 0.002, Procedure: models.HandoverInitiated},
		{To: models.Deregistered, Probability: 0.001, Procedure: models.LossOfConnection}, //loss of connectivity/ UL failure
	},

	models.Handover: {
		{To: models.Connected, Probability: 0.99, Procedure: models.HandoverSuccessful},
		{To: models.Deregistered, Probability: 0.01, Procedure: models.HandoverFailure}, //handover failed
	},
}

func NextState(current models.UeState) (models.UeState, models.UeProcedure) {
	rnd := rand.Float64()
	cumulative := 0.0
	for _, t := range transitions[current] {
		cumulative += t.Probability
		if rnd < cumulative {
			return t.To, t.Procedure
		}
	}
	//TODO use handlers instead of procedure code
	return current, models.NoProcedure // fallback
}
