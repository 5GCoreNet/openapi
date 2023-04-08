/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the RoamingStatusReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoamingStatusReport{}

// RoamingStatusReport struct for RoamingStatusReport
type RoamingStatusReport struct {
	Roaming bool `json:"roaming"`
	NewServingPlmn PlmnId `json:"newServingPlmn"`
	AccessType *AccessType `json:"accessType,omitempty"`
}

// NewRoamingStatusReport instantiates a new RoamingStatusReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoamingStatusReport(roaming bool, newServingPlmn PlmnId) *RoamingStatusReport {
	this := RoamingStatusReport{}
	this.Roaming = roaming
	this.NewServingPlmn = newServingPlmn
	return &this
}

// NewRoamingStatusReportWithDefaults instantiates a new RoamingStatusReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoamingStatusReportWithDefaults() *RoamingStatusReport {
	this := RoamingStatusReport{}
	return &this
}

// GetRoaming returns the Roaming field value
func (o *RoamingStatusReport) GetRoaming() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Roaming
}

// GetRoamingOk returns a tuple with the Roaming field value
// and a boolean to check if the value has been set.
func (o *RoamingStatusReport) GetRoamingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roaming, true
}

// SetRoaming sets field value
func (o *RoamingStatusReport) SetRoaming(v bool) {
	o.Roaming = v
}

// GetNewServingPlmn returns the NewServingPlmn field value
func (o *RoamingStatusReport) GetNewServingPlmn() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.NewServingPlmn
}

// GetNewServingPlmnOk returns a tuple with the NewServingPlmn field value
// and a boolean to check if the value has been set.
func (o *RoamingStatusReport) GetNewServingPlmnOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewServingPlmn, true
}

// SetNewServingPlmn sets field value
func (o *RoamingStatusReport) SetNewServingPlmn(v PlmnId) {
	o.NewServingPlmn = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *RoamingStatusReport) GetAccessType() AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoamingStatusReport) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *RoamingStatusReport) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *RoamingStatusReport) SetAccessType(v AccessType) {
	o.AccessType = &v
}

func (o RoamingStatusReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoamingStatusReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roaming"] = o.Roaming
	toSerialize["newServingPlmn"] = o.NewServingPlmn
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	return toSerialize, nil
}

type NullableRoamingStatusReport struct {
	value *RoamingStatusReport
	isSet bool
}

func (v NullableRoamingStatusReport) Get() *RoamingStatusReport {
	return v.value
}

func (v *NullableRoamingStatusReport) Set(val *RoamingStatusReport) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingStatusReport) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingStatusReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingStatusReport(val *RoamingStatusReport) *NullableRoamingStatusReport {
	return &NullableRoamingStatusReport{value: val, isSet: true}
}

func (v NullableRoamingStatusReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingStatusReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


