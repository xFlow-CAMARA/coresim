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
	"time"
)

// checks if the TscaiInputContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TscaiInputContainer{}

// TscaiInputContainer Indicates TSC Traffic pattern.
type TscaiInputContainer struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Periodicity *int32 `json:"periodicity,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	BurstArrivalTime *time.Time `json:"burstArrivalTime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInNumMsg *int32 `json:"surTimeInNumMsg,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInTime       *int32                   `json:"surTimeInTime,omitempty"`
	BurstArrivalTimeWnd *TimeWindow              `json:"burstArrivalTimeWnd,omitempty"`
	PeriodicityRange    NullablePeriodicityRange `json:"periodicityRange,omitempty"`
}

// NewTscaiInputContainer instantiates a new TscaiInputContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTscaiInputContainer() *TscaiInputContainer {
	this := TscaiInputContainer{}
	return &this
}

// NewTscaiInputContainerWithDefaults instantiates a new TscaiInputContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTscaiInputContainerWithDefaults() *TscaiInputContainer {
	this := TscaiInputContainer{}
	return &this
}

// GetPeriodicity returns the Periodicity field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetPeriodicity() int32 {
	if o == nil || IsNil(o.Periodicity) {
		var ret int32
		return ret
	}
	return *o.Periodicity
}

// GetPeriodicityOk returns a tuple with the Periodicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetPeriodicityOk() (*int32, bool) {
	if o == nil || IsNil(o.Periodicity) {
		return nil, false
	}
	return o.Periodicity, true
}

// HasPeriodicity returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasPeriodicity() bool {
	if o != nil && !IsNil(o.Periodicity) {
		return true
	}

	return false
}

// SetPeriodicity gets a reference to the given int32 and assigns it to the Periodicity field.
func (o *TscaiInputContainer) SetPeriodicity(v int32) {
	o.Periodicity = &v
}

// GetBurstArrivalTime returns the BurstArrivalTime field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetBurstArrivalTime() time.Time {
	if o == nil || IsNil(o.BurstArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.BurstArrivalTime
}

// GetBurstArrivalTimeOk returns a tuple with the BurstArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetBurstArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BurstArrivalTime) {
		return nil, false
	}
	return o.BurstArrivalTime, true
}

// HasBurstArrivalTime returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasBurstArrivalTime() bool {
	if o != nil && !IsNil(o.BurstArrivalTime) {
		return true
	}

	return false
}

// SetBurstArrivalTime gets a reference to the given time.Time and assigns it to the BurstArrivalTime field.
func (o *TscaiInputContainer) SetBurstArrivalTime(v time.Time) {
	o.BurstArrivalTime = &v
}

// GetSurTimeInNumMsg returns the SurTimeInNumMsg field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetSurTimeInNumMsg() int32 {
	if o == nil || IsNil(o.SurTimeInNumMsg) {
		var ret int32
		return ret
	}
	return *o.SurTimeInNumMsg
}

// GetSurTimeInNumMsgOk returns a tuple with the SurTimeInNumMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetSurTimeInNumMsgOk() (*int32, bool) {
	if o == nil || IsNil(o.SurTimeInNumMsg) {
		return nil, false
	}
	return o.SurTimeInNumMsg, true
}

// HasSurTimeInNumMsg returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasSurTimeInNumMsg() bool {
	if o != nil && !IsNil(o.SurTimeInNumMsg) {
		return true
	}

	return false
}

// SetSurTimeInNumMsg gets a reference to the given int32 and assigns it to the SurTimeInNumMsg field.
func (o *TscaiInputContainer) SetSurTimeInNumMsg(v int32) {
	o.SurTimeInNumMsg = &v
}

// GetSurTimeInTime returns the SurTimeInTime field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetSurTimeInTime() int32 {
	if o == nil || IsNil(o.SurTimeInTime) {
		var ret int32
		return ret
	}
	return *o.SurTimeInTime
}

// GetSurTimeInTimeOk returns a tuple with the SurTimeInTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetSurTimeInTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.SurTimeInTime) {
		return nil, false
	}
	return o.SurTimeInTime, true
}

// HasSurTimeInTime returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasSurTimeInTime() bool {
	if o != nil && !IsNil(o.SurTimeInTime) {
		return true
	}

	return false
}

// SetSurTimeInTime gets a reference to the given int32 and assigns it to the SurTimeInTime field.
func (o *TscaiInputContainer) SetSurTimeInTime(v int32) {
	o.SurTimeInTime = &v
}

// GetBurstArrivalTimeWnd returns the BurstArrivalTimeWnd field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetBurstArrivalTimeWnd() TimeWindow {
	if o == nil || IsNil(o.BurstArrivalTimeWnd) {
		var ret TimeWindow
		return ret
	}
	return *o.BurstArrivalTimeWnd
}

// GetBurstArrivalTimeWndOk returns a tuple with the BurstArrivalTimeWnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetBurstArrivalTimeWndOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.BurstArrivalTimeWnd) {
		return nil, false
	}
	return o.BurstArrivalTimeWnd, true
}

// HasBurstArrivalTimeWnd returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasBurstArrivalTimeWnd() bool {
	if o != nil && !IsNil(o.BurstArrivalTimeWnd) {
		return true
	}

	return false
}

// SetBurstArrivalTimeWnd gets a reference to the given TimeWindow and assigns it to the BurstArrivalTimeWnd field.
func (o *TscaiInputContainer) SetBurstArrivalTimeWnd(v TimeWindow) {
	o.BurstArrivalTimeWnd = &v
}

// GetPeriodicityRange returns the PeriodicityRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TscaiInputContainer) GetPeriodicityRange() PeriodicityRange {
	if o == nil || IsNil(o.PeriodicityRange.Get()) {
		var ret PeriodicityRange
		return ret
	}
	return *o.PeriodicityRange.Get()
}

// GetPeriodicityRangeOk returns a tuple with the PeriodicityRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TscaiInputContainer) GetPeriodicityRangeOk() (*PeriodicityRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodicityRange.Get(), o.PeriodicityRange.IsSet()
}

// HasPeriodicityRange returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasPeriodicityRange() bool {
	if o != nil && o.PeriodicityRange.IsSet() {
		return true
	}

	return false
}

// SetPeriodicityRange gets a reference to the given NullablePeriodicityRange and assigns it to the PeriodicityRange field.
func (o *TscaiInputContainer) SetPeriodicityRange(v PeriodicityRange) {
	o.PeriodicityRange.Set(&v)
}

// SetPeriodicityRangeNil sets the value for PeriodicityRange to be an explicit nil
func (o *TscaiInputContainer) SetPeriodicityRangeNil() {
	o.PeriodicityRange.Set(nil)
}

// UnsetPeriodicityRange ensures that no value is present for PeriodicityRange, not even an explicit nil
func (o *TscaiInputContainer) UnsetPeriodicityRange() {
	o.PeriodicityRange.Unset()
}

func (o TscaiInputContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TscaiInputContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Periodicity) {
		toSerialize["periodicity"] = o.Periodicity
	}
	if !IsNil(o.BurstArrivalTime) {
		toSerialize["burstArrivalTime"] = o.BurstArrivalTime
	}
	if !IsNil(o.SurTimeInNumMsg) {
		toSerialize["surTimeInNumMsg"] = o.SurTimeInNumMsg
	}
	if !IsNil(o.SurTimeInTime) {
		toSerialize["surTimeInTime"] = o.SurTimeInTime
	}
	if !IsNil(o.BurstArrivalTimeWnd) {
		toSerialize["burstArrivalTimeWnd"] = o.BurstArrivalTimeWnd
	}
	if o.PeriodicityRange.IsSet() {
		toSerialize["periodicityRange"] = o.PeriodicityRange.Get()
	}
	return toSerialize, nil
}

type NullableTscaiInputContainer struct {
	value *TscaiInputContainer
	isSet bool
}

func (v NullableTscaiInputContainer) Get() *TscaiInputContainer {
	return v.value
}

func (v *NullableTscaiInputContainer) Set(val *TscaiInputContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableTscaiInputContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableTscaiInputContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscaiInputContainer(val *TscaiInputContainer) *NullableTscaiInputContainer {
	return &NullableTscaiInputContainer{value: val, isSet: true}
}

func (v NullableTscaiInputContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscaiInputContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
