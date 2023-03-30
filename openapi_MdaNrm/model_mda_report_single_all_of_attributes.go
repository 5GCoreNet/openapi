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

// checks if the MDAReportSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAReportSingleAllOfAttributes{}

// MDAReportSingleAllOfAttributes struct for MDAReportSingleAllOfAttributes
type MDAReportSingleAllOfAttributes struct {
	MDAReportID *string `json:"mDAReportID,omitempty"`
	MDAOutputs *MDAOutputs1 `json:"mDAOutputs,omitempty"`
}

// NewMDAReportSingleAllOfAttributes instantiates a new MDAReportSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAReportSingleAllOfAttributes() *MDAReportSingleAllOfAttributes {
	this := MDAReportSingleAllOfAttributes{}
	return &this
}

// NewMDAReportSingleAllOfAttributesWithDefaults instantiates a new MDAReportSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAReportSingleAllOfAttributesWithDefaults() *MDAReportSingleAllOfAttributes {
	this := MDAReportSingleAllOfAttributes{}
	return &this
}

// GetMDAReportID returns the MDAReportID field value if set, zero value otherwise.
func (o *MDAReportSingleAllOfAttributes) GetMDAReportID() string {
	if o == nil || IsNil(o.MDAReportID) {
		var ret string
		return ret
	}
	return *o.MDAReportID
}

// GetMDAReportIDOk returns a tuple with the MDAReportID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportSingleAllOfAttributes) GetMDAReportIDOk() (*string, bool) {
	if o == nil || IsNil(o.MDAReportID) {
		return nil, false
	}
	return o.MDAReportID, true
}

// HasMDAReportID returns a boolean if a field has been set.
func (o *MDAReportSingleAllOfAttributes) HasMDAReportID() bool {
	if o != nil && !IsNil(o.MDAReportID) {
		return true
	}

	return false
}

// SetMDAReportID gets a reference to the given string and assigns it to the MDAReportID field.
func (o *MDAReportSingleAllOfAttributes) SetMDAReportID(v string) {
	o.MDAReportID = &v
}

// GetMDAOutputs returns the MDAOutputs field value if set, zero value otherwise.
func (o *MDAReportSingleAllOfAttributes) GetMDAOutputs() MDAOutputs1 {
	if o == nil || IsNil(o.MDAOutputs) {
		var ret MDAOutputs1
		return ret
	}
	return *o.MDAOutputs
}

// GetMDAOutputsOk returns a tuple with the MDAOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportSingleAllOfAttributes) GetMDAOutputsOk() (*MDAOutputs1, bool) {
	if o == nil || IsNil(o.MDAOutputs) {
		return nil, false
	}
	return o.MDAOutputs, true
}

// HasMDAOutputs returns a boolean if a field has been set.
func (o *MDAReportSingleAllOfAttributes) HasMDAOutputs() bool {
	if o != nil && !IsNil(o.MDAOutputs) {
		return true
	}

	return false
}

// SetMDAOutputs gets a reference to the given MDAOutputs1 and assigns it to the MDAOutputs field.
func (o *MDAReportSingleAllOfAttributes) SetMDAOutputs(v MDAOutputs1) {
	o.MDAOutputs = &v
}

func (o MDAReportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAReportSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAReportID) {
		toSerialize["mDAReportID"] = o.MDAReportID
	}
	if !IsNil(o.MDAOutputs) {
		toSerialize["mDAOutputs"] = o.MDAOutputs
	}
	return toSerialize, nil
}

type NullableMDAReportSingleAllOfAttributes struct {
	value *MDAReportSingleAllOfAttributes
	isSet bool
}

func (v NullableMDAReportSingleAllOfAttributes) Get() *MDAReportSingleAllOfAttributes {
	return v.value
}

func (v *NullableMDAReportSingleAllOfAttributes) Set(val *MDAReportSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAReportSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAReportSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAReportSingleAllOfAttributes(val *MDAReportSingleAllOfAttributes) *NullableMDAReportSingleAllOfAttributes {
	return &NullableMDAReportSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableMDAReportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAReportSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


