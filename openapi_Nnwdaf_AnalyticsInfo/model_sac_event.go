/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the SACEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SACEvent{}

// SACEvent Describes an event to be subscribed
type SACEvent struct {
	EventType SACEventType `json:"eventType"`
	EventTrigger *SACEventTrigger `json:"eventTrigger,omitempty"`
	EventFilter []Snssai `json:"eventFilter"`
	// indicating a time in seconds.
	NotificationPeriod *int32 `json:"notificationPeriod,omitempty"`
	NotifThreshold *SACInfo `json:"notifThreshold,omitempty"`
	ImmediateFlag *bool `json:"immediateFlag,omitempty"`
}

// NewSACEvent instantiates a new SACEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACEvent(eventType SACEventType, eventFilter []Snssai) *SACEvent {
	this := SACEvent{}
	this.EventType = eventType
	this.EventFilter = eventFilter
	var immediateFlag bool = false
	this.ImmediateFlag = &immediateFlag
	return &this
}

// NewSACEventWithDefaults instantiates a new SACEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACEventWithDefaults() *SACEvent {
	this := SACEvent{}
	var immediateFlag bool = false
	this.ImmediateFlag = &immediateFlag
	return &this
}

// GetEventType returns the EventType field value
func (o *SACEvent) GetEventType() SACEventType {
	if o == nil {
		var ret SACEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SACEvent) GetEventTypeOk() (*SACEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SACEvent) SetEventType(v SACEventType) {
	o.EventType = v
}

// GetEventTrigger returns the EventTrigger field value if set, zero value otherwise.
func (o *SACEvent) GetEventTrigger() SACEventTrigger {
	if o == nil || IsNil(o.EventTrigger) {
		var ret SACEventTrigger
		return ret
	}
	return *o.EventTrigger
}

// GetEventTriggerOk returns a tuple with the EventTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEvent) GetEventTriggerOk() (*SACEventTrigger, bool) {
	if o == nil || IsNil(o.EventTrigger) {
		return nil, false
	}
	return o.EventTrigger, true
}

// HasEventTrigger returns a boolean if a field has been set.
func (o *SACEvent) HasEventTrigger() bool {
	if o != nil && !IsNil(o.EventTrigger) {
		return true
	}

	return false
}

// SetEventTrigger gets a reference to the given SACEventTrigger and assigns it to the EventTrigger field.
func (o *SACEvent) SetEventTrigger(v SACEventTrigger) {
	o.EventTrigger = &v
}

// GetEventFilter returns the EventFilter field value
func (o *SACEvent) GetEventFilter() []Snssai {
	if o == nil {
		var ret []Snssai
		return ret
	}

	return o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value
// and a boolean to check if the value has been set.
func (o *SACEvent) GetEventFilterOk() ([]Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventFilter, true
}

// SetEventFilter sets field value
func (o *SACEvent) SetEventFilter(v []Snssai) {
	o.EventFilter = v
}

// GetNotificationPeriod returns the NotificationPeriod field value if set, zero value otherwise.
func (o *SACEvent) GetNotificationPeriod() int32 {
	if o == nil || IsNil(o.NotificationPeriod) {
		var ret int32
		return ret
	}
	return *o.NotificationPeriod
}

// GetNotificationPeriodOk returns a tuple with the NotificationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEvent) GetNotificationPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.NotificationPeriod) {
		return nil, false
	}
	return o.NotificationPeriod, true
}

// HasNotificationPeriod returns a boolean if a field has been set.
func (o *SACEvent) HasNotificationPeriod() bool {
	if o != nil && !IsNil(o.NotificationPeriod) {
		return true
	}

	return false
}

// SetNotificationPeriod gets a reference to the given int32 and assigns it to the NotificationPeriod field.
func (o *SACEvent) SetNotificationPeriod(v int32) {
	o.NotificationPeriod = &v
}

// GetNotifThreshold returns the NotifThreshold field value if set, zero value otherwise.
func (o *SACEvent) GetNotifThreshold() SACInfo {
	if o == nil || IsNil(o.NotifThreshold) {
		var ret SACInfo
		return ret
	}
	return *o.NotifThreshold
}

// GetNotifThresholdOk returns a tuple with the NotifThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEvent) GetNotifThresholdOk() (*SACInfo, bool) {
	if o == nil || IsNil(o.NotifThreshold) {
		return nil, false
	}
	return o.NotifThreshold, true
}

// HasNotifThreshold returns a boolean if a field has been set.
func (o *SACEvent) HasNotifThreshold() bool {
	if o != nil && !IsNil(o.NotifThreshold) {
		return true
	}

	return false
}

// SetNotifThreshold gets a reference to the given SACInfo and assigns it to the NotifThreshold field.
func (o *SACEvent) SetNotifThreshold(v SACInfo) {
	o.NotifThreshold = &v
}

// GetImmediateFlag returns the ImmediateFlag field value if set, zero value otherwise.
func (o *SACEvent) GetImmediateFlag() bool {
	if o == nil || IsNil(o.ImmediateFlag) {
		var ret bool
		return ret
	}
	return *o.ImmediateFlag
}

// GetImmediateFlagOk returns a tuple with the ImmediateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEvent) GetImmediateFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmediateFlag) {
		return nil, false
	}
	return o.ImmediateFlag, true
}

// HasImmediateFlag returns a boolean if a field has been set.
func (o *SACEvent) HasImmediateFlag() bool {
	if o != nil && !IsNil(o.ImmediateFlag) {
		return true
	}

	return false
}

// SetImmediateFlag gets a reference to the given bool and assigns it to the ImmediateFlag field.
func (o *SACEvent) SetImmediateFlag(v bool) {
	o.ImmediateFlag = &v
}

func (o SACEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SACEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	if !IsNil(o.EventTrigger) {
		toSerialize["eventTrigger"] = o.EventTrigger
	}
	toSerialize["eventFilter"] = o.EventFilter
	if !IsNil(o.NotificationPeriod) {
		toSerialize["notificationPeriod"] = o.NotificationPeriod
	}
	if !IsNil(o.NotifThreshold) {
		toSerialize["notifThreshold"] = o.NotifThreshold
	}
	if !IsNil(o.ImmediateFlag) {
		toSerialize["immediateFlag"] = o.ImmediateFlag
	}
	return toSerialize, nil
}

type NullableSACEvent struct {
	value *SACEvent
	isSet bool
}

func (v NullableSACEvent) Get() *SACEvent {
	return v.value
}

func (v *NullableSACEvent) Set(val *SACEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEvent(val *SACEvent) *NullableSACEvent {
	return &NullableSACEvent{value: val, isSet: true}
}

func (v NullableSACEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


