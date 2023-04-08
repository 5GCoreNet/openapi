/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the NwdafInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafInfo{}

// NwdafInfo Information of a NWDAF NF Instance
type NwdafInfo struct {
	EventIds []EventId `json:"eventIds,omitempty"`
	NwdafEvents []NwdafEvent `json:"nwdafEvents,omitempty"`
}

// NewNwdafInfo instantiates a new NwdafInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafInfo() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// NewNwdafInfoWithDefaults instantiates a new NwdafInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafInfoWithDefaults() *NwdafInfo {
	this := NwdafInfo{}
	return &this
}

// GetEventIds returns the EventIds field value if set, zero value otherwise.
func (o *NwdafInfo) GetEventIds() []EventId {
	if o == nil || isNil(o.EventIds) {
		var ret []EventId
		return ret
	}
	return o.EventIds
}

// GetEventIdsOk returns a tuple with the EventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetEventIdsOk() ([]EventId, bool) {
	if o == nil || isNil(o.EventIds) {
		return nil, false
	}
	return o.EventIds, true
}

// HasEventIds returns a boolean if a field has been set.
func (o *NwdafInfo) HasEventIds() bool {
	if o != nil && !isNil(o.EventIds) {
		return true
	}

	return false
}

// SetEventIds gets a reference to the given []EventId and assigns it to the EventIds field.
func (o *NwdafInfo) SetEventIds(v []EventId) {
	o.EventIds = v
}

// GetNwdafEvents returns the NwdafEvents field value if set, zero value otherwise.
func (o *NwdafInfo) GetNwdafEvents() []NwdafEvent {
	if o == nil || isNil(o.NwdafEvents) {
		var ret []NwdafEvent
		return ret
	}
	return o.NwdafEvents
}

// GetNwdafEventsOk returns a tuple with the NwdafEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafInfo) GetNwdafEventsOk() ([]NwdafEvent, bool) {
	if o == nil || isNil(o.NwdafEvents) {
		return nil, false
	}
	return o.NwdafEvents, true
}

// HasNwdafEvents returns a boolean if a field has been set.
func (o *NwdafInfo) HasNwdafEvents() bool {
	if o != nil && !isNil(o.NwdafEvents) {
		return true
	}

	return false
}

// SetNwdafEvents gets a reference to the given []NwdafEvent and assigns it to the NwdafEvents field.
func (o *NwdafInfo) SetNwdafEvents(v []NwdafEvent) {
	o.NwdafEvents = v
}

func (o NwdafInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventIds) {
		toSerialize["eventIds"] = o.EventIds
	}
	if !isNil(o.NwdafEvents) {
		toSerialize["nwdafEvents"] = o.NwdafEvents
	}
	return toSerialize, nil
}

type NullableNwdafInfo struct {
	value *NwdafInfo
	isSet bool
}

func (v NullableNwdafInfo) Get() *NwdafInfo {
	return v.value
}

func (v *NullableNwdafInfo) Set(val *NwdafInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafInfo(val *NwdafInfo) *NullableNwdafInfo {
	return &NullableNwdafInfo{value: val, isSet: true}
}

func (v NullableNwdafInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


