/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the UeRegStatusUpdateReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeRegStatusUpdateReqData{}

// UeRegStatusUpdateReqData Data within a UE registration status update request to indicate a completion of transferring at a target AMF
type UeRegStatusUpdateReqData struct {
	TransferStatus UeContextTransferStatus `json:"transferStatus"`
	ToReleaseSessionList []int32 `json:"toReleaseSessionList,omitempty"`
	PcfReselectedInd *bool `json:"pcfReselectedInd,omitempty"`
	SmfChangeInfoList []SmfChangeInfo `json:"smfChangeInfoList,omitempty"`
	AnalyticsNotUsedList []string `json:"analyticsNotUsedList,omitempty"`
	ToReleaseSessionInfo []ReleaseSessionInfo `json:"toReleaseSessionInfo,omitempty"`
}

// NewUeRegStatusUpdateReqData instantiates a new UeRegStatusUpdateReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeRegStatusUpdateReqData(transferStatus UeContextTransferStatus) *UeRegStatusUpdateReqData {
	this := UeRegStatusUpdateReqData{}
	this.TransferStatus = transferStatus
	return &this
}

// NewUeRegStatusUpdateReqDataWithDefaults instantiates a new UeRegStatusUpdateReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeRegStatusUpdateReqDataWithDefaults() *UeRegStatusUpdateReqData {
	this := UeRegStatusUpdateReqData{}
	return &this
}

// GetTransferStatus returns the TransferStatus field value
func (o *UeRegStatusUpdateReqData) GetTransferStatus() UeContextTransferStatus {
	if o == nil {
		var ret UeContextTransferStatus
		return ret
	}

	return o.TransferStatus
}

// GetTransferStatusOk returns a tuple with the TransferStatus field value
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetTransferStatusOk() (*UeContextTransferStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferStatus, true
}

// SetTransferStatus sets field value
func (o *UeRegStatusUpdateReqData) SetTransferStatus(v UeContextTransferStatus) {
	o.TransferStatus = v
}

// GetToReleaseSessionList returns the ToReleaseSessionList field value if set, zero value otherwise.
func (o *UeRegStatusUpdateReqData) GetToReleaseSessionList() []int32 {
	if o == nil || IsNil(o.ToReleaseSessionList) {
		var ret []int32
		return ret
	}
	return o.ToReleaseSessionList
}

// GetToReleaseSessionListOk returns a tuple with the ToReleaseSessionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetToReleaseSessionListOk() ([]int32, bool) {
	if o == nil || IsNil(o.ToReleaseSessionList) {
		return nil, false
	}
	return o.ToReleaseSessionList, true
}

// HasToReleaseSessionList returns a boolean if a field has been set.
func (o *UeRegStatusUpdateReqData) HasToReleaseSessionList() bool {
	if o != nil && !IsNil(o.ToReleaseSessionList) {
		return true
	}

	return false
}

// SetToReleaseSessionList gets a reference to the given []int32 and assigns it to the ToReleaseSessionList field.
func (o *UeRegStatusUpdateReqData) SetToReleaseSessionList(v []int32) {
	o.ToReleaseSessionList = v
}

// GetPcfReselectedInd returns the PcfReselectedInd field value if set, zero value otherwise.
func (o *UeRegStatusUpdateReqData) GetPcfReselectedInd() bool {
	if o == nil || IsNil(o.PcfReselectedInd) {
		var ret bool
		return ret
	}
	return *o.PcfReselectedInd
}

// GetPcfReselectedIndOk returns a tuple with the PcfReselectedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetPcfReselectedIndOk() (*bool, bool) {
	if o == nil || IsNil(o.PcfReselectedInd) {
		return nil, false
	}
	return o.PcfReselectedInd, true
}

// HasPcfReselectedInd returns a boolean if a field has been set.
func (o *UeRegStatusUpdateReqData) HasPcfReselectedInd() bool {
	if o != nil && !IsNil(o.PcfReselectedInd) {
		return true
	}

	return false
}

// SetPcfReselectedInd gets a reference to the given bool and assigns it to the PcfReselectedInd field.
func (o *UeRegStatusUpdateReqData) SetPcfReselectedInd(v bool) {
	o.PcfReselectedInd = &v
}

// GetSmfChangeInfoList returns the SmfChangeInfoList field value if set, zero value otherwise.
func (o *UeRegStatusUpdateReqData) GetSmfChangeInfoList() []SmfChangeInfo {
	if o == nil || IsNil(o.SmfChangeInfoList) {
		var ret []SmfChangeInfo
		return ret
	}
	return o.SmfChangeInfoList
}

// GetSmfChangeInfoListOk returns a tuple with the SmfChangeInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetSmfChangeInfoListOk() ([]SmfChangeInfo, bool) {
	if o == nil || IsNil(o.SmfChangeInfoList) {
		return nil, false
	}
	return o.SmfChangeInfoList, true
}

// HasSmfChangeInfoList returns a boolean if a field has been set.
func (o *UeRegStatusUpdateReqData) HasSmfChangeInfoList() bool {
	if o != nil && !IsNil(o.SmfChangeInfoList) {
		return true
	}

	return false
}

// SetSmfChangeInfoList gets a reference to the given []SmfChangeInfo and assigns it to the SmfChangeInfoList field.
func (o *UeRegStatusUpdateReqData) SetSmfChangeInfoList(v []SmfChangeInfo) {
	o.SmfChangeInfoList = v
}

// GetAnalyticsNotUsedList returns the AnalyticsNotUsedList field value if set, zero value otherwise.
func (o *UeRegStatusUpdateReqData) GetAnalyticsNotUsedList() []string {
	if o == nil || IsNil(o.AnalyticsNotUsedList) {
		var ret []string
		return ret
	}
	return o.AnalyticsNotUsedList
}

// GetAnalyticsNotUsedListOk returns a tuple with the AnalyticsNotUsedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetAnalyticsNotUsedListOk() ([]string, bool) {
	if o == nil || IsNil(o.AnalyticsNotUsedList) {
		return nil, false
	}
	return o.AnalyticsNotUsedList, true
}

// HasAnalyticsNotUsedList returns a boolean if a field has been set.
func (o *UeRegStatusUpdateReqData) HasAnalyticsNotUsedList() bool {
	if o != nil && !IsNil(o.AnalyticsNotUsedList) {
		return true
	}

	return false
}

// SetAnalyticsNotUsedList gets a reference to the given []string and assigns it to the AnalyticsNotUsedList field.
func (o *UeRegStatusUpdateReqData) SetAnalyticsNotUsedList(v []string) {
	o.AnalyticsNotUsedList = v
}

// GetToReleaseSessionInfo returns the ToReleaseSessionInfo field value if set, zero value otherwise.
func (o *UeRegStatusUpdateReqData) GetToReleaseSessionInfo() []ReleaseSessionInfo {
	if o == nil || IsNil(o.ToReleaseSessionInfo) {
		var ret []ReleaseSessionInfo
		return ret
	}
	return o.ToReleaseSessionInfo
}

// GetToReleaseSessionInfoOk returns a tuple with the ToReleaseSessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRegStatusUpdateReqData) GetToReleaseSessionInfoOk() ([]ReleaseSessionInfo, bool) {
	if o == nil || IsNil(o.ToReleaseSessionInfo) {
		return nil, false
	}
	return o.ToReleaseSessionInfo, true
}

// HasToReleaseSessionInfo returns a boolean if a field has been set.
func (o *UeRegStatusUpdateReqData) HasToReleaseSessionInfo() bool {
	if o != nil && !IsNil(o.ToReleaseSessionInfo) {
		return true
	}

	return false
}

// SetToReleaseSessionInfo gets a reference to the given []ReleaseSessionInfo and assigns it to the ToReleaseSessionInfo field.
func (o *UeRegStatusUpdateReqData) SetToReleaseSessionInfo(v []ReleaseSessionInfo) {
	o.ToReleaseSessionInfo = v
}

func (o UeRegStatusUpdateReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeRegStatusUpdateReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transferStatus"] = o.TransferStatus
	if !IsNil(o.ToReleaseSessionList) {
		toSerialize["toReleaseSessionList"] = o.ToReleaseSessionList
	}
	if !IsNil(o.PcfReselectedInd) {
		toSerialize["pcfReselectedInd"] = o.PcfReselectedInd
	}
	if !IsNil(o.SmfChangeInfoList) {
		toSerialize["smfChangeInfoList"] = o.SmfChangeInfoList
	}
	if !IsNil(o.AnalyticsNotUsedList) {
		toSerialize["analyticsNotUsedList"] = o.AnalyticsNotUsedList
	}
	if !IsNil(o.ToReleaseSessionInfo) {
		toSerialize["toReleaseSessionInfo"] = o.ToReleaseSessionInfo
	}
	return toSerialize, nil
}

type NullableUeRegStatusUpdateReqData struct {
	value *UeRegStatusUpdateReqData
	isSet bool
}

func (v NullableUeRegStatusUpdateReqData) Get() *UeRegStatusUpdateReqData {
	return v.value
}

func (v *NullableUeRegStatusUpdateReqData) Set(val *UeRegStatusUpdateReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeRegStatusUpdateReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeRegStatusUpdateReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeRegStatusUpdateReqData(val *UeRegStatusUpdateReqData) *NullableUeRegStatusUpdateReqData {
	return &NullableUeRegStatusUpdateReqData{value: val, isSet: true}
}

func (v NullableUeRegStatusUpdateReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeRegStatusUpdateReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


