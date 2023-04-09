/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the NssaafInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NssaafInfo{}

// NssaafInfo Information of a NSSAAF Instance
type NssaafInfo struct {
	SupiRanges                     []SupiRange            `json:"supiRanges,omitempty"`
	InternalGroupIdentifiersRanges []InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty"`
}

// NewNssaafInfo instantiates a new NssaafInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssaafInfo() *NssaafInfo {
	this := NssaafInfo{}
	return &this
}

// NewNssaafInfoWithDefaults instantiates a new NssaafInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssaafInfoWithDefaults() *NssaafInfo {
	this := NssaafInfo{}
	return &this
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *NssaafInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssaafInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *NssaafInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *NssaafInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *NssaafInfo) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		var ret []InternalGroupIdRange
		return ret
	}
	return o.InternalGroupIdentifiersRanges
}

// GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssaafInfo) GetInternalGroupIdentifiersRangesOk() ([]InternalGroupIdRange, bool) {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.InternalGroupIdentifiersRanges, true
}

// HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *NssaafInfo) HasInternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.InternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetInternalGroupIdentifiersRanges gets a reference to the given []InternalGroupIdRange and assigns it to the InternalGroupIdentifiersRanges field.
func (o *NssaafInfo) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange) {
	o.InternalGroupIdentifiersRanges = v
}

func (o NssaafInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NssaafInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.InternalGroupIdentifiersRanges) {
		toSerialize["internalGroupIdentifiersRanges"] = o.InternalGroupIdentifiersRanges
	}
	return toSerialize, nil
}

type NullableNssaafInfo struct {
	value *NssaafInfo
	isSet bool
}

func (v NullableNssaafInfo) Get() *NssaafInfo {
	return v.value
}

func (v *NullableNssaafInfo) Set(val *NssaafInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNssaafInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNssaafInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssaafInfo(val *NssaafInfo) *NullableNssaafInfo {
	return &NullableNssaafInfo{value: val, isSet: true}
}

func (v NullableNssaafInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssaafInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
