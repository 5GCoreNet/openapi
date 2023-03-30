/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"time"
)

// checks if the AmfEventMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventMode{}

// AmfEventMode Describes how the reports shall be generated by a subscribed event
type AmfEventMode struct {
	Trigger AmfEventTrigger `json:"trigger"`
	MaxReports *int32 `json:"maxReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio *int32 `json:"sampRatio,omitempty"`
	PartitioningCriteria []PartitioningCriteria `json:"partitioningCriteria,omitempty"`
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`
}

// NewAmfEventMode instantiates a new AmfEventMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventMode(trigger AmfEventTrigger) *AmfEventMode {
	this := AmfEventMode{}
	this.Trigger = trigger
	return &this
}

// NewAmfEventModeWithDefaults instantiates a new AmfEventMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventModeWithDefaults() *AmfEventMode {
	this := AmfEventMode{}
	return &this
}

// GetTrigger returns the Trigger field value
func (o *AmfEventMode) GetTrigger() AmfEventTrigger {
	if o == nil {
		var ret AmfEventTrigger
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetTriggerOk() (*AmfEventTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *AmfEventMode) SetTrigger(v AmfEventTrigger) {
	o.Trigger = v
}

// GetMaxReports returns the MaxReports field value if set, zero value otherwise.
func (o *AmfEventMode) GetMaxReports() int32 {
	if o == nil || IsNil(o.MaxReports) {
		var ret int32
		return ret
	}
	return *o.MaxReports
}

// GetMaxReportsOk returns a tuple with the MaxReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetMaxReportsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReports) {
		return nil, false
	}
	return o.MaxReports, true
}

// HasMaxReports returns a boolean if a field has been set.
func (o *AmfEventMode) HasMaxReports() bool {
	if o != nil && !IsNil(o.MaxReports) {
		return true
	}

	return false
}

// SetMaxReports gets a reference to the given int32 and assigns it to the MaxReports field.
func (o *AmfEventMode) SetMaxReports(v int32) {
	o.MaxReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *AmfEventMode) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *AmfEventMode) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *AmfEventMode) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *AmfEventMode) GetRepPeriod() int32 {
	if o == nil || IsNil(o.RepPeriod) {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RepPeriod) {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *AmfEventMode) HasRepPeriod() bool {
	if o != nil && !IsNil(o.RepPeriod) {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *AmfEventMode) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *AmfEventMode) GetSampRatio() int32 {
	if o == nil || IsNil(o.SampRatio) {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetSampRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.SampRatio) {
		return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *AmfEventMode) HasSampRatio() bool {
	if o != nil && !IsNil(o.SampRatio) {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *AmfEventMode) SetSampRatio(v int32) {
	o.SampRatio = &v
}

// GetPartitioningCriteria returns the PartitioningCriteria field value if set, zero value otherwise.
func (o *AmfEventMode) GetPartitioningCriteria() []PartitioningCriteria {
	if o == nil || IsNil(o.PartitioningCriteria) {
		var ret []PartitioningCriteria
		return ret
	}
	return o.PartitioningCriteria
}

// GetPartitioningCriteriaOk returns a tuple with the PartitioningCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetPartitioningCriteriaOk() ([]PartitioningCriteria, bool) {
	if o == nil || IsNil(o.PartitioningCriteria) {
		return nil, false
	}
	return o.PartitioningCriteria, true
}

// HasPartitioningCriteria returns a boolean if a field has been set.
func (o *AmfEventMode) HasPartitioningCriteria() bool {
	if o != nil && !IsNil(o.PartitioningCriteria) {
		return true
	}

	return false
}

// SetPartitioningCriteria gets a reference to the given []PartitioningCriteria and assigns it to the PartitioningCriteria field.
func (o *AmfEventMode) SetPartitioningCriteria(v []PartitioningCriteria) {
	o.PartitioningCriteria = v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *AmfEventMode) GetNotifFlag() NotificationFlag {
	if o == nil || IsNil(o.NotifFlag) {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventMode) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || IsNil(o.NotifFlag) {
		return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *AmfEventMode) HasNotifFlag() bool {
	if o != nil && !IsNil(o.NotifFlag) {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *AmfEventMode) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

func (o AmfEventMode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trigger"] = o.Trigger
	if !IsNil(o.MaxReports) {
		toSerialize["maxReports"] = o.MaxReports
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.RepPeriod) {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if !IsNil(o.SampRatio) {
		toSerialize["sampRatio"] = o.SampRatio
	}
	if !IsNil(o.PartitioningCriteria) {
		toSerialize["partitioningCriteria"] = o.PartitioningCriteria
	}
	if !IsNil(o.NotifFlag) {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	return toSerialize, nil
}

type NullableAmfEventMode struct {
	value *AmfEventMode
	isSet bool
}

func (v NullableAmfEventMode) Get() *AmfEventMode {
	return v.value
}

func (v *NullableAmfEventMode) Set(val *AmfEventMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventMode(val *AmfEventMode) *NullableAmfEventMode {
	return &NullableAmfEventMode{value: val, isSet: true}
}

func (v NullableAmfEventMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


