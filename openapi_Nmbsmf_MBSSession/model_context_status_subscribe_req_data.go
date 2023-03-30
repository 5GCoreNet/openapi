/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the ContextStatusSubscribeReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextStatusSubscribeReqData{}

// ContextStatusSubscribeReqData Data within ContextStatusSubscribe Request
type ContextStatusSubscribeReqData struct {
	Subscription ContextStatusSubscription `json:"subscription"`
}

// NewContextStatusSubscribeReqData instantiates a new ContextStatusSubscribeReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextStatusSubscribeReqData(subscription ContextStatusSubscription) *ContextStatusSubscribeReqData {
	this := ContextStatusSubscribeReqData{}
	this.Subscription = subscription
	return &this
}

// NewContextStatusSubscribeReqDataWithDefaults instantiates a new ContextStatusSubscribeReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextStatusSubscribeReqDataWithDefaults() *ContextStatusSubscribeReqData {
	this := ContextStatusSubscribeReqData{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *ContextStatusSubscribeReqData) GetSubscription() ContextStatusSubscription {
	if o == nil {
		var ret ContextStatusSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *ContextStatusSubscribeReqData) GetSubscriptionOk() (*ContextStatusSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *ContextStatusSubscribeReqData) SetSubscription(v ContextStatusSubscription) {
	o.Subscription = v
}

func (o ContextStatusSubscribeReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextStatusSubscribeReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	return toSerialize, nil
}

type NullableContextStatusSubscribeReqData struct {
	value *ContextStatusSubscribeReqData
	isSet bool
}

func (v NullableContextStatusSubscribeReqData) Get() *ContextStatusSubscribeReqData {
	return v.value
}

func (v *NullableContextStatusSubscribeReqData) Set(val *ContextStatusSubscribeReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusSubscribeReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusSubscribeReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusSubscribeReqData(val *ContextStatusSubscribeReqData) *NullableContextStatusSubscribeReqData {
	return &NullableContextStatusSubscribeReqData{value: val, isSet: true}
}

func (v NullableContextStatusSubscribeReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusSubscribeReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


