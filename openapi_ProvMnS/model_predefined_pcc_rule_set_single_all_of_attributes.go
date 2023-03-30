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

// checks if the PredefinedPccRuleSetSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredefinedPccRuleSetSingleAllOfAttributes{}

// PredefinedPccRuleSetSingleAllOfAttributes struct for PredefinedPccRuleSetSingleAllOfAttributes
type PredefinedPccRuleSetSingleAllOfAttributes struct {
	PredefinedPccRules []PccRule `json:"predefinedPccRules,omitempty"`
}

// NewPredefinedPccRuleSetSingleAllOfAttributes instantiates a new PredefinedPccRuleSetSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredefinedPccRuleSetSingleAllOfAttributes() *PredefinedPccRuleSetSingleAllOfAttributes {
	this := PredefinedPccRuleSetSingleAllOfAttributes{}
	return &this
}

// NewPredefinedPccRuleSetSingleAllOfAttributesWithDefaults instantiates a new PredefinedPccRuleSetSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredefinedPccRuleSetSingleAllOfAttributesWithDefaults() *PredefinedPccRuleSetSingleAllOfAttributes {
	this := PredefinedPccRuleSetSingleAllOfAttributes{}
	return &this
}

// GetPredefinedPccRules returns the PredefinedPccRules field value if set, zero value otherwise.
func (o *PredefinedPccRuleSetSingleAllOfAttributes) GetPredefinedPccRules() []PccRule {
	if o == nil || IsNil(o.PredefinedPccRules) {
		var ret []PccRule
		return ret
	}
	return o.PredefinedPccRules
}

// GetPredefinedPccRulesOk returns a tuple with the PredefinedPccRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredefinedPccRuleSetSingleAllOfAttributes) GetPredefinedPccRulesOk() ([]PccRule, bool) {
	if o == nil || IsNil(o.PredefinedPccRules) {
		return nil, false
	}
	return o.PredefinedPccRules, true
}

// HasPredefinedPccRules returns a boolean if a field has been set.
func (o *PredefinedPccRuleSetSingleAllOfAttributes) HasPredefinedPccRules() bool {
	if o != nil && !IsNil(o.PredefinedPccRules) {
		return true
	}

	return false
}

// SetPredefinedPccRules gets a reference to the given []PccRule and assigns it to the PredefinedPccRules field.
func (o *PredefinedPccRuleSetSingleAllOfAttributes) SetPredefinedPccRules(v []PccRule) {
	o.PredefinedPccRules = v
}

func (o PredefinedPccRuleSetSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredefinedPccRuleSetSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PredefinedPccRules) {
		toSerialize["predefinedPccRules"] = o.PredefinedPccRules
	}
	return toSerialize, nil
}

type NullablePredefinedPccRuleSetSingleAllOfAttributes struct {
	value *PredefinedPccRuleSetSingleAllOfAttributes
	isSet bool
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributes) Get() *PredefinedPccRuleSetSingleAllOfAttributes {
	return v.value
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributes) Set(val *PredefinedPccRuleSetSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredefinedPccRuleSetSingleAllOfAttributes(val *PredefinedPccRuleSetSingleAllOfAttributes) *NullablePredefinedPccRuleSetSingleAllOfAttributes {
	return &NullablePredefinedPccRuleSetSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullablePredefinedPccRuleSetSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredefinedPccRuleSetSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


