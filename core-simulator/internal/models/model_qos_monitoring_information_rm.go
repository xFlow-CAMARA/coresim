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

// checks if the QosMonitoringInformationRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosMonitoringInformationRm{}

// QosMonitoringInformationRm This data type is defined in the same way as the QosMonitoringInformation data type, but with the OpenAPI nullable property set to true.
type QosMonitoringInformationRm struct {
	RepThreshDl *int32 `json:"repThreshDl,omitempty"`
	RepThreshUl *int32 `json:"repThreshUl,omitempty"`
	RepThreshRp *int32 `json:"repThreshRp,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateUl NullableString `json:"repThreshDatRateUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	RepThreshDatRateDl NullableString `json:"repThreshDatRateDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshDl NullableInt32 `json:"conThreshDl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property.
	ConThreshUl NullableInt32 `json:"conThreshUl,omitempty"`
}

// NewQosMonitoringInformationRm instantiates a new QosMonitoringInformationRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringInformationRm() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// NewQosMonitoringInformationRmWithDefaults instantiates a new QosMonitoringInformationRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringInformationRmWithDefaults() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// GetRepThreshDl returns the RepThreshDl field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshDl() int32 {
	if o == nil || IsNil(o.RepThreshDl) {
		var ret int32
		return ret
	}
	return *o.RepThreshDl
}

// GetRepThreshDlOk returns a tuple with the RepThreshDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshDlOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshDl) {
		return nil, false
	}
	return o.RepThreshDl, true
}

// HasRepThreshDl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshDl() bool {
	if o != nil && !IsNil(o.RepThreshDl) {
		return true
	}

	return false
}

// SetRepThreshDl gets a reference to the given int32 and assigns it to the RepThreshDl field.
func (o *QosMonitoringInformationRm) SetRepThreshDl(v int32) {
	o.RepThreshDl = &v
}

// GetRepThreshUl returns the RepThreshUl field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshUl() int32 {
	if o == nil || IsNil(o.RepThreshUl) {
		var ret int32
		return ret
	}
	return *o.RepThreshUl
}

// GetRepThreshUlOk returns a tuple with the RepThreshUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshUlOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshUl) {
		return nil, false
	}
	return o.RepThreshUl, true
}

// HasRepThreshUl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshUl() bool {
	if o != nil && !IsNil(o.RepThreshUl) {
		return true
	}

	return false
}

// SetRepThreshUl gets a reference to the given int32 and assigns it to the RepThreshUl field.
func (o *QosMonitoringInformationRm) SetRepThreshUl(v int32) {
	o.RepThreshUl = &v
}

// GetRepThreshRp returns the RepThreshRp field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshRp() int32 {
	if o == nil || IsNil(o.RepThreshRp) {
		var ret int32
		return ret
	}
	return *o.RepThreshRp
}

// GetRepThreshRpOk returns a tuple with the RepThreshRp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshRpOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshRp) {
		return nil, false
	}
	return o.RepThreshRp, true
}

// HasRepThreshRp returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshRp() bool {
	if o != nil && !IsNil(o.RepThreshRp) {
		return true
	}

	return false
}

// SetRepThreshRp gets a reference to the given int32 and assigns it to the RepThreshRp field.
func (o *QosMonitoringInformationRm) SetRepThreshRp(v int32) {
	o.RepThreshRp = &v
}

// GetRepThreshDatRateUl returns the RepThreshDatRateUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepThreshDatRateUl() string {
	if o == nil || IsNil(o.RepThreshDatRateUl.Get()) {
		var ret string
		return ret
	}
	return *o.RepThreshDatRateUl.Get()
}

// GetRepThreshDatRateUlOk returns a tuple with the RepThreshDatRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepThreshDatRateUlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepThreshDatRateUl.Get(), o.RepThreshDatRateUl.IsSet()
}

// HasRepThreshDatRateUl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshDatRateUl() bool {
	if o != nil && o.RepThreshDatRateUl.IsSet() {
		return true
	}

	return false
}

// SetRepThreshDatRateUl gets a reference to the given NullableString and assigns it to the RepThreshDatRateUl field.
func (o *QosMonitoringInformationRm) SetRepThreshDatRateUl(v string) {
	o.RepThreshDatRateUl.Set(&v)
}

// SetRepThreshDatRateUlNil sets the value for RepThreshDatRateUl to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepThreshDatRateUlNil() {
	o.RepThreshDatRateUl.Set(nil)
}

// UnsetRepThreshDatRateUl ensures that no value is present for RepThreshDatRateUl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepThreshDatRateUl() {
	o.RepThreshDatRateUl.Unset()
}

// GetRepThreshDatRateDl returns the RepThreshDatRateDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepThreshDatRateDl() string {
	if o == nil || IsNil(o.RepThreshDatRateDl.Get()) {
		var ret string
		return ret
	}
	return *o.RepThreshDatRateDl.Get()
}

// GetRepThreshDatRateDlOk returns a tuple with the RepThreshDatRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepThreshDatRateDlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepThreshDatRateDl.Get(), o.RepThreshDatRateDl.IsSet()
}

// HasRepThreshDatRateDl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshDatRateDl() bool {
	if o != nil && o.RepThreshDatRateDl.IsSet() {
		return true
	}

	return false
}

// SetRepThreshDatRateDl gets a reference to the given NullableString and assigns it to the RepThreshDatRateDl field.
func (o *QosMonitoringInformationRm) SetRepThreshDatRateDl(v string) {
	o.RepThreshDatRateDl.Set(&v)
}

// SetRepThreshDatRateDlNil sets the value for RepThreshDatRateDl to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepThreshDatRateDlNil() {
	o.RepThreshDatRateDl.Set(nil)
}

// UnsetRepThreshDatRateDl ensures that no value is present for RepThreshDatRateDl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepThreshDatRateDl() {
	o.RepThreshDatRateDl.Unset()
}

// GetConThreshDl returns the ConThreshDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetConThreshDl() int32 {
	if o == nil || IsNil(o.ConThreshDl.Get()) {
		var ret int32
		return ret
	}
	return *o.ConThreshDl.Get()
}

// GetConThreshDlOk returns a tuple with the ConThreshDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetConThreshDlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConThreshDl.Get(), o.ConThreshDl.IsSet()
}

// HasConThreshDl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasConThreshDl() bool {
	if o != nil && o.ConThreshDl.IsSet() {
		return true
	}

	return false
}

// SetConThreshDl gets a reference to the given NullableInt32 and assigns it to the ConThreshDl field.
func (o *QosMonitoringInformationRm) SetConThreshDl(v int32) {
	o.ConThreshDl.Set(&v)
}

// SetConThreshDlNil sets the value for ConThreshDl to be an explicit nil
func (o *QosMonitoringInformationRm) SetConThreshDlNil() {
	o.ConThreshDl.Set(nil)
}

// UnsetConThreshDl ensures that no value is present for ConThreshDl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetConThreshDl() {
	o.ConThreshDl.Unset()
}

// GetConThreshUl returns the ConThreshUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetConThreshUl() int32 {
	if o == nil || IsNil(o.ConThreshUl.Get()) {
		var ret int32
		return ret
	}
	return *o.ConThreshUl.Get()
}

// GetConThreshUlOk returns a tuple with the ConThreshUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetConThreshUlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConThreshUl.Get(), o.ConThreshUl.IsSet()
}

// HasConThreshUl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasConThreshUl() bool {
	if o != nil && o.ConThreshUl.IsSet() {
		return true
	}

	return false
}

// SetConThreshUl gets a reference to the given NullableInt32 and assigns it to the ConThreshUl field.
func (o *QosMonitoringInformationRm) SetConThreshUl(v int32) {
	o.ConThreshUl.Set(&v)
}

// SetConThreshUlNil sets the value for ConThreshUl to be an explicit nil
func (o *QosMonitoringInformationRm) SetConThreshUlNil() {
	o.ConThreshUl.Set(nil)
}

// UnsetConThreshUl ensures that no value is present for ConThreshUl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetConThreshUl() {
	o.ConThreshUl.Unset()
}

func (o QosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosMonitoringInformationRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepThreshDl) {
		toSerialize["repThreshDl"] = o.RepThreshDl
	}
	if !IsNil(o.RepThreshUl) {
		toSerialize["repThreshUl"] = o.RepThreshUl
	}
	if !IsNil(o.RepThreshRp) {
		toSerialize["repThreshRp"] = o.RepThreshRp
	}
	if o.RepThreshDatRateUl.IsSet() {
		toSerialize["repThreshDatRateUl"] = o.RepThreshDatRateUl.Get()
	}
	if o.RepThreshDatRateDl.IsSet() {
		toSerialize["repThreshDatRateDl"] = o.RepThreshDatRateDl.Get()
	}
	if o.ConThreshDl.IsSet() {
		toSerialize["conThreshDl"] = o.ConThreshDl.Get()
	}
	if o.ConThreshUl.IsSet() {
		toSerialize["conThreshUl"] = o.ConThreshUl.Get()
	}
	return toSerialize, nil
}

type NullableQosMonitoringInformationRm struct {
	value *QosMonitoringInformationRm
	isSet bool
}

func (v NullableQosMonitoringInformationRm) Get() *QosMonitoringInformationRm {
	return v.value
}

func (v *NullableQosMonitoringInformationRm) Set(val *QosMonitoringInformationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringInformationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringInformationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringInformationRm(val *QosMonitoringInformationRm) *NullableQosMonitoringInformationRm {
	return &NullableQosMonitoringInformationRm{value: val, isSet: true}
}

func (v NullableQosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringInformationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
