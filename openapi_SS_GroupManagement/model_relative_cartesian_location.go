/*
SS_GroupManagement

API for SEAL Group management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_GroupManagement

import (
	"encoding/json"
)

// checks if the RelativeCartesianLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelativeCartesianLocation{}

// RelativeCartesianLocation Relative Cartesian Location
type RelativeCartesianLocation struct {
	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x"`
	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y"`
	// string with format 'float' as defined in OpenAPI.
	Z *float32 `json:"z,omitempty"`
}

// NewRelativeCartesianLocation instantiates a new RelativeCartesianLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelativeCartesianLocation(x float32, y float32) *RelativeCartesianLocation {
	this := RelativeCartesianLocation{}
	this.X = x
	this.Y = y
	return &this
}

// NewRelativeCartesianLocationWithDefaults instantiates a new RelativeCartesianLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelativeCartesianLocationWithDefaults() *RelativeCartesianLocation {
	this := RelativeCartesianLocation{}
	return &this
}

// GetX returns the X field value
func (o *RelativeCartesianLocation) GetX() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetXOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *RelativeCartesianLocation) SetX(v float32) {
	o.X = v
}

// GetY returns the Y field value
func (o *RelativeCartesianLocation) GetY() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetYOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *RelativeCartesianLocation) SetY(v float32) {
	o.Y = v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *RelativeCartesianLocation) GetZ() float32 {
	if o == nil || IsNil(o.Z) {
		var ret float32
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeCartesianLocation) GetZOk() (*float32, bool) {
	if o == nil || IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *RelativeCartesianLocation) HasZ() bool {
	if o != nil && !IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given float32 and assigns it to the Z field.
func (o *RelativeCartesianLocation) SetZ(v float32) {
	o.Z = &v
}

func (o RelativeCartesianLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelativeCartesianLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	if !IsNil(o.Z) {
		toSerialize["z"] = o.Z
	}
	return toSerialize, nil
}

type NullableRelativeCartesianLocation struct {
	value *RelativeCartesianLocation
	isSet bool
}

func (v NullableRelativeCartesianLocation) Get() *RelativeCartesianLocation {
	return v.value
}

func (v *NullableRelativeCartesianLocation) Set(val *RelativeCartesianLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRelativeCartesianLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRelativeCartesianLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelativeCartesianLocation(val *RelativeCartesianLocation) *NullableRelativeCartesianLocation {
	return &NullableRelativeCartesianLocation{value: val, isSet: true}
}

func (v NullableRelativeCartesianLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelativeCartesianLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


