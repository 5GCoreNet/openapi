/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the SubNetworkSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubNetworkSingleAllOf{}

// SubNetworkSingleAllOf struct for SubNetworkSingleAllOf
type SubNetworkSingleAllOf struct {
	Attributes *SubNetworkSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewSubNetworkSingleAllOf instantiates a new SubNetworkSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingleAllOf() *SubNetworkSingleAllOf {
	this := SubNetworkSingleAllOf{}
	return &this
}

// NewSubNetworkSingleAllOfWithDefaults instantiates a new SubNetworkSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleAllOfWithDefaults() *SubNetworkSingleAllOf {
	this := SubNetworkSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubNetworkSingleAllOf) GetAttributes() SubNetworkSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubNetworkSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingleAllOf) GetAttributesOk() (*SubNetworkSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubNetworkSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubNetworkSingleAllOfAttributes and assigns it to the Attributes field.
func (o *SubNetworkSingleAllOf) SetAttributes(v SubNetworkSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o SubNetworkSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubNetworkSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSubNetworkSingleAllOf struct {
	value *SubNetworkSingleAllOf
	isSet bool
}

func (v NullableSubNetworkSingleAllOf) Get() *SubNetworkSingleAllOf {
	return v.value
}

func (v *NullableSubNetworkSingleAllOf) Set(val *SubNetworkSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingleAllOf(val *SubNetworkSingleAllOf) *NullableSubNetworkSingleAllOf {
	return &NullableSubNetworkSingleAllOf{value: val, isSet: true}
}

func (v NullableSubNetworkSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


