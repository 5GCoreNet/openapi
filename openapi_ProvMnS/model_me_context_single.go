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

// checks if the MeContextSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeContextSingle{}

// MeContextSingle struct for MeContextSingle
type MeContextSingle struct {
	Top
	Attributes *MeContextSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewMeContextSingle instantiates a new MeContextSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeContextSingle(id NullableString) *MeContextSingle {
	this := MeContextSingle{}
	this.Id = id
	return &this
}

// NewMeContextSingleWithDefaults instantiates a new MeContextSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeContextSingleWithDefaults() *MeContextSingle {
	this := MeContextSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MeContextSingle) GetAttributes() MeContextSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret MeContextSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeContextSingle) GetAttributesOk() (*MeContextSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MeContextSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MeContextSingleAllOfAttributes and assigns it to the Attributes field.
func (o *MeContextSingle) SetAttributes(v MeContextSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o MeContextSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeContextSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMeContextSingle struct {
	value *MeContextSingle
	isSet bool
}

func (v NullableMeContextSingle) Get() *MeContextSingle {
	return v.value
}

func (v *NullableMeContextSingle) Set(val *MeContextSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableMeContextSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableMeContextSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeContextSingle(val *MeContextSingle) *NullableMeContextSingle {
	return &NullableMeContextSingle{value: val, isSet: true}
}

func (v NullableMeContextSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeContextSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
