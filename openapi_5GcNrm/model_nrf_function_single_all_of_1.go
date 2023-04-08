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

// checks if the NrfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrfFunctionSingleAllOf1{}

// NrfFunctionSingleAllOf1 struct for NrfFunctionSingleAllOf1
type NrfFunctionSingleAllOf1 struct {
	EPN27 []EPN27Single `json:"EP_N27,omitempty"`
}

// NewNrfFunctionSingleAllOf1 instantiates a new NrfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrfFunctionSingleAllOf1() *NrfFunctionSingleAllOf1 {
	this := NrfFunctionSingleAllOf1{}
	return &this
}

// NewNrfFunctionSingleAllOf1WithDefaults instantiates a new NrfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrfFunctionSingleAllOf1WithDefaults() *NrfFunctionSingleAllOf1 {
	this := NrfFunctionSingleAllOf1{}
	return &this
}

// GetEPN27 returns the EPN27 field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOf1) GetEPN27() []EPN27Single {
	if o == nil || isNil(o.EPN27) {
		var ret []EPN27Single
		return ret
	}
	return o.EPN27
}

// GetEPN27Ok returns a tuple with the EPN27 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOf1) GetEPN27Ok() ([]EPN27Single, bool) {
	if o == nil || isNil(o.EPN27) {
		return nil, false
	}
	return o.EPN27, true
}

// HasEPN27 returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOf1) HasEPN27() bool {
	if o != nil && !isNil(o.EPN27) {
		return true
	}

	return false
}

// SetEPN27 gets a reference to the given []EPN27Single and assigns it to the EPN27 field.
func (o *NrfFunctionSingleAllOf1) SetEPN27(v []EPN27Single) {
	o.EPN27 = v
}

func (o NrfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EPN27) {
		toSerialize["EP_N27"] = o.EPN27
	}
	return toSerialize, nil
}

type NullableNrfFunctionSingleAllOf1 struct {
	value *NrfFunctionSingleAllOf1
	isSet bool
}

func (v NullableNrfFunctionSingleAllOf1) Get() *NrfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableNrfFunctionSingleAllOf1) Set(val *NrfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfFunctionSingleAllOf1(val *NrfFunctionSingleAllOf1) *NullableNrfFunctionSingleAllOf1 {
	return &NullableNrfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableNrfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


