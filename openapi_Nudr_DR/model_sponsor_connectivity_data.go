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

// checks if the SponsorConnectivityData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsorConnectivityData{}

// SponsorConnectivityData Contains the sponsored data connectivity related information for a sponsor identifier. 
type SponsorConnectivityData struct {
	AspIds []string `json:"aspIds"`
}

// NewSponsorConnectivityData instantiates a new SponsorConnectivityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsorConnectivityData(aspIds []string) *SponsorConnectivityData {
	this := SponsorConnectivityData{}
	this.AspIds = aspIds
	return &this
}

// NewSponsorConnectivityDataWithDefaults instantiates a new SponsorConnectivityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsorConnectivityDataWithDefaults() *SponsorConnectivityData {
	this := SponsorConnectivityData{}
	return &this
}

// GetAspIds returns the AspIds field value
func (o *SponsorConnectivityData) GetAspIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AspIds
}

// GetAspIdsOk returns a tuple with the AspIds field value
// and a boolean to check if the value has been set.
func (o *SponsorConnectivityData) GetAspIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AspIds, true
}

// SetAspIds sets field value
func (o *SponsorConnectivityData) SetAspIds(v []string) {
	o.AspIds = v
}

func (o SponsorConnectivityData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SponsorConnectivityData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aspIds"] = o.AspIds
	return toSerialize, nil
}

type NullableSponsorConnectivityData struct {
	value *SponsorConnectivityData
	isSet bool
}

func (v NullableSponsorConnectivityData) Get() *SponsorConnectivityData {
	return v.value
}

func (v *NullableSponsorConnectivityData) Set(val *SponsorConnectivityData) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsorConnectivityData) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsorConnectivityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsorConnectivityData(val *SponsorConnectivityData) *NullableSponsorConnectivityData {
	return &NullableSponsorConnectivityData{value: val, isSet: true}
}

func (v NullableSponsorConnectivityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSponsorConnectivityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


