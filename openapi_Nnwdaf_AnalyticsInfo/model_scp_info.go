/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the ScpInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScpInfo{}

// ScpInfo Information of an SCP Instance
type ScpInfo struct {
	// A map (list of key-value pairs) where the key of the map shall be the string identifying an SCP domain
	ScpDomainInfoList *map[string]ScpDomainInfo `json:"scpDomainInfoList,omitempty"`
	ScpPrefix         *string                   `json:"scpPrefix,omitempty"`
	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\".
	ScpPorts          *map[string]int32  `json:"scpPorts,omitempty"`
	AddressDomains    []string           `json:"addressDomains,omitempty"`
	Ipv4Addresses     []string           `json:"ipv4Addresses,omitempty"`
	Ipv6Prefixes      []Ipv6Prefix       `json:"ipv6Prefixes,omitempty"`
	Ipv4AddrRanges    []Ipv4AddressRange `json:"ipv4AddrRanges,omitempty"`
	Ipv6PrefixRanges  []Ipv6PrefixRange  `json:"ipv6PrefixRanges,omitempty"`
	ServedNfSetIdList []string           `json:"servedNfSetIdList,omitempty"`
	RemotePlmnList    []PlmnId           `json:"remotePlmnList,omitempty"`
	RemoteSnpnList    []PlmnIdNid        `json:"remoteSnpnList,omitempty"`
	IpReachability    *IpReachability    `json:"ipReachability,omitempty"`
	ScpCapabilities   []ScpCapability    `json:"scpCapabilities,omitempty"`
}

// NewScpInfo instantiates a new ScpInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScpInfo() *ScpInfo {
	this := ScpInfo{}
	return &this
}

// NewScpInfoWithDefaults instantiates a new ScpInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScpInfoWithDefaults() *ScpInfo {
	this := ScpInfo{}
	return &this
}

// GetScpDomainInfoList returns the ScpDomainInfoList field value if set, zero value otherwise.
func (o *ScpInfo) GetScpDomainInfoList() map[string]ScpDomainInfo {
	if o == nil || IsNil(o.ScpDomainInfoList) {
		var ret map[string]ScpDomainInfo
		return ret
	}
	return *o.ScpDomainInfoList
}

// GetScpDomainInfoListOk returns a tuple with the ScpDomainInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetScpDomainInfoListOk() (*map[string]ScpDomainInfo, bool) {
	if o == nil || IsNil(o.ScpDomainInfoList) {
		return nil, false
	}
	return o.ScpDomainInfoList, true
}

// HasScpDomainInfoList returns a boolean if a field has been set.
func (o *ScpInfo) HasScpDomainInfoList() bool {
	if o != nil && !IsNil(o.ScpDomainInfoList) {
		return true
	}

	return false
}

// SetScpDomainInfoList gets a reference to the given map[string]ScpDomainInfo and assigns it to the ScpDomainInfoList field.
func (o *ScpInfo) SetScpDomainInfoList(v map[string]ScpDomainInfo) {
	o.ScpDomainInfoList = &v
}

// GetScpPrefix returns the ScpPrefix field value if set, zero value otherwise.
func (o *ScpInfo) GetScpPrefix() string {
	if o == nil || IsNil(o.ScpPrefix) {
		var ret string
		return ret
	}
	return *o.ScpPrefix
}

// GetScpPrefixOk returns a tuple with the ScpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetScpPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ScpPrefix) {
		return nil, false
	}
	return o.ScpPrefix, true
}

// HasScpPrefix returns a boolean if a field has been set.
func (o *ScpInfo) HasScpPrefix() bool {
	if o != nil && !IsNil(o.ScpPrefix) {
		return true
	}

	return false
}

// SetScpPrefix gets a reference to the given string and assigns it to the ScpPrefix field.
func (o *ScpInfo) SetScpPrefix(v string) {
	o.ScpPrefix = &v
}

// GetScpPorts returns the ScpPorts field value if set, zero value otherwise.
func (o *ScpInfo) GetScpPorts() map[string]int32 {
	if o == nil || IsNil(o.ScpPorts) {
		var ret map[string]int32
		return ret
	}
	return *o.ScpPorts
}

// GetScpPortsOk returns a tuple with the ScpPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetScpPortsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.ScpPorts) {
		return nil, false
	}
	return o.ScpPorts, true
}

// HasScpPorts returns a boolean if a field has been set.
func (o *ScpInfo) HasScpPorts() bool {
	if o != nil && !IsNil(o.ScpPorts) {
		return true
	}

	return false
}

// SetScpPorts gets a reference to the given map[string]int32 and assigns it to the ScpPorts field.
func (o *ScpInfo) SetScpPorts(v map[string]int32) {
	o.ScpPorts = &v
}

// GetAddressDomains returns the AddressDomains field value if set, zero value otherwise.
func (o *ScpInfo) GetAddressDomains() []string {
	if o == nil || IsNil(o.AddressDomains) {
		var ret []string
		return ret
	}
	return o.AddressDomains
}

// GetAddressDomainsOk returns a tuple with the AddressDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetAddressDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddressDomains) {
		return nil, false
	}
	return o.AddressDomains, true
}

// HasAddressDomains returns a boolean if a field has been set.
func (o *ScpInfo) HasAddressDomains() bool {
	if o != nil && !IsNil(o.AddressDomains) {
		return true
	}

	return false
}

// SetAddressDomains gets a reference to the given []string and assigns it to the AddressDomains field.
func (o *ScpInfo) SetAddressDomains(v []string) {
	o.AddressDomains = v
}

// GetIpv4Addresses returns the Ipv4Addresses field value if set, zero value otherwise.
func (o *ScpInfo) GetIpv4Addresses() []string {
	if o == nil || IsNil(o.Ipv4Addresses) {
		var ret []string
		return ret
	}
	return o.Ipv4Addresses
}

// GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetIpv4AddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4Addresses) {
		return nil, false
	}
	return o.Ipv4Addresses, true
}

// HasIpv4Addresses returns a boolean if a field has been set.
func (o *ScpInfo) HasIpv4Addresses() bool {
	if o != nil && !IsNil(o.Ipv4Addresses) {
		return true
	}

	return false
}

// SetIpv4Addresses gets a reference to the given []string and assigns it to the Ipv4Addresses field.
func (o *ScpInfo) SetIpv4Addresses(v []string) {
	o.Ipv4Addresses = v
}

// GetIpv6Prefixes returns the Ipv6Prefixes field value if set, zero value otherwise.
func (o *ScpInfo) GetIpv6Prefixes() []Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6Prefixes) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.Ipv6Prefixes
}

// GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetIpv6PrefixesOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6Prefixes) {
		return nil, false
	}
	return o.Ipv6Prefixes, true
}

// HasIpv6Prefixes returns a boolean if a field has been set.
func (o *ScpInfo) HasIpv6Prefixes() bool {
	if o != nil && !IsNil(o.Ipv6Prefixes) {
		return true
	}

	return false
}

// SetIpv6Prefixes gets a reference to the given []Ipv6Prefix and assigns it to the Ipv6Prefixes field.
func (o *ScpInfo) SetIpv6Prefixes(v []Ipv6Prefix) {
	o.Ipv6Prefixes = v
}

// GetIpv4AddrRanges returns the Ipv4AddrRanges field value if set, zero value otherwise.
func (o *ScpInfo) GetIpv4AddrRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.Ipv4AddrRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.Ipv4AddrRanges
}

// GetIpv4AddrRangesOk returns a tuple with the Ipv4AddrRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetIpv4AddrRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.Ipv4AddrRanges) {
		return nil, false
	}
	return o.Ipv4AddrRanges, true
}

// HasIpv4AddrRanges returns a boolean if a field has been set.
func (o *ScpInfo) HasIpv4AddrRanges() bool {
	if o != nil && !IsNil(o.Ipv4AddrRanges) {
		return true
	}

	return false
}

// SetIpv4AddrRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddrRanges field.
func (o *ScpInfo) SetIpv4AddrRanges(v []Ipv4AddressRange) {
	o.Ipv4AddrRanges = v
}

// GetIpv6PrefixRanges returns the Ipv6PrefixRanges field value if set, zero value otherwise.
func (o *ScpInfo) GetIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		var ret []Ipv6PrefixRange
		return ret
	}
	return o.Ipv6PrefixRanges
}

// GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetIpv6PrefixRangesOk() ([]Ipv6PrefixRange, bool) {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		return nil, false
	}
	return o.Ipv6PrefixRanges, true
}

// HasIpv6PrefixRanges returns a boolean if a field has been set.
func (o *ScpInfo) HasIpv6PrefixRanges() bool {
	if o != nil && !IsNil(o.Ipv6PrefixRanges) {
		return true
	}

	return false
}

// SetIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the Ipv6PrefixRanges field.
func (o *ScpInfo) SetIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.Ipv6PrefixRanges = v
}

// GetServedNfSetIdList returns the ServedNfSetIdList field value if set, zero value otherwise.
func (o *ScpInfo) GetServedNfSetIdList() []string {
	if o == nil || IsNil(o.ServedNfSetIdList) {
		var ret []string
		return ret
	}
	return o.ServedNfSetIdList
}

// GetServedNfSetIdListOk returns a tuple with the ServedNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetServedNfSetIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.ServedNfSetIdList) {
		return nil, false
	}
	return o.ServedNfSetIdList, true
}

// HasServedNfSetIdList returns a boolean if a field has been set.
func (o *ScpInfo) HasServedNfSetIdList() bool {
	if o != nil && !IsNil(o.ServedNfSetIdList) {
		return true
	}

	return false
}

// SetServedNfSetIdList gets a reference to the given []string and assigns it to the ServedNfSetIdList field.
func (o *ScpInfo) SetServedNfSetIdList(v []string) {
	o.ServedNfSetIdList = v
}

// GetRemotePlmnList returns the RemotePlmnList field value if set, zero value otherwise.
func (o *ScpInfo) GetRemotePlmnList() []PlmnId {
	if o == nil || IsNil(o.RemotePlmnList) {
		var ret []PlmnId
		return ret
	}
	return o.RemotePlmnList
}

// GetRemotePlmnListOk returns a tuple with the RemotePlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetRemotePlmnListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.RemotePlmnList) {
		return nil, false
	}
	return o.RemotePlmnList, true
}

// HasRemotePlmnList returns a boolean if a field has been set.
func (o *ScpInfo) HasRemotePlmnList() bool {
	if o != nil && !IsNil(o.RemotePlmnList) {
		return true
	}

	return false
}

// SetRemotePlmnList gets a reference to the given []PlmnId and assigns it to the RemotePlmnList field.
func (o *ScpInfo) SetRemotePlmnList(v []PlmnId) {
	o.RemotePlmnList = v
}

// GetRemoteSnpnList returns the RemoteSnpnList field value if set, zero value otherwise.
func (o *ScpInfo) GetRemoteSnpnList() []PlmnIdNid {
	if o == nil || IsNil(o.RemoteSnpnList) {
		var ret []PlmnIdNid
		return ret
	}
	return o.RemoteSnpnList
}

// GetRemoteSnpnListOk returns a tuple with the RemoteSnpnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetRemoteSnpnListOk() ([]PlmnIdNid, bool) {
	if o == nil || IsNil(o.RemoteSnpnList) {
		return nil, false
	}
	return o.RemoteSnpnList, true
}

// HasRemoteSnpnList returns a boolean if a field has been set.
func (o *ScpInfo) HasRemoteSnpnList() bool {
	if o != nil && !IsNil(o.RemoteSnpnList) {
		return true
	}

	return false
}

// SetRemoteSnpnList gets a reference to the given []PlmnIdNid and assigns it to the RemoteSnpnList field.
func (o *ScpInfo) SetRemoteSnpnList(v []PlmnIdNid) {
	o.RemoteSnpnList = v
}

// GetIpReachability returns the IpReachability field value if set, zero value otherwise.
func (o *ScpInfo) GetIpReachability() IpReachability {
	if o == nil || IsNil(o.IpReachability) {
		var ret IpReachability
		return ret
	}
	return *o.IpReachability
}

// GetIpReachabilityOk returns a tuple with the IpReachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetIpReachabilityOk() (*IpReachability, bool) {
	if o == nil || IsNil(o.IpReachability) {
		return nil, false
	}
	return o.IpReachability, true
}

// HasIpReachability returns a boolean if a field has been set.
func (o *ScpInfo) HasIpReachability() bool {
	if o != nil && !IsNil(o.IpReachability) {
		return true
	}

	return false
}

// SetIpReachability gets a reference to the given IpReachability and assigns it to the IpReachability field.
func (o *ScpInfo) SetIpReachability(v IpReachability) {
	o.IpReachability = &v
}

// GetScpCapabilities returns the ScpCapabilities field value if set, zero value otherwise.
func (o *ScpInfo) GetScpCapabilities() []ScpCapability {
	if o == nil || IsNil(o.ScpCapabilities) {
		var ret []ScpCapability
		return ret
	}
	return o.ScpCapabilities
}

// GetScpCapabilitiesOk returns a tuple with the ScpCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScpInfo) GetScpCapabilitiesOk() ([]ScpCapability, bool) {
	if o == nil || IsNil(o.ScpCapabilities) {
		return nil, false
	}
	return o.ScpCapabilities, true
}

// HasScpCapabilities returns a boolean if a field has been set.
func (o *ScpInfo) HasScpCapabilities() bool {
	if o != nil && !IsNil(o.ScpCapabilities) {
		return true
	}

	return false
}

// SetScpCapabilities gets a reference to the given []ScpCapability and assigns it to the ScpCapabilities field.
func (o *ScpInfo) SetScpCapabilities(v []ScpCapability) {
	o.ScpCapabilities = v
}

func (o ScpInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScpInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScpDomainInfoList) {
		toSerialize["scpDomainInfoList"] = o.ScpDomainInfoList
	}
	if !IsNil(o.ScpPrefix) {
		toSerialize["scpPrefix"] = o.ScpPrefix
	}
	if !IsNil(o.ScpPorts) {
		toSerialize["scpPorts"] = o.ScpPorts
	}
	if !IsNil(o.AddressDomains) {
		toSerialize["addressDomains"] = o.AddressDomains
	}
	if !IsNil(o.Ipv4Addresses) {
		toSerialize["ipv4Addresses"] = o.Ipv4Addresses
	}
	if !IsNil(o.Ipv6Prefixes) {
		toSerialize["ipv6Prefixes"] = o.Ipv6Prefixes
	}
	if !IsNil(o.Ipv4AddrRanges) {
		toSerialize["ipv4AddrRanges"] = o.Ipv4AddrRanges
	}
	if !IsNil(o.Ipv6PrefixRanges) {
		toSerialize["ipv6PrefixRanges"] = o.Ipv6PrefixRanges
	}
	if !IsNil(o.ServedNfSetIdList) {
		toSerialize["servedNfSetIdList"] = o.ServedNfSetIdList
	}
	if !IsNil(o.RemotePlmnList) {
		toSerialize["remotePlmnList"] = o.RemotePlmnList
	}
	if !IsNil(o.RemoteSnpnList) {
		toSerialize["remoteSnpnList"] = o.RemoteSnpnList
	}
	if !IsNil(o.IpReachability) {
		toSerialize["ipReachability"] = o.IpReachability
	}
	if !IsNil(o.ScpCapabilities) {
		toSerialize["scpCapabilities"] = o.ScpCapabilities
	}
	return toSerialize, nil
}

type NullableScpInfo struct {
	value *ScpInfo
	isSet bool
}

func (v NullableScpInfo) Get() *ScpInfo {
	return v.value
}

func (v *NullableScpInfo) Set(val *ScpInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableScpInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableScpInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpInfo(val *ScpInfo) *NullableScpInfo {
	return &NullableScpInfo{value: val, isSet: true}
}

func (v NullableScpInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
