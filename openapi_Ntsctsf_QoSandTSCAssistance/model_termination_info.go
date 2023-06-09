/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
)

// checks if the TerminationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminationInfo{}

// TerminationInfo Indicates the cause for requesting the deletion of the Individual Application Session Context resource.
type TerminationInfo struct {
	TermCause TerminationCause `json:"termCause"`
	// String providing an URI formatted according to RFC 3986.
	ResUri string `json:"resUri"`
}

// NewTerminationInfo instantiates a new TerminationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminationInfo(termCause TerminationCause, resUri string) *TerminationInfo {
	this := TerminationInfo{}
	this.TermCause = termCause
	this.ResUri = resUri
	return &this
}

// NewTerminationInfoWithDefaults instantiates a new TerminationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminationInfoWithDefaults() *TerminationInfo {
	this := TerminationInfo{}
	return &this
}

// GetTermCause returns the TermCause field value
func (o *TerminationInfo) GetTermCause() TerminationCause {
	if o == nil {
		var ret TerminationCause
		return ret
	}

	return o.TermCause
}

// GetTermCauseOk returns a tuple with the TermCause field value
// and a boolean to check if the value has been set.
func (o *TerminationInfo) GetTermCauseOk() (*TerminationCause, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermCause, true
}

// SetTermCause sets field value
func (o *TerminationInfo) SetTermCause(v TerminationCause) {
	o.TermCause = v
}

// GetResUri returns the ResUri field value
func (o *TerminationInfo) GetResUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResUri
}

// GetResUriOk returns a tuple with the ResUri field value
// and a boolean to check if the value has been set.
func (o *TerminationInfo) GetResUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResUri, true
}

// SetResUri sets field value
func (o *TerminationInfo) SetResUri(v string) {
	o.ResUri = v
}

func (o TerminationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["termCause"] = o.TermCause
	toSerialize["resUri"] = o.ResUri
	return toSerialize, nil
}

type NullableTerminationInfo struct {
	value *TerminationInfo
	isSet bool
}

func (v NullableTerminationInfo) Get() *TerminationInfo {
	return v.value
}

func (v *NullableTerminationInfo) Set(val *TerminationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationInfo(val *TerminationInfo) *NullableTerminationInfo {
	return &NullableTerminationInfo{value: val, isSet: true}
}

func (v NullableTerminationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
