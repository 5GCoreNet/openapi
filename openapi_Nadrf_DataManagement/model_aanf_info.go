/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the AanfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AanfInfo{}

// AanfInfo Represents the information relative to an AAnF NF Instance.
type AanfInfo struct {
	RoutingIndicators []string `json:"routingIndicators,omitempty"`
}

// NewAanfInfo instantiates a new AanfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAanfInfo() *AanfInfo {
	this := AanfInfo{}
	return &this
}

// NewAanfInfoWithDefaults instantiates a new AanfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAanfInfoWithDefaults() *AanfInfo {
	this := AanfInfo{}
	return &this
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *AanfInfo) GetRoutingIndicators() []string {
	if o == nil || IsNil(o.RoutingIndicators) {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AanfInfo) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoutingIndicators) {
		return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *AanfInfo) HasRoutingIndicators() bool {
	if o != nil && !IsNil(o.RoutingIndicators) {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *AanfInfo) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

func (o AanfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AanfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoutingIndicators) {
		toSerialize["routingIndicators"] = o.RoutingIndicators
	}
	return toSerialize, nil
}

type NullableAanfInfo struct {
	value *AanfInfo
	isSet bool
}

func (v NullableAanfInfo) Get() *AanfInfo {
	return v.value
}

func (v *NullableAanfInfo) Set(val *AanfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAanfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAanfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAanfInfo(val *AanfInfo) *NullableAanfInfo {
	return &NullableAanfInfo{value: val, isSet: true}
}

func (v NullableAanfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAanfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


