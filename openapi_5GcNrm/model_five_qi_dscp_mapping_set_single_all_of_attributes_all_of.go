/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the FiveQiDscpMappingSetSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiveQiDscpMappingSetSingleAllOfAttributesAllOf{}

// FiveQiDscpMappingSetSingleAllOfAttributesAllOf struct for FiveQiDscpMappingSetSingleAllOfAttributesAllOf
type FiveQiDscpMappingSetSingleAllOfAttributesAllOf struct {
	FiveQiDscpMappingList []FiveQiDscpMapping `json:"FiveQiDscpMappingList,omitempty"`
}

// NewFiveQiDscpMappingSetSingleAllOfAttributesAllOf instantiates a new FiveQiDscpMappingSetSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiveQiDscpMappingSetSingleAllOfAttributesAllOf() *FiveQiDscpMappingSetSingleAllOfAttributesAllOf {
	this := FiveQiDscpMappingSetSingleAllOfAttributesAllOf{}
	return &this
}

// NewFiveQiDscpMappingSetSingleAllOfAttributesAllOfWithDefaults instantiates a new FiveQiDscpMappingSetSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiveQiDscpMappingSetSingleAllOfAttributesAllOfWithDefaults() *FiveQiDscpMappingSetSingleAllOfAttributesAllOf {
	this := FiveQiDscpMappingSetSingleAllOfAttributesAllOf{}
	return &this
}

// GetFiveQiDscpMappingList returns the FiveQiDscpMappingList field value if set, zero value otherwise.
func (o *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) GetFiveQiDscpMappingList() []FiveQiDscpMapping {
	if o == nil || IsNil(o.FiveQiDscpMappingList) {
		var ret []FiveQiDscpMapping
		return ret
	}
	return o.FiveQiDscpMappingList
}

// GetFiveQiDscpMappingListOk returns a tuple with the FiveQiDscpMappingList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) GetFiveQiDscpMappingListOk() ([]FiveQiDscpMapping, bool) {
	if o == nil || IsNil(o.FiveQiDscpMappingList) {
		return nil, false
	}
	return o.FiveQiDscpMappingList, true
}

// HasFiveQiDscpMappingList returns a boolean if a field has been set.
func (o *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) HasFiveQiDscpMappingList() bool {
	if o != nil && !IsNil(o.FiveQiDscpMappingList) {
		return true
	}

	return false
}

// SetFiveQiDscpMappingList gets a reference to the given []FiveQiDscpMapping and assigns it to the FiveQiDscpMappingList field.
func (o *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) SetFiveQiDscpMappingList(v []FiveQiDscpMapping) {
	o.FiveQiDscpMappingList = v
}

func (o FiveQiDscpMappingSetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiveQiDscpMappingSetSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiveQiDscpMappingList) {
		toSerialize["FiveQiDscpMappingList"] = o.FiveQiDscpMappingList
	}
	return toSerialize, nil
}

type NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf struct {
	value *FiveQiDscpMappingSetSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) Get() *FiveQiDscpMappingSetSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) Set(val *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf(val *FiveQiDscpMappingSetSingleAllOfAttributesAllOf) *NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf {
	return &NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiveQiDscpMappingSetSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
