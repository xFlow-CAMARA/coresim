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

// NotificationMethod Possible values are - PERIODIC - ONE_TIME - ON_EVENT_DETECTION
type NotificationMethod struct {
	NotificationMethodAnyOf *NotificationMethodAnyOf
	string                  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationMethodAnyOf
	err = json.Unmarshal(data, &dst.NotificationMethodAnyOf)
	if err == nil {
		jsonNotificationMethodAnyOf, _ := json.Marshal(dst.NotificationMethodAnyOf)
		if string(jsonNotificationMethodAnyOf) == "{}" { // empty struct
			dst.NotificationMethodAnyOf = nil
		} else {
			return nil // data stored in dst.NotificationMethodAnyOf, return on the first match
		}
	} else {
		dst.NotificationMethodAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationMethod) MarshalJSON() ([]byte, error) {
	if src.NotificationMethodAnyOf != nil {
		return json.Marshal(&src.NotificationMethodAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationMethod struct {
	value *NotificationMethod
	isSet bool
}

func (v NullableNotificationMethod) Get() *NotificationMethod {
	return v.value
}

func (v *NullableNotificationMethod) Set(val *NotificationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMethod(val *NotificationMethod) *NullableNotificationMethod {
	return &NullableNotificationMethod{value: val, isSet: true}
}

func (v NullableNotificationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
