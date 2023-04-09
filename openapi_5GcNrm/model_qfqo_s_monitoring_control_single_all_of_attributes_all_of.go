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

// checks if the QFQoSMonitoringControlSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QFQoSMonitoringControlSingleAllOfAttributesAllOf{}

// QFQoSMonitoringControlSingleAllOfAttributesAllOf struct for QFQoSMonitoringControlSingleAllOfAttributesAllOf
type QFQoSMonitoringControlSingleAllOfAttributesAllOf struct {
	QFQoSMonitoringState                   *string                      `json:"qFQoSMonitoringState,omitempty"`
	QFMonitoredSNSSAIs                     []Snssai                     `json:"qFMonitoredSNSSAIs,omitempty"`
	QFMonitored5QIs                        []int32                      `json:"qFMonitored5QIs,omitempty"`
	IsEventTriggeredQFMonitoringSupported  *bool                        `json:"isEventTriggeredQFMonitoringSupported,omitempty"`
	IsPeriodicQFMonitoringSupported        *bool                        `json:"isPeriodicQFMonitoringSupported,omitempty"`
	IsSessionReleasedQFMonitoringSupported *bool                        `json:"isSessionReleasedQFMonitoringSupported,omitempty"`
	QFPacketDelayThresholds                *QFPacketDelayThresholdsType `json:"qFPacketDelayThresholds,omitempty"`
	QFMinimumWaitTime                      *int32                       `json:"qFMinimumWaitTime,omitempty"`
	QFMeasurementPeriod                    *int32                       `json:"qFMeasurementPeriod,omitempty"`
}

// NewQFQoSMonitoringControlSingleAllOfAttributesAllOf instantiates a new QFQoSMonitoringControlSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQFQoSMonitoringControlSingleAllOfAttributesAllOf() *QFQoSMonitoringControlSingleAllOfAttributesAllOf {
	this := QFQoSMonitoringControlSingleAllOfAttributesAllOf{}
	return &this
}

// NewQFQoSMonitoringControlSingleAllOfAttributesAllOfWithDefaults instantiates a new QFQoSMonitoringControlSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQFQoSMonitoringControlSingleAllOfAttributesAllOfWithDefaults() *QFQoSMonitoringControlSingleAllOfAttributesAllOf {
	this := QFQoSMonitoringControlSingleAllOfAttributesAllOf{}
	return &this
}

// GetQFQoSMonitoringState returns the QFQoSMonitoringState field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFQoSMonitoringState() string {
	if o == nil || IsNil(o.QFQoSMonitoringState) {
		var ret string
		return ret
	}
	return *o.QFQoSMonitoringState
}

// GetQFQoSMonitoringStateOk returns a tuple with the QFQoSMonitoringState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFQoSMonitoringStateOk() (*string, bool) {
	if o == nil || IsNil(o.QFQoSMonitoringState) {
		return nil, false
	}
	return o.QFQoSMonitoringState, true
}

// HasQFQoSMonitoringState returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFQoSMonitoringState() bool {
	if o != nil && !IsNil(o.QFQoSMonitoringState) {
		return true
	}

	return false
}

// SetQFQoSMonitoringState gets a reference to the given string and assigns it to the QFQoSMonitoringState field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFQoSMonitoringState(v string) {
	o.QFQoSMonitoringState = &v
}

// GetQFMonitoredSNSSAIs returns the QFMonitoredSNSSAIs field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMonitoredSNSSAIs() []Snssai {
	if o == nil || IsNil(o.QFMonitoredSNSSAIs) {
		var ret []Snssai
		return ret
	}
	return o.QFMonitoredSNSSAIs
}

// GetQFMonitoredSNSSAIsOk returns a tuple with the QFMonitoredSNSSAIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMonitoredSNSSAIsOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.QFMonitoredSNSSAIs) {
		return nil, false
	}
	return o.QFMonitoredSNSSAIs, true
}

// HasQFMonitoredSNSSAIs returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFMonitoredSNSSAIs() bool {
	if o != nil && !IsNil(o.QFMonitoredSNSSAIs) {
		return true
	}

	return false
}

// SetQFMonitoredSNSSAIs gets a reference to the given []Snssai and assigns it to the QFMonitoredSNSSAIs field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFMonitoredSNSSAIs(v []Snssai) {
	o.QFMonitoredSNSSAIs = v
}

// GetQFMonitored5QIs returns the QFMonitored5QIs field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMonitored5QIs() []int32 {
	if o == nil || IsNil(o.QFMonitored5QIs) {
		var ret []int32
		return ret
	}
	return o.QFMonitored5QIs
}

// GetQFMonitored5QIsOk returns a tuple with the QFMonitored5QIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMonitored5QIsOk() ([]int32, bool) {
	if o == nil || IsNil(o.QFMonitored5QIs) {
		return nil, false
	}
	return o.QFMonitored5QIs, true
}

// HasQFMonitored5QIs returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFMonitored5QIs() bool {
	if o != nil && !IsNil(o.QFMonitored5QIs) {
		return true
	}

	return false
}

// SetQFMonitored5QIs gets a reference to the given []int32 and assigns it to the QFMonitored5QIs field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFMonitored5QIs(v []int32) {
	o.QFMonitored5QIs = v
}

// GetIsEventTriggeredQFMonitoringSupported returns the IsEventTriggeredQFMonitoringSupported field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsEventTriggeredQFMonitoringSupported() bool {
	if o == nil || IsNil(o.IsEventTriggeredQFMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsEventTriggeredQFMonitoringSupported
}

// GetIsEventTriggeredQFMonitoringSupportedOk returns a tuple with the IsEventTriggeredQFMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsEventTriggeredQFMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEventTriggeredQFMonitoringSupported) {
		return nil, false
	}
	return o.IsEventTriggeredQFMonitoringSupported, true
}

// HasIsEventTriggeredQFMonitoringSupported returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsEventTriggeredQFMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsEventTriggeredQFMonitoringSupported) {
		return true
	}

	return false
}

// SetIsEventTriggeredQFMonitoringSupported gets a reference to the given bool and assigns it to the IsEventTriggeredQFMonitoringSupported field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsEventTriggeredQFMonitoringSupported(v bool) {
	o.IsEventTriggeredQFMonitoringSupported = &v
}

// GetIsPeriodicQFMonitoringSupported returns the IsPeriodicQFMonitoringSupported field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsPeriodicQFMonitoringSupported() bool {
	if o == nil || IsNil(o.IsPeriodicQFMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsPeriodicQFMonitoringSupported
}

// GetIsPeriodicQFMonitoringSupportedOk returns a tuple with the IsPeriodicQFMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsPeriodicQFMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPeriodicQFMonitoringSupported) {
		return nil, false
	}
	return o.IsPeriodicQFMonitoringSupported, true
}

// HasIsPeriodicQFMonitoringSupported returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsPeriodicQFMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsPeriodicQFMonitoringSupported) {
		return true
	}

	return false
}

// SetIsPeriodicQFMonitoringSupported gets a reference to the given bool and assigns it to the IsPeriodicQFMonitoringSupported field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsPeriodicQFMonitoringSupported(v bool) {
	o.IsPeriodicQFMonitoringSupported = &v
}

// GetIsSessionReleasedQFMonitoringSupported returns the IsSessionReleasedQFMonitoringSupported field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsSessionReleasedQFMonitoringSupported() bool {
	if o == nil || IsNil(o.IsSessionReleasedQFMonitoringSupported) {
		var ret bool
		return ret
	}
	return *o.IsSessionReleasedQFMonitoringSupported
}

// GetIsSessionReleasedQFMonitoringSupportedOk returns a tuple with the IsSessionReleasedQFMonitoringSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetIsSessionReleasedQFMonitoringSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSessionReleasedQFMonitoringSupported) {
		return nil, false
	}
	return o.IsSessionReleasedQFMonitoringSupported, true
}

// HasIsSessionReleasedQFMonitoringSupported returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasIsSessionReleasedQFMonitoringSupported() bool {
	if o != nil && !IsNil(o.IsSessionReleasedQFMonitoringSupported) {
		return true
	}

	return false
}

// SetIsSessionReleasedQFMonitoringSupported gets a reference to the given bool and assigns it to the IsSessionReleasedQFMonitoringSupported field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetIsSessionReleasedQFMonitoringSupported(v bool) {
	o.IsSessionReleasedQFMonitoringSupported = &v
}

// GetQFPacketDelayThresholds returns the QFPacketDelayThresholds field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFPacketDelayThresholds() QFPacketDelayThresholdsType {
	if o == nil || IsNil(o.QFPacketDelayThresholds) {
		var ret QFPacketDelayThresholdsType
		return ret
	}
	return *o.QFPacketDelayThresholds
}

// GetQFPacketDelayThresholdsOk returns a tuple with the QFPacketDelayThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFPacketDelayThresholdsOk() (*QFPacketDelayThresholdsType, bool) {
	if o == nil || IsNil(o.QFPacketDelayThresholds) {
		return nil, false
	}
	return o.QFPacketDelayThresholds, true
}

// HasQFPacketDelayThresholds returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFPacketDelayThresholds() bool {
	if o != nil && !IsNil(o.QFPacketDelayThresholds) {
		return true
	}

	return false
}

// SetQFPacketDelayThresholds gets a reference to the given QFPacketDelayThresholdsType and assigns it to the QFPacketDelayThresholds field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFPacketDelayThresholds(v QFPacketDelayThresholdsType) {
	o.QFPacketDelayThresholds = &v
}

// GetQFMinimumWaitTime returns the QFMinimumWaitTime field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMinimumWaitTime() int32 {
	if o == nil || IsNil(o.QFMinimumWaitTime) {
		var ret int32
		return ret
	}
	return *o.QFMinimumWaitTime
}

// GetQFMinimumWaitTimeOk returns a tuple with the QFMinimumWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMinimumWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.QFMinimumWaitTime) {
		return nil, false
	}
	return o.QFMinimumWaitTime, true
}

// HasQFMinimumWaitTime returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFMinimumWaitTime() bool {
	if o != nil && !IsNil(o.QFMinimumWaitTime) {
		return true
	}

	return false
}

// SetQFMinimumWaitTime gets a reference to the given int32 and assigns it to the QFMinimumWaitTime field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFMinimumWaitTime(v int32) {
	o.QFMinimumWaitTime = &v
}

// GetQFMeasurementPeriod returns the QFMeasurementPeriod field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMeasurementPeriod() int32 {
	if o == nil || IsNil(o.QFMeasurementPeriod) {
		var ret int32
		return ret
	}
	return *o.QFMeasurementPeriod
}

// GetQFMeasurementPeriodOk returns a tuple with the QFMeasurementPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) GetQFMeasurementPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.QFMeasurementPeriod) {
		return nil, false
	}
	return o.QFMeasurementPeriod, true
}

// HasQFMeasurementPeriod returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) HasQFMeasurementPeriod() bool {
	if o != nil && !IsNil(o.QFMeasurementPeriod) {
		return true
	}

	return false
}

// SetQFMeasurementPeriod gets a reference to the given int32 and assigns it to the QFMeasurementPeriod field.
func (o *QFQoSMonitoringControlSingleAllOfAttributesAllOf) SetQFMeasurementPeriod(v int32) {
	o.QFMeasurementPeriod = &v
}

func (o QFQoSMonitoringControlSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QFQoSMonitoringControlSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QFQoSMonitoringState) {
		toSerialize["qFQoSMonitoringState"] = o.QFQoSMonitoringState
	}
	if !IsNil(o.QFMonitoredSNSSAIs) {
		toSerialize["qFMonitoredSNSSAIs"] = o.QFMonitoredSNSSAIs
	}
	if !IsNil(o.QFMonitored5QIs) {
		toSerialize["qFMonitored5QIs"] = o.QFMonitored5QIs
	}
	if !IsNil(o.IsEventTriggeredQFMonitoringSupported) {
		toSerialize["isEventTriggeredQFMonitoringSupported"] = o.IsEventTriggeredQFMonitoringSupported
	}
	if !IsNil(o.IsPeriodicQFMonitoringSupported) {
		toSerialize["isPeriodicQFMonitoringSupported"] = o.IsPeriodicQFMonitoringSupported
	}
	if !IsNil(o.IsSessionReleasedQFMonitoringSupported) {
		toSerialize["isSessionReleasedQFMonitoringSupported"] = o.IsSessionReleasedQFMonitoringSupported
	}
	if !IsNil(o.QFPacketDelayThresholds) {
		toSerialize["qFPacketDelayThresholds"] = o.QFPacketDelayThresholds
	}
	if !IsNil(o.QFMinimumWaitTime) {
		toSerialize["qFMinimumWaitTime"] = o.QFMinimumWaitTime
	}
	if !IsNil(o.QFMeasurementPeriod) {
		toSerialize["qFMeasurementPeriod"] = o.QFMeasurementPeriod
	}
	return toSerialize, nil
}

type NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf struct {
	value *QFQoSMonitoringControlSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) Get() *QFQoSMonitoringControlSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) Set(val *QFQoSMonitoringControlSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQFQoSMonitoringControlSingleAllOfAttributesAllOf(val *QFQoSMonitoringControlSingleAllOfAttributesAllOf) *NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf {
	return &NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQFQoSMonitoringControlSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
