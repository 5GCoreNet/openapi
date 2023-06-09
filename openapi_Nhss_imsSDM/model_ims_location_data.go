/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the ImsLocationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImsLocationData{}

// ImsLocationData IMS Location Data (S-CSCF name)
type ImsLocationData struct {
	ScscfName string `json:"scscfName"`
}

// NewImsLocationData instantiates a new ImsLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImsLocationData(scscfName string) *ImsLocationData {
	this := ImsLocationData{}
	this.ScscfName = scscfName
	return &this
}

// NewImsLocationDataWithDefaults instantiates a new ImsLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImsLocationDataWithDefaults() *ImsLocationData {
	this := ImsLocationData{}
	return &this
}

// GetScscfName returns the ScscfName field value
func (o *ImsLocationData) GetScscfName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScscfName
}

// GetScscfNameOk returns a tuple with the ScscfName field value
// and a boolean to check if the value has been set.
func (o *ImsLocationData) GetScscfNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScscfName, true
}

// SetScscfName sets field value
func (o *ImsLocationData) SetScscfName(v string) {
	o.ScscfName = v
}

func (o ImsLocationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImsLocationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scscfName"] = o.ScscfName
	return toSerialize, nil
}

type NullableImsLocationData struct {
	value *ImsLocationData
	isSet bool
}

func (v NullableImsLocationData) Get() *ImsLocationData {
	return v.value
}

func (v *NullableImsLocationData) Set(val *ImsLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableImsLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableImsLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsLocationData(val *ImsLocationData) *NullableImsLocationData {
	return &NullableImsLocationData{value: val, isSet: true}
}

func (v NullableImsLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
