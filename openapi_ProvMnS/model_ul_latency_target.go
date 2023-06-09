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

// checks if the ULLatencyTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ULLatencyTarget{}

// ULLatencyTarget This data type is the \"ExpectationTarget\" data type with specialisations for ULLatencyTarget
type ULLatencyTarget struct {
	TargetName       *string `json:"targetName,omitempty"`
	TargetCondition  *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32  `json:"targetValueRange,omitempty"`
}

// NewULLatencyTarget instantiates a new ULLatencyTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewULLatencyTarget() *ULLatencyTarget {
	this := ULLatencyTarget{}
	return &this
}

// NewULLatencyTargetWithDefaults instantiates a new ULLatencyTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewULLatencyTargetWithDefaults() *ULLatencyTarget {
	this := ULLatencyTarget{}
	return &this
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *ULLatencyTarget) GetTargetName() string {
	if o == nil || IsNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ULLatencyTarget) GetTargetNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *ULLatencyTarget) HasTargetName() bool {
	if o != nil && !IsNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *ULLatencyTarget) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *ULLatencyTarget) GetTargetCondition() string {
	if o == nil || IsNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ULLatencyTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCondition) {
		return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *ULLatencyTarget) HasTargetCondition() bool {
	if o != nil && !IsNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *ULLatencyTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *ULLatencyTarget) GetTargetValueRange() int32 {
	if o == nil || IsNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ULLatencyTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetValueRange) {
		return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *ULLatencyTarget) HasTargetValueRange() bool {
	if o != nil && !IsNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *ULLatencyTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

func (o ULLatencyTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ULLatencyTarget) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableULLatencyTarget struct {
	value *ULLatencyTarget
	isSet bool
}

func (v NullableULLatencyTarget) Get() *ULLatencyTarget {
	return v.value
}

func (v *NullableULLatencyTarget) Set(val *ULLatencyTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableULLatencyTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableULLatencyTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableULLatencyTarget(val *ULLatencyTarget) *NullableULLatencyTarget {
	return &NullableULLatencyTarget{value: val, isSet: true}
}

func (v NullableULLatencyTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableULLatencyTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
