/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
)

// checks if the Model3GAkaAv type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model3GAkaAv{}

// Model3GAkaAv Contains RAND, XRES, AUTN, CK, and IK
type Model3GAkaAv struct {
	Rand string `json:"rand"`
	Xres string `json:"xres"`
	Autn string `json:"autn"`
	Ck   string `json:"ck"`
	Ik   string `json:"ik"`
}

// NewModel3GAkaAv instantiates a new Model3GAkaAv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel3GAkaAv(rand string, xres string, autn string, ck string, ik string) *Model3GAkaAv {
	this := Model3GAkaAv{}
	this.Rand = rand
	this.Xres = xres
	this.Autn = autn
	this.Ck = ck
	this.Ik = ik
	return &this
}

// NewModel3GAkaAvWithDefaults instantiates a new Model3GAkaAv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel3GAkaAvWithDefaults() *Model3GAkaAv {
	this := Model3GAkaAv{}
	return &this
}

// GetRand returns the Rand field value
func (o *Model3GAkaAv) GetRand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rand
}

// GetRandOk returns a tuple with the Rand field value
// and a boolean to check if the value has been set.
func (o *Model3GAkaAv) GetRandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rand, true
}

// SetRand sets field value
func (o *Model3GAkaAv) SetRand(v string) {
	o.Rand = v
}

// GetXres returns the Xres field value
func (o *Model3GAkaAv) GetXres() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Xres
}

// GetXresOk returns a tuple with the Xres field value
// and a boolean to check if the value has been set.
func (o *Model3GAkaAv) GetXresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Xres, true
}

// SetXres sets field value
func (o *Model3GAkaAv) SetXres(v string) {
	o.Xres = v
}

// GetAutn returns the Autn field value
func (o *Model3GAkaAv) GetAutn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Autn
}

// GetAutnOk returns a tuple with the Autn field value
// and a boolean to check if the value has been set.
func (o *Model3GAkaAv) GetAutnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Autn, true
}

// SetAutn sets field value
func (o *Model3GAkaAv) SetAutn(v string) {
	o.Autn = v
}

// GetCk returns the Ck field value
func (o *Model3GAkaAv) GetCk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ck
}

// GetCkOk returns a tuple with the Ck field value
// and a boolean to check if the value has been set.
func (o *Model3GAkaAv) GetCkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ck, true
}

// SetCk sets field value
func (o *Model3GAkaAv) SetCk(v string) {
	o.Ck = v
}

// GetIk returns the Ik field value
func (o *Model3GAkaAv) GetIk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ik
}

// GetIkOk returns a tuple with the Ik field value
// and a boolean to check if the value has been set.
func (o *Model3GAkaAv) GetIkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ik, true
}

// SetIk sets field value
func (o *Model3GAkaAv) SetIk(v string) {
	o.Ik = v
}

func (o Model3GAkaAv) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model3GAkaAv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rand"] = o.Rand
	toSerialize["xres"] = o.Xres
	toSerialize["autn"] = o.Autn
	toSerialize["ck"] = o.Ck
	toSerialize["ik"] = o.Ik
	return toSerialize, nil
}

type NullableModel3GAkaAv struct {
	value *Model3GAkaAv
	isSet bool
}

func (v NullableModel3GAkaAv) Get() *Model3GAkaAv {
	return v.value
}

func (v *NullableModel3GAkaAv) Set(val *Model3GAkaAv) {
	v.value = val
	v.isSet = true
}

func (v NullableModel3GAkaAv) IsSet() bool {
	return v.isSet
}

func (v *NullableModel3GAkaAv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel3GAkaAv(val *Model3GAkaAv) *NullableModel3GAkaAv {
	return &NullableModel3GAkaAv{value: val, isSet: true}
}

func (v NullableModel3GAkaAv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel3GAkaAv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
