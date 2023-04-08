/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the ServiceAreaRestriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAreaRestriction{}

// ServiceAreaRestriction Provides information about allowed or not allowed areas.
type ServiceAreaRestriction struct {
	RestrictionType *RestrictionType `json:"restrictionType,omitempty"`
	Areas []Area `json:"areas,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAs *int32 `json:"maxNumOfTAs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTAsForNotAllowedAreas *int32 `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}

// NewServiceAreaRestriction instantiates a new ServiceAreaRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAreaRestriction() *ServiceAreaRestriction {
	this := ServiceAreaRestriction{}
	return &this
}

// NewServiceAreaRestrictionWithDefaults instantiates a new ServiceAreaRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaRestrictionWithDefaults() *ServiceAreaRestriction {
	this := ServiceAreaRestriction{}
	return &this
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *ServiceAreaRestriction) GetRestrictionType() RestrictionType {
	if o == nil || isNil(o.RestrictionType) {
		var ret RestrictionType
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction) GetRestrictionTypeOk() (*RestrictionType, bool) {
	if o == nil || isNil(o.RestrictionType) {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *ServiceAreaRestriction) HasRestrictionType() bool {
	if o != nil && !isNil(o.RestrictionType) {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given RestrictionType and assigns it to the RestrictionType field.
func (o *ServiceAreaRestriction) SetRestrictionType(v RestrictionType) {
	o.RestrictionType = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *ServiceAreaRestriction) GetAreas() []Area {
	if o == nil || isNil(o.Areas) {
		var ret []Area
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction) GetAreasOk() ([]Area, bool) {
	if o == nil || isNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *ServiceAreaRestriction) HasAreas() bool {
	if o != nil && !isNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []Area and assigns it to the Areas field.
func (o *ServiceAreaRestriction) SetAreas(v []Area) {
	o.Areas = v
}

// GetMaxNumOfTAs returns the MaxNumOfTAs field value if set, zero value otherwise.
func (o *ServiceAreaRestriction) GetMaxNumOfTAs() int32 {
	if o == nil || isNil(o.MaxNumOfTAs) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTAs
}

// GetMaxNumOfTAsOk returns a tuple with the MaxNumOfTAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction) GetMaxNumOfTAsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumOfTAs) {
		return nil, false
	}
	return o.MaxNumOfTAs, true
}

// HasMaxNumOfTAs returns a boolean if a field has been set.
func (o *ServiceAreaRestriction) HasMaxNumOfTAs() bool {
	if o != nil && !isNil(o.MaxNumOfTAs) {
		return true
	}

	return false
}

// SetMaxNumOfTAs gets a reference to the given int32 and assigns it to the MaxNumOfTAs field.
func (o *ServiceAreaRestriction) SetMaxNumOfTAs(v int32) {
	o.MaxNumOfTAs = &v
}

// GetMaxNumOfTAsForNotAllowedAreas returns the MaxNumOfTAsForNotAllowedAreas field value if set, zero value otherwise.
func (o *ServiceAreaRestriction) GetMaxNumOfTAsForNotAllowedAreas() int32 {
	if o == nil || isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTAsForNotAllowedAreas
}

// GetMaxNumOfTAsForNotAllowedAreasOk returns a tuple with the MaxNumOfTAsForNotAllowedAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAreaRestriction) GetMaxNumOfTAsForNotAllowedAreasOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		return nil, false
	}
	return o.MaxNumOfTAsForNotAllowedAreas, true
}

// HasMaxNumOfTAsForNotAllowedAreas returns a boolean if a field has been set.
func (o *ServiceAreaRestriction) HasMaxNumOfTAsForNotAllowedAreas() bool {
	if o != nil && !isNil(o.MaxNumOfTAsForNotAllowedAreas) {
		return true
	}

	return false
}

// SetMaxNumOfTAsForNotAllowedAreas gets a reference to the given int32 and assigns it to the MaxNumOfTAsForNotAllowedAreas field.
func (o *ServiceAreaRestriction) SetMaxNumOfTAsForNotAllowedAreas(v int32) {
	o.MaxNumOfTAsForNotAllowedAreas = &v
}

func (o ServiceAreaRestriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAreaRestriction) ToMap() (map[string]interface{}, error) {
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

type NullableServiceAreaRestriction struct {
	value *ServiceAreaRestriction
	isSet bool
}

func (v NullableServiceAreaRestriction) Get() *ServiceAreaRestriction {
	return v.value
}

func (v *NullableServiceAreaRestriction) Set(val *ServiceAreaRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAreaRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAreaRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAreaRestriction(val *ServiceAreaRestriction) *NullableServiceAreaRestriction {
	return &NullableServiceAreaRestriction{value: val, isSet: true}
}

func (v NullableServiceAreaRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAreaRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


