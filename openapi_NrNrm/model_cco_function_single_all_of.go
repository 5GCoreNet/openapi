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

// checks if the CCOFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CCOFunctionSingleAllOf{}

// CCOFunctionSingleAllOf struct for CCOFunctionSingleAllOf
type CCOFunctionSingleAllOf struct {
	Attributes *CCOFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewCCOFunctionSingleAllOf instantiates a new CCOFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOFunctionSingleAllOf() *CCOFunctionSingleAllOf {
	this := CCOFunctionSingleAllOf{}
	return &this
}

// NewCCOFunctionSingleAllOfWithDefaults instantiates a new CCOFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOFunctionSingleAllOfWithDefaults() *CCOFunctionSingleAllOf {
	this := CCOFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CCOFunctionSingleAllOf) GetAttributes() CCOFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CCOFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOFunctionSingleAllOf) GetAttributesOk() (*CCOFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CCOFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CCOFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *CCOFunctionSingleAllOf) SetAttributes(v CCOFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o CCOFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CCOFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableCCOFunctionSingleAllOf struct {
	value *CCOFunctionSingleAllOf
	isSet bool
}

func (v NullableCCOFunctionSingleAllOf) Get() *CCOFunctionSingleAllOf {
	return v.value
}

func (v *NullableCCOFunctionSingleAllOf) Set(val *CCOFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOFunctionSingleAllOf(val *CCOFunctionSingleAllOf) *NullableCCOFunctionSingleAllOf {
	return &NullableCCOFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableCCOFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
