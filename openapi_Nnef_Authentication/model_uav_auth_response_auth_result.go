/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_Authentication

import (
	"encoding/json"
)

// checks if the UAVAuthResponseAuthResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UAVAuthResponseAuthResult{}

// UAVAuthResponseAuthResult struct for UAVAuthResponseAuthResult
type UAVAuthResponseAuthResult struct {
	AuthResult
}

// NewUAVAuthResponseAuthResult instantiates a new UAVAuthResponseAuthResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAVAuthResponseAuthResult() *UAVAuthResponseAuthResult {
	this := UAVAuthResponseAuthResult{}
	return &this
}

// NewUAVAuthResponseAuthResultWithDefaults instantiates a new UAVAuthResponseAuthResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUAVAuthResponseAuthResultWithDefaults() *UAVAuthResponseAuthResult {
	this := UAVAuthResponseAuthResult{}
	return &this
}

func (o UAVAuthResponseAuthResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UAVAuthResponseAuthResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthResult, errAuthResult := json.Marshal(o.AuthResult)
	if errAuthResult != nil {
		return map[string]interface{}{}, errAuthResult
	}
	errAuthResult = json.Unmarshal([]byte(serializedAuthResult), &toSerialize)
	if errAuthResult != nil {
		return map[string]interface{}{}, errAuthResult
	}
	return toSerialize, nil
}

type NullableUAVAuthResponseAuthResult struct {
	value *UAVAuthResponseAuthResult
	isSet bool
}

func (v NullableUAVAuthResponseAuthResult) Get() *UAVAuthResponseAuthResult {
	return v.value
}

func (v *NullableUAVAuthResponseAuthResult) Set(val *UAVAuthResponseAuthResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUAVAuthResponseAuthResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUAVAuthResponseAuthResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUAVAuthResponseAuthResult(val *UAVAuthResponseAuthResult) *NullableUAVAuthResponseAuthResult {
	return &NullableUAVAuthResponseAuthResult{value: val, isSet: true}
}

func (v NullableUAVAuthResponseAuthResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUAVAuthResponseAuthResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
