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

// DlDataDeliveryStatusAnyOf the model 'DlDataDeliveryStatusAnyOf'
type DlDataDeliveryStatusAnyOf string

// List of DlDataDeliveryStatus_anyOf
const (
	DLDATADELIVERYSTATUSANYOF_BUFFERED    DlDataDeliveryStatusAnyOf = "BUFFERED"
	DLDATADELIVERYSTATUSANYOF_TRANSMITTED DlDataDeliveryStatusAnyOf = "TRANSMITTED"
	DLDATADELIVERYSTATUSANYOF_DISCARDED   DlDataDeliveryStatusAnyOf = "DISCARDED"
)

// All allowed values of DlDataDeliveryStatusAnyOf enum
var AllowedDlDataDeliveryStatusAnyOfEnumValues = []DlDataDeliveryStatusAnyOf{
	"BUFFERED",
	"TRANSMITTED",
	"DISCARDED",
}

func (v *DlDataDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DlDataDeliveryStatusAnyOf(value)
	for _, existing := range AllowedDlDataDeliveryStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DlDataDeliveryStatusAnyOf", value)
}

// NewDlDataDeliveryStatusAnyOfFromValue returns a pointer to a valid DlDataDeliveryStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDlDataDeliveryStatusAnyOfFromValue(v string) (*DlDataDeliveryStatusAnyOf, error) {
	ev := DlDataDeliveryStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DlDataDeliveryStatusAnyOf: valid values are %v", v, AllowedDlDataDeliveryStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DlDataDeliveryStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedDlDataDeliveryStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DlDataDeliveryStatus_anyOf value
func (v DlDataDeliveryStatusAnyOf) Ptr() *DlDataDeliveryStatusAnyOf {
	return &v
}

type NullableDlDataDeliveryStatusAnyOf struct {
	value *DlDataDeliveryStatusAnyOf
	isSet bool
}

func (v NullableDlDataDeliveryStatusAnyOf) Get() *DlDataDeliveryStatusAnyOf {
	return v.value
}

func (v *NullableDlDataDeliveryStatusAnyOf) Set(val *DlDataDeliveryStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDlDataDeliveryStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDlDataDeliveryStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlDataDeliveryStatusAnyOf(val *DlDataDeliveryStatusAnyOf) *NullableDlDataDeliveryStatusAnyOf {
	return &NullableDlDataDeliveryStatusAnyOf{value: val, isSet: true}
}

func (v NullableDlDataDeliveryStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlDataDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
