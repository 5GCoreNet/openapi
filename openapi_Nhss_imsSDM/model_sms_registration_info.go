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

// checks if the SmsRegistrationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsRegistrationInfo{}

// SmsRegistrationInfo SMS Registration Information (IP-SM-GW number and SC address)
type SmsRegistrationInfo struct {
	// String containing an additional or basic MSISDN
	IpSmGwNumber string `json:"ipSmGwNumber"`
	// String containing an additional or basic MSISDN
	ScAddress *string `json:"scAddress,omitempty"`
}

// NewSmsRegistrationInfo instantiates a new SmsRegistrationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsRegistrationInfo(ipSmGwNumber string) *SmsRegistrationInfo {
	this := SmsRegistrationInfo{}
	this.IpSmGwNumber = ipSmGwNumber
	return &this
}

// NewSmsRegistrationInfoWithDefaults instantiates a new SmsRegistrationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsRegistrationInfoWithDefaults() *SmsRegistrationInfo {
	this := SmsRegistrationInfo{}
	return &this
}

// GetIpSmGwNumber returns the IpSmGwNumber field value
func (o *SmsRegistrationInfo) GetIpSmGwNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpSmGwNumber
}

// GetIpSmGwNumberOk returns a tuple with the IpSmGwNumber field value
// and a boolean to check if the value has been set.
func (o *SmsRegistrationInfo) GetIpSmGwNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpSmGwNumber, true
}

// SetIpSmGwNumber sets field value
func (o *SmsRegistrationInfo) SetIpSmGwNumber(v string) {
	o.IpSmGwNumber = v
}

// GetScAddress returns the ScAddress field value if set, zero value otherwise.
func (o *SmsRegistrationInfo) GetScAddress() string {
	if o == nil || isNil(o.ScAddress) {
		var ret string
		return ret
	}
	return *o.ScAddress
}

// GetScAddressOk returns a tuple with the ScAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsRegistrationInfo) GetScAddressOk() (*string, bool) {
	if o == nil || isNil(o.ScAddress) {
		return nil, false
	}
	return o.ScAddress, true
}

// HasScAddress returns a boolean if a field has been set.
func (o *SmsRegistrationInfo) HasScAddress() bool {
	if o != nil && !isNil(o.ScAddress) {
		return true
	}

	return false
}

// SetScAddress gets a reference to the given string and assigns it to the ScAddress field.
func (o *SmsRegistrationInfo) SetScAddress(v string) {
	o.ScAddress = &v
}

func (o SmsRegistrationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsRegistrationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ipSmGwNumber"] = o.IpSmGwNumber
	if !isNil(o.ScAddress) {
		toSerialize["scAddress"] = o.ScAddress
	}
	return toSerialize, nil
}

type NullableSmsRegistrationInfo struct {
	value *SmsRegistrationInfo
	isSet bool
}

func (v NullableSmsRegistrationInfo) Get() *SmsRegistrationInfo {
	return v.value
}

func (v *NullableSmsRegistrationInfo) Set(val *SmsRegistrationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsRegistrationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsRegistrationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsRegistrationInfo(val *SmsRegistrationInfo) *NullableSmsRegistrationInfo {
	return &NullableSmsRegistrationInfo{value: val, isSet: true}
}

func (v NullableSmsRegistrationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsRegistrationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


