/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the SmfSubscriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfSubscriptionInfo{}

// SmfSubscriptionInfo Information related to active subscriptions at the SMF(s)
type SmfSubscriptionInfo struct {
	SmfSubscriptionList []SmfSubscriptionItem `json:"smfSubscriptionList"`
}

// NewSmfSubscriptionInfo instantiates a new SmfSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfSubscriptionInfo(smfSubscriptionList []SmfSubscriptionItem) *SmfSubscriptionInfo {
	this := SmfSubscriptionInfo{}
	this.SmfSubscriptionList = smfSubscriptionList
	return &this
}

// NewSmfSubscriptionInfoWithDefaults instantiates a new SmfSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfSubscriptionInfoWithDefaults() *SmfSubscriptionInfo {
	this := SmfSubscriptionInfo{}
	return &this
}

// GetSmfSubscriptionList returns the SmfSubscriptionList field value
func (o *SmfSubscriptionInfo) GetSmfSubscriptionList() []SmfSubscriptionItem {
	if o == nil {
		var ret []SmfSubscriptionItem
		return ret
	}

	return o.SmfSubscriptionList
}

// GetSmfSubscriptionListOk returns a tuple with the SmfSubscriptionList field value
// and a boolean to check if the value has been set.
func (o *SmfSubscriptionInfo) GetSmfSubscriptionListOk() ([]SmfSubscriptionItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmfSubscriptionList, true
}

// SetSmfSubscriptionList sets field value
func (o *SmfSubscriptionInfo) SetSmfSubscriptionList(v []SmfSubscriptionItem) {
	o.SmfSubscriptionList = v
}

func (o SmfSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfSubscriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smfSubscriptionList"] = o.SmfSubscriptionList
	return toSerialize, nil
}

type NullableSmfSubscriptionInfo struct {
	value *SmfSubscriptionInfo
	isSet bool
}

func (v NullableSmfSubscriptionInfo) Get() *SmfSubscriptionInfo {
	return v.value
}

func (v *NullableSmfSubscriptionInfo) Set(val *SmfSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfSubscriptionInfo(val *SmfSubscriptionInfo) *NullableSmfSubscriptionInfo {
	return &NullableSmfSubscriptionInfo{value: val, isSet: true}
}

func (v NullableSmfSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
