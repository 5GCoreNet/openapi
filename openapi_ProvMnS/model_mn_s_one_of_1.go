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

// checks if the MnSOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MnSOneOf1{}

// MnSOneOf1 struct for MnSOneOf1
type MnSOneOf1 struct {
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
}

// NewMnSOneOf1 instantiates a new MnSOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnSOneOf1() *MnSOneOf1 {
	this := MnSOneOf1{}
	return &this
}

// NewMnSOneOf1WithDefaults instantiates a new MnSOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnSOneOf1WithDefaults() *MnSOneOf1 {
	this := MnSOneOf1{}
	return &this
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *MnSOneOf1) GetManagedElement() []ManagedElementSingle {
	if o == nil || IsNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnSOneOf1) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || IsNil(o.ManagedElement) {
		return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *MnSOneOf1) HasManagedElement() bool {
	if o != nil && !IsNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *MnSOneOf1) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

func (o MnSOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MnSOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	return toSerialize, nil
}

type NullableMnSOneOf1 struct {
	value *MnSOneOf1
	isSet bool
}

func (v NullableMnSOneOf1) Get() *MnSOneOf1 {
	return v.value
}

func (v *NullableMnSOneOf1) Set(val *MnSOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableMnSOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableMnSOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnSOneOf1(val *MnSOneOf1) *NullableMnSOneOf1 {
	return &NullableMnSOneOf1{value: val, isSet: true}
}

func (v NullableMnSOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnSOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


