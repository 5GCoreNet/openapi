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

// checks if the PredefinedPccRuleSetSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPccRuleSetSingleAllOf{}

// PredefinedPccRuleSetSingleAllOf struct for PredefinedPccRuleSetSingleAllOf
type PredefinedPccRuleSetSingleAllOf struct {
	Attributes *PredefinedPccRuleSetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewPredefinedPccRuleSetSingleAllOf instantiates a new PredefinedPccRuleSetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPccRuleSetSingleAllOf() *PredefinedPccRuleSetSingleAllOf {
	this := PredefinedPccRuleSetSingleAllOf{}
	return &this
}

// NewPredefinedPccRuleSetSingleAllOfWithDefaults instantiates a new PredefinedPccRuleSetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPccRuleSetSingleAllOfWithDefaults() *PredefinedPccRuleSetSingleAllOf {
	this := PredefinedPccRuleSetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingleAllOf) GetAttributes() PredefinedPccRuleSetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PredefinedPccRuleSetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingleAllOf) GetAttributesOk() (*PredefinedPccRuleSetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PredefinedPccRuleSetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *PredefinedPccRuleSetSingleAllOf) SetAttributes(v PredefinedPccRuleSetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o PredefinedPccRuleSetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPccRuleSetSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullablePredefinedPccRuleSetSingleAllOf struct {
	value *PredefinedPccRuleSetSingleAllOf
	isSet bool
}

func (v NullablePredefinedPccRuleSetSingleAllOf) Get() *PredefinedPccRuleSetSingleAllOf {
	return v.value
}

func (v *NullablePredefinedPccRuleSetSingleAllOf) Set(val *PredefinedPccRuleSetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPccRuleSetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPccRuleSetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPccRuleSetSingleAllOf(val *PredefinedPccRuleSetSingleAllOf) *NullablePredefinedPccRuleSetSingleAllOf {
	return &NullablePredefinedPccRuleSetSingleAllOf{value: val, isSet: true}
}

func (v NullablePredefinedPccRuleSetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPccRuleSetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

