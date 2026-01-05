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

// checks if the MediaSubComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaSubComponent{}

// MediaSubComponent Identifies a media subcomponent.
type MediaSubComponent struct {
	AfSigProtocol NullableAfSigProtocol `json:"afSigProtocol,omitempty"`
	EthfDescs     []EthFlowDescription  `json:"ethfDescs,omitempty"`
	FNum          int32                 `json:"fNum"`
	FDescs        []string              `json:"fDescs,omitempty"`
	// Represents additional flow description information (flow label and IPsec SPI) per Uplink and/or Downlink IP flows.
	AddInfoFlowDescs []AddFlowDescriptionInfo `json:"addInfoFlowDescs,omitempty"`
	FStatus          *FlowStatus              `json:"fStatus,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl *string `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl *string `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// 2-octet string, where each octet is encoded in hexadecimal representation. The first octet contains the IPv4 Type-of-Service or the IPv6 Traffic-Class field and the second octet contains the ToS/Traffic Class mask field.
	TosTrCl   *string             `json:"tosTrCl,omitempty"`
	FlowUsage *string             `json:"flowUsage,omitempty"`
	EvSubsc   *EventsSubscReqData `json:"evSubsc,omitempty"`
}

type _MediaSubComponent MediaSubComponent

// NewMediaSubComponent instantiates a new MediaSubComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaSubComponent(fNum int32) *MediaSubComponent {
	this := MediaSubComponent{}
	this.FNum = fNum
	return &this
}

// NewMediaSubComponentWithDefaults instantiates a new MediaSubComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaSubComponentWithDefaults() *MediaSubComponent {
	this := MediaSubComponent{}
	return &this
}

// GetAfSigProtocol returns the AfSigProtocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaSubComponent) GetAfSigProtocol() AfSigProtocol {
	if o == nil || IsNil(o.AfSigProtocol.Get()) {
		var ret AfSigProtocol
		return ret
	}
	return *o.AfSigProtocol.Get()
}

// GetAfSigProtocolOk returns a tuple with the AfSigProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaSubComponent) GetAfSigProtocolOk() (*AfSigProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSigProtocol.Get(), o.AfSigProtocol.IsSet()
}

// HasAfSigProtocol returns a boolean if a field has been set.
func (o *MediaSubComponent) HasAfSigProtocol() bool {
	if o != nil && o.AfSigProtocol.IsSet() {
		return true
	}

	return false
}

// SetAfSigProtocol gets a reference to the given NullableAfSigProtocol and assigns it to the AfSigProtocol field.
func (o *MediaSubComponent) SetAfSigProtocol(v AfSigProtocol) {
	o.AfSigProtocol.Set(&v)
}

// SetAfSigProtocolNil sets the value for AfSigProtocol to be an explicit nil
func (o *MediaSubComponent) SetAfSigProtocolNil() {
	o.AfSigProtocol.Set(nil)
}

// UnsetAfSigProtocol ensures that no value is present for AfSigProtocol, not even an explicit nil
func (o *MediaSubComponent) UnsetAfSigProtocol() {
	o.AfSigProtocol.Unset()
}

// GetEthfDescs returns the EthfDescs field value if set, zero value otherwise.
func (o *MediaSubComponent) GetEthfDescs() []EthFlowDescription {
	if o == nil || IsNil(o.EthfDescs) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthfDescs
}

// GetEthfDescsOk returns a tuple with the EthfDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetEthfDescsOk() ([]EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthfDescs) {
		return nil, false
	}
	return o.EthfDescs, true
}

// HasEthfDescs returns a boolean if a field has been set.
func (o *MediaSubComponent) HasEthfDescs() bool {
	if o != nil && !IsNil(o.EthfDescs) {
		return true
	}

	return false
}

// SetEthfDescs gets a reference to the given []EthFlowDescription and assigns it to the EthfDescs field.
func (o *MediaSubComponent) SetEthfDescs(v []EthFlowDescription) {
	o.EthfDescs = v
}

// GetFNum returns the FNum field value
func (o *MediaSubComponent) GetFNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FNum
}

// GetFNumOk returns a tuple with the FNum field value
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetFNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FNum, true
}

// SetFNum sets field value
func (o *MediaSubComponent) SetFNum(v int32) {
	o.FNum = v
}

// GetFDescs returns the FDescs field value if set, zero value otherwise.
func (o *MediaSubComponent) GetFDescs() []string {
	if o == nil || IsNil(o.FDescs) {
		var ret []string
		return ret
	}
	return o.FDescs
}

// GetFDescsOk returns a tuple with the FDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetFDescsOk() ([]string, bool) {
	if o == nil || IsNil(o.FDescs) {
		return nil, false
	}
	return o.FDescs, true
}

// HasFDescs returns a boolean if a field has been set.
func (o *MediaSubComponent) HasFDescs() bool {
	if o != nil && !IsNil(o.FDescs) {
		return true
	}

	return false
}

// SetFDescs gets a reference to the given []string and assigns it to the FDescs field.
func (o *MediaSubComponent) SetFDescs(v []string) {
	o.FDescs = v
}

// GetAddInfoFlowDescs returns the AddInfoFlowDescs field value if set, zero value otherwise.
func (o *MediaSubComponent) GetAddInfoFlowDescs() []AddFlowDescriptionInfo {
	if o == nil || IsNil(o.AddInfoFlowDescs) {
		var ret []AddFlowDescriptionInfo
		return ret
	}
	return o.AddInfoFlowDescs
}

// GetAddInfoFlowDescsOk returns a tuple with the AddInfoFlowDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetAddInfoFlowDescsOk() ([]AddFlowDescriptionInfo, bool) {
	if o == nil || IsNil(o.AddInfoFlowDescs) {
		return nil, false
	}
	return o.AddInfoFlowDescs, true
}

// HasAddInfoFlowDescs returns a boolean if a field has been set.
func (o *MediaSubComponent) HasAddInfoFlowDescs() bool {
	if o != nil && !IsNil(o.AddInfoFlowDescs) {
		return true
	}

	return false
}

// SetAddInfoFlowDescs gets a reference to the given []AddFlowDescriptionInfo and assigns it to the AddInfoFlowDescs field.
func (o *MediaSubComponent) SetAddInfoFlowDescs(v []AddFlowDescriptionInfo) {
	o.AddInfoFlowDescs = v
}

// GetFStatus returns the FStatus field value if set, zero value otherwise.
func (o *MediaSubComponent) GetFStatus() FlowStatus {
	if o == nil || IsNil(o.FStatus) {
		var ret FlowStatus
		return ret
	}
	return *o.FStatus
}

// GetFStatusOk returns a tuple with the FStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetFStatusOk() (*FlowStatus, bool) {
	if o == nil || IsNil(o.FStatus) {
		return nil, false
	}
	return o.FStatus, true
}

// HasFStatus returns a boolean if a field has been set.
func (o *MediaSubComponent) HasFStatus() bool {
	if o != nil && !IsNil(o.FStatus) {
		return true
	}

	return false
}

// SetFStatus gets a reference to the given FlowStatus and assigns it to the FStatus field.
func (o *MediaSubComponent) SetFStatus(v FlowStatus) {
	o.FStatus = &v
}

// GetMarBwDl returns the MarBwDl field value if set, zero value otherwise.
func (o *MediaSubComponent) GetMarBwDl() string {
	if o == nil || IsNil(o.MarBwDl) {
		var ret string
		return ret
	}
	return *o.MarBwDl
}

// GetMarBwDlOk returns a tuple with the MarBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetMarBwDlOk() (*string, bool) {
	if o == nil || IsNil(o.MarBwDl) {
		return nil, false
	}
	return o.MarBwDl, true
}

// HasMarBwDl returns a boolean if a field has been set.
func (o *MediaSubComponent) HasMarBwDl() bool {
	if o != nil && !IsNil(o.MarBwDl) {
		return true
	}

	return false
}

// SetMarBwDl gets a reference to the given string and assigns it to the MarBwDl field.
func (o *MediaSubComponent) SetMarBwDl(v string) {
	o.MarBwDl = &v
}

// GetMarBwUl returns the MarBwUl field value if set, zero value otherwise.
func (o *MediaSubComponent) GetMarBwUl() string {
	if o == nil || IsNil(o.MarBwUl) {
		var ret string
		return ret
	}
	return *o.MarBwUl
}

// GetMarBwUlOk returns a tuple with the MarBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetMarBwUlOk() (*string, bool) {
	if o == nil || IsNil(o.MarBwUl) {
		return nil, false
	}
	return o.MarBwUl, true
}

// HasMarBwUl returns a boolean if a field has been set.
func (o *MediaSubComponent) HasMarBwUl() bool {
	if o != nil && !IsNil(o.MarBwUl) {
		return true
	}

	return false
}

// SetMarBwUl gets a reference to the given string and assigns it to the MarBwUl field.
func (o *MediaSubComponent) SetMarBwUl(v string) {
	o.MarBwUl = &v
}

// GetTosTrCl returns the TosTrCl field value if set, zero value otherwise.
func (o *MediaSubComponent) GetTosTrCl() string {
	if o == nil || IsNil(o.TosTrCl) {
		var ret string
		return ret
	}
	return *o.TosTrCl
}

// GetTosTrClOk returns a tuple with the TosTrCl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetTosTrClOk() (*string, bool) {
	if o == nil || IsNil(o.TosTrCl) {
		return nil, false
	}
	return o.TosTrCl, true
}

// HasTosTrCl returns a boolean if a field has been set.
func (o *MediaSubComponent) HasTosTrCl() bool {
	if o != nil && !IsNil(o.TosTrCl) {
		return true
	}

	return false
}

// SetTosTrCl gets a reference to the given string and assigns it to the TosTrCl field.
func (o *MediaSubComponent) SetTosTrCl(v string) {
	o.TosTrCl = &v
}

// GetFlowUsage returns the FlowUsage field value if set, zero value otherwise.
func (o *MediaSubComponent) GetFlowUsage() string {
	if o == nil || IsNil(o.FlowUsage) {
		var ret string
		return ret
	}
	return *o.FlowUsage
}

// GetFlowUsageOk returns a tuple with the FlowUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetFlowUsageOk() (*string, bool) {
	if o == nil || IsNil(o.FlowUsage) {
		return nil, false
	}
	return o.FlowUsage, true
}

// HasFlowUsage returns a boolean if a field has been set.
func (o *MediaSubComponent) HasFlowUsage() bool {
	if o != nil && !IsNil(o.FlowUsage) {
		return true
	}

	return false
}

// SetFlowUsage gets a reference to the given FlowUsage and assigns it to the FlowUsage field.
func (o *MediaSubComponent) SetFlowUsage(v string) {
	o.FlowUsage = &v
}

// GetEvSubsc returns the EvSubsc field value if set, zero value otherwise.
func (o *MediaSubComponent) GetEvSubsc() EventsSubscReqData {
	if o == nil || IsNil(o.EvSubsc) {
		var ret EventsSubscReqData
		return ret
	}
	return *o.EvSubsc
}

// GetEvSubscOk returns a tuple with the EvSubsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSubComponent) GetEvSubscOk() (*EventsSubscReqData, bool) {
	if o == nil || IsNil(o.EvSubsc) {
		return nil, false
	}
	return o.EvSubsc, true
}

// HasEvSubsc returns a boolean if a field has been set.
func (o *MediaSubComponent) HasEvSubsc() bool {
	if o != nil && !IsNil(o.EvSubsc) {
		return true
	}

	return false
}

// SetEvSubsc gets a reference to the given EventsSubscReqData and assigns it to the EvSubsc field.
func (o *MediaSubComponent) SetEvSubsc(v EventsSubscReqData) {
	o.EvSubsc = &v
}

func (o MediaSubComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaSubComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AfSigProtocol.IsSet() {
		toSerialize["afSigProtocol"] = o.AfSigProtocol.Get()
	}
	if !IsNil(o.EthfDescs) {
		toSerialize["ethfDescs"] = o.EthfDescs
	}
	toSerialize["fNum"] = o.FNum
	if !IsNil(o.FDescs) {
		toSerialize["fDescs"] = o.FDescs
	}
	if !IsNil(o.AddInfoFlowDescs) {
		toSerialize["addInfoFlowDescs"] = o.AddInfoFlowDescs
	}
	if !IsNil(o.FStatus) {
		toSerialize["fStatus"] = o.FStatus
	}
	if !IsNil(o.MarBwDl) {
		toSerialize["marBwDl"] = o.MarBwDl
	}
	if !IsNil(o.MarBwUl) {
		toSerialize["marBwUl"] = o.MarBwUl
	}
	if !IsNil(o.TosTrCl) {
		toSerialize["tosTrCl"] = o.TosTrCl
	}
	if !IsNil(o.FlowUsage) {
		toSerialize["flowUsage"] = o.FlowUsage
	}
	if !IsNil(o.EvSubsc) {
		toSerialize["evSubsc"] = o.EvSubsc
	}
	return toSerialize, nil
}

func (o *MediaSubComponent) UnmarshalJSON(data []byte) (err error) {
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

	varMediaSubComponent := _MediaSubComponent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaSubComponent)

	if err != nil {
		return err
	}

	*o = MediaSubComponent(varMediaSubComponent)

	return err
}

type NullableMediaSubComponent struct {
	value *MediaSubComponent
	isSet bool
}

func (v NullableMediaSubComponent) Get() *MediaSubComponent {
	return v.value
}

func (v *NullableMediaSubComponent) Set(val *MediaSubComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSubComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSubComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSubComponent(val *MediaSubComponent) *NullableMediaSubComponent {
	return &NullableMediaSubComponent{value: val, isSet: true}
}

func (v NullableMediaSubComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSubComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
