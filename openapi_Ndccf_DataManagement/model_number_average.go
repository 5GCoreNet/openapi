/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the NumberAverage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumberAverage{}

// NumberAverage Represents average and variance information.
type NumberAverage struct {
	// string with format 'float' as defined in OpenAPI.
	Number float32 `json:"number"`
	// string with format 'float' as defined in OpenAPI.
	Variance float32 `json:"variance"`
	// string with format 'float' as defined in OpenAPI.
	Skewness *float32 `json:"skewness,omitempty"`
}

// NewNumberAverage instantiates a new NumberAverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberAverage(number float32, variance float32) *NumberAverage {
	this := NumberAverage{}
	this.Number = number
	this.Variance = variance
	return &this
}

// NewNumberAverageWithDefaults instantiates a new NumberAverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberAverageWithDefaults() *NumberAverage {
	this := NumberAverage{}
	return &this
}

// GetNumber returns the Number field value
func (o *NumberAverage) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *NumberAverage) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *NumberAverage) SetNumber(v float32) {
	o.Number = v
}

// GetVariance returns the Variance field value
func (o *NumberAverage) GetVariance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Variance
}

// GetVarianceOk returns a tuple with the Variance field value
// and a boolean to check if the value has been set.
func (o *NumberAverage) GetVarianceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variance, true
}

// SetVariance sets field value
func (o *NumberAverage) SetVariance(v float32) {
	o.Variance = v
}

// GetSkewness returns the Skewness field value if set, zero value otherwise.
func (o *NumberAverage) GetSkewness() float32 {
	if o == nil || IsNil(o.Skewness) {
		var ret float32
		return ret
	}
	return *o.Skewness
}

// GetSkewnessOk returns a tuple with the Skewness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberAverage) GetSkewnessOk() (*float32, bool) {
	if o == nil || IsNil(o.Skewness) {
		return nil, false
	}
	return o.Skewness, true
}

// HasSkewness returns a boolean if a field has been set.
func (o *NumberAverage) HasSkewness() bool {
	if o != nil && !IsNil(o.Skewness) {
		return true
	}

	return false
}

// SetSkewness gets a reference to the given float32 and assigns it to the Skewness field.
func (o *NumberAverage) SetSkewness(v float32) {
	o.Skewness = &v
}

func (o NumberAverage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberAverage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["variance"] = o.Variance
	if !IsNil(o.Skewness) {
		toSerialize["skewness"] = o.Skewness
	}
	return toSerialize, nil
}

type NullableNumberAverage struct {
	value *NumberAverage
	isSet bool
}

func (v NullableNumberAverage) Get() *NumberAverage {
	return v.value
}

func (v *NullableNumberAverage) Set(val *NumberAverage) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberAverage) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberAverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberAverage(val *NumberAverage) *NullableNumberAverage {
	return &NullableNumberAverage{value: val, isSet: true}
}

func (v NullableNumberAverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberAverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
