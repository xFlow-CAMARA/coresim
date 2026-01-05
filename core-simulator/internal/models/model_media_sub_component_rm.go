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

// checks if the MediaSubComponentRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaSubComponentRm{}

// MediaSubComponentRm This data type is defined in the same way as the MediaSubComponent data type, but with the OpenAPI nullable property set to true. Removable attributes marBwDl and marBwUl are defined with the corresponding removable data type.
type MediaSubComponentRm struct {
	AfSigProtocol NullableAfSigProtocol `json:"afSigProtocol,omitempty"`
	EthfDescs     []EthFlowDescription  `json:"ethfDescs,omitempty"`
	FNum          int32                 `json:"fNum"`
	FDescs        []string              `json:"fDescs,omitempty"`
	// Represents additional flow description information (flow label and IPsec SPI) per Uplink and/or Downlink IP flows.
	AddInfoFlowDescs []AddFlowDescriptionInfo `json:"addInfoFlowDescs,omitempty"`
	FStatus          *FlowStatus              `json:"fStatus,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwDl NullableString `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwUl NullableString `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// This data type is defined in the same way as the TosTrafficClass data type, but with the OpenAPI nullable property set to true.
	TosTrCl   NullableString               `json:"tosTrCl,omitempty"`
	FlowUsage *FlowUsage                   `json:"flowUsage,omitempty"`
	EvSubsc   NullableEventsSubscReqDataRm `json:"evSubsc,omitempty"`
}

type _MediaSubComponentRm MediaSubComponentRm

// NewMediaSubComponentRm instantiates a new MediaSubComponentRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaSubComponentRm(fNum int32) *MediaSubComponentRm {
	this := MediaSubComponentRm{}
	this.FNum = fNum
	return &this
}

// NewMediaSubComponentRmWithDefaults instantiates a new MediaSubComponentRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaSubComponentRmWithDefaults() *MediaSubComponentRm {
	this := MediaSubComponentRm{}
	return &this
}

// GetAfSigProtocol returns the AfSigProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetAfSigProtocol() AfSigProtocol {
	if o == nil || IsNil(o.AfSigProtocol.Get()) {
		var ret AfSigProtocol
		return ret
	}
	return *o.AfSigProtocol.Get()
}

// GetAfSigProtocolOk returns a tuple with the AfSigProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetAfSigProtocolOk() (*AfSigProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSigProtocol.Get(), o.AfSigProtocol.IsSet()
}

// HasAfSigProtocol returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasAfSigProtocol() bool {
	if o != nil && o.AfSigProtocol.IsSet() {
		return true
	}

	return false
}

// SetAfSigProtocol gets a reference to the given NullableAfSigProtocol and assigns it to the AfSigProtocol field.
func (o *MediaSubComponentRm) SetAfSigProtocol(v AfSigProtocol) {
	o.AfSigProtocol.Set(&v)
}

// SetAfSigProtocolNil sets the value for AfSigProtocol to be an explicit nil
func (o *MediaSubComponentRm) SetAfSigProtocolNil() {
	o.AfSigProtocol.Set(nil)
}

// UnsetAfSigProtocol ensures that no value is present for AfSigProtocol, not even an explicit nil
func (o *MediaSubComponentRm) UnsetAfSigProtocol() {
	o.AfSigProtocol.Unset()
}

// GetEthfDescs returns the EthfDescs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetEthfDescs() []EthFlowDescription {
	if o == nil {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthfDescs
}

// GetEthfDescsOk returns a tuple with the EthfDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetEthfDescsOk() ([]EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthfDescs) {
		return nil, false
	}
	return o.EthfDescs, true
}

// HasEthfDescs returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasEthfDescs() bool {
	if o != nil && !IsNil(o.EthfDescs) {
		return true
	}

	return false
}

// SetEthfDescs gets a reference to the given []EthFlowDescription and assigns it to the EthfDescs field.
func (o *MediaSubComponentRm) SetEthfDescs(v []EthFlowDescription) {
	o.EthfDescs = v
}

// GetFNum returns the FNum field value
func (o *MediaSubComponentRm) GetFNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FNum
}

// GetFNumOk returns a tuple with the FNum field value
// and a boolean to check if the value has been set.
func (o *MediaSubComponentRm) GetFNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FNum, true
}

// SetFNum sets field value
func (o *MediaSubComponentRm) SetFNum(v int32) {
	o.FNum = v
}

// GetFDescs returns the FDescs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetFDescs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FDescs
}

// GetFDescsOk returns a tuple with the FDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetFDescsOk() ([]string, bool) {
	if o == nil || IsNil(o.FDescs) {
		return nil, false
	}
	return o.FDescs, true
}

// HasFDescs returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasFDescs() bool {
	if o != nil && !IsNil(o.FDescs) {
		return true
	}

	return false
}

// SetFDescs gets a reference to the given []string and assigns it to the FDescs field.
func (o *MediaSubComponentRm) SetFDescs(v []string) {
	o.FDescs = v
}

// GetAddInfoFlowDescs returns the AddInfoFlowDescs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetAddInfoFlowDescs() []AddFlowDescriptionInfo {
	if o == nil {
		var ret []AddFlowDescriptionInfo
		return ret
	}
	return o.AddInfoFlowDescs
}

// GetAddInfoFlowDescsOk returns a tuple with the AddInfoFlowDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetAddInfoFlowDescsOk() ([]AddFlowDescriptionInfo, bool) {
	if o == nil || IsNil(o.AddInfoFlowDescs) {
		return nil, false
	}
	return o.AddInfoFlowDescs, true
}

// HasAddInfoFlowDescs returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasAddInfoFlowDescs() bool {
	if o != nil && !IsNil(o.AddInfoFlowDescs) {
		return true
	}

	return false
}

// SetAddInfoFlowDescs gets a reference to the given []AddFlowDescriptionInfo and assigns it to the AddInfoFlowDescs field.
func (o *MediaSubComponentRm) SetAddInfoFlowDescs(v []AddFlowDescriptionInfo) {
	o.AddInfoFlowDescs = v
}

// GetFStatus returns the FStatus field value if set, zero value otherwise.
func (o *MediaSubComponentRm) GetFStatus() FlowStatus {
	if o == nil || IsNil(o.FStatus) {
		var ret FlowStatus
		return ret
	}
	return *o.FStatus
}

// GetFStatusOk returns a tuple with the FStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponentRm) GetFStatusOk() (*FlowStatus, bool) {
	if o == nil || IsNil(o.FStatus) {
		return nil, false
	}
	return o.FStatus, true
}

// HasFStatus returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasFStatus() bool {
	if o != nil && !IsNil(o.FStatus) {
		return true
	}

	return false
}

// SetFStatus gets a reference to the given FlowStatus and assigns it to the FStatus field.
func (o *MediaSubComponentRm) SetFStatus(v FlowStatus) {
	o.FStatus = &v
}

// GetMarBwDl returns the MarBwDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetMarBwDl() string {
	if o == nil || IsNil(o.MarBwDl.Get()) {
		var ret string
		return ret
	}
	return *o.MarBwDl.Get()
}

// GetMarBwDlOk returns a tuple with the MarBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetMarBwDlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarBwDl.Get(), o.MarBwDl.IsSet()
}

// HasMarBwDl returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasMarBwDl() bool {
	if o != nil && o.MarBwDl.IsSet() {
		return true
	}

	return false
}

// SetMarBwDl gets a reference to the given NullableString and assigns it to the MarBwDl field.
func (o *MediaSubComponentRm) SetMarBwDl(v string) {
	o.MarBwDl.Set(&v)
}

// SetMarBwDlNil sets the value for MarBwDl to be an explicit nil
func (o *MediaSubComponentRm) SetMarBwDlNil() {
	o.MarBwDl.Set(nil)
}

// UnsetMarBwDl ensures that no value is present for MarBwDl, not even an explicit nil
func (o *MediaSubComponentRm) UnsetMarBwDl() {
	o.MarBwDl.Unset()
}

// GetMarBwUl returns the MarBwUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetMarBwUl() string {
	if o == nil || IsNil(o.MarBwUl.Get()) {
		var ret string
		return ret
	}
	return *o.MarBwUl.Get()
}

// GetMarBwUlOk returns a tuple with the MarBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetMarBwUlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarBwUl.Get(), o.MarBwUl.IsSet()
}

// HasMarBwUl returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasMarBwUl() bool {
	if o != nil && o.MarBwUl.IsSet() {
		return true
	}

	return false
}

// SetMarBwUl gets a reference to the given NullableString and assigns it to the MarBwUl field.
func (o *MediaSubComponentRm) SetMarBwUl(v string) {
	o.MarBwUl.Set(&v)
}

// SetMarBwUlNil sets the value for MarBwUl to be an explicit nil
func (o *MediaSubComponentRm) SetMarBwUlNil() {
	o.MarBwUl.Set(nil)
}

// UnsetMarBwUl ensures that no value is present for MarBwUl, not even an explicit nil
func (o *MediaSubComponentRm) UnsetMarBwUl() {
	o.MarBwUl.Unset()
}

// GetTosTrCl returns the TosTrCl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetTosTrCl() string {
	if o == nil || IsNil(o.TosTrCl.Get()) {
		var ret string
		return ret
	}
	return *o.TosTrCl.Get()
}

// GetTosTrClOk returns a tuple with the TosTrCl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetTosTrClOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TosTrCl.Get(), o.TosTrCl.IsSet()
}

// HasTosTrCl returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasTosTrCl() bool {
	if o != nil && o.TosTrCl.IsSet() {
		return true
	}

	return false
}

// SetTosTrCl gets a reference to the given NullableString and assigns it to the TosTrCl field.
func (o *MediaSubComponentRm) SetTosTrCl(v string) {
	o.TosTrCl.Set(&v)
}

// SetTosTrClNil sets the value for TosTrCl to be an explicit nil
func (o *MediaSubComponentRm) SetTosTrClNil() {
	o.TosTrCl.Set(nil)
}

// UnsetTosTrCl ensures that no value is present for TosTrCl, not even an explicit nil
func (o *MediaSubComponentRm) UnsetTosTrCl() {
	o.TosTrCl.Unset()
}

// GetFlowUsage returns the FlowUsage field value if set, zero value otherwise.
func (o *MediaSubComponentRm) GetFlowUsage() FlowUsage {
	if o == nil || IsNil(o.FlowUsage) {
		var ret FlowUsage
		return ret
	}
	return *o.FlowUsage
}

// GetFlowUsageOk returns a tuple with the FlowUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponentRm) GetFlowUsageOk() (*FlowUsage, bool) {
	if o == nil || IsNil(o.FlowUsage) {
		return nil, false
	}
	return o.FlowUsage, true
}

// HasFlowUsage returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasFlowUsage() bool {
	if o != nil && !IsNil(o.FlowUsage) {
		return true
	}

	return false
}

// SetFlowUsage gets a reference to the given FlowUsage and assigns it to the FlowUsage field.
func (o *MediaSubComponentRm) SetFlowUsage(v FlowUsage) {
	o.FlowUsage = &v
}

// GetEvSubsc returns the EvSubsc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponentRm) GetEvSubsc() EventsSubscReqDataRm {
	if o == nil || IsNil(o.EvSubsc.Get()) {
		var ret EventsSubscReqDataRm
		return ret
	}
	return *o.EvSubsc.Get()
}

// GetEvSubscOk returns a tuple with the EvSubsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponentRm) GetEvSubscOk() (*EventsSubscReqDataRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvSubsc.Get(), o.EvSubsc.IsSet()
}

// HasEvSubsc returns a boolean if a field has been set.
func (o *MediaSubComponentRm) HasEvSubsc() bool {
	if o != nil && o.EvSubsc.IsSet() {
		return true
	}

	return false
}

// SetEvSubsc gets a reference to the given NullableEventsSubscReqDataRm and assigns it to the EvSubsc field.
func (o *MediaSubComponentRm) SetEvSubsc(v EventsSubscReqDataRm) {
	o.EvSubsc.Set(&v)
}

// SetEvSubscNil sets the value for EvSubsc to be an explicit nil
func (o *MediaSubComponentRm) SetEvSubscNil() {
	o.EvSubsc.Set(nil)
}

// UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil
func (o *MediaSubComponentRm) UnsetEvSubsc() {
	o.EvSubsc.Unset()
}

func (o MediaSubComponentRm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaSubComponentRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AfSigProtocol.IsSet() {
		toSerialize["afSigProtocol"] = o.AfSigProtocol.Get()
	}
	if o.EthfDescs != nil {
		toSerialize["ethfDescs"] = o.EthfDescs
	}
	toSerialize["fNum"] = o.FNum
	if o.FDescs != nil {
		toSerialize["fDescs"] = o.FDescs
	}
	if o.AddInfoFlowDescs != nil {
		toSerialize["addInfoFlowDescs"] = o.AddInfoFlowDescs
	}
	if !IsNil(o.FStatus) {
		toSerialize["fStatus"] = o.FStatus
	}
	if o.MarBwDl.IsSet() {
		toSerialize["marBwDl"] = o.MarBwDl.Get()
	}
	if o.MarBwUl.IsSet() {
		toSerialize["marBwUl"] = o.MarBwUl.Get()
	}
	if o.TosTrCl.IsSet() {
		toSerialize["tosTrCl"] = o.TosTrCl.Get()
	}
	if !IsNil(o.FlowUsage) {
		toSerialize["flowUsage"] = o.FlowUsage
	}
	if o.EvSubsc.IsSet() {
		toSerialize["evSubsc"] = o.EvSubsc.Get()
	}
	return toSerialize, nil
}

func (o *MediaSubComponentRm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fNum",
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

	varMediaSubComponentRm := _MediaSubComponentRm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaSubComponentRm)

	if err != nil {
		return err
	}

	*o = MediaSubComponentRm(varMediaSubComponentRm)

	return err
}

type NullableMediaSubComponentRm struct {
	value *MediaSubComponentRm
	isSet bool
}

func (v NullableMediaSubComponentRm) Get() *MediaSubComponentRm {
	return v.value
}

func (v *NullableMediaSubComponentRm) Set(val *MediaSubComponentRm) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSubComponentRm) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSubComponentRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSubComponentRm(val *MediaSubComponentRm) *NullableMediaSubComponentRm {
	return &NullableMediaSubComponentRm{value: val, isSet: true}
}

func (v NullableMediaSubComponentRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSubComponentRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
