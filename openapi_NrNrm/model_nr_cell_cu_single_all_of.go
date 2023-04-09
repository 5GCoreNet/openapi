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

// checks if the NrCellCuSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrCellCuSingleAllOf{}

// NrCellCuSingleAllOf struct for NrCellCuSingleAllOf
type NrCellCuSingleAllOf struct {
	Attributes *NrCellCuSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNrCellCuSingleAllOf instantiates a new NrCellCuSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellCuSingleAllOf() *NrCellCuSingleAllOf {
	this := NrCellCuSingleAllOf{}
	return &this
}

// NewNrCellCuSingleAllOfWithDefaults instantiates a new NrCellCuSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellCuSingleAllOfWithDefaults() *NrCellCuSingleAllOf {
	this := NrCellCuSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf) GetAttributes() NrCellCuSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NrCellCuSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf) GetAttributesOk() (*NrCellCuSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NrCellCuSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NrCellCuSingleAllOf) SetAttributes(v NrCellCuSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NrCellCuSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrCellCuSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableNrCellCuSingleAllOf struct {
	value *NrCellCuSingleAllOf
	isSet bool
}

func (v NullableNrCellCuSingleAllOf) Get() *NrCellCuSingleAllOf {
	return v.value
}

func (v *NullableNrCellCuSingleAllOf) Set(val *NrCellCuSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellCuSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellCuSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellCuSingleAllOf(val *NrCellCuSingleAllOf) *NullableNrCellCuSingleAllOf {
	return &NullableNrCellCuSingleAllOf{value: val, isSet: true}
}

func (v NullableNrCellCuSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellCuSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
