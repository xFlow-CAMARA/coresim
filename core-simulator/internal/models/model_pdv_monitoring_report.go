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

// checks if the PdvMonitoringReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PdvMonitoringReport{}

// PdvMonitoringReport Packet Delay Variation reporting information.
type PdvMonitoringReport struct {
	// Identification of the flows.
	Flows []Flows `json:"flows,omitempty"`
	// Uplink packet delay variation in units of milliseconds.
	UlPdv *int32 `json:"ulPdv,omitempty"`
	// Downlink packet delay variation in units of milliseconds.
	DlPdv *int32 `json:"dlPdv,omitempty"`
	// Round trip packet delay variation in units of milliseconds.
	RtPdv *int32 `json:"rtPdv,omitempty"`
}

// NewPdvMonitoringReport instantiates a new PdvMonitoringReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdvMonitoringReport() *PdvMonitoringReport {
	this := PdvMonitoringReport{}
	return &this
}

// NewPdvMonitoringReportWithDefaults instantiates a new PdvMonitoringReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdvMonitoringReportWithDefaults() *PdvMonitoringReport {
	this := PdvMonitoringReport{}
	return &this
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *PdvMonitoringReport) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdvMonitoringReport) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *PdvMonitoringReport) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *PdvMonitoringReport) SetFlows(v []Flows) {
	o.Flows = v
}

// GetUlPdv returns the UlPdv field value if set, zero value otherwise.
func (o *PdvMonitoringReport) GetUlPdv() int32 {
	if o == nil || IsNil(o.UlPdv) {
		var ret int32
		return ret
	}
	return *o.UlPdv
}

// GetUlPdvOk returns a tuple with the UlPdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdvMonitoringReport) GetUlPdvOk() (*int32, bool) {
	if o == nil || IsNil(o.UlPdv) {
		return nil, false
	}
	return o.UlPdv, true
}

// HasUlPdv returns a boolean if a field has been set.
func (o *PdvMonitoringReport) HasUlPdv() bool {
	if o != nil && !IsNil(o.UlPdv) {
		return true
	}

	return false
}

// SetUlPdv gets a reference to the given int32 and assigns it to the UlPdv field.
func (o *PdvMonitoringReport) SetUlPdv(v int32) {
	o.UlPdv = &v
}

// GetDlPdv returns the DlPdv field value if set, zero value otherwise.
func (o *PdvMonitoringReport) GetDlPdv() int32 {
	if o == nil || IsNil(o.DlPdv) {
		var ret int32
		return ret
	}
	return *o.DlPdv
}

// GetDlPdvOk returns a tuple with the DlPdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdvMonitoringReport) GetDlPdvOk() (*int32, bool) {
	if o == nil || IsNil(o.DlPdv) {
		return nil, false
	}
	return o.DlPdv, true
}

// HasDlPdv returns a boolean if a field has been set.
func (o *PdvMonitoringReport) HasDlPdv() bool {
	if o != nil && !IsNil(o.DlPdv) {
		return true
	}

	return false
}

// SetDlPdv gets a reference to the given int32 and assigns it to the DlPdv field.
func (o *PdvMonitoringReport) SetDlPdv(v int32) {
	o.DlPdv = &v
}

// GetRtPdv returns the RtPdv field value if set, zero value otherwise.
func (o *PdvMonitoringReport) GetRtPdv() int32 {
	if o == nil || IsNil(o.RtPdv) {
		var ret int32
		return ret
	}
	return *o.RtPdv
}

// GetRtPdvOk returns a tuple with the RtPdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdvMonitoringReport) GetRtPdvOk() (*int32, bool) {
	if o == nil || IsNil(o.RtPdv) {
		return nil, false
	}
	return o.RtPdv, true
}

// HasRtPdv returns a boolean if a field has been set.
func (o *PdvMonitoringReport) HasRtPdv() bool {
	if o != nil && !IsNil(o.RtPdv) {
		return true
	}

	return false
}

// SetRtPdv gets a reference to the given int32 and assigns it to the RtPdv field.
func (o *PdvMonitoringReport) SetRtPdv(v int32) {
	o.RtPdv = &v
}

func (o PdvMonitoringReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PdvMonitoringReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !IsNil(o.UlPdv) {
		toSerialize["ulPdv"] = o.UlPdv
	}
	if !IsNil(o.DlPdv) {
		toSerialize["dlPdv"] = o.DlPdv
	}
	if !IsNil(o.RtPdv) {
		toSerialize["rtPdv"] = o.RtPdv
	}
	return toSerialize, nil
}

type NullablePdvMonitoringReport struct {
	value *PdvMonitoringReport
	isSet bool
}

func (v NullablePdvMonitoringReport) Get() *PdvMonitoringReport {
	return v.value
}

func (v *NullablePdvMonitoringReport) Set(val *PdvMonitoringReport) {
	v.value = val
	v.isSet = true
}

func (v NullablePdvMonitoringReport) IsSet() bool {
	return v.isSet
}

func (v *NullablePdvMonitoringReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdvMonitoringReport(val *PdvMonitoringReport) *NullablePdvMonitoringReport {
	return &NullablePdvMonitoringReport{value: val, isSet: true}
}

func (v NullablePdvMonitoringReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdvMonitoringReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
