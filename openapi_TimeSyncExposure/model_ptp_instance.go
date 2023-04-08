/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
)

// checks if the PtpInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PtpInstance{}

// PtpInstance Contains PTP instance configuration and activation requested by the AF.
type PtpInstance struct {
	InstanceType InstanceType `json:"instanceType"`
	Protocol Protocol `json:"protocol"`
	PtpProfile string `json:"ptpProfile"`
	PortConfigs []ConfigForPort `json:"portConfigs,omitempty"`
}

// NewPtpInstance instantiates a new PtpInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtpInstance(instanceType InstanceType, protocol Protocol, ptpProfile string) *PtpInstance {
	this := PtpInstance{}
	this.InstanceType = instanceType
	this.Protocol = protocol
	this.PtpProfile = ptpProfile
	return &this
}

// NewPtpInstanceWithDefaults instantiates a new PtpInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtpInstanceWithDefaults() *PtpInstance {
	this := PtpInstance{}
	return &this
}

// GetInstanceType returns the InstanceType field value
func (o *PtpInstance) GetInstanceType() InstanceType {
	if o == nil {
		var ret InstanceType
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *PtpInstance) GetInstanceTypeOk() (*InstanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *PtpInstance) SetInstanceType(v InstanceType) {
	o.InstanceType = v
}

// GetProtocol returns the Protocol field value
func (o *PtpInstance) GetProtocol() Protocol {
	if o == nil {
		var ret Protocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *PtpInstance) GetProtocolOk() (*Protocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *PtpInstance) SetProtocol(v Protocol) {
	o.Protocol = v
}

// GetPtpProfile returns the PtpProfile field value
func (o *PtpInstance) GetPtpProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PtpProfile
}

// GetPtpProfileOk returns a tuple with the PtpProfile field value
// and a boolean to check if the value has been set.
func (o *PtpInstance) GetPtpProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PtpProfile, true
}

// SetPtpProfile sets field value
func (o *PtpInstance) SetPtpProfile(v string) {
	o.PtpProfile = v
}

// GetPortConfigs returns the PortConfigs field value if set, zero value otherwise.
func (o *PtpInstance) GetPortConfigs() []ConfigForPort {
	if o == nil || isNil(o.PortConfigs) {
		var ret []ConfigForPort
		return ret
	}
	return o.PortConfigs
}

// GetPortConfigsOk returns a tuple with the PortConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtpInstance) GetPortConfigsOk() ([]ConfigForPort, bool) {
	if o == nil || isNil(o.PortConfigs) {
		return nil, false
	}
	return o.PortConfigs, true
}

// HasPortConfigs returns a boolean if a field has been set.
func (o *PtpInstance) HasPortConfigs() bool {
	if o != nil && !isNil(o.PortConfigs) {
		return true
	}

	return false
}

// SetPortConfigs gets a reference to the given []ConfigForPort and assigns it to the PortConfigs field.
func (o *PtpInstance) SetPortConfigs(v []ConfigForPort) {
	o.PortConfigs = v
}

func (o PtpInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PtpInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceType"] = o.InstanceType
	toSerialize["protocol"] = o.Protocol
	toSerialize["ptpProfile"] = o.PtpProfile
	if !isNil(o.PortConfigs) {
		toSerialize["portConfigs"] = o.PortConfigs
	}
	return toSerialize, nil
}

type NullablePtpInstance struct {
	value *PtpInstance
	isSet bool
}

func (v NullablePtpInstance) Get() *PtpInstance {
	return v.value
}

func (v *NullablePtpInstance) Set(val *PtpInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpInstance(val *PtpInstance) *NullablePtpInstance {
	return &NullablePtpInstance{value: val, isSet: true}
}

func (v NullablePtpInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


