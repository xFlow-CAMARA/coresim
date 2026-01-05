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

// checks if the TrafficCorrelationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficCorrelationInfo{}

// TrafficCorrelationInfo Contains the information for traffic correlation.
type TrafficCorrelationInfo struct {
	CorrType *CorrelationType `json:"corrType,omitempty"`
	// Identification of a set of UEs accessing the application identified by the  Application Identifier or traffic filtering information.
	TfcCorrId *string `json:"tfcCorrId,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166 with the OpenAPI defined 'nullable: true' property.
	ComEasIpv4Addr NullableString            `json:"comEasIpv4Addr,omitempty" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	ComEasIpv6Addr NullableIpv6AddrRm        `json:"comEasIpv6Addr,omitempty"`
	FqdnRange      []FqdnPatternMatchingRule `json:"fqdnRange,omitempty"`
	// String providing an URI formatted according to RFC 3986 with the OpenAPI 'nullable: true' property.
	NotifUri NullableString `json:"notifUri,omitempty"`
	// Notification correlation identifier.
	NotifCorrId NullableString `json:"notifCorrId,omitempty"`
}

// NewTrafficCorrelationInfo instantiates a new TrafficCorrelationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficCorrelationInfo() *TrafficCorrelationInfo {
	this := TrafficCorrelationInfo{}
	return &this
}

// NewTrafficCorrelationInfoWithDefaults instantiates a new TrafficCorrelationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficCorrelationInfoWithDefaults() *TrafficCorrelationInfo {
	this := TrafficCorrelationInfo{}
	return &this
}

// GetCorrType returns the CorrType field value if set, zero value otherwise.
func (o *TrafficCorrelationInfo) GetCorrType() CorrelationType {
	if o == nil || IsNil(o.CorrType) {
		var ret CorrelationType
		return ret
	}
	return *o.CorrType
}

// GetCorrTypeOk returns a tuple with the CorrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCorrelationInfo) GetCorrTypeOk() (*CorrelationType, bool) {
	if o == nil || IsNil(o.CorrType) {
		return nil, false
	}
	return o.CorrType, true
}

// HasCorrType returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasCorrType() bool {
	if o != nil && !IsNil(o.CorrType) {
		return true
	}

	return false
}

// SetCorrType gets a reference to the given CorrelationType and assigns it to the CorrType field.
func (o *TrafficCorrelationInfo) SetCorrType(v CorrelationType) {
	o.CorrType = &v
}

// GetTfcCorrId returns the TfcCorrId field value if set, zero value otherwise.
func (o *TrafficCorrelationInfo) GetTfcCorrId() string {
	if o == nil || IsNil(o.TfcCorrId) {
		var ret string
		return ret
	}
	return *o.TfcCorrId
}

// GetTfcCorrIdOk returns a tuple with the TfcCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCorrelationInfo) GetTfcCorrIdOk() (*string, bool) {
	if o == nil || IsNil(o.TfcCorrId) {
		return nil, false
	}
	return o.TfcCorrId, true
}

// HasTfcCorrId returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasTfcCorrId() bool {
	if o != nil && !IsNil(o.TfcCorrId) {
		return true
	}

	return false
}

// SetTfcCorrId gets a reference to the given string and assigns it to the TfcCorrId field.
func (o *TrafficCorrelationInfo) SetTfcCorrId(v string) {
	o.TfcCorrId = &v
}

// GetComEasIpv4Addr returns the ComEasIpv4Addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficCorrelationInfo) GetComEasIpv4Addr() string {
	if o == nil || IsNil(o.ComEasIpv4Addr.Get()) {
		var ret string
		return ret
	}
	return *o.ComEasIpv4Addr.Get()
}

// GetComEasIpv4AddrOk returns a tuple with the ComEasIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficCorrelationInfo) GetComEasIpv4AddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComEasIpv4Addr.Get(), o.ComEasIpv4Addr.IsSet()
}

// HasComEasIpv4Addr returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasComEasIpv4Addr() bool {
	if o != nil && o.ComEasIpv4Addr.IsSet() {
		return true
	}

	return false
}

// SetComEasIpv4Addr gets a reference to the given NullableString and assigns it to the ComEasIpv4Addr field.
func (o *TrafficCorrelationInfo) SetComEasIpv4Addr(v string) {
	o.ComEasIpv4Addr.Set(&v)
}

// SetComEasIpv4AddrNil sets the value for ComEasIpv4Addr to be an explicit nil
func (o *TrafficCorrelationInfo) SetComEasIpv4AddrNil() {
	o.ComEasIpv4Addr.Set(nil)
}

// UnsetComEasIpv4Addr ensures that no value is present for ComEasIpv4Addr, not even an explicit nil
func (o *TrafficCorrelationInfo) UnsetComEasIpv4Addr() {
	o.ComEasIpv4Addr.Unset()
}

// GetComEasIpv6Addr returns the ComEasIpv6Addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficCorrelationInfo) GetComEasIpv6Addr() Ipv6AddrRm {
	if o == nil || IsNil(o.ComEasIpv6Addr.Get()) {
		var ret Ipv6AddrRm
		return ret
	}
	return *o.ComEasIpv6Addr.Get()
}

// GetComEasIpv6AddrOk returns a tuple with the ComEasIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficCorrelationInfo) GetComEasIpv6AddrOk() (*Ipv6AddrRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComEasIpv6Addr.Get(), o.ComEasIpv6Addr.IsSet()
}

// HasComEasIpv6Addr returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasComEasIpv6Addr() bool {
	if o != nil && o.ComEasIpv6Addr.IsSet() {
		return true
	}

	return false
}

// SetComEasIpv6Addr gets a reference to the given NullableIpv6AddrRm and assigns it to the ComEasIpv6Addr field.
func (o *TrafficCorrelationInfo) SetComEasIpv6Addr(v Ipv6AddrRm) {
	o.ComEasIpv6Addr.Set(&v)
}

// SetComEasIpv6AddrNil sets the value for ComEasIpv6Addr to be an explicit nil
func (o *TrafficCorrelationInfo) SetComEasIpv6AddrNil() {
	o.ComEasIpv6Addr.Set(nil)
}

// UnsetComEasIpv6Addr ensures that no value is present for ComEasIpv6Addr, not even an explicit nil
func (o *TrafficCorrelationInfo) UnsetComEasIpv6Addr() {
	o.ComEasIpv6Addr.Unset()
}

// GetFqdnRange returns the FqdnRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficCorrelationInfo) GetFqdnRange() []FqdnPatternMatchingRule {
	if o == nil {
		var ret []FqdnPatternMatchingRule
		return ret
	}
	return o.FqdnRange
}

// GetFqdnRangeOk returns a tuple with the FqdnRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficCorrelationInfo) GetFqdnRangeOk() ([]FqdnPatternMatchingRule, bool) {
	if o == nil || IsNil(o.FqdnRange) {
		return nil, false
	}
	return o.FqdnRange, true
}

// HasFqdnRange returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasFqdnRange() bool {
	if o != nil && !IsNil(o.FqdnRange) {
		return true
	}

	return false
}

// SetFqdnRange gets a reference to the given []FqdnPatternMatchingRule and assigns it to the FqdnRange field.
func (o *TrafficCorrelationInfo) SetFqdnRange(v []FqdnPatternMatchingRule) {
	o.FqdnRange = v
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficCorrelationInfo) GetNotifUri() string {
	if o == nil || IsNil(o.NotifUri.Get()) {
		var ret string
		return ret
	}
	return *o.NotifUri.Get()
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficCorrelationInfo) GetNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifUri.Get(), o.NotifUri.IsSet()
}

// HasNotifUri returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasNotifUri() bool {
	if o != nil && o.NotifUri.IsSet() {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given NullableString and assigns it to the NotifUri field.
func (o *TrafficCorrelationInfo) SetNotifUri(v string) {
	o.NotifUri.Set(&v)
}

// SetNotifUriNil sets the value for NotifUri to be an explicit nil
func (o *TrafficCorrelationInfo) SetNotifUriNil() {
	o.NotifUri.Set(nil)
}

// UnsetNotifUri ensures that no value is present for NotifUri, not even an explicit nil
func (o *TrafficCorrelationInfo) UnsetNotifUri() {
	o.NotifUri.Unset()
}

// GetNotifCorrId returns the NotifCorrId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficCorrelationInfo) GetNotifCorrId() string {
	if o == nil || IsNil(o.NotifCorrId.Get()) {
		var ret string
		return ret
	}
	return *o.NotifCorrId.Get()
}

// GetNotifCorrIdOk returns a tuple with the NotifCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficCorrelationInfo) GetNotifCorrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifCorrId.Get(), o.NotifCorrId.IsSet()
}

// HasNotifCorrId returns a boolean if a field has been set.
func (o *TrafficCorrelationInfo) HasNotifCorrId() bool {
	if o != nil && o.NotifCorrId.IsSet() {
		return true
	}

	return false
}

// SetNotifCorrId gets a reference to the given NullableString and assigns it to the NotifCorrId field.
func (o *TrafficCorrelationInfo) SetNotifCorrId(v string) {
	o.NotifCorrId.Set(&v)
}

// SetNotifCorrIdNil sets the value for NotifCorrId to be an explicit nil
func (o *TrafficCorrelationInfo) SetNotifCorrIdNil() {
	o.NotifCorrId.Set(nil)
}

// UnsetNotifCorrId ensures that no value is present for NotifCorrId, not even an explicit nil
func (o *TrafficCorrelationInfo) UnsetNotifCorrId() {
	o.NotifCorrId.Unset()
}

func (o TrafficCorrelationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficCorrelationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrType) {
		toSerialize["corrType"] = o.CorrType
	}
	if !IsNil(o.TfcCorrId) {
		toSerialize["tfcCorrId"] = o.TfcCorrId
	}
	if o.ComEasIpv4Addr.IsSet() {
		toSerialize["comEasIpv4Addr"] = o.ComEasIpv4Addr.Get()
	}
	if o.ComEasIpv6Addr.IsSet() {
		toSerialize["comEasIpv6Addr"] = o.ComEasIpv6Addr.Get()
	}
	if o.FqdnRange != nil {
		toSerialize["fqdnRange"] = o.FqdnRange
	}
	if o.NotifUri.IsSet() {
		toSerialize["notifUri"] = o.NotifUri.Get()
	}
	if o.NotifCorrId.IsSet() {
		toSerialize["notifCorrId"] = o.NotifCorrId.Get()
	}
	return toSerialize, nil
}

type NullableTrafficCorrelationInfo struct {
	value *TrafficCorrelationInfo
	isSet bool
}

func (v NullableTrafficCorrelationInfo) Get() *TrafficCorrelationInfo {
	return v.value
}

func (v *NullableTrafficCorrelationInfo) Set(val *TrafficCorrelationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficCorrelationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficCorrelationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficCorrelationInfo(val *TrafficCorrelationInfo) *NullableTrafficCorrelationInfo {
	return &NullableTrafficCorrelationInfo{value: val, isSet: true}
}

func (v NullableTrafficCorrelationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficCorrelationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
