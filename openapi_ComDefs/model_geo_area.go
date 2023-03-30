/*
Common Type Definitions

OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ComDefs

import (
	"encoding/json"
)

// checks if the GeoArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoArea{}

// GeoArea struct for GeoArea
type GeoArea struct {
	ConvexGeoPolygon []GeoCoordinate `json:"convexGeoPolygon,omitempty"`
}

// NewGeoArea instantiates a new GeoArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoArea() *GeoArea {
	this := GeoArea{}
	return &this
}

// NewGeoAreaWithDefaults instantiates a new GeoArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoAreaWithDefaults() *GeoArea {
	this := GeoArea{}
	return &this
}

// GetConvexGeoPolygon returns the ConvexGeoPolygon field value if set, zero value otherwise.
func (o *GeoArea) GetConvexGeoPolygon() []GeoCoordinate {
	if o == nil || IsNil(o.ConvexGeoPolygon) {
		var ret []GeoCoordinate
		return ret
	}
	return o.ConvexGeoPolygon
}

// GetConvexGeoPolygonOk returns a tuple with the ConvexGeoPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoArea) GetConvexGeoPolygonOk() ([]GeoCoordinate, bool) {
	if o == nil || IsNil(o.ConvexGeoPolygon) {
		return nil, false
	}
	return o.ConvexGeoPolygon, true
}

// HasConvexGeoPolygon returns a boolean if a field has been set.
func (o *GeoArea) HasConvexGeoPolygon() bool {
	if o != nil && !IsNil(o.ConvexGeoPolygon) {
		return true
	}

	return false
}

// SetConvexGeoPolygon gets a reference to the given []GeoCoordinate and assigns it to the ConvexGeoPolygon field.
func (o *GeoArea) SetConvexGeoPolygon(v []GeoCoordinate) {
	o.ConvexGeoPolygon = v
}

func (o GeoArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConvexGeoPolygon) {
		toSerialize["convexGeoPolygon"] = o.ConvexGeoPolygon
	}
	return toSerialize, nil
}

type NullableGeoArea struct {
	value *GeoArea
	isSet bool
}

func (v NullableGeoArea) Get() *GeoArea {
	return v.value
}

func (v *NullableGeoArea) Set(val *GeoArea) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoArea) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoArea(val *GeoArea) *NullableGeoArea {
	return &NullableGeoArea{value: val, isSet: true}
}

func (v NullableGeoArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


