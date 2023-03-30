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

// checks if the SteeringInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SteeringInfo{}

// SteeringInfo Contains a combination of one PLMN identity and zero or more access technologies.
type SteeringInfo struct {
	PlmnId PlmnId `json:"plmnId"`
	AccessTechList []AccessTech `json:"accessTechList,omitempty"`
}

// NewSteeringInfo instantiates a new SteeringInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSteeringInfo(plmnId PlmnId) *SteeringInfo {
	this := SteeringInfo{}
	this.PlmnId = plmnId
	return &this
}

// NewSteeringInfoWithDefaults instantiates a new SteeringInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSteeringInfoWithDefaults() *SteeringInfo {
	this := SteeringInfo{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *SteeringInfo) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *SteeringInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *SteeringInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetAccessTechList returns the AccessTechList field value if set, zero value otherwise.
func (o *SteeringInfo) GetAccessTechList() []AccessTech {
	if o == nil || IsNil(o.AccessTechList) {
		var ret []AccessTech
		return ret
	}
	return o.AccessTechList
}

// GetAccessTechListOk returns a tuple with the AccessTechList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SteeringInfo) GetAccessTechListOk() ([]AccessTech, bool) {
	if o == nil || IsNil(o.AccessTechList) {
		return nil, false
	}
	return o.AccessTechList, true
}

// HasAccessTechList returns a boolean if a field has been set.
func (o *SteeringInfo) HasAccessTechList() bool {
	if o != nil && !IsNil(o.AccessTechList) {
		return true
	}

	return false
}

// SetAccessTechList gets a reference to the given []AccessTech and assigns it to the AccessTechList field.
func (o *SteeringInfo) SetAccessTechList(v []AccessTech) {
	o.AccessTechList = v
}

func (o SteeringInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SteeringInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	if !IsNil(o.AccessTechList) {
		toSerialize["accessTechList"] = o.AccessTechList
	}
	return toSerialize, nil
}

type NullableSteeringInfo struct {
	value *SteeringInfo
	isSet bool
}

func (v NullableSteeringInfo) Get() *SteeringInfo {
	return v.value
}

func (v *NullableSteeringInfo) Set(val *SteeringInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSteeringInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSteeringInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteeringInfo(val *SteeringInfo) *NullableSteeringInfo {
	return &NullableSteeringInfo{value: val, isSet: true}
}

func (v NullableSteeringInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteeringInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


