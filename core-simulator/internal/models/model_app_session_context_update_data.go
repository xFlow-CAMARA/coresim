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

// checks if the AppSessionContextUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSessionContextUpdateData{}

// AppSessionContextUpdateData Identifies the modifications to the \"ascReqData\" property of an Individual Application Session Context which may include the modifications to the sub-resource Events Subscription.
type AppSessionContextUpdateData struct {
	// Contains an AF application identifier.
	AfAppId   *string                        `json:"afAppId,omitempty"`
	AfRoutReq NullableAfRoutingRequirementRm `json:"afRoutReq,omitempty"`
	AfSfcReq  NullableAfSfcRequirement       `json:"afSfcReq,omitempty"`
	// Contains an identity of an application service provider.
	AspId *string `json:"aspId,omitempty"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId *string                      `json:"bdtRefId,omitempty"`
	EvSubsc  NullableEventsSubscReqDataRm `json:"evSubsc,omitempty"`
	// Indication of MCPTT service request.
	McpttId *string `json:"mcpttId,omitempty"`
	// Indication of modification of MCVideo service.
	McVideoId *string `json:"mcVideoId,omitempty"`
	// Contains media component information. The key of the map is the medCompN attribute.
	MedComponents *map[string]MediaComponentRm `json:"medComponents,omitempty"`
	MpsAction     *MpsAction                   `json:"mpsAction,omitempty"`
	// Indication of MPS service request.
	MpsId *string `json:"mpsId,omitempty"`
	// Indication of MCS service request.
	McsId              *string                              `json:"mcsId,omitempty"`
	PreemptControlInfo NullablePreemptionControlInformation `json:"preemptControlInfo,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	QosDuration NullableInt32 `json:"qosDuration,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	QosInactInt   NullableInt32         `json:"qosInactInt,omitempty"`
	ResPrio       *ReservPriority       `json:"resPrio,omitempty"`
	ServInfStatus *ServiceInfoStatus    `json:"servInfStatus,omitempty"`
	SipForkInd    *SipForkingIndication `json:"sipForkInd,omitempty"`
	// Contains an identity of a sponsor.
	SponId              *string                    `json:"sponId,omitempty"`
	SponStatus          *SponsoringStatus          `json:"sponStatus,omitempty"`
	TsnBridgeManCont    *BridgeManagementContainer `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt  *PortManagementContainer   `json:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts []PortManagementContainer  `json:"tsnPortManContNwtts,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	TscNotifUri *string `json:"tscNotifUri,omitempty"`
	// Correlation identifier for TSC management information notifications.
	TscNotifCorreId *string `json:"tscNotifCorreId,omitempty"`
}

// NewAppSessionContextUpdateData instantiates a new AppSessionContextUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSessionContextUpdateData() *AppSessionContextUpdateData {
	this := AppSessionContextUpdateData{}
	return &this
}

// NewAppSessionContextUpdateDataWithDefaults instantiates a new AppSessionContextUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSessionContextUpdateDataWithDefaults() *AppSessionContextUpdateData {
	this := AppSessionContextUpdateData{}
	return &this
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetAfAppId() string {
	if o == nil || IsNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetAfAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasAfAppId() bool {
	if o != nil && !IsNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *AppSessionContextUpdateData) SetAfAppId(v string) {
	o.AfAppId = &v
}

// GetAfRoutReq returns the AfRoutReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetAfRoutReq() AfRoutingRequirementRm {
	if o == nil || IsNil(o.AfRoutReq.Get()) {
		var ret AfRoutingRequirementRm
		return ret
	}
	return *o.AfRoutReq.Get()
}

// GetAfRoutReqOk returns a tuple with the AfRoutReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetAfRoutReqOk() (*AfRoutingRequirementRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfRoutReq.Get(), o.AfRoutReq.IsSet()
}

// HasAfRoutReq returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasAfRoutReq() bool {
	if o != nil && o.AfRoutReq.IsSet() {
		return true
	}

	return false
}

// SetAfRoutReq gets a reference to the given NullableAfRoutingRequirementRm and assigns it to the AfRoutReq field.
func (o *AppSessionContextUpdateData) SetAfRoutReq(v AfRoutingRequirementRm) {
	o.AfRoutReq.Set(&v)
}

// SetAfRoutReqNil sets the value for AfRoutReq to be an explicit nil
func (o *AppSessionContextUpdateData) SetAfRoutReqNil() {
	o.AfRoutReq.Set(nil)
}

// UnsetAfRoutReq ensures that no value is present for AfRoutReq, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetAfRoutReq() {
	o.AfRoutReq.Unset()
}

// GetAfSfcReq returns the AfSfcReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetAfSfcReq() AfSfcRequirement {
	if o == nil || IsNil(o.AfSfcReq.Get()) {
		var ret AfSfcRequirement
		return ret
	}
	return *o.AfSfcReq.Get()
}

// GetAfSfcReqOk returns a tuple with the AfSfcReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetAfSfcReqOk() (*AfSfcRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSfcReq.Get(), o.AfSfcReq.IsSet()
}

// HasAfSfcReq returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasAfSfcReq() bool {
	if o != nil && o.AfSfcReq.IsSet() {
		return true
	}

	return false
}

// SetAfSfcReq gets a reference to the given NullableAfSfcRequirement and assigns it to the AfSfcReq field.
func (o *AppSessionContextUpdateData) SetAfSfcReq(v AfSfcRequirement) {
	o.AfSfcReq.Set(&v)
}

// SetAfSfcReqNil sets the value for AfSfcReq to be an explicit nil
func (o *AppSessionContextUpdateData) SetAfSfcReqNil() {
	o.AfSfcReq.Set(nil)
}

// UnsetAfSfcReq ensures that no value is present for AfSfcReq, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetAfSfcReq() {
	o.AfSfcReq.Unset()
}

// GetAspId returns the AspId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetAspId() string {
	if o == nil || IsNil(o.AspId) {
		var ret string
		return ret
	}
	return *o.AspId
}

// GetAspIdOk returns a tuple with the AspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetAspIdOk() (*string, bool) {
	if o == nil || IsNil(o.AspId) {
		return nil, false
	}
	return o.AspId, true
}

// HasAspId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasAspId() bool {
	if o != nil && !IsNil(o.AspId) {
		return true
	}

	return false
}

// SetAspId gets a reference to the given string and assigns it to the AspId field.
func (o *AppSessionContextUpdateData) SetAspId(v string) {
	o.AspId = &v
}

// GetBdtRefId returns the BdtRefId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetBdtRefId() string {
	if o == nil || IsNil(o.BdtRefId) {
		var ret string
		return ret
	}
	return *o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetBdtRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.BdtRefId) {
		return nil, false
	}
	return o.BdtRefId, true
}

// HasBdtRefId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasBdtRefId() bool {
	if o != nil && !IsNil(o.BdtRefId) {
		return true
	}

	return false
}

// SetBdtRefId gets a reference to the given string and assigns it to the BdtRefId field.
func (o *AppSessionContextUpdateData) SetBdtRefId(v string) {
	o.BdtRefId = &v
}

// GetEvSubsc returns the EvSubsc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetEvSubsc() EventsSubscReqDataRm {
	if o == nil || IsNil(o.EvSubsc.Get()) {
		var ret EventsSubscReqDataRm
		return ret
	}
	return *o.EvSubsc.Get()
}

// GetEvSubscOk returns a tuple with the EvSubsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetEvSubscOk() (*EventsSubscReqDataRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvSubsc.Get(), o.EvSubsc.IsSet()
}

// HasEvSubsc returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasEvSubsc() bool {
	if o != nil && o.EvSubsc.IsSet() {
		return true
	}

	return false
}

// SetEvSubsc gets a reference to the given NullableEventsSubscReqDataRm and assigns it to the EvSubsc field.
func (o *AppSessionContextUpdateData) SetEvSubsc(v EventsSubscReqDataRm) {
	o.EvSubsc.Set(&v)
}

// SetEvSubscNil sets the value for EvSubsc to be an explicit nil
func (o *AppSessionContextUpdateData) SetEvSubscNil() {
	o.EvSubsc.Set(nil)
}

// UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetEvSubsc() {
	o.EvSubsc.Unset()
}

// GetMcpttId returns the McpttId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMcpttId() string {
	if o == nil || IsNil(o.McpttId) {
		var ret string
		return ret
	}
	return *o.McpttId
}

// GetMcpttIdOk returns a tuple with the McpttId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMcpttIdOk() (*string, bool) {
	if o == nil || IsNil(o.McpttId) {
		return nil, false
	}
	return o.McpttId, true
}

// HasMcpttId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMcpttId() bool {
	if o != nil && !IsNil(o.McpttId) {
		return true
	}

	return false
}

// SetMcpttId gets a reference to the given string and assigns it to the McpttId field.
func (o *AppSessionContextUpdateData) SetMcpttId(v string) {
	o.McpttId = &v
}

// GetMcVideoId returns the McVideoId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMcVideoId() string {
	if o == nil || IsNil(o.McVideoId) {
		var ret string
		return ret
	}
	return *o.McVideoId
}

// GetMcVideoIdOk returns a tuple with the McVideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMcVideoIdOk() (*string, bool) {
	if o == nil || IsNil(o.McVideoId) {
		return nil, false
	}
	return o.McVideoId, true
}

// HasMcVideoId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMcVideoId() bool {
	if o != nil && !IsNil(o.McVideoId) {
		return true
	}

	return false
}

// SetMcVideoId gets a reference to the given string and assigns it to the McVideoId field.
func (o *AppSessionContextUpdateData) SetMcVideoId(v string) {
	o.McVideoId = &v
}

// GetMedComponents returns the MedComponents field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMedComponents() map[string]MediaComponentRm {
	if o == nil || IsNil(o.MedComponents) {
		var ret map[string]MediaComponentRm
		return ret
	}
	return *o.MedComponents
}

// GetMedComponentsOk returns a tuple with the MedComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMedComponentsOk() (*map[string]MediaComponentRm, bool) {
	if o == nil || IsNil(o.MedComponents) {
		return nil, false
	}
	return o.MedComponents, true
}

// HasMedComponents returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMedComponents() bool {
	if o != nil && !IsNil(o.MedComponents) {
		return true
	}

	return false
}

// SetMedComponents gets a reference to the given map[string]MediaComponentRm and assigns it to the MedComponents field.
func (o *AppSessionContextUpdateData) SetMedComponents(v map[string]MediaComponentRm) {
	o.MedComponents = &v
}

// GetMpsAction returns the MpsAction field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMpsAction() MpsAction {
	if o == nil || IsNil(o.MpsAction) {
		var ret MpsAction
		return ret
	}
	return *o.MpsAction
}

// GetMpsActionOk returns a tuple with the MpsAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMpsActionOk() (*MpsAction, bool) {
	if o == nil || IsNil(o.MpsAction) {
		return nil, false
	}
	return o.MpsAction, true
}

// HasMpsAction returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMpsAction() bool {
	if o != nil && !IsNil(o.MpsAction) {
		return true
	}

	return false
}

// SetMpsAction gets a reference to the given MpsAction and assigns it to the MpsAction field.
func (o *AppSessionContextUpdateData) SetMpsAction(v MpsAction) {
	o.MpsAction = &v
}

// GetMpsId returns the MpsId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMpsId() string {
	if o == nil || IsNil(o.MpsId) {
		var ret string
		return ret
	}
	return *o.MpsId
}

// GetMpsIdOk returns a tuple with the MpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMpsIdOk() (*string, bool) {
	if o == nil || IsNil(o.MpsId) {
		return nil, false
	}
	return o.MpsId, true
}

// HasMpsId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMpsId() bool {
	if o != nil && !IsNil(o.MpsId) {
		return true
	}

	return false
}

// SetMpsId gets a reference to the given string and assigns it to the MpsId field.
func (o *AppSessionContextUpdateData) SetMpsId(v string) {
	o.MpsId = &v
}

// GetMcsId returns the McsId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetMcsId() string {
	if o == nil || IsNil(o.McsId) {
		var ret string
		return ret
	}
	return *o.McsId
}

// GetMcsIdOk returns a tuple with the McsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetMcsIdOk() (*string, bool) {
	if o == nil || IsNil(o.McsId) {
		return nil, false
	}
	return o.McsId, true
}

// HasMcsId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasMcsId() bool {
	if o != nil && !IsNil(o.McsId) {
		return true
	}

	return false
}

// SetMcsId gets a reference to the given string and assigns it to the McsId field.
func (o *AppSessionContextUpdateData) SetMcsId(v string) {
	o.McsId = &v
}

// GetPreemptControlInfo returns the PreemptControlInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetPreemptControlInfo() PreemptionControlInformation {
	if o == nil || IsNil(o.PreemptControlInfo.Get()) {
		var ret PreemptionControlInformation
		return ret
	}
	return *o.PreemptControlInfo.Get()
}

// GetPreemptControlInfoOk returns a tuple with the PreemptControlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetPreemptControlInfoOk() (*PreemptionControlInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreemptControlInfo.Get(), o.PreemptControlInfo.IsSet()
}

// HasPreemptControlInfo returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasPreemptControlInfo() bool {
	if o != nil && o.PreemptControlInfo.IsSet() {
		return true
	}

	return false
}

// SetPreemptControlInfo gets a reference to the given NullablePreemptionControlInformation and assigns it to the PreemptControlInfo field.
func (o *AppSessionContextUpdateData) SetPreemptControlInfo(v PreemptionControlInformation) {
	o.PreemptControlInfo.Set(&v)
}

// SetPreemptControlInfoNil sets the value for PreemptControlInfo to be an explicit nil
func (o *AppSessionContextUpdateData) SetPreemptControlInfoNil() {
	o.PreemptControlInfo.Set(nil)
}

// UnsetPreemptControlInfo ensures that no value is present for PreemptControlInfo, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetPreemptControlInfo() {
	o.PreemptControlInfo.Unset()
}

// GetQosDuration returns the QosDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetQosDuration() int32 {
	if o == nil || IsNil(o.QosDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.QosDuration.Get()
}

// GetQosDurationOk returns a tuple with the QosDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetQosDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosDuration.Get(), o.QosDuration.IsSet()
}

// HasQosDuration returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasQosDuration() bool {
	if o != nil && o.QosDuration.IsSet() {
		return true
	}

	return false
}

// SetQosDuration gets a reference to the given NullableInt32 and assigns it to the QosDuration field.
func (o *AppSessionContextUpdateData) SetQosDuration(v int32) {
	o.QosDuration.Set(&v)
}

// SetQosDurationNil sets the value for QosDuration to be an explicit nil
func (o *AppSessionContextUpdateData) SetQosDurationNil() {
	o.QosDuration.Set(nil)
}

// UnsetQosDuration ensures that no value is present for QosDuration, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetQosDuration() {
	o.QosDuration.Unset()
}

// GetQosInactInt returns the QosInactInt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextUpdateData) GetQosInactInt() int32 {
	if o == nil || IsNil(o.QosInactInt.Get()) {
		var ret int32
		return ret
	}
	return *o.QosInactInt.Get()
}

// GetQosInactIntOk returns a tuple with the QosInactInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextUpdateData) GetQosInactIntOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosInactInt.Get(), o.QosInactInt.IsSet()
}

// HasQosInactInt returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasQosInactInt() bool {
	if o != nil && o.QosInactInt.IsSet() {
		return true
	}

	return false
}

// SetQosInactInt gets a reference to the given NullableInt32 and assigns it to the QosInactInt field.
func (o *AppSessionContextUpdateData) SetQosInactInt(v int32) {
	o.QosInactInt.Set(&v)
}

// SetQosInactIntNil sets the value for QosInactInt to be an explicit nil
func (o *AppSessionContextUpdateData) SetQosInactIntNil() {
	o.QosInactInt.Set(nil)
}

// UnsetQosInactInt ensures that no value is present for QosInactInt, not even an explicit nil
func (o *AppSessionContextUpdateData) UnsetQosInactInt() {
	o.QosInactInt.Unset()
}

// GetResPrio returns the ResPrio field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetResPrio() ReservPriority {
	if o == nil || IsNil(o.ResPrio) {
		var ret ReservPriority
		return ret
	}
	return *o.ResPrio
}

// GetResPrioOk returns a tuple with the ResPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetResPrioOk() (*ReservPriority, bool) {
	if o == nil || IsNil(o.ResPrio) {
		return nil, false
	}
	return o.ResPrio, true
}

// HasResPrio returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasResPrio() bool {
	if o != nil && !IsNil(o.ResPrio) {
		return true
	}

	return false
}

// SetResPrio gets a reference to the given ReservPriority and assigns it to the ResPrio field.
func (o *AppSessionContextUpdateData) SetResPrio(v ReservPriority) {
	o.ResPrio = &v
}

// GetServInfStatus returns the ServInfStatus field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetServInfStatus() ServiceInfoStatus {
	if o == nil || IsNil(o.ServInfStatus) {
		var ret ServiceInfoStatus
		return ret
	}
	return *o.ServInfStatus
}

// GetServInfStatusOk returns a tuple with the ServInfStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetServInfStatusOk() (*ServiceInfoStatus, bool) {
	if o == nil || IsNil(o.ServInfStatus) {
		return nil, false
	}
	return o.ServInfStatus, true
}

// HasServInfStatus returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasServInfStatus() bool {
	if o != nil && !IsNil(o.ServInfStatus) {
		return true
	}

	return false
}

// SetServInfStatus gets a reference to the given ServiceInfoStatus and assigns it to the ServInfStatus field.
func (o *AppSessionContextUpdateData) SetServInfStatus(v ServiceInfoStatus) {
	o.ServInfStatus = &v
}

// GetSipForkInd returns the SipForkInd field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetSipForkInd() SipForkingIndication {
	if o == nil || IsNil(o.SipForkInd) {
		var ret SipForkingIndication
		return ret
	}
	return *o.SipForkInd
}

// GetSipForkIndOk returns a tuple with the SipForkInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetSipForkIndOk() (*SipForkingIndication, bool) {
	if o == nil || IsNil(o.SipForkInd) {
		return nil, false
	}
	return o.SipForkInd, true
}

// HasSipForkInd returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasSipForkInd() bool {
	if o != nil && !IsNil(o.SipForkInd) {
		return true
	}

	return false
}

// SetSipForkInd gets a reference to the given SipForkingIndication and assigns it to the SipForkInd field.
func (o *AppSessionContextUpdateData) SetSipForkInd(v SipForkingIndication) {
	o.SipForkInd = &v
}

// GetSponId returns the SponId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetSponId() string {
	if o == nil || IsNil(o.SponId) {
		var ret string
		return ret
	}
	return *o.SponId
}

// GetSponIdOk returns a tuple with the SponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetSponIdOk() (*string, bool) {
	if o == nil || IsNil(o.SponId) {
		return nil, false
	}
	return o.SponId, true
}

// HasSponId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasSponId() bool {
	if o != nil && !IsNil(o.SponId) {
		return true
	}

	return false
}

// SetSponId gets a reference to the given string and assigns it to the SponId field.
func (o *AppSessionContextUpdateData) SetSponId(v string) {
	o.SponId = &v
}

// GetSponStatus returns the SponStatus field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetSponStatus() SponsoringStatus {
	if o == nil || IsNil(o.SponStatus) {
		var ret SponsoringStatus
		return ret
	}
	return *o.SponStatus
}

// GetSponStatusOk returns a tuple with the SponStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetSponStatusOk() (*SponsoringStatus, bool) {
	if o == nil || IsNil(o.SponStatus) {
		return nil, false
	}
	return o.SponStatus, true
}

// HasSponStatus returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasSponStatus() bool {
	if o != nil && !IsNil(o.SponStatus) {
		return true
	}

	return false
}

// SetSponStatus gets a reference to the given SponsoringStatus and assigns it to the SponStatus field.
func (o *AppSessionContextUpdateData) SetSponStatus(v SponsoringStatus) {
	o.SponStatus = &v
}

// GetTsnBridgeManCont returns the TsnBridgeManCont field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetTsnBridgeManCont() BridgeManagementContainer {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		var ret BridgeManagementContainer
		return ret
	}
	return *o.TsnBridgeManCont
}

// GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool) {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		return nil, false
	}
	return o.TsnBridgeManCont, true
}

// HasTsnBridgeManCont returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasTsnBridgeManCont() bool {
	if o != nil && !IsNil(o.TsnBridgeManCont) {
		return true
	}

	return false
}

// SetTsnBridgeManCont gets a reference to the given BridgeManagementContainer and assigns it to the TsnBridgeManCont field.
func (o *AppSessionContextUpdateData) SetTsnBridgeManCont(v BridgeManagementContainer) {
	o.TsnBridgeManCont = &v
}

// GetTsnPortManContDstt returns the TsnPortManContDstt field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetTsnPortManContDstt() PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		var ret PortManagementContainer
		return ret
	}
	return *o.TsnPortManContDstt
}

// GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetTsnPortManContDsttOk() (*PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		return nil, false
	}
	return o.TsnPortManContDstt, true
}

// HasTsnPortManContDstt returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasTsnPortManContDstt() bool {
	if o != nil && !IsNil(o.TsnPortManContDstt) {
		return true
	}

	return false
}

// SetTsnPortManContDstt gets a reference to the given PortManagementContainer and assigns it to the TsnPortManContDstt field.
func (o *AppSessionContextUpdateData) SetTsnPortManContDstt(v PortManagementContainer) {
	o.TsnPortManContDstt = &v
}

// GetTsnPortManContNwtts returns the TsnPortManContNwtts field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetTsnPortManContNwtts() []PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		var ret []PortManagementContainer
		return ret
	}
	return o.TsnPortManContNwtts
}

// GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetTsnPortManContNwttsOk() ([]PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		return nil, false
	}
	return o.TsnPortManContNwtts, true
}

// HasTsnPortManContNwtts returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasTsnPortManContNwtts() bool {
	if o != nil && !IsNil(o.TsnPortManContNwtts) {
		return true
	}

	return false
}

// SetTsnPortManContNwtts gets a reference to the given []PortManagementContainer and assigns it to the TsnPortManContNwtts field.
func (o *AppSessionContextUpdateData) SetTsnPortManContNwtts(v []PortManagementContainer) {
	o.TsnPortManContNwtts = v
}

// GetTscNotifUri returns the TscNotifUri field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetTscNotifUri() string {
	if o == nil || IsNil(o.TscNotifUri) {
		var ret string
		return ret
	}
	return *o.TscNotifUri
}

// GetTscNotifUriOk returns a tuple with the TscNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetTscNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.TscNotifUri) {
		return nil, false
	}
	return o.TscNotifUri, true
}

// HasTscNotifUri returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasTscNotifUri() bool {
	if o != nil && !IsNil(o.TscNotifUri) {
		return true
	}

	return false
}

// SetTscNotifUri gets a reference to the given string and assigns it to the TscNotifUri field.
func (o *AppSessionContextUpdateData) SetTscNotifUri(v string) {
	o.TscNotifUri = &v
}

// GetTscNotifCorreId returns the TscNotifCorreId field value if set, zero value otherwise.
func (o *AppSessionContextUpdateData) GetTscNotifCorreId() string {
	if o == nil || IsNil(o.TscNotifCorreId) {
		var ret string
		return ret
	}
	return *o.TscNotifCorreId
}

// GetTscNotifCorreIdOk returns a tuple with the TscNotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextUpdateData) GetTscNotifCorreIdOk() (*string, bool) {
	if o == nil || IsNil(o.TscNotifCorreId) {
		return nil, false
	}
	return o.TscNotifCorreId, true
}

// HasTscNotifCorreId returns a boolean if a field has been set.
func (o *AppSessionContextUpdateData) HasTscNotifCorreId() bool {
	if o != nil && !IsNil(o.TscNotifCorreId) {
		return true
	}

	return false
}

// SetTscNotifCorreId gets a reference to the given string and assigns it to the TscNotifCorreId field.
func (o *AppSessionContextUpdateData) SetTscNotifCorreId(v string) {
	o.TscNotifCorreId = &v
}

func (o AppSessionContextUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSessionContextUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	if o.AfRoutReq.IsSet() {
		toSerialize["afRoutReq"] = o.AfRoutReq.Get()
	}
	if o.AfSfcReq.IsSet() {
		toSerialize["afSfcReq"] = o.AfSfcReq.Get()
	}
	if !IsNil(o.AspId) {
		toSerialize["aspId"] = o.AspId
	}
	if !IsNil(o.BdtRefId) {
		toSerialize["bdtRefId"] = o.BdtRefId
	}
	if o.EvSubsc.IsSet() {
		toSerialize["evSubsc"] = o.EvSubsc.Get()
	}
	if !IsNil(o.McpttId) {
		toSerialize["mcpttId"] = o.McpttId
	}
	if !IsNil(o.McVideoId) {
		toSerialize["mcVideoId"] = o.McVideoId
	}
	if !IsNil(o.MedComponents) {
		toSerialize["medComponents"] = o.MedComponents
	}
	if !IsNil(o.MpsAction) {
		toSerialize["mpsAction"] = o.MpsAction
	}
	if !IsNil(o.MpsId) {
		toSerialize["mpsId"] = o.MpsId
	}
	if !IsNil(o.McsId) {
		toSerialize["mcsId"] = o.McsId
	}
	if o.PreemptControlInfo.IsSet() {
		toSerialize["preemptControlInfo"] = o.PreemptControlInfo.Get()
	}
	if o.QosDuration.IsSet() {
		toSerialize["qosDuration"] = o.QosDuration.Get()
	}
	if o.QosInactInt.IsSet() {
		toSerialize["qosInactInt"] = o.QosInactInt.Get()
	}
	if !IsNil(o.ResPrio) {
		toSerialize["resPrio"] = o.ResPrio
	}
	if !IsNil(o.ServInfStatus) {
		toSerialize["servInfStatus"] = o.ServInfStatus
	}
	if !IsNil(o.SipForkInd) {
		toSerialize["sipForkInd"] = o.SipForkInd
	}
	if !IsNil(o.SponId) {
		toSerialize["sponId"] = o.SponId
	}
	if !IsNil(o.SponStatus) {
		toSerialize["sponStatus"] = o.SponStatus
	}
	if !IsNil(o.TsnBridgeManCont) {
		toSerialize["tsnBridgeManCont"] = o.TsnBridgeManCont
	}
	if !IsNil(o.TsnPortManContDstt) {
		toSerialize["tsnPortManContDstt"] = o.TsnPortManContDstt
	}
	if !IsNil(o.TsnPortManContNwtts) {
		toSerialize["tsnPortManContNwtts"] = o.TsnPortManContNwtts
	}
	if !IsNil(o.TscNotifUri) {
		toSerialize["tscNotifUri"] = o.TscNotifUri
	}
	if !IsNil(o.TscNotifCorreId) {
		toSerialize["tscNotifCorreId"] = o.TscNotifCorreId
	}
	return toSerialize, nil
}

type NullableAppSessionContextUpdateData struct {
	value *AppSessionContextUpdateData
	isSet bool
}

func (v NullableAppSessionContextUpdateData) Get() *AppSessionContextUpdateData {
	return v.value
}

func (v *NullableAppSessionContextUpdateData) Set(val *AppSessionContextUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContextUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContextUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContextUpdateData(val *AppSessionContextUpdateData) *NullableAppSessionContextUpdateData {
	return &NullableAppSessionContextUpdateData{value: val, isSet: true}
}

func (v NullableAppSessionContextUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContextUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
