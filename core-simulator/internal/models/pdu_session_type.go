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

package models

import (
	"encoding/json"
	"fmt"
)

// PduSessionType PduSessionType indicates the type of a PDU session. It shall comply with the provisions defined in table 5.4.3.3-1.
type PduSessionType struct {
	PduSessionTypeAnyOf *PduSessionTypeAnyOf
	string              *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PduSessionTypeAnyOf
	err = json.Unmarshal(data, &dst.PduSessionTypeAnyOf)
	if err == nil {
		jsonPduSessionTypeAnyOf, _ := json.Marshal(dst.PduSessionTypeAnyOf)
		if string(jsonPduSessionTypeAnyOf) == "{}" { // empty struct
			dst.PduSessionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.PduSessionTypeAnyOf, return on the first match
		}
	} else {
		dst.PduSessionTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionType) MarshalJSON() ([]byte, error) {
	if src.PduSessionTypeAnyOf != nil {
		return json.Marshal(&src.PduSessionTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionType struct {
	value *PduSessionType
	isSet bool
}

func (v NullablePduSessionType) Get() *PduSessionType {
	return v.value
}

func (v *NullablePduSessionType) Set(val *PduSessionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionType(val *PduSessionType) *NullablePduSessionType {
	return &NullablePduSessionType{value: val, isSet: true}
}

func (v NullablePduSessionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
