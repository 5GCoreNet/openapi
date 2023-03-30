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

// checks if the UpfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfFunctionSingleAllOf{}

// UpfFunctionSingleAllOf struct for UpfFunctionSingleAllOf
type UpfFunctionSingleAllOf struct {
	Attributes *UpfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewUpfFunctionSingleAllOf instantiates a new UpfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfFunctionSingleAllOf() *UpfFunctionSingleAllOf {
	this := UpfFunctionSingleAllOf{}
	return &this
}

// NewUpfFunctionSingleAllOfWithDefaults instantiates a new UpfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfFunctionSingleAllOfWithDefaults() *UpfFunctionSingleAllOf {
	this := UpfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOf) GetAttributes() UpfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret UpfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOf) GetAttributesOk() (*UpfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UpfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *UpfFunctionSingleAllOf) SetAttributes(v UpfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o UpfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableUpfFunctionSingleAllOf struct {
	value *UpfFunctionSingleAllOf
	isSet bool
}

func (v NullableUpfFunctionSingleAllOf) Get() *UpfFunctionSingleAllOf {
	return v.value
}

func (v *NullableUpfFunctionSingleAllOf) Set(val *UpfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfFunctionSingleAllOf(val *UpfFunctionSingleAllOf) *NullableUpfFunctionSingleAllOf {
	return &NullableUpfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableUpfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


