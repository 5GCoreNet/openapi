/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"time"
)

// checks if the MDARequestSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDARequestSingleAllOfAttributesAllOf{}

// MDARequestSingleAllOfAttributesAllOf struct for MDARequestSingleAllOfAttributesAllOf
type MDARequestSingleAllOfAttributesAllOf struct {
	RequestedMDAOutputs []MDAOutputPerMDAType `json:"requestedMDAOutputs,omitempty"`
	ReportingMethod     *ReportingMethod      `json:"reportingMethod,omitempty"`
	ReportingTarget     *string               `json:"reportingTarget,omitempty"`
	AnalyticsScope      *AnalyticsScopeType   `json:"analyticsScope,omitempty"`
	StartTime           *time.Time            `json:"startTime,omitempty"`
	StopTime            *time.Time            `json:"stopTime,omitempty"`
}

// NewMDARequestSingleAllOfAttributesAllOf instantiates a new MDARequestSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDARequestSingleAllOfAttributesAllOf() *MDARequestSingleAllOfAttributesAllOf {
	this := MDARequestSingleAllOfAttributesAllOf{}
	return &this
}

// NewMDARequestSingleAllOfAttributesAllOfWithDefaults instantiates a new MDARequestSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDARequestSingleAllOfAttributesAllOfWithDefaults() *MDARequestSingleAllOfAttributesAllOf {
	this := MDARequestSingleAllOfAttributesAllOf{}
	return &this
}

// GetRequestedMDAOutputs returns the RequestedMDAOutputs field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetRequestedMDAOutputs() []MDAOutputPerMDAType {
	if o == nil || IsNil(o.RequestedMDAOutputs) {
		var ret []MDAOutputPerMDAType
		return ret
	}
	return o.RequestedMDAOutputs
}

// GetRequestedMDAOutputsOk returns a tuple with the RequestedMDAOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetRequestedMDAOutputsOk() ([]MDAOutputPerMDAType, bool) {
	if o == nil || IsNil(o.RequestedMDAOutputs) {
		return nil, false
	}
	return o.RequestedMDAOutputs, true
}

// HasRequestedMDAOutputs returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasRequestedMDAOutputs() bool {
	if o != nil && !IsNil(o.RequestedMDAOutputs) {
		return true
	}

	return false
}

// SetRequestedMDAOutputs gets a reference to the given []MDAOutputPerMDAType and assigns it to the RequestedMDAOutputs field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetRequestedMDAOutputs(v []MDAOutputPerMDAType) {
	o.RequestedMDAOutputs = v
}

// GetReportingMethod returns the ReportingMethod field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetReportingMethod() ReportingMethod {
	if o == nil || IsNil(o.ReportingMethod) {
		var ret ReportingMethod
		return ret
	}
	return *o.ReportingMethod
}

// GetReportingMethodOk returns a tuple with the ReportingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetReportingMethodOk() (*ReportingMethod, bool) {
	if o == nil || IsNil(o.ReportingMethod) {
		return nil, false
	}
	return o.ReportingMethod, true
}

// HasReportingMethod returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasReportingMethod() bool {
	if o != nil && !IsNil(o.ReportingMethod) {
		return true
	}

	return false
}

// SetReportingMethod gets a reference to the given ReportingMethod and assigns it to the ReportingMethod field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetReportingMethod(v ReportingMethod) {
	o.ReportingMethod = &v
}

// GetReportingTarget returns the ReportingTarget field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetReportingTarget() string {
	if o == nil || IsNil(o.ReportingTarget) {
		var ret string
		return ret
	}
	return *o.ReportingTarget
}

// GetReportingTargetOk returns a tuple with the ReportingTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetReportingTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingTarget) {
		return nil, false
	}
	return o.ReportingTarget, true
}

// HasReportingTarget returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasReportingTarget() bool {
	if o != nil && !IsNil(o.ReportingTarget) {
		return true
	}

	return false
}

// SetReportingTarget gets a reference to the given string and assigns it to the ReportingTarget field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetReportingTarget(v string) {
	o.ReportingTarget = &v
}

// GetAnalyticsScope returns the AnalyticsScope field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetAnalyticsScope() AnalyticsScopeType {
	if o == nil || IsNil(o.AnalyticsScope) {
		var ret AnalyticsScopeType
		return ret
	}
	return *o.AnalyticsScope
}

// GetAnalyticsScopeOk returns a tuple with the AnalyticsScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetAnalyticsScopeOk() (*AnalyticsScopeType, bool) {
	if o == nil || IsNil(o.AnalyticsScope) {
		return nil, false
	}
	return o.AnalyticsScope, true
}

// HasAnalyticsScope returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasAnalyticsScope() bool {
	if o != nil && !IsNil(o.AnalyticsScope) {
		return true
	}

	return false
}

// SetAnalyticsScope gets a reference to the given AnalyticsScopeType and assigns it to the AnalyticsScope field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetAnalyticsScope(v AnalyticsScopeType) {
	o.AnalyticsScope = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *MDARequestSingleAllOfAttributesAllOf) GetStopTime() time.Time {
	if o == nil || IsNil(o.StopTime) {
		var ret time.Time
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) GetStopTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StopTime) {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *MDARequestSingleAllOfAttributesAllOf) HasStopTime() bool {
	if o != nil && !IsNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given time.Time and assigns it to the StopTime field.
func (o *MDARequestSingleAllOfAttributesAllOf) SetStopTime(v time.Time) {
	o.StopTime = &v
}

func (o MDARequestSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDARequestSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestedMDAOutputs) {
		toSerialize["requestedMDAOutputs"] = o.RequestedMDAOutputs
	}
	if !IsNil(o.ReportingMethod) {
		toSerialize["reportingMethod"] = o.ReportingMethod
	}
	if !IsNil(o.ReportingTarget) {
		toSerialize["reportingTarget"] = o.ReportingTarget
	}
	if !IsNil(o.AnalyticsScope) {
		toSerialize["analyticsScope"] = o.AnalyticsScope
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	return toSerialize, nil
}

type NullableMDARequestSingleAllOfAttributesAllOf struct {
	value *MDARequestSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableMDARequestSingleAllOfAttributesAllOf) Get() *MDARequestSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableMDARequestSingleAllOfAttributesAllOf) Set(val *MDARequestSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMDARequestSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMDARequestSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDARequestSingleAllOfAttributesAllOf(val *MDARequestSingleAllOfAttributesAllOf) *NullableMDARequestSingleAllOfAttributesAllOf {
	return &NullableMDARequestSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableMDARequestSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDARequestSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
