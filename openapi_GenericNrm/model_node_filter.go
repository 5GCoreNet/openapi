/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the NodeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeFilter{}

// NodeFilter struct for NodeFilter
type NodeFilter struct {
	AreaOfInterest *AreaOfInterest `json:"areaOfInterest,omitempty"`
	NetworkDomain *string `json:"networkDomain,omitempty"`
	CpUpType *string `json:"cpUpType,omitempty"`
	Sst *int32 `json:"sst,omitempty"`
}

// NewNodeFilter instantiates a new NodeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeFilter() *NodeFilter {
	this := NodeFilter{}
	return &this
}

// NewNodeFilterWithDefaults instantiates a new NodeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeFilterWithDefaults() *NodeFilter {
	this := NodeFilter{}
	return &this
}

// GetAreaOfInterest returns the AreaOfInterest field value if set, zero value otherwise.
func (o *NodeFilter) GetAreaOfInterest() AreaOfInterest {
	if o == nil || IsNil(o.AreaOfInterest) {
		var ret AreaOfInterest
		return ret
	}
	return *o.AreaOfInterest
}

// GetAreaOfInterestOk returns a tuple with the AreaOfInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFilter) GetAreaOfInterestOk() (*AreaOfInterest, bool) {
	if o == nil || IsNil(o.AreaOfInterest) {
		return nil, false
	}
	return o.AreaOfInterest, true
}

// HasAreaOfInterest returns a boolean if a field has been set.
func (o *NodeFilter) HasAreaOfInterest() bool {
	if o != nil && !IsNil(o.AreaOfInterest) {
		return true
	}

	return false
}

// SetAreaOfInterest gets a reference to the given AreaOfInterest and assigns it to the AreaOfInterest field.
func (o *NodeFilter) SetAreaOfInterest(v AreaOfInterest) {
	o.AreaOfInterest = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *NodeFilter) GetNetworkDomain() string {
	if o == nil || IsNil(o.NetworkDomain) {
		var ret string
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFilter) GetNetworkDomainOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkDomain) {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *NodeFilter) HasNetworkDomain() bool {
	if o != nil && !IsNil(o.NetworkDomain) {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given string and assigns it to the NetworkDomain field.
func (o *NodeFilter) SetNetworkDomain(v string) {
	o.NetworkDomain = &v
}

// GetCpUpType returns the CpUpType field value if set, zero value otherwise.
func (o *NodeFilter) GetCpUpType() string {
	if o == nil || IsNil(o.CpUpType) {
		var ret string
		return ret
	}
	return *o.CpUpType
}

// GetCpUpTypeOk returns a tuple with the CpUpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFilter) GetCpUpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CpUpType) {
		return nil, false
	}
	return o.CpUpType, true
}

// HasCpUpType returns a boolean if a field has been set.
func (o *NodeFilter) HasCpUpType() bool {
	if o != nil && !IsNil(o.CpUpType) {
		return true
	}

	return false
}

// SetCpUpType gets a reference to the given string and assigns it to the CpUpType field.
func (o *NodeFilter) SetCpUpType(v string) {
	o.CpUpType = &v
}

// GetSst returns the Sst field value if set, zero value otherwise.
func (o *NodeFilter) GetSst() int32 {
	if o == nil || IsNil(o.Sst) {
		var ret int32
		return ret
	}
	return *o.Sst
}

// GetSstOk returns a tuple with the Sst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFilter) GetSstOk() (*int32, bool) {
	if o == nil || IsNil(o.Sst) {
		return nil, false
	}
	return o.Sst, true
}

// HasSst returns a boolean if a field has been set.
func (o *NodeFilter) HasSst() bool {
	if o != nil && !IsNil(o.Sst) {
		return true
	}

	return false
}

// SetSst gets a reference to the given int32 and assigns it to the Sst field.
func (o *NodeFilter) SetSst(v int32) {
	o.Sst = &v
}

func (o NodeFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaOfInterest) {
		toSerialize["areaOfInterest"] = o.AreaOfInterest
	}
	if !IsNil(o.NetworkDomain) {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	if !IsNil(o.CpUpType) {
		toSerialize["cpUpType"] = o.CpUpType
	}
	if !IsNil(o.Sst) {
		toSerialize["sst"] = o.Sst
	}
	return toSerialize, nil
}

type NullableNodeFilter struct {
	value *NodeFilter
	isSet bool
}

func (v NullableNodeFilter) Get() *NodeFilter {
	return v.value
}

func (v *NullableNodeFilter) Set(val *NodeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeFilter(val *NodeFilter) *NullableNodeFilter {
	return &NullableNodeFilter{value: val, isSet: true}
}

func (v NullableNodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


