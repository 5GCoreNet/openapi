/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the EpsInterworkingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpsInterworkingInfo{}

// EpsInterworkingInfo struct for EpsInterworkingInfo
type EpsInterworkingInfo struct {
	// A map (list of key-value pairs where Dnn serves as key) of EpsIwkPgws
	EpsIwkPgws *map[string]EpsIwkPgw `json:"epsIwkPgws,omitempty"`
}

// NewEpsInterworkingInfo instantiates a new EpsInterworkingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpsInterworkingInfo() *EpsInterworkingInfo {
	this := EpsInterworkingInfo{}
	return &this
}

// NewEpsInterworkingInfoWithDefaults instantiates a new EpsInterworkingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpsInterworkingInfoWithDefaults() *EpsInterworkingInfo {
	this := EpsInterworkingInfo{}
	return &this
}

// GetEpsIwkPgws returns the EpsIwkPgws field value if set, zero value otherwise.
func (o *EpsInterworkingInfo) GetEpsIwkPgws() map[string]EpsIwkPgw {
	if o == nil || IsNil(o.EpsIwkPgws) {
		var ret map[string]EpsIwkPgw
		return ret
	}
	return *o.EpsIwkPgws
}

// GetEpsIwkPgwsOk returns a tuple with the EpsIwkPgws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpsInterworkingInfo) GetEpsIwkPgwsOk() (*map[string]EpsIwkPgw, bool) {
	if o == nil || IsNil(o.EpsIwkPgws) {
		return nil, false
	}
	return o.EpsIwkPgws, true
}

// HasEpsIwkPgws returns a boolean if a field has been set.
func (o *EpsInterworkingInfo) HasEpsIwkPgws() bool {
	if o != nil && !IsNil(o.EpsIwkPgws) {
		return true
	}

	return false
}

// SetEpsIwkPgws gets a reference to the given map[string]EpsIwkPgw and assigns it to the EpsIwkPgws field.
func (o *EpsInterworkingInfo) SetEpsIwkPgws(v map[string]EpsIwkPgw) {
	o.EpsIwkPgws = &v
}

func (o EpsInterworkingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpsInterworkingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EpsIwkPgws) {
		toSerialize["epsIwkPgws"] = o.EpsIwkPgws
	}
	return toSerialize, nil
}

type NullableEpsInterworkingInfo struct {
	value *EpsInterworkingInfo
	isSet bool
}

func (v NullableEpsInterworkingInfo) Get() *EpsInterworkingInfo {
	return v.value
}

func (v *NullableEpsInterworkingInfo) Set(val *EpsInterworkingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsInterworkingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsInterworkingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsInterworkingInfo(val *EpsInterworkingInfo) *NullableEpsInterworkingInfo {
	return &NullableEpsInterworkingInfo{value: val, isSet: true}
}

func (v NullableEpsInterworkingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsInterworkingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
