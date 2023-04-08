/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the InterFreqTargetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterFreqTargetInfo{}

// InterFreqTargetInfo Indicates the Inter Frequency Target information.
type InterFreqTargetInfo struct {
	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331. 
	DlCarrierFreq int32 `json:"dlCarrierFreq"`
	// When present, this IE shall contain a list of the physical cell identities where the UE is requested to perform measurement logging for the indicated frequency. 
	CellIdList []int32 `json:"cellIdList,omitempty"`
}

// NewInterFreqTargetInfo instantiates a new InterFreqTargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterFreqTargetInfo(dlCarrierFreq int32) *InterFreqTargetInfo {
	this := InterFreqTargetInfo{}
	this.DlCarrierFreq = dlCarrierFreq
	return &this
}

// NewInterFreqTargetInfoWithDefaults instantiates a new InterFreqTargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterFreqTargetInfoWithDefaults() *InterFreqTargetInfo {
	this := InterFreqTargetInfo{}
	return &this
}

// GetDlCarrierFreq returns the DlCarrierFreq field value
func (o *InterFreqTargetInfo) GetDlCarrierFreq() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DlCarrierFreq
}

// GetDlCarrierFreqOk returns a tuple with the DlCarrierFreq field value
// and a boolean to check if the value has been set.
func (o *InterFreqTargetInfo) GetDlCarrierFreqOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DlCarrierFreq, true
}

// SetDlCarrierFreq sets field value
func (o *InterFreqTargetInfo) SetDlCarrierFreq(v int32) {
	o.DlCarrierFreq = v
}

// GetCellIdList returns the CellIdList field value if set, zero value otherwise.
func (o *InterFreqTargetInfo) GetCellIdList() []int32 {
	if o == nil || isNil(o.CellIdList) {
		var ret []int32
		return ret
	}
	return o.CellIdList
}

// GetCellIdListOk returns a tuple with the CellIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterFreqTargetInfo) GetCellIdListOk() ([]int32, bool) {
	if o == nil || isNil(o.CellIdList) {
		return nil, false
	}
	return o.CellIdList, true
}

// HasCellIdList returns a boolean if a field has been set.
func (o *InterFreqTargetInfo) HasCellIdList() bool {
	if o != nil && !isNil(o.CellIdList) {
		return true
	}

	return false
}

// SetCellIdList gets a reference to the given []int32 and assigns it to the CellIdList field.
func (o *InterFreqTargetInfo) SetCellIdList(v []int32) {
	o.CellIdList = v
}

func (o InterFreqTargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterFreqTargetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dlCarrierFreq"] = o.DlCarrierFreq
	if !isNil(o.CellIdList) {
		toSerialize["cellIdList"] = o.CellIdList
	}
	return toSerialize, nil
}

type NullableInterFreqTargetInfo struct {
	value *InterFreqTargetInfo
	isSet bool
}

func (v NullableInterFreqTargetInfo) Get() *InterFreqTargetInfo {
	return v.value
}

func (v *NullableInterFreqTargetInfo) Set(val *InterFreqTargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInterFreqTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInterFreqTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterFreqTargetInfo(val *InterFreqTargetInfo) *NullableInterFreqTargetInfo {
	return &NullableInterFreqTargetInfo{value: val, isSet: true}
}

func (v NullableInterFreqTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterFreqTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


