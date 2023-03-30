/*
Nsoraf_SOR

Nsoraf Steering Of Roaming Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsoraf_SOR

import (
	"encoding/json"
	"time"
)

// checks if the SorAckInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SorAckInfo{}

// SorAckInfo Represents an indication to the SOR-AF on the reception status of the acknowledgment of successful reception of SoR Information by a UE. 
type SorAckInfo struct {
	SorAckStatus SorAckStatus `json:"sorAckStatus"`
	// string with format 'date-time' as defined in OpenAPI.
	SorSendingTime time.Time `json:"sorSendingTime"`
	MeSupportOfSorCmci *bool `json:"meSupportOfSorCmci,omitempty"`
}

// NewSorAckInfo instantiates a new SorAckInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorAckInfo(sorAckStatus SorAckStatus, sorSendingTime time.Time) *SorAckInfo {
	this := SorAckInfo{}
	this.SorAckStatus = sorAckStatus
	this.SorSendingTime = sorSendingTime
	return &this
}

// NewSorAckInfoWithDefaults instantiates a new SorAckInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorAckInfoWithDefaults() *SorAckInfo {
	this := SorAckInfo{}
	return &this
}

// GetSorAckStatus returns the SorAckStatus field value
func (o *SorAckInfo) GetSorAckStatus() SorAckStatus {
	if o == nil {
		var ret SorAckStatus
		return ret
	}

	return o.SorAckStatus
}

// GetSorAckStatusOk returns a tuple with the SorAckStatus field value
// and a boolean to check if the value has been set.
func (o *SorAckInfo) GetSorAckStatusOk() (*SorAckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SorAckStatus, true
}

// SetSorAckStatus sets field value
func (o *SorAckInfo) SetSorAckStatus(v SorAckStatus) {
	o.SorAckStatus = v
}

// GetSorSendingTime returns the SorSendingTime field value
func (o *SorAckInfo) GetSorSendingTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SorSendingTime
}

// GetSorSendingTimeOk returns a tuple with the SorSendingTime field value
// and a boolean to check if the value has been set.
func (o *SorAckInfo) GetSorSendingTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SorSendingTime, true
}

// SetSorSendingTime sets field value
func (o *SorAckInfo) SetSorSendingTime(v time.Time) {
	o.SorSendingTime = v
}

// GetMeSupportOfSorCmci returns the MeSupportOfSorCmci field value if set, zero value otherwise.
func (o *SorAckInfo) GetMeSupportOfSorCmci() bool {
	if o == nil || IsNil(o.MeSupportOfSorCmci) {
		var ret bool
		return ret
	}
	return *o.MeSupportOfSorCmci
}

// GetMeSupportOfSorCmciOk returns a tuple with the MeSupportOfSorCmci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorAckInfo) GetMeSupportOfSorCmciOk() (*bool, bool) {
	if o == nil || IsNil(o.MeSupportOfSorCmci) {
		return nil, false
	}
	return o.MeSupportOfSorCmci, true
}

// HasMeSupportOfSorCmci returns a boolean if a field has been set.
func (o *SorAckInfo) HasMeSupportOfSorCmci() bool {
	if o != nil && !IsNil(o.MeSupportOfSorCmci) {
		return true
	}

	return false
}

// SetMeSupportOfSorCmci gets a reference to the given bool and assigns it to the MeSupportOfSorCmci field.
func (o *SorAckInfo) SetMeSupportOfSorCmci(v bool) {
	o.MeSupportOfSorCmci = &v
}

func (o SorAckInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorAckInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sorAckStatus"] = o.SorAckStatus
	toSerialize["sorSendingTime"] = o.SorSendingTime
	if !IsNil(o.MeSupportOfSorCmci) {
		toSerialize["meSupportOfSorCmci"] = o.MeSupportOfSorCmci
	}
	return toSerialize, nil
}

type NullableSorAckInfo struct {
	value *SorAckInfo
	isSet bool
}

func (v NullableSorAckInfo) Get() *SorAckInfo {
	return v.value
}

func (v *NullableSorAckInfo) Set(val *SorAckInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSorAckInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSorAckInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorAckInfo(val *SorAckInfo) *NullableSorAckInfo {
	return &NullableSorAckInfo{value: val, isSet: true}
}

func (v NullableSorAckInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorAckInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


