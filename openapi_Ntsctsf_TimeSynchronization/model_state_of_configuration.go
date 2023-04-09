/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
)

// checks if the StateOfConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateOfConfiguration{}

// StateOfConfiguration Contains the state of the time synchronization configuration.
type StateOfConfiguration struct {
	// When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for NW-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for NW-TT port is inactive. Default value is false.
	StateNwtt *bool `json:"stateNwtt,omitempty"`
	// Contains the PTP port states of the DS-TT(s).
	StateOfDstts []StateOfDstt `json:"stateOfDstts,omitempty"`
}

// NewStateOfConfiguration instantiates a new StateOfConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateOfConfiguration() *StateOfConfiguration {
	this := StateOfConfiguration{}
	return &this
}

// NewStateOfConfigurationWithDefaults instantiates a new StateOfConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateOfConfigurationWithDefaults() *StateOfConfiguration {
	this := StateOfConfiguration{}
	return &this
}

// GetStateNwtt returns the StateNwtt field value if set, zero value otherwise.
func (o *StateOfConfiguration) GetStateNwtt() bool {
	if o == nil || IsNil(o.StateNwtt) {
		var ret bool
		return ret
	}
	return *o.StateNwtt
}

// GetStateNwttOk returns a tuple with the StateNwtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateOfConfiguration) GetStateNwttOk() (*bool, bool) {
	if o == nil || IsNil(o.StateNwtt) {
		return nil, false
	}
	return o.StateNwtt, true
}

// HasStateNwtt returns a boolean if a field has been set.
func (o *StateOfConfiguration) HasStateNwtt() bool {
	if o != nil && !IsNil(o.StateNwtt) {
		return true
	}

	return false
}

// SetStateNwtt gets a reference to the given bool and assigns it to the StateNwtt field.
func (o *StateOfConfiguration) SetStateNwtt(v bool) {
	o.StateNwtt = &v
}

// GetStateOfDstts returns the StateOfDstts field value if set, zero value otherwise.
func (o *StateOfConfiguration) GetStateOfDstts() []StateOfDstt {
	if o == nil || IsNil(o.StateOfDstts) {
		var ret []StateOfDstt
		return ret
	}
	return o.StateOfDstts
}

// GetStateOfDsttsOk returns a tuple with the StateOfDstts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateOfConfiguration) GetStateOfDsttsOk() ([]StateOfDstt, bool) {
	if o == nil || IsNil(o.StateOfDstts) {
		return nil, false
	}
	return o.StateOfDstts, true
}

// HasStateOfDstts returns a boolean if a field has been set.
func (o *StateOfConfiguration) HasStateOfDstts() bool {
	if o != nil && !IsNil(o.StateOfDstts) {
		return true
	}

	return false
}

// SetStateOfDstts gets a reference to the given []StateOfDstt and assigns it to the StateOfDstts field.
func (o *StateOfConfiguration) SetStateOfDstts(v []StateOfDstt) {
	o.StateOfDstts = v
}

func (o StateOfConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateOfConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StateNwtt) {
		toSerialize["stateNwtt"] = o.StateNwtt
	}
	if !IsNil(o.StateOfDstts) {
		toSerialize["stateOfDstts"] = o.StateOfDstts
	}
	return toSerialize, nil
}

type NullableStateOfConfiguration struct {
	value *StateOfConfiguration
	isSet bool
}

func (v NullableStateOfConfiguration) Get() *StateOfConfiguration {
	return v.value
}

func (v *NullableStateOfConfiguration) Set(val *StateOfConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStateOfConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStateOfConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateOfConfiguration(val *StateOfConfiguration) *NullableStateOfConfiguration {
	return &NullableStateOfConfiguration{value: val, isSet: true}
}

func (v NullableStateOfConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateOfConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
