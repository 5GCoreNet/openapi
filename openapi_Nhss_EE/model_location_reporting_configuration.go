/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
)

// checks if the LocationReportingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationReportingConfiguration{}

// LocationReportingConfiguration Contains data needed for a Monitoring Configuration, specific to the 'Location Reporting' event type
type LocationReportingConfiguration struct {
	CurrentLocation bool              `json:"currentLocation"`
	Accuracy        *LocationAccuracy `json:"accuracy,omitempty"`
}

// NewLocationReportingConfiguration instantiates a new LocationReportingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationReportingConfiguration(currentLocation bool) *LocationReportingConfiguration {
	this := LocationReportingConfiguration{}
	this.CurrentLocation = currentLocation
	return &this
}

// NewLocationReportingConfigurationWithDefaults instantiates a new LocationReportingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationReportingConfigurationWithDefaults() *LocationReportingConfiguration {
	this := LocationReportingConfiguration{}
	return &this
}

// GetCurrentLocation returns the CurrentLocation field value
func (o *LocationReportingConfiguration) GetCurrentLocation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CurrentLocation
}

// GetCurrentLocationOk returns a tuple with the CurrentLocation field value
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetCurrentLocationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLocation, true
}

// SetCurrentLocation sets field value
func (o *LocationReportingConfiguration) SetCurrentLocation(v bool) {
	o.CurrentLocation = v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *LocationReportingConfiguration) GetAccuracy() LocationAccuracy {
	if o == nil || IsNil(o.Accuracy) {
		var ret LocationAccuracy
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationReportingConfiguration) GetAccuracyOk() (*LocationAccuracy, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *LocationReportingConfiguration) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given LocationAccuracy and assigns it to the Accuracy field.
func (o *LocationReportingConfiguration) SetAccuracy(v LocationAccuracy) {
	o.Accuracy = &v
}

func (o LocationReportingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationReportingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentLocation"] = o.CurrentLocation
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return toSerialize, nil
}

type NullableLocationReportingConfiguration struct {
	value *LocationReportingConfiguration
	isSet bool
}

func (v NullableLocationReportingConfiguration) Get() *LocationReportingConfiguration {
	return v.value
}

func (v *NullableLocationReportingConfiguration) Set(val *LocationReportingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationReportingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationReportingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationReportingConfiguration(val *LocationReportingConfiguration) *NullableLocationReportingConfiguration {
	return &NullableLocationReportingConfiguration{value: val, isSet: true}
}

func (v NullableLocationReportingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationReportingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
