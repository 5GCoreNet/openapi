/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
)

// checks if the DnsQueryMdt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsQueryMdt{}

// DnsQueryMdt DNS Query message detection template
type DnsQueryMdt struct {
	MdtId string  `json:"mdtId"`
	Label *string `json:"label,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SourceIpv4Addr   *string                   `json:"sourceIpv4Addr,omitempty"`
	SourceIpv6Prefix *Ipv6Prefix               `json:"sourceIpv6Prefix,omitempty"`
	FqdnPatternList  []FqdnPatternMatchingRule `json:"fqdnPatternList,omitempty"`
}

// NewDnsQueryMdt instantiates a new DnsQueryMdt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsQueryMdt(mdtId string) *DnsQueryMdt {
	this := DnsQueryMdt{}
	this.MdtId = mdtId
	return &this
}

// NewDnsQueryMdtWithDefaults instantiates a new DnsQueryMdt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsQueryMdtWithDefaults() *DnsQueryMdt {
	this := DnsQueryMdt{}
	return &this
}

// GetMdtId returns the MdtId field value
func (o *DnsQueryMdt) GetMdtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MdtId
}

// GetMdtIdOk returns a tuple with the MdtId field value
// and a boolean to check if the value has been set.
func (o *DnsQueryMdt) GetMdtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MdtId, true
}

// SetMdtId sets field value
func (o *DnsQueryMdt) SetMdtId(v string) {
	o.MdtId = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DnsQueryMdt) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsQueryMdt) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DnsQueryMdt) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DnsQueryMdt) SetLabel(v string) {
	o.Label = &v
}

// GetSourceIpv4Addr returns the SourceIpv4Addr field value if set, zero value otherwise.
func (o *DnsQueryMdt) GetSourceIpv4Addr() string {
	if o == nil || IsNil(o.SourceIpv4Addr) {
		var ret string
		return ret
	}
	return *o.SourceIpv4Addr
}

// GetSourceIpv4AddrOk returns a tuple with the SourceIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsQueryMdt) GetSourceIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.SourceIpv4Addr) {
		return nil, false
	}
	return o.SourceIpv4Addr, true
}

// HasSourceIpv4Addr returns a boolean if a field has been set.
func (o *DnsQueryMdt) HasSourceIpv4Addr() bool {
	if o != nil && !IsNil(o.SourceIpv4Addr) {
		return true
	}

	return false
}

// SetSourceIpv4Addr gets a reference to the given string and assigns it to the SourceIpv4Addr field.
func (o *DnsQueryMdt) SetSourceIpv4Addr(v string) {
	o.SourceIpv4Addr = &v
}

// GetSourceIpv6Prefix returns the SourceIpv6Prefix field value if set, zero value otherwise.
func (o *DnsQueryMdt) GetSourceIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.SourceIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.SourceIpv6Prefix
}

// GetSourceIpv6PrefixOk returns a tuple with the SourceIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsQueryMdt) GetSourceIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.SourceIpv6Prefix) {
		return nil, false
	}
	return o.SourceIpv6Prefix, true
}

// HasSourceIpv6Prefix returns a boolean if a field has been set.
func (o *DnsQueryMdt) HasSourceIpv6Prefix() bool {
	if o != nil && !IsNil(o.SourceIpv6Prefix) {
		return true
	}

	return false
}

// SetSourceIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the SourceIpv6Prefix field.
func (o *DnsQueryMdt) SetSourceIpv6Prefix(v Ipv6Prefix) {
	o.SourceIpv6Prefix = &v
}

// GetFqdnPatternList returns the FqdnPatternList field value if set, zero value otherwise.
func (o *DnsQueryMdt) GetFqdnPatternList() []FqdnPatternMatchingRule {
	if o == nil || IsNil(o.FqdnPatternList) {
		var ret []FqdnPatternMatchingRule
		return ret
	}
	return o.FqdnPatternList
}

// GetFqdnPatternListOk returns a tuple with the FqdnPatternList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsQueryMdt) GetFqdnPatternListOk() ([]FqdnPatternMatchingRule, bool) {
	if o == nil || IsNil(o.FqdnPatternList) {
		return nil, false
	}
	return o.FqdnPatternList, true
}

// HasFqdnPatternList returns a boolean if a field has been set.
func (o *DnsQueryMdt) HasFqdnPatternList() bool {
	if o != nil && !IsNil(o.FqdnPatternList) {
		return true
	}

	return false
}

// SetFqdnPatternList gets a reference to the given []FqdnPatternMatchingRule and assigns it to the FqdnPatternList field.
func (o *DnsQueryMdt) SetFqdnPatternList(v []FqdnPatternMatchingRule) {
	o.FqdnPatternList = v
}

func (o DnsQueryMdt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsQueryMdt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mdtId"] = o.MdtId
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.SourceIpv4Addr) {
		toSerialize["sourceIpv4Addr"] = o.SourceIpv4Addr
	}
	if !IsNil(o.SourceIpv6Prefix) {
		toSerialize["sourceIpv6Prefix"] = o.SourceIpv6Prefix
	}
	if !IsNil(o.FqdnPatternList) {
		toSerialize["fqdnPatternList"] = o.FqdnPatternList
	}
	return toSerialize, nil
}

type NullableDnsQueryMdt struct {
	value *DnsQueryMdt
	isSet bool
}

func (v NullableDnsQueryMdt) Get() *DnsQueryMdt {
	return v.value
}

func (v *NullableDnsQueryMdt) Set(val *DnsQueryMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsQueryMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsQueryMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsQueryMdt(val *DnsQueryMdt) *NullableDnsQueryMdt {
	return &NullableDnsQueryMdt{value: val, isSet: true}
}

func (v NullableDnsQueryMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsQueryMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
