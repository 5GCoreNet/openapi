/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the MbsSessionStatusNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionStatusNotif{}

// MbsSessionStatusNotif Represents an MBS Session Status notification.
type MbsSessionStatusNotif struct {
	EventList MbsSessionEventReportList `json:"eventList"`
}

// NewMbsSessionStatusNotif instantiates a new MbsSessionStatusNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionStatusNotif(eventList MbsSessionEventReportList) *MbsSessionStatusNotif {
	this := MbsSessionStatusNotif{}
	this.EventList = eventList
	return &this
}

// NewMbsSessionStatusNotifWithDefaults instantiates a new MbsSessionStatusNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionStatusNotifWithDefaults() *MbsSessionStatusNotif {
	this := MbsSessionStatusNotif{}
	return &this
}

// GetEventList returns the EventList field value
func (o *MbsSessionStatusNotif) GetEventList() MbsSessionEventReportList {
	if o == nil {
		var ret MbsSessionEventReportList
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *MbsSessionStatusNotif) GetEventListOk() (*MbsSessionEventReportList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventList, true
}

// SetEventList sets field value
func (o *MbsSessionStatusNotif) SetEventList(v MbsSessionEventReportList) {
	o.EventList = v
}

func (o MbsSessionStatusNotif) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionStatusNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventList"] = o.EventList
	return toSerialize, nil
}

type NullableMbsSessionStatusNotif struct {
	value *MbsSessionStatusNotif
	isSet bool
}

func (v NullableMbsSessionStatusNotif) Get() *MbsSessionStatusNotif {
	return v.value
}

func (v *NullableMbsSessionStatusNotif) Set(val *MbsSessionStatusNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionStatusNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionStatusNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionStatusNotif(val *MbsSessionStatusNotif) *NullableMbsSessionStatusNotif {
	return &NullableMbsSessionStatusNotif{value: val, isSet: true}
}

func (v NullableMbsSessionStatusNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionStatusNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
