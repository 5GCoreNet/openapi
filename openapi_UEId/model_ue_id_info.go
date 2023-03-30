/*
3gpp-ueid

API for UE ID service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UEId

import (
	"encoding/json"
)

// checks if the UeIdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeIdInfo{}

// UeIdInfo Represents UE ID information.
type UeIdInfo struct {
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId"`
}

// NewUeIdInfo instantiates a new UeIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeIdInfo(externalId string) *UeIdInfo {
	this := UeIdInfo{}
	this.ExternalId = externalId
	return &this
}

// NewUeIdInfoWithDefaults instantiates a new UeIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeIdInfoWithDefaults() *UeIdInfo {
	this := UeIdInfo{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *UeIdInfo) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *UeIdInfo) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *UeIdInfo) SetExternalId(v string) {
	o.ExternalId = v
}

func (o UeIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeIdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	return toSerialize, nil
}

type NullableUeIdInfo struct {
	value *UeIdInfo
	isSet bool
}

func (v NullableUeIdInfo) Get() *UeIdInfo {
	return v.value
}

func (v *NullableUeIdInfo) Set(val *UeIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeIdInfo(val *UeIdInfo) *NullableUeIdInfo {
	return &NullableUeIdInfo{value: val, isSet: true}
}

func (v NullableUeIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

