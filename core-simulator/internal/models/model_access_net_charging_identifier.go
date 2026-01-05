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
	"encoding/json"
)

// checks if the AccessNetChargingIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessNetChargingIdentifier{}

// AccessNetChargingIdentifier Describes the access network charging identifier.
type AccessNetChargingIdentifier struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	// Deprecated
	AccNetChaIdValue *int32
	// A character string containing the access network charging identifier.
	AccNetChargIdString *string
	Flows               []Flows
}

// NewAccessNetChargingIdentifier instantiates a new AccessNetChargingIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessNetChargingIdentifier() *AccessNetChargingIdentifier {
	this := AccessNetChargingIdentifier{}
	return &this
}

// NewAccessNetChargingIdentifierWithDefaults instantiates a new AccessNetChargingIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessNetChargingIdentifierWithDefaults() *AccessNetChargingIdentifier {
	this := AccessNetChargingIdentifier{}
	return &this
}

// GetAccNetChaIdValue returns the AccNetChaIdValue field value if set, zero value otherwise.
// Deprecated
func (o *AccessNetChargingIdentifier) GetAccNetChaIdValue() int32 {
	if o == nil || IsNil(o.AccNetChaIdValue) {
		var ret int32
		return ret
	}
	return *o.AccNetChaIdValue
}

// GetAccNetChaIdValueOk returns a tuple with the AccNetChaIdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccessNetChargingIdentifier) GetAccNetChaIdValueOk() (*int32, bool) {
	if o == nil || IsNil(o.AccNetChaIdValue) {
		return nil, false
	}
	return o.AccNetChaIdValue, true
}

// HasAccNetChaIdValue returns a boolean if a field has been set.
func (o *AccessNetChargingIdentifier) HasAccNetChaIdValue() bool {
	if o != nil && !IsNil(o.AccNetChaIdValue) {
		return true
	}

	return false
}

// SetAccNetChaIdValue gets a reference to the given int32 and assigns it to the AccNetChaIdValue field.
// Deprecated
func (o *AccessNetChargingIdentifier) SetAccNetChaIdValue(v int32) {
	o.AccNetChaIdValue = &v
}

// GetAccNetChargIdString returns the AccNetChargIdString field value if set, zero value otherwise.
func (o *AccessNetChargingIdentifier) GetAccNetChargIdString() string {
	if o == nil || IsNil(o.AccNetChargIdString) {
		var ret string
		return ret
	}
	return *o.AccNetChargIdString
}

// GetAccNetChargIdStringOk returns a tuple with the AccNetChargIdString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessNetChargingIdentifier) GetAccNetChargIdStringOk() (*string, bool) {
	if o == nil || IsNil(o.AccNetChargIdString) {
		return nil, false
	}
	return o.AccNetChargIdString, true
}

// HasAccNetChargIdString returns a boolean if a field has been set.
func (o *AccessNetChargingIdentifier) HasAccNetChargIdString() bool {
	if o != nil && !IsNil(o.AccNetChargIdString) {
		return true
	}

	return false
}

// SetAccNetChargIdString gets a reference to the given string and assigns it to the AccNetChargIdString field.
func (o *AccessNetChargingIdentifier) SetAccNetChargIdString(v string) {
	o.AccNetChargIdString = &v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *AccessNetChargingIdentifier) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessNetChargingIdentifier) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *AccessNetChargingIdentifier) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *AccessNetChargingIdentifier) SetFlows(v []Flows) {
	o.Flows = v
}

func (o AccessNetChargingIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessNetChargingIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccNetChaIdValue) {
		toSerialize["accNetChaIdValue"] = o.AccNetChaIdValue
	}
	if !IsNil(o.AccNetChargIdString) {
		toSerialize["accNetChargIdString"] = o.AccNetChargIdString
	}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

type NullableAccessNetChargingIdentifier struct {
	value *AccessNetChargingIdentifier
	isSet bool
}

func (v NullableAccessNetChargingIdentifier) Get() *AccessNetChargingIdentifier {
	return v.value
}

func (v *NullableAccessNetChargingIdentifier) Set(val *AccessNetChargingIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessNetChargingIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessNetChargingIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessNetChargingIdentifier(val *AccessNetChargingIdentifier) *NullableAccessNetChargingIdentifier {
	return &NullableAccessNetChargingIdentifier{value: val, isSet: true}
}

func (v NullableAccessNetChargingIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessNetChargingIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
