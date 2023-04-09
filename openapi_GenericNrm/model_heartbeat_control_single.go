/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the HeartbeatControlSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeartbeatControlSingle{}

// HeartbeatControlSingle struct for HeartbeatControlSingle
type HeartbeatControlSingle struct {
	Top
	Attributes *HeartbeatControlSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewHeartbeatControlSingle instantiates a new HeartbeatControlSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeartbeatControlSingle(id NullableString) *HeartbeatControlSingle {
	this := HeartbeatControlSingle{}
	this.Id = id
	return &this
}

// NewHeartbeatControlSingleWithDefaults instantiates a new HeartbeatControlSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeartbeatControlSingleWithDefaults() *HeartbeatControlSingle {
	this := HeartbeatControlSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *HeartbeatControlSingle) GetAttributes() HeartbeatControlSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret HeartbeatControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatControlSingle) GetAttributesOk() (*HeartbeatControlSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *HeartbeatControlSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given HeartbeatControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *HeartbeatControlSingle) SetAttributes(v HeartbeatControlSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o HeartbeatControlSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeartbeatControlSingle) ToMap() (map[string]interface{}, error) {
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

type NullableHeartbeatControlSingle struct {
	value *HeartbeatControlSingle
	isSet bool
}

func (v NullableHeartbeatControlSingle) Get() *HeartbeatControlSingle {
	return v.value
}

func (v *NullableHeartbeatControlSingle) Set(val *HeartbeatControlSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartbeatControlSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartbeatControlSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartbeatControlSingle(val *HeartbeatControlSingle) *NullableHeartbeatControlSingle {
	return &NullableHeartbeatControlSingle{value: val, isSet: true}
}

func (v NullableHeartbeatControlSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartbeatControlSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
