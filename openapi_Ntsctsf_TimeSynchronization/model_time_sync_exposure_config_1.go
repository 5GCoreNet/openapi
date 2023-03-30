/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
)

// checks if the TimeSyncExposureConfig1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSyncExposureConfig1{}

// TimeSyncExposureConfig1 Contains the Time Synchronization Configuration parameters.
type TimeSyncExposureConfig1 struct {
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UpNodeId int32 `json:"upNodeId"`
	ReqPtpIns PtpInstance `json:"reqPtpIns"`
	// Indicates that the AF requests 5GS to act as a grandmaster for PTP or gPTP if it is included and set to true. 
	GmEnable *bool `json:"gmEnable,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	GmPrio *int32 `json:"gmPrio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeDom int32 `json:"timeDom"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeSyncErrBdgt *int32 `json:"timeSyncErrBdgt,omitempty"`
	// Notification Correlation ID assigned by the NF service consumer.
	ConfigNotifId string `json:"configNotifId"`
	// String providing an URI formatted according to RFC 3986.
	ConfigNotifUri string `json:"configNotifUri"`
	TempValidity *TemporalValidity `json:"tempValidity,omitempty"`
}

// NewTimeSyncExposureConfig1 instantiates a new TimeSyncExposureConfig1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSyncExposureConfig1(upNodeId int32, reqPtpIns PtpInstance, timeDom int32, configNotifId string, configNotifUri string) *TimeSyncExposureConfig1 {
	this := TimeSyncExposureConfig1{}
	this.UpNodeId = upNodeId
	this.ReqPtpIns = reqPtpIns
	this.TimeDom = timeDom
	this.ConfigNotifId = configNotifId
	this.ConfigNotifUri = configNotifUri
	return &this
}

// NewTimeSyncExposureConfig1WithDefaults instantiates a new TimeSyncExposureConfig1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSyncExposureConfig1WithDefaults() *TimeSyncExposureConfig1 {
	this := TimeSyncExposureConfig1{}
	return &this
}

// GetUpNodeId returns the UpNodeId field value
func (o *TimeSyncExposureConfig1) GetUpNodeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpNodeId
}

// GetUpNodeIdOk returns a tuple with the UpNodeId field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetUpNodeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpNodeId, true
}

// SetUpNodeId sets field value
func (o *TimeSyncExposureConfig1) SetUpNodeId(v int32) {
	o.UpNodeId = v
}

// GetReqPtpIns returns the ReqPtpIns field value
func (o *TimeSyncExposureConfig1) GetReqPtpIns() PtpInstance {
	if o == nil {
		var ret PtpInstance
		return ret
	}

	return o.ReqPtpIns
}

// GetReqPtpInsOk returns a tuple with the ReqPtpIns field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetReqPtpInsOk() (*PtpInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqPtpIns, true
}

// SetReqPtpIns sets field value
func (o *TimeSyncExposureConfig1) SetReqPtpIns(v PtpInstance) {
	o.ReqPtpIns = v
}

// GetGmEnable returns the GmEnable field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig1) GetGmEnable() bool {
	if o == nil || IsNil(o.GmEnable) {
		var ret bool
		return ret
	}
	return *o.GmEnable
}

// GetGmEnableOk returns a tuple with the GmEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetGmEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.GmEnable) {
		return nil, false
	}
	return o.GmEnable, true
}

// HasGmEnable returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig1) HasGmEnable() bool {
	if o != nil && !IsNil(o.GmEnable) {
		return true
	}

	return false
}

// SetGmEnable gets a reference to the given bool and assigns it to the GmEnable field.
func (o *TimeSyncExposureConfig1) SetGmEnable(v bool) {
	o.GmEnable = &v
}

// GetGmPrio returns the GmPrio field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig1) GetGmPrio() int32 {
	if o == nil || IsNil(o.GmPrio) {
		var ret int32
		return ret
	}
	return *o.GmPrio
}

// GetGmPrioOk returns a tuple with the GmPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetGmPrioOk() (*int32, bool) {
	if o == nil || IsNil(o.GmPrio) {
		return nil, false
	}
	return o.GmPrio, true
}

// HasGmPrio returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig1) HasGmPrio() bool {
	if o != nil && !IsNil(o.GmPrio) {
		return true
	}

	return false
}

// SetGmPrio gets a reference to the given int32 and assigns it to the GmPrio field.
func (o *TimeSyncExposureConfig1) SetGmPrio(v int32) {
	o.GmPrio = &v
}

// GetTimeDom returns the TimeDom field value
func (o *TimeSyncExposureConfig1) GetTimeDom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeDom
}

// GetTimeDomOk returns a tuple with the TimeDom field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetTimeDomOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeDom, true
}

// SetTimeDom sets field value
func (o *TimeSyncExposureConfig1) SetTimeDom(v int32) {
	o.TimeDom = v
}

// GetTimeSyncErrBdgt returns the TimeSyncErrBdgt field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig1) GetTimeSyncErrBdgt() int32 {
	if o == nil || IsNil(o.TimeSyncErrBdgt) {
		var ret int32
		return ret
	}
	return *o.TimeSyncErrBdgt
}

// GetTimeSyncErrBdgtOk returns a tuple with the TimeSyncErrBdgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetTimeSyncErrBdgtOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeSyncErrBdgt) {
		return nil, false
	}
	return o.TimeSyncErrBdgt, true
}

// HasTimeSyncErrBdgt returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig1) HasTimeSyncErrBdgt() bool {
	if o != nil && !IsNil(o.TimeSyncErrBdgt) {
		return true
	}

	return false
}

// SetTimeSyncErrBdgt gets a reference to the given int32 and assigns it to the TimeSyncErrBdgt field.
func (o *TimeSyncExposureConfig1) SetTimeSyncErrBdgt(v int32) {
	o.TimeSyncErrBdgt = &v
}

// GetConfigNotifId returns the ConfigNotifId field value
func (o *TimeSyncExposureConfig1) GetConfigNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigNotifId
}

// GetConfigNotifIdOk returns a tuple with the ConfigNotifId field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetConfigNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigNotifId, true
}

// SetConfigNotifId sets field value
func (o *TimeSyncExposureConfig1) SetConfigNotifId(v string) {
	o.ConfigNotifId = v
}

// GetConfigNotifUri returns the ConfigNotifUri field value
func (o *TimeSyncExposureConfig1) GetConfigNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigNotifUri
}

// GetConfigNotifUriOk returns a tuple with the ConfigNotifUri field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetConfigNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigNotifUri, true
}

// SetConfigNotifUri sets field value
func (o *TimeSyncExposureConfig1) SetConfigNotifUri(v string) {
	o.ConfigNotifUri = v
}

// GetTempValidity returns the TempValidity field value if set, zero value otherwise.
func (o *TimeSyncExposureConfig1) GetTempValidity() TemporalValidity {
	if o == nil || IsNil(o.TempValidity) {
		var ret TemporalValidity
		return ret
	}
	return *o.TempValidity
}

// GetTempValidityOk returns a tuple with the TempValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureConfig1) GetTempValidityOk() (*TemporalValidity, bool) {
	if o == nil || IsNil(o.TempValidity) {
		return nil, false
	}
	return o.TempValidity, true
}

// HasTempValidity returns a boolean if a field has been set.
func (o *TimeSyncExposureConfig1) HasTempValidity() bool {
	if o != nil && !IsNil(o.TempValidity) {
		return true
	}

	return false
}

// SetTempValidity gets a reference to the given TemporalValidity and assigns it to the TempValidity field.
func (o *TimeSyncExposureConfig1) SetTempValidity(v TemporalValidity) {
	o.TempValidity = &v
}

func (o TimeSyncExposureConfig1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSyncExposureConfig1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upNodeId"] = o.UpNodeId
	toSerialize["reqPtpIns"] = o.ReqPtpIns
	if !IsNil(o.GmEnable) {
		toSerialize["gmEnable"] = o.GmEnable
	}
	if !IsNil(o.GmPrio) {
		toSerialize["gmPrio"] = o.GmPrio
	}
	toSerialize["timeDom"] = o.TimeDom
	if !IsNil(o.TimeSyncErrBdgt) {
		toSerialize["timeSyncErrBdgt"] = o.TimeSyncErrBdgt
	}
	toSerialize["configNotifId"] = o.ConfigNotifId
	toSerialize["configNotifUri"] = o.ConfigNotifUri
	if !IsNil(o.TempValidity) {
		toSerialize["tempValidity"] = o.TempValidity
	}
	return toSerialize, nil
}

type NullableTimeSyncExposureConfig1 struct {
	value *TimeSyncExposureConfig1
	isSet bool
}

func (v NullableTimeSyncExposureConfig1) Get() *TimeSyncExposureConfig1 {
	return v.value
}

func (v *NullableTimeSyncExposureConfig1) Set(val *TimeSyncExposureConfig1) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSyncExposureConfig1) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSyncExposureConfig1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSyncExposureConfig1(val *TimeSyncExposureConfig1) *NullableTimeSyncExposureConfig1 {
	return &NullableTimeSyncExposureConfig1{value: val, isSet: true}
}

func (v NullableTimeSyncExposureConfig1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSyncExposureConfig1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


