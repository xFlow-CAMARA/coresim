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

// checks if the TsnBridgeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TsnBridgeInfo{}

// TsnBridgeInfo Contains parameters that describe and identify the TSC user plane node.
type TsnBridgeInfo struct {
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	BridgeId *int32 `json:"bridgeId,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	DsttAddr *string `json:"dsttAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttPortNum *int32 `json:"dsttPortNum,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttResidTime *int32 `json:"dsttResidTime,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	MtuIpv4 *int32 `json:"mtuIpv4,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	MtuIpv6 *int32 `json:"mtuIpv6,omitempty"`
}

// NewTsnBridgeInfo instantiates a new TsnBridgeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTsnBridgeInfo() *TsnBridgeInfo {
	this := TsnBridgeInfo{}
	return &this
}

// NewTsnBridgeInfoWithDefaults instantiates a new TsnBridgeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTsnBridgeInfoWithDefaults() *TsnBridgeInfo {
	this := TsnBridgeInfo{}
	return &this
}

// GetBridgeId returns the BridgeId field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetBridgeId() int32 {
	if o == nil || IsNil(o.BridgeId) {
		var ret int32
		return ret
	}
	return *o.BridgeId
}

// GetBridgeIdOk returns a tuple with the BridgeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetBridgeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BridgeId) {
		return nil, false
	}
	return o.BridgeId, true
}

// HasBridgeId returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasBridgeId() bool {
	if o != nil && !IsNil(o.BridgeId) {
		return true
	}

	return false
}

// SetBridgeId gets a reference to the given int32 and assigns it to the BridgeId field.
func (o *TsnBridgeInfo) SetBridgeId(v int32) {
	o.BridgeId = &v
}

// GetDsttAddr returns the DsttAddr field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttAddr() string {
	if o == nil || IsNil(o.DsttAddr) {
		var ret string
		return ret
	}
	return *o.DsttAddr
}

// GetDsttAddrOk returns a tuple with the DsttAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttAddrOk() (*string, bool) {
	if o == nil || IsNil(o.DsttAddr) {
		return nil, false
	}
	return o.DsttAddr, true
}

// HasDsttAddr returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttAddr() bool {
	if o != nil && !IsNil(o.DsttAddr) {
		return true
	}

	return false
}

// SetDsttAddr gets a reference to the given string and assigns it to the DsttAddr field.
func (o *TsnBridgeInfo) SetDsttAddr(v string) {
	o.DsttAddr = &v
}

// GetDsttPortNum returns the DsttPortNum field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttPortNum() int32 {
	if o == nil || IsNil(o.DsttPortNum) {
		var ret int32
		return ret
	}
	return *o.DsttPortNum
}

// GetDsttPortNumOk returns a tuple with the DsttPortNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttPortNumOk() (*int32, bool) {
	if o == nil || IsNil(o.DsttPortNum) {
		return nil, false
	}
	return o.DsttPortNum, true
}

// HasDsttPortNum returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttPortNum() bool {
	if o != nil && !IsNil(o.DsttPortNum) {
		return true
	}

	return false
}

// SetDsttPortNum gets a reference to the given int32 and assigns it to the DsttPortNum field.
func (o *TsnBridgeInfo) SetDsttPortNum(v int32) {
	o.DsttPortNum = &v
}

// GetDsttResidTime returns the DsttResidTime field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttResidTime() int32 {
	if o == nil || IsNil(o.DsttResidTime) {
		var ret int32
		return ret
	}
	return *o.DsttResidTime
}

// GetDsttResidTimeOk returns a tuple with the DsttResidTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttResidTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.DsttResidTime) {
		return nil, false
	}
	return o.DsttResidTime, true
}

// HasDsttResidTime returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttResidTime() bool {
	if o != nil && !IsNil(o.DsttResidTime) {
		return true
	}

	return false
}

// SetDsttResidTime gets a reference to the given int32 and assigns it to the DsttResidTime field.
func (o *TsnBridgeInfo) SetDsttResidTime(v int32) {
	o.DsttResidTime = &v
}

// GetMtuIpv4 returns the MtuIpv4 field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetMtuIpv4() int32 {
	if o == nil || IsNil(o.MtuIpv4) {
		var ret int32
		return ret
	}
	return *o.MtuIpv4
}

// GetMtuIpv4Ok returns a tuple with the MtuIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetMtuIpv4Ok() (*int32, bool) {
	if o == nil || IsNil(o.MtuIpv4) {
		return nil, false
	}
	return o.MtuIpv4, true
}

// HasMtuIpv4 returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasMtuIpv4() bool {
	if o != nil && !IsNil(o.MtuIpv4) {
		return true
	}

	return false
}

// SetMtuIpv4 gets a reference to the given int32 and assigns it to the MtuIpv4 field.
func (o *TsnBridgeInfo) SetMtuIpv4(v int32) {
	o.MtuIpv4 = &v
}

// GetMtuIpv6 returns the MtuIpv6 field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetMtuIpv6() int32 {
	if o == nil || IsNil(o.MtuIpv6) {
		var ret int32
		return ret
	}
	return *o.MtuIpv6
}

// GetMtuIpv6Ok returns a tuple with the MtuIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetMtuIpv6Ok() (*int32, bool) {
	if o == nil || IsNil(o.MtuIpv6) {
		return nil, false
	}
	return o.MtuIpv6, true
}

// HasMtuIpv6 returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasMtuIpv6() bool {
	if o != nil && !IsNil(o.MtuIpv6) {
		return true
	}

	return false
}

// SetMtuIpv6 gets a reference to the given int32 and assigns it to the MtuIpv6 field.
func (o *TsnBridgeInfo) SetMtuIpv6(v int32) {
	o.MtuIpv6 = &v
}

func (o TsnBridgeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TsnBridgeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BridgeId) {
		toSerialize["bridgeId"] = o.BridgeId
	}
	if !IsNil(o.DsttAddr) {
		toSerialize["dsttAddr"] = o.DsttAddr
	}
	if !IsNil(o.DsttPortNum) {
		toSerialize["dsttPortNum"] = o.DsttPortNum
	}
	if !IsNil(o.DsttResidTime) {
		toSerialize["dsttResidTime"] = o.DsttResidTime
	}
	if !IsNil(o.MtuIpv4) {
		toSerialize["mtuIpv4"] = o.MtuIpv4
	}
	if !IsNil(o.MtuIpv6) {
		toSerialize["mtuIpv6"] = o.MtuIpv6
	}
	return toSerialize, nil
}

type NullableTsnBridgeInfo struct {
	value *TsnBridgeInfo
	isSet bool
}

func (v NullableTsnBridgeInfo) Get() *TsnBridgeInfo {
	return v.value
}

func (v *NullableTsnBridgeInfo) Set(val *TsnBridgeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTsnBridgeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTsnBridgeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTsnBridgeInfo(val *TsnBridgeInfo) *NullableTsnBridgeInfo {
	return &NullableTsnBridgeInfo{value: val, isSet: true}
}

func (v NullableTsnBridgeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTsnBridgeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
