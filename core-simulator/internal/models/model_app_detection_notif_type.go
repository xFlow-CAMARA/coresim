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

// AppDetectionNotifType Indicates the notification type for Application Detection Control.
type AppDetectionNotifType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppDetectionNotifType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AppDetectionNotifType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppDetectionNotifType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppDetectionNotifType struct {
	value *AppDetectionNotifType
	isSet bool
}

func (v NullableAppDetectionNotifType) Get() *AppDetectionNotifType {
	return v.value
}

func (v *NullableAppDetectionNotifType) Set(val *AppDetectionNotifType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetectionNotifType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetectionNotifType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetectionNotifType(val *AppDetectionNotifType) *NullableAppDetectionNotifType {
	return &NullableAppDetectionNotifType{value: val, isSet: true}
}

func (v NullableAppDetectionNotifType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetectionNotifType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
