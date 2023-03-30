/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
)

// checks if the GeographicalArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeographicalArea{}

// GeographicalArea Contains geographical area information (e.g.a civic address or shapes).
type GeographicalArea struct {
	CivicAddress *CivicAddress `json:"civicAddress,omitempty"`
	Shapes *GeographicArea `json:"shapes,omitempty"`
}

// NewGeographicalArea instantiates a new GeographicalArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicalArea() *GeographicalArea {
	this := GeographicalArea{}
	return &this
}

// NewGeographicalAreaWithDefaults instantiates a new GeographicalArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicalAreaWithDefaults() *GeographicalArea {
	this := GeographicalArea{}
	return &this
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *GeographicalArea) GetCivicAddress() CivicAddress {
	if o == nil || IsNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalArea) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *GeographicalArea) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *GeographicalArea) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetShapes returns the Shapes field value if set, zero value otherwise.
func (o *GeographicalArea) GetShapes() GeographicArea {
	if o == nil || IsNil(o.Shapes) {
		var ret GeographicArea
		return ret
	}
	return *o.Shapes
}

// GetShapesOk returns a tuple with the Shapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalArea) GetShapesOk() (*GeographicArea, bool) {
	if o == nil || IsNil(o.Shapes) {
		return nil, false
	}
	return o.Shapes, true
}

// HasShapes returns a boolean if a field has been set.
func (o *GeographicalArea) HasShapes() bool {
	if o != nil && !IsNil(o.Shapes) {
		return true
	}

	return false
}

// SetShapes gets a reference to the given GeographicArea and assigns it to the Shapes field.
func (o *GeographicalArea) SetShapes(v GeographicArea) {
	o.Shapes = &v
}

func (o GeographicalArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeographicalArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if !IsNil(o.Shapes) {
		toSerialize["shapes"] = o.Shapes
	}
	return toSerialize, nil
}

type NullableGeographicalArea struct {
	value *GeographicalArea
	isSet bool
}

func (v NullableGeographicalArea) Get() *GeographicalArea {
	return v.value
}

func (v *NullableGeographicalArea) Set(val *GeographicalArea) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicalArea) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicalArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicalArea(val *GeographicalArea) *NullableGeographicalArea {
	return &NullableGeographicalArea{value: val, isSet: true}
}

func (v NullableGeographicalArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicalArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


