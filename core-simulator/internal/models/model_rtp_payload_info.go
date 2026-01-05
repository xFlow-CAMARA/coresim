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

// checks if the RtpPayloadInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RtpPayloadInfo{}

// RtpPayloadInfo RtpPayloadInfo contains Rtp payload type and format.
type RtpPayloadInfo struct {
	RtpPayloadTypeList []int32           `json:"rtpPayloadTypeList,omitempty"`
	RtpPayloadFormat   *RtpPayloadFormat `json:"rtpPayloadFormat,omitempty"`
}

// NewRtpPayloadInfo instantiates a new RtpPayloadInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRtpPayloadInfo() *RtpPayloadInfo {
	this := RtpPayloadInfo{}
	return &this
}

// NewRtpPayloadInfoWithDefaults instantiates a new RtpPayloadInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRtpPayloadInfoWithDefaults() *RtpPayloadInfo {
	this := RtpPayloadInfo{}
	return &this
}

// GetRtpPayloadTypeList returns the RtpPayloadTypeList field value if set, zero value otherwise.
func (o *RtpPayloadInfo) GetRtpPayloadTypeList() []int32 {
	if o == nil || IsNil(o.RtpPayloadTypeList) {
		var ret []int32
		return ret
	}
	return o.RtpPayloadTypeList
}

// GetRtpPayloadTypeListOk returns a tuple with the RtpPayloadTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpPayloadInfo) GetRtpPayloadTypeListOk() ([]int32, bool) {
	if o == nil || IsNil(o.RtpPayloadTypeList) {
		return nil, false
	}
	return o.RtpPayloadTypeList, true
}

// HasRtpPayloadTypeList returns a boolean if a field has been set.
func (o *RtpPayloadInfo) HasRtpPayloadTypeList() bool {
	if o != nil && !IsNil(o.RtpPayloadTypeList) {
		return true
	}

	return false
}

// SetRtpPayloadTypeList gets a reference to the given []int32 and assigns it to the RtpPayloadTypeList field.
func (o *RtpPayloadInfo) SetRtpPayloadTypeList(v []int32) {
	o.RtpPayloadTypeList = v
}

// GetRtpPayloadFormat returns the RtpPayloadFormat field value if set, zero value otherwise.
func (o *RtpPayloadInfo) GetRtpPayloadFormat() RtpPayloadFormat {
	if o == nil || IsNil(o.RtpPayloadFormat) {
		var ret RtpPayloadFormat
		return ret
	}
	return *o.RtpPayloadFormat
}

// GetRtpPayloadFormatOk returns a tuple with the RtpPayloadFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpPayloadInfo) GetRtpPayloadFormatOk() (*RtpPayloadFormat, bool) {
	if o == nil || IsNil(o.RtpPayloadFormat) {
		return nil, false
	}
	return o.RtpPayloadFormat, true
}

// HasRtpPayloadFormat returns a boolean if a field has been set.
func (o *RtpPayloadInfo) HasRtpPayloadFormat() bool {
	if o != nil && !IsNil(o.RtpPayloadFormat) {
		return true
	}

	return false
}

// SetRtpPayloadFormat gets a reference to the given RtpPayloadFormat and assigns it to the RtpPayloadFormat field.
func (o *RtpPayloadInfo) SetRtpPayloadFormat(v RtpPayloadFormat) {
	o.RtpPayloadFormat = &v
}

func (o RtpPayloadInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RtpPayloadInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RtpPayloadTypeList) {
		toSerialize["rtpPayloadTypeList"] = o.RtpPayloadTypeList
	}
	if !IsNil(o.RtpPayloadFormat) {
		toSerialize["rtpPayloadFormat"] = o.RtpPayloadFormat
	}
	return toSerialize, nil
}

type NullableRtpPayloadInfo struct {
	value *RtpPayloadInfo
	isSet bool
}

func (v NullableRtpPayloadInfo) Get() *RtpPayloadInfo {
	return v.value
}

func (v *NullableRtpPayloadInfo) Set(val *RtpPayloadInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRtpPayloadInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRtpPayloadInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtpPayloadInfo(val *RtpPayloadInfo) *NullableRtpPayloadInfo {
	return &NullableRtpPayloadInfo{value: val, isSet: true}
}

func (v NullableRtpPayloadInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtpPayloadInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
