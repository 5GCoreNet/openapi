/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
)

// checks if the ACTResultInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACTResultInfo{}

// ACTResultInfo Represents the result of ACT and the related information.
type ACTResultInfo struct {
	ActResult ACTResult `json:"actResult"`
	ActFailureCause *ACTFailureCause `json:"actFailureCause,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId string `json:"ueId"`
	EasEndPoint EndPoint `json:"easEndPoint"`
}

// NewACTResultInfo instantiates a new ACTResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACTResultInfo(actResult ACTResult, ueId string, easEndPoint EndPoint) *ACTResultInfo {
	this := ACTResultInfo{}
	this.ActResult = actResult
	this.UeId = ueId
	this.EasEndPoint = easEndPoint
	return &this
}

// NewACTResultInfoWithDefaults instantiates a new ACTResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACTResultInfoWithDefaults() *ACTResultInfo {
	this := ACTResultInfo{}
	return &this
}

// GetActResult returns the ActResult field value
func (o *ACTResultInfo) GetActResult() ACTResult {
	if o == nil {
		var ret ACTResult
		return ret
	}

	return o.ActResult
}

// GetActResultOk returns a tuple with the ActResult field value
// and a boolean to check if the value has been set.
func (o *ACTResultInfo) GetActResultOk() (*ACTResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActResult, true
}

// SetActResult sets field value
func (o *ACTResultInfo) SetActResult(v ACTResult) {
	o.ActResult = v
}

// GetActFailureCause returns the ActFailureCause field value if set, zero value otherwise.
func (o *ACTResultInfo) GetActFailureCause() ACTFailureCause {
	if o == nil || IsNil(o.ActFailureCause) {
		var ret ACTFailureCause
		return ret
	}
	return *o.ActFailureCause
}

// GetActFailureCauseOk returns a tuple with the ActFailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACTResultInfo) GetActFailureCauseOk() (*ACTFailureCause, bool) {
	if o == nil || IsNil(o.ActFailureCause) {
		return nil, false
	}
	return o.ActFailureCause, true
}

// HasActFailureCause returns a boolean if a field has been set.
func (o *ACTResultInfo) HasActFailureCause() bool {
	if o != nil && !IsNil(o.ActFailureCause) {
		return true
	}

	return false
}

// SetActFailureCause gets a reference to the given ACTFailureCause and assigns it to the ActFailureCause field.
func (o *ACTResultInfo) SetActFailureCause(v ACTFailureCause) {
	o.ActFailureCause = &v
}

// GetUeId returns the UeId field value
func (o *ACTResultInfo) GetUeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *ACTResultInfo) GetUeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *ACTResultInfo) SetUeId(v string) {
	o.UeId = v
}

// GetEasEndPoint returns the EasEndPoint field value
func (o *ACTResultInfo) GetEasEndPoint() EndPoint {
	if o == nil {
		var ret EndPoint
		return ret
	}

	return o.EasEndPoint
}

// GetEasEndPointOk returns a tuple with the EasEndPoint field value
// and a boolean to check if the value has been set.
func (o *ACTResultInfo) GetEasEndPointOk() (*EndPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasEndPoint, true
}

// SetEasEndPoint sets field value
func (o *ACTResultInfo) SetEasEndPoint(v EndPoint) {
	o.EasEndPoint = v
}

func (o ACTResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACTResultInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actResult"] = o.ActResult
	if !IsNil(o.ActFailureCause) {
		toSerialize["actFailureCause"] = o.ActFailureCause
	}
	toSerialize["ueId"] = o.UeId
	toSerialize["easEndPoint"] = o.EasEndPoint
	return toSerialize, nil
}

type NullableACTResultInfo struct {
	value *ACTResultInfo
	isSet bool
}

func (v NullableACTResultInfo) Get() *ACTResultInfo {
	return v.value
}

func (v *NullableACTResultInfo) Set(val *ACTResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableACTResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableACTResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACTResultInfo(val *ACTResultInfo) *NullableACTResultInfo {
	return &NullableACTResultInfo{value: val, isSet: true}
}

func (v NullableACTResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACTResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

