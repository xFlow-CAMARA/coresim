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

// checks if the BatOffsetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatOffsetInfo{}

// BatOffsetInfo Indicates the offset of the BAT and the optionally adjusted periodicity.
type BatOffsetInfo struct {
	// Indicates the BAT offset of the arrival time of the data burst in units of milliseconds.
	RanBatOffsetNotif int32 `json:"ranBatOffsetNotif"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AdjPeriod *int32 `json:"adjPeriod,omitempty"`
	// Identification of the flows. If no flows are provided, the BAT offset applies for all flows of the AF session.
	Flows []Flows `json:"flows,omitempty"`
}

type _BatOffsetInfo BatOffsetInfo

// NewBatOffsetInfo instantiates a new BatOffsetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatOffsetInfo(ranBatOffsetNotif int32) *BatOffsetInfo {
	this := BatOffsetInfo{}
	this.RanBatOffsetNotif = ranBatOffsetNotif
	return &this
}

// NewBatOffsetInfoWithDefaults instantiates a new BatOffsetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatOffsetInfoWithDefaults() *BatOffsetInfo {
	this := BatOffsetInfo{}
	return &this
}

// GetRanBatOffsetNotif returns the RanBatOffsetNotif field value
func (o *BatOffsetInfo) GetRanBatOffsetNotif() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RanBatOffsetNotif
}

// GetRanBatOffsetNotifOk returns a tuple with the RanBatOffsetNotif field value
// and a boolean to check if the value has been set.
func (o *BatOffsetInfo) GetRanBatOffsetNotifOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RanBatOffsetNotif, true
}

// SetRanBatOffsetNotif sets field value
func (o *BatOffsetInfo) SetRanBatOffsetNotif(v int32) {
	o.RanBatOffsetNotif = v
}

// GetAdjPeriod returns the AdjPeriod field value if set, zero value otherwise.
func (o *BatOffsetInfo) GetAdjPeriod() int32 {
	if o == nil || IsNil(o.AdjPeriod) {
		var ret int32
		return ret
	}
	return *o.AdjPeriod
}

// GetAdjPeriodOk returns a tuple with the AdjPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatOffsetInfo) GetAdjPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.AdjPeriod) {
		return nil, false
	}
	return o.AdjPeriod, true
}

// HasAdjPeriod returns a boolean if a field has been set.
func (o *BatOffsetInfo) HasAdjPeriod() bool {
	if o != nil && !IsNil(o.AdjPeriod) {
		return true
	}

	return false
}

// SetAdjPeriod gets a reference to the given int32 and assigns it to the AdjPeriod field.
func (o *BatOffsetInfo) SetAdjPeriod(v int32) {
	o.AdjPeriod = &v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *BatOffsetInfo) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatOffsetInfo) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *BatOffsetInfo) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *BatOffsetInfo) SetFlows(v []Flows) {
	o.Flows = v
}

func (o BatOffsetInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatOffsetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ranBatOffsetNotif"] = o.RanBatOffsetNotif
	if !IsNil(o.AdjPeriod) {
		toSerialize["adjPeriod"] = o.AdjPeriod
	}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return toSerialize, nil
}

func (o *BatOffsetInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ranBatOffsetNotif",
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

	varBatOffsetInfo := _BatOffsetInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatOffsetInfo)

	if err != nil {
		return err
	}

	*o = BatOffsetInfo(varBatOffsetInfo)

	return err
}

type NullableBatOffsetInfo struct {
	value *BatOffsetInfo
	isSet bool
}

func (v NullableBatOffsetInfo) Get() *BatOffsetInfo {
	return v.value
}

func (v *NullableBatOffsetInfo) Set(val *BatOffsetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBatOffsetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBatOffsetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatOffsetInfo(val *BatOffsetInfo) *NullableBatOffsetInfo {
	return &NullableBatOffsetInfo{value: val, isSet: true}
}

func (v NullableBatOffsetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatOffsetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
