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

// checks if the PcscfRestorationRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcscfRestorationRequestData{}

// PcscfRestorationRequestData Indicates P-CSCF restoration.
type PcscfRestorationRequestData struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn       *string
	IpDomain  *string
	SliceInfo *Snssai
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	UeIpv4 *string
	UeIpv6 *Ipv6Addr
}

// NewPcscfRestorationRequestData instantiates a new PcscfRestorationRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcscfRestorationRequestData() *PcscfRestorationRequestData {
	this := PcscfRestorationRequestData{}
	return &this
}

// NewPcscfRestorationRequestDataWithDefaults instantiates a new PcscfRestorationRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcscfRestorationRequestDataWithDefaults() *PcscfRestorationRequestData {
	this := PcscfRestorationRequestData{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PcscfRestorationRequestData) SetDnn(v string) {
	o.Dnn = &v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetIpDomain() string {
	if o == nil || IsNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetIpDomainOk() (*string, bool) {
	if o == nil || IsNil(o.IpDomain) {
		return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasIpDomain() bool {
	if o != nil && !IsNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *PcscfRestorationRequestData) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetSliceInfo returns the SliceInfo field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetSliceInfo() Snssai {
	if o == nil || IsNil(o.SliceInfo) {
		var ret Snssai
		return ret
	}
	return *o.SliceInfo
}

// GetSliceInfoOk returns a tuple with the SliceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetSliceInfoOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SliceInfo) {
		return nil, false
	}
	return o.SliceInfo, true
}

// HasSliceInfo returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasSliceInfo() bool {
	if o != nil && !IsNil(o.SliceInfo) {
		return true
	}

	return false
}

// SetSliceInfo gets a reference to the given Snssai and assigns it to the SliceInfo field.
func (o *PcscfRestorationRequestData) SetSliceInfo(v Snssai) {
	o.SliceInfo = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *PcscfRestorationRequestData) SetSupi(v string) {
	o.Supi = &v
}

// GetUeIpv4 returns the UeIpv4 field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetUeIpv4() string {
	if o == nil || IsNil(o.UeIpv4) {
		var ret string
		return ret
	}
	return *o.UeIpv4
}

// GetUeIpv4Ok returns a tuple with the UeIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetUeIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.UeIpv4) {
		return nil, false
	}
	return o.UeIpv4, true
}

// HasUeIpv4 returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasUeIpv4() bool {
	if o != nil && !IsNil(o.UeIpv4) {
		return true
	}

	return false
}

// SetUeIpv4 gets a reference to the given string and assigns it to the UeIpv4 field.
func (o *PcscfRestorationRequestData) SetUeIpv4(v string) {
	o.UeIpv4 = &v
}

// GetUeIpv6 returns the UeIpv6 field value if set, zero value otherwise.
func (o *PcscfRestorationRequestData) GetUeIpv6() Ipv6Addr {
	if o == nil || IsNil(o.UeIpv6) {
		var ret Ipv6Addr
		return ret
	}
	return *o.UeIpv6
}

// GetUeIpv6Ok returns a tuple with the UeIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfRestorationRequestData) GetUeIpv6Ok() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.UeIpv6) {
		return nil, false
	}
	return o.UeIpv6, true
}

// HasUeIpv6 returns a boolean if a field has been set.
func (o *PcscfRestorationRequestData) HasUeIpv6() bool {
	if o != nil && !IsNil(o.UeIpv6) {
		return true
	}

	return false
}

// SetUeIpv6 gets a reference to the given Ipv6Addr and assigns it to the UeIpv6 field.
func (o *PcscfRestorationRequestData) SetUeIpv6(v Ipv6Addr) {
	o.UeIpv6 = &v
}

func (o PcscfRestorationRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcscfRestorationRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
	}
	if !IsNil(o.SliceInfo) {
		toSerialize["sliceInfo"] = o.SliceInfo
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.UeIpv4) {
		toSerialize["ueIpv4"] = o.UeIpv4
	}
	if !IsNil(o.UeIpv6) {
		toSerialize["ueIpv6"] = o.UeIpv6
	}
	return toSerialize, nil
}

type NullablePcscfRestorationRequestData struct {
	value *PcscfRestorationRequestData
	isSet bool
}

func (v NullablePcscfRestorationRequestData) Get() *PcscfRestorationRequestData {
	return v.value
}

func (v *NullablePcscfRestorationRequestData) Set(val *PcscfRestorationRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePcscfRestorationRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePcscfRestorationRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcscfRestorationRequestData(val *PcscfRestorationRequestData) *NullablePcscfRestorationRequestData {
	return &NullablePcscfRestorationRequestData{value: val, isSet: true}
}

func (v NullablePcscfRestorationRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcscfRestorationRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
