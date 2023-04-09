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

// checks if the ProvMnSOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvMnSOneOf1{}

// ProvMnSOneOf1 struct for ProvMnSOneOf1
type ProvMnSOneOf1 struct {
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
}

// NewProvMnSOneOf1 instantiates a new ProvMnSOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvMnSOneOf1() *ProvMnSOneOf1 {
	this := ProvMnSOneOf1{}
	return &this
}

// NewProvMnSOneOf1WithDefaults instantiates a new ProvMnSOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvMnSOneOf1WithDefaults() *ProvMnSOneOf1 {
	this := ProvMnSOneOf1{}
	return &this
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *ProvMnSOneOf1) GetManagedElement() []ManagedElementSingle {
	if o == nil || IsNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvMnSOneOf1) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || IsNil(o.ManagedElement) {
		return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *ProvMnSOneOf1) HasManagedElement() bool {
	if o != nil && !IsNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *ProvMnSOneOf1) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

func (o ProvMnSOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvMnSOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	return toSerialize, nil
}

type NullableProvMnSOneOf1 struct {
	value *ProvMnSOneOf1
	isSet bool
}

func (v NullableProvMnSOneOf1) Get() *ProvMnSOneOf1 {
	return v.value
}

func (v *NullableProvMnSOneOf1) Set(val *ProvMnSOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableProvMnSOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableProvMnSOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvMnSOneOf1(val *ProvMnSOneOf1) *NullableProvMnSOneOf1 {
	return &NullableProvMnSOneOf1{value: val, isSet: true}
}

func (v NullableProvMnSOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvMnSOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
