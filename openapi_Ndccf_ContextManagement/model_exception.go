/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the Exception type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Exception{}

// Exception Represents the Exception information.
type Exception struct {
	ExcepId    ExceptionId     `json:"excepId"`
	ExcepLevel *int32          `json:"excepLevel,omitempty"`
	ExcepTrend *ExceptionTrend `json:"excepTrend,omitempty"`
}

// NewException instantiates a new Exception object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewException(excepId ExceptionId) *Exception {
	this := Exception{}
	this.ExcepId = excepId
	return &this
}

// NewExceptionWithDefaults instantiates a new Exception object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionWithDefaults() *Exception {
	this := Exception{}
	return &this
}

// GetExcepId returns the ExcepId field value
func (o *Exception) GetExcepId() ExceptionId {
	if o == nil {
		var ret ExceptionId
		return ret
	}

	return o.ExcepId
}

// GetExcepIdOk returns a tuple with the ExcepId field value
// and a boolean to check if the value has been set.
func (o *Exception) GetExcepIdOk() (*ExceptionId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcepId, true
}

// SetExcepId sets field value
func (o *Exception) SetExcepId(v ExceptionId) {
	o.ExcepId = v
}

// GetExcepLevel returns the ExcepLevel field value if set, zero value otherwise.
func (o *Exception) GetExcepLevel() int32 {
	if o == nil || IsNil(o.ExcepLevel) {
		var ret int32
		return ret
	}
	return *o.ExcepLevel
}

// GetExcepLevelOk returns a tuple with the ExcepLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Exception) GetExcepLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ExcepLevel) {
		return nil, false
	}
	return o.ExcepLevel, true
}

// HasExcepLevel returns a boolean if a field has been set.
func (o *Exception) HasExcepLevel() bool {
	if o != nil && !IsNil(o.ExcepLevel) {
		return true
	}

	return false
}

// SetExcepLevel gets a reference to the given int32 and assigns it to the ExcepLevel field.
func (o *Exception) SetExcepLevel(v int32) {
	o.ExcepLevel = &v
}

// GetExcepTrend returns the ExcepTrend field value if set, zero value otherwise.
func (o *Exception) GetExcepTrend() ExceptionTrend {
	if o == nil || IsNil(o.ExcepTrend) {
		var ret ExceptionTrend
		return ret
	}
	return *o.ExcepTrend
}

// GetExcepTrendOk returns a tuple with the ExcepTrend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Exception) GetExcepTrendOk() (*ExceptionTrend, bool) {
	if o == nil || IsNil(o.ExcepTrend) {
		return nil, false
	}
	return o.ExcepTrend, true
}

// HasExcepTrend returns a boolean if a field has been set.
func (o *Exception) HasExcepTrend() bool {
	if o != nil && !IsNil(o.ExcepTrend) {
		return true
	}

	return false
}

// SetExcepTrend gets a reference to the given ExceptionTrend and assigns it to the ExcepTrend field.
func (o *Exception) SetExcepTrend(v ExceptionTrend) {
	o.ExcepTrend = &v
}

func (o Exception) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Exception) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["excepId"] = o.ExcepId
	if !IsNil(o.ExcepLevel) {
		toSerialize["excepLevel"] = o.ExcepLevel
	}
	if !IsNil(o.ExcepTrend) {
		toSerialize["excepTrend"] = o.ExcepTrend
	}
	return toSerialize, nil
}

type NullableException struct {
	value *Exception
	isSet bool
}

func (v NullableException) Get() *Exception {
	return v.value
}

func (v *NullableException) Set(val *Exception) {
	v.value = val
	v.isSet = true
}

func (v NullableException) IsSet() bool {
	return v.isSet
}

func (v *NullableException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableException(val *Exception) *NullableException {
	return &NullableException{value: val, isSet: true}
}

func (v NullableException) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
