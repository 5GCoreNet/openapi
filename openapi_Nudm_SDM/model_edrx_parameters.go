/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the EdrxParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdrxParameters{}

// EdrxParameters struct for EdrxParameters
type EdrxParameters struct {
	RatType RatType `json:"ratType"`
	EdrxValue string `json:"edrxValue"`
}

// NewEdrxParameters instantiates a new EdrxParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdrxParameters(ratType RatType, edrxValue string) *EdrxParameters {
	this := EdrxParameters{}
	this.RatType = ratType
	this.EdrxValue = edrxValue
	return &this
}

// NewEdrxParametersWithDefaults instantiates a new EdrxParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdrxParametersWithDefaults() *EdrxParameters {
	this := EdrxParameters{}
	return &this
}

// GetRatType returns the RatType field value
func (o *EdrxParameters) GetRatType() RatType {
	if o == nil {
		var ret RatType
		return ret
	}

	return o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value
// and a boolean to check if the value has been set.
func (o *EdrxParameters) GetRatTypeOk() (*RatType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatType, true
}

// SetRatType sets field value
func (o *EdrxParameters) SetRatType(v RatType) {
	o.RatType = v
}

// GetEdrxValue returns the EdrxValue field value
func (o *EdrxParameters) GetEdrxValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdrxValue
}

// GetEdrxValueOk returns a tuple with the EdrxValue field value
// and a boolean to check if the value has been set.
func (o *EdrxParameters) GetEdrxValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdrxValue, true
}

// SetEdrxValue sets field value
func (o *EdrxParameters) SetEdrxValue(v string) {
	o.EdrxValue = v
}

func (o EdrxParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdrxParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ratType"] = o.RatType
	toSerialize["edrxValue"] = o.EdrxValue
	return toSerialize, nil
}

type NullableEdrxParameters struct {
	value *EdrxParameters
	isSet bool
}

func (v NullableEdrxParameters) Get() *EdrxParameters {
	return v.value
}

func (v *NullableEdrxParameters) Set(val *EdrxParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEdrxParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEdrxParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdrxParameters(val *EdrxParameters) *NullableEdrxParameters {
	return &NullableEdrxParameters{value: val, isSet: true}
}

func (v NullableEdrxParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdrxParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


