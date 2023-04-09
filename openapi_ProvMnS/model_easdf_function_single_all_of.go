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

// checks if the EASDFFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EASDFFunctionSingleAllOf{}

// EASDFFunctionSingleAllOf struct for EASDFFunctionSingleAllOf
type EASDFFunctionSingleAllOf struct {
	Attributes *EASDFFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEASDFFunctionSingleAllOf instantiates a new EASDFFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASDFFunctionSingleAllOf() *EASDFFunctionSingleAllOf {
	this := EASDFFunctionSingleAllOf{}
	return &this
}

// NewEASDFFunctionSingleAllOfWithDefaults instantiates a new EASDFFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASDFFunctionSingleAllOfWithDefaults() *EASDFFunctionSingleAllOf {
	this := EASDFFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOf) GetAttributes() EASDFFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EASDFFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOf) GetAttributesOk() (*EASDFFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EASDFFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EASDFFunctionSingleAllOf) SetAttributes(v EASDFFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EASDFFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EASDFFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEASDFFunctionSingleAllOf struct {
	value *EASDFFunctionSingleAllOf
	isSet bool
}

func (v NullableEASDFFunctionSingleAllOf) Get() *EASDFFunctionSingleAllOf {
	return v.value
}

func (v *NullableEASDFFunctionSingleAllOf) Set(val *EASDFFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEASDFFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEASDFFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASDFFunctionSingleAllOf(val *EASDFFunctionSingleAllOf) *NullableEASDFFunctionSingleAllOf {
	return &NullableEASDFFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableEASDFFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASDFFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
