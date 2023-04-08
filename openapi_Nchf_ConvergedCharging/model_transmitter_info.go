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

// checks if the TransmitterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransmitterInfo{}

// TransmitterInfo struct for TransmitterInfo
type TransmitterInfo struct {
	ProseSourceIPAddress *IpAddr `json:"proseSourceIPAddress,omitempty"`
	ProseSourceL2Id *string `json:"proseSourceL2Id,omitempty"`
}

// NewTransmitterInfo instantiates a new TransmitterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransmitterInfo() *TransmitterInfo {
	this := TransmitterInfo{}
	return &this
}

// NewTransmitterInfoWithDefaults instantiates a new TransmitterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransmitterInfoWithDefaults() *TransmitterInfo {
	this := TransmitterInfo{}
	return &this
}

// GetProseSourceIPAddress returns the ProseSourceIPAddress field value if set, zero value otherwise.
func (o *TransmitterInfo) GetProseSourceIPAddress() IpAddr {
	if o == nil || isNil(o.ProseSourceIPAddress) {
		var ret IpAddr
		return ret
	}
	return *o.ProseSourceIPAddress
}

// GetProseSourceIPAddressOk returns a tuple with the ProseSourceIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransmitterInfo) GetProseSourceIPAddressOk() (*IpAddr, bool) {
	if o == nil || isNil(o.ProseSourceIPAddress) {
		return nil, false
	}
	return o.ProseSourceIPAddress, true
}

// HasProseSourceIPAddress returns a boolean if a field has been set.
func (o *TransmitterInfo) HasProseSourceIPAddress() bool {
	if o != nil && !isNil(o.ProseSourceIPAddress) {
		return true
	}

	return false
}

// SetProseSourceIPAddress gets a reference to the given IpAddr and assigns it to the ProseSourceIPAddress field.
func (o *TransmitterInfo) SetProseSourceIPAddress(v IpAddr) {
	o.ProseSourceIPAddress = &v
}

// GetProseSourceL2Id returns the ProseSourceL2Id field value if set, zero value otherwise.
func (o *TransmitterInfo) GetProseSourceL2Id() string {
	if o == nil || isNil(o.ProseSourceL2Id) {
		var ret string
		return ret
	}
	return *o.ProseSourceL2Id
}

// GetProseSourceL2IdOk returns a tuple with the ProseSourceL2Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransmitterInfo) GetProseSourceL2IdOk() (*string, bool) {
	if o == nil || isNil(o.ProseSourceL2Id) {
		return nil, false
	}
	return o.ProseSourceL2Id, true
}

// HasProseSourceL2Id returns a boolean if a field has been set.
func (o *TransmitterInfo) HasProseSourceL2Id() bool {
	if o != nil && !isNil(o.ProseSourceL2Id) {
		return true
	}

	return false
}

// SetProseSourceL2Id gets a reference to the given string and assigns it to the ProseSourceL2Id field.
func (o *TransmitterInfo) SetProseSourceL2Id(v string) {
	o.ProseSourceL2Id = &v
}

func (o TransmitterInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransmitterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProseSourceIPAddress) {
		toSerialize["proseSourceIPAddress"] = o.ProseSourceIPAddress
	}
	if !isNil(o.ProseSourceL2Id) {
		toSerialize["proseSourceL2Id"] = o.ProseSourceL2Id
	}
	return toSerialize, nil
}

type NullableTransmitterInfo struct {
	value *TransmitterInfo
	isSet bool
}

func (v NullableTransmitterInfo) Get() *TransmitterInfo {
	return v.value
}

func (v *NullableTransmitterInfo) Set(val *TransmitterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransmitterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransmitterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransmitterInfo(val *TransmitterInfo) *NullableTransmitterInfo {
	return &NullableTransmitterInfo{value: val, isSet: true}
}

func (v NullableTransmitterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransmitterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


