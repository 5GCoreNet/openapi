/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the ExtSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtSnssai{}

// ExtSnssai The sdRanges and wildcardSd attributes shall be exclusive from each other. If one of these attributes is present,  the sd attribute shall also be present and it shall contain one Slice Differentiator value within the range of SD  (if the sdRanges attribute is present) or with any value (if the wildcardSd attribute is present).
type ExtSnssai struct {
	Snssai
	// When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type
	SdRanges []SdRange `json:"sdRanges,omitempty"`
	// When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type.
	WildcardSd *bool `json:"wildcardSd,omitempty"`
}

// NewExtSnssai instantiates a new ExtSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtSnssai(sst int32) *ExtSnssai {
	this := ExtSnssai{}
	this.Sst = sst
	return &this
}

// NewExtSnssaiWithDefaults instantiates a new ExtSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtSnssaiWithDefaults() *ExtSnssai {
	this := ExtSnssai{}
	return &this
}

// GetSdRanges returns the SdRanges field value if set, zero value otherwise.
func (o *ExtSnssai) GetSdRanges() []SdRange {
	if o == nil || IsNil(o.SdRanges) {
		var ret []SdRange
		return ret
	}
	return o.SdRanges
}

// GetSdRangesOk returns a tuple with the SdRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtSnssai) GetSdRangesOk() ([]SdRange, bool) {
	if o == nil || IsNil(o.SdRanges) {
		return nil, false
	}
	return o.SdRanges, true
}

// HasSdRanges returns a boolean if a field has been set.
func (o *ExtSnssai) HasSdRanges() bool {
	if o != nil && !IsNil(o.SdRanges) {
		return true
	}

	return false
}

// SetSdRanges gets a reference to the given []SdRange and assigns it to the SdRanges field.
func (o *ExtSnssai) SetSdRanges(v []SdRange) {
	o.SdRanges = v
}

// GetWildcardSd returns the WildcardSd field value if set, zero value otherwise.
func (o *ExtSnssai) GetWildcardSd() bool {
	if o == nil || IsNil(o.WildcardSd) {
		var ret bool
		return ret
	}
	return *o.WildcardSd
}

// GetWildcardSdOk returns a tuple with the WildcardSd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtSnssai) GetWildcardSdOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardSd) {
		return nil, false
	}
	return o.WildcardSd, true
}

// HasWildcardSd returns a boolean if a field has been set.
func (o *ExtSnssai) HasWildcardSd() bool {
	if o != nil && !IsNil(o.WildcardSd) {
		return true
	}

	return false
}

// SetWildcardSd gets a reference to the given bool and assigns it to the WildcardSd field.
func (o *ExtSnssai) SetWildcardSd(v bool) {
	o.WildcardSd = &v
}

func (o ExtSnssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSnssai, errSnssai := json.Marshal(o.Snssai)
	if errSnssai != nil {
		return map[string]interface{}{}, errSnssai
	}
	errSnssai = json.Unmarshal([]byte(serializedSnssai), &toSerialize)
	if errSnssai != nil {
		return map[string]interface{}{}, errSnssai
	}
	if !IsNil(o.SdRanges) {
		toSerialize["sdRanges"] = o.SdRanges
	}
	if !IsNil(o.WildcardSd) {
		toSerialize["wildcardSd"] = o.WildcardSd
	}
	return toSerialize, nil
}

type NullableExtSnssai struct {
	value *ExtSnssai
	isSet bool
}

func (v NullableExtSnssai) Get() *ExtSnssai {
	return v.value
}

func (v *NullableExtSnssai) Set(val *ExtSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableExtSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableExtSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtSnssai(val *ExtSnssai) *NullableExtSnssai {
	return &NullableExtSnssai{value: val, isSet: true}
}

func (v NullableExtSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
