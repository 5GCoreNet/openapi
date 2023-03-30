/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// checks if the Tmgi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tmgi{}

// Tmgi Temporary Mobile Group Identity
type Tmgi struct {
	// MBS Service ID
	MbsServiceId string `json:"mbsServiceId"`
	PlmnId PlmnId `json:"plmnId"`
}

// NewTmgi instantiates a new Tmgi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgi(mbsServiceId string, plmnId PlmnId) *Tmgi {
	this := Tmgi{}
	this.MbsServiceId = mbsServiceId
	this.PlmnId = plmnId
	return &this
}

// NewTmgiWithDefaults instantiates a new Tmgi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiWithDefaults() *Tmgi {
	this := Tmgi{}
	return &this
}

// GetMbsServiceId returns the MbsServiceId field value
func (o *Tmgi) GetMbsServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsServiceId
}

// GetMbsServiceIdOk returns a tuple with the MbsServiceId field value
// and a boolean to check if the value has been set.
func (o *Tmgi) GetMbsServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsServiceId, true
}

// SetMbsServiceId sets field value
func (o *Tmgi) SetMbsServiceId(v string) {
	o.MbsServiceId = v
}

// GetPlmnId returns the PlmnId field value
func (o *Tmgi) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Tmgi) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Tmgi) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

func (o Tmgi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tmgi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsServiceId"] = o.MbsServiceId
	toSerialize["plmnId"] = o.PlmnId
	return toSerialize, nil
}

type NullableTmgi struct {
	value *Tmgi
	isSet bool
}

func (v NullableTmgi) Get() *Tmgi {
	return v.value
}

func (v *NullableTmgi) Set(val *Tmgi) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgi) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgi(val *Tmgi) *NullableTmgi {
	return &NullableTmgi{value: val, isSet: true}
}

func (v NullableTmgi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


