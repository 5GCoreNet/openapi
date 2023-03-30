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

// checks if the UpdateSmContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSmContextRequest{}

// UpdateSmContextRequest struct for UpdateSmContextRequest
type UpdateSmContextRequest struct {
	JsonData *SmContextUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage **os.File `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation **os.File `json:"binaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationExt1 **os.File `json:"binaryDataN2SmInformationExt1,omitempty"`
}

// NewUpdateSmContextRequest instantiates a new UpdateSmContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSmContextRequest() *UpdateSmContextRequest {
	this := UpdateSmContextRequest{}
	return &this
}

// NewUpdateSmContextRequestWithDefaults instantiates a new UpdateSmContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSmContextRequestWithDefaults() *UpdateSmContextRequest {
	this := UpdateSmContextRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *UpdateSmContextRequest) GetJsonData() SmContextUpdateData {
	if o == nil || IsNil(o.JsonData) {
		var ret SmContextUpdateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContextRequest) GetJsonDataOk() (*SmContextUpdateData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *UpdateSmContextRequest) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmContextUpdateData and assigns it to the JsonData field.
func (o *UpdateSmContextRequest) SetJsonData(v SmContextUpdateData) {
	o.JsonData = &v
}

// GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field value if set, zero value otherwise.
func (o *UpdateSmContextRequest) GetBinaryDataN1SmMessage() *os.File {
	if o == nil || IsNil(o.BinaryDataN1SmMessage) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmMessage
}

// GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContextRequest) GetBinaryDataN1SmMessageOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN1SmMessage) {
		return nil, false
	}
	return o.BinaryDataN1SmMessage, true
}

// HasBinaryDataN1SmMessage returns a boolean if a field has been set.
func (o *UpdateSmContextRequest) HasBinaryDataN1SmMessage() bool {
	if o != nil && !IsNil(o.BinaryDataN1SmMessage) {
		return true
	}

	return false
}

// SetBinaryDataN1SmMessage gets a reference to the given *os.File and assigns it to the BinaryDataN1SmMessage field.
func (o *UpdateSmContextRequest) SetBinaryDataN1SmMessage(v *os.File) {
	o.BinaryDataN1SmMessage = &v
}

// GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field value if set, zero value otherwise.
func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformation() *os.File {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformation
}

// GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		return nil, false
	}
	return o.BinaryDataN2SmInformation, true
}

// HasBinaryDataN2SmInformation returns a boolean if a field has been set.
func (o *UpdateSmContextRequest) HasBinaryDataN2SmInformation() bool {
	if o != nil && !IsNil(o.BinaryDataN2SmInformation) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformation gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformation field.
func (o *UpdateSmContextRequest) SetBinaryDataN2SmInformation(v *os.File) {
	o.BinaryDataN2SmInformation = &v
}

// GetBinaryDataN2SmInformationExt1 returns the BinaryDataN2SmInformationExt1 field value if set, zero value otherwise.
func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationExt1() *os.File {
	if o == nil || IsNil(o.BinaryDataN2SmInformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformationExt1
}

// GetBinaryDataN2SmInformationExt1Ok returns a tuple with the BinaryDataN2SmInformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContextRequest) GetBinaryDataN2SmInformationExt1Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2SmInformationExt1) {
		return nil, false
	}
	return o.BinaryDataN2SmInformationExt1, true
}

// HasBinaryDataN2SmInformationExt1 returns a boolean if a field has been set.
func (o *UpdateSmContextRequest) HasBinaryDataN2SmInformationExt1() bool {
	if o != nil && !IsNil(o.BinaryDataN2SmInformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformationExt1 field.
func (o *UpdateSmContextRequest) SetBinaryDataN2SmInformationExt1(v *os.File) {
	o.BinaryDataN2SmInformationExt1 = &v
}

func (o UpdateSmContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSmContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN1SmMessage) {
		toSerialize["binaryDataN1SmMessage"] = o.BinaryDataN1SmMessage
	}
	if !IsNil(o.BinaryDataN2SmInformation) {
		toSerialize["binaryDataN2SmInformation"] = o.BinaryDataN2SmInformation
	}
	if !IsNil(o.BinaryDataN2SmInformationExt1) {
		toSerialize["binaryDataN2SmInformationExt1"] = o.BinaryDataN2SmInformationExt1
	}
	return toSerialize, nil
}

type NullableUpdateSmContextRequest struct {
	value *UpdateSmContextRequest
	isSet bool
}

func (v NullableUpdateSmContextRequest) Get() *UpdateSmContextRequest {
	return v.value
}

func (v *NullableUpdateSmContextRequest) Set(val *UpdateSmContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSmContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSmContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSmContextRequest(val *UpdateSmContextRequest) *NullableUpdateSmContextRequest {
	return &NullableUpdateSmContextRequest{value: val, isSet: true}
}

func (v NullableUpdateSmContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSmContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


