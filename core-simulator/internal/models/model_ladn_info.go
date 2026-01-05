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
)

// LadnInfo LADN Information
type LadnInfo struct {
	Ladn     string         `json:"ladn"`
	Presence *PresenceState `json:"presence,omitempty"`
}

// NewLadnInfo instantiates a new LadnInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLadnInfo(ladn string) *LadnInfo {
	this := LadnInfo{}
	this.Ladn = ladn
	return &this
}

// NewLadnInfoWithDefaults instantiates a new LadnInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLadnInfoWithDefaults() *LadnInfo {
	this := LadnInfo{}
	return &this
}

// GetLadn returns the Ladn field value
func (o *LadnInfo) GetLadn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ladn
}

// GetLadnOk returns a tuple with the Ladn field value
// and a boolean to check if the value has been set.
func (o *LadnInfo) GetLadnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ladn, true
}

// SetLadn sets field value
func (o *LadnInfo) SetLadn(v string) {
	o.Ladn = v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *LadnInfo) GetPresence() PresenceState {
	if o == nil || o.Presence == nil {
		var ret PresenceState
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LadnInfo) GetPresenceOk() (*PresenceState, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *LadnInfo) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given PresenceState and assigns it to the Presence field.
func (o *LadnInfo) SetPresence(v PresenceState) {
	o.Presence = &v
}

func (o LadnInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ladn"] = o.Ladn
	}
	if o.Presence != nil {
		toSerialize["presence"] = o.Presence
	}
	return json.Marshal(toSerialize)
}

type NullableLadnInfo struct {
	value *LadnInfo
	isSet bool
}

func (v NullableLadnInfo) Get() *LadnInfo {
	return v.value
}

func (v *NullableLadnInfo) Set(val *LadnInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLadnInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLadnInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLadnInfo(val *LadnInfo) *NullableLadnInfo {
	return &NullableLadnInfo{value: val, isSet: true}
}

func (v NullableLadnInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLadnInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
