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

// checks if the GeoAreaToCellMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoAreaToCellMapping{}

// GeoAreaToCellMapping struct for GeoAreaToCellMapping
type GeoAreaToCellMapping struct {
	GeoArea *GeoArea `json:"geoArea,omitempty"`
	AssociationThreshold *int32 `json:"associationThreshold,omitempty"`
}

// NewGeoAreaToCellMapping instantiates a new GeoAreaToCellMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoAreaToCellMapping() *GeoAreaToCellMapping {
	this := GeoAreaToCellMapping{}
	return &this
}

// NewGeoAreaToCellMappingWithDefaults instantiates a new GeoAreaToCellMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoAreaToCellMappingWithDefaults() *GeoAreaToCellMapping {
	this := GeoAreaToCellMapping{}
	return &this
}

// GetGeoArea returns the GeoArea field value if set, zero value otherwise.
func (o *GeoAreaToCellMapping) GetGeoArea() GeoArea {
	if o == nil || IsNil(o.GeoArea) {
		var ret GeoArea
		return ret
	}
	return *o.GeoArea
}

// GetGeoAreaOk returns a tuple with the GeoArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoAreaToCellMapping) GetGeoAreaOk() (*GeoArea, bool) {
	if o == nil || IsNil(o.GeoArea) {
		return nil, false
	}
	return o.GeoArea, true
}

// HasGeoArea returns a boolean if a field has been set.
func (o *GeoAreaToCellMapping) HasGeoArea() bool {
	if o != nil && !IsNil(o.GeoArea) {
		return true
	}

	return false
}

// SetGeoArea gets a reference to the given GeoArea and assigns it to the GeoArea field.
func (o *GeoAreaToCellMapping) SetGeoArea(v GeoArea) {
	o.GeoArea = &v
}

// GetAssociationThreshold returns the AssociationThreshold field value if set, zero value otherwise.
func (o *GeoAreaToCellMapping) GetAssociationThreshold() int32 {
	if o == nil || IsNil(o.AssociationThreshold) {
		var ret int32
		return ret
	}
	return *o.AssociationThreshold
}

// GetAssociationThresholdOk returns a tuple with the AssociationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoAreaToCellMapping) GetAssociationThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.AssociationThreshold) {
		return nil, false
	}
	return o.AssociationThreshold, true
}

// HasAssociationThreshold returns a boolean if a field has been set.
func (o *GeoAreaToCellMapping) HasAssociationThreshold() bool {
	if o != nil && !IsNil(o.AssociationThreshold) {
		return true
	}

	return false
}

// SetAssociationThreshold gets a reference to the given int32 and assigns it to the AssociationThreshold field.
func (o *GeoAreaToCellMapping) SetAssociationThreshold(v int32) {
	o.AssociationThreshold = &v
}

func (o GeoAreaToCellMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoAreaToCellMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeoArea) {
		toSerialize["geoArea"] = o.GeoArea
	}
	if !IsNil(o.AssociationThreshold) {
		toSerialize["associationThreshold"] = o.AssociationThreshold
	}
	return toSerialize, nil
}

type NullableGeoAreaToCellMapping struct {
	value *GeoAreaToCellMapping
	isSet bool
}

func (v NullableGeoAreaToCellMapping) Get() *GeoAreaToCellMapping {
	return v.value
}

func (v *NullableGeoAreaToCellMapping) Set(val *GeoAreaToCellMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoAreaToCellMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoAreaToCellMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoAreaToCellMapping(val *GeoAreaToCellMapping) *NullableGeoAreaToCellMapping {
	return &NullableGeoAreaToCellMapping{value: val, isSet: true}
}

func (v NullableGeoAreaToCellMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoAreaToCellMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


