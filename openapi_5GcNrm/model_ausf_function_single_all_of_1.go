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

// checks if the AusfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AusfFunctionSingleAllOf1{}

// AusfFunctionSingleAllOf1 struct for AusfFunctionSingleAllOf1
type AusfFunctionSingleAllOf1 struct {
	EPN12 []EPN12Single `json:"EP_N12,omitempty"`
	EPN13 []EPN13Single `json:"EP_N13,omitempty"`
}

// NewAusfFunctionSingleAllOf1 instantiates a new AusfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfFunctionSingleAllOf1() *AusfFunctionSingleAllOf1 {
	this := AusfFunctionSingleAllOf1{}
	return &this
}

// NewAusfFunctionSingleAllOf1WithDefaults instantiates a new AusfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfFunctionSingleAllOf1WithDefaults() *AusfFunctionSingleAllOf1 {
	this := AusfFunctionSingleAllOf1{}
	return &this
}

// GetEPN12 returns the EPN12 field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOf1) GetEPN12() []EPN12Single {
	if o == nil || IsNil(o.EPN12) {
		var ret []EPN12Single
		return ret
	}
	return o.EPN12
}

// GetEPN12Ok returns a tuple with the EPN12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOf1) GetEPN12Ok() ([]EPN12Single, bool) {
	if o == nil || IsNil(o.EPN12) {
		return nil, false
	}
	return o.EPN12, true
}

// HasEPN12 returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOf1) HasEPN12() bool {
	if o != nil && !IsNil(o.EPN12) {
		return true
	}

	return false
}

// SetEPN12 gets a reference to the given []EPN12Single and assigns it to the EPN12 field.
func (o *AusfFunctionSingleAllOf1) SetEPN12(v []EPN12Single) {
	o.EPN12 = v
}

// GetEPN13 returns the EPN13 field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOf1) GetEPN13() []EPN13Single {
	if o == nil || IsNil(o.EPN13) {
		var ret []EPN13Single
		return ret
	}
	return o.EPN13
}

// GetEPN13Ok returns a tuple with the EPN13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOf1) GetEPN13Ok() ([]EPN13Single, bool) {
	if o == nil || IsNil(o.EPN13) {
		return nil, false
	}
	return o.EPN13, true
}

// HasEPN13 returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOf1) HasEPN13() bool {
	if o != nil && !IsNil(o.EPN13) {
		return true
	}

	return false
}

// SetEPN13 gets a reference to the given []EPN13Single and assigns it to the EPN13 field.
func (o *AusfFunctionSingleAllOf1) SetEPN13(v []EPN13Single) {
	o.EPN13 = v
}

func (o AusfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AusfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN12) {
		toSerialize["EP_N12"] = o.EPN12
	}
	if !IsNil(o.EPN13) {
		toSerialize["EP_N13"] = o.EPN13
	}
	return toSerialize, nil
}

type NullableAusfFunctionSingleAllOf1 struct {
	value *AusfFunctionSingleAllOf1
	isSet bool
}

func (v NullableAusfFunctionSingleAllOf1) Get() *AusfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableAusfFunctionSingleAllOf1) Set(val *AusfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfFunctionSingleAllOf1(val *AusfFunctionSingleAllOf1) *NullableAusfFunctionSingleAllOf1 {
	return &NullableAusfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableAusfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
