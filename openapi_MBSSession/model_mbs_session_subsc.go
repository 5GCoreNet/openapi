/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the MbsSessionSubsc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionSubsc{}

// MbsSessionSubsc Represents an MBS Session Subscription.
type MbsSessionSubsc struct {
	AfId string `json:"afId"`
	Subscription MbsSessionSubscription `json:"subscription"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewMbsSessionSubsc instantiates a new MbsSessionSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionSubsc(afId string, subscription MbsSessionSubscription) *MbsSessionSubsc {
	this := MbsSessionSubsc{}
	this.AfId = afId
	this.Subscription = subscription
	return &this
}

// NewMbsSessionSubscWithDefaults instantiates a new MbsSessionSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionSubscWithDefaults() *MbsSessionSubsc {
	this := MbsSessionSubsc{}
	return &this
}

// GetAfId returns the AfId field value
func (o *MbsSessionSubsc) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *MbsSessionSubsc) GetAfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *MbsSessionSubsc) SetAfId(v string) {
	o.AfId = v
}

// GetSubscription returns the Subscription field value
func (o *MbsSessionSubsc) GetSubscription() MbsSessionSubscription {
	if o == nil {
		var ret MbsSessionSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *MbsSessionSubsc) GetSubscriptionOk() (*MbsSessionSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *MbsSessionSubsc) SetSubscription(v MbsSessionSubscription) {
	o.Subscription = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *MbsSessionSubsc) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionSubsc) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *MbsSessionSubsc) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *MbsSessionSubsc) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o MbsSessionSubsc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionSubsc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afId"] = o.AfId
	toSerialize["subscription"] = o.Subscription
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableMbsSessionSubsc struct {
	value *MbsSessionSubsc
	isSet bool
}

func (v NullableMbsSessionSubsc) Get() *MbsSessionSubsc {
	return v.value
}

func (v *NullableMbsSessionSubsc) Set(val *MbsSessionSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionSubsc(val *MbsSessionSubsc) *NullableMbsSessionSubsc {
	return &NullableMbsSessionSubsc{value: val, isSet: true}
}

func (v NullableMbsSessionSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


