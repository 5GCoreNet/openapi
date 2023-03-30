/*
5GMS Common Data Types

5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the TypedLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypedLocation{}

// TypedLocation struct for TypedLocation
type TypedLocation struct {
	LocationIdentifierType CellIdentifierType `json:"locationIdentifierType"`
	Location string `json:"location"`
}

// NewTypedLocation instantiates a new TypedLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypedLocation(locationIdentifierType CellIdentifierType, location string) *TypedLocation {
	this := TypedLocation{}
	this.LocationIdentifierType = locationIdentifierType
	this.Location = location
	return &this
}

// NewTypedLocationWithDefaults instantiates a new TypedLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypedLocationWithDefaults() *TypedLocation {
	this := TypedLocation{}
	return &this
}

// GetLocationIdentifierType returns the LocationIdentifierType field value
func (o *TypedLocation) GetLocationIdentifierType() CellIdentifierType {
	if o == nil {
		var ret CellIdentifierType
		return ret
	}

	return o.LocationIdentifierType
}

// GetLocationIdentifierTypeOk returns a tuple with the LocationIdentifierType field value
// and a boolean to check if the value has been set.
func (o *TypedLocation) GetLocationIdentifierTypeOk() (*CellIdentifierType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationIdentifierType, true
}

// SetLocationIdentifierType sets field value
func (o *TypedLocation) SetLocationIdentifierType(v CellIdentifierType) {
	o.LocationIdentifierType = v
}

// GetLocation returns the Location field value
func (o *TypedLocation) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *TypedLocation) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *TypedLocation) SetLocation(v string) {
	o.Location = v
}

func (o TypedLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypedLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locationIdentifierType"] = o.LocationIdentifierType
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableTypedLocation struct {
	value *TypedLocation
	isSet bool
}

func (v NullableTypedLocation) Get() *TypedLocation {
	return v.value
}

func (v *NullableTypedLocation) Set(val *TypedLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTypedLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTypedLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypedLocation(val *TypedLocation) *NullableTypedLocation {
	return &NullableTypedLocation{value: val, isSet: true}
}

func (v NullableTypedLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypedLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

