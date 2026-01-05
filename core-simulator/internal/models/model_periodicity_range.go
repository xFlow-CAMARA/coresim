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

// checks if the PeriodicityRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeriodicityRange{}

// PeriodicityRange Contains the acceptable range (which is formulated as lower bound and upper bound of the periodicity of the start twobursts in reference to the external GM) or acceptable periodicity value(s) (which is formulated as a list of values for the periodicity).
type PeriodicityRange struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	LowerBound *int32
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	UpperBound   *int32
	PeriodicVals []int32
}

// NewPeriodicityRange instantiates a new PeriodicityRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodicityRange() *PeriodicityRange {
	this := PeriodicityRange{}
	return &this
}

// NewPeriodicityRangeWithDefaults instantiates a new PeriodicityRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodicityRangeWithDefaults() *PeriodicityRange {
	this := PeriodicityRange{}
	return &this
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *PeriodicityRange) GetLowerBound() int32 {
	if o == nil || IsNil(o.LowerBound) {
		var ret int32
		return ret
	}
	return *o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicityRange) GetLowerBoundOk() (*int32, bool) {
	if o == nil || IsNil(o.LowerBound) {
		return nil, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *PeriodicityRange) HasLowerBound() bool {
	if o != nil && !IsNil(o.LowerBound) {
		return true
	}

	return false
}

// SetLowerBound gets a reference to the given int32 and assigns it to the LowerBound field.
func (o *PeriodicityRange) SetLowerBound(v int32) {
	o.LowerBound = &v
}

// GetUpperBound returns the UpperBound field value if set, zero value otherwise.
func (o *PeriodicityRange) GetUpperBound() int32 {
	if o == nil || IsNil(o.UpperBound) {
		var ret int32
		return ret
	}
	return *o.UpperBound
}

// GetUpperBoundOk returns a tuple with the UpperBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicityRange) GetUpperBoundOk() (*int32, bool) {
	if o == nil || IsNil(o.UpperBound) {
		return nil, false
	}
	return o.UpperBound, true
}

// HasUpperBound returns a boolean if a field has been set.
func (o *PeriodicityRange) HasUpperBound() bool {
	if o != nil && !IsNil(o.UpperBound) {
		return true
	}

	return false
}

// SetUpperBound gets a reference to the given int32 and assigns it to the UpperBound field.
func (o *PeriodicityRange) SetUpperBound(v int32) {
	o.UpperBound = &v
}

// GetPeriodicVals returns the PeriodicVals field value if set, zero value otherwise.
func (o *PeriodicityRange) GetPeriodicVals() []int32 {
	if o == nil || IsNil(o.PeriodicVals) {
		var ret []int32
		return ret
	}
	return o.PeriodicVals
}

// GetPeriodicValsOk returns a tuple with the PeriodicVals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicityRange) GetPeriodicValsOk() ([]int32, bool) {
	if o == nil || IsNil(o.PeriodicVals) {
		return nil, false
	}
	return o.PeriodicVals, true
}

// HasPeriodicVals returns a boolean if a field has been set.
func (o *PeriodicityRange) HasPeriodicVals() bool {
	if o != nil && !IsNil(o.PeriodicVals) {
		return true
	}

	return false
}

// SetPeriodicVals gets a reference to the given []int32 and assigns it to the PeriodicVals field.
func (o *PeriodicityRange) SetPeriodicVals(v []int32) {
	o.PeriodicVals = v
}

func (o PeriodicityRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeriodicityRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LowerBound) {
		toSerialize["lowerBound"] = o.LowerBound
	}
	if !IsNil(o.UpperBound) {
		toSerialize["upperBound"] = o.UpperBound
	}
	if !IsNil(o.PeriodicVals) {
		toSerialize["periodicVals"] = o.PeriodicVals
	}
	return toSerialize, nil
}

type NullablePeriodicityRange struct {
	value *PeriodicityRange
	isSet bool
}

func (v NullablePeriodicityRange) Get() *PeriodicityRange {
	return v.value
}

func (v *NullablePeriodicityRange) Set(val *PeriodicityRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicityRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicityRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicityRange(val *PeriodicityRange) *NullablePeriodicityRange {
	return &NullablePeriodicityRange{value: val, isSet: true}
}

func (v NullablePeriodicityRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicityRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
