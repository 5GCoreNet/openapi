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

// checks if the PcscfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcscfInfo{}

// PcscfInfo Information of a P-CSCF NF Instance
type PcscfInfo struct {
	AccessType []AccessType `json:"accessType,omitempty"`
	DnnList []string `json:"dnnList,omitempty"`
	// Fully Qualified Domain Name
	GmFqdn *string `json:"gmFqdn,omitempty"`
	GmIpv4Addresses []string `json:"gmIpv4Addresses,omitempty"`
	GmIpv6Addresses []Ipv6Addr `json:"gmIpv6Addresses,omitempty"`
	// Fully Qualified Domain Name
	MwFqdn *string `json:"mwFqdn,omitempty"`
	MwIpv4Addresses []string `json:"mwIpv4Addresses,omitempty"`
	MwIpv6Addresses []Ipv6Addr `json:"mwIpv6Addresses,omitempty"`
	ServedIpv4AddressRanges []Ipv4AddressRange `json:"servedIpv4AddressRanges,omitempty"`
	ServedIpv6PrefixRanges []Ipv6PrefixRange `json:"servedIpv6PrefixRanges,omitempty"`
}

// NewPcscfInfo instantiates a new PcscfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcscfInfo() *PcscfInfo {
	this := PcscfInfo{}
	return &this
}

// NewPcscfInfoWithDefaults instantiates a new PcscfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcscfInfoWithDefaults() *PcscfInfo {
	this := PcscfInfo{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *PcscfInfo) GetAccessType() []AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret []AccessType
		return ret
	}
	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetAccessTypeOk() ([]AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *PcscfInfo) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given []AccessType and assigns it to the AccessType field.
func (o *PcscfInfo) SetAccessType(v []AccessType) {
	o.AccessType = v
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *PcscfInfo) GetDnnList() []string {
	if o == nil || IsNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetDnnListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *PcscfInfo) HasDnnList() bool {
	if o != nil && !IsNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *PcscfInfo) SetDnnList(v []string) {
	o.DnnList = v
}

// GetGmFqdn returns the GmFqdn field value if set, zero value otherwise.
func (o *PcscfInfo) GetGmFqdn() string {
	if o == nil || IsNil(o.GmFqdn) {
		var ret string
		return ret
	}
	return *o.GmFqdn
}

// GetGmFqdnOk returns a tuple with the GmFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetGmFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.GmFqdn) {
		return nil, false
	}
	return o.GmFqdn, true
}

// HasGmFqdn returns a boolean if a field has been set.
func (o *PcscfInfo) HasGmFqdn() bool {
	if o != nil && !IsNil(o.GmFqdn) {
		return true
	}

	return false
}

// SetGmFqdn gets a reference to the given string and assigns it to the GmFqdn field.
func (o *PcscfInfo) SetGmFqdn(v string) {
	o.GmFqdn = &v
}

// GetGmIpv4Addresses returns the GmIpv4Addresses field value if set, zero value otherwise.
func (o *PcscfInfo) GetGmIpv4Addresses() []string {
	if o == nil || IsNil(o.GmIpv4Addresses) {
		var ret []string
		return ret
	}
	return o.GmIpv4Addresses
}

// GetGmIpv4AddressesOk returns a tuple with the GmIpv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetGmIpv4AddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.GmIpv4Addresses) {
		return nil, false
	}
	return o.GmIpv4Addresses, true
}

// HasGmIpv4Addresses returns a boolean if a field has been set.
func (o *PcscfInfo) HasGmIpv4Addresses() bool {
	if o != nil && !IsNil(o.GmIpv4Addresses) {
		return true
	}

	return false
}

// SetGmIpv4Addresses gets a reference to the given []string and assigns it to the GmIpv4Addresses field.
func (o *PcscfInfo) SetGmIpv4Addresses(v []string) {
	o.GmIpv4Addresses = v
}

// GetGmIpv6Addresses returns the GmIpv6Addresses field value if set, zero value otherwise.
func (o *PcscfInfo) GetGmIpv6Addresses() []Ipv6Addr {
	if o == nil || IsNil(o.GmIpv6Addresses) {
		var ret []Ipv6Addr
		return ret
	}
	return o.GmIpv6Addresses
}

// GetGmIpv6AddressesOk returns a tuple with the GmIpv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetGmIpv6AddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.GmIpv6Addresses) {
		return nil, false
	}
	return o.GmIpv6Addresses, true
}

// HasGmIpv6Addresses returns a boolean if a field has been set.
func (o *PcscfInfo) HasGmIpv6Addresses() bool {
	if o != nil && !IsNil(o.GmIpv6Addresses) {
		return true
	}

	return false
}

// SetGmIpv6Addresses gets a reference to the given []Ipv6Addr and assigns it to the GmIpv6Addresses field.
func (o *PcscfInfo) SetGmIpv6Addresses(v []Ipv6Addr) {
	o.GmIpv6Addresses = v
}

// GetMwFqdn returns the MwFqdn field value if set, zero value otherwise.
func (o *PcscfInfo) GetMwFqdn() string {
	if o == nil || IsNil(o.MwFqdn) {
		var ret string
		return ret
	}
	return *o.MwFqdn
}

// GetMwFqdnOk returns a tuple with the MwFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetMwFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.MwFqdn) {
		return nil, false
	}
	return o.MwFqdn, true
}

// HasMwFqdn returns a boolean if a field has been set.
func (o *PcscfInfo) HasMwFqdn() bool {
	if o != nil && !IsNil(o.MwFqdn) {
		return true
	}

	return false
}

// SetMwFqdn gets a reference to the given string and assigns it to the MwFqdn field.
func (o *PcscfInfo) SetMwFqdn(v string) {
	o.MwFqdn = &v
}

// GetMwIpv4Addresses returns the MwIpv4Addresses field value if set, zero value otherwise.
func (o *PcscfInfo) GetMwIpv4Addresses() []string {
	if o == nil || IsNil(o.MwIpv4Addresses) {
		var ret []string
		return ret
	}
	return o.MwIpv4Addresses
}

// GetMwIpv4AddressesOk returns a tuple with the MwIpv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetMwIpv4AddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.MwIpv4Addresses) {
		return nil, false
	}
	return o.MwIpv4Addresses, true
}

// HasMwIpv4Addresses returns a boolean if a field has been set.
func (o *PcscfInfo) HasMwIpv4Addresses() bool {
	if o != nil && !IsNil(o.MwIpv4Addresses) {
		return true
	}

	return false
}

// SetMwIpv4Addresses gets a reference to the given []string and assigns it to the MwIpv4Addresses field.
func (o *PcscfInfo) SetMwIpv4Addresses(v []string) {
	o.MwIpv4Addresses = v
}

// GetMwIpv6Addresses returns the MwIpv6Addresses field value if set, zero value otherwise.
func (o *PcscfInfo) GetMwIpv6Addresses() []Ipv6Addr {
	if o == nil || IsNil(o.MwIpv6Addresses) {
		var ret []Ipv6Addr
		return ret
	}
	return o.MwIpv6Addresses
}

// GetMwIpv6AddressesOk returns a tuple with the MwIpv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetMwIpv6AddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.MwIpv6Addresses) {
		return nil, false
	}
	return o.MwIpv6Addresses, true
}

// HasMwIpv6Addresses returns a boolean if a field has been set.
func (o *PcscfInfo) HasMwIpv6Addresses() bool {
	if o != nil && !IsNil(o.MwIpv6Addresses) {
		return true
	}

	return false
}

// SetMwIpv6Addresses gets a reference to the given []Ipv6Addr and assigns it to the MwIpv6Addresses field.
func (o *PcscfInfo) SetMwIpv6Addresses(v []Ipv6Addr) {
	o.MwIpv6Addresses = v
}

// GetServedIpv4AddressRanges returns the ServedIpv4AddressRanges field value if set, zero value otherwise.
func (o *PcscfInfo) GetServedIpv4AddressRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.ServedIpv4AddressRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.ServedIpv4AddressRanges
}

// GetServedIpv4AddressRangesOk returns a tuple with the ServedIpv4AddressRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetServedIpv4AddressRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.ServedIpv4AddressRanges) {
		return nil, false
	}
	return o.ServedIpv4AddressRanges, true
}

// HasServedIpv4AddressRanges returns a boolean if a field has been set.
func (o *PcscfInfo) HasServedIpv4AddressRanges() bool {
	if o != nil && !IsNil(o.ServedIpv4AddressRanges) {
		return true
	}

	return false
}

// SetServedIpv4AddressRanges gets a reference to the given []Ipv4AddressRange and assigns it to the ServedIpv4AddressRanges field.
func (o *PcscfInfo) SetServedIpv4AddressRanges(v []Ipv4AddressRange) {
	o.ServedIpv4AddressRanges = v
}

// GetServedIpv6PrefixRanges returns the ServedIpv6PrefixRanges field value if set, zero value otherwise.
func (o *PcscfInfo) GetServedIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || IsNil(o.ServedIpv6PrefixRanges) {
		var ret []Ipv6PrefixRange
		return ret
	}
	return o.ServedIpv6PrefixRanges
}

// GetServedIpv6PrefixRangesOk returns a tuple with the ServedIpv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfInfo) GetServedIpv6PrefixRangesOk() ([]Ipv6PrefixRange, bool) {
	if o == nil || IsNil(o.ServedIpv6PrefixRanges) {
		return nil, false
	}
	return o.ServedIpv6PrefixRanges, true
}

// HasServedIpv6PrefixRanges returns a boolean if a field has been set.
func (o *PcscfInfo) HasServedIpv6PrefixRanges() bool {
	if o != nil && !IsNil(o.ServedIpv6PrefixRanges) {
		return true
	}

	return false
}

// SetServedIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the ServedIpv6PrefixRanges field.
func (o *PcscfInfo) SetServedIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.ServedIpv6PrefixRanges = v
}

func (o PcscfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcscfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !IsNil(o.GmFqdn) {
		toSerialize["gmFqdn"] = o.GmFqdn
	}
	if !IsNil(o.GmIpv4Addresses) {
		toSerialize["gmIpv4Addresses"] = o.GmIpv4Addresses
	}
	if !IsNil(o.GmIpv6Addresses) {
		toSerialize["gmIpv6Addresses"] = o.GmIpv6Addresses
	}
	if !IsNil(o.MwFqdn) {
		toSerialize["mwFqdn"] = o.MwFqdn
	}
	if !IsNil(o.MwIpv4Addresses) {
		toSerialize["mwIpv4Addresses"] = o.MwIpv4Addresses
	}
	if !IsNil(o.MwIpv6Addresses) {
		toSerialize["mwIpv6Addresses"] = o.MwIpv6Addresses
	}
	if !IsNil(o.ServedIpv4AddressRanges) {
		toSerialize["servedIpv4AddressRanges"] = o.ServedIpv4AddressRanges
	}
	if !IsNil(o.ServedIpv6PrefixRanges) {
		toSerialize["servedIpv6PrefixRanges"] = o.ServedIpv6PrefixRanges
	}
	return toSerialize, nil
}

type NullablePcscfInfo struct {
	value *PcscfInfo
	isSet bool
}

func (v NullablePcscfInfo) Get() *PcscfInfo {
	return v.value
}

func (v *NullablePcscfInfo) Set(val *PcscfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcscfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcscfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcscfInfo(val *PcscfInfo) *NullablePcscfInfo {
	return &NullablePcscfInfo{value: val, isSet: true}
}

func (v NullablePcscfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcscfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


