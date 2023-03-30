/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
)

// checks if the SubscribedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribedEvent{}

// SubscribedEvent Represents a subscribed MBS User Data Ingest Session Status event and the related  information. 
type SubscribedEvent struct {
	StatusEvent Event `json:"statusEvent"`
	MbsDistSessionId *string `json:"mbsDistSessionId,omitempty"`
}

// NewSubscribedEvent instantiates a new SubscribedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribedEvent(statusEvent Event) *SubscribedEvent {
	this := SubscribedEvent{}
	this.StatusEvent = statusEvent
	return &this
}

// NewSubscribedEventWithDefaults instantiates a new SubscribedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribedEventWithDefaults() *SubscribedEvent {
	this := SubscribedEvent{}
	return &this
}

// GetStatusEvent returns the StatusEvent field value
func (o *SubscribedEvent) GetStatusEvent() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.StatusEvent
}

// GetStatusEventOk returns a tuple with the StatusEvent field value
// and a boolean to check if the value has been set.
func (o *SubscribedEvent) GetStatusEventOk() (*Event, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusEvent, true
}

// SetStatusEvent sets field value
func (o *SubscribedEvent) SetStatusEvent(v Event) {
	o.StatusEvent = v
}

// GetMbsDistSessionId returns the MbsDistSessionId field value if set, zero value otherwise.
func (o *SubscribedEvent) GetMbsDistSessionId() string {
	if o == nil || IsNil(o.MbsDistSessionId) {
		var ret string
		return ret
	}
	return *o.MbsDistSessionId
}

// GetMbsDistSessionIdOk returns a tuple with the MbsDistSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribedEvent) GetMbsDistSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.MbsDistSessionId) {
		return nil, false
	}
	return o.MbsDistSessionId, true
}

// HasMbsDistSessionId returns a boolean if a field has been set.
func (o *SubscribedEvent) HasMbsDistSessionId() bool {
	if o != nil && !IsNil(o.MbsDistSessionId) {
		return true
	}

	return false
}

// SetMbsDistSessionId gets a reference to the given string and assigns it to the MbsDistSessionId field.
func (o *SubscribedEvent) SetMbsDistSessionId(v string) {
	o.MbsDistSessionId = &v
}

func (o SubscribedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusEvent"] = o.StatusEvent
	if !IsNil(o.MbsDistSessionId) {
		toSerialize["mbsDistSessionId"] = o.MbsDistSessionId
	}
	return toSerialize, nil
}

type NullableSubscribedEvent struct {
	value *SubscribedEvent
	isSet bool
}

func (v NullableSubscribedEvent) Get() *SubscribedEvent {
	return v.value
}

func (v *NullableSubscribedEvent) Set(val *SubscribedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedEvent(val *SubscribedEvent) *NullableSubscribedEvent {
	return &NullableSubscribedEvent{value: val, isSet: true}
}

func (v NullableSubscribedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

