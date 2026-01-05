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

// checks if the RedundantPduSessionInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundantPduSessionInformation{}

// RedundantPduSessionInformation Redundant PDU Session Information
type RedundantPduSessionInformation struct {
	Rsn              Rsn    `json:"rsn"`
	PduSessionPairId *int32 `json:"pduSessionPairId,omitempty"`
}

type _RedundantPduSessionInformation RedundantPduSessionInformation

// NewRedundantPduSessionInformation instantiates a new RedundantPduSessionInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundantPduSessionInformation(rsn Rsn) *RedundantPduSessionInformation {
	this := RedundantPduSessionInformation{}
	this.Rsn = rsn
	return &this
}

// NewRedundantPduSessionInformationWithDefaults instantiates a new RedundantPduSessionInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundantPduSessionInformationWithDefaults() *RedundantPduSessionInformation {
	this := RedundantPduSessionInformation{}
	return &this
}

// GetRsn returns the Rsn field value
func (o *RedundantPduSessionInformation) GetRsn() Rsn {
	if o == nil {
		var ret Rsn
		return ret
	}

	return o.Rsn
}

// GetRsnOk returns a tuple with the Rsn field value
// and a boolean to check if the value has been set.
func (o *RedundantPduSessionInformation) GetRsnOk() (*Rsn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rsn, true
}

// SetRsn sets field value
func (o *RedundantPduSessionInformation) SetRsn(v Rsn) {
	o.Rsn = v
}

// GetPduSessionPairId returns the PduSessionPairId field value if set, zero value otherwise.
func (o *RedundantPduSessionInformation) GetPduSessionPairId() int32 {
	if o == nil || IsNil(o.PduSessionPairId) {
		var ret int32
		return ret
	}
	return *o.PduSessionPairId
}

// GetPduSessionPairIdOk returns a tuple with the PduSessionPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantPduSessionInformation) GetPduSessionPairIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSessionPairId) {
		return nil, false
	}
	return o.PduSessionPairId, true
}

// HasPduSessionPairId returns a boolean if a field has been set.
func (o *RedundantPduSessionInformation) HasPduSessionPairId() bool {
	if o != nil && !IsNil(o.PduSessionPairId) {
		return true
	}

	return false
}

// SetPduSessionPairId gets a reference to the given int32 and assigns it to the PduSessionPairId field.
func (o *RedundantPduSessionInformation) SetPduSessionPairId(v int32) {
	o.PduSessionPairId = &v
}

func (o RedundantPduSessionInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundantPduSessionInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rsn"] = o.Rsn
	if !IsNil(o.PduSessionPairId) {
		toSerialize["pduSessionPairId"] = o.PduSessionPairId
	}
	return toSerialize, nil
}

func (o *RedundantPduSessionInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rsn",
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

	varRedundantPduSessionInformation := _RedundantPduSessionInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRedundantPduSessionInformation)

	if err != nil {
		return err
	}

	*o = RedundantPduSessionInformation(varRedundantPduSessionInformation)

	return err
}

type NullableRedundantPduSessionInformation struct {
	value *RedundantPduSessionInformation
	isSet bool
}

func (v NullableRedundantPduSessionInformation) Get() *RedundantPduSessionInformation {
	return v.value
}

func (v *NullableRedundantPduSessionInformation) Set(val *RedundantPduSessionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantPduSessionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantPduSessionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantPduSessionInformation(val *RedundantPduSessionInformation) *NullableRedundantPduSessionInformation {
	return &NullableRedundantPduSessionInformation{value: val, isSet: true}
}

func (v NullableRedundantPduSessionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantPduSessionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
