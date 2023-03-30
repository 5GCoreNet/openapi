/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"time"
)

// checks if the UeTrajectoryCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeTrajectoryCollection{}

// UeTrajectoryCollection Contains UE trajectory information associated with an application.
type UeTrajectoryCollection struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`
	LocArea LocationArea5G `json:"locArea"`
}

// NewUeTrajectoryCollection instantiates a new UeTrajectoryCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeTrajectoryCollection(ts time.Time, locArea LocationArea5G) *UeTrajectoryCollection {
	this := UeTrajectoryCollection{}
	this.Ts = ts
	this.LocArea = locArea
	return &this
}

// NewUeTrajectoryCollectionWithDefaults instantiates a new UeTrajectoryCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeTrajectoryCollectionWithDefaults() *UeTrajectoryCollection {
	this := UeTrajectoryCollection{}
	return &this
}

// GetTs returns the Ts field value
func (o *UeTrajectoryCollection) GetTs() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ts
}

// GetTsOk returns a tuple with the Ts field value
// and a boolean to check if the value has been set.
func (o *UeTrajectoryCollection) GetTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ts, true
}

// SetTs sets field value
func (o *UeTrajectoryCollection) SetTs(v time.Time) {
	o.Ts = v
}

// GetLocArea returns the LocArea field value
func (o *UeTrajectoryCollection) GetLocArea() LocationArea5G {
	if o == nil {
		var ret LocationArea5G
		return ret
	}

	return o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value
// and a boolean to check if the value has been set.
func (o *UeTrajectoryCollection) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocArea, true
}

// SetLocArea sets field value
func (o *UeTrajectoryCollection) SetLocArea(v LocationArea5G) {
	o.LocArea = v
}

func (o UeTrajectoryCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeTrajectoryCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ts"] = o.Ts
	toSerialize["locArea"] = o.LocArea
	return toSerialize, nil
}

type NullableUeTrajectoryCollection struct {
	value *UeTrajectoryCollection
	isSet bool
}

func (v NullableUeTrajectoryCollection) Get() *UeTrajectoryCollection {
	return v.value
}

func (v *NullableUeTrajectoryCollection) Set(val *UeTrajectoryCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUeTrajectoryCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUeTrajectoryCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeTrajectoryCollection(val *UeTrajectoryCollection) *NullableUeTrajectoryCollection {
	return &NullableUeTrajectoryCollection{value: val, isSet: true}
}

func (v NullableUeTrajectoryCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeTrajectoryCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


