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

// checks if the UdrFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdrFunctionSingleAllOf{}

// UdrFunctionSingleAllOf struct for UdrFunctionSingleAllOf
type UdrFunctionSingleAllOf struct {
	Attributes *UdrFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewUdrFunctionSingleAllOf instantiates a new UdrFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdrFunctionSingleAllOf() *UdrFunctionSingleAllOf {
	this := UdrFunctionSingleAllOf{}
	return &this
}

// NewUdrFunctionSingleAllOfWithDefaults instantiates a new UdrFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdrFunctionSingleAllOfWithDefaults() *UdrFunctionSingleAllOf {
	this := UdrFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UdrFunctionSingleAllOf) GetAttributes() UdrFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret UdrFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrFunctionSingleAllOf) GetAttributesOk() (*UdrFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UdrFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UdrFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *UdrFunctionSingleAllOf) SetAttributes(v UdrFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o UdrFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdrFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableUdrFunctionSingleAllOf struct {
	value *UdrFunctionSingleAllOf
	isSet bool
}

func (v NullableUdrFunctionSingleAllOf) Get() *UdrFunctionSingleAllOf {
	return v.value
}

func (v *NullableUdrFunctionSingleAllOf) Set(val *UdrFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUdrFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUdrFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdrFunctionSingleAllOf(val *UdrFunctionSingleAllOf) *NullableUdrFunctionSingleAllOf {
	return &NullableUdrFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableUdrFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdrFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
