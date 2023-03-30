/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the AfEventExposureNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfEventExposureNotif{}

// AfEventExposureNotif Represents notifications on application event(s) that occurred for an Individual Application Event Subscription resource. 
type AfEventExposureNotif struct {
	NotifId string `json:"notifId"`
	EventNotifs []AfEventNotification `json:"eventNotifs"`
}

// NewAfEventExposureNotif instantiates a new AfEventExposureNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventExposureNotif(notifId string, eventNotifs []AfEventNotification) *AfEventExposureNotif {
	this := AfEventExposureNotif{}
	this.NotifId = notifId
	this.EventNotifs = eventNotifs
	return &this
}

// NewAfEventExposureNotifWithDefaults instantiates a new AfEventExposureNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventExposureNotifWithDefaults() *AfEventExposureNotif {
	this := AfEventExposureNotif{}
	return &this
}

// GetNotifId returns the NotifId field value
func (o *AfEventExposureNotif) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureNotif) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *AfEventExposureNotif) SetNotifId(v string) {
	o.NotifId = v
}

// GetEventNotifs returns the EventNotifs field value
func (o *AfEventExposureNotif) GetEventNotifs() []AfEventNotification {
	if o == nil {
		var ret []AfEventNotification
		return ret
	}

	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value
// and a boolean to check if the value has been set.
func (o *AfEventExposureNotif) GetEventNotifsOk() ([]AfEventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventNotifs, true
}

// SetEventNotifs sets field value
func (o *AfEventExposureNotif) SetEventNotifs(v []AfEventNotification) {
	o.EventNotifs = v
}

func (o AfEventExposureNotif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfEventExposureNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifId"] = o.NotifId
	toSerialize["eventNotifs"] = o.EventNotifs
	return toSerialize, nil
}

type NullableAfEventExposureNotif struct {
	value *AfEventExposureNotif
	isSet bool
}

func (v NullableAfEventExposureNotif) Get() *AfEventExposureNotif {
	return v.value
}

func (v *NullableAfEventExposureNotif) Set(val *AfEventExposureNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventExposureNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventExposureNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventExposureNotif(val *AfEventExposureNotif) *NullableAfEventExposureNotif {
	return &NullableAfEventExposureNotif{value: val, isSet: true}
}

func (v NullableAfEventExposureNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventExposureNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


