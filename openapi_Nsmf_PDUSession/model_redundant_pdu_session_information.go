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

// checks if the RedundantPduSessionInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundantPduSessionInformation{}

// RedundantPduSessionInformation Redundant PDU Session Information
type RedundantPduSessionInformation struct {
	Rsn Rsn `json:"rsn"`
	PduSessionPairId *int32 `json:"pduSessionPairId,omitempty"`
}

// NewRedundantPduSessionInformation instantiates a new RedundantPduSessionInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundantPduSessionInformation(rsn Rsn) *RedundantPduSessionInformation {
	this := RedundantPduSessionInformation{}
	this.Rsn = rsn
	return &this
}

// NewRedundantPduSessionInformationWithDefaults instantiates a new RedundantPduSessionInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundantPduSessionInformationWithDefaults() *RedundantPduSessionInformation {
	this := RedundantPduSessionInformation{}
	return &this
}

// GetRsn returns the Rsn field value
func (o *RedundantPduSessionInformation) GetRsn() Rsn {
	if o == nil {
		var ret Rsn
		return ret
	}

	return o.Rsn
}

// GetRsnOk returns a tuple with the Rsn field value
// and a boolean to check if the value has been set.
func (o *RedundantPduSessionInformation) GetRsnOk() (*Rsn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rsn, true
}

// SetRsn sets field value
func (o *RedundantPduSessionInformation) SetRsn(v Rsn) {
	o.Rsn = v
}

// GetPduSessionPairId returns the PduSessionPairId field value if set, zero value otherwise.
func (o *RedundantPduSessionInformation) GetPduSessionPairId() int32 {
	if o == nil || IsNil(o.PduSessionPairId) {
		var ret int32
		return ret
	}
	return *o.PduSessionPairId
}

// GetPduSessionPairIdOk returns a tuple with the PduSessionPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantPduSessionInformation) GetPduSessionPairIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSessionPairId) {
		return nil, false
	}
	return o.PduSessionPairId, true
}

// HasPduSessionPairId returns a boolean if a field has been set.
func (o *RedundantPduSessionInformation) HasPduSessionPairId() bool {
	if o != nil && !IsNil(o.PduSessionPairId) {
		return true
	}

	return false
}

// SetPduSessionPairId gets a reference to the given int32 and assigns it to the PduSessionPairId field.
func (o *RedundantPduSessionInformation) SetPduSessionPairId(v int32) {
	o.PduSessionPairId = &v
}

func (o RedundantPduSessionInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundantPduSessionInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rsn"] = o.Rsn
	if !IsNil(o.PduSessionPairId) {
		toSerialize["pduSessionPairId"] = o.PduSessionPairId
	}
	return toSerialize, nil
}

type NullableRedundantPduSessionInformation struct {
	value *RedundantPduSessionInformation
	isSet bool
}

func (v NullableRedundantPduSessionInformation) Get() *RedundantPduSessionInformation {
	return v.value
}

func (v *NullableRedundantPduSessionInformation) Set(val *RedundantPduSessionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantPduSessionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantPduSessionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantPduSessionInformation(val *RedundantPduSessionInformation) *NullableRedundantPduSessionInformation {
	return &NullableRedundantPduSessionInformation{value: val, isSet: true}
}

func (v NullableRedundantPduSessionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantPduSessionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


