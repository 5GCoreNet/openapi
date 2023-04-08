/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the NtfSubscriptionControlSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtfSubscriptionControlSingleAllOf{}

// NtfSubscriptionControlSingleAllOf struct for NtfSubscriptionControlSingleAllOf
type NtfSubscriptionControlSingleAllOf struct {
	Attributes *NtfSubscriptionControlSingleAllOfAttributes `json:"attributes,omitempty"`
	HeartbeatControl *HeartbeatControlSingle `json:"HeartbeatControl,omitempty"`
}

// NewNtfSubscriptionControlSingleAllOf instantiates a new NtfSubscriptionControlSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtfSubscriptionControlSingleAllOf() *NtfSubscriptionControlSingleAllOf {
	this := NtfSubscriptionControlSingleAllOf{}
	return &this
}

// NewNtfSubscriptionControlSingleAllOfWithDefaults instantiates a new NtfSubscriptionControlSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtfSubscriptionControlSingleAllOfWithDefaults() *NtfSubscriptionControlSingleAllOf {
	this := NtfSubscriptionControlSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOf) GetAttributes() NtfSubscriptionControlSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret NtfSubscriptionControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOf) GetAttributesOk() (*NtfSubscriptionControlSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NtfSubscriptionControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *NtfSubscriptionControlSingleAllOf) SetAttributes(v NtfSubscriptionControlSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetHeartbeatControl returns the HeartbeatControl field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOf) GetHeartbeatControl() HeartbeatControlSingle {
	if o == nil || isNil(o.HeartbeatControl) {
		var ret HeartbeatControlSingle
		return ret
	}
	return *o.HeartbeatControl
}

// GetHeartbeatControlOk returns a tuple with the HeartbeatControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOf) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool) {
	if o == nil || isNil(o.HeartbeatControl) {
		return nil, false
	}
	return o.HeartbeatControl, true
}

// HasHeartbeatControl returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOf) HasHeartbeatControl() bool {
	if o != nil && !isNil(o.HeartbeatControl) {
		return true
	}

	return false
}

// SetHeartbeatControl gets a reference to the given HeartbeatControlSingle and assigns it to the HeartbeatControl field.
func (o *NtfSubscriptionControlSingleAllOf) SetHeartbeatControl(v HeartbeatControlSingle) {
	o.HeartbeatControl = &v
}

func (o NtfSubscriptionControlSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtfSubscriptionControlSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.HeartbeatControl) {
		toSerialize["HeartbeatControl"] = o.HeartbeatControl
	}
	return toSerialize, nil
}

type NullableNtfSubscriptionControlSingleAllOf struct {
	value *NtfSubscriptionControlSingleAllOf
	isSet bool
}

func (v NullableNtfSubscriptionControlSingleAllOf) Get() *NtfSubscriptionControlSingleAllOf {
	return v.value
}

func (v *NullableNtfSubscriptionControlSingleAllOf) Set(val *NtfSubscriptionControlSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNtfSubscriptionControlSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNtfSubscriptionControlSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtfSubscriptionControlSingleAllOf(val *NtfSubscriptionControlSingleAllOf) *NullableNtfSubscriptionControlSingleAllOf {
	return &NullableNtfSubscriptionControlSingleAllOf{value: val, isSet: true}
}

func (v NullableNtfSubscriptionControlSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtfSubscriptionControlSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


