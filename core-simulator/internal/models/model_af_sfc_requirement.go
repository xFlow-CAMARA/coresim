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

// checks if the AfSfcRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfSfcRequirement{}

// AfSfcRequirement Describes AF requirements on steering traffic to N6-LAN.
type AfSfcRequirement struct {
	// Reference to a pre-configured SFC for downlink traffic.
	SfcIdDl NullableString `json:"sfcIdDl,omitempty"`
	// Reference to a pre-configured SFC for uplink traffic.
	SfcIdUl NullableString            `json:"sfcIdUl,omitempty"`
	SpVal   NullableSpatialValidityRm `json:"spVal,omitempty"`
	// A String which is transparently passed to the UPF to be applied for traffic to SFC.
	Metadata NullableString `json:"metadata,omitempty"`
}

// NewAfSfcRequirement instantiates a new AfSfcRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfSfcRequirement() *AfSfcRequirement {
	this := AfSfcRequirement{}
	return &this
}

// NewAfSfcRequirementWithDefaults instantiates a new AfSfcRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfSfcRequirementWithDefaults() *AfSfcRequirement {
	this := AfSfcRequirement{}
	return &this
}

// GetSfcIdDl returns the SfcIdDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AfSfcRequirement) GetSfcIdDl() string {
	if o == nil || IsNil(o.SfcIdDl.Get()) {
		var ret string
		return ret
	}
	return *o.SfcIdDl.Get()
}

// GetSfcIdDlOk returns a tuple with the SfcIdDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AfSfcRequirement) GetSfcIdDlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SfcIdDl.Get(), o.SfcIdDl.IsSet()
}

// HasSfcIdDl returns a boolean if a field has been set.
func (o *AfSfcRequirement) HasSfcIdDl() bool {
	if o != nil && o.SfcIdDl.IsSet() {
		return true
	}

	return false
}

// SetSfcIdDl gets a reference to the given NullableString and assigns it to the SfcIdDl field.
func (o *AfSfcRequirement) SetSfcIdDl(v string) {
	o.SfcIdDl.Set(&v)
}

// SetSfcIdDlNil sets the value for SfcIdDl to be an explicit nil
func (o *AfSfcRequirement) SetSfcIdDlNil() {
	o.SfcIdDl.Set(nil)
}

// UnsetSfcIdDl ensures that no value is present for SfcIdDl, not even an explicit nil
func (o *AfSfcRequirement) UnsetSfcIdDl() {
	o.SfcIdDl.Unset()
}

// GetSfcIdUl returns the SfcIdUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AfSfcRequirement) GetSfcIdUl() string {
	if o == nil || IsNil(o.SfcIdUl.Get()) {
		var ret string
		return ret
	}
	return *o.SfcIdUl.Get()
}

// GetSfcIdUlOk returns a tuple with the SfcIdUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AfSfcRequirement) GetSfcIdUlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SfcIdUl.Get(), o.SfcIdUl.IsSet()
}

// HasSfcIdUl returns a boolean if a field has been set.
func (o *AfSfcRequirement) HasSfcIdUl() bool {
	if o != nil && o.SfcIdUl.IsSet() {
		return true
	}

	return false
}

// SetSfcIdUl gets a reference to the given NullableString and assigns it to the SfcIdUl field.
func (o *AfSfcRequirement) SetSfcIdUl(v string) {
	o.SfcIdUl.Set(&v)
}

// SetSfcIdUlNil sets the value for SfcIdUl to be an explicit nil
func (o *AfSfcRequirement) SetSfcIdUlNil() {
	o.SfcIdUl.Set(nil)
}

// UnsetSfcIdUl ensures that no value is present for SfcIdUl, not even an explicit nil
func (o *AfSfcRequirement) UnsetSfcIdUl() {
	o.SfcIdUl.Unset()
}

// GetSpVal returns the SpVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AfSfcRequirement) GetSpVal() SpatialValidityRm {
	if o == nil || IsNil(o.SpVal.Get()) {
		var ret SpatialValidityRm
		return ret
	}
	return *o.SpVal.Get()
}

// GetSpValOk returns a tuple with the SpVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AfSfcRequirement) GetSpValOk() (*SpatialValidityRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpVal.Get(), o.SpVal.IsSet()
}

// HasSpVal returns a boolean if a field has been set.
func (o *AfSfcRequirement) HasSpVal() bool {
	if o != nil && o.SpVal.IsSet() {
		return true
	}

	return false
}

// SetSpVal gets a reference to the given NullableSpatialValidityRm and assigns it to the SpVal field.
func (o *AfSfcRequirement) SetSpVal(v SpatialValidityRm) {
	o.SpVal.Set(&v)
}

// SetSpValNil sets the value for SpVal to be an explicit nil
func (o *AfSfcRequirement) SetSpValNil() {
	o.SpVal.Set(nil)
}

// UnsetSpVal ensures that no value is present for SpVal, not even an explicit nil
func (o *AfSfcRequirement) UnsetSpVal() {
	o.SpVal.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AfSfcRequirement) GetMetadata() string {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret string
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AfSfcRequirement) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *AfSfcRequirement) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableString and assigns it to the Metadata field.
func (o *AfSfcRequirement) SetMetadata(v string) {
	o.Metadata.Set(&v)
}

// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *AfSfcRequirement) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *AfSfcRequirement) UnsetMetadata() {
	o.Metadata.Unset()
}

func (o AfSfcRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfSfcRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SfcIdDl.IsSet() {
		toSerialize["sfcIdDl"] = o.SfcIdDl.Get()
	}
	if o.SfcIdUl.IsSet() {
		toSerialize["sfcIdUl"] = o.SfcIdUl.Get()
	}
	if o.SpVal.IsSet() {
		toSerialize["spVal"] = o.SpVal.Get()
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	return toSerialize, nil
}

type NullableAfSfcRequirement struct {
	value *AfSfcRequirement
	isSet bool
}

func (v NullableAfSfcRequirement) Get() *AfSfcRequirement {
	return v.value
}

func (v *NullableAfSfcRequirement) Set(val *AfSfcRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableAfSfcRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableAfSfcRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfSfcRequirement(val *AfSfcRequirement) *NullableAfSfcRequirement {
	return &NullableAfSfcRequirement{value: val, isSet: true}
}

func (v NullableAfSfcRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfSfcRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
