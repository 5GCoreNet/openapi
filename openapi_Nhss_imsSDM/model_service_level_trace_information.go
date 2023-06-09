/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the ServiceLevelTraceInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLevelTraceInformation{}

// ServiceLevelTraceInformation IMS Service Level Trace Information
type ServiceLevelTraceInformation struct {
	ServiceLevelTraceInfo *string `json:"serviceLevelTraceInfo,omitempty"`
}

// NewServiceLevelTraceInformation instantiates a new ServiceLevelTraceInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelTraceInformation() *ServiceLevelTraceInformation {
	this := ServiceLevelTraceInformation{}
	return &this
}

// NewServiceLevelTraceInformationWithDefaults instantiates a new ServiceLevelTraceInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelTraceInformationWithDefaults() *ServiceLevelTraceInformation {
	this := ServiceLevelTraceInformation{}
	return &this
}

// GetServiceLevelTraceInfo returns the ServiceLevelTraceInfo field value if set, zero value otherwise.
func (o *ServiceLevelTraceInformation) GetServiceLevelTraceInfo() string {
	if o == nil || IsNil(o.ServiceLevelTraceInfo) {
		var ret string
		return ret
	}
	return *o.ServiceLevelTraceInfo
}

// GetServiceLevelTraceInfoOk returns a tuple with the ServiceLevelTraceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelTraceInformation) GetServiceLevelTraceInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevelTraceInfo) {
		return nil, false
	}
	return o.ServiceLevelTraceInfo, true
}

// HasServiceLevelTraceInfo returns a boolean if a field has been set.
func (o *ServiceLevelTraceInformation) HasServiceLevelTraceInfo() bool {
	if o != nil && !IsNil(o.ServiceLevelTraceInfo) {
		return true
	}

	return false
}

// SetServiceLevelTraceInfo gets a reference to the given string and assigns it to the ServiceLevelTraceInfo field.
func (o *ServiceLevelTraceInformation) SetServiceLevelTraceInfo(v string) {
	o.ServiceLevelTraceInfo = &v
}

func (o ServiceLevelTraceInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLevelTraceInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceLevelTraceInfo) {
		toSerialize["serviceLevelTraceInfo"] = o.ServiceLevelTraceInfo
	}
	return toSerialize, nil
}

type NullableServiceLevelTraceInformation struct {
	value *ServiceLevelTraceInformation
	isSet bool
}

func (v NullableServiceLevelTraceInformation) Get() *ServiceLevelTraceInformation {
	return v.value
}

func (v *NullableServiceLevelTraceInformation) Set(val *ServiceLevelTraceInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelTraceInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLevelTraceInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelTraceInformation(val *ServiceLevelTraceInformation) *NullableServiceLevelTraceInformation {
	return &NullableServiceLevelTraceInformation{value: val, isSet: true}
}

func (v NullableServiceLevelTraceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelTraceInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
