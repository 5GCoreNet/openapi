/*
MSGS_ASRegistration

API for MSGS AS Registration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGS_ASRegistration

import (
	"encoding/json"
)

// checks if the ASRegistrationAck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASRegistrationAck{}

// ASRegistrationAck AS registration response data
type ASRegistrationAck struct {
	AsSvcId string         `json:"asSvcId"`
	Result  ProblemDetails `json:"result"`
}

// NewASRegistrationAck instantiates a new ASRegistrationAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASRegistrationAck(asSvcId string, result ProblemDetails) *ASRegistrationAck {
	this := ASRegistrationAck{}
	this.AsSvcId = asSvcId
	this.Result = result
	return &this
}

// NewASRegistrationAckWithDefaults instantiates a new ASRegistrationAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASRegistrationAckWithDefaults() *ASRegistrationAck {
	this := ASRegistrationAck{}
	return &this
}

// GetAsSvcId returns the AsSvcId field value
func (o *ASRegistrationAck) GetAsSvcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AsSvcId
}

// GetAsSvcIdOk returns a tuple with the AsSvcId field value
// and a boolean to check if the value has been set.
func (o *ASRegistrationAck) GetAsSvcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsSvcId, true
}

// SetAsSvcId sets field value
func (o *ASRegistrationAck) SetAsSvcId(v string) {
	o.AsSvcId = v
}

// GetResult returns the Result field value
func (o *ASRegistrationAck) GetResult() ProblemDetails {
	if o == nil {
		var ret ProblemDetails
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ASRegistrationAck) GetResultOk() (*ProblemDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ASRegistrationAck) SetResult(v ProblemDetails) {
	o.Result = v
}

func (o ASRegistrationAck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ASRegistrationAck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asSvcId"] = o.AsSvcId
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableASRegistrationAck struct {
	value *ASRegistrationAck
	isSet bool
}

func (v NullableASRegistrationAck) Get() *ASRegistrationAck {
	return v.value
}

func (v *NullableASRegistrationAck) Set(val *ASRegistrationAck) {
	v.value = val
	v.isSet = true
}

func (v NullableASRegistrationAck) IsSet() bool {
	return v.isSet
}

func (v *NullableASRegistrationAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASRegistrationAck(val *ASRegistrationAck) *NullableASRegistrationAck {
	return &NullableASRegistrationAck{value: val, isSet: true}
}

func (v NullableASRegistrationAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASRegistrationAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
