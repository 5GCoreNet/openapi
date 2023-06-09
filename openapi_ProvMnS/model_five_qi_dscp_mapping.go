/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the FiveQiDscpMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiveQiDscpMapping{}

// FiveQiDscpMapping struct for FiveQiDscpMapping
type FiveQiDscpMapping struct {
	FiveQIValues []int32 `json:"fiveQIValues,omitempty"`
	Dscp         *int32  `json:"dscp,omitempty"`
}

// NewFiveQiDscpMapping instantiates a new FiveQiDscpMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiveQiDscpMapping() *FiveQiDscpMapping {
	this := FiveQiDscpMapping{}
	return &this
}

// NewFiveQiDscpMappingWithDefaults instantiates a new FiveQiDscpMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiveQiDscpMappingWithDefaults() *FiveQiDscpMapping {
	this := FiveQiDscpMapping{}
	return &this
}

// GetFiveQIValues returns the FiveQIValues field value if set, zero value otherwise.
func (o *FiveQiDscpMapping) GetFiveQIValues() []int32 {
	if o == nil || IsNil(o.FiveQIValues) {
		var ret []int32
		return ret
	}
	return o.FiveQIValues
}

// GetFiveQIValuesOk returns a tuple with the FiveQIValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQiDscpMapping) GetFiveQIValuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.FiveQIValues) {
		return nil, false
	}
	return o.FiveQIValues, true
}

// HasFiveQIValues returns a boolean if a field has been set.
func (o *FiveQiDscpMapping) HasFiveQIValues() bool {
	if o != nil && !IsNil(o.FiveQIValues) {
		return true
	}

	return false
}

// SetFiveQIValues gets a reference to the given []int32 and assigns it to the FiveQIValues field.
func (o *FiveQiDscpMapping) SetFiveQIValues(v []int32) {
	o.FiveQIValues = v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *FiveQiDscpMapping) GetDscp() int32 {
	if o == nil || IsNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQiDscpMapping) GetDscpOk() (*int32, bool) {
	if o == nil || IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *FiveQiDscpMapping) HasDscp() bool {
	if o != nil && !IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *FiveQiDscpMapping) SetDscp(v int32) {
	o.Dscp = &v
}

func (o FiveQiDscpMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiveQiDscpMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiveQIValues) {
		toSerialize["fiveQIValues"] = o.FiveQIValues
	}
	if !IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return toSerialize, nil
}

type NullableFiveQiDscpMapping struct {
	value *FiveQiDscpMapping
	isSet bool
}

func (v NullableFiveQiDscpMapping) Get() *FiveQiDscpMapping {
	return v.value
}

func (v *NullableFiveQiDscpMapping) Set(val *FiveQiDscpMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableFiveQiDscpMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableFiveQiDscpMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiveQiDscpMapping(val *FiveQiDscpMapping) *NullableFiveQiDscpMapping {
	return &NullableFiveQiDscpMapping{value: val, isSet: true}
}

func (v NullableFiveQiDscpMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiveQiDscpMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
