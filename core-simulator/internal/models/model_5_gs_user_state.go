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

// Model5GsUserState Describes the 5GS User State of a UE
type Model5GsUserState struct {
	Model5GsUserStateAnyOf *Model5GsUserStateAnyOf
	string                 *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Model5GsUserState) UnmarshalJSON(data []byte) error {
	// try to unmarshal JSON data into 5GsUserStateAnyOf
	err := json.Unmarshal(data, &dst.Model5GsUserStateAnyOf)
	if err == nil {
		json5GsUserStateAnyOf, _ := json.Marshal(dst.Model5GsUserStateAnyOf)
		if string(json5GsUserStateAnyOf) == "{}" { // empty struct
			dst.Model5GsUserStateAnyOf = nil
		} else {
			return nil // data stored in dst.Model5GsUserStateAnyOf, return on the first match
		}
	} else {
		dst.Model5GsUserStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Model5GsUserState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Model5GsUserState) MarshalJSON() ([]byte, error) {
	if src.Model5GsUserStateAnyOf != nil {
		return json.Marshal(&src.Model5GsUserStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableModel5GsUserState struct {
	value *Model5GsUserState
	isSet bool
}

func (v NullableModel5GsUserState) Get() *Model5GsUserState {
	return v.value
}

func (v *NullableModel5GsUserState) Set(val *Model5GsUserState) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GsUserState) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GsUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GsUserState(val *Model5GsUserState) *NullableModel5GsUserState {
	return &NullableModel5GsUserState{value: val, isSet: true}
}

func (v NullableModel5GsUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GsUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
