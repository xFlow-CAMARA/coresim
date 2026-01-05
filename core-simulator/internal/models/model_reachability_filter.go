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
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ReachabilityFilter Event filter for REACHABILITY_REPORT event type
type ReachabilityFilter struct {
	ReachabilityFilterAnyOf *ReachabilityFilterAnyOf
	string                  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityFilter) UnmarshalJSON(data []byte) error {
	// try to unmarshal JSON data into ReachabilityFilterAnyOf
	err := json.Unmarshal(data, &dst.ReachabilityFilterAnyOf)
	if err == nil {
		jsonReachabilityFilterAnyOf, _ := json.Marshal(dst.ReachabilityFilterAnyOf)
		if string(jsonReachabilityFilterAnyOf) == "{}" { // empty struct
			dst.ReachabilityFilterAnyOf = nil
		} else {
			return nil // data stored in dst.ReachabilityFilterAnyOf, return on the first match
		}
	} else {
		dst.ReachabilityFilterAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityFilter) MarshalJSON() ([]byte, error) {
	if src.ReachabilityFilterAnyOf != nil {
		return json.Marshal(&src.ReachabilityFilterAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityFilter struct {
	value *ReachabilityFilter
	isSet bool
}

func (v NullableReachabilityFilter) Get() *ReachabilityFilter {
	return v.value
}

func (v *NullableReachabilityFilter) Set(val *ReachabilityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityFilter(val *ReachabilityFilter) *NullableReachabilityFilter {
	return &NullableReachabilityFilter{value: val, isSet: true}
}

func (v NullableReachabilityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
