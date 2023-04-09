/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// checks if the SubsEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubsEventNotification{}

// SubsEventNotification Notifications about subscribed Individual Events.
type SubsEventNotification struct {
	Event         SubscribedEvent      `json:"event"`
	TimeSyncCapas []TimeSyncCapability `json:"timeSyncCapas,omitempty"`
}

// NewSubsEventNotification instantiates a new SubsEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubsEventNotification(event SubscribedEvent) *SubsEventNotification {
	this := SubsEventNotification{}
	this.Event = event
	return &this
}

// NewSubsEventNotificationWithDefaults instantiates a new SubsEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubsEventNotificationWithDefaults() *SubsEventNotification {
	this := SubsEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *SubsEventNotification) GetEvent() SubscribedEvent {
	if o == nil {
		var ret SubscribedEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *SubsEventNotification) GetEventOk() (*SubscribedEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *SubsEventNotification) SetEvent(v SubscribedEvent) {
	o.Event = v
}

// GetTimeSyncCapas returns the TimeSyncCapas field value if set, zero value otherwise.
func (o *SubsEventNotification) GetTimeSyncCapas() []TimeSyncCapability {
	if o == nil || IsNil(o.TimeSyncCapas) {
		var ret []TimeSyncCapability
		return ret
	}
	return o.TimeSyncCapas
}

// GetTimeSyncCapasOk returns a tuple with the TimeSyncCapas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubsEventNotification) GetTimeSyncCapasOk() ([]TimeSyncCapability, bool) {
	if o == nil || IsNil(o.TimeSyncCapas) {
		return nil, false
	}
	return o.TimeSyncCapas, true
}

// HasTimeSyncCapas returns a boolean if a field has been set.
func (o *SubsEventNotification) HasTimeSyncCapas() bool {
	if o != nil && !IsNil(o.TimeSyncCapas) {
		return true
	}

	return false
}

// SetTimeSyncCapas gets a reference to the given []TimeSyncCapability and assigns it to the TimeSyncCapas field.
func (o *SubsEventNotification) SetTimeSyncCapas(v []TimeSyncCapability) {
	o.TimeSyncCapas = v
}

func (o SubsEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubsEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.TimeSyncCapas) {
		toSerialize["timeSyncCapas"] = o.TimeSyncCapas
	}
	return toSerialize, nil
}

type NullableSubsEventNotification struct {
	value *SubsEventNotification
	isSet bool
}

func (v NullableSubsEventNotification) Get() *SubsEventNotification {
	return v.value
}

func (v *NullableSubsEventNotification) Set(val *SubsEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSubsEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSubsEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubsEventNotification(val *SubsEventNotification) *NullableSubsEventNotification {
	return &NullableSubsEventNotification{value: val, isSet: true}
}

func (v NullableSubsEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubsEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
