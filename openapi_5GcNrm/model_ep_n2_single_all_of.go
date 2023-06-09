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

// checks if the EPN2SingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN2SingleAllOf{}

// EPN2SingleAllOf struct for EPN2SingleAllOf
type EPN2SingleAllOf struct {
	Attributes *EPN2SingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPN2SingleAllOf instantiates a new EPN2SingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN2SingleAllOf() *EPN2SingleAllOf {
	this := EPN2SingleAllOf{}
	return &this
}

// NewEPN2SingleAllOfWithDefaults instantiates a new EPN2SingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN2SingleAllOfWithDefaults() *EPN2SingleAllOf {
	this := EPN2SingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN2SingleAllOf) GetAttributes() EPN2SingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPN2SingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN2SingleAllOf) GetAttributesOk() (*EPN2SingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN2SingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPN2SingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPN2SingleAllOf) SetAttributes(v EPN2SingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPN2SingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN2SingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPN2SingleAllOf struct {
	value *EPN2SingleAllOf
	isSet bool
}

func (v NullableEPN2SingleAllOf) Get() *EPN2SingleAllOf {
	return v.value
}

func (v *NullableEPN2SingleAllOf) Set(val *EPN2SingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN2SingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN2SingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN2SingleAllOf(val *EPN2SingleAllOf) *NullableEPN2SingleAllOf {
	return &NullableEPN2SingleAllOf{value: val, isSet: true}
}

func (v NullableEPN2SingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN2SingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
