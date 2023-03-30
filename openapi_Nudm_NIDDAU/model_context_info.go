/*
Nudm_NIDDAU

Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_NIDDAU

import (
	"encoding/json"
)

// checks if the ContextInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextInfo{}

// ContextInfo struct for ContextInfo
type ContextInfo struct {
	OrigHeaders []string `json:"origHeaders,omitempty"`
	RequestHeaders []string `json:"requestHeaders,omitempty"`
}

// NewContextInfo instantiates a new ContextInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextInfo() *ContextInfo {
	this := ContextInfo{}
	return &this
}

// NewContextInfoWithDefaults instantiates a new ContextInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextInfoWithDefaults() *ContextInfo {
	this := ContextInfo{}
	return &this
}

// GetOrigHeaders returns the OrigHeaders field value if set, zero value otherwise.
func (o *ContextInfo) GetOrigHeaders() []string {
	if o == nil || IsNil(o.OrigHeaders) {
		var ret []string
		return ret
	}
	return o.OrigHeaders
}

// GetOrigHeadersOk returns a tuple with the OrigHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextInfo) GetOrigHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.OrigHeaders) {
		return nil, false
	}
	return o.OrigHeaders, true
}

// HasOrigHeaders returns a boolean if a field has been set.
func (o *ContextInfo) HasOrigHeaders() bool {
	if o != nil && !IsNil(o.OrigHeaders) {
		return true
	}

	return false
}

// SetOrigHeaders gets a reference to the given []string and assigns it to the OrigHeaders field.
func (o *ContextInfo) SetOrigHeaders(v []string) {
	o.OrigHeaders = v
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise.
func (o *ContextInfo) GetRequestHeaders() []string {
	if o == nil || IsNil(o.RequestHeaders) {
		var ret []string
		return ret
	}
	return o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextInfo) GetRequestHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestHeaders) {
		return nil, false
	}
	return o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *ContextInfo) HasRequestHeaders() bool {
	if o != nil && !IsNil(o.RequestHeaders) {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given []string and assigns it to the RequestHeaders field.
func (o *ContextInfo) SetRequestHeaders(v []string) {
	o.RequestHeaders = v
}

func (o ContextInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrigHeaders) {
		toSerialize["origHeaders"] = o.OrigHeaders
	}
	if !IsNil(o.RequestHeaders) {
		toSerialize["requestHeaders"] = o.RequestHeaders
	}
	return toSerialize, nil
}

type NullableContextInfo struct {
	value *ContextInfo
	isSet bool
}

func (v NullableContextInfo) Get() *ContextInfo {
	return v.value
}

func (v *NullableContextInfo) Set(val *ContextInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableContextInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableContextInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextInfo(val *ContextInfo) *NullableContextInfo {
	return &NullableContextInfo{value: val, isSet: true}
}

func (v NullableContextInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


