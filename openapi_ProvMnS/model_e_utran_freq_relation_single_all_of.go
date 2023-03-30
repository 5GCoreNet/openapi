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

// checks if the EUtranFreqRelationSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EUtranFreqRelationSingleAllOf{}

// EUtranFreqRelationSingleAllOf struct for EUtranFreqRelationSingleAllOf
type EUtranFreqRelationSingleAllOf struct {
	Attributes *EUtranFreqRelationSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEUtranFreqRelationSingleAllOf instantiates a new EUtranFreqRelationSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEUtranFreqRelationSingleAllOf() *EUtranFreqRelationSingleAllOf {
	this := EUtranFreqRelationSingleAllOf{}
	return &this
}

// NewEUtranFreqRelationSingleAllOfWithDefaults instantiates a new EUtranFreqRelationSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEUtranFreqRelationSingleAllOfWithDefaults() *EUtranFreqRelationSingleAllOf {
	this := EUtranFreqRelationSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EUtranFreqRelationSingleAllOf) GetAttributes() EUtranFreqRelationSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EUtranFreqRelationSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFreqRelationSingleAllOf) GetAttributesOk() (*EUtranFreqRelationSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EUtranFreqRelationSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EUtranFreqRelationSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EUtranFreqRelationSingleAllOf) SetAttributes(v EUtranFreqRelationSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EUtranFreqRelationSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EUtranFreqRelationSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEUtranFreqRelationSingleAllOf struct {
	value *EUtranFreqRelationSingleAllOf
	isSet bool
}

func (v NullableEUtranFreqRelationSingleAllOf) Get() *EUtranFreqRelationSingleAllOf {
	return v.value
}

func (v *NullableEUtranFreqRelationSingleAllOf) Set(val *EUtranFreqRelationSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEUtranFreqRelationSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEUtranFreqRelationSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEUtranFreqRelationSingleAllOf(val *EUtranFreqRelationSingleAllOf) *NullableEUtranFreqRelationSingleAllOf {
	return &NullableEUtranFreqRelationSingleAllOf{value: val, isSet: true}
}

func (v NullableEUtranFreqRelationSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEUtranFreqRelationSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


