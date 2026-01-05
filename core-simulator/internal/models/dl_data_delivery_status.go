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

// DlDataDeliveryStatus Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is transmitted after previous buffering or discarding of corresponding packet(s) because the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is discarded because the Extended Buffering time, as determined by the SMF, expires or the amount of downlink data to be buffered is exceeded.
type DlDataDeliveryStatus struct {
	DlDataDeliveryStatusAnyOf *DlDataDeliveryStatusAnyOf
	string                    *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DlDataDeliveryStatus) UnmarshalJSON(data []byte) error {

	// try to unmarshal JSON data into DlDataDeliveryStatusAnyOf
	err := json.Unmarshal(data, &dst.DlDataDeliveryStatusAnyOf)
	if err == nil {
		jsonDlDataDeliveryStatusAnyOf, _ := json.Marshal(dst.DlDataDeliveryStatusAnyOf)
		if string(jsonDlDataDeliveryStatusAnyOf) == "{}" { // empty struct
			dst.DlDataDeliveryStatusAnyOf = nil
		} else {
			return nil // data stored in dst.DlDataDeliveryStatusAnyOf, return on the first match
		}
	} else {
		dst.DlDataDeliveryStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DlDataDeliveryStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DlDataDeliveryStatus) MarshalJSON() ([]byte, error) {
	if src.DlDataDeliveryStatusAnyOf != nil {
		return json.Marshal(&src.DlDataDeliveryStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDlDataDeliveryStatus struct {
	value *DlDataDeliveryStatus
	isSet bool
}

func (v NullableDlDataDeliveryStatus) Get() *DlDataDeliveryStatus {
	return v.value
}

func (v *NullableDlDataDeliveryStatus) Set(val *DlDataDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDlDataDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDlDataDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlDataDeliveryStatus(val *DlDataDeliveryStatus) *NullableDlDataDeliveryStatus {
	return &NullableDlDataDeliveryStatus{value: val, isSet: true}
}

func (v NullableDlDataDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlDataDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
