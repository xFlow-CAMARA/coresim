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

// RtpHeaderExtType The enumeration indicates the type of Rtp Header Extension type
type RtpHeaderExtType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RtpHeaderExtType) UnmarshalJSON(data []byte) error {

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

	return fmt.Errorf("data failed to match schemas in anyOf(RtpHeaderExtType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RtpHeaderExtType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRtpHeaderExtType struct {
	value *RtpHeaderExtType
	isSet bool
}

func (v NullableRtpHeaderExtType) Get() *RtpHeaderExtType {
	return v.value
}

func (v *NullableRtpHeaderExtType) Set(val *RtpHeaderExtType) {
	v.value = val
	v.isSet = true
}

func (v NullableRtpHeaderExtType) IsSet() bool {
	return v.isSet
}

func (v *NullableRtpHeaderExtType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtpHeaderExtType(val *RtpHeaderExtType) *NullableRtpHeaderExtType {
	return &NullableRtpHeaderExtType{value: val, isSet: true}
}

func (v NullableRtpHeaderExtType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtpHeaderExtType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
