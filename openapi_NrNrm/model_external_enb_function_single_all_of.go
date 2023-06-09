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

// checks if the ExternalENBFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalENBFunctionSingleAllOf{}

// ExternalENBFunctionSingleAllOf struct for ExternalENBFunctionSingleAllOf
type ExternalENBFunctionSingleAllOf struct {
	Attributes *ExternalENBFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalENBFunctionSingleAllOf instantiates a new ExternalENBFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalENBFunctionSingleAllOf() *ExternalENBFunctionSingleAllOf {
	this := ExternalENBFunctionSingleAllOf{}
	return &this
}

// NewExternalENBFunctionSingleAllOfWithDefaults instantiates a new ExternalENBFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalENBFunctionSingleAllOfWithDefaults() *ExternalENBFunctionSingleAllOf {
	this := ExternalENBFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalENBFunctionSingleAllOf) GetAttributes() ExternalENBFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalENBFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalENBFunctionSingleAllOf) GetAttributesOk() (*ExternalENBFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalENBFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalENBFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalENBFunctionSingleAllOf) SetAttributes(v ExternalENBFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalENBFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalENBFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalENBFunctionSingleAllOf struct {
	value *ExternalENBFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalENBFunctionSingleAllOf) Get() *ExternalENBFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalENBFunctionSingleAllOf) Set(val *ExternalENBFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalENBFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalENBFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalENBFunctionSingleAllOf(val *ExternalENBFunctionSingleAllOf) *NullableExternalENBFunctionSingleAllOf {
	return &NullableExternalENBFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalENBFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalENBFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
