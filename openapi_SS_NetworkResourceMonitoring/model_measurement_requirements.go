/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
)

// checks if the MeasurementRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeasurementRequirements{}

// MeasurementRequirements Indicates the measurement requirements.
type MeasurementRequirements struct {
	// Indicates the required the QoS measurement data types.
	MeasDataTypes []MeasurementDataType `json:"measDataTypes"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	MeasAggrGranWnd *int32             `json:"measAggrGranWnd,omitempty"`
	MeasPeriod      *MeasurementPeriod `json:"measPeriod,omitempty"`
}

// NewMeasurementRequirements instantiates a new MeasurementRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementRequirements(measDataTypes []MeasurementDataType) *MeasurementRequirements {
	this := MeasurementRequirements{}
	this.MeasDataTypes = measDataTypes
	var measAggrGranWnd int32 = 2000
	this.MeasAggrGranWnd = &measAggrGranWnd
	return &this
}

// NewMeasurementRequirementsWithDefaults instantiates a new MeasurementRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementRequirementsWithDefaults() *MeasurementRequirements {
	this := MeasurementRequirements{}
	var measAggrGranWnd int32 = 2000
	this.MeasAggrGranWnd = &measAggrGranWnd
	return &this
}

// GetMeasDataTypes returns the MeasDataTypes field value
func (o *MeasurementRequirements) GetMeasDataTypes() []MeasurementDataType {
	if o == nil {
		var ret []MeasurementDataType
		return ret
	}

	return o.MeasDataTypes
}

// GetMeasDataTypesOk returns a tuple with the MeasDataTypes field value
// and a boolean to check if the value has been set.
func (o *MeasurementRequirements) GetMeasDataTypesOk() ([]MeasurementDataType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeasDataTypes, true
}

// SetMeasDataTypes sets field value
func (o *MeasurementRequirements) SetMeasDataTypes(v []MeasurementDataType) {
	o.MeasDataTypes = v
}

// GetMeasAggrGranWnd returns the MeasAggrGranWnd field value if set, zero value otherwise.
func (o *MeasurementRequirements) GetMeasAggrGranWnd() int32 {
	if o == nil || IsNil(o.MeasAggrGranWnd) {
		var ret int32
		return ret
	}
	return *o.MeasAggrGranWnd
}

// GetMeasAggrGranWndOk returns a tuple with the MeasAggrGranWnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementRequirements) GetMeasAggrGranWndOk() (*int32, bool) {
	if o == nil || IsNil(o.MeasAggrGranWnd) {
		return nil, false
	}
	return o.MeasAggrGranWnd, true
}

// HasMeasAggrGranWnd returns a boolean if a field has been set.
func (o *MeasurementRequirements) HasMeasAggrGranWnd() bool {
	if o != nil && !IsNil(o.MeasAggrGranWnd) {
		return true
	}

	return false
}

// SetMeasAggrGranWnd gets a reference to the given int32 and assigns it to the MeasAggrGranWnd field.
func (o *MeasurementRequirements) SetMeasAggrGranWnd(v int32) {
	o.MeasAggrGranWnd = &v
}

// GetMeasPeriod returns the MeasPeriod field value if set, zero value otherwise.
func (o *MeasurementRequirements) GetMeasPeriod() MeasurementPeriod {
	if o == nil || IsNil(o.MeasPeriod) {
		var ret MeasurementPeriod
		return ret
	}
	return *o.MeasPeriod
}

// GetMeasPeriodOk returns a tuple with the MeasPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementRequirements) GetMeasPeriodOk() (*MeasurementPeriod, bool) {
	if o == nil || IsNil(o.MeasPeriod) {
		return nil, false
	}
	return o.MeasPeriod, true
}

// HasMeasPeriod returns a boolean if a field has been set.
func (o *MeasurementRequirements) HasMeasPeriod() bool {
	if o != nil && !IsNil(o.MeasPeriod) {
		return true
	}

	return false
}

// SetMeasPeriod gets a reference to the given MeasurementPeriod and assigns it to the MeasPeriod field.
func (o *MeasurementRequirements) SetMeasPeriod(v MeasurementPeriod) {
	o.MeasPeriod = &v
}

func (o MeasurementRequirements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeasurementRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["measDataTypes"] = o.MeasDataTypes
	if !IsNil(o.MeasAggrGranWnd) {
		toSerialize["measAggrGranWnd"] = o.MeasAggrGranWnd
	}
	if !IsNil(o.MeasPeriod) {
		toSerialize["measPeriod"] = o.MeasPeriod
	}
	return toSerialize, nil
}

type NullableMeasurementRequirements struct {
	value *MeasurementRequirements
	isSet bool
}

func (v NullableMeasurementRequirements) Get() *MeasurementRequirements {
	return v.value
}

func (v *NullableMeasurementRequirements) Set(val *MeasurementRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementRequirements(val *MeasurementRequirements) *NullableMeasurementRequirements {
	return &NullableMeasurementRequirements{value: val, isSet: true}
}

func (v NullableMeasurementRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
