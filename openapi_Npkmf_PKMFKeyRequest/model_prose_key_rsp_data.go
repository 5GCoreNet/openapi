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

// checks if the ProseKeyRspData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseKeyRspData{}

// ProseKeyRspData Representation of the successfully requested keying material.
type ProseKeyRspData struct {
	// Key for NR PC5
	Knrp string `json:"knrp"`
	// KNRP Freshness Parameter 2
	KnrpFreshness2 string `json:"knrpFreshness2"`
	// GBA Pushing Information
	Gpi *string `json:"gpi,omitempty"`
}

// NewProseKeyRspData instantiates a new ProseKeyRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseKeyRspData(knrp string, knrpFreshness2 string) *ProseKeyRspData {
	this := ProseKeyRspData{}
	this.Knrp = knrp
	this.KnrpFreshness2 = knrpFreshness2
	return &this
}

// NewProseKeyRspDataWithDefaults instantiates a new ProseKeyRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseKeyRspDataWithDefaults() *ProseKeyRspData {
	this := ProseKeyRspData{}
	return &this
}

// GetKnrp returns the Knrp field value
func (o *ProseKeyRspData) GetKnrp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Knrp
}

// GetKnrpOk returns a tuple with the Knrp field value
// and a boolean to check if the value has been set.
func (o *ProseKeyRspData) GetKnrpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Knrp, true
}

// SetKnrp sets field value
func (o *ProseKeyRspData) SetKnrp(v string) {
	o.Knrp = v
}

// GetKnrpFreshness2 returns the KnrpFreshness2 field value
func (o *ProseKeyRspData) GetKnrpFreshness2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KnrpFreshness2
}

// GetKnrpFreshness2Ok returns a tuple with the KnrpFreshness2 field value
// and a boolean to check if the value has been set.
func (o *ProseKeyRspData) GetKnrpFreshness2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KnrpFreshness2, true
}

// SetKnrpFreshness2 sets field value
func (o *ProseKeyRspData) SetKnrpFreshness2(v string) {
	o.KnrpFreshness2 = v
}

// GetGpi returns the Gpi field value if set, zero value otherwise.
func (o *ProseKeyRspData) GetGpi() string {
	if o == nil || IsNil(o.Gpi) {
		var ret string
		return ret
	}
	return *o.Gpi
}

// GetGpiOk returns a tuple with the Gpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseKeyRspData) GetGpiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpi) {
		return nil, false
	}
	return o.Gpi, true
}

// HasGpi returns a boolean if a field has been set.
func (o *ProseKeyRspData) HasGpi() bool {
	if o != nil && !IsNil(o.Gpi) {
		return true
	}

	return false
}

// SetGpi gets a reference to the given string and assigns it to the Gpi field.
func (o *ProseKeyRspData) SetGpi(v string) {
	o.Gpi = &v
}

func (o ProseKeyRspData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseKeyRspData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["knrp"] = o.Knrp
	toSerialize["knrpFreshness2"] = o.KnrpFreshness2
	if !IsNil(o.Gpi) {
		toSerialize["gpi"] = o.Gpi
	}
	return toSerialize, nil
}

type NullableProseKeyRspData struct {
	value *ProseKeyRspData
	isSet bool
}

func (v NullableProseKeyRspData) Get() *ProseKeyRspData {
	return v.value
}

func (v *NullableProseKeyRspData) Set(val *ProseKeyRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableProseKeyRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableProseKeyRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseKeyRspData(val *ProseKeyRspData) *NullableProseKeyRspData {
	return &NullableProseKeyRspData{value: val, isSet: true}
}

func (v NullableProseKeyRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseKeyRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
