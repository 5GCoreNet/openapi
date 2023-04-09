/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
)

// checks if the ContextUpdateReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextUpdateReqData{}

// ContextUpdateReqData Data within ContextUpdate Request
type ContextUpdateReqData struct {
	MbsServiceArea         *MbsServiceArea      `json:"mbsServiceArea,omitempty"`
	MbsServiceAreaInfoList []MbsServiceAreaInfo `json:"mbsServiceAreaInfoList,omitempty"`
	N2MbsSmInfo            *N2MbsSmInfo         `json:"n2MbsSmInfo,omitempty"`
	RanIdList              []GlobalRanNodeId    `json:"ranIdList,omitempty"`
	NoNgapSignallingInd    *bool                `json:"noNgapSignallingInd,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifyUri *string `json:"notifyUri,omitempty"`
	// indicating a time in seconds.
	MaxResponseTime    *int32 `json:"maxResponseTime,omitempty"`
	N2MbsInfoChangeInd *bool  `json:"n2MbsInfoChangeInd,omitempty"`
}

// NewContextUpdateReqData instantiates a new ContextUpdateReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdateReqData() *ContextUpdateReqData {
	this := ContextUpdateReqData{}
	return &this
}

// NewContextUpdateReqDataWithDefaults instantiates a new ContextUpdateReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdateReqDataWithDefaults() *ContextUpdateReqData {
	this := ContextUpdateReqData{}
	return &this
}

// GetMbsServiceArea returns the MbsServiceArea field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetMbsServiceArea() MbsServiceArea {
	if o == nil || IsNil(o.MbsServiceArea) {
		var ret MbsServiceArea
		return ret
	}
	return *o.MbsServiceArea
}

// GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetMbsServiceAreaOk() (*MbsServiceArea, bool) {
	if o == nil || IsNil(o.MbsServiceArea) {
		return nil, false
	}
	return o.MbsServiceArea, true
}

// HasMbsServiceArea returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasMbsServiceArea() bool {
	if o != nil && !IsNil(o.MbsServiceArea) {
		return true
	}

	return false
}

// SetMbsServiceArea gets a reference to the given MbsServiceArea and assigns it to the MbsServiceArea field.
func (o *ContextUpdateReqData) SetMbsServiceArea(v MbsServiceArea) {
	o.MbsServiceArea = &v
}

// GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetMbsServiceAreaInfoList() []MbsServiceAreaInfo {
	if o == nil || IsNil(o.MbsServiceAreaInfoList) {
		var ret []MbsServiceAreaInfo
		return ret
	}
	return o.MbsServiceAreaInfoList
}

// GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetMbsServiceAreaInfoListOk() ([]MbsServiceAreaInfo, bool) {
	if o == nil || IsNil(o.MbsServiceAreaInfoList) {
		return nil, false
	}
	return o.MbsServiceAreaInfoList, true
}

// HasMbsServiceAreaInfoList returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasMbsServiceAreaInfoList() bool {
	if o != nil && !IsNil(o.MbsServiceAreaInfoList) {
		return true
	}

	return false
}

// SetMbsServiceAreaInfoList gets a reference to the given []MbsServiceAreaInfo and assigns it to the MbsServiceAreaInfoList field.
func (o *ContextUpdateReqData) SetMbsServiceAreaInfoList(v []MbsServiceAreaInfo) {
	o.MbsServiceAreaInfoList = v
}

// GetN2MbsSmInfo returns the N2MbsSmInfo field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetN2MbsSmInfo() N2MbsSmInfo {
	if o == nil || IsNil(o.N2MbsSmInfo) {
		var ret N2MbsSmInfo
		return ret
	}
	return *o.N2MbsSmInfo
}

// GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool) {
	if o == nil || IsNil(o.N2MbsSmInfo) {
		return nil, false
	}
	return o.N2MbsSmInfo, true
}

// HasN2MbsSmInfo returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasN2MbsSmInfo() bool {
	if o != nil && !IsNil(o.N2MbsSmInfo) {
		return true
	}

	return false
}

// SetN2MbsSmInfo gets a reference to the given N2MbsSmInfo and assigns it to the N2MbsSmInfo field.
func (o *ContextUpdateReqData) SetN2MbsSmInfo(v N2MbsSmInfo) {
	o.N2MbsSmInfo = &v
}

// GetRanIdList returns the RanIdList field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetRanIdList() []GlobalRanNodeId {
	if o == nil || IsNil(o.RanIdList) {
		var ret []GlobalRanNodeId
		return ret
	}
	return o.RanIdList
}

// GetRanIdListOk returns a tuple with the RanIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetRanIdListOk() ([]GlobalRanNodeId, bool) {
	if o == nil || IsNil(o.RanIdList) {
		return nil, false
	}
	return o.RanIdList, true
}

// HasRanIdList returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasRanIdList() bool {
	if o != nil && !IsNil(o.RanIdList) {
		return true
	}

	return false
}

// SetRanIdList gets a reference to the given []GlobalRanNodeId and assigns it to the RanIdList field.
func (o *ContextUpdateReqData) SetRanIdList(v []GlobalRanNodeId) {
	o.RanIdList = v
}

// GetNoNgapSignallingInd returns the NoNgapSignallingInd field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetNoNgapSignallingInd() bool {
	if o == nil || IsNil(o.NoNgapSignallingInd) {
		var ret bool
		return ret
	}
	return *o.NoNgapSignallingInd
}

// GetNoNgapSignallingIndOk returns a tuple with the NoNgapSignallingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetNoNgapSignallingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.NoNgapSignallingInd) {
		return nil, false
	}
	return o.NoNgapSignallingInd, true
}

// HasNoNgapSignallingInd returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasNoNgapSignallingInd() bool {
	if o != nil && !IsNil(o.NoNgapSignallingInd) {
		return true
	}

	return false
}

// SetNoNgapSignallingInd gets a reference to the given bool and assigns it to the NoNgapSignallingInd field.
func (o *ContextUpdateReqData) SetNoNgapSignallingInd(v bool) {
	o.NoNgapSignallingInd = &v
}

// GetNotifyUri returns the NotifyUri field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetNotifyUri() string {
	if o == nil || IsNil(o.NotifyUri) {
		var ret string
		return ret
	}
	return *o.NotifyUri
}

// GetNotifyUriOk returns a tuple with the NotifyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetNotifyUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUri) {
		return nil, false
	}
	return o.NotifyUri, true
}

// HasNotifyUri returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasNotifyUri() bool {
	if o != nil && !IsNil(o.NotifyUri) {
		return true
	}

	return false
}

// SetNotifyUri gets a reference to the given string and assigns it to the NotifyUri field.
func (o *ContextUpdateReqData) SetNotifyUri(v string) {
	o.NotifyUri = &v
}

// GetMaxResponseTime returns the MaxResponseTime field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetMaxResponseTime() int32 {
	if o == nil || IsNil(o.MaxResponseTime) {
		var ret int32
		return ret
	}
	return *o.MaxResponseTime
}

// GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetMaxResponseTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResponseTime) {
		return nil, false
	}
	return o.MaxResponseTime, true
}

// HasMaxResponseTime returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasMaxResponseTime() bool {
	if o != nil && !IsNil(o.MaxResponseTime) {
		return true
	}

	return false
}

// SetMaxResponseTime gets a reference to the given int32 and assigns it to the MaxResponseTime field.
func (o *ContextUpdateReqData) SetMaxResponseTime(v int32) {
	o.MaxResponseTime = &v
}

// GetN2MbsInfoChangeInd returns the N2MbsInfoChangeInd field value if set, zero value otherwise.
func (o *ContextUpdateReqData) GetN2MbsInfoChangeInd() bool {
	if o == nil || IsNil(o.N2MbsInfoChangeInd) {
		var ret bool
		return ret
	}
	return *o.N2MbsInfoChangeInd
}

// GetN2MbsInfoChangeIndOk returns a tuple with the N2MbsInfoChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateReqData) GetN2MbsInfoChangeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.N2MbsInfoChangeInd) {
		return nil, false
	}
	return o.N2MbsInfoChangeInd, true
}

// HasN2MbsInfoChangeInd returns a boolean if a field has been set.
func (o *ContextUpdateReqData) HasN2MbsInfoChangeInd() bool {
	if o != nil && !IsNil(o.N2MbsInfoChangeInd) {
		return true
	}

	return false
}

// SetN2MbsInfoChangeInd gets a reference to the given bool and assigns it to the N2MbsInfoChangeInd field.
func (o *ContextUpdateReqData) SetN2MbsInfoChangeInd(v bool) {
	o.N2MbsInfoChangeInd = &v
}

func (o ContextUpdateReqData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextUpdateReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsServiceArea) {
		toSerialize["mbsServiceArea"] = o.MbsServiceArea
	}
	if !IsNil(o.MbsServiceAreaInfoList) {
		toSerialize["mbsServiceAreaInfoList"] = o.MbsServiceAreaInfoList
	}
	if !IsNil(o.N2MbsSmInfo) {
		toSerialize["n2MbsSmInfo"] = o.N2MbsSmInfo
	}
	if !IsNil(o.RanIdList) {
		toSerialize["ranIdList"] = o.RanIdList
	}
	if !IsNil(o.NoNgapSignallingInd) {
		toSerialize["noNgapSignallingInd"] = o.NoNgapSignallingInd
	}
	if !IsNil(o.NotifyUri) {
		toSerialize["notifyUri"] = o.NotifyUri
	}
	if !IsNil(o.MaxResponseTime) {
		toSerialize["maxResponseTime"] = o.MaxResponseTime
	}
	if !IsNil(o.N2MbsInfoChangeInd) {
		toSerialize["n2MbsInfoChangeInd"] = o.N2MbsInfoChangeInd
	}
	return toSerialize, nil
}

type NullableContextUpdateReqData struct {
	value *ContextUpdateReqData
	isSet bool
}

func (v NullableContextUpdateReqData) Get() *ContextUpdateReqData {
	return v.value
}

func (v *NullableContextUpdateReqData) Set(val *ContextUpdateReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdateReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdateReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdateReqData(val *ContextUpdateReqData) *NullableContextUpdateReqData {
	return &NullableContextUpdateReqData{value: val, isSet: true}
}

func (v NullableContextUpdateReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdateReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
