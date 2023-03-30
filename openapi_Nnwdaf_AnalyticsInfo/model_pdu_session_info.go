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

// checks if the PduSessionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSessionInfo{}

// PduSessionInfo Represents session information.
type PduSessionInfo struct {
	// The identifier of the N4 session for the reported PDU Session.
	N4SessId *string `json:"n4SessId,omitempty"`
	// indicating a time in seconds.
	SessInactiveTimer *int32 `json:"sessInactiveTimer,omitempty"`
	PduSessStatus *PduSessionStatus `json:"pduSessStatus,omitempty"`
}

// NewPduSessionInfo instantiates a new PduSessionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionInfo() *PduSessionInfo {
	this := PduSessionInfo{}
	return &this
}

// NewPduSessionInfoWithDefaults instantiates a new PduSessionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionInfoWithDefaults() *PduSessionInfo {
	this := PduSessionInfo{}
	return &this
}

// GetN4SessId returns the N4SessId field value if set, zero value otherwise.
func (o *PduSessionInfo) GetN4SessId() string {
	if o == nil || IsNil(o.N4SessId) {
		var ret string
		return ret
	}
	return *o.N4SessId
}

// GetN4SessIdOk returns a tuple with the N4SessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionInfo) GetN4SessIdOk() (*string, bool) {
	if o == nil || IsNil(o.N4SessId) {
		return nil, false
	}
	return o.N4SessId, true
}

// HasN4SessId returns a boolean if a field has been set.
func (o *PduSessionInfo) HasN4SessId() bool {
	if o != nil && !IsNil(o.N4SessId) {
		return true
	}

	return false
}

// SetN4SessId gets a reference to the given string and assigns it to the N4SessId field.
func (o *PduSessionInfo) SetN4SessId(v string) {
	o.N4SessId = &v
}

// GetSessInactiveTimer returns the SessInactiveTimer field value if set, zero value otherwise.
func (o *PduSessionInfo) GetSessInactiveTimer() int32 {
	if o == nil || IsNil(o.SessInactiveTimer) {
		var ret int32
		return ret
	}
	return *o.SessInactiveTimer
}

// GetSessInactiveTimerOk returns a tuple with the SessInactiveTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionInfo) GetSessInactiveTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.SessInactiveTimer) {
		return nil, false
	}
	return o.SessInactiveTimer, true
}

// HasSessInactiveTimer returns a boolean if a field has been set.
func (o *PduSessionInfo) HasSessInactiveTimer() bool {
	if o != nil && !IsNil(o.SessInactiveTimer) {
		return true
	}

	return false
}

// SetSessInactiveTimer gets a reference to the given int32 and assigns it to the SessInactiveTimer field.
func (o *PduSessionInfo) SetSessInactiveTimer(v int32) {
	o.SessInactiveTimer = &v
}

// GetPduSessStatus returns the PduSessStatus field value if set, zero value otherwise.
func (o *PduSessionInfo) GetPduSessStatus() PduSessionStatus {
	if o == nil || IsNil(o.PduSessStatus) {
		var ret PduSessionStatus
		return ret
	}
	return *o.PduSessStatus
}

// GetPduSessStatusOk returns a tuple with the PduSessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionInfo) GetPduSessStatusOk() (*PduSessionStatus, bool) {
	if o == nil || IsNil(o.PduSessStatus) {
		return nil, false
	}
	return o.PduSessStatus, true
}

// HasPduSessStatus returns a boolean if a field has been set.
func (o *PduSessionInfo) HasPduSessStatus() bool {
	if o != nil && !IsNil(o.PduSessStatus) {
		return true
	}

	return false
}

// SetPduSessStatus gets a reference to the given PduSessionStatus and assigns it to the PduSessStatus field.
func (o *PduSessionInfo) SetPduSessStatus(v PduSessionStatus) {
	o.PduSessStatus = &v
}

func (o PduSessionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSessionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.N4SessId) {
		toSerialize["n4SessId"] = o.N4SessId
	}
	if !IsNil(o.SessInactiveTimer) {
		toSerialize["sessInactiveTimer"] = o.SessInactiveTimer
	}
	if !IsNil(o.PduSessStatus) {
		toSerialize["pduSessStatus"] = o.PduSessStatus
	}
	return toSerialize, nil
}

type NullablePduSessionInfo struct {
	value *PduSessionInfo
	isSet bool
}

func (v NullablePduSessionInfo) Get() *PduSessionInfo {
	return v.value
}

func (v *NullablePduSessionInfo) Set(val *PduSessionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionInfo(val *PduSessionInfo) *NullablePduSessionInfo {
	return &NullablePduSessionInfo{value: val, isSet: true}
}

func (v NullablePduSessionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


