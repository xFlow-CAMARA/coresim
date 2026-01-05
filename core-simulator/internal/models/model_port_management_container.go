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

// checks if the PortManagementContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortManagementContainer{}

// PortManagementContainer Contains the port management information container for a port.
type PortManagementContainer struct {
	// string with format 'bytes' as defined in OpenAPI
	PortManCont string `json:"portManCont"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNum int32 `json:"portNum"`
}

type _PortManagementContainer PortManagementContainer

// NewPortManagementContainer instantiates a new PortManagementContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortManagementContainer(portManCont string, portNum int32) *PortManagementContainer {
	this := PortManagementContainer{}
	this.PortManCont = portManCont
	this.PortNum = portNum
	return &this
}

// NewPortManagementContainerWithDefaults instantiates a new PortManagementContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortManagementContainerWithDefaults() *PortManagementContainer {
	this := PortManagementContainer{}
	return &this
}

// GetPortManCont returns the PortManCont field value
func (o *PortManagementContainer) GetPortManCont() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortManCont
}

// GetPortManContOk returns a tuple with the PortManCont field value
// and a boolean to check if the value has been set.
func (o *PortManagementContainer) GetPortManContOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortManCont, true
}

// SetPortManCont sets field value
func (o *PortManagementContainer) SetPortManCont(v string) {
	o.PortManCont = v
}

// GetPortNum returns the PortNum field value
func (o *PortManagementContainer) GetPortNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortNum
}

// GetPortNumOk returns a tuple with the PortNum field value
// and a boolean to check if the value has been set.
func (o *PortManagementContainer) GetPortNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortNum, true
}

// SetPortNum sets field value
func (o *PortManagementContainer) SetPortNum(v int32) {
	o.PortNum = v
}

func (o PortManagementContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortManagementContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["portManCont"] = o.PortManCont
	toSerialize["portNum"] = o.PortNum
	return toSerialize, nil
}

func (o *PortManagementContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"portManCont",
		"portNum",
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

	varPortManagementContainer := _PortManagementContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortManagementContainer)

	if err != nil {
		return err
	}

	*o = PortManagementContainer(varPortManagementContainer)

	return err
}

type NullablePortManagementContainer struct {
	value *PortManagementContainer
	isSet bool
}

func (v NullablePortManagementContainer) Get() *PortManagementContainer {
	return v.value
}

func (v *NullablePortManagementContainer) Set(val *PortManagementContainer) {
	v.value = val
	v.isSet = true
}

func (v NullablePortManagementContainer) IsSet() bool {
	return v.isSet
}

func (v *NullablePortManagementContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortManagementContainer(val *PortManagementContainer) *NullablePortManagementContainer {
	return &NullablePortManagementContainer{value: val, isSet: true}
}

func (v NullablePortManagementContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortManagementContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
