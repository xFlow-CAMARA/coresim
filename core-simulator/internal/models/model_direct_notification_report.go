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

// checks if the DirectNotificationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectNotificationReport{}

// DirectNotificationReport Represents the QoS monitoring parameters that cannot be directly notified for the indicated flows.
type DirectNotificationReport struct {
	QosMonParamType QosMonitoringParamType `json:"qosMonParamType"`
	Flows           []Flows                `json:"flows,omitempty"`
}

type _DirectNotificationReport DirectNotificationReport

// NewDirectNotificationReport instantiates a new DirectNotificationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectNotificationReport(qosMonParamType QosMonitoringParamType) *DirectNotificationReport {
	this := DirectNotificationReport{}
	this.QosMonParamType = qosMonParamType
	return &this
}

// NewDirectNotificationReportWithDefaults instantiates a new DirectNotificationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectNotificationReportWithDefaults() *DirectNotificationReport {
	this := DirectNotificationReport{}
	return &this
}

// GetQosMonParamType returns the QosMonParamType field value
func (o *DirectNotificationReport) GetQosMonParamType() QosMonitoringParamType {
	if o == nil {
		var ret QosMonitoringParamType
		return ret
	}

	return o.QosMonParamType
}

// GetQosMonParamTypeOk returns a tuple with the QosMonParamType field value
// and a boolean to check if the value has been set.
func (o *DirectNotificationReport) GetQosMonParamTypeOk() (*QosMonitoringParamType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QosMonParamType, true
}

// SetQosMonParamType sets field value
func (o *DirectNotificationReport) SetQosMonParamType(v QosMonitoringParamType) {
	o.QosMonParamType = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *DirectNotificationReport) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectNotificationReport) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *DirectNotificationReport) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *DirectNotificationReport) SetFlows(v []Flows) {
	o.Flows = v
}

func (o DirectNotificationReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectNotificationReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qosMonParamType"] = o.QosMonParamType
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

func (o *DirectNotificationReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"qosMonParamType",
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

	varDirectNotificationReport := _DirectNotificationReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDirectNotificationReport)

	if err != nil {
		return err
	}

	*o = DirectNotificationReport(varDirectNotificationReport)

	return err
}

type NullableDirectNotificationReport struct {
	value *DirectNotificationReport
	isSet bool
}

func (v NullableDirectNotificationReport) Get() *DirectNotificationReport {
	return v.value
}

func (v *NullableDirectNotificationReport) Set(val *DirectNotificationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectNotificationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectNotificationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectNotificationReport(val *DirectNotificationReport) *NullableDirectNotificationReport {
	return &NullableDirectNotificationReport{value: val, isSet: true}
}

func (v NullableDirectNotificationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectNotificationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
