/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnsacf_NSAC

import (
	"encoding/json"
)

// checks if the UeACRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeACRequestData{}

// UeACRequestData struct for UeACRequestData
type UeACRequestData struct {
	UeACRequestInfo []UeACRequestInfo `json:"ueACRequestInfo"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfId string `json:"nfId"`
	NfType *NFType `json:"nfType,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	EacNotificationUri *string `json:"eacNotificationUri,omitempty"`
}

// NewUeACRequestData instantiates a new UeACRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeACRequestData(ueACRequestInfo []UeACRequestInfo, nfId string) *UeACRequestData {
	this := UeACRequestData{}
	this.UeACRequestInfo = ueACRequestInfo
	this.NfId = nfId
	return &this
}

// NewUeACRequestDataWithDefaults instantiates a new UeACRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeACRequestDataWithDefaults() *UeACRequestData {
	this := UeACRequestData{}
	return &this
}

// GetUeACRequestInfo returns the UeACRequestInfo field value
func (o *UeACRequestData) GetUeACRequestInfo() []UeACRequestInfo {
	if o == nil {
		var ret []UeACRequestInfo
		return ret
	}

	return o.UeACRequestInfo
}

// GetUeACRequestInfoOk returns a tuple with the UeACRequestInfo field value
// and a boolean to check if the value has been set.
func (o *UeACRequestData) GetUeACRequestInfoOk() ([]UeACRequestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.UeACRequestInfo, true
}

// SetUeACRequestInfo sets field value
func (o *UeACRequestData) SetUeACRequestInfo(v []UeACRequestInfo) {
	o.UeACRequestInfo = v
}

// GetNfId returns the NfId field value
func (o *UeACRequestData) GetNfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value
// and a boolean to check if the value has been set.
func (o *UeACRequestData) GetNfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfId, true
}

// SetNfId sets field value
func (o *UeACRequestData) SetNfId(v string) {
	o.NfId = v
}

// GetNfType returns the NfType field value if set, zero value otherwise.
func (o *UeACRequestData) GetNfType() NFType {
	if o == nil || IsNil(o.NfType) {
		var ret NFType
		return ret
	}
	return *o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeACRequestData) GetNfTypeOk() (*NFType, bool) {
	if o == nil || IsNil(o.NfType) {
		return nil, false
	}
	return o.NfType, true
}

// HasNfType returns a boolean if a field has been set.
func (o *UeACRequestData) HasNfType() bool {
	if o != nil && !IsNil(o.NfType) {
		return true
	}

	return false
}

// SetNfType gets a reference to the given NFType and assigns it to the NfType field.
func (o *UeACRequestData) SetNfType(v NFType) {
	o.NfType = &v
}

// GetEacNotificationUri returns the EacNotificationUri field value if set, zero value otherwise.
func (o *UeACRequestData) GetEacNotificationUri() string {
	if o == nil || IsNil(o.EacNotificationUri) {
		var ret string
		return ret
	}
	return *o.EacNotificationUri
}

// GetEacNotificationUriOk returns a tuple with the EacNotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeACRequestData) GetEacNotificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.EacNotificationUri) {
		return nil, false
	}
	return o.EacNotificationUri, true
}

// HasEacNotificationUri returns a boolean if a field has been set.
func (o *UeACRequestData) HasEacNotificationUri() bool {
	if o != nil && !IsNil(o.EacNotificationUri) {
		return true
	}

	return false
}

// SetEacNotificationUri gets a reference to the given string and assigns it to the EacNotificationUri field.
func (o *UeACRequestData) SetEacNotificationUri(v string) {
	o.EacNotificationUri = &v
}

func (o UeACRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeACRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueACRequestInfo"] = o.UeACRequestInfo
	toSerialize["nfId"] = o.NfId
	if !IsNil(o.NfType) {
		toSerialize["nfType"] = o.NfType
	}
	if !IsNil(o.EacNotificationUri) {
		toSerialize["eacNotificationUri"] = o.EacNotificationUri
	}
	return toSerialize, nil
}

type NullableUeACRequestData struct {
	value *UeACRequestData
	isSet bool
}

func (v NullableUeACRequestData) Get() *UeACRequestData {
	return v.value
}

func (v *NullableUeACRequestData) Set(val *UeACRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeACRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeACRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeACRequestData(val *UeACRequestData) *NullableUeACRequestData {
	return &NullableUeACRequestData{value: val, isSet: true}
}

func (v NullableUeACRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeACRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


