/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// checks if the EventNotifyDataExt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventNotifyDataExt{}

// EventNotifyDataExt Extended Event Notify Data for UEs of a target group
type EventNotifyDataExt struct {
	EventNotifyData
	AddEventDataList []EventNotifyData `json:"addEventDataList,omitempty"`
}

// NewEventNotifyDataExt instantiates a new EventNotifyDataExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotifyDataExt(ldrReference string, eventNotifyDataType EventNotifyDataType) *EventNotifyDataExt {
	this := EventNotifyDataExt{}
	this.LdrReference = ldrReference
	this.EventNotifyDataType = eventNotifyDataType
	return &this
}

// NewEventNotifyDataExtWithDefaults instantiates a new EventNotifyDataExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotifyDataExtWithDefaults() *EventNotifyDataExt {
	this := EventNotifyDataExt{}
	return &this
}

// GetAddEventDataList returns the AddEventDataList field value if set, zero value otherwise.
func (o *EventNotifyDataExt) GetAddEventDataList() []EventNotifyData {
	if o == nil || IsNil(o.AddEventDataList) {
		var ret []EventNotifyData
		return ret
	}
	return o.AddEventDataList
}

// GetAddEventDataListOk returns a tuple with the AddEventDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyDataExt) GetAddEventDataListOk() ([]EventNotifyData, bool) {
	if o == nil || IsNil(o.AddEventDataList) {
		return nil, false
	}
	return o.AddEventDataList, true
}

// HasAddEventDataList returns a boolean if a field has been set.
func (o *EventNotifyDataExt) HasAddEventDataList() bool {
	if o != nil && !IsNil(o.AddEventDataList) {
		return true
	}

	return false
}

// SetAddEventDataList gets a reference to the given []EventNotifyData and assigns it to the AddEventDataList field.
func (o *EventNotifyDataExt) SetAddEventDataList(v []EventNotifyData) {
	o.AddEventDataList = v
}

func (o EventNotifyDataExt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventNotifyDataExt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEventNotifyData, errEventNotifyData := json.Marshal(o.EventNotifyData)
	if errEventNotifyData != nil {
		return map[string]interface{}{}, errEventNotifyData
	}
	errEventNotifyData = json.Unmarshal([]byte(serializedEventNotifyData), &toSerialize)
	if errEventNotifyData != nil {
		return map[string]interface{}{}, errEventNotifyData
	}
	if !IsNil(o.AddEventDataList) {
		toSerialize["addEventDataList"] = o.AddEventDataList
	}
	return toSerialize, nil
}

type NullableEventNotifyDataExt struct {
	value *EventNotifyDataExt
	isSet bool
}

func (v NullableEventNotifyDataExt) Get() *EventNotifyDataExt {
	return v.value
}

func (v *NullableEventNotifyDataExt) Set(val *EventNotifyDataExt) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotifyDataExt) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotifyDataExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotifyDataExt(val *EventNotifyDataExt) *NullableEventNotifyDataExt {
	return &NullableEventNotifyDataExt{value: val, isSet: true}
}

func (v NullableEventNotifyDataExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotifyDataExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
