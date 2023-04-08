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

// checks if the HeartbeatControlSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeartbeatControlSingleAllOfAttributes{}

// HeartbeatControlSingleAllOfAttributes struct for HeartbeatControlSingleAllOfAttributes
type HeartbeatControlSingleAllOfAttributes struct {
	HeartbeatNtfPeriod *int32 `json:"heartbeatNtfPeriod,omitempty"`
	TriggerHeartbeatNtf *bool `json:"triggerHeartbeatNtf,omitempty"`
}

// NewHeartbeatControlSingleAllOfAttributes instantiates a new HeartbeatControlSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeartbeatControlSingleAllOfAttributes() *HeartbeatControlSingleAllOfAttributes {
	this := HeartbeatControlSingleAllOfAttributes{}
	return &this
}

// NewHeartbeatControlSingleAllOfAttributesWithDefaults instantiates a new HeartbeatControlSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeartbeatControlSingleAllOfAttributesWithDefaults() *HeartbeatControlSingleAllOfAttributes {
	this := HeartbeatControlSingleAllOfAttributes{}
	return &this
}

// GetHeartbeatNtfPeriod returns the HeartbeatNtfPeriod field value if set, zero value otherwise.
func (o *HeartbeatControlSingleAllOfAttributes) GetHeartbeatNtfPeriod() int32 {
	if o == nil || isNil(o.HeartbeatNtfPeriod) {
		var ret int32
		return ret
	}
	return *o.HeartbeatNtfPeriod
}

// GetHeartbeatNtfPeriodOk returns a tuple with the HeartbeatNtfPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatControlSingleAllOfAttributes) GetHeartbeatNtfPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.HeartbeatNtfPeriod) {
		return nil, false
	}
	return o.HeartbeatNtfPeriod, true
}

// HasHeartbeatNtfPeriod returns a boolean if a field has been set.
func (o *HeartbeatControlSingleAllOfAttributes) HasHeartbeatNtfPeriod() bool {
	if o != nil && !isNil(o.HeartbeatNtfPeriod) {
		return true
	}

	return false
}

// SetHeartbeatNtfPeriod gets a reference to the given int32 and assigns it to the HeartbeatNtfPeriod field.
func (o *HeartbeatControlSingleAllOfAttributes) SetHeartbeatNtfPeriod(v int32) {
	o.HeartbeatNtfPeriod = &v
}

// GetTriggerHeartbeatNtf returns the TriggerHeartbeatNtf field value if set, zero value otherwise.
func (o *HeartbeatControlSingleAllOfAttributes) GetTriggerHeartbeatNtf() bool {
	if o == nil || isNil(o.TriggerHeartbeatNtf) {
		var ret bool
		return ret
	}
	return *o.TriggerHeartbeatNtf
}

// GetTriggerHeartbeatNtfOk returns a tuple with the TriggerHeartbeatNtf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatControlSingleAllOfAttributes) GetTriggerHeartbeatNtfOk() (*bool, bool) {
	if o == nil || isNil(o.TriggerHeartbeatNtf) {
		return nil, false
	}
	return o.TriggerHeartbeatNtf, true
}

// HasTriggerHeartbeatNtf returns a boolean if a field has been set.
func (o *HeartbeatControlSingleAllOfAttributes) HasTriggerHeartbeatNtf() bool {
	if o != nil && !isNil(o.TriggerHeartbeatNtf) {
		return true
	}

	return false
}

// SetTriggerHeartbeatNtf gets a reference to the given bool and assigns it to the TriggerHeartbeatNtf field.
func (o *HeartbeatControlSingleAllOfAttributes) SetTriggerHeartbeatNtf(v bool) {
	o.TriggerHeartbeatNtf = &v
}

func (o HeartbeatControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeartbeatControlSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HeartbeatNtfPeriod) {
		toSerialize["heartbeatNtfPeriod"] = o.HeartbeatNtfPeriod
	}
	if !isNil(o.TriggerHeartbeatNtf) {
		toSerialize["triggerHeartbeatNtf"] = o.TriggerHeartbeatNtf
	}
	return toSerialize, nil
}

type NullableHeartbeatControlSingleAllOfAttributes struct {
	value *HeartbeatControlSingleAllOfAttributes
	isSet bool
}

func (v NullableHeartbeatControlSingleAllOfAttributes) Get() *HeartbeatControlSingleAllOfAttributes {
	return v.value
}

func (v *NullableHeartbeatControlSingleAllOfAttributes) Set(val *HeartbeatControlSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartbeatControlSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartbeatControlSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartbeatControlSingleAllOfAttributes(val *HeartbeatControlSingleAllOfAttributes) *NullableHeartbeatControlSingleAllOfAttributes {
	return &NullableHeartbeatControlSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableHeartbeatControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartbeatControlSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


