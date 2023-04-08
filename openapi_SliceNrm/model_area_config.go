/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the AreaConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaConfig{}

// AreaConfig struct for AreaConfig
type AreaConfig struct {
	FreqInfo *FreqInfo `json:"freqInfo,omitempty"`
	PciList []int32 `json:"pciList,omitempty"`
}

// NewAreaConfig instantiates a new AreaConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaConfig() *AreaConfig {
	this := AreaConfig{}
	return &this
}

// NewAreaConfigWithDefaults instantiates a new AreaConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaConfigWithDefaults() *AreaConfig {
	this := AreaConfig{}
	return &this
}

// GetFreqInfo returns the FreqInfo field value if set, zero value otherwise.
func (o *AreaConfig) GetFreqInfo() FreqInfo {
	if o == nil || isNil(o.FreqInfo) {
		var ret FreqInfo
		return ret
	}
	return *o.FreqInfo
}

// GetFreqInfoOk returns a tuple with the FreqInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetFreqInfoOk() (*FreqInfo, bool) {
	if o == nil || isNil(o.FreqInfo) {
		return nil, false
	}
	return o.FreqInfo, true
}

// HasFreqInfo returns a boolean if a field has been set.
func (o *AreaConfig) HasFreqInfo() bool {
	if o != nil && !isNil(o.FreqInfo) {
		return true
	}

	return false
}

// SetFreqInfo gets a reference to the given FreqInfo and assigns it to the FreqInfo field.
func (o *AreaConfig) SetFreqInfo(v FreqInfo) {
	o.FreqInfo = &v
}

// GetPciList returns the PciList field value if set, zero value otherwise.
func (o *AreaConfig) GetPciList() []int32 {
	if o == nil || isNil(o.PciList) {
		var ret []int32
		return ret
	}
	return o.PciList
}

// GetPciListOk returns a tuple with the PciList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaConfig) GetPciListOk() ([]int32, bool) {
	if o == nil || isNil(o.PciList) {
		return nil, false
	}
	return o.PciList, true
}

// HasPciList returns a boolean if a field has been set.
func (o *AreaConfig) HasPciList() bool {
	if o != nil && !isNil(o.PciList) {
		return true
	}

	return false
}

// SetPciList gets a reference to the given []int32 and assigns it to the PciList field.
func (o *AreaConfig) SetPciList(v []int32) {
	o.PciList = v
}

func (o AreaConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AreaConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FreqInfo) {
		toSerialize["freqInfo"] = o.FreqInfo
	}
	if !isNil(o.PciList) {
		toSerialize["pciList"] = o.PciList
	}
	return toSerialize, nil
}

type NullableAreaConfig struct {
	value *AreaConfig
	isSet bool
}

func (v NullableAreaConfig) Get() *AreaConfig {
	return v.value
}

func (v *NullableAreaConfig) Set(val *AreaConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaConfig(val *AreaConfig) *NullableAreaConfig {
	return &NullableAreaConfig{value: val, isSet: true}
}

func (v NullableAreaConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


