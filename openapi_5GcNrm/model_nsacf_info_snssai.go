/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the NsacfInfoSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsacfInfoSnssai{}

// NsacfInfoSnssai struct for NsacfInfoSnssai
type NsacfInfoSnssai struct {
	SnssaiInfo             *SnssaiInfo `json:"SnssaiInfo,omitempty"`
	IsSubjectToNsac        *bool       `json:"isSubjectToNsac,omitempty"`
	MaxNumberofUEs         *int32      `json:"maxNumberofUEs,omitempty"`
	EACMode                *string     `json:"eACMode,omitempty"`
	ActiveEacThreshhold    *int32      `json:"activeEacThreshhold,omitempty"`
	DeactiveEacThreshhold  *int32      `json:"deactiveEacThreshhold,omitempty"`
	NumberofUEs            *int32      `json:"numberofUEs,omitempty"`
	UEIdList               []string    `json:"uEIdList,omitempty"`
	MaxNumberofPDUSessions *int32      `json:"maxNumberofPDUSessions,omitempty"`
}

// NewNsacfInfoSnssai instantiates a new NsacfInfoSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfInfoSnssai() *NsacfInfoSnssai {
	this := NsacfInfoSnssai{}
	return &this
}

// NewNsacfInfoSnssaiWithDefaults instantiates a new NsacfInfoSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfInfoSnssaiWithDefaults() *NsacfInfoSnssai {
	this := NsacfInfoSnssai{}
	return &this
}

// GetSnssaiInfo returns the SnssaiInfo field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetSnssaiInfo() SnssaiInfo {
	if o == nil || IsNil(o.SnssaiInfo) {
		var ret SnssaiInfo
		return ret
	}
	return *o.SnssaiInfo
}

// GetSnssaiInfoOk returns a tuple with the SnssaiInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetSnssaiInfoOk() (*SnssaiInfo, bool) {
	if o == nil || IsNil(o.SnssaiInfo) {
		return nil, false
	}
	return o.SnssaiInfo, true
}

// HasSnssaiInfo returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasSnssaiInfo() bool {
	if o != nil && !IsNil(o.SnssaiInfo) {
		return true
	}

	return false
}

// SetSnssaiInfo gets a reference to the given SnssaiInfo and assigns it to the SnssaiInfo field.
func (o *NsacfInfoSnssai) SetSnssaiInfo(v SnssaiInfo) {
	o.SnssaiInfo = &v
}

// GetIsSubjectToNsac returns the IsSubjectToNsac field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetIsSubjectToNsac() bool {
	if o == nil || IsNil(o.IsSubjectToNsac) {
		var ret bool
		return ret
	}
	return *o.IsSubjectToNsac
}

// GetIsSubjectToNsacOk returns a tuple with the IsSubjectToNsac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetIsSubjectToNsacOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubjectToNsac) {
		return nil, false
	}
	return o.IsSubjectToNsac, true
}

// HasIsSubjectToNsac returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasIsSubjectToNsac() bool {
	if o != nil && !IsNil(o.IsSubjectToNsac) {
		return true
	}

	return false
}

// SetIsSubjectToNsac gets a reference to the given bool and assigns it to the IsSubjectToNsac field.
func (o *NsacfInfoSnssai) SetIsSubjectToNsac(v bool) {
	o.IsSubjectToNsac = &v
}

// GetMaxNumberofUEs returns the MaxNumberofUEs field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetMaxNumberofUEs() int32 {
	if o == nil || IsNil(o.MaxNumberofUEs) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofUEs
}

// GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetMaxNumberofUEsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberofUEs) {
		return nil, false
	}
	return o.MaxNumberofUEs, true
}

// HasMaxNumberofUEs returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasMaxNumberofUEs() bool {
	if o != nil && !IsNil(o.MaxNumberofUEs) {
		return true
	}

	return false
}

// SetMaxNumberofUEs gets a reference to the given int32 and assigns it to the MaxNumberofUEs field.
func (o *NsacfInfoSnssai) SetMaxNumberofUEs(v int32) {
	o.MaxNumberofUEs = &v
}

// GetEACMode returns the EACMode field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetEACMode() string {
	if o == nil || IsNil(o.EACMode) {
		var ret string
		return ret
	}
	return *o.EACMode
}

// GetEACModeOk returns a tuple with the EACMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetEACModeOk() (*string, bool) {
	if o == nil || IsNil(o.EACMode) {
		return nil, false
	}
	return o.EACMode, true
}

// HasEACMode returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasEACMode() bool {
	if o != nil && !IsNil(o.EACMode) {
		return true
	}

	return false
}

// SetEACMode gets a reference to the given string and assigns it to the EACMode field.
func (o *NsacfInfoSnssai) SetEACMode(v string) {
	o.EACMode = &v
}

// GetActiveEacThreshhold returns the ActiveEacThreshhold field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetActiveEacThreshhold() int32 {
	if o == nil || IsNil(o.ActiveEacThreshhold) {
		var ret int32
		return ret
	}
	return *o.ActiveEacThreshhold
}

// GetActiveEacThreshholdOk returns a tuple with the ActiveEacThreshhold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetActiveEacThreshholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveEacThreshhold) {
		return nil, false
	}
	return o.ActiveEacThreshhold, true
}

// HasActiveEacThreshhold returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasActiveEacThreshhold() bool {
	if o != nil && !IsNil(o.ActiveEacThreshhold) {
		return true
	}

	return false
}

// SetActiveEacThreshhold gets a reference to the given int32 and assigns it to the ActiveEacThreshhold field.
func (o *NsacfInfoSnssai) SetActiveEacThreshhold(v int32) {
	o.ActiveEacThreshhold = &v
}

// GetDeactiveEacThreshhold returns the DeactiveEacThreshhold field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetDeactiveEacThreshhold() int32 {
	if o == nil || IsNil(o.DeactiveEacThreshhold) {
		var ret int32
		return ret
	}
	return *o.DeactiveEacThreshhold
}

// GetDeactiveEacThreshholdOk returns a tuple with the DeactiveEacThreshhold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetDeactiveEacThreshholdOk() (*int32, bool) {
	if o == nil || IsNil(o.DeactiveEacThreshhold) {
		return nil, false
	}
	return o.DeactiveEacThreshhold, true
}

// HasDeactiveEacThreshhold returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasDeactiveEacThreshhold() bool {
	if o != nil && !IsNil(o.DeactiveEacThreshhold) {
		return true
	}

	return false
}

// SetDeactiveEacThreshhold gets a reference to the given int32 and assigns it to the DeactiveEacThreshhold field.
func (o *NsacfInfoSnssai) SetDeactiveEacThreshhold(v int32) {
	o.DeactiveEacThreshhold = &v
}

// GetNumberofUEs returns the NumberofUEs field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetNumberofUEs() int32 {
	if o == nil || IsNil(o.NumberofUEs) {
		var ret int32
		return ret
	}
	return *o.NumberofUEs
}

// GetNumberofUEsOk returns a tuple with the NumberofUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetNumberofUEsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberofUEs) {
		return nil, false
	}
	return o.NumberofUEs, true
}

// HasNumberofUEs returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasNumberofUEs() bool {
	if o != nil && !IsNil(o.NumberofUEs) {
		return true
	}

	return false
}

// SetNumberofUEs gets a reference to the given int32 and assigns it to the NumberofUEs field.
func (o *NsacfInfoSnssai) SetNumberofUEs(v int32) {
	o.NumberofUEs = &v
}

// GetUEIdList returns the UEIdList field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetUEIdList() []string {
	if o == nil || IsNil(o.UEIdList) {
		var ret []string
		return ret
	}
	return o.UEIdList
}

// GetUEIdListOk returns a tuple with the UEIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetUEIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.UEIdList) {
		return nil, false
	}
	return o.UEIdList, true
}

// HasUEIdList returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasUEIdList() bool {
	if o != nil && !IsNil(o.UEIdList) {
		return true
	}

	return false
}

// SetUEIdList gets a reference to the given []string and assigns it to the UEIdList field.
func (o *NsacfInfoSnssai) SetUEIdList(v []string) {
	o.UEIdList = v
}

// GetMaxNumberofPDUSessions returns the MaxNumberofPDUSessions field value if set, zero value otherwise.
func (o *NsacfInfoSnssai) GetMaxNumberofPDUSessions() int32 {
	if o == nil || IsNil(o.MaxNumberofPDUSessions) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofPDUSessions
}

// GetMaxNumberofPDUSessionsOk returns a tuple with the MaxNumberofPDUSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfoSnssai) GetMaxNumberofPDUSessionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumberofPDUSessions) {
		return nil, false
	}
	return o.MaxNumberofPDUSessions, true
}

// HasMaxNumberofPDUSessions returns a boolean if a field has been set.
func (o *NsacfInfoSnssai) HasMaxNumberofPDUSessions() bool {
	if o != nil && !IsNil(o.MaxNumberofPDUSessions) {
		return true
	}

	return false
}

// SetMaxNumberofPDUSessions gets a reference to the given int32 and assigns it to the MaxNumberofPDUSessions field.
func (o *NsacfInfoSnssai) SetMaxNumberofPDUSessions(v int32) {
	o.MaxNumberofPDUSessions = &v
}

func (o NsacfInfoSnssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsacfInfoSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnssaiInfo) {
		toSerialize["SnssaiInfo"] = o.SnssaiInfo
	}
	if !IsNil(o.IsSubjectToNsac) {
		toSerialize["isSubjectToNsac"] = o.IsSubjectToNsac
	}
	if !IsNil(o.MaxNumberofUEs) {
		toSerialize["maxNumberofUEs"] = o.MaxNumberofUEs
	}
	if !IsNil(o.EACMode) {
		toSerialize["eACMode"] = o.EACMode
	}
	if !IsNil(o.ActiveEacThreshhold) {
		toSerialize["activeEacThreshhold"] = o.ActiveEacThreshhold
	}
	if !IsNil(o.DeactiveEacThreshhold) {
		toSerialize["deactiveEacThreshhold"] = o.DeactiveEacThreshhold
	}
	if !IsNil(o.NumberofUEs) {
		toSerialize["numberofUEs"] = o.NumberofUEs
	}
	if !IsNil(o.UEIdList) {
		toSerialize["uEIdList"] = o.UEIdList
	}
	if !IsNil(o.MaxNumberofPDUSessions) {
		toSerialize["maxNumberofPDUSessions"] = o.MaxNumberofPDUSessions
	}
	return toSerialize, nil
}

type NullableNsacfInfoSnssai struct {
	value *NsacfInfoSnssai
	isSet bool
}

func (v NullableNsacfInfoSnssai) Get() *NsacfInfoSnssai {
	return v.value
}

func (v *NullableNsacfInfoSnssai) Set(val *NsacfInfoSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfInfoSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfInfoSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfInfoSnssai(val *NsacfInfoSnssai) *NullableNsacfInfoSnssai {
	return &NullableNsacfInfoSnssai{value: val, isSet: true}
}

func (v NullableNsacfInfoSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfInfoSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
