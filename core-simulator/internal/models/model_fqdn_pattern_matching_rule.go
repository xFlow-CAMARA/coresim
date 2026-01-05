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

// checks if the FqdnPatternMatchingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FqdnPatternMatchingRule{}

// FqdnPatternMatchingRule a matching rule for a FQDN pattern
type FqdnPatternMatchingRule struct {
	Regex              *string
	StringMatchingRule *StringMatchingRule
}

// NewFqdnPatternMatchingRule instantiates a new FqdnPatternMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFqdnPatternMatchingRule() *FqdnPatternMatchingRule {
	this := FqdnPatternMatchingRule{}
	return &this
}

// NewFqdnPatternMatchingRuleWithDefaults instantiates a new FqdnPatternMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFqdnPatternMatchingRuleWithDefaults() *FqdnPatternMatchingRule {
	this := FqdnPatternMatchingRule{}
	return &this
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *FqdnPatternMatchingRule) GetRegex() string {
	if o == nil || IsNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FqdnPatternMatchingRule) GetRegexOk() (*string, bool) {
	if o == nil || IsNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *FqdnPatternMatchingRule) HasRegex() bool {
	if o != nil && !IsNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *FqdnPatternMatchingRule) SetRegex(v string) {
	o.Regex = &v
}

// GetStringMatchingRule returns the StringMatchingRule field value if set, zero value otherwise.
func (o *FqdnPatternMatchingRule) GetStringMatchingRule() StringMatchingRule {
	if o == nil || IsNil(o.StringMatchingRule) {
		var ret StringMatchingRule
		return ret
	}
	return *o.StringMatchingRule
}

// GetStringMatchingRuleOk returns a tuple with the StringMatchingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FqdnPatternMatchingRule) GetStringMatchingRuleOk() (*StringMatchingRule, bool) {
	if o == nil || IsNil(o.StringMatchingRule) {
		return nil, false
	}
	return o.StringMatchingRule, true
}

// HasStringMatchingRule returns a boolean if a field has been set.
func (o *FqdnPatternMatchingRule) HasStringMatchingRule() bool {
	if o != nil && !IsNil(o.StringMatchingRule) {
		return true
	}

	return false
}

// SetStringMatchingRule gets a reference to the given StringMatchingRule and assigns it to the StringMatchingRule field.
func (o *FqdnPatternMatchingRule) SetStringMatchingRule(v StringMatchingRule) {
	o.StringMatchingRule = &v
}

func (o FqdnPatternMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FqdnPatternMatchingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	if !IsNil(o.StringMatchingRule) {
		toSerialize["stringMatchingRule"] = o.StringMatchingRule
	}
	return toSerialize, nil
}

type NullableFqdnPatternMatchingRule struct {
	value *FqdnPatternMatchingRule
	isSet bool
}

func (v NullableFqdnPatternMatchingRule) Get() *FqdnPatternMatchingRule {
	return v.value
}

func (v *NullableFqdnPatternMatchingRule) Set(val *FqdnPatternMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFqdnPatternMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFqdnPatternMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFqdnPatternMatchingRule(val *FqdnPatternMatchingRule) *NullableFqdnPatternMatchingRule {
	return &NullableFqdnPatternMatchingRule{value: val, isSet: true}
}

func (v NullableFqdnPatternMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFqdnPatternMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
