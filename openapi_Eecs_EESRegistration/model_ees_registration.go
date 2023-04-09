/*
ECS EES Registration_API

API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_EESRegistration

import (
	"encoding/json"
	"time"
)

// checks if the EESRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EESRegistration{}

// EESRegistration Represents an EES registration information.
type EESRegistration struct {
	EesProf EESProfile `json:"eesProf"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewEESRegistration instantiates a new EESRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEESRegistration(eesProf EESProfile) *EESRegistration {
	this := EESRegistration{}
	this.EesProf = eesProf
	return &this
}

// NewEESRegistrationWithDefaults instantiates a new EESRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEESRegistrationWithDefaults() *EESRegistration {
	this := EESRegistration{}
	return &this
}

// GetEesProf returns the EesProf field value
func (o *EESRegistration) GetEesProf() EESProfile {
	if o == nil {
		var ret EESProfile
		return ret
	}

	return o.EesProf
}

// GetEesProfOk returns a tuple with the EesProf field value
// and a boolean to check if the value has been set.
func (o *EESRegistration) GetEesProfOk() (*EESProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EesProf, true
}

// SetEesProf sets field value
func (o *EESRegistration) SetEesProf(v EESProfile) {
	o.EesProf = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *EESRegistration) GetExpTime() time.Time {
	if o == nil || IsNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESRegistration) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *EESRegistration) HasExpTime() bool {
	if o != nil && !IsNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *EESRegistration) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *EESRegistration) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESRegistration) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *EESRegistration) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *EESRegistration) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o EESRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EESRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eesProf"] = o.EesProf
	if !IsNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableEESRegistration struct {
	value *EESRegistration
	isSet bool
}

func (v NullableEESRegistration) Get() *EESRegistration {
	return v.value
}

func (v *NullableEESRegistration) Set(val *EESRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableEESRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableEESRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEESRegistration(val *EESRegistration) *NullableEESRegistration {
	return &NullableEESRegistration{value: val, isSet: true}
}

func (v NullableEESRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEESRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
