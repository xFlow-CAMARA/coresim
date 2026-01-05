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

// checks if the AppDetectionReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDetectionReport{}

// AppDetectionReport Indicates the start or stop of the detected application traffic and the application identifier of the detected application traffic.
type AppDetectionReport struct {
	AdNotifType AppDetectionNotifType `json:"adNotifType"`
	// Contains an AF application identifier.
	AfAppId string `json:"afAppId"`
}

type _AppDetectionReport AppDetectionReport

// NewAppDetectionReport instantiates a new AppDetectionReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDetectionReport(adNotifType AppDetectionNotifType, afAppId string) *AppDetectionReport {
	this := AppDetectionReport{}
	this.AdNotifType = adNotifType
	this.AfAppId = afAppId
	return &this
}

// NewAppDetectionReportWithDefaults instantiates a new AppDetectionReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDetectionReportWithDefaults() *AppDetectionReport {
	this := AppDetectionReport{}
	return &this
}

// GetAdNotifType returns the AdNotifType field value
func (o *AppDetectionReport) GetAdNotifType() AppDetectionNotifType {
	if o == nil {
		var ret AppDetectionNotifType
		return ret
	}

	return o.AdNotifType
}

// GetAdNotifTypeOk returns a tuple with the AdNotifType field value
// and a boolean to check if the value has been set.
func (o *AppDetectionReport) GetAdNotifTypeOk() (*AppDetectionNotifType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdNotifType, true
}

// SetAdNotifType sets field value
func (o *AppDetectionReport) SetAdNotifType(v AppDetectionNotifType) {
	o.AdNotifType = v
}

// GetAfAppId returns the AfAppId field value
func (o *AppDetectionReport) GetAfAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value
// and a boolean to check if the value has been set.
func (o *AppDetectionReport) GetAfAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfAppId, true
}

// SetAfAppId sets field value
func (o *AppDetectionReport) SetAfAppId(v string) {
	o.AfAppId = v
}

func (o AppDetectionReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDetectionReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adNotifType"] = o.AdNotifType
	toSerialize["afAppId"] = o.AfAppId
	return toSerialize, nil
}

func (o *AppDetectionReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"adNotifType",
		"afAppId",
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

	varAppDetectionReport := _AppDetectionReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppDetectionReport)

	if err != nil {
		return err
	}

	*o = AppDetectionReport(varAppDetectionReport)

	return err
}

type NullableAppDetectionReport struct {
	value *AppDetectionReport
	isSet bool
}

func (v NullableAppDetectionReport) Get() *AppDetectionReport {
	return v.value
}

func (v *NullableAppDetectionReport) Set(val *AppDetectionReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetectionReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetectionReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetectionReport(val *AppDetectionReport) *NullableAppDetectionReport {
	return &NullableAppDetectionReport{value: val, isSet: true}
}

func (v NullableAppDetectionReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetectionReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
