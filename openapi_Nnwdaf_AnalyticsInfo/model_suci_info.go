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

// checks if the SuciInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuciInfo{}

// SuciInfo SUCI information containing Routing Indicator and Home Network Public Key ID
type SuciInfo struct {
	RoutingInds []string `json:"routingInds,omitempty"`
	HNwPubKeyIds []int32 `json:"hNwPubKeyIds,omitempty"`
}

// NewSuciInfo instantiates a new SuciInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuciInfo() *SuciInfo {
	this := SuciInfo{}
	return &this
}

// NewSuciInfoWithDefaults instantiates a new SuciInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuciInfoWithDefaults() *SuciInfo {
	this := SuciInfo{}
	return &this
}

// GetRoutingInds returns the RoutingInds field value if set, zero value otherwise.
func (o *SuciInfo) GetRoutingInds() []string {
	if o == nil || IsNil(o.RoutingInds) {
		var ret []string
		return ret
	}
	return o.RoutingInds
}

// GetRoutingIndsOk returns a tuple with the RoutingInds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuciInfo) GetRoutingIndsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoutingInds) {
		return nil, false
	}
	return o.RoutingInds, true
}

// HasRoutingInds returns a boolean if a field has been set.
func (o *SuciInfo) HasRoutingInds() bool {
	if o != nil && !IsNil(o.RoutingInds) {
		return true
	}

	return false
}

// SetRoutingInds gets a reference to the given []string and assigns it to the RoutingInds field.
func (o *SuciInfo) SetRoutingInds(v []string) {
	o.RoutingInds = v
}

// GetHNwPubKeyIds returns the HNwPubKeyIds field value if set, zero value otherwise.
func (o *SuciInfo) GetHNwPubKeyIds() []int32 {
	if o == nil || IsNil(o.HNwPubKeyIds) {
		var ret []int32
		return ret
	}
	return o.HNwPubKeyIds
}

// GetHNwPubKeyIdsOk returns a tuple with the HNwPubKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuciInfo) GetHNwPubKeyIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.HNwPubKeyIds) {
		return nil, false
	}
	return o.HNwPubKeyIds, true
}

// HasHNwPubKeyIds returns a boolean if a field has been set.
func (o *SuciInfo) HasHNwPubKeyIds() bool {
	if o != nil && !IsNil(o.HNwPubKeyIds) {
		return true
	}

	return false
}

// SetHNwPubKeyIds gets a reference to the given []int32 and assigns it to the HNwPubKeyIds field.
func (o *SuciInfo) SetHNwPubKeyIds(v []int32) {
	o.HNwPubKeyIds = v
}

func (o SuciInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuciInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoutingInds) {
		toSerialize["routingInds"] = o.RoutingInds
	}
	if !IsNil(o.HNwPubKeyIds) {
		toSerialize["hNwPubKeyIds"] = o.HNwPubKeyIds
	}
	return toSerialize, nil
}

type NullableSuciInfo struct {
	value *SuciInfo
	isSet bool
}

func (v NullableSuciInfo) Get() *SuciInfo {
	return v.value
}

func (v *NullableSuciInfo) Set(val *SuciInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSuciInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSuciInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuciInfo(val *SuciInfo) *NullableSuciInfo {
	return &NullableSuciInfo{value: val, isSet: true}
}

func (v NullableSuciInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuciInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


