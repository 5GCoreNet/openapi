/*
MDA Report

OAS 3.0.1 specification of the MDA Report © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaReport

import (
	"encoding/json"
)

// checks if the MDAReportAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAReportAllOfAttributes{}

// MDAReportAllOfAttributes struct for MDAReportAllOfAttributes
type MDAReportAllOfAttributes struct {
	MDAReportID *string     `json:"mDAReportID,omitempty"`
	MDAOutputs  *MDAOutputs `json:"mDAOutputs,omitempty"`
}

// NewMDAReportAllOfAttributes instantiates a new MDAReportAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAReportAllOfAttributes() *MDAReportAllOfAttributes {
	this := MDAReportAllOfAttributes{}
	return &this
}

// NewMDAReportAllOfAttributesWithDefaults instantiates a new MDAReportAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAReportAllOfAttributesWithDefaults() *MDAReportAllOfAttributes {
	this := MDAReportAllOfAttributes{}
	return &this
}

// GetMDAReportID returns the MDAReportID field value if set, zero value otherwise.
func (o *MDAReportAllOfAttributes) GetMDAReportID() string {
	if o == nil || IsNil(o.MDAReportID) {
		var ret string
		return ret
	}
	return *o.MDAReportID
}

// GetMDAReportIDOk returns a tuple with the MDAReportID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportAllOfAttributes) GetMDAReportIDOk() (*string, bool) {
	if o == nil || IsNil(o.MDAReportID) {
		return nil, false
	}
	return o.MDAReportID, true
}

// HasMDAReportID returns a boolean if a field has been set.
func (o *MDAReportAllOfAttributes) HasMDAReportID() bool {
	if o != nil && !IsNil(o.MDAReportID) {
		return true
	}

	return false
}

// SetMDAReportID gets a reference to the given string and assigns it to the MDAReportID field.
func (o *MDAReportAllOfAttributes) SetMDAReportID(v string) {
	o.MDAReportID = &v
}

// GetMDAOutputs returns the MDAOutputs field value if set, zero value otherwise.
func (o *MDAReportAllOfAttributes) GetMDAOutputs() MDAOutputs {
	if o == nil || IsNil(o.MDAOutputs) {
		var ret MDAOutputs
		return ret
	}
	return *o.MDAOutputs
}

// GetMDAOutputsOk returns a tuple with the MDAOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportAllOfAttributes) GetMDAOutputsOk() (*MDAOutputs, bool) {
	if o == nil || IsNil(o.MDAOutputs) {
		return nil, false
	}
	return o.MDAOutputs, true
}

// HasMDAOutputs returns a boolean if a field has been set.
func (o *MDAReportAllOfAttributes) HasMDAOutputs() bool {
	if o != nil && !IsNil(o.MDAOutputs) {
		return true
	}

	return false
}

// SetMDAOutputs gets a reference to the given MDAOutputs and assigns it to the MDAOutputs field.
func (o *MDAReportAllOfAttributes) SetMDAOutputs(v MDAOutputs) {
	o.MDAOutputs = &v
}

func (o MDAReportAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAReportAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAReportID) {
		toSerialize["mDAReportID"] = o.MDAReportID
	}
	if !IsNil(o.MDAOutputs) {
		toSerialize["mDAOutputs"] = o.MDAOutputs
	}
	return toSerialize, nil
}

type NullableMDAReportAllOfAttributes struct {
	value *MDAReportAllOfAttributes
	isSet bool
}

func (v NullableMDAReportAllOfAttributes) Get() *MDAReportAllOfAttributes {
	return v.value
}

func (v *NullableMDAReportAllOfAttributes) Set(val *MDAReportAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAReportAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAReportAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAReportAllOfAttributes(val *MDAReportAllOfAttributes) *NullableMDAReportAllOfAttributes {
	return &NullableMDAReportAllOfAttributes{value: val, isSet: true}
}

func (v NullableMDAReportAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAReportAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
