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

// checks if the RedirectServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectServer{}

// RedirectServer struct for RedirectServer
type RedirectServer struct {
	RedirectAddressType RedirectAddressType `json:"redirectAddressType"`
	RedirectServerAddress string `json:"redirectServerAddress"`
}

// NewRedirectServer instantiates a new RedirectServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectServer(redirectAddressType RedirectAddressType, redirectServerAddress string) *RedirectServer {
	this := RedirectServer{}
	this.RedirectAddressType = redirectAddressType
	this.RedirectServerAddress = redirectServerAddress
	return &this
}

// NewRedirectServerWithDefaults instantiates a new RedirectServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectServerWithDefaults() *RedirectServer {
	this := RedirectServer{}
	return &this
}

// GetRedirectAddressType returns the RedirectAddressType field value
func (o *RedirectServer) GetRedirectAddressType() RedirectAddressType {
	if o == nil {
		var ret RedirectAddressType
		return ret
	}

	return o.RedirectAddressType
}

// GetRedirectAddressTypeOk returns a tuple with the RedirectAddressType field value
// and a boolean to check if the value has been set.
func (o *RedirectServer) GetRedirectAddressTypeOk() (*RedirectAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectAddressType, true
}

// SetRedirectAddressType sets field value
func (o *RedirectServer) SetRedirectAddressType(v RedirectAddressType) {
	o.RedirectAddressType = v
}

// GetRedirectServerAddress returns the RedirectServerAddress field value
func (o *RedirectServer) GetRedirectServerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectServerAddress
}

// GetRedirectServerAddressOk returns a tuple with the RedirectServerAddress field value
// and a boolean to check if the value has been set.
func (o *RedirectServer) GetRedirectServerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectServerAddress, true
}

// SetRedirectServerAddress sets field value
func (o *RedirectServer) SetRedirectServerAddress(v string) {
	o.RedirectServerAddress = v
}

func (o RedirectServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["redirectAddressType"] = o.RedirectAddressType
	toSerialize["redirectServerAddress"] = o.RedirectServerAddress
	return toSerialize, nil
}

type NullableRedirectServer struct {
	value *RedirectServer
	isSet bool
}

func (v NullableRedirectServer) Get() *RedirectServer {
	return v.value
}

func (v *NullableRedirectServer) Set(val *RedirectServer) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectServer) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectServer(val *RedirectServer) *NullableRedirectServer {
	return &NullableRedirectServer{value: val, isSet: true}
}

func (v NullableRedirectServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


