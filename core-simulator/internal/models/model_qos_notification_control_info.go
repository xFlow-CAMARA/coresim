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

// checks if the QosNotificationControlInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosNotificationControlInfo{}

// QosNotificationControlInfo Indicates whether the QoS targets for a GRB flow are not guaranteed or guaranteed again.
type QosNotificationControlInfo struct {
	NotifType QosNotifType `json:"notifType"`
	Flows     []Flows      `json:"flows,omitempty"`
	// Indicates the alternative service requirement NG-RAN can guarantee. When it is omitted and the notifType attribute is set to NOT_GUAARANTEED it indicates that the lowest priority alternative alternative service requirement could not be fulfilled by NG-RAN.
	AltSerReq *string `json:"altSerReq,omitempty"`
	// When present and set to true it indicates that Alternative Service Requirements are not  supported by NG-RAN.
	AltSerReqNotSuppInd *bool `json:"altSerReqNotSuppInd,omitempty"`
}

type _QosNotificationControlInfo QosNotificationControlInfo

// NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosNotificationControlInfo(notifType QosNotifType) *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	this.NotifType = notifType
	return &this
}

// NewQosNotificationControlInfoWithDefaults instantiates a new QosNotificationControlInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosNotificationControlInfoWithDefaults() *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	return &this
}

// GetNotifType returns the NotifType field value
func (o *QosNotificationControlInfo) GetNotifType() QosNotifType {
	if o == nil {
		var ret QosNotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetNotifTypeOk() (*QosNotifType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *QosNotificationControlInfo) SetNotifType(v QosNotifType) {
	o.NotifType = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *QosNotificationControlInfo) SetFlows(v []Flows) {
	o.Flows = v
}

// GetAltSerReq returns the AltSerReq field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltSerReq() string {
	if o == nil || IsNil(o.AltSerReq) {
		var ret string
		return ret
	}
	return *o.AltSerReq
}

// GetAltSerReqOk returns a tuple with the AltSerReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltSerReqOk() (*string, bool) {
	if o == nil || IsNil(o.AltSerReq) {
		return nil, false
	}
	return o.AltSerReq, true
}

// HasAltSerReq returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltSerReq() bool {
	if o != nil && !IsNil(o.AltSerReq) {
		return true
	}

	return false
}

// SetAltSerReq gets a reference to the given string and assigns it to the AltSerReq field.
func (o *QosNotificationControlInfo) SetAltSerReq(v string) {
	o.AltSerReq = &v
}

// GetAltSerReqNotSuppInd returns the AltSerReqNotSuppInd field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltSerReqNotSuppInd() bool {
	if o == nil || IsNil(o.AltSerReqNotSuppInd) {
		var ret bool
		return ret
	}
	return *o.AltSerReqNotSuppInd
}

// GetAltSerReqNotSuppIndOk returns a tuple with the AltSerReqNotSuppInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltSerReqNotSuppIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AltSerReqNotSuppInd) {
		return nil, false
	}
	return o.AltSerReqNotSuppInd, true
}

// HasAltSerReqNotSuppInd returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltSerReqNotSuppInd() bool {
	if o != nil && !IsNil(o.AltSerReqNotSuppInd) {
		return true
	}

	return false
}

// SetAltSerReqNotSuppInd gets a reference to the given bool and assigns it to the AltSerReqNotSuppInd field.
func (o *QosNotificationControlInfo) SetAltSerReqNotSuppInd(v bool) {
	o.AltSerReqNotSuppInd = &v
}

func (o QosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosNotificationControlInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifType"] = o.NotifType
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !IsNil(o.AltSerReq) {
		toSerialize["altSerReq"] = o.AltSerReq
	}
	if !IsNil(o.AltSerReqNotSuppInd) {
		toSerialize["altSerReqNotSuppInd"] = o.AltSerReqNotSuppInd
	}
	return toSerialize, nil
}

func (o *QosNotificationControlInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifType",
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

	varQosNotificationControlInfo := _QosNotificationControlInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQosNotificationControlInfo)

	if err != nil {
		return err
	}

	*o = QosNotificationControlInfo(varQosNotificationControlInfo)

	return err
}

type NullableQosNotificationControlInfo struct {
	value *QosNotificationControlInfo
	isSet bool
}

func (v NullableQosNotificationControlInfo) Get() *QosNotificationControlInfo {
	return v.value
}

func (v *NullableQosNotificationControlInfo) Set(val *QosNotificationControlInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosNotificationControlInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosNotificationControlInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosNotificationControlInfo(val *QosNotificationControlInfo) *NullableQosNotificationControlInfo {
	return &NullableQosNotificationControlInfo{value: val, isSet: true}
}

func (v NullableQosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosNotificationControlInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
