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

// checks if the RimRSSetSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RimRSSetSingleAllOf{}

// RimRSSetSingleAllOf struct for RimRSSetSingleAllOf
type RimRSSetSingleAllOf struct {
	Attributes *RimRSSetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewRimRSSetSingleAllOf instantiates a new RimRSSetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSSetSingleAllOf() *RimRSSetSingleAllOf {
	this := RimRSSetSingleAllOf{}
	return &this
}

// NewRimRSSetSingleAllOfWithDefaults instantiates a new RimRSSetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSSetSingleAllOfWithDefaults() *RimRSSetSingleAllOf {
	this := RimRSSetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RimRSSetSingleAllOf) GetAttributes() RimRSSetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret RimRSSetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSSetSingleAllOf) GetAttributesOk() (*RimRSSetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RimRSSetSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given RimRSSetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *RimRSSetSingleAllOf) SetAttributes(v RimRSSetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o RimRSSetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RimRSSetSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableRimRSSetSingleAllOf struct {
	value *RimRSSetSingleAllOf
	isSet bool
}

func (v NullableRimRSSetSingleAllOf) Get() *RimRSSetSingleAllOf {
	return v.value
}

func (v *NullableRimRSSetSingleAllOf) Set(val *RimRSSetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSSetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSSetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSSetSingleAllOf(val *RimRSSetSingleAllOf) *NullableRimRSSetSingleAllOf {
	return &NullableRimRSSetSingleAllOf{value: val, isSet: true}
}

func (v NullableRimRSSetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSSetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
