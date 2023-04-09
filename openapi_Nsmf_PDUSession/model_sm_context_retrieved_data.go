/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the SmContextRetrievedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmContextRetrievedData{}

// SmContextRetrievedData Data within Retrieve SM Context Response
type SmContextRetrievedData struct {
	// UE EPS PDN Connection container from SMF to AMF
	UeEpsPdnConnection  string               `json:"ueEpsPdnConnection"`
	SmContext           *SmContext           `json:"smContext,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	DlDataWaitingInd    *bool                `json:"dlDataWaitingInd,omitempty"`
	AfCoordinationInfo  *AfCoordinationInfo  `json:"afCoordinationInfo,omitempty"`
}

// NewSmContextRetrievedData instantiates a new SmContextRetrievedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextRetrievedData(ueEpsPdnConnection string) *SmContextRetrievedData {
	this := SmContextRetrievedData{}
	this.UeEpsPdnConnection = ueEpsPdnConnection
	var dlDataWaitingInd bool = false
	this.DlDataWaitingInd = &dlDataWaitingInd
	return &this
}

// NewSmContextRetrievedDataWithDefaults instantiates a new SmContextRetrievedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextRetrievedDataWithDefaults() *SmContextRetrievedData {
	this := SmContextRetrievedData{}
	var dlDataWaitingInd bool = false
	this.DlDataWaitingInd = &dlDataWaitingInd
	return &this
}

// GetUeEpsPdnConnection returns the UeEpsPdnConnection field value
func (o *SmContextRetrievedData) GetUeEpsPdnConnection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeEpsPdnConnection
}

// GetUeEpsPdnConnectionOk returns a tuple with the UeEpsPdnConnection field value
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetUeEpsPdnConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeEpsPdnConnection, true
}

// SetUeEpsPdnConnection sets field value
func (o *SmContextRetrievedData) SetUeEpsPdnConnection(v string) {
	o.UeEpsPdnConnection = v
}

// GetSmContext returns the SmContext field value if set, zero value otherwise.
func (o *SmContextRetrievedData) GetSmContext() SmContext {
	if o == nil || IsNil(o.SmContext) {
		var ret SmContext
		return ret
	}
	return *o.SmContext
}

// GetSmContextOk returns a tuple with the SmContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetSmContextOk() (*SmContext, bool) {
	if o == nil || IsNil(o.SmContext) {
		return nil, false
	}
	return o.SmContext, true
}

// HasSmContext returns a boolean if a field has been set.
func (o *SmContextRetrievedData) HasSmContext() bool {
	if o != nil && !IsNil(o.SmContext) {
		return true
	}

	return false
}

// SetSmContext gets a reference to the given SmContext and assigns it to the SmContext field.
func (o *SmContextRetrievedData) SetSmContext(v SmContext) {
	o.SmContext = &v
}

// GetSmallDataRateStatus returns the SmallDataRateStatus field value if set, zero value otherwise.
func (o *SmContextRetrievedData) GetSmallDataRateStatus() SmallDataRateStatus {
	if o == nil || IsNil(o.SmallDataRateStatus) {
		var ret SmallDataRateStatus
		return ret
	}
	return *o.SmallDataRateStatus
}

// GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool) {
	if o == nil || IsNil(o.SmallDataRateStatus) {
		return nil, false
	}
	return o.SmallDataRateStatus, true
}

// HasSmallDataRateStatus returns a boolean if a field has been set.
func (o *SmContextRetrievedData) HasSmallDataRateStatus() bool {
	if o != nil && !IsNil(o.SmallDataRateStatus) {
		return true
	}

	return false
}

// SetSmallDataRateStatus gets a reference to the given SmallDataRateStatus and assigns it to the SmallDataRateStatus field.
func (o *SmContextRetrievedData) SetSmallDataRateStatus(v SmallDataRateStatus) {
	o.SmallDataRateStatus = &v
}

// GetApnRateStatus returns the ApnRateStatus field value if set, zero value otherwise.
func (o *SmContextRetrievedData) GetApnRateStatus() ApnRateStatus {
	if o == nil || IsNil(o.ApnRateStatus) {
		var ret ApnRateStatus
		return ret
	}
	return *o.ApnRateStatus
}

// GetApnRateStatusOk returns a tuple with the ApnRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetApnRateStatusOk() (*ApnRateStatus, bool) {
	if o == nil || IsNil(o.ApnRateStatus) {
		return nil, false
	}
	return o.ApnRateStatus, true
}

// HasApnRateStatus returns a boolean if a field has been set.
func (o *SmContextRetrievedData) HasApnRateStatus() bool {
	if o != nil && !IsNil(o.ApnRateStatus) {
		return true
	}

	return false
}

// SetApnRateStatus gets a reference to the given ApnRateStatus and assigns it to the ApnRateStatus field.
func (o *SmContextRetrievedData) SetApnRateStatus(v ApnRateStatus) {
	o.ApnRateStatus = &v
}

// GetDlDataWaitingInd returns the DlDataWaitingInd field value if set, zero value otherwise.
func (o *SmContextRetrievedData) GetDlDataWaitingInd() bool {
	if o == nil || IsNil(o.DlDataWaitingInd) {
		var ret bool
		return ret
	}
	return *o.DlDataWaitingInd
}

// GetDlDataWaitingIndOk returns a tuple with the DlDataWaitingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetDlDataWaitingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DlDataWaitingInd) {
		return nil, false
	}
	return o.DlDataWaitingInd, true
}

// HasDlDataWaitingInd returns a boolean if a field has been set.
func (o *SmContextRetrievedData) HasDlDataWaitingInd() bool {
	if o != nil && !IsNil(o.DlDataWaitingInd) {
		return true
	}

	return false
}

// SetDlDataWaitingInd gets a reference to the given bool and assigns it to the DlDataWaitingInd field.
func (o *SmContextRetrievedData) SetDlDataWaitingInd(v bool) {
	o.DlDataWaitingInd = &v
}

// GetAfCoordinationInfo returns the AfCoordinationInfo field value if set, zero value otherwise.
func (o *SmContextRetrievedData) GetAfCoordinationInfo() AfCoordinationInfo {
	if o == nil || IsNil(o.AfCoordinationInfo) {
		var ret AfCoordinationInfo
		return ret
	}
	return *o.AfCoordinationInfo
}

// GetAfCoordinationInfoOk returns a tuple with the AfCoordinationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrievedData) GetAfCoordinationInfoOk() (*AfCoordinationInfo, bool) {
	if o == nil || IsNil(o.AfCoordinationInfo) {
		return nil, false
	}
	return o.AfCoordinationInfo, true
}

// HasAfCoordinationInfo returns a boolean if a field has been set.
func (o *SmContextRetrievedData) HasAfCoordinationInfo() bool {
	if o != nil && !IsNil(o.AfCoordinationInfo) {
		return true
	}

	return false
}

// SetAfCoordinationInfo gets a reference to the given AfCoordinationInfo and assigns it to the AfCoordinationInfo field.
func (o *SmContextRetrievedData) SetAfCoordinationInfo(v AfCoordinationInfo) {
	o.AfCoordinationInfo = &v
}

func (o SmContextRetrievedData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmContextRetrievedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueEpsPdnConnection"] = o.UeEpsPdnConnection
	if !IsNil(o.SmContext) {
		toSerialize["smContext"] = o.SmContext
	}
	if !IsNil(o.SmallDataRateStatus) {
		toSerialize["smallDataRateStatus"] = o.SmallDataRateStatus
	}
	if !IsNil(o.ApnRateStatus) {
		toSerialize["apnRateStatus"] = o.ApnRateStatus
	}
	if !IsNil(o.DlDataWaitingInd) {
		toSerialize["dlDataWaitingInd"] = o.DlDataWaitingInd
	}
	if !IsNil(o.AfCoordinationInfo) {
		toSerialize["afCoordinationInfo"] = o.AfCoordinationInfo
	}
	return toSerialize, nil
}

type NullableSmContextRetrievedData struct {
	value *SmContextRetrievedData
	isSet bool
}

func (v NullableSmContextRetrievedData) Get() *SmContextRetrievedData {
	return v.value
}

func (v *NullableSmContextRetrievedData) Set(val *SmContextRetrievedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextRetrievedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextRetrievedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextRetrievedData(val *SmContextRetrievedData) *NullableSmContextRetrievedData {
	return &NullableSmContextRetrievedData{value: val, isSet: true}
}

func (v NullableSmContextRetrievedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextRetrievedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
