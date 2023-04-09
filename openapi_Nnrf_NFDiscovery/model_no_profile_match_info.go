/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the NoProfileMatchInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoProfileMatchInfo{}

// NoProfileMatchInfo Provides the reason for not finding NF matching the search criteria
type NoProfileMatchInfo struct {
	Reason                    NoProfileMatchReason    `json:"reason"`
	QueryParamCombinationList []QueryParamCombination `json:"queryParamCombinationList,omitempty"`
}

// NewNoProfileMatchInfo instantiates a new NoProfileMatchInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoProfileMatchInfo(reason NoProfileMatchReason) *NoProfileMatchInfo {
	this := NoProfileMatchInfo{}
	this.Reason = reason
	return &this
}

// NewNoProfileMatchInfoWithDefaults instantiates a new NoProfileMatchInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoProfileMatchInfoWithDefaults() *NoProfileMatchInfo {
	this := NoProfileMatchInfo{}
	return &this
}

// GetReason returns the Reason field value
func (o *NoProfileMatchInfo) GetReason() NoProfileMatchReason {
	if o == nil {
		var ret NoProfileMatchReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *NoProfileMatchInfo) GetReasonOk() (*NoProfileMatchReason, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *NoProfileMatchInfo) SetReason(v NoProfileMatchReason) {
	o.Reason = v
}

// GetQueryParamCombinationList returns the QueryParamCombinationList field value if set, zero value otherwise.
func (o *NoProfileMatchInfo) GetQueryParamCombinationList() []QueryParamCombination {
	if o == nil || IsNil(o.QueryParamCombinationList) {
		var ret []QueryParamCombination
		return ret
	}
	return o.QueryParamCombinationList
}

// GetQueryParamCombinationListOk returns a tuple with the QueryParamCombinationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoProfileMatchInfo) GetQueryParamCombinationListOk() ([]QueryParamCombination, bool) {
	if o == nil || IsNil(o.QueryParamCombinationList) {
		return nil, false
	}
	return o.QueryParamCombinationList, true
}

// HasQueryParamCombinationList returns a boolean if a field has been set.
func (o *NoProfileMatchInfo) HasQueryParamCombinationList() bool {
	if o != nil && !IsNil(o.QueryParamCombinationList) {
		return true
	}

	return false
}

// SetQueryParamCombinationList gets a reference to the given []QueryParamCombination and assigns it to the QueryParamCombinationList field.
func (o *NoProfileMatchInfo) SetQueryParamCombinationList(v []QueryParamCombination) {
	o.QueryParamCombinationList = v
}

func (o NoProfileMatchInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoProfileMatchInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.QueryParamCombinationList) {
		toSerialize["queryParamCombinationList"] = o.QueryParamCombinationList
	}
	return toSerialize, nil
}

type NullableNoProfileMatchInfo struct {
	value *NoProfileMatchInfo
	isSet bool
}

func (v NullableNoProfileMatchInfo) Get() *NoProfileMatchInfo {
	return v.value
}

func (v *NullableNoProfileMatchInfo) Set(val *NoProfileMatchInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNoProfileMatchInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNoProfileMatchInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoProfileMatchInfo(val *NoProfileMatchInfo) *NullableNoProfileMatchInfo {
	return &NullableNoProfileMatchInfo{value: val, isSet: true}
}

func (v NullableNoProfileMatchInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoProfileMatchInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
