/*
Niwmsc_SMService

SMS-IWMSC Short Message Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Niwmsc_SMService

import (
	"encoding/json"
	"os"
)

// checks if the SendSMS200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendSMS200Response{}

// SendSMS200Response struct for SendSMS200Response
type SendSMS200Response struct {
	JsonData      *SmsDeliveryData `json:"jsonData,omitempty"`
	BinaryPayload **os.File        `json:"binaryPayload,omitempty"`
}

// NewSendSMS200Response instantiates a new SendSMS200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendSMS200Response() *SendSMS200Response {
	this := SendSMS200Response{}
	return &this
}

// NewSendSMS200ResponseWithDefaults instantiates a new SendSMS200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendSMS200ResponseWithDefaults() *SendSMS200Response {
	this := SendSMS200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *SendSMS200Response) GetJsonData() SmsDeliveryData {
	if o == nil || IsNil(o.JsonData) {
		var ret SmsDeliveryData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSMS200Response) GetJsonDataOk() (*SmsDeliveryData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *SendSMS200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmsDeliveryData and assigns it to the JsonData field.
func (o *SendSMS200Response) SetJsonData(v SmsDeliveryData) {
	o.JsonData = &v
}

// GetBinaryPayload returns the BinaryPayload field value if set, zero value otherwise.
func (o *SendSMS200Response) GetBinaryPayload() *os.File {
	if o == nil || IsNil(o.BinaryPayload) {
		var ret *os.File
		return ret
	}
	return *o.BinaryPayload
}

// GetBinaryPayloadOk returns a tuple with the BinaryPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSMS200Response) GetBinaryPayloadOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryPayload) {
		return nil, false
	}
	return o.BinaryPayload, true
}

// HasBinaryPayload returns a boolean if a field has been set.
func (o *SendSMS200Response) HasBinaryPayload() bool {
	if o != nil && !IsNil(o.BinaryPayload) {
		return true
	}

	return false
}

// SetBinaryPayload gets a reference to the given *os.File and assigns it to the BinaryPayload field.
func (o *SendSMS200Response) SetBinaryPayload(v *os.File) {
	o.BinaryPayload = &v
}

func (o SendSMS200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendSMS200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryPayload) {
		toSerialize["binaryPayload"] = o.BinaryPayload
	}
	return toSerialize, nil
}

type NullableSendSMS200Response struct {
	value *SendSMS200Response
	isSet bool
}

func (v NullableSendSMS200Response) Get() *SendSMS200Response {
	return v.value
}

func (v *NullableSendSMS200Response) Set(val *SendSMS200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSendSMS200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSendSMS200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendSMS200Response(val *SendSMS200Response) *NullableSendSMS200Response {
	return &NullableSendSMS200Response{value: val, isSet: true}
}

func (v NullableSendSMS200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendSMS200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
