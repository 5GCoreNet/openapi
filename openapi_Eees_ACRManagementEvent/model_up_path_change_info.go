/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
)

// checks if the UpPathChangeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpPathChangeInfo{}

// UpPathChangeInfo Represents user plane path change information.
type UpPathChangeInfo struct {
	UeId               IndUeIdentification     `json:"ueId"`
	DnaiChgType        DnaiChangeType          `json:"dnaiChgType"`
	SourceTrafficRoute NullableRouteToLocation `json:"sourceTrafficRoute,omitempty"`
	TargetTrafficRoute NullableRouteToLocation `json:"targetTrafficRoute,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai *string `json:"sourceDnai,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	TargetDnai *string `json:"targetDnai,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	SrcUeIpv4Addr   *string     `json:"srcUeIpv4Addr,omitempty"`
	SrcUeIpv6Prefix *Ipv6Prefix `json:"srcUeIpv6Prefix,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	TgtUeIpv4Addr   *string     `json:"tgtUeIpv4Addr,omitempty"`
	TgtUeIpv6Prefix *Ipv6Prefix `json:"tgtUeIpv6Prefix,omitempty"`
}

// NewUpPathChangeInfo instantiates a new UpPathChangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpPathChangeInfo(ueId IndUeIdentification, dnaiChgType DnaiChangeType) *UpPathChangeInfo {
	this := UpPathChangeInfo{}
	this.UeId = ueId
	this.DnaiChgType = dnaiChgType
	return &this
}

// NewUpPathChangeInfoWithDefaults instantiates a new UpPathChangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpPathChangeInfoWithDefaults() *UpPathChangeInfo {
	this := UpPathChangeInfo{}
	return &this
}

// GetUeId returns the UeId field value
func (o *UpPathChangeInfo) GetUeId() IndUeIdentification {
	if o == nil {
		var ret IndUeIdentification
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetUeIdOk() (*IndUeIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *UpPathChangeInfo) SetUeId(v IndUeIdentification) {
	o.UeId = v
}

// GetDnaiChgType returns the DnaiChgType field value
func (o *UpPathChangeInfo) GetDnaiChgType() DnaiChangeType {
	if o == nil {
		var ret DnaiChangeType
		return ret
	}

	return o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnaiChgType, true
}

// SetDnaiChgType sets field value
func (o *UpPathChangeInfo) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = v
}

// GetSourceTrafficRoute returns the SourceTrafficRoute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpPathChangeInfo) GetSourceTrafficRoute() RouteToLocation {
	if o == nil || IsNil(o.SourceTrafficRoute.Get()) {
		var ret RouteToLocation
		return ret
	}
	return *o.SourceTrafficRoute.Get()
}

// GetSourceTrafficRouteOk returns a tuple with the SourceTrafficRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpPathChangeInfo) GetSourceTrafficRouteOk() (*RouteToLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceTrafficRoute.Get(), o.SourceTrafficRoute.IsSet()
}

// HasSourceTrafficRoute returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasSourceTrafficRoute() bool {
	if o != nil && o.SourceTrafficRoute.IsSet() {
		return true
	}

	return false
}

// SetSourceTrafficRoute gets a reference to the given NullableRouteToLocation and assigns it to the SourceTrafficRoute field.
func (o *UpPathChangeInfo) SetSourceTrafficRoute(v RouteToLocation) {
	o.SourceTrafficRoute.Set(&v)
}

// SetSourceTrafficRouteNil sets the value for SourceTrafficRoute to be an explicit nil
func (o *UpPathChangeInfo) SetSourceTrafficRouteNil() {
	o.SourceTrafficRoute.Set(nil)
}

// UnsetSourceTrafficRoute ensures that no value is present for SourceTrafficRoute, not even an explicit nil
func (o *UpPathChangeInfo) UnsetSourceTrafficRoute() {
	o.SourceTrafficRoute.Unset()
}

// GetTargetTrafficRoute returns the TargetTrafficRoute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpPathChangeInfo) GetTargetTrafficRoute() RouteToLocation {
	if o == nil || IsNil(o.TargetTrafficRoute.Get()) {
		var ret RouteToLocation
		return ret
	}
	return *o.TargetTrafficRoute.Get()
}

// GetTargetTrafficRouteOk returns a tuple with the TargetTrafficRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpPathChangeInfo) GetTargetTrafficRouteOk() (*RouteToLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetTrafficRoute.Get(), o.TargetTrafficRoute.IsSet()
}

// HasTargetTrafficRoute returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasTargetTrafficRoute() bool {
	if o != nil && o.TargetTrafficRoute.IsSet() {
		return true
	}

	return false
}

// SetTargetTrafficRoute gets a reference to the given NullableRouteToLocation and assigns it to the TargetTrafficRoute field.
func (o *UpPathChangeInfo) SetTargetTrafficRoute(v RouteToLocation) {
	o.TargetTrafficRoute.Set(&v)
}

// SetTargetTrafficRouteNil sets the value for TargetTrafficRoute to be an explicit nil
func (o *UpPathChangeInfo) SetTargetTrafficRouteNil() {
	o.TargetTrafficRoute.Set(nil)
}

// UnsetTargetTrafficRoute ensures that no value is present for TargetTrafficRoute, not even an explicit nil
func (o *UpPathChangeInfo) UnsetTargetTrafficRoute() {
	o.TargetTrafficRoute.Unset()
}

// GetSourceDnai returns the SourceDnai field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetSourceDnai() string {
	if o == nil || IsNil(o.SourceDnai) {
		var ret string
		return ret
	}
	return *o.SourceDnai
}

// GetSourceDnaiOk returns a tuple with the SourceDnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetSourceDnaiOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDnai) {
		return nil, false
	}
	return o.SourceDnai, true
}

// HasSourceDnai returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasSourceDnai() bool {
	if o != nil && !IsNil(o.SourceDnai) {
		return true
	}

	return false
}

// SetSourceDnai gets a reference to the given string and assigns it to the SourceDnai field.
func (o *UpPathChangeInfo) SetSourceDnai(v string) {
	o.SourceDnai = &v
}

// GetTargetDnai returns the TargetDnai field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetTargetDnai() string {
	if o == nil || IsNil(o.TargetDnai) {
		var ret string
		return ret
	}
	return *o.TargetDnai
}

// GetTargetDnaiOk returns a tuple with the TargetDnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetTargetDnaiOk() (*string, bool) {
	if o == nil || IsNil(o.TargetDnai) {
		return nil, false
	}
	return o.TargetDnai, true
}

// HasTargetDnai returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasTargetDnai() bool {
	if o != nil && !IsNil(o.TargetDnai) {
		return true
	}

	return false
}

// SetTargetDnai gets a reference to the given string and assigns it to the TargetDnai field.
func (o *UpPathChangeInfo) SetTargetDnai(v string) {
	o.TargetDnai = &v
}

// GetSrcUeIpv4Addr returns the SrcUeIpv4Addr field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetSrcUeIpv4Addr() string {
	if o == nil || IsNil(o.SrcUeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.SrcUeIpv4Addr
}

// GetSrcUeIpv4AddrOk returns a tuple with the SrcUeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetSrcUeIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.SrcUeIpv4Addr) {
		return nil, false
	}
	return o.SrcUeIpv4Addr, true
}

// HasSrcUeIpv4Addr returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasSrcUeIpv4Addr() bool {
	if o != nil && !IsNil(o.SrcUeIpv4Addr) {
		return true
	}

	return false
}

// SetSrcUeIpv4Addr gets a reference to the given string and assigns it to the SrcUeIpv4Addr field.
func (o *UpPathChangeInfo) SetSrcUeIpv4Addr(v string) {
	o.SrcUeIpv4Addr = &v
}

// GetSrcUeIpv6Prefix returns the SrcUeIpv6Prefix field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetSrcUeIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.SrcUeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.SrcUeIpv6Prefix
}

// GetSrcUeIpv6PrefixOk returns a tuple with the SrcUeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetSrcUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.SrcUeIpv6Prefix) {
		return nil, false
	}
	return o.SrcUeIpv6Prefix, true
}

// HasSrcUeIpv6Prefix returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasSrcUeIpv6Prefix() bool {
	if o != nil && !IsNil(o.SrcUeIpv6Prefix) {
		return true
	}

	return false
}

// SetSrcUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the SrcUeIpv6Prefix field.
func (o *UpPathChangeInfo) SetSrcUeIpv6Prefix(v Ipv6Prefix) {
	o.SrcUeIpv6Prefix = &v
}

// GetTgtUeIpv4Addr returns the TgtUeIpv4Addr field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetTgtUeIpv4Addr() string {
	if o == nil || IsNil(o.TgtUeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.TgtUeIpv4Addr
}

// GetTgtUeIpv4AddrOk returns a tuple with the TgtUeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetTgtUeIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.TgtUeIpv4Addr) {
		return nil, false
	}
	return o.TgtUeIpv4Addr, true
}

// HasTgtUeIpv4Addr returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasTgtUeIpv4Addr() bool {
	if o != nil && !IsNil(o.TgtUeIpv4Addr) {
		return true
	}

	return false
}

// SetTgtUeIpv4Addr gets a reference to the given string and assigns it to the TgtUeIpv4Addr field.
func (o *UpPathChangeInfo) SetTgtUeIpv4Addr(v string) {
	o.TgtUeIpv4Addr = &v
}

// GetTgtUeIpv6Prefix returns the TgtUeIpv6Prefix field value if set, zero value otherwise.
func (o *UpPathChangeInfo) GetTgtUeIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.TgtUeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.TgtUeIpv6Prefix
}

// GetTgtUeIpv6PrefixOk returns a tuple with the TgtUeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpPathChangeInfo) GetTgtUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.TgtUeIpv6Prefix) {
		return nil, false
	}
	return o.TgtUeIpv6Prefix, true
}

// HasTgtUeIpv6Prefix returns a boolean if a field has been set.
func (o *UpPathChangeInfo) HasTgtUeIpv6Prefix() bool {
	if o != nil && !IsNil(o.TgtUeIpv6Prefix) {
		return true
	}

	return false
}

// SetTgtUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the TgtUeIpv6Prefix field.
func (o *UpPathChangeInfo) SetTgtUeIpv6Prefix(v Ipv6Prefix) {
	o.TgtUeIpv6Prefix = &v
}

func (o UpPathChangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpPathChangeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueId"] = o.UeId
	toSerialize["dnaiChgType"] = o.DnaiChgType
	if o.SourceTrafficRoute.IsSet() {
		toSerialize["sourceTrafficRoute"] = o.SourceTrafficRoute.Get()
	}
	if o.TargetTrafficRoute.IsSet() {
		toSerialize["targetTrafficRoute"] = o.TargetTrafficRoute.Get()
	}
	if !IsNil(o.SourceDnai) {
		toSerialize["sourceDnai"] = o.SourceDnai
	}
	if !IsNil(o.TargetDnai) {
		toSerialize["targetDnai"] = o.TargetDnai
	}
	if !IsNil(o.SrcUeIpv4Addr) {
		toSerialize["srcUeIpv4Addr"] = o.SrcUeIpv4Addr
	}
	if !IsNil(o.SrcUeIpv6Prefix) {
		toSerialize["srcUeIpv6Prefix"] = o.SrcUeIpv6Prefix
	}
	if !IsNil(o.TgtUeIpv4Addr) {
		toSerialize["tgtUeIpv4Addr"] = o.TgtUeIpv4Addr
	}
	if !IsNil(o.TgtUeIpv6Prefix) {
		toSerialize["tgtUeIpv6Prefix"] = o.TgtUeIpv6Prefix
	}
	return toSerialize, nil
}

type NullableUpPathChangeInfo struct {
	value *UpPathChangeInfo
	isSet bool
}

func (v NullableUpPathChangeInfo) Get() *UpPathChangeInfo {
	return v.value
}

func (v *NullableUpPathChangeInfo) Set(val *UpPathChangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpPathChangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpPathChangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpPathChangeInfo(val *UpPathChangeInfo) *NullableUpPathChangeInfo {
	return &NullableUpPathChangeInfo{value: val, isSet: true}
}

func (v NullableUpPathChangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpPathChangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
