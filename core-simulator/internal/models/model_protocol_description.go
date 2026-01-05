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

// checks if the ProtocolDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolDescription{}

// ProtocolDescription ProtocolDescription contains information to derive PDU set information.
type ProtocolDescription struct {
	TransportProto     *MediaTransportProto `json:"transportProto,omitempty"`
	RtpHeaderExtInfo   *RtpHeaderExtInfo    `json:"rtpHeaderExtInfo,omitempty"`
	RtpPayloadInfoList []RtpPayloadInfo     `json:"rtpPayloadInfoList,omitempty"`
}

// NewProtocolDescription instantiates a new ProtocolDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolDescription() *ProtocolDescription {
	this := ProtocolDescription{}
	return &this
}

// NewProtocolDescriptionWithDefaults instantiates a new ProtocolDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolDescriptionWithDefaults() *ProtocolDescription {
	this := ProtocolDescription{}
	return &this
}

// GetTransportProto returns the TransportProto field value if set, zero value otherwise.
func (o *ProtocolDescription) GetTransportProto() MediaTransportProto {
	if o == nil || IsNil(o.TransportProto) {
		var ret MediaTransportProto
		return ret
	}
	return *o.TransportProto
}

// GetTransportProtoOk returns a tuple with the TransportProto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolDescription) GetTransportProtoOk() (*MediaTransportProto, bool) {
	if o == nil || IsNil(o.TransportProto) {
		return nil, false
	}
	return o.TransportProto, true
}

// HasTransportProto returns a boolean if a field has been set.
func (o *ProtocolDescription) HasTransportProto() bool {
	if o != nil && !IsNil(o.TransportProto) {
		return true
	}

	return false
}

// SetTransportProto gets a reference to the given MediaTransportProto and assigns it to the TransportProto field.
func (o *ProtocolDescription) SetTransportProto(v MediaTransportProto) {
	o.TransportProto = &v
}

// GetRtpHeaderExtInfo returns the RtpHeaderExtInfo field value if set, zero value otherwise.
func (o *ProtocolDescription) GetRtpHeaderExtInfo() RtpHeaderExtInfo {
	if o == nil || IsNil(o.RtpHeaderExtInfo) {
		var ret RtpHeaderExtInfo
		return ret
	}
	return *o.RtpHeaderExtInfo
}

// GetRtpHeaderExtInfoOk returns a tuple with the RtpHeaderExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolDescription) GetRtpHeaderExtInfoOk() (*RtpHeaderExtInfo, bool) {
	if o == nil || IsNil(o.RtpHeaderExtInfo) {
		return nil, false
	}
	return o.RtpHeaderExtInfo, true
}

// HasRtpHeaderExtInfo returns a boolean if a field has been set.
func (o *ProtocolDescription) HasRtpHeaderExtInfo() bool {
	if o != nil && !IsNil(o.RtpHeaderExtInfo) {
		return true
	}

	return false
}

// SetRtpHeaderExtInfo gets a reference to the given RtpHeaderExtInfo and assigns it to the RtpHeaderExtInfo field.
func (o *ProtocolDescription) SetRtpHeaderExtInfo(v RtpHeaderExtInfo) {
	o.RtpHeaderExtInfo = &v
}

// GetRtpPayloadInfoList returns the RtpPayloadInfoList field value if set, zero value otherwise.
func (o *ProtocolDescription) GetRtpPayloadInfoList() []RtpPayloadInfo {
	if o == nil || IsNil(o.RtpPayloadInfoList) {
		var ret []RtpPayloadInfo
		return ret
	}
	return o.RtpPayloadInfoList
}

// GetRtpPayloadInfoListOk returns a tuple with the RtpPayloadInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolDescription) GetRtpPayloadInfoListOk() ([]RtpPayloadInfo, bool) {
	if o == nil || IsNil(o.RtpPayloadInfoList) {
		return nil, false
	}
	return o.RtpPayloadInfoList, true
}

// HasRtpPayloadInfoList returns a boolean if a field has been set.
func (o *ProtocolDescription) HasRtpPayloadInfoList() bool {
	if o != nil && !IsNil(o.RtpPayloadInfoList) {
		return true
	}

	return false
}

// SetRtpPayloadInfoList gets a reference to the given []RtpPayloadInfo and assigns it to the RtpPayloadInfoList field.
func (o *ProtocolDescription) SetRtpPayloadInfoList(v []RtpPayloadInfo) {
	o.RtpPayloadInfoList = v
}

func (o ProtocolDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransportProto) {
		toSerialize["transportProto"] = o.TransportProto
	}
	if !IsNil(o.RtpHeaderExtInfo) {
		toSerialize["rtpHeaderExtInfo"] = o.RtpHeaderExtInfo
	}
	if !IsNil(o.RtpPayloadInfoList) {
		toSerialize["rtpPayloadInfoList"] = o.RtpPayloadInfoList
	}
	return toSerialize, nil
}

type NullableProtocolDescription struct {
	value *ProtocolDescription
	isSet bool
}

func (v NullableProtocolDescription) Get() *ProtocolDescription {
	return v.value
}

func (v *NullableProtocolDescription) Set(val *ProtocolDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolDescription(val *ProtocolDescription) *NullableProtocolDescription {
	return &NullableProtocolDescription{value: val, isSet: true}
}

func (v NullableProtocolDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
