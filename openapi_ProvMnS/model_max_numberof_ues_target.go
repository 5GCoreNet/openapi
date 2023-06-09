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

// checks if the MaxNumberofUEsTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxNumberofUEsTarget{}

// MaxNumberofUEsTarget This data type is the \"ExpectationTarget\" data type with specialisations for MaxNumberofUEsTarget
type MaxNumberofUEsTarget struct {
	TargetAttribute  *string `json:"targetAttribute,omitempty"`
	TargetCondition  *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32  `json:"targetValueRange,omitempty"`
}

// NewMaxNumberofUEsTarget instantiates a new MaxNumberofUEsTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxNumberofUEsTarget() *MaxNumberofUEsTarget {
	this := MaxNumberofUEsTarget{}
	return &this
}

// NewMaxNumberofUEsTargetWithDefaults instantiates a new MaxNumberofUEsTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxNumberofUEsTargetWithDefaults() *MaxNumberofUEsTarget {
	this := MaxNumberofUEsTarget{}
	return &this
}

// GetTargetAttribute returns the TargetAttribute field value if set, zero value otherwise.
func (o *MaxNumberofUEsTarget) GetTargetAttribute() string {
	if o == nil || IsNil(o.TargetAttribute) {
		var ret string
		return ret
	}
	return *o.TargetAttribute
}

// GetTargetAttributeOk returns a tuple with the TargetAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEsTarget) GetTargetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAttribute) {
		return nil, false
	}
	return o.TargetAttribute, true
}

// HasTargetAttribute returns a boolean if a field has been set.
func (o *MaxNumberofUEsTarget) HasTargetAttribute() bool {
	if o != nil && !IsNil(o.TargetAttribute) {
		return true
	}

	return false
}

// SetTargetAttribute gets a reference to the given string and assigns it to the TargetAttribute field.
func (o *MaxNumberofUEsTarget) SetTargetAttribute(v string) {
	o.TargetAttribute = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *MaxNumberofUEsTarget) GetTargetCondition() string {
	if o == nil || IsNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEsTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *MaxNumberofUEsTarget) HasTargetCondition() bool {
	if o != nil && !IsNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *MaxNumberofUEsTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *MaxNumberofUEsTarget) GetTargetValueRange() int32 {
	if o == nil || IsNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaxNumberofUEsTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *MaxNumberofUEsTarget) HasTargetValueRange() bool {
	if o != nil && !IsNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *MaxNumberofUEsTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

func (o MaxNumberofUEsTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxNumberofUEsTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetAttribute) {
		toSerialize["targetAttribute"] = o.TargetAttribute
	}
	if !IsNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !IsNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	return toSerialize, nil
}

type NullableMaxNumberofUEsTarget struct {
	value *MaxNumberofUEsTarget
	isSet bool
}

func (v NullableMaxNumberofUEsTarget) Get() *MaxNumberofUEsTarget {
	return v.value
}

func (v *NullableMaxNumberofUEsTarget) Set(val *MaxNumberofUEsTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxNumberofUEsTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxNumberofUEsTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxNumberofUEsTarget(val *MaxNumberofUEsTarget) *NullableMaxNumberofUEsTarget {
	return &NullableMaxNumberofUEsTarget{value: val, isSet: true}
}

func (v NullableMaxNumberofUEsTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxNumberofUEsTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
