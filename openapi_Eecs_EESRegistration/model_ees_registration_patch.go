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

// checks if the EESRegistrationPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EESRegistrationPatch{}

// EESRegistrationPatch Represents partial update request of individual EES registration information.
type EESRegistrationPatch struct {
	EesProf *EESProfile `json:"eesProf,omitempty"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	ExpTime NullableTime `json:"expTime,omitempty"`
}

// NewEESRegistrationPatch instantiates a new EESRegistrationPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEESRegistrationPatch() *EESRegistrationPatch {
	this := EESRegistrationPatch{}
	return &this
}

// NewEESRegistrationPatchWithDefaults instantiates a new EESRegistrationPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEESRegistrationPatchWithDefaults() *EESRegistrationPatch {
	this := EESRegistrationPatch{}
	return &this
}

// GetEesProf returns the EesProf field value if set, zero value otherwise.
func (o *EESRegistrationPatch) GetEesProf() EESProfile {
	if o == nil || IsNil(o.EesProf) {
		var ret EESProfile
		return ret
	}
	return *o.EesProf
}

// GetEesProfOk returns a tuple with the EesProf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESRegistrationPatch) GetEesProfOk() (*EESProfile, bool) {
	if o == nil || IsNil(o.EesProf) {
		return nil, false
	}
	return o.EesProf, true
}

// HasEesProf returns a boolean if a field has been set.
func (o *EESRegistrationPatch) HasEesProf() bool {
	if o != nil && !IsNil(o.EesProf) {
		return true
	}

	return false
}

// SetEesProf gets a reference to the given EESProfile and assigns it to the EesProf field.
func (o *EESRegistrationPatch) SetEesProf(v EESProfile) {
	o.EesProf = &v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EESRegistrationPatch) GetExpTime() time.Time {
	if o == nil || IsNil(o.ExpTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime.Get()
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EESRegistrationPatch) GetExpTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpTime.Get(), o.ExpTime.IsSet()
}

// HasExpTime returns a boolean if a field has been set.
func (o *EESRegistrationPatch) HasExpTime() bool {
	if o != nil && o.ExpTime.IsSet() {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given NullableTime and assigns it to the ExpTime field.
func (o *EESRegistrationPatch) SetExpTime(v time.Time) {
	o.ExpTime.Set(&v)
}
// SetExpTimeNil sets the value for ExpTime to be an explicit nil
func (o *EESRegistrationPatch) SetExpTimeNil() {
	o.ExpTime.Set(nil)
}

// UnsetExpTime ensures that no value is present for ExpTime, not even an explicit nil
func (o *EESRegistrationPatch) UnsetExpTime() {
	o.ExpTime.Unset()
}

func (o EESRegistrationPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EESRegistrationPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EesProf) {
		toSerialize["eesProf"] = o.EesProf
	}
	if o.ExpTime.IsSet() {
		toSerialize["expTime"] = o.ExpTime.Get()
	}
	return toSerialize, nil
}

type NullableEESRegistrationPatch struct {
	value *EESRegistrationPatch
	isSet bool
}

func (v NullableEESRegistrationPatch) Get() *EESRegistrationPatch {
	return v.value
}

func (v *NullableEESRegistrationPatch) Set(val *EESRegistrationPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEESRegistrationPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEESRegistrationPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEESRegistrationPatch(val *EESRegistrationPatch) *NullableEESRegistrationPatch {
	return &NullableEESRegistrationPatch{value: val, isSet: true}
}

func (v NullableEESRegistrationPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEESRegistrationPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

