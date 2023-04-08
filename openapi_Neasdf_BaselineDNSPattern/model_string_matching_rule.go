/*
Neasdf_BaselineDNSPattern

EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_BaselineDNSPattern

import (
	"encoding/json"
)

// checks if the StringMatchingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringMatchingRule{}

// StringMatchingRule A list of conditions for string matching
type StringMatchingRule struct {
	StringMatchingConditions []StringMatchingCondition `json:"stringMatchingConditions,omitempty"`
}

// NewStringMatchingRule instantiates a new StringMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringMatchingRule() *StringMatchingRule {
	this := StringMatchingRule{}
	return &this
}

// NewStringMatchingRuleWithDefaults instantiates a new StringMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringMatchingRuleWithDefaults() *StringMatchingRule {
	this := StringMatchingRule{}
	return &this
}

// GetStringMatchingConditions returns the StringMatchingConditions field value if set, zero value otherwise.
func (o *StringMatchingRule) GetStringMatchingConditions() []StringMatchingCondition {
	if o == nil || isNil(o.StringMatchingConditions) {
		var ret []StringMatchingCondition
		return ret
	}
	return o.StringMatchingConditions
}

// GetStringMatchingConditionsOk returns a tuple with the StringMatchingConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMatchingRule) GetStringMatchingConditionsOk() ([]StringMatchingCondition, bool) {
	if o == nil || isNil(o.StringMatchingConditions) {
		return nil, false
	}
	return o.StringMatchingConditions, true
}

// HasStringMatchingConditions returns a boolean if a field has been set.
func (o *StringMatchingRule) HasStringMatchingConditions() bool {
	if o != nil && !isNil(o.StringMatchingConditions) {
		return true
	}

	return false
}

// SetStringMatchingConditions gets a reference to the given []StringMatchingCondition and assigns it to the StringMatchingConditions field.
func (o *StringMatchingRule) SetStringMatchingConditions(v []StringMatchingCondition) {
	o.StringMatchingConditions = v
}

func (o StringMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringMatchingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StringMatchingConditions) {
		toSerialize["stringMatchingConditions"] = o.StringMatchingConditions
	}
	return toSerialize, nil
}

type NullableStringMatchingRule struct {
	value *StringMatchingRule
	isSet bool
}

func (v NullableStringMatchingRule) Get() *StringMatchingRule {
	return v.value
}

func (v *NullableStringMatchingRule) Set(val *StringMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableStringMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableStringMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringMatchingRule(val *StringMatchingRule) *NullableStringMatchingRule {
	return &NullableStringMatchingRule{value: val, isSet: true}
}

func (v NullableStringMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


