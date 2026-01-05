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

// checks if the AppSessionContextReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSessionContextReqData{}

// AppSessionContextReqData Identifies the service requirements of an Individual Application Session Context.
type AppSessionContextReqData struct {
	// Contains an AF application identifier.
	AfAppId *string
	// Application provided charging identifier allowing correlation of charging information.
	AfChargId *string
	AfReqData *AfRequestedData
	AfRoutReq *AfRoutingRequirement
	AfSfcReq  NullableAfSfcRequirement
	// Contains an identity of an application service provider.
	AspId *string
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId *string
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn     *string
	EvSubsc *EventsSubscReqData
	// Indication of MCPTT service request.
	McpttId *string
	// Indication of MCVideo service request.
	McVideoId *string
	// Contains media component information. The key of the map is the medCompN attribute.
	MedComponents *map[string]MediaComponent
	// This data type contains a multi-modal service identifier.
	MultiModalId *string
	IpDomain     *string
	MpsAction    *MpsAction
	// Indication of MPS service request.
	MpsId *string
	// Indication of MCS service request.
	McsId              *string
	PreemptControlInfo *PreemptionControlInformation
	// indicating a time in seconds.
	QosDuration *int32
	// indicating a time in seconds.
	QosInactInt   *int32
	ResPrio       *ReservPriority
	ServInfStatus *ServiceInfoStatus
	// String providing an URI formatted according to RFC 3986.
	NotifUri string
	// Contains values of the service URN and may include subservices.
	ServUrn   *string
	SliceInfo *Snssai
	// Contains an identity of a sponsor.
	SponId     *string
	SponStatus *SponsoringStatus
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	UeIpv4 *string
	UeIpv6 *Ipv6Addr
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	UeMac               *string
	TsnBridgeManCont    *BridgeManagementContainer
	TsnPortManContDstt  *PortManagementContainer
	TsnPortManContNwtts []PortManagementContainer
	// String providing an URI formatted according to RFC 3986.
	TscNotifUri *string
	// Correlation identifier for TSC management information notifications.
	TscNotifCorreId *string
}

type _AppSessionContextReqData AppSessionContextReqData

// NewAppSessionContextReqData instantiates a new AppSessionContextReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSessionContextReqData(notifUri string, suppFeat string) *AppSessionContextReqData {
	this := AppSessionContextReqData{}
	return &this
}

// NewAppSessionContextReqDataWithDefaults instantiates a new AppSessionContextReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSessionContextReqDataWithDefaults() *AppSessionContextReqData {
	this := AppSessionContextReqData{}
	return &this
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetAfAppId() string {
	if o == nil || IsNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetAfAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAfAppId() bool {
	if o != nil && !IsNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *AppSessionContextReqData) SetAfAppId(v string) {
	o.AfAppId = &v
}

// GetAfChargId returns the AfChargId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetAfChargId() string {
	if o == nil || IsNil(o.AfChargId) {
		var ret string
		return ret
	}
	return *o.AfChargId
}

// GetAfChargIdOk returns a tuple with the AfChargId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetAfChargIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfChargId) {
		return nil, false
	}
	return o.AfChargId, true
}

// HasAfChargId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAfChargId() bool {
	if o != nil && !IsNil(o.AfChargId) {
		return true
	}

	return false
}

// SetAfChargId gets a reference to the given string and assigns it to the AfChargId field.
func (o *AppSessionContextReqData) SetAfChargId(v string) {
	o.AfChargId = &v
}

// GetAfReqData returns the AfReqData field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetAfReqData() AfRequestedData {
	if o == nil || IsNil(o.AfReqData) {
		var ret AfRequestedData
		return ret
	}
	return *o.AfReqData
}

// GetAfReqDataOk returns a tuple with the AfReqData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetAfReqDataOk() (*AfRequestedData, bool) {
	if o == nil || IsNil(o.AfReqData) {
		return nil, false
	}
	return o.AfReqData, true
}

// HasAfReqData returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAfReqData() bool {
	if o != nil && !IsNil(o.AfReqData) {
		return true
	}

	return false
}

// SetAfReqData gets a reference to the given AfRequestedData and assigns it to the AfReqData field.
func (o *AppSessionContextReqData) SetAfReqData(v AfRequestedData) {
	o.AfReqData = &v
}

// GetAfRoutReq returns the AfRoutReq field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetAfRoutReq() AfRoutingRequirement {
	if o == nil || IsNil(o.AfRoutReq) {
		var ret AfRoutingRequirement
		return ret
	}
	return *o.AfRoutReq
}

// GetAfRoutReqOk returns a tuple with the AfRoutReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetAfRoutReqOk() (*AfRoutingRequirement, bool) {
	if o == nil || IsNil(o.AfRoutReq) {
		return nil, false
	}
	return o.AfRoutReq, true
}

// HasAfRoutReq returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAfRoutReq() bool {
	if o != nil && !IsNil(o.AfRoutReq) {
		return true
	}

	return false
}

// SetAfRoutReq gets a reference to the given AfRoutingRequirement and assigns it to the AfRoutReq field.
func (o *AppSessionContextReqData) SetAfRoutReq(v AfRoutingRequirement) {
	o.AfRoutReq = &v
}

// GetAfSfcReq returns the AfSfcReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSessionContextReqData) GetAfSfcReq() AfSfcRequirement {
	if o == nil || IsNil(o.AfSfcReq.Get()) {
		var ret AfSfcRequirement
		return ret
	}
	return *o.AfSfcReq.Get()
}

// GetAfSfcReqOk returns a tuple with the AfSfcReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSessionContextReqData) GetAfSfcReqOk() (*AfSfcRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AfSfcReq.Get(), o.AfSfcReq.IsSet()
}

// HasAfSfcReq returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAfSfcReq() bool {
	if o != nil && o.AfSfcReq.IsSet() {
		return true
	}

	return false
}

// SetAfSfcReq gets a reference to the given NullableAfSfcRequirement and assigns it to the AfSfcReq field.
func (o *AppSessionContextReqData) SetAfSfcReq(v AfSfcRequirement) {
	o.AfSfcReq.Set(&v)
}

// SetAfSfcReqNil sets the value for AfSfcReq to be an explicit nil
func (o *AppSessionContextReqData) SetAfSfcReqNil() {
	o.AfSfcReq.Set(nil)
}

// UnsetAfSfcReq ensures that no value is present for AfSfcReq, not even an explicit nil
func (o *AppSessionContextReqData) UnsetAfSfcReq() {
	o.AfSfcReq.Unset()
}

// GetAspId returns the AspId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetAspId() string {
	if o == nil || IsNil(o.AspId) {
		var ret string
		return ret
	}
	return *o.AspId
}

// GetAspIdOk returns a tuple with the AspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetAspIdOk() (*string, bool) {
	if o == nil || IsNil(o.AspId) {
		return nil, false
	}
	return o.AspId, true
}

// HasAspId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasAspId() bool {
	if o != nil && !IsNil(o.AspId) {
		return true
	}

	return false
}

// SetAspId gets a reference to the given string and assigns it to the AspId field.
func (o *AppSessionContextReqData) SetAspId(v string) {
	o.AspId = &v
}

// GetBdtRefId returns the BdtRefId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetBdtRefId() string {
	if o == nil || IsNil(o.BdtRefId) {
		var ret string
		return ret
	}
	return *o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetBdtRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.BdtRefId) {
		return nil, false
	}
	return o.BdtRefId, true
}

// HasBdtRefId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasBdtRefId() bool {
	if o != nil && !IsNil(o.BdtRefId) {
		return true
	}

	return false
}

// SetBdtRefId gets a reference to the given string and assigns it to the BdtRefId field.
func (o *AppSessionContextReqData) SetBdtRefId(v string) {
	o.BdtRefId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *AppSessionContextReqData) SetDnn(v string) {
	o.Dnn = &v
}

// GetEvSubsc returns the EvSubsc field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetEvSubsc() EventsSubscReqData {
	if o == nil || IsNil(o.EvSubsc) {
		var ret EventsSubscReqData
		return ret
	}
	return *o.EvSubsc
}

// GetEvSubscOk returns a tuple with the EvSubsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetEvSubscOk() (*EventsSubscReqData, bool) {
	if o == nil || IsNil(o.EvSubsc) {
		return nil, false
	}
	return o.EvSubsc, true
}

// HasEvSubsc returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasEvSubsc() bool {
	if o != nil && !IsNil(o.EvSubsc) {
		return true
	}

	return false
}

// SetEvSubsc gets a reference to the given EventsSubscReqData and assigns it to the EvSubsc field.
func (o *AppSessionContextReqData) SetEvSubsc(v EventsSubscReqData) {
	o.EvSubsc = &v
}

// GetMcpttId returns the McpttId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMcpttId() string {
	if o == nil || IsNil(o.McpttId) {
		var ret string
		return ret
	}
	return *o.McpttId
}

// GetMcpttIdOk returns a tuple with the McpttId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMcpttIdOk() (*string, bool) {
	if o == nil || IsNil(o.McpttId) {
		return nil, false
	}
	return o.McpttId, true
}

// HasMcpttId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMcpttId() bool {
	if o != nil && !IsNil(o.McpttId) {
		return true
	}

	return false
}

// SetMcpttId gets a reference to the given string and assigns it to the McpttId field.
func (o *AppSessionContextReqData) SetMcpttId(v string) {
	o.McpttId = &v
}

// GetMcVideoId returns the McVideoId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMcVideoId() string {
	if o == nil || IsNil(o.McVideoId) {
		var ret string
		return ret
	}
	return *o.McVideoId
}

// GetMcVideoIdOk returns a tuple with the McVideoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMcVideoIdOk() (*string, bool) {
	if o == nil || IsNil(o.McVideoId) {
		return nil, false
	}
	return o.McVideoId, true
}

// HasMcVideoId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMcVideoId() bool {
	if o != nil && !IsNil(o.McVideoId) {
		return true
	}

	return false
}

// SetMcVideoId gets a reference to the given string and assigns it to the McVideoId field.
func (o *AppSessionContextReqData) SetMcVideoId(v string) {
	o.McVideoId = &v
}

// GetMedComponents returns the MedComponents field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMedComponents() map[string]MediaComponent {
	if o == nil || IsNil(o.MedComponents) {
		var ret map[string]MediaComponent
		return ret
	}
	return *o.MedComponents
}

// GetMedComponentsOk returns a tuple with the MedComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMedComponentsOk() (*map[string]MediaComponent, bool) {
	if o == nil || IsNil(o.MedComponents) {
		return nil, false
	}
	return o.MedComponents, true
}

// HasMedComponents returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMedComponents() bool {
	if o != nil && !IsNil(o.MedComponents) {
		return true
	}

	return false
}

// SetMedComponents gets a reference to the given map[string]MediaComponent and assigns it to the MedComponents field.
func (o *AppSessionContextReqData) SetMedComponents(v map[string]MediaComponent) {
	o.MedComponents = &v
}

// GetMultiModalId returns the MultiModalId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMultiModalId() string {
	if o == nil || IsNil(o.MultiModalId) {
		var ret string
		return ret
	}
	return *o.MultiModalId
}

// GetMultiModalIdOk returns a tuple with the MultiModalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMultiModalIdOk() (*string, bool) {
	if o == nil || IsNil(o.MultiModalId) {
		return nil, false
	}
	return o.MultiModalId, true
}

// HasMultiModalId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMultiModalId() bool {
	if o != nil && !IsNil(o.MultiModalId) {
		return true
	}

	return false
}

// SetMultiModalId gets a reference to the given string and assigns it to the MultiModalId field.
func (o *AppSessionContextReqData) SetMultiModalId(v string) {
	o.MultiModalId = &v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetIpDomain() string {
	if o == nil || IsNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetIpDomainOk() (*string, bool) {
	if o == nil || IsNil(o.IpDomain) {
		return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasIpDomain() bool {
	if o != nil && !IsNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *AppSessionContextReqData) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetMpsAction returns the MpsAction field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMpsAction() MpsAction {
	if o == nil || IsNil(o.MpsAction) {
		var ret MpsAction
		return ret
	}
	return *o.MpsAction
}

// GetMpsActionOk returns a tuple with the MpsAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMpsActionOk() (*MpsAction, bool) {
	if o == nil || IsNil(o.MpsAction) {
		return nil, false
	}
	return o.MpsAction, true
}

// HasMpsAction returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMpsAction() bool {
	if o != nil && !IsNil(o.MpsAction) {
		return true
	}

	return false
}

// SetMpsAction gets a reference to the given MpsAction and assigns it to the MpsAction field.
func (o *AppSessionContextReqData) SetMpsAction(v MpsAction) {
	o.MpsAction = &v
}

// GetMpsId returns the MpsId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMpsId() string {
	if o == nil || IsNil(o.MpsId) {
		var ret string
		return ret
	}
	return *o.MpsId
}

// GetMpsIdOk returns a tuple with the MpsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMpsIdOk() (*string, bool) {
	if o == nil || IsNil(o.MpsId) {
		return nil, false
	}
	return o.MpsId, true
}

// HasMpsId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMpsId() bool {
	if o != nil && !IsNil(o.MpsId) {
		return true
	}

	return false
}

// SetMpsId gets a reference to the given string and assigns it to the MpsId field.
func (o *AppSessionContextReqData) SetMpsId(v string) {
	o.MpsId = &v
}

// GetMcsId returns the McsId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetMcsId() string {
	if o == nil || IsNil(o.McsId) {
		var ret string
		return ret
	}
	return *o.McsId
}

// GetMcsIdOk returns a tuple with the McsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetMcsIdOk() (*string, bool) {
	if o == nil || IsNil(o.McsId) {
		return nil, false
	}
	return o.McsId, true
}

// HasMcsId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasMcsId() bool {
	if o != nil && !IsNil(o.McsId) {
		return true
	}

	return false
}

// SetMcsId gets a reference to the given string and assigns it to the McsId field.
func (o *AppSessionContextReqData) SetMcsId(v string) {
	o.McsId = &v
}

// GetPreemptControlInfo returns the PreemptControlInfo field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetPreemptControlInfo() PreemptionControlInformation {
	if o == nil || IsNil(o.PreemptControlInfo) {
		var ret PreemptionControlInformation
		return ret
	}
	return *o.PreemptControlInfo
}

// GetPreemptControlInfoOk returns a tuple with the PreemptControlInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetPreemptControlInfoOk() (*PreemptionControlInformation, bool) {
	if o == nil || IsNil(o.PreemptControlInfo) {
		return nil, false
	}
	return o.PreemptControlInfo, true
}

// HasPreemptControlInfo returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasPreemptControlInfo() bool {
	if o != nil && !IsNil(o.PreemptControlInfo) {
		return true
	}

	return false
}

// SetPreemptControlInfo gets a reference to the given PreemptionControlInformation and assigns it to the PreemptControlInfo field.
func (o *AppSessionContextReqData) SetPreemptControlInfo(v PreemptionControlInformation) {
	o.PreemptControlInfo = &v
}

// GetQosDuration returns the QosDuration field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetQosDuration() int32 {
	if o == nil || IsNil(o.QosDuration) {
		var ret int32
		return ret
	}
	return *o.QosDuration
}

// GetQosDurationOk returns a tuple with the QosDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetQosDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.QosDuration) {
		return nil, false
	}
	return o.QosDuration, true
}

// HasQosDuration returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasQosDuration() bool {
	if o != nil && !IsNil(o.QosDuration) {
		return true
	}

	return false
}

// SetQosDuration gets a reference to the given int32 and assigns it to the QosDuration field.
func (o *AppSessionContextReqData) SetQosDuration(v int32) {
	o.QosDuration = &v
}

// GetQosInactInt returns the QosInactInt field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetQosInactInt() int32 {
	if o == nil || IsNil(o.QosInactInt) {
		var ret int32
		return ret
	}
	return *o.QosInactInt
}

// GetQosInactIntOk returns a tuple with the QosInactInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetQosInactIntOk() (*int32, bool) {
	if o == nil || IsNil(o.QosInactInt) {
		return nil, false
	}
	return o.QosInactInt, true
}

// HasQosInactInt returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasQosInactInt() bool {
	if o != nil && !IsNil(o.QosInactInt) {
		return true
	}

	return false
}

// SetQosInactInt gets a reference to the given int32 and assigns it to the QosInactInt field.
func (o *AppSessionContextReqData) SetQosInactInt(v int32) {
	o.QosInactInt = &v
}

// GetResPrio returns the ResPrio field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetResPrio() ReservPriority {
	if o == nil || IsNil(o.ResPrio) {
		var ret ReservPriority
		return ret
	}
	return *o.ResPrio
}

// GetResPrioOk returns a tuple with the ResPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetResPrioOk() (*ReservPriority, bool) {
	if o == nil || IsNil(o.ResPrio) {
		return nil, false
	}
	return o.ResPrio, true
}

// HasResPrio returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasResPrio() bool {
	if o != nil && !IsNil(o.ResPrio) {
		return true
	}

	return false
}

// SetResPrio gets a reference to the given ReservPriority and assigns it to the ResPrio field.
func (o *AppSessionContextReqData) SetResPrio(v ReservPriority) {
	o.ResPrio = &v
}

// GetServInfStatus returns the ServInfStatus field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetServInfStatus() ServiceInfoStatus {
	if o == nil || IsNil(o.ServInfStatus) {
		var ret ServiceInfoStatus
		return ret
	}
	return *o.ServInfStatus
}

// GetServInfStatusOk returns a tuple with the ServInfStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetServInfStatusOk() (*ServiceInfoStatus, bool) {
	if o == nil || IsNil(o.ServInfStatus) {
		return nil, false
	}
	return o.ServInfStatus, true
}

// HasServInfStatus returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasServInfStatus() bool {
	if o != nil && !IsNil(o.ServInfStatus) {
		return true
	}

	return false
}

// SetServInfStatus gets a reference to the given ServiceInfoStatus and assigns it to the ServInfStatus field.
func (o *AppSessionContextReqData) SetServInfStatus(v ServiceInfoStatus) {
	o.ServInfStatus = &v
}

// GetNotifUri returns the NotifUri field value
func (o *AppSessionContextReqData) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *AppSessionContextReqData) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetServUrn returns the ServUrn field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetServUrn() string {
	if o == nil || IsNil(o.ServUrn) {
		var ret string
		return ret
	}
	return *o.ServUrn
}

// GetServUrnOk returns a tuple with the ServUrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetServUrnOk() (*string, bool) {
	if o == nil || IsNil(o.ServUrn) {
		return nil, false
	}
	return o.ServUrn, true
}

// HasServUrn returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasServUrn() bool {
	if o != nil && !IsNil(o.ServUrn) {
		return true
	}

	return false
}

// SetServUrn gets a reference to the given string and assigns it to the ServUrn field.
func (o *AppSessionContextReqData) SetServUrn(v string) {
	o.ServUrn = &v
}

// GetSliceInfo returns the SliceInfo field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetSliceInfo() Snssai {
	if o == nil || IsNil(o.SliceInfo) {
		var ret Snssai
		return ret
	}
	return *o.SliceInfo
}

// GetSliceInfoOk returns a tuple with the SliceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetSliceInfoOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SliceInfo) {
		return nil, false
	}
	return o.SliceInfo, true
}

// HasSliceInfo returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasSliceInfo() bool {
	if o != nil && !IsNil(o.SliceInfo) {
		return true
	}

	return false
}

// SetSliceInfo gets a reference to the given Snssai and assigns it to the SliceInfo field.
func (o *AppSessionContextReqData) SetSliceInfo(v Snssai) {
	o.SliceInfo = &v
}

// GetSponId returns the SponId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetSponId() string {
	if o == nil || IsNil(o.SponId) {
		var ret string
		return ret
	}
	return *o.SponId
}

// GetSponIdOk returns a tuple with the SponId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetSponIdOk() (*string, bool) {
	if o == nil || IsNil(o.SponId) {
		return nil, false
	}
	return o.SponId, true
}

// HasSponId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasSponId() bool {
	if o != nil && !IsNil(o.SponId) {
		return true
	}

	return false
}

// SetSponId gets a reference to the given string and assigns it to the SponId field.
func (o *AppSessionContextReqData) SetSponId(v string) {
	o.SponId = &v
}

// GetSponStatus returns the SponStatus field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetSponStatus() SponsoringStatus {
	if o == nil || IsNil(o.SponStatus) {
		var ret SponsoringStatus
		return ret
	}
	return *o.SponStatus
}

// GetSponStatusOk returns a tuple with the SponStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetSponStatusOk() (*SponsoringStatus, bool) {
	if o == nil || IsNil(o.SponStatus) {
		return nil, false
	}
	return o.SponStatus, true
}

// HasSponStatus returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasSponStatus() bool {
	if o != nil && !IsNil(o.SponStatus) {
		return true
	}

	return false
}

// SetSponStatus gets a reference to the given SponsoringStatus and assigns it to the SponStatus field.
func (o *AppSessionContextReqData) SetSponStatus(v SponsoringStatus) {
	o.SponStatus = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AppSessionContextReqData) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *AppSessionContextReqData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSuppFeat returns the SuppFeat field value
func (o *AppSessionContextReqData) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *AppSessionContextReqData) SetSuppFeat(v string) {
	o.SuppFeat = v
}

// GetUeIpv4 returns the UeIpv4 field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetUeIpv4() string {
	if o == nil || IsNil(o.UeIpv4) {
		var ret string
		return ret
	}
	return *o.UeIpv4
}

// GetUeIpv4Ok returns a tuple with the UeIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetUeIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.UeIpv4) {
		return nil, false
	}
	return o.UeIpv4, true
}

// HasUeIpv4 returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasUeIpv4() bool {
	if o != nil && !IsNil(o.UeIpv4) {
		return true
	}

	return false
}

// SetUeIpv4 gets a reference to the given string and assigns it to the UeIpv4 field.
func (o *AppSessionContextReqData) SetUeIpv4(v string) {
	o.UeIpv4 = &v
}

// GetUeIpv6 returns the UeIpv6 field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetUeIpv6() Ipv6Addr {
	if o == nil || IsNil(o.UeIpv6) {
		var ret Ipv6Addr
		return ret
	}
	return *o.UeIpv6
}

// GetUeIpv6Ok returns a tuple with the UeIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetUeIpv6Ok() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.UeIpv6) {
		return nil, false
	}
	return o.UeIpv6, true
}

// HasUeIpv6 returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasUeIpv6() bool {
	if o != nil && !IsNil(o.UeIpv6) {
		return true
	}

	return false
}

// SetUeIpv6 gets a reference to the given Ipv6Addr and assigns it to the UeIpv6 field.
func (o *AppSessionContextReqData) SetUeIpv6(v Ipv6Addr) {
	o.UeIpv6 = &v
}

// GetUeMac returns the UeMac field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetUeMac() string {
	if o == nil || IsNil(o.UeMac) {
		var ret string
		return ret
	}
	return *o.UeMac
}

// GetUeMacOk returns a tuple with the UeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetUeMacOk() (*string, bool) {
	if o == nil || IsNil(o.UeMac) {
		return nil, false
	}
	return o.UeMac, true
}

// HasUeMac returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasUeMac() bool {
	if o != nil && !IsNil(o.UeMac) {
		return true
	}

	return false
}

// SetUeMac gets a reference to the given string and assigns it to the UeMac field.
func (o *AppSessionContextReqData) SetUeMac(v string) {
	o.UeMac = &v
}

// GetTsnBridgeManCont returns the TsnBridgeManCont field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetTsnBridgeManCont() BridgeManagementContainer {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		var ret BridgeManagementContainer
		return ret
	}
	return *o.TsnBridgeManCont
}

// GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool) {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		return nil, false
	}
	return o.TsnBridgeManCont, true
}

// HasTsnBridgeManCont returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasTsnBridgeManCont() bool {
	if o != nil && !IsNil(o.TsnBridgeManCont) {
		return true
	}

	return false
}

// SetTsnBridgeManCont gets a reference to the given BridgeManagementContainer and assigns it to the TsnBridgeManCont field.
func (o *AppSessionContextReqData) SetTsnBridgeManCont(v BridgeManagementContainer) {
	o.TsnBridgeManCont = &v
}

// GetTsnPortManContDstt returns the TsnPortManContDstt field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetTsnPortManContDstt() PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		var ret PortManagementContainer
		return ret
	}
	return *o.TsnPortManContDstt
}

// GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetTsnPortManContDsttOk() (*PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		return nil, false
	}
	return o.TsnPortManContDstt, true
}

// HasTsnPortManContDstt returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasTsnPortManContDstt() bool {
	if o != nil && !IsNil(o.TsnPortManContDstt) {
		return true
	}

	return false
}

// SetTsnPortManContDstt gets a reference to the given PortManagementContainer and assigns it to the TsnPortManContDstt field.
func (o *AppSessionContextReqData) SetTsnPortManContDstt(v PortManagementContainer) {
	o.TsnPortManContDstt = &v
}

// GetTsnPortManContNwtts returns the TsnPortManContNwtts field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetTsnPortManContNwtts() []PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		var ret []PortManagementContainer
		return ret
	}
	return o.TsnPortManContNwtts
}

// GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetTsnPortManContNwttsOk() ([]PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		return nil, false
	}
	return o.TsnPortManContNwtts, true
}

// HasTsnPortManContNwtts returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasTsnPortManContNwtts() bool {
	if o != nil && !IsNil(o.TsnPortManContNwtts) {
		return true
	}

	return false
}

// SetTsnPortManContNwtts gets a reference to the given []PortManagementContainer and assigns it to the TsnPortManContNwtts field.
func (o *AppSessionContextReqData) SetTsnPortManContNwtts(v []PortManagementContainer) {
	o.TsnPortManContNwtts = v
}

// GetTscNotifUri returns the TscNotifUri field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetTscNotifUri() string {
	if o == nil || IsNil(o.TscNotifUri) {
		var ret string
		return ret
	}
	return *o.TscNotifUri
}

// GetTscNotifUriOk returns a tuple with the TscNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetTscNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.TscNotifUri) {
		return nil, false
	}
	return o.TscNotifUri, true
}

// HasTscNotifUri returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasTscNotifUri() bool {
	if o != nil && !IsNil(o.TscNotifUri) {
		return true
	}

	return false
}

// SetTscNotifUri gets a reference to the given string and assigns it to the TscNotifUri field.
func (o *AppSessionContextReqData) SetTscNotifUri(v string) {
	o.TscNotifUri = &v
}

// GetTscNotifCorreId returns the TscNotifCorreId field value if set, zero value otherwise.
func (o *AppSessionContextReqData) GetTscNotifCorreId() string {
	if o == nil || IsNil(o.TscNotifCorreId) {
		var ret string
		return ret
	}
	return *o.TscNotifCorreId
}

// GetTscNotifCorreIdOk returns a tuple with the TscNotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContextReqData) GetTscNotifCorreIdOk() (*string, bool) {
	if o == nil || IsNil(o.TscNotifCorreId) {
		return nil, false
	}
	return o.TscNotifCorreId, true
}

// HasTscNotifCorreId returns a boolean if a field has been set.
func (o *AppSessionContextReqData) HasTscNotifCorreId() bool {
	if o != nil && !IsNil(o.TscNotifCorreId) {
		return true
	}

	return false
}

// SetTscNotifCorreId gets a reference to the given string and assigns it to the TscNotifCorreId field.
func (o *AppSessionContextReqData) SetTscNotifCorreId(v string) {
	o.TscNotifCorreId = &v
}

func (o AppSessionContextReqData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSessionContextReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	if !IsNil(o.AfChargId) {
		toSerialize["afChargId"] = o.AfChargId
	}
	if !IsNil(o.AfReqData) {
		toSerialize["afReqData"] = o.AfReqData
	}
	if !IsNil(o.AfRoutReq) {
		toSerialize["afRoutReq"] = o.AfRoutReq
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
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.EvSubsc) {
		toSerialize["evSubsc"] = o.EvSubsc
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
	if !IsNil(o.MultiModalId) {
		toSerialize["multiModalId"] = o.MultiModalId
	}
	if !IsNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
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
	if !IsNil(o.PreemptControlInfo) {
		toSerialize["preemptControlInfo"] = o.PreemptControlInfo
	}
	if !IsNil(o.QosDuration) {
		toSerialize["qosDuration"] = o.QosDuration
	}
	if !IsNil(o.QosInactInt) {
		toSerialize["qosInactInt"] = o.QosInactInt
	}
	if !IsNil(o.ResPrio) {
		toSerialize["resPrio"] = o.ResPrio
	}
	if !IsNil(o.ServInfStatus) {
		toSerialize["servInfStatus"] = o.ServInfStatus
	}
	toSerialize["notifUri"] = o.NotifUri
	if !IsNil(o.ServUrn) {
		toSerialize["servUrn"] = o.ServUrn
	}
	if !IsNil(o.SliceInfo) {
		toSerialize["sliceInfo"] = o.SliceInfo
	}
	if !IsNil(o.SponId) {
		toSerialize["sponId"] = o.SponId
	}
	if !IsNil(o.SponStatus) {
		toSerialize["sponStatus"] = o.SponStatus
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	toSerialize["suppFeat"] = o.SuppFeat
	if !IsNil(o.UeIpv4) {
		toSerialize["ueIpv4"] = o.UeIpv4
	}
	if !IsNil(o.UeIpv6) {
		toSerialize["ueIpv6"] = o.UeIpv6
	}
	if !IsNil(o.UeMac) {
		toSerialize["ueMac"] = o.UeMac
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

func (o *AppSessionContextReqData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifUri",
		"suppFeat",
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

	varAppSessionContextReqData := _AppSessionContextReqData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppSessionContextReqData)

	if err != nil {
		return err
	}

	*o = AppSessionContextReqData(varAppSessionContextReqData)

	return err
}

type NullableAppSessionContextReqData struct {
	value *AppSessionContextReqData
	isSet bool
}

func (v NullableAppSessionContextReqData) Get() *AppSessionContextReqData {
	return v.value
}

func (v *NullableAppSessionContextReqData) Set(val *AppSessionContextReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContextReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContextReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContextReqData(val *AppSessionContextReqData) *NullableAppSessionContextReqData {
	return &NullableAppSessionContextReqData{value: val, isSet: true}
}

func (v NullableAppSessionContextReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContextReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
