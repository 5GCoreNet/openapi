/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EventExposure

import (
	"encoding/json"
	"time"
)

// checks if the ReportingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingInformation{}

// ReportingInformation Represents the type of reporting that the subscription requires.
type ReportingInformation struct {
	ImmRep      *bool               `json:"immRep,omitempty"`
	NotifMethod *NotificationMethod `json:"notifMethod,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr *int32 `json:"maxReportNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MonDur *time.Time `json:"monDur,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio *int32 `json:"sampRatio,omitempty"`
	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`
	// indicating a time in seconds.
	GrpRepTime *int32            `json:"grpRepTime,omitempty"`
	NotifFlag  *NotificationFlag `json:"notifFlag,omitempty"`
}

// NewReportingInformation instantiates a new ReportingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingInformation() *ReportingInformation {
	this := ReportingInformation{}
	return &this
}

// NewReportingInformationWithDefaults instantiates a new ReportingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingInformationWithDefaults() *ReportingInformation {
	this := ReportingInformation{}
	return &this
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *ReportingInformation) GetImmRep() bool {
	if o == nil || IsNil(o.ImmRep) {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetImmRepOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmRep) {
		return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *ReportingInformation) HasImmRep() bool {
	if o != nil && !IsNil(o.ImmRep) {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *ReportingInformation) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *ReportingInformation) GetNotifMethod() NotificationMethod {
	if o == nil || IsNil(o.NotifMethod) {
		var ret NotificationMethod
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetNotifMethodOk() (*NotificationMethod, bool) {
	if o == nil || IsNil(o.NotifMethod) {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *ReportingInformation) HasNotifMethod() bool {
	if o != nil && !IsNil(o.NotifMethod) {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given NotificationMethod and assigns it to the NotifMethod field.
func (o *ReportingInformation) SetNotifMethod(v NotificationMethod) {
	o.NotifMethod = &v
}

// GetMaxReportNbr returns the MaxReportNbr field value if set, zero value otherwise.
func (o *ReportingInformation) GetMaxReportNbr() int32 {
	if o == nil || IsNil(o.MaxReportNbr) {
		var ret int32
		return ret
	}
	return *o.MaxReportNbr
}

// GetMaxReportNbrOk returns a tuple with the MaxReportNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetMaxReportNbrOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReportNbr) {
		return nil, false
	}
	return o.MaxReportNbr, true
}

// HasMaxReportNbr returns a boolean if a field has been set.
func (o *ReportingInformation) HasMaxReportNbr() bool {
	if o != nil && !IsNil(o.MaxReportNbr) {
		return true
	}

	return false
}

// SetMaxReportNbr gets a reference to the given int32 and assigns it to the MaxReportNbr field.
func (o *ReportingInformation) SetMaxReportNbr(v int32) {
	o.MaxReportNbr = &v
}

// GetMonDur returns the MonDur field value if set, zero value otherwise.
func (o *ReportingInformation) GetMonDur() time.Time {
	if o == nil || IsNil(o.MonDur) {
		var ret time.Time
		return ret
	}
	return *o.MonDur
}

// GetMonDurOk returns a tuple with the MonDur field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetMonDurOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MonDur) {
		return nil, false
	}
	return o.MonDur, true
}

// HasMonDur returns a boolean if a field has been set.
func (o *ReportingInformation) HasMonDur() bool {
	if o != nil && !IsNil(o.MonDur) {
		return true
	}

	return false
}

// SetMonDur gets a reference to the given time.Time and assigns it to the MonDur field.
func (o *ReportingInformation) SetMonDur(v time.Time) {
	o.MonDur = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *ReportingInformation) GetRepPeriod() int32 {
	if o == nil || IsNil(o.RepPeriod) {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RepPeriod) {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *ReportingInformation) HasRepPeriod() bool {
	if o != nil && !IsNil(o.RepPeriod) {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *ReportingInformation) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *ReportingInformation) GetSampRatio() int32 {
	if o == nil || IsNil(o.SampRatio) {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetSampRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.SampRatio) {
		return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *ReportingInformation) HasSampRatio() bool {
	if o != nil && !IsNil(o.SampRatio) {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *ReportingInformation) SetSampRatio(v int32) {
	o.SampRatio = &v
}

// GetPartitionCriteria returns the PartitionCriteria field value if set, zero value otherwise.
func (o *ReportingInformation) GetPartitionCriteria() []PartitioningCriteria {
	if o == nil || IsNil(o.PartitionCriteria) {
		var ret []PartitioningCriteria
		return ret
	}
	return o.PartitionCriteria
}

// GetPartitionCriteriaOk returns a tuple with the PartitionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetPartitionCriteriaOk() ([]PartitioningCriteria, bool) {
	if o == nil || IsNil(o.PartitionCriteria) {
		return nil, false
	}
	return o.PartitionCriteria, true
}

// HasPartitionCriteria returns a boolean if a field has been set.
func (o *ReportingInformation) HasPartitionCriteria() bool {
	if o != nil && !IsNil(o.PartitionCriteria) {
		return true
	}

	return false
}

// SetPartitionCriteria gets a reference to the given []PartitioningCriteria and assigns it to the PartitionCriteria field.
func (o *ReportingInformation) SetPartitionCriteria(v []PartitioningCriteria) {
	o.PartitionCriteria = v
}

// GetGrpRepTime returns the GrpRepTime field value if set, zero value otherwise.
func (o *ReportingInformation) GetGrpRepTime() int32 {
	if o == nil || IsNil(o.GrpRepTime) {
		var ret int32
		return ret
	}
	return *o.GrpRepTime
}

// GetGrpRepTimeOk returns a tuple with the GrpRepTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetGrpRepTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GrpRepTime) {
		return nil, false
	}
	return o.GrpRepTime, true
}

// HasGrpRepTime returns a boolean if a field has been set.
func (o *ReportingInformation) HasGrpRepTime() bool {
	if o != nil && !IsNil(o.GrpRepTime) {
		return true
	}

	return false
}

// SetGrpRepTime gets a reference to the given int32 and assigns it to the GrpRepTime field.
func (o *ReportingInformation) SetGrpRepTime(v int32) {
	o.GrpRepTime = &v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *ReportingInformation) GetNotifFlag() NotificationFlag {
	if o == nil || IsNil(o.NotifFlag) {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || IsNil(o.NotifFlag) {
		return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *ReportingInformation) HasNotifFlag() bool {
	if o != nil && !IsNil(o.NotifFlag) {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *ReportingInformation) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

func (o ReportingInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImmRep) {
		toSerialize["immRep"] = o.ImmRep
	}
	if !IsNil(o.NotifMethod) {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if !IsNil(o.MaxReportNbr) {
		toSerialize["maxReportNbr"] = o.MaxReportNbr
	}
	if !IsNil(o.MonDur) {
		toSerialize["monDur"] = o.MonDur
	}
	if !IsNil(o.RepPeriod) {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if !IsNil(o.SampRatio) {
		toSerialize["sampRatio"] = o.SampRatio
	}
	if !IsNil(o.PartitionCriteria) {
		toSerialize["partitionCriteria"] = o.PartitionCriteria
	}
	if !IsNil(o.GrpRepTime) {
		toSerialize["grpRepTime"] = o.GrpRepTime
	}
	if !IsNil(o.NotifFlag) {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	return toSerialize, nil
}

type NullableReportingInformation struct {
	value *ReportingInformation
	isSet bool
}

func (v NullableReportingInformation) Get() *ReportingInformation {
	return v.value
}

func (v *NullableReportingInformation) Set(val *ReportingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingInformation(val *ReportingInformation) *NullableReportingInformation {
	return &NullableReportingInformation{value: val, isSet: true}
}

func (v NullableReportingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
