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

// checks if the AppDetectionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDetectionInfo{}

// AppDetectionInfo Contains the detected application's traffic information.
type AppDetectionInfo struct {
	// A reference to the application detection filter configured at the UPF
	AppId string `json:"appId"`
	// Identifier sent by the SMF in order to allow correlation of application Start and Stop events to the specific service data flow description, if service data flow descriptions are deducible. 
	InstanceId *string `json:"instanceId,omitempty"`
	// Contains the detected service data flow descriptions if they are deducible.
	SdfDescriptions []FlowInformation `json:"sdfDescriptions,omitempty"`
}

// NewAppDetectionInfo instantiates a new AppDetectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDetectionInfo(appId string) *AppDetectionInfo {
	this := AppDetectionInfo{}
	this.AppId = appId
	return &this
}

// NewAppDetectionInfoWithDefaults instantiates a new AppDetectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDetectionInfoWithDefaults() *AppDetectionInfo {
	this := AppDetectionInfo{}
	return &this
}

// GetAppId returns the AppId field value
func (o *AppDetectionInfo) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AppDetectionInfo) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppDetectionInfo) SetAppId(v string) {
	o.AppId = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *AppDetectionInfo) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetectionInfo) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *AppDetectionInfo) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *AppDetectionInfo) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetSdfDescriptions returns the SdfDescriptions field value if set, zero value otherwise.
func (o *AppDetectionInfo) GetSdfDescriptions() []FlowInformation {
	if o == nil || IsNil(o.SdfDescriptions) {
		var ret []FlowInformation
		return ret
	}
	return o.SdfDescriptions
}

// GetSdfDescriptionsOk returns a tuple with the SdfDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDetectionInfo) GetSdfDescriptionsOk() ([]FlowInformation, bool) {
	if o == nil || IsNil(o.SdfDescriptions) {
		return nil, false
	}
	return o.SdfDescriptions, true
}

// HasSdfDescriptions returns a boolean if a field has been set.
func (o *AppDetectionInfo) HasSdfDescriptions() bool {
	if o != nil && !IsNil(o.SdfDescriptions) {
		return true
	}

	return false
}

// SetSdfDescriptions gets a reference to the given []FlowInformation and assigns it to the SdfDescriptions field.
func (o *AppDetectionInfo) SetSdfDescriptions(v []FlowInformation) {
	o.SdfDescriptions = v
}

func (o AppDetectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDetectionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.SdfDescriptions) {
		toSerialize["sdfDescriptions"] = o.SdfDescriptions
	}
	return toSerialize, nil
}

type NullableAppDetectionInfo struct {
	value *AppDetectionInfo
	isSet bool
}

func (v NullableAppDetectionInfo) Get() *AppDetectionInfo {
	return v.value
}

func (v *NullableAppDetectionInfo) Set(val *AppDetectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetectionInfo(val *AppDetectionInfo) *NullableAppDetectionInfo {
	return &NullableAppDetectionInfo{value: val, isSet: true}
}

func (v NullableAppDetectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


