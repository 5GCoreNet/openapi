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

// checks if the N3iwfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N3iwfFunctionSingleAllOf1{}

// N3iwfFunctionSingleAllOf1 struct for N3iwfFunctionSingleAllOf1
type N3iwfFunctionSingleAllOf1 struct {
	EPN3 []EPN3Single `json:"EP_N3,omitempty"`
	EPN4 []EPN4Single `json:"EP_N4,omitempty"`
}

// NewN3iwfFunctionSingleAllOf1 instantiates a new N3iwfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN3iwfFunctionSingleAllOf1() *N3iwfFunctionSingleAllOf1 {
	this := N3iwfFunctionSingleAllOf1{}
	return &this
}

// NewN3iwfFunctionSingleAllOf1WithDefaults instantiates a new N3iwfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN3iwfFunctionSingleAllOf1WithDefaults() *N3iwfFunctionSingleAllOf1 {
	this := N3iwfFunctionSingleAllOf1{}
	return &this
}

// GetEPN3 returns the EPN3 field value if set, zero value otherwise.
func (o *N3iwfFunctionSingleAllOf1) GetEPN3() []EPN3Single {
	if o == nil || IsNil(o.EPN3) {
		var ret []EPN3Single
		return ret
	}
	return o.EPN3
}

// GetEPN3Ok returns a tuple with the EPN3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3iwfFunctionSingleAllOf1) GetEPN3Ok() ([]EPN3Single, bool) {
	if o == nil || IsNil(o.EPN3) {
		return nil, false
	}
	return o.EPN3, true
}

// HasEPN3 returns a boolean if a field has been set.
func (o *N3iwfFunctionSingleAllOf1) HasEPN3() bool {
	if o != nil && !IsNil(o.EPN3) {
		return true
	}

	return false
}

// SetEPN3 gets a reference to the given []EPN3Single and assigns it to the EPN3 field.
func (o *N3iwfFunctionSingleAllOf1) SetEPN3(v []EPN3Single) {
	o.EPN3 = v
}

// GetEPN4 returns the EPN4 field value if set, zero value otherwise.
func (o *N3iwfFunctionSingleAllOf1) GetEPN4() []EPN4Single {
	if o == nil || IsNil(o.EPN4) {
		var ret []EPN4Single
		return ret
	}
	return o.EPN4
}

// GetEPN4Ok returns a tuple with the EPN4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3iwfFunctionSingleAllOf1) GetEPN4Ok() ([]EPN4Single, bool) {
	if o == nil || IsNil(o.EPN4) {
		return nil, false
	}
	return o.EPN4, true
}

// HasEPN4 returns a boolean if a field has been set.
func (o *N3iwfFunctionSingleAllOf1) HasEPN4() bool {
	if o != nil && !IsNil(o.EPN4) {
		return true
	}

	return false
}

// SetEPN4 gets a reference to the given []EPN4Single and assigns it to the EPN4 field.
func (o *N3iwfFunctionSingleAllOf1) SetEPN4(v []EPN4Single) {
	o.EPN4 = v
}

func (o N3iwfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N3iwfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN3) {
		toSerialize["EP_N3"] = o.EPN3
	}
	if !IsNil(o.EPN4) {
		toSerialize["EP_N4"] = o.EPN4
	}
	return toSerialize, nil
}

type NullableN3iwfFunctionSingleAllOf1 struct {
	value *N3iwfFunctionSingleAllOf1
	isSet bool
}

func (v NullableN3iwfFunctionSingleAllOf1) Get() *N3iwfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableN3iwfFunctionSingleAllOf1) Set(val *N3iwfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableN3iwfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableN3iwfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN3iwfFunctionSingleAllOf1(val *N3iwfFunctionSingleAllOf1) *NullableN3iwfFunctionSingleAllOf1 {
	return &NullableN3iwfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableN3iwfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN3iwfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
