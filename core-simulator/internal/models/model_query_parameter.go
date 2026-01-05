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

// checks if the QueryParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryParameter{}

// QueryParameter Contains the name and value of a query parameter
type QueryParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type _QueryParameter QueryParameter

// NewQueryParameter instantiates a new QueryParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParameter(name string, value string) *QueryParameter {
	this := QueryParameter{}
	this.Name = name
	this.Value = value
	return &this
}

// NewQueryParameterWithDefaults instantiates a new QueryParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParameterWithDefaults() *QueryParameter {
	this := QueryParameter{}
	return &this
}

// GetName returns the Name field value
func (o *QueryParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QueryParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *QueryParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *QueryParameter) SetValue(v string) {
	o.Value = v
}

func (o QueryParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *QueryParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varQueryParameter := _QueryParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryParameter)

	if err != nil {
		return err
	}

	*o = QueryParameter(varQueryParameter)

	return err
}

type NullableQueryParameter struct {
	value *QueryParameter
	isSet bool
}

func (v NullableQueryParameter) Get() *QueryParameter {
	return v.value
}

func (v *NullableQueryParameter) Set(val *QueryParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParameter(val *QueryParameter) *NullableQueryParameter {
	return &NullableQueryParameter{value: val, isSet: true}
}

func (v NullableQueryParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
