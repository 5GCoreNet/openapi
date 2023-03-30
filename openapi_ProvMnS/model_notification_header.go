/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"time"
)

// checks if the NotificationHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationHeader{}

// NotificationHeader struct for NotificationHeader
type NotificationHeader struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
}

// NewNotificationHeader instantiates a new NotificationHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationHeader(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotificationHeader {
	this := NotificationHeader{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	return &this
}

// NewNotificationHeaderWithDefaults instantiates a new NotificationHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationHeaderWithDefaults() *NotificationHeader {
	this := NotificationHeader{}
	return &this
}

// GetHref returns the Href field value
func (o *NotificationHeader) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotificationHeader) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotificationHeader) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotificationHeader) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotificationHeader) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotificationHeader) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotificationHeader) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotificationHeader) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotificationHeader) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotificationHeader) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotificationHeader) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotificationHeader) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotificationHeader) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotificationHeader) GetSystemDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotificationHeader) SetSystemDN(v string) {
	o.SystemDN = v
}

func (o NotificationHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["notificationId"] = o.NotificationId
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["eventTime"] = o.EventTime
	toSerialize["systemDN"] = o.SystemDN
	return toSerialize, nil
}

type NullableNotificationHeader struct {
	value *NotificationHeader
	isSet bool
}

func (v NullableNotificationHeader) Get() *NotificationHeader {
	return v.value
}

func (v *NullableNotificationHeader) Set(val *NotificationHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationHeader(val *NotificationHeader) *NullableNotificationHeader {
	return &NullableNotificationHeader{value: val, isSet: true}
}

func (v NullableNotificationHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


