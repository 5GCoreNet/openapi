/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
	"time"
)

// checks if the AuthDataForRestricted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthDataForRestricted{}

// AuthDataForRestricted Represents obtained authorization Data for restricted discovery for a discoverer UE
type AuthDataForRestricted struct {
	ProseQueryCodes []string `json:"proseQueryCodes"`
	// Contains the ProSe Response Code.
	ProseRespCode string `json:"proseRespCode"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime"`
}

// NewAuthDataForRestricted instantiates a new AuthDataForRestricted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthDataForRestricted(proseQueryCodes []string, proseRespCode string, validityTime time.Time) *AuthDataForRestricted {
	this := AuthDataForRestricted{}
	this.ProseQueryCodes = proseQueryCodes
	this.ProseRespCode = proseRespCode
	this.ValidityTime = validityTime
	return &this
}

// NewAuthDataForRestrictedWithDefaults instantiates a new AuthDataForRestricted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDataForRestrictedWithDefaults() *AuthDataForRestricted {
	this := AuthDataForRestricted{}
	return &this
}

// GetProseQueryCodes returns the ProseQueryCodes field value
func (o *AuthDataForRestricted) GetProseQueryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProseQueryCodes
}

// GetProseQueryCodesOk returns a tuple with the ProseQueryCodes field value
// and a boolean to check if the value has been set.
func (o *AuthDataForRestricted) GetProseQueryCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProseQueryCodes, true
}

// SetProseQueryCodes sets field value
func (o *AuthDataForRestricted) SetProseQueryCodes(v []string) {
	o.ProseQueryCodes = v
}

// GetProseRespCode returns the ProseRespCode field value
func (o *AuthDataForRestricted) GetProseRespCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProseRespCode
}

// GetProseRespCodeOk returns a tuple with the ProseRespCode field value
// and a boolean to check if the value has been set.
func (o *AuthDataForRestricted) GetProseRespCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProseRespCode, true
}

// SetProseRespCode sets field value
func (o *AuthDataForRestricted) SetProseRespCode(v string) {
	o.ProseRespCode = v
}

// GetValidityTime returns the ValidityTime field value
func (o *AuthDataForRestricted) GetValidityTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value
// and a boolean to check if the value has been set.
func (o *AuthDataForRestricted) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidityTime, true
}

// SetValidityTime sets field value
func (o *AuthDataForRestricted) SetValidityTime(v time.Time) {
	o.ValidityTime = v
}

func (o AuthDataForRestricted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthDataForRestricted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["proseQueryCodes"] = o.ProseQueryCodes
	toSerialize["proseRespCode"] = o.ProseRespCode
	toSerialize["validityTime"] = o.ValidityTime
	return toSerialize, nil
}

type NullableAuthDataForRestricted struct {
	value *AuthDataForRestricted
	isSet bool
}

func (v NullableAuthDataForRestricted) Get() *AuthDataForRestricted {
	return v.value
}

func (v *NullableAuthDataForRestricted) Set(val *AuthDataForRestricted) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthDataForRestricted) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthDataForRestricted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthDataForRestricted(val *AuthDataForRestricted) *NullableAuthDataForRestricted {
	return &NullableAuthDataForRestricted{value: val, isSet: true}
}

func (v NullableAuthDataForRestricted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthDataForRestricted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


