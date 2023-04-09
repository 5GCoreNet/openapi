/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the ScpDomainInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScpDomainInfo{}

// ScpDomainInfo SCP Domain specific information
type ScpDomainInfo struct {
	// Fully Qualified Domain Name
	ScpFqdn        *string      `json:"scpFqdn,omitempty"`
	ScpIpEndPoints []IpEndPoint `json:"scpIpEndPoints,omitempty"`
	ScpPrefix      *string      `json:"scpPrefix,omitempty"`
	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\".
	ScpPorts *map[string]int32 `json:"scpPorts,omitempty"`
}

// NewScpDomainInfo instantiates a new ScpDomainInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpDomainInfo() *ScpDomainInfo {
	this := ScpDomainInfo{}
	return &this
}

// NewScpDomainInfoWithDefaults instantiates a new ScpDomainInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpDomainInfoWithDefaults() *ScpDomainInfo {
	this := ScpDomainInfo{}
	return &this
}

// GetScpFqdn returns the ScpFqdn field value if set, zero value otherwise.
func (o *ScpDomainInfo) GetScpFqdn() string {
	if o == nil || IsNil(o.ScpFqdn) {
		var ret string
		return ret
	}
	return *o.ScpFqdn
}

// GetScpFqdnOk returns a tuple with the ScpFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpDomainInfo) GetScpFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ScpFqdn) {
		return nil, false
	}
	return o.ScpFqdn, true
}

// HasScpFqdn returns a boolean if a field has been set.
func (o *ScpDomainInfo) HasScpFqdn() bool {
	if o != nil && !IsNil(o.ScpFqdn) {
		return true
	}

	return false
}

// SetScpFqdn gets a reference to the given string and assigns it to the ScpFqdn field.
func (o *ScpDomainInfo) SetScpFqdn(v string) {
	o.ScpFqdn = &v
}

// GetScpIpEndPoints returns the ScpIpEndPoints field value if set, zero value otherwise.
func (o *ScpDomainInfo) GetScpIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.ScpIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.ScpIpEndPoints
}

// GetScpIpEndPointsOk returns a tuple with the ScpIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpDomainInfo) GetScpIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.ScpIpEndPoints) {
		return nil, false
	}
	return o.ScpIpEndPoints, true
}

// HasScpIpEndPoints returns a boolean if a field has been set.
func (o *ScpDomainInfo) HasScpIpEndPoints() bool {
	if o != nil && !IsNil(o.ScpIpEndPoints) {
		return true
	}

	return false
}

// SetScpIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the ScpIpEndPoints field.
func (o *ScpDomainInfo) SetScpIpEndPoints(v []IpEndPoint) {
	o.ScpIpEndPoints = v
}

// GetScpPrefix returns the ScpPrefix field value if set, zero value otherwise.
func (o *ScpDomainInfo) GetScpPrefix() string {
	if o == nil || IsNil(o.ScpPrefix) {
		var ret string
		return ret
	}
	return *o.ScpPrefix
}

// GetScpPrefixOk returns a tuple with the ScpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpDomainInfo) GetScpPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ScpPrefix) {
		return nil, false
	}
	return o.ScpPrefix, true
}

// HasScpPrefix returns a boolean if a field has been set.
func (o *ScpDomainInfo) HasScpPrefix() bool {
	if o != nil && !IsNil(o.ScpPrefix) {
		return true
	}

	return false
}

// SetScpPrefix gets a reference to the given string and assigns it to the ScpPrefix field.
func (o *ScpDomainInfo) SetScpPrefix(v string) {
	o.ScpPrefix = &v
}

// GetScpPorts returns the ScpPorts field value if set, zero value otherwise.
func (o *ScpDomainInfo) GetScpPorts() map[string]int32 {
	if o == nil || IsNil(o.ScpPorts) {
		var ret map[string]int32
		return ret
	}
	return *o.ScpPorts
}

// GetScpPortsOk returns a tuple with the ScpPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpDomainInfo) GetScpPortsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.ScpPorts) {
		return nil, false
	}
	return o.ScpPorts, true
}

// HasScpPorts returns a boolean if a field has been set.
func (o *ScpDomainInfo) HasScpPorts() bool {
	if o != nil && !IsNil(o.ScpPorts) {
		return true
	}

	return false
}

// SetScpPorts gets a reference to the given map[string]int32 and assigns it to the ScpPorts field.
func (o *ScpDomainInfo) SetScpPorts(v map[string]int32) {
	o.ScpPorts = &v
}

func (o ScpDomainInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScpDomainInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScpFqdn) {
		toSerialize["scpFqdn"] = o.ScpFqdn
	}
	if !IsNil(o.ScpIpEndPoints) {
		toSerialize["scpIpEndPoints"] = o.ScpIpEndPoints
	}
	if !IsNil(o.ScpPrefix) {
		toSerialize["scpPrefix"] = o.ScpPrefix
	}
	if !IsNil(o.ScpPorts) {
		toSerialize["scpPorts"] = o.ScpPorts
	}
	return toSerialize, nil
}

type NullableScpDomainInfo struct {
	value *ScpDomainInfo
	isSet bool
}

func (v NullableScpDomainInfo) Get() *ScpDomainInfo {
	return v.value
}

func (v *NullableScpDomainInfo) Set(val *ScpDomainInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableScpDomainInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableScpDomainInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpDomainInfo(val *ScpDomainInfo) *NullableScpDomainInfo {
	return &NullableScpDomainInfo{value: val, isSet: true}
}

func (v NullableScpDomainInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpDomainInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
