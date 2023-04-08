/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
)

// checks if the ServiceIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceIdentification{}

// ServiceIdentification Identifies the service to which the subscription applies.
type ServiceIdentification struct {
	ServEthFlows []EthernetFlowInfo `json:"servEthFlows,omitempty"`
	ServIpFlows []IpFlowInfo `json:"servIpFlows,omitempty"`
	// Contains an AF application identifier.
	AfAppId *string `json:"afAppId,omitempty"`
}

// NewServiceIdentification instantiates a new ServiceIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceIdentification() *ServiceIdentification {
	this := ServiceIdentification{}
	return &this
}

// NewServiceIdentificationWithDefaults instantiates a new ServiceIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceIdentificationWithDefaults() *ServiceIdentification {
	this := ServiceIdentification{}
	return &this
}

// GetServEthFlows returns the ServEthFlows field value if set, zero value otherwise.
func (o *ServiceIdentification) GetServEthFlows() []EthernetFlowInfo {
	if o == nil || isNil(o.ServEthFlows) {
		var ret []EthernetFlowInfo
		return ret
	}
	return o.ServEthFlows
}

// GetServEthFlowsOk returns a tuple with the ServEthFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIdentification) GetServEthFlowsOk() ([]EthernetFlowInfo, bool) {
	if o == nil || isNil(o.ServEthFlows) {
		return nil, false
	}
	return o.ServEthFlows, true
}

// HasServEthFlows returns a boolean if a field has been set.
func (o *ServiceIdentification) HasServEthFlows() bool {
	if o != nil && !isNil(o.ServEthFlows) {
		return true
	}

	return false
}

// SetServEthFlows gets a reference to the given []EthernetFlowInfo and assigns it to the ServEthFlows field.
func (o *ServiceIdentification) SetServEthFlows(v []EthernetFlowInfo) {
	o.ServEthFlows = v
}

// GetServIpFlows returns the ServIpFlows field value if set, zero value otherwise.
func (o *ServiceIdentification) GetServIpFlows() []IpFlowInfo {
	if o == nil || isNil(o.ServIpFlows) {
		var ret []IpFlowInfo
		return ret
	}
	return o.ServIpFlows
}

// GetServIpFlowsOk returns a tuple with the ServIpFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIdentification) GetServIpFlowsOk() ([]IpFlowInfo, bool) {
	if o == nil || isNil(o.ServIpFlows) {
		return nil, false
	}
	return o.ServIpFlows, true
}

// HasServIpFlows returns a boolean if a field has been set.
func (o *ServiceIdentification) HasServIpFlows() bool {
	if o != nil && !isNil(o.ServIpFlows) {
		return true
	}

	return false
}

// SetServIpFlows gets a reference to the given []IpFlowInfo and assigns it to the ServIpFlows field.
func (o *ServiceIdentification) SetServIpFlows(v []IpFlowInfo) {
	o.ServIpFlows = v
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *ServiceIdentification) GetAfAppId() string {
	if o == nil || isNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIdentification) GetAfAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *ServiceIdentification) HasAfAppId() bool {
	if o != nil && !isNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *ServiceIdentification) SetAfAppId(v string) {
	o.AfAppId = &v
}

func (o ServiceIdentification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServEthFlows) {
		toSerialize["servEthFlows"] = o.ServEthFlows
	}
	if !isNil(o.ServIpFlows) {
		toSerialize["servIpFlows"] = o.ServIpFlows
	}
	if !isNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	return toSerialize, nil
}

type NullableServiceIdentification struct {
	value *ServiceIdentification
	isSet bool
}

func (v NullableServiceIdentification) Get() *ServiceIdentification {
	return v.value
}

func (v *NullableServiceIdentification) Set(val *ServiceIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceIdentification(val *ServiceIdentification) *NullableServiceIdentification {
	return &NullableServiceIdentification{value: val, isSet: true}
}

func (v NullableServiceIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


