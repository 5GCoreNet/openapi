/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the SnssaiTaiMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiTaiMapping{}

// SnssaiTaiMapping List of restricted or unrestricted S-NSSAIs per TAI(s)
type SnssaiTaiMapping struct {
	ReportingArea TargetArea `json:"reportingArea"`
	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`
	SupportedSnssaiList []SupportedSnssai `json:"supportedSnssaiList,omitempty"`
}

// NewSnssaiTaiMapping instantiates a new SnssaiTaiMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiTaiMapping(reportingArea TargetArea) *SnssaiTaiMapping {
	this := SnssaiTaiMapping{}
	this.ReportingArea = reportingArea
	return &this
}

// NewSnssaiTaiMappingWithDefaults instantiates a new SnssaiTaiMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiTaiMappingWithDefaults() *SnssaiTaiMapping {
	this := SnssaiTaiMapping{}
	return &this
}

// GetReportingArea returns the ReportingArea field value
func (o *SnssaiTaiMapping) GetReportingArea() TargetArea {
	if o == nil {
		var ret TargetArea
		return ret
	}

	return o.ReportingArea
}

// GetReportingAreaOk returns a tuple with the ReportingArea field value
// and a boolean to check if the value has been set.
func (o *SnssaiTaiMapping) GetReportingAreaOk() (*TargetArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingArea, true
}

// SetReportingArea sets field value
func (o *SnssaiTaiMapping) SetReportingArea(v TargetArea) {
	o.ReportingArea = v
}

// GetAccessTypeList returns the AccessTypeList field value if set, zero value otherwise.
func (o *SnssaiTaiMapping) GetAccessTypeList() []AccessType {
	if o == nil || IsNil(o.AccessTypeList) {
		var ret []AccessType
		return ret
	}
	return o.AccessTypeList
}

// GetAccessTypeListOk returns a tuple with the AccessTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiTaiMapping) GetAccessTypeListOk() ([]AccessType, bool) {
	if o == nil || IsNil(o.AccessTypeList) {
		return nil, false
	}
	return o.AccessTypeList, true
}

// HasAccessTypeList returns a boolean if a field has been set.
func (o *SnssaiTaiMapping) HasAccessTypeList() bool {
	if o != nil && !IsNil(o.AccessTypeList) {
		return true
	}

	return false
}

// SetAccessTypeList gets a reference to the given []AccessType and assigns it to the AccessTypeList field.
func (o *SnssaiTaiMapping) SetAccessTypeList(v []AccessType) {
	o.AccessTypeList = v
}

// GetSupportedSnssaiList returns the SupportedSnssaiList field value if set, zero value otherwise.
func (o *SnssaiTaiMapping) GetSupportedSnssaiList() []SupportedSnssai {
	if o == nil || IsNil(o.SupportedSnssaiList) {
		var ret []SupportedSnssai
		return ret
	}
	return o.SupportedSnssaiList
}

// GetSupportedSnssaiListOk returns a tuple with the SupportedSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiTaiMapping) GetSupportedSnssaiListOk() ([]SupportedSnssai, bool) {
	if o == nil || IsNil(o.SupportedSnssaiList) {
		return nil, false
	}
	return o.SupportedSnssaiList, true
}

// HasSupportedSnssaiList returns a boolean if a field has been set.
func (o *SnssaiTaiMapping) HasSupportedSnssaiList() bool {
	if o != nil && !IsNil(o.SupportedSnssaiList) {
		return true
	}

	return false
}

// SetSupportedSnssaiList gets a reference to the given []SupportedSnssai and assigns it to the SupportedSnssaiList field.
func (o *SnssaiTaiMapping) SetSupportedSnssaiList(v []SupportedSnssai) {
	o.SupportedSnssaiList = v
}

func (o SnssaiTaiMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiTaiMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportingArea"] = o.ReportingArea
	if !IsNil(o.AccessTypeList) {
		toSerialize["accessTypeList"] = o.AccessTypeList
	}
	if !IsNil(o.SupportedSnssaiList) {
		toSerialize["supportedSnssaiList"] = o.SupportedSnssaiList
	}
	return toSerialize, nil
}

type NullableSnssaiTaiMapping struct {
	value *SnssaiTaiMapping
	isSet bool
}

func (v NullableSnssaiTaiMapping) Get() *SnssaiTaiMapping {
	return v.value
}

func (v *NullableSnssaiTaiMapping) Set(val *SnssaiTaiMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiTaiMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiTaiMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiTaiMapping(val *SnssaiTaiMapping) *NullableSnssaiTaiMapping {
	return &NullableSnssaiTaiMapping{value: val, isSet: true}
}

func (v NullableSnssaiTaiMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiTaiMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

