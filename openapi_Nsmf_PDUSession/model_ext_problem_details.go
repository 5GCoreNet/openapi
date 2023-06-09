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

// checks if the ExtProblemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtProblemDetails{}

// ExtProblemDetails Extended Problem Details
type ExtProblemDetails struct {
	ProblemDetails
	RemoteError *bool `json:"remoteError,omitempty"`
}

// NewExtProblemDetails instantiates a new ExtProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtProblemDetails() *ExtProblemDetails {
	this := ExtProblemDetails{}
	return &this
}

// NewExtProblemDetailsWithDefaults instantiates a new ExtProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtProblemDetailsWithDefaults() *ExtProblemDetails {
	this := ExtProblemDetails{}
	return &this
}

// GetRemoteError returns the RemoteError field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetRemoteError() bool {
	if o == nil || IsNil(o.RemoteError) {
		var ret bool
		return ret
	}
	return *o.RemoteError
}

// GetRemoteErrorOk returns a tuple with the RemoteError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetRemoteErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteError) {
		return nil, false
	}
	return o.RemoteError, true
}

// HasRemoteError returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasRemoteError() bool {
	if o != nil && !IsNil(o.RemoteError) {
		return true
	}

	return false
}

// SetRemoteError gets a reference to the given bool and assigns it to the RemoteError field.
func (o *ExtProblemDetails) SetRemoteError(v bool) {
	o.RemoteError = &v
}

func (o ExtProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtProblemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProblemDetails, errProblemDetails := json.Marshal(o.ProblemDetails)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	errProblemDetails = json.Unmarshal([]byte(serializedProblemDetails), &toSerialize)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	if !IsNil(o.RemoteError) {
		toSerialize["remoteError"] = o.RemoteError
	}
	return toSerialize, nil
}

type NullableExtProblemDetails struct {
	value *ExtProblemDetails
	isSet bool
}

func (v NullableExtProblemDetails) Get() *ExtProblemDetails {
	return v.value
}

func (v *NullableExtProblemDetails) Set(val *ExtProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableExtProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableExtProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtProblemDetails(val *ExtProblemDetails) *NullableExtProblemDetails {
	return &NullableExtProblemDetails{value: val, isSet: true}
}

func (v NullableExtProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
