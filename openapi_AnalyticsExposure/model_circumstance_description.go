/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"time"
)

// checks if the CircumstanceDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircumstanceDescription{}

// CircumstanceDescription Contains the description of a circumstance.
type CircumstanceDescription struct {
	// string with format 'float' as defined in OpenAPI.
	Freq *float32 `json:"freq,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Tm *time.Time `json:"tm,omitempty"`
	LocArea *NetworkAreaInfo `json:"locArea,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	Vol *int64 `json:"vol,omitempty"`
}

// NewCircumstanceDescription instantiates a new CircumstanceDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircumstanceDescription() *CircumstanceDescription {
	this := CircumstanceDescription{}
	return &this
}

// NewCircumstanceDescriptionWithDefaults instantiates a new CircumstanceDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircumstanceDescriptionWithDefaults() *CircumstanceDescription {
	this := CircumstanceDescription{}
	return &this
}

// GetFreq returns the Freq field value if set, zero value otherwise.
func (o *CircumstanceDescription) GetFreq() float32 {
	if o == nil || IsNil(o.Freq) {
		var ret float32
		return ret
	}
	return *o.Freq
}

// GetFreqOk returns a tuple with the Freq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircumstanceDescription) GetFreqOk() (*float32, bool) {
	if o == nil || IsNil(o.Freq) {
		return nil, false
	}
	return o.Freq, true
}

// HasFreq returns a boolean if a field has been set.
func (o *CircumstanceDescription) HasFreq() bool {
	if o != nil && !IsNil(o.Freq) {
		return true
	}

	return false
}

// SetFreq gets a reference to the given float32 and assigns it to the Freq field.
func (o *CircumstanceDescription) SetFreq(v float32) {
	o.Freq = &v
}

// GetTm returns the Tm field value if set, zero value otherwise.
func (o *CircumstanceDescription) GetTm() time.Time {
	if o == nil || IsNil(o.Tm) {
		var ret time.Time
		return ret
	}
	return *o.Tm
}

// GetTmOk returns a tuple with the Tm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircumstanceDescription) GetTmOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Tm) {
		return nil, false
	}
	return o.Tm, true
}

// HasTm returns a boolean if a field has been set.
func (o *CircumstanceDescription) HasTm() bool {
	if o != nil && !IsNil(o.Tm) {
		return true
	}

	return false
}

// SetTm gets a reference to the given time.Time and assigns it to the Tm field.
func (o *CircumstanceDescription) SetTm(v time.Time) {
	o.Tm = &v
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *CircumstanceDescription) GetLocArea() NetworkAreaInfo {
	if o == nil || IsNil(o.LocArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircumstanceDescription) GetLocAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.LocArea) {
		return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *CircumstanceDescription) HasLocArea() bool {
	if o != nil && !IsNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given NetworkAreaInfo and assigns it to the LocArea field.
func (o *CircumstanceDescription) SetLocArea(v NetworkAreaInfo) {
	o.LocArea = &v
}

// GetVol returns the Vol field value if set, zero value otherwise.
func (o *CircumstanceDescription) GetVol() int64 {
	if o == nil || IsNil(o.Vol) {
		var ret int64
		return ret
	}
	return *o.Vol
}

// GetVolOk returns a tuple with the Vol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircumstanceDescription) GetVolOk() (*int64, bool) {
	if o == nil || IsNil(o.Vol) {
		return nil, false
	}
	return o.Vol, true
}

// HasVol returns a boolean if a field has been set.
func (o *CircumstanceDescription) HasVol() bool {
	if o != nil && !IsNil(o.Vol) {
		return true
	}

	return false
}

// SetVol gets a reference to the given int64 and assigns it to the Vol field.
func (o *CircumstanceDescription) SetVol(v int64) {
	o.Vol = &v
}

func (o CircumstanceDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircumstanceDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Freq) {
		toSerialize["freq"] = o.Freq
	}
	if !IsNil(o.Tm) {
		toSerialize["tm"] = o.Tm
	}
	if !IsNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !IsNil(o.Vol) {
		toSerialize["vol"] = o.Vol
	}
	return toSerialize, nil
}

type NullableCircumstanceDescription struct {
	value *CircumstanceDescription
	isSet bool
}

func (v NullableCircumstanceDescription) Get() *CircumstanceDescription {
	return v.value
}

func (v *NullableCircumstanceDescription) Set(val *CircumstanceDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableCircumstanceDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableCircumstanceDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircumstanceDescription(val *CircumstanceDescription) *NullableCircumstanceDescription {
	return &NullableCircumstanceDescription{value: val, isSet: true}
}

func (v NullableCircumstanceDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircumstanceDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


