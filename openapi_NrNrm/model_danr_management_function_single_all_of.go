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

// checks if the DANRManagementFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DANRManagementFunctionSingleAllOf{}

// DANRManagementFunctionSingleAllOf struct for DANRManagementFunctionSingleAllOf
type DANRManagementFunctionSingleAllOf struct {
	Attributes *DANRManagementFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDANRManagementFunctionSingleAllOf instantiates a new DANRManagementFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDANRManagementFunctionSingleAllOf() *DANRManagementFunctionSingleAllOf {
	this := DANRManagementFunctionSingleAllOf{}
	return &this
}

// NewDANRManagementFunctionSingleAllOfWithDefaults instantiates a new DANRManagementFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDANRManagementFunctionSingleAllOfWithDefaults() *DANRManagementFunctionSingleAllOf {
	this := DANRManagementFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DANRManagementFunctionSingleAllOf) GetAttributes() DANRManagementFunctionSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret DANRManagementFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DANRManagementFunctionSingleAllOf) GetAttributesOk() (*DANRManagementFunctionSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DANRManagementFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DANRManagementFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *DANRManagementFunctionSingleAllOf) SetAttributes(v DANRManagementFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o DANRManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DANRManagementFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDANRManagementFunctionSingleAllOf struct {
	value *DANRManagementFunctionSingleAllOf
	isSet bool
}

func (v NullableDANRManagementFunctionSingleAllOf) Get() *DANRManagementFunctionSingleAllOf {
	return v.value
}

func (v *NullableDANRManagementFunctionSingleAllOf) Set(val *DANRManagementFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDANRManagementFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDANRManagementFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDANRManagementFunctionSingleAllOf(val *DANRManagementFunctionSingleAllOf) *NullableDANRManagementFunctionSingleAllOf {
	return &NullableDANRManagementFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableDANRManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDANRManagementFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


