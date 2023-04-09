/*
3gpp-traffic-influence

API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TrafficInfluence

import (
	"encoding/json"
)

// checks if the AfAckInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfAckInfo{}

// AfAckInfo Represents acknowledgement information of a traffic influence event notification.
type AfAckInfo struct {
	AfTransId *string      `json:"afTransId,omitempty"`
	AckResult AfResultInfo `json:"ackResult"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string `json:"gpsi,omitempty"`
}

// NewAfAckInfo instantiates a new AfAckInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfAckInfo(ackResult AfResultInfo) *AfAckInfo {
	this := AfAckInfo{}
	this.AckResult = ackResult
	return &this
}

// NewAfAckInfoWithDefaults instantiates a new AfAckInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfAckInfoWithDefaults() *AfAckInfo {
	this := AfAckInfo{}
	return &this
}

// GetAfTransId returns the AfTransId field value if set, zero value otherwise.
func (o *AfAckInfo) GetAfTransId() string {
	if o == nil || IsNil(o.AfTransId) {
		var ret string
		return ret
	}
	return *o.AfTransId
}

// GetAfTransIdOk returns a tuple with the AfTransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfAckInfo) GetAfTransIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfTransId) {
		return nil, false
	}
	return o.AfTransId, true
}

// HasAfTransId returns a boolean if a field has been set.
func (o *AfAckInfo) HasAfTransId() bool {
	if o != nil && !IsNil(o.AfTransId) {
		return true
	}

	return false
}

// SetAfTransId gets a reference to the given string and assigns it to the AfTransId field.
func (o *AfAckInfo) SetAfTransId(v string) {
	o.AfTransId = &v
}

// GetAckResult returns the AckResult field value
func (o *AfAckInfo) GetAckResult() AfResultInfo {
	if o == nil {
		var ret AfResultInfo
		return ret
	}

	return o.AckResult
}

// GetAckResultOk returns a tuple with the AckResult field value
// and a boolean to check if the value has been set.
func (o *AfAckInfo) GetAckResultOk() (*AfResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckResult, true
}

// SetAckResult sets field value
func (o *AfAckInfo) SetAckResult(v AfResultInfo) {
	o.AckResult = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *AfAckInfo) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfAckInfo) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *AfAckInfo) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *AfAckInfo) SetGpsi(v string) {
	o.Gpsi = &v
}

func (o AfAckInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfAckInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AfTransId) {
		toSerialize["afTransId"] = o.AfTransId
	}
	toSerialize["ackResult"] = o.AckResult
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	return toSerialize, nil
}

type NullableAfAckInfo struct {
	value *AfAckInfo
	isSet bool
}

func (v NullableAfAckInfo) Get() *AfAckInfo {
	return v.value
}

func (v *NullableAfAckInfo) Set(val *AfAckInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAfAckInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAfAckInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfAckInfo(val *AfAckInfo) *NullableAfAckInfo {
	return &NullableAfAckInfo{value: val, isSet: true}
}

func (v NullableAfAckInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfAckInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
