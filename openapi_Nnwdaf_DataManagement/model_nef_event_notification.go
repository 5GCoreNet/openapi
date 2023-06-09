/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the NefEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NefEventNotification{}

// NefEventNotification Represents information related to an event to be reported.
type NefEventNotification struct {
	Event NefEvent `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp        time.Time                         `json:"timeStamp"`
	SvcExprcInfos    []ServiceExperienceInfo1          `json:"svcExprcInfos,omitempty"`
	UeMobilityInfos  []UeMobilityInfo                  `json:"ueMobilityInfos,omitempty"`
	UeCommInfos      []UeCommunicationInfo             `json:"ueCommInfos,omitempty"`
	ExcepInfos       []ExceptionInfo                   `json:"excepInfos,omitempty"`
	CongestionInfos  []UserDataCongestionCollection    `json:"congestionInfos,omitempty"`
	PerfDataInfos    []PerformanceDataInfo             `json:"perfDataInfos,omitempty"`
	DispersionInfos  []DispersionCollection1           `json:"dispersionInfos,omitempty"`
	CollBhvrInfs     []CollectiveBehaviourInfo         `json:"collBhvrInfs,omitempty"`
	MsQoeMetrInfos   []MsQoeMetricsCollection          `json:"msQoeMetrInfos,omitempty"`
	MsConsumpInfos   []MsConsumptionCollection         `json:"msConsumpInfos,omitempty"`
	MsNetAssInvInfos []MsNetAssInvocationCollection    `json:"msNetAssInvInfos,omitempty"`
	MsDynPlyInvInfos []MsDynPolicyInvocationCollection `json:"msDynPlyInvInfos,omitempty"`
	MsAccActInfos    []MSAccessActivityCollection      `json:"msAccActInfos,omitempty"`
}

// NewNefEventNotification instantiates a new NefEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefEventNotification(event NefEvent, timeStamp time.Time) *NefEventNotification {
	this := NefEventNotification{}
	this.Event = event
	this.TimeStamp = timeStamp
	return &this
}

// NewNefEventNotificationWithDefaults instantiates a new NefEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefEventNotificationWithDefaults() *NefEventNotification {
	this := NefEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *NefEventNotification) GetEvent() NefEvent {
	if o == nil {
		var ret NefEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetEventOk() (*NefEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NefEventNotification) SetEvent(v NefEvent) {
	o.Event = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *NefEventNotification) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *NefEventNotification) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetSvcExprcInfos returns the SvcExprcInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetSvcExprcInfos() []ServiceExperienceInfo1 {
	if o == nil || IsNil(o.SvcExprcInfos) {
		var ret []ServiceExperienceInfo1
		return ret
	}
	return o.SvcExprcInfos
}

// GetSvcExprcInfosOk returns a tuple with the SvcExprcInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetSvcExprcInfosOk() ([]ServiceExperienceInfo1, bool) {
	if o == nil || IsNil(o.SvcExprcInfos) {
		return nil, false
	}
	return o.SvcExprcInfos, true
}

// HasSvcExprcInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasSvcExprcInfos() bool {
	if o != nil && !IsNil(o.SvcExprcInfos) {
		return true
	}

	return false
}

// SetSvcExprcInfos gets a reference to the given []ServiceExperienceInfo1 and assigns it to the SvcExprcInfos field.
func (o *NefEventNotification) SetSvcExprcInfos(v []ServiceExperienceInfo1) {
	o.SvcExprcInfos = v
}

// GetUeMobilityInfos returns the UeMobilityInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetUeMobilityInfos() []UeMobilityInfo {
	if o == nil || IsNil(o.UeMobilityInfos) {
		var ret []UeMobilityInfo
		return ret
	}
	return o.UeMobilityInfos
}

// GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetUeMobilityInfosOk() ([]UeMobilityInfo, bool) {
	if o == nil || IsNil(o.UeMobilityInfos) {
		return nil, false
	}
	return o.UeMobilityInfos, true
}

// HasUeMobilityInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasUeMobilityInfos() bool {
	if o != nil && !IsNil(o.UeMobilityInfos) {
		return true
	}

	return false
}

// SetUeMobilityInfos gets a reference to the given []UeMobilityInfo and assigns it to the UeMobilityInfos field.
func (o *NefEventNotification) SetUeMobilityInfos(v []UeMobilityInfo) {
	o.UeMobilityInfos = v
}

// GetUeCommInfos returns the UeCommInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetUeCommInfos() []UeCommunicationInfo {
	if o == nil || IsNil(o.UeCommInfos) {
		var ret []UeCommunicationInfo
		return ret
	}
	return o.UeCommInfos
}

// GetUeCommInfosOk returns a tuple with the UeCommInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetUeCommInfosOk() ([]UeCommunicationInfo, bool) {
	if o == nil || IsNil(o.UeCommInfos) {
		return nil, false
	}
	return o.UeCommInfos, true
}

// HasUeCommInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasUeCommInfos() bool {
	if o != nil && !IsNil(o.UeCommInfos) {
		return true
	}

	return false
}

// SetUeCommInfos gets a reference to the given []UeCommunicationInfo and assigns it to the UeCommInfos field.
func (o *NefEventNotification) SetUeCommInfos(v []UeCommunicationInfo) {
	o.UeCommInfos = v
}

// GetExcepInfos returns the ExcepInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetExcepInfos() []ExceptionInfo {
	if o == nil || IsNil(o.ExcepInfos) {
		var ret []ExceptionInfo
		return ret
	}
	return o.ExcepInfos
}

// GetExcepInfosOk returns a tuple with the ExcepInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetExcepInfosOk() ([]ExceptionInfo, bool) {
	if o == nil || IsNil(o.ExcepInfos) {
		return nil, false
	}
	return o.ExcepInfos, true
}

// HasExcepInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasExcepInfos() bool {
	if o != nil && !IsNil(o.ExcepInfos) {
		return true
	}

	return false
}

// SetExcepInfos gets a reference to the given []ExceptionInfo and assigns it to the ExcepInfos field.
func (o *NefEventNotification) SetExcepInfos(v []ExceptionInfo) {
	o.ExcepInfos = v
}

// GetCongestionInfos returns the CongestionInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetCongestionInfos() []UserDataCongestionCollection {
	if o == nil || IsNil(o.CongestionInfos) {
		var ret []UserDataCongestionCollection
		return ret
	}
	return o.CongestionInfos
}

// GetCongestionInfosOk returns a tuple with the CongestionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetCongestionInfosOk() ([]UserDataCongestionCollection, bool) {
	if o == nil || IsNil(o.CongestionInfos) {
		return nil, false
	}
	return o.CongestionInfos, true
}

// HasCongestionInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasCongestionInfos() bool {
	if o != nil && !IsNil(o.CongestionInfos) {
		return true
	}

	return false
}

// SetCongestionInfos gets a reference to the given []UserDataCongestionCollection and assigns it to the CongestionInfos field.
func (o *NefEventNotification) SetCongestionInfos(v []UserDataCongestionCollection) {
	o.CongestionInfos = v
}

// GetPerfDataInfos returns the PerfDataInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetPerfDataInfos() []PerformanceDataInfo {
	if o == nil || IsNil(o.PerfDataInfos) {
		var ret []PerformanceDataInfo
		return ret
	}
	return o.PerfDataInfos
}

// GetPerfDataInfosOk returns a tuple with the PerfDataInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetPerfDataInfosOk() ([]PerformanceDataInfo, bool) {
	if o == nil || IsNil(o.PerfDataInfos) {
		return nil, false
	}
	return o.PerfDataInfos, true
}

// HasPerfDataInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasPerfDataInfos() bool {
	if o != nil && !IsNil(o.PerfDataInfos) {
		return true
	}

	return false
}

// SetPerfDataInfos gets a reference to the given []PerformanceDataInfo and assigns it to the PerfDataInfos field.
func (o *NefEventNotification) SetPerfDataInfos(v []PerformanceDataInfo) {
	o.PerfDataInfos = v
}

// GetDispersionInfos returns the DispersionInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetDispersionInfos() []DispersionCollection1 {
	if o == nil || IsNil(o.DispersionInfos) {
		var ret []DispersionCollection1
		return ret
	}
	return o.DispersionInfos
}

// GetDispersionInfosOk returns a tuple with the DispersionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetDispersionInfosOk() ([]DispersionCollection1, bool) {
	if o == nil || IsNil(o.DispersionInfos) {
		return nil, false
	}
	return o.DispersionInfos, true
}

// HasDispersionInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasDispersionInfos() bool {
	if o != nil && !IsNil(o.DispersionInfos) {
		return true
	}

	return false
}

// SetDispersionInfos gets a reference to the given []DispersionCollection1 and assigns it to the DispersionInfos field.
func (o *NefEventNotification) SetDispersionInfos(v []DispersionCollection1) {
	o.DispersionInfos = v
}

// GetCollBhvrInfs returns the CollBhvrInfs field value if set, zero value otherwise.
func (o *NefEventNotification) GetCollBhvrInfs() []CollectiveBehaviourInfo {
	if o == nil || IsNil(o.CollBhvrInfs) {
		var ret []CollectiveBehaviourInfo
		return ret
	}
	return o.CollBhvrInfs
}

// GetCollBhvrInfsOk returns a tuple with the CollBhvrInfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetCollBhvrInfsOk() ([]CollectiveBehaviourInfo, bool) {
	if o == nil || IsNil(o.CollBhvrInfs) {
		return nil, false
	}
	return o.CollBhvrInfs, true
}

// HasCollBhvrInfs returns a boolean if a field has been set.
func (o *NefEventNotification) HasCollBhvrInfs() bool {
	if o != nil && !IsNil(o.CollBhvrInfs) {
		return true
	}

	return false
}

// SetCollBhvrInfs gets a reference to the given []CollectiveBehaviourInfo and assigns it to the CollBhvrInfs field.
func (o *NefEventNotification) SetCollBhvrInfs(v []CollectiveBehaviourInfo) {
	o.CollBhvrInfs = v
}

// GetMsQoeMetrInfos returns the MsQoeMetrInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetMsQoeMetrInfos() []MsQoeMetricsCollection {
	if o == nil || IsNil(o.MsQoeMetrInfos) {
		var ret []MsQoeMetricsCollection
		return ret
	}
	return o.MsQoeMetrInfos
}

// GetMsQoeMetrInfosOk returns a tuple with the MsQoeMetrInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetMsQoeMetrInfosOk() ([]MsQoeMetricsCollection, bool) {
	if o == nil || IsNil(o.MsQoeMetrInfos) {
		return nil, false
	}
	return o.MsQoeMetrInfos, true
}

// HasMsQoeMetrInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasMsQoeMetrInfos() bool {
	if o != nil && !IsNil(o.MsQoeMetrInfos) {
		return true
	}

	return false
}

// SetMsQoeMetrInfos gets a reference to the given []MsQoeMetricsCollection and assigns it to the MsQoeMetrInfos field.
func (o *NefEventNotification) SetMsQoeMetrInfos(v []MsQoeMetricsCollection) {
	o.MsQoeMetrInfos = v
}

// GetMsConsumpInfos returns the MsConsumpInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetMsConsumpInfos() []MsConsumptionCollection {
	if o == nil || IsNil(o.MsConsumpInfos) {
		var ret []MsConsumptionCollection
		return ret
	}
	return o.MsConsumpInfos
}

// GetMsConsumpInfosOk returns a tuple with the MsConsumpInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetMsConsumpInfosOk() ([]MsConsumptionCollection, bool) {
	if o == nil || IsNil(o.MsConsumpInfos) {
		return nil, false
	}
	return o.MsConsumpInfos, true
}

// HasMsConsumpInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasMsConsumpInfos() bool {
	if o != nil && !IsNil(o.MsConsumpInfos) {
		return true
	}

	return false
}

// SetMsConsumpInfos gets a reference to the given []MsConsumptionCollection and assigns it to the MsConsumpInfos field.
func (o *NefEventNotification) SetMsConsumpInfos(v []MsConsumptionCollection) {
	o.MsConsumpInfos = v
}

// GetMsNetAssInvInfos returns the MsNetAssInvInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetMsNetAssInvInfos() []MsNetAssInvocationCollection {
	if o == nil || IsNil(o.MsNetAssInvInfos) {
		var ret []MsNetAssInvocationCollection
		return ret
	}
	return o.MsNetAssInvInfos
}

// GetMsNetAssInvInfosOk returns a tuple with the MsNetAssInvInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetMsNetAssInvInfosOk() ([]MsNetAssInvocationCollection, bool) {
	if o == nil || IsNil(o.MsNetAssInvInfos) {
		return nil, false
	}
	return o.MsNetAssInvInfos, true
}

// HasMsNetAssInvInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasMsNetAssInvInfos() bool {
	if o != nil && !IsNil(o.MsNetAssInvInfos) {
		return true
	}

	return false
}

// SetMsNetAssInvInfos gets a reference to the given []MsNetAssInvocationCollection and assigns it to the MsNetAssInvInfos field.
func (o *NefEventNotification) SetMsNetAssInvInfos(v []MsNetAssInvocationCollection) {
	o.MsNetAssInvInfos = v
}

// GetMsDynPlyInvInfos returns the MsDynPlyInvInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetMsDynPlyInvInfos() []MsDynPolicyInvocationCollection {
	if o == nil || IsNil(o.MsDynPlyInvInfos) {
		var ret []MsDynPolicyInvocationCollection
		return ret
	}
	return o.MsDynPlyInvInfos
}

// GetMsDynPlyInvInfosOk returns a tuple with the MsDynPlyInvInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetMsDynPlyInvInfosOk() ([]MsDynPolicyInvocationCollection, bool) {
	if o == nil || IsNil(o.MsDynPlyInvInfos) {
		return nil, false
	}
	return o.MsDynPlyInvInfos, true
}

// HasMsDynPlyInvInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasMsDynPlyInvInfos() bool {
	if o != nil && !IsNil(o.MsDynPlyInvInfos) {
		return true
	}

	return false
}

// SetMsDynPlyInvInfos gets a reference to the given []MsDynPolicyInvocationCollection and assigns it to the MsDynPlyInvInfos field.
func (o *NefEventNotification) SetMsDynPlyInvInfos(v []MsDynPolicyInvocationCollection) {
	o.MsDynPlyInvInfos = v
}

// GetMsAccActInfos returns the MsAccActInfos field value if set, zero value otherwise.
func (o *NefEventNotification) GetMsAccActInfos() []MSAccessActivityCollection {
	if o == nil || IsNil(o.MsAccActInfos) {
		var ret []MSAccessActivityCollection
		return ret
	}
	return o.MsAccActInfos
}

// GetMsAccActInfosOk returns a tuple with the MsAccActInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventNotification) GetMsAccActInfosOk() ([]MSAccessActivityCollection, bool) {
	if o == nil || IsNil(o.MsAccActInfos) {
		return nil, false
	}
	return o.MsAccActInfos, true
}

// HasMsAccActInfos returns a boolean if a field has been set.
func (o *NefEventNotification) HasMsAccActInfos() bool {
	if o != nil && !IsNil(o.MsAccActInfos) {
		return true
	}

	return false
}

// SetMsAccActInfos gets a reference to the given []MSAccessActivityCollection and assigns it to the MsAccActInfos field.
func (o *NefEventNotification) SetMsAccActInfos(v []MSAccessActivityCollection) {
	o.MsAccActInfos = v
}

func (o NefEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NefEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["timeStamp"] = o.TimeStamp
	if !IsNil(o.SvcExprcInfos) {
		toSerialize["svcExprcInfos"] = o.SvcExprcInfos
	}
	if !IsNil(o.UeMobilityInfos) {
		toSerialize["ueMobilityInfos"] = o.UeMobilityInfos
	}
	if !IsNil(o.UeCommInfos) {
		toSerialize["ueCommInfos"] = o.UeCommInfos
	}
	if !IsNil(o.ExcepInfos) {
		toSerialize["excepInfos"] = o.ExcepInfos
	}
	if !IsNil(o.CongestionInfos) {
		toSerialize["congestionInfos"] = o.CongestionInfos
	}
	if !IsNil(o.PerfDataInfos) {
		toSerialize["perfDataInfos"] = o.PerfDataInfos
	}
	if !IsNil(o.DispersionInfos) {
		toSerialize["dispersionInfos"] = o.DispersionInfos
	}
	if !IsNil(o.CollBhvrInfs) {
		toSerialize["collBhvrInfs"] = o.CollBhvrInfs
	}
	if !IsNil(o.MsQoeMetrInfos) {
		toSerialize["msQoeMetrInfos"] = o.MsQoeMetrInfos
	}
	if !IsNil(o.MsConsumpInfos) {
		toSerialize["msConsumpInfos"] = o.MsConsumpInfos
	}
	if !IsNil(o.MsNetAssInvInfos) {
		toSerialize["msNetAssInvInfos"] = o.MsNetAssInvInfos
	}
	if !IsNil(o.MsDynPlyInvInfos) {
		toSerialize["msDynPlyInvInfos"] = o.MsDynPlyInvInfos
	}
	if !IsNil(o.MsAccActInfos) {
		toSerialize["msAccActInfos"] = o.MsAccActInfos
	}
	return toSerialize, nil
}

type NullableNefEventNotification struct {
	value *NefEventNotification
	isSet bool
}

func (v NullableNefEventNotification) Get() *NefEventNotification {
	return v.value
}

func (v *NullableNefEventNotification) Set(val *NefEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEventNotification(val *NefEventNotification) *NullableNefEventNotification {
	return &NullableNefEventNotification{value: val, isSet: true}
}

func (v NullableNefEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
