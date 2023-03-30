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

// checks if the BsfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BsfInfo{}

// BsfInfo Information of a BSF NF Instance
type BsfInfo struct {
	DnnList []string `json:"dnnList,omitempty"`
	IpDomainList []string `json:"ipDomainList,omitempty"`
	Ipv4AddressRanges []Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`
	Ipv6PrefixRanges []Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`
	// Fully Qualified Domain Name
	RxDiamHost *string `json:"rxDiamHost,omitempty"`
	// Fully Qualified Domain Name
	RxDiamRealm *string `json:"rxDiamRealm,omitempty"`
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
}

// NewBsfInfo instantiates a new BsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsfInfo() *BsfInfo {
	this := BsfInfo{}
	return &this
}

// NewBsfInfoWithDefaults instantiates a new BsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsfInfoWithDefaults() *BsfInfo {
	this := BsfInfo{}
	return &this
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *BsfInfo) GetDnnList() []string {
	if o == nil || IsNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetDnnListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *BsfInfo) HasDnnList() bool {
	if o != nil && !IsNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *BsfInfo) SetDnnList(v []string) {
	o.DnnList = v
}

// GetIpDomainList returns the IpDomainList field value if set, zero value otherwise.
func (o *BsfInfo) GetIpDomainList() []string {
	if o == nil || IsNil(o.IpDomainList) {
		var ret []string
		return ret
	}
	return o.IpDomainList
}

// GetIpDomainListOk returns a tuple with the IpDomainList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpDomainListOk() ([]string, bool) {
	if o == nil || IsNil(o.IpDomainList) {
		return nil, false
	}
	return o.IpDomainList, true
}

// HasIpDomainList returns a boolean if a field has been set.
func (o *BsfInfo) HasIpDomainList() bool {
	if o != nil && !IsNil(o.IpDomainList) {
		return true
	}

	return false
}

// SetIpDomainList gets a reference to the given []string and assigns it to the IpDomainList field.
func (o *BsfInfo) SetIpDomainList(v []string) {
	o.IpDomainList = v
}

// GetIpv4AddressRanges returns the Ipv4AddressRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetIpv4AddressRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.Ipv4AddressRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.Ipv4AddressRanges
}

// GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpv4AddressRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.Ipv4AddressRanges) {
		return nil, false
	}
	return o.Ipv4AddressRanges, true
}

// HasIpv4AddressRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasIpv4AddressRanges() bool {
	if o != nil && !IsNil(o.Ipv4AddressRanges) {
		return true
	}

	return false
}

// SetIpv4AddressRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddressRanges field.
func (o *BsfInfo) SetIpv4AddressRanges(v []Ipv4AddressRange) {
	o.Ipv4AddressRanges = v
}

// GetIpv6PrefixRanges returns the Ipv6PrefixRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		var ret []Ipv6PrefixRange
		return ret
	}
	return o.Ipv6PrefixRanges
}

// GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetIpv6PrefixRangesOk() ([]Ipv6PrefixRange, bool) {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		return nil, false
	}
	return o.Ipv6PrefixRanges, true
}

// HasIpv6PrefixRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasIpv6PrefixRanges() bool {
	if o != nil && !IsNil(o.Ipv6PrefixRanges) {
		return true
	}

	return false
}

// SetIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the Ipv6PrefixRanges field.
func (o *BsfInfo) SetIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.Ipv6PrefixRanges = v
}

// GetRxDiamHost returns the RxDiamHost field value if set, zero value otherwise.
func (o *BsfInfo) GetRxDiamHost() string {
	if o == nil || IsNil(o.RxDiamHost) {
		var ret string
		return ret
	}
	return *o.RxDiamHost
}

// GetRxDiamHostOk returns a tuple with the RxDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetRxDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamHost) {
		return nil, false
	}
	return o.RxDiamHost, true
}

// HasRxDiamHost returns a boolean if a field has been set.
func (o *BsfInfo) HasRxDiamHost() bool {
	if o != nil && !IsNil(o.RxDiamHost) {
		return true
	}

	return false
}

// SetRxDiamHost gets a reference to the given string and assigns it to the RxDiamHost field.
func (o *BsfInfo) SetRxDiamHost(v string) {
	o.RxDiamHost = &v
}

// GetRxDiamRealm returns the RxDiamRealm field value if set, zero value otherwise.
func (o *BsfInfo) GetRxDiamRealm() string {
	if o == nil || IsNil(o.RxDiamRealm) {
		var ret string
		return ret
	}
	return *o.RxDiamRealm
}

// GetRxDiamRealmOk returns a tuple with the RxDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetRxDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamRealm) {
		return nil, false
	}
	return o.RxDiamRealm, true
}

// HasRxDiamRealm returns a boolean if a field has been set.
func (o *BsfInfo) HasRxDiamRealm() bool {
	if o != nil && !IsNil(o.RxDiamRealm) {
		return true
	}

	return false
}

// SetRxDiamRealm gets a reference to the given string and assigns it to the RxDiamRealm field.
func (o *BsfInfo) SetRxDiamRealm(v string) {
	o.RxDiamRealm = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *BsfInfo) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *BsfInfo) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *BsfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *BsfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *BsfInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *BsfInfo) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *BsfInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

func (o BsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BsfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !IsNil(o.IpDomainList) {
		toSerialize["ipDomainList"] = o.IpDomainList
	}
	if !IsNil(o.Ipv4AddressRanges) {
		toSerialize["ipv4AddressRanges"] = o.Ipv4AddressRanges
	}
	if !IsNil(o.Ipv6PrefixRanges) {
		toSerialize["ipv6PrefixRanges"] = o.Ipv6PrefixRanges
	}
	if !IsNil(o.RxDiamHost) {
		toSerialize["rxDiamHost"] = o.RxDiamHost
	}
	if !IsNil(o.RxDiamRealm) {
		toSerialize["rxDiamRealm"] = o.RxDiamRealm
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	return toSerialize, nil
}

type NullableBsfInfo struct {
	value *BsfInfo
	isSet bool
}

func (v NullableBsfInfo) Get() *BsfInfo {
	return v.value
}

func (v *NullableBsfInfo) Set(val *BsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfInfo(val *BsfInfo) *NullableBsfInfo {
	return &NullableBsfInfo{value: val, isSet: true}
}

func (v NullableBsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


