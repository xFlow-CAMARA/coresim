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

/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AmfEventTriggerAnyOf the model 'AmfEventTriggerAnyOf'
type AmfEventTriggerAnyOf string

// List of AmfEventTrigger_anyOf
const (
	AMFEVENTTRIGGERANYOF_ONE_TIME   AmfEventTriggerAnyOf = "ONE_TIME"
	AMFEVENTTRIGGERANYOF_CONTINUOUS AmfEventTriggerAnyOf = "CONTINUOUS"
	AMFEVENTTRIGGERANYOF_PERIODIC   AmfEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of AmfEventTriggerAnyOf enum
var AllowedAmfEventTriggerAnyOfEnumValues = []AmfEventTriggerAnyOf{
	"ONE_TIME",
	"CONTINUOUS",
	"PERIODIC",
}

func (v *AmfEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmfEventTriggerAnyOf(value)
	for _, existing := range AllowedAmfEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmfEventTriggerAnyOf", value)
}

// NewAmfEventTriggerAnyOfFromValue returns a pointer to a valid AmfEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmfEventTriggerAnyOfFromValue(v string) (*AmfEventTriggerAnyOf, error) {
	ev := AmfEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmfEventTriggerAnyOf: valid values are %v", v, AllowedAmfEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmfEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedAmfEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AmfEventTrigger_anyOf value
func (v AmfEventTriggerAnyOf) Ptr() *AmfEventTriggerAnyOf {
	return &v
}

type NullableAmfEventTriggerAnyOf struct {
	value *AmfEventTriggerAnyOf
	isSet bool
}

func (v NullableAmfEventTriggerAnyOf) Get() *AmfEventTriggerAnyOf {
	return v.value
}

func (v *NullableAmfEventTriggerAnyOf) Set(val *AmfEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventTriggerAnyOf(val *AmfEventTriggerAnyOf) *NullableAmfEventTriggerAnyOf {
	return &NullableAmfEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableAmfEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
