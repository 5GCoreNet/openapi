/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_SMContext

import (
	"encoding/json"
)

// checks if the SmContextReleaseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmContextReleaseData{}

// SmContextReleaseData Representation of the information to release the Individual SM context.
type SmContextReleaseData struct {
	Cause ReleaseCause `json:"cause"`
}

// NewSmContextReleaseData instantiates a new SmContextReleaseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextReleaseData(cause ReleaseCause) *SmContextReleaseData {
	this := SmContextReleaseData{}
	this.Cause = cause
	return &this
}

// NewSmContextReleaseDataWithDefaults instantiates a new SmContextReleaseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextReleaseDataWithDefaults() *SmContextReleaseData {
	this := SmContextReleaseData{}
	return &this
}

// GetCause returns the Cause field value
func (o *SmContextReleaseData) GetCause() ReleaseCause {
	if o == nil {
		var ret ReleaseCause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *SmContextReleaseData) GetCauseOk() (*ReleaseCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *SmContextReleaseData) SetCause(v ReleaseCause) {
	o.Cause = v
}

func (o SmContextReleaseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmContextReleaseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cause"] = o.Cause
	return toSerialize, nil
}

type NullableSmContextReleaseData struct {
	value *SmContextReleaseData
	isSet bool
}

func (v NullableSmContextReleaseData) Get() *SmContextReleaseData {
	return v.value
}

func (v *NullableSmContextReleaseData) Set(val *SmContextReleaseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextReleaseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextReleaseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextReleaseData(val *SmContextReleaseData) *NullableSmContextReleaseData {
	return &NullableSmContextReleaseData{value: val, isSet: true}
}

func (v NullableSmContextReleaseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextReleaseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
