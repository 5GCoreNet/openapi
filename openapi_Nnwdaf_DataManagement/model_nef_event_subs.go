/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the NefEventSubs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NefEventSubs{}

// NefEventSubs Represents an event to be subscribed and the related event filter information.
type NefEventSubs struct {
	Event NefEvent `json:"event"`
	EventFilter *NefEventFilter `json:"eventFilter,omitempty"`
}

// NewNefEventSubs instantiates a new NefEventSubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefEventSubs(event NefEvent) *NefEventSubs {
	this := NefEventSubs{}
	this.Event = event
	return &this
}

// NewNefEventSubsWithDefaults instantiates a new NefEventSubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefEventSubsWithDefaults() *NefEventSubs {
	this := NefEventSubs{}
	return &this
}

// GetEvent returns the Event field value
func (o *NefEventSubs) GetEvent() NefEvent {
	if o == nil {
		var ret NefEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NefEventSubs) GetEventOk() (*NefEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NefEventSubs) SetEvent(v NefEvent) {
	o.Event = v
}

// GetEventFilter returns the EventFilter field value if set, zero value otherwise.
func (o *NefEventSubs) GetEventFilter() NefEventFilter {
	if o == nil || IsNil(o.EventFilter) {
		var ret NefEventFilter
		return ret
	}
	return *o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventSubs) GetEventFilterOk() (*NefEventFilter, bool) {
	if o == nil || IsNil(o.EventFilter) {
		return nil, false
	}
	return o.EventFilter, true
}

// HasEventFilter returns a boolean if a field has been set.
func (o *NefEventSubs) HasEventFilter() bool {
	if o != nil && !IsNil(o.EventFilter) {
		return true
	}

	return false
}

// SetEventFilter gets a reference to the given NefEventFilter and assigns it to the EventFilter field.
func (o *NefEventSubs) SetEventFilter(v NefEventFilter) {
	o.EventFilter = &v
}

func (o NefEventSubs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NefEventSubs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.EventFilter) {
		toSerialize["eventFilter"] = o.EventFilter
	}
	return toSerialize, nil
}

type NullableNefEventSubs struct {
	value *NefEventSubs
	isSet bool
}

func (v NullableNefEventSubs) Get() *NefEventSubs {
	return v.value
}

func (v *NullableNefEventSubs) Set(val *NefEventSubs) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEventSubs) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEventSubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEventSubs(val *NefEventSubs) *NullableNefEventSubs {
	return &NullableNefEventSubs{value: val, isSet: true}
}

func (v NullableNefEventSubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEventSubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


