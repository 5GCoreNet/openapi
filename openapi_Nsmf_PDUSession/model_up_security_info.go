/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the UpSecurityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpSecurityInfo{}

// UpSecurityInfo User Plane Security Information
type UpSecurityInfo struct {
	UpSecurity UpSecurity `json:"upSecurity"`
	MaxIntegrityProtectedDataRateUl *MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl *MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	SecurityResult *SecurityResult `json:"securityResult,omitempty"`
}

// NewUpSecurityInfo instantiates a new UpSecurityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpSecurityInfo(upSecurity UpSecurity) *UpSecurityInfo {
	this := UpSecurityInfo{}
	this.UpSecurity = upSecurity
	return &this
}

// NewUpSecurityInfoWithDefaults instantiates a new UpSecurityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpSecurityInfoWithDefaults() *UpSecurityInfo {
	this := UpSecurityInfo{}
	return &this
}

// GetUpSecurity returns the UpSecurity field value
func (o *UpSecurityInfo) GetUpSecurity() UpSecurity {
	if o == nil {
		var ret UpSecurity
		return ret
	}

	return o.UpSecurity
}

// GetUpSecurityOk returns a tuple with the UpSecurity field value
// and a boolean to check if the value has been set.
func (o *UpSecurityInfo) GetUpSecurityOk() (*UpSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpSecurity, true
}

// SetUpSecurity sets field value
func (o *UpSecurityInfo) SetUpSecurity(v UpSecurity) {
	o.UpSecurity = v
}

// GetMaxIntegrityProtectedDataRateUl returns the MaxIntegrityProtectedDataRateUl field value if set, zero value otherwise.
func (o *UpSecurityInfo) GetMaxIntegrityProtectedDataRateUl() MaxIntegrityProtectedDataRate {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateUl) {
		var ret MaxIntegrityProtectedDataRate
		return ret
	}
	return *o.MaxIntegrityProtectedDataRateUl
}

// GetMaxIntegrityProtectedDataRateUlOk returns a tuple with the MaxIntegrityProtectedDataRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpSecurityInfo) GetMaxIntegrityProtectedDataRateUlOk() (*MaxIntegrityProtectedDataRate, bool) {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateUl) {
		return nil, false
	}
	return o.MaxIntegrityProtectedDataRateUl, true
}

// HasMaxIntegrityProtectedDataRateUl returns a boolean if a field has been set.
func (o *UpSecurityInfo) HasMaxIntegrityProtectedDataRateUl() bool {
	if o != nil && !IsNil(o.MaxIntegrityProtectedDataRateUl) {
		return true
	}

	return false
}

// SetMaxIntegrityProtectedDataRateUl gets a reference to the given MaxIntegrityProtectedDataRate and assigns it to the MaxIntegrityProtectedDataRateUl field.
func (o *UpSecurityInfo) SetMaxIntegrityProtectedDataRateUl(v MaxIntegrityProtectedDataRate) {
	o.MaxIntegrityProtectedDataRateUl = &v
}

// GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field value if set, zero value otherwise.
func (o *UpSecurityInfo) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateDl) {
		var ret MaxIntegrityProtectedDataRate
		return ret
	}
	return *o.MaxIntegrityProtectedDataRateDl
}

// GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpSecurityInfo) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool) {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateDl) {
		return nil, false
	}
	return o.MaxIntegrityProtectedDataRateDl, true
}

// HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.
func (o *UpSecurityInfo) HasMaxIntegrityProtectedDataRateDl() bool {
	if o != nil && !IsNil(o.MaxIntegrityProtectedDataRateDl) {
		return true
	}

	return false
}

// SetMaxIntegrityProtectedDataRateDl gets a reference to the given MaxIntegrityProtectedDataRate and assigns it to the MaxIntegrityProtectedDataRateDl field.
func (o *UpSecurityInfo) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate) {
	o.MaxIntegrityProtectedDataRateDl = &v
}

// GetSecurityResult returns the SecurityResult field value if set, zero value otherwise.
func (o *UpSecurityInfo) GetSecurityResult() SecurityResult {
	if o == nil || IsNil(o.SecurityResult) {
		var ret SecurityResult
		return ret
	}
	return *o.SecurityResult
}

// GetSecurityResultOk returns a tuple with the SecurityResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpSecurityInfo) GetSecurityResultOk() (*SecurityResult, bool) {
	if o == nil || IsNil(o.SecurityResult) {
		return nil, false
	}
	return o.SecurityResult, true
}

// HasSecurityResult returns a boolean if a field has been set.
func (o *UpSecurityInfo) HasSecurityResult() bool {
	if o != nil && !IsNil(o.SecurityResult) {
		return true
	}

	return false
}

// SetSecurityResult gets a reference to the given SecurityResult and assigns it to the SecurityResult field.
func (o *UpSecurityInfo) SetSecurityResult(v SecurityResult) {
	o.SecurityResult = &v
}

func (o UpSecurityInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpSecurityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upSecurity"] = o.UpSecurity
	if !IsNil(o.MaxIntegrityProtectedDataRateUl) {
		toSerialize["maxIntegrityProtectedDataRateUl"] = o.MaxIntegrityProtectedDataRateUl
	}
	if !IsNil(o.MaxIntegrityProtectedDataRateDl) {
		toSerialize["maxIntegrityProtectedDataRateDl"] = o.MaxIntegrityProtectedDataRateDl
	}
	if !IsNil(o.SecurityResult) {
		toSerialize["securityResult"] = o.SecurityResult
	}
	return toSerialize, nil
}

type NullableUpSecurityInfo struct {
	value *UpSecurityInfo
	isSet bool
}

func (v NullableUpSecurityInfo) Get() *UpSecurityInfo {
	return v.value
}

func (v *NullableUpSecurityInfo) Set(val *UpSecurityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpSecurityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpSecurityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpSecurityInfo(val *UpSecurityInfo) *NullableUpSecurityInfo {
	return &NullableUpSecurityInfo{value: val, isSet: true}
}

func (v NullableUpSecurityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpSecurityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

