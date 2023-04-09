/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the Snssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Snssai{}

// Snssai struct for Snssai
type Snssai struct {
	Sst *int32  `json:"sst,omitempty"`
	Sd  *string `json:"sd,omitempty"`
}

// NewSnssai instantiates a new Snssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssai() *Snssai {
	this := Snssai{}
	return &this
}

// NewSnssaiWithDefaults instantiates a new Snssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiWithDefaults() *Snssai {
	this := Snssai{}
	return &this
}

// GetSst returns the Sst field value if set, zero value otherwise.
func (o *Snssai) GetSst() int32 {
	if o == nil || IsNil(o.Sst) {
		var ret int32
		return ret
	}
	return *o.Sst
}

// GetSstOk returns a tuple with the Sst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetSstOk() (*int32, bool) {
	if o == nil || IsNil(o.Sst) {
		return nil, false
	}
	return o.Sst, true
}

// HasSst returns a boolean if a field has been set.
func (o *Snssai) HasSst() bool {
	if o != nil && !IsNil(o.Sst) {
		return true
	}

	return false
}

// SetSst gets a reference to the given int32 and assigns it to the Sst field.
func (o *Snssai) SetSst(v int32) {
	o.Sst = &v
}

// GetSd returns the Sd field value if set, zero value otherwise.
func (o *Snssai) GetSd() string {
	if o == nil || IsNil(o.Sd) {
		var ret string
		return ret
	}
	return *o.Sd
}

// GetSdOk returns a tuple with the Sd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Snssai) GetSdOk() (*string, bool) {
	if o == nil || IsNil(o.Sd) {
		return nil, false
	}
	return o.Sd, true
}

// HasSd returns a boolean if a field has been set.
func (o *Snssai) HasSd() bool {
	if o != nil && !IsNil(o.Sd) {
		return true
	}

	return false
}

// SetSd gets a reference to the given string and assigns it to the Sd field.
func (o *Snssai) SetSd(v string) {
	o.Sd = &v
}

func (o Snssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Snssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sst) {
		toSerialize["sst"] = o.Sst
	}
	if !IsNil(o.Sd) {
		toSerialize["sd"] = o.Sd
	}
	return toSerialize, nil
}

type NullableSnssai struct {
	value *Snssai
	isSet bool
}

func (v NullableSnssai) Get() *Snssai {
	return v.value
}

func (v *NullableSnssai) Set(val *Snssai) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssai(val *Snssai) *NullableSnssai {
	return &NullableSnssai{value: val, isSet: true}
}

func (v NullableSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
