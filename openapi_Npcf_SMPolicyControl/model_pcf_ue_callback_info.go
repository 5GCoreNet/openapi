/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the PcfUeCallbackInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfUeCallbackInfo{}

// PcfUeCallbackInfo Contains the PCF for the UE information necessary for the PCF for the PDU session to send  SM Policy Association Establishment and Termination events. 
type PcfUeCallbackInfo struct {
	// String providing an URI formatted according to RFC 3986.
	CallbackUri string `json:"callbackUri"`
	BindingInfo *string `json:"bindingInfo,omitempty"`
}

// NewPcfUeCallbackInfo instantiates a new PcfUeCallbackInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfUeCallbackInfo(callbackUri string) *PcfUeCallbackInfo {
	this := PcfUeCallbackInfo{}
	this.CallbackUri = callbackUri
	return &this
}

// NewPcfUeCallbackInfoWithDefaults instantiates a new PcfUeCallbackInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfUeCallbackInfoWithDefaults() *PcfUeCallbackInfo {
	this := PcfUeCallbackInfo{}
	return &this
}

// GetCallbackUri returns the CallbackUri field value
func (o *PcfUeCallbackInfo) GetCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUri
}

// GetCallbackUriOk returns a tuple with the CallbackUri field value
// and a boolean to check if the value has been set.
func (o *PcfUeCallbackInfo) GetCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUri, true
}

// SetCallbackUri sets field value
func (o *PcfUeCallbackInfo) SetCallbackUri(v string) {
	o.CallbackUri = v
}

// GetBindingInfo returns the BindingInfo field value if set, zero value otherwise.
func (o *PcfUeCallbackInfo) GetBindingInfo() string {
	if o == nil || IsNil(o.BindingInfo) {
		var ret string
		return ret
	}
	return *o.BindingInfo
}

// GetBindingInfoOk returns a tuple with the BindingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfUeCallbackInfo) GetBindingInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BindingInfo) {
		return nil, false
	}
	return o.BindingInfo, true
}

// HasBindingInfo returns a boolean if a field has been set.
func (o *PcfUeCallbackInfo) HasBindingInfo() bool {
	if o != nil && !IsNil(o.BindingInfo) {
		return true
	}

	return false
}

// SetBindingInfo gets a reference to the given string and assigns it to the BindingInfo field.
func (o *PcfUeCallbackInfo) SetBindingInfo(v string) {
	o.BindingInfo = &v
}

func (o PcfUeCallbackInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfUeCallbackInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["callbackUri"] = o.CallbackUri
	if !IsNil(o.BindingInfo) {
		toSerialize["bindingInfo"] = o.BindingInfo
	}
	return toSerialize, nil
}

type NullablePcfUeCallbackInfo struct {
	value *PcfUeCallbackInfo
	isSet bool
}

func (v NullablePcfUeCallbackInfo) Get() *PcfUeCallbackInfo {
	return v.value
}

func (v *NullablePcfUeCallbackInfo) Set(val *PcfUeCallbackInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfUeCallbackInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfUeCallbackInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfUeCallbackInfo(val *PcfUeCallbackInfo) *NullablePcfUeCallbackInfo {
	return &NullablePcfUeCallbackInfo{value: val, isSet: true}
}

func (v NullablePcfUeCallbackInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfUeCallbackInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


