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

// checks if the SorData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SorData{}

// SorData Used to store the status of the latest SOR data update.
type SorData struct {
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`
	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorXmacIue *string `json:"sorXmacIue,omitempty"`
	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIue *string `json:"sorMacIue,omitempty"`
	MeSupportOfSorCmci *bool `json:"meSupportOfSorCmci,omitempty"`
}

// NewSorData instantiates a new SorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus) *SorData {
	this := SorData{}
	this.ProvisioningTime = provisioningTime
	this.UeUpdateStatus = ueUpdateStatus
	return &this
}

// NewSorDataWithDefaults instantiates a new SorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorDataWithDefaults() *SorData {
	this := SorData{}
	return &this
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *SorData) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *SorData) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *SorData) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetUeUpdateStatus returns the UeUpdateStatus field value
func (o *SorData) GetUeUpdateStatus() UeUpdateStatus {
	if o == nil {
		var ret UeUpdateStatus
		return ret
	}

	return o.UeUpdateStatus
}

// GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field value
// and a boolean to check if the value has been set.
func (o *SorData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeUpdateStatus, true
}

// SetUeUpdateStatus sets field value
func (o *SorData) SetUeUpdateStatus(v UeUpdateStatus) {
	o.UeUpdateStatus = v
}

// GetSorXmacIue returns the SorXmacIue field value if set, zero value otherwise.
func (o *SorData) GetSorXmacIue() string {
	if o == nil || isNil(o.SorXmacIue) {
		var ret string
		return ret
	}
	return *o.SorXmacIue
}

// GetSorXmacIueOk returns a tuple with the SorXmacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorData) GetSorXmacIueOk() (*string, bool) {
	if o == nil || isNil(o.SorXmacIue) {
		return nil, false
	}
	return o.SorXmacIue, true
}

// HasSorXmacIue returns a boolean if a field has been set.
func (o *SorData) HasSorXmacIue() bool {
	if o != nil && !isNil(o.SorXmacIue) {
		return true
	}

	return false
}

// SetSorXmacIue gets a reference to the given string and assigns it to the SorXmacIue field.
func (o *SorData) SetSorXmacIue(v string) {
	o.SorXmacIue = &v
}

// GetSorMacIue returns the SorMacIue field value if set, zero value otherwise.
func (o *SorData) GetSorMacIue() string {
	if o == nil || isNil(o.SorMacIue) {
		var ret string
		return ret
	}
	return *o.SorMacIue
}

// GetSorMacIueOk returns a tuple with the SorMacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorData) GetSorMacIueOk() (*string, bool) {
	if o == nil || isNil(o.SorMacIue) {
		return nil, false
	}
	return o.SorMacIue, true
}

// HasSorMacIue returns a boolean if a field has been set.
func (o *SorData) HasSorMacIue() bool {
	if o != nil && !isNil(o.SorMacIue) {
		return true
	}

	return false
}

// SetSorMacIue gets a reference to the given string and assigns it to the SorMacIue field.
func (o *SorData) SetSorMacIue(v string) {
	o.SorMacIue = &v
}

// GetMeSupportOfSorCmci returns the MeSupportOfSorCmci field value if set, zero value otherwise.
func (o *SorData) GetMeSupportOfSorCmci() bool {
	if o == nil || isNil(o.MeSupportOfSorCmci) {
		var ret bool
		return ret
	}
	return *o.MeSupportOfSorCmci
}

// GetMeSupportOfSorCmciOk returns a tuple with the MeSupportOfSorCmci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorData) GetMeSupportOfSorCmciOk() (*bool, bool) {
	if o == nil || isNil(o.MeSupportOfSorCmci) {
		return nil, false
	}
	return o.MeSupportOfSorCmci, true
}

// HasMeSupportOfSorCmci returns a boolean if a field has been set.
func (o *SorData) HasMeSupportOfSorCmci() bool {
	if o != nil && !isNil(o.MeSupportOfSorCmci) {
		return true
	}

	return false
}

// SetMeSupportOfSorCmci gets a reference to the given bool and assigns it to the MeSupportOfSorCmci field.
func (o *SorData) SetMeSupportOfSorCmci(v bool) {
	o.MeSupportOfSorCmci = &v
}

func (o SorData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provisioningTime"] = o.ProvisioningTime
	toSerialize["ueUpdateStatus"] = o.UeUpdateStatus
	if !isNil(o.SorXmacIue) {
		toSerialize["sorXmacIue"] = o.SorXmacIue
	}
	if !isNil(o.SorMacIue) {
		toSerialize["sorMacIue"] = o.SorMacIue
	}
	if !isNil(o.MeSupportOfSorCmci) {
		toSerialize["meSupportOfSorCmci"] = o.MeSupportOfSorCmci
	}
	return toSerialize, nil
}

type NullableSorData struct {
	value *SorData
	isSet bool
}

func (v NullableSorData) Get() *SorData {
	return v.value
}

func (v *NullableSorData) Set(val *SorData) {
	v.value = val
	v.isSet = true
}

func (v NullableSorData) IsSet() bool {
	return v.isSet
}

func (v *NullableSorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorData(val *SorData) *NullableSorData {
	return &NullableSorData{value: val, isSet: true}
}

func (v NullableSorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


