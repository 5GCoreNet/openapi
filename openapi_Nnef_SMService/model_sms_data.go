/*
Nnef_SMService

Nnef SMService Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_SMService

import (
	"encoding/json"
)

// checks if the SmsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsData{}

// SmsData Information within request message for delivering SMS.
type SmsData struct {
	SmsPayload RefToBinaryData `json:"smsPayload"`
}

// NewSmsData instantiates a new SmsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsData(smsPayload RefToBinaryData) *SmsData {
	this := SmsData{}
	this.SmsPayload = smsPayload
	return &this
}

// NewSmsDataWithDefaults instantiates a new SmsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsDataWithDefaults() *SmsData {
	this := SmsData{}
	return &this
}

// GetSmsPayload returns the SmsPayload field value
func (o *SmsData) GetSmsPayload() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.SmsPayload
}

// GetSmsPayloadOk returns a tuple with the SmsPayload field value
// and a boolean to check if the value has been set.
func (o *SmsData) GetSmsPayloadOk() (*RefToBinaryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmsPayload, true
}

// SetSmsPayload sets field value
func (o *SmsData) SetSmsPayload(v RefToBinaryData) {
	o.SmsPayload = v
}

func (o SmsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smsPayload"] = o.SmsPayload
	return toSerialize, nil
}

type NullableSmsData struct {
	value *SmsData
	isSet bool
}

func (v NullableSmsData) Get() *SmsData {
	return v.value
}

func (v *NullableSmsData) Set(val *SmsData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsData(val *SmsData) *NullableSmsData {
	return &NullableSmsData{value: val, isSet: true}
}

func (v NullableSmsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
