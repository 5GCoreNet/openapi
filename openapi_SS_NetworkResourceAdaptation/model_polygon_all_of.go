/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// checks if the PolygonAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolygonAllOf{}

// PolygonAllOf struct for PolygonAllOf
type PolygonAllOf struct {
	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}

// NewPolygonAllOf instantiates a new PolygonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolygonAllOf(pointList []GeographicalCoordinates) *PolygonAllOf {
	this := PolygonAllOf{}
	this.PointList = pointList
	return &this
}

// NewPolygonAllOfWithDefaults instantiates a new PolygonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolygonAllOfWithDefaults() *PolygonAllOf {
	this := PolygonAllOf{}
	return &this
}

// GetPointList returns the PointList field value
func (o *PolygonAllOf) GetPointList() []GeographicalCoordinates {
	if o == nil {
		var ret []GeographicalCoordinates
		return ret
	}

	return o.PointList
}

// GetPointListOk returns a tuple with the PointList field value
// and a boolean to check if the value has been set.
func (o *PolygonAllOf) GetPointListOk() ([]GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return o.PointList, true
}

// SetPointList sets field value
func (o *PolygonAllOf) SetPointList(v []GeographicalCoordinates) {
	o.PointList = v
}

func (o PolygonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolygonAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pointList"] = o.PointList
	return toSerialize, nil
}

type NullablePolygonAllOf struct {
	value *PolygonAllOf
	isSet bool
}

func (v NullablePolygonAllOf) Get() *PolygonAllOf {
	return v.value
}

func (v *NullablePolygonAllOf) Set(val *PolygonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolygonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolygonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolygonAllOf(val *PolygonAllOf) *NullablePolygonAllOf {
	return &NullablePolygonAllOf{value: val, isSet: true}
}

func (v NullablePolygonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolygonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
