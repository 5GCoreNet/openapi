/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
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
	ReqAnis         []RequiredAccessInfo              `json:"reqAnis,omitempty"`
	UsgThres        *UsageThreshold                   `json:"usgThres,omitempty"`
	NotifCorreId    *string                           `json:"notifCorreId,omitempty"`
	AfAppIds        []string                          `json:"afAppIds,omitempty"`
	DirectNotifInd  *bool                             `json:"directNotifInd,omitempty"`
}

// NewEventsSubscReqData instantiates a new EventsSubscReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsSubscReqData(events []AfEventSubscription) *EventsSubscReqData {
	this := EventsSubscReqData{}
	this.Events = events
	return &this
}

// NewEventsSubscReqDataWithDefaults instantiates a new EventsSubscReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsSubscReqDataWithDefaults() *EventsSubscReqData {
	this := EventsSubscReqData{}
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
	return toSerialize, nil
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
