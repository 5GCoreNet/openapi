/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
	"time"
)

// checks if the AcrMgntEventReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcrMgntEventReport{}

// AcrMgntEventReport Represents an ACR management event report.
type AcrMgntEventReport struct {
	Event AcrMgntEvent `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp     *time.Time        `json:"timeStamp,omitempty"`
	UpPathChgInfo *UpPathChangeInfo `json:"upPathChgInfo,omitempty"`
	EasEndPoint   *EndPoint         `json:"easEndPoint,omitempty"`
	ActStatus     *ActStatus        `json:"actStatus,omitempty"`
}

// NewAcrMgntEventReport instantiates a new AcrMgntEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcrMgntEventReport(event AcrMgntEvent) *AcrMgntEventReport {
	this := AcrMgntEventReport{}
	this.Event = event
	return &this
}

// NewAcrMgntEventReportWithDefaults instantiates a new AcrMgntEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcrMgntEventReportWithDefaults() *AcrMgntEventReport {
	this := AcrMgntEventReport{}
	return &this
}

// GetEvent returns the Event field value
func (o *AcrMgntEventReport) GetEvent() AcrMgntEvent {
	if o == nil {
		var ret AcrMgntEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AcrMgntEventReport) GetEventOk() (*AcrMgntEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AcrMgntEventReport) SetEvent(v AcrMgntEvent) {
	o.Event = v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *AcrMgntEventReport) GetTimeStamp() time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventReport) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *AcrMgntEventReport) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *AcrMgntEventReport) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetUpPathChgInfo returns the UpPathChgInfo field value if set, zero value otherwise.
func (o *AcrMgntEventReport) GetUpPathChgInfo() UpPathChangeInfo {
	if o == nil || IsNil(o.UpPathChgInfo) {
		var ret UpPathChangeInfo
		return ret
	}
	return *o.UpPathChgInfo
}

// GetUpPathChgInfoOk returns a tuple with the UpPathChgInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventReport) GetUpPathChgInfoOk() (*UpPathChangeInfo, bool) {
	if o == nil || IsNil(o.UpPathChgInfo) {
		return nil, false
	}
	return o.UpPathChgInfo, true
}

// HasUpPathChgInfo returns a boolean if a field has been set.
func (o *AcrMgntEventReport) HasUpPathChgInfo() bool {
	if o != nil && !IsNil(o.UpPathChgInfo) {
		return true
	}

	return false
}

// SetUpPathChgInfo gets a reference to the given UpPathChangeInfo and assigns it to the UpPathChgInfo field.
func (o *AcrMgntEventReport) SetUpPathChgInfo(v UpPathChangeInfo) {
	o.UpPathChgInfo = &v
}

// GetEasEndPoint returns the EasEndPoint field value if set, zero value otherwise.
func (o *AcrMgntEventReport) GetEasEndPoint() EndPoint {
	if o == nil || IsNil(o.EasEndPoint) {
		var ret EndPoint
		return ret
	}
	return *o.EasEndPoint
}

// GetEasEndPointOk returns a tuple with the EasEndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventReport) GetEasEndPointOk() (*EndPoint, bool) {
	if o == nil || IsNil(o.EasEndPoint) {
		return nil, false
	}
	return o.EasEndPoint, true
}

// HasEasEndPoint returns a boolean if a field has been set.
func (o *AcrMgntEventReport) HasEasEndPoint() bool {
	if o != nil && !IsNil(o.EasEndPoint) {
		return true
	}

	return false
}

// SetEasEndPoint gets a reference to the given EndPoint and assigns it to the EasEndPoint field.
func (o *AcrMgntEventReport) SetEasEndPoint(v EndPoint) {
	o.EasEndPoint = &v
}

// GetActStatus returns the ActStatus field value if set, zero value otherwise.
func (o *AcrMgntEventReport) GetActStatus() ActStatus {
	if o == nil || IsNil(o.ActStatus) {
		var ret ActStatus
		return ret
	}
	return *o.ActStatus
}

// GetActStatusOk returns a tuple with the ActStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcrMgntEventReport) GetActStatusOk() (*ActStatus, bool) {
	if o == nil || IsNil(o.ActStatus) {
		return nil, false
	}
	return o.ActStatus, true
}

// HasActStatus returns a boolean if a field has been set.
func (o *AcrMgntEventReport) HasActStatus() bool {
	if o != nil && !IsNil(o.ActStatus) {
		return true
	}

	return false
}

// SetActStatus gets a reference to the given ActStatus and assigns it to the ActStatus field.
func (o *AcrMgntEventReport) SetActStatus(v ActStatus) {
	o.ActStatus = &v
}

func (o AcrMgntEventReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcrMgntEventReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !IsNil(o.UpPathChgInfo) {
		toSerialize["upPathChgInfo"] = o.UpPathChgInfo
	}
	if !IsNil(o.EasEndPoint) {
		toSerialize["easEndPoint"] = o.EasEndPoint
	}
	if !IsNil(o.ActStatus) {
		toSerialize["actStatus"] = o.ActStatus
	}
	return toSerialize, nil
}

type NullableAcrMgntEventReport struct {
	value *AcrMgntEventReport
	isSet bool
}

func (v NullableAcrMgntEventReport) Get() *AcrMgntEventReport {
	return v.value
}

func (v *NullableAcrMgntEventReport) Set(val *AcrMgntEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrMgntEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrMgntEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrMgntEventReport(val *AcrMgntEventReport) *NullableAcrMgntEventReport {
	return &NullableAcrMgntEventReport{value: val, isSet: true}
}

func (v NullableAcrMgntEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrMgntEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
