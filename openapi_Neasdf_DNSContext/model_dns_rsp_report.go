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

// checks if the DnsRspReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRspReport{}

// DnsRspReport DNS Response Event Report
type DnsRspReport struct {
	// Fully Qualified Domain Name
	Fqdn             *string    `json:"fqdn,omitempty"`
	EasIpv4Addresses []string   `json:"easIpv4Addresses,omitempty"`
	EasIpv6Addresses []Ipv6Addr `json:"easIpv6Addresses,omitempty"`
	EcsOption        *EcsOption `json:"ecsOption,omitempty"`
}

// NewDnsRspReport instantiates a new DnsRspReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRspReport() *DnsRspReport {
	this := DnsRspReport{}
	return &this
}

// NewDnsRspReportWithDefaults instantiates a new DnsRspReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRspReportWithDefaults() *DnsRspReport {
	this := DnsRspReport{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *DnsRspReport) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRspReport) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *DnsRspReport) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *DnsRspReport) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetEasIpv4Addresses returns the EasIpv4Addresses field value if set, zero value otherwise.
func (o *DnsRspReport) GetEasIpv4Addresses() []string {
	if o == nil || IsNil(o.EasIpv4Addresses) {
		var ret []string
		return ret
	}
	return o.EasIpv4Addresses
}

// GetEasIpv4AddressesOk returns a tuple with the EasIpv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRspReport) GetEasIpv4AddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.EasIpv4Addresses) {
		return nil, false
	}
	return o.EasIpv4Addresses, true
}

// HasEasIpv4Addresses returns a boolean if a field has been set.
func (o *DnsRspReport) HasEasIpv4Addresses() bool {
	if o != nil && !IsNil(o.EasIpv4Addresses) {
		return true
	}

	return false
}

// SetEasIpv4Addresses gets a reference to the given []string and assigns it to the EasIpv4Addresses field.
func (o *DnsRspReport) SetEasIpv4Addresses(v []string) {
	o.EasIpv4Addresses = v
}

// GetEasIpv6Addresses returns the EasIpv6Addresses field value if set, zero value otherwise.
func (o *DnsRspReport) GetEasIpv6Addresses() []Ipv6Addr {
	if o == nil || IsNil(o.EasIpv6Addresses) {
		var ret []Ipv6Addr
		return ret
	}
	return o.EasIpv6Addresses
}

// GetEasIpv6AddressesOk returns a tuple with the EasIpv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRspReport) GetEasIpv6AddressesOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.EasIpv6Addresses) {
		return nil, false
	}
	return o.EasIpv6Addresses, true
}

// HasEasIpv6Addresses returns a boolean if a field has been set.
func (o *DnsRspReport) HasEasIpv6Addresses() bool {
	if o != nil && !IsNil(o.EasIpv6Addresses) {
		return true
	}

	return false
}

// SetEasIpv6Addresses gets a reference to the given []Ipv6Addr and assigns it to the EasIpv6Addresses field.
func (o *DnsRspReport) SetEasIpv6Addresses(v []Ipv6Addr) {
	o.EasIpv6Addresses = v
}

// GetEcsOption returns the EcsOption field value if set, zero value otherwise.
func (o *DnsRspReport) GetEcsOption() EcsOption {
	if o == nil || IsNil(o.EcsOption) {
		var ret EcsOption
		return ret
	}
	return *o.EcsOption
}

// GetEcsOptionOk returns a tuple with the EcsOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRspReport) GetEcsOptionOk() (*EcsOption, bool) {
	if o == nil || IsNil(o.EcsOption) {
		return nil, false
	}
	return o.EcsOption, true
}

// HasEcsOption returns a boolean if a field has been set.
func (o *DnsRspReport) HasEcsOption() bool {
	if o != nil && !IsNil(o.EcsOption) {
		return true
	}

	return false
}

// SetEcsOption gets a reference to the given EcsOption and assigns it to the EcsOption field.
func (o *DnsRspReport) SetEcsOption(v EcsOption) {
	o.EcsOption = &v
}

func (o DnsRspReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRspReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.EasIpv4Addresses) {
		toSerialize["easIpv4Addresses"] = o.EasIpv4Addresses
	}
	if !IsNil(o.EasIpv6Addresses) {
		toSerialize["easIpv6Addresses"] = o.EasIpv6Addresses
	}
	if !IsNil(o.EcsOption) {
		toSerialize["ecsOption"] = o.EcsOption
	}
	return toSerialize, nil
}

type NullableDnsRspReport struct {
	value *DnsRspReport
	isSet bool
}

func (v NullableDnsRspReport) Get() *DnsRspReport {
	return v.value
}

func (v *NullableDnsRspReport) Set(val *DnsRspReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRspReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRspReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRspReport(val *DnsRspReport) *NullableDnsRspReport {
	return &NullableDnsRspReport{value: val, isSet: true}
}

func (v NullableDnsRspReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRspReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
