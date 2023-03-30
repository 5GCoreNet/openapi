/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the FormattingInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormattingInstruction{}

// FormattingInstruction Contains data or analytics formatting instructions.
type FormattingInstruction struct {
	// Indicates that notifications shall be buffered until the NF service consumer requests their delivery. 
	ConsTrigNotif *bool `json:"consTrigNotif,omitempty"`
	ReportingOptions *ReportingOptions `json:"reportingOptions,omitempty"`
}

// NewFormattingInstruction instantiates a new FormattingInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormattingInstruction() *FormattingInstruction {
	this := FormattingInstruction{}
	return &this
}

// NewFormattingInstructionWithDefaults instantiates a new FormattingInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormattingInstructionWithDefaults() *FormattingInstruction {
	this := FormattingInstruction{}
	return &this
}

// GetConsTrigNotif returns the ConsTrigNotif field value if set, zero value otherwise.
func (o *FormattingInstruction) GetConsTrigNotif() bool {
	if o == nil || IsNil(o.ConsTrigNotif) {
		var ret bool
		return ret
	}
	return *o.ConsTrigNotif
}

// GetConsTrigNotifOk returns a tuple with the ConsTrigNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattingInstruction) GetConsTrigNotifOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsTrigNotif) {
		return nil, false
	}
	return o.ConsTrigNotif, true
}

// HasConsTrigNotif returns a boolean if a field has been set.
func (o *FormattingInstruction) HasConsTrigNotif() bool {
	if o != nil && !IsNil(o.ConsTrigNotif) {
		return true
	}

	return false
}

// SetConsTrigNotif gets a reference to the given bool and assigns it to the ConsTrigNotif field.
func (o *FormattingInstruction) SetConsTrigNotif(v bool) {
	o.ConsTrigNotif = &v
}

// GetReportingOptions returns the ReportingOptions field value if set, zero value otherwise.
func (o *FormattingInstruction) GetReportingOptions() ReportingOptions {
	if o == nil || IsNil(o.ReportingOptions) {
		var ret ReportingOptions
		return ret
	}
	return *o.ReportingOptions
}

// GetReportingOptionsOk returns a tuple with the ReportingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattingInstruction) GetReportingOptionsOk() (*ReportingOptions, bool) {
	if o == nil || IsNil(o.ReportingOptions) {
		return nil, false
	}
	return o.ReportingOptions, true
}

// HasReportingOptions returns a boolean if a field has been set.
func (o *FormattingInstruction) HasReportingOptions() bool {
	if o != nil && !IsNil(o.ReportingOptions) {
		return true
	}

	return false
}

// SetReportingOptions gets a reference to the given ReportingOptions and assigns it to the ReportingOptions field.
func (o *FormattingInstruction) SetReportingOptions(v ReportingOptions) {
	o.ReportingOptions = &v
}

func (o FormattingInstruction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormattingInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsTrigNotif) {
		toSerialize["consTrigNotif"] = o.ConsTrigNotif
	}
	if !IsNil(o.ReportingOptions) {
		toSerialize["reportingOptions"] = o.ReportingOptions
	}
	return toSerialize, nil
}

type NullableFormattingInstruction struct {
	value *FormattingInstruction
	isSet bool
}

func (v NullableFormattingInstruction) Get() *FormattingInstruction {
	return v.value
}

func (v *NullableFormattingInstruction) Set(val *FormattingInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableFormattingInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableFormattingInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormattingInstruction(val *FormattingInstruction) *NullableFormattingInstruction {
	return &NullableFormattingInstruction{value: val, isSet: true}
}

func (v NullableFormattingInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormattingInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


