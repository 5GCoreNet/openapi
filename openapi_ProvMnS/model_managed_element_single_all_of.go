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

// checks if the ManagedElementSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementSingleAllOf{}

// ManagedElementSingleAllOf struct for ManagedElementSingleAllOf
type ManagedElementSingleAllOf struct {
	Attributes *ManagedElementAttr `json:"attributes,omitempty"`
}

// NewManagedElementSingleAllOf instantiates a new ManagedElementSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingleAllOf() *ManagedElementSingleAllOf {
	this := ManagedElementSingleAllOf{}
	return &this
}

// NewManagedElementSingleAllOfWithDefaults instantiates a new ManagedElementSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingleAllOfWithDefaults() *ManagedElementSingleAllOf {
	this := ManagedElementSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf) GetAttributes() ManagedElementAttr {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedElementAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf) GetAttributesOk() (*ManagedElementAttr, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementAttr and assigns it to the Attributes field.
func (o *ManagedElementSingleAllOf) SetAttributes(v ManagedElementAttr) {
	o.Attributes = &v
}

func (o ManagedElementSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagedElementSingleAllOf struct {
	value *ManagedElementSingleAllOf
	isSet bool
}

func (v NullableManagedElementSingleAllOf) Get() *ManagedElementSingleAllOf {
	return v.value
}

func (v *NullableManagedElementSingleAllOf) Set(val *ManagedElementSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingleAllOf(val *ManagedElementSingleAllOf) *NullableManagedElementSingleAllOf {
	return &NullableManagedElementSingleAllOf{value: val, isSet: true}
}

func (v NullableManagedElementSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
