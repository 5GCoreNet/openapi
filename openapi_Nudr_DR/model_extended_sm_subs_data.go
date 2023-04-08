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

// checks if the ExtendedSmSubsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendedSmSubsData{}

// ExtendedSmSubsData Contains identifiers of shared Session Management Subscription Data and optionally individual Session Management Subscription Data.
type ExtendedSmSubsData struct {
	SharedSmSubsDataIds []string `json:"sharedSmSubsDataIds"`
	IndividualSmSubsData []SessionManagementSubscriptionData `json:"individualSmSubsData,omitempty"`
}

// NewExtendedSmSubsData instantiates a new ExtendedSmSubsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedSmSubsData(sharedSmSubsDataIds []string) *ExtendedSmSubsData {
	this := ExtendedSmSubsData{}
	this.SharedSmSubsDataIds = sharedSmSubsDataIds
	return &this
}

// NewExtendedSmSubsDataWithDefaults instantiates a new ExtendedSmSubsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedSmSubsDataWithDefaults() *ExtendedSmSubsData {
	this := ExtendedSmSubsData{}
	return &this
}

// GetSharedSmSubsDataIds returns the SharedSmSubsDataIds field value
func (o *ExtendedSmSubsData) GetSharedSmSubsDataIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SharedSmSubsDataIds
}

// GetSharedSmSubsDataIdsOk returns a tuple with the SharedSmSubsDataIds field value
// and a boolean to check if the value has been set.
func (o *ExtendedSmSubsData) GetSharedSmSubsDataIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedSmSubsDataIds, true
}

// SetSharedSmSubsDataIds sets field value
func (o *ExtendedSmSubsData) SetSharedSmSubsDataIds(v []string) {
	o.SharedSmSubsDataIds = v
}

// GetIndividualSmSubsData returns the IndividualSmSubsData field value if set, zero value otherwise.
func (o *ExtendedSmSubsData) GetIndividualSmSubsData() []SessionManagementSubscriptionData {
	if o == nil || isNil(o.IndividualSmSubsData) {
		var ret []SessionManagementSubscriptionData
		return ret
	}
	return o.IndividualSmSubsData
}

// GetIndividualSmSubsDataOk returns a tuple with the IndividualSmSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedSmSubsData) GetIndividualSmSubsDataOk() ([]SessionManagementSubscriptionData, bool) {
	if o == nil || isNil(o.IndividualSmSubsData) {
		return nil, false
	}
	return o.IndividualSmSubsData, true
}

// HasIndividualSmSubsData returns a boolean if a field has been set.
func (o *ExtendedSmSubsData) HasIndividualSmSubsData() bool {
	if o != nil && !isNil(o.IndividualSmSubsData) {
		return true
	}

	return false
}

// SetIndividualSmSubsData gets a reference to the given []SessionManagementSubscriptionData and assigns it to the IndividualSmSubsData field.
func (o *ExtendedSmSubsData) SetIndividualSmSubsData(v []SessionManagementSubscriptionData) {
	o.IndividualSmSubsData = v
}

func (o ExtendedSmSubsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedSmSubsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sharedSmSubsDataIds"] = o.SharedSmSubsDataIds
	if !isNil(o.IndividualSmSubsData) {
		toSerialize["individualSmSubsData"] = o.IndividualSmSubsData
	}
	return toSerialize, nil
}

type NullableExtendedSmSubsData struct {
	value *ExtendedSmSubsData
	isSet bool
}

func (v NullableExtendedSmSubsData) Get() *ExtendedSmSubsData {
	return v.value
}

func (v *NullableExtendedSmSubsData) Set(val *ExtendedSmSubsData) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedSmSubsData) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedSmSubsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedSmSubsData(val *ExtendedSmSubsData) *NullableExtendedSmSubsData {
	return &NullableExtendedSmSubsData{value: val, isSet: true}
}

func (v NullableExtendedSmSubsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedSmSubsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


