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

package models

import (
	"encoding/json"
)

// UserLocation At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be present
type UserLocation struct {
	NrLocation *NrLocation `json:"nrLocation,omitempty"`
}

// NewUserLocation instantiates a new UserLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLocation() *UserLocation {
	this := UserLocation{}
	return &this
}

// NewUserLocationWithDefaults instantiates a new UserLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLocationWithDefaults() *UserLocation {
	this := UserLocation{}
	return &this
}

// GetNrLocation returns the NrLocation field value if set, zero value otherwise.
func (o *UserLocation) GetNrLocation() NrLocation {
	if o == nil || o.NrLocation == nil {
		var ret NrLocation
		return ret
	}
	return *o.NrLocation
}

// GetNrLocationOk returns a tuple with the NrLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetNrLocationOk() (*NrLocation, bool) {
	if o == nil || o.NrLocation == nil {
		return nil, false
	}
	return o.NrLocation, true
}

// HasNrLocation returns a boolean if a field has been set.
func (o *UserLocation) HasNrLocation() bool {
	if o != nil && o.NrLocation != nil {
		return true
	}

	return false
}

// SetNrLocation gets a reference to the given NrLocation and assigns it to the NrLocation field.
func (o *UserLocation) SetNrLocation(v NrLocation) {
	o.NrLocation = &v
}

type NullableUserLocation struct {
	value *UserLocation
	isSet bool
}

func (v NullableUserLocation) Get() *UserLocation {
	return v.value
}

func (v *NullableUserLocation) Set(val *UserLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLocation(val *UserLocation) *NullableUserLocation {
	return &NullableUserLocation{value: val, isSet: true}
}

func (v NullableUserLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
