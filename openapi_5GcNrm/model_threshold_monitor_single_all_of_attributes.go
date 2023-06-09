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

// checks if the ThresholdMonitorSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThresholdMonitorSingleAllOfAttributes{}

// ThresholdMonitorSingleAllOfAttributes struct for ThresholdMonitorSingleAllOfAttributes
type ThresholdMonitorSingleAllOfAttributes struct {
	AdministrativeState      *AdministrativeState `json:"administrativeState,omitempty"`
	OperationalState         *OperationalState    `json:"operationalState,omitempty"`
	PerformanceMetrics       []string             `json:"performanceMetrics,omitempty"`
	ThresholdInfoList        []ThresholdInfo      `json:"thresholdInfoList,omitempty"`
	MonitorGranularityPeriod *int32               `json:"monitorGranularityPeriod,omitempty"`
	ObjectInstances          []string             `json:"objectInstances,omitempty"`
	RootObjectInstances      []string             `json:"rootObjectInstances,omitempty"`
}

// NewThresholdMonitorSingleAllOfAttributes instantiates a new ThresholdMonitorSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdMonitorSingleAllOfAttributes() *ThresholdMonitorSingleAllOfAttributes {
	this := ThresholdMonitorSingleAllOfAttributes{}
	return &this
}

// NewThresholdMonitorSingleAllOfAttributesWithDefaults instantiates a new ThresholdMonitorSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdMonitorSingleAllOfAttributesWithDefaults() *ThresholdMonitorSingleAllOfAttributes {
	this := ThresholdMonitorSingleAllOfAttributes{}
	return &this
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || IsNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetPerformanceMetrics returns the PerformanceMetrics field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetPerformanceMetrics() []string {
	if o == nil || IsNil(o.PerformanceMetrics) {
		var ret []string
		return ret
	}
	return o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetPerformanceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.PerformanceMetrics) {
		return nil, false
	}
	return o.PerformanceMetrics, true
}

// HasPerformanceMetrics returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasPerformanceMetrics() bool {
	if o != nil && !IsNil(o.PerformanceMetrics) {
		return true
	}

	return false
}

// SetPerformanceMetrics gets a reference to the given []string and assigns it to the PerformanceMetrics field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetPerformanceMetrics(v []string) {
	o.PerformanceMetrics = v
}

// GetThresholdInfoList returns the ThresholdInfoList field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetThresholdInfoList() []ThresholdInfo {
	if o == nil || IsNil(o.ThresholdInfoList) {
		var ret []ThresholdInfo
		return ret
	}
	return o.ThresholdInfoList
}

// GetThresholdInfoListOk returns a tuple with the ThresholdInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetThresholdInfoListOk() ([]ThresholdInfo, bool) {
	if o == nil || IsNil(o.ThresholdInfoList) {
		return nil, false
	}
	return o.ThresholdInfoList, true
}

// HasThresholdInfoList returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasThresholdInfoList() bool {
	if o != nil && !IsNil(o.ThresholdInfoList) {
		return true
	}

	return false
}

// SetThresholdInfoList gets a reference to the given []ThresholdInfo and assigns it to the ThresholdInfoList field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetThresholdInfoList(v []ThresholdInfo) {
	o.ThresholdInfoList = v
}

// GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetMonitorGranularityPeriod() int32 {
	if o == nil || IsNil(o.MonitorGranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.MonitorGranularityPeriod
}

// GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetMonitorGranularityPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.MonitorGranularityPeriod) {
		return nil, false
	}
	return o.MonitorGranularityPeriod, true
}

// HasMonitorGranularityPeriod returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasMonitorGranularityPeriod() bool {
	if o != nil && !IsNil(o.MonitorGranularityPeriod) {
		return true
	}

	return false
}

// SetMonitorGranularityPeriod gets a reference to the given int32 and assigns it to the MonitorGranularityPeriod field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetMonitorGranularityPeriod(v int32) {
	o.MonitorGranularityPeriod = &v
}

// GetObjectInstances returns the ObjectInstances field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetObjectInstances() []string {
	if o == nil || IsNil(o.ObjectInstances) {
		var ret []string
		return ret
	}
	return o.ObjectInstances
}

// GetObjectInstancesOk returns a tuple with the ObjectInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetObjectInstancesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectInstances) {
		return nil, false
	}
	return o.ObjectInstances, true
}

// HasObjectInstances returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasObjectInstances() bool {
	if o != nil && !IsNil(o.ObjectInstances) {
		return true
	}

	return false
}

// SetObjectInstances gets a reference to the given []string and assigns it to the ObjectInstances field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetObjectInstances(v []string) {
	o.ObjectInstances = v
}

// GetRootObjectInstances returns the RootObjectInstances field value if set, zero value otherwise.
func (o *ThresholdMonitorSingleAllOfAttributes) GetRootObjectInstances() []string {
	if o == nil || IsNil(o.RootObjectInstances) {
		var ret []string
		return ret
	}
	return o.RootObjectInstances
}

// GetRootObjectInstancesOk returns a tuple with the RootObjectInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) GetRootObjectInstancesOk() ([]string, bool) {
	if o == nil || IsNil(o.RootObjectInstances) {
		return nil, false
	}
	return o.RootObjectInstances, true
}

// HasRootObjectInstances returns a boolean if a field has been set.
func (o *ThresholdMonitorSingleAllOfAttributes) HasRootObjectInstances() bool {
	if o != nil && !IsNil(o.RootObjectInstances) {
		return true
	}

	return false
}

// SetRootObjectInstances gets a reference to the given []string and assigns it to the RootObjectInstances field.
func (o *ThresholdMonitorSingleAllOfAttributes) SetRootObjectInstances(v []string) {
	o.RootObjectInstances = v
}

func (o ThresholdMonitorSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThresholdMonitorSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !IsNil(o.PerformanceMetrics) {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	if !IsNil(o.ThresholdInfoList) {
		toSerialize["thresholdInfoList"] = o.ThresholdInfoList
	}
	if !IsNil(o.MonitorGranularityPeriod) {
		toSerialize["monitorGranularityPeriod"] = o.MonitorGranularityPeriod
	}
	if !IsNil(o.ObjectInstances) {
		toSerialize["objectInstances"] = o.ObjectInstances
	}
	if !IsNil(o.RootObjectInstances) {
		toSerialize["rootObjectInstances"] = o.RootObjectInstances
	}
	return toSerialize, nil
}

type NullableThresholdMonitorSingleAllOfAttributes struct {
	value *ThresholdMonitorSingleAllOfAttributes
	isSet bool
}

func (v NullableThresholdMonitorSingleAllOfAttributes) Get() *ThresholdMonitorSingleAllOfAttributes {
	return v.value
}

func (v *NullableThresholdMonitorSingleAllOfAttributes) Set(val *ThresholdMonitorSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdMonitorSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdMonitorSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdMonitorSingleAllOfAttributes(val *ThresholdMonitorSingleAllOfAttributes) *NullableThresholdMonitorSingleAllOfAttributes {
	return &NullableThresholdMonitorSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableThresholdMonitorSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdMonitorSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
