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

// ServiceName Service names known to NRF
type ServiceName struct {
	ServiceNameAnyOf *ServiceNameAnyOf
	string           *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceName) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ServiceNameAnyOf
	err = json.Unmarshal(data, &dst.ServiceNameAnyOf)
	if err == nil {
		jsonServiceNameAnyOf, _ := json.Marshal(dst.ServiceNameAnyOf)
		if string(jsonServiceNameAnyOf) == "{}" { // empty struct
			dst.ServiceNameAnyOf = nil
		} else {
			return nil // data stored in dst.ServiceNameAnyOf, return on the first match
		}
	} else {
		dst.ServiceNameAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceName)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceName) MarshalJSON() ([]byte, error) {
	if src.ServiceNameAnyOf != nil {
		return json.Marshal(&src.ServiceNameAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceName struct {
	value *ServiceName
	isSet bool
}

func (v NullableServiceName) Get() *ServiceName {
	return v.value
}

func (v *NullableServiceName) Set(val *ServiceName) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceName) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceName(val *ServiceName) *NullableServiceName {
	return &NullableServiceName{value: val, isSet: true}
}

func (v NullableServiceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
