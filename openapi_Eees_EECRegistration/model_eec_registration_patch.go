/*
Eees_EECRegistration

API for EEC registration. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EECRegistration

import (
	"encoding/json"
	"time"
)

// checks if the EECRegistrationPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EECRegistrationPatch{}

// EECRegistrationPatch Describes the parameters to perform EEC Registration update.
type EECRegistrationPatch struct {
	// Profiles of ACs for which the EEC provides edge enabling services.
	AcProfs []ACProfile `json:"acProfs,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
	UnfulfilledAcProfs *UnfulfilledAcProfile `json:"unfulfilledAcProfs,omitempty"`
}

// NewEECRegistrationPatch instantiates a new EECRegistrationPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEECRegistrationPatch() *EECRegistrationPatch {
	this := EECRegistrationPatch{}
	return &this
}

// NewEECRegistrationPatchWithDefaults instantiates a new EECRegistrationPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEECRegistrationPatchWithDefaults() *EECRegistrationPatch {
	this := EECRegistrationPatch{}
	return &this
}

// GetAcProfs returns the AcProfs field value if set, zero value otherwise.
func (o *EECRegistrationPatch) GetAcProfs() []ACProfile {
	if o == nil || isNil(o.AcProfs) {
		var ret []ACProfile
		return ret
	}
	return o.AcProfs
}

// GetAcProfsOk returns a tuple with the AcProfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EECRegistrationPatch) GetAcProfsOk() ([]ACProfile, bool) {
	if o == nil || isNil(o.AcProfs) {
		return nil, false
	}
	return o.AcProfs, true
}

// HasAcProfs returns a boolean if a field has been set.
func (o *EECRegistrationPatch) HasAcProfs() bool {
	if o != nil && !isNil(o.AcProfs) {
		return true
	}

	return false
}

// SetAcProfs gets a reference to the given []ACProfile and assigns it to the AcProfs field.
func (o *EECRegistrationPatch) SetAcProfs(v []ACProfile) {
	o.AcProfs = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *EECRegistrationPatch) GetExpTime() time.Time {
	if o == nil || isNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EECRegistrationPatch) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *EECRegistrationPatch) HasExpTime() bool {
	if o != nil && !isNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *EECRegistrationPatch) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetUnfulfilledAcProfs returns the UnfulfilledAcProfs field value if set, zero value otherwise.
func (o *EECRegistrationPatch) GetUnfulfilledAcProfs() UnfulfilledAcProfile {
	if o == nil || isNil(o.UnfulfilledAcProfs) {
		var ret UnfulfilledAcProfile
		return ret
	}
	return *o.UnfulfilledAcProfs
}

// GetUnfulfilledAcProfsOk returns a tuple with the UnfulfilledAcProfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EECRegistrationPatch) GetUnfulfilledAcProfsOk() (*UnfulfilledAcProfile, bool) {
	if o == nil || isNil(o.UnfulfilledAcProfs) {
		return nil, false
	}
	return o.UnfulfilledAcProfs, true
}

// HasUnfulfilledAcProfs returns a boolean if a field has been set.
func (o *EECRegistrationPatch) HasUnfulfilledAcProfs() bool {
	if o != nil && !isNil(o.UnfulfilledAcProfs) {
		return true
	}

	return false
}

// SetUnfulfilledAcProfs gets a reference to the given UnfulfilledAcProfile and assigns it to the UnfulfilledAcProfs field.
func (o *EECRegistrationPatch) SetUnfulfilledAcProfs(v UnfulfilledAcProfile) {
	o.UnfulfilledAcProfs = &v
}

func (o EECRegistrationPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EECRegistrationPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AcProfs) {
		toSerialize["acProfs"] = o.AcProfs
	}
	if !isNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !isNil(o.UnfulfilledAcProfs) {
		toSerialize["unfulfilledAcProfs"] = o.UnfulfilledAcProfs
	}
	return toSerialize, nil
}

type NullableEECRegistrationPatch struct {
	value *EECRegistrationPatch
	isSet bool
}

func (v NullableEECRegistrationPatch) Get() *EECRegistrationPatch {
	return v.value
}

func (v *NullableEECRegistrationPatch) Set(val *EECRegistrationPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEECRegistrationPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEECRegistrationPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEECRegistrationPatch(val *EECRegistrationPatch) *NullableEECRegistrationPatch {
	return &NullableEECRegistrationPatch{value: val, isSet: true}
}

func (v NullableEECRegistrationPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEECRegistrationPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


