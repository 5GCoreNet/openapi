/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the QosMonitoringInformationRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosMonitoringInformationRm{}

// QosMonitoringInformationRm This data type is defined in the same way as the QosMonitoringInformation data type, but with the OpenAPI nullable property set to true.
type QosMonitoringInformationRm struct {
	RepThreshDl *int32 `json:"repThreshDl,omitempty"`
	RepThreshUl *int32 `json:"repThreshUl,omitempty"`
	RepThreshRp *int32 `json:"repThreshRp,omitempty"`
}

// NewQosMonitoringInformationRm instantiates a new QosMonitoringInformationRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringInformationRm() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// NewQosMonitoringInformationRmWithDefaults instantiates a new QosMonitoringInformationRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringInformationRmWithDefaults() *QosMonitoringInformationRm {
	this := QosMonitoringInformationRm{}
	return &this
}

// GetRepThreshDl returns the RepThreshDl field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshDl() int32 {
	if o == nil || IsNil(o.RepThreshDl) {
		var ret int32
		return ret
	}
	return *o.RepThreshDl
}

// GetRepThreshDlOk returns a tuple with the RepThreshDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshDlOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshDl) {
		return nil, false
	}
	return o.RepThreshDl, true
}

// HasRepThreshDl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshDl() bool {
	if o != nil && !IsNil(o.RepThreshDl) {
		return true
	}

	return false
}

// SetRepThreshDl gets a reference to the given int32 and assigns it to the RepThreshDl field.
func (o *QosMonitoringInformationRm) SetRepThreshDl(v int32) {
	o.RepThreshDl = &v
}

// GetRepThreshUl returns the RepThreshUl field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshUl() int32 {
	if o == nil || IsNil(o.RepThreshUl) {
		var ret int32
		return ret
	}
	return *o.RepThreshUl
}

// GetRepThreshUlOk returns a tuple with the RepThreshUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshUlOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshUl) {
		return nil, false
	}
	return o.RepThreshUl, true
}

// HasRepThreshUl returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshUl() bool {
	if o != nil && !IsNil(o.RepThreshUl) {
		return true
	}

	return false
}

// SetRepThreshUl gets a reference to the given int32 and assigns it to the RepThreshUl field.
func (o *QosMonitoringInformationRm) SetRepThreshUl(v int32) {
	o.RepThreshUl = &v
}

// GetRepThreshRp returns the RepThreshRp field value if set, zero value otherwise.
func (o *QosMonitoringInformationRm) GetRepThreshRp() int32 {
	if o == nil || IsNil(o.RepThreshRp) {
		var ret int32
		return ret
	}
	return *o.RepThreshRp
}

// GetRepThreshRpOk returns a tuple with the RepThreshRp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInformationRm) GetRepThreshRpOk() (*int32, bool) {
	if o == nil || IsNil(o.RepThreshRp) {
		return nil, false
	}
	return o.RepThreshRp, true
}

// HasRepThreshRp returns a boolean if a field has been set.
func (o *QosMonitoringInformationRm) HasRepThreshRp() bool {
	if o != nil && !IsNil(o.RepThreshRp) {
		return true
	}

	return false
}

// SetRepThreshRp gets a reference to the given int32 and assigns it to the RepThreshRp field.
func (o *QosMonitoringInformationRm) SetRepThreshRp(v int32) {
	o.RepThreshRp = &v
}

func (o QosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosMonitoringInformationRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepThreshDl) {
		toSerialize["repThreshDl"] = o.RepThreshDl
	}
	if !IsNil(o.RepThreshUl) {
		toSerialize["repThreshUl"] = o.RepThreshUl
	}
	if !IsNil(o.RepThreshRp) {
		toSerialize["repThreshRp"] = o.RepThreshRp
	}
	return toSerialize, nil
}

type NullableQosMonitoringInformationRm struct {
	value *QosMonitoringInformationRm
	isSet bool
}

func (v NullableQosMonitoringInformationRm) Get() *QosMonitoringInformationRm {
	return v.value
}

func (v *NullableQosMonitoringInformationRm) Set(val *QosMonitoringInformationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringInformationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringInformationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringInformationRm(val *QosMonitoringInformationRm) *NullableQosMonitoringInformationRm {
	return &NullableQosMonitoringInformationRm{value: val, isSet: true}
}

func (v NullableQosMonitoringInformationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringInformationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
