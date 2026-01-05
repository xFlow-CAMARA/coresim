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

// checks if the StringMatchingCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringMatchingCondition{}

// StringMatchingCondition A String with Matching Operator
type StringMatchingCondition struct {
	MatchingString   *string          `json:"matchingString,omitempty"`
	MatchingOperator MatchingOperator `json:"matchingOperator"`
}

type _StringMatchingCondition StringMatchingCondition

// NewStringMatchingCondition instantiates a new StringMatchingCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringMatchingCondition(matchingOperator MatchingOperator) *StringMatchingCondition {
	this := StringMatchingCondition{}
	this.MatchingOperator = matchingOperator
	return &this
}

// NewStringMatchingConditionWithDefaults instantiates a new StringMatchingCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringMatchingConditionWithDefaults() *StringMatchingCondition {
	this := StringMatchingCondition{}
	return &this
}

// GetMatchingString returns the MatchingString field value if set, zero value otherwise.
func (o *StringMatchingCondition) GetMatchingString() string {
	if o == nil || IsNil(o.MatchingString) {
		var ret string
		return ret
	}
	return *o.MatchingString
}

// GetMatchingStringOk returns a tuple with the MatchingString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMatchingCondition) GetMatchingStringOk() (*string, bool) {
	if o == nil || IsNil(o.MatchingString) {
		return nil, false
	}
	return o.MatchingString, true
}

// HasMatchingString returns a boolean if a field has been set.
func (o *StringMatchingCondition) HasMatchingString() bool {
	if o != nil && !IsNil(o.MatchingString) {
		return true
	}

	return false
}

// SetMatchingString gets a reference to the given string and assigns it to the MatchingString field.
func (o *StringMatchingCondition) SetMatchingString(v string) {
	o.MatchingString = &v
}

// GetMatchingOperator returns the MatchingOperator field value
func (o *StringMatchingCondition) GetMatchingOperator() MatchingOperator {
	if o == nil {
		var ret MatchingOperator
		return ret
	}

	return o.MatchingOperator
}

// GetMatchingOperatorOk returns a tuple with the MatchingOperator field value
// and a boolean to check if the value has been set.
func (o *StringMatchingCondition) GetMatchingOperatorOk() (*MatchingOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingOperator, true
}

// SetMatchingOperator sets field value
func (o *StringMatchingCondition) SetMatchingOperator(v MatchingOperator) {
	o.MatchingOperator = v
}

func (o StringMatchingCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringMatchingCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchingString) {
		toSerialize["matchingString"] = o.MatchingString
	}
	toSerialize["matchingOperator"] = o.MatchingOperator
	return toSerialize, nil
}

func (o *StringMatchingCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"matchingOperator",
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

	varStringMatchingCondition := _StringMatchingCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringMatchingCondition)

	if err != nil {
		return err
	}

	*o = StringMatchingCondition(varStringMatchingCondition)

	return err
}

type NullableStringMatchingCondition struct {
	value *StringMatchingCondition
	isSet bool
}

func (v NullableStringMatchingCondition) Get() *StringMatchingCondition {
	return v.value
}

func (v *NullableStringMatchingCondition) Set(val *StringMatchingCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableStringMatchingCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableStringMatchingCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringMatchingCondition(val *StringMatchingCondition) *NullableStringMatchingCondition {
	return &NullableStringMatchingCondition{value: val, isSet: true}
}

func (v NullableStringMatchingCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringMatchingCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
