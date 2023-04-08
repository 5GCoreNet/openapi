/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
)

// checks if the RoamingInfoUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoamingInfoUpdate{}

// RoamingInfoUpdate Contains the Roaming Information Update
type RoamingInfoUpdate struct {
	Roaming *bool `json:"roaming,omitempty"`
	ServingPlmn PlmnId `json:"servingPlmn"`
}

// NewRoamingInfoUpdate instantiates a new RoamingInfoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoamingInfoUpdate(servingPlmn PlmnId) *RoamingInfoUpdate {
	this := RoamingInfoUpdate{}
	this.ServingPlmn = servingPlmn
	return &this
}

// NewRoamingInfoUpdateWithDefaults instantiates a new RoamingInfoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoamingInfoUpdateWithDefaults() *RoamingInfoUpdate {
	this := RoamingInfoUpdate{}
	return &this
}

// GetRoaming returns the Roaming field value if set, zero value otherwise.
func (o *RoamingInfoUpdate) GetRoaming() bool {
	if o == nil || isNil(o.Roaming) {
		var ret bool
		return ret
	}
	return *o.Roaming
}

// GetRoamingOk returns a tuple with the Roaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoamingInfoUpdate) GetRoamingOk() (*bool, bool) {
	if o == nil || isNil(o.Roaming) {
		return nil, false
	}
	return o.Roaming, true
}

// HasRoaming returns a boolean if a field has been set.
func (o *RoamingInfoUpdate) HasRoaming() bool {
	if o != nil && !isNil(o.Roaming) {
		return true
	}

	return false
}

// SetRoaming gets a reference to the given bool and assigns it to the Roaming field.
func (o *RoamingInfoUpdate) SetRoaming(v bool) {
	o.Roaming = &v
}

// GetServingPlmn returns the ServingPlmn field value
func (o *RoamingInfoUpdate) GetServingPlmn() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.ServingPlmn
}

// GetServingPlmnOk returns a tuple with the ServingPlmn field value
// and a boolean to check if the value has been set.
func (o *RoamingInfoUpdate) GetServingPlmnOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingPlmn, true
}

// SetServingPlmn sets field value
func (o *RoamingInfoUpdate) SetServingPlmn(v PlmnId) {
	o.ServingPlmn = v
}

func (o RoamingInfoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoamingInfoUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Roaming) {
		toSerialize["roaming"] = o.Roaming
	}
	toSerialize["servingPlmn"] = o.ServingPlmn
	return toSerialize, nil
}

type NullableRoamingInfoUpdate struct {
	value *RoamingInfoUpdate
	isSet bool
}

func (v NullableRoamingInfoUpdate) Get() *RoamingInfoUpdate {
	return v.value
}

func (v *NullableRoamingInfoUpdate) Set(val *RoamingInfoUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingInfoUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingInfoUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingInfoUpdate(val *RoamingInfoUpdate) *NullableRoamingInfoUpdate {
	return &NullableRoamingInfoUpdate{value: val, isSet: true}
}

func (v NullableRoamingInfoUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingInfoUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


