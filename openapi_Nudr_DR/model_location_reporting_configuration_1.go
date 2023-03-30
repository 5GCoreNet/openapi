/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the LocationReportingConfiguration1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationReportingConfiguration1{}

// LocationReportingConfiguration1 struct for LocationReportingConfiguration1
type LocationReportingConfiguration1 struct {
	CurrentLocation bool `json:"currentLocation"`
	OneTime *bool `json:"oneTime,omitempty"`
	Accuracy *LocationAccuracy `json:"accuracy,omitempty"`
	N3gppAccuracy *LocationAccuracy `json:"n3gppAccuracy,omitempty"`
}

// NewLocationReportingConfiguration1 instantiates a new LocationReportingConfiguration1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationReportingConfiguration1(currentLocation bool) *LocationReportingConfiguration1 {
	this := LocationReportingConfiguration1{}
	this.CurrentLocation = currentLocation
	return &this
}

// NewLocationReportingConfiguration1WithDefaults instantiates a new LocationReportingConfiguration1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationReportingConfiguration1WithDefaults() *LocationReportingConfiguration1 {
	this := LocationReportingConfiguration1{}
	return &this
}

// GetCurrentLocation returns the CurrentLocation field value
func (o *LocationReportingConfiguration1) GetCurrentLocation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CurrentLocation
}

// GetCurrentLocationOk returns a tuple with the CurrentLocation field value
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration1) GetCurrentLocationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLocation, true
}

// SetCurrentLocation sets field value
func (o *LocationReportingConfiguration1) SetCurrentLocation(v bool) {
	o.CurrentLocation = v
}

// GetOneTime returns the OneTime field value if set, zero value otherwise.
func (o *LocationReportingConfiguration1) GetOneTime() bool {
	if o == nil || IsNil(o.OneTime) {
		var ret bool
		return ret
	}
	return *o.OneTime
}

// GetOneTimeOk returns a tuple with the OneTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration1) GetOneTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.OneTime) {
		return nil, false
	}
	return o.OneTime, true
}

// HasOneTime returns a boolean if a field has been set.
func (o *LocationReportingConfiguration1) HasOneTime() bool {
	if o != nil && !IsNil(o.OneTime) {
		return true
	}

	return false
}

// SetOneTime gets a reference to the given bool and assigns it to the OneTime field.
func (o *LocationReportingConfiguration1) SetOneTime(v bool) {
	o.OneTime = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *LocationReportingConfiguration1) GetAccuracy() LocationAccuracy {
	if o == nil || IsNil(o.Accuracy) {
		var ret LocationAccuracy
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration1) GetAccuracyOk() (*LocationAccuracy, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *LocationReportingConfiguration1) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given LocationAccuracy and assigns it to the Accuracy field.
func (o *LocationReportingConfiguration1) SetAccuracy(v LocationAccuracy) {
	o.Accuracy = &v
}

// GetN3gppAccuracy returns the N3gppAccuracy field value if set, zero value otherwise.
func (o *LocationReportingConfiguration1) GetN3gppAccuracy() LocationAccuracy {
	if o == nil || IsNil(o.N3gppAccuracy) {
		var ret LocationAccuracy
		return ret
	}
	return *o.N3gppAccuracy
}

// GetN3gppAccuracyOk returns a tuple with the N3gppAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration1) GetN3gppAccuracyOk() (*LocationAccuracy, bool) {
	if o == nil || IsNil(o.N3gppAccuracy) {
		return nil, false
	}
	return o.N3gppAccuracy, true
}

// HasN3gppAccuracy returns a boolean if a field has been set.
func (o *LocationReportingConfiguration1) HasN3gppAccuracy() bool {
	if o != nil && !IsNil(o.N3gppAccuracy) {
		return true
	}

	return false
}

// SetN3gppAccuracy gets a reference to the given LocationAccuracy and assigns it to the N3gppAccuracy field.
func (o *LocationReportingConfiguration1) SetN3gppAccuracy(v LocationAccuracy) {
	o.N3gppAccuracy = &v
}

func (o LocationReportingConfiguration1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationReportingConfiguration1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentLocation"] = o.CurrentLocation
	if !IsNil(o.OneTime) {
		toSerialize["oneTime"] = o.OneTime
	}
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	if !IsNil(o.N3gppAccuracy) {
		toSerialize["n3gppAccuracy"] = o.N3gppAccuracy
	}
	return toSerialize, nil
}

type NullableLocationReportingConfiguration1 struct {
	value *LocationReportingConfiguration1
	isSet bool
}

func (v NullableLocationReportingConfiguration1) Get() *LocationReportingConfiguration1 {
	return v.value
}

func (v *NullableLocationReportingConfiguration1) Set(val *LocationReportingConfiguration1) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationReportingConfiguration1) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationReportingConfiguration1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationReportingConfiguration1(val *LocationReportingConfiguration1) *NullableLocationReportingConfiguration1 {
	return &NullableLocationReportingConfiguration1{value: val, isSet: true}
}

func (v NullableLocationReportingConfiguration1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationReportingConfiguration1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


