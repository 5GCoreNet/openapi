/*
Heartbeat notification

OAS 3.0.1 definition of the heartbeat notification © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_HeartbeatNtf

import (
	"encoding/json"
)

// checks if the NotifyHeartbeatAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyHeartbeatAllOf{}

// NotifyHeartbeatAllOf struct for NotifyHeartbeatAllOf
type NotifyHeartbeatAllOf struct {
	HeartbeatNtfPeriod *int32 `json:"heartbeatNtfPeriod,omitempty"`
}

// NewNotifyHeartbeatAllOf instantiates a new NotifyHeartbeatAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyHeartbeatAllOf() *NotifyHeartbeatAllOf {
	this := NotifyHeartbeatAllOf{}
	return &this
}

// NewNotifyHeartbeatAllOfWithDefaults instantiates a new NotifyHeartbeatAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyHeartbeatAllOfWithDefaults() *NotifyHeartbeatAllOf {
	this := NotifyHeartbeatAllOf{}
	return &this
}

// GetHeartbeatNtfPeriod returns the HeartbeatNtfPeriod field value if set, zero value otherwise.
func (o *NotifyHeartbeatAllOf) GetHeartbeatNtfPeriod() int32 {
	if o == nil || IsNil(o.HeartbeatNtfPeriod) {
		var ret int32
		return ret
	}
	return *o.HeartbeatNtfPeriod
}

// GetHeartbeatNtfPeriodOk returns a tuple with the HeartbeatNtfPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyHeartbeatAllOf) GetHeartbeatNtfPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.HeartbeatNtfPeriod) {
		return nil, false
	}
	return o.HeartbeatNtfPeriod, true
}

// HasHeartbeatNtfPeriod returns a boolean if a field has been set.
func (o *NotifyHeartbeatAllOf) HasHeartbeatNtfPeriod() bool {
	if o != nil && !IsNil(o.HeartbeatNtfPeriod) {
		return true
	}

	return false
}

// SetHeartbeatNtfPeriod gets a reference to the given int32 and assigns it to the HeartbeatNtfPeriod field.
func (o *NotifyHeartbeatAllOf) SetHeartbeatNtfPeriod(v int32) {
	o.HeartbeatNtfPeriod = &v
}

func (o NotifyHeartbeatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyHeartbeatAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeartbeatNtfPeriod) {
		toSerialize["heartbeatNtfPeriod"] = o.HeartbeatNtfPeriod
	}
	return toSerialize, nil
}

type NullableNotifyHeartbeatAllOf struct {
	value *NotifyHeartbeatAllOf
	isSet bool
}

func (v NullableNotifyHeartbeatAllOf) Get() *NotifyHeartbeatAllOf {
	return v.value
}

func (v *NullableNotifyHeartbeatAllOf) Set(val *NotifyHeartbeatAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyHeartbeatAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyHeartbeatAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyHeartbeatAllOf(val *NotifyHeartbeatAllOf) *NullableNotifyHeartbeatAllOf {
	return &NullableNotifyHeartbeatAllOf{value: val, isSet: true}
}

func (v NullableNotifyHeartbeatAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyHeartbeatAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
