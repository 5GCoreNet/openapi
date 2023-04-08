/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the TrafficInfluDataNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficInfluDataNotif{}

// TrafficInfluDataNotif Represents traffic influence data for notification.
type TrafficInfluDataNotif struct {
	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`
	TrafficInfluData *TrafficInfluData `json:"trafficInfluData,omitempty"`
}

// NewTrafficInfluDataNotif instantiates a new TrafficInfluDataNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficInfluDataNotif(resUri string) *TrafficInfluDataNotif {
	this := TrafficInfluDataNotif{}
	this.ResUri = resUri
	return &this
}

// NewTrafficInfluDataNotifWithDefaults instantiates a new TrafficInfluDataNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficInfluDataNotifWithDefaults() *TrafficInfluDataNotif {
	this := TrafficInfluDataNotif{}
	return &this
}

// GetResUri returns the ResUri field value
func (o *TrafficInfluDataNotif) GetResUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResUri
}

// GetResUriOk returns a tuple with the ResUri field value
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataNotif) GetResUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResUri, true
}

// SetResUri sets field value
func (o *TrafficInfluDataNotif) SetResUri(v string) {
	o.ResUri = v
}

// GetTrafficInfluData returns the TrafficInfluData field value if set, zero value otherwise.
func (o *TrafficInfluDataNotif) GetTrafficInfluData() TrafficInfluData {
	if o == nil || isNil(o.TrafficInfluData) {
		var ret TrafficInfluData
		return ret
	}
	return *o.TrafficInfluData
}

// GetTrafficInfluDataOk returns a tuple with the TrafficInfluData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataNotif) GetTrafficInfluDataOk() (*TrafficInfluData, bool) {
	if o == nil || isNil(o.TrafficInfluData) {
		return nil, false
	}
	return o.TrafficInfluData, true
}

// HasTrafficInfluData returns a boolean if a field has been set.
func (o *TrafficInfluDataNotif) HasTrafficInfluData() bool {
	if o != nil && !isNil(o.TrafficInfluData) {
		return true
	}

	return false
}

// SetTrafficInfluData gets a reference to the given TrafficInfluData and assigns it to the TrafficInfluData field.
func (o *TrafficInfluDataNotif) SetTrafficInfluData(v TrafficInfluData) {
	o.TrafficInfluData = &v
}

func (o TrafficInfluDataNotif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficInfluDataNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resUri"] = o.ResUri
	if !isNil(o.TrafficInfluData) {
		toSerialize["trafficInfluData"] = o.TrafficInfluData
	}
	return toSerialize, nil
}

type NullableTrafficInfluDataNotif struct {
	value *TrafficInfluDataNotif
	isSet bool
}

func (v NullableTrafficInfluDataNotif) Get() *TrafficInfluDataNotif {
	return v.value
}

func (v *NullableTrafficInfluDataNotif) Set(val *TrafficInfluDataNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficInfluDataNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficInfluDataNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficInfluDataNotif(val *TrafficInfluDataNotif) *NullableTrafficInfluDataNotif {
	return &NullableTrafficInfluDataNotif{value: val, isSet: true}
}

func (v NullableTrafficInfluDataNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficInfluDataNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


