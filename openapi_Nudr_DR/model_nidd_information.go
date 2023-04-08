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

// checks if the NiddInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiddInformation{}

// NiddInformation struct for NiddInformation
type NiddInformation struct {
	AfId string `json:"afId"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGroupId *string `json:"extGroupId,omitempty"`
}

// NewNiddInformation instantiates a new NiddInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiddInformation(afId string) *NiddInformation {
	this := NiddInformation{}
	this.AfId = afId
	return &this
}

// NewNiddInformationWithDefaults instantiates a new NiddInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiddInformationWithDefaults() *NiddInformation {
	this := NiddInformation{}
	return &this
}

// GetAfId returns the AfId field value
func (o *NiddInformation) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *NiddInformation) GetAfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *NiddInformation) SetAfId(v string) {
	o.AfId = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *NiddInformation) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddInformation) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *NiddInformation) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *NiddInformation) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExtGroupId returns the ExtGroupId field value if set, zero value otherwise.
func (o *NiddInformation) GetExtGroupId() string {
	if o == nil || isNil(o.ExtGroupId) {
		var ret string
		return ret
	}
	return *o.ExtGroupId
}

// GetExtGroupIdOk returns a tuple with the ExtGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiddInformation) GetExtGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtGroupId) {
		return nil, false
	}
	return o.ExtGroupId, true
}

// HasExtGroupId returns a boolean if a field has been set.
func (o *NiddInformation) HasExtGroupId() bool {
	if o != nil && !isNil(o.ExtGroupId) {
		return true
	}

	return false
}

// SetExtGroupId gets a reference to the given string and assigns it to the ExtGroupId field.
func (o *NiddInformation) SetExtGroupId(v string) {
	o.ExtGroupId = &v
}

func (o NiddInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiddInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afId"] = o.AfId
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.ExtGroupId) {
		toSerialize["extGroupId"] = o.ExtGroupId
	}
	return toSerialize, nil
}

type NullableNiddInformation struct {
	value *NiddInformation
	isSet bool
}

func (v NullableNiddInformation) Get() *NiddInformation {
	return v.value
}

func (v *NullableNiddInformation) Set(val *NiddInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddInformation(val *NiddInformation) *NullableNiddInformation {
	return &NullableNiddInformation{value: val, isSet: true}
}

func (v NullableNiddInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


