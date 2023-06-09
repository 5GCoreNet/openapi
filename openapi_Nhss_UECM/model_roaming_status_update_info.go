/*
Nhss_UECM

HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UECM

import (
	"encoding/json"
)

// checks if the RoamingStatusUpdateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoamingStatusUpdateInfo{}

// RoamingStatusUpdateInfo It represents the request body of the Roaming Status Update request sent by UDM to HSS, and contains the IMSI of the UE and the new PLMN-ID
type RoamingStatusUpdateInfo struct {
	Imsi   string `json:"imsi"`
	PlmnId PlmnId `json:"plmnId"`
}

// NewRoamingStatusUpdateInfo instantiates a new RoamingStatusUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoamingStatusUpdateInfo(imsi string, plmnId PlmnId) *RoamingStatusUpdateInfo {
	this := RoamingStatusUpdateInfo{}
	this.Imsi = imsi
	this.PlmnId = plmnId
	return &this
}

// NewRoamingStatusUpdateInfoWithDefaults instantiates a new RoamingStatusUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoamingStatusUpdateInfoWithDefaults() *RoamingStatusUpdateInfo {
	this := RoamingStatusUpdateInfo{}
	return &this
}

// GetImsi returns the Imsi field value
func (o *RoamingStatusUpdateInfo) GetImsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value
// and a boolean to check if the value has been set.
func (o *RoamingStatusUpdateInfo) GetImsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imsi, true
}

// SetImsi sets field value
func (o *RoamingStatusUpdateInfo) SetImsi(v string) {
	o.Imsi = v
}

// GetPlmnId returns the PlmnId field value
func (o *RoamingStatusUpdateInfo) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *RoamingStatusUpdateInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *RoamingStatusUpdateInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

func (o RoamingStatusUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoamingStatusUpdateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imsi"] = o.Imsi
	toSerialize["plmnId"] = o.PlmnId
	return toSerialize, nil
}

type NullableRoamingStatusUpdateInfo struct {
	value *RoamingStatusUpdateInfo
	isSet bool
}

func (v NullableRoamingStatusUpdateInfo) Get() *RoamingStatusUpdateInfo {
	return v.value
}

func (v *NullableRoamingStatusUpdateInfo) Set(val *RoamingStatusUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingStatusUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingStatusUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingStatusUpdateInfo(val *RoamingStatusUpdateInfo) *NullableRoamingStatusUpdateInfo {
	return &NullableRoamingStatusUpdateInfo{value: val, isSet: true}
}

func (v NullableRoamingStatusUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingStatusUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
