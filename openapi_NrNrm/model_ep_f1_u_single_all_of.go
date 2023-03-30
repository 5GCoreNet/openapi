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

// checks if the EPF1USingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPF1USingleAllOf{}

// EPF1USingleAllOf struct for EPF1USingleAllOf
type EPF1USingleAllOf struct {
	Attributes *EPF1USingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPF1USingleAllOf instantiates a new EPF1USingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPF1USingleAllOf() *EPF1USingleAllOf {
	this := EPF1USingleAllOf{}
	return &this
}

// NewEPF1USingleAllOfWithDefaults instantiates a new EPF1USingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPF1USingleAllOfWithDefaults() *EPF1USingleAllOf {
	this := EPF1USingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPF1USingleAllOf) GetAttributes() EPF1USingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPF1USingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOf) GetAttributesOk() (*EPF1USingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPF1USingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPF1USingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPF1USingleAllOf) SetAttributes(v EPF1USingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPF1USingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPF1USingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPF1USingleAllOf struct {
	value *EPF1USingleAllOf
	isSet bool
}

func (v NullableEPF1USingleAllOf) Get() *EPF1USingleAllOf {
	return v.value
}

func (v *NullableEPF1USingleAllOf) Set(val *EPF1USingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPF1USingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPF1USingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPF1USingleAllOf(val *EPF1USingleAllOf) *NullableEPF1USingleAllOf {
	return &NullableEPF1USingleAllOf{value: val, isSet: true}
}

func (v NullableEPF1USingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPF1USingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

