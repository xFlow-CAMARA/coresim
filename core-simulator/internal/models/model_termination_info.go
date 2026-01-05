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

// checks if the TerminationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminationInfo{}

// TerminationInfo Indicates the cause for requesting the deletion of the Individual Application Session Context resource.
type TerminationInfo struct {
	TermCause TerminationCause `json:"termCause"`
	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`
}

type _TerminationInfo TerminationInfo

// NewTerminationInfo instantiates a new TerminationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminationInfo(termCause TerminationCause, resUri string) *TerminationInfo {
	this := TerminationInfo{}
	this.TermCause = termCause
	this.ResUri = resUri
	return &this
}

// NewTerminationInfoWithDefaults instantiates a new TerminationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminationInfoWithDefaults() *TerminationInfo {
	this := TerminationInfo{}
	return &this
}

// GetTermCause returns the TermCause field value
func (o *TerminationInfo) GetTermCause() TerminationCause {
	if o == nil {
		var ret TerminationCause
		return ret
	}

	return o.TermCause
}

// GetTermCauseOk returns a tuple with the TermCause field value
// and a boolean to check if the value has been set.
func (o *TerminationInfo) GetTermCauseOk() (*TerminationCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermCause, true
}

// SetTermCause sets field value
func (o *TerminationInfo) SetTermCause(v TerminationCause) {
	o.TermCause = v
}

// GetResUri returns the ResUri field value
func (o *TerminationInfo) GetResUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResUri
}

// GetResUriOk returns a tuple with the ResUri field value
// and a boolean to check if the value has been set.
func (o *TerminationInfo) GetResUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResUri, true
}

// SetResUri sets field value
func (o *TerminationInfo) SetResUri(v string) {
	o.ResUri = v
}

func (o TerminationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["termCause"] = o.TermCause
	toSerialize["resUri"] = o.ResUri
	return toSerialize, nil
}

func (o *TerminationInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"termCause",
		"resUri",
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

	varTerminationInfo := _TerminationInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTerminationInfo)

	if err != nil {
		return err
	}

	*o = TerminationInfo(varTerminationInfo)

	return err
}

type NullableTerminationInfo struct {
	value *TerminationInfo
	isSet bool
}

func (v NullableTerminationInfo) Get() *TerminationInfo {
	return v.value
}

func (v *NullableTerminationInfo) Set(val *TerminationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationInfo(val *TerminationInfo) *NullableTerminationInfo {
	return &NullableTerminationInfo{value: val, isSet: true}
}

func (v NullableTerminationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
