/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"time"
)

// checks if the MbsContextInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsContextInfo{}

// MbsContextInfo MBS context information
type MbsContextInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	AnyUeInd  *bool      `json:"anyUeInd,omitempty"`
	LlSsm     *Ssm       `json:"llSsm,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	CTeid          *int32          `json:"cTeid,omitempty"`
	MbsServiceArea *MbsServiceArea `json:"mbsServiceArea,omitempty"`
	// A map (list of key-value pairs) where the key identifies an areaSessionId
	MbsServiceAreaInfoList *map[string]MbsServiceAreaInfo `json:"mbsServiceAreaInfoList,omitempty"`
}

// NewMbsContextInfo instantiates a new MbsContextInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsContextInfo() *MbsContextInfo {
	this := MbsContextInfo{}
	var anyUeInd bool = false
	this.AnyUeInd = &anyUeInd
	return &this
}

// NewMbsContextInfoWithDefaults instantiates a new MbsContextInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsContextInfoWithDefaults() *MbsContextInfo {
	this := MbsContextInfo{}
	var anyUeInd bool = false
	this.AnyUeInd = &anyUeInd
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MbsContextInfo) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MbsContextInfo) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *MbsContextInfo) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *MbsContextInfo) GetAnyUeInd() bool {
	if o == nil || IsNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AnyUeInd) {
		return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *MbsContextInfo) HasAnyUeInd() bool {
	if o != nil && !IsNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *MbsContextInfo) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetLlSsm returns the LlSsm field value if set, zero value otherwise.
func (o *MbsContextInfo) GetLlSsm() Ssm {
	if o == nil || IsNil(o.LlSsm) {
		var ret Ssm
		return ret
	}
	return *o.LlSsm
}

// GetLlSsmOk returns a tuple with the LlSsm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetLlSsmOk() (*Ssm, bool) {
	if o == nil || IsNil(o.LlSsm) {
		return nil, false
	}
	return o.LlSsm, true
}

// HasLlSsm returns a boolean if a field has been set.
func (o *MbsContextInfo) HasLlSsm() bool {
	if o != nil && !IsNil(o.LlSsm) {
		return true
	}

	return false
}

// SetLlSsm gets a reference to the given Ssm and assigns it to the LlSsm field.
func (o *MbsContextInfo) SetLlSsm(v Ssm) {
	o.LlSsm = &v
}

// GetCTeid returns the CTeid field value if set, zero value otherwise.
func (o *MbsContextInfo) GetCTeid() int32 {
	if o == nil || IsNil(o.CTeid) {
		var ret int32
		return ret
	}
	return *o.CTeid
}

// GetCTeidOk returns a tuple with the CTeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetCTeidOk() (*int32, bool) {
	if o == nil || IsNil(o.CTeid) {
		return nil, false
	}
	return o.CTeid, true
}

// HasCTeid returns a boolean if a field has been set.
func (o *MbsContextInfo) HasCTeid() bool {
	if o != nil && !IsNil(o.CTeid) {
		return true
	}

	return false
}

// SetCTeid gets a reference to the given int32 and assigns it to the CTeid field.
func (o *MbsContextInfo) SetCTeid(v int32) {
	o.CTeid = &v
}

// GetMbsServiceArea returns the MbsServiceArea field value if set, zero value otherwise.
func (o *MbsContextInfo) GetMbsServiceArea() MbsServiceArea {
	if o == nil || IsNil(o.MbsServiceArea) {
		var ret MbsServiceArea
		return ret
	}
	return *o.MbsServiceArea
}

// GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetMbsServiceAreaOk() (*MbsServiceArea, bool) {
	if o == nil || IsNil(o.MbsServiceArea) {
		return nil, false
	}
	return o.MbsServiceArea, true
}

// HasMbsServiceArea returns a boolean if a field has been set.
func (o *MbsContextInfo) HasMbsServiceArea() bool {
	if o != nil && !IsNil(o.MbsServiceArea) {
		return true
	}

	return false
}

// SetMbsServiceArea gets a reference to the given MbsServiceArea and assigns it to the MbsServiceArea field.
func (o *MbsContextInfo) SetMbsServiceArea(v MbsServiceArea) {
	o.MbsServiceArea = &v
}

// GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field value if set, zero value otherwise.
func (o *MbsContextInfo) GetMbsServiceAreaInfoList() map[string]MbsServiceAreaInfo {
	if o == nil || IsNil(o.MbsServiceAreaInfoList) {
		var ret map[string]MbsServiceAreaInfo
		return ret
	}
	return *o.MbsServiceAreaInfoList
}

// GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsContextInfo) GetMbsServiceAreaInfoListOk() (*map[string]MbsServiceAreaInfo, bool) {
	if o == nil || IsNil(o.MbsServiceAreaInfoList) {
		return nil, false
	}
	return o.MbsServiceAreaInfoList, true
}

// HasMbsServiceAreaInfoList returns a boolean if a field has been set.
func (o *MbsContextInfo) HasMbsServiceAreaInfoList() bool {
	if o != nil && !IsNil(o.MbsServiceAreaInfoList) {
		return true
	}

	return false
}

// SetMbsServiceAreaInfoList gets a reference to the given map[string]MbsServiceAreaInfo and assigns it to the MbsServiceAreaInfoList field.
func (o *MbsContextInfo) SetMbsServiceAreaInfoList(v map[string]MbsServiceAreaInfo) {
	o.MbsServiceAreaInfoList = &v
}

func (o MbsContextInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsContextInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !IsNil(o.LlSsm) {
		toSerialize["llSsm"] = o.LlSsm
	}
	if !IsNil(o.CTeid) {
		toSerialize["cTeid"] = o.CTeid
	}
	if !IsNil(o.MbsServiceArea) {
		toSerialize["mbsServiceArea"] = o.MbsServiceArea
	}
	if !IsNil(o.MbsServiceAreaInfoList) {
		toSerialize["mbsServiceAreaInfoList"] = o.MbsServiceAreaInfoList
	}
	return toSerialize, nil
}

type NullableMbsContextInfo struct {
	value *MbsContextInfo
	isSet bool
}

func (v NullableMbsContextInfo) Get() *MbsContextInfo {
	return v.value
}

func (v *NullableMbsContextInfo) Set(val *MbsContextInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsContextInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsContextInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsContextInfo(val *MbsContextInfo) *NullableMbsContextInfo {
	return &NullableMbsContextInfo{value: val, isSet: true}
}

func (v NullableMbsContextInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsContextInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
