/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the InvocationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvocationResult{}

// InvocationResult struct for InvocationResult
type InvocationResult struct {
	Error *ProblemDetails `json:"error,omitempty"`
	FailureHandling *FailureHandling `json:"failureHandling,omitempty"`
}

// NewInvocationResult instantiates a new InvocationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvocationResult() *InvocationResult {
	this := InvocationResult{}
	return &this
}

// NewInvocationResultWithDefaults instantiates a new InvocationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvocationResultWithDefaults() *InvocationResult {
	this := InvocationResult{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InvocationResult) GetError() ProblemDetails {
	if o == nil || IsNil(o.Error) {
		var ret ProblemDetails
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResult) GetErrorOk() (*ProblemDetails, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InvocationResult) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ProblemDetails and assigns it to the Error field.
func (o *InvocationResult) SetError(v ProblemDetails) {
	o.Error = &v
}

// GetFailureHandling returns the FailureHandling field value if set, zero value otherwise.
func (o *InvocationResult) GetFailureHandling() FailureHandling {
	if o == nil || IsNil(o.FailureHandling) {
		var ret FailureHandling
		return ret
	}
	return *o.FailureHandling
}

// GetFailureHandlingOk returns a tuple with the FailureHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvocationResult) GetFailureHandlingOk() (*FailureHandling, bool) {
	if o == nil || IsNil(o.FailureHandling) {
		return nil, false
	}
	return o.FailureHandling, true
}

// HasFailureHandling returns a boolean if a field has been set.
func (o *InvocationResult) HasFailureHandling() bool {
	if o != nil && !IsNil(o.FailureHandling) {
		return true
	}

	return false
}

// SetFailureHandling gets a reference to the given FailureHandling and assigns it to the FailureHandling field.
func (o *InvocationResult) SetFailureHandling(v FailureHandling) {
	o.FailureHandling = &v
}

func (o InvocationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvocationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.FailureHandling) {
		toSerialize["failureHandling"] = o.FailureHandling
	}
	return toSerialize, nil
}

type NullableInvocationResult struct {
	value *InvocationResult
	isSet bool
}

func (v NullableInvocationResult) Get() *InvocationResult {
	return v.value
}

func (v *NullableInvocationResult) Set(val *InvocationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInvocationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInvocationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvocationResult(val *InvocationResult) *NullableInvocationResult {
	return &NullableInvocationResult{value: val, isSet: true}
}

func (v NullableInvocationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvocationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


