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

// LineTypeAnyOf the model 'LineTypeAnyOf'
type LineTypeAnyOf string

// List of LineType_anyOf
const (
	LINETYPEANYOF_DSL LineTypeAnyOf = "DSL"
	LINETYPEANYOF_PON LineTypeAnyOf = "PON"
)

// All allowed values of LineTypeAnyOf enum
var AllowedLineTypeAnyOfEnumValues = []LineTypeAnyOf{
	"DSL",
	"PON",
}

func (v *LineTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LineTypeAnyOf(value)
	for _, existing := range AllowedLineTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LineTypeAnyOf", value)
}

// NewLineTypeAnyOfFromValue returns a pointer to a valid LineTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLineTypeAnyOfFromValue(v string) (*LineTypeAnyOf, error) {
	ev := LineTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LineTypeAnyOf: valid values are %v", v, AllowedLineTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LineTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedLineTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LineType_anyOf value
func (v LineTypeAnyOf) Ptr() *LineTypeAnyOf {
	return &v
}

type NullableLineTypeAnyOf struct {
	value *LineTypeAnyOf
	isSet bool
}

func (v NullableLineTypeAnyOf) Get() *LineTypeAnyOf {
	return v.value
}

func (v *NullableLineTypeAnyOf) Set(val *LineTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLineTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLineTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineTypeAnyOf(val *LineTypeAnyOf) *NullableLineTypeAnyOf {
	return &NullableLineTypeAnyOf{value: val, isSet: true}
}

func (v NullableLineTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
