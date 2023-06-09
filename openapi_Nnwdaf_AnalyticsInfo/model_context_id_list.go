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

// checks if the ContextIdList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextIdList{}

// ContextIdList Contains a list of context identifiers of context information of analytics subscriptions.
type ContextIdList struct {
	ContextIds []AnalyticsContextIdentifier `json:"contextIds"`
}

// NewContextIdList instantiates a new ContextIdList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextIdList(contextIds []AnalyticsContextIdentifier) *ContextIdList {
	this := ContextIdList{}
	this.ContextIds = contextIds
	return &this
}

// NewContextIdListWithDefaults instantiates a new ContextIdList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextIdListWithDefaults() *ContextIdList {
	this := ContextIdList{}
	return &this
}

// GetContextIds returns the ContextIds field value
func (o *ContextIdList) GetContextIds() []AnalyticsContextIdentifier {
	if o == nil {
		var ret []AnalyticsContextIdentifier
		return ret
	}

	return o.ContextIds
}

// GetContextIdsOk returns a tuple with the ContextIds field value
// and a boolean to check if the value has been set.
func (o *ContextIdList) GetContextIdsOk() ([]AnalyticsContextIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContextIds, true
}

// SetContextIds sets field value
func (o *ContextIdList) SetContextIds(v []AnalyticsContextIdentifier) {
	o.ContextIds = v
}

func (o ContextIdList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextIdList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextIds"] = o.ContextIds
	return toSerialize, nil
}

type NullableContextIdList struct {
	value *ContextIdList
	isSet bool
}

func (v NullableContextIdList) Get() *ContextIdList {
	return v.value
}

func (v *NullableContextIdList) Set(val *ContextIdList) {
	v.value = val
	v.isSet = true
}

func (v NullableContextIdList) IsSet() bool {
	return v.isSet
}

func (v *NullableContextIdList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextIdList(val *ContextIdList) *NullableContextIdList {
	return &NullableContextIdList{value: val, isSet: true}
}

func (v NullableContextIdList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextIdList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
