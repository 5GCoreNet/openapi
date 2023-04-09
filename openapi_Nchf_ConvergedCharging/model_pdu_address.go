/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the PDUAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PDUAddress{}

// PDUAddress struct for PDUAddress
type PDUAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	PduIPv4Address           *string      `json:"pduIPv4Address,omitempty"`
	PduIPv6AddresswithPrefix *Ipv6Addr    `json:"pduIPv6AddresswithPrefix,omitempty"`
	PduAddressprefixlength   *int32       `json:"pduAddressprefixlength,omitempty"`
	IPv4dynamicAddressFlag   *bool        `json:"iPv4dynamicAddressFlag,omitempty"`
	IPv6dynamicPrefixFlag    *bool        `json:"iPv6dynamicPrefixFlag,omitempty"`
	AddIpv6AddrPrefixes      *Ipv6Prefix  `json:"addIpv6AddrPrefixes,omitempty"`
	AddIpv6AddrPrefixList    []Ipv6Prefix `json:"addIpv6AddrPrefixList,omitempty"`
}

// NewPDUAddress instantiates a new PDUAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPDUAddress() *PDUAddress {
	this := PDUAddress{}
	return &this
}

// NewPDUAddressWithDefaults instantiates a new PDUAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPDUAddressWithDefaults() *PDUAddress {
	this := PDUAddress{}
	return &this
}

// GetPduIPv4Address returns the PduIPv4Address field value if set, zero value otherwise.
func (o *PDUAddress) GetPduIPv4Address() string {
	if o == nil || IsNil(o.PduIPv4Address) {
		var ret string
		return ret
	}
	return *o.PduIPv4Address
}

// GetPduIPv4AddressOk returns a tuple with the PduIPv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetPduIPv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.PduIPv4Address) {
		return nil, false
	}
	return o.PduIPv4Address, true
}

// HasPduIPv4Address returns a boolean if a field has been set.
func (o *PDUAddress) HasPduIPv4Address() bool {
	if o != nil && !IsNil(o.PduIPv4Address) {
		return true
	}

	return false
}

// SetPduIPv4Address gets a reference to the given string and assigns it to the PduIPv4Address field.
func (o *PDUAddress) SetPduIPv4Address(v string) {
	o.PduIPv4Address = &v
}

// GetPduIPv6AddresswithPrefix returns the PduIPv6AddresswithPrefix field value if set, zero value otherwise.
func (o *PDUAddress) GetPduIPv6AddresswithPrefix() Ipv6Addr {
	if o == nil || IsNil(o.PduIPv6AddresswithPrefix) {
		var ret Ipv6Addr
		return ret
	}
	return *o.PduIPv6AddresswithPrefix
}

// GetPduIPv6AddresswithPrefixOk returns a tuple with the PduIPv6AddresswithPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetPduIPv6AddresswithPrefixOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.PduIPv6AddresswithPrefix) {
		return nil, false
	}
	return o.PduIPv6AddresswithPrefix, true
}

// HasPduIPv6AddresswithPrefix returns a boolean if a field has been set.
func (o *PDUAddress) HasPduIPv6AddresswithPrefix() bool {
	if o != nil && !IsNil(o.PduIPv6AddresswithPrefix) {
		return true
	}

	return false
}

// SetPduIPv6AddresswithPrefix gets a reference to the given Ipv6Addr and assigns it to the PduIPv6AddresswithPrefix field.
func (o *PDUAddress) SetPduIPv6AddresswithPrefix(v Ipv6Addr) {
	o.PduIPv6AddresswithPrefix = &v
}

// GetPduAddressprefixlength returns the PduAddressprefixlength field value if set, zero value otherwise.
func (o *PDUAddress) GetPduAddressprefixlength() int32 {
	if o == nil || IsNil(o.PduAddressprefixlength) {
		var ret int32
		return ret
	}
	return *o.PduAddressprefixlength
}

// GetPduAddressprefixlengthOk returns a tuple with the PduAddressprefixlength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetPduAddressprefixlengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PduAddressprefixlength) {
		return nil, false
	}
	return o.PduAddressprefixlength, true
}

// HasPduAddressprefixlength returns a boolean if a field has been set.
func (o *PDUAddress) HasPduAddressprefixlength() bool {
	if o != nil && !IsNil(o.PduAddressprefixlength) {
		return true
	}

	return false
}

// SetPduAddressprefixlength gets a reference to the given int32 and assigns it to the PduAddressprefixlength field.
func (o *PDUAddress) SetPduAddressprefixlength(v int32) {
	o.PduAddressprefixlength = &v
}

// GetIPv4dynamicAddressFlag returns the IPv4dynamicAddressFlag field value if set, zero value otherwise.
func (o *PDUAddress) GetIPv4dynamicAddressFlag() bool {
	if o == nil || IsNil(o.IPv4dynamicAddressFlag) {
		var ret bool
		return ret
	}
	return *o.IPv4dynamicAddressFlag
}

// GetIPv4dynamicAddressFlagOk returns a tuple with the IPv4dynamicAddressFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetIPv4dynamicAddressFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.IPv4dynamicAddressFlag) {
		return nil, false
	}
	return o.IPv4dynamicAddressFlag, true
}

// HasIPv4dynamicAddressFlag returns a boolean if a field has been set.
func (o *PDUAddress) HasIPv4dynamicAddressFlag() bool {
	if o != nil && !IsNil(o.IPv4dynamicAddressFlag) {
		return true
	}

	return false
}

// SetIPv4dynamicAddressFlag gets a reference to the given bool and assigns it to the IPv4dynamicAddressFlag field.
func (o *PDUAddress) SetIPv4dynamicAddressFlag(v bool) {
	o.IPv4dynamicAddressFlag = &v
}

// GetIPv6dynamicPrefixFlag returns the IPv6dynamicPrefixFlag field value if set, zero value otherwise.
func (o *PDUAddress) GetIPv6dynamicPrefixFlag() bool {
	if o == nil || IsNil(o.IPv6dynamicPrefixFlag) {
		var ret bool
		return ret
	}
	return *o.IPv6dynamicPrefixFlag
}

// GetIPv6dynamicPrefixFlagOk returns a tuple with the IPv6dynamicPrefixFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetIPv6dynamicPrefixFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.IPv6dynamicPrefixFlag) {
		return nil, false
	}
	return o.IPv6dynamicPrefixFlag, true
}

// HasIPv6dynamicPrefixFlag returns a boolean if a field has been set.
func (o *PDUAddress) HasIPv6dynamicPrefixFlag() bool {
	if o != nil && !IsNil(o.IPv6dynamicPrefixFlag) {
		return true
	}

	return false
}

// SetIPv6dynamicPrefixFlag gets a reference to the given bool and assigns it to the IPv6dynamicPrefixFlag field.
func (o *PDUAddress) SetIPv6dynamicPrefixFlag(v bool) {
	o.IPv6dynamicPrefixFlag = &v
}

// GetAddIpv6AddrPrefixes returns the AddIpv6AddrPrefixes field value if set, zero value otherwise.
func (o *PDUAddress) GetAddIpv6AddrPrefixes() Ipv6Prefix {
	if o == nil || IsNil(o.AddIpv6AddrPrefixes) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.AddIpv6AddrPrefixes
}

// GetAddIpv6AddrPrefixesOk returns a tuple with the AddIpv6AddrPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetAddIpv6AddrPrefixesOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.AddIpv6AddrPrefixes) {
		return nil, false
	}
	return o.AddIpv6AddrPrefixes, true
}

// HasAddIpv6AddrPrefixes returns a boolean if a field has been set.
func (o *PDUAddress) HasAddIpv6AddrPrefixes() bool {
	if o != nil && !IsNil(o.AddIpv6AddrPrefixes) {
		return true
	}

	return false
}

// SetAddIpv6AddrPrefixes gets a reference to the given Ipv6Prefix and assigns it to the AddIpv6AddrPrefixes field.
func (o *PDUAddress) SetAddIpv6AddrPrefixes(v Ipv6Prefix) {
	o.AddIpv6AddrPrefixes = &v
}

// GetAddIpv6AddrPrefixList returns the AddIpv6AddrPrefixList field value if set, zero value otherwise.
func (o *PDUAddress) GetAddIpv6AddrPrefixList() []Ipv6Prefix {
	if o == nil || IsNil(o.AddIpv6AddrPrefixList) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.AddIpv6AddrPrefixList
}

// GetAddIpv6AddrPrefixListOk returns a tuple with the AddIpv6AddrPrefixList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PDUAddress) GetAddIpv6AddrPrefixListOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.AddIpv6AddrPrefixList) {
		return nil, false
	}
	return o.AddIpv6AddrPrefixList, true
}

// HasAddIpv6AddrPrefixList returns a boolean if a field has been set.
func (o *PDUAddress) HasAddIpv6AddrPrefixList() bool {
	if o != nil && !IsNil(o.AddIpv6AddrPrefixList) {
		return true
	}

	return false
}

// SetAddIpv6AddrPrefixList gets a reference to the given []Ipv6Prefix and assigns it to the AddIpv6AddrPrefixList field.
func (o *PDUAddress) SetAddIpv6AddrPrefixList(v []Ipv6Prefix) {
	o.AddIpv6AddrPrefixList = v
}

func (o PDUAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PDUAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PduIPv4Address) {
		toSerialize["pduIPv4Address"] = o.PduIPv4Address
	}
	if !IsNil(o.PduIPv6AddresswithPrefix) {
		toSerialize["pduIPv6AddresswithPrefix"] = o.PduIPv6AddresswithPrefix
	}
	if !IsNil(o.PduAddressprefixlength) {
		toSerialize["pduAddressprefixlength"] = o.PduAddressprefixlength
	}
	if !IsNil(o.IPv4dynamicAddressFlag) {
		toSerialize["iPv4dynamicAddressFlag"] = o.IPv4dynamicAddressFlag
	}
	if !IsNil(o.IPv6dynamicPrefixFlag) {
		toSerialize["iPv6dynamicPrefixFlag"] = o.IPv6dynamicPrefixFlag
	}
	if !IsNil(o.AddIpv6AddrPrefixes) {
		toSerialize["addIpv6AddrPrefixes"] = o.AddIpv6AddrPrefixes
	}
	if !IsNil(o.AddIpv6AddrPrefixList) {
		toSerialize["addIpv6AddrPrefixList"] = o.AddIpv6AddrPrefixList
	}
	return toSerialize, nil
}

type NullablePDUAddress struct {
	value *PDUAddress
	isSet bool
}

func (v NullablePDUAddress) Get() *PDUAddress {
	return v.value
}

func (v *NullablePDUAddress) Set(val *PDUAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePDUAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePDUAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePDUAddress(val *PDUAddress) *NullablePDUAddress {
	return &NullablePDUAddress{value: val, isSet: true}
}

func (v NullablePDUAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePDUAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
