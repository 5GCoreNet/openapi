/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// checks if the N2InfoContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N2InfoContent{}

// N2InfoContent Represents a transparent N2 information content to be relayed by AMF
type N2InfoContent struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NgapMessageType *int32          `json:"ngapMessageType,omitempty"`
	NgapIeType      *NgapIeType     `json:"ngapIeType,omitempty"`
	NgapData        RefToBinaryData `json:"ngapData"`
}

// NewN2InfoContent instantiates a new N2InfoContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2InfoContent(ngapData RefToBinaryData) *N2InfoContent {
	this := N2InfoContent{}
	this.NgapData = ngapData
	return &this
}

// NewN2InfoContentWithDefaults instantiates a new N2InfoContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2InfoContentWithDefaults() *N2InfoContent {
	this := N2InfoContent{}
	return &this
}

// GetNgapMessageType returns the NgapMessageType field value if set, zero value otherwise.
func (o *N2InfoContent) GetNgapMessageType() int32 {
	if o == nil || IsNil(o.NgapMessageType) {
		var ret int32
		return ret
	}
	return *o.NgapMessageType
}

// GetNgapMessageTypeOk returns a tuple with the NgapMessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContent) GetNgapMessageTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.NgapMessageType) {
		return nil, false
	}
	return o.NgapMessageType, true
}

// HasNgapMessageType returns a boolean if a field has been set.
func (o *N2InfoContent) HasNgapMessageType() bool {
	if o != nil && !IsNil(o.NgapMessageType) {
		return true
	}

	return false
}

// SetNgapMessageType gets a reference to the given int32 and assigns it to the NgapMessageType field.
func (o *N2InfoContent) SetNgapMessageType(v int32) {
	o.NgapMessageType = &v
}

// GetNgapIeType returns the NgapIeType field value if set, zero value otherwise.
func (o *N2InfoContent) GetNgapIeType() NgapIeType {
	if o == nil || IsNil(o.NgapIeType) {
		var ret NgapIeType
		return ret
	}
	return *o.NgapIeType
}

// GetNgapIeTypeOk returns a tuple with the NgapIeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContent) GetNgapIeTypeOk() (*NgapIeType, bool) {
	if o == nil || IsNil(o.NgapIeType) {
		return nil, false
	}
	return o.NgapIeType, true
}

// HasNgapIeType returns a boolean if a field has been set.
func (o *N2InfoContent) HasNgapIeType() bool {
	if o != nil && !IsNil(o.NgapIeType) {
		return true
	}

	return false
}

// SetNgapIeType gets a reference to the given NgapIeType and assigns it to the NgapIeType field.
func (o *N2InfoContent) SetNgapIeType(v NgapIeType) {
	o.NgapIeType = &v
}

// GetNgapData returns the NgapData field value
func (o *N2InfoContent) GetNgapData() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.NgapData
}

// GetNgapDataOk returns a tuple with the NgapData field value
// and a boolean to check if the value has been set.
func (o *N2InfoContent) GetNgapDataOk() (*RefToBinaryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NgapData, true
}

// SetNgapData sets field value
func (o *N2InfoContent) SetNgapData(v RefToBinaryData) {
	o.NgapData = v
}

func (o N2InfoContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N2InfoContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NgapMessageType) {
		toSerialize["ngapMessageType"] = o.NgapMessageType
	}
	if !IsNil(o.NgapIeType) {
		toSerialize["ngapIeType"] = o.NgapIeType
	}
	toSerialize["ngapData"] = o.NgapData
	return toSerialize, nil
}

type NullableN2InfoContent struct {
	value *N2InfoContent
	isSet bool
}

func (v NullableN2InfoContent) Get() *N2InfoContent {
	return v.value
}

func (v *NullableN2InfoContent) Set(val *N2InfoContent) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InfoContent) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InfoContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InfoContent(val *N2InfoContent) *NullableN2InfoContent {
	return &NullableN2InfoContent{value: val, isSet: true}
}

func (v NullableN2InfoContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InfoContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
