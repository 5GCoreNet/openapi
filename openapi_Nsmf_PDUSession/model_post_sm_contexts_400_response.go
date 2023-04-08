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

// checks if the PostSmContexts400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSmContexts400Response{}

// PostSmContexts400Response struct for PostSmContexts400Response
type PostSmContexts400Response struct {
	JsonData *SmContextCreateError `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage *os.File `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmMessage *os.File `json:"binaryDataN2SmMessage,omitempty"`
}

// NewPostSmContexts400Response instantiates a new PostSmContexts400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSmContexts400Response() *PostSmContexts400Response {
	this := PostSmContexts400Response{}
	return &this
}

// NewPostSmContexts400ResponseWithDefaults instantiates a new PostSmContexts400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSmContexts400ResponseWithDefaults() *PostSmContexts400Response {
	this := PostSmContexts400Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *PostSmContexts400Response) GetJsonData() SmContextCreateError {
	if o == nil || isNil(o.JsonData) {
		var ret SmContextCreateError
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContexts400Response) GetJsonDataOk() (*SmContextCreateError, bool) {
	if o == nil || isNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *PostSmContexts400Response) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmContextCreateError and assigns it to the JsonData field.
func (o *PostSmContexts400Response) SetJsonData(v SmContextCreateError) {
	o.JsonData = &v
}

// GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field value if set, zero value otherwise.
func (o *PostSmContexts400Response) GetBinaryDataN1SmMessage() os.File {
	if o == nil || isNil(o.BinaryDataN1SmMessage) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataN1SmMessage
}

// GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContexts400Response) GetBinaryDataN1SmMessageOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataN1SmMessage) {
		return nil, false
	}
	return o.BinaryDataN1SmMessage, true
}

// HasBinaryDataN1SmMessage returns a boolean if a field has been set.
func (o *PostSmContexts400Response) HasBinaryDataN1SmMessage() bool {
	if o != nil && !isNil(o.BinaryDataN1SmMessage) {
		return true
	}

	return false
}

// SetBinaryDataN1SmMessage gets a reference to the given os.File and assigns it to the BinaryDataN1SmMessage field.
func (o *PostSmContexts400Response) SetBinaryDataN1SmMessage(v os.File) {
	o.BinaryDataN1SmMessage = &v
}

// GetBinaryDataN2SmMessage returns the BinaryDataN2SmMessage field value if set, zero value otherwise.
func (o *PostSmContexts400Response) GetBinaryDataN2SmMessage() os.File {
	if o == nil || isNil(o.BinaryDataN2SmMessage) {
		var ret os.File
		return ret
	}
	return *o.BinaryDataN2SmMessage
}

// GetBinaryDataN2SmMessageOk returns a tuple with the BinaryDataN2SmMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContexts400Response) GetBinaryDataN2SmMessageOk() (*os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2SmMessage) {
		return nil, false
	}
	return o.BinaryDataN2SmMessage, true
}

// HasBinaryDataN2SmMessage returns a boolean if a field has been set.
func (o *PostSmContexts400Response) HasBinaryDataN2SmMessage() bool {
	if o != nil && !isNil(o.BinaryDataN2SmMessage) {
		return true
	}

	return false
}

// SetBinaryDataN2SmMessage gets a reference to the given os.File and assigns it to the BinaryDataN2SmMessage field.
func (o *PostSmContexts400Response) SetBinaryDataN2SmMessage(v os.File) {
	o.BinaryDataN2SmMessage = &v
}

func (o PostSmContexts400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSmContexts400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN1SmMessage) {
		toSerialize["binaryDataN1SmMessage"] = o.BinaryDataN1SmMessage
	}
	if !isNil(o.BinaryDataN2SmMessage) {
		toSerialize["binaryDataN2SmMessage"] = o.BinaryDataN2SmMessage
	}
	return toSerialize, nil
}

type NullablePostSmContexts400Response struct {
	value *PostSmContexts400Response
	isSet bool
}

func (v NullablePostSmContexts400Response) Get() *PostSmContexts400Response {
	return v.value
}

func (v *NullablePostSmContexts400Response) Set(val *PostSmContexts400Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSmContexts400Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSmContexts400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSmContexts400Response(val *PostSmContexts400Response) *NullablePostSmContexts400Response {
	return &NullablePostSmContexts400Response{value: val, isSet: true}
}

func (v NullablePostSmContexts400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSmContexts400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


