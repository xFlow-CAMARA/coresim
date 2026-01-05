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

// checks if the RtpHeaderExtInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RtpHeaderExtInfo{}

// RtpHeaderExtInfo RTP Header Extension information
type RtpHeaderExtInfo struct {
	RtpHeaderExtType *RtpHeaderExtType `json:"rtpHeaderExtType,omitempty"`
	RtpHeaderExtId   *int32            `json:"rtpHeaderExtId,omitempty"`
	LongFormat       *bool             `json:"longFormat,omitempty"`
	PduSetSizeActive *bool             `json:"pduSetSizeActive,omitempty"`
}

// NewRtpHeaderExtInfo instantiates a new RtpHeaderExtInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRtpHeaderExtInfo() *RtpHeaderExtInfo {
	this := RtpHeaderExtInfo{}
	return &this
}

// NewRtpHeaderExtInfoWithDefaults instantiates a new RtpHeaderExtInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRtpHeaderExtInfoWithDefaults() *RtpHeaderExtInfo {
	this := RtpHeaderExtInfo{}
	return &this
}

// GetRtpHeaderExtType returns the RtpHeaderExtType field value if set, zero value otherwise.
func (o *RtpHeaderExtInfo) GetRtpHeaderExtType() RtpHeaderExtType {
	if o == nil || IsNil(o.RtpHeaderExtType) {
		var ret RtpHeaderExtType
		return ret
	}
	return *o.RtpHeaderExtType
}

// GetRtpHeaderExtTypeOk returns a tuple with the RtpHeaderExtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpHeaderExtInfo) GetRtpHeaderExtTypeOk() (*RtpHeaderExtType, bool) {
	if o == nil || IsNil(o.RtpHeaderExtType) {
		return nil, false
	}
	return o.RtpHeaderExtType, true
}

// HasRtpHeaderExtType returns a boolean if a field has been set.
func (o *RtpHeaderExtInfo) HasRtpHeaderExtType() bool {
	if o != nil && !IsNil(o.RtpHeaderExtType) {
		return true
	}

	return false
}

// SetRtpHeaderExtType gets a reference to the given RtpHeaderExtType and assigns it to the RtpHeaderExtType field.
func (o *RtpHeaderExtInfo) SetRtpHeaderExtType(v RtpHeaderExtType) {
	o.RtpHeaderExtType = &v
}

// GetRtpHeaderExtId returns the RtpHeaderExtId field value if set, zero value otherwise.
func (o *RtpHeaderExtInfo) GetRtpHeaderExtId() int32 {
	if o == nil || IsNil(o.RtpHeaderExtId) {
		var ret int32
		return ret
	}
	return *o.RtpHeaderExtId
}

// GetRtpHeaderExtIdOk returns a tuple with the RtpHeaderExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpHeaderExtInfo) GetRtpHeaderExtIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RtpHeaderExtId) {
		return nil, false
	}
	return o.RtpHeaderExtId, true
}

// HasRtpHeaderExtId returns a boolean if a field has been set.
func (o *RtpHeaderExtInfo) HasRtpHeaderExtId() bool {
	if o != nil && !IsNil(o.RtpHeaderExtId) {
		return true
	}

	return false
}

// SetRtpHeaderExtId gets a reference to the given int32 and assigns it to the RtpHeaderExtId field.
func (o *RtpHeaderExtInfo) SetRtpHeaderExtId(v int32) {
	o.RtpHeaderExtId = &v
}

// GetLongFormat returns the LongFormat field value if set, zero value otherwise.
func (o *RtpHeaderExtInfo) GetLongFormat() bool {
	if o == nil || IsNil(o.LongFormat) {
		var ret bool
		return ret
	}
	return *o.LongFormat
}

// GetLongFormatOk returns a tuple with the LongFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpHeaderExtInfo) GetLongFormatOk() (*bool, bool) {
	if o == nil || IsNil(o.LongFormat) {
		return nil, false
	}
	return o.LongFormat, true
}

// HasLongFormat returns a boolean if a field has been set.
func (o *RtpHeaderExtInfo) HasLongFormat() bool {
	if o != nil && !IsNil(o.LongFormat) {
		return true
	}

	return false
}

// SetLongFormat gets a reference to the given bool and assigns it to the LongFormat field.
func (o *RtpHeaderExtInfo) SetLongFormat(v bool) {
	o.LongFormat = &v
}

// GetPduSetSizeActive returns the PduSetSizeActive field value if set, zero value otherwise.
func (o *RtpHeaderExtInfo) GetPduSetSizeActive() bool {
	if o == nil || IsNil(o.PduSetSizeActive) {
		var ret bool
		return ret
	}
	return *o.PduSetSizeActive
}

// GetPduSetSizeActiveOk returns a tuple with the PduSetSizeActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RtpHeaderExtInfo) GetPduSetSizeActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.PduSetSizeActive) {
		return nil, false
	}
	return o.PduSetSizeActive, true
}

// HasPduSetSizeActive returns a boolean if a field has been set.
func (o *RtpHeaderExtInfo) HasPduSetSizeActive() bool {
	if o != nil && !IsNil(o.PduSetSizeActive) {
		return true
	}

	return false
}

// SetPduSetSizeActive gets a reference to the given bool and assigns it to the PduSetSizeActive field.
func (o *RtpHeaderExtInfo) SetPduSetSizeActive(v bool) {
	o.PduSetSizeActive = &v
}

func (o RtpHeaderExtInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RtpHeaderExtInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RtpHeaderExtType) {
		toSerialize["rtpHeaderExtType"] = o.RtpHeaderExtType
	}
	if !IsNil(o.RtpHeaderExtId) {
		toSerialize["rtpHeaderExtId"] = o.RtpHeaderExtId
	}
	if !IsNil(o.LongFormat) {
		toSerialize["longFormat"] = o.LongFormat
	}
	if !IsNil(o.PduSetSizeActive) {
		toSerialize["pduSetSizeActive"] = o.PduSetSizeActive
	}
	return toSerialize, nil
}

type NullableRtpHeaderExtInfo struct {
	value *RtpHeaderExtInfo
	isSet bool
}

func (v NullableRtpHeaderExtInfo) Get() *RtpHeaderExtInfo {
	return v.value
}

func (v *NullableRtpHeaderExtInfo) Set(val *RtpHeaderExtInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRtpHeaderExtInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRtpHeaderExtInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtpHeaderExtInfo(val *RtpHeaderExtInfo) *NullableRtpHeaderExtInfo {
	return &NullableRtpHeaderExtInfo{value: val, isSet: true}
}

func (v NullableRtpHeaderExtInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtpHeaderExtInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
