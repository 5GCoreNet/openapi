/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the FlowInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowInfo{}

// FlowInfo Represents IP flow information.
type FlowInfo struct {
	// Indicates the IP flow identifier.
	FlowId int32 `json:"flowId"`
	// Indicates the packet filters of the IP flow. Refer to clause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
}

// NewFlowInfo instantiates a new FlowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInfo(flowId int32) *FlowInfo {
	this := FlowInfo{}
	this.FlowId = flowId
	return &this
}

// NewFlowInfoWithDefaults instantiates a new FlowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInfoWithDefaults() *FlowInfo {
	this := FlowInfo{}
	return &this
}

// GetFlowId returns the FlowId field value
func (o *FlowInfo) GetFlowId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value
// and a boolean to check if the value has been set.
func (o *FlowInfo) GetFlowIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowId, true
}

// SetFlowId sets field value
func (o *FlowInfo) SetFlowId(v int32) {
	o.FlowId = v
}

// GetFlowDescriptions returns the FlowDescriptions field value if set, zero value otherwise.
func (o *FlowInfo) GetFlowDescriptions() []string {
	if o == nil || IsNil(o.FlowDescriptions) {
		var ret []string
		return ret
	}
	return o.FlowDescriptions
}

// GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInfo) GetFlowDescriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.FlowDescriptions) {
		return nil, false
	}
	return o.FlowDescriptions, true
}

// HasFlowDescriptions returns a boolean if a field has been set.
func (o *FlowInfo) HasFlowDescriptions() bool {
	if o != nil && !IsNil(o.FlowDescriptions) {
		return true
	}

	return false
}

// SetFlowDescriptions gets a reference to the given []string and assigns it to the FlowDescriptions field.
func (o *FlowInfo) SetFlowDescriptions(v []string) {
	o.FlowDescriptions = v
}

func (o FlowInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flowId"] = o.FlowId
	if !IsNil(o.FlowDescriptions) {
		toSerialize["flowDescriptions"] = o.FlowDescriptions
	}
	return toSerialize, nil
}

type NullableFlowInfo struct {
	value *FlowInfo
	isSet bool
}

func (v NullableFlowInfo) Get() *FlowInfo {
	return v.value
}

func (v *NullableFlowInfo) Set(val *FlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInfo(val *FlowInfo) *NullableFlowInfo {
	return &NullableFlowInfo{value: val, isSet: true}
}

func (v NullableFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
