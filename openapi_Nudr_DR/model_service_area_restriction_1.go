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

// checks if the ServiceAreaRestriction1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAreaRestriction1{}

// ServiceAreaRestriction1 Provides information about allowed or not allowed areas.
type ServiceAreaRestriction1 struct {
	RestrictionType *RestrictionType `json:"restrictionType,omitempty"`
	Areas []Area1 `json:"areas,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAs *int32 `json:"maxNumOfTAs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAsForNotAllowedAreas *int32 `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}

// NewServiceAreaRestriction1 instantiates a new ServiceAreaRestriction1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAreaRestriction1() *ServiceAreaRestriction1 {
	this := ServiceAreaRestriction1{}
	return &this
}

// NewServiceAreaRestriction1WithDefaults instantiates a new ServiceAreaRestriction1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaRestriction1WithDefaults() *ServiceAreaRestriction1 {
	this := ServiceAreaRestriction1{}
	return &this
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *ServiceAreaRestriction1) GetRestrictionType() RestrictionType {
	if o == nil || isNil(o.RestrictionType) {
		var ret RestrictionType
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction1) GetRestrictionTypeOk() (*RestrictionType, bool) {
	if o == nil || isNil(o.RestrictionType) {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *ServiceAreaRestriction1) HasRestrictionType() bool {
	if o != nil && !isNil(o.RestrictionType) {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given RestrictionType and assigns it to the RestrictionType field.
func (o *ServiceAreaRestriction1) SetRestrictionType(v RestrictionType) {
	o.RestrictionType = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *ServiceAreaRestriction1) GetAreas() []Area1 {
	if o == nil || isNil(o.Areas) {
		var ret []Area1
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction1) GetAreasOk() ([]Area1, bool) {
	if o == nil || isNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *ServiceAreaRestriction1) HasAreas() bool {
	if o != nil && !isNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []Area1 and assigns it to the Areas field.
func (o *ServiceAreaRestriction1) SetAreas(v []Area1) {
	o.Areas = v
}

// GetMaxNumOfTAs returns the MaxNumOfTAs field value if set, zero value otherwise.
func (o *ServiceAreaRestriction1) GetMaxNumOfTAs() int32 {
	if o == nil || isNil(o.MaxNumOfTAs) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTAs
}

// GetMaxNumOfTAsOk returns a tuple with the MaxNumOfTAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction1) GetMaxNumOfTAsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumOfTAs) {
		return nil, false
	}
	return o.MaxNumOfTAs, true
}

// HasMaxNumOfTAs returns a boolean if a field has been set.
func (o *ServiceAreaRestriction1) HasMaxNumOfTAs() bool {
	if o != nil && !isNil(o.MaxNumOfTAs) {
		return true
	}

	return false
}

// SetMaxNumOfTAs gets a reference to the given int32 and assigns it to the MaxNumOfTAs field.
func (o *ServiceAreaRestriction1) SetMaxNumOfTAs(v int32) {
	o.MaxNumOfTAs = &v
}

// GetMaxNumOfTAsForNotAllowedAreas returns the MaxNumOfTAsForNotAllowedAreas field value if set, zero value otherwise.
func (o *ServiceAreaRestriction1) GetMaxNumOfTAsForNotAllowedAreas() int32 {
	if o == nil || isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTAsForNotAllowedAreas
}

// GetMaxNumOfTAsForNotAllowedAreasOk returns a tuple with the MaxNumOfTAsForNotAllowedAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction1) GetMaxNumOfTAsForNotAllowedAreasOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		return nil, false
	}
	return o.MaxNumOfTAsForNotAllowedAreas, true
}

// HasMaxNumOfTAsForNotAllowedAreas returns a boolean if a field has been set.
func (o *ServiceAreaRestriction1) HasMaxNumOfTAsForNotAllowedAreas() bool {
	if o != nil && !isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		return true
	}

	return false
}

// SetMaxNumOfTAsForNotAllowedAreas gets a reference to the given int32 and assigns it to the MaxNumOfTAsForNotAllowedAreas field.
func (o *ServiceAreaRestriction1) SetMaxNumOfTAsForNotAllowedAreas(v int32) {
	o.MaxNumOfTAsForNotAllowedAreas = &v
}

func (o ServiceAreaRestriction1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAreaRestriction1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RestrictionType) {
		toSerialize["restrictionType"] = o.RestrictionType
	}
	if !isNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	if !isNil(o.MaxNumOfTAs) {
		toSerialize["maxNumOfTAs"] = o.MaxNumOfTAs
	}
	if !isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		toSerialize["maxNumOfTAsForNotAllowedAreas"] = o.MaxNumOfTAsForNotAllowedAreas
	}
	return toSerialize, nil
}

type NullableServiceAreaRestriction1 struct {
	value *ServiceAreaRestriction1
	isSet bool
}

func (v NullableServiceAreaRestriction1) Get() *ServiceAreaRestriction1 {
	return v.value
}

func (v *NullableServiceAreaRestriction1) Set(val *ServiceAreaRestriction1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAreaRestriction1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAreaRestriction1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAreaRestriction1(val *ServiceAreaRestriction1) *NullableServiceAreaRestriction1 {
	return &NullableServiceAreaRestriction1{value: val, isSet: true}
}

func (v NullableServiceAreaRestriction1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAreaRestriction1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


