/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EventExposure

import (
	"encoding/json"
)

// checks if the ServiceExperienceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceExperienceInfo{}

// ServiceExperienceInfo Contains service experience information associated with an application.
type ServiceExperienceInfo struct {
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	Supis []string `json:"supis,omitempty"`
	SvcExpPerFlows []ServiceExperienceInfoPerFlow `json:"svcExpPerFlows"`
}

// NewServiceExperienceInfo instantiates a new ServiceExperienceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceInfo(svcExpPerFlows []ServiceExperienceInfoPerFlow) *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	this.SvcExpPerFlows = svcExpPerFlows
	return &this
}

// NewServiceExperienceInfoWithDefaults instantiates a new ServiceExperienceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceInfoWithDefaults() *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceExperienceInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *ServiceExperienceInfo) SetSupis(v []string) {
	o.Supis = v
}

// GetSvcExpPerFlows returns the SvcExpPerFlows field value
func (o *ServiceExperienceInfo) GetSvcExpPerFlows() []ServiceExperienceInfoPerFlow {
	if o == nil {
		var ret []ServiceExperienceInfoPerFlow
		return ret
	}

	return o.SvcExpPerFlows
}

// GetSvcExpPerFlowsOk returns a tuple with the SvcExpPerFlows field value
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSvcExpPerFlowsOk() ([]ServiceExperienceInfoPerFlow, bool) {
	if o == nil {
		return nil, false
	}
	return o.SvcExpPerFlows, true
}

// SetSvcExpPerFlows sets field value
func (o *ServiceExperienceInfo) SetSvcExpPerFlows(v []ServiceExperienceInfoPerFlow) {
	o.SvcExpPerFlows = v
}

func (o ServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceExperienceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	toSerialize["svcExpPerFlows"] = o.SvcExpPerFlows
	return toSerialize, nil
}

type NullableServiceExperienceInfo struct {
	value *ServiceExperienceInfo
	isSet bool
}

func (v NullableServiceExperienceInfo) Get() *ServiceExperienceInfo {
	return v.value
}

func (v *NullableServiceExperienceInfo) Set(val *ServiceExperienceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceInfo(val *ServiceExperienceInfo) *NullableServiceExperienceInfo {
	return &NullableServiceExperienceInfo{value: val, isSet: true}
}

func (v NullableServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

