/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
)

// checks if the EventsNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsNotification{}

// EventsNotification Describes the notification of matched events.
type EventsNotification struct {
	NotifCorreId string `json:"notifCorreId"`
	Events []EventNotification `json:"events"`
}

// NewEventsNotification instantiates a new EventsNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsNotification(notifCorreId string, events []EventNotification) *EventsNotification {
	this := EventsNotification{}
	this.NotifCorreId = notifCorreId
	this.Events = events
	return &this
}

// NewEventsNotificationWithDefaults instantiates a new EventsNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsNotificationWithDefaults() *EventsNotification {
	this := EventsNotification{}
	return &this
}

// GetNotifCorreId returns the NotifCorreId field value
func (o *EventsNotification) GetNotifCorreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetNotifCorreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifCorreId, true
}

// SetNotifCorreId sets field value
func (o *EventsNotification) SetNotifCorreId(v string) {
	o.NotifCorreId = v
}

// GetEvents returns the Events field value
func (o *EventsNotification) GetEvents() []EventNotification {
	if o == nil {
		var ret []EventNotification
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsNotification) GetEventsOk() ([]EventNotification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsNotification) SetEvents(v []EventNotification) {
	o.Events = v
}

func (o EventsNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifCorreId"] = o.NotifCorreId
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

type NullableEventsNotification struct {
	value *EventsNotification
	isSet bool
}

func (v NullableEventsNotification) Get() *EventsNotification {
	return v.value
}

func (v *NullableEventsNotification) Set(val *EventsNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsNotification(val *EventsNotification) *NullableEventsNotification {
	return &NullableEventsNotification{value: val, isSet: true}
}

func (v NullableEventsNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


