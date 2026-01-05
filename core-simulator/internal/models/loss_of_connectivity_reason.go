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

// LossOfConnectivityReason Describes the reason for loss of connectivity
type LossOfConnectivityReason struct {
	LossOfConnectivityReasonAnyOf *LossOfConnectivityReasonAnyOf
	string                        *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LossOfConnectivityReason) UnmarshalJSON(data []byte) error {
	// try to unmarshal JSON data into LossOfConnectivityReasonAnyOf
	err := json.Unmarshal(data, &dst.LossOfConnectivityReasonAnyOf)
	if err == nil {
		jsonLossOfConnectivityReasonAnyOf, _ := json.Marshal(dst.LossOfConnectivityReasonAnyOf)
		if string(jsonLossOfConnectivityReasonAnyOf) == "{}" { // empty struct
			dst.LossOfConnectivityReasonAnyOf = nil
		} else {
			return nil // data stored in dst.LossOfConnectivityReasonAnyOf, return on the first match
		}
	} else {
		dst.LossOfConnectivityReasonAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(LossOfConnectivityReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LossOfConnectivityReason) MarshalJSON() ([]byte, error) {
	if src.LossOfConnectivityReasonAnyOf != nil {
		return json.Marshal(&src.LossOfConnectivityReasonAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLossOfConnectivityReason struct {
	value *LossOfConnectivityReason
	isSet bool
}

func (v NullableLossOfConnectivityReason) Get() *LossOfConnectivityReason {
	return v.value
}

func (v *NullableLossOfConnectivityReason) Set(val *LossOfConnectivityReason) {
	v.value = val
	v.isSet = true
}

func (v NullableLossOfConnectivityReason) IsSet() bool {
	return v.isSet
}

func (v *NullableLossOfConnectivityReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossOfConnectivityReason(val *LossOfConnectivityReason) *NullableLossOfConnectivityReason {
	return &NullableLossOfConnectivityReason{value: val, isSet: true}
}

func (v NullableLossOfConnectivityReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossOfConnectivityReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
