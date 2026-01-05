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

// SipForkingIndication Indicates whether several SIP dialogues are related to an \"Individual Application Session Context\" resource.
type SipForkingIndication struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SipForkingIndication) UnmarshalJSON(data []byte) error {

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

	return fmt.Errorf("data failed to match schemas in anyOf(SipForkingIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SipForkingIndication) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSipForkingIndication struct {
	value *SipForkingIndication
	isSet bool
}

func (v NullableSipForkingIndication) Get() *SipForkingIndication {
	return v.value
}

func (v *NullableSipForkingIndication) Set(val *SipForkingIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableSipForkingIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableSipForkingIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipForkingIndication(val *SipForkingIndication) *NullableSipForkingIndication {
	return &NullableSipForkingIndication{value: val, isSet: true}
}

func (v NullableSipForkingIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipForkingIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
