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

// checks if the ExpectationTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpectationTarget{}

// ExpectationTarget This data type is the \"ExpectationTarget\" data type without specialisations
type ExpectationTarget struct {
	TargetName *string `json:"targetName,omitempty"`
	TargetCondition *Condition `json:"targetCondition,omitempty"`
	TargetValueRange *float32 `json:"targetValueRange,omitempty"`
	TargetContexts []TargetContext `json:"targetContexts,omitempty"`
}

// NewExpectationTarget instantiates a new ExpectationTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpectationTarget() *ExpectationTarget {
	this := ExpectationTarget{}
	return &this
}

// NewExpectationTargetWithDefaults instantiates a new ExpectationTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpectationTargetWithDefaults() *ExpectationTarget {
	this := ExpectationTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *ExpectationTarget) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *ExpectationTarget) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *ExpectationTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *ExpectationTarget) GetTargetCondition() Condition {
	if o == nil || IsNil(o.TargetCondition) {
		var ret Condition
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationTarget) GetTargetConditionOk() (*Condition, bool) {
	if o == nil || IsNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *ExpectationTarget) HasTargetCondition() bool {
	if o != nil && !IsNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given Condition and assigns it to the TargetCondition field.
func (o *ExpectationTarget) SetTargetCondition(v Condition) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *ExpectationTarget) GetTargetValueRange() float32 {
	if o == nil || IsNil(o.TargetValueRange) {
		var ret float32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationTarget) GetTargetValueRangeOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *ExpectationTarget) HasTargetValueRange() bool {
	if o != nil && !IsNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given float32 and assigns it to the TargetValueRange field.
func (o *ExpectationTarget) SetTargetValueRange(v float32) {
	o.TargetValueRange = &v
}

// GetTargetContexts returns the TargetContexts field value if set, zero value otherwise.
func (o *ExpectationTarget) GetTargetContexts() []TargetContext {
	if o == nil || IsNil(o.TargetContexts) {
		var ret []TargetContext
		return ret
	}
	return o.TargetContexts
}

// GetTargetContextsOk returns a tuple with the TargetContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectationTarget) GetTargetContextsOk() ([]TargetContext, bool) {
	if o == nil || IsNil(o.TargetContexts) {
		return nil, false
	}
	return o.TargetContexts, true
}

// HasTargetContexts returns a boolean if a field has been set.
func (o *ExpectationTarget) HasTargetContexts() bool {
	if o != nil && !IsNil(o.TargetContexts) {
		return true
	}

	return false
}

// SetTargetContexts gets a reference to the given []TargetContext and assigns it to the TargetContexts field.
func (o *ExpectationTarget) SetTargetContexts(v []TargetContext) {
	o.TargetContexts = v
}

func (o ExpectationTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpectationTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetName) {
		toSerialize["targetName"] = o.TargetName
	}
	if !IsNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !IsNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	if !IsNil(o.TargetContexts) {
		toSerialize["targetContexts"] = o.TargetContexts
	}
	return toSerialize, nil
}

type NullableExpectationTarget struct {
	value *ExpectationTarget
	isSet bool
}

func (v NullableExpectationTarget) Get() *ExpectationTarget {
	return v.value
}

func (v *NullableExpectationTarget) Set(val *ExpectationTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectationTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectationTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectationTarget(val *ExpectationTarget) *NullableExpectationTarget {
	return &NullableExpectationTarget{value: val, isSet: true}
}

func (v NullableExpectationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectationTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


