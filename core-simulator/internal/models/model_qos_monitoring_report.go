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

// checks if the QosMonitoringReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosMonitoringReport{}

// QosMonitoringReport QoS Monitoring reporting information.
type QosMonitoringReport struct {
	Flows    []Flows `json:"flows,omitempty"`
	UlDelays []int32 `json:"ulDelays,omitempty"`
	DlDelays []int32 `json:"dlDelays,omitempty"`
	RtDelays []int32 `json:"rtDelays,omitempty"`
	// Represents the packet delay measurement failure indicator.
	Pdmf      *bool   `json:"pdmf,omitempty"`
	UlConInfo []int32 `json:"ulConInfo,omitempty"`
	DlConInfo []int32 `json:"dlConInfo,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	UlDataRate *string `json:"ulDataRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	DlDataRate *string `json:"dlDataRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
}

// NewQosMonitoringReport instantiates a new QosMonitoringReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringReport() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// NewQosMonitoringReportWithDefaults instantiates a new QosMonitoringReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringReportWithDefaults() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *QosMonitoringReport) SetFlows(v []Flows) {
	o.Flows = v
}

// GetUlDelays returns the UlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetUlDelays() []int32 {
	if o == nil || IsNil(o.UlDelays) {
		var ret []int32
		return ret
	}
	return o.UlDelays
}

// GetUlDelaysOk returns a tuple with the UlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetUlDelaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.UlDelays) {
		return nil, false
	}
	return o.UlDelays, true
}

// HasUlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasUlDelays() bool {
	if o != nil && !IsNil(o.UlDelays) {
		return true
	}

	return false
}

// SetUlDelays gets a reference to the given []int32 and assigns it to the UlDelays field.
func (o *QosMonitoringReport) SetUlDelays(v []int32) {
	o.UlDelays = v
}

// GetDlDelays returns the DlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetDlDelays() []int32 {
	if o == nil || IsNil(o.DlDelays) {
		var ret []int32
		return ret
	}
	return o.DlDelays
}

// GetDlDelaysOk returns a tuple with the DlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetDlDelaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.DlDelays) {
		return nil, false
	}
	return o.DlDelays, true
}

// HasDlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasDlDelays() bool {
	if o != nil && !IsNil(o.DlDelays) {
		return true
	}

	return false
}

// SetDlDelays gets a reference to the given []int32 and assigns it to the DlDelays field.
func (o *QosMonitoringReport) SetDlDelays(v []int32) {
	o.DlDelays = v
}

// GetRtDelays returns the RtDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetRtDelays() []int32 {
	if o == nil || IsNil(o.RtDelays) {
		var ret []int32
		return ret
	}
	return o.RtDelays
}

// GetRtDelaysOk returns a tuple with the RtDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetRtDelaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.RtDelays) {
		return nil, false
	}
	return o.RtDelays, true
}

// HasRtDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasRtDelays() bool {
	if o != nil && !IsNil(o.RtDelays) {
		return true
	}

	return false
}

// SetRtDelays gets a reference to the given []int32 and assigns it to the RtDelays field.
func (o *QosMonitoringReport) SetRtDelays(v []int32) {
	o.RtDelays = v
}

// GetPdmf returns the Pdmf field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetPdmf() bool {
	if o == nil || IsNil(o.Pdmf) {
		var ret bool
		return ret
	}
	return *o.Pdmf
}

// GetPdmfOk returns a tuple with the Pdmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetPdmfOk() (*bool, bool) {
	if o == nil || IsNil(o.Pdmf) {
		return nil, false
	}
	return o.Pdmf, true
}

// HasPdmf returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasPdmf() bool {
	if o != nil && !IsNil(o.Pdmf) {
		return true
	}

	return false
}

// SetPdmf gets a reference to the given bool and assigns it to the Pdmf field.
func (o *QosMonitoringReport) SetPdmf(v bool) {
	o.Pdmf = &v
}

// GetUlConInfo returns the UlConInfo field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetUlConInfo() []int32 {
	if o == nil || IsNil(o.UlConInfo) {
		var ret []int32
		return ret
	}
	return o.UlConInfo
}

// GetUlConInfoOk returns a tuple with the UlConInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetUlConInfoOk() ([]int32, bool) {
	if o == nil || IsNil(o.UlConInfo) {
		return nil, false
	}
	return o.UlConInfo, true
}

// HasUlConInfo returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasUlConInfo() bool {
	if o != nil && !IsNil(o.UlConInfo) {
		return true
	}

	return false
}

// SetUlConInfo gets a reference to the given []int32 and assigns it to the UlConInfo field.
func (o *QosMonitoringReport) SetUlConInfo(v []int32) {
	o.UlConInfo = v
}

// GetDlConInfo returns the DlConInfo field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetDlConInfo() []int32 {
	if o == nil || IsNil(o.DlConInfo) {
		var ret []int32
		return ret
	}
	return o.DlConInfo
}

// GetDlConInfoOk returns a tuple with the DlConInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetDlConInfoOk() ([]int32, bool) {
	if o == nil || IsNil(o.DlConInfo) {
		return nil, false
	}
	return o.DlConInfo, true
}

// HasDlConInfo returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasDlConInfo() bool {
	if o != nil && !IsNil(o.DlConInfo) {
		return true
	}

	return false
}

// SetDlConInfo gets a reference to the given []int32 and assigns it to the DlConInfo field.
func (o *QosMonitoringReport) SetDlConInfo(v []int32) {
	o.DlConInfo = v
}

// GetUlDataRate returns the UlDataRate field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetUlDataRate() string {
	if o == nil || IsNil(o.UlDataRate) {
		var ret string
		return ret
	}
	return *o.UlDataRate
}

// GetUlDataRateOk returns a tuple with the UlDataRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetUlDataRateOk() (*string, bool) {
	if o == nil || IsNil(o.UlDataRate) {
		return nil, false
	}
	return o.UlDataRate, true
}

// HasUlDataRate returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasUlDataRate() bool {
	if o != nil && !IsNil(o.UlDataRate) {
		return true
	}

	return false
}

// SetUlDataRate gets a reference to the given string and assigns it to the UlDataRate field.
func (o *QosMonitoringReport) SetUlDataRate(v string) {
	o.UlDataRate = &v
}

// GetDlDataRate returns the DlDataRate field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetDlDataRate() string {
	if o == nil || IsNil(o.DlDataRate) {
		var ret string
		return ret
	}
	return *o.DlDataRate
}

// GetDlDataRateOk returns a tuple with the DlDataRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetDlDataRateOk() (*string, bool) {
	if o == nil || IsNil(o.DlDataRate) {
		return nil, false
	}
	return o.DlDataRate, true
}

// HasDlDataRate returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasDlDataRate() bool {
	if o != nil && !IsNil(o.DlDataRate) {
		return true
	}

	return false
}

// SetDlDataRate gets a reference to the given string and assigns it to the DlDataRate field.
func (o *QosMonitoringReport) SetDlDataRate(v string) {
	o.DlDataRate = &v
}

func (o QosMonitoringReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosMonitoringReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !IsNil(o.UlDelays) {
		toSerialize["ulDelays"] = o.UlDelays
	}
	if !IsNil(o.DlDelays) {
		toSerialize["dlDelays"] = o.DlDelays
	}
	if !IsNil(o.RtDelays) {
		toSerialize["rtDelays"] = o.RtDelays
	}
	if !IsNil(o.Pdmf) {
		toSerialize["pdmf"] = o.Pdmf
	}
	if !IsNil(o.UlConInfo) {
		toSerialize["ulConInfo"] = o.UlConInfo
	}
	if !IsNil(o.DlConInfo) {
		toSerialize["dlConInfo"] = o.DlConInfo
	}
	if !IsNil(o.UlDataRate) {
		toSerialize["ulDataRate"] = o.UlDataRate
	}
	if !IsNil(o.DlDataRate) {
		toSerialize["dlDataRate"] = o.DlDataRate
	}
	return toSerialize, nil
}

type NullableQosMonitoringReport struct {
	value *QosMonitoringReport
	isSet bool
}

func (v NullableQosMonitoringReport) Get() *QosMonitoringReport {
	return v.value
}

func (v *NullableQosMonitoringReport) Set(val *QosMonitoringReport) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringReport) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringReport(val *QosMonitoringReport) *NullableQosMonitoringReport {
	return &NullableQosMonitoringReport{value: val, isSet: true}
}

func (v NullableQosMonitoringReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
