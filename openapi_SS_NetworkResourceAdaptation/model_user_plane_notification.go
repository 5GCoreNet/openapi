/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// checks if the UserPlaneNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPlaneNotification{}

// UserPlaneNotification Represents a notification on User Plane events.
type UserPlaneNotification struct {
	// String providing an URI formatted according to RFC 3986.
	NotifId     string                 `json:"notifId"`
	EventNotifs []NrmEventNotification `json:"eventNotifs"`
}

// NewUserPlaneNotification instantiates a new UserPlaneNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPlaneNotification(notifId string, eventNotifs []NrmEventNotification) *UserPlaneNotification {
	this := UserPlaneNotification{}
	this.NotifId = notifId
	this.EventNotifs = eventNotifs
	return &this
}

// NewUserPlaneNotificationWithDefaults instantiates a new UserPlaneNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPlaneNotificationWithDefaults() *UserPlaneNotification {
	this := UserPlaneNotification{}
	return &this
}

// GetNotifId returns the NotifId field value
func (o *UserPlaneNotification) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *UserPlaneNotification) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *UserPlaneNotification) SetNotifId(v string) {
	o.NotifId = v
}

// GetEventNotifs returns the EventNotifs field value
func (o *UserPlaneNotification) GetEventNotifs() []NrmEventNotification {
	if o == nil {
		var ret []NrmEventNotification
		return ret
	}

	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value
// and a boolean to check if the value has been set.
func (o *UserPlaneNotification) GetEventNotifsOk() ([]NrmEventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventNotifs, true
}

// SetEventNotifs sets field value
func (o *UserPlaneNotification) SetEventNotifs(v []NrmEventNotification) {
	o.EventNotifs = v
}

func (o UserPlaneNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPlaneNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifId"] = o.NotifId
	toSerialize["eventNotifs"] = o.EventNotifs
	return toSerialize, nil
}

type NullableUserPlaneNotification struct {
	value *UserPlaneNotification
	isSet bool
}

func (v NullableUserPlaneNotification) Get() *UserPlaneNotification {
	return v.value
}

func (v *NullableUserPlaneNotification) Set(val *UserPlaneNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlaneNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlaneNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlaneNotification(val *UserPlaneNotification) *NullableUserPlaneNotification {
	return &NullableUserPlaneNotification{value: val, isSet: true}
}

func (v NullableUserPlaneNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlaneNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
