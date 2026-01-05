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

// checks if the AdditionalAccessInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalAccessInfo{}

// AdditionalAccessInfo Indicates the combination of additional Access Type and RAT Type for a MA PDU session.
type AdditionalAccessInfo struct {
	AccessType AccessType `json:"accessType"`
	RatType    *RatType   `json:"ratType,omitempty"`
}

type _AdditionalAccessInfo AdditionalAccessInfo

// NewAdditionalAccessInfo instantiates a new AdditionalAccessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalAccessInfo(accessType AccessType) *AdditionalAccessInfo {
	this := AdditionalAccessInfo{}
	this.AccessType = accessType
	return &this
}

// NewAdditionalAccessInfoWithDefaults instantiates a new AdditionalAccessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalAccessInfoWithDefaults() *AdditionalAccessInfo {
	this := AdditionalAccessInfo{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *AdditionalAccessInfo) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *AdditionalAccessInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *AdditionalAccessInfo) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *AdditionalAccessInfo) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalAccessInfo) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *AdditionalAccessInfo) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *AdditionalAccessInfo) SetRatType(v RatType) {
	o.RatType = &v
}

func (o AdditionalAccessInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalAccessInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessType"] = o.AccessType
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	return toSerialize, nil
}

func (o *AdditionalAccessInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessType",
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

	varAdditionalAccessInfo := _AdditionalAccessInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdditionalAccessInfo)

	if err != nil {
		return err
	}

	*o = AdditionalAccessInfo(varAdditionalAccessInfo)

	return err
}

type NullableAdditionalAccessInfo struct {
	value *AdditionalAccessInfo
	isSet bool
}

func (v NullableAdditionalAccessInfo) Get() *AdditionalAccessInfo {
	return v.value
}

func (v *NullableAdditionalAccessInfo) Set(val *AdditionalAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalAccessInfo(val *AdditionalAccessInfo) *NullableAdditionalAccessInfo {
	return &NullableAdditionalAccessInfo{value: val, isSet: true}
}

func (v NullableAdditionalAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
