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

// checks if the AmfEventSubscriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventSubscriptionInfo{}

// AmfEventSubscriptionInfo Individual AMF Event Subscription Information
type AmfEventSubscriptionInfo struct {
	// String providing an URI formatted according to RFC 3986.
	SubId string `json:"subId"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	RefIdList []int32 `json:"refIdList"`
	// String providing an URI formatted according to RFC 3986.
	OldSubId *string `json:"oldSubId,omitempty"`
}

// NewAmfEventSubscriptionInfo instantiates a new AmfEventSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventSubscriptionInfo(subId string, refIdList []int32) *AmfEventSubscriptionInfo {
	this := AmfEventSubscriptionInfo{}
	this.SubId = subId
	this.RefIdList = refIdList
	return &this
}

// NewAmfEventSubscriptionInfoWithDefaults instantiates a new AmfEventSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventSubscriptionInfoWithDefaults() *AmfEventSubscriptionInfo {
	this := AmfEventSubscriptionInfo{}
	return &this
}

// GetSubId returns the SubId field value
func (o *AmfEventSubscriptionInfo) GetSubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscriptionInfo) GetSubIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubId, true
}

// SetSubId sets field value
func (o *AmfEventSubscriptionInfo) SetSubId(v string) {
	o.SubId = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *AmfEventSubscriptionInfo) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscriptionInfo) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *AmfEventSubscriptionInfo) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *AmfEventSubscriptionInfo) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetRefIdList returns the RefIdList field value
func (o *AmfEventSubscriptionInfo) GetRefIdList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RefIdList
}

// GetRefIdListOk returns a tuple with the RefIdList field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscriptionInfo) GetRefIdListOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefIdList, true
}

// SetRefIdList sets field value
func (o *AmfEventSubscriptionInfo) SetRefIdList(v []int32) {
	o.RefIdList = v
}

// GetOldSubId returns the OldSubId field value if set, zero value otherwise.
func (o *AmfEventSubscriptionInfo) GetOldSubId() string {
	if o == nil || IsNil(o.OldSubId) {
		var ret string
		return ret
	}
	return *o.OldSubId
}

// GetOldSubIdOk returns a tuple with the OldSubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscriptionInfo) GetOldSubIdOk() (*string, bool) {
	if o == nil || IsNil(o.OldSubId) {
		return nil, false
	}
	return o.OldSubId, true
}

// HasOldSubId returns a boolean if a field has been set.
func (o *AmfEventSubscriptionInfo) HasOldSubId() bool {
	if o != nil && !IsNil(o.OldSubId) {
		return true
	}

	return false
}

// SetOldSubId gets a reference to the given string and assigns it to the OldSubId field.
func (o *AmfEventSubscriptionInfo) SetOldSubId(v string) {
	o.OldSubId = &v
}

func (o AmfEventSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventSubscriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subId"] = o.SubId
	if !IsNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	toSerialize["refIdList"] = o.RefIdList
	if !IsNil(o.OldSubId) {
		toSerialize["oldSubId"] = o.OldSubId
	}
	return toSerialize, nil
}

type NullableAmfEventSubscriptionInfo struct {
	value *AmfEventSubscriptionInfo
	isSet bool
}

func (v NullableAmfEventSubscriptionInfo) Get() *AmfEventSubscriptionInfo {
	return v.value
}

func (v *NullableAmfEventSubscriptionInfo) Set(val *AmfEventSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventSubscriptionInfo(val *AmfEventSubscriptionInfo) *NullableAmfEventSubscriptionInfo {
	return &NullableAmfEventSubscriptionInfo{value: val, isSet: true}
}

func (v NullableAmfEventSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


