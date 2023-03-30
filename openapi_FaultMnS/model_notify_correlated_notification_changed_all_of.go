/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// checks if the NotifyCorrelatedNotificationChangedAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyCorrelatedNotificationChangedAllOf{}

// NotifyCorrelatedNotificationChangedAllOf struct for NotifyCorrelatedNotificationChangedAllOf
type NotifyCorrelatedNotificationChangedAllOf struct {
	AlarmId string `json:"alarmId"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications"`
	RootCauseIndicator *bool `json:"rootCauseIndicator,omitempty"`
}

// NewNotifyCorrelatedNotificationChangedAllOf instantiates a new NotifyCorrelatedNotificationChangedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyCorrelatedNotificationChangedAllOf(alarmId string, correlatedNotifications []CorrelatedNotification) *NotifyCorrelatedNotificationChangedAllOf {
	this := NotifyCorrelatedNotificationChangedAllOf{}
	this.AlarmId = alarmId
	this.CorrelatedNotifications = correlatedNotifications
	return &this
}

// NewNotifyCorrelatedNotificationChangedAllOfWithDefaults instantiates a new NotifyCorrelatedNotificationChangedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyCorrelatedNotificationChangedAllOfWithDefaults() *NotifyCorrelatedNotificationChangedAllOf {
	this := NotifyCorrelatedNotificationChangedAllOf{}
	return &this
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyCorrelatedNotificationChangedAllOf) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyCorrelatedNotificationChangedAllOf) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyCorrelatedNotificationChangedAllOf) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value
func (o *NotifyCorrelatedNotificationChangedAllOf) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil {
		var ret []CorrelatedNotification
		return ret
	}

	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value
// and a boolean to check if the value has been set.
func (o *NotifyCorrelatedNotificationChangedAllOf) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// SetCorrelatedNotifications sets field value
func (o *NotifyCorrelatedNotificationChangedAllOf) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetRootCauseIndicator returns the RootCauseIndicator field value if set, zero value otherwise.
func (o *NotifyCorrelatedNotificationChangedAllOf) GetRootCauseIndicator() bool {
	if o == nil || IsNil(o.RootCauseIndicator) {
		var ret bool
		return ret
	}
	return *o.RootCauseIndicator
}

// GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyCorrelatedNotificationChangedAllOf) GetRootCauseIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.RootCauseIndicator) {
		return nil, false
	}
	return o.RootCauseIndicator, true
}

// HasRootCauseIndicator returns a boolean if a field has been set.
func (o *NotifyCorrelatedNotificationChangedAllOf) HasRootCauseIndicator() bool {
	if o != nil && !IsNil(o.RootCauseIndicator) {
		return true
	}

	return false
}

// SetRootCauseIndicator gets a reference to the given bool and assigns it to the RootCauseIndicator field.
func (o *NotifyCorrelatedNotificationChangedAllOf) SetRootCauseIndicator(v bool) {
	o.RootCauseIndicator = &v
}

func (o NotifyCorrelatedNotificationChangedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyCorrelatedNotificationChangedAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	if !IsNil(o.RootCauseIndicator) {
		toSerialize["rootCauseIndicator"] = o.RootCauseIndicator
	}
	return toSerialize, nil
}

type NullableNotifyCorrelatedNotificationChangedAllOf struct {
	value *NotifyCorrelatedNotificationChangedAllOf
	isSet bool
}

func (v NullableNotifyCorrelatedNotificationChangedAllOf) Get() *NotifyCorrelatedNotificationChangedAllOf {
	return v.value
}

func (v *NullableNotifyCorrelatedNotificationChangedAllOf) Set(val *NotifyCorrelatedNotificationChangedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyCorrelatedNotificationChangedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyCorrelatedNotificationChangedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyCorrelatedNotificationChangedAllOf(val *NotifyCorrelatedNotificationChangedAllOf) *NullableNotifyCorrelatedNotificationChangedAllOf {
	return &NullableNotifyCorrelatedNotificationChangedAllOf{value: val, isSet: true}
}

func (v NullableNotifyCorrelatedNotificationChangedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyCorrelatedNotificationChangedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

