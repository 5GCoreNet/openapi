/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
)

// checks if the ReportingCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingCondition{}

// ReportingCondition A condition that triggers data reporting by a data collection client to the Data Collection AF.
type ReportingCondition struct {
	Type ReportingConditionType `json:"type"`
	// indicating a time in seconds.
	Period *int32 `json:"period,omitempty"`
	Parameter *string `json:"parameter,omitempty"`
	Threshold *ReportingConditionThreshold `json:"threshold,omitempty"`
	ReportWhenBelow *bool `json:"reportWhenBelow,omitempty"`
	EventTrigger *ReportingEventTrigger `json:"eventTrigger,omitempty"`
}

// NewReportingCondition instantiates a new ReportingCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingCondition(type_ ReportingConditionType) *ReportingCondition {
	this := ReportingCondition{}
	this.Type = type_
	return &this
}

// NewReportingConditionWithDefaults instantiates a new ReportingCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingConditionWithDefaults() *ReportingCondition {
	this := ReportingCondition{}
	return &this
}

// GetType returns the Type field value
func (o *ReportingCondition) GetType() ReportingConditionType {
	if o == nil {
		var ret ReportingConditionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetTypeOk() (*ReportingConditionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReportingCondition) SetType(v ReportingConditionType) {
	o.Type = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ReportingCondition) GetPeriod() int32 {
	if o == nil || isNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ReportingCondition) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *ReportingCondition) SetPeriod(v int32) {
	o.Period = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ReportingCondition) GetParameter() string {
	if o == nil || isNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetParameterOk() (*string, bool) {
	if o == nil || isNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ReportingCondition) HasParameter() bool {
	if o != nil && !isNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ReportingCondition) SetParameter(v string) {
	o.Parameter = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ReportingCondition) GetThreshold() ReportingConditionThreshold {
	if o == nil || isNil(o.Threshold) {
		var ret ReportingConditionThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetThresholdOk() (*ReportingConditionThreshold, bool) {
	if o == nil || isNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ReportingCondition) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given ReportingConditionThreshold and assigns it to the Threshold field.
func (o *ReportingCondition) SetThreshold(v ReportingConditionThreshold) {
	o.Threshold = &v
}

// GetReportWhenBelow returns the ReportWhenBelow field value if set, zero value otherwise.
func (o *ReportingCondition) GetReportWhenBelow() bool {
	if o == nil || isNil(o.ReportWhenBelow) {
		var ret bool
		return ret
	}
	return *o.ReportWhenBelow
}

// GetReportWhenBelowOk returns a tuple with the ReportWhenBelow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetReportWhenBelowOk() (*bool, bool) {
	if o == nil || isNil(o.ReportWhenBelow) {
		return nil, false
	}
	return o.ReportWhenBelow, true
}

// HasReportWhenBelow returns a boolean if a field has been set.
func (o *ReportingCondition) HasReportWhenBelow() bool {
	if o != nil && !isNil(o.ReportWhenBelow) {
		return true
	}

	return false
}

// SetReportWhenBelow gets a reference to the given bool and assigns it to the ReportWhenBelow field.
func (o *ReportingCondition) SetReportWhenBelow(v bool) {
	o.ReportWhenBelow = &v
}

// GetEventTrigger returns the EventTrigger field value if set, zero value otherwise.
func (o *ReportingCondition) GetEventTrigger() ReportingEventTrigger {
	if o == nil || isNil(o.EventTrigger) {
		var ret ReportingEventTrigger
		return ret
	}
	return *o.EventTrigger
}

// GetEventTriggerOk returns a tuple with the EventTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingCondition) GetEventTriggerOk() (*ReportingEventTrigger, bool) {
	if o == nil || isNil(o.EventTrigger) {
		return nil, false
	}
	return o.EventTrigger, true
}

// HasEventTrigger returns a boolean if a field has been set.
func (o *ReportingCondition) HasEventTrigger() bool {
	if o != nil && !isNil(o.EventTrigger) {
		return true
	}

	return false
}

// SetEventTrigger gets a reference to the given ReportingEventTrigger and assigns it to the EventTrigger field.
func (o *ReportingCondition) SetEventTrigger(v ReportingEventTrigger) {
	o.EventTrigger = &v
}

func (o ReportingCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !isNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !isNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !isNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !isNil(o.ReportWhenBelow) {
		toSerialize["reportWhenBelow"] = o.ReportWhenBelow
	}
	if !isNil(o.EventTrigger) {
		toSerialize["eventTrigger"] = o.EventTrigger
	}
	return toSerialize, nil
}

type NullableReportingCondition struct {
	value *ReportingCondition
	isSet bool
}

func (v NullableReportingCondition) Get() *ReportingCondition {
	return v.value
}

func (v *NullableReportingCondition) Set(val *ReportingCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingCondition(val *ReportingCondition) *NullableReportingCondition {
	return &NullableReportingCondition{value: val, isSet: true}
}

func (v NullableReportingCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


