/*
Nnef_EASDeployment

NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EASDeployment

import (
	"encoding/json"
)

// checks if the EasDepNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDepNotification{}

// EasDepNotification Represents the EAS Deployment Notifcation.
type EasDepNotification struct {
	EasDepInfo EasDeployInfoData `json:"easDepInfo"`
	EventId    EasEvent          `json:"eventId"`
}

// NewEasDepNotification instantiates a new EasDepNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDepNotification(easDepInfo EasDeployInfoData, eventId EasEvent) *EasDepNotification {
	this := EasDepNotification{}
	this.EasDepInfo = easDepInfo
	this.EventId = eventId
	return &this
}

// NewEasDepNotificationWithDefaults instantiates a new EasDepNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDepNotificationWithDefaults() *EasDepNotification {
	this := EasDepNotification{}
	return &this
}

// GetEasDepInfo returns the EasDepInfo field value
func (o *EasDepNotification) GetEasDepInfo() EasDeployInfoData {
	if o == nil {
		var ret EasDeployInfoData
		return ret
	}

	return o.EasDepInfo
}

// GetEasDepInfoOk returns a tuple with the EasDepInfo field value
// and a boolean to check if the value has been set.
func (o *EasDepNotification) GetEasDepInfoOk() (*EasDeployInfoData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasDepInfo, true
}

// SetEasDepInfo sets field value
func (o *EasDepNotification) SetEasDepInfo(v EasDeployInfoData) {
	o.EasDepInfo = v
}

// GetEventId returns the EventId field value
func (o *EasDepNotification) GetEventId() EasEvent {
	if o == nil {
		var ret EasEvent
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EasDepNotification) GetEventIdOk() (*EasEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EasDepNotification) SetEventId(v EasEvent) {
	o.EventId = v
}

func (o EasDepNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDepNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["easDepInfo"] = o.EasDepInfo
	toSerialize["eventId"] = o.EventId
	return toSerialize, nil
}

type NullableEasDepNotification struct {
	value *EasDepNotification
	isSet bool
}

func (v NullableEasDepNotification) Get() *EasDepNotification {
	return v.value
}

func (v *NullableEasDepNotification) Set(val *EasDepNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDepNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDepNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDepNotification(val *EasDepNotification) *NullableEasDepNotification {
	return &NullableEasDepNotification{value: val, isSet: true}
}

func (v NullableEasDepNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDepNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
