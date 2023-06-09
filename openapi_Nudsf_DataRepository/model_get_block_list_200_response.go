/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
)

// checks if the GetBlockList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBlockList200Response{}

// GetBlockList200Response struct for GetBlockList200Response
type GetBlockList200Response struct {
	// an array of Block parts, can be empty
	Blocks []interface{} `json:"blocks,omitempty"`
}

// NewGetBlockList200Response instantiates a new GetBlockList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockList200Response() *GetBlockList200Response {
	this := GetBlockList200Response{}
	return &this
}

// NewGetBlockList200ResponseWithDefaults instantiates a new GetBlockList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockList200ResponseWithDefaults() *GetBlockList200Response {
	this := GetBlockList200Response{}
	return &this
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *GetBlockList200Response) GetBlocks() []interface{} {
	if o == nil || IsNil(o.Blocks) {
		var ret []interface{}
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBlockList200Response) GetBlocksOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *GetBlockList200Response) HasBlocks() bool {
	if o != nil && !IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []interface{} and assigns it to the Blocks field.
func (o *GetBlockList200Response) SetBlocks(v []interface{}) {
	o.Blocks = v
}

func (o GetBlockList200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBlockList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	return toSerialize, nil
}

type NullableGetBlockList200Response struct {
	value *GetBlockList200Response
	isSet bool
}

func (v NullableGetBlockList200Response) Get() *GetBlockList200Response {
	return v.value
}

func (v *NullableGetBlockList200Response) Set(val *GetBlockList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockList200Response(val *GetBlockList200Response) *NullableGetBlockList200Response {
	return &NullableGetBlockList200Response{value: val, isSet: true}
}

func (v NullableGetBlockList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
