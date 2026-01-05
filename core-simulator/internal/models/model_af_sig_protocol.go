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

// AfSigProtocol Indicates the protocol used for signalling between the UE and the AF.   Possible values are - NO_INFORMATION: Indicate that no information about the AF signalling protocol is being provided. - SIP: Indicate that the signalling protocol is Session Initiation Protocol.
type AfSigProtocol struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfSigProtocol) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfSigProtocol)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfSigProtocol) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfSigProtocol struct {
	value *AfSigProtocol
	isSet bool
}

func (v NullableAfSigProtocol) Get() *AfSigProtocol {
	return v.value
}

func (v *NullableAfSigProtocol) Set(val *AfSigProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableAfSigProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableAfSigProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfSigProtocol(val *AfSigProtocol) *NullableAfSigProtocol {
	return &NullableAfSigProtocol{value: val, isSet: true}
}

func (v NullableAfSigProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfSigProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
