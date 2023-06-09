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

// checks if the NotifyPotentialFaultyAlarmList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyPotentialFaultyAlarmList{}

// NotifyPotentialFaultyAlarmList struct for NotifyPotentialFaultyAlarmList
type NotifyPotentialFaultyAlarmList struct {
	NotificationHeader
	Reason string `json:"reason"`
}

// NewNotifyPotentialFaultyAlarmList instantiates a new NotifyPotentialFaultyAlarmList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyPotentialFaultyAlarmList(reason string, href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyPotentialFaultyAlarmList {
	this := NotifyPotentialFaultyAlarmList{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.Reason = reason
	return &this
}

// NewNotifyPotentialFaultyAlarmListWithDefaults instantiates a new NotifyPotentialFaultyAlarmList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyPotentialFaultyAlarmListWithDefaults() *NotifyPotentialFaultyAlarmList {
	this := NotifyPotentialFaultyAlarmList{}
	return &this
}

// GetReason returns the Reason field value
func (o *NotifyPotentialFaultyAlarmList) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *NotifyPotentialFaultyAlarmList) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *NotifyPotentialFaultyAlarmList) SetReason(v string) {
	o.Reason = v
}

func (o NotifyPotentialFaultyAlarmList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyPotentialFaultyAlarmList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationHeader, errNotificationHeader := json.Marshal(o.NotificationHeader)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	errNotificationHeader = json.Unmarshal([]byte(serializedNotificationHeader), &toSerialize)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

type NullableNotifyPotentialFaultyAlarmList struct {
	value *NotifyPotentialFaultyAlarmList
	isSet bool
}

func (v NullableNotifyPotentialFaultyAlarmList) Get() *NotifyPotentialFaultyAlarmList {
	return v.value
}

func (v *NullableNotifyPotentialFaultyAlarmList) Set(val *NotifyPotentialFaultyAlarmList) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyPotentialFaultyAlarmList) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyPotentialFaultyAlarmList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyPotentialFaultyAlarmList(val *NotifyPotentialFaultyAlarmList) *NullableNotifyPotentialFaultyAlarmList {
	return &NullableNotifyPotentialFaultyAlarmList{value: val, isSet: true}
}

func (v NullableNotifyPotentialFaultyAlarmList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyPotentialFaultyAlarmList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
