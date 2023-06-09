/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"os"
)

// checks if the TransferMtDataIsmfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferMtDataIsmfRequest{}

// TransferMtDataIsmfRequest struct for TransferMtDataIsmfRequest
type TransferMtDataIsmfRequest struct {
	JsonData     *TransferMtDataReqData `json:"jsonData,omitempty"`
	BinaryMtData **os.File              `json:"binaryMtData,omitempty"`
}

// NewTransferMtDataIsmfRequest instantiates a new TransferMtDataIsmfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMtDataIsmfRequest() *TransferMtDataIsmfRequest {
	this := TransferMtDataIsmfRequest{}
	return &this
}

// NewTransferMtDataIsmfRequestWithDefaults instantiates a new TransferMtDataIsmfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMtDataIsmfRequestWithDefaults() *TransferMtDataIsmfRequest {
	this := TransferMtDataIsmfRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *TransferMtDataIsmfRequest) GetJsonData() TransferMtDataReqData {
	if o == nil || IsNil(o.JsonData) {
		var ret TransferMtDataReqData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataIsmfRequest) GetJsonDataOk() (*TransferMtDataReqData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *TransferMtDataIsmfRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given TransferMtDataReqData and assigns it to the JsonData field.
func (o *TransferMtDataIsmfRequest) SetJsonData(v TransferMtDataReqData) {
	o.JsonData = &v
}

// GetBinaryMtData returns the BinaryMtData field value if set, zero value otherwise.
func (o *TransferMtDataIsmfRequest) GetBinaryMtData() *os.File {
	if o == nil || IsNil(o.BinaryMtData) {
		var ret *os.File
		return ret
	}
	return *o.BinaryMtData
}

// GetBinaryMtDataOk returns a tuple with the BinaryMtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataIsmfRequest) GetBinaryMtDataOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryMtData) {
		return nil, false
	}
	return o.BinaryMtData, true
}

// HasBinaryMtData returns a boolean if a field has been set.
func (o *TransferMtDataIsmfRequest) HasBinaryMtData() bool {
	if o != nil && !IsNil(o.BinaryMtData) {
		return true
	}

	return false
}

// SetBinaryMtData gets a reference to the given *os.File and assigns it to the BinaryMtData field.
func (o *TransferMtDataIsmfRequest) SetBinaryMtData(v *os.File) {
	o.BinaryMtData = &v
}

func (o TransferMtDataIsmfRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferMtDataIsmfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryMtData) {
		toSerialize["binaryMtData"] = o.BinaryMtData
	}
	return toSerialize, nil
}

type NullableTransferMtDataIsmfRequest struct {
	value *TransferMtDataIsmfRequest
	isSet bool
}

func (v NullableTransferMtDataIsmfRequest) Get() *TransferMtDataIsmfRequest {
	return v.value
}

func (v *NullableTransferMtDataIsmfRequest) Set(val *TransferMtDataIsmfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMtDataIsmfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMtDataIsmfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMtDataIsmfRequest(val *TransferMtDataIsmfRequest) *NullableTransferMtDataIsmfRequest {
	return &NullableTransferMtDataIsmfRequest{value: val, isSet: true}
}

func (v NullableTransferMtDataIsmfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMtDataIsmfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
