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

// checks if the OutOfCreditInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutOfCreditInformation{}

// OutOfCreditInformation Indicates the SDFs without available credit and the corresponding termination action.
type OutOfCreditInformation struct {
	FinUnitAct FinalUnitAction `json:"finUnitAct"`
	Flows      []Flows         `json:"flows,omitempty"`
}

type _OutOfCreditInformation OutOfCreditInformation

// NewOutOfCreditInformation instantiates a new OutOfCreditInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutOfCreditInformation(finUnitAct FinalUnitAction) *OutOfCreditInformation {
	this := OutOfCreditInformation{}
	this.FinUnitAct = finUnitAct
	return &this
}

// NewOutOfCreditInformationWithDefaults instantiates a new OutOfCreditInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutOfCreditInformationWithDefaults() *OutOfCreditInformation {
	this := OutOfCreditInformation{}
	return &this
}

// GetFinUnitAct returns the FinUnitAct field value
func (o *OutOfCreditInformation) GetFinUnitAct() FinalUnitAction {
	if o == nil {
		var ret FinalUnitAction
		return ret
	}

	return o.FinUnitAct
}

// GetFinUnitActOk returns a tuple with the FinUnitAct field value
// and a boolean to check if the value has been set.
func (o *OutOfCreditInformation) GetFinUnitActOk() (*FinalUnitAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinUnitAct, true
}

// SetFinUnitAct sets field value
func (o *OutOfCreditInformation) SetFinUnitAct(v FinalUnitAction) {
	o.FinUnitAct = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *OutOfCreditInformation) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutOfCreditInformation) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *OutOfCreditInformation) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *OutOfCreditInformation) SetFlows(v []Flows) {
	o.Flows = v
}

func (o OutOfCreditInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutOfCreditInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finUnitAct"] = o.FinUnitAct
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

func (o *OutOfCreditInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finUnitAct",
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

	varOutOfCreditInformation := _OutOfCreditInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutOfCreditInformation)

	if err != nil {
		return err
	}

	*o = OutOfCreditInformation(varOutOfCreditInformation)

	return err
}

type NullableOutOfCreditInformation struct {
	value *OutOfCreditInformation
	isSet bool
}

func (v NullableOutOfCreditInformation) Get() *OutOfCreditInformation {
	return v.value
}

func (v *NullableOutOfCreditInformation) Set(val *OutOfCreditInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableOutOfCreditInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableOutOfCreditInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutOfCreditInformation(val *OutOfCreditInformation) *NullableOutOfCreditInformation {
	return &NullableOutOfCreditInformation{value: val, isSet: true}
}

func (v NullableOutOfCreditInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutOfCreditInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
