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
	"time"
)

// checks if the EventsNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsNotification{}

// EventsNotification Describes the notification of a matched event.
type EventsNotification struct {
	// Includes the detected application report.
	AdReports     []AppDetectionReport          `json:"adReports,omitempty"`
	AccessType    *AccessType                   `json:"accessType,omitempty"`
	AddAccessInfo *AdditionalAccessInfo         `json:"addAccessInfo,omitempty"`
	RelAccessInfo *AdditionalAccessInfo         `json:"relAccessInfo,omitempty"`
	AnChargAddr   NullableAccNetChargingAddress `json:"anChargAddr,omitempty"`
	AnChargIds    []AccessNetChargingIdentifier `json:"anChargIds,omitempty"`
	AnGwAddr      NullableAnGwAddress           `json:"anGwAddr,omitempty"`
	L4sReports    []L4sSupport                  `json:"l4sReports,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	EvSubsUri                 string                       `json:"evSubsUri"`
	EvNotifs                  []AfEventNotification        `json:"evNotifs"`
	FailedResourcAllocReports []ResourcesAllocationInfo    `json:"failedResourcAllocReports,omitempty"`
	SuccResourcAllocReports   []ResourcesAllocationInfo    `json:"succResourcAllocReports,omitempty"`
	NoNetLocSupp              *NetLocAccessSupport         `json:"noNetLocSupp,omitempty"`
	OutOfCredReports          []OutOfCreditInformation     `json:"outOfCredReports,omitempty"`
	PlmnId                    *PlmnIdNid                   `json:"plmnId,omitempty"`
	QncReports                []QosNotificationControlInfo `json:"qncReports,omitempty"`
	QosMonReports             []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	QosMonDatRateReps         []QosMonitoringReport        `json:"qosMonDatRateReps,omitempty"`
	PdvMonReports             []PdvMonitoringReport        `json:"pdvMonReports,omitempty"`
	CongestReports            []QosMonitoringReport        `json:"congestReports,omitempty"`
	RttMonReports             []QosMonitoringReport        `json:"rttMonReports,omitempty"`
	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses     []RanNasRelCause           `json:"ranNasRelCauses,omitempty"`
	RatType             *RatType                   `json:"ratType,omitempty"`
	SatBackhaulCategory *SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	UeLoc               *UserLocation              `json:"ueLoc,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UeLocTime *time.Time `json:"ueLocTime,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone *string           `json:"ueTimeZone,omitempty"`
	UsgRep     *AccumulatedUsage `json:"usgRep,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UrspEnfRep *string  `json:"urspEnfRep,omitempty"`
	SscMode    *SscMode `json:"sscMode,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	UeReqDnn                *string                         `json:"ueReqDnn,omitempty"`
	RedundantPduSessionInfo *RedundantPduSessionInformation `json:"redundantPduSessionInfo,omitempty"`
	TsnBridgeManCont        *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt      *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts     []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	Ipv4AddrList            []string                        `json:"ipv4AddrList,omitempty"`
	Ipv6PrefixList          []Ipv6Prefix                    `json:"ipv6PrefixList,omitempty"`
	BatOffsetInfo           *BatOffsetInfo                  `json:"batOffsetInfo,omitempty"`
}

type _EventsNotification EventsNotification

// NewEventsNotification instantiates a new EventsNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsNotification(evSubsUri string, evNotifs []AfEventNotification) *EventsNotification {
	this := EventsNotification{}
	this.EvSubsUri = evSubsUri
	this.EvNotifs = evNotifs
	return &this
}

// NewEventsNotificationWithDefaults instantiates a new EventsNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsNotificationWithDefaults() *EventsNotification {
	this := EventsNotification{}
	return &this
}

// GetAdReports returns the AdReports field value if set, zero value otherwise.
func (o *EventsNotification) GetAdReports() []AppDetectionReport {
	if o == nil || IsNil(o.AdReports) {
		var ret []AppDetectionReport
		return ret
	}
	return o.AdReports
}

// GetAdReportsOk returns a tuple with the AdReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetAdReportsOk() ([]AppDetectionReport, bool) {
	if o == nil || IsNil(o.AdReports) {
		return nil, false
	}
	return o.AdReports, true
}

// HasAdReports returns a boolean if a field has been set.
func (o *EventsNotification) HasAdReports() bool {
	if o != nil && !IsNil(o.AdReports) {
		return true
	}

	return false
}

// SetAdReports gets a reference to the given []AppDetectionReport and assigns it to the AdReports field.
func (o *EventsNotification) SetAdReports(v []AppDetectionReport) {
	o.AdReports = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *EventsNotification) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *EventsNotification) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *EventsNotification) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetAddAccessInfo returns the AddAccessInfo field value if set, zero value otherwise.
func (o *EventsNotification) GetAddAccessInfo() AdditionalAccessInfo {
	if o == nil || IsNil(o.AddAccessInfo) {
		var ret AdditionalAccessInfo
		return ret
	}
	return *o.AddAccessInfo
}

// GetAddAccessInfoOk returns a tuple with the AddAccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool) {
	if o == nil || IsNil(o.AddAccessInfo) {
		return nil, false
	}
	return o.AddAccessInfo, true
}

// HasAddAccessInfo returns a boolean if a field has been set.
func (o *EventsNotification) HasAddAccessInfo() bool {
	if o != nil && !IsNil(o.AddAccessInfo) {
		return true
	}

	return false
}

// SetAddAccessInfo gets a reference to the given AdditionalAccessInfo and assigns it to the AddAccessInfo field.
func (o *EventsNotification) SetAddAccessInfo(v AdditionalAccessInfo) {
	o.AddAccessInfo = &v
}

// GetRelAccessInfo returns the RelAccessInfo field value if set, zero value otherwise.
func (o *EventsNotification) GetRelAccessInfo() AdditionalAccessInfo {
	if o == nil || IsNil(o.RelAccessInfo) {
		var ret AdditionalAccessInfo
		return ret
	}
	return *o.RelAccessInfo
}

// GetRelAccessInfoOk returns a tuple with the RelAccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetRelAccessInfoOk() (*AdditionalAccessInfo, bool) {
	if o == nil || IsNil(o.RelAccessInfo) {
		return nil, false
	}
	return o.RelAccessInfo, true
}

// HasRelAccessInfo returns a boolean if a field has been set.
func (o *EventsNotification) HasRelAccessInfo() bool {
	if o != nil && !IsNil(o.RelAccessInfo) {
		return true
	}

	return false
}

// SetRelAccessInfo gets a reference to the given AdditionalAccessInfo and assigns it to the RelAccessInfo field.
func (o *EventsNotification) SetRelAccessInfo(v AdditionalAccessInfo) {
	o.RelAccessInfo = &v
}

// GetAnChargAddr returns the AnChargAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsNotification) GetAnChargAddr() AccNetChargingAddress {
	if o == nil || IsNil(o.AnChargAddr.Get()) {
		var ret AccNetChargingAddress
		return ret
	}
	return *o.AnChargAddr.Get()
}

// GetAnChargAddrOk returns a tuple with the AnChargAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsNotification) GetAnChargAddrOk() (*AccNetChargingAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnChargAddr.Get(), o.AnChargAddr.IsSet()
}

// HasAnChargAddr returns a boolean if a field has been set.
func (o *EventsNotification) HasAnChargAddr() bool {
	if o != nil && o.AnChargAddr.IsSet() {
		return true
	}

	return false
}

// SetAnChargAddr gets a reference to the given NullableAccNetChargingAddress and assigns it to the AnChargAddr field.
func (o *EventsNotification) SetAnChargAddr(v AccNetChargingAddress) {
	o.AnChargAddr.Set(&v)
}

// SetAnChargAddrNil sets the value for AnChargAddr to be an explicit nil
func (o *EventsNotification) SetAnChargAddrNil() {
	o.AnChargAddr.Set(nil)
}

// UnsetAnChargAddr ensures that no value is present for AnChargAddr, not even an explicit nil
func (o *EventsNotification) UnsetAnChargAddr() {
	o.AnChargAddr.Unset()
}

// GetAnChargIds returns the AnChargIds field value if set, zero value otherwise.
func (o *EventsNotification) GetAnChargIds() []AccessNetChargingIdentifier {
	if o == nil || IsNil(o.AnChargIds) {
		var ret []AccessNetChargingIdentifier
		return ret
	}
	return o.AnChargIds
}

// GetAnChargIdsOk returns a tuple with the AnChargIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetAnChargIdsOk() ([]AccessNetChargingIdentifier, bool) {
	if o == nil || IsNil(o.AnChargIds) {
		return nil, false
	}
	return o.AnChargIds, true
}

// HasAnChargIds returns a boolean if a field has been set.
func (o *EventsNotification) HasAnChargIds() bool {
	if o != nil && !IsNil(o.AnChargIds) {
		return true
	}

	return false
}

// SetAnChargIds gets a reference to the given []AccessNetChargingIdentifier and assigns it to the AnChargIds field.
func (o *EventsNotification) SetAnChargIds(v []AccessNetChargingIdentifier) {
	o.AnChargIds = v
}

// GetAnGwAddr returns the AnGwAddr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsNotification) GetAnGwAddr() AnGwAddress {
	if o == nil || IsNil(o.AnGwAddr.Get()) {
		var ret AnGwAddress
		return ret
	}
	return *o.AnGwAddr.Get()
}

// GetAnGwAddrOk returns a tuple with the AnGwAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsNotification) GetAnGwAddrOk() (*AnGwAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnGwAddr.Get(), o.AnGwAddr.IsSet()
}

// HasAnGwAddr returns a boolean if a field has been set.
func (o *EventsNotification) HasAnGwAddr() bool {
	if o != nil && o.AnGwAddr.IsSet() {
		return true
	}

	return false
}

// SetAnGwAddr gets a reference to the given NullableAnGwAddress and assigns it to the AnGwAddr field.
func (o *EventsNotification) SetAnGwAddr(v AnGwAddress) {
	o.AnGwAddr.Set(&v)
}

// SetAnGwAddrNil sets the value for AnGwAddr to be an explicit nil
func (o *EventsNotification) SetAnGwAddrNil() {
	o.AnGwAddr.Set(nil)
}

// UnsetAnGwAddr ensures that no value is present for AnGwAddr, not even an explicit nil
func (o *EventsNotification) UnsetAnGwAddr() {
	o.AnGwAddr.Unset()
}

// GetL4sReports returns the L4sReports field value if set, zero value otherwise.
func (o *EventsNotification) GetL4sReports() []L4sSupport {
	if o == nil || IsNil(o.L4sReports) {
		var ret []L4sSupport
		return ret
	}
	return o.L4sReports
}

// GetL4sReportsOk returns a tuple with the L4sReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetL4sReportsOk() ([]L4sSupport, bool) {
	if o == nil || IsNil(o.L4sReports) {
		return nil, false
	}
	return o.L4sReports, true
}

// HasL4sReports returns a boolean if a field has been set.
func (o *EventsNotification) HasL4sReports() bool {
	if o != nil && !IsNil(o.L4sReports) {
		return true
	}

	return false
}

// SetL4sReports gets a reference to the given []L4sSupport and assigns it to the L4sReports field.
func (o *EventsNotification) SetL4sReports(v []L4sSupport) {
	o.L4sReports = v
}

// GetEvSubsUri returns the EvSubsUri field value
func (o *EventsNotification) GetEvSubsUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvSubsUri
}

// GetEvSubsUriOk returns a tuple with the EvSubsUri field value
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetEvSubsUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvSubsUri, true
}

// SetEvSubsUri sets field value
func (o *EventsNotification) SetEvSubsUri(v string) {
	o.EvSubsUri = v
}

// GetEvNotifs returns the EvNotifs field value
func (o *EventsNotification) GetEvNotifs() []AfEventNotification {
	if o == nil {
		var ret []AfEventNotification
		return ret
	}

	return o.EvNotifs
}

// GetEvNotifsOk returns a tuple with the EvNotifs field value
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetEvNotifsOk() ([]AfEventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvNotifs, true
}

// SetEvNotifs sets field value
func (o *EventsNotification) SetEvNotifs(v []AfEventNotification) {
	o.EvNotifs = v
}

// GetFailedResourcAllocReports returns the FailedResourcAllocReports field value if set, zero value otherwise.
func (o *EventsNotification) GetFailedResourcAllocReports() []ResourcesAllocationInfo {
	if o == nil || IsNil(o.FailedResourcAllocReports) {
		var ret []ResourcesAllocationInfo
		return ret
	}
	return o.FailedResourcAllocReports
}

// GetFailedResourcAllocReportsOk returns a tuple with the FailedResourcAllocReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetFailedResourcAllocReportsOk() ([]ResourcesAllocationInfo, bool) {
	if o == nil || IsNil(o.FailedResourcAllocReports) {
		return nil, false
	}
	return o.FailedResourcAllocReports, true
}

// HasFailedResourcAllocReports returns a boolean if a field has been set.
func (o *EventsNotification) HasFailedResourcAllocReports() bool {
	if o != nil && !IsNil(o.FailedResourcAllocReports) {
		return true
	}

	return false
}

// SetFailedResourcAllocReports gets a reference to the given []ResourcesAllocationInfo and assigns it to the FailedResourcAllocReports field.
func (o *EventsNotification) SetFailedResourcAllocReports(v []ResourcesAllocationInfo) {
	o.FailedResourcAllocReports = v
}

// GetSuccResourcAllocReports returns the SuccResourcAllocReports field value if set, zero value otherwise.
func (o *EventsNotification) GetSuccResourcAllocReports() []ResourcesAllocationInfo {
	if o == nil || IsNil(o.SuccResourcAllocReports) {
		var ret []ResourcesAllocationInfo
		return ret
	}
	return o.SuccResourcAllocReports
}

// GetSuccResourcAllocReportsOk returns a tuple with the SuccResourcAllocReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetSuccResourcAllocReportsOk() ([]ResourcesAllocationInfo, bool) {
	if o == nil || IsNil(o.SuccResourcAllocReports) {
		return nil, false
	}
	return o.SuccResourcAllocReports, true
}

// HasSuccResourcAllocReports returns a boolean if a field has been set.
func (o *EventsNotification) HasSuccResourcAllocReports() bool {
	if o != nil && !IsNil(o.SuccResourcAllocReports) {
		return true
	}

	return false
}

// SetSuccResourcAllocReports gets a reference to the given []ResourcesAllocationInfo and assigns it to the SuccResourcAllocReports field.
func (o *EventsNotification) SetSuccResourcAllocReports(v []ResourcesAllocationInfo) {
	o.SuccResourcAllocReports = v
}

// GetNoNetLocSupp returns the NoNetLocSupp field value if set, zero value otherwise.
func (o *EventsNotification) GetNoNetLocSupp() NetLocAccessSupport {
	if o == nil || IsNil(o.NoNetLocSupp) {
		var ret NetLocAccessSupport
		return ret
	}
	return *o.NoNetLocSupp
}

// GetNoNetLocSuppOk returns a tuple with the NoNetLocSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetNoNetLocSuppOk() (*NetLocAccessSupport, bool) {
	if o == nil || IsNil(o.NoNetLocSupp) {
		return nil, false
	}
	return o.NoNetLocSupp, true
}

// HasNoNetLocSupp returns a boolean if a field has been set.
func (o *EventsNotification) HasNoNetLocSupp() bool {
	if o != nil && !IsNil(o.NoNetLocSupp) {
		return true
	}

	return false
}

// SetNoNetLocSupp gets a reference to the given NetLocAccessSupport and assigns it to the NoNetLocSupp field.
func (o *EventsNotification) SetNoNetLocSupp(v NetLocAccessSupport) {
	o.NoNetLocSupp = &v
}

// GetOutOfCredReports returns the OutOfCredReports field value if set, zero value otherwise.
func (o *EventsNotification) GetOutOfCredReports() []OutOfCreditInformation {
	if o == nil || IsNil(o.OutOfCredReports) {
		var ret []OutOfCreditInformation
		return ret
	}
	return o.OutOfCredReports
}

// GetOutOfCredReportsOk returns a tuple with the OutOfCredReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetOutOfCredReportsOk() ([]OutOfCreditInformation, bool) {
	if o == nil || IsNil(o.OutOfCredReports) {
		return nil, false
	}
	return o.OutOfCredReports, true
}

// HasOutOfCredReports returns a boolean if a field has been set.
func (o *EventsNotification) HasOutOfCredReports() bool {
	if o != nil && !IsNil(o.OutOfCredReports) {
		return true
	}

	return false
}

// SetOutOfCredReports gets a reference to the given []OutOfCreditInformation and assigns it to the OutOfCredReports field.
func (o *EventsNotification) SetOutOfCredReports(v []OutOfCreditInformation) {
	o.OutOfCredReports = v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *EventsNotification) GetPlmnId() PlmnIdNid {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnIdNid
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetPlmnIdOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *EventsNotification) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnIdNid and assigns it to the PlmnId field.
func (o *EventsNotification) SetPlmnId(v PlmnIdNid) {
	o.PlmnId = &v
}

// GetQncReports returns the QncReports field value if set, zero value otherwise.
func (o *EventsNotification) GetQncReports() []QosNotificationControlInfo {
	if o == nil || IsNil(o.QncReports) {
		var ret []QosNotificationControlInfo
		return ret
	}
	return o.QncReports
}

// GetQncReportsOk returns a tuple with the QncReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetQncReportsOk() ([]QosNotificationControlInfo, bool) {
	if o == nil || IsNil(o.QncReports) {
		return nil, false
	}
	return o.QncReports, true
}

// HasQncReports returns a boolean if a field has been set.
func (o *EventsNotification) HasQncReports() bool {
	if o != nil && !IsNil(o.QncReports) {
		return true
	}

	return false
}

// SetQncReports gets a reference to the given []QosNotificationControlInfo and assigns it to the QncReports field.
func (o *EventsNotification) SetQncReports(v []QosNotificationControlInfo) {
	o.QncReports = v
}

// GetQosMonReports returns the QosMonReports field value if set, zero value otherwise.
func (o *EventsNotification) GetQosMonReports() []QosMonitoringReport {
	if o == nil || IsNil(o.QosMonReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.QosMonReports
}

// GetQosMonReportsOk returns a tuple with the QosMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetQosMonReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || IsNil(o.QosMonReports) {
		return nil, false
	}
	return o.QosMonReports, true
}

// HasQosMonReports returns a boolean if a field has been set.
func (o *EventsNotification) HasQosMonReports() bool {
	if o != nil && !IsNil(o.QosMonReports) {
		return true
	}

	return false
}

// SetQosMonReports gets a reference to the given []QosMonitoringReport and assigns it to the QosMonReports field.
func (o *EventsNotification) SetQosMonReports(v []QosMonitoringReport) {
	o.QosMonReports = v
}

// GetQosMonDatRateReps returns the QosMonDatRateReps field value if set, zero value otherwise.
func (o *EventsNotification) GetQosMonDatRateReps() []QosMonitoringReport {
	if o == nil || IsNil(o.QosMonDatRateReps) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.QosMonDatRateReps
}

// GetQosMonDatRateRepsOk returns a tuple with the QosMonDatRateReps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetQosMonDatRateRepsOk() ([]QosMonitoringReport, bool) {
	if o == nil || IsNil(o.QosMonDatRateReps) {
		return nil, false
	}
	return o.QosMonDatRateReps, true
}

// HasQosMonDatRateReps returns a boolean if a field has been set.
func (o *EventsNotification) HasQosMonDatRateReps() bool {
	if o != nil && !IsNil(o.QosMonDatRateReps) {
		return true
	}

	return false
}

// SetQosMonDatRateReps gets a reference to the given []QosMonitoringReport and assigns it to the QosMonDatRateReps field.
func (o *EventsNotification) SetQosMonDatRateReps(v []QosMonitoringReport) {
	o.QosMonDatRateReps = v
}

// GetPdvMonReports returns the PdvMonReports field value if set, zero value otherwise.
func (o *EventsNotification) GetPdvMonReports() []PdvMonitoringReport {
	if o == nil || IsNil(o.PdvMonReports) {
		var ret []PdvMonitoringReport
		return ret
	}
	return o.PdvMonReports
}

// GetPdvMonReportsOk returns a tuple with the PdvMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetPdvMonReportsOk() ([]PdvMonitoringReport, bool) {
	if o == nil || IsNil(o.PdvMonReports) {
		return nil, false
	}
	return o.PdvMonReports, true
}

// HasPdvMonReports returns a boolean if a field has been set.
func (o *EventsNotification) HasPdvMonReports() bool {
	if o != nil && !IsNil(o.PdvMonReports) {
		return true
	}

	return false
}

// SetPdvMonReports gets a reference to the given []PdvMonitoringReport and assigns it to the PdvMonReports field.
func (o *EventsNotification) SetPdvMonReports(v []PdvMonitoringReport) {
	o.PdvMonReports = v
}

// GetCongestReports returns the CongestReports field value if set, zero value otherwise.
func (o *EventsNotification) GetCongestReports() []QosMonitoringReport {
	if o == nil || IsNil(o.CongestReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.CongestReports
}

// GetCongestReportsOk returns a tuple with the CongestReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetCongestReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || IsNil(o.CongestReports) {
		return nil, false
	}
	return o.CongestReports, true
}

// HasCongestReports returns a boolean if a field has been set.
func (o *EventsNotification) HasCongestReports() bool {
	if o != nil && !IsNil(o.CongestReports) {
		return true
	}

	return false
}

// SetCongestReports gets a reference to the given []QosMonitoringReport and assigns it to the CongestReports field.
func (o *EventsNotification) SetCongestReports(v []QosMonitoringReport) {
	o.CongestReports = v
}

// GetRttMonReports returns the RttMonReports field value if set, zero value otherwise.
func (o *EventsNotification) GetRttMonReports() []QosMonitoringReport {
	if o == nil || IsNil(o.RttMonReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.RttMonReports
}

// GetRttMonReportsOk returns a tuple with the RttMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetRttMonReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || IsNil(o.RttMonReports) {
		return nil, false
	}
	return o.RttMonReports, true
}

// HasRttMonReports returns a boolean if a field has been set.
func (o *EventsNotification) HasRttMonReports() bool {
	if o != nil && !IsNil(o.RttMonReports) {
		return true
	}

	return false
}

// SetRttMonReports gets a reference to the given []QosMonitoringReport and assigns it to the RttMonReports field.
func (o *EventsNotification) SetRttMonReports(v []QosMonitoringReport) {
	o.RttMonReports = v
}

// GetRanNasRelCauses returns the RanNasRelCauses field value if set, zero value otherwise.
func (o *EventsNotification) GetRanNasRelCauses() []RanNasRelCause {
	if o == nil || IsNil(o.RanNasRelCauses) {
		var ret []RanNasRelCause
		return ret
	}
	return o.RanNasRelCauses
}

// GetRanNasRelCausesOk returns a tuple with the RanNasRelCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetRanNasRelCausesOk() ([]RanNasRelCause, bool) {
	if o == nil || IsNil(o.RanNasRelCauses) {
		return nil, false
	}
	return o.RanNasRelCauses, true
}

// HasRanNasRelCauses returns a boolean if a field has been set.
func (o *EventsNotification) HasRanNasRelCauses() bool {
	if o != nil && !IsNil(o.RanNasRelCauses) {
		return true
	}

	return false
}

// SetRanNasRelCauses gets a reference to the given []RanNasRelCause and assigns it to the RanNasRelCauses field.
func (o *EventsNotification) SetRanNasRelCauses(v []RanNasRelCause) {
	o.RanNasRelCauses = v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *EventsNotification) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *EventsNotification) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *EventsNotification) SetRatType(v RatType) {
	o.RatType = &v
}

// GetSatBackhaulCategory returns the SatBackhaulCategory field value if set, zero value otherwise.
func (o *EventsNotification) GetSatBackhaulCategory() SatelliteBackhaulCategory {
	if o == nil || IsNil(o.SatBackhaulCategory) {
		var ret SatelliteBackhaulCategory
		return ret
	}
	return *o.SatBackhaulCategory
}

// GetSatBackhaulCategoryOk returns a tuple with the SatBackhaulCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetSatBackhaulCategoryOk() (*SatelliteBackhaulCategory, bool) {
	if o == nil || IsNil(o.SatBackhaulCategory) {
		return nil, false
	}
	return o.SatBackhaulCategory, true
}

// HasSatBackhaulCategory returns a boolean if a field has been set.
func (o *EventsNotification) HasSatBackhaulCategory() bool {
	if o != nil && !IsNil(o.SatBackhaulCategory) {
		return true
	}

	return false
}

// SetSatBackhaulCategory gets a reference to the given SatelliteBackhaulCategory and assigns it to the SatBackhaulCategory field.
func (o *EventsNotification) SetSatBackhaulCategory(v SatelliteBackhaulCategory) {
	o.SatBackhaulCategory = &v
}

// GetUeLoc returns the UeLoc field value if set, zero value otherwise.
func (o *EventsNotification) GetUeLoc() UserLocation {
	if o == nil || IsNil(o.UeLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UeLoc
}

// GetUeLocOk returns a tuple with the UeLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUeLocOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UeLoc) {
		return nil, false
	}
	return o.UeLoc, true
}

// HasUeLoc returns a boolean if a field has been set.
func (o *EventsNotification) HasUeLoc() bool {
	if o != nil && !IsNil(o.UeLoc) {
		return true
	}

	return false
}

// SetUeLoc gets a reference to the given UserLocation and assigns it to the UeLoc field.
func (o *EventsNotification) SetUeLoc(v UserLocation) {
	o.UeLoc = &v
}

// GetUeLocTime returns the UeLocTime field value if set, zero value otherwise.
func (o *EventsNotification) GetUeLocTime() time.Time {
	if o == nil || IsNil(o.UeLocTime) {
		var ret time.Time
		return ret
	}
	return *o.UeLocTime
}

// GetUeLocTimeOk returns a tuple with the UeLocTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUeLocTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UeLocTime) {
		return nil, false
	}
	return o.UeLocTime, true
}

// HasUeLocTime returns a boolean if a field has been set.
func (o *EventsNotification) HasUeLocTime() bool {
	if o != nil && !IsNil(o.UeLocTime) {
		return true
	}

	return false
}

// SetUeLocTime gets a reference to the given time.Time and assigns it to the UeLocTime field.
func (o *EventsNotification) SetUeLocTime(v time.Time) {
	o.UeLocTime = &v
}

// GetUeTimeZone returns the UeTimeZone field value if set, zero value otherwise.
func (o *EventsNotification) GetUeTimeZone() string {
	if o == nil || IsNil(o.UeTimeZone) {
		var ret string
		return ret
	}
	return *o.UeTimeZone
}

// GetUeTimeZoneOk returns a tuple with the UeTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUeTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UeTimeZone) {
		return nil, false
	}
	return o.UeTimeZone, true
}

// HasUeTimeZone returns a boolean if a field has been set.
func (o *EventsNotification) HasUeTimeZone() bool {
	if o != nil && !IsNil(o.UeTimeZone) {
		return true
	}

	return false
}

// SetUeTimeZone gets a reference to the given string and assigns it to the UeTimeZone field.
func (o *EventsNotification) SetUeTimeZone(v string) {
	o.UeTimeZone = &v
}

// GetUsgRep returns the UsgRep field value if set, zero value otherwise.
func (o *EventsNotification) GetUsgRep() AccumulatedUsage {
	if o == nil || IsNil(o.UsgRep) {
		var ret AccumulatedUsage
		return ret
	}
	return *o.UsgRep
}

// GetUsgRepOk returns a tuple with the UsgRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUsgRepOk() (*AccumulatedUsage, bool) {
	if o == nil || IsNil(o.UsgRep) {
		return nil, false
	}
	return o.UsgRep, true
}

// HasUsgRep returns a boolean if a field has been set.
func (o *EventsNotification) HasUsgRep() bool {
	if o != nil && !IsNil(o.UsgRep) {
		return true
	}

	return false
}

// SetUsgRep gets a reference to the given AccumulatedUsage and assigns it to the UsgRep field.
func (o *EventsNotification) SetUsgRep(v AccumulatedUsage) {
	o.UsgRep = &v
}

// GetUrspEnfRep returns the UrspEnfRep field value if set, zero value otherwise.
func (o *EventsNotification) GetUrspEnfRep() string {
	if o == nil || IsNil(o.UrspEnfRep) {
		var ret string
		return ret
	}
	return *o.UrspEnfRep
}

// GetUrspEnfRepOk returns a tuple with the UrspEnfRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUrspEnfRepOk() (*string, bool) {
	if o == nil || IsNil(o.UrspEnfRep) {
		return nil, false
	}
	return o.UrspEnfRep, true
}

// HasUrspEnfRep returns a boolean if a field has been set.
func (o *EventsNotification) HasUrspEnfRep() bool {
	if o != nil && !IsNil(o.UrspEnfRep) {
		return true
	}

	return false
}

// SetUrspEnfRep gets a reference to the given string and assigns it to the UrspEnfRep field.
func (o *EventsNotification) SetUrspEnfRep(v string) {
	o.UrspEnfRep = &v
}

// GetSscMode returns the SscMode field value if set, zero value otherwise.
func (o *EventsNotification) GetSscMode() SscMode {
	if o == nil || IsNil(o.SscMode) {
		var ret SscMode
		return ret
	}
	return *o.SscMode
}

// GetSscModeOk returns a tuple with the SscMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetSscModeOk() (*SscMode, bool) {
	if o == nil || IsNil(o.SscMode) {
		return nil, false
	}
	return o.SscMode, true
}

// HasSscMode returns a boolean if a field has been set.
func (o *EventsNotification) HasSscMode() bool {
	if o != nil && !IsNil(o.SscMode) {
		return true
	}

	return false
}

// SetSscMode gets a reference to the given SscMode and assigns it to the SscMode field.
func (o *EventsNotification) SetSscMode(v SscMode) {
	o.SscMode = &v
}

// GetUeReqDnn returns the UeReqDnn field value if set, zero value otherwise.
func (o *EventsNotification) GetUeReqDnn() string {
	if o == nil || IsNil(o.UeReqDnn) {
		var ret string
		return ret
	}
	return *o.UeReqDnn
}

// GetUeReqDnnOk returns a tuple with the UeReqDnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetUeReqDnnOk() (*string, bool) {
	if o == nil || IsNil(o.UeReqDnn) {
		return nil, false
	}
	return o.UeReqDnn, true
}

// HasUeReqDnn returns a boolean if a field has been set.
func (o *EventsNotification) HasUeReqDnn() bool {
	if o != nil && !IsNil(o.UeReqDnn) {
		return true
	}

	return false
}

// SetUeReqDnn gets a reference to the given string and assigns it to the UeReqDnn field.
func (o *EventsNotification) SetUeReqDnn(v string) {
	o.UeReqDnn = &v
}

// GetRedundantPduSessionInfo returns the RedundantPduSessionInfo field value if set, zero value otherwise.
func (o *EventsNotification) GetRedundantPduSessionInfo() RedundantPduSessionInformation {
	if o == nil || IsNil(o.RedundantPduSessionInfo) {
		var ret RedundantPduSessionInformation
		return ret
	}
	return *o.RedundantPduSessionInfo
}

// GetRedundantPduSessionInfoOk returns a tuple with the RedundantPduSessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetRedundantPduSessionInfoOk() (*RedundantPduSessionInformation, bool) {
	if o == nil || IsNil(o.RedundantPduSessionInfo) {
		return nil, false
	}
	return o.RedundantPduSessionInfo, true
}

// HasRedundantPduSessionInfo returns a boolean if a field has been set.
func (o *EventsNotification) HasRedundantPduSessionInfo() bool {
	if o != nil && !IsNil(o.RedundantPduSessionInfo) {
		return true
	}

	return false
}

// SetRedundantPduSessionInfo gets a reference to the given RedundantPduSessionInformation and assigns it to the RedundantPduSessionInfo field.
func (o *EventsNotification) SetRedundantPduSessionInfo(v RedundantPduSessionInformation) {
	o.RedundantPduSessionInfo = &v
}

// GetTsnBridgeManCont returns the TsnBridgeManCont field value if set, zero value otherwise.
func (o *EventsNotification) GetTsnBridgeManCont() BridgeManagementContainer {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		var ret BridgeManagementContainer
		return ret
	}
	return *o.TsnBridgeManCont
}

// GetTsnBridgeManContOk returns a tuple with the TsnBridgeManCont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetTsnBridgeManContOk() (*BridgeManagementContainer, bool) {
	if o == nil || IsNil(o.TsnBridgeManCont) {
		return nil, false
	}
	return o.TsnBridgeManCont, true
}

// HasTsnBridgeManCont returns a boolean if a field has been set.
func (o *EventsNotification) HasTsnBridgeManCont() bool {
	if o != nil && !IsNil(o.TsnBridgeManCont) {
		return true
	}

	return false
}

// SetTsnBridgeManCont gets a reference to the given BridgeManagementContainer and assigns it to the TsnBridgeManCont field.
func (o *EventsNotification) SetTsnBridgeManCont(v BridgeManagementContainer) {
	o.TsnBridgeManCont = &v
}

// GetTsnPortManContDstt returns the TsnPortManContDstt field value if set, zero value otherwise.
func (o *EventsNotification) GetTsnPortManContDstt() PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		var ret PortManagementContainer
		return ret
	}
	return *o.TsnPortManContDstt
}

// GetTsnPortManContDsttOk returns a tuple with the TsnPortManContDstt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetTsnPortManContDsttOk() (*PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContDstt) {
		return nil, false
	}
	return o.TsnPortManContDstt, true
}

// HasTsnPortManContDstt returns a boolean if a field has been set.
func (o *EventsNotification) HasTsnPortManContDstt() bool {
	if o != nil && !IsNil(o.TsnPortManContDstt) {
		return true
	}

	return false
}

// SetTsnPortManContDstt gets a reference to the given PortManagementContainer and assigns it to the TsnPortManContDstt field.
func (o *EventsNotification) SetTsnPortManContDstt(v PortManagementContainer) {
	o.TsnPortManContDstt = &v
}

// GetTsnPortManContNwtts returns the TsnPortManContNwtts field value if set, zero value otherwise.
func (o *EventsNotification) GetTsnPortManContNwtts() []PortManagementContainer {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		var ret []PortManagementContainer
		return ret
	}
	return o.TsnPortManContNwtts
}

// GetTsnPortManContNwttsOk returns a tuple with the TsnPortManContNwtts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetTsnPortManContNwttsOk() ([]PortManagementContainer, bool) {
	if o == nil || IsNil(o.TsnPortManContNwtts) {
		return nil, false
	}
	return o.TsnPortManContNwtts, true
}

// HasTsnPortManContNwtts returns a boolean if a field has been set.
func (o *EventsNotification) HasTsnPortManContNwtts() bool {
	if o != nil && !IsNil(o.TsnPortManContNwtts) {
		return true
	}

	return false
}

// SetTsnPortManContNwtts gets a reference to the given []PortManagementContainer and assigns it to the TsnPortManContNwtts field.
func (o *EventsNotification) SetTsnPortManContNwtts(v []PortManagementContainer) {
	o.TsnPortManContNwtts = v
}

// GetIpv4AddrList returns the Ipv4AddrList field value if set, zero value otherwise.
func (o *EventsNotification) GetIpv4AddrList() []string {
	if o == nil || IsNil(o.Ipv4AddrList) {
		var ret []string
		return ret
	}
	return o.Ipv4AddrList
}

// GetIpv4AddrListOk returns a tuple with the Ipv4AddrList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetIpv4AddrListOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4AddrList) {
		return nil, false
	}
	return o.Ipv4AddrList, true
}

// HasIpv4AddrList returns a boolean if a field has been set.
func (o *EventsNotification) HasIpv4AddrList() bool {
	if o != nil && !IsNil(o.Ipv4AddrList) {
		return true
	}

	return false
}

// SetIpv4AddrList gets a reference to the given []string and assigns it to the Ipv4AddrList field.
func (o *EventsNotification) SetIpv4AddrList(v []string) {
	o.Ipv4AddrList = v
}

// GetIpv6PrefixList returns the Ipv6PrefixList field value if set, zero value otherwise.
func (o *EventsNotification) GetIpv6PrefixList() []Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6PrefixList) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.Ipv6PrefixList
}

// GetIpv6PrefixListOk returns a tuple with the Ipv6PrefixList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetIpv6PrefixListOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6PrefixList) {
		return nil, false
	}
	return o.Ipv6PrefixList, true
}

// HasIpv6PrefixList returns a boolean if a field has been set.
func (o *EventsNotification) HasIpv6PrefixList() bool {
	if o != nil && !IsNil(o.Ipv6PrefixList) {
		return true
	}

	return false
}

// SetIpv6PrefixList gets a reference to the given []Ipv6Prefix and assigns it to the Ipv6PrefixList field.
func (o *EventsNotification) SetIpv6PrefixList(v []Ipv6Prefix) {
	o.Ipv6PrefixList = v
}

// GetBatOffsetInfo returns the BatOffsetInfo field value if set, zero value otherwise.
func (o *EventsNotification) GetBatOffsetInfo() BatOffsetInfo {
	if o == nil || IsNil(o.BatOffsetInfo) {
		var ret BatOffsetInfo
		return ret
	}
	return *o.BatOffsetInfo
}

// GetBatOffsetInfoOk returns a tuple with the BatOffsetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetBatOffsetInfoOk() (*BatOffsetInfo, bool) {
	if o == nil || IsNil(o.BatOffsetInfo) {
		return nil, false
	}
	return o.BatOffsetInfo, true
}

// HasBatOffsetInfo returns a boolean if a field has been set.
func (o *EventsNotification) HasBatOffsetInfo() bool {
	if o != nil && !IsNil(o.BatOffsetInfo) {
		return true
	}

	return false
}

// SetBatOffsetInfo gets a reference to the given BatOffsetInfo and assigns it to the BatOffsetInfo field.
func (o *EventsNotification) SetBatOffsetInfo(v BatOffsetInfo) {
	o.BatOffsetInfo = &v
}

func (o EventsNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdReports) {
		toSerialize["adReports"] = o.AdReports
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.AddAccessInfo) {
		toSerialize["addAccessInfo"] = o.AddAccessInfo
	}
	if !IsNil(o.RelAccessInfo) {
		toSerialize["relAccessInfo"] = o.RelAccessInfo
	}
	if o.AnChargAddr.IsSet() {
		toSerialize["anChargAddr"] = o.AnChargAddr.Get()
	}
	if !IsNil(o.AnChargIds) {
		toSerialize["anChargIds"] = o.AnChargIds
	}
	if o.AnGwAddr.IsSet() {
		toSerialize["anGwAddr"] = o.AnGwAddr.Get()
	}
	if !IsNil(o.L4sReports) {
		toSerialize["l4sReports"] = o.L4sReports
	}
	toSerialize["evSubsUri"] = o.EvSubsUri
	toSerialize["evNotifs"] = o.EvNotifs
	if !IsNil(o.FailedResourcAllocReports) {
		toSerialize["failedResourcAllocReports"] = o.FailedResourcAllocReports
	}
	if !IsNil(o.SuccResourcAllocReports) {
		toSerialize["succResourcAllocReports"] = o.SuccResourcAllocReports
	}
	if !IsNil(o.NoNetLocSupp) {
		toSerialize["noNetLocSupp"] = o.NoNetLocSupp
	}
	if !IsNil(o.OutOfCredReports) {
		toSerialize["outOfCredReports"] = o.OutOfCredReports
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.QncReports) {
		toSerialize["qncReports"] = o.QncReports
	}
	if !IsNil(o.QosMonReports) {
		toSerialize["qosMonReports"] = o.QosMonReports
	}
	if !IsNil(o.QosMonDatRateReps) {
		toSerialize["qosMonDatRateReps"] = o.QosMonDatRateReps
	}
	if !IsNil(o.PdvMonReports) {
		toSerialize["pdvMonReports"] = o.PdvMonReports
	}
	if !IsNil(o.CongestReports) {
		toSerialize["congestReports"] = o.CongestReports
	}
	if !IsNil(o.RttMonReports) {
		toSerialize["rttMonReports"] = o.RttMonReports
	}
	if !IsNil(o.RanNasRelCauses) {
		toSerialize["ranNasRelCauses"] = o.RanNasRelCauses
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !IsNil(o.SatBackhaulCategory) {
		toSerialize["satBackhaulCategory"] = o.SatBackhaulCategory
	}
	if !IsNil(o.UeLoc) {
		toSerialize["ueLoc"] = o.UeLoc
	}
	if !IsNil(o.UeLocTime) {
		toSerialize["ueLocTime"] = o.UeLocTime
	}
	if !IsNil(o.UeTimeZone) {
		toSerialize["ueTimeZone"] = o.UeTimeZone
	}
	if !IsNil(o.UsgRep) {
		toSerialize["usgRep"] = o.UsgRep
	}
	if !IsNil(o.UrspEnfRep) {
		toSerialize["urspEnfRep"] = o.UrspEnfRep
	}
	if !IsNil(o.SscMode) {
		toSerialize["sscMode"] = o.SscMode
	}
	if !IsNil(o.UeReqDnn) {
		toSerialize["ueReqDnn"] = o.UeReqDnn
	}
	if !IsNil(o.RedundantPduSessionInfo) {
		toSerialize["redundantPduSessionInfo"] = o.RedundantPduSessionInfo
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
	if !IsNil(o.Ipv4AddrList) {
		toSerialize["ipv4AddrList"] = o.Ipv4AddrList
	}
	if !IsNil(o.Ipv6PrefixList) {
		toSerialize["ipv6PrefixList"] = o.Ipv6PrefixList
	}
	if !IsNil(o.BatOffsetInfo) {
		toSerialize["batOffsetInfo"] = o.BatOffsetInfo
	}
	return toSerialize, nil
}

func (o *EventsNotification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"evSubsUri",
		"evNotifs",
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

	varEventsNotification := _EventsNotification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsNotification)

	if err != nil {
		return err
	}

	*o = EventsNotification(varEventsNotification)

	return err
}

type NullableEventsNotification struct {
	value *EventsNotification
	isSet bool
}

func (v NullableEventsNotification) Get() *EventsNotification {
	return v.value
}

func (v *NullableEventsNotification) Set(val *EventsNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsNotification(val *EventsNotification) *NullableEventsNotification {
	return &NullableEventsNotification{value: val, isSet: true}
}

func (v NullableEventsNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
