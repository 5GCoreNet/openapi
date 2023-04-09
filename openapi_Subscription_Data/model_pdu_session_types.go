/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the PduSessionTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSessionTypes{}

// PduSessionTypes struct for PduSessionTypes
type PduSessionTypes struct {
	DefaultSessionType  *PduSessionType  `json:"defaultSessionType,omitempty"`
	AllowedSessionTypes []PduSessionType `json:"allowedSessionTypes,omitempty"`
}

// NewPduSessionTypes instantiates a new PduSessionTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionTypes() *PduSessionTypes {
	this := PduSessionTypes{}
	return &this
}

// NewPduSessionTypesWithDefaults instantiates a new PduSessionTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionTypesWithDefaults() *PduSessionTypes {
	this := PduSessionTypes{}
	return &this
}

// GetDefaultSessionType returns the DefaultSessionType field value if set, zero value otherwise.
func (o *PduSessionTypes) GetDefaultSessionType() PduSessionType {
	if o == nil || IsNil(o.DefaultSessionType) {
		var ret PduSessionType
		return ret
	}
	return *o.DefaultSessionType
}

// GetDefaultSessionTypeOk returns a tuple with the DefaultSessionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes) GetDefaultSessionTypeOk() (*PduSessionType, bool) {
	if o == nil || IsNil(o.DefaultSessionType) {
		return nil, false
	}
	return o.DefaultSessionType, true
}

// HasDefaultSessionType returns a boolean if a field has been set.
func (o *PduSessionTypes) HasDefaultSessionType() bool {
	if o != nil && !IsNil(o.DefaultSessionType) {
		return true
	}

	return false
}

// SetDefaultSessionType gets a reference to the given PduSessionType and assigns it to the DefaultSessionType field.
func (o *PduSessionTypes) SetDefaultSessionType(v PduSessionType) {
	o.DefaultSessionType = &v
}

// GetAllowedSessionTypes returns the AllowedSessionTypes field value if set, zero value otherwise.
func (o *PduSessionTypes) GetAllowedSessionTypes() []PduSessionType {
	if o == nil || IsNil(o.AllowedSessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.AllowedSessionTypes
}

// GetAllowedSessionTypesOk returns a tuple with the AllowedSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionTypes) GetAllowedSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || IsNil(o.AllowedSessionTypes) {
		return nil, false
	}
	return o.AllowedSessionTypes, true
}

// HasAllowedSessionTypes returns a boolean if a field has been set.
func (o *PduSessionTypes) HasAllowedSessionTypes() bool {
	if o != nil && !IsNil(o.AllowedSessionTypes) {
		return true
	}

	return false
}

// SetAllowedSessionTypes gets a reference to the given []PduSessionType and assigns it to the AllowedSessionTypes field.
func (o *PduSessionTypes) SetAllowedSessionTypes(v []PduSessionType) {
	o.AllowedSessionTypes = v
}

func (o PduSessionTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSessionTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultSessionType) {
		toSerialize["defaultSessionType"] = o.DefaultSessionType
	}
	if !IsNil(o.AllowedSessionTypes) {
		toSerialize["allowedSessionTypes"] = o.AllowedSessionTypes
	}
	return toSerialize, nil
}

type NullablePduSessionTypes struct {
	value *PduSessionTypes
	isSet bool
}

func (v NullablePduSessionTypes) Get() *PduSessionTypes {
	return v.value
}

func (v *NullablePduSessionTypes) Set(val *PduSessionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionTypes(val *PduSessionTypes) *NullablePduSessionTypes {
	return &NullablePduSessionTypes{value: val, isSet: true}
}

func (v NullablePduSessionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
