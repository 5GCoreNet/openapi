/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
)

// checks if the ServiceAreaCoverageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAreaCoverageInfo{}

// ServiceAreaCoverageInfo It represents a list of Tracking Areas within a serving network.
type ServiceAreaCoverageInfo struct {
	// Indicates a list of Tracking Areas where the service is allowed.
	TacList []string `json:"tacList"`
	ServingNetwork *PlmnIdNid `json:"servingNetwork,omitempty"`
}

// NewServiceAreaCoverageInfo instantiates a new ServiceAreaCoverageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAreaCoverageInfo(tacList []string) *ServiceAreaCoverageInfo {
	this := ServiceAreaCoverageInfo{}
	this.TacList = tacList
	return &this
}

// NewServiceAreaCoverageInfoWithDefaults instantiates a new ServiceAreaCoverageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaCoverageInfoWithDefaults() *ServiceAreaCoverageInfo {
	this := ServiceAreaCoverageInfo{}
	return &this
}

// GetTacList returns the TacList field value
func (o *ServiceAreaCoverageInfo) GetTacList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TacList
}

// GetTacListOk returns a tuple with the TacList field value
// and a boolean to check if the value has been set.
func (o *ServiceAreaCoverageInfo) GetTacListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TacList, true
}

// SetTacList sets field value
func (o *ServiceAreaCoverageInfo) SetTacList(v []string) {
	o.TacList = v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *ServiceAreaCoverageInfo) GetServingNetwork() PlmnIdNid {
	if o == nil || IsNil(o.ServingNetwork) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaCoverageInfo) GetServingNetworkOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.ServingNetwork) {
		return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *ServiceAreaCoverageInfo) HasServingNetwork() bool {
	if o != nil && !IsNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given PlmnIdNid and assigns it to the ServingNetwork field.
func (o *ServiceAreaCoverageInfo) SetServingNetwork(v PlmnIdNid) {
	o.ServingNetwork = &v
}

func (o ServiceAreaCoverageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAreaCoverageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tacList"] = o.TacList
	if !IsNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	return toSerialize, nil
}

type NullableServiceAreaCoverageInfo struct {
	value *ServiceAreaCoverageInfo
	isSet bool
}

func (v NullableServiceAreaCoverageInfo) Get() *ServiceAreaCoverageInfo {
	return v.value
}

func (v *NullableServiceAreaCoverageInfo) Set(val *ServiceAreaCoverageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAreaCoverageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAreaCoverageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAreaCoverageInfo(val *ServiceAreaCoverageInfo) *NullableServiceAreaCoverageInfo {
	return &NullableServiceAreaCoverageInfo{value: val, isSet: true}
}

func (v NullableServiceAreaCoverageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAreaCoverageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


