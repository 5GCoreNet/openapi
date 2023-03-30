/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the PlmnId1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnId1{}

// PlmnId1 struct for PlmnId1
type PlmnId1 struct {
	Mcc *string `json:"mcc,omitempty"`
	Mnc *string `json:"mnc,omitempty"`
}

// NewPlmnId1 instantiates a new PlmnId1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnId1() *PlmnId1 {
	this := PlmnId1{}
	return &this
}

// NewPlmnId1WithDefaults instantiates a new PlmnId1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnId1WithDefaults() *PlmnId1 {
	this := PlmnId1{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *PlmnId1) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnId1) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *PlmnId1) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *PlmnId1) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *PlmnId1) GetMnc() string {
	if o == nil || IsNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnId1) GetMncOk() (*string, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *PlmnId1) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *PlmnId1) SetMnc(v string) {
	o.Mnc = &v
}

func (o PlmnId1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnId1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	return toSerialize, nil
}

type NullablePlmnId1 struct {
	value *PlmnId1
	isSet bool
}

func (v NullablePlmnId1) Get() *PlmnId1 {
	return v.value
}

func (v *NullablePlmnId1) Set(val *PlmnId1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnId1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnId1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnId1(val *PlmnId1) *NullablePlmnId1 {
	return &NullablePlmnId1{value: val, isSet: true}
}

func (v NullablePlmnId1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnId1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


