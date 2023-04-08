/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AsSessionWithQoS

import (
	"encoding/json"
)

// checks if the QosMonitoringInformationRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosMonitoringInformationRm{}

// QosMonitoringInformationRm Represents the same as the QosMonitoringInformation data type but with the nullable:true property.
type QosMonitoringInformationRm struct {
	ReqQosMonParams []RequestedQosMonitoringParameter `json:"reqQosMonParams,omitempty"`
	RepFreqs []ReportingFrequency `json:"repFreqs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	RepThreshDl NullableInt32 `json:"repThreshDl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	RepThreshUl NullableInt32 `json:"repThreshUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	RepThreshRp NullableInt32 `json:"repThreshRp,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	WaitTime NullableInt32 `json:"waitTime,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	RepPeriod NullableInt32 `json:"repPeriod,omitempty"`
}

// NewQosMonitoringInformationRm instantiates a new QosMonitoringInformationRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringInformationRm() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// NewQosMonitoringInformationRmWithDefaults instantiates a new QosMonitoringInformationRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringInformationRmWithDefaults() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// GetReqQosMonParams returns the ReqQosMonParams field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetReqQosMonParams() []RequestedQosMonitoringParameter {
	if o == nil || isNil(o.ReqQosMonParams) {
		var ret []RequestedQosMonitoringParameter
		return ret
	}
	return o.ReqQosMonParams
}

// GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetReqQosMonParamsOk() ([]RequestedQosMonitoringParameter, bool) {
	if o == nil || isNil(o.ReqQosMonParams) {
		return nil, false
	}
	return o.ReqQosMonParams, true
}

// HasReqQosMonParams returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasReqQosMonParams() bool {
	if o != nil && !isNil(o.ReqQosMonParams) {
		return true
	}

	return false
}

// SetReqQosMonParams gets a reference to the given []RequestedQosMonitoringParameter and assigns it to the ReqQosMonParams field.
func (o *QosMonitoringInformationRm) SetReqQosMonParams(v []RequestedQosMonitoringParameter) {
	o.ReqQosMonParams = v
}

// GetRepFreqs returns the RepFreqs field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepFreqs() []ReportingFrequency {
	if o == nil || isNil(o.RepFreqs) {
		var ret []ReportingFrequency
		return ret
	}
	return o.RepFreqs
}

// GetRepFreqsOk returns a tuple with the RepFreqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepFreqsOk() ([]ReportingFrequency, bool) {
	if o == nil || isNil(o.RepFreqs) {
		return nil, false
	}
	return o.RepFreqs, true
}

// HasRepFreqs returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepFreqs() bool {
	if o != nil && !isNil(o.RepFreqs) {
		return true
	}

	return false
}

// SetRepFreqs gets a reference to the given []ReportingFrequency and assigns it to the RepFreqs field.
func (o *QosMonitoringInformationRm) SetRepFreqs(v []ReportingFrequency) {
	o.RepFreqs = v
}

// GetRepThreshDl returns the RepThreshDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepThreshDl() int32 {
	if o == nil || isNil(o.RepThreshDl.Get()) {
		var ret int32
		return ret
	}
	return *o.RepThreshDl.Get()
}

// GetRepThreshDlOk returns a tuple with the RepThreshDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepThreshDlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepThreshDl.Get(), o.RepThreshDl.IsSet()
}

// HasRepThreshDl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshDl() bool {
	if o != nil && o.RepThreshDl.IsSet() {
		return true
	}

	return false
}

// SetRepThreshDl gets a reference to the given NullableInt32 and assigns it to the RepThreshDl field.
func (o *QosMonitoringInformationRm) SetRepThreshDl(v int32) {
	o.RepThreshDl.Set(&v)
}
// SetRepThreshDlNil sets the value for RepThreshDl to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepThreshDlNil() {
	o.RepThreshDl.Set(nil)
}

// UnsetRepThreshDl ensures that no value is present for RepThreshDl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepThreshDl() {
	o.RepThreshDl.Unset()
}

// GetRepThreshUl returns the RepThreshUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepThreshUl() int32 {
	if o == nil || isNil(o.RepThreshUl.Get()) {
		var ret int32
		return ret
	}
	return *o.RepThreshUl.Get()
}

// GetRepThreshUlOk returns a tuple with the RepThreshUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepThreshUlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepThreshUl.Get(), o.RepThreshUl.IsSet()
}

// HasRepThreshUl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshUl() bool {
	if o != nil && o.RepThreshUl.IsSet() {
		return true
	}

	return false
}

// SetRepThreshUl gets a reference to the given NullableInt32 and assigns it to the RepThreshUl field.
func (o *QosMonitoringInformationRm) SetRepThreshUl(v int32) {
	o.RepThreshUl.Set(&v)
}
// SetRepThreshUlNil sets the value for RepThreshUl to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepThreshUlNil() {
	o.RepThreshUl.Set(nil)
}

// UnsetRepThreshUl ensures that no value is present for RepThreshUl, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepThreshUl() {
	o.RepThreshUl.Unset()
}

// GetRepThreshRp returns the RepThreshRp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepThreshRp() int32 {
	if o == nil || isNil(o.RepThreshRp.Get()) {
		var ret int32
		return ret
	}
	return *o.RepThreshRp.Get()
}

// GetRepThreshRpOk returns a tuple with the RepThreshRp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepThreshRpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepThreshRp.Get(), o.RepThreshRp.IsSet()
}

// HasRepThreshRp returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshRp() bool {
	if o != nil && o.RepThreshRp.IsSet() {
		return true
	}

	return false
}

// SetRepThreshRp gets a reference to the given NullableInt32 and assigns it to the RepThreshRp field.
func (o *QosMonitoringInformationRm) SetRepThreshRp(v int32) {
	o.RepThreshRp.Set(&v)
}
// SetRepThreshRpNil sets the value for RepThreshRp to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepThreshRpNil() {
	o.RepThreshRp.Set(nil)
}

// UnsetRepThreshRp ensures that no value is present for RepThreshRp, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepThreshRp() {
	o.RepThreshRp.Unset()
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetWaitTime() int32 {
	if o == nil || isNil(o.WaitTime.Get()) {
		var ret int32
		return ret
	}
	return *o.WaitTime.Get()
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetWaitTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WaitTime.Get(), o.WaitTime.IsSet()
}

// HasWaitTime returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasWaitTime() bool {
	if o != nil && o.WaitTime.IsSet() {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given NullableInt32 and assigns it to the WaitTime field.
func (o *QosMonitoringInformationRm) SetWaitTime(v int32) {
	o.WaitTime.Set(&v)
}
// SetWaitTimeNil sets the value for WaitTime to be an explicit nil
func (o *QosMonitoringInformationRm) SetWaitTimeNil() {
	o.WaitTime.Set(nil)
}

// UnsetWaitTime ensures that no value is present for WaitTime, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetWaitTime() {
	o.WaitTime.Unset()
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosMonitoringInformationRm) GetRepPeriod() int32 {
	if o == nil || isNil(o.RepPeriod.Get()) {
		var ret int32
		return ret
	}
	return *o.RepPeriod.Get()
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosMonitoringInformationRm) GetRepPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepPeriod.Get(), o.RepPeriod.IsSet()
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepPeriod() bool {
	if o != nil && o.RepPeriod.IsSet() {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given NullableInt32 and assigns it to the RepPeriod field.
func (o *QosMonitoringInformationRm) SetRepPeriod(v int32) {
	o.RepPeriod.Set(&v)
}
// SetRepPeriodNil sets the value for RepPeriod to be an explicit nil
func (o *QosMonitoringInformationRm) SetRepPeriodNil() {
	o.RepPeriod.Set(nil)
}

// UnsetRepPeriod ensures that no value is present for RepPeriod, not even an explicit nil
func (o *QosMonitoringInformationRm) UnsetRepPeriod() {
	o.RepPeriod.Unset()
}

func (o QosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosMonitoringInformationRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReqQosMonParams) {
		toSerialize["reqQosMonParams"] = o.ReqQosMonParams
	}
	if !isNil(o.RepFreqs) {
		toSerialize["repFreqs"] = o.RepFreqs
	}
	if o.RepThreshDl.IsSet() {
		toSerialize["repThreshDl"] = o.RepThreshDl.Get()
	}
	if o.RepThreshUl.IsSet() {
		toSerialize["repThreshUl"] = o.RepThreshUl.Get()
	}
	if o.RepThreshRp.IsSet() {
		toSerialize["repThreshRp"] = o.RepThreshRp.Get()
	}
	if o.WaitTime.IsSet() {
		toSerialize["waitTime"] = o.WaitTime.Get()
	}
	if o.RepPeriod.IsSet() {
		toSerialize["repPeriod"] = o.RepPeriod.Get()
	}
	return toSerialize, nil
}

type NullableQosMonitoringInformationRm struct {
	value *QosMonitoringInformationRm
	isSet bool
}

func (v NullableQosMonitoringInformationRm) Get() *QosMonitoringInformationRm {
	return v.value
}

func (v *NullableQosMonitoringInformationRm) Set(val *QosMonitoringInformationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringInformationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringInformationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringInformationRm(val *QosMonitoringInformationRm) *NullableQosMonitoringInformationRm {
	return &NullableQosMonitoringInformationRm{value: val, isSet: true}
}

func (v NullableQosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringInformationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


