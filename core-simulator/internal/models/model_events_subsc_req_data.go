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

// checks if the EventsSubscReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsSubscReqData{}

// EventsSubscReqData Identifies the events the application subscribes to.
type EventsSubscReqData struct {
	Events []AfEventSubscription `json:"events"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri        *string                           `json:"notifUri,omitempty"`
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty"`
	QosMon          *QosMonitoringInformation         `json:"qosMon,omitempty"`
	QosMonDatRate   *QosMonitoringInformation         `json:"qosMonDatRate,omitempty"`
	PdvReqMonParams []RequestedQosMonitoringParameter `json:"pdvReqMonParams,omitempty"`
	PdvMon          *QosMonitoringInformation         `json:"pdvMon,omitempty"`
	CongestMon      *QosMonitoringInformation         `json:"congestMon,omitempty"`
	RttMon          *QosMonitoringInformation         `json:"rttMon,omitempty"`
	ReqAnis         []RequiredAccessInfo              `json:"reqAnis,omitempty"`
	UsgThres        *UsageThreshold                   `json:"usgThres,omitempty"`
	NotifCorreId    *string                           `json:"notifCorreId,omitempty"`
	AfAppIds        []string                          `json:"afAppIds,omitempty"`
	// Indicates whether the direct event notification is requested (true) or not (false) for the provided QoS monitoring parameters. Default value is false.
	DirectNotifInd *bool `json:"directNotifInd,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AvrgWndw *int32 `json:"avrgWndw,omitempty"`
}

type _EventsSubscReqData EventsSubscReqData

// NewEventsSubscReqData instantiates a new EventsSubscReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsSubscReqData(events []AfEventSubscription) *EventsSubscReqData {
	this := EventsSubscReqData{}
	this.Events = events
	var avrgWndw int32 = 2000
	this.AvrgWndw = &avrgWndw
	return &this
}

// NewEventsSubscReqDataWithDefaults instantiates a new EventsSubscReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsSubscReqDataWithDefaults() *EventsSubscReqData {
	this := EventsSubscReqData{}
	var avrgWndw int32 = 2000
	this.AvrgWndw = &avrgWndw
	return &this
}

// GetEvents returns the Events field value
func (o *EventsSubscReqData) GetEvents() []AfEventSubscription {
	if o == nil {
		var ret []AfEventSubscription
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetEventsOk() ([]AfEventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsSubscReqData) SetEvents(v []AfEventSubscription) {
	o.Events = v
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetNotifUri() string {
	if o == nil || IsNil(o.NotifUri) {
		var ret string
		return ret
	}
	return *o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotifUri) {
		return nil, false
	}
	return o.NotifUri, true
}

// HasNotifUri returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasNotifUri() bool {
	if o != nil && !IsNil(o.NotifUri) {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given string and assigns it to the NotifUri field.
func (o *EventsSubscReqData) SetNotifUri(v string) {
	o.NotifUri = &v
}

// GetReqQosMonParams returns the ReqQosMonParams field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetReqQosMonParams() []RequestedQosMonitoringParameter {
	if o == nil || IsNil(o.ReqQosMonParams) {
		var ret []RequestedQosMonitoringParameter
		return ret
	}
	return o.ReqQosMonParams
}

// GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetReqQosMonParamsOk() ([]RequestedQosMonitoringParameter, bool) {
	if o == nil || IsNil(o.ReqQosMonParams) {
		return nil, false
	}
	return o.ReqQosMonParams, true
}

// HasReqQosMonParams returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasReqQosMonParams() bool {
	if o != nil && !IsNil(o.ReqQosMonParams) {
		return true
	}

	return false
}

// SetReqQosMonParams gets a reference to the given []RequestedQosMonitoringParameter and assigns it to the ReqQosMonParams field.
func (o *EventsSubscReqData) SetReqQosMonParams(v []RequestedQosMonitoringParameter) {
	o.ReqQosMonParams = v
}

// GetQosMon returns the QosMon field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetQosMon() QosMonitoringInformation {
	if o == nil || IsNil(o.QosMon) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.QosMon
}

// GetQosMonOk returns a tuple with the QosMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetQosMonOk() (*QosMonitoringInformation, bool) {
	if o == nil || IsNil(o.QosMon) {
		return nil, false
	}
	return o.QosMon, true
}

// HasQosMon returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasQosMon() bool {
	if o != nil && !IsNil(o.QosMon) {
		return true
	}

	return false
}

// SetQosMon gets a reference to the given QosMonitoringInformation and assigns it to the QosMon field.
func (o *EventsSubscReqData) SetQosMon(v QosMonitoringInformation) {
	o.QosMon = &v
}

// GetQosMonDatRate returns the QosMonDatRate field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetQosMonDatRate() QosMonitoringInformation {
	if o == nil || IsNil(o.QosMonDatRate) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.QosMonDatRate
}

// GetQosMonDatRateOk returns a tuple with the QosMonDatRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetQosMonDatRateOk() (*QosMonitoringInformation, bool) {
	if o == nil || IsNil(o.QosMonDatRate) {
		return nil, false
	}
	return o.QosMonDatRate, true
}

// HasQosMonDatRate returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasQosMonDatRate() bool {
	if o != nil && !IsNil(o.QosMonDatRate) {
		return true
	}

	return false
}

// SetQosMonDatRate gets a reference to the given QosMonitoringInformation and assigns it to the QosMonDatRate field.
func (o *EventsSubscReqData) SetQosMonDatRate(v QosMonitoringInformation) {
	o.QosMonDatRate = &v
}

// GetPdvReqMonParams returns the PdvReqMonParams field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetPdvReqMonParams() []RequestedQosMonitoringParameter {
	if o == nil || IsNil(o.PdvReqMonParams) {
		var ret []RequestedQosMonitoringParameter
		return ret
	}
	return o.PdvReqMonParams
}

// GetPdvReqMonParamsOk returns a tuple with the PdvReqMonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetPdvReqMonParamsOk() ([]RequestedQosMonitoringParameter, bool) {
	if o == nil || IsNil(o.PdvReqMonParams) {
		return nil, false
	}
	return o.PdvReqMonParams, true
}

// HasPdvReqMonParams returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasPdvReqMonParams() bool {
	if o != nil && !IsNil(o.PdvReqMonParams) {
		return true
	}

	return false
}

// SetPdvReqMonParams gets a reference to the given []RequestedQosMonitoringParameter and assigns it to the PdvReqMonParams field.
func (o *EventsSubscReqData) SetPdvReqMonParams(v []RequestedQosMonitoringParameter) {
	o.PdvReqMonParams = v
}

// GetPdvMon returns the PdvMon field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetPdvMon() QosMonitoringInformation {
	if o == nil || IsNil(o.PdvMon) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.PdvMon
}

// GetPdvMonOk returns a tuple with the PdvMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetPdvMonOk() (*QosMonitoringInformation, bool) {
	if o == nil || IsNil(o.PdvMon) {
		return nil, false
	}
	return o.PdvMon, true
}

// HasPdvMon returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasPdvMon() bool {
	if o != nil && !IsNil(o.PdvMon) {
		return true
	}

	return false
}

// SetPdvMon gets a reference to the given QosMonitoringInformation and assigns it to the PdvMon field.
func (o *EventsSubscReqData) SetPdvMon(v QosMonitoringInformation) {
	o.PdvMon = &v
}

// GetCongestMon returns the CongestMon field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetCongestMon() QosMonitoringInformation {
	if o == nil || IsNil(o.CongestMon) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.CongestMon
}

// GetCongestMonOk returns a tuple with the CongestMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetCongestMonOk() (*QosMonitoringInformation, bool) {
	if o == nil || IsNil(o.CongestMon) {
		return nil, false
	}
	return o.CongestMon, true
}

// HasCongestMon returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasCongestMon() bool {
	if o != nil && !IsNil(o.CongestMon) {
		return true
	}

	return false
}

// SetCongestMon gets a reference to the given QosMonitoringInformation and assigns it to the CongestMon field.
func (o *EventsSubscReqData) SetCongestMon(v QosMonitoringInformation) {
	o.CongestMon = &v
}

// GetRttMon returns the RttMon field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetRttMon() QosMonitoringInformation {
	if o == nil || IsNil(o.RttMon) {
		var ret QosMonitoringInformation
		return ret
	}
	return *o.RttMon
}

// GetRttMonOk returns a tuple with the RttMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetRttMonOk() (*QosMonitoringInformation, bool) {
	if o == nil || IsNil(o.RttMon) {
		return nil, false
	}
	return o.RttMon, true
}

// HasRttMon returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasRttMon() bool {
	if o != nil && !IsNil(o.RttMon) {
		return true
	}

	return false
}

// SetRttMon gets a reference to the given QosMonitoringInformation and assigns it to the RttMon field.
func (o *EventsSubscReqData) SetRttMon(v QosMonitoringInformation) {
	o.RttMon = &v
}

// GetReqAnis returns the ReqAnis field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetReqAnis() []RequiredAccessInfo {
	if o == nil || IsNil(o.ReqAnis) {
		var ret []RequiredAccessInfo
		return ret
	}
	return o.ReqAnis
}

// GetReqAnisOk returns a tuple with the ReqAnis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetReqAnisOk() ([]RequiredAccessInfo, bool) {
	if o == nil || IsNil(o.ReqAnis) {
		return nil, false
	}
	return o.ReqAnis, true
}

// HasReqAnis returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasReqAnis() bool {
	if o != nil && !IsNil(o.ReqAnis) {
		return true
	}

	return false
}

// SetReqAnis gets a reference to the given []RequiredAccessInfo and assigns it to the ReqAnis field.
func (o *EventsSubscReqData) SetReqAnis(v []RequiredAccessInfo) {
	o.ReqAnis = v
}

// GetUsgThres returns the UsgThres field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetUsgThres() UsageThreshold {
	if o == nil || IsNil(o.UsgThres) {
		var ret UsageThreshold
		return ret
	}
	return *o.UsgThres
}

// GetUsgThresOk returns a tuple with the UsgThres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetUsgThresOk() (*UsageThreshold, bool) {
	if o == nil || IsNil(o.UsgThres) {
		return nil, false
	}
	return o.UsgThres, true
}

// HasUsgThres returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasUsgThres() bool {
	if o != nil && !IsNil(o.UsgThres) {
		return true
	}

	return false
}

// SetUsgThres gets a reference to the given UsageThreshold and assigns it to the UsgThres field.
func (o *EventsSubscReqData) SetUsgThres(v UsageThreshold) {
	o.UsgThres = &v
}

// GetNotifCorreId returns the NotifCorreId field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetNotifCorreId() string {
	if o == nil || IsNil(o.NotifCorreId) {
		var ret string
		return ret
	}
	return *o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetNotifCorreIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifCorreId) {
		return nil, false
	}
	return o.NotifCorreId, true
}

// HasNotifCorreId returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasNotifCorreId() bool {
	if o != nil && !IsNil(o.NotifCorreId) {
		return true
	}

	return false
}

// SetNotifCorreId gets a reference to the given string and assigns it to the NotifCorreId field.
func (o *EventsSubscReqData) SetNotifCorreId(v string) {
	o.NotifCorreId = &v
}

// GetAfAppIds returns the AfAppIds field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetAfAppIds() []string {
	if o == nil || IsNil(o.AfAppIds) {
		var ret []string
		return ret
	}
	return o.AfAppIds
}

// GetAfAppIdsOk returns a tuple with the AfAppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetAfAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AfAppIds) {
		return nil, false
	}
	return o.AfAppIds, true
}

// HasAfAppIds returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasAfAppIds() bool {
	if o != nil && !IsNil(o.AfAppIds) {
		return true
	}

	return false
}

// SetAfAppIds gets a reference to the given []string and assigns it to the AfAppIds field.
func (o *EventsSubscReqData) SetAfAppIds(v []string) {
	o.AfAppIds = v
}

// GetDirectNotifInd returns the DirectNotifInd field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetDirectNotifInd() bool {
	if o == nil || IsNil(o.DirectNotifInd) {
		var ret bool
		return ret
	}
	return *o.DirectNotifInd
}

// GetDirectNotifIndOk returns a tuple with the DirectNotifInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetDirectNotifIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectNotifInd) {
		return nil, false
	}
	return o.DirectNotifInd, true
}

// HasDirectNotifInd returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasDirectNotifInd() bool {
	if o != nil && !IsNil(o.DirectNotifInd) {
		return true
	}

	return false
}

// SetDirectNotifInd gets a reference to the given bool and assigns it to the DirectNotifInd field.
func (o *EventsSubscReqData) SetDirectNotifInd(v bool) {
	o.DirectNotifInd = &v
}

// GetAvrgWndw returns the AvrgWndw field value if set, zero value otherwise.
func (o *EventsSubscReqData) GetAvrgWndw() int32 {
	if o == nil || IsNil(o.AvrgWndw) {
		var ret int32
		return ret
	}
	return *o.AvrgWndw
}

// GetAvrgWndwOk returns a tuple with the AvrgWndw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqData) GetAvrgWndwOk() (*int32, bool) {
	if o == nil || IsNil(o.AvrgWndw) {
		return nil, false
	}
	return o.AvrgWndw, true
}

// HasAvrgWndw returns a boolean if a field has been set.
func (o *EventsSubscReqData) HasAvrgWndw() bool {
	if o != nil && !IsNil(o.AvrgWndw) {
		return true
	}

	return false
}

// SetAvrgWndw gets a reference to the given int32 and assigns it to the AvrgWndw field.
func (o *EventsSubscReqData) SetAvrgWndw(v int32) {
	o.AvrgWndw = &v
}

func (o EventsSubscReqData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsSubscReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.NotifUri) {
		toSerialize["notifUri"] = o.NotifUri
	}
	if !IsNil(o.ReqQosMonParams) {
		toSerialize["reqQosMonParams"] = o.ReqQosMonParams
	}
	if !IsNil(o.QosMon) {
		toSerialize["qosMon"] = o.QosMon
	}
	if !IsNil(o.QosMonDatRate) {
		toSerialize["qosMonDatRate"] = o.QosMonDatRate
	}
	if !IsNil(o.PdvReqMonParams) {
		toSerialize["pdvReqMonParams"] = o.PdvReqMonParams
	}
	if !IsNil(o.PdvMon) {
		toSerialize["pdvMon"] = o.PdvMon
	}
	if !IsNil(o.CongestMon) {
		toSerialize["congestMon"] = o.CongestMon
	}
	if !IsNil(o.RttMon) {
		toSerialize["rttMon"] = o.RttMon
	}
	if !IsNil(o.ReqAnis) {
		toSerialize["reqAnis"] = o.ReqAnis
	}
	if !IsNil(o.UsgThres) {
		toSerialize["usgThres"] = o.UsgThres
	}
	if !IsNil(o.NotifCorreId) {
		toSerialize["notifCorreId"] = o.NotifCorreId
	}
	if !IsNil(o.AfAppIds) {
		toSerialize["afAppIds"] = o.AfAppIds
	}
	if !IsNil(o.DirectNotifInd) {
		toSerialize["directNotifInd"] = o.DirectNotifInd
	}
	if !IsNil(o.AvrgWndw) {
		toSerialize["avrgWndw"] = o.AvrgWndw
	}
	return toSerialize, nil
}

func (o *EventsSubscReqData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varEventsSubscReqData := _EventsSubscReqData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsSubscReqData)

	if err != nil {
		return err
	}

	*o = EventsSubscReqData(varEventsSubscReqData)

	return err
}

type NullableEventsSubscReqData struct {
	value *EventsSubscReqData
	isSet bool
}

func (v NullableEventsSubscReqData) Get() *EventsSubscReqData {
	return v.value
}

func (v *NullableEventsSubscReqData) Set(val *EventsSubscReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsSubscReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsSubscReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsSubscReqData(val *EventsSubscReqData) *NullableEventsSubscReqData {
	return &NullableEventsSubscReqData{value: val, isSet: true}
}

func (v NullableEventsSubscReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsSubscReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
