/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the RankingCriterion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankingCriterion{}

// RankingCriterion Indicates the usage ranking criterion between the high, medium and low usage UE.
type RankingCriterion struct {
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	HighBase int32 `json:"highBase"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	LowBase int32 `json:"lowBase"`
}

// NewRankingCriterion instantiates a new RankingCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankingCriterion(highBase int32, lowBase int32) *RankingCriterion {
	this := RankingCriterion{}
	this.HighBase = highBase
	this.LowBase = lowBase
	return &this
}

// NewRankingCriterionWithDefaults instantiates a new RankingCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankingCriterionWithDefaults() *RankingCriterion {
	this := RankingCriterion{}
	return &this
}

// GetHighBase returns the HighBase field value
func (o *RankingCriterion) GetHighBase() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HighBase
}

// GetHighBaseOk returns a tuple with the HighBase field value
// and a boolean to check if the value has been set.
func (o *RankingCriterion) GetHighBaseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighBase, true
}

// SetHighBase sets field value
func (o *RankingCriterion) SetHighBase(v int32) {
	o.HighBase = v
}

// GetLowBase returns the LowBase field value
func (o *RankingCriterion) GetLowBase() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LowBase
}

// GetLowBaseOk returns a tuple with the LowBase field value
// and a boolean to check if the value has been set.
func (o *RankingCriterion) GetLowBaseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowBase, true
}

// SetLowBase sets field value
func (o *RankingCriterion) SetLowBase(v int32) {
	o.LowBase = v
}

func (o RankingCriterion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RankingCriterion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["highBase"] = o.HighBase
	toSerialize["lowBase"] = o.LowBase
	return toSerialize, nil
}

type NullableRankingCriterion struct {
	value *RankingCriterion
	isSet bool
}

func (v NullableRankingCriterion) Get() *RankingCriterion {
	return v.value
}

func (v *NullableRankingCriterion) Set(val *RankingCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableRankingCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableRankingCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankingCriterion(val *RankingCriterion) *NullableRankingCriterion {
	return &NullableRankingCriterion{value: val, isSet: true}
}

func (v NullableRankingCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRankingCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


