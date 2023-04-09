/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the GnbCuCpFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GnbCuCpFunctionSingleAllOf{}

// GnbCuCpFunctionSingleAllOf struct for GnbCuCpFunctionSingleAllOf
type GnbCuCpFunctionSingleAllOf struct {
	Attributes *GnbCuCpFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewGnbCuCpFunctionSingleAllOf instantiates a new GnbCuCpFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbCuCpFunctionSingleAllOf() *GnbCuCpFunctionSingleAllOf {
	this := GnbCuCpFunctionSingleAllOf{}
	return &this
}

// NewGnbCuCpFunctionSingleAllOfWithDefaults instantiates a new GnbCuCpFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbCuCpFunctionSingleAllOfWithDefaults() *GnbCuCpFunctionSingleAllOf {
	this := GnbCuCpFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf) GetAttributes() GnbCuCpFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GnbCuCpFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf) GetAttributesOk() (*GnbCuCpFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GnbCuCpFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *GnbCuCpFunctionSingleAllOf) SetAttributes(v GnbCuCpFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o GnbCuCpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GnbCuCpFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGnbCuCpFunctionSingleAllOf struct {
	value *GnbCuCpFunctionSingleAllOf
	isSet bool
}

func (v NullableGnbCuCpFunctionSingleAllOf) Get() *GnbCuCpFunctionSingleAllOf {
	return v.value
}

func (v *NullableGnbCuCpFunctionSingleAllOf) Set(val *GnbCuCpFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbCuCpFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbCuCpFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbCuCpFunctionSingleAllOf(val *GnbCuCpFunctionSingleAllOf) *NullableGnbCuCpFunctionSingleAllOf {
	return &NullableGnbCuCpFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableGnbCuCpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbCuCpFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
