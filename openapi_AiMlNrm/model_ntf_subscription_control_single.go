/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// checks if the NtfSubscriptionControlSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtfSubscriptionControlSingle{}

// NtfSubscriptionControlSingle struct for NtfSubscriptionControlSingle
type NtfSubscriptionControlSingle struct {
	Top
	Attributes       *NtfSubscriptionControlSingleAllOfAttributes `json:"attributes,omitempty"`
	HeartbeatControl *HeartbeatControlSingle                      `json:"HeartbeatControl,omitempty"`
}

// NewNtfSubscriptionControlSingle instantiates a new NtfSubscriptionControlSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtfSubscriptionControlSingle(id NullableString) *NtfSubscriptionControlSingle {
	this := NtfSubscriptionControlSingle{}
	this.Id = id
	return &this
}

// NewNtfSubscriptionControlSingleWithDefaults instantiates a new NtfSubscriptionControlSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtfSubscriptionControlSingleWithDefaults() *NtfSubscriptionControlSingle {
	this := NtfSubscriptionControlSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetAttributes() NtfSubscriptionControlSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret NtfSubscriptionControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetAttributesOk() (*NtfSubscriptionControlSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NtfSubscriptionControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NtfSubscriptionControlSingle) SetAttributes(v NtfSubscriptionControlSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetHeartbeatControl returns the HeartbeatControl field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingle) GetHeartbeatControl() HeartbeatControlSingle {
	if o == nil || IsNil(o.HeartbeatControl) {
		var ret HeartbeatControlSingle
		return ret
	}
	return *o.HeartbeatControl
}

// GetHeartbeatControlOk returns a tuple with the HeartbeatControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingle) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool) {
	if o == nil || IsNil(o.HeartbeatControl) {
		return nil, false
	}
	return o.HeartbeatControl, true
}

// HasHeartbeatControl returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingle) HasHeartbeatControl() bool {
	if o != nil && !IsNil(o.HeartbeatControl) {
		return true
	}

	return false
}

// SetHeartbeatControl gets a reference to the given HeartbeatControlSingle and assigns it to the HeartbeatControl field.
func (o *NtfSubscriptionControlSingle) SetHeartbeatControl(v HeartbeatControlSingle) {
	o.HeartbeatControl = &v
}

func (o NtfSubscriptionControlSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtfSubscriptionControlSingle) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.HeartbeatControl) {
		toSerialize["HeartbeatControl"] = o.HeartbeatControl
	}
	return toSerialize, nil
}

type NullableNtfSubscriptionControlSingle struct {
	value *NtfSubscriptionControlSingle
	isSet bool
}

func (v NullableNtfSubscriptionControlSingle) Get() *NtfSubscriptionControlSingle {
	return v.value
}

func (v *NullableNtfSubscriptionControlSingle) Set(val *NtfSubscriptionControlSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNtfSubscriptionControlSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNtfSubscriptionControlSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtfSubscriptionControlSingle(val *NtfSubscriptionControlSingle) *NullableNtfSubscriptionControlSingle {
	return &NullableNtfSubscriptionControlSingle{value: val, isSet: true}
}

func (v NullableNtfSubscriptionControlSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtfSubscriptionControlSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
