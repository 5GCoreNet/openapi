/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the SessInactTimerForUeComm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessInactTimerForUeComm{}

// SessInactTimerForUeComm Represents the N4 Session inactivity timer.
type SessInactTimerForUeComm struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	N4SessId int32 `json:"n4SessId"`
	// indicating a time in seconds.
	SessInactiveTimer int32 `json:"sessInactiveTimer"`
}

// NewSessInactTimerForUeComm instantiates a new SessInactTimerForUeComm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessInactTimerForUeComm(n4SessId int32, sessInactiveTimer int32) *SessInactTimerForUeComm {
	this := SessInactTimerForUeComm{}
	this.N4SessId = n4SessId
	this.SessInactiveTimer = sessInactiveTimer
	return &this
}

// NewSessInactTimerForUeCommWithDefaults instantiates a new SessInactTimerForUeComm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessInactTimerForUeCommWithDefaults() *SessInactTimerForUeComm {
	this := SessInactTimerForUeComm{}
	return &this
}

// GetN4SessId returns the N4SessId field value
func (o *SessInactTimerForUeComm) GetN4SessId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.N4SessId
}

// GetN4SessIdOk returns a tuple with the N4SessId field value
// and a boolean to check if the value has been set.
func (o *SessInactTimerForUeComm) GetN4SessIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N4SessId, true
}

// SetN4SessId sets field value
func (o *SessInactTimerForUeComm) SetN4SessId(v int32) {
	o.N4SessId = v
}

// GetSessInactiveTimer returns the SessInactiveTimer field value
func (o *SessInactTimerForUeComm) GetSessInactiveTimer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessInactiveTimer
}

// GetSessInactiveTimerOk returns a tuple with the SessInactiveTimer field value
// and a boolean to check if the value has been set.
func (o *SessInactTimerForUeComm) GetSessInactiveTimerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessInactiveTimer, true
}

// SetSessInactiveTimer sets field value
func (o *SessInactTimerForUeComm) SetSessInactiveTimer(v int32) {
	o.SessInactiveTimer = v
}

func (o SessInactTimerForUeComm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessInactTimerForUeComm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n4SessId"] = o.N4SessId
	toSerialize["sessInactiveTimer"] = o.SessInactiveTimer
	return toSerialize, nil
}

type NullableSessInactTimerForUeComm struct {
	value *SessInactTimerForUeComm
	isSet bool
}

func (v NullableSessInactTimerForUeComm) Get() *SessInactTimerForUeComm {
	return v.value
}

func (v *NullableSessInactTimerForUeComm) Set(val *SessInactTimerForUeComm) {
	v.value = val
	v.isSet = true
}

func (v NullableSessInactTimerForUeComm) IsSet() bool {
	return v.isSet
}

func (v *NullableSessInactTimerForUeComm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessInactTimerForUeComm(val *SessInactTimerForUeComm) *NullableSessInactTimerForUeComm {
	return &NullableSessInactTimerForUeComm{value: val, isSet: true}
}

func (v NullableSessInactTimerForUeComm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessInactTimerForUeComm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
