/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the Tai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tai{}

// Tai struct for Tai
type Tai struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
	Tac *string `json:"tac,omitempty"`
}

// NewTai instantiates a new Tai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTai() *Tai {
	this := Tai{}
	return &this
}

// NewTaiWithDefaults instantiates a new Tai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiWithDefaults() *Tai {
	this := Tai{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *Tai) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *Tai) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *Tai) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *Tai) GetMnc() string {
	if o == nil || IsNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetMncOk() (*string, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *Tai) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *Tai) SetMnc(v string) {
	o.Mnc = &v
}

// GetTac returns the Tac field value if set, zero value otherwise.
func (o *Tai) GetTac() string {
	if o == nil || IsNil(o.Tac) {
		var ret string
		return ret
	}
	return *o.Tac
}

// GetTacOk returns a tuple with the Tac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tai) GetTacOk() (*string, bool) {
	if o == nil || IsNil(o.Tac) {
		return nil, false
	}
	return o.Tac, true
}

// HasTac returns a boolean if a field has been set.
func (o *Tai) HasTac() bool {
	if o != nil && !IsNil(o.Tac) {
		return true
	}

	return false
}

// SetTac gets a reference to the given string and assigns it to the Tac field.
func (o *Tai) SetTac(v string) {
	o.Tac = &v
}

func (o Tai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !IsNil(o.Tac) {
		toSerialize["tac"] = o.Tac
	}
	return toSerialize, nil
}

type NullableTai struct {
	value *Tai
	isSet bool
}

func (v NullableTai) Get() *Tai {
	return v.value
}

func (v *NullableTai) Set(val *Tai) {
	v.value = val
	v.isSet = true
}

func (v NullableTai) IsSet() bool {
	return v.isSet
}

func (v *NullableTai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTai(val *Tai) *NullableTai {
	return &NullableTai{value: val, isSet: true}
}

func (v NullableTai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
