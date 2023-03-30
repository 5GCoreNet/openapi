/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// checks if the RadioParameterSetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadioParameterSetInfo{}

// RadioParameterSetInfo struct for RadioParameterSetInfo
type RadioParameterSetInfo struct {
	RadioParameterSetValues []string `json:"radioParameterSetValues,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTimestamp *time.Time `json:"changeTimestamp,omitempty"`
}

// NewRadioParameterSetInfo instantiates a new RadioParameterSetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadioParameterSetInfo() *RadioParameterSetInfo {
	this := RadioParameterSetInfo{}
	return &this
}

// NewRadioParameterSetInfoWithDefaults instantiates a new RadioParameterSetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadioParameterSetInfoWithDefaults() *RadioParameterSetInfo {
	this := RadioParameterSetInfo{}
	return &this
}

// GetRadioParameterSetValues returns the RadioParameterSetValues field value if set, zero value otherwise.
func (o *RadioParameterSetInfo) GetRadioParameterSetValues() []string {
	if o == nil || IsNil(o.RadioParameterSetValues) {
		var ret []string
		return ret
	}
	return o.RadioParameterSetValues
}

// GetRadioParameterSetValuesOk returns a tuple with the RadioParameterSetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioParameterSetInfo) GetRadioParameterSetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.RadioParameterSetValues) {
		return nil, false
	}
	return o.RadioParameterSetValues, true
}

// HasRadioParameterSetValues returns a boolean if a field has been set.
func (o *RadioParameterSetInfo) HasRadioParameterSetValues() bool {
	if o != nil && !IsNil(o.RadioParameterSetValues) {
		return true
	}

	return false
}

// SetRadioParameterSetValues gets a reference to the given []string and assigns it to the RadioParameterSetValues field.
func (o *RadioParameterSetInfo) SetRadioParameterSetValues(v []string) {
	o.RadioParameterSetValues = v
}

// GetChangeTimestamp returns the ChangeTimestamp field value if set, zero value otherwise.
func (o *RadioParameterSetInfo) GetChangeTimestamp() time.Time {
	if o == nil || IsNil(o.ChangeTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTimestamp
}

// GetChangeTimestampOk returns a tuple with the ChangeTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioParameterSetInfo) GetChangeTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ChangeTimestamp) {
		return nil, false
	}
	return o.ChangeTimestamp, true
}

// HasChangeTimestamp returns a boolean if a field has been set.
func (o *RadioParameterSetInfo) HasChangeTimestamp() bool {
	if o != nil && !IsNil(o.ChangeTimestamp) {
		return true
	}

	return false
}

// SetChangeTimestamp gets a reference to the given time.Time and assigns it to the ChangeTimestamp field.
func (o *RadioParameterSetInfo) SetChangeTimestamp(v time.Time) {
	o.ChangeTimestamp = &v
}

func (o RadioParameterSetInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadioParameterSetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RadioParameterSetValues) {
		toSerialize["radioParameterSetValues"] = o.RadioParameterSetValues
	}
	if !IsNil(o.ChangeTimestamp) {
		toSerialize["changeTimestamp"] = o.ChangeTimestamp
	}
	return toSerialize, nil
}

type NullableRadioParameterSetInfo struct {
	value *RadioParameterSetInfo
	isSet bool
}

func (v NullableRadioParameterSetInfo) Get() *RadioParameterSetInfo {
	return v.value
}

func (v *NullableRadioParameterSetInfo) Set(val *RadioParameterSetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioParameterSetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioParameterSetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioParameterSetInfo(val *RadioParameterSetInfo) *NullableRadioParameterSetInfo {
	return &NullableRadioParameterSetInfo{value: val, isSet: true}
}

func (v NullableRadioParameterSetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioParameterSetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

