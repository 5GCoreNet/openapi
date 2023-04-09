/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the UEIdExt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UEIdExt{}

// UEIdExt UE Identity
type UEIdExt struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string `json:"gpsi,omitempty"`
}

// NewUEIdExt instantiates a new UEIdExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUEIdExt() *UEIdExt {
	this := UEIdExt{}
	return &this
}

// NewUEIdExtWithDefaults instantiates a new UEIdExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUEIdExtWithDefaults() *UEIdExt {
	this := UEIdExt{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UEIdExt) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEIdExt) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UEIdExt) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UEIdExt) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UEIdExt) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEIdExt) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UEIdExt) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UEIdExt) SetGpsi(v string) {
	o.Gpsi = &v
}

func (o UEIdExt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UEIdExt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	return toSerialize, nil
}

type NullableUEIdExt struct {
	value *UEIdExt
	isSet bool
}

func (v NullableUEIdExt) Get() *UEIdExt {
	return v.value
}

func (v *NullableUEIdExt) Set(val *UEIdExt) {
	v.value = val
	v.isSet = true
}

func (v NullableUEIdExt) IsSet() bool {
	return v.isSet
}

func (v *NullableUEIdExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUEIdExt(val *UEIdExt) *NullableUEIdExt {
	return &NullableUEIdExt{value: val, isSet: true}
}

func (v NullableUEIdExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUEIdExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
