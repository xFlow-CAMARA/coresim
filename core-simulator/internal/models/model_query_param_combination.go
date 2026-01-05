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

// checks if the QueryParamCombination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryParamCombination{}

// QueryParamCombination Contains a list of Query Parameters
type QueryParamCombination struct {
	QueryParams []QueryParameter `json:"queryParams"`
}

type _QueryParamCombination QueryParamCombination

// NewQueryParamCombination instantiates a new QueryParamCombination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParamCombination(queryParams []QueryParameter) *QueryParamCombination {
	this := QueryParamCombination{}
	this.QueryParams = queryParams
	return &this
}

// NewQueryParamCombinationWithDefaults instantiates a new QueryParamCombination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParamCombinationWithDefaults() *QueryParamCombination {
	this := QueryParamCombination{}
	return &this
}

// GetQueryParams returns the QueryParams field value
func (o *QueryParamCombination) GetQueryParams() []QueryParameter {
	if o == nil {
		var ret []QueryParameter
		return ret
	}

	return o.QueryParams
}

// GetQueryParamsOk returns a tuple with the QueryParams field value
// and a boolean to check if the value has been set.
func (o *QueryParamCombination) GetQueryParamsOk() ([]QueryParameter, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryParams, true
}

// SetQueryParams sets field value
func (o *QueryParamCombination) SetQueryParams(v []QueryParameter) {
	o.QueryParams = v
}

func (o QueryParamCombination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryParamCombination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryParams"] = o.QueryParams
	return toSerialize, nil
}

func (o *QueryParamCombination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryParams",
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

	varQueryParamCombination := _QueryParamCombination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueryParamCombination)

	if err != nil {
		return err
	}

	*o = QueryParamCombination(varQueryParamCombination)

	return err
}

type NullableQueryParamCombination struct {
	value *QueryParamCombination
	isSet bool
}

func (v NullableQueryParamCombination) Get() *QueryParamCombination {
	return v.value
}

func (v *NullableQueryParamCombination) Set(val *QueryParamCombination) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParamCombination) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParamCombination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParamCombination(val *QueryParamCombination) *NullableQueryParamCombination {
	return &NullableQueryParamCombination{value: val, isSet: true}
}

func (v NullableQueryParamCombination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParamCombination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
