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

// checks if the GnbDuFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GnbDuFunctionSingleAllOf{}

// GnbDuFunctionSingleAllOf struct for GnbDuFunctionSingleAllOf
type GnbDuFunctionSingleAllOf struct {
	Attributes *GnbDuFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewGnbDuFunctionSingleAllOf instantiates a new GnbDuFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbDuFunctionSingleAllOf() *GnbDuFunctionSingleAllOf {
	this := GnbDuFunctionSingleAllOf{}
	return &this
}

// NewGnbDuFunctionSingleAllOfWithDefaults instantiates a new GnbDuFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbDuFunctionSingleAllOfWithDefaults() *GnbDuFunctionSingleAllOf {
	this := GnbDuFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GnbDuFunctionSingleAllOf) GetAttributes() GnbDuFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GnbDuFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbDuFunctionSingleAllOf) GetAttributesOk() (*GnbDuFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GnbDuFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GnbDuFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *GnbDuFunctionSingleAllOf) SetAttributes(v GnbDuFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o GnbDuFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GnbDuFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGnbDuFunctionSingleAllOf struct {
	value *GnbDuFunctionSingleAllOf
	isSet bool
}

func (v NullableGnbDuFunctionSingleAllOf) Get() *GnbDuFunctionSingleAllOf {
	return v.value
}

func (v *NullableGnbDuFunctionSingleAllOf) Set(val *GnbDuFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbDuFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbDuFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbDuFunctionSingleAllOf(val *GnbDuFunctionSingleAllOf) *NullableGnbDuFunctionSingleAllOf {
	return &NullableGnbDuFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableGnbDuFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbDuFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


