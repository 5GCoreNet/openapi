/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
)

// checks if the AvailabilityNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityNotif{}

// AvailabilityNotif Represents the availability information of user plane path management events monitoring via the 3GPP 5GC network.
type AvailabilityNotif struct {
	AvailabilityStatus AvailabilityStatus `json:"availabilityStatus"`
}

// NewAvailabilityNotif instantiates a new AvailabilityNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityNotif(availabilityStatus AvailabilityStatus) *AvailabilityNotif {
	this := AvailabilityNotif{}
	this.AvailabilityStatus = availabilityStatus
	return &this
}

// NewAvailabilityNotifWithDefaults instantiates a new AvailabilityNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityNotifWithDefaults() *AvailabilityNotif {
	this := AvailabilityNotif{}
	return &this
}

// GetAvailabilityStatus returns the AvailabilityStatus field value
func (o *AvailabilityNotif) GetAvailabilityStatus() AvailabilityStatus {
	if o == nil {
		var ret AvailabilityStatus
		return ret
	}

	return o.AvailabilityStatus
}

// GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field value
// and a boolean to check if the value has been set.
func (o *AvailabilityNotif) GetAvailabilityStatusOk() (*AvailabilityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityStatus, true
}

// SetAvailabilityStatus sets field value
func (o *AvailabilityNotif) SetAvailabilityStatus(v AvailabilityStatus) {
	o.AvailabilityStatus = v
}

func (o AvailabilityNotif) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["availabilityStatus"] = o.AvailabilityStatus
	return toSerialize, nil
}

type NullableAvailabilityNotif struct {
	value *AvailabilityNotif
	isSet bool
}

func (v NullableAvailabilityNotif) Get() *AvailabilityNotif {
	return v.value
}

func (v *NullableAvailabilityNotif) Set(val *AvailabilityNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityNotif(val *AvailabilityNotif) *NullableAvailabilityNotif {
	return &NullableAvailabilityNotif{value: val, isSet: true}
}

func (v NullableAvailabilityNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
