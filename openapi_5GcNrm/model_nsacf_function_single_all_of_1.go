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

// checks if the NsacfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsacfFunctionSingleAllOf1{}

// NsacfFunctionSingleAllOf1 struct for NsacfFunctionSingleAllOf1
type NsacfFunctionSingleAllOf1 struct {
	EPN60 []EPN60Single `json:"EP_N60,omitempty"`
}

// NewNsacfFunctionSingleAllOf1 instantiates a new NsacfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfFunctionSingleAllOf1() *NsacfFunctionSingleAllOf1 {
	this := NsacfFunctionSingleAllOf1{}
	return &this
}

// NewNsacfFunctionSingleAllOf1WithDefaults instantiates a new NsacfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfFunctionSingleAllOf1WithDefaults() *NsacfFunctionSingleAllOf1 {
	this := NsacfFunctionSingleAllOf1{}
	return &this
}

// GetEPN60 returns the EPN60 field value if set, zero value otherwise.
func (o *NsacfFunctionSingleAllOf1) GetEPN60() []EPN60Single {
	if o == nil || IsNil(o.EPN60) {
		var ret []EPN60Single
		return ret
	}
	return o.EPN60
}

// GetEPN60Ok returns a tuple with the EPN60 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfFunctionSingleAllOf1) GetEPN60Ok() ([]EPN60Single, bool) {
	if o == nil || IsNil(o.EPN60) {
		return nil, false
	}
	return o.EPN60, true
}

// HasEPN60 returns a boolean if a field has been set.
func (o *NsacfFunctionSingleAllOf1) HasEPN60() bool {
	if o != nil && !IsNil(o.EPN60) {
		return true
	}

	return false
}

// SetEPN60 gets a reference to the given []EPN60Single and assigns it to the EPN60 field.
func (o *NsacfFunctionSingleAllOf1) SetEPN60(v []EPN60Single) {
	o.EPN60 = v
}

func (o NsacfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsacfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN60) {
		toSerialize["EP_N60"] = o.EPN60
	}
	return toSerialize, nil
}

type NullableNsacfFunctionSingleAllOf1 struct {
	value *NsacfFunctionSingleAllOf1
	isSet bool
}

func (v NullableNsacfFunctionSingleAllOf1) Get() *NsacfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableNsacfFunctionSingleAllOf1) Set(val *NsacfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfFunctionSingleAllOf1(val *NsacfFunctionSingleAllOf1) *NullableNsacfFunctionSingleAllOf1 {
	return &NullableNsacfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableNsacfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


