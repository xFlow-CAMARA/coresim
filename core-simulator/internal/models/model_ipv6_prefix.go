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
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
)

// Ipv6Prefix String identifying an IPv6 address prefix formatted according to clause 4 of RFC 5952. IPv6Prefix data type may contain an individual /128 IPv6 address.
type Ipv6Prefix struct {
}

// NewIpv6Prefix instantiates a new Ipv6Prefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Prefix() *Ipv6Prefix {
	this := Ipv6Prefix{}
	return &this
}

// NewIpv6PrefixWithDefaults instantiates a new Ipv6Prefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6PrefixWithDefaults() *Ipv6Prefix {
	this := Ipv6Prefix{}
	return &this
}

func (o Ipv6Prefix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableIpv6Prefix struct {
	value *Ipv6Prefix
	isSet bool
}

func (v NullableIpv6Prefix) Get() *Ipv6Prefix {
	return v.value
}

func (v *NullableIpv6Prefix) Set(val *Ipv6Prefix) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Prefix) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Prefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Prefix(val *Ipv6Prefix) *NullableIpv6Prefix {
	return &NullableIpv6Prefix{value: val, isSet: true}
}

func (v NullableIpv6Prefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Prefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
