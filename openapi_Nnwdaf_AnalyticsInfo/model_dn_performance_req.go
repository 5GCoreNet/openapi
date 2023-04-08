/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the DnPerformanceReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnPerformanceReq{}

// DnPerformanceReq Represents other DN performance analytics requirements.
type DnPerformanceReq struct {
	DnPerfOrderCriter *DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty"`
	Order *MatchingDirection `json:"order,omitempty"`
	ReportThresholds []ThresholdLevel `json:"reportThresholds,omitempty"`
}

// NewDnPerformanceReq instantiates a new DnPerformanceReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnPerformanceReq() *DnPerformanceReq {
	this := DnPerformanceReq{}
	return &this
}

// NewDnPerformanceReqWithDefaults instantiates a new DnPerformanceReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnPerformanceReqWithDefaults() *DnPerformanceReq {
	this := DnPerformanceReq{}
	return &this
}

// GetDnPerfOrderCriter returns the DnPerfOrderCriter field value if set, zero value otherwise.
func (o *DnPerformanceReq) GetDnPerfOrderCriter() DnPerfOrderingCriterion {
	if o == nil || isNil(o.DnPerfOrderCriter) {
		var ret DnPerfOrderingCriterion
		return ret
	}
	return *o.DnPerfOrderCriter
}

// GetDnPerfOrderCriterOk returns a tuple with the DnPerfOrderCriter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerformanceReq) GetDnPerfOrderCriterOk() (*DnPerfOrderingCriterion, bool) {
	if o == nil || isNil(o.DnPerfOrderCriter) {
		return nil, false
	}
	return o.DnPerfOrderCriter, true
}

// HasDnPerfOrderCriter returns a boolean if a field has been set.
func (o *DnPerformanceReq) HasDnPerfOrderCriter() bool {
	if o != nil && !isNil(o.DnPerfOrderCriter) {
		return true
	}

	return false
}

// SetDnPerfOrderCriter gets a reference to the given DnPerfOrderingCriterion and assigns it to the DnPerfOrderCriter field.
func (o *DnPerformanceReq) SetDnPerfOrderCriter(v DnPerfOrderingCriterion) {
	o.DnPerfOrderCriter = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DnPerformanceReq) GetOrder() MatchingDirection {
	if o == nil || isNil(o.Order) {
		var ret MatchingDirection
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerformanceReq) GetOrderOk() (*MatchingDirection, bool) {
	if o == nil || isNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DnPerformanceReq) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given MatchingDirection and assigns it to the Order field.
func (o *DnPerformanceReq) SetOrder(v MatchingDirection) {
	o.Order = &v
}

// GetReportThresholds returns the ReportThresholds field value if set, zero value otherwise.
func (o *DnPerformanceReq) GetReportThresholds() []ThresholdLevel {
	if o == nil || isNil(o.ReportThresholds) {
		var ret []ThresholdLevel
		return ret
	}
	return o.ReportThresholds
}

// GetReportThresholdsOk returns a tuple with the ReportThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerformanceReq) GetReportThresholdsOk() ([]ThresholdLevel, bool) {
	if o == nil || isNil(o.ReportThresholds) {
		return nil, false
	}
	return o.ReportThresholds, true
}

// HasReportThresholds returns a boolean if a field has been set.
func (o *DnPerformanceReq) HasReportThresholds() bool {
	if o != nil && !isNil(o.ReportThresholds) {
		return true
	}

	return false
}

// SetReportThresholds gets a reference to the given []ThresholdLevel and assigns it to the ReportThresholds field.
func (o *DnPerformanceReq) SetReportThresholds(v []ThresholdLevel) {
	o.ReportThresholds = v
}

func (o DnPerformanceReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnPerformanceReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnPerfOrderCriter) {
		toSerialize["dnPerfOrderCriter"] = o.DnPerfOrderCriter
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.ReportThresholds) {
		toSerialize["reportThresholds"] = o.ReportThresholds
	}
	return toSerialize, nil
}

type NullableDnPerformanceReq struct {
	value *DnPerformanceReq
	isSet bool
}

func (v NullableDnPerformanceReq) Get() *DnPerformanceReq {
	return v.value
}

func (v *NullableDnPerformanceReq) Set(val *DnPerformanceReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDnPerformanceReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDnPerformanceReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnPerformanceReq(val *DnPerformanceReq) *NullableDnPerformanceReq {
	return &NullableDnPerformanceReq{value: val, isSet: true}
}

func (v NullableDnPerformanceReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnPerformanceReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


