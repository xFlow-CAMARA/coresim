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

// checks if the PduSetQosPara type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSetQosPara{}

// PduSetQosPara Represents the PDU Set level QoS parameters.
type PduSetQosPara struct {
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501 [8])), expressed in 0.01 milliseconds.
	PduSetDelayBudget *int32 `json:"pduSetDelayBudget,omitempty"`
	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	PduSetErrRate      *string             `json:"pduSetErrRate,omitempty" validate:"regexp=^([0-9]E-[0-9])$"`
	PduSetHandlingInfo *PduSetHandlingInfo `json:"pduSetHandlingInfo,omitempty"`
}

// NewPduSetQosPara instantiates a new PduSetQosPara object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSetQosPara() *PduSetQosPara {
	this := PduSetQosPara{}
	return &this
}

// NewPduSetQosParaWithDefaults instantiates a new PduSetQosPara object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSetQosParaWithDefaults() *PduSetQosPara {
	this := PduSetQosPara{}
	return &this
}

// GetPduSetDelayBudget returns the PduSetDelayBudget field value if set, zero value otherwise.
func (o *PduSetQosPara) GetPduSetDelayBudget() int32 {
	if o == nil || IsNil(o.PduSetDelayBudget) {
		var ret int32
		return ret
	}
	return *o.PduSetDelayBudget
}

// GetPduSetDelayBudgetOk returns a tuple with the PduSetDelayBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSetQosPara) GetPduSetDelayBudgetOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSetDelayBudget) {
		return nil, false
	}
	return o.PduSetDelayBudget, true
}

// HasPduSetDelayBudget returns a boolean if a field has been set.
func (o *PduSetQosPara) HasPduSetDelayBudget() bool {
	if o != nil && !IsNil(o.PduSetDelayBudget) {
		return true
	}

	return false
}

// SetPduSetDelayBudget gets a reference to the given int32 and assigns it to the PduSetDelayBudget field.
func (o *PduSetQosPara) SetPduSetDelayBudget(v int32) {
	o.PduSetDelayBudget = &v
}

// GetPduSetErrRate returns the PduSetErrRate field value if set, zero value otherwise.
func (o *PduSetQosPara) GetPduSetErrRate() string {
	if o == nil || IsNil(o.PduSetErrRate) {
		var ret string
		return ret
	}
	return *o.PduSetErrRate
}

// GetPduSetErrRateOk returns a tuple with the PduSetErrRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSetQosPara) GetPduSetErrRateOk() (*string, bool) {
	if o == nil || IsNil(o.PduSetErrRate) {
		return nil, false
	}
	return o.PduSetErrRate, true
}

// HasPduSetErrRate returns a boolean if a field has been set.
func (o *PduSetQosPara) HasPduSetErrRate() bool {
	if o != nil && !IsNil(o.PduSetErrRate) {
		return true
	}

	return false
}

// SetPduSetErrRate gets a reference to the given string and assigns it to the PduSetErrRate field.
func (o *PduSetQosPara) SetPduSetErrRate(v string) {
	o.PduSetErrRate = &v
}

// GetPduSetHandlingInfo returns the PduSetHandlingInfo field value if set, zero value otherwise.
func (o *PduSetQosPara) GetPduSetHandlingInfo() PduSetHandlingInfo {
	if o == nil || IsNil(o.PduSetHandlingInfo) {
		var ret PduSetHandlingInfo
		return ret
	}
	return *o.PduSetHandlingInfo
}

// GetPduSetHandlingInfoOk returns a tuple with the PduSetHandlingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSetQosPara) GetPduSetHandlingInfoOk() (*PduSetHandlingInfo, bool) {
	if o == nil || IsNil(o.PduSetHandlingInfo) {
		return nil, false
	}
	return o.PduSetHandlingInfo, true
}

// HasPduSetHandlingInfo returns a boolean if a field has been set.
func (o *PduSetQosPara) HasPduSetHandlingInfo() bool {
	if o != nil && !IsNil(o.PduSetHandlingInfo) {
		return true
	}

	return false
}

// SetPduSetHandlingInfo gets a reference to the given PduSetHandlingInfo and assigns it to the PduSetHandlingInfo field.
func (o *PduSetQosPara) SetPduSetHandlingInfo(v PduSetHandlingInfo) {
	o.PduSetHandlingInfo = &v
}

func (o PduSetQosPara) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSetQosPara) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PduSetDelayBudget) {
		toSerialize["pduSetDelayBudget"] = o.PduSetDelayBudget
	}
	if !IsNil(o.PduSetErrRate) {
		toSerialize["pduSetErrRate"] = o.PduSetErrRate
	}
	if !IsNil(o.PduSetHandlingInfo) {
		toSerialize["pduSetHandlingInfo"] = o.PduSetHandlingInfo
	}
	return toSerialize, nil
}

type NullablePduSetQosPara struct {
	value *PduSetQosPara
	isSet bool
}

func (v NullablePduSetQosPara) Get() *PduSetQosPara {
	return v.value
}

func (v *NullablePduSetQosPara) Set(val *PduSetQosPara) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSetQosPara) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSetQosPara) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSetQosPara(val *PduSetQosPara) *NullablePduSetQosPara {
	return &NullablePduSetQosPara{value: val, isSet: true}
}

func (v NullablePduSetQosPara) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSetQosPara) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
