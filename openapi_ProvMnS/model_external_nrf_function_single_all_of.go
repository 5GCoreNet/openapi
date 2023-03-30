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

// checks if the ExternalNrfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalNrfFunctionSingleAllOf{}

// ExternalNrfFunctionSingleAllOf struct for ExternalNrfFunctionSingleAllOf
type ExternalNrfFunctionSingleAllOf struct {
	Attributes *ExternalNrfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalNrfFunctionSingleAllOf instantiates a new ExternalNrfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalNrfFunctionSingleAllOf() *ExternalNrfFunctionSingleAllOf {
	this := ExternalNrfFunctionSingleAllOf{}
	return &this
}

// NewExternalNrfFunctionSingleAllOfWithDefaults instantiates a new ExternalNrfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalNrfFunctionSingleAllOfWithDefaults() *ExternalNrfFunctionSingleAllOf {
	this := ExternalNrfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalNrfFunctionSingleAllOf) GetAttributes() ExternalNrfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalNrfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrfFunctionSingleAllOf) GetAttributesOk() (*ExternalNrfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalNrfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalNrfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalNrfFunctionSingleAllOf) SetAttributes(v ExternalNrfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalNrfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalNrfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalNrfFunctionSingleAllOf struct {
	value *ExternalNrfFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalNrfFunctionSingleAllOf) Get() *ExternalNrfFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalNrfFunctionSingleAllOf) Set(val *ExternalNrfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalNrfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalNrfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalNrfFunctionSingleAllOf(val *ExternalNrfFunctionSingleAllOf) *NullableExternalNrfFunctionSingleAllOf {
	return &NullableExternalNrfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalNrfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalNrfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


