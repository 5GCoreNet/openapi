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

// checks if the ExternalSeppFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalSeppFunctionSingleAllOf{}

// ExternalSeppFunctionSingleAllOf struct for ExternalSeppFunctionSingleAllOf
type ExternalSeppFunctionSingleAllOf struct {
	Attributes *ExternalSeppFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalSeppFunctionSingleAllOf instantiates a new ExternalSeppFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSeppFunctionSingleAllOf() *ExternalSeppFunctionSingleAllOf {
	this := ExternalSeppFunctionSingleAllOf{}
	return &this
}

// NewExternalSeppFunctionSingleAllOfWithDefaults instantiates a new ExternalSeppFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSeppFunctionSingleAllOfWithDefaults() *ExternalSeppFunctionSingleAllOf {
	this := ExternalSeppFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalSeppFunctionSingleAllOf) GetAttributes() ExternalSeppFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalSeppFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSeppFunctionSingleAllOf) GetAttributesOk() (*ExternalSeppFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalSeppFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalSeppFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalSeppFunctionSingleAllOf) SetAttributes(v ExternalSeppFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalSeppFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalSeppFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalSeppFunctionSingleAllOf struct {
	value *ExternalSeppFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalSeppFunctionSingleAllOf) Get() *ExternalSeppFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalSeppFunctionSingleAllOf) Set(val *ExternalSeppFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSeppFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSeppFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSeppFunctionSingleAllOf(val *ExternalSeppFunctionSingleAllOf) *NullableExternalSeppFunctionSingleAllOf {
	return &NullableExternalSeppFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalSeppFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSeppFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
