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

// checks if the ActivityFactorTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityFactorTarget{}

// ActivityFactorTarget This data type is the \"ExpectationTarget\" data type with specialisations for ActivityFactorTarget       
type ActivityFactorTarget struct {
	TargetAttribute *string `json:"targetAttribute,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
}

// NewActivityFactorTarget instantiates a new ActivityFactorTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityFactorTarget() *ActivityFactorTarget {
	this := ActivityFactorTarget{}
	return &this
}

// NewActivityFactorTargetWithDefaults instantiates a new ActivityFactorTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityFactorTargetWithDefaults() *ActivityFactorTarget {
	this := ActivityFactorTarget{}
	return &this
}

// GetTargetAttribute returns the TargetAttribute field value if set, zero value otherwise.
func (o *ActivityFactorTarget) GetTargetAttribute() string {
	if o == nil || isNil(o.TargetAttribute) {
		var ret string
		return ret
	}
	return *o.TargetAttribute
}

// GetTargetAttributeOk returns a tuple with the TargetAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityFactorTarget) GetTargetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.TargetAttribute) {
		return nil, false
	}
	return o.TargetAttribute, true
}

// HasTargetAttribute returns a boolean if a field has been set.
func (o *ActivityFactorTarget) HasTargetAttribute() bool {
	if o != nil && !isNil(o.TargetAttribute) {
		return true
	}

	return false
}

// SetTargetAttribute gets a reference to the given string and assigns it to the TargetAttribute field.
func (o *ActivityFactorTarget) SetTargetAttribute(v string) {
	o.TargetAttribute = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *ActivityFactorTarget) GetTargetCondition() string {
	if o == nil || isNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityFactorTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || isNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *ActivityFactorTarget) HasTargetCondition() bool {
	if o != nil && !isNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *ActivityFactorTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *ActivityFactorTarget) GetTargetValueRange() int32 {
	if o == nil || isNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityFactorTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || isNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *ActivityFactorTarget) HasTargetValueRange() bool {
	if o != nil && !isNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *ActivityFactorTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

func (o ActivityFactorTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityFactorTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TargetAttribute) {
		toSerialize["targetAttribute"] = o.TargetAttribute
	}
	if !isNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !isNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	return toSerialize, nil
}

type NullableActivityFactorTarget struct {
	value *ActivityFactorTarget
	isSet bool
}

func (v NullableActivityFactorTarget) Get() *ActivityFactorTarget {
	return v.value
}

func (v *NullableActivityFactorTarget) Set(val *ActivityFactorTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityFactorTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityFactorTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityFactorTarget(val *ActivityFactorTarget) *NullableActivityFactorTarget {
	return &NullableActivityFactorTarget{value: val, isSet: true}
}

func (v NullableActivityFactorTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityFactorTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


