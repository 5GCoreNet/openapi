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

// checks if the NwdafFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafFunctionSingleAllOf{}

// NwdafFunctionSingleAllOf struct for NwdafFunctionSingleAllOf
type NwdafFunctionSingleAllOf struct {
	Attributes *NwdafFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNwdafFunctionSingleAllOf instantiates a new NwdafFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafFunctionSingleAllOf() *NwdafFunctionSingleAllOf {
	this := NwdafFunctionSingleAllOf{}
	return &this
}

// NewNwdafFunctionSingleAllOfWithDefaults instantiates a new NwdafFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafFunctionSingleAllOfWithDefaults() *NwdafFunctionSingleAllOf {
	this := NwdafFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOf) GetAttributes() NwdafFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NwdafFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOf) GetAttributesOk() (*NwdafFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NwdafFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NwdafFunctionSingleAllOf) SetAttributes(v NwdafFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NwdafFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableNwdafFunctionSingleAllOf struct {
	value *NwdafFunctionSingleAllOf
	isSet bool
}

func (v NullableNwdafFunctionSingleAllOf) Get() *NwdafFunctionSingleAllOf {
	return v.value
}

func (v *NullableNwdafFunctionSingleAllOf) Set(val *NwdafFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafFunctionSingleAllOf(val *NwdafFunctionSingleAllOf) *NullableNwdafFunctionSingleAllOf {
	return &NullableNwdafFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableNwdafFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


