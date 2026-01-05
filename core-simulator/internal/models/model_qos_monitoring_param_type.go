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

// QosMonitoringParamType Indicates the QoS monitoring parameter type.   Possible values are: - PACKET_DELAY: Indicates that the QoS monitoring parameter to be measured is packet delay. - CONGESTION: Indicates that the QoS monitoring parameter to be measured is congestion. - DATA_RATE: Indicates that the QoS monitoring parameter to be measured is data rate.
type QosMonitoringParamType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QosMonitoringParamType) UnmarshalJSON(data []byte) error {

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

	return fmt.Errorf("data failed to match schemas in anyOf(QosMonitoringParamType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *QosMonitoringParamType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQosMonitoringParamType struct {
	value *QosMonitoringParamType
	isSet bool
}

func (v NullableQosMonitoringParamType) Get() *QosMonitoringParamType {
	return v.value
}

func (v *NullableQosMonitoringParamType) Set(val *QosMonitoringParamType) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringParamType) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringParamType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringParamType(val *QosMonitoringParamType) *NullableQosMonitoringParamType {
	return &NullableQosMonitoringParamType{value: val, isSet: true}
}

func (v NullableQosMonitoringParamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringParamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
