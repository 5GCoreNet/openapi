/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the SupportedPerfMetricGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedPerfMetricGroup{}

// SupportedPerfMetricGroup struct for SupportedPerfMetricGroup
type SupportedPerfMetricGroup struct {
	PerformanceMetrics        []string `json:"performanceMetrics,omitempty"`
	GranularityPeriods        []int32  `json:"granularityPeriods,omitempty"`
	ReportingMethods          []string `json:"reportingMethods,omitempty"`
	MonitorGranularityPeriods []int32  `json:"monitorGranularityPeriods,omitempty"`
}

// NewSupportedPerfMetricGroup instantiates a new SupportedPerfMetricGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedPerfMetricGroup() *SupportedPerfMetricGroup {
	this := SupportedPerfMetricGroup{}
	return &this
}

// NewSupportedPerfMetricGroupWithDefaults instantiates a new SupportedPerfMetricGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedPerfMetricGroupWithDefaults() *SupportedPerfMetricGroup {
	this := SupportedPerfMetricGroup{}
	return &this
}

// GetPerformanceMetrics returns the PerformanceMetrics field value if set, zero value otherwise.
func (o *SupportedPerfMetricGroup) GetPerformanceMetrics() []string {
	if o == nil || IsNil(o.PerformanceMetrics) {
		var ret []string
		return ret
	}
	return o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPerfMetricGroup) GetPerformanceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.PerformanceMetrics) {
		return nil, false
	}
	return o.PerformanceMetrics, true
}

// HasPerformanceMetrics returns a boolean if a field has been set.
func (o *SupportedPerfMetricGroup) HasPerformanceMetrics() bool {
	if o != nil && !IsNil(o.PerformanceMetrics) {
		return true
	}

	return false
}

// SetPerformanceMetrics gets a reference to the given []string and assigns it to the PerformanceMetrics field.
func (o *SupportedPerfMetricGroup) SetPerformanceMetrics(v []string) {
	o.PerformanceMetrics = v
}

// GetGranularityPeriods returns the GranularityPeriods field value if set, zero value otherwise.
func (o *SupportedPerfMetricGroup) GetGranularityPeriods() []int32 {
	if o == nil || IsNil(o.GranularityPeriods) {
		var ret []int32
		return ret
	}
	return o.GranularityPeriods
}

// GetGranularityPeriodsOk returns a tuple with the GranularityPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPerfMetricGroup) GetGranularityPeriodsOk() ([]int32, bool) {
	if o == nil || IsNil(o.GranularityPeriods) {
		return nil, false
	}
	return o.GranularityPeriods, true
}

// HasGranularityPeriods returns a boolean if a field has been set.
func (o *SupportedPerfMetricGroup) HasGranularityPeriods() bool {
	if o != nil && !IsNil(o.GranularityPeriods) {
		return true
	}

	return false
}

// SetGranularityPeriods gets a reference to the given []int32 and assigns it to the GranularityPeriods field.
func (o *SupportedPerfMetricGroup) SetGranularityPeriods(v []int32) {
	o.GranularityPeriods = v
}

// GetReportingMethods returns the ReportingMethods field value if set, zero value otherwise.
func (o *SupportedPerfMetricGroup) GetReportingMethods() []string {
	if o == nil || IsNil(o.ReportingMethods) {
		var ret []string
		return ret
	}
	return o.ReportingMethods
}

// GetReportingMethodsOk returns a tuple with the ReportingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPerfMetricGroup) GetReportingMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReportingMethods) {
		return nil, false
	}
	return o.ReportingMethods, true
}

// HasReportingMethods returns a boolean if a field has been set.
func (o *SupportedPerfMetricGroup) HasReportingMethods() bool {
	if o != nil && !IsNil(o.ReportingMethods) {
		return true
	}

	return false
}

// SetReportingMethods gets a reference to the given []string and assigns it to the ReportingMethods field.
func (o *SupportedPerfMetricGroup) SetReportingMethods(v []string) {
	o.ReportingMethods = v
}

// GetMonitorGranularityPeriods returns the MonitorGranularityPeriods field value if set, zero value otherwise.
func (o *SupportedPerfMetricGroup) GetMonitorGranularityPeriods() []int32 {
	if o == nil || IsNil(o.MonitorGranularityPeriods) {
		var ret []int32
		return ret
	}
	return o.MonitorGranularityPeriods
}

// GetMonitorGranularityPeriodsOk returns a tuple with the MonitorGranularityPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPerfMetricGroup) GetMonitorGranularityPeriodsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonitorGranularityPeriods) {
		return nil, false
	}
	return o.MonitorGranularityPeriods, true
}

// HasMonitorGranularityPeriods returns a boolean if a field has been set.
func (o *SupportedPerfMetricGroup) HasMonitorGranularityPeriods() bool {
	if o != nil && !IsNil(o.MonitorGranularityPeriods) {
		return true
	}

	return false
}

// SetMonitorGranularityPeriods gets a reference to the given []int32 and assigns it to the MonitorGranularityPeriods field.
func (o *SupportedPerfMetricGroup) SetMonitorGranularityPeriods(v []int32) {
	o.MonitorGranularityPeriods = v
}

func (o SupportedPerfMetricGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedPerfMetricGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PerformanceMetrics) {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	if !IsNil(o.GranularityPeriods) {
		toSerialize["granularityPeriods"] = o.GranularityPeriods
	}
	if !IsNil(o.ReportingMethods) {
		toSerialize["reportingMethods"] = o.ReportingMethods
	}
	if !IsNil(o.MonitorGranularityPeriods) {
		toSerialize["monitorGranularityPeriods"] = o.MonitorGranularityPeriods
	}
	return toSerialize, nil
}

type NullableSupportedPerfMetricGroup struct {
	value *SupportedPerfMetricGroup
	isSet bool
}

func (v NullableSupportedPerfMetricGroup) Get() *SupportedPerfMetricGroup {
	return v.value
}

func (v *NullableSupportedPerfMetricGroup) Set(val *SupportedPerfMetricGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedPerfMetricGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedPerfMetricGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedPerfMetricGroup(val *SupportedPerfMetricGroup) *NullableSupportedPerfMetricGroup {
	return &NullableSupportedPerfMetricGroup{value: val, isSet: true}
}

func (v NullableSupportedPerfMetricGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedPerfMetricGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
