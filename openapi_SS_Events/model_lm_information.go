/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
	"time"
)

// checks if the LMInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LMInformation{}

// LMInformation Represents the location information for a VAL User ID or a VAL UE ID.
type LMInformation struct {
	ValTgtUe ValTargetUe `json:"valTgtUe"`
	LocInfo LocationInfo `json:"locInfo"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
	// Identity of the VAL service
	ValSvcId *string `json:"valSvcId,omitempty"`
}

// NewLMInformation instantiates a new LMInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLMInformation(valTgtUe ValTargetUe, locInfo LocationInfo) *LMInformation {
	this := LMInformation{}
	this.ValTgtUe = valTgtUe
	this.LocInfo = locInfo
	return &this
}

// NewLMInformationWithDefaults instantiates a new LMInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLMInformationWithDefaults() *LMInformation {
	this := LMInformation{}
	return &this
}

// GetValTgtUe returns the ValTgtUe field value
func (o *LMInformation) GetValTgtUe() ValTargetUe {
	if o == nil {
		var ret ValTargetUe
		return ret
	}

	return o.ValTgtUe
}

// GetValTgtUeOk returns a tuple with the ValTgtUe field value
// and a boolean to check if the value has been set.
func (o *LMInformation) GetValTgtUeOk() (*ValTargetUe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValTgtUe, true
}

// SetValTgtUe sets field value
func (o *LMInformation) SetValTgtUe(v ValTargetUe) {
	o.ValTgtUe = v
}

// GetLocInfo returns the LocInfo field value
func (o *LMInformation) GetLocInfo() LocationInfo {
	if o == nil {
		var ret LocationInfo
		return ret
	}

	return o.LocInfo
}

// GetLocInfoOk returns a tuple with the LocInfo field value
// and a boolean to check if the value has been set.
func (o *LMInformation) GetLocInfoOk() (*LocationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocInfo, true
}

// SetLocInfo sets field value
func (o *LMInformation) SetLocInfo(v LocationInfo) {
	o.LocInfo = v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *LMInformation) GetTimeStamp() time.Time {
	if o == nil || isNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LMInformation) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *LMInformation) HasTimeStamp() bool {
	if o != nil && !isNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *LMInformation) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetValSvcId returns the ValSvcId field value if set, zero value otherwise.
func (o *LMInformation) GetValSvcId() string {
	if o == nil || isNil(o.ValSvcId) {
		var ret string
		return ret
	}
	return *o.ValSvcId
}

// GetValSvcIdOk returns a tuple with the ValSvcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LMInformation) GetValSvcIdOk() (*string, bool) {
	if o == nil || isNil(o.ValSvcId) {
		return nil, false
	}
	return o.ValSvcId, true
}

// HasValSvcId returns a boolean if a field has been set.
func (o *LMInformation) HasValSvcId() bool {
	if o != nil && !isNil(o.ValSvcId) {
		return true
	}

	return false
}

// SetValSvcId gets a reference to the given string and assigns it to the ValSvcId field.
func (o *LMInformation) SetValSvcId(v string) {
	o.ValSvcId = &v
}

func (o LMInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LMInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valTgtUe"] = o.ValTgtUe
	toSerialize["locInfo"] = o.LocInfo
	if !isNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !isNil(o.ValSvcId) {
		toSerialize["valSvcId"] = o.ValSvcId
	}
	return toSerialize, nil
}

type NullableLMInformation struct {
	value *LMInformation
	isSet bool
}

func (v NullableLMInformation) Get() *LMInformation {
	return v.value
}

func (v *NullableLMInformation) Set(val *LMInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableLMInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableLMInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLMInformation(val *LMInformation) *NullableLMInformation {
	return &NullableLMInformation{value: val, isSet: true}
}

func (v NullableLMInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLMInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


