/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the MnsAgentSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MnsAgentSingle{}

// MnsAgentSingle struct for MnsAgentSingle
type MnsAgentSingle struct {
	Top
	Attributes *MnsAgentSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewMnsAgentSingle instantiates a new MnsAgentSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnsAgentSingle(id NullableString) *MnsAgentSingle {
	this := MnsAgentSingle{}
	this.Id = id
	return &this
}

// NewMnsAgentSingleWithDefaults instantiates a new MnsAgentSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnsAgentSingleWithDefaults() *MnsAgentSingle {
	this := MnsAgentSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MnsAgentSingle) GetAttributes() MnsAgentSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret MnsAgentSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsAgentSingle) GetAttributesOk() (*MnsAgentSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MnsAgentSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MnsAgentSingleAllOfAttributes and assigns it to the Attributes field.
func (o *MnsAgentSingle) SetAttributes(v MnsAgentSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o MnsAgentSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MnsAgentSingle) ToMap() (map[string]interface{}, error) {
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

type NullableMnsAgentSingle struct {
	value *MnsAgentSingle
	isSet bool
}

func (v NullableMnsAgentSingle) Get() *MnsAgentSingle {
	return v.value
}

func (v *NullableMnsAgentSingle) Set(val *MnsAgentSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableMnsAgentSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableMnsAgentSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnsAgentSingle(val *MnsAgentSingle) *NullableMnsAgentSingle {
	return &NullableMnsAgentSingle{value: val, isSet: true}
}

func (v NullableMnsAgentSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnsAgentSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
