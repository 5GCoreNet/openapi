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

// checks if the EasDeployInfoNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDeployInfoNotif{}

// EasDeployInfoNotif Represents notifications on EAS Deployment Information changes event(s) that occurred for an  Individual EAS Deployment Event Subscription resource.
type EasDeployInfoNotif struct {
	EasDepNotifs []EasDepNotification `json:"easDepNotifs"`
	NotifId      string               `json:"notifId"`
}

// NewEasDeployInfoNotif instantiates a new EasDeployInfoNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDeployInfoNotif(easDepNotifs []EasDepNotification, notifId string) *EasDeployInfoNotif {
	this := EasDeployInfoNotif{}
	this.EasDepNotifs = easDepNotifs
	this.NotifId = notifId
	return &this
}

// NewEasDeployInfoNotifWithDefaults instantiates a new EasDeployInfoNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDeployInfoNotifWithDefaults() *EasDeployInfoNotif {
	this := EasDeployInfoNotif{}
	return &this
}

// GetEasDepNotifs returns the EasDepNotifs field value
func (o *EasDeployInfoNotif) GetEasDepNotifs() []EasDepNotification {
	if o == nil {
		var ret []EasDepNotification
		return ret
	}

	return o.EasDepNotifs
}

// GetEasDepNotifsOk returns a tuple with the EasDepNotifs field value
// and a boolean to check if the value has been set.
func (o *EasDeployInfoNotif) GetEasDepNotifsOk() ([]EasDepNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.EasDepNotifs, true
}

// SetEasDepNotifs sets field value
func (o *EasDeployInfoNotif) SetEasDepNotifs(v []EasDepNotification) {
	o.EasDepNotifs = v
}

// GetNotifId returns the NotifId field value
func (o *EasDeployInfoNotif) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *EasDeployInfoNotif) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *EasDeployInfoNotif) SetNotifId(v string) {
	o.NotifId = v
}

func (o EasDeployInfoNotif) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDeployInfoNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["easDepNotifs"] = o.EasDepNotifs
	toSerialize["notifId"] = o.NotifId
	return toSerialize, nil
}

type NullableEasDeployInfoNotif struct {
	value *EasDeployInfoNotif
	isSet bool
}

func (v NullableEasDeployInfoNotif) Get() *EasDeployInfoNotif {
	return v.value
}

func (v *NullableEasDeployInfoNotif) Set(val *EasDeployInfoNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDeployInfoNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDeployInfoNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDeployInfoNotif(val *EasDeployInfoNotif) *NullableEasDeployInfoNotif {
	return &NullableEasDeployInfoNotif{value: val, isSet: true}
}

func (v NullableEasDeployInfoNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDeployInfoNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
