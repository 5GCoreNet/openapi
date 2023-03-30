/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// checks if the AreaEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaEventInfo{}

// AreaEventInfo Indicates the information of area based event reporting.
type AreaEventInfo struct {
	AreaDefinition []ReportingArea `json:"areaDefinition"`
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

// NewAreaEventInfo instantiates a new AreaEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaEventInfo(areaDefinition []ReportingArea) *AreaEventInfo {
	this := AreaEventInfo{}
	this.AreaDefinition = areaDefinition
	var reportingLocationReq bool = true
	this.ReportingLocationReq = &reportingLocationReq
	return &this
}

// NewAreaEventInfoWithDefaults instantiates a new AreaEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaEventInfoWithDefaults() *AreaEventInfo {
	this := AreaEventInfo{}
	var reportingLocationReq bool = true
	this.ReportingLocationReq = &reportingLocationReq
	return &this
}

// GetAreaDefinition returns the AreaDefinition field value
func (o *AreaEventInfo) GetAreaDefinition() []ReportingArea {
	if o == nil {
		var ret []ReportingArea
		return ret
	}

	return o.AreaDefinition
}

// GetAreaDefinitionOk returns a tuple with the AreaDefinition field value
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetAreaDefinitionOk() ([]ReportingArea, bool) {
	if o == nil {
		return nil, false
	}
	return o.AreaDefinition, true
}

// SetAreaDefinition sets field value
func (o *AreaEventInfo) SetAreaDefinition(v []ReportingArea) {
	o.AreaDefinition = v
}

// GetOccurrenceInfo returns the OccurrenceInfo field value if set, zero value otherwise.
func (o *AreaEventInfo) GetOccurrenceInfo() OccurrenceInfo {
	if o == nil || IsNil(o.OccurrenceInfo) {
		var ret OccurrenceInfo
		return ret
	}
	return *o.OccurrenceInfo
}

// GetOccurrenceInfoOk returns a tuple with the OccurrenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetOccurrenceInfoOk() (*OccurrenceInfo, bool) {
	if o == nil || IsNil(o.OccurrenceInfo) {
		return nil, false
	}
	return o.OccurrenceInfo, true
}

// HasOccurrenceInfo returns a boolean if a field has been set.
func (o *AreaEventInfo) HasOccurrenceInfo() bool {
	if o != nil && !IsNil(o.OccurrenceInfo) {
		return true
	}

	return false
}

// SetOccurrenceInfo gets a reference to the given OccurrenceInfo and assigns it to the OccurrenceInfo field.
func (o *AreaEventInfo) SetOccurrenceInfo(v OccurrenceInfo) {
	o.OccurrenceInfo = &v
}

// GetMinimumInterval returns the MinimumInterval field value if set, zero value otherwise.
func (o *AreaEventInfo) GetMinimumInterval() int32 {
	if o == nil || IsNil(o.MinimumInterval) {
		var ret int32
		return ret
	}
	return *o.MinimumInterval
}

// GetMinimumIntervalOk returns a tuple with the MinimumInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetMinimumIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumInterval) {
		return nil, false
	}
	return o.MinimumInterval, true
}

// HasMinimumInterval returns a boolean if a field has been set.
func (o *AreaEventInfo) HasMinimumInterval() bool {
	if o != nil && !IsNil(o.MinimumInterval) {
		return true
	}

	return false
}

// SetMinimumInterval gets a reference to the given int32 and assigns it to the MinimumInterval field.
func (o *AreaEventInfo) SetMinimumInterval(v int32) {
	o.MinimumInterval = &v
}

// GetMaximumInterval returns the MaximumInterval field value if set, zero value otherwise.
func (o *AreaEventInfo) GetMaximumInterval() int32 {
	if o == nil || IsNil(o.MaximumInterval) {
		var ret int32
		return ret
	}
	return *o.MaximumInterval
}

// GetMaximumIntervalOk returns a tuple with the MaximumInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetMaximumIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumInterval) {
		return nil, false
	}
	return o.MaximumInterval, true
}

// HasMaximumInterval returns a boolean if a field has been set.
func (o *AreaEventInfo) HasMaximumInterval() bool {
	if o != nil && !IsNil(o.MaximumInterval) {
		return true
	}

	return false
}

// SetMaximumInterval gets a reference to the given int32 and assigns it to the MaximumInterval field.
func (o *AreaEventInfo) SetMaximumInterval(v int32) {
	o.MaximumInterval = &v
}

// GetSamplingInterval returns the SamplingInterval field value if set, zero value otherwise.
func (o *AreaEventInfo) GetSamplingInterval() int32 {
	if o == nil || IsNil(o.SamplingInterval) {
		var ret int32
		return ret
	}
	return *o.SamplingInterval
}

// GetSamplingIntervalOk returns a tuple with the SamplingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetSamplingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.SamplingInterval) {
		return nil, false
	}
	return o.SamplingInterval, true
}

// HasSamplingInterval returns a boolean if a field has been set.
func (o *AreaEventInfo) HasSamplingInterval() bool {
	if o != nil && !IsNil(o.SamplingInterval) {
		return true
	}

	return false
}

// SetSamplingInterval gets a reference to the given int32 and assigns it to the SamplingInterval field.
func (o *AreaEventInfo) SetSamplingInterval(v int32) {
	o.SamplingInterval = &v
}

// GetReportingDuration returns the ReportingDuration field value if set, zero value otherwise.
func (o *AreaEventInfo) GetReportingDuration() int32 {
	if o == nil || IsNil(o.ReportingDuration) {
		var ret int32
		return ret
	}
	return *o.ReportingDuration
}

// GetReportingDurationOk returns a tuple with the ReportingDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetReportingDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportingDuration) {
		return nil, false
	}
	return o.ReportingDuration, true
}

// HasReportingDuration returns a boolean if a field has been set.
func (o *AreaEventInfo) HasReportingDuration() bool {
	if o != nil && !IsNil(o.ReportingDuration) {
		return true
	}

	return false
}

// SetReportingDuration gets a reference to the given int32 and assigns it to the ReportingDuration field.
func (o *AreaEventInfo) SetReportingDuration(v int32) {
	o.ReportingDuration = &v
}

// GetReportingLocationReq returns the ReportingLocationReq field value if set, zero value otherwise.
func (o *AreaEventInfo) GetReportingLocationReq() bool {
	if o == nil || IsNil(o.ReportingLocationReq) {
		var ret bool
		return ret
	}
	return *o.ReportingLocationReq
}

// GetReportingLocationReqOk returns a tuple with the ReportingLocationReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AreaEventInfo) GetReportingLocationReqOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportingLocationReq) {
		return nil, false
	}
	return o.ReportingLocationReq, true
}

// HasReportingLocationReq returns a boolean if a field has been set.
func (o *AreaEventInfo) HasReportingLocationReq() bool {
	if o != nil && !IsNil(o.ReportingLocationReq) {
		return true
	}

	return false
}

// SetReportingLocationReq gets a reference to the given bool and assigns it to the ReportingLocationReq field.
func (o *AreaEventInfo) SetReportingLocationReq(v bool) {
	o.ReportingLocationReq = &v
}

func (o AreaEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AreaEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["areaDefinition"] = o.AreaDefinition
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

type NullableAreaEventInfo struct {
	value *AreaEventInfo
	isSet bool
}

func (v NullableAreaEventInfo) Get() *AreaEventInfo {
	return v.value
}

func (v *NullableAreaEventInfo) Set(val *AreaEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaEventInfo(val *AreaEventInfo) *NullableAreaEventInfo {
	return &NullableAreaEventInfo{value: val, isSet: true}
}

func (v NullableAreaEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


