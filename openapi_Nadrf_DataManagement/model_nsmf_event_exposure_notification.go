/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the NsmfEventExposureNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsmfEventExposureNotification{}

// NsmfEventExposureNotification Represents notifications on events that occurred.
type NsmfEventExposureNotification struct {
	// Notification correlation ID
	NotifId string `json:"notifId"`
	// Notifications about Individual Events
	EventNotifs []EventNotification1 `json:"eventNotifs"`
	// String providing an URI formatted according to RFC 3986.
	AckUri *string `json:"ackUri,omitempty"`
}

// NewNsmfEventExposureNotification instantiates a new NsmfEventExposureNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsmfEventExposureNotification(notifId string, eventNotifs []EventNotification1) *NsmfEventExposureNotification {
	this := NsmfEventExposureNotification{}
	this.NotifId = notifId
	this.EventNotifs = eventNotifs
	return &this
}

// NewNsmfEventExposureNotificationWithDefaults instantiates a new NsmfEventExposureNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsmfEventExposureNotificationWithDefaults() *NsmfEventExposureNotification {
	this := NsmfEventExposureNotification{}
	return &this
}

// GetNotifId returns the NotifId field value
func (o *NsmfEventExposureNotification) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *NsmfEventExposureNotification) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *NsmfEventExposureNotification) SetNotifId(v string) {
	o.NotifId = v
}

// GetEventNotifs returns the EventNotifs field value
func (o *NsmfEventExposureNotification) GetEventNotifs() []EventNotification1 {
	if o == nil {
		var ret []EventNotification1
		return ret
	}

	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value
// and a boolean to check if the value has been set.
func (o *NsmfEventExposureNotification) GetEventNotifsOk() ([]EventNotification1, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventNotifs, true
}

// SetEventNotifs sets field value
func (o *NsmfEventExposureNotification) SetEventNotifs(v []EventNotification1) {
	o.EventNotifs = v
}

// GetAckUri returns the AckUri field value if set, zero value otherwise.
func (o *NsmfEventExposureNotification) GetAckUri() string {
	if o == nil || isNil(o.AckUri) {
		var ret string
		return ret
	}
	return *o.AckUri
}

// GetAckUriOk returns a tuple with the AckUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposureNotification) GetAckUriOk() (*string, bool) {
	if o == nil || isNil(o.AckUri) {
		return nil, false
	}
	return o.AckUri, true
}

// HasAckUri returns a boolean if a field has been set.
func (o *NsmfEventExposureNotification) HasAckUri() bool {
	if o != nil && !isNil(o.AckUri) {
		return true
	}

	return false
}

// SetAckUri gets a reference to the given string and assigns it to the AckUri field.
func (o *NsmfEventExposureNotification) SetAckUri(v string) {
	o.AckUri = &v
}

func (o NsmfEventExposureNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsmfEventExposureNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifId"] = o.NotifId
	toSerialize["eventNotifs"] = o.EventNotifs
	if !isNil(o.AckUri) {
		toSerialize["ackUri"] = o.AckUri
	}
	return toSerialize, nil
}

type NullableNsmfEventExposureNotification struct {
	value *NsmfEventExposureNotification
	isSet bool
}

func (v NullableNsmfEventExposureNotification) Get() *NsmfEventExposureNotification {
	return v.value
}

func (v *NullableNsmfEventExposureNotification) Set(val *NsmfEventExposureNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableNsmfEventExposureNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNsmfEventExposureNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsmfEventExposureNotification(val *NsmfEventExposureNotification) *NullableNsmfEventExposureNotification {
	return &NullableNsmfEventExposureNotification{value: val, isSet: true}
}

func (v NullableNsmfEventExposureNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsmfEventExposureNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


