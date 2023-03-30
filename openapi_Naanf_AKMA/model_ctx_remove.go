/*
3gpp-akma

API for Naanf_AKMA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naanf_AKMA

import (
	"encoding/json"
)

// checks if the CtxRemove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CtxRemove{}

// CtxRemove Parameters to request to delete the AKMA context for the UE, the \"supi\" attribute shall be  included. 
type CtxRemove struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
}

// NewCtxRemove instantiates a new CtxRemove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCtxRemove() *CtxRemove {
	this := CtxRemove{}
	return &this
}

// NewCtxRemoveWithDefaults instantiates a new CtxRemove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCtxRemoveWithDefaults() *CtxRemove {
	this := CtxRemove{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *CtxRemove) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CtxRemove) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *CtxRemove) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *CtxRemove) SetSupi(v string) {
	o.Supi = &v
}

func (o CtxRemove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CtxRemove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	return toSerialize, nil
}

type NullableCtxRemove struct {
	value *CtxRemove
	isSet bool
}

func (v NullableCtxRemove) Get() *CtxRemove {
	return v.value
}

func (v *NullableCtxRemove) Set(val *CtxRemove) {
	v.value = val
	v.isSet = true
}

func (v NullableCtxRemove) IsSet() bool {
	return v.isSet
}

func (v *NullableCtxRemove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCtxRemove(val *CtxRemove) *NullableCtxRemove {
	return &NullableCtxRemove{value: val, isSet: true}
}

func (v NullableCtxRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCtxRemove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


