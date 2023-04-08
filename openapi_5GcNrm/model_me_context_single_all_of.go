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

// checks if the MeContextSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeContextSingleAllOf{}

// MeContextSingleAllOf struct for MeContextSingleAllOf
type MeContextSingleAllOf struct {
	Attributes *MeContextSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewMeContextSingleAllOf instantiates a new MeContextSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeContextSingleAllOf() *MeContextSingleAllOf {
	this := MeContextSingleAllOf{}
	return &this
}

// NewMeContextSingleAllOfWithDefaults instantiates a new MeContextSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeContextSingleAllOfWithDefaults() *MeContextSingleAllOf {
	this := MeContextSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MeContextSingleAllOf) GetAttributes() MeContextSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret MeContextSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeContextSingleAllOf) GetAttributesOk() (*MeContextSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MeContextSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MeContextSingleAllOfAttributes and assigns it to the Attributes field.
func (o *MeContextSingleAllOf) SetAttributes(v MeContextSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o MeContextSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeContextSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMeContextSingleAllOf struct {
	value *MeContextSingleAllOf
	isSet bool
}

func (v NullableMeContextSingleAllOf) Get() *MeContextSingleAllOf {
	return v.value
}

func (v *NullableMeContextSingleAllOf) Set(val *MeContextSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMeContextSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMeContextSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeContextSingleAllOf(val *MeContextSingleAllOf) *NullableMeContextSingleAllOf {
	return &NullableMeContextSingleAllOf{value: val, isSet: true}
}

func (v NullableMeContextSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeContextSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


