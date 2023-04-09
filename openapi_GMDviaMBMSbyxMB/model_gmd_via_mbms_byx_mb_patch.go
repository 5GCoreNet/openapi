/*
GMDviaMBMSbyxMB

API for Group Message Delivery via MBMS by xMB   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GMDviaMBMSbyxMB

import (
	"encoding/json"
	"time"
)

// checks if the GMDViaMBMSByxMBPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GMDViaMBMSByxMBPatch{}

// GMDViaMBMSByxMBPatch Represents a modification request of a group message delivery via MBMS by xMB.
type GMDViaMBMSByxMBPatch struct {
	MbmsLocArea *MbmsLocArea `json:"mbmsLocArea,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStartTime *time.Time `json:"messageDeliveryStartTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	MessageDeliveryStopTime *time.Time `json:"messageDeliveryStopTime,omitempty"`
	// String with format \"byte\" as defined in OpenAPI Specification, i.e, base64-encoded characters.
	GroupMessagePayload *string `json:"groupMessagePayload,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
}

// NewGMDViaMBMSByxMBPatch instantiates a new GMDViaMBMSByxMBPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGMDViaMBMSByxMBPatch() *GMDViaMBMSByxMBPatch {
	this := GMDViaMBMSByxMBPatch{}
	return &this
}

// NewGMDViaMBMSByxMBPatchWithDefaults instantiates a new GMDViaMBMSByxMBPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGMDViaMBMSByxMBPatchWithDefaults() *GMDViaMBMSByxMBPatch {
	this := GMDViaMBMSByxMBPatch{}
	return &this
}

// GetMbmsLocArea returns the MbmsLocArea field value if set, zero value otherwise.
func (o *GMDViaMBMSByxMBPatch) GetMbmsLocArea() MbmsLocArea {
	if o == nil || IsNil(o.MbmsLocArea) {
		var ret MbmsLocArea
		return ret
	}
	return *o.MbmsLocArea
}

// GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByxMBPatch) GetMbmsLocAreaOk() (*MbmsLocArea, bool) {
	if o == nil || IsNil(o.MbmsLocArea) {
		return nil, false
	}
	return o.MbmsLocArea, true
}

// HasMbmsLocArea returns a boolean if a field has been set.
func (o *GMDViaMBMSByxMBPatch) HasMbmsLocArea() bool {
	if o != nil && !IsNil(o.MbmsLocArea) {
		return true
	}

	return false
}

// SetMbmsLocArea gets a reference to the given MbmsLocArea and assigns it to the MbmsLocArea field.
func (o *GMDViaMBMSByxMBPatch) SetMbmsLocArea(v MbmsLocArea) {
	o.MbmsLocArea = &v
}

// GetMessageDeliveryStartTime returns the MessageDeliveryStartTime field value if set, zero value otherwise.
func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStartTime() time.Time {
	if o == nil || IsNil(o.MessageDeliveryStartTime) {
		var ret time.Time
		return ret
	}
	return *o.MessageDeliveryStartTime
}

// GetMessageDeliveryStartTimeOk returns a tuple with the MessageDeliveryStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MessageDeliveryStartTime) {
		return nil, false
	}
	return o.MessageDeliveryStartTime, true
}

// HasMessageDeliveryStartTime returns a boolean if a field has been set.
func (o *GMDViaMBMSByxMBPatch) HasMessageDeliveryStartTime() bool {
	if o != nil && !IsNil(o.MessageDeliveryStartTime) {
		return true
	}

	return false
}

// SetMessageDeliveryStartTime gets a reference to the given time.Time and assigns it to the MessageDeliveryStartTime field.
func (o *GMDViaMBMSByxMBPatch) SetMessageDeliveryStartTime(v time.Time) {
	o.MessageDeliveryStartTime = &v
}

// GetMessageDeliveryStopTime returns the MessageDeliveryStopTime field value if set, zero value otherwise.
func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStopTime() time.Time {
	if o == nil || IsNil(o.MessageDeliveryStopTime) {
		var ret time.Time
		return ret
	}
	return *o.MessageDeliveryStopTime
}

// GetMessageDeliveryStopTimeOk returns a tuple with the MessageDeliveryStopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStopTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MessageDeliveryStopTime) {
		return nil, false
	}
	return o.MessageDeliveryStopTime, true
}

// HasMessageDeliveryStopTime returns a boolean if a field has been set.
func (o *GMDViaMBMSByxMBPatch) HasMessageDeliveryStopTime() bool {
	if o != nil && !IsNil(o.MessageDeliveryStopTime) {
		return true
	}

	return false
}

// SetMessageDeliveryStopTime gets a reference to the given time.Time and assigns it to the MessageDeliveryStopTime field.
func (o *GMDViaMBMSByxMBPatch) SetMessageDeliveryStopTime(v time.Time) {
	o.MessageDeliveryStopTime = &v
}

// GetGroupMessagePayload returns the GroupMessagePayload field value if set, zero value otherwise.
func (o *GMDViaMBMSByxMBPatch) GetGroupMessagePayload() string {
	if o == nil || IsNil(o.GroupMessagePayload) {
		var ret string
		return ret
	}
	return *o.GroupMessagePayload
}

// GetGroupMessagePayloadOk returns a tuple with the GroupMessagePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByxMBPatch) GetGroupMessagePayloadOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMessagePayload) {
		return nil, false
	}
	return o.GroupMessagePayload, true
}

// HasGroupMessagePayload returns a boolean if a field has been set.
func (o *GMDViaMBMSByxMBPatch) HasGroupMessagePayload() bool {
	if o != nil && !IsNil(o.GroupMessagePayload) {
		return true
	}

	return false
}

// SetGroupMessagePayload gets a reference to the given string and assigns it to the GroupMessagePayload field.
func (o *GMDViaMBMSByxMBPatch) SetGroupMessagePayload(v string) {
	o.GroupMessagePayload = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *GMDViaMBMSByxMBPatch) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMDViaMBMSByxMBPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *GMDViaMBMSByxMBPatch) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *GMDViaMBMSByxMBPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

func (o GMDViaMBMSByxMBPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GMDViaMBMSByxMBPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbmsLocArea) {
		toSerialize["mbmsLocArea"] = o.MbmsLocArea
	}
	if !IsNil(o.MessageDeliveryStartTime) {
		toSerialize["messageDeliveryStartTime"] = o.MessageDeliveryStartTime
	}
	if !IsNil(o.MessageDeliveryStopTime) {
		toSerialize["messageDeliveryStopTime"] = o.MessageDeliveryStopTime
	}
	if !IsNil(o.GroupMessagePayload) {
		toSerialize["groupMessagePayload"] = o.GroupMessagePayload
	}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	return toSerialize, nil
}

type NullableGMDViaMBMSByxMBPatch struct {
	value *GMDViaMBMSByxMBPatch
	isSet bool
}

func (v NullableGMDViaMBMSByxMBPatch) Get() *GMDViaMBMSByxMBPatch {
	return v.value
}

func (v *NullableGMDViaMBMSByxMBPatch) Set(val *GMDViaMBMSByxMBPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableGMDViaMBMSByxMBPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableGMDViaMBMSByxMBPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGMDViaMBMSByxMBPatch(val *GMDViaMBMSByxMBPatch) *NullableGMDViaMBMSByxMBPatch {
	return &NullableGMDViaMBMSByxMBPatch{value: val, isSet: true}
}

func (v NullableGMDViaMBMSByxMBPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGMDViaMBMSByxMBPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
