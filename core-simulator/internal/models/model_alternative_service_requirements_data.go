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

// checks if the AlternativeServiceRequirementsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeServiceRequirementsData{}

// AlternativeServiceRequirementsData Contains an alternative QoS related parameter set.
type AlternativeServiceRequirementsData struct {
	// Reference to this alternative QoS related parameter set.
	AltQosParamSetRef string `json:"altQosParamSetRef"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrUl *string `json:"gbrUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl *string `json:"gbrDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb *int32 `json:"pdb,omitempty"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	Per *string `json:"per,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`
}

type _AlternativeServiceRequirementsData AlternativeServiceRequirementsData

// NewAlternativeServiceRequirementsData instantiates a new AlternativeServiceRequirementsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeServiceRequirementsData(altQosParamSetRef string) *AlternativeServiceRequirementsData {
	this := AlternativeServiceRequirementsData{}
	this.AltQosParamSetRef = altQosParamSetRef
	return &this
}

// NewAlternativeServiceRequirementsDataWithDefaults instantiates a new AlternativeServiceRequirementsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeServiceRequirementsDataWithDefaults() *AlternativeServiceRequirementsData {
	this := AlternativeServiceRequirementsData{}
	return &this
}

// GetAltQosParamSetRef returns the AltQosParamSetRef field value
func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AltQosParamSetRef
}

// GetAltQosParamSetRefOk returns a tuple with the AltQosParamSetRef field value
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AltQosParamSetRef, true
}

// SetAltQosParamSetRef sets field value
func (o *AlternativeServiceRequirementsData) SetAltQosParamSetRef(v string) {
	o.AltQosParamSetRef = v
}

// GetGbrUl returns the GbrUl field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetGbrUl() string {
	if o == nil || IsNil(o.GbrUl) {
		var ret string
		return ret
	}
	return *o.GbrUl
}

// GetGbrUlOk returns a tuple with the GbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetGbrUlOk() (*string, bool) {
	if o == nil || IsNil(o.GbrUl) {
		return nil, false
	}
	return o.GbrUl, true
}

// HasGbrUl returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasGbrUl() bool {
	if o != nil && !IsNil(o.GbrUl) {
		return true
	}

	return false
}

// SetGbrUl gets a reference to the given string and assigns it to the GbrUl field.
func (o *AlternativeServiceRequirementsData) SetGbrUl(v string) {
	o.GbrUl = &v
}

// GetGbrDl returns the GbrDl field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetGbrDl() string {
	if o == nil || IsNil(o.GbrDl) {
		var ret string
		return ret
	}
	return *o.GbrDl
}

// GetGbrDlOk returns a tuple with the GbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetGbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.GbrDl) {
		return nil, false
	}
	return o.GbrDl, true
}

// HasGbrDl returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasGbrDl() bool {
	if o != nil && !IsNil(o.GbrDl) {
		return true
	}

	return false
}

// SetGbrDl gets a reference to the given string and assigns it to the GbrDl field.
func (o *AlternativeServiceRequirementsData) SetGbrDl(v string) {
	o.GbrDl = &v
}

// GetPdb returns the Pdb field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetPdb() int32 {
	if o == nil || IsNil(o.Pdb) {
		var ret int32
		return ret
	}
	return *o.Pdb
}

// GetPdbOk returns a tuple with the Pdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetPdbOk() (*int32, bool) {
	if o == nil || IsNil(o.Pdb) {
		return nil, false
	}
	return o.Pdb, true
}

// HasPdb returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasPdb() bool {
	if o != nil && !IsNil(o.Pdb) {
		return true
	}

	return false
}

// SetPdb gets a reference to the given int32 and assigns it to the Pdb field.
func (o *AlternativeServiceRequirementsData) SetPdb(v int32) {
	o.Pdb = &v
}

// GetPer returns the Per field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetPer() string {
	if o == nil || IsNil(o.Per) {
		var ret string
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetPerOk() (*string, bool) {
	if o == nil || IsNil(o.Per) {
		return nil, false
	}
	return o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasPer() bool {
	if o != nil && !IsNil(o.Per) {
		return true
	}

	return false
}

// SetPer gets a reference to the given string and assigns it to the Per field.
func (o *AlternativeServiceRequirementsData) SetPer(v string) {
	o.Per = &v
}

func (o AlternativeServiceRequirementsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeServiceRequirementsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["altQosParamSetRef"] = o.AltQosParamSetRef
	if !IsNil(o.GbrUl) {
		toSerialize["gbrUl"] = o.GbrUl
	}
	if !IsNil(o.GbrDl) {
		toSerialize["gbrDl"] = o.GbrDl
	}
	if !IsNil(o.Pdb) {
		toSerialize["pdb"] = o.Pdb
	}
	if !IsNil(o.Per) {
		toSerialize["per"] = o.Per
	}
	return toSerialize, nil
}

func (o *AlternativeServiceRequirementsData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"altQosParamSetRef",
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

	varAlternativeServiceRequirementsData := _AlternativeServiceRequirementsData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlternativeServiceRequirementsData)

	if err != nil {
		return err
	}

	*o = AlternativeServiceRequirementsData(varAlternativeServiceRequirementsData)

	return err
}

type NullableAlternativeServiceRequirementsData struct {
	value *AlternativeServiceRequirementsData
	isSet bool
}

func (v NullableAlternativeServiceRequirementsData) Get() *AlternativeServiceRequirementsData {
	return v.value
}

func (v *NullableAlternativeServiceRequirementsData) Set(val *AlternativeServiceRequirementsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeServiceRequirementsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeServiceRequirementsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeServiceRequirementsData(val *AlternativeServiceRequirementsData) *NullableAlternativeServiceRequirementsData {
	return &NullableAlternativeServiceRequirementsData{value: val, isSet: true}
}

func (v NullableAlternativeServiceRequirementsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeServiceRequirementsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
