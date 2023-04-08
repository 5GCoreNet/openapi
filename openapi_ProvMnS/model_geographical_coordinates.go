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

// checks if the GeographicalCoordinates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeographicalCoordinates{}

// GeographicalCoordinates struct for GeographicalCoordinates
type GeographicalCoordinates struct {
	Latitude *int32 `json:"latitude,omitempty"`
	Longitude *int32 `json:"longitude,omitempty"`
}

// NewGeographicalCoordinates instantiates a new GeographicalCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicalCoordinates() *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	return &this
}

// NewGeographicalCoordinatesWithDefaults instantiates a new GeographicalCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicalCoordinatesWithDefaults() *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GeographicalCoordinates) GetLatitude() int32 {
	if o == nil || isNil(o.Latitude) {
		var ret int32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLatitudeOk() (*int32, bool) {
	if o == nil || isNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GeographicalCoordinates) HasLatitude() bool {
	if o != nil && !isNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given int32 and assigns it to the Latitude field.
func (o *GeographicalCoordinates) SetLatitude(v int32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GeographicalCoordinates) GetLongitude() int32 {
	if o == nil || isNil(o.Longitude) {
		var ret int32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLongitudeOk() (*int32, bool) {
	if o == nil || isNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GeographicalCoordinates) HasLongitude() bool {
	if o != nil && !isNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given int32 and assigns it to the Longitude field.
func (o *GeographicalCoordinates) SetLongitude(v int32) {
	o.Longitude = &v
}

func (o GeographicalCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeographicalCoordinates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !isNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableGeographicalCoordinates struct {
	value *GeographicalCoordinates
	isSet bool
}

func (v NullableGeographicalCoordinates) Get() *GeographicalCoordinates {
	return v.value
}

func (v *NullableGeographicalCoordinates) Set(val *GeographicalCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicalCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicalCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicalCoordinates(val *GeographicalCoordinates) *NullableGeographicalCoordinates {
	return &NullableGeographicalCoordinates{value: val, isSet: true}
}

func (v NullableGeographicalCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicalCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


