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
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
)

// AmfEventSubsSyncInfo AMF Event Subscriptions Information for synchronization
type AmfEventSubsSyncInfo struct {
	SubscriptionList []AmfEventSubscriptionInfo `json:"subscriptionList"`
}

// NewAmfEventSubsSyncInfo instantiates a new AmfEventSubsSyncInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventSubsSyncInfo(subscriptionList []AmfEventSubscriptionInfo) *AmfEventSubsSyncInfo {
	this := AmfEventSubsSyncInfo{}
	this.SubscriptionList = subscriptionList
	return &this
}

// NewAmfEventSubsSyncInfoWithDefaults instantiates a new AmfEventSubsSyncInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventSubsSyncInfoWithDefaults() *AmfEventSubsSyncInfo {
	this := AmfEventSubsSyncInfo{}
	return &this
}

// GetSubscriptionList returns the SubscriptionList field value
func (o *AmfEventSubsSyncInfo) GetSubscriptionList() []AmfEventSubscriptionInfo {
	if o == nil {
		var ret []AmfEventSubscriptionInfo
		return ret
	}

	return o.SubscriptionList
}

// GetSubscriptionListOk returns a tuple with the SubscriptionList field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubsSyncInfo) GetSubscriptionListOk() ([]AmfEventSubscriptionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionList, true
}

// SetSubscriptionList sets field value
func (o *AmfEventSubsSyncInfo) SetSubscriptionList(v []AmfEventSubscriptionInfo) {
	o.SubscriptionList = v
}

func (o AmfEventSubsSyncInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionList"] = o.SubscriptionList
	}
	return json.Marshal(toSerialize)
}

type NullableAmfEventSubsSyncInfo struct {
	value *AmfEventSubsSyncInfo
	isSet bool
}

func (v NullableAmfEventSubsSyncInfo) Get() *AmfEventSubsSyncInfo {
	return v.value
}

func (v *NullableAmfEventSubsSyncInfo) Set(val *AmfEventSubsSyncInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventSubsSyncInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventSubsSyncInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventSubsSyncInfo(val *AmfEventSubsSyncInfo) *NullableAmfEventSubsSyncInfo {
	return &NullableAmfEventSubsSyncInfo{value: val, isSet: true}
}

func (v NullableAmfEventSubsSyncInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventSubsSyncInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
