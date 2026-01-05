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

// FinalUnitAction struct for FinalUnitAction
type FinalUnitAction struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FinalUnitAction) UnmarshalJSON(data []byte) error {
	// try to unmarshal JSON data into String
	err := json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FinalUnitAction)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FinalUnitAction) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFinalUnitAction struct {
	value *FinalUnitAction
	isSet bool
}

func (v NullableFinalUnitAction) Get() *FinalUnitAction {
	return v.value
}

func (v *NullableFinalUnitAction) Set(val *FinalUnitAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFinalUnitAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFinalUnitAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinalUnitAction(val *FinalUnitAction) *NullableFinalUnitAction {
	return &NullableFinalUnitAction{value: val, isSet: true}
}

func (v NullableFinalUnitAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinalUnitAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
