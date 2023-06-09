/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Broadcast

import (
	"encoding/json"
)

// checks if the CipherResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CipherResponseData{}

// CipherResponseData Information within Ciphering Key Data Response.
type CipherResponseData struct {
	DataAvailability DataAvailability `json:"dataAvailability"`
}

// NewCipherResponseData instantiates a new CipherResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCipherResponseData(dataAvailability DataAvailability) *CipherResponseData {
	this := CipherResponseData{}
	this.DataAvailability = dataAvailability
	return &this
}

// NewCipherResponseDataWithDefaults instantiates a new CipherResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCipherResponseDataWithDefaults() *CipherResponseData {
	this := CipherResponseData{}
	return &this
}

// GetDataAvailability returns the DataAvailability field value
func (o *CipherResponseData) GetDataAvailability() DataAvailability {
	if o == nil {
		var ret DataAvailability
		return ret
	}

	return o.DataAvailability
}

// GetDataAvailabilityOk returns a tuple with the DataAvailability field value
// and a boolean to check if the value has been set.
func (o *CipherResponseData) GetDataAvailabilityOk() (*DataAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataAvailability, true
}

// SetDataAvailability sets field value
func (o *CipherResponseData) SetDataAvailability(v DataAvailability) {
	o.DataAvailability = v
}

func (o CipherResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CipherResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataAvailability"] = o.DataAvailability
	return toSerialize, nil
}

type NullableCipherResponseData struct {
	value *CipherResponseData
	isSet bool
}

func (v NullableCipherResponseData) Get() *CipherResponseData {
	return v.value
}

func (v *NullableCipherResponseData) Set(val *CipherResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCipherResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCipherResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCipherResponseData(val *CipherResponseData) *NullableCipherResponseData {
	return &NullableCipherResponseData{value: val, isSet: true}
}

func (v NullableCipherResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCipherResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
