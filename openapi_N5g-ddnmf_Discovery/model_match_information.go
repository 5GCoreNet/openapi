/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
)

// checks if the MatchInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchInformation{}

// MatchInformation Represents a report including a matching result, and the information that can be used for charging purpose. 
type MatchInformation struct {
	DiscType DiscoveryType `json:"discType"`
	OpenMatchInfoForOpen *MatchInfoForOpen `json:"openMatchInfoForOpen,omitempty"`
	RestrictedMatchInfo *MatchInfoForRestricted `json:"restrictedMatchInfo,omitempty"`
}

// NewMatchInformation instantiates a new MatchInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchInformation(discType DiscoveryType) *MatchInformation {
	this := MatchInformation{}
	this.DiscType = discType
	return &this
}

// NewMatchInformationWithDefaults instantiates a new MatchInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchInformationWithDefaults() *MatchInformation {
	this := MatchInformation{}
	return &this
}

// GetDiscType returns the DiscType field value
func (o *MatchInformation) GetDiscType() DiscoveryType {
	if o == nil {
		var ret DiscoveryType
		return ret
	}

	return o.DiscType
}

// GetDiscTypeOk returns a tuple with the DiscType field value
// and a boolean to check if the value has been set.
func (o *MatchInformation) GetDiscTypeOk() (*DiscoveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscType, true
}

// SetDiscType sets field value
func (o *MatchInformation) SetDiscType(v DiscoveryType) {
	o.DiscType = v
}

// GetOpenMatchInfoForOpen returns the OpenMatchInfoForOpen field value if set, zero value otherwise.
func (o *MatchInformation) GetOpenMatchInfoForOpen() MatchInfoForOpen {
	if o == nil || isNil(o.OpenMatchInfoForOpen) {
		var ret MatchInfoForOpen
		return ret
	}
	return *o.OpenMatchInfoForOpen
}

// GetOpenMatchInfoForOpenOk returns a tuple with the OpenMatchInfoForOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchInformation) GetOpenMatchInfoForOpenOk() (*MatchInfoForOpen, bool) {
	if o == nil || isNil(o.OpenMatchInfoForOpen) {
		return nil, false
	}
	return o.OpenMatchInfoForOpen, true
}

// HasOpenMatchInfoForOpen returns a boolean if a field has been set.
func (o *MatchInformation) HasOpenMatchInfoForOpen() bool {
	if o != nil && !isNil(o.OpenMatchInfoForOpen) {
		return true
	}

	return false
}

// SetOpenMatchInfoForOpen gets a reference to the given MatchInfoForOpen and assigns it to the OpenMatchInfoForOpen field.
func (o *MatchInformation) SetOpenMatchInfoForOpen(v MatchInfoForOpen) {
	o.OpenMatchInfoForOpen = &v
}

// GetRestrictedMatchInfo returns the RestrictedMatchInfo field value if set, zero value otherwise.
func (o *MatchInformation) GetRestrictedMatchInfo() MatchInfoForRestricted {
	if o == nil || isNil(o.RestrictedMatchInfo) {
		var ret MatchInfoForRestricted
		return ret
	}
	return *o.RestrictedMatchInfo
}

// GetRestrictedMatchInfoOk returns a tuple with the RestrictedMatchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchInformation) GetRestrictedMatchInfoOk() (*MatchInfoForRestricted, bool) {
	if o == nil || isNil(o.RestrictedMatchInfo) {
		return nil, false
	}
	return o.RestrictedMatchInfo, true
}

// HasRestrictedMatchInfo returns a boolean if a field has been set.
func (o *MatchInformation) HasRestrictedMatchInfo() bool {
	if o != nil && !isNil(o.RestrictedMatchInfo) {
		return true
	}

	return false
}

// SetRestrictedMatchInfo gets a reference to the given MatchInfoForRestricted and assigns it to the RestrictedMatchInfo field.
func (o *MatchInformation) SetRestrictedMatchInfo(v MatchInfoForRestricted) {
	o.RestrictedMatchInfo = &v
}

func (o MatchInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discType"] = o.DiscType
	if !isNil(o.OpenMatchInfoForOpen) {
		toSerialize["openMatchInfoForOpen"] = o.OpenMatchInfoForOpen
	}
	if !isNil(o.RestrictedMatchInfo) {
		toSerialize["restrictedMatchInfo"] = o.RestrictedMatchInfo
	}
	return toSerialize, nil
}

type NullableMatchInformation struct {
	value *MatchInformation
	isSet bool
}

func (v NullableMatchInformation) Get() *MatchInformation {
	return v.value
}

func (v *NullableMatchInformation) Set(val *MatchInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchInformation(val *MatchInformation) *NullableMatchInformation {
	return &NullableMatchInformation{value: val, isSet: true}
}

func (v NullableMatchInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


