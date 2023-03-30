/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Events_API

import (
	"encoding/json"
)

// checks if the RoutingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingRule{}

// RoutingRule Represents an API routing rule.
type RoutingRule struct {
	Ipv4AddrRanges []Ipv4AddressRange `json:"ipv4AddrRanges,omitempty"`
	Ipv6AddrRanges []Ipv6AddressRange `json:"ipv6AddrRanges,omitempty"`
	AefProfile AefProfile `json:"aefProfile"`
}

// NewRoutingRule instantiates a new RoutingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRule(aefProfile AefProfile) *RoutingRule {
	this := RoutingRule{}
	this.AefProfile = aefProfile
	return &this
}

// NewRoutingRuleWithDefaults instantiates a new RoutingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRuleWithDefaults() *RoutingRule {
	this := RoutingRule{}
	return &this
}

// GetIpv4AddrRanges returns the Ipv4AddrRanges field value if set, zero value otherwise.
func (o *RoutingRule) GetIpv4AddrRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.Ipv4AddrRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.Ipv4AddrRanges
}

// GetIpv4AddrRangesOk returns a tuple with the Ipv4AddrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRule) GetIpv4AddrRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.Ipv4AddrRanges) {
		return nil, false
	}
	return o.Ipv4AddrRanges, true
}

// HasIpv4AddrRanges returns a boolean if a field has been set.
func (o *RoutingRule) HasIpv4AddrRanges() bool {
	if o != nil && !IsNil(o.Ipv4AddrRanges) {
		return true
	}

	return false
}

// SetIpv4AddrRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddrRanges field.
func (o *RoutingRule) SetIpv4AddrRanges(v []Ipv4AddressRange) {
	o.Ipv4AddrRanges = v
}

// GetIpv6AddrRanges returns the Ipv6AddrRanges field value if set, zero value otherwise.
func (o *RoutingRule) GetIpv6AddrRanges() []Ipv6AddressRange {
	if o == nil || IsNil(o.Ipv6AddrRanges) {
		var ret []Ipv6AddressRange
		return ret
	}
	return o.Ipv6AddrRanges
}

// GetIpv6AddrRangesOk returns a tuple with the Ipv6AddrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRule) GetIpv6AddrRangesOk() ([]Ipv6AddressRange, bool) {
	if o == nil || IsNil(o.Ipv6AddrRanges) {
		return nil, false
	}
	return o.Ipv6AddrRanges, true
}

// HasIpv6AddrRanges returns a boolean if a field has been set.
func (o *RoutingRule) HasIpv6AddrRanges() bool {
	if o != nil && !IsNil(o.Ipv6AddrRanges) {
		return true
	}

	return false
}

// SetIpv6AddrRanges gets a reference to the given []Ipv6AddressRange and assigns it to the Ipv6AddrRanges field.
func (o *RoutingRule) SetIpv6AddrRanges(v []Ipv6AddressRange) {
	o.Ipv6AddrRanges = v
}

// GetAefProfile returns the AefProfile field value
func (o *RoutingRule) GetAefProfile() AefProfile {
	if o == nil {
		var ret AefProfile
		return ret
	}

	return o.AefProfile
}

// GetAefProfileOk returns a tuple with the AefProfile field value
// and a boolean to check if the value has been set.
func (o *RoutingRule) GetAefProfileOk() (*AefProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AefProfile, true
}

// SetAefProfile sets field value
func (o *RoutingRule) SetAefProfile(v AefProfile) {
	o.AefProfile = v
}

func (o RoutingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4AddrRanges) {
		toSerialize["ipv4AddrRanges"] = o.Ipv4AddrRanges
	}
	if !IsNil(o.Ipv6AddrRanges) {
		toSerialize["ipv6AddrRanges"] = o.Ipv6AddrRanges
	}
	toSerialize["aefProfile"] = o.AefProfile
	return toSerialize, nil
}

type NullableRoutingRule struct {
	value *RoutingRule
	isSet bool
}

func (v NullableRoutingRule) Get() *RoutingRule {
	return v.value
}

func (v *NullableRoutingRule) Set(val *RoutingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRule(val *RoutingRule) *NullableRoutingRule {
	return &NullableRoutingRule{value: val, isSet: true}
}

func (v NullableRoutingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

