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

// checks if the ServiceExperienceRecordAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceExperienceRecordAllOf{}

// ServiceExperienceRecordAllOf struct for ServiceExperienceRecordAllOf
type ServiceExperienceRecordAllOf struct {
	ServiceExperienceInfos PerFlowServiceExperienceInfo `json:"serviceExperienceInfos"`
}

// NewServiceExperienceRecordAllOf instantiates a new ServiceExperienceRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceRecordAllOf(serviceExperienceInfos PerFlowServiceExperienceInfo) *ServiceExperienceRecordAllOf {
	this := ServiceExperienceRecordAllOf{}
	this.ServiceExperienceInfos = serviceExperienceInfos
	return &this
}

// NewServiceExperienceRecordAllOfWithDefaults instantiates a new ServiceExperienceRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceRecordAllOfWithDefaults() *ServiceExperienceRecordAllOf {
	this := ServiceExperienceRecordAllOf{}
	return &this
}

// GetServiceExperienceInfos returns the ServiceExperienceInfos field value
func (o *ServiceExperienceRecordAllOf) GetServiceExperienceInfos() PerFlowServiceExperienceInfo {
	if o == nil {
		var ret PerFlowServiceExperienceInfo
		return ret
	}

	return o.ServiceExperienceInfos
}

// GetServiceExperienceInfosOk returns a tuple with the ServiceExperienceInfos field value
// and a boolean to check if the value has been set.
func (o *ServiceExperienceRecordAllOf) GetServiceExperienceInfosOk() (*PerFlowServiceExperienceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceExperienceInfos, true
}

// SetServiceExperienceInfos sets field value
func (o *ServiceExperienceRecordAllOf) SetServiceExperienceInfos(v PerFlowServiceExperienceInfo) {
	o.ServiceExperienceInfos = v
}

func (o ServiceExperienceRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceExperienceRecordAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceExperienceInfos"] = o.ServiceExperienceInfos
	return toSerialize, nil
}

type NullableServiceExperienceRecordAllOf struct {
	value *ServiceExperienceRecordAllOf
	isSet bool
}

func (v NullableServiceExperienceRecordAllOf) Get() *ServiceExperienceRecordAllOf {
	return v.value
}

func (v *NullableServiceExperienceRecordAllOf) Set(val *ServiceExperienceRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceRecordAllOf(val *ServiceExperienceRecordAllOf) *NullableServiceExperienceRecordAllOf {
	return &NullableServiceExperienceRecordAllOf{value: val, isSet: true}
}

func (v NullableServiceExperienceRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


