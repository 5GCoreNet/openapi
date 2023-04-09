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

// checks if the UPFConnectionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UPFConnectionInfo{}

// UPFConnectionInfo struct for UPFConnectionInfo
type UPFConnectionInfo struct {
	UPFIpAddress *string `json:"uPFIpAddress,omitempty"`
	UPFRef       *string `json:"uPFRef,omitempty"`
}

// NewUPFConnectionInfo instantiates a new UPFConnectionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUPFConnectionInfo() *UPFConnectionInfo {
	this := UPFConnectionInfo{}
	return &this
}

// NewUPFConnectionInfoWithDefaults instantiates a new UPFConnectionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUPFConnectionInfoWithDefaults() *UPFConnectionInfo {
	this := UPFConnectionInfo{}
	return &this
}

// GetUPFIpAddress returns the UPFIpAddress field value if set, zero value otherwise.
func (o *UPFConnectionInfo) GetUPFIpAddress() string {
	if o == nil || IsNil(o.UPFIpAddress) {
		var ret string
		return ret
	}
	return *o.UPFIpAddress
}

// GetUPFIpAddressOk returns a tuple with the UPFIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPFConnectionInfo) GetUPFIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.UPFIpAddress) {
		return nil, false
	}
	return o.UPFIpAddress, true
}

// HasUPFIpAddress returns a boolean if a field has been set.
func (o *UPFConnectionInfo) HasUPFIpAddress() bool {
	if o != nil && !IsNil(o.UPFIpAddress) {
		return true
	}

	return false
}

// SetUPFIpAddress gets a reference to the given string and assigns it to the UPFIpAddress field.
func (o *UPFConnectionInfo) SetUPFIpAddress(v string) {
	o.UPFIpAddress = &v
}

// GetUPFRef returns the UPFRef field value if set, zero value otherwise.
func (o *UPFConnectionInfo) GetUPFRef() string {
	if o == nil || IsNil(o.UPFRef) {
		var ret string
		return ret
	}
	return *o.UPFRef
}

// GetUPFRefOk returns a tuple with the UPFRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPFConnectionInfo) GetUPFRefOk() (*string, bool) {
	if o == nil || IsNil(o.UPFRef) {
		return nil, false
	}
	return o.UPFRef, true
}

// HasUPFRef returns a boolean if a field has been set.
func (o *UPFConnectionInfo) HasUPFRef() bool {
	if o != nil && !IsNil(o.UPFRef) {
		return true
	}

	return false
}

// SetUPFRef gets a reference to the given string and assigns it to the UPFRef field.
func (o *UPFConnectionInfo) SetUPFRef(v string) {
	o.UPFRef = &v
}

func (o UPFConnectionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UPFConnectionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UPFIpAddress) {
		toSerialize["uPFIpAddress"] = o.UPFIpAddress
	}
	if !IsNil(o.UPFRef) {
		toSerialize["uPFRef"] = o.UPFRef
	}
	return toSerialize, nil
}

type NullableUPFConnectionInfo struct {
	value *UPFConnectionInfo
	isSet bool
}

func (v NullableUPFConnectionInfo) Get() *UPFConnectionInfo {
	return v.value
}

func (v *NullableUPFConnectionInfo) Set(val *UPFConnectionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUPFConnectionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUPFConnectionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPFConnectionInfo(val *UPFConnectionInfo) *NullableUPFConnectionInfo {
	return &NullableUPFConnectionInfo{value: val, isSet: true}
}

func (v NullableUPFConnectionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPFConnectionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
