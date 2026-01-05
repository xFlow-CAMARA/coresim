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

// checks if the AppSessionContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSessionContext{}

// AppSessionContext Represents an Individual Application Session Context resource.
type AppSessionContext struct {
	AscReqData  NullableAppSessionContextReqData `json:"ascReqData,omitempty"`
	AscRespData *AppSessionContextRespData       `json:"ascRespData,omitempty"`
	EvsNotif    *EventsNotification              `json:"evsNotif,omitempty"`
}

// NewAppSessionContext instantiates a new AppSessionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSessionContext() *AppSessionContext {
	this := AppSessionContext{}
	return &this
}

// NewAppSessionContextWithDefaults instantiates a new AppSessionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSessionContextWithDefaults() *AppSessionContext {
	this := AppSessionContext{}
	return &this
}

// GetAscReqData returns the AscReqData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContext) GetAscReqData() AppSessionContextReqData {
	if o == nil || IsNil(o.AscReqData.Get()) {
		var ret AppSessionContextReqData
		return ret
	}
	return *o.AscReqData.Get()
}

// GetAscReqDataOk returns a tuple with the AscReqData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContext) GetAscReqDataOk() (*AppSessionContextReqData, bool) {
	if o == nil {
		return nil, false
	}
	return o.AscReqData.Get(), o.AscReqData.IsSet()
}

// HasAscReqData returns a boolean if a field has been set.
func (o *AppSessionContext) HasAscReqData() bool {
	if o != nil && o.AscReqData.IsSet() {
		return true
	}

	return false
}

// SetAscReqData gets a reference to the given NullableAppSessionContextReqData and assigns it to the AscReqData field.
func (o *AppSessionContext) SetAscReqData(v AppSessionContextReqData) {
	o.AscReqData.Set(&v)
}

// SetAscReqDataNil sets the value for AscReqData to be an explicit nil
func (o *AppSessionContext) SetAscReqDataNil() {
	o.AscReqData.Set(nil)
}

// UnsetAscReqData ensures that no value is present for AscReqData, not even an explicit nil
func (o *AppSessionContext) UnsetAscReqData() {
	o.AscReqData.Unset()
}

// GetAscRespData returns the AscRespData field value if set, zero value otherwise.
func (o *AppSessionContext) GetAscRespData() AppSessionContextRespData {
	if o == nil || IsNil(o.AscRespData) {
		var ret AppSessionContextRespData
		return ret
	}
	return *o.AscRespData
}

// GetAscRespDataOk returns a tuple with the AscRespData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContext) GetAscRespDataOk() (*AppSessionContextRespData, bool) {
	if o == nil || IsNil(o.AscRespData) {
		return nil, false
	}
	return o.AscRespData, true
}

// HasAscRespData returns a boolean if a field has been set.
func (o *AppSessionContext) HasAscRespData() bool {
	if o != nil && !IsNil(o.AscRespData) {
		return true
	}

	return false
}

// SetAscRespData gets a reference to the given AppSessionContextRespData and assigns it to the AscRespData field.
func (o *AppSessionContext) SetAscRespData(v AppSessionContextRespData) {
	o.AscRespData = &v
}

// GetEvsNotif returns the EvsNotif field value if set, zero value otherwise.
func (o *AppSessionContext) GetEvsNotif() EventsNotification {
	if o == nil || IsNil(o.EvsNotif) {
		var ret EventsNotification
		return ret
	}
	return *o.EvsNotif
}

// GetEvsNotifOk returns a tuple with the EvsNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContext) GetEvsNotifOk() (*EventsNotification, bool) {
	if o == nil || IsNil(o.EvsNotif) {
		return nil, false
	}
	return o.EvsNotif, true
}

// HasEvsNotif returns a boolean if a field has been set.
func (o *AppSessionContext) HasEvsNotif() bool {
	if o != nil && !IsNil(o.EvsNotif) {
		return true
	}

	return false
}

// SetEvsNotif gets a reference to the given EventsNotification and assigns it to the EvsNotif field.
func (o *AppSessionContext) SetEvsNotif(v EventsNotification) {
	o.EvsNotif = &v
}

func (o AppSessionContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSessionContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AscReqData.IsSet() {
		toSerialize["ascReqData"] = o.AscReqData.Get()
	}
	if !IsNil(o.AscRespData) {
		toSerialize["ascRespData"] = o.AscRespData
	}
	if !IsNil(o.EvsNotif) {
		toSerialize["evsNotif"] = o.EvsNotif
	}
	return toSerialize, nil
}

type NullableAppSessionContext struct {
	value *AppSessionContext
	isSet bool
}

func (v NullableAppSessionContext) Get() *AppSessionContext {
	return v.value
}

func (v *NullableAppSessionContext) Set(val *AppSessionContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContext(val *AppSessionContext) *NullableAppSessionContext {
	return &NullableAppSessionContext{value: val, isSet: true}
}

func (v NullableAppSessionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
