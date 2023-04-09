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

// checks if the NgeirFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgeirFunctionSingleAllOf{}

// NgeirFunctionSingleAllOf struct for NgeirFunctionSingleAllOf
type NgeirFunctionSingleAllOf struct {
	Attributes *NgeirFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNgeirFunctionSingleAllOf instantiates a new NgeirFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgeirFunctionSingleAllOf() *NgeirFunctionSingleAllOf {
	this := NgeirFunctionSingleAllOf{}
	return &this
}

// NewNgeirFunctionSingleAllOfWithDefaults instantiates a new NgeirFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgeirFunctionSingleAllOfWithDefaults() *NgeirFunctionSingleAllOf {
	this := NgeirFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NgeirFunctionSingleAllOf) GetAttributes() NgeirFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NgeirFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgeirFunctionSingleAllOf) GetAttributesOk() (*NgeirFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NgeirFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NgeirFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NgeirFunctionSingleAllOf) SetAttributes(v NgeirFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NgeirFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgeirFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableNgeirFunctionSingleAllOf struct {
	value *NgeirFunctionSingleAllOf
	isSet bool
}

func (v NullableNgeirFunctionSingleAllOf) Get() *NgeirFunctionSingleAllOf {
	return v.value
}

func (v *NullableNgeirFunctionSingleAllOf) Set(val *NgeirFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgeirFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgeirFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgeirFunctionSingleAllOf(val *NgeirFunctionSingleAllOf) *NullableNgeirFunctionSingleAllOf {
	return &NullableNgeirFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableNgeirFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgeirFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
