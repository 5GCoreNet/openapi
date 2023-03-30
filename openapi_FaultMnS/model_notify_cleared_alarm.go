/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"time"
)

// checks if the NotifyClearedAlarm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyClearedAlarm{}

// NotifyClearedAlarm struct for NotifyClearedAlarm
type NotifyClearedAlarm struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	AlarmId string `json:"alarmId"`
	AlarmType AlarmType `json:"alarmType"`
	ProbableCause ProbableCause `json:"probableCause"`
	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	ClearUserId *string `json:"clearUserId,omitempty"`
	ClearSystemId *string `json:"clearSystemId,omitempty"`
}

// NewNotifyClearedAlarm instantiates a new NotifyClearedAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyClearedAlarm(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity) *NotifyClearedAlarm {
	this := NotifyClearedAlarm{}
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

// NewNotifyClearedAlarmWithDefaults instantiates a new NotifyClearedAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyClearedAlarmWithDefaults() *NotifyClearedAlarm {
	this := NotifyClearedAlarm{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyClearedAlarm) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyClearedAlarm) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyClearedAlarm) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyClearedAlarm) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyClearedAlarm) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyClearedAlarm) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyClearedAlarm) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyClearedAlarm) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyClearedAlarm) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetSystemDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyClearedAlarm) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetAlarmId returns the AlarmId field value
func (o *NotifyClearedAlarm) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetAlarmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *NotifyClearedAlarm) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetAlarmType returns the AlarmType field value
func (o *NotifyClearedAlarm) GetAlarmType() AlarmType {
	if o == nil {
		var ret AlarmType
		return ret
	}

	return o.AlarmType
}

// GetAlarmTypeOk returns a tuple with the AlarmType field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetAlarmTypeOk() (*AlarmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlarmType, true
}

// SetAlarmType sets field value
func (o *NotifyClearedAlarm) SetAlarmType(v AlarmType) {
	o.AlarmType = v
}

// GetProbableCause returns the ProbableCause field value
func (o *NotifyClearedAlarm) GetProbableCause() ProbableCause {
	if o == nil {
		var ret ProbableCause
		return ret
	}

	return o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetProbableCauseOk() (*ProbableCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbableCause, true
}

// SetProbableCause sets field value
func (o *NotifyClearedAlarm) SetProbableCause(v ProbableCause) {
	o.ProbableCause = v
}

// GetPerceivedSeverity returns the PerceivedSeverity field value
func (o *NotifyClearedAlarm) GetPerceivedSeverity() PerceivedSeverity {
	if o == nil {
		var ret PerceivedSeverity
		return ret
	}

	return o.PerceivedSeverity
}

// GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field value
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerceivedSeverity, true
}

// SetPerceivedSeverity sets field value
func (o *NotifyClearedAlarm) SetPerceivedSeverity(v PerceivedSeverity) {
	o.PerceivedSeverity = v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyClearedAlarm) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyClearedAlarm) HasCorrelatedNotifications() bool {
	if o != nil && !IsNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyClearedAlarm) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetClearUserId returns the ClearUserId field value if set, zero value otherwise.
func (o *NotifyClearedAlarm) GetClearUserId() string {
	if o == nil || IsNil(o.ClearUserId) {
		var ret string
		return ret
	}
	return *o.ClearUserId
}

// GetClearUserIdOk returns a tuple with the ClearUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetClearUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClearUserId) {
		return nil, false
	}
	return o.ClearUserId, true
}

// HasClearUserId returns a boolean if a field has been set.
func (o *NotifyClearedAlarm) HasClearUserId() bool {
	if o != nil && !IsNil(o.ClearUserId) {
		return true
	}

	return false
}

// SetClearUserId gets a reference to the given string and assigns it to the ClearUserId field.
func (o *NotifyClearedAlarm) SetClearUserId(v string) {
	o.ClearUserId = &v
}

// GetClearSystemId returns the ClearSystemId field value if set, zero value otherwise.
func (o *NotifyClearedAlarm) GetClearSystemId() string {
	if o == nil || IsNil(o.ClearSystemId) {
		var ret string
		return ret
	}
	return *o.ClearSystemId
}

// GetClearSystemIdOk returns a tuple with the ClearSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyClearedAlarm) GetClearSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClearSystemId) {
		return nil, false
	}
	return o.ClearSystemId, true
}

// HasClearSystemId returns a boolean if a field has been set.
func (o *NotifyClearedAlarm) HasClearSystemId() bool {
	if o != nil && !IsNil(o.ClearSystemId) {
		return true
	}

	return false
}

// SetClearSystemId gets a reference to the given string and assigns it to the ClearSystemId field.
func (o *NotifyClearedAlarm) SetClearSystemId(v string) {
	o.ClearSystemId = &v
}

func (o NotifyClearedAlarm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyClearedAlarm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["notificationId"] = o.NotificationId
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["eventTime"] = o.EventTime
	toSerialize["systemDN"] = o.SystemDN
	toSerialize["alarmId"] = o.AlarmId
	toSerialize["alarmType"] = o.AlarmType
	toSerialize["probableCause"] = o.ProbableCause
	toSerialize["perceivedSeverity"] = o.PerceivedSeverity
	if !IsNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !IsNil(o.ClearUserId) {
		toSerialize["clearUserId"] = o.ClearUserId
	}
	if !IsNil(o.ClearSystemId) {
		toSerialize["clearSystemId"] = o.ClearSystemId
	}
	return toSerialize, nil
}

type NullableNotifyClearedAlarm struct {
	value *NotifyClearedAlarm
	isSet bool
}

func (v NullableNotifyClearedAlarm) Get() *NotifyClearedAlarm {
	return v.value
}

func (v *NullableNotifyClearedAlarm) Set(val *NotifyClearedAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyClearedAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyClearedAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyClearedAlarm(val *NotifyClearedAlarm) *NullableNotifyClearedAlarm {
	return &NullableNotifyClearedAlarm{value: val, isSet: true}
}

func (v NullableNotifyClearedAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyClearedAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

