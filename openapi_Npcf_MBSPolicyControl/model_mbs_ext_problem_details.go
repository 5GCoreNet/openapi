/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// checks if the MbsExtProblemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsExtProblemDetails{}

// MbsExtProblemDetails Identifies the MBS related extensions to the ProblemDetails data structure.
type MbsExtProblemDetails struct {
	ProblemDetails
}

// NewMbsExtProblemDetails instantiates a new MbsExtProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsExtProblemDetails() *MbsExtProblemDetails {
	this := MbsExtProblemDetails{}
	return &this
}

// NewMbsExtProblemDetailsWithDefaults instantiates a new MbsExtProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsExtProblemDetailsWithDefaults() *MbsExtProblemDetails {
	this := MbsExtProblemDetails{}
	return &this
}

func (o MbsExtProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsExtProblemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProblemDetails, errProblemDetails := json.Marshal(o.ProblemDetails)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	errProblemDetails = json.Unmarshal([]byte(serializedProblemDetails), &toSerialize)
	if errProblemDetails != nil {
		return map[string]interface{}{}, errProblemDetails
	}
	return toSerialize, nil
}

type NullableMbsExtProblemDetails struct {
	value *MbsExtProblemDetails
	isSet bool
}

func (v NullableMbsExtProblemDetails) Get() *MbsExtProblemDetails {
	return v.value
}

func (v *NullableMbsExtProblemDetails) Set(val *MbsExtProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsExtProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsExtProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsExtProblemDetails(val *MbsExtProblemDetails) *NullableMbsExtProblemDetails {
	return &NullableMbsExtProblemDetails{value: val, isSet: true}
}

func (v NullableMbsExtProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsExtProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
