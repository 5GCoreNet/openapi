/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// checks if the AnalyticsContextIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsContextIdentifier{}

// AnalyticsContextIdentifier Contains information about available analytics contexts.
type AnalyticsContextIdentifier struct {
	// The identifier of a subscription.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// List of analytics types for which NF related analytics contexts can be retrieved.
	NfAnaCtxts []NwdafEvent `json:"nfAnaCtxts,omitempty"`
	// List of objects that indicate for which SUPI and analytics types combinations analytics  context can be retrieved.
	UeAnaCtxts []UeAnalyticsContextDescriptor `json:"ueAnaCtxts,omitempty"`
}

// NewAnalyticsContextIdentifier instantiates a new AnalyticsContextIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsContextIdentifier() *AnalyticsContextIdentifier {
	this := AnalyticsContextIdentifier{}
	return &this
}

// NewAnalyticsContextIdentifierWithDefaults instantiates a new AnalyticsContextIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsContextIdentifierWithDefaults() *AnalyticsContextIdentifier {
	this := AnalyticsContextIdentifier{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AnalyticsContextIdentifier) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsContextIdentifier) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AnalyticsContextIdentifier) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AnalyticsContextIdentifier) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetNfAnaCtxts returns the NfAnaCtxts field value if set, zero value otherwise.
func (o *AnalyticsContextIdentifier) GetNfAnaCtxts() []NwdafEvent {
	if o == nil || IsNil(o.NfAnaCtxts) {
		var ret []NwdafEvent
		return ret
	}
	return o.NfAnaCtxts
}

// GetNfAnaCtxtsOk returns a tuple with the NfAnaCtxts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsContextIdentifier) GetNfAnaCtxtsOk() ([]NwdafEvent, bool) {
	if o == nil || IsNil(o.NfAnaCtxts) {
		return nil, false
	}
	return o.NfAnaCtxts, true
}

// HasNfAnaCtxts returns a boolean if a field has been set.
func (o *AnalyticsContextIdentifier) HasNfAnaCtxts() bool {
	if o != nil && !IsNil(o.NfAnaCtxts) {
		return true
	}

	return false
}

// SetNfAnaCtxts gets a reference to the given []NwdafEvent and assigns it to the NfAnaCtxts field.
func (o *AnalyticsContextIdentifier) SetNfAnaCtxts(v []NwdafEvent) {
	o.NfAnaCtxts = v
}

// GetUeAnaCtxts returns the UeAnaCtxts field value if set, zero value otherwise.
func (o *AnalyticsContextIdentifier) GetUeAnaCtxts() []UeAnalyticsContextDescriptor {
	if o == nil || IsNil(o.UeAnaCtxts) {
		var ret []UeAnalyticsContextDescriptor
		return ret
	}
	return o.UeAnaCtxts
}

// GetUeAnaCtxtsOk returns a tuple with the UeAnaCtxts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsContextIdentifier) GetUeAnaCtxtsOk() ([]UeAnalyticsContextDescriptor, bool) {
	if o == nil || IsNil(o.UeAnaCtxts) {
		return nil, false
	}
	return o.UeAnaCtxts, true
}

// HasUeAnaCtxts returns a boolean if a field has been set.
func (o *AnalyticsContextIdentifier) HasUeAnaCtxts() bool {
	if o != nil && !IsNil(o.UeAnaCtxts) {
		return true
	}

	return false
}

// SetUeAnaCtxts gets a reference to the given []UeAnalyticsContextDescriptor and assigns it to the UeAnaCtxts field.
func (o *AnalyticsContextIdentifier) SetUeAnaCtxts(v []UeAnalyticsContextDescriptor) {
	o.UeAnaCtxts = v
}

func (o AnalyticsContextIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsContextIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.NfAnaCtxts) {
		toSerialize["nfAnaCtxts"] = o.NfAnaCtxts
	}
	if !IsNil(o.UeAnaCtxts) {
		toSerialize["ueAnaCtxts"] = o.UeAnaCtxts
	}
	return toSerialize, nil
}

type NullableAnalyticsContextIdentifier struct {
	value *AnalyticsContextIdentifier
	isSet bool
}

func (v NullableAnalyticsContextIdentifier) Get() *AnalyticsContextIdentifier {
	return v.value
}

func (v *NullableAnalyticsContextIdentifier) Set(val *AnalyticsContextIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsContextIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsContextIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsContextIdentifier(val *AnalyticsContextIdentifier) *NullableAnalyticsContextIdentifier {
	return &NullableAnalyticsContextIdentifier{value: val, isSet: true}
}

func (v NullableAnalyticsContextIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsContextIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
