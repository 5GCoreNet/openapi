/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
)

// checks if the ACRCompleteEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACRCompleteEventInfo{}

// ACRCompleteEventInfo Indicates the completed ACR result and target EAS endpoint info.
type ACRCompleteEventInfo struct {
	// Indicates whether the ACR is successful or failure.
	AcrRes bool `json:"acrRes"`
	TEasEndpoint EndPoint `json:"tEasEndpoint"`
	// Indicates the cause information for the failure.
	FailReason *string `json:"failReason,omitempty"`
}

// NewACRCompleteEventInfo instantiates a new ACRCompleteEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACRCompleteEventInfo(acrRes bool, tEasEndpoint EndPoint) *ACRCompleteEventInfo {
	this := ACRCompleteEventInfo{}
	this.AcrRes = acrRes
	this.TEasEndpoint = tEasEndpoint
	return &this
}

// NewACRCompleteEventInfoWithDefaults instantiates a new ACRCompleteEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACRCompleteEventInfoWithDefaults() *ACRCompleteEventInfo {
	this := ACRCompleteEventInfo{}
	return &this
}

// GetAcrRes returns the AcrRes field value
func (o *ACRCompleteEventInfo) GetAcrRes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcrRes
}

// GetAcrResOk returns a tuple with the AcrRes field value
// and a boolean to check if the value has been set.
func (o *ACRCompleteEventInfo) GetAcrResOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcrRes, true
}

// SetAcrRes sets field value
func (o *ACRCompleteEventInfo) SetAcrRes(v bool) {
	o.AcrRes = v
}

// GetTEasEndpoint returns the TEasEndpoint field value
func (o *ACRCompleteEventInfo) GetTEasEndpoint() EndPoint {
	if o == nil {
		var ret EndPoint
		return ret
	}

	return o.TEasEndpoint
}

// GetTEasEndpointOk returns a tuple with the TEasEndpoint field value
// and a boolean to check if the value has been set.
func (o *ACRCompleteEventInfo) GetTEasEndpointOk() (*EndPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEasEndpoint, true
}

// SetTEasEndpoint sets field value
func (o *ACRCompleteEventInfo) SetTEasEndpoint(v EndPoint) {
	o.TEasEndpoint = v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *ACRCompleteEventInfo) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACRCompleteEventInfo) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *ACRCompleteEventInfo) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *ACRCompleteEventInfo) SetFailReason(v string) {
	o.FailReason = &v
}

func (o ACRCompleteEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACRCompleteEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acrRes"] = o.AcrRes
	toSerialize["tEasEndpoint"] = o.TEasEndpoint
	if !IsNil(o.FailReason) {
		toSerialize["failReason"] = o.FailReason
	}
	return toSerialize, nil
}

type NullableACRCompleteEventInfo struct {
	value *ACRCompleteEventInfo
	isSet bool
}

func (v NullableACRCompleteEventInfo) Get() *ACRCompleteEventInfo {
	return v.value
}

func (v *NullableACRCompleteEventInfo) Set(val *ACRCompleteEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableACRCompleteEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableACRCompleteEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRCompleteEventInfo(val *ACRCompleteEventInfo) *NullableACRCompleteEventInfo {
	return &NullableACRCompleteEventInfo{value: val, isSet: true}
}

func (v NullableACRCompleteEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRCompleteEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


