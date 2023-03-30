/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"time"
)

// checks if the ApnRateStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApnRateStatus{}

// ApnRateStatus Contains the APN rate control status e.g. of the AMF.
type ApnRateStatus struct {
	// When present, it shall contain the number of packets the UE is allowed to send uplink in the given time unit for the given APN (all PDN connections of the UE to this APN see clause 4.7.7.3 in 3GPP TS 23.401. 
	RemainPacketsUl *int32 `json:"remainPacketsUl,omitempty"`
	// When present, it shall contain the number of packets the UE is allowed to send uplink in the given time unit for the given APN (all PDN connections of the UE to this APN see clause 4.7.7.3 in 3GPP TS 23.401. 
	RemainPacketsDl *int32 `json:"remainPacketsDl,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// When present, it shall indicate the number of additional exception reports the UE is allowed to send uplink in the given time unit for the given APN (all PDN connections of the UE to this APN, see clause 4.7.7.3 in 3GPP TS 23.401. 
	RemainExReportsUl *int32 `json:"remainExReportsUl,omitempty"`
	// When present, it shall indicate the number of additional exception reports the AF is allowed to send downlink in the  given time unit for the given APN (all PDN connections of the UE to this APN, see clause 4.7.7.3 in 3GPP TS 23.401. 
	RemainExReportsDl *int32 `json:"remainExReportsDl,omitempty"`
}

// NewApnRateStatus instantiates a new ApnRateStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApnRateStatus() *ApnRateStatus {
	this := ApnRateStatus{}
	return &this
}

// NewApnRateStatusWithDefaults instantiates a new ApnRateStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApnRateStatusWithDefaults() *ApnRateStatus {
	this := ApnRateStatus{}
	return &this
}

// GetRemainPacketsUl returns the RemainPacketsUl field value if set, zero value otherwise.
func (o *ApnRateStatus) GetRemainPacketsUl() int32 {
	if o == nil || IsNil(o.RemainPacketsUl) {
		var ret int32
		return ret
	}
	return *o.RemainPacketsUl
}

// GetRemainPacketsUlOk returns a tuple with the RemainPacketsUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApnRateStatus) GetRemainPacketsUlOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainPacketsUl) {
		return nil, false
	}
	return o.RemainPacketsUl, true
}

// HasRemainPacketsUl returns a boolean if a field has been set.
func (o *ApnRateStatus) HasRemainPacketsUl() bool {
	if o != nil && !IsNil(o.RemainPacketsUl) {
		return true
	}

	return false
}

// SetRemainPacketsUl gets a reference to the given int32 and assigns it to the RemainPacketsUl field.
func (o *ApnRateStatus) SetRemainPacketsUl(v int32) {
	o.RemainPacketsUl = &v
}

// GetRemainPacketsDl returns the RemainPacketsDl field value if set, zero value otherwise.
func (o *ApnRateStatus) GetRemainPacketsDl() int32 {
	if o == nil || IsNil(o.RemainPacketsDl) {
		var ret int32
		return ret
	}
	return *o.RemainPacketsDl
}

// GetRemainPacketsDlOk returns a tuple with the RemainPacketsDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApnRateStatus) GetRemainPacketsDlOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainPacketsDl) {
		return nil, false
	}
	return o.RemainPacketsDl, true
}

// HasRemainPacketsDl returns a boolean if a field has been set.
func (o *ApnRateStatus) HasRemainPacketsDl() bool {
	if o != nil && !IsNil(o.RemainPacketsDl) {
		return true
	}

	return false
}

// SetRemainPacketsDl gets a reference to the given int32 and assigns it to the RemainPacketsDl field.
func (o *ApnRateStatus) SetRemainPacketsDl(v int32) {
	o.RemainPacketsDl = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *ApnRateStatus) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApnRateStatus) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *ApnRateStatus) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *ApnRateStatus) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetRemainExReportsUl returns the RemainExReportsUl field value if set, zero value otherwise.
func (o *ApnRateStatus) GetRemainExReportsUl() int32 {
	if o == nil || IsNil(o.RemainExReportsUl) {
		var ret int32
		return ret
	}
	return *o.RemainExReportsUl
}

// GetRemainExReportsUlOk returns a tuple with the RemainExReportsUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApnRateStatus) GetRemainExReportsUlOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainExReportsUl) {
		return nil, false
	}
	return o.RemainExReportsUl, true
}

// HasRemainExReportsUl returns a boolean if a field has been set.
func (o *ApnRateStatus) HasRemainExReportsUl() bool {
	if o != nil && !IsNil(o.RemainExReportsUl) {
		return true
	}

	return false
}

// SetRemainExReportsUl gets a reference to the given int32 and assigns it to the RemainExReportsUl field.
func (o *ApnRateStatus) SetRemainExReportsUl(v int32) {
	o.RemainExReportsUl = &v
}

// GetRemainExReportsDl returns the RemainExReportsDl field value if set, zero value otherwise.
func (o *ApnRateStatus) GetRemainExReportsDl() int32 {
	if o == nil || IsNil(o.RemainExReportsDl) {
		var ret int32
		return ret
	}
	return *o.RemainExReportsDl
}

// GetRemainExReportsDlOk returns a tuple with the RemainExReportsDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApnRateStatus) GetRemainExReportsDlOk() (*int32, bool) {
	if o == nil || IsNil(o.RemainExReportsDl) {
		return nil, false
	}
	return o.RemainExReportsDl, true
}

// HasRemainExReportsDl returns a boolean if a field has been set.
func (o *ApnRateStatus) HasRemainExReportsDl() bool {
	if o != nil && !IsNil(o.RemainExReportsDl) {
		return true
	}

	return false
}

// SetRemainExReportsDl gets a reference to the given int32 and assigns it to the RemainExReportsDl field.
func (o *ApnRateStatus) SetRemainExReportsDl(v int32) {
	o.RemainExReportsDl = &v
}

func (o ApnRateStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApnRateStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemainPacketsUl) {
		toSerialize["remainPacketsUl"] = o.RemainPacketsUl
	}
	if !IsNil(o.RemainPacketsDl) {
		toSerialize["remainPacketsDl"] = o.RemainPacketsDl
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.RemainExReportsUl) {
		toSerialize["remainExReportsUl"] = o.RemainExReportsUl
	}
	if !IsNil(o.RemainExReportsDl) {
		toSerialize["remainExReportsDl"] = o.RemainExReportsDl
	}
	return toSerialize, nil
}

type NullableApnRateStatus struct {
	value *ApnRateStatus
	isSet bool
}

func (v NullableApnRateStatus) Get() *ApnRateStatus {
	return v.value
}

func (v *NullableApnRateStatus) Set(val *ApnRateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApnRateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApnRateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApnRateStatus(val *ApnRateStatus) *NullableApnRateStatus {
	return &NullableApnRateStatus{value: val, isSet: true}
}

func (v NullableApnRateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApnRateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


