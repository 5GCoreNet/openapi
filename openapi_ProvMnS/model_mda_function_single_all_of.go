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

// checks if the MDAFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAFunctionSingleAllOf{}

// MDAFunctionSingleAllOf struct for MDAFunctionSingleAllOf
type MDAFunctionSingleAllOf struct {
	Attributes *MDAFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewMDAFunctionSingleAllOf instantiates a new MDAFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAFunctionSingleAllOf() *MDAFunctionSingleAllOf {
	this := MDAFunctionSingleAllOf{}
	return &this
}

// NewMDAFunctionSingleAllOfWithDefaults instantiates a new MDAFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAFunctionSingleAllOfWithDefaults() *MDAFunctionSingleAllOf {
	this := MDAFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MDAFunctionSingleAllOf) GetAttributes() MDAFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret MDAFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAFunctionSingleAllOf) GetAttributesOk() (*MDAFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MDAFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MDAFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *MDAFunctionSingleAllOf) SetAttributes(v MDAFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o MDAFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMDAFunctionSingleAllOf struct {
	value *MDAFunctionSingleAllOf
	isSet bool
}

func (v NullableMDAFunctionSingleAllOf) Get() *MDAFunctionSingleAllOf {
	return v.value
}

func (v *NullableMDAFunctionSingleAllOf) Set(val *MDAFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAFunctionSingleAllOf(val *MDAFunctionSingleAllOf) *NullableMDAFunctionSingleAllOf {
	return &NullableMDAFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableMDAFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


