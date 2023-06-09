/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
)

// checks if the Report type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Report{}

// Report Contains data for a given Monitoring Event Report
type Report struct {
	ReachabilityForSmsReport  *ReachabilityForSmsReport  `json:"reachabilityForSmsReport,omitempty"`
	ReachabilityForDataReport *ReachabilityForDataReport `json:"reachabilityForDataReport,omitempty"`
	LossConnectivityReport    *LossConnectivityReport    `json:"lossConnectivityReport,omitempty"`
	LocationReport            *LocationReport            `json:"locationReport,omitempty"`
	PdnConnectivityStatReport *PdnConnectivityStatReport `json:"pdnConnectivityStatReport,omitempty"`
}

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport() *Report {
	this := Report{}
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}
	return &this
}

// GetReachabilityForSmsReport returns the ReachabilityForSmsReport field value if set, zero value otherwise.
func (o *Report) GetReachabilityForSmsReport() ReachabilityForSmsReport {
	if o == nil || IsNil(o.ReachabilityForSmsReport) {
		var ret ReachabilityForSmsReport
		return ret
	}
	return *o.ReachabilityForSmsReport
}

// GetReachabilityForSmsReportOk returns a tuple with the ReachabilityForSmsReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetReachabilityForSmsReportOk() (*ReachabilityForSmsReport, bool) {
	if o == nil || IsNil(o.ReachabilityForSmsReport) {
		return nil, false
	}
	return o.ReachabilityForSmsReport, true
}

// HasReachabilityForSmsReport returns a boolean if a field has been set.
func (o *Report) HasReachabilityForSmsReport() bool {
	if o != nil && !IsNil(o.ReachabilityForSmsReport) {
		return true
	}

	return false
}

// SetReachabilityForSmsReport gets a reference to the given ReachabilityForSmsReport and assigns it to the ReachabilityForSmsReport field.
func (o *Report) SetReachabilityForSmsReport(v ReachabilityForSmsReport) {
	o.ReachabilityForSmsReport = &v
}

// GetReachabilityForDataReport returns the ReachabilityForDataReport field value if set, zero value otherwise.
func (o *Report) GetReachabilityForDataReport() ReachabilityForDataReport {
	if o == nil || IsNil(o.ReachabilityForDataReport) {
		var ret ReachabilityForDataReport
		return ret
	}
	return *o.ReachabilityForDataReport
}

// GetReachabilityForDataReportOk returns a tuple with the ReachabilityForDataReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetReachabilityForDataReportOk() (*ReachabilityForDataReport, bool) {
	if o == nil || IsNil(o.ReachabilityForDataReport) {
		return nil, false
	}
	return o.ReachabilityForDataReport, true
}

// HasReachabilityForDataReport returns a boolean if a field has been set.
func (o *Report) HasReachabilityForDataReport() bool {
	if o != nil && !IsNil(o.ReachabilityForDataReport) {
		return true
	}

	return false
}

// SetReachabilityForDataReport gets a reference to the given ReachabilityForDataReport and assigns it to the ReachabilityForDataReport field.
func (o *Report) SetReachabilityForDataReport(v ReachabilityForDataReport) {
	o.ReachabilityForDataReport = &v
}

// GetLossConnectivityReport returns the LossConnectivityReport field value if set, zero value otherwise.
func (o *Report) GetLossConnectivityReport() LossConnectivityReport {
	if o == nil || IsNil(o.LossConnectivityReport) {
		var ret LossConnectivityReport
		return ret
	}
	return *o.LossConnectivityReport
}

// GetLossConnectivityReportOk returns a tuple with the LossConnectivityReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetLossConnectivityReportOk() (*LossConnectivityReport, bool) {
	if o == nil || IsNil(o.LossConnectivityReport) {
		return nil, false
	}
	return o.LossConnectivityReport, true
}

// HasLossConnectivityReport returns a boolean if a field has been set.
func (o *Report) HasLossConnectivityReport() bool {
	if o != nil && !IsNil(o.LossConnectivityReport) {
		return true
	}

	return false
}

// SetLossConnectivityReport gets a reference to the given LossConnectivityReport and assigns it to the LossConnectivityReport field.
func (o *Report) SetLossConnectivityReport(v LossConnectivityReport) {
	o.LossConnectivityReport = &v
}

// GetLocationReport returns the LocationReport field value if set, zero value otherwise.
func (o *Report) GetLocationReport() LocationReport {
	if o == nil || IsNil(o.LocationReport) {
		var ret LocationReport
		return ret
	}
	return *o.LocationReport
}

// GetLocationReportOk returns a tuple with the LocationReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetLocationReportOk() (*LocationReport, bool) {
	if o == nil || IsNil(o.LocationReport) {
		return nil, false
	}
	return o.LocationReport, true
}

// HasLocationReport returns a boolean if a field has been set.
func (o *Report) HasLocationReport() bool {
	if o != nil && !IsNil(o.LocationReport) {
		return true
	}

	return false
}

// SetLocationReport gets a reference to the given LocationReport and assigns it to the LocationReport field.
func (o *Report) SetLocationReport(v LocationReport) {
	o.LocationReport = &v
}

// GetPdnConnectivityStatReport returns the PdnConnectivityStatReport field value if set, zero value otherwise.
func (o *Report) GetPdnConnectivityStatReport() PdnConnectivityStatReport {
	if o == nil || IsNil(o.PdnConnectivityStatReport) {
		var ret PdnConnectivityStatReport
		return ret
	}
	return *o.PdnConnectivityStatReport
}

// GetPdnConnectivityStatReportOk returns a tuple with the PdnConnectivityStatReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetPdnConnectivityStatReportOk() (*PdnConnectivityStatReport, bool) {
	if o == nil || IsNil(o.PdnConnectivityStatReport) {
		return nil, false
	}
	return o.PdnConnectivityStatReport, true
}

// HasPdnConnectivityStatReport returns a boolean if a field has been set.
func (o *Report) HasPdnConnectivityStatReport() bool {
	if o != nil && !IsNil(o.PdnConnectivityStatReport) {
		return true
	}

	return false
}

// SetPdnConnectivityStatReport gets a reference to the given PdnConnectivityStatReport and assigns it to the PdnConnectivityStatReport field.
func (o *Report) SetPdnConnectivityStatReport(v PdnConnectivityStatReport) {
	o.PdnConnectivityStatReport = &v
}

func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Report) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReachabilityForSmsReport) {
		toSerialize["reachabilityForSmsReport"] = o.ReachabilityForSmsReport
	}
	if !IsNil(o.ReachabilityForDataReport) {
		toSerialize["reachabilityForDataReport"] = o.ReachabilityForDataReport
	}
	if !IsNil(o.LossConnectivityReport) {
		toSerialize["lossConnectivityReport"] = o.LossConnectivityReport
	}
	if !IsNil(o.LocationReport) {
		toSerialize["locationReport"] = o.LocationReport
	}
	if !IsNil(o.PdnConnectivityStatReport) {
		toSerialize["pdnConnectivityStatReport"] = o.PdnConnectivityStatReport
	}
	return toSerialize, nil
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
