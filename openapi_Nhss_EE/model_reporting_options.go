/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
	"time"
)

// checks if the ReportingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingOptions{}

// ReportingOptions Contains the different reporting options associated to a given subscription created in HSS
type ReportingOptions struct {
	// Maximum number of events to be reported for events in a given subscription
	MaxNumOfReports *int32 `json:"maxNumOfReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// indicating a time in seconds.
	ReportPeriod *int32 `json:"reportPeriod,omitempty"`
}

// NewReportingOptions instantiates a new ReportingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingOptions() *ReportingOptions {
	this := ReportingOptions{}
	return &this
}

// NewReportingOptionsWithDefaults instantiates a new ReportingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingOptionsWithDefaults() *ReportingOptions {
	this := ReportingOptions{}
	return &this
}

// GetMaxNumOfReports returns the MaxNumOfReports field value if set, zero value otherwise.
func (o *ReportingOptions) GetMaxNumOfReports() int32 {
	if o == nil || IsNil(o.MaxNumOfReports) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfReports
}

// GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetMaxNumOfReportsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfReports) {
		return nil, false
	}
	return o.MaxNumOfReports, true
}

// HasMaxNumOfReports returns a boolean if a field has been set.
func (o *ReportingOptions) HasMaxNumOfReports() bool {
	if o != nil && !IsNil(o.MaxNumOfReports) {
		return true
	}

	return false
}

// SetMaxNumOfReports gets a reference to the given int32 and assigns it to the MaxNumOfReports field.
func (o *ReportingOptions) SetMaxNumOfReports(v int32) {
	o.MaxNumOfReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ReportingOptions) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ReportingOptions) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *ReportingOptions) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetReportPeriod returns the ReportPeriod field value if set, zero value otherwise.
func (o *ReportingOptions) GetReportPeriod() int32 {
	if o == nil || IsNil(o.ReportPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportPeriod
}

// GetReportPeriodOk returns a tuple with the ReportPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetReportPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportPeriod) {
		return nil, false
	}
	return o.ReportPeriod, true
}

// HasReportPeriod returns a boolean if a field has been set.
func (o *ReportingOptions) HasReportPeriod() bool {
	if o != nil && !IsNil(o.ReportPeriod) {
		return true
	}

	return false
}

// SetReportPeriod gets a reference to the given int32 and assigns it to the ReportPeriod field.
func (o *ReportingOptions) SetReportPeriod(v int32) {
	o.ReportPeriod = &v
}

func (o ReportingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxNumOfReports) {
		toSerialize["maxNumOfReports"] = o.MaxNumOfReports
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.ReportPeriod) {
		toSerialize["reportPeriod"] = o.ReportPeriod
	}
	return toSerialize, nil
}

type NullableReportingOptions struct {
	value *ReportingOptions
	isSet bool
}

func (v NullableReportingOptions) Get() *ReportingOptions {
	return v.value
}

func (v *NullableReportingOptions) Set(val *ReportingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingOptions(val *ReportingOptions) *NullableReportingOptions {
	return &NullableReportingOptions{value: val, isSet: true}
}

func (v NullableReportingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

