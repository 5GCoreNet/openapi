/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// checks if the DataToIntegrityProtectAndCipherBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataToIntegrityProtectAndCipherBlock{}

// DataToIntegrityProtectAndCipherBlock HTTP header to be encrypted or the value of a JSON attribute to be encrypted
type DataToIntegrityProtectAndCipherBlock struct {
	DataToEncrypt []interface{} `json:"dataToEncrypt"`
}

// NewDataToIntegrityProtectAndCipherBlock instantiates a new DataToIntegrityProtectAndCipherBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataToIntegrityProtectAndCipherBlock(dataToEncrypt []interface{}) *DataToIntegrityProtectAndCipherBlock {
	this := DataToIntegrityProtectAndCipherBlock{}
	this.DataToEncrypt = dataToEncrypt
	return &this
}

// NewDataToIntegrityProtectAndCipherBlockWithDefaults instantiates a new DataToIntegrityProtectAndCipherBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataToIntegrityProtectAndCipherBlockWithDefaults() *DataToIntegrityProtectAndCipherBlock {
	this := DataToIntegrityProtectAndCipherBlock{}
	return &this
}

// GetDataToEncrypt returns the DataToEncrypt field value
func (o *DataToIntegrityProtectAndCipherBlock) GetDataToEncrypt() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.DataToEncrypt
}

// GetDataToEncryptOk returns a tuple with the DataToEncrypt field value
// and a boolean to check if the value has been set.
func (o *DataToIntegrityProtectAndCipherBlock) GetDataToEncryptOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataToEncrypt, true
}

// SetDataToEncrypt sets field value
func (o *DataToIntegrityProtectAndCipherBlock) SetDataToEncrypt(v []interface{}) {
	o.DataToEncrypt = v
}

func (o DataToIntegrityProtectAndCipherBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataToIntegrityProtectAndCipherBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataToEncrypt"] = o.DataToEncrypt
	return toSerialize, nil
}

type NullableDataToIntegrityProtectAndCipherBlock struct {
	value *DataToIntegrityProtectAndCipherBlock
	isSet bool
}

func (v NullableDataToIntegrityProtectAndCipherBlock) Get() *DataToIntegrityProtectAndCipherBlock {
	return v.value
}

func (v *NullableDataToIntegrityProtectAndCipherBlock) Set(val *DataToIntegrityProtectAndCipherBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableDataToIntegrityProtectAndCipherBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableDataToIntegrityProtectAndCipherBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataToIntegrityProtectAndCipherBlock(val *DataToIntegrityProtectAndCipherBlock) *NullableDataToIntegrityProtectAndCipherBlock {
	return &NullableDataToIntegrityProtectAndCipherBlock{value: val, isSet: true}
}

func (v NullableDataToIntegrityProtectAndCipherBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataToIntegrityProtectAndCipherBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
