/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the UpfInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfInformation{}

// UpfInformation Represents the ID/address/FQDN of the UPF.
type UpfInformation struct {
	UpfId   *string   `json:"upfId,omitempty"`
	UpfAddr *AddrFqdn `json:"upfAddr,omitempty"`
}

// NewUpfInformation instantiates a new UpfInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfInformation() *UpfInformation {
	this := UpfInformation{}
	return &this
}

// NewUpfInformationWithDefaults instantiates a new UpfInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfInformationWithDefaults() *UpfInformation {
	this := UpfInformation{}
	return &this
}

// GetUpfId returns the UpfId field value if set, zero value otherwise.
func (o *UpfInformation) GetUpfId() string {
	if o == nil || IsNil(o.UpfId) {
		var ret string
		return ret
	}
	return *o.UpfId
}

// GetUpfIdOk returns a tuple with the UpfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInformation) GetUpfIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpfId) {
		return nil, false
	}
	return o.UpfId, true
}

// HasUpfId returns a boolean if a field has been set.
func (o *UpfInformation) HasUpfId() bool {
	if o != nil && !IsNil(o.UpfId) {
		return true
	}

	return false
}

// SetUpfId gets a reference to the given string and assigns it to the UpfId field.
func (o *UpfInformation) SetUpfId(v string) {
	o.UpfId = &v
}

// GetUpfAddr returns the UpfAddr field value if set, zero value otherwise.
func (o *UpfInformation) GetUpfAddr() AddrFqdn {
	if o == nil || IsNil(o.UpfAddr) {
		var ret AddrFqdn
		return ret
	}
	return *o.UpfAddr
}

// GetUpfAddrOk returns a tuple with the UpfAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInformation) GetUpfAddrOk() (*AddrFqdn, bool) {
	if o == nil || IsNil(o.UpfAddr) {
		return nil, false
	}
	return o.UpfAddr, true
}

// HasUpfAddr returns a boolean if a field has been set.
func (o *UpfInformation) HasUpfAddr() bool {
	if o != nil && !IsNil(o.UpfAddr) {
		return true
	}

	return false
}

// SetUpfAddr gets a reference to the given AddrFqdn and assigns it to the UpfAddr field.
func (o *UpfInformation) SetUpfAddr(v AddrFqdn) {
	o.UpfAddr = &v
}

func (o UpfInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpfId) {
		toSerialize["upfId"] = o.UpfId
	}
	if !IsNil(o.UpfAddr) {
		toSerialize["upfAddr"] = o.UpfAddr
	}
	return toSerialize, nil
}

type NullableUpfInformation struct {
	value *UpfInformation
	isSet bool
}

func (v NullableUpfInformation) Get() *UpfInformation {
	return v.value
}

func (v *NullableUpfInformation) Set(val *UpfInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfInformation(val *UpfInformation) *NullableUpfInformation {
	return &NullableUpfInformation{value: val, isSet: true}
}

func (v NullableUpfInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
