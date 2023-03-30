/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// checks if the NotifyPotentialFaultyAlarmListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyPotentialFaultyAlarmListAllOf{}

// NotifyPotentialFaultyAlarmListAllOf struct for NotifyPotentialFaultyAlarmListAllOf
type NotifyPotentialFaultyAlarmListAllOf struct {
	Reason string `json:"reason"`
}

// NewNotifyPotentialFaultyAlarmListAllOf instantiates a new NotifyPotentialFaultyAlarmListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyPotentialFaultyAlarmListAllOf(reason string) *NotifyPotentialFaultyAlarmListAllOf {
	this := NotifyPotentialFaultyAlarmListAllOf{}
	this.Reason = reason
	return &this
}

// NewNotifyPotentialFaultyAlarmListAllOfWithDefaults instantiates a new NotifyPotentialFaultyAlarmListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyPotentialFaultyAlarmListAllOfWithDefaults() *NotifyPotentialFaultyAlarmListAllOf {
	this := NotifyPotentialFaultyAlarmListAllOf{}
	return &this
}

// GetReason returns the Reason field value
func (o *NotifyPotentialFaultyAlarmListAllOf) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *NotifyPotentialFaultyAlarmListAllOf) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *NotifyPotentialFaultyAlarmListAllOf) SetReason(v string) {
	o.Reason = v
}

func (o NotifyPotentialFaultyAlarmListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyPotentialFaultyAlarmListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

type NullableNotifyPotentialFaultyAlarmListAllOf struct {
	value *NotifyPotentialFaultyAlarmListAllOf
	isSet bool
}

func (v NullableNotifyPotentialFaultyAlarmListAllOf) Get() *NotifyPotentialFaultyAlarmListAllOf {
	return v.value
}

func (v *NullableNotifyPotentialFaultyAlarmListAllOf) Set(val *NotifyPotentialFaultyAlarmListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyPotentialFaultyAlarmListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyPotentialFaultyAlarmListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyPotentialFaultyAlarmListAllOf(val *NotifyPotentialFaultyAlarmListAllOf) *NullableNotifyPotentialFaultyAlarmListAllOf {
	return &NullableNotifyPotentialFaultyAlarmListAllOf{value: val, isSet: true}
}

func (v NullableNotifyPotentialFaultyAlarmListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyPotentialFaultyAlarmListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


