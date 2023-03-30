/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the GtpUPathQoSMonitoringControlSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GtpUPathQoSMonitoringControlSingleAllOfAttributes{}

// GtpUPathQoSMonitoringControlSingleAllOfAttributes struct for GtpUPathQoSMonitoringControlSingleAllOfAttributes
type GtpUPathQoSMonitoringControlSingleAllOfAttributes struct {
	GtpUPathQoSMonitoringState *string `json:"gtpUPathQoSMonitoringState,omitempty"`
	GtpUPathMonitoredSNSSAIs []Snssai `json:"gtpUPathMonitoredSNSSAIs,omitempty"`
	MonitoredDSCPs []int32 `json:"monitoredDSCPs,omitempty"`
	IsEventTriggeredGtpUPathMonitoringSupported *bool `json:"isEventTriggeredGtpUPathMonitoringSupported,omitempty"`
	IsPeriodicGtpUMonitoringSupported *bool `json:"isPeriodicGtpUMonitoringSupported,omitempty"`
	IsImmediateGtpUMonitoringSupported *bool `json:"isImmediateGtpUMonitoringSupported,omitempty"`
	GtpUPathDelayThresholds *GtpUPathDelayThresholdsType `json:"gtpUPathDelayThresholds,omitempty"`
	GtpUPathMinimumWaitTime *int32 `json:"gtpUPathMinimumWaitTime,omitempty"`
	GtpUPathMeasurementPeriod *int32 `json:"gtpUPathMeasurementPeriod,omitempty"`
}

// NewGtpUPathQoSMonitoringControlSingleAllOfAttributes instantiates a new GtpUPathQoSMonitoringControlSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGtpUPathQoSMonitoringControlSingleAllOfAttributes() *GtpUPathQoSMonitoringControlSingleAllOfAttributes {
	this := GtpUPathQoSMonitoringControlSingleAllOfAttributes{}
	return &this
}

// NewGtpUPathQoSMonitoringControlSingleAllOfAttributesWithDefaults instantiates a new GtpUPathQoSMonitoringControlSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGtpUPathQoSMonitoringControlSingleAllOfAttributesWithDefaults() *GtpUPathQoSMonitoringControlSingleAllOfAttributes {
	this := GtpUPathQoSMonitoringControlSingleAllOfAttributes{}
	return &this
}

// GetGtpUPathQoSMonitoringState returns the GtpUPathQoSMonitoringState field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathQoSMonitoringState() string {
	if o == nil || IsNil(o.GtpUPathQoSMonitoringState) {
		var ret string
		return ret
	}
	return *o.GtpUPathQoSMonitoringState
}

// GetGtpUPathQoSMonitoringStateOk returns a tuple with the GtpUPathQoSMonitoringState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathQoSMonitoringStateOk() (*string, bool) {
	if o == nil || IsNil(o.GtpUPathQoSMonitoringState) {
		return nil, false
	}
	return o.GtpUPathQoSMonitoringState, true
}

// HasGtpUPathQoSMonitoringState returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasGtpUPathQoSMonitoringState() bool {
	if o != nil && !IsNil(o.GtpUPathQoSMonitoringState) {
		return true
	}

	return false
}

// SetGtpUPathQoSMonitoringState gets a reference to the given string and assigns it to the GtpUPathQoSMonitoringState field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetGtpUPathQoSMonitoringState(v string) {
	o.GtpUPathQoSMonitoringState = &v
}

// GetGtpUPathMonitoredSNSSAIs returns the GtpUPathMonitoredSNSSAIs field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMonitoredSNSSAIs() []Snssai {
	if o == nil || IsNil(o.GtpUPathMonitoredSNSSAIs) {
		var ret []Snssai
		return ret
	}
	return o.GtpUPathMonitoredSNSSAIs
}

// GetGtpUPathMonitoredSNSSAIsOk returns a tuple with the GtpUPathMonitoredSNSSAIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMonitoredSNSSAIsOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.GtpUPathMonitoredSNSSAIs) {
		return nil, false
	}
	return o.GtpUPathMonitoredSNSSAIs, true
}

// HasGtpUPathMonitoredSNSSAIs returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasGtpUPathMonitoredSNSSAIs() bool {
	if o != nil && !IsNil(o.GtpUPathMonitoredSNSSAIs) {
		return true
	}

	return false
}

// SetGtpUPathMonitoredSNSSAIs gets a reference to the given []Snssai and assigns it to the GtpUPathMonitoredSNSSAIs field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetGtpUPathMonitoredSNSSAIs(v []Snssai) {
	o.GtpUPathMonitoredSNSSAIs = v
}

// GetMonitoredDSCPs returns the MonitoredDSCPs field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetMonitoredDSCPs() []int32 {
	if o == nil || IsNil(o.MonitoredDSCPs) {
		var ret []int32
		return ret
	}
	return o.MonitoredDSCPs
}

// GetMonitoredDSCPsOk returns a tuple with the MonitoredDSCPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetMonitoredDSCPsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonitoredDSCPs) {
		return nil, false
	}
	return o.MonitoredDSCPs, true
}

// HasMonitoredDSCPs returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasMonitoredDSCPs() bool {
	if o != nil && !IsNil(o.MonitoredDSCPs) {
		return true
	}

	return false
}

// SetMonitoredDSCPs gets a reference to the given []int32 and assigns it to the MonitoredDSCPs field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetMonitoredDSCPs(v []int32) {
	o.MonitoredDSCPs = v
}

// GetIsEventTriggeredGtpUPathMonitoringSupported returns the IsEventTriggeredGtpUPathMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsEventTriggeredGtpUPathMonitoringSupported() bool {
	if o == nil || IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsEventTriggeredGtpUPathMonitoringSupported
}

// GetIsEventTriggeredGtpUPathMonitoringSupportedOk returns a tuple with the IsEventTriggeredGtpUPathMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsEventTriggeredGtpUPathMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		return nil, false
	}
	return o.IsEventTriggeredGtpUPathMonitoringSupported, true
}

// HasIsEventTriggeredGtpUPathMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasIsEventTriggeredGtpUPathMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		return true
	}

	return false
}

// SetIsEventTriggeredGtpUPathMonitoringSupported gets a reference to the given bool and assigns it to the IsEventTriggeredGtpUPathMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetIsEventTriggeredGtpUPathMonitoringSupported(v bool) {
	o.IsEventTriggeredGtpUPathMonitoringSupported = &v
}

// GetIsPeriodicGtpUMonitoringSupported returns the IsPeriodicGtpUMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsPeriodicGtpUMonitoringSupported() bool {
	if o == nil || IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsPeriodicGtpUMonitoringSupported
}

// GetIsPeriodicGtpUMonitoringSupportedOk returns a tuple with the IsPeriodicGtpUMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsPeriodicGtpUMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		return nil, false
	}
	return o.IsPeriodicGtpUMonitoringSupported, true
}

// HasIsPeriodicGtpUMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasIsPeriodicGtpUMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		return true
	}

	return false
}

// SetIsPeriodicGtpUMonitoringSupported gets a reference to the given bool and assigns it to the IsPeriodicGtpUMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetIsPeriodicGtpUMonitoringSupported(v bool) {
	o.IsPeriodicGtpUMonitoringSupported = &v
}

// GetIsImmediateGtpUMonitoringSupported returns the IsImmediateGtpUMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsImmediateGtpUMonitoringSupported() bool {
	if o == nil || IsNil(o.IsImmediateGtpUMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsImmediateGtpUMonitoringSupported
}

// GetIsImmediateGtpUMonitoringSupportedOk returns a tuple with the IsImmediateGtpUMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetIsImmediateGtpUMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImmediateGtpUMonitoringSupported) {
		return nil, false
	}
	return o.IsImmediateGtpUMonitoringSupported, true
}

// HasIsImmediateGtpUMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasIsImmediateGtpUMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsImmediateGtpUMonitoringSupported) {
		return true
	}

	return false
}

// SetIsImmediateGtpUMonitoringSupported gets a reference to the given bool and assigns it to the IsImmediateGtpUMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetIsImmediateGtpUMonitoringSupported(v bool) {
	o.IsImmediateGtpUMonitoringSupported = &v
}

// GetGtpUPathDelayThresholds returns the GtpUPathDelayThresholds field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathDelayThresholds() GtpUPathDelayThresholdsType {
	if o == nil || IsNil(o.GtpUPathDelayThresholds) {
		var ret GtpUPathDelayThresholdsType
		return ret
	}
	return *o.GtpUPathDelayThresholds
}

// GetGtpUPathDelayThresholdsOk returns a tuple with the GtpUPathDelayThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathDelayThresholdsOk() (*GtpUPathDelayThresholdsType, bool) {
	if o == nil || IsNil(o.GtpUPathDelayThresholds) {
		return nil, false
	}
	return o.GtpUPathDelayThresholds, true
}

// HasGtpUPathDelayThresholds returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasGtpUPathDelayThresholds() bool {
	if o != nil && !IsNil(o.GtpUPathDelayThresholds) {
		return true
	}

	return false
}

// SetGtpUPathDelayThresholds gets a reference to the given GtpUPathDelayThresholdsType and assigns it to the GtpUPathDelayThresholds field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetGtpUPathDelayThresholds(v GtpUPathDelayThresholdsType) {
	o.GtpUPathDelayThresholds = &v
}

// GetGtpUPathMinimumWaitTime returns the GtpUPathMinimumWaitTime field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMinimumWaitTime() int32 {
	if o == nil || IsNil(o.GtpUPathMinimumWaitTime) {
		var ret int32
		return ret
	}
	return *o.GtpUPathMinimumWaitTime
}

// GetGtpUPathMinimumWaitTimeOk returns a tuple with the GtpUPathMinimumWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMinimumWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GtpUPathMinimumWaitTime) {
		return nil, false
	}
	return o.GtpUPathMinimumWaitTime, true
}

// HasGtpUPathMinimumWaitTime returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasGtpUPathMinimumWaitTime() bool {
	if o != nil && !IsNil(o.GtpUPathMinimumWaitTime) {
		return true
	}

	return false
}

// SetGtpUPathMinimumWaitTime gets a reference to the given int32 and assigns it to the GtpUPathMinimumWaitTime field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetGtpUPathMinimumWaitTime(v int32) {
	o.GtpUPathMinimumWaitTime = &v
}

// GetGtpUPathMeasurementPeriod returns the GtpUPathMeasurementPeriod field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMeasurementPeriod() int32 {
	if o == nil || IsNil(o.GtpUPathMeasurementPeriod) {
		var ret int32
		return ret
	}
	return *o.GtpUPathMeasurementPeriod
}

// GetGtpUPathMeasurementPeriodOk returns a tuple with the GtpUPathMeasurementPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) GetGtpUPathMeasurementPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.GtpUPathMeasurementPeriod) {
		return nil, false
	}
	return o.GtpUPathMeasurementPeriod, true
}

// HasGtpUPathMeasurementPeriod returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) HasGtpUPathMeasurementPeriod() bool {
	if o != nil && !IsNil(o.GtpUPathMeasurementPeriod) {
		return true
	}

	return false
}

// SetGtpUPathMeasurementPeriod gets a reference to the given int32 and assigns it to the GtpUPathMeasurementPeriod field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributes) SetGtpUPathMeasurementPeriod(v int32) {
	o.GtpUPathMeasurementPeriod = &v
}

func (o GtpUPathQoSMonitoringControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GtpUPathQoSMonitoringControlSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GtpUPathQoSMonitoringState) {
		toSerialize["gtpUPathQoSMonitoringState"] = o.GtpUPathQoSMonitoringState
	}
	if !IsNil(o.GtpUPathMonitoredSNSSAIs) {
		toSerialize["gtpUPathMonitoredSNSSAIs"] = o.GtpUPathMonitoredSNSSAIs
	}
	if !IsNil(o.MonitoredDSCPs) {
		toSerialize["monitoredDSCPs"] = o.MonitoredDSCPs
	}
	if !IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		toSerialize["isEventTriggeredGtpUPathMonitoringSupported"] = o.IsEventTriggeredGtpUPathMonitoringSupported
	}
	if !IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		toSerialize["isPeriodicGtpUMonitoringSupported"] = o.IsPeriodicGtpUMonitoringSupported
	}
	if !IsNil(o.IsImmediateGtpUMonitoringSupported) {
		toSerialize["isImmediateGtpUMonitoringSupported"] = o.IsImmediateGtpUMonitoringSupported
	}
	if !IsNil(o.GtpUPathDelayThresholds) {
		toSerialize["gtpUPathDelayThresholds"] = o.GtpUPathDelayThresholds
	}
	if !IsNil(o.GtpUPathMinimumWaitTime) {
		toSerialize["gtpUPathMinimumWaitTime"] = o.GtpUPathMinimumWaitTime
	}
	if !IsNil(o.GtpUPathMeasurementPeriod) {
		toSerialize["gtpUPathMeasurementPeriod"] = o.GtpUPathMeasurementPeriod
	}
	return toSerialize, nil
}

type NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes struct {
	value *GtpUPathQoSMonitoringControlSingleAllOfAttributes
	isSet bool
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) Get() *GtpUPathQoSMonitoringControlSingleAllOfAttributes {
	return v.value
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) Set(val *GtpUPathQoSMonitoringControlSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGtpUPathQoSMonitoringControlSingleAllOfAttributes(val *GtpUPathQoSMonitoringControlSingleAllOfAttributes) *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes {
	return &NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


