/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the NRFreqRelationSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NRFreqRelationSingle{}

// NRFreqRelationSingle struct for NRFreqRelationSingle
type NRFreqRelationSingle struct {
	Top
	Attributes *NRFreqRelationSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewNRFreqRelationSingle instantiates a new NRFreqRelationSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNRFreqRelationSingle(id NullableString) *NRFreqRelationSingle {
	this := NRFreqRelationSingle{}
	this.Id = id
	return &this
}

// NewNRFreqRelationSingleWithDefaults instantiates a new NRFreqRelationSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNRFreqRelationSingleWithDefaults() *NRFreqRelationSingle {
	this := NRFreqRelationSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NRFreqRelationSingle) GetAttributes() NRFreqRelationSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NRFreqRelationSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRFreqRelationSingle) GetAttributesOk() (*NRFreqRelationSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NRFreqRelationSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NRFreqRelationSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NRFreqRelationSingle) SetAttributes(v NRFreqRelationSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o NRFreqRelationSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NRFreqRelationSingle) ToMap() (map[string]interface{}, error) {
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

type NullableNRFreqRelationSingle struct {
	value *NRFreqRelationSingle
	isSet bool
}

func (v NullableNRFreqRelationSingle) Get() *NRFreqRelationSingle {
	return v.value
}

func (v *NullableNRFreqRelationSingle) Set(val *NRFreqRelationSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNRFreqRelationSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNRFreqRelationSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNRFreqRelationSingle(val *NRFreqRelationSingle) *NullableNRFreqRelationSingle {
	return &NullableNRFreqRelationSingle{value: val, isSet: true}
}

func (v NullableNRFreqRelationSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNRFreqRelationSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
