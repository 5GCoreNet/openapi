/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the AmfEventSubsSyncInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventSubsSyncInfo{}

// AmfEventSubsSyncInfo AMF Event Subscriptions Information for synchronization
type AmfEventSubsSyncInfo struct {
	SubscriptionList []AmfEventSubscriptionInfo `json:"subscriptionList"`
}

// NewAmfEventSubsSyncInfo instantiates a new AmfEventSubsSyncInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventSubsSyncInfo(subscriptionList []AmfEventSubscriptionInfo) *AmfEventSubsSyncInfo {
	this := AmfEventSubsSyncInfo{}
	this.SubscriptionList = subscriptionList
	return &this
}

// NewAmfEventSubsSyncInfoWithDefaults instantiates a new AmfEventSubsSyncInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventSubsSyncInfoWithDefaults() *AmfEventSubsSyncInfo {
	this := AmfEventSubsSyncInfo{}
	return &this
}

// GetSubscriptionList returns the SubscriptionList field value
func (o *AmfEventSubsSyncInfo) GetSubscriptionList() []AmfEventSubscriptionInfo {
	if o == nil {
		var ret []AmfEventSubscriptionInfo
		return ret
	}

	return o.SubscriptionList
}

// GetSubscriptionListOk returns a tuple with the SubscriptionList field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubsSyncInfo) GetSubscriptionListOk() ([]AmfEventSubscriptionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionList, true
}

// SetSubscriptionList sets field value
func (o *AmfEventSubsSyncInfo) SetSubscriptionList(v []AmfEventSubscriptionInfo) {
	o.SubscriptionList = v
}

func (o AmfEventSubsSyncInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventSubsSyncInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionList"] = o.SubscriptionList
	return toSerialize, nil
}

type NullableAmfEventSubsSyncInfo struct {
	value *AmfEventSubsSyncInfo
	isSet bool
}

func (v NullableAmfEventSubsSyncInfo) Get() *AmfEventSubsSyncInfo {
	return v.value
}

func (v *NullableAmfEventSubsSyncInfo) Set(val *AmfEventSubsSyncInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventSubsSyncInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventSubsSyncInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventSubsSyncInfo(val *AmfEventSubsSyncInfo) *NullableAmfEventSubsSyncInfo {
	return &NullableAmfEventSubsSyncInfo{value: val, isSet: true}
}

func (v NullableAmfEventSubsSyncInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventSubsSyncInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

