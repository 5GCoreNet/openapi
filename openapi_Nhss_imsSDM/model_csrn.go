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

// checks if the Csrn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Csrn{}

// Csrn CS domain routeing number
type Csrn struct {
	// String containing an additional or basic MSISDN
	Csrn string `json:"csrn"`
}

// NewCsrn instantiates a new Csrn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrn(csrn string) *Csrn {
	this := Csrn{}
	this.Csrn = csrn
	return &this
}

// NewCsrnWithDefaults instantiates a new Csrn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrnWithDefaults() *Csrn {
	this := Csrn{}
	return &this
}

// GetCsrn returns the Csrn field value
func (o *Csrn) GetCsrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Csrn
}

// GetCsrnOk returns a tuple with the Csrn field value
// and a boolean to check if the value has been set.
func (o *Csrn) GetCsrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Csrn, true
}

// SetCsrn sets field value
func (o *Csrn) SetCsrn(v string) {
	o.Csrn = v
}

func (o Csrn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Csrn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["csrn"] = o.Csrn
	return toSerialize, nil
}

type NullableCsrn struct {
	value *Csrn
	isSet bool
}

func (v NullableCsrn) Get() *Csrn {
	return v.value
}

func (v *NullableCsrn) Set(val *Csrn) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrn) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrn(val *Csrn) *NullableCsrn {
	return &NullableCsrn{value: val, isSet: true}
}

func (v NullableCsrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

