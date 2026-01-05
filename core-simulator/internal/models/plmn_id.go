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
)

// PlmnId When PlmnId needs to be converted to string (e.g. when used in maps as key), the string  shall be composed of three digits \"mcc\" followed by \"-\" and two or three digits \"mnc\".
type PlmnId struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mcc string `json:"mcc"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
}

// NewPlmnId instantiates a new PlmnId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnId(mcc string, mnc string) *PlmnId {
	this := PlmnId{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnIdWithDefaults instantiates a new PlmnId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdWithDefaults() *PlmnId {
	this := PlmnId{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnId) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnId) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnId) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnId) SetMnc(v string) {
	o.Mnc = v
}

func (o PlmnId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mcc"] = o.Mcc
	}
	if true {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnId struct {
	value *PlmnId
	isSet bool
}

func (v NullablePlmnId) Get() *PlmnId {
	return v.value
}

func (v *NullablePlmnId) Set(val *PlmnId) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnId) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnId(val *PlmnId) *NullablePlmnId {
	return &NullablePlmnId{value: val, isSet: true}
}

func (v NullablePlmnId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
