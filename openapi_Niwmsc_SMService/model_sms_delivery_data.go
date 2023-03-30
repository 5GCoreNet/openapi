/*
Niwmsc_SMService

SMS-IWMSC Short Message Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Niwmsc_SMService

import (
	"encoding/json"
)

// checks if the SmsDeliveryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsDeliveryData{}

// SmsDeliveryData Information within response message invoking MoForwardSm service operation, for delivering MO SMS Delivery Report. 
type SmsDeliveryData struct {
	SmsPayload RefToBinaryData `json:"smsPayload"`
}

// NewSmsDeliveryData instantiates a new SmsDeliveryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsDeliveryData(smsPayload RefToBinaryData) *SmsDeliveryData {
	this := SmsDeliveryData{}
	this.SmsPayload = smsPayload
	return &this
}

// NewSmsDeliveryDataWithDefaults instantiates a new SmsDeliveryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsDeliveryDataWithDefaults() *SmsDeliveryData {
	this := SmsDeliveryData{}
	return &this
}

// GetSmsPayload returns the SmsPayload field value
func (o *SmsDeliveryData) GetSmsPayload() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.SmsPayload
}

// GetSmsPayloadOk returns a tuple with the SmsPayload field value
// and a boolean to check if the value has been set.
func (o *SmsDeliveryData) GetSmsPayloadOk() (*RefToBinaryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmsPayload, true
}

// SetSmsPayload sets field value
func (o *SmsDeliveryData) SetSmsPayload(v RefToBinaryData) {
	o.SmsPayload = v
}

func (o SmsDeliveryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsDeliveryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smsPayload"] = o.SmsPayload
	return toSerialize, nil
}

type NullableSmsDeliveryData struct {
	value *SmsDeliveryData
	isSet bool
}

func (v NullableSmsDeliveryData) Get() *SmsDeliveryData {
	return v.value
}

func (v *NullableSmsDeliveryData) Set(val *SmsDeliveryData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsDeliveryData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsDeliveryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsDeliveryData(val *SmsDeliveryData) *NullableSmsDeliveryData {
	return &NullableSmsDeliveryData{value: val, isSet: true}
}

func (v NullableSmsDeliveryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsDeliveryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

