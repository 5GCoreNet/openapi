/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the ReportingCtrlOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingCtrlOneOf{}

// ReportingCtrlOneOf struct for ReportingCtrlOneOf
type ReportingCtrlOneOf struct {
	FileReportingPeriod *int32 `json:"fileReportingPeriod,omitempty"`
}

// NewReportingCtrlOneOf instantiates a new ReportingCtrlOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingCtrlOneOf() *ReportingCtrlOneOf {
	this := ReportingCtrlOneOf{}
	return &this
}

// NewReportingCtrlOneOfWithDefaults instantiates a new ReportingCtrlOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingCtrlOneOfWithDefaults() *ReportingCtrlOneOf {
	this := ReportingCtrlOneOf{}
	return &this
}

// GetFileReportingPeriod returns the FileReportingPeriod field value if set, zero value otherwise.
func (o *ReportingCtrlOneOf) GetFileReportingPeriod() int32 {
	if o == nil || IsNil(o.FileReportingPeriod) {
		var ret int32
		return ret
	}
	return *o.FileReportingPeriod
}

// GetFileReportingPeriodOk returns a tuple with the FileReportingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCtrlOneOf) GetFileReportingPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.FileReportingPeriod) {
		return nil, false
	}
	return o.FileReportingPeriod, true
}

// HasFileReportingPeriod returns a boolean if a field has been set.
func (o *ReportingCtrlOneOf) HasFileReportingPeriod() bool {
	if o != nil && !IsNil(o.FileReportingPeriod) {
		return true
	}

	return false
}

// SetFileReportingPeriod gets a reference to the given int32 and assigns it to the FileReportingPeriod field.
func (o *ReportingCtrlOneOf) SetFileReportingPeriod(v int32) {
	o.FileReportingPeriod = &v
}

func (o ReportingCtrlOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingCtrlOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileReportingPeriod) {
		toSerialize["fileReportingPeriod"] = o.FileReportingPeriod
	}
	return toSerialize, nil
}

type NullableReportingCtrlOneOf struct {
	value *ReportingCtrlOneOf
	isSet bool
}

func (v NullableReportingCtrlOneOf) Get() *ReportingCtrlOneOf {
	return v.value
}

func (v *NullableReportingCtrlOneOf) Set(val *ReportingCtrlOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingCtrlOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingCtrlOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingCtrlOneOf(val *ReportingCtrlOneOf) *NullableReportingCtrlOneOf {
	return &NullableReportingCtrlOneOf{value: val, isSet: true}
}

func (v NullableReportingCtrlOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingCtrlOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

