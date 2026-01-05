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

// checks if the ResourcesAllocationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcesAllocationInfo{}

// ResourcesAllocationInfo Describes the status of the PCC rule(s) related to certain media components.
type ResourcesAllocationInfo struct {
	McResourcStatus *MediaComponentResourcesStatus `json:"mcResourcStatus,omitempty"`
	Flows           []Flows                        `json:"flows,omitempty"`
	// Indicates whether NG-RAN supports alternative QoS parameters. The default value false shall apply if the attribute is not present. It shall be set to false to indicate that the lowest priority alternative QoS profile could not be fulfilled.
	AltSerReq *string `json:"altSerReq,omitempty"`
}

// NewResourcesAllocationInfo instantiates a new ResourcesAllocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcesAllocationInfo() *ResourcesAllocationInfo {
	this := ResourcesAllocationInfo{}
	return &this
}

// NewResourcesAllocationInfoWithDefaults instantiates a new ResourcesAllocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesAllocationInfoWithDefaults() *ResourcesAllocationInfo {
	this := ResourcesAllocationInfo{}
	return &this
}

// GetMcResourcStatus returns the McResourcStatus field value if set, zero value otherwise.
func (o *ResourcesAllocationInfo) GetMcResourcStatus() MediaComponentResourcesStatus {
	if o == nil || IsNil(o.McResourcStatus) {
		var ret MediaComponentResourcesStatus
		return ret
	}
	return *o.McResourcStatus
}

// GetMcResourcStatusOk returns a tuple with the McResourcStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesAllocationInfo) GetMcResourcStatusOk() (*MediaComponentResourcesStatus, bool) {
	if o == nil || IsNil(o.McResourcStatus) {
		return nil, false
	}
	return o.McResourcStatus, true
}

// HasMcResourcStatus returns a boolean if a field has been set.
func (o *ResourcesAllocationInfo) HasMcResourcStatus() bool {
	if o != nil && !IsNil(o.McResourcStatus) {
		return true
	}

	return false
}

// SetMcResourcStatus gets a reference to the given MediaComponentResourcesStatus and assigns it to the McResourcStatus field.
func (o *ResourcesAllocationInfo) SetMcResourcStatus(v MediaComponentResourcesStatus) {
	o.McResourcStatus = &v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *ResourcesAllocationInfo) GetFlows() []Flows {
	if o == nil || IsNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesAllocationInfo) GetFlowsOk() ([]Flows, bool) {
	if o == nil || IsNil(o.Flows) {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *ResourcesAllocationInfo) HasFlows() bool {
	if o != nil && !IsNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *ResourcesAllocationInfo) SetFlows(v []Flows) {
	o.Flows = v
}

// GetAltSerReq returns the AltSerReq field value if set, zero value otherwise.
func (o *ResourcesAllocationInfo) GetAltSerReq() string {
	if o == nil || IsNil(o.AltSerReq) {
		var ret string
		return ret
	}
	return *o.AltSerReq
}

// GetAltSerReqOk returns a tuple with the AltSerReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcesAllocationInfo) GetAltSerReqOk() (*string, bool) {
	if o == nil || IsNil(o.AltSerReq) {
		return nil, false
	}
	return o.AltSerReq, true
}

// HasAltSerReq returns a boolean if a field has been set.
func (o *ResourcesAllocationInfo) HasAltSerReq() bool {
	if o != nil && !IsNil(o.AltSerReq) {
		return true
	}

	return false
}

// SetAltSerReq gets a reference to the given string and assigns it to the AltSerReq field.
func (o *ResourcesAllocationInfo) SetAltSerReq(v string) {
	o.AltSerReq = &v
}

func (o ResourcesAllocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcesAllocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.McResourcStatus) {
		toSerialize["mcResourcStatus"] = o.McResourcStatus
	}
	if !IsNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !IsNil(o.AltSerReq) {
		toSerialize["altSerReq"] = o.AltSerReq
	}
	return toSerialize, nil
}

type NullableResourcesAllocationInfo struct {
	value *ResourcesAllocationInfo
	isSet bool
}

func (v NullableResourcesAllocationInfo) Get() *ResourcesAllocationInfo {
	return v.value
}

func (v *NullableResourcesAllocationInfo) Set(val *ResourcesAllocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcesAllocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcesAllocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcesAllocationInfo(val *ResourcesAllocationInfo) *NullableResourcesAllocationInfo {
	return &NullableResourcesAllocationInfo{value: val, isSet: true}
}

func (v NullableResourcesAllocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcesAllocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
