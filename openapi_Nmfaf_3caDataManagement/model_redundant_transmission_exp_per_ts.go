/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"time"
)

// checks if the RedundantTransmissionExpPerTS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundantTransmissionExpPerTS{}

// RedundantTransmissionExpPerTS The redundant transmission experience per Time Slot.
type RedundantTransmissionExpPerTS struct {
	// string with format 'date-time' as defined in OpenAPI.
	TsStart time.Time `json:"tsStart"`
	// indicating a time in seconds.
	TsDuration      int32                     `json:"tsDuration"`
	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`
	// Redundant Transmission Status. Set to \"true\" if redundant transmission was activated, otherwise set to \"false\". Default value is \"false\" if omitted.
	RedTransStatus *bool `json:"redTransStatus,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	UeRatio *int32 `json:"ueRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewRedundantTransmissionExpPerTS instantiates a new RedundantTransmissionExpPerTS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundantTransmissionExpPerTS(tsStart time.Time, tsDuration int32, obsvRedTransExp ObservedRedundantTransExp) *RedundantTransmissionExpPerTS {
	this := RedundantTransmissionExpPerTS{}
	this.TsStart = tsStart
	this.TsDuration = tsDuration
	this.ObsvRedTransExp = obsvRedTransExp
	return &this
}

// NewRedundantTransmissionExpPerTSWithDefaults instantiates a new RedundantTransmissionExpPerTS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundantTransmissionExpPerTSWithDefaults() *RedundantTransmissionExpPerTS {
	this := RedundantTransmissionExpPerTS{}
	return &this
}

// GetTsStart returns the TsStart field value
func (o *RedundantTransmissionExpPerTS) GetTsStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TsStart
}

// GetTsStartOk returns a tuple with the TsStart field value
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetTsStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsStart, true
}

// SetTsStart sets field value
func (o *RedundantTransmissionExpPerTS) SetTsStart(v time.Time) {
	o.TsStart = v
}

// GetTsDuration returns the TsDuration field value
func (o *RedundantTransmissionExpPerTS) GetTsDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TsDuration
}

// GetTsDurationOk returns a tuple with the TsDuration field value
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetTsDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsDuration, true
}

// SetTsDuration sets field value
func (o *RedundantTransmissionExpPerTS) SetTsDuration(v int32) {
	o.TsDuration = v
}

// GetObsvRedTransExp returns the ObsvRedTransExp field value
func (o *RedundantTransmissionExpPerTS) GetObsvRedTransExp() ObservedRedundantTransExp {
	if o == nil {
		var ret ObservedRedundantTransExp
		return ret
	}

	return o.ObsvRedTransExp
}

// GetObsvRedTransExpOk returns a tuple with the ObsvRedTransExp field value
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetObsvRedTransExpOk() (*ObservedRedundantTransExp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObsvRedTransExp, true
}

// SetObsvRedTransExp sets field value
func (o *RedundantTransmissionExpPerTS) SetObsvRedTransExp(v ObservedRedundantTransExp) {
	o.ObsvRedTransExp = v
}

// GetRedTransStatus returns the RedTransStatus field value if set, zero value otherwise.
func (o *RedundantTransmissionExpPerTS) GetRedTransStatus() bool {
	if o == nil || IsNil(o.RedTransStatus) {
		var ret bool
		return ret
	}
	return *o.RedTransStatus
}

// GetRedTransStatusOk returns a tuple with the RedTransStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetRedTransStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.RedTransStatus) {
		return nil, false
	}
	return o.RedTransStatus, true
}

// HasRedTransStatus returns a boolean if a field has been set.
func (o *RedundantTransmissionExpPerTS) HasRedTransStatus() bool {
	if o != nil && !IsNil(o.RedTransStatus) {
		return true
	}

	return false
}

// SetRedTransStatus gets a reference to the given bool and assigns it to the RedTransStatus field.
func (o *RedundantTransmissionExpPerTS) SetRedTransStatus(v bool) {
	o.RedTransStatus = &v
}

// GetUeRatio returns the UeRatio field value if set, zero value otherwise.
func (o *RedundantTransmissionExpPerTS) GetUeRatio() int32 {
	if o == nil || IsNil(o.UeRatio) {
		var ret int32
		return ret
	}
	return *o.UeRatio
}

// GetUeRatioOk returns a tuple with the UeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetUeRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.UeRatio) {
		return nil, false
	}
	return o.UeRatio, true
}

// HasUeRatio returns a boolean if a field has been set.
func (o *RedundantTransmissionExpPerTS) HasUeRatio() bool {
	if o != nil && !IsNil(o.UeRatio) {
		return true
	}

	return false
}

// SetUeRatio gets a reference to the given int32 and assigns it to the UeRatio field.
func (o *RedundantTransmissionExpPerTS) SetUeRatio(v int32) {
	o.UeRatio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *RedundantTransmissionExpPerTS) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundantTransmissionExpPerTS) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *RedundantTransmissionExpPerTS) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *RedundantTransmissionExpPerTS) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o RedundantTransmissionExpPerTS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundantTransmissionExpPerTS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tsStart"] = o.TsStart
	toSerialize["tsDuration"] = o.TsDuration
	toSerialize["obsvRedTransExp"] = o.ObsvRedTransExp
	if !IsNil(o.RedTransStatus) {
		toSerialize["redTransStatus"] = o.RedTransStatus
	}
	if !IsNil(o.UeRatio) {
		toSerialize["ueRatio"] = o.UeRatio
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableRedundantTransmissionExpPerTS struct {
	value *RedundantTransmissionExpPerTS
	isSet bool
}

func (v NullableRedundantTransmissionExpPerTS) Get() *RedundantTransmissionExpPerTS {
	return v.value
}

func (v *NullableRedundantTransmissionExpPerTS) Set(val *RedundantTransmissionExpPerTS) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantTransmissionExpPerTS) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantTransmissionExpPerTS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantTransmissionExpPerTS(val *RedundantTransmissionExpPerTS) *NullableRedundantTransmissionExpPerTS {
	return &NullableRedundantTransmissionExpPerTS{value: val, isSet: true}
}

func (v NullableRedundantTransmissionExpPerTS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantTransmissionExpPerTS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
