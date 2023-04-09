/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the NsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsInfo{}

// NsInfo struct for NsInfo
type NsInfo struct {
	NsInstanceId *string `json:"nsInstanceId,omitempty"`
	NsName       *string `json:"nsName,omitempty"`
}

// NewNsInfo instantiates a new NsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsInfo() *NsInfo {
	this := NsInfo{}
	return &this
}

// NewNsInfoWithDefaults instantiates a new NsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsInfoWithDefaults() *NsInfo {
	this := NsInfo{}
	return &this
}

// GetNsInstanceId returns the NsInstanceId field value if set, zero value otherwise.
func (o *NsInfo) GetNsInstanceId() string {
	if o == nil || IsNil(o.NsInstanceId) {
		var ret string
		return ret
	}
	return *o.NsInstanceId
}

// GetNsInstanceIdOk returns a tuple with the NsInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsInfo) GetNsInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NsInstanceId) {
		return nil, false
	}
	return o.NsInstanceId, true
}

// HasNsInstanceId returns a boolean if a field has been set.
func (o *NsInfo) HasNsInstanceId() bool {
	if o != nil && !IsNil(o.NsInstanceId) {
		return true
	}

	return false
}

// SetNsInstanceId gets a reference to the given string and assigns it to the NsInstanceId field.
func (o *NsInfo) SetNsInstanceId(v string) {
	o.NsInstanceId = &v
}

// GetNsName returns the NsName field value if set, zero value otherwise.
func (o *NsInfo) GetNsName() string {
	if o == nil || IsNil(o.NsName) {
		var ret string
		return ret
	}
	return *o.NsName
}

// GetNsNameOk returns a tuple with the NsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsInfo) GetNsNameOk() (*string, bool) {
	if o == nil || IsNil(o.NsName) {
		return nil, false
	}
	return o.NsName, true
}

// HasNsName returns a boolean if a field has been set.
func (o *NsInfo) HasNsName() bool {
	if o != nil && !IsNil(o.NsName) {
		return true
	}

	return false
}

// SetNsName gets a reference to the given string and assigns it to the NsName field.
func (o *NsInfo) SetNsName(v string) {
	o.NsName = &v
}

func (o NsInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NsInstanceId) {
		toSerialize["nsInstanceId"] = o.NsInstanceId
	}
	if !IsNil(o.NsName) {
		toSerialize["nsName"] = o.NsName
	}
	return toSerialize, nil
}

type NullableNsInfo struct {
	value *NsInfo
	isSet bool
}

func (v NullableNsInfo) Get() *NsInfo {
	return v.value
}

func (v *NullableNsInfo) Set(val *NsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsInfo(val *NsInfo) *NullableNsInfo {
	return &NullableNsInfo{value: val, isSet: true}
}

func (v NullableNsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
