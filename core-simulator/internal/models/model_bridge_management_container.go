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

// checks if the BridgeManagementContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BridgeManagementContainer{}

// BridgeManagementContainer Contains the UMIC.
type BridgeManagementContainer struct {
	// string with format 'bytes' as defined in OpenAPI
	BridgeManCont string `json:"bridgeManCont"`
}

type _BridgeManagementContainer BridgeManagementContainer

// NewBridgeManagementContainer instantiates a new BridgeManagementContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBridgeManagementContainer(bridgeManCont string) *BridgeManagementContainer {
	this := BridgeManagementContainer{}
	this.BridgeManCont = bridgeManCont
	return &this
}

// NewBridgeManagementContainerWithDefaults instantiates a new BridgeManagementContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBridgeManagementContainerWithDefaults() *BridgeManagementContainer {
	this := BridgeManagementContainer{}
	return &this
}

// GetBridgeManCont returns the BridgeManCont field value
func (o *BridgeManagementContainer) GetBridgeManCont() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BridgeManCont
}

// GetBridgeManContOk returns a tuple with the BridgeManCont field value
// and a boolean to check if the value has been set.
func (o *BridgeManagementContainer) GetBridgeManContOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BridgeManCont, true
}

// SetBridgeManCont sets field value
func (o *BridgeManagementContainer) SetBridgeManCont(v string) {
	o.BridgeManCont = v
}

func (o BridgeManagementContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BridgeManagementContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bridgeManCont"] = o.BridgeManCont
	return toSerialize, nil
}

func (o *BridgeManagementContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bridgeManCont",
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

	varBridgeManagementContainer := _BridgeManagementContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBridgeManagementContainer)

	if err != nil {
		return err
	}

	*o = BridgeManagementContainer(varBridgeManagementContainer)

	return err
}

type NullableBridgeManagementContainer struct {
	value *BridgeManagementContainer
	isSet bool
}

func (v NullableBridgeManagementContainer) Get() *BridgeManagementContainer {
	return v.value
}

func (v *NullableBridgeManagementContainer) Set(val *BridgeManagementContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableBridgeManagementContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableBridgeManagementContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBridgeManagementContainer(val *BridgeManagementContainer) *NullableBridgeManagementContainer {
	return &NullableBridgeManagementContainer{value: val, isSet: true}
}

func (v NullableBridgeManagementContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBridgeManagementContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
