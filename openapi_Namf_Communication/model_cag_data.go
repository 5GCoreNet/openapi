/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"time"
)

// checks if the CagData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CagData{}

// CagData struct for CagData
type CagData struct {
	// A map (list of key-value pairs where PlmnId serves as key) of CagInfo
	CagInfos map[string]CagInfo `json:"cagInfos"`
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime *time.Time `json:"provisioningTime,omitempty"`
}

// NewCagData instantiates a new CagData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCagData(cagInfos map[string]CagInfo) *CagData {
	this := CagData{}
	this.CagInfos = cagInfos
	return &this
}

// NewCagDataWithDefaults instantiates a new CagData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCagDataWithDefaults() *CagData {
	this := CagData{}
	return &this
}

// GetCagInfos returns the CagInfos field value
func (o *CagData) GetCagInfos() map[string]CagInfo {
	if o == nil {
		var ret map[string]CagInfo
		return ret
	}

	return o.CagInfos
}

// GetCagInfosOk returns a tuple with the CagInfos field value
// and a boolean to check if the value has been set.
func (o *CagData) GetCagInfosOk() (*map[string]CagInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CagInfos, true
}

// SetCagInfos sets field value
func (o *CagData) SetCagInfos(v map[string]CagInfo) {
	o.CagInfos = v
}

// GetProvisioningTime returns the ProvisioningTime field value if set, zero value otherwise.
func (o *CagData) GetProvisioningTime() time.Time {
	if o == nil || IsNil(o.ProvisioningTime) {
		var ret time.Time
		return ret
	}
	return *o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CagData) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProvisioningTime) {
		return nil, false
	}
	return o.ProvisioningTime, true
}

// HasProvisioningTime returns a boolean if a field has been set.
func (o *CagData) HasProvisioningTime() bool {
	if o != nil && !IsNil(o.ProvisioningTime) {
		return true
	}

	return false
}

// SetProvisioningTime gets a reference to the given time.Time and assigns it to the ProvisioningTime field.
func (o *CagData) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = &v
}

func (o CagData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CagData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cagInfos"] = o.CagInfos
	if !IsNil(o.ProvisioningTime) {
		toSerialize["provisioningTime"] = o.ProvisioningTime
	}
	return toSerialize, nil
}

type NullableCagData struct {
	value *CagData
	isSet bool
}

func (v NullableCagData) Get() *CagData {
	return v.value
}

func (v *NullableCagData) Set(val *CagData) {
	v.value = val
	v.isSet = true
}

func (v NullableCagData) IsSet() bool {
	return v.isSet
}

func (v *NullableCagData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCagData(val *CagData) *NullableCagData {
	return &NullableCagData{value: val, isSet: true}
}

func (v NullableCagData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCagData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
