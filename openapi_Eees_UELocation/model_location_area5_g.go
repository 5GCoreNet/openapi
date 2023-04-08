/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
)

// checks if the LocationArea5G type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationArea5G{}

// LocationArea5G Represents a user location area when the UE is attached to 5G.
type LocationArea5G struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`
	NwAreaInfo *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
}

// NewLocationArea5G instantiates a new LocationArea5G object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationArea5G() *LocationArea5G {
	this := LocationArea5G{}
	return &this
}

// NewLocationArea5GWithDefaults instantiates a new LocationArea5G object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationArea5GWithDefaults() *LocationArea5G {
	this := LocationArea5G{}
	return &this
}

// GetGeographicAreas returns the GeographicAreas field value if set, zero value otherwise.
func (o *LocationArea5G) GetGeographicAreas() []GeographicArea {
	if o == nil || isNil(o.GeographicAreas) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreas
}

// GetGeographicAreasOk returns a tuple with the GeographicAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea5G) GetGeographicAreasOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeographicAreas) {
		return nil, false
	}
	return o.GeographicAreas, true
}

// HasGeographicAreas returns a boolean if a field has been set.
func (o *LocationArea5G) HasGeographicAreas() bool {
	if o != nil && !isNil(o.GeographicAreas) {
		return true
	}

	return false
}

// SetGeographicAreas gets a reference to the given []GeographicArea and assigns it to the GeographicAreas field.
func (o *LocationArea5G) SetGeographicAreas(v []GeographicArea) {
	o.GeographicAreas = v
}

// GetCivicAddresses returns the CivicAddresses field value if set, zero value otherwise.
func (o *LocationArea5G) GetCivicAddresses() []CivicAddress {
	if o == nil || isNil(o.CivicAddresses) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddresses
}

// GetCivicAddressesOk returns a tuple with the CivicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea5G) GetCivicAddressesOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddresses) {
		return nil, false
	}
	return o.CivicAddresses, true
}

// HasCivicAddresses returns a boolean if a field has been set.
func (o *LocationArea5G) HasCivicAddresses() bool {
	if o != nil && !isNil(o.CivicAddresses) {
		return true
	}

	return false
}

// SetCivicAddresses gets a reference to the given []CivicAddress and assigns it to the CivicAddresses field.
func (o *LocationArea5G) SetCivicAddresses(v []CivicAddress) {
	o.CivicAddresses = v
}

// GetNwAreaInfo returns the NwAreaInfo field value if set, zero value otherwise.
func (o *LocationArea5G) GetNwAreaInfo() NetworkAreaInfo {
	if o == nil || isNil(o.NwAreaInfo) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NwAreaInfo
}

// GetNwAreaInfoOk returns a tuple with the NwAreaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea5G) GetNwAreaInfoOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.NwAreaInfo) {
		return nil, false
	}
	return o.NwAreaInfo, true
}

// HasNwAreaInfo returns a boolean if a field has been set.
func (o *LocationArea5G) HasNwAreaInfo() bool {
	if o != nil && !isNil(o.NwAreaInfo) {
		return true
	}

	return false
}

// SetNwAreaInfo gets a reference to the given NetworkAreaInfo and assigns it to the NwAreaInfo field.
func (o *LocationArea5G) SetNwAreaInfo(v NetworkAreaInfo) {
	o.NwAreaInfo = &v
}

func (o LocationArea5G) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationArea5G) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeographicAreas) {
		toSerialize["geographicAreas"] = o.GeographicAreas
	}
	if !isNil(o.CivicAddresses) {
		toSerialize["civicAddresses"] = o.CivicAddresses
	}
	if !isNil(o.NwAreaInfo) {
		toSerialize["nwAreaInfo"] = o.NwAreaInfo
	}
	return toSerialize, nil
}

type NullableLocationArea5G struct {
	value *LocationArea5G
	isSet bool
}

func (v NullableLocationArea5G) Get() *LocationArea5G {
	return v.value
}

func (v *NullableLocationArea5G) Set(val *LocationArea5G) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationArea5G) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationArea5G) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationArea5G(val *LocationArea5G) *NullableLocationArea5G {
	return &NullableLocationArea5G{value: val, isSet: true}
}

func (v NullableLocationArea5G) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationArea5G) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


