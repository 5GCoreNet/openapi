/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the GeoCoordinate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoCoordinate{}

// GeoCoordinate struct for GeoCoordinate
type GeoCoordinate struct {
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
}

// NewGeoCoordinate instantiates a new GeoCoordinate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoCoordinate() *GeoCoordinate {
	this := GeoCoordinate{}
	return &this
}

// NewGeoCoordinateWithDefaults instantiates a new GeoCoordinate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoCoordinateWithDefaults() *GeoCoordinate {
	this := GeoCoordinate{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GeoCoordinate) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoCoordinate) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GeoCoordinate) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *GeoCoordinate) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GeoCoordinate) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoCoordinate) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GeoCoordinate) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *GeoCoordinate) SetLongitude(v float32) {
	o.Longitude = &v
}

func (o GeoCoordinate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoCoordinate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableGeoCoordinate struct {
	value *GeoCoordinate
	isSet bool
}

func (v NullableGeoCoordinate) Get() *GeoCoordinate {
	return v.value
}

func (v *NullableGeoCoordinate) Set(val *GeoCoordinate) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoCoordinate) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoCoordinate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoCoordinate(val *GeoCoordinate) *NullableGeoCoordinate {
	return &NullableGeoCoordinate{value: val, isSet: true}
}

func (v NullableGeoCoordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoCoordinate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

