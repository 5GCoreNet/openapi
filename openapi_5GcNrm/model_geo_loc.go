/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the GeoLoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoLoc{}

// GeoLoc struct for GeoLoc
type GeoLoc struct {
	GeographicalCoordinates *GeographicalCoordinates `json:"geographicalCoordinates,omitempty"`
	CivicLocation           *string                  `json:"civicLocation,omitempty"`
}

// NewGeoLoc instantiates a new GeoLoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoLoc() *GeoLoc {
	this := GeoLoc{}
	return &this
}

// NewGeoLocWithDefaults instantiates a new GeoLoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoLocWithDefaults() *GeoLoc {
	this := GeoLoc{}
	return &this
}

// GetGeographicalCoordinates returns the GeographicalCoordinates field value if set, zero value otherwise.
func (o *GeoLoc) GetGeographicalCoordinates() GeographicalCoordinates {
	if o == nil || IsNil(o.GeographicalCoordinates) {
		var ret GeographicalCoordinates
		return ret
	}
	return *o.GeographicalCoordinates
}

// GetGeographicalCoordinatesOk returns a tuple with the GeographicalCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLoc) GetGeographicalCoordinatesOk() (*GeographicalCoordinates, bool) {
	if o == nil || IsNil(o.GeographicalCoordinates) {
		return nil, false
	}
	return o.GeographicalCoordinates, true
}

// HasGeographicalCoordinates returns a boolean if a field has been set.
func (o *GeoLoc) HasGeographicalCoordinates() bool {
	if o != nil && !IsNil(o.GeographicalCoordinates) {
		return true
	}

	return false
}

// SetGeographicalCoordinates gets a reference to the given GeographicalCoordinates and assigns it to the GeographicalCoordinates field.
func (o *GeoLoc) SetGeographicalCoordinates(v GeographicalCoordinates) {
	o.GeographicalCoordinates = &v
}

// GetCivicLocation returns the CivicLocation field value if set, zero value otherwise.
func (o *GeoLoc) GetCivicLocation() string {
	if o == nil || IsNil(o.CivicLocation) {
		var ret string
		return ret
	}
	return *o.CivicLocation
}

// GetCivicLocationOk returns a tuple with the CivicLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLoc) GetCivicLocationOk() (*string, bool) {
	if o == nil || IsNil(o.CivicLocation) {
		return nil, false
	}
	return o.CivicLocation, true
}

// HasCivicLocation returns a boolean if a field has been set.
func (o *GeoLoc) HasCivicLocation() bool {
	if o != nil && !IsNil(o.CivicLocation) {
		return true
	}

	return false
}

// SetCivicLocation gets a reference to the given string and assigns it to the CivicLocation field.
func (o *GeoLoc) SetCivicLocation(v string) {
	o.CivicLocation = &v
}

func (o GeoLoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoLoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeographicalCoordinates) {
		toSerialize["geographicalCoordinates"] = o.GeographicalCoordinates
	}
	if !IsNil(o.CivicLocation) {
		toSerialize["civicLocation"] = o.CivicLocation
	}
	return toSerialize, nil
}

type NullableGeoLoc struct {
	value *GeoLoc
	isSet bool
}

func (v NullableGeoLoc) Get() *GeoLoc {
	return v.value
}

func (v *NullableGeoLoc) Set(val *GeoLoc) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoLoc) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoLoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoLoc(val *GeoLoc) *NullableGeoLoc {
	return &NullableGeoLoc{value: val, isSet: true}
}

func (v NullableGeoLoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoLoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
