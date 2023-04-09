/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the MDAReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAReport{}

// MDAReport struct for MDAReport
type MDAReport struct {
	Top1
	Attributes *MDAReportAllOfAttributes `json:"attributes,omitempty"`
}

// NewMDAReport instantiates a new MDAReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAReport(id NullableString) *MDAReport {
	this := MDAReport{}
	this.Id = id
	return &this
}

// NewMDAReportWithDefaults instantiates a new MDAReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAReportWithDefaults() *MDAReport {
	this := MDAReport{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MDAReport) GetAttributes() MDAReportAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret MDAReportAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReport) GetAttributesOk() (*MDAReportAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MDAReport) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MDAReportAllOfAttributes and assigns it to the Attributes field.
func (o *MDAReport) SetAttributes(v MDAReportAllOfAttributes) {
	o.Attributes = &v
}

func (o MDAReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop1, errTop1 := json.Marshal(o.Top1)
	if errTop1 != nil {
		return map[string]interface{}{}, errTop1
	}
	errTop1 = json.Unmarshal([]byte(serializedTop1), &toSerialize)
	if errTop1 != nil {
		return map[string]interface{}{}, errTop1
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMDAReport struct {
	value *MDAReport
	isSet bool
}

func (v NullableMDAReport) Get() *MDAReport {
	return v.value
}

func (v *NullableMDAReport) Set(val *MDAReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAReport(val *MDAReport) *NullableMDAReport {
	return &NullableMDAReport{value: val, isSet: true}
}

func (v NullableMDAReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
