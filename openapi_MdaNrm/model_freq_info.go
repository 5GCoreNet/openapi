/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the FreqInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreqInfo{}

// FreqInfo specifies the carrier frequency and bands used in a cell.
type FreqInfo struct {
	Arfcn     *int32  `json:"arfcn,omitempty"`
	FreqBands []int32 `json:"freqBands,omitempty"`
}

// NewFreqInfo instantiates a new FreqInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreqInfo() *FreqInfo {
	this := FreqInfo{}
	return &this
}

// NewFreqInfoWithDefaults instantiates a new FreqInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreqInfoWithDefaults() *FreqInfo {
	this := FreqInfo{}
	return &this
}

// GetArfcn returns the Arfcn field value if set, zero value otherwise.
func (o *FreqInfo) GetArfcn() int32 {
	if o == nil || IsNil(o.Arfcn) {
		var ret int32
		return ret
	}
	return *o.Arfcn
}

// GetArfcnOk returns a tuple with the Arfcn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreqInfo) GetArfcnOk() (*int32, bool) {
	if o == nil || IsNil(o.Arfcn) {
		return nil, false
	}
	return o.Arfcn, true
}

// HasArfcn returns a boolean if a field has been set.
func (o *FreqInfo) HasArfcn() bool {
	if o != nil && !IsNil(o.Arfcn) {
		return true
	}

	return false
}

// SetArfcn gets a reference to the given int32 and assigns it to the Arfcn field.
func (o *FreqInfo) SetArfcn(v int32) {
	o.Arfcn = &v
}

// GetFreqBands returns the FreqBands field value if set, zero value otherwise.
func (o *FreqInfo) GetFreqBands() []int32 {
	if o == nil || IsNil(o.FreqBands) {
		var ret []int32
		return ret
	}
	return o.FreqBands
}

// GetFreqBandsOk returns a tuple with the FreqBands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreqInfo) GetFreqBandsOk() ([]int32, bool) {
	if o == nil || IsNil(o.FreqBands) {
		return nil, false
	}
	return o.FreqBands, true
}

// HasFreqBands returns a boolean if a field has been set.
func (o *FreqInfo) HasFreqBands() bool {
	if o != nil && !IsNil(o.FreqBands) {
		return true
	}

	return false
}

// SetFreqBands gets a reference to the given []int32 and assigns it to the FreqBands field.
func (o *FreqInfo) SetFreqBands(v []int32) {
	o.FreqBands = v
}

func (o FreqInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreqInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Arfcn) {
		toSerialize["arfcn"] = o.Arfcn
	}
	if !IsNil(o.FreqBands) {
		toSerialize["freqBands"] = o.FreqBands
	}
	return toSerialize, nil
}

type NullableFreqInfo struct {
	value *FreqInfo
	isSet bool
}

func (v NullableFreqInfo) Get() *FreqInfo {
	return v.value
}

func (v *NullableFreqInfo) Set(val *FreqInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFreqInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFreqInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreqInfo(val *FreqInfo) *NullableFreqInfo {
	return &NullableFreqInfo{value: val, isSet: true}
}

func (v NullableFreqInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreqInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
