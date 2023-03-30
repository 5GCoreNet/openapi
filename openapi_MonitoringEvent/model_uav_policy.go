/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the UavPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UavPolicy{}

// UavPolicy Represents the policy information included in the UAV presence monitoring request.
type UavPolicy struct {
	UavMoveInd bool `json:"uavMoveInd"`
	RevokeInd bool `json:"revokeInd"`
}

// NewUavPolicy instantiates a new UavPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUavPolicy(uavMoveInd bool, revokeInd bool) *UavPolicy {
	this := UavPolicy{}
	this.UavMoveInd = uavMoveInd
	this.RevokeInd = revokeInd
	return &this
}

// NewUavPolicyWithDefaults instantiates a new UavPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUavPolicyWithDefaults() *UavPolicy {
	this := UavPolicy{}
	return &this
}

// GetUavMoveInd returns the UavMoveInd field value
func (o *UavPolicy) GetUavMoveInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UavMoveInd
}

// GetUavMoveIndOk returns a tuple with the UavMoveInd field value
// and a boolean to check if the value has been set.
func (o *UavPolicy) GetUavMoveIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UavMoveInd, true
}

// SetUavMoveInd sets field value
func (o *UavPolicy) SetUavMoveInd(v bool) {
	o.UavMoveInd = v
}

// GetRevokeInd returns the RevokeInd field value
func (o *UavPolicy) GetRevokeInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RevokeInd
}

// GetRevokeIndOk returns a tuple with the RevokeInd field value
// and a boolean to check if the value has been set.
func (o *UavPolicy) GetRevokeIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevokeInd, true
}

// SetRevokeInd sets field value
func (o *UavPolicy) SetRevokeInd(v bool) {
	o.RevokeInd = v
}

func (o UavPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UavPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uavMoveInd"] = o.UavMoveInd
	toSerialize["revokeInd"] = o.RevokeInd
	return toSerialize, nil
}

type NullableUavPolicy struct {
	value *UavPolicy
	isSet bool
}

func (v NullableUavPolicy) Get() *UavPolicy {
	return v.value
}

func (v *NullableUavPolicy) Set(val *UavPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableUavPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableUavPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUavPolicy(val *UavPolicy) *NullableUavPolicy {
	return &NullableUavPolicy{value: val, isSet: true}
}

func (v NullableUavPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUavPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

