/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the PerFlowServiceExperienceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerFlowServiceExperienceInfo{}

// PerFlowServiceExperienceInfo struct for PerFlowServiceExperienceInfo
type PerFlowServiceExperienceInfo struct {
	ServiceExperience SvcExperience `json:"serviceExperience"`
	TimeInterval      TimeWindow    `json:"timeInterval"`
	RemoteEndpoint    AddrFqdn      `json:"remoteEndpoint"`
}

// NewPerFlowServiceExperienceInfo instantiates a new PerFlowServiceExperienceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerFlowServiceExperienceInfo(serviceExperience SvcExperience, timeInterval TimeWindow, remoteEndpoint AddrFqdn) *PerFlowServiceExperienceInfo {
	this := PerFlowServiceExperienceInfo{}
	this.ServiceExperience = serviceExperience
	this.TimeInterval = timeInterval
	this.RemoteEndpoint = remoteEndpoint
	return &this
}

// NewPerFlowServiceExperienceInfoWithDefaults instantiates a new PerFlowServiceExperienceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerFlowServiceExperienceInfoWithDefaults() *PerFlowServiceExperienceInfo {
	this := PerFlowServiceExperienceInfo{}
	return &this
}

// GetServiceExperience returns the ServiceExperience field value
func (o *PerFlowServiceExperienceInfo) GetServiceExperience() SvcExperience {
	if o == nil {
		var ret SvcExperience
		return ret
	}

	return o.ServiceExperience
}

// GetServiceExperienceOk returns a tuple with the ServiceExperience field value
// and a boolean to check if the value has been set.
func (o *PerFlowServiceExperienceInfo) GetServiceExperienceOk() (*SvcExperience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceExperience, true
}

// SetServiceExperience sets field value
func (o *PerFlowServiceExperienceInfo) SetServiceExperience(v SvcExperience) {
	o.ServiceExperience = v
}

// GetTimeInterval returns the TimeInterval field value
func (o *PerFlowServiceExperienceInfo) GetTimeInterval() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *PerFlowServiceExperienceInfo) GetTimeIntervalOk() (*TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value
func (o *PerFlowServiceExperienceInfo) SetTimeInterval(v TimeWindow) {
	o.TimeInterval = v
}

// GetRemoteEndpoint returns the RemoteEndpoint field value
func (o *PerFlowServiceExperienceInfo) GetRemoteEndpoint() AddrFqdn {
	if o == nil {
		var ret AddrFqdn
		return ret
	}

	return o.RemoteEndpoint
}

// GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field value
// and a boolean to check if the value has been set.
func (o *PerFlowServiceExperienceInfo) GetRemoteEndpointOk() (*AddrFqdn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteEndpoint, true
}

// SetRemoteEndpoint sets field value
func (o *PerFlowServiceExperienceInfo) SetRemoteEndpoint(v AddrFqdn) {
	o.RemoteEndpoint = v
}

func (o PerFlowServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerFlowServiceExperienceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceExperience"] = o.ServiceExperience
	toSerialize["timeInterval"] = o.TimeInterval
	toSerialize["remoteEndpoint"] = o.RemoteEndpoint
	return toSerialize, nil
}

type NullablePerFlowServiceExperienceInfo struct {
	value *PerFlowServiceExperienceInfo
	isSet bool
}

func (v NullablePerFlowServiceExperienceInfo) Get() *PerFlowServiceExperienceInfo {
	return v.value
}

func (v *NullablePerFlowServiceExperienceInfo) Set(val *PerFlowServiceExperienceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePerFlowServiceExperienceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePerFlowServiceExperienceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerFlowServiceExperienceInfo(val *PerFlowServiceExperienceInfo) *NullablePerFlowServiceExperienceInfo {
	return &NullablePerFlowServiceExperienceInfo{value: val, isSet: true}
}

func (v NullablePerFlowServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerFlowServiceExperienceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
