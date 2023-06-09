/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
)

// checks if the Uss type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Uss{}

// Uss User Security Settings for a given GAA Service
type Uss struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	GsId int32 `json:"gsId"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	GsType int32       `json:"gsType"`
	UeIds  []UeIdsItem `json:"ueIds"`
	// Character string representing a NAF Group
	NafGroup  *string     `json:"nafGroup,omitempty"`
	Flags     []FlagsItem `json:"flags,omitempty"`
	KeyChoice *KeyChoice  `json:"keyChoice,omitempty"`
}

// NewUss instantiates a new Uss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUss(gsId int32, gsType int32, ueIds []UeIdsItem) *Uss {
	this := Uss{}
	this.GsId = gsId
	this.GsType = gsType
	this.UeIds = ueIds
	return &this
}

// NewUssWithDefaults instantiates a new Uss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUssWithDefaults() *Uss {
	this := Uss{}
	return &this
}

// GetGsId returns the GsId field value
func (o *Uss) GetGsId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GsId
}

// GetGsIdOk returns a tuple with the GsId field value
// and a boolean to check if the value has been set.
func (o *Uss) GetGsIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GsId, true
}

// SetGsId sets field value
func (o *Uss) SetGsId(v int32) {
	o.GsId = v
}

// GetGsType returns the GsType field value
func (o *Uss) GetGsType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GsType
}

// GetGsTypeOk returns a tuple with the GsType field value
// and a boolean to check if the value has been set.
func (o *Uss) GetGsTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GsType, true
}

// SetGsType sets field value
func (o *Uss) SetGsType(v int32) {
	o.GsType = v
}

// GetUeIds returns the UeIds field value
func (o *Uss) GetUeIds() []UeIdsItem {
	if o == nil {
		var ret []UeIdsItem
		return ret
	}

	return o.UeIds
}

// GetUeIdsOk returns a tuple with the UeIds field value
// and a boolean to check if the value has been set.
func (o *Uss) GetUeIdsOk() ([]UeIdsItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.UeIds, true
}

// SetUeIds sets field value
func (o *Uss) SetUeIds(v []UeIdsItem) {
	o.UeIds = v
}

// GetNafGroup returns the NafGroup field value if set, zero value otherwise.
func (o *Uss) GetNafGroup() string {
	if o == nil || IsNil(o.NafGroup) {
		var ret string
		return ret
	}
	return *o.NafGroup
}

// GetNafGroupOk returns a tuple with the NafGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Uss) GetNafGroupOk() (*string, bool) {
	if o == nil || IsNil(o.NafGroup) {
		return nil, false
	}
	return o.NafGroup, true
}

// HasNafGroup returns a boolean if a field has been set.
func (o *Uss) HasNafGroup() bool {
	if o != nil && !IsNil(o.NafGroup) {
		return true
	}

	return false
}

// SetNafGroup gets a reference to the given string and assigns it to the NafGroup field.
func (o *Uss) SetNafGroup(v string) {
	o.NafGroup = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *Uss) GetFlags() []FlagsItem {
	if o == nil || IsNil(o.Flags) {
		var ret []FlagsItem
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Uss) GetFlagsOk() ([]FlagsItem, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *Uss) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []FlagsItem and assigns it to the Flags field.
func (o *Uss) SetFlags(v []FlagsItem) {
	o.Flags = v
}

// GetKeyChoice returns the KeyChoice field value if set, zero value otherwise.
func (o *Uss) GetKeyChoice() KeyChoice {
	if o == nil || IsNil(o.KeyChoice) {
		var ret KeyChoice
		return ret
	}
	return *o.KeyChoice
}

// GetKeyChoiceOk returns a tuple with the KeyChoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Uss) GetKeyChoiceOk() (*KeyChoice, bool) {
	if o == nil || IsNil(o.KeyChoice) {
		return nil, false
	}
	return o.KeyChoice, true
}

// HasKeyChoice returns a boolean if a field has been set.
func (o *Uss) HasKeyChoice() bool {
	if o != nil && !IsNil(o.KeyChoice) {
		return true
	}

	return false
}

// SetKeyChoice gets a reference to the given KeyChoice and assigns it to the KeyChoice field.
func (o *Uss) SetKeyChoice(v KeyChoice) {
	o.KeyChoice = &v
}

func (o Uss) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Uss) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gsId"] = o.GsId
	toSerialize["gsType"] = o.GsType
	toSerialize["ueIds"] = o.UeIds
	if !IsNil(o.NafGroup) {
		toSerialize["nafGroup"] = o.NafGroup
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.KeyChoice) {
		toSerialize["keyChoice"] = o.KeyChoice
	}
	return toSerialize, nil
}

type NullableUss struct {
	value *Uss
	isSet bool
}

func (v NullableUss) Get() *Uss {
	return v.value
}

func (v *NullableUss) Set(val *Uss) {
	v.value = val
	v.isSet = true
}

func (v NullableUss) IsSet() bool {
	return v.isSet
}

func (v *NullableUss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUss(val *Uss) *NullableUss {
	return &NullableUss{value: val, isSet: true}
}

func (v NullableUss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
