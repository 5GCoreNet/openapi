/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
)

// checks if the VsDataContainerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsDataContainerType{}

// VsDataContainerType container for vendor specific data (see 3GPP TS 28.622)
type VsDataContainerType struct {
	VsDataType          *string `json:"vsDataType,omitempty"`
	VsData              *string `json:"vsData,omitempty"`
	VsDataFormatVersion *string `json:"vsDataFormatVersion,omitempty"`
}

// NewVsDataContainerType instantiates a new VsDataContainerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsDataContainerType() *VsDataContainerType {
	this := VsDataContainerType{}
	return &this
}

// NewVsDataContainerTypeWithDefaults instantiates a new VsDataContainerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsDataContainerTypeWithDefaults() *VsDataContainerType {
	this := VsDataContainerType{}
	return &this
}

// GetVsDataType returns the VsDataType field value if set, zero value otherwise.
func (o *VsDataContainerType) GetVsDataType() string {
	if o == nil || IsNil(o.VsDataType) {
		var ret string
		return ret
	}
	return *o.VsDataType
}

// GetVsDataTypeOk returns a tuple with the VsDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerType) GetVsDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VsDataType) {
		return nil, false
	}
	return o.VsDataType, true
}

// HasVsDataType returns a boolean if a field has been set.
func (o *VsDataContainerType) HasVsDataType() bool {
	if o != nil && !IsNil(o.VsDataType) {
		return true
	}

	return false
}

// SetVsDataType gets a reference to the given string and assigns it to the VsDataType field.
func (o *VsDataContainerType) SetVsDataType(v string) {
	o.VsDataType = &v
}

// GetVsData returns the VsData field value if set, zero value otherwise.
func (o *VsDataContainerType) GetVsData() string {
	if o == nil || IsNil(o.VsData) {
		var ret string
		return ret
	}
	return *o.VsData
}

// GetVsDataOk returns a tuple with the VsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerType) GetVsDataOk() (*string, bool) {
	if o == nil || IsNil(o.VsData) {
		return nil, false
	}
	return o.VsData, true
}

// HasVsData returns a boolean if a field has been set.
func (o *VsDataContainerType) HasVsData() bool {
	if o != nil && !IsNil(o.VsData) {
		return true
	}

	return false
}

// SetVsData gets a reference to the given string and assigns it to the VsData field.
func (o *VsDataContainerType) SetVsData(v string) {
	o.VsData = &v
}

// GetVsDataFormatVersion returns the VsDataFormatVersion field value if set, zero value otherwise.
func (o *VsDataContainerType) GetVsDataFormatVersion() string {
	if o == nil || IsNil(o.VsDataFormatVersion) {
		var ret string
		return ret
	}
	return *o.VsDataFormatVersion
}

// GetVsDataFormatVersionOk returns a tuple with the VsDataFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerType) GetVsDataFormatVersionOk() (*string, bool) {
	if o == nil || IsNil(o.VsDataFormatVersion) {
		return nil, false
	}
	return o.VsDataFormatVersion, true
}

// HasVsDataFormatVersion returns a boolean if a field has been set.
func (o *VsDataContainerType) HasVsDataFormatVersion() bool {
	if o != nil && !IsNil(o.VsDataFormatVersion) {
		return true
	}

	return false
}

// SetVsDataFormatVersion gets a reference to the given string and assigns it to the VsDataFormatVersion field.
func (o *VsDataContainerType) SetVsDataFormatVersion(v string) {
	o.VsDataFormatVersion = &v
}

func (o VsDataContainerType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsDataContainerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VsDataType) {
		toSerialize["vsDataType"] = o.VsDataType
	}
	if !IsNil(o.VsData) {
		toSerialize["vsData"] = o.VsData
	}
	if !IsNil(o.VsDataFormatVersion) {
		toSerialize["vsDataFormatVersion"] = o.VsDataFormatVersion
	}
	return toSerialize, nil
}

type NullableVsDataContainerType struct {
	value *VsDataContainerType
	isSet bool
}

func (v NullableVsDataContainerType) Get() *VsDataContainerType {
	return v.value
}

func (v *NullableVsDataContainerType) Set(val *VsDataContainerType) {
	v.value = val
	v.isSet = true
}

func (v NullableVsDataContainerType) IsSet() bool {
	return v.isSet
}

func (v *NullableVsDataContainerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsDataContainerType(val *VsDataContainerType) *NullableVsDataContainerType {
	return &NullableVsDataContainerType{value: val, isSet: true}
}

func (v NullableVsDataContainerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsDataContainerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
