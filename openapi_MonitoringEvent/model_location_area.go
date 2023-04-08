/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the LocationArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationArea{}

// LocationArea Represents a user location area.
type LocationArea struct {
	// Indicates a list of Cell Global Identities of the user which identifies the cell the UE is registered.
	CellIds []string `json:"cellIds,omitempty"`
	// Indicates a list of eNodeB identities in which the UE is currently located.
	EnodeBIds []string `json:"enodeBIds,omitempty"`
	// Identifies a list of Routing Area Identities of the user where the UE is located.
	RoutingAreaIds []string `json:"routingAreaIds,omitempty"`
	// Identifies a list of Tracking Area Identities of the user where the UE is located.
	TrackingAreaIds []string `json:"trackingAreaIds,omitempty"`
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`
}

// NewLocationArea instantiates a new LocationArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationArea() *LocationArea {
	this := LocationArea{}
	return &this
}

// NewLocationAreaWithDefaults instantiates a new LocationArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAreaWithDefaults() *LocationArea {
	this := LocationArea{}
	return &this
}

// GetCellIds returns the CellIds field value if set, zero value otherwise.
func (o *LocationArea) GetCellIds() []string {
	if o == nil || isNil(o.CellIds) {
		var ret []string
		return ret
	}
	return o.CellIds
}

// GetCellIdsOk returns a tuple with the CellIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetCellIdsOk() ([]string, bool) {
	if o == nil || isNil(o.CellIds) {
		return nil, false
	}
	return o.CellIds, true
}

// HasCellIds returns a boolean if a field has been set.
func (o *LocationArea) HasCellIds() bool {
	if o != nil && !isNil(o.CellIds) {
		return true
	}

	return false
}

// SetCellIds gets a reference to the given []string and assigns it to the CellIds field.
func (o *LocationArea) SetCellIds(v []string) {
	o.CellIds = v
}

// GetEnodeBIds returns the EnodeBIds field value if set, zero value otherwise.
func (o *LocationArea) GetEnodeBIds() []string {
	if o == nil || isNil(o.EnodeBIds) {
		var ret []string
		return ret
	}
	return o.EnodeBIds
}

// GetEnodeBIdsOk returns a tuple with the EnodeBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetEnodeBIdsOk() ([]string, bool) {
	if o == nil || isNil(o.EnodeBIds) {
		return nil, false
	}
	return o.EnodeBIds, true
}

// HasEnodeBIds returns a boolean if a field has been set.
func (o *LocationArea) HasEnodeBIds() bool {
	if o != nil && !isNil(o.EnodeBIds) {
		return true
	}

	return false
}

// SetEnodeBIds gets a reference to the given []string and assigns it to the EnodeBIds field.
func (o *LocationArea) SetEnodeBIds(v []string) {
	o.EnodeBIds = v
}

// GetRoutingAreaIds returns the RoutingAreaIds field value if set, zero value otherwise.
func (o *LocationArea) GetRoutingAreaIds() []string {
	if o == nil || isNil(o.RoutingAreaIds) {
		var ret []string
		return ret
	}
	return o.RoutingAreaIds
}

// GetRoutingAreaIdsOk returns a tuple with the RoutingAreaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetRoutingAreaIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RoutingAreaIds) {
		return nil, false
	}
	return o.RoutingAreaIds, true
}

// HasRoutingAreaIds returns a boolean if a field has been set.
func (o *LocationArea) HasRoutingAreaIds() bool {
	if o != nil && !isNil(o.RoutingAreaIds) {
		return true
	}

	return false
}

// SetRoutingAreaIds gets a reference to the given []string and assigns it to the RoutingAreaIds field.
func (o *LocationArea) SetRoutingAreaIds(v []string) {
	o.RoutingAreaIds = v
}

// GetTrackingAreaIds returns the TrackingAreaIds field value if set, zero value otherwise.
func (o *LocationArea) GetTrackingAreaIds() []string {
	if o == nil || isNil(o.TrackingAreaIds) {
		var ret []string
		return ret
	}
	return o.TrackingAreaIds
}

// GetTrackingAreaIdsOk returns a tuple with the TrackingAreaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetTrackingAreaIdsOk() ([]string, bool) {
	if o == nil || isNil(o.TrackingAreaIds) {
		return nil, false
	}
	return o.TrackingAreaIds, true
}

// HasTrackingAreaIds returns a boolean if a field has been set.
func (o *LocationArea) HasTrackingAreaIds() bool {
	if o != nil && !isNil(o.TrackingAreaIds) {
		return true
	}

	return false
}

// SetTrackingAreaIds gets a reference to the given []string and assigns it to the TrackingAreaIds field.
func (o *LocationArea) SetTrackingAreaIds(v []string) {
	o.TrackingAreaIds = v
}

// GetGeographicAreas returns the GeographicAreas field value if set, zero value otherwise.
func (o *LocationArea) GetGeographicAreas() []GeographicArea {
	if o == nil || isNil(o.GeographicAreas) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreas
}

// GetGeographicAreasOk returns a tuple with the GeographicAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetGeographicAreasOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeographicAreas) {
		return nil, false
	}
	return o.GeographicAreas, true
}

// HasGeographicAreas returns a boolean if a field has been set.
func (o *LocationArea) HasGeographicAreas() bool {
	if o != nil && !isNil(o.GeographicAreas) {
		return true
	}

	return false
}

// SetGeographicAreas gets a reference to the given []GeographicArea and assigns it to the GeographicAreas field.
func (o *LocationArea) SetGeographicAreas(v []GeographicArea) {
	o.GeographicAreas = v
}

// GetCivicAddresses returns the CivicAddresses field value if set, zero value otherwise.
func (o *LocationArea) GetCivicAddresses() []CivicAddress {
	if o == nil || isNil(o.CivicAddresses) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddresses
}

// GetCivicAddressesOk returns a tuple with the CivicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetCivicAddressesOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddresses) {
		return nil, false
	}
	return o.CivicAddresses, true
}

// HasCivicAddresses returns a boolean if a field has been set.
func (o *LocationArea) HasCivicAddresses() bool {
	if o != nil && !isNil(o.CivicAddresses) {
		return true
	}

	return false
}

// SetCivicAddresses gets a reference to the given []CivicAddress and assigns it to the CivicAddresses field.
func (o *LocationArea) SetCivicAddresses(v []CivicAddress) {
	o.CivicAddresses = v
}

func (o LocationArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CellIds) {
		toSerialize["cellIds"] = o.CellIds
	}
	if !isNil(o.EnodeBIds) {
		toSerialize["enodeBIds"] = o.EnodeBIds
	}
	if !isNil(o.RoutingAreaIds) {
		toSerialize["routingAreaIds"] = o.RoutingAreaIds
	}
	if !isNil(o.TrackingAreaIds) {
		toSerialize["trackingAreaIds"] = o.TrackingAreaIds
	}
	if !isNil(o.GeographicAreas) {
		toSerialize["geographicAreas"] = o.GeographicAreas
	}
	if !isNil(o.CivicAddresses) {
		toSerialize["civicAddresses"] = o.CivicAddresses
	}
	return toSerialize, nil
}

type NullableLocationArea struct {
	value *LocationArea
	isSet bool
}

func (v NullableLocationArea) Get() *LocationArea {
	return v.value
}

func (v *NullableLocationArea) Set(val *LocationArea) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationArea) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationArea(val *LocationArea) *NullableLocationArea {
	return &NullableLocationArea{value: val, isSet: true}
}

func (v NullableLocationArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


