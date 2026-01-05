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

// AppliedSmccType Possible values are - DNN_CC: Indicates the DNN based congestion control. - SNSSAI_CC: Indicates the S-NSSAI based congestion control.
type AppliedSmccType struct {
	AppliedSmccTypeAnyOf *AppliedSmccTypeAnyOf
	string               *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppliedSmccType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AppliedSmccTypeAnyOf
	err = json.Unmarshal(data, &dst.AppliedSmccTypeAnyOf)
	if err == nil {
		jsonAppliedSmccTypeAnyOf, _ := json.Marshal(dst.AppliedSmccTypeAnyOf)
		if string(jsonAppliedSmccTypeAnyOf) == "{}" { // empty struct
			dst.AppliedSmccTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AppliedSmccTypeAnyOf, return on the first match
		}
	} else {
		dst.AppliedSmccTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AppliedSmccType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppliedSmccType) MarshalJSON() ([]byte, error) {
	if src.AppliedSmccTypeAnyOf != nil {
		return json.Marshal(&src.AppliedSmccTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppliedSmccType struct {
	value *AppliedSmccType
	isSet bool
}

func (v NullableAppliedSmccType) Get() *AppliedSmccType {
	return v.value
}

func (v *NullableAppliedSmccType) Set(val *AppliedSmccType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedSmccType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedSmccType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedSmccType(val *AppliedSmccType) *NullableAppliedSmccType {
	return &NullableAppliedSmccType{value: val, isSet: true}
}

func (v NullableAppliedSmccType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedSmccType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
