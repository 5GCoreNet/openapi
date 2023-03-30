/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the RoamingRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoamingRestrictions{}

// RoamingRestrictions Indicates if access is allowed to a given serving network, e.g. a PLMN (MCC, MNC) or an  SNPN (MCC, MNC, NID). 
type RoamingRestrictions struct {
	AccessAllowed *bool `json:"accessAllowed,omitempty"`
}

// NewRoamingRestrictions instantiates a new RoamingRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoamingRestrictions() *RoamingRestrictions {
	this := RoamingRestrictions{}
	return &this
}

// NewRoamingRestrictionsWithDefaults instantiates a new RoamingRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoamingRestrictionsWithDefaults() *RoamingRestrictions {
	this := RoamingRestrictions{}
	return &this
}

// GetAccessAllowed returns the AccessAllowed field value if set, zero value otherwise.
func (o *RoamingRestrictions) GetAccessAllowed() bool {
	if o == nil || IsNil(o.AccessAllowed) {
		var ret bool
		return ret
	}
	return *o.AccessAllowed
}

// GetAccessAllowedOk returns a tuple with the AccessAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoamingRestrictions) GetAccessAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessAllowed) {
		return nil, false
	}
	return o.AccessAllowed, true
}

// HasAccessAllowed returns a boolean if a field has been set.
func (o *RoamingRestrictions) HasAccessAllowed() bool {
	if o != nil && !IsNil(o.AccessAllowed) {
		return true
	}

	return false
}

// SetAccessAllowed gets a reference to the given bool and assigns it to the AccessAllowed field.
func (o *RoamingRestrictions) SetAccessAllowed(v bool) {
	o.AccessAllowed = &v
}

func (o RoamingRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoamingRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessAllowed) {
		toSerialize["accessAllowed"] = o.AccessAllowed
	}
	return toSerialize, nil
}

type NullableRoamingRestrictions struct {
	value *RoamingRestrictions
	isSet bool
}

func (v NullableRoamingRestrictions) Get() *RoamingRestrictions {
	return v.value
}

func (v *NullableRoamingRestrictions) Set(val *RoamingRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingRestrictions(val *RoamingRestrictions) *NullableRoamingRestrictions {
	return &NullableRoamingRestrictions{value: val, isSet: true}
}

func (v NullableRoamingRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


