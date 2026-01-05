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
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
	"fmt"
)

// AfResultStatus Possible values are: - SUCCESS: The application layer is ready or the relocation is completed. - TEMPORARY_CONGESTION: The application relocation fails due to temporary congestion. - RELOC_NO_ALLOWED: The application relocation fails because application relocation is not allowed. - OTHER: The application relocation fails due to other reason.
type AfResultStatus struct {
	AfResultStatusAnyOf *AfResultStatusAnyOf
	string              *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfResultStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AfResultStatusAnyOf
	err = json.Unmarshal(data, &dst.AfResultStatusAnyOf)
	if err == nil {
		jsonAfResultStatusAnyOf, _ := json.Marshal(dst.AfResultStatusAnyOf)
		if string(jsonAfResultStatusAnyOf) == "{}" { // empty struct
			dst.AfResultStatusAnyOf = nil
		} else {
			return nil // data stored in dst.AfResultStatusAnyOf, return on the first match
		}
	} else {
		dst.AfResultStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfResultStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfResultStatus) MarshalJSON() ([]byte, error) {
	if src.AfResultStatusAnyOf != nil {
		return json.Marshal(&src.AfResultStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfResultStatus struct {
	value *AfResultStatus
	isSet bool
}

func (v NullableAfResultStatus) Get() *AfResultStatus {
	return v.value
}

func (v *NullableAfResultStatus) Set(val *AfResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAfResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAfResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfResultStatus(val *AfResultStatus) *NullableAfResultStatus {
	return &NullableAfResultStatus{value: val, isSet: true}
}

func (v NullableAfResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
