/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the ProfileDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileDoc{}

// ProfileDoc Represents the profile information associated with a VAL user ID or a VAL UE ID.
type ProfileDoc struct {
	// Profile information associated with the valUserId or valUeId.
	ProfileInformation string      `json:"profileInformation"`
	ValTgtUe           ValTargetUe `json:"valTgtUe"`
}

// NewProfileDoc instantiates a new ProfileDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileDoc(profileInformation string, valTgtUe ValTargetUe) *ProfileDoc {
	this := ProfileDoc{}
	this.ProfileInformation = profileInformation
	this.ValTgtUe = valTgtUe
	return &this
}

// NewProfileDocWithDefaults instantiates a new ProfileDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileDocWithDefaults() *ProfileDoc {
	this := ProfileDoc{}
	return &this
}

// GetProfileInformation returns the ProfileInformation field value
func (o *ProfileDoc) GetProfileInformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileInformation
}

// GetProfileInformationOk returns a tuple with the ProfileInformation field value
// and a boolean to check if the value has been set.
func (o *ProfileDoc) GetProfileInformationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileInformation, true
}

// SetProfileInformation sets field value
func (o *ProfileDoc) SetProfileInformation(v string) {
	o.ProfileInformation = v
}

// GetValTgtUe returns the ValTgtUe field value
func (o *ProfileDoc) GetValTgtUe() ValTargetUe {
	if o == nil {
		var ret ValTargetUe
		return ret
	}

	return o.ValTgtUe
}

// GetValTgtUeOk returns a tuple with the ValTgtUe field value
// and a boolean to check if the value has been set.
func (o *ProfileDoc) GetValTgtUeOk() (*ValTargetUe, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValTgtUe, true
}

// SetValTgtUe sets field value
func (o *ProfileDoc) SetValTgtUe(v ValTargetUe) {
	o.ValTgtUe = v
}

func (o ProfileDoc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profileInformation"] = o.ProfileInformation
	toSerialize["valTgtUe"] = o.ValTgtUe
	return toSerialize, nil
}

type NullableProfileDoc struct {
	value *ProfileDoc
	isSet bool
}

func (v NullableProfileDoc) Get() *ProfileDoc {
	return v.value
}

func (v *NullableProfileDoc) Set(val *ProfileDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileDoc(val *ProfileDoc) *NullableProfileDoc {
	return &NullableProfileDoc{value: val, isSet: true}
}

func (v NullableProfileDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
