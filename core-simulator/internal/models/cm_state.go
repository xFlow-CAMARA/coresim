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
	"encoding/json"
	"fmt"
)

type CmState string

const (
	CmStateIdle      CmState = "IDLE"
	CmStateConnected CmState = "CONNECTED"
)

func (c *CmState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case string(CmStateIdle), string(CmStateConnected):
		*c = CmState(s)
		return nil
	default:
		return fmt.Errorf("invalid CmState: %q", s)
	}
}

func (c CmState) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}
