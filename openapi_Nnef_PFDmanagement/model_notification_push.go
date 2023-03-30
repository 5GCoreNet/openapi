/*
Nnef_PFDmanagement Service API

Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_PFDmanagement

import (
	"encoding/json"
)

// checks if the NotificationPush type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationPush{}

// NotificationPush Represents the information to be used by the NF service consumer to retrieve the PFDs and/or remove the PFDs of the applicable application identifier(s). 
type NotificationPush struct {
	AppIds []string `json:"appIds"`
	// indicating a time in seconds.
	AllowedDelay *int32 `json:"allowedDelay,omitempty"`
	PfdOp *PfdOperation `json:"pfdOp,omitempty"`
}

// NewNotificationPush instantiates a new NotificationPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationPush(appIds []string) *NotificationPush {
	this := NotificationPush{}
	this.AppIds = appIds
	return &this
}

// NewNotificationPushWithDefaults instantiates a new NotificationPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationPushWithDefaults() *NotificationPush {
	this := NotificationPush{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *NotificationPush) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *NotificationPush) GetAppIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *NotificationPush) SetAppIds(v []string) {
	o.AppIds = v
}

// GetAllowedDelay returns the AllowedDelay field value if set, zero value otherwise.
func (o *NotificationPush) GetAllowedDelay() int32 {
	if o == nil || IsNil(o.AllowedDelay) {
		var ret int32
		return ret
	}
	return *o.AllowedDelay
}

// GetAllowedDelayOk returns a tuple with the AllowedDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPush) GetAllowedDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedDelay) {
		return nil, false
	}
	return o.AllowedDelay, true
}

// HasAllowedDelay returns a boolean if a field has been set.
func (o *NotificationPush) HasAllowedDelay() bool {
	if o != nil && !IsNil(o.AllowedDelay) {
		return true
	}

	return false
}

// SetAllowedDelay gets a reference to the given int32 and assigns it to the AllowedDelay field.
func (o *NotificationPush) SetAllowedDelay(v int32) {
	o.AllowedDelay = &v
}

// GetPfdOp returns the PfdOp field value if set, zero value otherwise.
func (o *NotificationPush) GetPfdOp() PfdOperation {
	if o == nil || IsNil(o.PfdOp) {
		var ret PfdOperation
		return ret
	}
	return *o.PfdOp
}

// GetPfdOpOk returns a tuple with the PfdOp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationPush) GetPfdOpOk() (*PfdOperation, bool) {
	if o == nil || IsNil(o.PfdOp) {
		return nil, false
	}
	return o.PfdOp, true
}

// HasPfdOp returns a boolean if a field has been set.
func (o *NotificationPush) HasPfdOp() bool {
	if o != nil && !IsNil(o.PfdOp) {
		return true
	}

	return false
}

// SetPfdOp gets a reference to the given PfdOperation and assigns it to the PfdOp field.
func (o *NotificationPush) SetPfdOp(v PfdOperation) {
	o.PfdOp = &v
}

func (o NotificationPush) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationPush) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appIds"] = o.AppIds
	if !IsNil(o.AllowedDelay) {
		toSerialize["allowedDelay"] = o.AllowedDelay
	}
	if !IsNil(o.PfdOp) {
		toSerialize["pfdOp"] = o.PfdOp
	}
	return toSerialize, nil
}

type NullableNotificationPush struct {
	value *NotificationPush
	isSet bool
}

func (v NullableNotificationPush) Get() *NotificationPush {
	return v.value
}

func (v *NullableNotificationPush) Set(val *NotificationPush) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationPush) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationPush(val *NotificationPush) *NullableNotificationPush {
	return &NullableNotificationPush{value: val, isSet: true}
}

func (v NullableNotificationPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

