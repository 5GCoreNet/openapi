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

// checks if the NgeirFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgeirFunctionSingleAllOf1{}

// NgeirFunctionSingleAllOf1 struct for NgeirFunctionSingleAllOf1
type NgeirFunctionSingleAllOf1 struct {
	EPN17 []EPN17Single `json:"EP_N17,omitempty"`
}

// NewNgeirFunctionSingleAllOf1 instantiates a new NgeirFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgeirFunctionSingleAllOf1() *NgeirFunctionSingleAllOf1 {
	this := NgeirFunctionSingleAllOf1{}
	return &this
}

// NewNgeirFunctionSingleAllOf1WithDefaults instantiates a new NgeirFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgeirFunctionSingleAllOf1WithDefaults() *NgeirFunctionSingleAllOf1 {
	this := NgeirFunctionSingleAllOf1{}
	return &this
}

// GetEPN17 returns the EPN17 field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOf1) GetEPN17() []EPN17Single {
	if o == nil || IsNil(o.EPN17) {
		var ret []EPN17Single
		return ret
	}
	return o.EPN17
}

// GetEPN17Ok returns a tuple with the EPN17 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOf1) GetEPN17Ok() ([]EPN17Single, bool) {
	if o == nil || IsNil(o.EPN17) {
		return nil, false
	}
	return o.EPN17, true
}

// HasEPN17 returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOf1) HasEPN17() bool {
	if o != nil && !IsNil(o.EPN17) {
		return true
	}

	return false
}

// SetEPN17 gets a reference to the given []EPN17Single and assigns it to the EPN17 field.
func (o *NgeirFunctionSingleAllOf1) SetEPN17(v []EPN17Single) {
	o.EPN17 = v
}

func (o NgeirFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgeirFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN17) {
		toSerialize["EP_N17"] = o.EPN17
	}
	return toSerialize, nil
}

type NullableNgeirFunctionSingleAllOf1 struct {
	value *NgeirFunctionSingleAllOf1
	isSet bool
}

func (v NullableNgeirFunctionSingleAllOf1) Get() *NgeirFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableNgeirFunctionSingleAllOf1) Set(val *NgeirFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNgeirFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNgeirFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgeirFunctionSingleAllOf1(val *NgeirFunctionSingleAllOf1) *NullableNgeirFunctionSingleAllOf1 {
	return &NullableNgeirFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableNgeirFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgeirFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


