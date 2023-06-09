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

// checks if the ManagedElementSingle1AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementSingle1AllOf{}

// ManagedElementSingle1AllOf struct for ManagedElementSingle1AllOf
type ManagedElementSingle1AllOf struct {
	Attributes *ManagedElementSingle1AllOfAttributes `json:"attributes,omitempty"`
}

// NewManagedElementSingle1AllOf instantiates a new ManagedElementSingle1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingle1AllOf() *ManagedElementSingle1AllOf {
	this := ManagedElementSingle1AllOf{}
	return &this
}

// NewManagedElementSingle1AllOfWithDefaults instantiates a new ManagedElementSingle1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingle1AllOfWithDefaults() *ManagedElementSingle1AllOf {
	this := ManagedElementSingle1AllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingle1AllOf) GetAttributes() ManagedElementSingle1AllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedElementSingle1AllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle1AllOf) GetAttributesOk() (*ManagedElementSingle1AllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingle1AllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementSingle1AllOfAttributes and assigns it to the Attributes field.
func (o *ManagedElementSingle1AllOf) SetAttributes(v ManagedElementSingle1AllOfAttributes) {
	o.Attributes = &v
}

func (o ManagedElementSingle1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementSingle1AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagedElementSingle1AllOf struct {
	value *ManagedElementSingle1AllOf
	isSet bool
}

func (v NullableManagedElementSingle1AllOf) Get() *ManagedElementSingle1AllOf {
	return v.value
}

func (v *NullableManagedElementSingle1AllOf) Set(val *ManagedElementSingle1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingle1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingle1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingle1AllOf(val *ManagedElementSingle1AllOf) *NullableManagedElementSingle1AllOf {
	return &NullableManagedElementSingle1AllOf{value: val, isSet: true}
}

func (v NullableManagedElementSingle1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingle1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
