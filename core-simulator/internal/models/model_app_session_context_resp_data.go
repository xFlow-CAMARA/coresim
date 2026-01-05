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

// checks if the AppSessionContextRespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSessionContextRespData{}

// AppSessionContextRespData Describes the authorization data of an Individual Application Session Context created by the PCF.
type AppSessionContextRespData struct {
	ServAuthInfo *ServAuthInfo `json:"servAuthInfo,omitempty"`
	// QoS monitoring parameter(s) that cannot be directly notified for the indicated flows.
	DirectNotifReports []DirectNotificationReport `json:"directNotifReports,omitempty"`
	UeIds              []UeIdentityInfo           `json:"ueIds,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
}

// NewAppSessionContextRespData instantiates a new AppSessionContextRespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSessionContextRespData() *AppSessionContextRespData {
	this := AppSessionContextRespData{}
	return &this
}

// NewAppSessionContextRespDataWithDefaults instantiates a new AppSessionContextRespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSessionContextRespDataWithDefaults() *AppSessionContextRespData {
	this := AppSessionContextRespData{}
	return &this
}

// GetServAuthInfo returns the ServAuthInfo field value if set, zero value otherwise.
func (o *AppSessionContextRespData) GetServAuthInfo() ServAuthInfo {
	if o == nil || IsNil(o.ServAuthInfo) {
		var ret ServAuthInfo
		return ret
	}
	return *o.ServAuthInfo
}

// GetServAuthInfoOk returns a tuple with the ServAuthInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextRespData) GetServAuthInfoOk() (*ServAuthInfo, bool) {
	if o == nil || IsNil(o.ServAuthInfo) {
		return nil, false
	}
	return o.ServAuthInfo, true
}

// HasServAuthInfo returns a boolean if a field has been set.
func (o *AppSessionContextRespData) HasServAuthInfo() bool {
	if o != nil && !IsNil(o.ServAuthInfo) {
		return true
	}

	return false
}

// SetServAuthInfo gets a reference to the given ServAuthInfo and assigns it to the ServAuthInfo field.
func (o *AppSessionContextRespData) SetServAuthInfo(v ServAuthInfo) {
	o.ServAuthInfo = &v
}

// GetDirectNotifReports returns the DirectNotifReports field value if set, zero value otherwise.
func (o *AppSessionContextRespData) GetDirectNotifReports() []DirectNotificationReport {
	if o == nil || IsNil(o.DirectNotifReports) {
		var ret []DirectNotificationReport
		return ret
	}
	return o.DirectNotifReports
}

// GetDirectNotifReportsOk returns a tuple with the DirectNotifReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextRespData) GetDirectNotifReportsOk() ([]DirectNotificationReport, bool) {
	if o == nil || IsNil(o.DirectNotifReports) {
		return nil, false
	}
	return o.DirectNotifReports, true
}

// HasDirectNotifReports returns a boolean if a field has been set.
func (o *AppSessionContextRespData) HasDirectNotifReports() bool {
	if o != nil && !IsNil(o.DirectNotifReports) {
		return true
	}

	return false
}

// SetDirectNotifReports gets a reference to the given []DirectNotificationReport and assigns it to the DirectNotifReports field.
func (o *AppSessionContextRespData) SetDirectNotifReports(v []DirectNotificationReport) {
	o.DirectNotifReports = v
}

// GetUeIds returns the UeIds field value if set, zero value otherwise.
func (o *AppSessionContextRespData) GetUeIds() []UeIdentityInfo {
	if o == nil || IsNil(o.UeIds) {
		var ret []UeIdentityInfo
		return ret
	}
	return o.UeIds
}

// GetUeIdsOk returns a tuple with the UeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextRespData) GetUeIdsOk() ([]UeIdentityInfo, bool) {
	if o == nil || IsNil(o.UeIds) {
		return nil, false
	}
	return o.UeIds, true
}

// HasUeIds returns a boolean if a field has been set.
func (o *AppSessionContextRespData) HasUeIds() bool {
	if o != nil && !IsNil(o.UeIds) {
		return true
	}

	return false
}

// SetUeIds gets a reference to the given []UeIdentityInfo and assigns it to the UeIds field.
func (o *AppSessionContextRespData) SetUeIds(v []UeIdentityInfo) {
	o.UeIds = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AppSessionContextRespData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextRespData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AppSessionContextRespData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AppSessionContextRespData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o AppSessionContextRespData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSessionContextRespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServAuthInfo) {
		toSerialize["servAuthInfo"] = o.ServAuthInfo
	}
	if !IsNil(o.DirectNotifReports) {
		toSerialize["directNotifReports"] = o.DirectNotifReports
	}
	if !IsNil(o.UeIds) {
		toSerialize["ueIds"] = o.UeIds
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableAppSessionContextRespData struct {
	value *AppSessionContextRespData
	isSet bool
}

func (v NullableAppSessionContextRespData) Get() *AppSessionContextRespData {
	return v.value
}

func (v *NullableAppSessionContextRespData) Set(val *AppSessionContextRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContextRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContextRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContextRespData(val *AppSessionContextRespData) *NullableAppSessionContextRespData {
	return &NullableAppSessionContextRespData{value: val, isSet: true}
}

func (v NullableAppSessionContextRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContextRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
