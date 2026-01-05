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

// AccessType Indicates wether the access is  via 3GPP or via non-3GPP.
type AccessType string

// List of AccessType
const (
	ACCESSTYPE__3_GPP_ACCESS    AccessType = "3GPP_ACCESS"
	ACCESSTYPE_NON_3_GPP_ACCESS AccessType = "NON_3GPP_ACCESS"
)

// All allowed values of AccessType enum
var AllowedAccessTypeEnumValues = []AccessType{
	"3GPP_ACCESS",
	"NON_3GPP_ACCESS",
}

func (v *AccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessType(value)
	for _, existing := range AllowedAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessType", value)
}

// NewAccessTypeFromValue returns a pointer to a valid AccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessTypeFromValue(v string) (*AccessType, error) {
	ev := AccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessType: valid values are %v", v, AllowedAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessType) IsValid() bool {
	for _, existing := range AllowedAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessType value
func (v AccessType) Ptr() *AccessType {
	return &v
}

type NullableAccessType struct {
	value *AccessType
	isSet bool
}

func (v NullableAccessType) Get() *AccessType {
	return v.value
}

func (v *NullableAccessType) Set(val *AccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessType(val *AccessType) *NullableAccessType {
	return &NullableAccessType{value: val, isSet: true}
}

func (v NullableAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
