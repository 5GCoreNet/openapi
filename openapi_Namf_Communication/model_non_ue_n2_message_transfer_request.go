/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"os"
)

// checks if the NonUeN2MessageTransferRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonUeN2MessageTransferRequest{}

// NonUeN2MessageTransferRequest struct for NonUeN2MessageTransferRequest
type NonUeN2MessageTransferRequest struct {
	JsonData                *N2InformationTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN2Information **os.File                     `json:"binaryDataN2Information,omitempty"`
}

// NewNonUeN2MessageTransferRequest instantiates a new NonUeN2MessageTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonUeN2MessageTransferRequest() *NonUeN2MessageTransferRequest {
	this := NonUeN2MessageTransferRequest{}
	return &this
}

// NewNonUeN2MessageTransferRequestWithDefaults instantiates a new NonUeN2MessageTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonUeN2MessageTransferRequestWithDefaults() *NonUeN2MessageTransferRequest {
	this := NonUeN2MessageTransferRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *NonUeN2MessageTransferRequest) GetJsonData() N2InformationTransferReqData {
	if o == nil || IsNil(o.JsonData) {
		var ret N2InformationTransferReqData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2MessageTransferRequest) GetJsonDataOk() (*N2InformationTransferReqData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *NonUeN2MessageTransferRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given N2InformationTransferReqData and assigns it to the JsonData field.
func (o *NonUeN2MessageTransferRequest) SetJsonData(v N2InformationTransferReqData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information returns the BinaryDataN2Information field value if set, zero value otherwise.
func (o *NonUeN2MessageTransferRequest) GetBinaryDataN2Information() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information
}

// GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonUeN2MessageTransferRequest) GetBinaryDataN2InformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information) {
		return nil, false
	}
	return o.BinaryDataN2Information, true
}

// HasBinaryDataN2Information returns a boolean if a field has been set.
func (o *NonUeN2MessageTransferRequest) HasBinaryDataN2Information() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information) {
		return true
	}

	return false
}

// SetBinaryDataN2Information gets a reference to the given *os.File and assigns it to the BinaryDataN2Information field.
func (o *NonUeN2MessageTransferRequest) SetBinaryDataN2Information(v *os.File) {
	o.BinaryDataN2Information = &v
}

func (o NonUeN2MessageTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonUeN2MessageTransferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN2Information) {
		toSerialize["binaryDataN2Information"] = o.BinaryDataN2Information
	}
	return toSerialize, nil
}

type NullableNonUeN2MessageTransferRequest struct {
	value *NonUeN2MessageTransferRequest
	isSet bool
}

func (v NullableNonUeN2MessageTransferRequest) Get() *NonUeN2MessageTransferRequest {
	return v.value
}

func (v *NullableNonUeN2MessageTransferRequest) Set(val *NonUeN2MessageTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNonUeN2MessageTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNonUeN2MessageTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonUeN2MessageTransferRequest(val *NonUeN2MessageTransferRequest) *NullableNonUeN2MessageTransferRequest {
	return &NullableNonUeN2MessageTransferRequest{value: val, isSet: true}
}

func (v NullableNonUeN2MessageTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonUeN2MessageTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
