/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// checks if the BindingResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindingResp{}

// BindingResp Contains the binding information for a PCF for a PDU Session.
type BindingResp struct {
	// Fully Qualified Domain Name
	PcfSmFqdn *string `json:"pcfSmFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_SMPolicyControl service.
	PcfSmIpEndPoints []IpEndPoint `json:"pcfSmIpEndPoints,omitempty"`
}

// NewBindingResp instantiates a new BindingResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindingResp() *BindingResp {
	this := BindingResp{}
	return &this
}

// NewBindingRespWithDefaults instantiates a new BindingResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindingRespWithDefaults() *BindingResp {
	this := BindingResp{}
	return &this
}

// GetPcfSmFqdn returns the PcfSmFqdn field value if set, zero value otherwise.
func (o *BindingResp) GetPcfSmFqdn() string {
	if o == nil || IsNil(o.PcfSmFqdn) {
		var ret string
		return ret
	}
	return *o.PcfSmFqdn
}

// GetPcfSmFqdnOk returns a tuple with the PcfSmFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindingResp) GetPcfSmFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PcfSmFqdn) {
		return nil, false
	}
	return o.PcfSmFqdn, true
}

// HasPcfSmFqdn returns a boolean if a field has been set.
func (o *BindingResp) HasPcfSmFqdn() bool {
	if o != nil && !IsNil(o.PcfSmFqdn) {
		return true
	}

	return false
}

// SetPcfSmFqdn gets a reference to the given string and assigns it to the PcfSmFqdn field.
func (o *BindingResp) SetPcfSmFqdn(v string) {
	o.PcfSmFqdn = &v
}

// GetPcfSmIpEndPoints returns the PcfSmIpEndPoints field value if set, zero value otherwise.
func (o *BindingResp) GetPcfSmIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.PcfSmIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfSmIpEndPoints
}

// GetPcfSmIpEndPointsOk returns a tuple with the PcfSmIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindingResp) GetPcfSmIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.PcfSmIpEndPoints) {
		return nil, false
	}
	return o.PcfSmIpEndPoints, true
}

// HasPcfSmIpEndPoints returns a boolean if a field has been set.
func (o *BindingResp) HasPcfSmIpEndPoints() bool {
	if o != nil && !IsNil(o.PcfSmIpEndPoints) {
		return true
	}

	return false
}

// SetPcfSmIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfSmIpEndPoints field.
func (o *BindingResp) SetPcfSmIpEndPoints(v []IpEndPoint) {
	o.PcfSmIpEndPoints = v
}

func (o BindingResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindingResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PcfSmFqdn) {
		toSerialize["pcfSmFqdn"] = o.PcfSmFqdn
	}
	if !IsNil(o.PcfSmIpEndPoints) {
		toSerialize["pcfSmIpEndPoints"] = o.PcfSmIpEndPoints
	}
	return toSerialize, nil
}

type NullableBindingResp struct {
	value *BindingResp
	isSet bool
}

func (v NullableBindingResp) Get() *BindingResp {
	return v.value
}

func (v *NullableBindingResp) Set(val *BindingResp) {
	v.value = val
	v.isSet = true
}

func (v NullableBindingResp) IsSet() bool {
	return v.isSet
}

func (v *NullableBindingResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindingResp(val *BindingResp) *NullableBindingResp {
	return &NullableBindingResp{value: val, isSet: true}
}

func (v NullableBindingResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindingResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
