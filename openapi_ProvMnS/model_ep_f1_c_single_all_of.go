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

// checks if the EPF1CSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPF1CSingleAllOf{}

// EPF1CSingleAllOf struct for EPF1CSingleAllOf
type EPF1CSingleAllOf struct {
	Attributes *EPF1CSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPF1CSingleAllOf instantiates a new EPF1CSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPF1CSingleAllOf() *EPF1CSingleAllOf {
	this := EPF1CSingleAllOf{}
	return &this
}

// NewEPF1CSingleAllOfWithDefaults instantiates a new EPF1CSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPF1CSingleAllOfWithDefaults() *EPF1CSingleAllOf {
	this := EPF1CSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPF1CSingleAllOf) GetAttributes() EPF1CSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPF1CSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1CSingleAllOf) GetAttributesOk() (*EPF1CSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPF1CSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPF1CSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPF1CSingleAllOf) SetAttributes(v EPF1CSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPF1CSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPF1CSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPF1CSingleAllOf struct {
	value *EPF1CSingleAllOf
	isSet bool
}

func (v NullableEPF1CSingleAllOf) Get() *EPF1CSingleAllOf {
	return v.value
}

func (v *NullableEPF1CSingleAllOf) Set(val *EPF1CSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPF1CSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPF1CSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPF1CSingleAllOf(val *EPF1CSingleAllOf) *NullableEPF1CSingleAllOf {
	return &NullableEPF1CSingleAllOf{value: val, isSet: true}
}

func (v NullableEPF1CSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPF1CSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


