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

// checks if the EventsSubscReqDataRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsSubscReqDataRm{}

// EventsSubscReqDataRm This data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type EventsSubscReqDataRm struct {
	Events []AfEventSubscription `json:"events"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri        *string                            `json:"notifUri,omitempty"`
	ReqQosMonParams []RequestedQosMonitoringParameter  `json:"reqQosMonParams,omitempty"`
	QosMon          NullableQosMonitoringInformationRm `json:"qosMon,omitempty"`
	QosMonDatRate   NullableQosMonitoringInformationRm `json:"qosMonDatRate,omitempty"`
	PdvReqMonParams []RequestedQosMonitoringParameter  `json:"pdvReqMonParams,omitempty"`
	PdvMon          NullableQosMonitoringInformationRm `json:"pdvMon,omitempty"`
	CongestMon      NullableQosMonitoringInformationRm `json:"congestMon,omitempty"`
	ReqAnis         []RequiredAccessInfo               `json:"reqAnis,omitempty"`
	UsgThres        NullableUsageThresholdRm           `json:"usgThres,omitempty"`
	NotifCorreId    *string                            `json:"notifCorreId,omitempty"`
	// Indicates whether the direct event notification is requested (true) or not (false) for the provided and/or previously provided QoS monitoring parameters.
	DirectNotifInd NullableBool `json:"directNotifInd,omitempty"`
	// This data type is defined in the same way as the 'AverWindow' data type, but with the OpenAPI 'nullable: true' property.
	AvrgWndw NullableInt32 `json:"avrgWndw,omitempty"`
}

type _EventsSubscReqDataRm EventsSubscReqDataRm

// NewEventsSubscReqDataRm instantiates a new EventsSubscReqDataRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsSubscReqDataRm(events []AfEventSubscription) *EventsSubscReqDataRm {
	this := EventsSubscReqDataRm{}
	this.Events = events
	var avrgWndw int32 = 2000
	this.AvrgWndw = *NewNullableInt32(&avrgWndw)
	return &this
}

// NewEventsSubscReqDataRmWithDefaults instantiates a new EventsSubscReqDataRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsSubscReqDataRmWithDefaults() *EventsSubscReqDataRm {
	this := EventsSubscReqDataRm{}
	var avrgWndw int32 = 2000
	this.AvrgWndw = *NewNullableInt32(&avrgWndw)
	return &this
}

// GetEvents returns the Events field value
func (o *EventsSubscReqDataRm) GetEvents() []AfEventSubscription {
	if o == nil {
		var ret []AfEventSubscription
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetEventsOk() ([]AfEventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsSubscReqDataRm) SetEvents(v []AfEventSubscription) {
	o.Events = v
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise.
func (o *EventsSubscReqDataRm) GetNotifUri() string {
	if o == nil || IsNil(o.NotifUri) {
		var ret string
		return ret
	}
	return *o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotifUri) {
		return nil, false
	}
	return o.NotifUri, true
}

// HasNotifUri returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasNotifUri() bool {
	if o != nil && !IsNil(o.NotifUri) {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given string and assigns it to the NotifUri field.
func (o *EventsSubscReqDataRm) SetNotifUri(v string) {
	o.NotifUri = &v
}

// GetReqQosMonParams returns the ReqQosMonParams field value if set, zero value otherwise.
func (o *EventsSubscReqDataRm) GetReqQosMonParams() []RequestedQosMonitoringParameter {
	if o == nil || IsNil(o.ReqQosMonParams) {
		var ret []RequestedQosMonitoringParameter
		return ret
	}
	return o.ReqQosMonParams
}

// GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetReqQosMonParamsOk() ([]RequestedQosMonitoringParameter, bool) {
	if o == nil || IsNil(o.ReqQosMonParams) {
		return nil, false
	}
	return o.ReqQosMonParams, true
}

// HasReqQosMonParams returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasReqQosMonParams() bool {
	if o != nil && !IsNil(o.ReqQosMonParams) {
		return true
	}

	return false
}

// SetReqQosMonParams gets a reference to the given []RequestedQosMonitoringParameter and assigns it to the ReqQosMonParams field.
func (o *EventsSubscReqDataRm) SetReqQosMonParams(v []RequestedQosMonitoringParameter) {
	o.ReqQosMonParams = v
}

// GetQosMon returns the QosMon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetQosMon() QosMonitoringInformationRm {
	if o == nil || IsNil(o.QosMon.Get()) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.QosMon.Get()
}

// GetQosMonOk returns a tuple with the QosMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetQosMonOk() (*QosMonitoringInformationRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosMon.Get(), o.QosMon.IsSet()
}

// HasQosMon returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasQosMon() bool {
	if o != nil && o.QosMon.IsSet() {
		return true
	}

	return false
}

// SetQosMon gets a reference to the given NullableQosMonitoringInformationRm and assigns it to the QosMon field.
func (o *EventsSubscReqDataRm) SetQosMon(v QosMonitoringInformationRm) {
	o.QosMon.Set(&v)
}

// SetQosMonNil sets the value for QosMon to be an explicit nil
func (o *EventsSubscReqDataRm) SetQosMonNil() {
	o.QosMon.Set(nil)
}

// UnsetQosMon ensures that no value is present for QosMon, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetQosMon() {
	o.QosMon.Unset()
}

// GetQosMonDatRate returns the QosMonDatRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetQosMonDatRate() QosMonitoringInformationRm {
	if o == nil || IsNil(o.QosMonDatRate.Get()) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.QosMonDatRate.Get()
}

// GetQosMonDatRateOk returns a tuple with the QosMonDatRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetQosMonDatRateOk() (*QosMonitoringInformationRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.QosMonDatRate.Get(), o.QosMonDatRate.IsSet()
}

// HasQosMonDatRate returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasQosMonDatRate() bool {
	if o != nil && o.QosMonDatRate.IsSet() {
		return true
	}

	return false
}

// SetQosMonDatRate gets a reference to the given NullableQosMonitoringInformationRm and assigns it to the QosMonDatRate field.
func (o *EventsSubscReqDataRm) SetQosMonDatRate(v QosMonitoringInformationRm) {
	o.QosMonDatRate.Set(&v)
}

// SetQosMonDatRateNil sets the value for QosMonDatRate to be an explicit nil
func (o *EventsSubscReqDataRm) SetQosMonDatRateNil() {
	o.QosMonDatRate.Set(nil)
}

// UnsetQosMonDatRate ensures that no value is present for QosMonDatRate, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetQosMonDatRate() {
	o.QosMonDatRate.Unset()
}

// GetPdvReqMonParams returns the PdvReqMonParams field value if set, zero value otherwise.
func (o *EventsSubscReqDataRm) GetPdvReqMonParams() []RequestedQosMonitoringParameter {
	if o == nil || IsNil(o.PdvReqMonParams) {
		var ret []RequestedQosMonitoringParameter
		return ret
	}
	return o.PdvReqMonParams
}

// GetPdvReqMonParamsOk returns a tuple with the PdvReqMonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetPdvReqMonParamsOk() ([]RequestedQosMonitoringParameter, bool) {
	if o == nil || IsNil(o.PdvReqMonParams) {
		return nil, false
	}
	return o.PdvReqMonParams, true
}

// HasPdvReqMonParams returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasPdvReqMonParams() bool {
	if o != nil && !IsNil(o.PdvReqMonParams) {
		return true
	}

	return false
}

// SetPdvReqMonParams gets a reference to the given []RequestedQosMonitoringParameter and assigns it to the PdvReqMonParams field.
func (o *EventsSubscReqDataRm) SetPdvReqMonParams(v []RequestedQosMonitoringParameter) {
	o.PdvReqMonParams = v
}

// GetPdvMon returns the PdvMon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetPdvMon() QosMonitoringInformationRm {
	if o == nil || IsNil(o.PdvMon.Get()) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.PdvMon.Get()
}

// GetPdvMonOk returns a tuple with the PdvMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetPdvMonOk() (*QosMonitoringInformationRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.PdvMon.Get(), o.PdvMon.IsSet()
}

// HasPdvMon returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasPdvMon() bool {
	if o != nil && o.PdvMon.IsSet() {
		return true
	}

	return false
}

// SetPdvMon gets a reference to the given NullableQosMonitoringInformationRm and assigns it to the PdvMon field.
func (o *EventsSubscReqDataRm) SetPdvMon(v QosMonitoringInformationRm) {
	o.PdvMon.Set(&v)
}

// SetPdvMonNil sets the value for PdvMon to be an explicit nil
func (o *EventsSubscReqDataRm) SetPdvMonNil() {
	o.PdvMon.Set(nil)
}

// UnsetPdvMon ensures that no value is present for PdvMon, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetPdvMon() {
	o.PdvMon.Unset()
}

// GetCongestMon returns the CongestMon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetCongestMon() QosMonitoringInformationRm {
	if o == nil || IsNil(o.CongestMon.Get()) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.CongestMon.Get()
}

// GetCongestMonOk returns a tuple with the CongestMon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetCongestMonOk() (*QosMonitoringInformationRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.CongestMon.Get(), o.CongestMon.IsSet()
}

// HasCongestMon returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasCongestMon() bool {
	if o != nil && o.CongestMon.IsSet() {
		return true
	}

	return false
}

// SetCongestMon gets a reference to the given NullableQosMonitoringInformationRm and assigns it to the CongestMon field.
func (o *EventsSubscReqDataRm) SetCongestMon(v QosMonitoringInformationRm) {
	o.CongestMon.Set(&v)
}

// SetCongestMonNil sets the value for CongestMon to be an explicit nil
func (o *EventsSubscReqDataRm) SetCongestMonNil() {
	o.CongestMon.Set(nil)
}

// UnsetCongestMon ensures that no value is present for CongestMon, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetCongestMon() {
	o.CongestMon.Unset()
}

// GetReqAnis returns the ReqAnis field value if set, zero value otherwise.
func (o *EventsSubscReqDataRm) GetReqAnis() []RequiredAccessInfo {
	if o == nil || IsNil(o.ReqAnis) {
		var ret []RequiredAccessInfo
		return ret
	}
	return o.ReqAnis
}

// GetReqAnisOk returns a tuple with the ReqAnis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetReqAnisOk() ([]RequiredAccessInfo, bool) {
	if o == nil || IsNil(o.ReqAnis) {
		return nil, false
	}
	return o.ReqAnis, true
}

// HasReqAnis returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasReqAnis() bool {
	if o != nil && !IsNil(o.ReqAnis) {
		return true
	}

	return false
}

// SetReqAnis gets a reference to the given []RequiredAccessInfo and assigns it to the ReqAnis field.
func (o *EventsSubscReqDataRm) SetReqAnis(v []RequiredAccessInfo) {
	o.ReqAnis = v
}

// GetUsgThres returns the UsgThres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetUsgThres() UsageThresholdRm {
	if o == nil || IsNil(o.UsgThres.Get()) {
		var ret UsageThresholdRm
		return ret
	}
	return *o.UsgThres.Get()
}

// GetUsgThresOk returns a tuple with the UsgThres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetUsgThresOk() (*UsageThresholdRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsgThres.Get(), o.UsgThres.IsSet()
}

// HasUsgThres returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasUsgThres() bool {
	if o != nil && o.UsgThres.IsSet() {
		return true
	}

	return false
}

// SetUsgThres gets a reference to the given NullableUsageThresholdRm and assigns it to the UsgThres field.
func (o *EventsSubscReqDataRm) SetUsgThres(v UsageThresholdRm) {
	o.UsgThres.Set(&v)
}

// SetUsgThresNil sets the value for UsgThres to be an explicit nil
func (o *EventsSubscReqDataRm) SetUsgThresNil() {
	o.UsgThres.Set(nil)
}

// UnsetUsgThres ensures that no value is present for UsgThres, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetUsgThres() {
	o.UsgThres.Unset()
}

// GetNotifCorreId returns the NotifCorreId field value if set, zero value otherwise.
func (o *EventsSubscReqDataRm) GetNotifCorreId() string {
	if o == nil || IsNil(o.NotifCorreId) {
		var ret string
		return ret
	}
	return *o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsSubscReqDataRm) GetNotifCorreIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifCorreId) {
		return nil, false
	}
	return o.NotifCorreId, true
}

// HasNotifCorreId returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasNotifCorreId() bool {
	if o != nil && !IsNil(o.NotifCorreId) {
		return true
	}

	return false
}

// SetNotifCorreId gets a reference to the given string and assigns it to the NotifCorreId field.
func (o *EventsSubscReqDataRm) SetNotifCorreId(v string) {
	o.NotifCorreId = &v
}

// GetDirectNotifInd returns the DirectNotifInd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetDirectNotifInd() bool {
	if o == nil || IsNil(o.DirectNotifInd.Get()) {
		var ret bool
		return ret
	}
	return *o.DirectNotifInd.Get()
}

// GetDirectNotifIndOk returns a tuple with the DirectNotifInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetDirectNotifIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectNotifInd.Get(), o.DirectNotifInd.IsSet()
}

// HasDirectNotifInd returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasDirectNotifInd() bool {
	if o != nil && o.DirectNotifInd.IsSet() {
		return true
	}

	return false
}

// SetDirectNotifInd gets a reference to the given NullableBool and assigns it to the DirectNotifInd field.
func (o *EventsSubscReqDataRm) SetDirectNotifInd(v bool) {
	o.DirectNotifInd.Set(&v)
}

// SetDirectNotifIndNil sets the value for DirectNotifInd to be an explicit nil
func (o *EventsSubscReqDataRm) SetDirectNotifIndNil() {
	o.DirectNotifInd.Set(nil)
}

// UnsetDirectNotifInd ensures that no value is present for DirectNotifInd, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetDirectNotifInd() {
	o.DirectNotifInd.Unset()
}

// GetAvrgWndw returns the AvrgWndw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventsSubscReqDataRm) GetAvrgWndw() int32 {
	if o == nil || IsNil(o.AvrgWndw.Get()) {
		var ret int32
		return ret
	}
	return *o.AvrgWndw.Get()
}

// GetAvrgWndwOk returns a tuple with the AvrgWndw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventsSubscReqDataRm) GetAvrgWndwOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvrgWndw.Get(), o.AvrgWndw.IsSet()
}

// HasAvrgWndw returns a boolean if a field has been set.
func (o *EventsSubscReqDataRm) HasAvrgWndw() bool {
	if o != nil && o.AvrgWndw.IsSet() {
		return true
	}

	return false
}

// SetAvrgWndw gets a reference to the given NullableInt32 and assigns it to the AvrgWndw field.
func (o *EventsSubscReqDataRm) SetAvrgWndw(v int32) {
	o.AvrgWndw.Set(&v)
}

// SetAvrgWndwNil sets the value for AvrgWndw to be an explicit nil
func (o *EventsSubscReqDataRm) SetAvrgWndwNil() {
	o.AvrgWndw.Set(nil)
}

// UnsetAvrgWndw ensures that no value is present for AvrgWndw, not even an explicit nil
func (o *EventsSubscReqDataRm) UnsetAvrgWndw() {
	o.AvrgWndw.Unset()
}

func (o EventsSubscReqDataRm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsSubscReqDataRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.NotifUri) {
		toSerialize["notifUri"] = o.NotifUri
	}
	if !IsNil(o.ReqQosMonParams) {
		toSerialize["reqQosMonParams"] = o.ReqQosMonParams
	}
	if o.QosMon.IsSet() {
		toSerialize["qosMon"] = o.QosMon.Get()
	}
	if o.QosMonDatRate.IsSet() {
		toSerialize["qosMonDatRate"] = o.QosMonDatRate.Get()
	}
	if !IsNil(o.PdvReqMonParams) {
		toSerialize["pdvReqMonParams"] = o.PdvReqMonParams
	}
	if o.PdvMon.IsSet() {
		toSerialize["pdvMon"] = o.PdvMon.Get()
	}
	if o.CongestMon.IsSet() {
		toSerialize["congestMon"] = o.CongestMon.Get()
	}
	if !IsNil(o.ReqAnis) {
		toSerialize["reqAnis"] = o.ReqAnis
	}
	if o.UsgThres.IsSet() {
		toSerialize["usgThres"] = o.UsgThres.Get()
	}
	if !IsNil(o.NotifCorreId) {
		toSerialize["notifCorreId"] = o.NotifCorreId
	}
	if o.DirectNotifInd.IsSet() {
		toSerialize["directNotifInd"] = o.DirectNotifInd.Get()
	}
	if o.AvrgWndw.IsSet() {
		toSerialize["avrgWndw"] = o.AvrgWndw.Get()
	}
	return toSerialize, nil
}

func (o *EventsSubscReqDataRm) UnmarshalJSON(data []byte) (err error) {
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

	varEventsSubscReqDataRm := _EventsSubscReqDataRm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventsSubscReqDataRm)

	if err != nil {
		return err
	}

	*o = EventsSubscReqDataRm(varEventsSubscReqDataRm)

	return err
}

type NullableEventsSubscReqDataRm struct {
	value *EventsSubscReqDataRm
	isSet bool
}

func (v NullableEventsSubscReqDataRm) Get() *EventsSubscReqDataRm {
	return v.value
}

func (v *NullableEventsSubscReqDataRm) Set(val *EventsSubscReqDataRm) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsSubscReqDataRm) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsSubscReqDataRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsSubscReqDataRm(val *EventsSubscReqDataRm) *NullableEventsSubscReqDataRm {
	return &NullableEventsSubscReqDataRm{value: val, isSet: true}
}

func (v NullableEventsSubscReqDataRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsSubscReqDataRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
