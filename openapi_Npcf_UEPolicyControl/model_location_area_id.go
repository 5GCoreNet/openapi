/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// checks if the LocationAreaId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationAreaId{}

// LocationAreaId Contains a Location area identification as defined in 3GPP TS 23.003, clause 4.1.
type LocationAreaId struct {
	PlmnId PlmnId `json:"plmnId"`
	// Location Area Code.
	Lac string `json:"lac"`
}

// NewLocationAreaId instantiates a new LocationAreaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAreaId(plmnId PlmnId, lac string) *LocationAreaId {
	this := LocationAreaId{}
	this.PlmnId = plmnId
	this.Lac = lac
	return &this
}

// NewLocationAreaIdWithDefaults instantiates a new LocationAreaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAreaIdWithDefaults() *LocationAreaId {
	this := LocationAreaId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *LocationAreaId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *LocationAreaId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *LocationAreaId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *LocationAreaId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *LocationAreaId) GetLacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *LocationAreaId) SetLac(v string) {
	o.Lac = v
}

func (o LocationAreaId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationAreaId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["lac"] = o.Lac
	return toSerialize, nil
}

type NullableLocationAreaId struct {
	value *LocationAreaId
	isSet bool
}

func (v NullableLocationAreaId) Get() *LocationAreaId {
	return v.value
}

func (v *NullableLocationAreaId) Set(val *LocationAreaId) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAreaId) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAreaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAreaId(val *LocationAreaId) *NullableLocationAreaId {
	return &NullableLocationAreaId{value: val, isSet: true}
}

func (v NullableLocationAreaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAreaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
