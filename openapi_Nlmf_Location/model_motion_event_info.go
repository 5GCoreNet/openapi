/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
)

// checks if the MotionEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MotionEventInfo{}

// MotionEventInfo Indicates the information of motion based event reporting.
type MotionEventInfo struct {
	// Minimum straight line distance moved by a UE to trigger a motion event report.
	LinearDistance int32 `json:"linearDistance"`
	OccurrenceInfo *OccurrenceInfo `json:"occurrenceInfo,omitempty"`
	// Minimum interval between event reports.
	MinimumInterval *int32 `json:"minimumInterval,omitempty"`
	// Maximum interval between event reports.
	MaximumInterval *int32 `json:"maximumInterval,omitempty"`
	// Maximum time interval between consecutive evaluations by a UE of a trigger event.
	SamplingInterval *int32 `json:"samplingInterval,omitempty"`
	// Maximum duration of event reporting.
	ReportingDuration *int32 `json:"reportingDuration,omitempty"`
	ReportingLocationReq *bool `json:"reportingLocationReq,omitempty"`
}

// NewMotionEventInfo instantiates a new MotionEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotionEventInfo(linearDistance int32) *MotionEventInfo {
	this := MotionEventInfo{}
	this.LinearDistance = linearDistance
	var reportingLocationReq bool = true
	this.ReportingLocationReq = &reportingLocationReq
	return &this
}

// NewMotionEventInfoWithDefaults instantiates a new MotionEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotionEventInfoWithDefaults() *MotionEventInfo {
	this := MotionEventInfo{}
	var reportingLocationReq bool = true
	this.ReportingLocationReq = &reportingLocationReq
	return &this
}

// GetLinearDistance returns the LinearDistance field value
func (o *MotionEventInfo) GetLinearDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LinearDistance
}

// GetLinearDistanceOk returns a tuple with the LinearDistance field value
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetLinearDistanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinearDistance, true
}

// SetLinearDistance sets field value
func (o *MotionEventInfo) SetLinearDistance(v int32) {
	o.LinearDistance = v
}

// GetOccurrenceInfo returns the OccurrenceInfo field value if set, zero value otherwise.
func (o *MotionEventInfo) GetOccurrenceInfo() OccurrenceInfo {
	if o == nil || IsNil(o.OccurrenceInfo) {
		var ret OccurrenceInfo
		return ret
	}
	return *o.OccurrenceInfo
}

// GetOccurrenceInfoOk returns a tuple with the OccurrenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetOccurrenceInfoOk() (*OccurrenceInfo, bool) {
	if o == nil || IsNil(o.OccurrenceInfo) {
		return nil, false
	}
	return o.OccurrenceInfo, true
}

// HasOccurrenceInfo returns a boolean if a field has been set.
func (o *MotionEventInfo) HasOccurrenceInfo() bool {
	if o != nil && !IsNil(o.OccurrenceInfo) {
		return true
	}

	return false
}

// SetOccurrenceInfo gets a reference to the given OccurrenceInfo and assigns it to the OccurrenceInfo field.
func (o *MotionEventInfo) SetOccurrenceInfo(v OccurrenceInfo) {
	o.OccurrenceInfo = &v
}

// GetMinimumInterval returns the MinimumInterval field value if set, zero value otherwise.
func (o *MotionEventInfo) GetMinimumInterval() int32 {
	if o == nil || IsNil(o.MinimumInterval) {
		var ret int32
		return ret
	}
	return *o.MinimumInterval
}

// GetMinimumIntervalOk returns a tuple with the MinimumInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetMinimumIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumInterval) {
		return nil, false
	}
	return o.MinimumInterval, true
}

// HasMinimumInterval returns a boolean if a field has been set.
func (o *MotionEventInfo) HasMinimumInterval() bool {
	if o != nil && !IsNil(o.MinimumInterval) {
		return true
	}

	return false
}

// SetMinimumInterval gets a reference to the given int32 and assigns it to the MinimumInterval field.
func (o *MotionEventInfo) SetMinimumInterval(v int32) {
	o.MinimumInterval = &v
}

// GetMaximumInterval returns the MaximumInterval field value if set, zero value otherwise.
func (o *MotionEventInfo) GetMaximumInterval() int32 {
	if o == nil || IsNil(o.MaximumInterval) {
		var ret int32
		return ret
	}
	return *o.MaximumInterval
}

// GetMaximumIntervalOk returns a tuple with the MaximumInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetMaximumIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumInterval) {
		return nil, false
	}
	return o.MaximumInterval, true
}

// HasMaximumInterval returns a boolean if a field has been set.
func (o *MotionEventInfo) HasMaximumInterval() bool {
	if o != nil && !IsNil(o.MaximumInterval) {
		return true
	}

	return false
}

// SetMaximumInterval gets a reference to the given int32 and assigns it to the MaximumInterval field.
func (o *MotionEventInfo) SetMaximumInterval(v int32) {
	o.MaximumInterval = &v
}

// GetSamplingInterval returns the SamplingInterval field value if set, zero value otherwise.
func (o *MotionEventInfo) GetSamplingInterval() int32 {
	if o == nil || IsNil(o.SamplingInterval) {
		var ret int32
		return ret
	}
	return *o.SamplingInterval
}

// GetSamplingIntervalOk returns a tuple with the SamplingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetSamplingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.SamplingInterval) {
		return nil, false
	}
	return o.SamplingInterval, true
}

// HasSamplingInterval returns a boolean if a field has been set.
func (o *MotionEventInfo) HasSamplingInterval() bool {
	if o != nil && !IsNil(o.SamplingInterval) {
		return true
	}

	return false
}

// SetSamplingInterval gets a reference to the given int32 and assigns it to the SamplingInterval field.
func (o *MotionEventInfo) SetSamplingInterval(v int32) {
	o.SamplingInterval = &v
}

// GetReportingDuration returns the ReportingDuration field value if set, zero value otherwise.
func (o *MotionEventInfo) GetReportingDuration() int32 {
	if o == nil || IsNil(o.ReportingDuration) {
		var ret int32
		return ret
	}
	return *o.ReportingDuration
}

// GetReportingDurationOk returns a tuple with the ReportingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetReportingDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportingDuration) {
		return nil, false
	}
	return o.ReportingDuration, true
}

// HasReportingDuration returns a boolean if a field has been set.
func (o *MotionEventInfo) HasReportingDuration() bool {
	if o != nil && !IsNil(o.ReportingDuration) {
		return true
	}

	return false
}

// SetReportingDuration gets a reference to the given int32 and assigns it to the ReportingDuration field.
func (o *MotionEventInfo) SetReportingDuration(v int32) {
	o.ReportingDuration = &v
}

// GetReportingLocationReq returns the ReportingLocationReq field value if set, zero value otherwise.
func (o *MotionEventInfo) GetReportingLocationReq() bool {
	if o == nil || IsNil(o.ReportingLocationReq) {
		var ret bool
		return ret
	}
	return *o.ReportingLocationReq
}

// GetReportingLocationReqOk returns a tuple with the ReportingLocationReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotionEventInfo) GetReportingLocationReqOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportingLocationReq) {
		return nil, false
	}
	return o.ReportingLocationReq, true
}

// HasReportingLocationReq returns a boolean if a field has been set.
func (o *MotionEventInfo) HasReportingLocationReq() bool {
	if o != nil && !IsNil(o.ReportingLocationReq) {
		return true
	}

	return false
}

// SetReportingLocationReq gets a reference to the given bool and assigns it to the ReportingLocationReq field.
func (o *MotionEventInfo) SetReportingLocationReq(v bool) {
	o.ReportingLocationReq = &v
}

func (o MotionEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MotionEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linearDistance"] = o.LinearDistance
	if !IsNil(o.OccurrenceInfo) {
		toSerialize["occurrenceInfo"] = o.OccurrenceInfo
	}
	if !IsNil(o.MinimumInterval) {
		toSerialize["minimumInterval"] = o.MinimumInterval
	}
	if !IsNil(o.MaximumInterval) {
		toSerialize["maximumInterval"] = o.MaximumInterval
	}
	if !IsNil(o.SamplingInterval) {
		toSerialize["samplingInterval"] = o.SamplingInterval
	}
	if !IsNil(o.ReportingDuration) {
		toSerialize["reportingDuration"] = o.ReportingDuration
	}
	if !IsNil(o.ReportingLocationReq) {
		toSerialize["reportingLocationReq"] = o.ReportingLocationReq
	}
	return toSerialize, nil
}

type NullableMotionEventInfo struct {
	value *MotionEventInfo
	isSet bool
}

func (v NullableMotionEventInfo) Get() *MotionEventInfo {
	return v.value
}

func (v *NullableMotionEventInfo) Set(val *MotionEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMotionEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMotionEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotionEventInfo(val *MotionEventInfo) *NullableMotionEventInfo {
	return &NullableMotionEventInfo{value: val, isSet: true}
}

func (v NullableMotionEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotionEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

