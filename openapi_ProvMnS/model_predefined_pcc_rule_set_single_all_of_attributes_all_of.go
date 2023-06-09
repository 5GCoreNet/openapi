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

// checks if the PredefinedPccRuleSetSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPccRuleSetSingleAllOfAttributesAllOf{}

// PredefinedPccRuleSetSingleAllOfAttributesAllOf struct for PredefinedPccRuleSetSingleAllOfAttributesAllOf
type PredefinedPccRuleSetSingleAllOfAttributesAllOf struct {
	PredefinedPccRules []PccRule `json:"predefinedPccRules,omitempty"`
}

// NewPredefinedPccRuleSetSingleAllOfAttributesAllOf instantiates a new PredefinedPccRuleSetSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPccRuleSetSingleAllOfAttributesAllOf() *PredefinedPccRuleSetSingleAllOfAttributesAllOf {
	this := PredefinedPccRuleSetSingleAllOfAttributesAllOf{}
	return &this
}

// NewPredefinedPccRuleSetSingleAllOfAttributesAllOfWithDefaults instantiates a new PredefinedPccRuleSetSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPccRuleSetSingleAllOfAttributesAllOfWithDefaults() *PredefinedPccRuleSetSingleAllOfAttributesAllOf {
	this := PredefinedPccRuleSetSingleAllOfAttributesAllOf{}
	return &this
}

// GetPredefinedPccRules returns the PredefinedPccRules field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingleAllOfAttributesAllOf) GetPredefinedPccRules() []PccRule {
	if o == nil || IsNil(o.PredefinedPccRules) {
		var ret []PccRule
		return ret
	}
	return o.PredefinedPccRules
}

// GetPredefinedPccRulesOk returns a tuple with the PredefinedPccRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingleAllOfAttributesAllOf) GetPredefinedPccRulesOk() ([]PccRule, bool) {
	if o == nil || IsNil(o.PredefinedPccRules) {
		return nil, false
	}
	return o.PredefinedPccRules, true
}

// HasPredefinedPccRules returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingleAllOfAttributesAllOf) HasPredefinedPccRules() bool {
	if o != nil && !IsNil(o.PredefinedPccRules) {
		return true
	}

	return false
}

// SetPredefinedPccRules gets a reference to the given []PccRule and assigns it to the PredefinedPccRules field.
func (o *PredefinedPccRuleSetSingleAllOfAttributesAllOf) SetPredefinedPccRules(v []PccRule) {
	o.PredefinedPccRules = v
}

func (o PredefinedPccRuleSetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPccRuleSetSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PredefinedPccRules) {
		toSerialize["predefinedPccRules"] = o.PredefinedPccRules
	}
	return toSerialize, nil
}

type NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf struct {
	value *PredefinedPccRuleSetSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) Get() *PredefinedPccRuleSetSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) Set(val *PredefinedPccRuleSetSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPccRuleSetSingleAllOfAttributesAllOf(val *PredefinedPccRuleSetSingleAllOfAttributesAllOf) *NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf {
	return &NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
