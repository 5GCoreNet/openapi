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

// checks if the ExternalNrCellCuSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalNrCellCuSingleAllOf{}

// ExternalNrCellCuSingleAllOf struct for ExternalNrCellCuSingleAllOf
type ExternalNrCellCuSingleAllOf struct {
	Attributes *ExternalNrCellCuSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalNrCellCuSingleAllOf instantiates a new ExternalNrCellCuSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalNrCellCuSingleAllOf() *ExternalNrCellCuSingleAllOf {
	this := ExternalNrCellCuSingleAllOf{}
	return &this
}

// NewExternalNrCellCuSingleAllOfWithDefaults instantiates a new ExternalNrCellCuSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalNrCellCuSingleAllOfWithDefaults() *ExternalNrCellCuSingleAllOf {
	this := ExternalNrCellCuSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalNrCellCuSingleAllOf) GetAttributes() ExternalNrCellCuSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalNrCellCuSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalNrCellCuSingleAllOf) GetAttributesOk() (*ExternalNrCellCuSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalNrCellCuSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalNrCellCuSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalNrCellCuSingleAllOf) SetAttributes(v ExternalNrCellCuSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalNrCellCuSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalNrCellCuSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalNrCellCuSingleAllOf struct {
	value *ExternalNrCellCuSingleAllOf
	isSet bool
}

func (v NullableExternalNrCellCuSingleAllOf) Get() *ExternalNrCellCuSingleAllOf {
	return v.value
}

func (v *NullableExternalNrCellCuSingleAllOf) Set(val *ExternalNrCellCuSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalNrCellCuSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalNrCellCuSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalNrCellCuSingleAllOf(val *ExternalNrCellCuSingleAllOf) *NullableExternalNrCellCuSingleAllOf {
	return &NullableExternalNrCellCuSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalNrCellCuSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalNrCellCuSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
