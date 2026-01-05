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

// checks if the AddFlowDescriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFlowDescriptionInfo{}

// AddFlowDescriptionInfo Contains additional flow description information.
type AddFlowDescriptionInfo struct {
	// 4-octet string representing the security parameter index of the IPSec packet in hexadecimal representation.
	Spi *string `json:"spi,omitempty"`
	// 3-octet string representing the IPv6 flow label header field in hexadecimal representation.
	FlowLabel *string        `json:"flowLabel,omitempty"`
	FlowDir   *FlowDirection `json:"flowDir,omitempty"`
}

// NewAddFlowDescriptionInfo instantiates a new AddFlowDescriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFlowDescriptionInfo() *AddFlowDescriptionInfo {
	this := AddFlowDescriptionInfo{}
	return &this
}

// NewAddFlowDescriptionInfoWithDefaults instantiates a new AddFlowDescriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFlowDescriptionInfoWithDefaults() *AddFlowDescriptionInfo {
	this := AddFlowDescriptionInfo{}
	return &this
}

// GetSpi returns the Spi field value if set, zero value otherwise.
func (o *AddFlowDescriptionInfo) GetSpi() string {
	if o == nil || IsNil(o.Spi) {
		var ret string
		return ret
	}
	return *o.Spi
}

// GetSpiOk returns a tuple with the Spi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFlowDescriptionInfo) GetSpiOk() (*string, bool) {
	if o == nil || IsNil(o.Spi) {
		return nil, false
	}
	return o.Spi, true
}

// HasSpi returns a boolean if a field has been set.
func (o *AddFlowDescriptionInfo) HasSpi() bool {
	if o != nil && !IsNil(o.Spi) {
		return true
	}

	return false
}

// SetSpi gets a reference to the given string and assigns it to the Spi field.
func (o *AddFlowDescriptionInfo) SetSpi(v string) {
	o.Spi = &v
}

// GetFlowLabel returns the FlowLabel field value if set, zero value otherwise.
func (o *AddFlowDescriptionInfo) GetFlowLabel() string {
	if o == nil || IsNil(o.FlowLabel) {
		var ret string
		return ret
	}
	return *o.FlowLabel
}

// GetFlowLabelOk returns a tuple with the FlowLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFlowDescriptionInfo) GetFlowLabelOk() (*string, bool) {
	if o == nil || IsNil(o.FlowLabel) {
		return nil, false
	}
	return o.FlowLabel, true
}

// HasFlowLabel returns a boolean if a field has been set.
func (o *AddFlowDescriptionInfo) HasFlowLabel() bool {
	if o != nil && !IsNil(o.FlowLabel) {
		return true
	}

	return false
}

// SetFlowLabel gets a reference to the given string and assigns it to the FlowLabel field.
func (o *AddFlowDescriptionInfo) SetFlowLabel(v string) {
	o.FlowLabel = &v
}

// GetFlowDir returns the FlowDir field value if set, zero value otherwise.
func (o *AddFlowDescriptionInfo) GetFlowDir() FlowDirection {
	if o == nil || IsNil(o.FlowDir) {
		var ret FlowDirection
		return ret
	}
	return *o.FlowDir
}

// GetFlowDirOk returns a tuple with the FlowDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFlowDescriptionInfo) GetFlowDirOk() (*FlowDirection, bool) {
	if o == nil || IsNil(o.FlowDir) {
		return nil, false
	}
	return o.FlowDir, true
}

// HasFlowDir returns a boolean if a field has been set.
func (o *AddFlowDescriptionInfo) HasFlowDir() bool {
	if o != nil && !IsNil(o.FlowDir) {
		return true
	}

	return false
}

// SetFlowDir gets a reference to the given FlowDirection and assigns it to the FlowDir field.
func (o *AddFlowDescriptionInfo) SetFlowDir(v FlowDirection) {
	o.FlowDir = &v
}

func (o AddFlowDescriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFlowDescriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Spi) {
		toSerialize["spi"] = o.Spi
	}
	if !IsNil(o.FlowLabel) {
		toSerialize["flowLabel"] = o.FlowLabel
	}
	if !IsNil(o.FlowDir) {
		toSerialize["flowDir"] = o.FlowDir
	}
	return toSerialize, nil
}

type NullableAddFlowDescriptionInfo struct {
	value *AddFlowDescriptionInfo
	isSet bool
}

func (v NullableAddFlowDescriptionInfo) Get() *AddFlowDescriptionInfo {
	return v.value
}

func (v *NullableAddFlowDescriptionInfo) Set(val *AddFlowDescriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFlowDescriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFlowDescriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFlowDescriptionInfo(val *AddFlowDescriptionInfo) *NullableAddFlowDescriptionInfo {
	return &NullableAddFlowDescriptionInfo{value: val, isSet: true}
}

func (v NullableAddFlowDescriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFlowDescriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
