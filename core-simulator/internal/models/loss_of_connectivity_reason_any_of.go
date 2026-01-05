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

// LossOfConnectivityReasonAnyOf the model 'LossOfConnectivityReasonAnyOf'
type LossOfConnectivityReasonAnyOf string

// List of LossOfConnectivityReason_anyOf
const (
	LOSSOFCONNECTIVITYREASONANYOF_DEREGISTERED               LossOfConnectivityReasonAnyOf = "DEREGISTERED"
	LOSSOFCONNECTIVITYREASONANYOF_MAX_DETECTION_TIME_EXPIRED LossOfConnectivityReasonAnyOf = "MAX_DETECTION_TIME_EXPIRED"
	LOSSOFCONNECTIVITYREASONANYOF_PURGED                     LossOfConnectivityReasonAnyOf = "PURGED"
)

// All allowed values of LossOfConnectivityReasonAnyOf enum
var AllowedLossOfConnectivityReasonAnyOfEnumValues = []LossOfConnectivityReasonAnyOf{
	"DEREGISTERED",
	"MAX_DETECTION_TIME_EXPIRED",
	"PURGED",
}

func (v *LossOfConnectivityReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LossOfConnectivityReasonAnyOf(value)
	for _, existing := range AllowedLossOfConnectivityReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LossOfConnectivityReasonAnyOf", value)
}

// NewLossOfConnectivityReasonAnyOfFromValue returns a pointer to a valid LossOfConnectivityReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLossOfConnectivityReasonAnyOfFromValue(v string) (*LossOfConnectivityReasonAnyOf, error) {
	ev := LossOfConnectivityReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LossOfConnectivityReasonAnyOf: valid values are %v", v, AllowedLossOfConnectivityReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LossOfConnectivityReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedLossOfConnectivityReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LossOfConnectivityReason_anyOf value
func (v LossOfConnectivityReasonAnyOf) Ptr() *LossOfConnectivityReasonAnyOf {
	return &v
}

type NullableLossOfConnectivityReasonAnyOf struct {
	value *LossOfConnectivityReasonAnyOf
	isSet bool
}

func (v NullableLossOfConnectivityReasonAnyOf) Get() *LossOfConnectivityReasonAnyOf {
	return v.value
}

func (v *NullableLossOfConnectivityReasonAnyOf) Set(val *LossOfConnectivityReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLossOfConnectivityReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLossOfConnectivityReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossOfConnectivityReasonAnyOf(val *LossOfConnectivityReasonAnyOf) *NullableLossOfConnectivityReasonAnyOf {
	return &NullableLossOfConnectivityReasonAnyOf{value: val, isSet: true}
}

func (v NullableLossOfConnectivityReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossOfConnectivityReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
