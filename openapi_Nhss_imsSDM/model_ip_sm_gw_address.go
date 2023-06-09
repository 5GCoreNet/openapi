/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the IpSmGwAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpSmGwAddress{}

// IpSmGwAddress IP-SM-GW number and diameter URI/realm
type IpSmGwAddress struct {
	// String containing an additional or basic MSISDN
	IpSmGwNumber string `json:"ipSmGwNumber"`
	// Fully Qualified Domain Name
	IpSmGwDiaUri *string `json:"ipSmGwDiaUri,omitempty"`
	// Fully Qualified Domain Name
	IpSmGwDiaRealm  *string `json:"ipSmGwDiaRealm,omitempty"`
	IpSmGwSbiSupInd *bool   `json:"ipSmGwSbiSupInd,omitempty"`
}

// NewIpSmGwAddress instantiates a new IpSmGwAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpSmGwAddress(ipSmGwNumber string) *IpSmGwAddress {
	this := IpSmGwAddress{}
	this.IpSmGwNumber = ipSmGwNumber
	var ipSmGwSbiSupInd bool = false
	this.IpSmGwSbiSupInd = &ipSmGwSbiSupInd
	return &this
}

// NewIpSmGwAddressWithDefaults instantiates a new IpSmGwAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpSmGwAddressWithDefaults() *IpSmGwAddress {
	this := IpSmGwAddress{}
	var ipSmGwSbiSupInd bool = false
	this.IpSmGwSbiSupInd = &ipSmGwSbiSupInd
	return &this
}

// GetIpSmGwNumber returns the IpSmGwNumber field value
func (o *IpSmGwAddress) GetIpSmGwNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpSmGwNumber
}

// GetIpSmGwNumberOk returns a tuple with the IpSmGwNumber field value
// and a boolean to check if the value has been set.
func (o *IpSmGwAddress) GetIpSmGwNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpSmGwNumber, true
}

// SetIpSmGwNumber sets field value
func (o *IpSmGwAddress) SetIpSmGwNumber(v string) {
	o.IpSmGwNumber = v
}

// GetIpSmGwDiaUri returns the IpSmGwDiaUri field value if set, zero value otherwise.
func (o *IpSmGwAddress) GetIpSmGwDiaUri() string {
	if o == nil || IsNil(o.IpSmGwDiaUri) {
		var ret string
		return ret
	}
	return *o.IpSmGwDiaUri
}

// GetIpSmGwDiaUriOk returns a tuple with the IpSmGwDiaUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpSmGwAddress) GetIpSmGwDiaUriOk() (*string, bool) {
	if o == nil || IsNil(o.IpSmGwDiaUri) {
		return nil, false
	}
	return o.IpSmGwDiaUri, true
}

// HasIpSmGwDiaUri returns a boolean if a field has been set.
func (o *IpSmGwAddress) HasIpSmGwDiaUri() bool {
	if o != nil && !IsNil(o.IpSmGwDiaUri) {
		return true
	}

	return false
}

// SetIpSmGwDiaUri gets a reference to the given string and assigns it to the IpSmGwDiaUri field.
func (o *IpSmGwAddress) SetIpSmGwDiaUri(v string) {
	o.IpSmGwDiaUri = &v
}

// GetIpSmGwDiaRealm returns the IpSmGwDiaRealm field value if set, zero value otherwise.
func (o *IpSmGwAddress) GetIpSmGwDiaRealm() string {
	if o == nil || IsNil(o.IpSmGwDiaRealm) {
		var ret string
		return ret
	}
	return *o.IpSmGwDiaRealm
}

// GetIpSmGwDiaRealmOk returns a tuple with the IpSmGwDiaRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpSmGwAddress) GetIpSmGwDiaRealmOk() (*string, bool) {
	if o == nil || IsNil(o.IpSmGwDiaRealm) {
		return nil, false
	}
	return o.IpSmGwDiaRealm, true
}

// HasIpSmGwDiaRealm returns a boolean if a field has been set.
func (o *IpSmGwAddress) HasIpSmGwDiaRealm() bool {
	if o != nil && !IsNil(o.IpSmGwDiaRealm) {
		return true
	}

	return false
}

// SetIpSmGwDiaRealm gets a reference to the given string and assigns it to the IpSmGwDiaRealm field.
func (o *IpSmGwAddress) SetIpSmGwDiaRealm(v string) {
	o.IpSmGwDiaRealm = &v
}

// GetIpSmGwSbiSupInd returns the IpSmGwSbiSupInd field value if set, zero value otherwise.
func (o *IpSmGwAddress) GetIpSmGwSbiSupInd() bool {
	if o == nil || IsNil(o.IpSmGwSbiSupInd) {
		var ret bool
		return ret
	}
	return *o.IpSmGwSbiSupInd
}

// GetIpSmGwSbiSupIndOk returns a tuple with the IpSmGwSbiSupInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpSmGwAddress) GetIpSmGwSbiSupIndOk() (*bool, bool) {
	if o == nil || IsNil(o.IpSmGwSbiSupInd) {
		return nil, false
	}
	return o.IpSmGwSbiSupInd, true
}

// HasIpSmGwSbiSupInd returns a boolean if a field has been set.
func (o *IpSmGwAddress) HasIpSmGwSbiSupInd() bool {
	if o != nil && !IsNil(o.IpSmGwSbiSupInd) {
		return true
	}

	return false
}

// SetIpSmGwSbiSupInd gets a reference to the given bool and assigns it to the IpSmGwSbiSupInd field.
func (o *IpSmGwAddress) SetIpSmGwSbiSupInd(v bool) {
	o.IpSmGwSbiSupInd = &v
}

func (o IpSmGwAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpSmGwAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ipSmGwNumber"] = o.IpSmGwNumber
	if !IsNil(o.IpSmGwDiaUri) {
		toSerialize["ipSmGwDiaUri"] = o.IpSmGwDiaUri
	}
	if !IsNil(o.IpSmGwDiaRealm) {
		toSerialize["ipSmGwDiaRealm"] = o.IpSmGwDiaRealm
	}
	if !IsNil(o.IpSmGwSbiSupInd) {
		toSerialize["ipSmGwSbiSupInd"] = o.IpSmGwSbiSupInd
	}
	return toSerialize, nil
}

type NullableIpSmGwAddress struct {
	value *IpSmGwAddress
	isSet bool
}

func (v NullableIpSmGwAddress) Get() *IpSmGwAddress {
	return v.value
}

func (v *NullableIpSmGwAddress) Set(val *IpSmGwAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIpSmGwAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIpSmGwAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpSmGwAddress(val *IpSmGwAddress) *NullableIpSmGwAddress {
	return &NullableIpSmGwAddress{value: val, isSet: true}
}

func (v NullableIpSmGwAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpSmGwAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
