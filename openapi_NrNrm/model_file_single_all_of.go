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

// checks if the FileSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSingleAllOf{}

// FileSingleAllOf struct for FileSingleAllOf
type FileSingleAllOf struct {
	Attributes *FileSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFileSingleAllOf instantiates a new FileSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSingleAllOf() *FileSingleAllOf {
	this := FileSingleAllOf{}
	return &this
}

// NewFileSingleAllOfWithDefaults instantiates a new FileSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSingleAllOfWithDefaults() *FileSingleAllOf {
	this := FileSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FileSingleAllOf) GetAttributes() FileSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret FileSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOf) GetAttributesOk() (*FileSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FileSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FileSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FileSingleAllOf) SetAttributes(v FileSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FileSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableFileSingleAllOf struct {
	value *FileSingleAllOf
	isSet bool
}

func (v NullableFileSingleAllOf) Get() *FileSingleAllOf {
	return v.value
}

func (v *NullableFileSingleAllOf) Set(val *FileSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSingleAllOf(val *FileSingleAllOf) *NullableFileSingleAllOf {
	return &NullableFileSingleAllOf{value: val, isSet: true}
}

func (v NullableFileSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
