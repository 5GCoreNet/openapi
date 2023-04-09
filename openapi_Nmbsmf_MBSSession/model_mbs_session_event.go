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

// checks if the MbsSessionEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionEvent{}

// MbsSessionEvent MBS session event
type MbsSessionEvent struct {
	EventType MbsSessionEventType `json:"eventType"`
}

// NewMbsSessionEvent instantiates a new MbsSessionEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionEvent(eventType MbsSessionEventType) *MbsSessionEvent {
	this := MbsSessionEvent{}
	this.EventType = eventType
	return &this
}

// NewMbsSessionEventWithDefaults instantiates a new MbsSessionEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionEventWithDefaults() *MbsSessionEvent {
	this := MbsSessionEvent{}
	return &this
}

// GetEventType returns the EventType field value
func (o *MbsSessionEvent) GetEventType() MbsSessionEventType {
	if o == nil {
		var ret MbsSessionEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MbsSessionEvent) GetEventTypeOk() (*MbsSessionEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MbsSessionEvent) SetEventType(v MbsSessionEventType) {
	o.EventType = v
}

func (o MbsSessionEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	return toSerialize, nil
}

type NullableMbsSessionEvent struct {
	value *MbsSessionEvent
	isSet bool
}

func (v NullableMbsSessionEvent) Get() *MbsSessionEvent {
	return v.value
}

func (v *NullableMbsSessionEvent) Set(val *MbsSessionEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionEvent(val *MbsSessionEvent) *NullableMbsSessionEvent {
	return &NullableMbsSessionEvent{value: val, isSet: true}
}

func (v NullableMbsSessionEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
