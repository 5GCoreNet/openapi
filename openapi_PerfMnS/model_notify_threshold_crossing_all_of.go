/*
TS 28.532 Performance Threshold Monitoring MnS

OAS 3.0.1 definition of the Performance Threshold Monitoring MnS © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMnS

import (
	"encoding/json"
)

// checks if the NotifyThresholdCrossingAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyThresholdCrossingAllOf{}

// NotifyThresholdCrossingAllOf struct for NotifyThresholdCrossingAllOf
type NotifyThresholdCrossingAllOf struct {
	ObservedPerfMetricName      *string              `json:"observedPerfMetricName,omitempty"`
	ObservedPerfMetricValue     *PerfMetricValue     `json:"observedPerfMetricValue,omitempty"`
	ObservedPerfMetricDirection *PerfMetricDirection `json:"observedPerfMetricDirection,omitempty"`
	ThresholdValue              *PerfMetricValue     `json:"thresholdValue,omitempty"`
	Hysteresis                  *PerfMetricValue     `json:"hysteresis,omitempty"`
	MonitorGranularityPeriod    *int32               `json:"monitorGranularityPeriod,omitempty"`
	AdditionalText              *string              `json:"additionalText,omitempty"`
}

// NewNotifyThresholdCrossingAllOf instantiates a new NotifyThresholdCrossingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyThresholdCrossingAllOf() *NotifyThresholdCrossingAllOf {
	this := NotifyThresholdCrossingAllOf{}
	return &this
}

// NewNotifyThresholdCrossingAllOfWithDefaults instantiates a new NotifyThresholdCrossingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyThresholdCrossingAllOfWithDefaults() *NotifyThresholdCrossingAllOf {
	this := NotifyThresholdCrossingAllOf{}
	return &this
}

// GetObservedPerfMetricName returns the ObservedPerfMetricName field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricName() string {
	if o == nil || IsNil(o.ObservedPerfMetricName) {
		var ret string
		return ret
	}
	return *o.ObservedPerfMetricName
}

// GetObservedPerfMetricNameOk returns a tuple with the ObservedPerfMetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObservedPerfMetricName) {
		return nil, false
	}
	return o.ObservedPerfMetricName, true
}

// HasObservedPerfMetricName returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricName() bool {
	if o != nil && !IsNil(o.ObservedPerfMetricName) {
		return true
	}

	return false
}

// SetObservedPerfMetricName gets a reference to the given string and assigns it to the ObservedPerfMetricName field.
func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricName(v string) {
	o.ObservedPerfMetricName = &v
}

// GetObservedPerfMetricValue returns the ObservedPerfMetricValue field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricValue() PerfMetricValue {
	if o == nil || IsNil(o.ObservedPerfMetricValue) {
		var ret PerfMetricValue
		return ret
	}
	return *o.ObservedPerfMetricValue
}

// GetObservedPerfMetricValueOk returns a tuple with the ObservedPerfMetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricValueOk() (*PerfMetricValue, bool) {
	if o == nil || IsNil(o.ObservedPerfMetricValue) {
		return nil, false
	}
	return o.ObservedPerfMetricValue, true
}

// HasObservedPerfMetricValue returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricValue() bool {
	if o != nil && !IsNil(o.ObservedPerfMetricValue) {
		return true
	}

	return false
}

// SetObservedPerfMetricValue gets a reference to the given PerfMetricValue and assigns it to the ObservedPerfMetricValue field.
func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricValue(v PerfMetricValue) {
	o.ObservedPerfMetricValue = &v
}

// GetObservedPerfMetricDirection returns the ObservedPerfMetricDirection field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricDirection() PerfMetricDirection {
	if o == nil || IsNil(o.ObservedPerfMetricDirection) {
		var ret PerfMetricDirection
		return ret
	}
	return *o.ObservedPerfMetricDirection
}

// GetObservedPerfMetricDirectionOk returns a tuple with the ObservedPerfMetricDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricDirectionOk() (*PerfMetricDirection, bool) {
	if o == nil || IsNil(o.ObservedPerfMetricDirection) {
		return nil, false
	}
	return o.ObservedPerfMetricDirection, true
}

// HasObservedPerfMetricDirection returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricDirection() bool {
	if o != nil && !IsNil(o.ObservedPerfMetricDirection) {
		return true
	}

	return false
}

// SetObservedPerfMetricDirection gets a reference to the given PerfMetricDirection and assigns it to the ObservedPerfMetricDirection field.
func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricDirection(v PerfMetricDirection) {
	o.ObservedPerfMetricDirection = &v
}

// GetThresholdValue returns the ThresholdValue field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetThresholdValue() PerfMetricValue {
	if o == nil || IsNil(o.ThresholdValue) {
		var ret PerfMetricValue
		return ret
	}
	return *o.ThresholdValue
}

// GetThresholdValueOk returns a tuple with the ThresholdValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetThresholdValueOk() (*PerfMetricValue, bool) {
	if o == nil || IsNil(o.ThresholdValue) {
		return nil, false
	}
	return o.ThresholdValue, true
}

// HasThresholdValue returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasThresholdValue() bool {
	if o != nil && !IsNil(o.ThresholdValue) {
		return true
	}

	return false
}

// SetThresholdValue gets a reference to the given PerfMetricValue and assigns it to the ThresholdValue field.
func (o *NotifyThresholdCrossingAllOf) SetThresholdValue(v PerfMetricValue) {
	o.ThresholdValue = &v
}

// GetHysteresis returns the Hysteresis field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetHysteresis() PerfMetricValue {
	if o == nil || IsNil(o.Hysteresis) {
		var ret PerfMetricValue
		return ret
	}
	return *o.Hysteresis
}

// GetHysteresisOk returns a tuple with the Hysteresis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetHysteresisOk() (*PerfMetricValue, bool) {
	if o == nil || IsNil(o.Hysteresis) {
		return nil, false
	}
	return o.Hysteresis, true
}

// HasHysteresis returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasHysteresis() bool {
	if o != nil && !IsNil(o.Hysteresis) {
		return true
	}

	return false
}

// SetHysteresis gets a reference to the given PerfMetricValue and assigns it to the Hysteresis field.
func (o *NotifyThresholdCrossingAllOf) SetHysteresis(v PerfMetricValue) {
	o.Hysteresis = &v
}

// GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetMonitorGranularityPeriod() int32 {
	if o == nil || IsNil(o.MonitorGranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.MonitorGranularityPeriod
}

// GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetMonitorGranularityPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.MonitorGranularityPeriod) {
		return nil, false
	}
	return o.MonitorGranularityPeriod, true
}

// HasMonitorGranularityPeriod returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasMonitorGranularityPeriod() bool {
	if o != nil && !IsNil(o.MonitorGranularityPeriod) {
		return true
	}

	return false
}

// SetMonitorGranularityPeriod gets a reference to the given int32 and assigns it to the MonitorGranularityPeriod field.
func (o *NotifyThresholdCrossingAllOf) SetMonitorGranularityPeriod(v int32) {
	o.MonitorGranularityPeriod = &v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyThresholdCrossingAllOf) GetAdditionalText() string {
	if o == nil || IsNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyThresholdCrossingAllOf) GetAdditionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyThresholdCrossingAllOf) HasAdditionalText() bool {
	if o != nil && !IsNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyThresholdCrossingAllOf) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

func (o NotifyThresholdCrossingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyThresholdCrossingAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObservedPerfMetricName) {
		toSerialize["observedPerfMetricName"] = o.ObservedPerfMetricName
	}
	if !IsNil(o.ObservedPerfMetricValue) {
		toSerialize["observedPerfMetricValue"] = o.ObservedPerfMetricValue
	}
	if !IsNil(o.ObservedPerfMetricDirection) {
		toSerialize["observedPerfMetricDirection"] = o.ObservedPerfMetricDirection
	}
	if !IsNil(o.ThresholdValue) {
		toSerialize["thresholdValue"] = o.ThresholdValue
	}
	if !IsNil(o.Hysteresis) {
		toSerialize["hysteresis"] = o.Hysteresis
	}
	if !IsNil(o.MonitorGranularityPeriod) {
		toSerialize["monitorGranularityPeriod"] = o.MonitorGranularityPeriod
	}
	if !IsNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	return toSerialize, nil
}

type NullableNotifyThresholdCrossingAllOf struct {
	value *NotifyThresholdCrossingAllOf
	isSet bool
}

func (v NullableNotifyThresholdCrossingAllOf) Get() *NotifyThresholdCrossingAllOf {
	return v.value
}

func (v *NullableNotifyThresholdCrossingAllOf) Set(val *NotifyThresholdCrossingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyThresholdCrossingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyThresholdCrossingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyThresholdCrossingAllOf(val *NotifyThresholdCrossingAllOf) *NullableNotifyThresholdCrossingAllOf {
	return &NullableNotifyThresholdCrossingAllOf{value: val, isSet: true}
}

func (v NullableNotifyThresholdCrossingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyThresholdCrossingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
