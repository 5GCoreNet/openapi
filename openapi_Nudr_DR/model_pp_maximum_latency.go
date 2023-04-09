/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"time"
)

// checks if the PpMaximumLatency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PpMaximumLatency{}

// PpMaximumLatency struct for PpMaximumLatency
type PpMaximumLatency struct {
	// indicating a time in seconds.
	MaximumLatency int32  `json:"maximumLatency"`
	AfInstanceId   string `json:"afInstanceId"`
	ReferenceId    int32  `json:"referenceId"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

// NewPpMaximumLatency instantiates a new PpMaximumLatency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpMaximumLatency(maximumLatency int32, afInstanceId string, referenceId int32) *PpMaximumLatency {
	this := PpMaximumLatency{}
	this.MaximumLatency = maximumLatency
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewPpMaximumLatencyWithDefaults instantiates a new PpMaximumLatency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpMaximumLatencyWithDefaults() *PpMaximumLatency {
	this := PpMaximumLatency{}
	return &this
}

// GetMaximumLatency returns the MaximumLatency field value
func (o *PpMaximumLatency) GetMaximumLatency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumLatency
}

// GetMaximumLatencyOk returns a tuple with the MaximumLatency field value
// and a boolean to check if the value has been set.
func (o *PpMaximumLatency) GetMaximumLatencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumLatency, true
}

// SetMaximumLatency sets field value
func (o *PpMaximumLatency) SetMaximumLatency(v int32) {
	o.MaximumLatency = v
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *PpMaximumLatency) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *PpMaximumLatency) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *PpMaximumLatency) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *PpMaximumLatency) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *PpMaximumLatency) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *PpMaximumLatency) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *PpMaximumLatency) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpMaximumLatency) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *PpMaximumLatency) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *PpMaximumLatency) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *PpMaximumLatency) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpMaximumLatency) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *PpMaximumLatency) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *PpMaximumLatency) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o PpMaximumLatency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PpMaximumLatency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maximumLatency"] = o.MaximumLatency
	toSerialize["afInstanceId"] = o.AfInstanceId
	toSerialize["referenceId"] = o.ReferenceId
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return toSerialize, nil
}

type NullablePpMaximumLatency struct {
	value *PpMaximumLatency
	isSet bool
}

func (v NullablePpMaximumLatency) Get() *PpMaximumLatency {
	return v.value
}

func (v *NullablePpMaximumLatency) Set(val *PpMaximumLatency) {
	v.value = val
	v.isSet = true
}

func (v NullablePpMaximumLatency) IsSet() bool {
	return v.isSet
}

func (v *NullablePpMaximumLatency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpMaximumLatency(val *PpMaximumLatency) *NullablePpMaximumLatency {
	return &NullablePpMaximumLatency{value: val, isSet: true}
}

func (v NullablePpMaximumLatency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpMaximumLatency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
