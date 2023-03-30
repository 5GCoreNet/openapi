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

// checks if the ParameterRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterRange{}

// ParameterRange struct for ParameterRange
type ParameterRange struct {
	MaxValue *int32 `json:"maxValue,omitempty"`
	MinValue *int32 `json:"minValue,omitempty"`
}

// NewParameterRange instantiates a new ParameterRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterRange() *ParameterRange {
	this := ParameterRange{}
	return &this
}

// NewParameterRangeWithDefaults instantiates a new ParameterRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterRangeWithDefaults() *ParameterRange {
	this := ParameterRange{}
	return &this
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ParameterRange) GetMaxValue() int32 {
	if o == nil || IsNil(o.MaxValue) {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterRange) GetMaxValueOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxValue) {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ParameterRange) HasMaxValue() bool {
	if o != nil && !IsNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *ParameterRange) SetMaxValue(v int32) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ParameterRange) GetMinValue() int32 {
	if o == nil || IsNil(o.MinValue) {
		var ret int32
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterRange) GetMinValueOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValue) {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ParameterRange) HasMinValue() bool {
	if o != nil && !IsNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int32 and assigns it to the MinValue field.
func (o *ParameterRange) SetMinValue(v int32) {
	o.MinValue = &v
}

func (o ParameterRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxValue) {
		toSerialize["maxValue"] = o.MaxValue
	}
	if !IsNil(o.MinValue) {
		toSerialize["minValue"] = o.MinValue
	}
	return toSerialize, nil
}

type NullableParameterRange struct {
	value *ParameterRange
	isSet bool
}

func (v NullableParameterRange) Get() *ParameterRange {
	return v.value
}

func (v *NullableParameterRange) Set(val *ParameterRange) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterRange) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterRange(val *ParameterRange) *NullableParameterRange {
	return &NullableParameterRange{value: val, isSet: true}
}

func (v NullableParameterRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


