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

// checks if the PredefinedPccRuleSetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPccRuleSetSingle{}

// PredefinedPccRuleSetSingle struct for PredefinedPccRuleSetSingle
type PredefinedPccRuleSetSingle struct {
	Top
	Attributes *PredefinedPccRuleSetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewPredefinedPccRuleSetSingle instantiates a new PredefinedPccRuleSetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPccRuleSetSingle(id NullableString) *PredefinedPccRuleSetSingle {
	this := PredefinedPccRuleSetSingle{}
	this.Id = id
	return &this
}

// NewPredefinedPccRuleSetSingleWithDefaults instantiates a new PredefinedPccRuleSetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPccRuleSetSingleWithDefaults() *PredefinedPccRuleSetSingle {
	this := PredefinedPccRuleSetSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingle) GetAttributes() PredefinedPccRuleSetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PredefinedPccRuleSetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingle) GetAttributesOk() (*PredefinedPccRuleSetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PredefinedPccRuleSetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *PredefinedPccRuleSetSingle) SetAttributes(v PredefinedPccRuleSetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o PredefinedPccRuleSetSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPccRuleSetSingle) ToMap() (map[string]interface{}, error) {
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

type NullablePredefinedPccRuleSetSingle struct {
	value *PredefinedPccRuleSetSingle
	isSet bool
}

func (v NullablePredefinedPccRuleSetSingle) Get() *PredefinedPccRuleSetSingle {
	return v.value
}

func (v *NullablePredefinedPccRuleSetSingle) Set(val *PredefinedPccRuleSetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPccRuleSetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPccRuleSetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPccRuleSetSingle(val *PredefinedPccRuleSetSingle) *NullablePredefinedPccRuleSetSingle {
	return &NullablePredefinedPccRuleSetSingle{value: val, isSet: true}
}

func (v NullablePredefinedPccRuleSetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPccRuleSetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
