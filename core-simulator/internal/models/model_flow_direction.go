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

// FlowDirection Possible values are - DOWNLINK: The corresponding filter applies for traffic to the UE. - UPLINK: The corresponding filter applies for traffic from the UE. - BIDIRECTIONAL: The corresponding filter applies for traffic both to and from the UE. - UNSPECIFIED: The corresponding filter applies for traffic to the UE (downlink), but has no specific direction declared. The service data flow detection shall apply the filter for uplink traffic as if the filter was bidirectional. The PCF shall not use the value UNSPECIFIED in filters created by the network in NW-initiated procedures. The PCF shall only include the value UNSPECIFIED in filters in UE-initiated procedures if the same value is received from the SMF.
type FlowDirection struct {
	FlowDirectionAnyOf *FlowDirectionAnyOf
	string             *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FlowDirection) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FlowDirectionAnyOf
	err = json.Unmarshal(data, &dst.FlowDirectionAnyOf)
	if err == nil {
		jsonFlowDirectionAnyOf, _ := json.Marshal(dst.FlowDirectionAnyOf)
		if string(jsonFlowDirectionAnyOf) == "{}" { // empty struct
			dst.FlowDirectionAnyOf = nil
		} else {
			return nil // data stored in dst.FlowDirectionAnyOf, return on the first match
		}
	} else {
		dst.FlowDirectionAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FlowDirection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FlowDirection) MarshalJSON() ([]byte, error) {
	if src.FlowDirectionAnyOf != nil {
		return json.Marshal(&src.FlowDirectionAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFlowDirection struct {
	value *FlowDirection
	isSet bool
}

func (v NullableFlowDirection) Get() *FlowDirection {
	return v.value
}

func (v *NullableFlowDirection) Set(val *FlowDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDirection(val *FlowDirection) *NullableFlowDirection {
	return &NullableFlowDirection{value: val, isSet: true}
}

func (v NullableFlowDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
