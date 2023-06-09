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

// checks if the NotifyAlarmListRebuilt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyAlarmListRebuilt{}

// NotifyAlarmListRebuilt struct for NotifyAlarmListRebuilt
type NotifyAlarmListRebuilt struct {
	NotificationHeader
	Reason                        string                         `json:"reason"`
	AlarmListAlignmentRequirement *AlarmListAlignmentRequirement `json:"alarmListAlignmentRequirement,omitempty"`
}

// NewNotifyAlarmListRebuilt instantiates a new NotifyAlarmListRebuilt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyAlarmListRebuilt(reason string, href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyAlarmListRebuilt {
	this := NotifyAlarmListRebuilt{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.Reason = reason
	return &this
}

// NewNotifyAlarmListRebuiltWithDefaults instantiates a new NotifyAlarmListRebuilt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyAlarmListRebuiltWithDefaults() *NotifyAlarmListRebuilt {
	this := NotifyAlarmListRebuilt{}
	return &this
}

// GetReason returns the Reason field value
func (o *NotifyAlarmListRebuilt) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *NotifyAlarmListRebuilt) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *NotifyAlarmListRebuilt) SetReason(v string) {
	o.Reason = v
}

// GetAlarmListAlignmentRequirement returns the AlarmListAlignmentRequirement field value if set, zero value otherwise.
func (o *NotifyAlarmListRebuilt) GetAlarmListAlignmentRequirement() AlarmListAlignmentRequirement {
	if o == nil || IsNil(o.AlarmListAlignmentRequirement) {
		var ret AlarmListAlignmentRequirement
		return ret
	}
	return *o.AlarmListAlignmentRequirement
}

// GetAlarmListAlignmentRequirementOk returns a tuple with the AlarmListAlignmentRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyAlarmListRebuilt) GetAlarmListAlignmentRequirementOk() (*AlarmListAlignmentRequirement, bool) {
	if o == nil || IsNil(o.AlarmListAlignmentRequirement) {
		return nil, false
	}
	return o.AlarmListAlignmentRequirement, true
}

// HasAlarmListAlignmentRequirement returns a boolean if a field has been set.
func (o *NotifyAlarmListRebuilt) HasAlarmListAlignmentRequirement() bool {
	if o != nil && !IsNil(o.AlarmListAlignmentRequirement) {
		return true
	}

	return false
}

// SetAlarmListAlignmentRequirement gets a reference to the given AlarmListAlignmentRequirement and assigns it to the AlarmListAlignmentRequirement field.
func (o *NotifyAlarmListRebuilt) SetAlarmListAlignmentRequirement(v AlarmListAlignmentRequirement) {
	o.AlarmListAlignmentRequirement = &v
}

func (o NotifyAlarmListRebuilt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyAlarmListRebuilt) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AlarmListAlignmentRequirement) {
		toSerialize["alarmListAlignmentRequirement"] = o.AlarmListAlignmentRequirement
	}
	return toSerialize, nil
}

type NullableNotifyAlarmListRebuilt struct {
	value *NotifyAlarmListRebuilt
	isSet bool
}

func (v NullableNotifyAlarmListRebuilt) Get() *NotifyAlarmListRebuilt {
	return v.value
}

func (v *NullableNotifyAlarmListRebuilt) Set(val *NotifyAlarmListRebuilt) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyAlarmListRebuilt) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyAlarmListRebuilt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyAlarmListRebuilt(val *NotifyAlarmListRebuilt) *NullableNotifyAlarmListRebuilt {
	return &NullableNotifyAlarmListRebuilt{value: val, isSet: true}
}

func (v NullableNotifyAlarmListRebuilt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyAlarmListRebuilt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
