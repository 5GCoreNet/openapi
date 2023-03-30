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

// checks if the CagData1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CagData1{}

// CagData1 struct for CagData1
type CagData1 struct {
	// A map (list of key-value pairs where PlmnId serves as key) of CagInfo
	CagInfos map[string]CagInfo1 `json:"cagInfos"`
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime *time.Time `json:"provisioningTime,omitempty"`
}

// NewCagData1 instantiates a new CagData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCagData1(cagInfos map[string]CagInfo1) *CagData1 {
	this := CagData1{}
	this.CagInfos = cagInfos
	return &this
}

// NewCagData1WithDefaults instantiates a new CagData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCagData1WithDefaults() *CagData1 {
	this := CagData1{}
	return &this
}

// GetCagInfos returns the CagInfos field value
func (o *CagData1) GetCagInfos() map[string]CagInfo1 {
	if o == nil {
		var ret map[string]CagInfo1
		return ret
	}

	return o.CagInfos
}

// GetCagInfosOk returns a tuple with the CagInfos field value
// and a boolean to check if the value has been set.
func (o *CagData1) GetCagInfosOk() (*map[string]CagInfo1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CagInfos, true
}

// SetCagInfos sets field value
func (o *CagData1) SetCagInfos(v map[string]CagInfo1) {
	o.CagInfos = v
}

// GetProvisioningTime returns the ProvisioningTime field value if set, zero value otherwise.
func (o *CagData1) GetProvisioningTime() time.Time {
	if o == nil || IsNil(o.ProvisioningTime) {
		var ret time.Time
		return ret
	}
	return *o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CagData1) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProvisioningTime) {
		return nil, false
	}
	return o.ProvisioningTime, true
}

// HasProvisioningTime returns a boolean if a field has been set.
func (o *CagData1) HasProvisioningTime() bool {
	if o != nil && !IsNil(o.ProvisioningTime) {
		return true
	}

	return false
}

// SetProvisioningTime gets a reference to the given time.Time and assigns it to the ProvisioningTime field.
func (o *CagData1) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = &v
}

func (o CagData1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CagData1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cagInfos"] = o.CagInfos
	if !IsNil(o.ProvisioningTime) {
		toSerialize["provisioningTime"] = o.ProvisioningTime
	}
	return toSerialize, nil
}

type NullableCagData1 struct {
	value *CagData1
	isSet bool
}

func (v NullableCagData1) Get() *CagData1 {
	return v.value
}

func (v *NullableCagData1) Set(val *CagData1) {
	v.value = val
	v.isSet = true
}

func (v NullableCagData1) IsSet() bool {
	return v.isSet
}

func (v *NullableCagData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCagData1(val *CagData1) *NullableCagData1 {
	return &NullableCagData1{value: val, isSet: true}
}

func (v NullableCagData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCagData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


