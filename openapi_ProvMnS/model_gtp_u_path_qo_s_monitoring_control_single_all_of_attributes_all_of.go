/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf{}

// GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf struct for GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf
type GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf struct {
	GtpUPathQoSMonitoringState                  *string                      `json:"gtpUPathQoSMonitoringState,omitempty"`
	GtpUPathMonitoredSNSSAIs                    []Snssai                     `json:"gtpUPathMonitoredSNSSAIs,omitempty"`
	MonitoredDSCPs                              []int32                      `json:"monitoredDSCPs,omitempty"`
	IsEventTriggeredGtpUPathMonitoringSupported *bool                        `json:"isEventTriggeredGtpUPathMonitoringSupported,omitempty"`
	IsPeriodicGtpUMonitoringSupported           *bool                        `json:"isPeriodicGtpUMonitoringSupported,omitempty"`
	IsImmediateGtpUMonitoringSupported          *bool                        `json:"isImmediateGtpUMonitoringSupported,omitempty"`
	GtpUPathDelayThresholds                     *GtpUPathDelayThresholdsType `json:"gtpUPathDelayThresholds,omitempty"`
	GtpUPathMinimumWaitTime                     *int32                       `json:"gtpUPathMinimumWaitTime,omitempty"`
	GtpUPathMeasurementPeriod                   *int32                       `json:"gtpUPathMeasurementPeriod,omitempty"`
}

// NewGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf instantiates a new GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf() *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf {
	this := GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf{}
	return &this
}

// NewGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOfWithDefaults instantiates a new GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOfWithDefaults() *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf {
	this := GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf{}
	return &this
}

// GetGtpUPathQoSMonitoringState returns the GtpUPathQoSMonitoringState field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathQoSMonitoringState() string {
	if o == nil || IsNil(o.GtpUPathQoSMonitoringState) {
		var ret string
		return ret
	}
	return *o.GtpUPathQoSMonitoringState
}

// GetGtpUPathQoSMonitoringStateOk returns a tuple with the GtpUPathQoSMonitoringState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathQoSMonitoringStateOk() (*string, bool) {
	if o == nil || IsNil(o.GtpUPathQoSMonitoringState) {
		return nil, false
	}
	return o.GtpUPathQoSMonitoringState, true
}

// HasGtpUPathQoSMonitoringState returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasGtpUPathQoSMonitoringState() bool {
	if o != nil && !IsNil(o.GtpUPathQoSMonitoringState) {
		return true
	}

	return false
}

// SetGtpUPathQoSMonitoringState gets a reference to the given string and assigns it to the GtpUPathQoSMonitoringState field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetGtpUPathQoSMonitoringState(v string) {
	o.GtpUPathQoSMonitoringState = &v
}

// GetGtpUPathMonitoredSNSSAIs returns the GtpUPathMonitoredSNSSAIs field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMonitoredSNSSAIs() []Snssai {
	if o == nil || IsNil(o.GtpUPathMonitoredSNSSAIs) {
		var ret []Snssai
		return ret
	}
	return o.GtpUPathMonitoredSNSSAIs
}

// GetGtpUPathMonitoredSNSSAIsOk returns a tuple with the GtpUPathMonitoredSNSSAIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMonitoredSNSSAIsOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.GtpUPathMonitoredSNSSAIs) {
		return nil, false
	}
	return o.GtpUPathMonitoredSNSSAIs, true
}

// HasGtpUPathMonitoredSNSSAIs returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasGtpUPathMonitoredSNSSAIs() bool {
	if o != nil && !IsNil(o.GtpUPathMonitoredSNSSAIs) {
		return true
	}

	return false
}

// SetGtpUPathMonitoredSNSSAIs gets a reference to the given []Snssai and assigns it to the GtpUPathMonitoredSNSSAIs field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetGtpUPathMonitoredSNSSAIs(v []Snssai) {
	o.GtpUPathMonitoredSNSSAIs = v
}

// GetMonitoredDSCPs returns the MonitoredDSCPs field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetMonitoredDSCPs() []int32 {
	if o == nil || IsNil(o.MonitoredDSCPs) {
		var ret []int32
		return ret
	}
	return o.MonitoredDSCPs
}

// GetMonitoredDSCPsOk returns a tuple with the MonitoredDSCPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetMonitoredDSCPsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonitoredDSCPs) {
		return nil, false
	}
	return o.MonitoredDSCPs, true
}

// HasMonitoredDSCPs returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasMonitoredDSCPs() bool {
	if o != nil && !IsNil(o.MonitoredDSCPs) {
		return true
	}

	return false
}

// SetMonitoredDSCPs gets a reference to the given []int32 and assigns it to the MonitoredDSCPs field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetMonitoredDSCPs(v []int32) {
	o.MonitoredDSCPs = v
}

// GetIsEventTriggeredGtpUPathMonitoringSupported returns the IsEventTriggeredGtpUPathMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsEventTriggeredGtpUPathMonitoringSupported() bool {
	if o == nil || IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsEventTriggeredGtpUPathMonitoringSupported
}

// GetIsEventTriggeredGtpUPathMonitoringSupportedOk returns a tuple with the IsEventTriggeredGtpUPathMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsEventTriggeredGtpUPathMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		return nil, false
	}
	return o.IsEventTriggeredGtpUPathMonitoringSupported, true
}

// HasIsEventTriggeredGtpUPathMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsEventTriggeredGtpUPathMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsEventTriggeredGtpUPathMonitoringSupported) {
		return true
	}

	return false
}

// SetIsEventTriggeredGtpUPathMonitoringSupported gets a reference to the given bool and assigns it to the IsEventTriggeredGtpUPathMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsEventTriggeredGtpUPathMonitoringSupported(v bool) {
	o.IsEventTriggeredGtpUPathMonitoringSupported = &v
}

// GetIsPeriodicGtpUMonitoringSupported returns the IsPeriodicGtpUMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsPeriodicGtpUMonitoringSupported() bool {
	if o == nil || IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsPeriodicGtpUMonitoringSupported
}

// GetIsPeriodicGtpUMonitoringSupportedOk returns a tuple with the IsPeriodicGtpUMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsPeriodicGtpUMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		return nil, false
	}
	return o.IsPeriodicGtpUMonitoringSupported, true
}

// HasIsPeriodicGtpUMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsPeriodicGtpUMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsPeriodicGtpUMonitoringSupported) {
		return true
	}

	return false
}

// SetIsPeriodicGtpUMonitoringSupported gets a reference to the given bool and assigns it to the IsPeriodicGtpUMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsPeriodicGtpUMonitoringSupported(v bool) {
	o.IsPeriodicGtpUMonitoringSupported = &v
}

// GetIsImmediateGtpUMonitoringSupported returns the IsImmediateGtpUMonitoringSupported field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsImmediateGtpUMonitoringSupported() bool {
	if o == nil || IsNil(o.IsImmediateGtpUMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsImmediateGtpUMonitoringSupported
}

// GetIsImmediateGtpUMonitoringSupportedOk returns a tuple with the IsImmediateGtpUMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsImmediateGtpUMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImmediateGtpUMonitoringSupported) {
		return nil, false
	}
	return o.IsImmediateGtpUMonitoringSupported, true
}

// HasIsImmediateGtpUMonitoringSupported returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsImmediateGtpUMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsImmediateGtpUMonitoringSupported) {
		return true
	}

	return false
}

// SetIsImmediateGtpUMonitoringSupported gets a reference to the given bool and assigns it to the IsImmediateGtpUMonitoringSupported field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsImmediateGtpUMonitoringSupported(v bool) {
	o.IsImmediateGtpUMonitoringSupported = &v
}

// GetGtpUPathDelayThresholds returns the GtpUPathDelayThresholds field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathDelayThresholds() GtpUPathDelayThresholdsType {
	if o == nil || IsNil(o.GtpUPathDelayThresholds) {
		var ret GtpUPathDelayThresholdsType
		return ret
	}
	return *o.GtpUPathDelayThresholds
}

// GetGtpUPathDelayThresholdsOk returns a tuple with the GtpUPathDelayThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathDelayThresholdsOk() (*GtpUPathDelayThresholdsType, bool) {
	if o == nil || IsNil(o.GtpUPathDelayThresholds) {
		return nil, false
	}
	return o.GtpUPathDelayThresholds, true
}

// HasGtpUPathDelayThresholds returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasGtpUPathDelayThresholds() bool {
	if o != nil && !IsNil(o.GtpUPathDelayThresholds) {
		return true
	}

	return false
}

// SetGtpUPathDelayThresholds gets a reference to the given GtpUPathDelayThresholdsType and assigns it to the GtpUPathDelayThresholds field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetGtpUPathDelayThresholds(v GtpUPathDelayThresholdsType) {
	o.GtpUPathDelayThresholds = &v
}

// GetGtpUPathMinimumWaitTime returns the GtpUPathMinimumWaitTime field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMinimumWaitTime() int32 {
	if o == nil || IsNil(o.GtpUPathMinimumWaitTime) {
		var ret int32
		return ret
	}
	return *o.GtpUPathMinimumWaitTime
}

// GetGtpUPathMinimumWaitTimeOk returns a tuple with the GtpUPathMinimumWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMinimumWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GtpUPathMinimumWaitTime) {
		return nil, false
	}
	return o.GtpUPathMinimumWaitTime, true
}

// HasGtpUPathMinimumWaitTime returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasGtpUPathMinimumWaitTime() bool {
	if o != nil && !IsNil(o.GtpUPathMinimumWaitTime) {
		return true
	}

	return false
}

// SetGtpUPathMinimumWaitTime gets a reference to the given int32 and assigns it to the GtpUPathMinimumWaitTime field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetGtpUPathMinimumWaitTime(v int32) {
	o.GtpUPathMinimumWaitTime = &v
}

// GetGtpUPathMeasurementPeriod returns the GtpUPathMeasurementPeriod field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMeasurementPeriod() int32 {
	if o == nil || IsNil(o.GtpUPathMeasurementPeriod) {
		var ret int32
		return ret
	}
	return *o.GtpUPathMeasurementPeriod
}

// GetGtpUPathMeasurementPeriodOk returns a tuple with the GtpUPathMeasurementPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) GetGtpUPathMeasurementPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.GtpUPathMeasurementPeriod) {
		return nil, false
	}
	return o.GtpUPathMeasurementPeriod, true
}

// HasGtpUPathMeasurementPeriod returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) HasGtpUPathMeasurementPeriod() bool {
	if o != nil && !IsNil(o.GtpUPathMeasurementPeriod) {
		return true
	}

	return false
}

// SetGtpUPathMeasurementPeriod gets a reference to the given int32 and assigns it to the GtpUPathMeasurementPeriod field.
func (o *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) SetGtpUPathMeasurementPeriod(v int32) {
	o.GtpUPathMeasurementPeriod = &v
}

func (o GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
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

type NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf struct {
	value *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) Get() *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) Set(val *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf(val *GtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf {
	return &NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
