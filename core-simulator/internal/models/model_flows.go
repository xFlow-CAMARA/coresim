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

// checks if the Flows type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Flows{}

// Flows Identifies the flows.
type Flows struct {
	ContVers []int32 `json:"contVers,omitempty"`
	FNums    []int32 `json:"fNums,omitempty"`
	MedCompN int32   `json:"medCompN"`
}

type _Flows Flows

// NewFlows instantiates a new Flows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlows(medCompN int32) *Flows {
	this := Flows{}
	this.MedCompN = medCompN
	return &this
}

// NewFlowsWithDefaults instantiates a new Flows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowsWithDefaults() *Flows {
	this := Flows{}
	return &this
}

// GetContVers returns the ContVers field value if set, zero value otherwise.
func (o *Flows) GetContVers() []int32 {
	if o == nil || IsNil(o.ContVers) {
		var ret []int32
		return ret
	}
	return o.ContVers
}

// GetContVersOk returns a tuple with the ContVers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flows) GetContVersOk() ([]int32, bool) {
	if o == nil || IsNil(o.ContVers) {
		return nil, false
	}
	return o.ContVers, true
}

// HasContVers returns a boolean if a field has been set.
func (o *Flows) HasContVers() bool {
	if o != nil && !IsNil(o.ContVers) {
		return true
	}

	return false
}

// SetContVers gets a reference to the given []int32 and assigns it to the ContVers field.
func (o *Flows) SetContVers(v []int32) {
	o.ContVers = v
}

// GetFNums returns the FNums field value if set, zero value otherwise.
func (o *Flows) GetFNums() []int32 {
	if o == nil || IsNil(o.FNums) {
		var ret []int32
		return ret
	}
	return o.FNums
}

// GetFNumsOk returns a tuple with the FNums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flows) GetFNumsOk() ([]int32, bool) {
	if o == nil || IsNil(o.FNums) {
		return nil, false
	}
	return o.FNums, true
}

// HasFNums returns a boolean if a field has been set.
func (o *Flows) HasFNums() bool {
	if o != nil && !IsNil(o.FNums) {
		return true
	}

	return false
}

// SetFNums gets a reference to the given []int32 and assigns it to the FNums field.
func (o *Flows) SetFNums(v []int32) {
	o.FNums = v
}

// GetMedCompN returns the MedCompN field value
func (o *Flows) GetMedCompN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MedCompN
}

// GetMedCompNOk returns a tuple with the MedCompN field value
// and a boolean to check if the value has been set.
func (o *Flows) GetMedCompNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MedCompN, true
}

// SetMedCompN sets field value
func (o *Flows) SetMedCompN(v int32) {
	o.MedCompN = v
}

func (o Flows) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Flows) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContVers) {
		toSerialize["contVers"] = o.ContVers
	}
	if !IsNil(o.FNums) {
		toSerialize["fNums"] = o.FNums
	}
	toSerialize["medCompN"] = o.MedCompN
	return toSerialize, nil
}

func (o *Flows) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"medCompN",
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

	varFlows := _Flows{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlows)

	if err != nil {
		return err
	}

	*o = Flows(varFlows)

	return err
}

type NullableFlows struct {
	value *Flows
	isSet bool
}

func (v NullableFlows) Get() *Flows {
	return v.value
}

func (v *NullableFlows) Set(val *Flows) {
	v.value = val
	v.isSet = true
}

func (v NullableFlows) IsSet() bool {
	return v.isSet
}

func (v *NullableFlows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlows(val *Flows) *NullableFlows {
	return &NullableFlows{value: val, isSet: true}
}

func (v NullableFlows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
