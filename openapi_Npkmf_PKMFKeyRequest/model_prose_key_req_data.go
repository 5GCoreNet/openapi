/*
Npkmf_PKMFKeyRequest

PKMF KeyRequest Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npkmf_PKMFKeyRequest

import (
	"encoding/json"
)

// checks if the ProseKeyReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseKeyReqData{}

// ProseKeyReqData Representation of the input to request the keying material.
type ProseKeyReqData struct {
	// Relay Service Code to identify a connectivity service provided by the UE-to-Network relay. 
	RelayServCode int32 `json:"relayServCode"`
	// KNRP Freshness Parameter 1
	KnrpFreshness1 string `json:"knrpFreshness1"`
	ResyncInfo *ResynchronizationInfo `json:"resyncInfo,omitempty"`
	// User Plane Prose Remote User Key ID
	PrukId *string `json:"prukId,omitempty"`
	// Contains the SUCI.
	Suci *string `json:"suci,omitempty"`
}

// NewProseKeyReqData instantiates a new ProseKeyReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseKeyReqData(relayServCode int32, knrpFreshness1 string) *ProseKeyReqData {
	this := ProseKeyReqData{}
	this.RelayServCode = relayServCode
	this.KnrpFreshness1 = knrpFreshness1
	return &this
}

// NewProseKeyReqDataWithDefaults instantiates a new ProseKeyReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseKeyReqDataWithDefaults() *ProseKeyReqData {
	this := ProseKeyReqData{}
	return &this
}

// GetRelayServCode returns the RelayServCode field value
func (o *ProseKeyReqData) GetRelayServCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelayServCode
}

// GetRelayServCodeOk returns a tuple with the RelayServCode field value
// and a boolean to check if the value has been set.
func (o *ProseKeyReqData) GetRelayServCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelayServCode, true
}

// SetRelayServCode sets field value
func (o *ProseKeyReqData) SetRelayServCode(v int32) {
	o.RelayServCode = v
}

// GetKnrpFreshness1 returns the KnrpFreshness1 field value
func (o *ProseKeyReqData) GetKnrpFreshness1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KnrpFreshness1
}

// GetKnrpFreshness1Ok returns a tuple with the KnrpFreshness1 field value
// and a boolean to check if the value has been set.
func (o *ProseKeyReqData) GetKnrpFreshness1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnrpFreshness1, true
}

// SetKnrpFreshness1 sets field value
func (o *ProseKeyReqData) SetKnrpFreshness1(v string) {
	o.KnrpFreshness1 = v
}

// GetResyncInfo returns the ResyncInfo field value if set, zero value otherwise.
func (o *ProseKeyReqData) GetResyncInfo() ResynchronizationInfo {
	if o == nil || IsNil(o.ResyncInfo) {
		var ret ResynchronizationInfo
		return ret
	}
	return *o.ResyncInfo
}

// GetResyncInfoOk returns a tuple with the ResyncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseKeyReqData) GetResyncInfoOk() (*ResynchronizationInfo, bool) {
	if o == nil || IsNil(o.ResyncInfo) {
		return nil, false
	}
	return o.ResyncInfo, true
}

// HasResyncInfo returns a boolean if a field has been set.
func (o *ProseKeyReqData) HasResyncInfo() bool {
	if o != nil && !IsNil(o.ResyncInfo) {
		return true
	}

	return false
}

// SetResyncInfo gets a reference to the given ResynchronizationInfo and assigns it to the ResyncInfo field.
func (o *ProseKeyReqData) SetResyncInfo(v ResynchronizationInfo) {
	o.ResyncInfo = &v
}

// GetPrukId returns the PrukId field value if set, zero value otherwise.
func (o *ProseKeyReqData) GetPrukId() string {
	if o == nil || IsNil(o.PrukId) {
		var ret string
		return ret
	}
	return *o.PrukId
}

// GetPrukIdOk returns a tuple with the PrukId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseKeyReqData) GetPrukIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrukId) {
		return nil, false
	}
	return o.PrukId, true
}

// HasPrukId returns a boolean if a field has been set.
func (o *ProseKeyReqData) HasPrukId() bool {
	if o != nil && !IsNil(o.PrukId) {
		return true
	}

	return false
}

// SetPrukId gets a reference to the given string and assigns it to the PrukId field.
func (o *ProseKeyReqData) SetPrukId(v string) {
	o.PrukId = &v
}

// GetSuci returns the Suci field value if set, zero value otherwise.
func (o *ProseKeyReqData) GetSuci() string {
	if o == nil || IsNil(o.Suci) {
		var ret string
		return ret
	}
	return *o.Suci
}

// GetSuciOk returns a tuple with the Suci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseKeyReqData) GetSuciOk() (*string, bool) {
	if o == nil || IsNil(o.Suci) {
		return nil, false
	}
	return o.Suci, true
}

// HasSuci returns a boolean if a field has been set.
func (o *ProseKeyReqData) HasSuci() bool {
	if o != nil && !IsNil(o.Suci) {
		return true
	}

	return false
}

// SetSuci gets a reference to the given string and assigns it to the Suci field.
func (o *ProseKeyReqData) SetSuci(v string) {
	o.Suci = &v
}

func (o ProseKeyReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseKeyReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relayServCode"] = o.RelayServCode
	toSerialize["knrpFreshness1"] = o.KnrpFreshness1
	if !IsNil(o.ResyncInfo) {
		toSerialize["resyncInfo"] = o.ResyncInfo
	}
	if !IsNil(o.PrukId) {
		toSerialize["prukId"] = o.PrukId
	}
	if !IsNil(o.Suci) {
		toSerialize["suci"] = o.Suci
	}
	return toSerialize, nil
}

type NullableProseKeyReqData struct {
	value *ProseKeyReqData
	isSet bool
}

func (v NullableProseKeyReqData) Get() *ProseKeyReqData {
	return v.value
}

func (v *NullableProseKeyReqData) Set(val *ProseKeyReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableProseKeyReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableProseKeyReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseKeyReqData(val *ProseKeyReqData) *NullableProseKeyReqData {
	return &NullableProseKeyReqData{value: val, isSet: true}
}

func (v NullableProseKeyReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseKeyReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


