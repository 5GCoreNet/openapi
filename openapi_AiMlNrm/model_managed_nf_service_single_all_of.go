/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// checks if the ManagedNFServiceSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedNFServiceSingleAllOf{}

// ManagedNFServiceSingleAllOf struct for ManagedNFServiceSingleAllOf
type ManagedNFServiceSingleAllOf struct {
	Attributes *ManagedNFServiceSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewManagedNFServiceSingleAllOf instantiates a new ManagedNFServiceSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedNFServiceSingleAllOf() *ManagedNFServiceSingleAllOf {
	this := ManagedNFServiceSingleAllOf{}
	return &this
}

// NewManagedNFServiceSingleAllOfWithDefaults instantiates a new ManagedNFServiceSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedNFServiceSingleAllOfWithDefaults() *ManagedNFServiceSingleAllOf {
	this := ManagedNFServiceSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOf) GetAttributes() ManagedNFServiceSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagedNFServiceSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOf) GetAttributesOk() (*ManagedNFServiceSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedNFServiceSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagedNFServiceSingleAllOf) SetAttributes(v ManagedNFServiceSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ManagedNFServiceSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedNFServiceSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableManagedNFServiceSingleAllOf struct {
	value *ManagedNFServiceSingleAllOf
	isSet bool
}

func (v NullableManagedNFServiceSingleAllOf) Get() *ManagedNFServiceSingleAllOf {
	return v.value
}

func (v *NullableManagedNFServiceSingleAllOf) Set(val *ManagedNFServiceSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedNFServiceSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedNFServiceSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedNFServiceSingleAllOf(val *ManagedNFServiceSingleAllOf) *NullableManagedNFServiceSingleAllOf {
	return &NullableManagedNFServiceSingleAllOf{value: val, isSet: true}
}

func (v NullableManagedNFServiceSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedNFServiceSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
