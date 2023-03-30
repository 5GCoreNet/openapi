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

// checks if the BwpSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BwpSingleAllOf{}

// BwpSingleAllOf struct for BwpSingleAllOf
type BwpSingleAllOf struct {
	Attributes *BwpSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewBwpSingleAllOf instantiates a new BwpSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBwpSingleAllOf() *BwpSingleAllOf {
	this := BwpSingleAllOf{}
	return &this
}

// NewBwpSingleAllOfWithDefaults instantiates a new BwpSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBwpSingleAllOfWithDefaults() *BwpSingleAllOf {
	this := BwpSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BwpSingleAllOf) GetAttributes() BwpSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BwpSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwpSingleAllOf) GetAttributesOk() (*BwpSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BwpSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BwpSingleAllOfAttributes and assigns it to the Attributes field.
func (o *BwpSingleAllOf) SetAttributes(v BwpSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o BwpSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BwpSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBwpSingleAllOf struct {
	value *BwpSingleAllOf
	isSet bool
}

func (v NullableBwpSingleAllOf) Get() *BwpSingleAllOf {
	return v.value
}

func (v *NullableBwpSingleAllOf) Set(val *BwpSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBwpSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBwpSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBwpSingleAllOf(val *BwpSingleAllOf) *NullableBwpSingleAllOf {
	return &NullableBwpSingleAllOf{value: val, isSet: true}
}

func (v NullableBwpSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBwpSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


