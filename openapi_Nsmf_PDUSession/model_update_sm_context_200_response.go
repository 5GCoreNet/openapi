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

// checks if the UpdateSmContext200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSmContext200Response{}

// UpdateSmContext200Response struct for UpdateSmContext200Response
type UpdateSmContext200Response struct {
	JsonData                  *SmContextUpdatedData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage     **os.File             `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation **os.File             `json:"binaryDataN2SmInformation,omitempty"`
}

// NewUpdateSmContext200Response instantiates a new UpdateSmContext200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSmContext200Response() *UpdateSmContext200Response {
	this := UpdateSmContext200Response{}
	return &this
}

// NewUpdateSmContext200ResponseWithDefaults instantiates a new UpdateSmContext200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSmContext200ResponseWithDefaults() *UpdateSmContext200Response {
	this := UpdateSmContext200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *UpdateSmContext200Response) GetJsonData() SmContextUpdatedData {
	if o == nil || IsNil(o.JsonData) {
		var ret SmContextUpdatedData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContext200Response) GetJsonDataOk() (*SmContextUpdatedData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *UpdateSmContext200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmContextUpdatedData and assigns it to the JsonData field.
func (o *UpdateSmContext200Response) SetJsonData(v SmContextUpdatedData) {
	o.JsonData = &v
}

// GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field value if set, zero value otherwise.
func (o *UpdateSmContext200Response) GetBinaryDataN1SmMessage() *os.File {
	if o == nil || IsNil(o.BinaryDataN1SmMessage) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmMessage
}

// GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContext200Response) GetBinaryDataN1SmMessageOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN1SmMessage) {
		return nil, false
	}
	return o.BinaryDataN1SmMessage, true
}

// HasBinaryDataN1SmMessage returns a boolean if a field has been set.
func (o *UpdateSmContext200Response) HasBinaryDataN1SmMessage() bool {
	if o != nil && !IsNil(o.BinaryDataN1SmMessage) {
		return true
	}

	return false
}

// SetBinaryDataN1SmMessage gets a reference to the given *os.File and assigns it to the BinaryDataN1SmMessage field.
func (o *UpdateSmContext200Response) SetBinaryDataN1SmMessage(v *os.File) {
	o.BinaryDataN1SmMessage = &v
}

// GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field value if set, zero value otherwise.
func (o *UpdateSmContext200Response) GetBinaryDataN2SmInformation() *os.File {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformation
}

// GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSmContext200Response) GetBinaryDataN2SmInformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2SmInformation) {
		return nil, false
	}
	return o.BinaryDataN2SmInformation, true
}

// HasBinaryDataN2SmInformation returns a boolean if a field has been set.
func (o *UpdateSmContext200Response) HasBinaryDataN2SmInformation() bool {
	if o != nil && !IsNil(o.BinaryDataN2SmInformation) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformation gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformation field.
func (o *UpdateSmContext200Response) SetBinaryDataN2SmInformation(v *os.File) {
	o.BinaryDataN2SmInformation = &v
}

func (o UpdateSmContext200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSmContext200Response) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableUpdateSmContext200Response struct {
	value *UpdateSmContext200Response
	isSet bool
}

func (v NullableUpdateSmContext200Response) Get() *UpdateSmContext200Response {
	return v.value
}

func (v *NullableUpdateSmContext200Response) Set(val *UpdateSmContext200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSmContext200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSmContext200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSmContext200Response(val *UpdateSmContext200Response) *NullableUpdateSmContext200Response {
	return &NullableUpdateSmContext200Response{value: val, isSet: true}
}

func (v NullableUpdateSmContext200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSmContext200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
