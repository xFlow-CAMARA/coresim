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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NtnTaiInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtnTaiInfo{}

// NtnTaiInfo struct for NtnTaiInfo
type NtnTaiInfo struct {
	PlmnId  PlmnIdNid `json:"plmnId"`
	TacList []string  `json:"tacList"`
	// 2 or 3-octet string identifying a tracking area code as specified in clause 9.3.3.10  of 3GPP TS 38.413, in hexadecimal representation. Each character in the string shall  take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the TAC shall  appear first in the string, and the character representing the 4 least significant bit  of the TAC shall appear last in the string.
	DerivedTac *string `json:"derivedTac,omitempty" validate:"regexp=(^[A-Fa-f0-9]{4}$)|(^[A-Fa-f0-9]{6}$)"`
}

type _NtnTaiInfo NtnTaiInfo

// NewNtnTaiInfo instantiates a new NtnTaiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtnTaiInfo(plmnId PlmnIdNid, tacList []string) *NtnTaiInfo {
	this := NtnTaiInfo{}
	this.PlmnId = plmnId
	this.TacList = tacList
	return &this
}

// NewNtnTaiInfoWithDefaults instantiates a new NtnTaiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtnTaiInfoWithDefaults() *NtnTaiInfo {
	this := NtnTaiInfo{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *NtnTaiInfo) GetPlmnId() PlmnIdNid {
	if o == nil {
		var ret PlmnIdNid
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *NtnTaiInfo) GetPlmnIdOk() (*PlmnIdNid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *NtnTaiInfo) SetPlmnId(v PlmnIdNid) {
	o.PlmnId = v
}

// GetTacList returns the TacList field value
func (o *NtnTaiInfo) GetTacList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TacList
}

// GetTacListOk returns a tuple with the TacList field value
// and a boolean to check if the value has been set.
func (o *NtnTaiInfo) GetTacListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TacList, true
}

// SetTacList sets field value
func (o *NtnTaiInfo) SetTacList(v []string) {
	o.TacList = v
}

// GetDerivedTac returns the DerivedTac field value if set, zero value otherwise.
func (o *NtnTaiInfo) GetDerivedTac() string {
	if o == nil || IsNil(o.DerivedTac) {
		var ret string
		return ret
	}
	return *o.DerivedTac
}

// GetDerivedTacOk returns a tuple with the DerivedTac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtnTaiInfo) GetDerivedTacOk() (*string, bool) {
	if o == nil || IsNil(o.DerivedTac) {
		return nil, false
	}
	return o.DerivedTac, true
}

// HasDerivedTac returns a boolean if a field has been set.
func (o *NtnTaiInfo) HasDerivedTac() bool {
	if o != nil && !IsNil(o.DerivedTac) {
		return true
	}

	return false
}

// SetDerivedTac gets a reference to the given string and assigns it to the DerivedTac field.
func (o *NtnTaiInfo) SetDerivedTac(v string) {
	o.DerivedTac = &v
}

func (o NtnTaiInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtnTaiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["tacList"] = o.TacList
	if !IsNil(o.DerivedTac) {
		toSerialize["derivedTac"] = o.DerivedTac
	}
	return toSerialize, nil
}

func (o *NtnTaiInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plmnId",
		"tacList",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNtnTaiInfo := _NtnTaiInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNtnTaiInfo)

	if err != nil {
		return err
	}

	*o = NtnTaiInfo(varNtnTaiInfo)

	return err
}

type NullableNtnTaiInfo struct {
	value *NtnTaiInfo
	isSet bool
}

func (v NullableNtnTaiInfo) Get() *NtnTaiInfo {
	return v.value
}

func (v *NullableNtnTaiInfo) Set(val *NtnTaiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNtnTaiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNtnTaiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtnTaiInfo(val *NtnTaiInfo) *NullableNtnTaiInfo {
	return &NullableNtnTaiInfo{value: val, isSet: true}
}

func (v NullableNtnTaiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtnTaiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
