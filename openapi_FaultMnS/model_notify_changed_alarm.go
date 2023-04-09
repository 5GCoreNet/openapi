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

// checks if the NotifyChangedAlarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyChangedAlarm{}

// NotifyChangedAlarm struct for NotifyChangedAlarm
type NotifyChangedAlarm struct {
	NotificationHeader
	AlarmId           string            `json:"alarmId"`
	AlarmType         AlarmType         `json:"alarmType"`
	ProbableCause     ProbableCause     `json:"probableCause"`
	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`
}

// NewNotifyChangedAlarm instantiates a new NotifyChangedAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyChangedAlarm(alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyChangedAlarm {
	this := NotifyChangedAlarm{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AlarmId = alarmId
	this.AlarmType = alarmType
	this.ProbableCause = probableCause
	this.PerceivedSeverity = perceivedSeverity
	return &this
}

// NewNotifyChangedAlarmWithDefaults instantiates a new NotifyChangedAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyChangedAlarmWithDefaults() *NotifyChangedAlarm {
	this := NotifyChangedAlarm{}
	return &this
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyChangedAlarm) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyChangedAlarm) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyChangedAlarm) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyChangedAlarm) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value
func (o *NotifyChangedAlarm) GetProbableCause() ProbableCause {
	if o == nil {
		var ret ProbableCause
		return ret
	}

	return o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbableCause, true
}

// SetProbableCause sets field value
func (o *NotifyChangedAlarm) SetProbableCause(v ProbableCause) {
	o.ProbableCause = v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value
func (o *NotifyChangedAlarm) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil {
		var ret PerceivedSeverity
		return ret
	}

	return o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value
// and a boolean to check if the value has been set.
func (o *NotifyChangedAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerceivedSeverity, true
}

// SetPerceivedSeverity sets field value
func (o *NotifyChangedAlarm) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = v
}

func (o NotifyChangedAlarm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyChangedAlarm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationHeader, errNotificationHeader := json.Marshal(o.NotificationHeader)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	errNotificationHeader = json.Unmarshal([]byte(serializedNotificationHeader), &toSerialize)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	toSerialize["probableCause"] = o.ProbableCause
	toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	return toSerialize, nil
}

type NullableNotifyChangedAlarm struct {
	value *NotifyChangedAlarm
	isSet bool
}

func (v NullableNotifyChangedAlarm) Get() *NotifyChangedAlarm {
	return v.value
}

func (v *NullableNotifyChangedAlarm) Set(val *NotifyChangedAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyChangedAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyChangedAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyChangedAlarm(val *NotifyChangedAlarm) *NullableNotifyChangedAlarm {
	return &NullableNotifyChangedAlarm{value: val, isSet: true}
}

func (v NullableNotifyChangedAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyChangedAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
