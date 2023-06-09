/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the ProblemDetailsAddInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetailsAddInfo{}

// ProblemDetailsAddInfo Problem Details Additional Information
type ProblemDetailsAddInfo struct {
	RemoteError *bool `json:"remoteError,omitempty"`
}

// NewProblemDetailsAddInfo instantiates a new ProblemDetailsAddInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetailsAddInfo() *ProblemDetailsAddInfo {
	this := ProblemDetailsAddInfo{}
	return &this
}

// NewProblemDetailsAddInfoWithDefaults instantiates a new ProblemDetailsAddInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailsAddInfoWithDefaults() *ProblemDetailsAddInfo {
	this := ProblemDetailsAddInfo{}
	return &this
}

// GetRemoteError returns the RemoteError field value if set, zero value otherwise.
func (o *ProblemDetailsAddInfo) GetRemoteError() bool {
	if o == nil || IsNil(o.RemoteError) {
		var ret bool
		return ret
	}
	return *o.RemoteError
}

// GetRemoteErrorOk returns a tuple with the RemoteError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAddInfo) GetRemoteErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteError) {
		return nil, false
	}
	return o.RemoteError, true
}

// HasRemoteError returns a boolean if a field has been set.
func (o *ProblemDetailsAddInfo) HasRemoteError() bool {
	if o != nil && !IsNil(o.RemoteError) {
		return true
	}

	return false
}

// SetRemoteError gets a reference to the given bool and assigns it to the RemoteError field.
func (o *ProblemDetailsAddInfo) SetRemoteError(v bool) {
	o.RemoteError = &v
}

func (o ProblemDetailsAddInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetailsAddInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemoteError) {
		toSerialize["remoteError"] = o.RemoteError
	}
	return toSerialize, nil
}

type NullableProblemDetailsAddInfo struct {
	value *ProblemDetailsAddInfo
	isSet bool
}

func (v NullableProblemDetailsAddInfo) Get() *ProblemDetailsAddInfo {
	return v.value
}

func (v *NullableProblemDetailsAddInfo) Set(val *ProblemDetailsAddInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetailsAddInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetailsAddInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetailsAddInfo(val *ProblemDetailsAddInfo) *NullableProblemDetailsAddInfo {
	return &NullableProblemDetailsAddInfo{value: val, isSet: true}
}

func (v NullableProblemDetailsAddInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetailsAddInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
