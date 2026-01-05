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

// checks if the SpatialValidity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpatialValidity{}

// SpatialValidity Describes explicitly the route to an Application location.
type SpatialValidity struct {
	// Defines the presence information provisioned by the AF. The praId attribute within the PresenceInfo data type is the key of the map.
	PresenceInfoList map[string]PresenceInfo `json:"presenceInfoList"`
}

type _SpatialValidity SpatialValidity

// NewSpatialValidity instantiates a new SpatialValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpatialValidity(presenceInfoList map[string]PresenceInfo) *SpatialValidity {
	this := SpatialValidity{}
	this.PresenceInfoList = presenceInfoList
	return &this
}

// NewSpatialValidityWithDefaults instantiates a new SpatialValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpatialValidityWithDefaults() *SpatialValidity {
	this := SpatialValidity{}
	return &this
}

// GetPresenceInfoList returns the PresenceInfoList field value
func (o *SpatialValidity) GetPresenceInfoList() map[string]PresenceInfo {
	if o == nil {
		var ret map[string]PresenceInfo
		return ret
	}

	return o.PresenceInfoList
}

// GetPresenceInfoListOk returns a tuple with the PresenceInfoList field value
// and a boolean to check if the value has been set.
func (o *SpatialValidity) GetPresenceInfoListOk() (*map[string]PresenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresenceInfoList, true
}

// SetPresenceInfoList sets field value
func (o *SpatialValidity) SetPresenceInfoList(v map[string]PresenceInfo) {
	o.PresenceInfoList = v
}

func (o SpatialValidity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpatialValidity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["presenceInfoList"] = o.PresenceInfoList
	return toSerialize, nil
}

func (o *SpatialValidity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"presenceInfoList",
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

	varSpatialValidity := _SpatialValidity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpatialValidity)

	if err != nil {
		return err
	}

	*o = SpatialValidity(varSpatialValidity)

	return err
}

type NullableSpatialValidity struct {
	value *SpatialValidity
	isSet bool
}

func (v NullableSpatialValidity) Get() *SpatialValidity {
	return v.value
}

func (v *NullableSpatialValidity) Set(val *SpatialValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableSpatialValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableSpatialValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpatialValidity(val *SpatialValidity) *NullableSpatialValidity {
	return &NullableSpatialValidity{value: val, isSet: true}
}

func (v NullableSpatialValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpatialValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
