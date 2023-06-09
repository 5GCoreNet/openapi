/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the ACCharacteristics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACCharacteristics{}

// ACCharacteristics Represents EAS dynamic information changes filter.
type ACCharacteristics struct {
	AcProf ACProfile `json:"acProf"`
}

// NewACCharacteristics instantiates a new ACCharacteristics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACCharacteristics(acProf ACProfile) *ACCharacteristics {
	this := ACCharacteristics{}
	this.AcProf = acProf
	return &this
}

// NewACCharacteristicsWithDefaults instantiates a new ACCharacteristics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACCharacteristicsWithDefaults() *ACCharacteristics {
	this := ACCharacteristics{}
	return &this
}

// GetAcProf returns the AcProf field value
func (o *ACCharacteristics) GetAcProf() ACProfile {
	if o == nil {
		var ret ACProfile
		return ret
	}

	return o.AcProf
}

// GetAcProfOk returns a tuple with the AcProf field value
// and a boolean to check if the value has been set.
func (o *ACCharacteristics) GetAcProfOk() (*ACProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcProf, true
}

// SetAcProf sets field value
func (o *ACCharacteristics) SetAcProf(v ACProfile) {
	o.AcProf = v
}

func (o ACCharacteristics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACCharacteristics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acProf"] = o.AcProf
	return toSerialize, nil
}

type NullableACCharacteristics struct {
	value *ACCharacteristics
	isSet bool
}

func (v NullableACCharacteristics) Get() *ACCharacteristics {
	return v.value
}

func (v *NullableACCharacteristics) Set(val *ACCharacteristics) {
	v.value = val
	v.isSet = true
}

func (v NullableACCharacteristics) IsSet() bool {
	return v.isSet
}

func (v *NullableACCharacteristics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACCharacteristics(val *ACCharacteristics) *NullableACCharacteristics {
	return &NullableACCharacteristics{value: val, isSet: true}
}

func (v NullableACCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACCharacteristics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
