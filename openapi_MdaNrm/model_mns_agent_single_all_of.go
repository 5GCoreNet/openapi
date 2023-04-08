/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the MnsAgentSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MnsAgentSingleAllOf{}

// MnsAgentSingleAllOf struct for MnsAgentSingleAllOf
type MnsAgentSingleAllOf struct {
	Attributes *MnsAgentSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewMnsAgentSingleAllOf instantiates a new MnsAgentSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnsAgentSingleAllOf() *MnsAgentSingleAllOf {
	this := MnsAgentSingleAllOf{}
	return &this
}

// NewMnsAgentSingleAllOfWithDefaults instantiates a new MnsAgentSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnsAgentSingleAllOfWithDefaults() *MnsAgentSingleAllOf {
	this := MnsAgentSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MnsAgentSingleAllOf) GetAttributes() MnsAgentSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret MnsAgentSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MnsAgentSingleAllOf) GetAttributesOk() (*MnsAgentSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MnsAgentSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MnsAgentSingleAllOfAttributes and assigns it to the Attributes field.
func (o *MnsAgentSingleAllOf) SetAttributes(v MnsAgentSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o MnsAgentSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MnsAgentSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMnsAgentSingleAllOf struct {
	value *MnsAgentSingleAllOf
	isSet bool
}

func (v NullableMnsAgentSingleAllOf) Get() *MnsAgentSingleAllOf {
	return v.value
}

func (v *NullableMnsAgentSingleAllOf) Set(val *MnsAgentSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMnsAgentSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMnsAgentSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnsAgentSingleAllOf(val *MnsAgentSingleAllOf) *NullableMnsAgentSingleAllOf {
	return &NullableMnsAgentSingleAllOf{value: val, isSet: true}
}

func (v NullableMnsAgentSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnsAgentSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


