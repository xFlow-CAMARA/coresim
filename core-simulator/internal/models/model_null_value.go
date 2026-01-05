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
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.6
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NullValue JSON's null value.
type NullValue string

// List of NullValue
const ()

// All allowed values of NullValue enum
var AllowedNullValueEnumValues = []NullValue{}

func (v *NullValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NullValue(value)
	for _, existing := range AllowedNullValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NullValue", value)
}

// NewNullValueFromValue returns a pointer to a valid NullValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNullValueFromValue(v string) (*NullValue, error) {
	ev := NullValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NullValue: valid values are %v", v, AllowedNullValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NullValue) IsValid() bool {
	for _, existing := range AllowedNullValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NullValue value
func (v NullValue) Ptr() *NullValue {
	return &v
}

type NullableNullValue struct {
	value *NullValue
	isSet bool
}

func (v NullableNullValue) Get() *NullValue {
	return v.value
}

func (v *NullableNullValue) Set(val *NullValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNullValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNullValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullValue(val *NullValue) *NullableNullValue {
	return &NullableNullValue{value: val, isSet: true}
}

func (v NullableNullValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
