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

// checks if the MDAOutputIEFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAOutputIEFilter{}

// MDAOutputIEFilter struct for MDAOutputIEFilter
type MDAOutputIEFilter struct {
	MDAOutputIEName *string `json:"mDAOutputIEName,omitempty"`
	FilterValue *string `json:"filterValue,omitempty"`
	Threshold *ThresholdInfo1 `json:"threshold,omitempty"`
	AnalyticsPeriod *AnalyticsSchedule `json:"analyticsPeriod,omitempty"`
	TimeOut *time.Time `json:"timeOut,omitempty"`
}

// NewMDAOutputIEFilter instantiates a new MDAOutputIEFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAOutputIEFilter() *MDAOutputIEFilter {
	this := MDAOutputIEFilter{}
	return &this
}

// NewMDAOutputIEFilterWithDefaults instantiates a new MDAOutputIEFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAOutputIEFilterWithDefaults() *MDAOutputIEFilter {
	this := MDAOutputIEFilter{}
	return &this
}

// GetMDAOutputIEName returns the MDAOutputIEName field value if set, zero value otherwise.
func (o *MDAOutputIEFilter) GetMDAOutputIEName() string {
	if o == nil || IsNil(o.MDAOutputIEName) {
		var ret string
		return ret
	}
	return *o.MDAOutputIEName
}

// GetMDAOutputIENameOk returns a tuple with the MDAOutputIEName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputIEFilter) GetMDAOutputIENameOk() (*string, bool) {
	if o == nil || IsNil(o.MDAOutputIEName) {
		return nil, false
	}
	return o.MDAOutputIEName, true
}

// HasMDAOutputIEName returns a boolean if a field has been set.
func (o *MDAOutputIEFilter) HasMDAOutputIEName() bool {
	if o != nil && !IsNil(o.MDAOutputIEName) {
		return true
	}

	return false
}

// SetMDAOutputIEName gets a reference to the given string and assigns it to the MDAOutputIEName field.
func (o *MDAOutputIEFilter) SetMDAOutputIEName(v string) {
	o.MDAOutputIEName = &v
}

// GetFilterValue returns the FilterValue field value if set, zero value otherwise.
func (o *MDAOutputIEFilter) GetFilterValue() string {
	if o == nil || IsNil(o.FilterValue) {
		var ret string
		return ret
	}
	return *o.FilterValue
}

// GetFilterValueOk returns a tuple with the FilterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputIEFilter) GetFilterValueOk() (*string, bool) {
	if o == nil || IsNil(o.FilterValue) {
		return nil, false
	}
	return o.FilterValue, true
}

// HasFilterValue returns a boolean if a field has been set.
func (o *MDAOutputIEFilter) HasFilterValue() bool {
	if o != nil && !IsNil(o.FilterValue) {
		return true
	}

	return false
}

// SetFilterValue gets a reference to the given string and assigns it to the FilterValue field.
func (o *MDAOutputIEFilter) SetFilterValue(v string) {
	o.FilterValue = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MDAOutputIEFilter) GetThreshold() ThresholdInfo1 {
	if o == nil || IsNil(o.Threshold) {
		var ret ThresholdInfo1
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputIEFilter) GetThresholdOk() (*ThresholdInfo1, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MDAOutputIEFilter) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given ThresholdInfo1 and assigns it to the Threshold field.
func (o *MDAOutputIEFilter) SetThreshold(v ThresholdInfo1) {
	o.Threshold = &v
}

// GetAnalyticsPeriod returns the AnalyticsPeriod field value if set, zero value otherwise.
func (o *MDAOutputIEFilter) GetAnalyticsPeriod() AnalyticsSchedule {
	if o == nil || IsNil(o.AnalyticsPeriod) {
		var ret AnalyticsSchedule
		return ret
	}
	return *o.AnalyticsPeriod
}

// GetAnalyticsPeriodOk returns a tuple with the AnalyticsPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputIEFilter) GetAnalyticsPeriodOk() (*AnalyticsSchedule, bool) {
	if o == nil || IsNil(o.AnalyticsPeriod) {
		return nil, false
	}
	return o.AnalyticsPeriod, true
}

// HasAnalyticsPeriod returns a boolean if a field has been set.
func (o *MDAOutputIEFilter) HasAnalyticsPeriod() bool {
	if o != nil && !IsNil(o.AnalyticsPeriod) {
		return true
	}

	return false
}

// SetAnalyticsPeriod gets a reference to the given AnalyticsSchedule and assigns it to the AnalyticsPeriod field.
func (o *MDAOutputIEFilter) SetAnalyticsPeriod(v AnalyticsSchedule) {
	o.AnalyticsPeriod = &v
}

// GetTimeOut returns the TimeOut field value if set, zero value otherwise.
func (o *MDAOutputIEFilter) GetTimeOut() time.Time {
	if o == nil || IsNil(o.TimeOut) {
		var ret time.Time
		return ret
	}
	return *o.TimeOut
}

// GetTimeOutOk returns a tuple with the TimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputIEFilter) GetTimeOutOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeOut) {
		return nil, false
	}
	return o.TimeOut, true
}

// HasTimeOut returns a boolean if a field has been set.
func (o *MDAOutputIEFilter) HasTimeOut() bool {
	if o != nil && !IsNil(o.TimeOut) {
		return true
	}

	return false
}

// SetTimeOut gets a reference to the given time.Time and assigns it to the TimeOut field.
func (o *MDAOutputIEFilter) SetTimeOut(v time.Time) {
	o.TimeOut = &v
}

func (o MDAOutputIEFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAOutputIEFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAOutputIEName) {
		toSerialize["mDAOutputIEName"] = o.MDAOutputIEName
	}
	if !IsNil(o.FilterValue) {
		toSerialize["filterValue"] = o.FilterValue
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.AnalyticsPeriod) {
		toSerialize["analyticsPeriod"] = o.AnalyticsPeriod
	}
	if !IsNil(o.TimeOut) {
		toSerialize["timeOut"] = o.TimeOut
	}
	return toSerialize, nil
}

type NullableMDAOutputIEFilter struct {
	value *MDAOutputIEFilter
	isSet bool
}

func (v NullableMDAOutputIEFilter) Get() *MDAOutputIEFilter {
	return v.value
}

func (v *NullableMDAOutputIEFilter) Set(val *MDAOutputIEFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAOutputIEFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAOutputIEFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAOutputIEFilter(val *MDAOutputIEFilter) *NullableMDAOutputIEFilter {
	return &NullableMDAOutputIEFilter{value: val, isSet: true}
}

func (v NullableMDAOutputIEFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAOutputIEFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

