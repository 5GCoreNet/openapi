/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
)

// checks if the ACRInfoNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACRInfoNotification{}

// ACRInfoNotification Notification of ACR events information.
type ACRInfoNotification struct {
	// String identifying the Individual ACR events subscription for which the ACT Information notification is delivered.
	SubId string `json:"subId"`
	// Application identifier of the EAS.
	EasId string `json:"easId"`
	// Identity of the AC.
	AcId         *string               `json:"acId,omitempty"`
	EventId      ACREventIDs           `json:"eventId"`
	TrgtInfo     *TargetInfo           `json:"trgtInfo,omitempty"`
	AcrStatus    *ACRCompleteEventInfo `json:"acrStatus,omitempty"`
	EecCtxtReloc *EecCtxtRelocStatus   `json:"eecCtxtReloc,omitempty"`
}

// NewACRInfoNotification instantiates a new ACRInfoNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACRInfoNotification(subId string, easId string, eventId ACREventIDs) *ACRInfoNotification {
	this := ACRInfoNotification{}
	this.SubId = subId
	this.EasId = easId
	this.EventId = eventId
	return &this
}

// NewACRInfoNotificationWithDefaults instantiates a new ACRInfoNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACRInfoNotificationWithDefaults() *ACRInfoNotification {
	this := ACRInfoNotification{}
	return &this
}

// GetSubId returns the SubId field value
func (o *ACRInfoNotification) GetSubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetSubIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubId, true
}

// SetSubId sets field value
func (o *ACRInfoNotification) SetSubId(v string) {
	o.SubId = v
}

// GetEasId returns the EasId field value
func (o *ACRInfoNotification) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEasIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *ACRInfoNotification) SetEasId(v string) {
	o.EasId = v
}

// GetAcId returns the AcId field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetAcId() string {
	if o == nil || IsNil(o.AcId) {
		var ret string
		return ret
	}
	return *o.AcId
}

// GetAcIdOk returns a tuple with the AcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetAcIdOk() (*string, bool) {
	if o == nil || IsNil(o.AcId) {
		return nil, false
	}
	return o.AcId, true
}

// HasAcId returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasAcId() bool {
	if o != nil && !IsNil(o.AcId) {
		return true
	}

	return false
}

// SetAcId gets a reference to the given string and assigns it to the AcId field.
func (o *ACRInfoNotification) SetAcId(v string) {
	o.AcId = &v
}

// GetEventId returns the EventId field value
func (o *ACRInfoNotification) GetEventId() ACREventIDs {
	if o == nil {
		var ret ACREventIDs
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEventIdOk() (*ACREventIDs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *ACRInfoNotification) SetEventId(v ACREventIDs) {
	o.EventId = v
}

// GetTrgtInfo returns the TrgtInfo field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetTrgtInfo() TargetInfo {
	if o == nil || IsNil(o.TrgtInfo) {
		var ret TargetInfo
		return ret
	}
	return *o.TrgtInfo
}

// GetTrgtInfoOk returns a tuple with the TrgtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetTrgtInfoOk() (*TargetInfo, bool) {
	if o == nil || IsNil(o.TrgtInfo) {
		return nil, false
	}
	return o.TrgtInfo, true
}

// HasTrgtInfo returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasTrgtInfo() bool {
	if o != nil && !IsNil(o.TrgtInfo) {
		return true
	}

	return false
}

// SetTrgtInfo gets a reference to the given TargetInfo and assigns it to the TrgtInfo field.
func (o *ACRInfoNotification) SetTrgtInfo(v TargetInfo) {
	o.TrgtInfo = &v
}

// GetAcrStatus returns the AcrStatus field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetAcrStatus() ACRCompleteEventInfo {
	if o == nil || IsNil(o.AcrStatus) {
		var ret ACRCompleteEventInfo
		return ret
	}
	return *o.AcrStatus
}

// GetAcrStatusOk returns a tuple with the AcrStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetAcrStatusOk() (*ACRCompleteEventInfo, bool) {
	if o == nil || IsNil(o.AcrStatus) {
		return nil, false
	}
	return o.AcrStatus, true
}

// HasAcrStatus returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasAcrStatus() bool {
	if o != nil && !IsNil(o.AcrStatus) {
		return true
	}

	return false
}

// SetAcrStatus gets a reference to the given ACRCompleteEventInfo and assigns it to the AcrStatus field.
func (o *ACRInfoNotification) SetAcrStatus(v ACRCompleteEventInfo) {
	o.AcrStatus = &v
}

// GetEecCtxtReloc returns the EecCtxtReloc field value if set, zero value otherwise.
func (o *ACRInfoNotification) GetEecCtxtReloc() EecCtxtRelocStatus {
	if o == nil || IsNil(o.EecCtxtReloc) {
		var ret EecCtxtRelocStatus
		return ret
	}
	return *o.EecCtxtReloc
}

// GetEecCtxtRelocOk returns a tuple with the EecCtxtReloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRInfoNotification) GetEecCtxtRelocOk() (*EecCtxtRelocStatus, bool) {
	if o == nil || IsNil(o.EecCtxtReloc) {
		return nil, false
	}
	return o.EecCtxtReloc, true
}

// HasEecCtxtReloc returns a boolean if a field has been set.
func (o *ACRInfoNotification) HasEecCtxtReloc() bool {
	if o != nil && !IsNil(o.EecCtxtReloc) {
		return true
	}

	return false
}

// SetEecCtxtReloc gets a reference to the given EecCtxtRelocStatus and assigns it to the EecCtxtReloc field.
func (o *ACRInfoNotification) SetEecCtxtReloc(v EecCtxtRelocStatus) {
	o.EecCtxtReloc = &v
}

func (o ACRInfoNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACRInfoNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subId"] = o.SubId
	toSerialize["easId"] = o.EasId
	if !IsNil(o.AcId) {
		toSerialize["acId"] = o.AcId
	}
	toSerialize["eventId"] = o.EventId
	if !IsNil(o.TrgtInfo) {
		toSerialize["trgtInfo"] = o.TrgtInfo
	}
	if !IsNil(o.AcrStatus) {
		toSerialize["acrStatus"] = o.AcrStatus
	}
	if !IsNil(o.EecCtxtReloc) {
		toSerialize["eecCtxtReloc"] = o.EecCtxtReloc
	}
	return toSerialize, nil
}

type NullableACRInfoNotification struct {
	value *ACRInfoNotification
	isSet bool
}

func (v NullableACRInfoNotification) Get() *ACRInfoNotification {
	return v.value
}

func (v *NullableACRInfoNotification) Set(val *ACRInfoNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableACRInfoNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableACRInfoNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRInfoNotification(val *ACRInfoNotification) *NullableACRInfoNotification {
	return &NullableACRInfoNotification{value: val, isSet: true}
}

func (v NullableACRInfoNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRInfoNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
