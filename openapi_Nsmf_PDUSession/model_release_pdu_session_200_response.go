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

// checks if the ReleasePduSession200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleasePduSession200Response{}

// ReleasePduSession200Response struct for ReleasePduSession200Response
type ReleasePduSession200Response struct {
	JsonData                    *ReleasedData `json:"jsonData,omitempty"`
	BinaryDataN4Information     **os.File     `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 **os.File     `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 **os.File     `json:"binaryDataN4InformationExt2,omitempty"`
}

// NewReleasePduSession200Response instantiates a new ReleasePduSession200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleasePduSession200Response() *ReleasePduSession200Response {
	this := ReleasePduSession200Response{}
	return &this
}

// NewReleasePduSession200ResponseWithDefaults instantiates a new ReleasePduSession200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleasePduSession200ResponseWithDefaults() *ReleasePduSession200Response {
	this := ReleasePduSession200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ReleasePduSession200Response) GetJsonData() ReleasedData {
	if o == nil || IsNil(o.JsonData) {
		var ret ReleasedData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSession200Response) GetJsonDataOk() (*ReleasedData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ReleasePduSession200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given ReleasedData and assigns it to the JsonData field.
func (o *ReleasePduSession200Response) SetJsonData(v ReleasedData) {
	o.JsonData = &v
}

// GetBinaryDataN4Information returns the BinaryDataN4Information field value if set, zero value otherwise.
func (o *ReleasePduSession200Response) GetBinaryDataN4Information() *os.File {
	if o == nil || IsNil(o.BinaryDataN4Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4Information
}

// GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSession200Response) GetBinaryDataN4InformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4Information) {
		return nil, false
	}
	return o.BinaryDataN4Information, true
}

// HasBinaryDataN4Information returns a boolean if a field has been set.
func (o *ReleasePduSession200Response) HasBinaryDataN4Information() bool {
	if o != nil && !IsNil(o.BinaryDataN4Information) {
		return true
	}

	return false
}

// SetBinaryDataN4Information gets a reference to the given *os.File and assigns it to the BinaryDataN4Information field.
func (o *ReleasePduSession200Response) SetBinaryDataN4Information(v *os.File) {
	o.BinaryDataN4Information = &v
}

// GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field value if set, zero value otherwise.
func (o *ReleasePduSession200Response) GetBinaryDataN4InformationExt1() *os.File {
	if o == nil || IsNil(o.BinaryDataN4InformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt1
}

// GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSession200Response) GetBinaryDataN4InformationExt1Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4InformationExt1) {
		return nil, false
	}
	return o.BinaryDataN4InformationExt1, true
}

// HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.
func (o *ReleasePduSession200Response) HasBinaryDataN4InformationExt1() bool {
	if o != nil && !IsNil(o.BinaryDataN4InformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt1 field.
func (o *ReleasePduSession200Response) SetBinaryDataN4InformationExt1(v *os.File) {
	o.BinaryDataN4InformationExt1 = &v
}

// GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field value if set, zero value otherwise.
func (o *ReleasePduSession200Response) GetBinaryDataN4InformationExt2() *os.File {
	if o == nil || IsNil(o.BinaryDataN4InformationExt2) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt2
}

// GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleasePduSession200Response) GetBinaryDataN4InformationExt2Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN4InformationExt2) {
		return nil, false
	}
	return o.BinaryDataN4InformationExt2, true
}

// HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.
func (o *ReleasePduSession200Response) HasBinaryDataN4InformationExt2() bool {
	if o != nil && !IsNil(o.BinaryDataN4InformationExt2) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt2 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt2 field.
func (o *ReleasePduSession200Response) SetBinaryDataN4InformationExt2(v *os.File) {
	o.BinaryDataN4InformationExt2 = &v
}

func (o ReleasePduSession200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleasePduSession200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN4Information) {
		toSerialize["binaryDataN4Information"] = o.BinaryDataN4Information
	}
	if !IsNil(o.BinaryDataN4InformationExt1) {
		toSerialize["binaryDataN4InformationExt1"] = o.BinaryDataN4InformationExt1
	}
	if !IsNil(o.BinaryDataN4InformationExt2) {
		toSerialize["binaryDataN4InformationExt2"] = o.BinaryDataN4InformationExt2
	}
	return toSerialize, nil
}

type NullableReleasePduSession200Response struct {
	value *ReleasePduSession200Response
	isSet bool
}

func (v NullableReleasePduSession200Response) Get() *ReleasePduSession200Response {
	return v.value
}

func (v *NullableReleasePduSession200Response) Set(val *ReleasePduSession200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReleasePduSession200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReleasePduSession200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleasePduSession200Response(val *ReleasePduSession200Response) *NullableReleasePduSession200Response {
	return &NullableReleasePduSession200Response{value: val, isSet: true}
}

func (v NullableReleasePduSession200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleasePduSession200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
