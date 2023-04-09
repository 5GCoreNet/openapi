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

// checks if the AlarmListSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmListSingle{}

// AlarmListSingle struct for AlarmListSingle
type AlarmListSingle struct {
	Top
	Attributes *AlarmListSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewAlarmListSingle instantiates a new AlarmListSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmListSingle(id NullableString) *AlarmListSingle {
	this := AlarmListSingle{}
	this.Id = id
	return &this
}

// NewAlarmListSingleWithDefaults instantiates a new AlarmListSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmListSingleWithDefaults() *AlarmListSingle {
	this := AlarmListSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AlarmListSingle) GetAttributes() AlarmListSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AlarmListSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListSingle) GetAttributesOk() (*AlarmListSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AlarmListSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AlarmListSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AlarmListSingle) SetAttributes(v AlarmListSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o AlarmListSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmListSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAlarmListSingle struct {
	value *AlarmListSingle
	isSet bool
}

func (v NullableAlarmListSingle) Get() *AlarmListSingle {
	return v.value
}

func (v *NullableAlarmListSingle) Set(val *AlarmListSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmListSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmListSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmListSingle(val *AlarmListSingle) *NullableAlarmListSingle {
	return &NullableAlarmListSingle{value: val, isSet: true}
}

func (v NullableAlarmListSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmListSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
