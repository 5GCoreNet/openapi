/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
)

// checks if the TransferPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferPolicy{}

// TransferPolicy Describes a transfer policy.
type TransferPolicy struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxBitRateDl *string `json:"maxBitRateDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxBitRateUl *string `json:"maxBitRateUl,omitempty"`
	// Indicates a rating group for the recommended time window.
	RatingGroup int32      `json:"ratingGroup"`
	RecTimeInt  TimeWindow `json:"recTimeInt"`
	// Contains an identity of a transfer policy.
	TransPolicyId int32 `json:"transPolicyId"`
}

// NewTransferPolicy instantiates a new TransferPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferPolicy(ratingGroup int32, recTimeInt TimeWindow, transPolicyId int32) *TransferPolicy {
	this := TransferPolicy{}
	this.RatingGroup = ratingGroup
	this.RecTimeInt = recTimeInt
	this.TransPolicyId = transPolicyId
	return &this
}

// NewTransferPolicyWithDefaults instantiates a new TransferPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferPolicyWithDefaults() *TransferPolicy {
	this := TransferPolicy{}
	return &this
}

// GetMaxBitRateDl returns the MaxBitRateDl field value if set, zero value otherwise.
func (o *TransferPolicy) GetMaxBitRateDl() string {
	if o == nil || IsNil(o.MaxBitRateDl) {
		var ret string
		return ret
	}
	return *o.MaxBitRateDl
}

// GetMaxBitRateDlOk returns a tuple with the MaxBitRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetMaxBitRateDlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBitRateDl) {
		return nil, false
	}
	return o.MaxBitRateDl, true
}

// HasMaxBitRateDl returns a boolean if a field has been set.
func (o *TransferPolicy) HasMaxBitRateDl() bool {
	if o != nil && !IsNil(o.MaxBitRateDl) {
		return true
	}

	return false
}

// SetMaxBitRateDl gets a reference to the given string and assigns it to the MaxBitRateDl field.
func (o *TransferPolicy) SetMaxBitRateDl(v string) {
	o.MaxBitRateDl = &v
}

// GetMaxBitRateUl returns the MaxBitRateUl field value if set, zero value otherwise.
func (o *TransferPolicy) GetMaxBitRateUl() string {
	if o == nil || IsNil(o.MaxBitRateUl) {
		var ret string
		return ret
	}
	return *o.MaxBitRateUl
}

// GetMaxBitRateUlOk returns a tuple with the MaxBitRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetMaxBitRateUlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBitRateUl) {
		return nil, false
	}
	return o.MaxBitRateUl, true
}

// HasMaxBitRateUl returns a boolean if a field has been set.
func (o *TransferPolicy) HasMaxBitRateUl() bool {
	if o != nil && !IsNil(o.MaxBitRateUl) {
		return true
	}

	return false
}

// SetMaxBitRateUl gets a reference to the given string and assigns it to the MaxBitRateUl field.
func (o *TransferPolicy) SetMaxBitRateUl(v string) {
	o.MaxBitRateUl = &v
}

// GetRatingGroup returns the RatingGroup field value
func (o *TransferPolicy) GetRatingGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RatingGroup
}

// GetRatingGroupOk returns a tuple with the RatingGroup field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetRatingGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingGroup, true
}

// SetRatingGroup sets field value
func (o *TransferPolicy) SetRatingGroup(v int32) {
	o.RatingGroup = v
}

// GetRecTimeInt returns the RecTimeInt field value
func (o *TransferPolicy) GetRecTimeInt() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.RecTimeInt
}

// GetRecTimeIntOk returns a tuple with the RecTimeInt field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetRecTimeIntOk() (*TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecTimeInt, true
}

// SetRecTimeInt sets field value
func (o *TransferPolicy) SetRecTimeInt(v TimeWindow) {
	o.RecTimeInt = v
}

// GetTransPolicyId returns the TransPolicyId field value
func (o *TransferPolicy) GetTransPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransPolicyId
}

// GetTransPolicyIdOk returns a tuple with the TransPolicyId field value
// and a boolean to check if the value has been set.
func (o *TransferPolicy) GetTransPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransPolicyId, true
}

// SetTransPolicyId sets field value
func (o *TransferPolicy) SetTransPolicyId(v int32) {
	o.TransPolicyId = v
}

func (o TransferPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxBitRateDl) {
		toSerialize["maxBitRateDl"] = o.MaxBitRateDl
	}
	if !IsNil(o.MaxBitRateUl) {
		toSerialize["maxBitRateUl"] = o.MaxBitRateUl
	}
	toSerialize["ratingGroup"] = o.RatingGroup
	toSerialize["recTimeInt"] = o.RecTimeInt
	toSerialize["transPolicyId"] = o.TransPolicyId
	return toSerialize, nil
}

type NullableTransferPolicy struct {
	value *TransferPolicy
	isSet bool
}

func (v NullableTransferPolicy) Get() *TransferPolicy {
	return v.value
}

func (v *NullableTransferPolicy) Set(val *TransferPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferPolicy(val *TransferPolicy) *NullableTransferPolicy {
	return &NullableTransferPolicy{value: val, isSet: true}
}

func (v NullableTransferPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
