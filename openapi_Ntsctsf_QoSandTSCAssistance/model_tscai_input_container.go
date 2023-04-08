/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
	"time"
)

// checks if the TscaiInputContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TscaiInputContainer{}

// TscaiInputContainer Indicates TSC Traffic pattern.
type TscaiInputContainer struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Periodicity *int32 `json:"periodicity,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	BurstArrivalTime *time.Time `json:"burstArrivalTime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInNumMsg *int32 `json:"surTimeInNumMsg,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInTime *int32 `json:"surTimeInTime,omitempty"`
}

// NewTscaiInputContainer instantiates a new TscaiInputContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTscaiInputContainer() *TscaiInputContainer {
	this := TscaiInputContainer{}
	return &this
}

// NewTscaiInputContainerWithDefaults instantiates a new TscaiInputContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTscaiInputContainerWithDefaults() *TscaiInputContainer {
	this := TscaiInputContainer{}
	return &this
}

// GetPeriodicity returns the Periodicity field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetPeriodicity() int32 {
	if o == nil || isNil(o.Periodicity) {
		var ret int32
		return ret
	}
	return *o.Periodicity
}

// GetPeriodicityOk returns a tuple with the Periodicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetPeriodicityOk() (*int32, bool) {
	if o == nil || isNil(o.Periodicity) {
		return nil, false
	}
	return o.Periodicity, true
}

// HasPeriodicity returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasPeriodicity() bool {
	if o != nil && !isNil(o.Periodicity) {
		return true
	}

	return false
}

// SetPeriodicity gets a reference to the given int32 and assigns it to the Periodicity field.
func (o *TscaiInputContainer) SetPeriodicity(v int32) {
	o.Periodicity = &v
}

// GetBurstArrivalTime returns the BurstArrivalTime field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetBurstArrivalTime() time.Time {
	if o == nil || isNil(o.BurstArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.BurstArrivalTime
}

// GetBurstArrivalTimeOk returns a tuple with the BurstArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetBurstArrivalTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.BurstArrivalTime) {
		return nil, false
	}
	return o.BurstArrivalTime, true
}

// HasBurstArrivalTime returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasBurstArrivalTime() bool {
	if o != nil && !isNil(o.BurstArrivalTime) {
		return true
	}

	return false
}

// SetBurstArrivalTime gets a reference to the given time.Time and assigns it to the BurstArrivalTime field.
func (o *TscaiInputContainer) SetBurstArrivalTime(v time.Time) {
	o.BurstArrivalTime = &v
}

// GetSurTimeInNumMsg returns the SurTimeInNumMsg field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetSurTimeInNumMsg() int32 {
	if o == nil || isNil(o.SurTimeInNumMsg) {
		var ret int32
		return ret
	}
	return *o.SurTimeInNumMsg
}

// GetSurTimeInNumMsgOk returns a tuple with the SurTimeInNumMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetSurTimeInNumMsgOk() (*int32, bool) {
	if o == nil || isNil(o.SurTimeInNumMsg) {
		return nil, false
	}
	return o.SurTimeInNumMsg, true
}

// HasSurTimeInNumMsg returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasSurTimeInNumMsg() bool {
	if o != nil && !isNil(o.SurTimeInNumMsg) {
		return true
	}

	return false
}

// SetSurTimeInNumMsg gets a reference to the given int32 and assigns it to the SurTimeInNumMsg field.
func (o *TscaiInputContainer) SetSurTimeInNumMsg(v int32) {
	o.SurTimeInNumMsg = &v
}

// GetSurTimeInTime returns the SurTimeInTime field value if set, zero value otherwise.
func (o *TscaiInputContainer) GetSurTimeInTime() int32 {
	if o == nil || isNil(o.SurTimeInTime) {
		var ret int32
		return ret
	}
	return *o.SurTimeInTime
}

// GetSurTimeInTimeOk returns a tuple with the SurTimeInTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscaiInputContainer) GetSurTimeInTimeOk() (*int32, bool) {
	if o == nil || isNil(o.SurTimeInTime) {
		return nil, false
	}
	return o.SurTimeInTime, true
}

// HasSurTimeInTime returns a boolean if a field has been set.
func (o *TscaiInputContainer) HasSurTimeInTime() bool {
	if o != nil && !isNil(o.SurTimeInTime) {
		return true
	}

	return false
}

// SetSurTimeInTime gets a reference to the given int32 and assigns it to the SurTimeInTime field.
func (o *TscaiInputContainer) SetSurTimeInTime(v int32) {
	o.SurTimeInTime = &v
}

func (o TscaiInputContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TscaiInputContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Periodicity) {
		toSerialize["periodicity"] = o.Periodicity
	}
	if !isNil(o.BurstArrivalTime) {
		toSerialize["burstArrivalTime"] = o.BurstArrivalTime
	}
	if !isNil(o.SurTimeInNumMsg) {
		toSerialize["surTimeInNumMsg"] = o.SurTimeInNumMsg
	}
	if !isNil(o.SurTimeInTime) {
		toSerialize["surTimeInTime"] = o.SurTimeInTime
	}
	return toSerialize, nil
}

type NullableTscaiInputContainer struct {
	value *TscaiInputContainer
	isSet bool
}

func (v NullableTscaiInputContainer) Get() *TscaiInputContainer {
	return v.value
}

func (v *NullableTscaiInputContainer) Set(val *TscaiInputContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableTscaiInputContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableTscaiInputContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscaiInputContainer(val *TscaiInputContainer) *NullableTscaiInputContainer {
	return &NullableTscaiInputContainer{value: val, isSet: true}
}

func (v NullableTscaiInputContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscaiInputContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


