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

// checks if the PcfBindingPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfBindingPatch{}

// PcfBindingPatch Identifies an Individual PCF binding used in an HTTP Patch method.
type PcfBindingPatch struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166 with the OpenAPI defined 'nullable: true' property.
	Ipv4Addr   NullableString       `json:"ipv4Addr,omitempty"`
	IpDomain   NullableString       `json:"ipDomain,omitempty"`
	Ipv6Prefix NullableIpv6PrefixRm `json:"ipv6Prefix,omitempty"`
	// The additional IPv6 Address Prefixes of the served UE.
	AddIpv6Prefixes []Ipv6Prefix `json:"addIpv6Prefixes,omitempty"`
	// \"String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042 with the OpenAPI 'nullable: true' property.\"
	MacAddr48 NullableString `json:"macAddr48,omitempty"`
	// The additional MAC Addresses of the served UE.
	AddMacAddrs []string `json:"addMacAddrs,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId *string `json:"pcfId,omitempty"`
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamHost *string `json:"pcfDiamHost,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamRealm *string `json:"pcfDiamRealm,omitempty"`
}

// NewPcfBindingPatch instantiates a new PcfBindingPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfBindingPatch() *PcfBindingPatch {
	this := PcfBindingPatch{}
	return &this
}

// NewPcfBindingPatchWithDefaults instantiates a new PcfBindingPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfBindingPatchWithDefaults() *PcfBindingPatch {
	this := PcfBindingPatch{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr.Get()
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetIpv4AddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv4Addr.Get(), o.Ipv4Addr.IsSet()
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasIpv4Addr() bool {
	if o != nil && o.Ipv4Addr.IsSet() {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given NullableString and assigns it to the Ipv4Addr field.
func (o *PcfBindingPatch) SetIpv4Addr(v string) {
	o.Ipv4Addr.Set(&v)
}

// SetIpv4AddrNil sets the value for Ipv4Addr to be an explicit nil
func (o *PcfBindingPatch) SetIpv4AddrNil() {
	o.Ipv4Addr.Set(nil)
}

// UnsetIpv4Addr ensures that no value is present for Ipv4Addr, not even an explicit nil
func (o *PcfBindingPatch) UnsetIpv4Addr() {
	o.Ipv4Addr.Unset()
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetIpDomain() string {
	if o == nil || IsNil(o.IpDomain.Get()) {
		var ret string
		return ret
	}
	return *o.IpDomain.Get()
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetIpDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpDomain.Get(), o.IpDomain.IsSet()
}

// HasIpDomain returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasIpDomain() bool {
	if o != nil && o.IpDomain.IsSet() {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given NullableString and assigns it to the IpDomain field.
func (o *PcfBindingPatch) SetIpDomain(v string) {
	o.IpDomain.Set(&v)
}

// SetIpDomainNil sets the value for IpDomain to be an explicit nil
func (o *PcfBindingPatch) SetIpDomainNil() {
	o.IpDomain.Set(nil)
}

// UnsetIpDomain ensures that no value is present for IpDomain, not even an explicit nil
func (o *PcfBindingPatch) UnsetIpDomain() {
	o.IpDomain.Unset()
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetIpv6Prefix() Ipv6PrefixRm {
	if o == nil || IsNil(o.Ipv6Prefix.Get()) {
		var ret Ipv6PrefixRm
		return ret
	}
	return *o.Ipv6Prefix.Get()
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetIpv6PrefixOk() (*Ipv6PrefixRm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6Prefix.Get(), o.Ipv6Prefix.IsSet()
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasIpv6Prefix() bool {
	if o != nil && o.Ipv6Prefix.IsSet() {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given NullableIpv6PrefixRm and assigns it to the Ipv6Prefix field.
func (o *PcfBindingPatch) SetIpv6Prefix(v Ipv6PrefixRm) {
	o.Ipv6Prefix.Set(&v)
}

// SetIpv6PrefixNil sets the value for Ipv6Prefix to be an explicit nil
func (o *PcfBindingPatch) SetIpv6PrefixNil() {
	o.Ipv6Prefix.Set(nil)
}

// UnsetIpv6Prefix ensures that no value is present for Ipv6Prefix, not even an explicit nil
func (o *PcfBindingPatch) UnsetIpv6Prefix() {
	o.Ipv6Prefix.Unset()
}

// GetAddIpv6Prefixes returns the AddIpv6Prefixes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetAddIpv6Prefixes() []Ipv6Prefix {
	if o == nil {
		var ret []Ipv6Prefix
		return ret
	}
	return o.AddIpv6Prefixes
}

// GetAddIpv6PrefixesOk returns a tuple with the AddIpv6Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetAddIpv6PrefixesOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.AddIpv6Prefixes) {
		return nil, false
	}
	return o.AddIpv6Prefixes, true
}

// HasAddIpv6Prefixes returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasAddIpv6Prefixes() bool {
	if o != nil && IsNil(o.AddIpv6Prefixes) {
		return true
	}

	return false
}

// SetAddIpv6Prefixes gets a reference to the given []Ipv6Prefix and assigns it to the AddIpv6Prefixes field.
func (o *PcfBindingPatch) SetAddIpv6Prefixes(v []Ipv6Prefix) {
	o.AddIpv6Prefixes = v
}

// GetMacAddr48 returns the MacAddr48 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetMacAddr48() string {
	if o == nil || IsNil(o.MacAddr48.Get()) {
		var ret string
		return ret
	}
	return *o.MacAddr48.Get()
}

// GetMacAddr48Ok returns a tuple with the MacAddr48 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetMacAddr48Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddr48.Get(), o.MacAddr48.IsSet()
}

// HasMacAddr48 returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasMacAddr48() bool {
	if o != nil && o.MacAddr48.IsSet() {
		return true
	}

	return false
}

// SetMacAddr48 gets a reference to the given NullableString and assigns it to the MacAddr48 field.
func (o *PcfBindingPatch) SetMacAddr48(v string) {
	o.MacAddr48.Set(&v)
}

// SetMacAddr48Nil sets the value for MacAddr48 to be an explicit nil
func (o *PcfBindingPatch) SetMacAddr48Nil() {
	o.MacAddr48.Set(nil)
}

// UnsetMacAddr48 ensures that no value is present for MacAddr48, not even an explicit nil
func (o *PcfBindingPatch) UnsetMacAddr48() {
	o.MacAddr48.Unset()
}

// GetAddMacAddrs returns the AddMacAddrs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PcfBindingPatch) GetAddMacAddrs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AddMacAddrs
}

// GetAddMacAddrsOk returns a tuple with the AddMacAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PcfBindingPatch) GetAddMacAddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddMacAddrs) {
		return nil, false
	}
	return o.AddMacAddrs, true
}

// HasAddMacAddrs returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasAddMacAddrs() bool {
	if o != nil && IsNil(o.AddMacAddrs) {
		return true
	}

	return false
}

// SetAddMacAddrs gets a reference to the given []string and assigns it to the AddMacAddrs field.
func (o *PcfBindingPatch) SetAddMacAddrs(v []string) {
	o.AddMacAddrs = v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *PcfBindingPatch) GetPcfId() string {
	if o == nil || IsNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBindingPatch) GetPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasPcfId() bool {
	if o != nil && !IsNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *PcfBindingPatch) SetPcfId(v string) {
	o.PcfId = &v
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *PcfBindingPatch) GetPcfFqdn() string {
	if o == nil || IsNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBindingPatch) GetPcfFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PcfFqdn) {
		return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasPcfFqdn() bool {
	if o != nil && !IsNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *PcfBindingPatch) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *PcfBindingPatch) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBindingPatch) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.PcfIpEndPoints) {
		return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasPcfIpEndPoints() bool {
	if o != nil && !IsNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *PcfBindingPatch) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

// GetPcfDiamHost returns the PcfDiamHost field value if set, zero value otherwise.
func (o *PcfBindingPatch) GetPcfDiamHost() string {
	if o == nil || IsNil(o.PcfDiamHost) {
		var ret string
		return ret
	}
	return *o.PcfDiamHost
}

// GetPcfDiamHostOk returns a tuple with the PcfDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBindingPatch) GetPcfDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.PcfDiamHost) {
		return nil, false
	}
	return o.PcfDiamHost, true
}

// HasPcfDiamHost returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasPcfDiamHost() bool {
	if o != nil && !IsNil(o.PcfDiamHost) {
		return true
	}

	return false
}

// SetPcfDiamHost gets a reference to the given string and assigns it to the PcfDiamHost field.
func (o *PcfBindingPatch) SetPcfDiamHost(v string) {
	o.PcfDiamHost = &v
}

// GetPcfDiamRealm returns the PcfDiamRealm field value if set, zero value otherwise.
func (o *PcfBindingPatch) GetPcfDiamRealm() string {
	if o == nil || IsNil(o.PcfDiamRealm) {
		var ret string
		return ret
	}
	return *o.PcfDiamRealm
}

// GetPcfDiamRealmOk returns a tuple with the PcfDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBindingPatch) GetPcfDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.PcfDiamRealm) {
		return nil, false
	}
	return o.PcfDiamRealm, true
}

// HasPcfDiamRealm returns a boolean if a field has been set.
func (o *PcfBindingPatch) HasPcfDiamRealm() bool {
	if o != nil && !IsNil(o.PcfDiamRealm) {
		return true
	}

	return false
}

// SetPcfDiamRealm gets a reference to the given string and assigns it to the PcfDiamRealm field.
func (o *PcfBindingPatch) SetPcfDiamRealm(v string) {
	o.PcfDiamRealm = &v
}

func (o PcfBindingPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfBindingPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipv4Addr.IsSet() {
		toSerialize["ipv4Addr"] = o.Ipv4Addr.Get()
	}
	if o.IpDomain.IsSet() {
		toSerialize["ipDomain"] = o.IpDomain.Get()
	}
	if o.Ipv6Prefix.IsSet() {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix.Get()
	}
	if o.AddIpv6Prefixes != nil {
		toSerialize["addIpv6Prefixes"] = o.AddIpv6Prefixes
	}
	if o.MacAddr48.IsSet() {
		toSerialize["macAddr48"] = o.MacAddr48.Get()
	}
	if o.AddMacAddrs != nil {
		toSerialize["addMacAddrs"] = o.AddMacAddrs
	}
	if !IsNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !IsNil(o.PcfFqdn) {
		toSerialize["pcfFqdn"] = o.PcfFqdn
	}
	if !IsNil(o.PcfIpEndPoints) {
		toSerialize["pcfIpEndPoints"] = o.PcfIpEndPoints
	}
	if !IsNil(o.PcfDiamHost) {
		toSerialize["pcfDiamHost"] = o.PcfDiamHost
	}
	if !IsNil(o.PcfDiamRealm) {
		toSerialize["pcfDiamRealm"] = o.PcfDiamRealm
	}
	return toSerialize, nil
}

type NullablePcfBindingPatch struct {
	value *PcfBindingPatch
	isSet bool
}

func (v NullablePcfBindingPatch) Get() *PcfBindingPatch {
	return v.value
}

func (v *NullablePcfBindingPatch) Set(val *PcfBindingPatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfBindingPatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfBindingPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfBindingPatch(val *PcfBindingPatch) *NullablePcfBindingPatch {
	return &NullablePcfBindingPatch{value: val, isSet: true}
}

func (v NullablePcfBindingPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfBindingPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
