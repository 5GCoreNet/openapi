/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the SeafData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeafData{}

// SeafData Represents SEAF data derived from data received from AUSF
type SeafData struct {
	NgKsi                NgKsi   `json:"ngKsi"`
	KeyAmf               KeyAmf  `json:"keyAmf"`
	Nh                   *string `json:"nh,omitempty"`
	Ncc                  *int32  `json:"ncc,omitempty"`
	KeyAmfChangeInd      *bool   `json:"keyAmfChangeInd,omitempty"`
	KeyAmfHDerivationInd *bool   `json:"keyAmfHDerivationInd,omitempty"`
}

// NewSeafData instantiates a new SeafData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeafData(ngKsi NgKsi, keyAmf KeyAmf) *SeafData {
	this := SeafData{}
	this.NgKsi = ngKsi
	this.KeyAmf = keyAmf
	return &this
}

// NewSeafDataWithDefaults instantiates a new SeafData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeafDataWithDefaults() *SeafData {
	this := SeafData{}
	return &this
}

// GetNgKsi returns the NgKsi field value
func (o *SeafData) GetNgKsi() NgKsi {
	if o == nil {
		var ret NgKsi
		return ret
	}

	return o.NgKsi
}

// GetNgKsiOk returns a tuple with the NgKsi field value
// and a boolean to check if the value has been set.
func (o *SeafData) GetNgKsiOk() (*NgKsi, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NgKsi, true
}

// SetNgKsi sets field value
func (o *SeafData) SetNgKsi(v NgKsi) {
	o.NgKsi = v
}

// GetKeyAmf returns the KeyAmf field value
func (o *SeafData) GetKeyAmf() KeyAmf {
	if o == nil {
		var ret KeyAmf
		return ret
	}

	return o.KeyAmf
}

// GetKeyAmfOk returns a tuple with the KeyAmf field value
// and a boolean to check if the value has been set.
func (o *SeafData) GetKeyAmfOk() (*KeyAmf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAmf, true
}

// SetKeyAmf sets field value
func (o *SeafData) SetKeyAmf(v KeyAmf) {
	o.KeyAmf = v
}

// GetNh returns the Nh field value if set, zero value otherwise.
func (o *SeafData) GetNh() string {
	if o == nil || IsNil(o.Nh) {
		var ret string
		return ret
	}
	return *o.Nh
}

// GetNhOk returns a tuple with the Nh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeafData) GetNhOk() (*string, bool) {
	if o == nil || IsNil(o.Nh) {
		return nil, false
	}
	return o.Nh, true
}

// HasNh returns a boolean if a field has been set.
func (o *SeafData) HasNh() bool {
	if o != nil && !IsNil(o.Nh) {
		return true
	}

	return false
}

// SetNh gets a reference to the given string and assigns it to the Nh field.
func (o *SeafData) SetNh(v string) {
	o.Nh = &v
}

// GetNcc returns the Ncc field value if set, zero value otherwise.
func (o *SeafData) GetNcc() int32 {
	if o == nil || IsNil(o.Ncc) {
		var ret int32
		return ret
	}
	return *o.Ncc
}

// GetNccOk returns a tuple with the Ncc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeafData) GetNccOk() (*int32, bool) {
	if o == nil || IsNil(o.Ncc) {
		return nil, false
	}
	return o.Ncc, true
}

// HasNcc returns a boolean if a field has been set.
func (o *SeafData) HasNcc() bool {
	if o != nil && !IsNil(o.Ncc) {
		return true
	}

	return false
}

// SetNcc gets a reference to the given int32 and assigns it to the Ncc field.
func (o *SeafData) SetNcc(v int32) {
	o.Ncc = &v
}

// GetKeyAmfChangeInd returns the KeyAmfChangeInd field value if set, zero value otherwise.
func (o *SeafData) GetKeyAmfChangeInd() bool {
	if o == nil || IsNil(o.KeyAmfChangeInd) {
		var ret bool
		return ret
	}
	return *o.KeyAmfChangeInd
}

// GetKeyAmfChangeIndOk returns a tuple with the KeyAmfChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeafData) GetKeyAmfChangeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyAmfChangeInd) {
		return nil, false
	}
	return o.KeyAmfChangeInd, true
}

// HasKeyAmfChangeInd returns a boolean if a field has been set.
func (o *SeafData) HasKeyAmfChangeInd() bool {
	if o != nil && !IsNil(o.KeyAmfChangeInd) {
		return true
	}

	return false
}

// SetKeyAmfChangeInd gets a reference to the given bool and assigns it to the KeyAmfChangeInd field.
func (o *SeafData) SetKeyAmfChangeInd(v bool) {
	o.KeyAmfChangeInd = &v
}

// GetKeyAmfHDerivationInd returns the KeyAmfHDerivationInd field value if set, zero value otherwise.
func (o *SeafData) GetKeyAmfHDerivationInd() bool {
	if o == nil || IsNil(o.KeyAmfHDerivationInd) {
		var ret bool
		return ret
	}
	return *o.KeyAmfHDerivationInd
}

// GetKeyAmfHDerivationIndOk returns a tuple with the KeyAmfHDerivationInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeafData) GetKeyAmfHDerivationIndOk() (*bool, bool) {
	if o == nil || IsNil(o.KeyAmfHDerivationInd) {
		return nil, false
	}
	return o.KeyAmfHDerivationInd, true
}

// HasKeyAmfHDerivationInd returns a boolean if a field has been set.
func (o *SeafData) HasKeyAmfHDerivationInd() bool {
	if o != nil && !IsNil(o.KeyAmfHDerivationInd) {
		return true
	}

	return false
}

// SetKeyAmfHDerivationInd gets a reference to the given bool and assigns it to the KeyAmfHDerivationInd field.
func (o *SeafData) SetKeyAmfHDerivationInd(v bool) {
	o.KeyAmfHDerivationInd = &v
}

func (o SeafData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeafData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ngKsi"] = o.NgKsi
	toSerialize["keyAmf"] = o.KeyAmf
	if !IsNil(o.Nh) {
		toSerialize["nh"] = o.Nh
	}
	if !IsNil(o.Ncc) {
		toSerialize["ncc"] = o.Ncc
	}
	if !IsNil(o.KeyAmfChangeInd) {
		toSerialize["keyAmfChangeInd"] = o.KeyAmfChangeInd
	}
	if !IsNil(o.KeyAmfHDerivationInd) {
		toSerialize["keyAmfHDerivationInd"] = o.KeyAmfHDerivationInd
	}
	return toSerialize, nil
}

type NullableSeafData struct {
	value *SeafData
	isSet bool
}

func (v NullableSeafData) Get() *SeafData {
	return v.value
}

func (v *NullableSeafData) Set(val *SeafData) {
	v.value = val
	v.isSet = true
}

func (v NullableSeafData) IsSet() bool {
	return v.isSet
}

func (v *NullableSeafData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeafData(val *SeafData) *NullableSeafData {
	return &NullableSeafData{value: val, isSet: true}
}

func (v NullableSeafData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeafData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
