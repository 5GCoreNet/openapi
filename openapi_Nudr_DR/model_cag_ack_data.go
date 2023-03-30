/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"time"
)

// checks if the CagAckData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CagAckData{}

// CagAckData Used to store the status of the latest CAG data update.
type CagAckData struct {
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`
}

// NewCagAckData instantiates a new CagAckData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCagAckData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus) *CagAckData {
	this := CagAckData{}
	this.ProvisioningTime = provisioningTime
	this.UeUpdateStatus = ueUpdateStatus
	return &this
}

// NewCagAckDataWithDefaults instantiates a new CagAckData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCagAckDataWithDefaults() *CagAckData {
	this := CagAckData{}
	return &this
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *CagAckData) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *CagAckData) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *CagAckData) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetUeUpdateStatus returns the UeUpdateStatus field value
func (o *CagAckData) GetUeUpdateStatus() UeUpdateStatus {
	if o == nil {
		var ret UeUpdateStatus
		return ret
	}

	return o.UeUpdateStatus
}

// GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field value
// and a boolean to check if the value has been set.
func (o *CagAckData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeUpdateStatus, true
}

// SetUeUpdateStatus sets field value
func (o *CagAckData) SetUeUpdateStatus(v UeUpdateStatus) {
	o.UeUpdateStatus = v
}

func (o CagAckData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CagAckData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provisioningTime"] = o.ProvisioningTime
	toSerialize["ueUpdateStatus"] = o.UeUpdateStatus
	return toSerialize, nil
}

type NullableCagAckData struct {
	value *CagAckData
	isSet bool
}

func (v NullableCagAckData) Get() *CagAckData {
	return v.value
}

func (v *NullableCagAckData) Set(val *CagAckData) {
	v.value = val
	v.isSet = true
}

func (v NullableCagAckData) IsSet() bool {
	return v.isSet
}

func (v *NullableCagAckData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCagAckData(val *CagAckData) *NullableCagAckData {
	return &NullableCagAckData{value: val, isSet: true}
}

func (v NullableCagAckData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCagAckData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


