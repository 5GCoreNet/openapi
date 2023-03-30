/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CpProvisioning

import (
	"encoding/json"
)

// checks if the CpReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpReport{}

// CpReport Represents a CP report indicating the CP set identifier(s) which CP parameter(s) are not added or modified successfully and the corresponding failure cause(s).
type CpReport struct {
	// Identifies the CP set identifier(s) which CP parameter(s) are not added or modified successfully
	SetIds []string `json:"setIds,omitempty"`
	FailureCode CpFailureCode `json:"failureCode"`
}

// NewCpReport instantiates a new CpReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpReport(failureCode CpFailureCode) *CpReport {
	this := CpReport{}
	this.FailureCode = failureCode
	return &this
}

// NewCpReportWithDefaults instantiates a new CpReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpReportWithDefaults() *CpReport {
	this := CpReport{}
	return &this
}

// GetSetIds returns the SetIds field value if set, zero value otherwise.
func (o *CpReport) GetSetIds() []string {
	if o == nil || IsNil(o.SetIds) {
		var ret []string
		return ret
	}
	return o.SetIds
}

// GetSetIdsOk returns a tuple with the SetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpReport) GetSetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SetIds) {
		return nil, false
	}
	return o.SetIds, true
}

// HasSetIds returns a boolean if a field has been set.
func (o *CpReport) HasSetIds() bool {
	if o != nil && !IsNil(o.SetIds) {
		return true
	}

	return false
}

// SetSetIds gets a reference to the given []string and assigns it to the SetIds field.
func (o *CpReport) SetSetIds(v []string) {
	o.SetIds = v
}

// GetFailureCode returns the FailureCode field value
func (o *CpReport) GetFailureCode() CpFailureCode {
	if o == nil {
		var ret CpFailureCode
		return ret
	}

	return o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
func (o *CpReport) GetFailureCodeOk() (*CpFailureCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCode, true
}

// SetFailureCode sets field value
func (o *CpReport) SetFailureCode(v CpFailureCode) {
	o.FailureCode = v
}

func (o CpReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SetIds) {
		toSerialize["setIds"] = o.SetIds
	}
	toSerialize["failureCode"] = o.FailureCode
	return toSerialize, nil
}

type NullableCpReport struct {
	value *CpReport
	isSet bool
}

func (v NullableCpReport) Get() *CpReport {
	return v.value
}

func (v *NullableCpReport) Set(val *CpReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCpReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCpReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpReport(val *CpReport) *NullableCpReport {
	return &NullableCpReport{value: val, isSet: true}
}

func (v NullableCpReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


