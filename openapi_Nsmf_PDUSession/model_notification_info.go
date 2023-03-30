/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the NotificationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationInfo{}

// NotificationInfo Notification Correlation ID and Notification URI provided by the NF service consumer 
type NotificationInfo struct {
	NotifId string `json:"notifId"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
	UpBufferInd *bool `json:"upBufferInd,omitempty"`
}

// NewNotificationInfo instantiates a new NotificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationInfo(notifId string, notifUri string) *NotificationInfo {
	this := NotificationInfo{}
	this.NotifId = notifId
	this.NotifUri = notifUri
	var upBufferInd bool = false
	this.UpBufferInd = &upBufferInd
	return &this
}

// NewNotificationInfoWithDefaults instantiates a new NotificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationInfoWithDefaults() *NotificationInfo {
	this := NotificationInfo{}
	var upBufferInd bool = false
	this.UpBufferInd = &upBufferInd
	return &this
}

// GetNotifId returns the NotifId field value
func (o *NotificationInfo) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *NotificationInfo) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *NotificationInfo) SetNotifId(v string) {
	o.NotifId = v
}

// GetNotifUri returns the NotifUri field value
func (o *NotificationInfo) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *NotificationInfo) GetNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *NotificationInfo) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetUpBufferInd returns the UpBufferInd field value if set, zero value otherwise.
func (o *NotificationInfo) GetUpBufferInd() bool {
	if o == nil || IsNil(o.UpBufferInd) {
		var ret bool
		return ret
	}
	return *o.UpBufferInd
}

// GetUpBufferIndOk returns a tuple with the UpBufferInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationInfo) GetUpBufferIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UpBufferInd) {
		return nil, false
	}
	return o.UpBufferInd, true
}

// HasUpBufferInd returns a boolean if a field has been set.
func (o *NotificationInfo) HasUpBufferInd() bool {
	if o != nil && !IsNil(o.UpBufferInd) {
		return true
	}

	return false
}

// SetUpBufferInd gets a reference to the given bool and assigns it to the UpBufferInd field.
func (o *NotificationInfo) SetUpBufferInd(v bool) {
	o.UpBufferInd = &v
}

func (o NotificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifId"] = o.NotifId
	toSerialize["notifUri"] = o.NotifUri
	if !IsNil(o.UpBufferInd) {
		toSerialize["upBufferInd"] = o.UpBufferInd
	}
	return toSerialize, nil
}

type NullableNotificationInfo struct {
	value *NotificationInfo
	isSet bool
}

func (v NullableNotificationInfo) Get() *NotificationInfo {
	return v.value
}

func (v *NullableNotificationInfo) Set(val *NotificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationInfo(val *NotificationInfo) *NullableNotificationInfo {
	return &NullableNotificationInfo{value: val, isSet: true}
}

func (v NullableNotificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


