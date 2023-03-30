/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
)

// checks if the ReportingRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingRequirements{}

// ReportingRequirements Indicates the requested frequency of reporting.
type ReportingRequirements struct {
	ReportingMode NotificationMethod `json:"reportingMode"`
	// indicating a time in seconds.
	ReportingPeriod *int32 `json:"reportingPeriod,omitempty"`
	ReportingThrs []ReportingThreshold `json:"reportingThrs,omitempty"`
	ImmRep *bool `json:"immRep,omitempty"`
	RepTerminMode *TerminationMode `json:"repTerminMode,omitempty"`
	// indicating a time in seconds.
	ExpirationTimer *int32 `json:"expirationTimer,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumRep *int32 `json:"maxNumRep,omitempty"`
	TermThr *MeasurementData `json:"termThr,omitempty"`
	TermThrMode *ThresholdHandlingMode `json:"termThrMode,omitempty"`
}

// NewReportingRequirements instantiates a new ReportingRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingRequirements(reportingMode NotificationMethod) *ReportingRequirements {
	this := ReportingRequirements{}
	this.ReportingMode = reportingMode
	return &this
}

// NewReportingRequirementsWithDefaults instantiates a new ReportingRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingRequirementsWithDefaults() *ReportingRequirements {
	this := ReportingRequirements{}
	return &this
}

// GetReportingMode returns the ReportingMode field value
func (o *ReportingRequirements) GetReportingMode() NotificationMethod {
	if o == nil {
		var ret NotificationMethod
		return ret
	}

	return o.ReportingMode
}

// GetReportingModeOk returns a tuple with the ReportingMode field value
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetReportingModeOk() (*NotificationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingMode, true
}

// SetReportingMode sets field value
func (o *ReportingRequirements) SetReportingMode(v NotificationMethod) {
	o.ReportingMode = v
}

// GetReportingPeriod returns the ReportingPeriod field value if set, zero value otherwise.
func (o *ReportingRequirements) GetReportingPeriod() int32 {
	if o == nil || IsNil(o.ReportingPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportingPeriod
}

// GetReportingPeriodOk returns a tuple with the ReportingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetReportingPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportingPeriod) {
		return nil, false
	}
	return o.ReportingPeriod, true
}

// HasReportingPeriod returns a boolean if a field has been set.
func (o *ReportingRequirements) HasReportingPeriod() bool {
	if o != nil && !IsNil(o.ReportingPeriod) {
		return true
	}

	return false
}

// SetReportingPeriod gets a reference to the given int32 and assigns it to the ReportingPeriod field.
func (o *ReportingRequirements) SetReportingPeriod(v int32) {
	o.ReportingPeriod = &v
}

// GetReportingThrs returns the ReportingThrs field value if set, zero value otherwise.
func (o *ReportingRequirements) GetReportingThrs() []ReportingThreshold {
	if o == nil || IsNil(o.ReportingThrs) {
		var ret []ReportingThreshold
		return ret
	}
	return o.ReportingThrs
}

// GetReportingThrsOk returns a tuple with the ReportingThrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetReportingThrsOk() ([]ReportingThreshold, bool) {
	if o == nil || IsNil(o.ReportingThrs) {
		return nil, false
	}
	return o.ReportingThrs, true
}

// HasReportingThrs returns a boolean if a field has been set.
func (o *ReportingRequirements) HasReportingThrs() bool {
	if o != nil && !IsNil(o.ReportingThrs) {
		return true
	}

	return false
}

// SetReportingThrs gets a reference to the given []ReportingThreshold and assigns it to the ReportingThrs field.
func (o *ReportingRequirements) SetReportingThrs(v []ReportingThreshold) {
	o.ReportingThrs = v
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *ReportingRequirements) GetImmRep() bool {
	if o == nil || IsNil(o.ImmRep) {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetImmRepOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmRep) {
		return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *ReportingRequirements) HasImmRep() bool {
	if o != nil && !IsNil(o.ImmRep) {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *ReportingRequirements) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetRepTerminMode returns the RepTerminMode field value if set, zero value otherwise.
func (o *ReportingRequirements) GetRepTerminMode() TerminationMode {
	if o == nil || IsNil(o.RepTerminMode) {
		var ret TerminationMode
		return ret
	}
	return *o.RepTerminMode
}

// GetRepTerminModeOk returns a tuple with the RepTerminMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetRepTerminModeOk() (*TerminationMode, bool) {
	if o == nil || IsNil(o.RepTerminMode) {
		return nil, false
	}
	return o.RepTerminMode, true
}

// HasRepTerminMode returns a boolean if a field has been set.
func (o *ReportingRequirements) HasRepTerminMode() bool {
	if o != nil && !IsNil(o.RepTerminMode) {
		return true
	}

	return false
}

// SetRepTerminMode gets a reference to the given TerminationMode and assigns it to the RepTerminMode field.
func (o *ReportingRequirements) SetRepTerminMode(v TerminationMode) {
	o.RepTerminMode = &v
}

// GetExpirationTimer returns the ExpirationTimer field value if set, zero value otherwise.
func (o *ReportingRequirements) GetExpirationTimer() int32 {
	if o == nil || IsNil(o.ExpirationTimer) {
		var ret int32
		return ret
	}
	return *o.ExpirationTimer
}

// GetExpirationTimerOk returns a tuple with the ExpirationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetExpirationTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationTimer) {
		return nil, false
	}
	return o.ExpirationTimer, true
}

// HasExpirationTimer returns a boolean if a field has been set.
func (o *ReportingRequirements) HasExpirationTimer() bool {
	if o != nil && !IsNil(o.ExpirationTimer) {
		return true
	}

	return false
}

// SetExpirationTimer gets a reference to the given int32 and assigns it to the ExpirationTimer field.
func (o *ReportingRequirements) SetExpirationTimer(v int32) {
	o.ExpirationTimer = &v
}

// GetMaxNumRep returns the MaxNumRep field value if set, zero value otherwise.
func (o *ReportingRequirements) GetMaxNumRep() int32 {
	if o == nil || IsNil(o.MaxNumRep) {
		var ret int32
		return ret
	}
	return *o.MaxNumRep
}

// GetMaxNumRepOk returns a tuple with the MaxNumRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetMaxNumRepOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumRep) {
		return nil, false
	}
	return o.MaxNumRep, true
}

// HasMaxNumRep returns a boolean if a field has been set.
func (o *ReportingRequirements) HasMaxNumRep() bool {
	if o != nil && !IsNil(o.MaxNumRep) {
		return true
	}

	return false
}

// SetMaxNumRep gets a reference to the given int32 and assigns it to the MaxNumRep field.
func (o *ReportingRequirements) SetMaxNumRep(v int32) {
	o.MaxNumRep = &v
}

// GetTermThr returns the TermThr field value if set, zero value otherwise.
func (o *ReportingRequirements) GetTermThr() MeasurementData {
	if o == nil || IsNil(o.TermThr) {
		var ret MeasurementData
		return ret
	}
	return *o.TermThr
}

// GetTermThrOk returns a tuple with the TermThr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetTermThrOk() (*MeasurementData, bool) {
	if o == nil || IsNil(o.TermThr) {
		return nil, false
	}
	return o.TermThr, true
}

// HasTermThr returns a boolean if a field has been set.
func (o *ReportingRequirements) HasTermThr() bool {
	if o != nil && !IsNil(o.TermThr) {
		return true
	}

	return false
}

// SetTermThr gets a reference to the given MeasurementData and assigns it to the TermThr field.
func (o *ReportingRequirements) SetTermThr(v MeasurementData) {
	o.TermThr = &v
}

// GetTermThrMode returns the TermThrMode field value if set, zero value otherwise.
func (o *ReportingRequirements) GetTermThrMode() ThresholdHandlingMode {
	if o == nil || IsNil(o.TermThrMode) {
		var ret ThresholdHandlingMode
		return ret
	}
	return *o.TermThrMode
}

// GetTermThrModeOk returns a tuple with the TermThrMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingRequirements) GetTermThrModeOk() (*ThresholdHandlingMode, bool) {
	if o == nil || IsNil(o.TermThrMode) {
		return nil, false
	}
	return o.TermThrMode, true
}

// HasTermThrMode returns a boolean if a field has been set.
func (o *ReportingRequirements) HasTermThrMode() bool {
	if o != nil && !IsNil(o.TermThrMode) {
		return true
	}

	return false
}

// SetTermThrMode gets a reference to the given ThresholdHandlingMode and assigns it to the TermThrMode field.
func (o *ReportingRequirements) SetTermThrMode(v ThresholdHandlingMode) {
	o.TermThrMode = &v
}

func (o ReportingRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportingMode"] = o.ReportingMode
	if !IsNil(o.ReportingPeriod) {
		toSerialize["reportingPeriod"] = o.ReportingPeriod
	}
	if !IsNil(o.ReportingThrs) {
		toSerialize["reportingThrs"] = o.ReportingThrs
	}
	if !IsNil(o.ImmRep) {
		toSerialize["immRep"] = o.ImmRep
	}
	if !IsNil(o.RepTerminMode) {
		toSerialize["repTerminMode"] = o.RepTerminMode
	}
	if !IsNil(o.ExpirationTimer) {
		toSerialize["expirationTimer"] = o.ExpirationTimer
	}
	if !IsNil(o.MaxNumRep) {
		toSerialize["maxNumRep"] = o.MaxNumRep
	}
	if !IsNil(o.TermThr) {
		toSerialize["termThr"] = o.TermThr
	}
	if !IsNil(o.TermThrMode) {
		toSerialize["termThrMode"] = o.TermThrMode
	}
	return toSerialize, nil
}

type NullableReportingRequirements struct {
	value *ReportingRequirements
	isSet bool
}

func (v NullableReportingRequirements) Get() *ReportingRequirements {
	return v.value
}

func (v *NullableReportingRequirements) Set(val *ReportingRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingRequirements(val *ReportingRequirements) *NullableReportingRequirements {
	return &NullableReportingRequirements{value: val, isSet: true}
}

func (v NullableReportingRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

