/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the MDAReportSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAReportSingleAllOfAttributesAllOf{}

// MDAReportSingleAllOfAttributesAllOf struct for MDAReportSingleAllOfAttributesAllOf
type MDAReportSingleAllOfAttributesAllOf struct {
	MDAReportID *string      `json:"mDAReportID,omitempty"`
	MDAOutputs  *MDAOutputs1 `json:"mDAOutputs,omitempty"`
}

// NewMDAReportSingleAllOfAttributesAllOf instantiates a new MDAReportSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAReportSingleAllOfAttributesAllOf() *MDAReportSingleAllOfAttributesAllOf {
	this := MDAReportSingleAllOfAttributesAllOf{}
	return &this
}

// NewMDAReportSingleAllOfAttributesAllOfWithDefaults instantiates a new MDAReportSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAReportSingleAllOfAttributesAllOfWithDefaults() *MDAReportSingleAllOfAttributesAllOf {
	this := MDAReportSingleAllOfAttributesAllOf{}
	return &this
}

// GetMDAReportID returns the MDAReportID field value if set, zero value otherwise.
func (o *MDAReportSingleAllOfAttributesAllOf) GetMDAReportID() string {
	if o == nil || IsNil(o.MDAReportID) {
		var ret string
		return ret
	}
	return *o.MDAReportID
}

// GetMDAReportIDOk returns a tuple with the MDAReportID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportSingleAllOfAttributesAllOf) GetMDAReportIDOk() (*string, bool) {
	if o == nil || IsNil(o.MDAReportID) {
		return nil, false
	}
	return o.MDAReportID, true
}

// HasMDAReportID returns a boolean if a field has been set.
func (o *MDAReportSingleAllOfAttributesAllOf) HasMDAReportID() bool {
	if o != nil && !IsNil(o.MDAReportID) {
		return true
	}

	return false
}

// SetMDAReportID gets a reference to the given string and assigns it to the MDAReportID field.
func (o *MDAReportSingleAllOfAttributesAllOf) SetMDAReportID(v string) {
	o.MDAReportID = &v
}

// GetMDAOutputs returns the MDAOutputs field value if set, zero value otherwise.
func (o *MDAReportSingleAllOfAttributesAllOf) GetMDAOutputs() MDAOutputs1 {
	if o == nil || IsNil(o.MDAOutputs) {
		var ret MDAOutputs1
		return ret
	}
	return *o.MDAOutputs
}

// GetMDAOutputsOk returns a tuple with the MDAOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportSingleAllOfAttributesAllOf) GetMDAOutputsOk() (*MDAOutputs1, bool) {
	if o == nil || IsNil(o.MDAOutputs) {
		return nil, false
	}
	return o.MDAOutputs, true
}

// HasMDAOutputs returns a boolean if a field has been set.
func (o *MDAReportSingleAllOfAttributesAllOf) HasMDAOutputs() bool {
	if o != nil && !IsNil(o.MDAOutputs) {
		return true
	}

	return false
}

// SetMDAOutputs gets a reference to the given MDAOutputs1 and assigns it to the MDAOutputs field.
func (o *MDAReportSingleAllOfAttributesAllOf) SetMDAOutputs(v MDAOutputs1) {
	o.MDAOutputs = &v
}

func (o MDAReportSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAReportSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAReportID) {
		toSerialize["mDAReportID"] = o.MDAReportID
	}
	if !IsNil(o.MDAOutputs) {
		toSerialize["mDAOutputs"] = o.MDAOutputs
	}
	return toSerialize, nil
}

type NullableMDAReportSingleAllOfAttributesAllOf struct {
	value *MDAReportSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableMDAReportSingleAllOfAttributesAllOf) Get() *MDAReportSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableMDAReportSingleAllOfAttributesAllOf) Set(val *MDAReportSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAReportSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAReportSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAReportSingleAllOfAttributesAllOf(val *MDAReportSingleAllOfAttributesAllOf) *NullableMDAReportSingleAllOfAttributesAllOf {
	return &NullableMDAReportSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableMDAReportSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAReportSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
