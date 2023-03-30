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

// checks if the TsctsfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TsctsfInfo{}

// TsctsfInfo Information of a TSCTSF NF Instance
type TsctsfInfo struct {
	// A map (list of key-value pairs) where a valid JSON string serves as key
	SNssaiInfoList *map[string]SnssaiTsctsfInfoItem `json:"sNssaiInfoList,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	InternalGroupIdentifiersRanges []InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty"`
}

// NewTsctsfInfo instantiates a new TsctsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTsctsfInfo() *TsctsfInfo {
	this := TsctsfInfo{}
	return &this
}

// NewTsctsfInfoWithDefaults instantiates a new TsctsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTsctsfInfoWithDefaults() *TsctsfInfo {
	this := TsctsfInfo{}
	return &this
}

// GetSNssaiInfoList returns the SNssaiInfoList field value if set, zero value otherwise.
func (o *TsctsfInfo) GetSNssaiInfoList() map[string]SnssaiTsctsfInfoItem {
	if o == nil || IsNil(o.SNssaiInfoList) {
		var ret map[string]SnssaiTsctsfInfoItem
		return ret
	}
	return *o.SNssaiInfoList
}

// GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsctsfInfo) GetSNssaiInfoListOk() (*map[string]SnssaiTsctsfInfoItem, bool) {
	if o == nil || IsNil(o.SNssaiInfoList) {
		return nil, false
	}
	return o.SNssaiInfoList, true
}

// HasSNssaiInfoList returns a boolean if a field has been set.
func (o *TsctsfInfo) HasSNssaiInfoList() bool {
	if o != nil && !IsNil(o.SNssaiInfoList) {
		return true
	}

	return false
}

// SetSNssaiInfoList gets a reference to the given map[string]SnssaiTsctsfInfoItem and assigns it to the SNssaiInfoList field.
func (o *TsctsfInfo) SetSNssaiInfoList(v map[string]SnssaiTsctsfInfoItem) {
	o.SNssaiInfoList = &v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *TsctsfInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsctsfInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *TsctsfInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *TsctsfInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *TsctsfInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsctsfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *TsctsfInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *TsctsfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *TsctsfInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsctsfInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *TsctsfInfo) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *TsctsfInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *TsctsfInfo) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		var ret []InternalGroupIdRange
		return ret
	}
	return o.InternalGroupIdentifiersRanges
}

// GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsctsfInfo) GetInternalGroupIdentifiersRangesOk() ([]InternalGroupIdRange, bool) {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.InternalGroupIdentifiersRanges, true
}

// HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *TsctsfInfo) HasInternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.InternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetInternalGroupIdentifiersRanges gets a reference to the given []InternalGroupIdRange and assigns it to the InternalGroupIdentifiersRanges field.
func (o *TsctsfInfo) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange) {
	o.InternalGroupIdentifiersRanges = v
}

func (o TsctsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TsctsfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SNssaiInfoList) {
		toSerialize["sNssaiInfoList"] = o.SNssaiInfoList
	}
	if !IsNil(o.ExternalGroupIdentifiersRanges) {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.InternalGroupIdentifiersRanges) {
		toSerialize["internalGroupIdentifiersRanges"] = o.InternalGroupIdentifiersRanges
	}
	return toSerialize, nil
}

type NullableTsctsfInfo struct {
	value *TsctsfInfo
	isSet bool
}

func (v NullableTsctsfInfo) Get() *TsctsfInfo {
	return v.value
}

func (v *NullableTsctsfInfo) Set(val *TsctsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTsctsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTsctsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTsctsfInfo(val *TsctsfInfo) *NullableTsctsfInfo {
	return &NullableTsctsfInfo{value: val, isSet: true}
}

func (v NullableTsctsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTsctsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


