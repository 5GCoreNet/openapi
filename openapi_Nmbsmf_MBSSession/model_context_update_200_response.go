/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"os"
)

// checks if the ContextUpdate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextUpdate200Response{}

// ContextUpdate200Response struct for ContextUpdate200Response
type ContextUpdate200Response struct {
	JsonData                *ContextUpdateRspData `json:"jsonData,omitempty"`
	BinaryDataN2Information **os.File             `json:"binaryDataN2Information,omitempty"`
}

// NewContextUpdate200Response instantiates a new ContextUpdate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdate200Response() *ContextUpdate200Response {
	this := ContextUpdate200Response{}
	return &this
}

// NewContextUpdate200ResponseWithDefaults instantiates a new ContextUpdate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdate200ResponseWithDefaults() *ContextUpdate200Response {
	this := ContextUpdate200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetJsonData() ContextUpdateRspData {
	if o == nil || IsNil(o.JsonData) {
		var ret ContextUpdateRspData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetJsonDataOk() (*ContextUpdateRspData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given ContextUpdateRspData and assigns it to the JsonData field.
func (o *ContextUpdate200Response) SetJsonData(v ContextUpdateRspData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information returns the BinaryDataN2Information field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information
}

// GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2InformationOk() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information) {
		return nil, false
	}
	return o.BinaryDataN2Information, true
}

// HasBinaryDataN2Information returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information) {
		return true
	}

	return false
}

// SetBinaryDataN2Information gets a reference to the given *os.File and assigns it to the BinaryDataN2Information field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information(v *os.File) {
	o.BinaryDataN2Information = &v
}

func (o ContextUpdate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextUpdate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN2Information) {
		toSerialize["binaryDataN2Information"] = o.BinaryDataN2Information
	}
	return toSerialize, nil
}

type NullableContextUpdate200Response struct {
	value *ContextUpdate200Response
	isSet bool
}

func (v NullableContextUpdate200Response) Get() *ContextUpdate200Response {
	return v.value
}

func (v *NullableContextUpdate200Response) Set(val *ContextUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdate200Response(val *ContextUpdate200Response) *NullableContextUpdate200Response {
	return &NullableContextUpdate200Response{value: val, isSet: true}
}

func (v NullableContextUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
