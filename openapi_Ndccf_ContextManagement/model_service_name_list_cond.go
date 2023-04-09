/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the ServiceNameListCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceNameListCond{}

// ServiceNameListCond Subscription to a set of NFs based on their support for a Service Name in the Servic Name list
type ServiceNameListCond struct {
	ConditionType   string        `json:"conditionType"`
	ServiceNameList []ServiceName `json:"serviceNameList"`
}

// NewServiceNameListCond instantiates a new ServiceNameListCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNameListCond(conditionType string, serviceNameList []ServiceName) *ServiceNameListCond {
	this := ServiceNameListCond{}
	this.ConditionType = conditionType
	this.ServiceNameList = serviceNameList
	return &this
}

// NewServiceNameListCondWithDefaults instantiates a new ServiceNameListCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNameListCondWithDefaults() *ServiceNameListCond {
	this := ServiceNameListCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *ServiceNameListCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *ServiceNameListCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *ServiceNameListCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetServiceNameList returns the ServiceNameList field value
func (o *ServiceNameListCond) GetServiceNameList() []ServiceName {
	if o == nil {
		var ret []ServiceName
		return ret
	}

	return o.ServiceNameList
}

// GetServiceNameListOk returns a tuple with the ServiceNameList field value
// and a boolean to check if the value has been set.
func (o *ServiceNameListCond) GetServiceNameListOk() ([]ServiceName, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceNameList, true
}

// SetServiceNameList sets field value
func (o *ServiceNameListCond) SetServiceNameList(v []ServiceName) {
	o.ServiceNameList = v
}

func (o ServiceNameListCond) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceNameListCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionType"] = o.ConditionType
	toSerialize["serviceNameList"] = o.ServiceNameList
	return toSerialize, nil
}

type NullableServiceNameListCond struct {
	value *ServiceNameListCond
	isSet bool
}

func (v NullableServiceNameListCond) Get() *ServiceNameListCond {
	return v.value
}

func (v *NullableServiceNameListCond) Set(val *ServiceNameListCond) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNameListCond) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNameListCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNameListCond(val *ServiceNameListCond) *NullableServiceNameListCond {
	return &NullableServiceNameListCond{value: val, isSet: true}
}

func (v NullableServiceNameListCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNameListCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
