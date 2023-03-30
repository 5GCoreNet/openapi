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

// checks if the SmfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfFunctionSingleAllOf{}

// SmfFunctionSingleAllOf struct for SmfFunctionSingleAllOf
type SmfFunctionSingleAllOf struct {
	Attributes *SmfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewSmfFunctionSingleAllOf instantiates a new SmfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfFunctionSingleAllOf() *SmfFunctionSingleAllOf {
	this := SmfFunctionSingleAllOf{}
	return &this
}

// NewSmfFunctionSingleAllOfWithDefaults instantiates a new SmfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfFunctionSingleAllOfWithDefaults() *SmfFunctionSingleAllOf {
	this := SmfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SmfFunctionSingleAllOf) GetAttributes() SmfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SmfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfFunctionSingleAllOf) GetAttributesOk() (*SmfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SmfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SmfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *SmfFunctionSingleAllOf) SetAttributes(v SmfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o SmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSmfFunctionSingleAllOf struct {
	value *SmfFunctionSingleAllOf
	isSet bool
}

func (v NullableSmfFunctionSingleAllOf) Get() *SmfFunctionSingleAllOf {
	return v.value
}

func (v *NullableSmfFunctionSingleAllOf) Set(val *SmfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfFunctionSingleAllOf(val *SmfFunctionSingleAllOf) *NullableSmfFunctionSingleAllOf {
	return &NullableSmfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableSmfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


