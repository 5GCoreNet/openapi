/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_Authentication

import (
	"encoding/json"
)

// checks if the SnssaiExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiExtension{}

// SnssaiExtension Extensions to the Snssai data type, sdRanges and wildcardSd shall not be present simultaneously 
type SnssaiExtension struct {
	// When present, it shall contain the range(s) of Slice Differentiator values supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type 
	SdRanges []SdRange `json:"sdRanges,omitempty"`
	// When present, it shall be set to true, to indicate that all SD values are supported for the Slice/Service Type value indicated in the sst attribute of the Snssai data type. 
	WildcardSd *bool `json:"wildcardSd,omitempty"`
}

// NewSnssaiExtension instantiates a new SnssaiExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiExtension() *SnssaiExtension {
	this := SnssaiExtension{}
	return &this
}

// NewSnssaiExtensionWithDefaults instantiates a new SnssaiExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiExtensionWithDefaults() *SnssaiExtension {
	this := SnssaiExtension{}
	return &this
}

// GetSdRanges returns the SdRanges field value if set, zero value otherwise.
func (o *SnssaiExtension) GetSdRanges() []SdRange {
	if o == nil || IsNil(o.SdRanges) {
		var ret []SdRange
		return ret
	}
	return o.SdRanges
}

// GetSdRangesOk returns a tuple with the SdRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiExtension) GetSdRangesOk() ([]SdRange, bool) {
	if o == nil || IsNil(o.SdRanges) {
		return nil, false
	}
	return o.SdRanges, true
}

// HasSdRanges returns a boolean if a field has been set.
func (o *SnssaiExtension) HasSdRanges() bool {
	if o != nil && !IsNil(o.SdRanges) {
		return true
	}

	return false
}

// SetSdRanges gets a reference to the given []SdRange and assigns it to the SdRanges field.
func (o *SnssaiExtension) SetSdRanges(v []SdRange) {
	o.SdRanges = v
}

// GetWildcardSd returns the WildcardSd field value if set, zero value otherwise.
func (o *SnssaiExtension) GetWildcardSd() bool {
	if o == nil || IsNil(o.WildcardSd) {
		var ret bool
		return ret
	}
	return *o.WildcardSd
}

// GetWildcardSdOk returns a tuple with the WildcardSd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiExtension) GetWildcardSdOk() (*bool, bool) {
	if o == nil || IsNil(o.WildcardSd) {
		return nil, false
	}
	return o.WildcardSd, true
}

// HasWildcardSd returns a boolean if a field has been set.
func (o *SnssaiExtension) HasWildcardSd() bool {
	if o != nil && !IsNil(o.WildcardSd) {
		return true
	}

	return false
}

// SetWildcardSd gets a reference to the given bool and assigns it to the WildcardSd field.
func (o *SnssaiExtension) SetWildcardSd(v bool) {
	o.WildcardSd = &v
}

func (o SnssaiExtension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SdRanges) {
		toSerialize["sdRanges"] = o.SdRanges
	}
	if !IsNil(o.WildcardSd) {
		toSerialize["wildcardSd"] = o.WildcardSd
	}
	return toSerialize, nil
}

type NullableSnssaiExtension struct {
	value *SnssaiExtension
	isSet bool
}

func (v NullableSnssaiExtension) Get() *SnssaiExtension {
	return v.value
}

func (v *NullableSnssaiExtension) Set(val *SnssaiExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiExtension(val *SnssaiExtension) *NullableSnssaiExtension {
	return &NullableSnssaiExtension{value: val, isSet: true}
}

func (v NullableSnssaiExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


