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

// checks if the UeIdentityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeIdentityInfo{}

// UeIdentityInfo Represents 5GS-Level UE identities.
type UeIdentityInfo struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei *string
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string
}

// NewUeIdentityInfo instantiates a new UeIdentityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeIdentityInfo() *UeIdentityInfo {
	this := UeIdentityInfo{}
	return &this
}

// NewUeIdentityInfoWithDefaults instantiates a new UeIdentityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeIdentityInfoWithDefaults() *UeIdentityInfo {
	this := UeIdentityInfo{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UeIdentityInfo) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeIdentityInfo) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UeIdentityInfo) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UeIdentityInfo) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *UeIdentityInfo) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeIdentityInfo) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *UeIdentityInfo) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *UeIdentityInfo) SetPei(v string) {
	o.Pei = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UeIdentityInfo) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeIdentityInfo) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UeIdentityInfo) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UeIdentityInfo) SetSupi(v string) {
	o.Supi = &v
}

func (o UeIdentityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeIdentityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	return toSerialize, nil
}

type NullableUeIdentityInfo struct {
	value *UeIdentityInfo
	isSet bool
}

func (v NullableUeIdentityInfo) Get() *UeIdentityInfo {
	return v.value
}

func (v *NullableUeIdentityInfo) Set(val *UeIdentityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeIdentityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeIdentityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeIdentityInfo(val *UeIdentityInfo) *NullableUeIdentityInfo {
	return &NullableUeIdentityInfo{value: val, isSet: true}
}

func (v NullableUeIdentityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeIdentityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
