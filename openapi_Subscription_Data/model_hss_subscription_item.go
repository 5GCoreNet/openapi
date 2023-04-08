/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the HssSubscriptionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HssSubscriptionItem{}

// HssSubscriptionItem Contains info about a single HSS event subscription
type HssSubscriptionItem struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	HssInstanceId string `json:"hssInstanceId"`
	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
}

// NewHssSubscriptionItem instantiates a new HssSubscriptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHssSubscriptionItem(hssInstanceId string, subscriptionId string) *HssSubscriptionItem {
	this := HssSubscriptionItem{}
	this.HssInstanceId = hssInstanceId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewHssSubscriptionItemWithDefaults instantiates a new HssSubscriptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHssSubscriptionItemWithDefaults() *HssSubscriptionItem {
	this := HssSubscriptionItem{}
	return &this
}

// GetHssInstanceId returns the HssInstanceId field value
func (o *HssSubscriptionItem) GetHssInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HssInstanceId
}

// GetHssInstanceIdOk returns a tuple with the HssInstanceId field value
// and a boolean to check if the value has been set.
func (o *HssSubscriptionItem) GetHssInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HssInstanceId, true
}

// SetHssInstanceId sets field value
func (o *HssSubscriptionItem) SetHssInstanceId(v string) {
	o.HssInstanceId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *HssSubscriptionItem) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *HssSubscriptionItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *HssSubscriptionItem) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *HssSubscriptionItem) GetContextInfo() ContextInfo {
	if o == nil || isNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HssSubscriptionItem) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || isNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *HssSubscriptionItem) HasContextInfo() bool {
	if o != nil && !isNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *HssSubscriptionItem) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

func (o HssSubscriptionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HssSubscriptionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hssInstanceId"] = o.HssInstanceId
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !isNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	return toSerialize, nil
}

type NullableHssSubscriptionItem struct {
	value *HssSubscriptionItem
	isSet bool
}

func (v NullableHssSubscriptionItem) Get() *HssSubscriptionItem {
	return v.value
}

func (v *NullableHssSubscriptionItem) Set(val *HssSubscriptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableHssSubscriptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableHssSubscriptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssSubscriptionItem(val *HssSubscriptionItem) *NullableHssSubscriptionItem {
	return &NullableHssSubscriptionItem{value: val, isSet: true}
}

func (v NullableHssSubscriptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssSubscriptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


