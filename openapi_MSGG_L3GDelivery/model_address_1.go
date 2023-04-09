/*
MSGG_L3GDelivery

API for MSGG L3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGG_L3GDelivery

import (
	"encoding/json"
)

// checks if the Address1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address1{}

// Address1 Contains the Message type data
type Address1 struct {
	AddrType AddressType `json:"addrType"`
	Addr     string      `json:"addr"`
}

// NewAddress1 instantiates a new Address1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress1(addrType AddressType, addr string) *Address1 {
	this := Address1{}
	this.AddrType = addrType
	this.Addr = addr
	return &this
}

// NewAddress1WithDefaults instantiates a new Address1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress1WithDefaults() *Address1 {
	this := Address1{}
	return &this
}

// GetAddrType returns the AddrType field value
func (o *Address1) GetAddrType() AddressType {
	if o == nil {
		var ret AddressType
		return ret
	}

	return o.AddrType
}

// GetAddrTypeOk returns a tuple with the AddrType field value
// and a boolean to check if the value has been set.
func (o *Address1) GetAddrTypeOk() (*AddressType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddrType, true
}

// SetAddrType sets field value
func (o *Address1) SetAddrType(v AddressType) {
	o.AddrType = v
}

// GetAddr returns the Addr field value
func (o *Address1) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *Address1) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *Address1) SetAddr(v string) {
	o.Addr = v
}

func (o Address1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addrType"] = o.AddrType
	toSerialize["addr"] = o.Addr
	return toSerialize, nil
}

type NullableAddress1 struct {
	value *Address1
	isSet bool
}

func (v NullableAddress1) Get() *Address1 {
	return v.value
}

func (v *NullableAddress1) Set(val *Address1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress1(val *Address1) *NullableAddress1 {
	return &NullableAddress1{value: val, isSet: true}
}

func (v NullableAddress1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
