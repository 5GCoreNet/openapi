/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the AfEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfEventSubscription{}

// AfEventSubscription Describes the event information delivered in the subscription.
type AfEventSubscription struct {
	Event AfEvent `json:"event"`
	NotifMethod *AfNotifMethod `json:"notifMethod,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// indicating a time in seconds.
	WaitTime *int32 `json:"waitTime,omitempty"`
}

// NewAfEventSubscription instantiates a new AfEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventSubscription(event AfEvent) *AfEventSubscription {
	this := AfEventSubscription{}
	this.Event = event
	return &this
}

// NewAfEventSubscriptionWithDefaults instantiates a new AfEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventSubscriptionWithDefaults() *AfEventSubscription {
	this := AfEventSubscription{}
	return &this
}

// GetEvent returns the Event field value
func (o *AfEventSubscription) GetEvent() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetEventOk() (*AfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AfEventSubscription) SetEvent(v AfEvent) {
	o.Event = v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *AfEventSubscription) GetNotifMethod() AfNotifMethod {
	if o == nil || IsNil(o.NotifMethod) {
		var ret AfNotifMethod
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetNotifMethodOk() (*AfNotifMethod, bool) {
	if o == nil || IsNil(o.NotifMethod) {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *AfEventSubscription) HasNotifMethod() bool {
	if o != nil && !IsNil(o.NotifMethod) {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given AfNotifMethod and assigns it to the NotifMethod field.
func (o *AfEventSubscription) SetNotifMethod(v AfNotifMethod) {
	o.NotifMethod = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *AfEventSubscription) GetRepPeriod() int32 {
	if o == nil || IsNil(o.RepPeriod) {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RepPeriod) {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *AfEventSubscription) HasRepPeriod() bool {
	if o != nil && !IsNil(o.RepPeriod) {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *AfEventSubscription) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *AfEventSubscription) GetWaitTime() int32 {
	if o == nil || IsNil(o.WaitTime) {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventSubscription) GetWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTime) {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *AfEventSubscription) HasWaitTime() bool {
	if o != nil && !IsNil(o.WaitTime) {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *AfEventSubscription) SetWaitTime(v int32) {
	o.WaitTime = &v
}

func (o AfEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.NotifMethod) {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if !IsNil(o.RepPeriod) {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if !IsNil(o.WaitTime) {
		toSerialize["waitTime"] = o.WaitTime
	}
	return toSerialize, nil
}

type NullableAfEventSubscription struct {
	value *AfEventSubscription
	isSet bool
}

func (v NullableAfEventSubscription) Get() *AfEventSubscription {
	return v.value
}

func (v *NullableAfEventSubscription) Set(val *AfEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventSubscription(val *AfEventSubscription) *NullableAfEventSubscription {
	return &NullableAfEventSubscription{value: val, isSet: true}
}

func (v NullableAfEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

