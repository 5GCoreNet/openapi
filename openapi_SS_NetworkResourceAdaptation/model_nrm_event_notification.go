/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"time"
)

// checks if the NrmEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrmEventNotification{}

// NrmEventNotification Represents a notification on an individual User Plane event.
type NrmEventNotification struct {
	Event NrmEvent `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`
	DeliveryMode *DeliveryMode `json:"deliveryMode,omitempty"`
	StreamIds []string `json:"streamIds,omitempty"`
}

// NewNrmEventNotification instantiates a new NrmEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrmEventNotification(event NrmEvent, ts time.Time) *NrmEventNotification {
	this := NrmEventNotification{}
	this.Event = event
	this.Ts = ts
	return &this
}

// NewNrmEventNotificationWithDefaults instantiates a new NrmEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrmEventNotificationWithDefaults() *NrmEventNotification {
	this := NrmEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *NrmEventNotification) GetEvent() NrmEvent {
	if o == nil {
		var ret NrmEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NrmEventNotification) GetEventOk() (*NrmEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NrmEventNotification) SetEvent(v NrmEvent) {
	o.Event = v
}

// GetTs returns the Ts field value
func (o *NrmEventNotification) GetTs() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ts
}

// GetTsOk returns a tuple with the Ts field value
// and a boolean to check if the value has been set.
func (o *NrmEventNotification) GetTsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ts, true
}

// SetTs sets field value
func (o *NrmEventNotification) SetTs(v time.Time) {
	o.Ts = v
}

// GetDeliveryMode returns the DeliveryMode field value if set, zero value otherwise.
func (o *NrmEventNotification) GetDeliveryMode() DeliveryMode {
	if o == nil || IsNil(o.DeliveryMode) {
		var ret DeliveryMode
		return ret
	}
	return *o.DeliveryMode
}

// GetDeliveryModeOk returns a tuple with the DeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrmEventNotification) GetDeliveryModeOk() (*DeliveryMode, bool) {
	if o == nil || IsNil(o.DeliveryMode) {
		return nil, false
	}
	return o.DeliveryMode, true
}

// HasDeliveryMode returns a boolean if a field has been set.
func (o *NrmEventNotification) HasDeliveryMode() bool {
	if o != nil && !IsNil(o.DeliveryMode) {
		return true
	}

	return false
}

// SetDeliveryMode gets a reference to the given DeliveryMode and assigns it to the DeliveryMode field.
func (o *NrmEventNotification) SetDeliveryMode(v DeliveryMode) {
	o.DeliveryMode = &v
}

// GetStreamIds returns the StreamIds field value if set, zero value otherwise.
func (o *NrmEventNotification) GetStreamIds() []string {
	if o == nil || IsNil(o.StreamIds) {
		var ret []string
		return ret
	}
	return o.StreamIds
}

// GetStreamIdsOk returns a tuple with the StreamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrmEventNotification) GetStreamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StreamIds) {
		return nil, false
	}
	return o.StreamIds, true
}

// HasStreamIds returns a boolean if a field has been set.
func (o *NrmEventNotification) HasStreamIds() bool {
	if o != nil && !IsNil(o.StreamIds) {
		return true
	}

	return false
}

// SetStreamIds gets a reference to the given []string and assigns it to the StreamIds field.
func (o *NrmEventNotification) SetStreamIds(v []string) {
	o.StreamIds = v
}

func (o NrmEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrmEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["ts"] = o.Ts
	if !IsNil(o.DeliveryMode) {
		toSerialize["deliveryMode"] = o.DeliveryMode
	}
	if !IsNil(o.StreamIds) {
		toSerialize["streamIds"] = o.StreamIds
	}
	return toSerialize, nil
}

type NullableNrmEventNotification struct {
	value *NrmEventNotification
	isSet bool
}

func (v NullableNrmEventNotification) Get() *NrmEventNotification {
	return v.value
}

func (v *NullableNrmEventNotification) Set(val *NrmEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableNrmEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNrmEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrmEventNotification(val *NrmEventNotification) *NullableNrmEventNotification {
	return &NullableNrmEventNotification{value: val, isSet: true}
}

func (v NullableNrmEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrmEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


