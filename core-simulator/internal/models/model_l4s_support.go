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

// checks if the L4sSupport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L4sSupport{}

// L4sSupport Indicates whether the ECN marking for L4S support is not available or available again in 5GS.
type L4sSupport struct {
	NotifType L4sNotifType `json:"notifType"`
	Flows     []Flows      `json:"flows,omitempty"`
}

type _L4sSupport L4sSupport

// NewL4sSupport instantiates a new L4sSupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL4sSupport(notifType L4sNotifType) *L4sSupport {
	this := L4sSupport{}
	this.NotifType = notifType
	return &this
}

// NewL4sSupportWithDefaults instantiates a new L4sSupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL4sSupportWithDefaults() *L4sSupport {
	this := L4sSupport{}
	return &this
}

// GetNotifType returns the NotifType field value
func (o *L4sSupport) GetNotifType() L4sNotifType {
	if o == nil {
		var ret L4sNotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *L4sSupport) GetNotifTypeOk() (*L4sNotifType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *L4sSupport) SetNotifType(v L4sNotifType) {
	o.NotifType = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *L4sSupport) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L4sSupport) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *L4sSupport) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *L4sSupport) SetFlows(v []Flows) {
	o.Flows = v
}

func (o L4sSupport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L4sSupport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifType"] = o.NotifType
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

func (o *L4sSupport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifType",
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

	varL4sSupport := _L4sSupport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varL4sSupport)

	if err != nil {
		return err
	}

	*o = L4sSupport(varL4sSupport)

	return err
}

type NullableL4sSupport struct {
	value *L4sSupport
	isSet bool
}

func (v NullableL4sSupport) Get() *L4sSupport {
	return v.value
}

func (v *NullableL4sSupport) Set(val *L4sSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableL4sSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableL4sSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL4sSupport(val *L4sSupport) *NullableL4sSupport {
	return &NullableL4sSupport{value: val, isSet: true}
}

func (v NullableL4sSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL4sSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
