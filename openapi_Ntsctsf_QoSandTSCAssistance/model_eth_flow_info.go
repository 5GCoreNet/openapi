/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
)

// checks if the EthFlowInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthFlowInfo{}

// EthFlowInfo Represents Ethernet flow information.
type EthFlowInfo struct {
	// Indicates the Ethernet flow identifier.
	FlowId int32 `json:"flowId"`
	// Indicates the packet filters of the Ethernet flow. It shall contain UL and/or DL Ethernet flow description.
	EthFlowDescriptions []EthFlowDescription `json:"ethFlowDescriptions,omitempty"`
}

// NewEthFlowInfo instantiates a new EthFlowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthFlowInfo(flowId int32) *EthFlowInfo {
	this := EthFlowInfo{}
	this.FlowId = flowId
	return &this
}

// NewEthFlowInfoWithDefaults instantiates a new EthFlowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthFlowInfoWithDefaults() *EthFlowInfo {
	this := EthFlowInfo{}
	return &this
}

// GetFlowId returns the FlowId field value
func (o *EthFlowInfo) GetFlowId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value
// and a boolean to check if the value has been set.
func (o *EthFlowInfo) GetFlowIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlowId, true
}

// SetFlowId sets field value
func (o *EthFlowInfo) SetFlowId(v int32) {
	o.FlowId = v
}

// GetEthFlowDescriptions returns the EthFlowDescriptions field value if set, zero value otherwise.
func (o *EthFlowInfo) GetEthFlowDescriptions() []EthFlowDescription {
	if o == nil || IsNil(o.EthFlowDescriptions) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthFlowDescriptions
}

// GetEthFlowDescriptionsOk returns a tuple with the EthFlowDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowInfo) GetEthFlowDescriptionsOk() ([]EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthFlowDescriptions) {
		return nil, false
	}
	return o.EthFlowDescriptions, true
}

// HasEthFlowDescriptions returns a boolean if a field has been set.
func (o *EthFlowInfo) HasEthFlowDescriptions() bool {
	if o != nil && !IsNil(o.EthFlowDescriptions) {
		return true
	}

	return false
}

// SetEthFlowDescriptions gets a reference to the given []EthFlowDescription and assigns it to the EthFlowDescriptions field.
func (o *EthFlowInfo) SetEthFlowDescriptions(v []EthFlowDescription) {
	o.EthFlowDescriptions = v
}

func (o EthFlowInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthFlowInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flowId"] = o.FlowId
	if !IsNil(o.EthFlowDescriptions) {
		toSerialize["ethFlowDescriptions"] = o.EthFlowDescriptions
	}
	return toSerialize, nil
}

type NullableEthFlowInfo struct {
	value *EthFlowInfo
	isSet bool
}

func (v NullableEthFlowInfo) Get() *EthFlowInfo {
	return v.value
}

func (v *NullableEthFlowInfo) Set(val *EthFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEthFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEthFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthFlowInfo(val *EthFlowInfo) *NullableEthFlowInfo {
	return &NullableEthFlowInfo{value: val, isSet: true}
}

func (v NullableEthFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
