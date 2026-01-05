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

// checks if the MediaComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaComponent{}

// MediaComponent Identifies a media component.
type MediaComponent struct {
	// Contains an AF application identifier.
	AfAppId      *string                  `json:"afAppId,omitempty"`
	AfRoutReq    *AfRoutingRequirement    `json:"afRoutReq,omitempty"`
	AfSfcReq     NullableAfSfcRequirement `json:"afSfcReq,omitempty"`
	QosReference *string                  `json:"qosReference,omitempty"`
	DisUeNotif   *bool                    `json:"disUeNotif,omitempty"`
	AltSerReqs   []string                 `json:"altSerReqs,omitempty"`
	// Contains alternative service requirements that include individual QoS parameter sets.
	AltSerReqsData []AlternativeServiceRequirementsData `json:"altSerReqsData,omitempty"`
	// Represents the content version of some content.
	ContVer *int32   `json:"contVer,omitempty"`
	Codecs  []string `json:"codecs,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DesMaxLatency *float32 `json:"desMaxLatency,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DesMaxLoss *float32 `json:"desMaxLoss,omitempty"`
	FlusId     *string  `json:"flusId,omitempty"`
	FStatus    *string  `json:"fStatus,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDl *string `json:"marBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUl *string `json:"marBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateDl NullableInt32 `json:"maxPacketLossRateDl,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property.
	MaxPacketLossRateUl NullableInt32 `json:"maxPacketLossRateUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwDl *string `json:"maxSuppBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxSuppBwUl *string `json:"maxSuppBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	MedCompN    int32   `json:"medCompN"`
	// Contains the requested bitrate and filters for the set of service data flows identified by their common flow identifier. The key of the map is the fNum attribute.
	MedSubComps *map[string]MediaSubComponent `json:"medSubComps,omitempty"`
	MedType     *string                       `json:"medType,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwDl *string `json:"minDesBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwUl *string `json:"minDesBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDl *string `json:"mirBwDl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUl        *string                   `json:"mirBwUl,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	PreemptCap     *PreemptionCapability     `json:"preemptCap,omitempty"`
	PreemptVuln    *PreemptionVulnerability  `json:"preemptVuln,omitempty"`
	PrioSharingInd *PrioritySharingIndicator `json:"prioSharingInd,omitempty"`
	ResPrio        *ReservPriority           `json:"resPrio,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RrBw *string `json:"rrBw,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RsBw *string `json:"rsBw,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyDl *int32 `json:"sharingKeyDl,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	SharingKeyUl *int32                      `json:"sharingKeyUl,omitempty"`
	TsnQos       *TsnQosContainer            `json:"tsnQos,omitempty"`
	TscaiInputDl NullableTscaiInputContainer `json:"tscaiInputDl,omitempty"`
	TscaiInputUl NullableTscaiInputContainer `json:"tscaiInputUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom *int32 `json:"tscaiTimeDom,omitempty"`
	// Indicates the capability for AF to adjust the burst sending time, when it is supported and set to \"true\". The default value is \"false\" if omitted.
	CapBatAdaptation *bool `json:"capBatAdaptation,omitempty"`
	// Indicates the service data flow needs to meet the Round-Trip (RT) latency requirement of the service, when it is included and set to \"true\". The default value is \"false\" if omitted.
	RTLatencyInd *bool                `json:"rTLatencyInd,omitempty"`
	PduSetQosDl  *PduSetQosPara       `json:"pduSetQosDl,omitempty"`
	PduSetQosUl  *PduSetQosPara       `json:"pduSetQosUl,omitempty"`
	ProtoDescDl  *ProtocolDescription `json:"protoDescDl,omitempty"`
	ProtoDescUl  *ProtocolDescription `json:"protoDescUl,omitempty"`
	// Indicates the time interval in units of milliseconds.
	PeriodUl *int32 `json:"periodUl,omitempty"`
	// Indicates the time interval in units of milliseconds.
	PeriodDl *int32                 `json:"periodDl,omitempty"`
	L4sInd   *UplinkDownlinkSupport `json:"l4sInd,omitempty"`
}

type _MediaComponent MediaComponent

// NewMediaComponent instantiates a new MediaComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaComponent(medCompN int32) *MediaComponent {
	this := MediaComponent{}
	return &this
}

// NewMediaComponentWithDefaults instantiates a new MediaComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaComponentWithDefaults() *MediaComponent {
	this := MediaComponent{}
	return &this
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *MediaComponent) GetAfAppId() string {
	if o == nil || IsNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetAfAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *MediaComponent) HasAfAppId() bool {
	if o != nil && !IsNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *MediaComponent) SetAfAppId(v string) {
	o.AfAppId = &v
}

// GetAfRoutReq returns the AfRoutReq field value if set, zero value otherwise.
func (o *MediaComponent) GetAfRoutReq() AfRoutingRequirement {
	if o == nil || IsNil(o.AfRoutReq) {
		var ret AfRoutingRequirement
		return ret
	}
	return *o.AfRoutReq
}

// GetAfRoutReqOk returns a tuple with the AfRoutReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetAfRoutReqOk() (*AfRoutingRequirement, bool) {
	if o == nil || IsNil(o.AfRoutReq) {
		return nil, false
	}
	return o.AfRoutReq, true
}

// HasAfRoutReq returns a boolean if a field has been set.
func (o *MediaComponent) HasAfRoutReq() bool {
	if o != nil && !IsNil(o.AfRoutReq) {
		return true
	}

	return false
}

// SetAfRoutReq gets a reference to the given AfRoutingRequirement and assigns it to the AfRoutReq field.
func (o *MediaComponent) SetAfRoutReq(v AfRoutingRequirement) {
	o.AfRoutReq = &v
}

// GetAfSfcReq returns the AfSfcReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaComponent) GetAfSfcReq() AfSfcRequirement {
	if o == nil || IsNil(o.AfSfcReq.Get()) {
		var ret AfSfcRequirement
		return ret
	}
	return *o.AfSfcReq.Get()
}

// GetAfSfcReqOk returns a tuple with the AfSfcReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaComponent) GetAfSfcReqOk() (*AfSfcRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSfcReq.Get(), o.AfSfcReq.IsSet()
}

// HasAfSfcReq returns a boolean if a field has been set.
func (o *MediaComponent) HasAfSfcReq() bool {
	if o != nil && o.AfSfcReq.IsSet() {
		return true
	}

	return false
}

// SetAfSfcReq gets a reference to the given NullableAfSfcRequirement and assigns it to the AfSfcReq field.
func (o *MediaComponent) SetAfSfcReq(v AfSfcRequirement) {
	o.AfSfcReq.Set(&v)
}

// SetAfSfcReqNil sets the value for AfSfcReq to be an explicit nil
func (o *MediaComponent) SetAfSfcReqNil() {
	o.AfSfcReq.Set(nil)
}

// UnsetAfSfcReq ensures that no value is present for AfSfcReq, not even an explicit nil
func (o *MediaComponent) UnsetAfSfcReq() {
	o.AfSfcReq.Unset()
}

// GetQosReference returns the QosReference field value if set, zero value otherwise.
func (o *MediaComponent) GetQosReference() string {
	if o == nil || IsNil(o.QosReference) {
		var ret string
		return ret
	}
	return *o.QosReference
}

// GetQosReferenceOk returns a tuple with the QosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetQosReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.QosReference) {
		return nil, false
	}
	return o.QosReference, true
}

// HasQosReference returns a boolean if a field has been set.
func (o *MediaComponent) HasQosReference() bool {
	if o != nil && !IsNil(o.QosReference) {
		return true
	}

	return false
}

// SetQosReference gets a reference to the given string and assigns it to the QosReference field.
func (o *MediaComponent) SetQosReference(v string) {
	o.QosReference = &v
}

// GetDisUeNotif returns the DisUeNotif field value if set, zero value otherwise.
func (o *MediaComponent) GetDisUeNotif() bool {
	if o == nil || IsNil(o.DisUeNotif) {
		var ret bool
		return ret
	}
	return *o.DisUeNotif
}

// GetDisUeNotifOk returns a tuple with the DisUeNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetDisUeNotifOk() (*bool, bool) {
	if o == nil || IsNil(o.DisUeNotif) {
		return nil, false
	}
	return o.DisUeNotif, true
}

// HasDisUeNotif returns a boolean if a field has been set.
func (o *MediaComponent) HasDisUeNotif() bool {
	if o != nil && !IsNil(o.DisUeNotif) {
		return true
	}

	return false
}

// SetDisUeNotif gets a reference to the given bool and assigns it to the DisUeNotif field.
func (o *MediaComponent) SetDisUeNotif(v bool) {
	o.DisUeNotif = &v
}

// GetAltSerReqs returns the AltSerReqs field value if set, zero value otherwise.
func (o *MediaComponent) GetAltSerReqs() []string {
	if o == nil || IsNil(o.AltSerReqs) {
		var ret []string
		return ret
	}
	return o.AltSerReqs
}

// GetAltSerReqsOk returns a tuple with the AltSerReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetAltSerReqsOk() ([]string, bool) {
	if o == nil || IsNil(o.AltSerReqs) {
		return nil, false
	}
	return o.AltSerReqs, true
}

// HasAltSerReqs returns a boolean if a field has been set.
func (o *MediaComponent) HasAltSerReqs() bool {
	if o != nil && !IsNil(o.AltSerReqs) {
		return true
	}

	return false
}

// SetAltSerReqs gets a reference to the given []string and assigns it to the AltSerReqs field.
func (o *MediaComponent) SetAltSerReqs(v []string) {
	o.AltSerReqs = v
}

// GetAltSerReqsData returns the AltSerReqsData field value if set, zero value otherwise.
func (o *MediaComponent) GetAltSerReqsData() []AlternativeServiceRequirementsData {
	if o == nil || IsNil(o.AltSerReqsData) {
		var ret []AlternativeServiceRequirementsData
		return ret
	}
	return o.AltSerReqsData
}

// GetAltSerReqsDataOk returns a tuple with the AltSerReqsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetAltSerReqsDataOk() ([]AlternativeServiceRequirementsData, bool) {
	if o == nil || IsNil(o.AltSerReqsData) {
		return nil, false
	}
	return o.AltSerReqsData, true
}

// HasAltSerReqsData returns a boolean if a field has been set.
func (o *MediaComponent) HasAltSerReqsData() bool {
	if o != nil && !IsNil(o.AltSerReqsData) {
		return true
	}

	return false
}

// SetAltSerReqsData gets a reference to the given []AlternativeServiceRequirementsData and assigns it to the AltSerReqsData field.
func (o *MediaComponent) SetAltSerReqsData(v []AlternativeServiceRequirementsData) {
	o.AltSerReqsData = v
}

// GetContVer returns the ContVer field value if set, zero value otherwise.
func (o *MediaComponent) GetContVer() int32 {
	if o == nil || IsNil(o.ContVer) {
		var ret int32
		return ret
	}
	return *o.ContVer
}

// GetContVerOk returns a tuple with the ContVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetContVerOk() (*int32, bool) {
	if o == nil || IsNil(o.ContVer) {
		return nil, false
	}
	return o.ContVer, true
}

// HasContVer returns a boolean if a field has been set.
func (o *MediaComponent) HasContVer() bool {
	if o != nil && !IsNil(o.ContVer) {
		return true
	}

	return false
}

// SetContVer gets a reference to the given int32 and assigns it to the ContVer field.
func (o *MediaComponent) SetContVer(v int32) {
	o.ContVer = &v
}

// GetCodecs returns the Codecs field value if set, zero value otherwise.
func (o *MediaComponent) GetCodecs() []string {
	if o == nil || IsNil(o.Codecs) {
		var ret []string
		return ret
	}
	return o.Codecs
}

// GetCodecsOk returns a tuple with the Codecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetCodecsOk() ([]string, bool) {
	if o == nil || IsNil(o.Codecs) {
		return nil, false
	}
	return o.Codecs, true
}

// HasCodecs returns a boolean if a field has been set.
func (o *MediaComponent) HasCodecs() bool {
	if o != nil && !IsNil(o.Codecs) {
		return true
	}

	return false
}

// SetCodecs gets a reference to the given []string and assigns it to the Codecs field.
func (o *MediaComponent) SetCodecs(v []string) {
	o.Codecs = v
}

// GetDesMaxLatency returns the DesMaxLatency field value if set, zero value otherwise.
func (o *MediaComponent) GetDesMaxLatency() float32 {
	if o == nil || IsNil(o.DesMaxLatency) {
		var ret float32
		return ret
	}
	return *o.DesMaxLatency
}

// GetDesMaxLatencyOk returns a tuple with the DesMaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetDesMaxLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.DesMaxLatency) {
		return nil, false
	}
	return o.DesMaxLatency, true
}

// HasDesMaxLatency returns a boolean if a field has been set.
func (o *MediaComponent) HasDesMaxLatency() bool {
	if o != nil && !IsNil(o.DesMaxLatency) {
		return true
	}

	return false
}

// SetDesMaxLatency gets a reference to the given float32 and assigns it to the DesMaxLatency field.
func (o *MediaComponent) SetDesMaxLatency(v float32) {
	o.DesMaxLatency = &v
}

// GetDesMaxLoss returns the DesMaxLoss field value if set, zero value otherwise.
func (o *MediaComponent) GetDesMaxLoss() float32 {
	if o == nil || IsNil(o.DesMaxLoss) {
		var ret float32
		return ret
	}
	return *o.DesMaxLoss
}

// GetDesMaxLossOk returns a tuple with the DesMaxLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetDesMaxLossOk() (*float32, bool) {
	if o == nil || IsNil(o.DesMaxLoss) {
		return nil, false
	}
	return o.DesMaxLoss, true
}

// HasDesMaxLoss returns a boolean if a field has been set.
func (o *MediaComponent) HasDesMaxLoss() bool {
	if o != nil && !IsNil(o.DesMaxLoss) {
		return true
	}

	return false
}

// SetDesMaxLoss gets a reference to the given float32 and assigns it to the DesMaxLoss field.
func (o *MediaComponent) SetDesMaxLoss(v float32) {
	o.DesMaxLoss = &v
}

// GetFlusId returns the FlusId field value if set, zero value otherwise.
func (o *MediaComponent) GetFlusId() string {
	if o == nil || IsNil(o.FlusId) {
		var ret string
		return ret
	}
	return *o.FlusId
}

// GetFlusIdOk returns a tuple with the FlusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetFlusIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlusId) {
		return nil, false
	}
	return o.FlusId, true
}

// HasFlusId returns a boolean if a field has been set.
func (o *MediaComponent) HasFlusId() bool {
	if o != nil && !IsNil(o.FlusId) {
		return true
	}

	return false
}

// SetFlusId gets a reference to the given string and assigns it to the FlusId field.
func (o *MediaComponent) SetFlusId(v string) {
	o.FlusId = &v
}

// GetFStatus returns the FStatus field value if set, zero value otherwise.
func (o *MediaComponent) GetFStatus() string {
	if o == nil || IsNil(o.FStatus) {
		var ret string
		return ret
	}
	return *o.FStatus
}

// GetFStatusOk returns a tuple with the FStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetFStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FStatus) {
		return nil, false
	}
	return o.FStatus, true
}

// HasFStatus returns a boolean if a field has been set.
func (o *MediaComponent) HasFStatus() bool {
	if o != nil && !IsNil(o.FStatus) {
		return true
	}

	return false
}

// SetFStatus gets a reference to the given FlowStatus and assigns it to the FStatus field.
func (o *MediaComponent) SetFStatus(v string) {
	o.FStatus = &v
}

// GetMarBwDl returns the MarBwDl field value if set, zero value otherwise.
func (o *MediaComponent) GetMarBwDl() string {
	if o == nil || IsNil(o.MarBwDl) {
		var ret string
		return ret
	}
	return *o.MarBwDl
}

// GetMarBwDlOk returns a tuple with the MarBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMarBwDlOk() (*string, bool) {
	if o == nil || IsNil(o.MarBwDl) {
		return nil, false
	}
	return o.MarBwDl, true
}

// HasMarBwDl returns a boolean if a field has been set.
func (o *MediaComponent) HasMarBwDl() bool {
	if o != nil && !IsNil(o.MarBwDl) {
		return true
	}

	return false
}

// SetMarBwDl gets a reference to the given string and assigns it to the MarBwDl field.
func (o *MediaComponent) SetMarBwDl(v string) {
	o.MarBwDl = &v
}

// GetMarBwUl returns the MarBwUl field value if set, zero value otherwise.
func (o *MediaComponent) GetMarBwUl() string {
	if o == nil || IsNil(o.MarBwUl) {
		var ret string
		return ret
	}
	return *o.MarBwUl
}

// GetMarBwUlOk returns a tuple with the MarBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMarBwUlOk() (*string, bool) {
	if o == nil || IsNil(o.MarBwUl) {
		return nil, false
	}
	return o.MarBwUl, true
}

// HasMarBwUl returns a boolean if a field has been set.
func (o *MediaComponent) HasMarBwUl() bool {
	if o != nil && !IsNil(o.MarBwUl) {
		return true
	}

	return false
}

// SetMarBwUl gets a reference to the given string and assigns it to the MarBwUl field.
func (o *MediaComponent) SetMarBwUl(v string) {
	o.MarBwUl = &v
}

// GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaComponent) GetMaxPacketLossRateDl() int32 {
	if o == nil || IsNil(o.MaxPacketLossRateDl.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPacketLossRateDl.Get()
}

// GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaComponent) GetMaxPacketLossRateDlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPacketLossRateDl.Get(), o.MaxPacketLossRateDl.IsSet()
}

// HasMaxPacketLossRateDl returns a boolean if a field has been set.
func (o *MediaComponent) HasMaxPacketLossRateDl() bool {
	if o != nil && o.MaxPacketLossRateDl.IsSet() {
		return true
	}

	return false
}

// SetMaxPacketLossRateDl gets a reference to the given NullableInt32 and assigns it to the MaxPacketLossRateDl field.
func (o *MediaComponent) SetMaxPacketLossRateDl(v int32) {
	o.MaxPacketLossRateDl.Set(&v)
}

// SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil
func (o *MediaComponent) SetMaxPacketLossRateDlNil() {
	o.MaxPacketLossRateDl.Set(nil)
}

// UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
func (o *MediaComponent) UnsetMaxPacketLossRateDl() {
	o.MaxPacketLossRateDl.Unset()
}

// GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaComponent) GetMaxPacketLossRateUl() int32 {
	if o == nil || IsNil(o.MaxPacketLossRateUl.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPacketLossRateUl.Get()
}

// GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaComponent) GetMaxPacketLossRateUlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPacketLossRateUl.Get(), o.MaxPacketLossRateUl.IsSet()
}

// HasMaxPacketLossRateUl returns a boolean if a field has been set.
func (o *MediaComponent) HasMaxPacketLossRateUl() bool {
	if o != nil && o.MaxPacketLossRateUl.IsSet() {
		return true
	}

	return false
}

// SetMaxPacketLossRateUl gets a reference to the given NullableInt32 and assigns it to the MaxPacketLossRateUl field.
func (o *MediaComponent) SetMaxPacketLossRateUl(v int32) {
	o.MaxPacketLossRateUl.Set(&v)
}

// SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil
func (o *MediaComponent) SetMaxPacketLossRateUlNil() {
	o.MaxPacketLossRateUl.Set(nil)
}

// UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
func (o *MediaComponent) UnsetMaxPacketLossRateUl() {
	o.MaxPacketLossRateUl.Unset()
}

// GetMaxSuppBwDl returns the MaxSuppBwDl field value if set, zero value otherwise.
func (o *MediaComponent) GetMaxSuppBwDl() string {
	if o == nil || IsNil(o.MaxSuppBwDl) {
		var ret string
		return ret
	}
	return *o.MaxSuppBwDl
}

// GetMaxSuppBwDlOk returns a tuple with the MaxSuppBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMaxSuppBwDlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSuppBwDl) {
		return nil, false
	}
	return o.MaxSuppBwDl, true
}

// HasMaxSuppBwDl returns a boolean if a field has been set.
func (o *MediaComponent) HasMaxSuppBwDl() bool {
	if o != nil && !IsNil(o.MaxSuppBwDl) {
		return true
	}

	return false
}

// SetMaxSuppBwDl gets a reference to the given string and assigns it to the MaxSuppBwDl field.
func (o *MediaComponent) SetMaxSuppBwDl(v string) {
	o.MaxSuppBwDl = &v
}

// GetMaxSuppBwUl returns the MaxSuppBwUl field value if set, zero value otherwise.
func (o *MediaComponent) GetMaxSuppBwUl() string {
	if o == nil || IsNil(o.MaxSuppBwUl) {
		var ret string
		return ret
	}
	return *o.MaxSuppBwUl
}

// GetMaxSuppBwUlOk returns a tuple with the MaxSuppBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMaxSuppBwUlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSuppBwUl) {
		return nil, false
	}
	return o.MaxSuppBwUl, true
}

// HasMaxSuppBwUl returns a boolean if a field has been set.
func (o *MediaComponent) HasMaxSuppBwUl() bool {
	if o != nil && !IsNil(o.MaxSuppBwUl) {
		return true
	}

	return false
}

// SetMaxSuppBwUl gets a reference to the given string and assigns it to the MaxSuppBwUl field.
func (o *MediaComponent) SetMaxSuppBwUl(v string) {
	o.MaxSuppBwUl = &v
}

// GetMedCompN returns the MedCompN field value
func (o *MediaComponent) GetMedCompN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MedCompN
}

// GetMedCompNOk returns a tuple with the MedCompN field value
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMedCompNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MedCompN, true
}

// SetMedCompN sets field value
func (o *MediaComponent) SetMedCompN(v int32) {
	o.MedCompN = v
}

// GetMedSubComps returns the MedSubComps field value if set, zero value otherwise.
func (o *MediaComponent) GetMedSubComps() map[string]MediaSubComponent {
	if o == nil || IsNil(o.MedSubComps) {
		var ret map[string]MediaSubComponent
		return ret
	}
	return *o.MedSubComps
}

// GetMedSubCompsOk returns a tuple with the MedSubComps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMedSubCompsOk() (*map[string]MediaSubComponent, bool) {
	if o == nil || IsNil(o.MedSubComps) {
		return nil, false
	}
	return o.MedSubComps, true
}

// HasMedSubComps returns a boolean if a field has been set.
func (o *MediaComponent) HasMedSubComps() bool {
	if o != nil && !IsNil(o.MedSubComps) {
		return true
	}

	return false
}

// SetMedSubComps gets a reference to the given map[string]MediaSubComponent and assigns it to the MedSubComps field.
func (o *MediaComponent) SetMedSubComps(v map[string]MediaSubComponent) {
	o.MedSubComps = &v
}

// GetMedType returns the MedType field value if set, zero value otherwise.
func (o *MediaComponent) GetMedType() string {
	if o == nil || IsNil(o.MedType) {
		var ret string
		return ret
	}
	return *o.MedType
}

// GetMedTypeOk returns a tuple with the MedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMedTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MedType) {
		return nil, false
	}
	return o.MedType, true
}

// HasMedType returns a boolean if a field has been set.
func (o *MediaComponent) HasMedType() bool {
	if o != nil && !IsNil(o.MedType) {
		return true
	}

	return false
}

// SetMedType gets a reference to the given MediaType and assigns it to the MedType field.
func (o *MediaComponent) SetMedType(v string) {
	o.MedType = &v
}

// GetMinDesBwDl returns the MinDesBwDl field value if set, zero value otherwise.
func (o *MediaComponent) GetMinDesBwDl() string {
	if o == nil || IsNil(o.MinDesBwDl) {
		var ret string
		return ret
	}
	return *o.MinDesBwDl
}

// GetMinDesBwDlOk returns a tuple with the MinDesBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMinDesBwDlOk() (*string, bool) {
	if o == nil || IsNil(o.MinDesBwDl) {
		return nil, false
	}
	return o.MinDesBwDl, true
}

// HasMinDesBwDl returns a boolean if a field has been set.
func (o *MediaComponent) HasMinDesBwDl() bool {
	if o != nil && !IsNil(o.MinDesBwDl) {
		return true
	}

	return false
}

// SetMinDesBwDl gets a reference to the given string and assigns it to the MinDesBwDl field.
func (o *MediaComponent) SetMinDesBwDl(v string) {
	o.MinDesBwDl = &v
}

// GetMinDesBwUl returns the MinDesBwUl field value if set, zero value otherwise.
func (o *MediaComponent) GetMinDesBwUl() string {
	if o == nil || IsNil(o.MinDesBwUl) {
		var ret string
		return ret
	}
	return *o.MinDesBwUl
}

// GetMinDesBwUlOk returns a tuple with the MinDesBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMinDesBwUlOk() (*string, bool) {
	if o == nil || IsNil(o.MinDesBwUl) {
		return nil, false
	}
	return o.MinDesBwUl, true
}

// HasMinDesBwUl returns a boolean if a field has been set.
func (o *MediaComponent) HasMinDesBwUl() bool {
	if o != nil && !IsNil(o.MinDesBwUl) {
		return true
	}

	return false
}

// SetMinDesBwUl gets a reference to the given string and assigns it to the MinDesBwUl field.
func (o *MediaComponent) SetMinDesBwUl(v string) {
	o.MinDesBwUl = &v
}

// GetMirBwDl returns the MirBwDl field value if set, zero value otherwise.
func (o *MediaComponent) GetMirBwDl() string {
	if o == nil || IsNil(o.MirBwDl) {
		var ret string
		return ret
	}
	return *o.MirBwDl
}

// GetMirBwDlOk returns a tuple with the MirBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMirBwDlOk() (*string, bool) {
	if o == nil || IsNil(o.MirBwDl) {
		return nil, false
	}
	return o.MirBwDl, true
}

// HasMirBwDl returns a boolean if a field has been set.
func (o *MediaComponent) HasMirBwDl() bool {
	if o != nil && !IsNil(o.MirBwDl) {
		return true
	}

	return false
}

// SetMirBwDl gets a reference to the given string and assigns it to the MirBwDl field.
func (o *MediaComponent) SetMirBwDl(v string) {
	o.MirBwDl = &v
}

// GetMirBwUl returns the MirBwUl field value if set, zero value otherwise.
func (o *MediaComponent) GetMirBwUl() string {
	if o == nil || IsNil(o.MirBwUl) {
		var ret string
		return ret
	}
	return *o.MirBwUl
}

// GetMirBwUlOk returns a tuple with the MirBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetMirBwUlOk() (*string, bool) {
	if o == nil || IsNil(o.MirBwUl) {
		return nil, false
	}
	return o.MirBwUl, true
}

// HasMirBwUl returns a boolean if a field has been set.
func (o *MediaComponent) HasMirBwUl() bool {
	if o != nil && !IsNil(o.MirBwUl) {
		return true
	}

	return false
}

// SetMirBwUl gets a reference to the given string and assigns it to the MirBwUl field.
func (o *MediaComponent) SetMirBwUl(v string) {
	o.MirBwUl = &v
}

// GetPreemptCap returns the PreemptCap field value if set, zero value otherwise.
func (o *MediaComponent) GetPreemptCap() PreemptionCapability {
	if o == nil || IsNil(o.PreemptCap) {
		var ret PreemptionCapability
		return ret
	}
	return *o.PreemptCap
}

// GetPreemptCapOk returns a tuple with the PreemptCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPreemptCapOk() (*PreemptionCapability, bool) {
	if o == nil || IsNil(o.PreemptCap) {
		return nil, false
	}
	return o.PreemptCap, true
}

// HasPreemptCap returns a boolean if a field has been set.
func (o *MediaComponent) HasPreemptCap() bool {
	if o != nil && !IsNil(o.PreemptCap) {
		return true
	}

	return false
}

// SetPreemptCap gets a reference to the given PreemptionCapability and assigns it to the PreemptCap field.
func (o *MediaComponent) SetPreemptCap(v PreemptionCapability) {
	o.PreemptCap = &v
}

// GetPreemptVuln returns the PreemptVuln field value if set, zero value otherwise.
func (o *MediaComponent) GetPreemptVuln() PreemptionVulnerability {
	if o == nil || IsNil(o.PreemptVuln) {
		var ret PreemptionVulnerability
		return ret
	}
	return *o.PreemptVuln
}

// GetPreemptVulnOk returns a tuple with the PreemptVuln field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPreemptVulnOk() (*PreemptionVulnerability, bool) {
	if o == nil || IsNil(o.PreemptVuln) {
		return nil, false
	}
	return o.PreemptVuln, true
}

// HasPreemptVuln returns a boolean if a field has been set.
func (o *MediaComponent) HasPreemptVuln() bool {
	if o != nil && !IsNil(o.PreemptVuln) {
		return true
	}

	return false
}

// SetPreemptVuln gets a reference to the given PreemptionVulnerability and assigns it to the PreemptVuln field.
func (o *MediaComponent) SetPreemptVuln(v PreemptionVulnerability) {
	o.PreemptVuln = &v
}

// GetPrioSharingInd returns the PrioSharingInd field value if set, zero value otherwise.
func (o *MediaComponent) GetPrioSharingInd() PrioritySharingIndicator {
	if o == nil || IsNil(o.PrioSharingInd) {
		var ret PrioritySharingIndicator
		return ret
	}
	return *o.PrioSharingInd
}

// GetPrioSharingIndOk returns a tuple with the PrioSharingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPrioSharingIndOk() (*PrioritySharingIndicator, bool) {
	if o == nil || IsNil(o.PrioSharingInd) {
		return nil, false
	}
	return o.PrioSharingInd, true
}

// HasPrioSharingInd returns a boolean if a field has been set.
func (o *MediaComponent) HasPrioSharingInd() bool {
	if o != nil && !IsNil(o.PrioSharingInd) {
		return true
	}

	return false
}

// SetPrioSharingInd gets a reference to the given PrioritySharingIndicator and assigns it to the PrioSharingInd field.
func (o *MediaComponent) SetPrioSharingInd(v PrioritySharingIndicator) {
	o.PrioSharingInd = &v
}

// GetResPrio returns the ResPrio field value if set, zero value otherwise.
func (o *MediaComponent) GetResPrio() ReservPriority {
	if o == nil || IsNil(o.ResPrio) {
		var ret ReservPriority
		return ret
	}
	return *o.ResPrio
}

// GetResPrioOk returns a tuple with the ResPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetResPrioOk() (*ReservPriority, bool) {
	if o == nil || IsNil(o.ResPrio) {
		return nil, false
	}
	return o.ResPrio, true
}

// HasResPrio returns a boolean if a field has been set.
func (o *MediaComponent) HasResPrio() bool {
	if o != nil && !IsNil(o.ResPrio) {
		return true
	}

	return false
}

// SetResPrio gets a reference to the given ReservPriority and assigns it to the ResPrio field.
func (o *MediaComponent) SetResPrio(v ReservPriority) {
	o.ResPrio = &v
}

// GetRrBw returns the RrBw field value if set, zero value otherwise.
func (o *MediaComponent) GetRrBw() string {
	if o == nil || IsNil(o.RrBw) {
		var ret string
		return ret
	}
	return *o.RrBw
}

// GetRrBwOk returns a tuple with the RrBw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetRrBwOk() (*string, bool) {
	if o == nil || IsNil(o.RrBw) {
		return nil, false
	}
	return o.RrBw, true
}

// HasRrBw returns a boolean if a field has been set.
func (o *MediaComponent) HasRrBw() bool {
	if o != nil && !IsNil(o.RrBw) {
		return true
	}

	return false
}

// SetRrBw gets a reference to the given string and assigns it to the RrBw field.
func (o *MediaComponent) SetRrBw(v string) {
	o.RrBw = &v
}

// GetRsBw returns the RsBw field value if set, zero value otherwise.
func (o *MediaComponent) GetRsBw() string {
	if o == nil || IsNil(o.RsBw) {
		var ret string
		return ret
	}
	return *o.RsBw
}

// GetRsBwOk returns a tuple with the RsBw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetRsBwOk() (*string, bool) {
	if o == nil || IsNil(o.RsBw) {
		return nil, false
	}
	return o.RsBw, true
}

// HasRsBw returns a boolean if a field has been set.
func (o *MediaComponent) HasRsBw() bool {
	if o != nil && !IsNil(o.RsBw) {
		return true
	}

	return false
}

// SetRsBw gets a reference to the given string and assigns it to the RsBw field.
func (o *MediaComponent) SetRsBw(v string) {
	o.RsBw = &v
}

// GetSharingKeyDl returns the SharingKeyDl field value if set, zero value otherwise.
func (o *MediaComponent) GetSharingKeyDl() int32 {
	if o == nil || IsNil(o.SharingKeyDl) {
		var ret int32
		return ret
	}
	return *o.SharingKeyDl
}

// GetSharingKeyDlOk returns a tuple with the SharingKeyDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetSharingKeyDlOk() (*int32, bool) {
	if o == nil || IsNil(o.SharingKeyDl) {
		return nil, false
	}
	return o.SharingKeyDl, true
}

// HasSharingKeyDl returns a boolean if a field has been set.
func (o *MediaComponent) HasSharingKeyDl() bool {
	if o != nil && !IsNil(o.SharingKeyDl) {
		return true
	}

	return false
}

// SetSharingKeyDl gets a reference to the given int32 and assigns it to the SharingKeyDl field.
func (o *MediaComponent) SetSharingKeyDl(v int32) {
	o.SharingKeyDl = &v
}

// GetSharingKeyUl returns the SharingKeyUl field value if set, zero value otherwise.
func (o *MediaComponent) GetSharingKeyUl() int32 {
	if o == nil || IsNil(o.SharingKeyUl) {
		var ret int32
		return ret
	}
	return *o.SharingKeyUl
}

// GetSharingKeyUlOk returns a tuple with the SharingKeyUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetSharingKeyUlOk() (*int32, bool) {
	if o == nil || IsNil(o.SharingKeyUl) {
		return nil, false
	}
	return o.SharingKeyUl, true
}

// HasSharingKeyUl returns a boolean if a field has been set.
func (o *MediaComponent) HasSharingKeyUl() bool {
	if o != nil && !IsNil(o.SharingKeyUl) {
		return true
	}

	return false
}

// SetSharingKeyUl gets a reference to the given int32 and assigns it to the SharingKeyUl field.
func (o *MediaComponent) SetSharingKeyUl(v int32) {
	o.SharingKeyUl = &v
}

// GetTsnQos returns the TsnQos field value if set, zero value otherwise.
func (o *MediaComponent) GetTsnQos() TsnQosContainer {
	if o == nil || IsNil(o.TsnQos) {
		var ret TsnQosContainer
		return ret
	}
	return *o.TsnQos
}

// GetTsnQosOk returns a tuple with the TsnQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetTsnQosOk() (*TsnQosContainer, bool) {
	if o == nil || IsNil(o.TsnQos) {
		return nil, false
	}
	return o.TsnQos, true
}

// HasTsnQos returns a boolean if a field has been set.
func (o *MediaComponent) HasTsnQos() bool {
	if o != nil && !IsNil(o.TsnQos) {
		return true
	}

	return false
}

// SetTsnQos gets a reference to the given TsnQosContainer and assigns it to the TsnQos field.
func (o *MediaComponent) SetTsnQos(v TsnQosContainer) {
	o.TsnQos = &v
}

// GetTscaiInputDl returns the TscaiInputDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaComponent) GetTscaiInputDl() TscaiInputContainer {
	if o == nil || IsNil(o.TscaiInputDl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputDl.Get()
}

// GetTscaiInputDlOk returns a tuple with the TscaiInputDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaComponent) GetTscaiInputDlOk() (*TscaiInputContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscaiInputDl.Get(), o.TscaiInputDl.IsSet()
}

// HasTscaiInputDl returns a boolean if a field has been set.
func (o *MediaComponent) HasTscaiInputDl() bool {
	if o != nil && o.TscaiInputDl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputDl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputDl field.
func (o *MediaComponent) SetTscaiInputDl(v TscaiInputContainer) {
	o.TscaiInputDl.Set(&v)
}

// SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil
func (o *MediaComponent) SetTscaiInputDlNil() {
	o.TscaiInputDl.Set(nil)
}

// UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
func (o *MediaComponent) UnsetTscaiInputDl() {
	o.TscaiInputDl.Unset()
}

// GetTscaiInputUl returns the TscaiInputUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaComponent) GetTscaiInputUl() TscaiInputContainer {
	if o == nil || IsNil(o.TscaiInputUl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputUl.Get()
}

// GetTscaiInputUlOk returns a tuple with the TscaiInputUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaComponent) GetTscaiInputUlOk() (*TscaiInputContainer, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscaiInputUl.Get(), o.TscaiInputUl.IsSet()
}

// HasTscaiInputUl returns a boolean if a field has been set.
func (o *MediaComponent) HasTscaiInputUl() bool {
	if o != nil && o.TscaiInputUl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputUl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputUl field.
func (o *MediaComponent) SetTscaiInputUl(v TscaiInputContainer) {
	o.TscaiInputUl.Set(&v)
}

// SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil
func (o *MediaComponent) SetTscaiInputUlNil() {
	o.TscaiInputUl.Set(nil)
}

// UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil
func (o *MediaComponent) UnsetTscaiInputUl() {
	o.TscaiInputUl.Unset()
}

// GetTscaiTimeDom returns the TscaiTimeDom field value if set, zero value otherwise.
func (o *MediaComponent) GetTscaiTimeDom() int32 {
	if o == nil || IsNil(o.TscaiTimeDom) {
		var ret int32
		return ret
	}
	return *o.TscaiTimeDom
}

// GetTscaiTimeDomOk returns a tuple with the TscaiTimeDom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetTscaiTimeDomOk() (*int32, bool) {
	if o == nil || IsNil(o.TscaiTimeDom) {
		return nil, false
	}
	return o.TscaiTimeDom, true
}

// HasTscaiTimeDom returns a boolean if a field has been set.
func (o *MediaComponent) HasTscaiTimeDom() bool {
	if o != nil && !IsNil(o.TscaiTimeDom) {
		return true
	}

	return false
}

// SetTscaiTimeDom gets a reference to the given int32 and assigns it to the TscaiTimeDom field.
func (o *MediaComponent) SetTscaiTimeDom(v int32) {
	o.TscaiTimeDom = &v
}

// GetCapBatAdaptation returns the CapBatAdaptation field value if set, zero value otherwise.
func (o *MediaComponent) GetCapBatAdaptation() bool {
	if o == nil || IsNil(o.CapBatAdaptation) {
		var ret bool
		return ret
	}
	return *o.CapBatAdaptation
}

// GetCapBatAdaptationOk returns a tuple with the CapBatAdaptation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetCapBatAdaptationOk() (*bool, bool) {
	if o == nil || IsNil(o.CapBatAdaptation) {
		return nil, false
	}
	return o.CapBatAdaptation, true
}

// HasCapBatAdaptation returns a boolean if a field has been set.
func (o *MediaComponent) HasCapBatAdaptation() bool {
	if o != nil && !IsNil(o.CapBatAdaptation) {
		return true
	}

	return false
}

// SetCapBatAdaptation gets a reference to the given bool and assigns it to the CapBatAdaptation field.
func (o *MediaComponent) SetCapBatAdaptation(v bool) {
	o.CapBatAdaptation = &v
}

// GetRTLatencyInd returns the RTLatencyInd field value if set, zero value otherwise.
func (o *MediaComponent) GetRTLatencyInd() bool {
	if o == nil || IsNil(o.RTLatencyInd) {
		var ret bool
		return ret
	}
	return *o.RTLatencyInd
}

// GetRTLatencyIndOk returns a tuple with the RTLatencyInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetRTLatencyIndOk() (*bool, bool) {
	if o == nil || IsNil(o.RTLatencyInd) {
		return nil, false
	}
	return o.RTLatencyInd, true
}

// HasRTLatencyInd returns a boolean if a field has been set.
func (o *MediaComponent) HasRTLatencyInd() bool {
	if o != nil && !IsNil(o.RTLatencyInd) {
		return true
	}

	return false
}

// SetRTLatencyInd gets a reference to the given bool and assigns it to the RTLatencyInd field.
func (o *MediaComponent) SetRTLatencyInd(v bool) {
	o.RTLatencyInd = &v
}

// GetPduSetQosDl returns the PduSetQosDl field value if set, zero value otherwise.
func (o *MediaComponent) GetPduSetQosDl() PduSetQosPara {
	if o == nil || IsNil(o.PduSetQosDl) {
		var ret PduSetQosPara
		return ret
	}
	return *o.PduSetQosDl
}

// GetPduSetQosDlOk returns a tuple with the PduSetQosDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPduSetQosDlOk() (*PduSetQosPara, bool) {
	if o == nil || IsNil(o.PduSetQosDl) {
		return nil, false
	}
	return o.PduSetQosDl, true
}

// HasPduSetQosDl returns a boolean if a field has been set.
func (o *MediaComponent) HasPduSetQosDl() bool {
	if o != nil && !IsNil(o.PduSetQosDl) {
		return true
	}

	return false
}

// SetPduSetQosDl gets a reference to the given PduSetQosPara and assigns it to the PduSetQosDl field.
func (o *MediaComponent) SetPduSetQosDl(v PduSetQosPara) {
	o.PduSetQosDl = &v
}

// GetPduSetQosUl returns the PduSetQosUl field value if set, zero value otherwise.
func (o *MediaComponent) GetPduSetQosUl() PduSetQosPara {
	if o == nil || IsNil(o.PduSetQosUl) {
		var ret PduSetQosPara
		return ret
	}
	return *o.PduSetQosUl
}

// GetPduSetQosUlOk returns a tuple with the PduSetQosUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPduSetQosUlOk() (*PduSetQosPara, bool) {
	if o == nil || IsNil(o.PduSetQosUl) {
		return nil, false
	}
	return o.PduSetQosUl, true
}

// HasPduSetQosUl returns a boolean if a field has been set.
func (o *MediaComponent) HasPduSetQosUl() bool {
	if o != nil && !IsNil(o.PduSetQosUl) {
		return true
	}

	return false
}

// SetPduSetQosUl gets a reference to the given PduSetQosPara and assigns it to the PduSetQosUl field.
func (o *MediaComponent) SetPduSetQosUl(v PduSetQosPara) {
	o.PduSetQosUl = &v
}

// GetProtoDescDl returns the ProtoDescDl field value if set, zero value otherwise.
func (o *MediaComponent) GetProtoDescDl() ProtocolDescription {
	if o == nil || IsNil(o.ProtoDescDl) {
		var ret ProtocolDescription
		return ret
	}
	return *o.ProtoDescDl
}

// GetProtoDescDlOk returns a tuple with the ProtoDescDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetProtoDescDlOk() (*ProtocolDescription, bool) {
	if o == nil || IsNil(o.ProtoDescDl) {
		return nil, false
	}
	return o.ProtoDescDl, true
}

// HasProtoDescDl returns a boolean if a field has been set.
func (o *MediaComponent) HasProtoDescDl() bool {
	if o != nil && !IsNil(o.ProtoDescDl) {
		return true
	}

	return false
}

// SetProtoDescDl gets a reference to the given ProtocolDescription and assigns it to the ProtoDescDl field.
func (o *MediaComponent) SetProtoDescDl(v ProtocolDescription) {
	o.ProtoDescDl = &v
}

// GetProtoDescUl returns the ProtoDescUl field value if set, zero value otherwise.
func (o *MediaComponent) GetProtoDescUl() ProtocolDescription {
	if o == nil || IsNil(o.ProtoDescUl) {
		var ret ProtocolDescription
		return ret
	}
	return *o.ProtoDescUl
}

// GetProtoDescUlOk returns a tuple with the ProtoDescUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetProtoDescUlOk() (*ProtocolDescription, bool) {
	if o == nil || IsNil(o.ProtoDescUl) {
		return nil, false
	}
	return o.ProtoDescUl, true
}

// HasProtoDescUl returns a boolean if a field has been set.
func (o *MediaComponent) HasProtoDescUl() bool {
	if o != nil && !IsNil(o.ProtoDescUl) {
		return true
	}

	return false
}

// SetProtoDescUl gets a reference to the given ProtocolDescription and assigns it to the ProtoDescUl field.
func (o *MediaComponent) SetProtoDescUl(v ProtocolDescription) {
	o.ProtoDescUl = &v
}

// GetPeriodUl returns the PeriodUl field value if set, zero value otherwise.
func (o *MediaComponent) GetPeriodUl() int32 {
	if o == nil || IsNil(o.PeriodUl) {
		var ret int32
		return ret
	}
	return *o.PeriodUl
}

// GetPeriodUlOk returns a tuple with the PeriodUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPeriodUlOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodUl) {
		return nil, false
	}
	return o.PeriodUl, true
}

// HasPeriodUl returns a boolean if a field has been set.
func (o *MediaComponent) HasPeriodUl() bool {
	if o != nil && !IsNil(o.PeriodUl) {
		return true
	}

	return false
}

// SetPeriodUl gets a reference to the given int32 and assigns it to the PeriodUl field.
func (o *MediaComponent) SetPeriodUl(v int32) {
	o.PeriodUl = &v
}

// GetPeriodDl returns the PeriodDl field value if set, zero value otherwise.
func (o *MediaComponent) GetPeriodDl() int32 {
	if o == nil || IsNil(o.PeriodDl) {
		var ret int32
		return ret
	}
	return *o.PeriodDl
}

// GetPeriodDlOk returns a tuple with the PeriodDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetPeriodDlOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodDl) {
		return nil, false
	}
	return o.PeriodDl, true
}

// HasPeriodDl returns a boolean if a field has been set.
func (o *MediaComponent) HasPeriodDl() bool {
	if o != nil && !IsNil(o.PeriodDl) {
		return true
	}

	return false
}

// SetPeriodDl gets a reference to the given int32 and assigns it to the PeriodDl field.
func (o *MediaComponent) SetPeriodDl(v int32) {
	o.PeriodDl = &v
}

// GetL4sInd returns the L4sInd field value if set, zero value otherwise.
func (o *MediaComponent) GetL4sInd() UplinkDownlinkSupport {
	if o == nil || IsNil(o.L4sInd) {
		var ret UplinkDownlinkSupport
		return ret
	}
	return *o.L4sInd
}

// GetL4sIndOk returns a tuple with the L4sInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaComponent) GetL4sIndOk() (*UplinkDownlinkSupport, bool) {
	if o == nil || IsNil(o.L4sInd) {
		return nil, false
	}
	return o.L4sInd, true
}

// HasL4sInd returns a boolean if a field has been set.
func (o *MediaComponent) HasL4sInd() bool {
	if o != nil && !IsNil(o.L4sInd) {
		return true
	}

	return false
}

// SetL4sInd gets a reference to the given UplinkDownlinkSupport and assigns it to the L4sInd field.
func (o *MediaComponent) SetL4sInd(v UplinkDownlinkSupport) {
	o.L4sInd = &v
}

func (o MediaComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	if !IsNil(o.AfRoutReq) {
		toSerialize["afRoutReq"] = o.AfRoutReq
	}
	if o.AfSfcReq.IsSet() {
		toSerialize["afSfcReq"] = o.AfSfcReq.Get()
	}
	if !IsNil(o.QosReference) {
		toSerialize["qosReference"] = o.QosReference
	}
	if !IsNil(o.DisUeNotif) {
		toSerialize["disUeNotif"] = o.DisUeNotif
	}
	if !IsNil(o.AltSerReqs) {
		toSerialize["altSerReqs"] = o.AltSerReqs
	}
	if !IsNil(o.AltSerReqsData) {
		toSerialize["altSerReqsData"] = o.AltSerReqsData
	}
	if !IsNil(o.ContVer) {
		toSerialize["contVer"] = o.ContVer
	}
	if !IsNil(o.Codecs) {
		toSerialize["codecs"] = o.Codecs
	}
	if !IsNil(o.DesMaxLatency) {
		toSerialize["desMaxLatency"] = o.DesMaxLatency
	}
	if !IsNil(o.DesMaxLoss) {
		toSerialize["desMaxLoss"] = o.DesMaxLoss
	}
	if !IsNil(o.FlusId) {
		toSerialize["flusId"] = o.FlusId
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
	if o.MaxPacketLossRateDl.IsSet() {
		toSerialize["maxPacketLossRateDl"] = o.MaxPacketLossRateDl.Get()
	}
	if o.MaxPacketLossRateUl.IsSet() {
		toSerialize["maxPacketLossRateUl"] = o.MaxPacketLossRateUl.Get()
	}
	if !IsNil(o.MaxSuppBwDl) {
		toSerialize["maxSuppBwDl"] = o.MaxSuppBwDl
	}
	if !IsNil(o.MaxSuppBwUl) {
		toSerialize["maxSuppBwUl"] = o.MaxSuppBwUl
	}
	toSerialize["medCompN"] = o.MedCompN
	if !IsNil(o.MedSubComps) {
		toSerialize["medSubComps"] = o.MedSubComps
	}
	if !IsNil(o.MedType) {
		toSerialize["medType"] = o.MedType
	}
	if !IsNil(o.MinDesBwDl) {
		toSerialize["minDesBwDl"] = o.MinDesBwDl
	}
	if !IsNil(o.MinDesBwUl) {
		toSerialize["minDesBwUl"] = o.MinDesBwUl
	}
	if !IsNil(o.MirBwDl) {
		toSerialize["mirBwDl"] = o.MirBwDl
	}
	if !IsNil(o.MirBwUl) {
		toSerialize["mirBwUl"] = o.MirBwUl
	}
	if !IsNil(o.PreemptCap) {
		toSerialize["preemptCap"] = o.PreemptCap
	}
	if !IsNil(o.PreemptVuln) {
		toSerialize["preemptVuln"] = o.PreemptVuln
	}
	if !IsNil(o.PrioSharingInd) {
		toSerialize["prioSharingInd"] = o.PrioSharingInd
	}
	if !IsNil(o.ResPrio) {
		toSerialize["resPrio"] = o.ResPrio
	}
	if !IsNil(o.RrBw) {
		toSerialize["rrBw"] = o.RrBw
	}
	if !IsNil(o.RsBw) {
		toSerialize["rsBw"] = o.RsBw
	}
	if !IsNil(o.SharingKeyDl) {
		toSerialize["sharingKeyDl"] = o.SharingKeyDl
	}
	if !IsNil(o.SharingKeyUl) {
		toSerialize["sharingKeyUl"] = o.SharingKeyUl
	}
	if !IsNil(o.TsnQos) {
		toSerialize["tsnQos"] = o.TsnQos
	}
	if o.TscaiInputDl.IsSet() {
		toSerialize["tscaiInputDl"] = o.TscaiInputDl.Get()
	}
	if o.TscaiInputUl.IsSet() {
		toSerialize["tscaiInputUl"] = o.TscaiInputUl.Get()
	}
	if !IsNil(o.TscaiTimeDom) {
		toSerialize["tscaiTimeDom"] = o.TscaiTimeDom
	}
	if !IsNil(o.CapBatAdaptation) {
		toSerialize["capBatAdaptation"] = o.CapBatAdaptation
	}
	if !IsNil(o.RTLatencyInd) {
		toSerialize["rTLatencyInd"] = o.RTLatencyInd
	}
	if !IsNil(o.PduSetQosDl) {
		toSerialize["pduSetQosDl"] = o.PduSetQosDl
	}
	if !IsNil(o.PduSetQosUl) {
		toSerialize["pduSetQosUl"] = o.PduSetQosUl
	}
	if !IsNil(o.ProtoDescDl) {
		toSerialize["protoDescDl"] = o.ProtoDescDl
	}
	if !IsNil(o.ProtoDescUl) {
		toSerialize["protoDescUl"] = o.ProtoDescUl
	}
	if !IsNil(o.PeriodUl) {
		toSerialize["periodUl"] = o.PeriodUl
	}
	if !IsNil(o.PeriodDl) {
		toSerialize["periodDl"] = o.PeriodDl
	}
	if !IsNil(o.L4sInd) {
		toSerialize["l4sInd"] = o.L4sInd
	}
	return toSerialize, nil
}

func (o *MediaComponent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"medCompN",
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

	varMediaComponent := _MediaComponent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaComponent)

	if err != nil {
		return err
	}

	*o = MediaComponent(varMediaComponent)

	return err
}

type NullableMediaComponent struct {
	value *MediaComponent
	isSet bool
}

func (v NullableMediaComponent) Get() *MediaComponent {
	return v.value
}

func (v *NullableMediaComponent) Set(val *MediaComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaComponent(val *MediaComponent) *NullableMediaComponent {
	return &NullableMediaComponent{value: val, isSet: true}
}

func (v NullableMediaComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
