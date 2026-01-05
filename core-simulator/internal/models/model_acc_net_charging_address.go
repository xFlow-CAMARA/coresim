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

// checks if the AccNetChargingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccNetChargingAddress{}

// AccNetChargingAddress Describes the network entity within the access network performing charging
type AccNetChargingAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AnChargIpv4Addr *string
	AnChargIpv6Addr *Ipv6Addr
}

// NewAccNetChargingAddress instantiates a new AccNetChargingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccNetChargingAddress() *AccNetChargingAddress {
	this := AccNetChargingAddress{}
	return &this
}

// NewAccNetChargingAddressWithDefaults instantiates a new AccNetChargingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccNetChargingAddressWithDefaults() *AccNetChargingAddress {
	this := AccNetChargingAddress{}
	return &this
}

// GetAnChargIpv4Addr returns the AnChargIpv4Addr field value if set, zero value otherwise.
func (o *AccNetChargingAddress) GetAnChargIpv4Addr() string {
	if o == nil || IsNil(o.AnChargIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AnChargIpv4Addr
}

// GetAnChargIpv4AddrOk returns a tuple with the AnChargIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChargingAddress) GetAnChargIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AnChargIpv4Addr) {
		return nil, false
	}
	return o.AnChargIpv4Addr, true
}

// HasAnChargIpv4Addr returns a boolean if a field has been set.
func (o *AccNetChargingAddress) HasAnChargIpv4Addr() bool {
	if o != nil && !IsNil(o.AnChargIpv4Addr) {
		return true
	}

	return false
}

// SetAnChargIpv4Addr gets a reference to the given string and assigns it to the AnChargIpv4Addr field.
func (o *AccNetChargingAddress) SetAnChargIpv4Addr(v string) {
	o.AnChargIpv4Addr = &v
}

// GetAnChargIpv6Addr returns the AnChargIpv6Addr field value if set, zero value otherwise.
func (o *AccNetChargingAddress) GetAnChargIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AnChargIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AnChargIpv6Addr
}

// GetAnChargIpv6AddrOk returns a tuple with the AnChargIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccNetChargingAddress) GetAnChargIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AnChargIpv6Addr) {
		return nil, false
	}
	return o.AnChargIpv6Addr, true
}

// HasAnChargIpv6Addr returns a boolean if a field has been set.
func (o *AccNetChargingAddress) HasAnChargIpv6Addr() bool {
	if o != nil && !IsNil(o.AnChargIpv6Addr) {
		return true
	}

	return false
}

// SetAnChargIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AnChargIpv6Addr field.
func (o *AccNetChargingAddress) SetAnChargIpv6Addr(v Ipv6Addr) {
	o.AnChargIpv6Addr = &v
}

func (o AccNetChargingAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccNetChargingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnChargIpv4Addr) {
		toSerialize["anChargIpv4Addr"] = o.AnChargIpv4Addr
	}
	if !IsNil(o.AnChargIpv6Addr) {
		toSerialize["anChargIpv6Addr"] = o.AnChargIpv6Addr
	}
	return toSerialize, nil
}

type NullableAccNetChargingAddress struct {
	value *AccNetChargingAddress
	isSet bool
}

func (v NullableAccNetChargingAddress) Get() *AccNetChargingAddress {
	return v.value
}

func (v *NullableAccNetChargingAddress) Set(val *AccNetChargingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccNetChargingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccNetChargingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccNetChargingAddress(val *AccNetChargingAddress) *NullableAccNetChargingAddress {
	return &NullableAccNetChargingAddress{value: val, isSet: true}
}

func (v NullableAccNetChargingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccNetChargingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
