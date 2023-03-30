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

// checks if the SpatialValidityCond1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpatialValidityCond1{}

// SpatialValidityCond1 Contains the Spatial Validity Condition.
type SpatialValidityCond1 struct {
	TrackingAreaList []Tai1 `json:"trackingAreaList,omitempty"`
	Countries []string `json:"countries,omitempty"`
	GeographicalServiceArea *GeoServiceArea1 `json:"geographicalServiceArea,omitempty"`
}

// NewSpatialValidityCond1 instantiates a new SpatialValidityCond1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpatialValidityCond1() *SpatialValidityCond1 {
	this := SpatialValidityCond1{}
	return &this
}

// NewSpatialValidityCond1WithDefaults instantiates a new SpatialValidityCond1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpatialValidityCond1WithDefaults() *SpatialValidityCond1 {
	this := SpatialValidityCond1{}
	return &this
}

// GetTrackingAreaList returns the TrackingAreaList field value if set, zero value otherwise.
func (o *SpatialValidityCond1) GetTrackingAreaList() []Tai1 {
	if o == nil || IsNil(o.TrackingAreaList) {
		var ret []Tai1
		return ret
	}
	return o.TrackingAreaList
}

// GetTrackingAreaListOk returns a tuple with the TrackingAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond1) GetTrackingAreaListOk() ([]Tai1, bool) {
	if o == nil || IsNil(o.TrackingAreaList) {
		return nil, false
	}
	return o.TrackingAreaList, true
}

// HasTrackingAreaList returns a boolean if a field has been set.
func (o *SpatialValidityCond1) HasTrackingAreaList() bool {
	if o != nil && !IsNil(o.TrackingAreaList) {
		return true
	}

	return false
}

// SetTrackingAreaList gets a reference to the given []Tai1 and assigns it to the TrackingAreaList field.
func (o *SpatialValidityCond1) SetTrackingAreaList(v []Tai1) {
	o.TrackingAreaList = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *SpatialValidityCond1) GetCountries() []string {
	if o == nil || IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond1) GetCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *SpatialValidityCond1) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *SpatialValidityCond1) SetCountries(v []string) {
	o.Countries = v
}

// GetGeographicalServiceArea returns the GeographicalServiceArea field value if set, zero value otherwise.
func (o *SpatialValidityCond1) GetGeographicalServiceArea() GeoServiceArea1 {
	if o == nil || IsNil(o.GeographicalServiceArea) {
		var ret GeoServiceArea1
		return ret
	}
	return *o.GeographicalServiceArea
}

// GetGeographicalServiceAreaOk returns a tuple with the GeographicalServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond1) GetGeographicalServiceAreaOk() (*GeoServiceArea1, bool) {
	if o == nil || IsNil(o.GeographicalServiceArea) {
		return nil, false
	}
	return o.GeographicalServiceArea, true
}

// HasGeographicalServiceArea returns a boolean if a field has been set.
func (o *SpatialValidityCond1) HasGeographicalServiceArea() bool {
	if o != nil && !IsNil(o.GeographicalServiceArea) {
		return true
	}

	return false
}

// SetGeographicalServiceArea gets a reference to the given GeoServiceArea1 and assigns it to the GeographicalServiceArea field.
func (o *SpatialValidityCond1) SetGeographicalServiceArea(v GeoServiceArea1) {
	o.GeographicalServiceArea = &v
}

func (o SpatialValidityCond1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpatialValidityCond1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingAreaList) {
		toSerialize["trackingAreaList"] = o.TrackingAreaList
	}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !IsNil(o.GeographicalServiceArea) {
		toSerialize["geographicalServiceArea"] = o.GeographicalServiceArea
	}
	return toSerialize, nil
}

type NullableSpatialValidityCond1 struct {
	value *SpatialValidityCond1
	isSet bool
}

func (v NullableSpatialValidityCond1) Get() *SpatialValidityCond1 {
	return v.value
}

func (v *NullableSpatialValidityCond1) Set(val *SpatialValidityCond1) {
	v.value = val
	v.isSet = true
}

func (v NullableSpatialValidityCond1) IsSet() bool {
	return v.isSet
}

func (v *NullableSpatialValidityCond1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpatialValidityCond1(val *SpatialValidityCond1) *NullableSpatialValidityCond1 {
	return &NullableSpatialValidityCond1{value: val, isSet: true}
}

func (v NullableSpatialValidityCond1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpatialValidityCond1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


