/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"time"
)

// checks if the UpuData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpuData{}

// UpuData Used to store the status of the latest UPU data update.
type UpuData struct {
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time      `json:"provisioningTime"`
	UeUpdateStatus   UeUpdateStatus `json:"ueUpdateStatus"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuXmacIue *string `json:"upuXmacIue,omitempty"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue *string `json:"upuMacIue,omitempty"`
}

// NewUpuData instantiates a new UpuData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpuData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus) *UpuData {
	this := UpuData{}
	this.ProvisioningTime = provisioningTime
	this.UeUpdateStatus = ueUpdateStatus
	return &this
}

// NewUpuDataWithDefaults instantiates a new UpuData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpuDataWithDefaults() *UpuData {
	this := UpuData{}
	return &this
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *UpuData) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *UpuData) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *UpuData) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetUeUpdateStatus returns the UeUpdateStatus field value
func (o *UpuData) GetUeUpdateStatus() UeUpdateStatus {
	if o == nil {
		var ret UeUpdateStatus
		return ret
	}

	return o.UeUpdateStatus
}

// GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field value
// and a boolean to check if the value has been set.
func (o *UpuData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeUpdateStatus, true
}

// SetUeUpdateStatus sets field value
func (o *UpuData) SetUeUpdateStatus(v UeUpdateStatus) {
	o.UeUpdateStatus = v
}

// GetUpuXmacIue returns the UpuXmacIue field value if set, zero value otherwise.
func (o *UpuData) GetUpuXmacIue() string {
	if o == nil || IsNil(o.UpuXmacIue) {
		var ret string
		return ret
	}
	return *o.UpuXmacIue
}

// GetUpuXmacIueOk returns a tuple with the UpuXmacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData) GetUpuXmacIueOk() (*string, bool) {
	if o == nil || IsNil(o.UpuXmacIue) {
		return nil, false
	}
	return o.UpuXmacIue, true
}

// HasUpuXmacIue returns a boolean if a field has been set.
func (o *UpuData) HasUpuXmacIue() bool {
	if o != nil && !IsNil(o.UpuXmacIue) {
		return true
	}

	return false
}

// SetUpuXmacIue gets a reference to the given string and assigns it to the UpuXmacIue field.
func (o *UpuData) SetUpuXmacIue(v string) {
	o.UpuXmacIue = &v
}

// GetUpuMacIue returns the UpuMacIue field value if set, zero value otherwise.
func (o *UpuData) GetUpuMacIue() string {
	if o == nil || IsNil(o.UpuMacIue) {
		var ret string
		return ret
	}
	return *o.UpuMacIue
}

// GetUpuMacIueOk returns a tuple with the UpuMacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData) GetUpuMacIueOk() (*string, bool) {
	if o == nil || IsNil(o.UpuMacIue) {
		return nil, false
	}
	return o.UpuMacIue, true
}

// HasUpuMacIue returns a boolean if a field has been set.
func (o *UpuData) HasUpuMacIue() bool {
	if o != nil && !IsNil(o.UpuMacIue) {
		return true
	}

	return false
}

// SetUpuMacIue gets a reference to the given string and assigns it to the UpuMacIue field.
func (o *UpuData) SetUpuMacIue(v string) {
	o.UpuMacIue = &v
}

func (o UpuData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpuData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provisioningTime"] = o.ProvisioningTime
	toSerialize["ueUpdateStatus"] = o.UeUpdateStatus
	if !IsNil(o.UpuXmacIue) {
		toSerialize["upuXmacIue"] = o.UpuXmacIue
	}
	if !IsNil(o.UpuMacIue) {
		toSerialize["upuMacIue"] = o.UpuMacIue
	}
	return toSerialize, nil
}

type NullableUpuData struct {
	value *UpuData
	isSet bool
}

func (v NullableUpuData) Get() *UpuData {
	return v.value
}

func (v *NullableUpuData) Set(val *UpuData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpuData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpuData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpuData(val *UpuData) *NullableUpuData {
	return &NullableUpuData{value: val, isSet: true}
}

func (v NullableUpuData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpuData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
