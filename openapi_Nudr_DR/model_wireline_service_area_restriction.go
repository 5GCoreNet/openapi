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

// checks if the WirelineServiceAreaRestriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelineServiceAreaRestriction{}

// WirelineServiceAreaRestriction The \"restrictionType\" attribute and the \"areas\" attribute shall be either both present or absent.  The empty array of areas is used when service is allowed/restricted nowhere.
type WirelineServiceAreaRestriction struct {
	RestrictionType *RestrictionType `json:"restrictionType,omitempty"`
	Areas           []WirelineArea   `json:"areas,omitempty"`
}

// NewWirelineServiceAreaRestriction instantiates a new WirelineServiceAreaRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelineServiceAreaRestriction() *WirelineServiceAreaRestriction {
	this := WirelineServiceAreaRestriction{}
	return &this
}

// NewWirelineServiceAreaRestrictionWithDefaults instantiates a new WirelineServiceAreaRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelineServiceAreaRestrictionWithDefaults() *WirelineServiceAreaRestriction {
	this := WirelineServiceAreaRestriction{}
	return &this
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *WirelineServiceAreaRestriction) GetRestrictionType() RestrictionType {
	if o == nil || IsNil(o.RestrictionType) {
		var ret RestrictionType
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineServiceAreaRestriction) GetRestrictionTypeOk() (*RestrictionType, bool) {
	if o == nil || IsNil(o.RestrictionType) {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *WirelineServiceAreaRestriction) HasRestrictionType() bool {
	if o != nil && !IsNil(o.RestrictionType) {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given RestrictionType and assigns it to the RestrictionType field.
func (o *WirelineServiceAreaRestriction) SetRestrictionType(v RestrictionType) {
	o.RestrictionType = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *WirelineServiceAreaRestriction) GetAreas() []WirelineArea {
	if o == nil || IsNil(o.Areas) {
		var ret []WirelineArea
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelineServiceAreaRestriction) GetAreasOk() ([]WirelineArea, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *WirelineServiceAreaRestriction) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []WirelineArea and assigns it to the Areas field.
func (o *WirelineServiceAreaRestriction) SetAreas(v []WirelineArea) {
	o.Areas = v
}

func (o WirelineServiceAreaRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelineServiceAreaRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestrictionType) {
		toSerialize["restrictionType"] = o.RestrictionType
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableWirelineServiceAreaRestriction struct {
	value *WirelineServiceAreaRestriction
	isSet bool
}

func (v NullableWirelineServiceAreaRestriction) Get() *WirelineServiceAreaRestriction {
	return v.value
}

func (v *NullableWirelineServiceAreaRestriction) Set(val *WirelineServiceAreaRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelineServiceAreaRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelineServiceAreaRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelineServiceAreaRestriction(val *WirelineServiceAreaRestriction) *NullableWirelineServiceAreaRestriction {
	return &NullableWirelineServiceAreaRestriction{value: val, isSet: true}
}

func (v NullableWirelineServiceAreaRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelineServiceAreaRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
