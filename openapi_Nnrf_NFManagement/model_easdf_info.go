/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the EasdfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasdfInfo{}

// EasdfInfo Information of an EASDF NF Instance
type EasdfInfo struct {
	SNssaiEasdfInfoList []SnssaiEasdfInfoItem `json:"sNssaiEasdfInfoList,omitempty"`
	EasdfN6IpAddressList []IpAddr `json:"easdfN6IpAddressList,omitempty"`
	UpfN6IpAddressList []IpAddr `json:"upfN6IpAddressList,omitempty"`
}

// NewEasdfInfo instantiates a new EasdfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasdfInfo() *EasdfInfo {
	this := EasdfInfo{}
	return &this
}

// NewEasdfInfoWithDefaults instantiates a new EasdfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasdfInfoWithDefaults() *EasdfInfo {
	this := EasdfInfo{}
	return &this
}

// GetSNssaiEasdfInfoList returns the SNssaiEasdfInfoList field value if set, zero value otherwise.
func (o *EasdfInfo) GetSNssaiEasdfInfoList() []SnssaiEasdfInfoItem {
	if o == nil || IsNil(o.SNssaiEasdfInfoList) {
		var ret []SnssaiEasdfInfoItem
		return ret
	}
	return o.SNssaiEasdfInfoList
}

// GetSNssaiEasdfInfoListOk returns a tuple with the SNssaiEasdfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasdfInfo) GetSNssaiEasdfInfoListOk() ([]SnssaiEasdfInfoItem, bool) {
	if o == nil || IsNil(o.SNssaiEasdfInfoList) {
		return nil, false
	}
	return o.SNssaiEasdfInfoList, true
}

// HasSNssaiEasdfInfoList returns a boolean if a field has been set.
func (o *EasdfInfo) HasSNssaiEasdfInfoList() bool {
	if o != nil && !IsNil(o.SNssaiEasdfInfoList) {
		return true
	}

	return false
}

// SetSNssaiEasdfInfoList gets a reference to the given []SnssaiEasdfInfoItem and assigns it to the SNssaiEasdfInfoList field.
func (o *EasdfInfo) SetSNssaiEasdfInfoList(v []SnssaiEasdfInfoItem) {
	o.SNssaiEasdfInfoList = v
}

// GetEasdfN6IpAddressList returns the EasdfN6IpAddressList field value if set, zero value otherwise.
func (o *EasdfInfo) GetEasdfN6IpAddressList() []IpAddr {
	if o == nil || IsNil(o.EasdfN6IpAddressList) {
		var ret []IpAddr
		return ret
	}
	return o.EasdfN6IpAddressList
}

// GetEasdfN6IpAddressListOk returns a tuple with the EasdfN6IpAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasdfInfo) GetEasdfN6IpAddressListOk() ([]IpAddr, bool) {
	if o == nil || IsNil(o.EasdfN6IpAddressList) {
		return nil, false
	}
	return o.EasdfN6IpAddressList, true
}

// HasEasdfN6IpAddressList returns a boolean if a field has been set.
func (o *EasdfInfo) HasEasdfN6IpAddressList() bool {
	if o != nil && !IsNil(o.EasdfN6IpAddressList) {
		return true
	}

	return false
}

// SetEasdfN6IpAddressList gets a reference to the given []IpAddr and assigns it to the EasdfN6IpAddressList field.
func (o *EasdfInfo) SetEasdfN6IpAddressList(v []IpAddr) {
	o.EasdfN6IpAddressList = v
}

// GetUpfN6IpAddressList returns the UpfN6IpAddressList field value if set, zero value otherwise.
func (o *EasdfInfo) GetUpfN6IpAddressList() []IpAddr {
	if o == nil || IsNil(o.UpfN6IpAddressList) {
		var ret []IpAddr
		return ret
	}
	return o.UpfN6IpAddressList
}

// GetUpfN6IpAddressListOk returns a tuple with the UpfN6IpAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasdfInfo) GetUpfN6IpAddressListOk() ([]IpAddr, bool) {
	if o == nil || IsNil(o.UpfN6IpAddressList) {
		return nil, false
	}
	return o.UpfN6IpAddressList, true
}

// HasUpfN6IpAddressList returns a boolean if a field has been set.
func (o *EasdfInfo) HasUpfN6IpAddressList() bool {
	if o != nil && !IsNil(o.UpfN6IpAddressList) {
		return true
	}

	return false
}

// SetUpfN6IpAddressList gets a reference to the given []IpAddr and assigns it to the UpfN6IpAddressList field.
func (o *EasdfInfo) SetUpfN6IpAddressList(v []IpAddr) {
	o.UpfN6IpAddressList = v
}

func (o EasdfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasdfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SNssaiEasdfInfoList) {
		toSerialize["sNssaiEasdfInfoList"] = o.SNssaiEasdfInfoList
	}
	if !IsNil(o.EasdfN6IpAddressList) {
		toSerialize["easdfN6IpAddressList"] = o.EasdfN6IpAddressList
	}
	if !IsNil(o.UpfN6IpAddressList) {
		toSerialize["upfN6IpAddressList"] = o.UpfN6IpAddressList
	}
	return toSerialize, nil
}

type NullableEasdfInfo struct {
	value *EasdfInfo
	isSet bool
}

func (v NullableEasdfInfo) Get() *EasdfInfo {
	return v.value
}

func (v *NullableEasdfInfo) Set(val *EasdfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEasdfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEasdfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasdfInfo(val *EasdfInfo) *NullableEasdfInfo {
	return &NullableEasdfInfo{value: val, isSet: true}
}

func (v NullableEasdfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasdfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


