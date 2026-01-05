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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AfEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfEventSubscription{}

// AfEventSubscription Describes the event information delivered in the subscription.
type AfEventSubscription struct {
	Event       AfEvent        `json:"event"`
	NotifMethod *AfNotifMethod `json:"notifMethod,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// indicating a time in seconds.
	WaitTime        *int32                  `json:"waitTime,omitempty"`
	QosMonParamType *QosMonitoringParamType `json:"qosMonParamType,omitempty"`
}

type _AfEventSubscription AfEventSubscription

// NewAfEventSubscription instantiates a new AfEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventSubscription(event AfEvent) *AfEventSubscription {
	this := AfEventSubscription{}
	this.Event = event
	return &this
}

// NewAfEventSubscriptionWithDefaults instantiates a new AfEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventSubscriptionWithDefaults() *AfEventSubscription {
	this := AfEventSubscription{}
	return &this
}

// GetEvent returns the Event field value
func (o *AfEventSubscription) GetEvent() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetEventOk() (*AfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AfEventSubscription) SetEvent(v AfEvent) {
	o.Event = v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *AfEventSubscription) GetNotifMethod() AfNotifMethod {
	if o == nil || IsNil(o.NotifMethod) {
		var ret AfNotifMethod
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetNotifMethodOk() (*AfNotifMethod, bool) {
	if o == nil || IsNil(o.NotifMethod) {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *AfEventSubscription) HasNotifMethod() bool {
	if o != nil && !IsNil(o.NotifMethod) {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given AfNotifMethod and assigns it to the NotifMethod field.
func (o *AfEventSubscription) SetNotifMethod(v AfNotifMethod) {
	o.NotifMethod = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *AfEventSubscription) GetRepPeriod() int32 {
	if o == nil || IsNil(o.RepPeriod) {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RepPeriod) {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *AfEventSubscription) HasRepPeriod() bool {
	if o != nil && !IsNil(o.RepPeriod) {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *AfEventSubscription) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *AfEventSubscription) GetWaitTime() int32 {
	if o == nil || IsNil(o.WaitTime) {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTime) {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *AfEventSubscription) HasWaitTime() bool {
	if o != nil && !IsNil(o.WaitTime) {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *AfEventSubscription) SetWaitTime(v int32) {
	o.WaitTime = &v
}

// GetQosMonParamType returns the QosMonParamType field value if set, zero value otherwise.
func (o *AfEventSubscription) GetQosMonParamType() QosMonitoringParamType {
	if o == nil || IsNil(o.QosMonParamType) {
		var ret QosMonitoringParamType
		return ret
	}
	return *o.QosMonParamType
}

// GetQosMonParamTypeOk returns a tuple with the QosMonParamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetQosMonParamTypeOk() (*QosMonitoringParamType, bool) {
	if o == nil || IsNil(o.QosMonParamType) {
		return nil, false
	}
	return o.QosMonParamType, true
}

// HasQosMonParamType returns a boolean if a field has been set.
func (o *AfEventSubscription) HasQosMonParamType() bool {
	if o != nil && !IsNil(o.QosMonParamType) {
		return true
	}

	return false
}

// SetQosMonParamType gets a reference to the given QosMonitoringParamType and assigns it to the QosMonParamType field.
func (o *AfEventSubscription) SetQosMonParamType(v QosMonitoringParamType) {
	o.QosMonParamType = &v
}

func (o AfEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.NotifMethod) {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if !IsNil(o.RepPeriod) {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if !IsNil(o.WaitTime) {
		toSerialize["waitTime"] = o.WaitTime
	}
	if !IsNil(o.QosMonParamType) {
		toSerialize["qosMonParamType"] = o.QosMonParamType
	}
	return toSerialize, nil
}

func (o *AfEventSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAfEventSubscription := _AfEventSubscription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAfEventSubscription)

	if err != nil {
		return err
	}

	*o = AfEventSubscription(varAfEventSubscription)

	return err
}

type NullableAfEventSubscription struct {
	value *AfEventSubscription
	isSet bool
}

func (v NullableAfEventSubscription) Get() *AfEventSubscription {
	return v.value
}

func (v *NullableAfEventSubscription) Set(val *AfEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventSubscription(val *AfEventSubscription) *NullableAfEventSubscription {
	return &NullableAfEventSubscription{value: val, isSet: true}
}

func (v NullableAfEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
