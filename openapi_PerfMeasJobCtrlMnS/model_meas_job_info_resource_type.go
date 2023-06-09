/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// checks if the MeasJobInfoResourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasJobInfoResourceType{}

// MeasJobInfoResourceType struct for MeasJobInfoResourceType
type MeasJobInfoResourceType struct {
	Href                    *string              `json:"href,omitempty"`
	IOCName                 *string              `json:"iOCName,omitempty"`
	IOCInstanceList         []string             `json:"iOCInstanceList,omitempty"`
	MeasurementCategoryList []string             `json:"measurementCategoryList,omitempty"`
	ReportingMethod         *ReportingMethodType `json:"reportingMethod,omitempty"`
	GranularityPeriod       *int32               `json:"granularityPeriod,omitempty"`
	ReportingPeriod         *int32               `json:"reportingPeriod,omitempty"`
	StartTime               *string              `json:"startTime,omitempty"`
	StopTime                *string              `json:"stopTime,omitempty"`
	Schedule                *ScheduleType        `json:"schedule,omitempty"`
	StreamTarget            *string              `json:"streamTarget,omitempty"`
	Priority                *PriorityType        `json:"priority,omitempty"`
	Reliability             *string              `json:"reliability,omitempty"`
}

// NewMeasJobInfoResourceType instantiates a new MeasJobInfoResourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasJobInfoResourceType() *MeasJobInfoResourceType {
	this := MeasJobInfoResourceType{}
	return &this
}

// NewMeasJobInfoResourceTypeWithDefaults instantiates a new MeasJobInfoResourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasJobInfoResourceTypeWithDefaults() *MeasJobInfoResourceType {
	this := MeasJobInfoResourceType{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *MeasJobInfoResourceType) SetHref(v string) {
	o.Href = &v
}

// GetIOCName returns the IOCName field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetIOCName() string {
	if o == nil || IsNil(o.IOCName) {
		var ret string
		return ret
	}
	return *o.IOCName
}

// GetIOCNameOk returns a tuple with the IOCName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetIOCNameOk() (*string, bool) {
	if o == nil || IsNil(o.IOCName) {
		return nil, false
	}
	return o.IOCName, true
}

// HasIOCName returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasIOCName() bool {
	if o != nil && !IsNil(o.IOCName) {
		return true
	}

	return false
}

// SetIOCName gets a reference to the given string and assigns it to the IOCName field.
func (o *MeasJobInfoResourceType) SetIOCName(v string) {
	o.IOCName = &v
}

// GetIOCInstanceList returns the IOCInstanceList field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetIOCInstanceList() []string {
	if o == nil || IsNil(o.IOCInstanceList) {
		var ret []string
		return ret
	}
	return o.IOCInstanceList
}

// GetIOCInstanceListOk returns a tuple with the IOCInstanceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetIOCInstanceListOk() ([]string, bool) {
	if o == nil || IsNil(o.IOCInstanceList) {
		return nil, false
	}
	return o.IOCInstanceList, true
}

// HasIOCInstanceList returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasIOCInstanceList() bool {
	if o != nil && !IsNil(o.IOCInstanceList) {
		return true
	}

	return false
}

// SetIOCInstanceList gets a reference to the given []string and assigns it to the IOCInstanceList field.
func (o *MeasJobInfoResourceType) SetIOCInstanceList(v []string) {
	o.IOCInstanceList = v
}

// GetMeasurementCategoryList returns the MeasurementCategoryList field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetMeasurementCategoryList() []string {
	if o == nil || IsNil(o.MeasurementCategoryList) {
		var ret []string
		return ret
	}
	return o.MeasurementCategoryList
}

// GetMeasurementCategoryListOk returns a tuple with the MeasurementCategoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetMeasurementCategoryListOk() ([]string, bool) {
	if o == nil || IsNil(o.MeasurementCategoryList) {
		return nil, false
	}
	return o.MeasurementCategoryList, true
}

// HasMeasurementCategoryList returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasMeasurementCategoryList() bool {
	if o != nil && !IsNil(o.MeasurementCategoryList) {
		return true
	}

	return false
}

// SetMeasurementCategoryList gets a reference to the given []string and assigns it to the MeasurementCategoryList field.
func (o *MeasJobInfoResourceType) SetMeasurementCategoryList(v []string) {
	o.MeasurementCategoryList = v
}

// GetReportingMethod returns the ReportingMethod field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetReportingMethod() ReportingMethodType {
	if o == nil || IsNil(o.ReportingMethod) {
		var ret ReportingMethodType
		return ret
	}
	return *o.ReportingMethod
}

// GetReportingMethodOk returns a tuple with the ReportingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetReportingMethodOk() (*ReportingMethodType, bool) {
	if o == nil || IsNil(o.ReportingMethod) {
		return nil, false
	}
	return o.ReportingMethod, true
}

// HasReportingMethod returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasReportingMethod() bool {
	if o != nil && !IsNil(o.ReportingMethod) {
		return true
	}

	return false
}

// SetReportingMethod gets a reference to the given ReportingMethodType and assigns it to the ReportingMethod field.
func (o *MeasJobInfoResourceType) SetReportingMethod(v ReportingMethodType) {
	o.ReportingMethod = &v
}

// GetGranularityPeriod returns the GranularityPeriod field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetGranularityPeriod() int32 {
	if o == nil || IsNil(o.GranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.GranularityPeriod
}

// GetGranularityPeriodOk returns a tuple with the GranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetGranularityPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.GranularityPeriod) {
		return nil, false
	}
	return o.GranularityPeriod, true
}

// HasGranularityPeriod returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasGranularityPeriod() bool {
	if o != nil && !IsNil(o.GranularityPeriod) {
		return true
	}

	return false
}

// SetGranularityPeriod gets a reference to the given int32 and assigns it to the GranularityPeriod field.
func (o *MeasJobInfoResourceType) SetGranularityPeriod(v int32) {
	o.GranularityPeriod = &v
}

// GetReportingPeriod returns the ReportingPeriod field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetReportingPeriod() int32 {
	if o == nil || IsNil(o.ReportingPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportingPeriod
}

// GetReportingPeriodOk returns a tuple with the ReportingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetReportingPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportingPeriod) {
		return nil, false
	}
	return o.ReportingPeriod, true
}

// HasReportingPeriod returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasReportingPeriod() bool {
	if o != nil && !IsNil(o.ReportingPeriod) {
		return true
	}

	return false
}

// SetReportingPeriod gets a reference to the given int32 and assigns it to the ReportingPeriod field.
func (o *MeasJobInfoResourceType) SetReportingPeriod(v int32) {
	o.ReportingPeriod = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *MeasJobInfoResourceType) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetStopTime() string {
	if o == nil || IsNil(o.StopTime) {
		var ret string
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetStopTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StopTime) {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasStopTime() bool {
	if o != nil && !IsNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given string and assigns it to the StopTime field.
func (o *MeasJobInfoResourceType) SetStopTime(v string) {
	o.StopTime = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetSchedule() ScheduleType {
	if o == nil || IsNil(o.Schedule) {
		var ret ScheduleType
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetScheduleOk() (*ScheduleType, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ScheduleType and assigns it to the Schedule field.
func (o *MeasJobInfoResourceType) SetSchedule(v ScheduleType) {
	o.Schedule = &v
}

// GetStreamTarget returns the StreamTarget field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetStreamTarget() string {
	if o == nil || IsNil(o.StreamTarget) {
		var ret string
		return ret
	}
	return *o.StreamTarget
}

// GetStreamTargetOk returns a tuple with the StreamTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetStreamTargetOk() (*string, bool) {
	if o == nil || IsNil(o.StreamTarget) {
		return nil, false
	}
	return o.StreamTarget, true
}

// HasStreamTarget returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasStreamTarget() bool {
	if o != nil && !IsNil(o.StreamTarget) {
		return true
	}

	return false
}

// SetStreamTarget gets a reference to the given string and assigns it to the StreamTarget field.
func (o *MeasJobInfoResourceType) SetStreamTarget(v string) {
	o.StreamTarget = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetPriority() PriorityType {
	if o == nil || IsNil(o.Priority) {
		var ret PriorityType
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetPriorityOk() (*PriorityType, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given PriorityType and assigns it to the Priority field.
func (o *MeasJobInfoResourceType) SetPriority(v PriorityType) {
	o.Priority = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *MeasJobInfoResourceType) GetReliability() string {
	if o == nil || IsNil(o.Reliability) {
		var ret string
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasJobInfoResourceType) GetReliabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Reliability) {
		return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *MeasJobInfoResourceType) HasReliability() bool {
	if o != nil && !IsNil(o.Reliability) {
		return true
	}

	return false
}

// SetReliability gets a reference to the given string and assigns it to the Reliability field.
func (o *MeasJobInfoResourceType) SetReliability(v string) {
	o.Reliability = &v
}

func (o MeasJobInfoResourceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasJobInfoResourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.IOCName) {
		toSerialize["iOCName"] = o.IOCName
	}
	if !IsNil(o.IOCInstanceList) {
		toSerialize["iOCInstanceList"] = o.IOCInstanceList
	}
	if !IsNil(o.MeasurementCategoryList) {
		toSerialize["measurementCategoryList"] = o.MeasurementCategoryList
	}
	if !IsNil(o.ReportingMethod) {
		toSerialize["reportingMethod"] = o.ReportingMethod
	}
	if !IsNil(o.GranularityPeriod) {
		toSerialize["granularityPeriod"] = o.GranularityPeriod
	}
	if !IsNil(o.ReportingPeriod) {
		toSerialize["reportingPeriod"] = o.ReportingPeriod
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.StreamTarget) {
		toSerialize["streamTarget"] = o.StreamTarget
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Reliability) {
		toSerialize["reliability"] = o.Reliability
	}
	return toSerialize, nil
}

type NullableMeasJobInfoResourceType struct {
	value *MeasJobInfoResourceType
	isSet bool
}

func (v NullableMeasJobInfoResourceType) Get() *MeasJobInfoResourceType {
	return v.value
}

func (v *NullableMeasJobInfoResourceType) Set(val *MeasJobInfoResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasJobInfoResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasJobInfoResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasJobInfoResourceType(val *MeasJobInfoResourceType) *NullableMeasJobInfoResourceType {
	return &NullableMeasJobInfoResourceType{value: val, isSet: true}
}

func (v NullableMeasJobInfoResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasJobInfoResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
