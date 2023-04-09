/*
Nhss_imsUEAU

Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUEAU

import (
	"encoding/json"
)

// checks if the DigestAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigestAuthentication{}

// DigestAuthentication Parameters used for the Digest authentication scheme
type DigestAuthentication struct {
	DigestRealm     string             `json:"digestRealm"`
	DigestAlgorithm SipDigestAlgorithm `json:"digestAlgorithm"`
	DigestQop       SipDigestQop       `json:"digestQop"`
	Ha1             string             `json:"ha1"`
}

// NewDigestAuthentication instantiates a new DigestAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigestAuthentication(digestRealm string, digestAlgorithm SipDigestAlgorithm, digestQop SipDigestQop, ha1 string) *DigestAuthentication {
	this := DigestAuthentication{}
	this.DigestRealm = digestRealm
	this.DigestAlgorithm = digestAlgorithm
	this.DigestQop = digestQop
	this.Ha1 = ha1
	return &this
}

// NewDigestAuthenticationWithDefaults instantiates a new DigestAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigestAuthenticationWithDefaults() *DigestAuthentication {
	this := DigestAuthentication{}
	return &this
}

// GetDigestRealm returns the DigestRealm field value
func (o *DigestAuthentication) GetDigestRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DigestRealm
}

// GetDigestRealmOk returns a tuple with the DigestRealm field value
// and a boolean to check if the value has been set.
func (o *DigestAuthentication) GetDigestRealmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigestRealm, true
}

// SetDigestRealm sets field value
func (o *DigestAuthentication) SetDigestRealm(v string) {
	o.DigestRealm = v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value
func (o *DigestAuthentication) GetDigestAlgorithm() SipDigestAlgorithm {
	if o == nil {
		var ret SipDigestAlgorithm
		return ret
	}

	return o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value
// and a boolean to check if the value has been set.
func (o *DigestAuthentication) GetDigestAlgorithmOk() (*SipDigestAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigestAlgorithm, true
}

// SetDigestAlgorithm sets field value
func (o *DigestAuthentication) SetDigestAlgorithm(v SipDigestAlgorithm) {
	o.DigestAlgorithm = v
}

// GetDigestQop returns the DigestQop field value
func (o *DigestAuthentication) GetDigestQop() SipDigestQop {
	if o == nil {
		var ret SipDigestQop
		return ret
	}

	return o.DigestQop
}

// GetDigestQopOk returns a tuple with the DigestQop field value
// and a boolean to check if the value has been set.
func (o *DigestAuthentication) GetDigestQopOk() (*SipDigestQop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigestQop, true
}

// SetDigestQop sets field value
func (o *DigestAuthentication) SetDigestQop(v SipDigestQop) {
	o.DigestQop = v
}

// GetHa1 returns the Ha1 field value
func (o *DigestAuthentication) GetHa1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ha1
}

// GetHa1Ok returns a tuple with the Ha1 field value
// and a boolean to check if the value has been set.
func (o *DigestAuthentication) GetHa1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ha1, true
}

// SetHa1 sets field value
func (o *DigestAuthentication) SetHa1(v string) {
	o.Ha1 = v
}

func (o DigestAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigestAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digestRealm"] = o.DigestRealm
	toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	toSerialize["digestQop"] = o.DigestQop
	toSerialize["ha1"] = o.Ha1
	return toSerialize, nil
}

type NullableDigestAuthentication struct {
	value *DigestAuthentication
	isSet bool
}

func (v NullableDigestAuthentication) Get() *DigestAuthentication {
	return v.value
}

func (v *NullableDigestAuthentication) Set(val *DigestAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableDigestAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableDigestAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigestAuthentication(val *DigestAuthentication) *NullableDigestAuthentication {
	return &NullableDigestAuthentication{value: val, isSet: true}
}

func (v NullableDigestAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigestAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
