/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the PatchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchResult{}

// PatchResult The execution report result on failed modification.
type PatchResult struct {
	// The execution report contains an array of report items. Each report item indicates one  failed modification. 
	Report []ReportItem `json:"report"`
}

// NewPatchResult instantiates a new PatchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchResult(report []ReportItem) *PatchResult {
	this := PatchResult{}
	this.Report = report
	return &this
}

// NewPatchResultWithDefaults instantiates a new PatchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchResultWithDefaults() *PatchResult {
	this := PatchResult{}
	return &this
}

// GetReport returns the Report field value
func (o *PatchResult) GetReport() []ReportItem {
	if o == nil {
		var ret []ReportItem
		return ret
	}

	return o.Report
}

// GetReportOk returns a tuple with the Report field value
// and a boolean to check if the value has been set.
func (o *PatchResult) GetReportOk() ([]ReportItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Report, true
}

// SetReport sets field value
func (o *PatchResult) SetReport(v []ReportItem) {
	o.Report = v
}

func (o PatchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["report"] = o.Report
	return toSerialize, nil
}

type NullablePatchResult struct {
	value *PatchResult
	isSet bool
}

func (v NullablePatchResult) Get() *PatchResult {
	return v.value
}

func (v *NullablePatchResult) Set(val *PatchResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchResult(val *PatchResult) *NullablePatchResult {
	return &NullablePatchResult{value: val, isSet: true}
}

func (v NullablePatchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


