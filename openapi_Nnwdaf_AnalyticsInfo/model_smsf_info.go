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

// checks if the SmsfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsfInfo{}

// SmsfInfo Specific Data for SMSF
type SmsfInfo struct {
	RoamingUeInd        *bool       `json:"roamingUeInd,omitempty"`
	RemotePlmnRangeList []PlmnRange `json:"remotePlmnRangeList,omitempty"`
}

// NewSmsfInfo instantiates a new SmsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsfInfo() *SmsfInfo {
	this := SmsfInfo{}
	return &this
}

// NewSmsfInfoWithDefaults instantiates a new SmsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsfInfoWithDefaults() *SmsfInfo {
	this := SmsfInfo{}
	return &this
}

// GetRoamingUeInd returns the RoamingUeInd field value if set, zero value otherwise.
func (o *SmsfInfo) GetRoamingUeInd() bool {
	if o == nil || IsNil(o.RoamingUeInd) {
		var ret bool
		return ret
	}
	return *o.RoamingUeInd
}

// GetRoamingUeIndOk returns a tuple with the RoamingUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfInfo) GetRoamingUeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.RoamingUeInd) {
		return nil, false
	}
	return o.RoamingUeInd, true
}

// HasRoamingUeInd returns a boolean if a field has been set.
func (o *SmsfInfo) HasRoamingUeInd() bool {
	if o != nil && !IsNil(o.RoamingUeInd) {
		return true
	}

	return false
}

// SetRoamingUeInd gets a reference to the given bool and assigns it to the RoamingUeInd field.
func (o *SmsfInfo) SetRoamingUeInd(v bool) {
	o.RoamingUeInd = &v
}

// GetRemotePlmnRangeList returns the RemotePlmnRangeList field value if set, zero value otherwise.
func (o *SmsfInfo) GetRemotePlmnRangeList() []PlmnRange {
	if o == nil || IsNil(o.RemotePlmnRangeList) {
		var ret []PlmnRange
		return ret
	}
	return o.RemotePlmnRangeList
}

// GetRemotePlmnRangeListOk returns a tuple with the RemotePlmnRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfInfo) GetRemotePlmnRangeListOk() ([]PlmnRange, bool) {
	if o == nil || IsNil(o.RemotePlmnRangeList) {
		return nil, false
	}
	return o.RemotePlmnRangeList, true
}

// HasRemotePlmnRangeList returns a boolean if a field has been set.
func (o *SmsfInfo) HasRemotePlmnRangeList() bool {
	if o != nil && !IsNil(o.RemotePlmnRangeList) {
		return true
	}

	return false
}

// SetRemotePlmnRangeList gets a reference to the given []PlmnRange and assigns it to the RemotePlmnRangeList field.
func (o *SmsfInfo) SetRemotePlmnRangeList(v []PlmnRange) {
	o.RemotePlmnRangeList = v
}

func (o SmsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoamingUeInd) {
		toSerialize["roamingUeInd"] = o.RoamingUeInd
	}
	if !IsNil(o.RemotePlmnRangeList) {
		toSerialize["remotePlmnRangeList"] = o.RemotePlmnRangeList
	}
	return toSerialize, nil
}

type NullableSmsfInfo struct {
	value *SmsfInfo
	isSet bool
}

func (v NullableSmsfInfo) Get() *SmsfInfo {
	return v.value
}

func (v *NullableSmsfInfo) Set(val *SmsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsfInfo(val *SmsfInfo) *NullableSmsfInfo {
	return &NullableSmsfInfo{value: val, isSet: true}
}

func (v NullableSmsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
