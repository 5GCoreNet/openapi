/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the RanNasRelCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RanNasRelCause{}

// RanNasRelCause Contains the RAN/NAS release cause.
type RanNasRelCause struct {
	NgApCause *NgApCause `json:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCause *int32 `json:"5gMmCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gSmCause *int32 `json:"5gSmCause,omitempty"`
	// Defines the EPS RAN/NAS release cause.
	EpsCause *string `json:"epsCause,omitempty"`
}

// NewRanNasRelCause instantiates a new RanNasRelCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRanNasRelCause() *RanNasRelCause {
	this := RanNasRelCause{}
	return &this
}

// NewRanNasRelCauseWithDefaults instantiates a new RanNasRelCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRanNasRelCauseWithDefaults() *RanNasRelCause {
	this := RanNasRelCause{}
	return &this
}

// GetNgApCause returns the NgApCause field value if set, zero value otherwise.
func (o *RanNasRelCause) GetNgApCause() NgApCause {
	if o == nil || isNil(o.NgApCause) {
		var ret NgApCause
		return ret
	}
	return *o.NgApCause
}

// GetNgApCauseOk returns a tuple with the NgApCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RanNasRelCause) GetNgApCauseOk() (*NgApCause, bool) {
	if o == nil || isNil(o.NgApCause) {
		return nil, false
	}
	return o.NgApCause, true
}

// HasNgApCause returns a boolean if a field has been set.
func (o *RanNasRelCause) HasNgApCause() bool {
	if o != nil && !isNil(o.NgApCause) {
		return true
	}

	return false
}

// SetNgApCause gets a reference to the given NgApCause and assigns it to the NgApCause field.
func (o *RanNasRelCause) SetNgApCause(v NgApCause) {
	o.NgApCause = &v
}

// GetVar5gMmCause returns the Var5gMmCause field value if set, zero value otherwise.
func (o *RanNasRelCause) GetVar5gMmCause() int32 {
	if o == nil || isNil(o.Var5gMmCause) {
		var ret int32
		return ret
	}
	return *o.Var5gMmCause
}

// GetVar5gMmCauseOk returns a tuple with the Var5gMmCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RanNasRelCause) GetVar5gMmCauseOk() (*int32, bool) {
	if o == nil || isNil(o.Var5gMmCause) {
		return nil, false
	}
	return o.Var5gMmCause, true
}

// HasVar5gMmCause returns a boolean if a field has been set.
func (o *RanNasRelCause) HasVar5gMmCause() bool {
	if o != nil && !isNil(o.Var5gMmCause) {
		return true
	}

	return false
}

// SetVar5gMmCause gets a reference to the given int32 and assigns it to the Var5gMmCause field.
func (o *RanNasRelCause) SetVar5gMmCause(v int32) {
	o.Var5gMmCause = &v
}

// GetVar5gSmCause returns the Var5gSmCause field value if set, zero value otherwise.
func (o *RanNasRelCause) GetVar5gSmCause() int32 {
	if o == nil || isNil(o.Var5gSmCause) {
		var ret int32
		return ret
	}
	return *o.Var5gSmCause
}

// GetVar5gSmCauseOk returns a tuple with the Var5gSmCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RanNasRelCause) GetVar5gSmCauseOk() (*int32, bool) {
	if o == nil || isNil(o.Var5gSmCause) {
		return nil, false
	}
	return o.Var5gSmCause, true
}

// HasVar5gSmCause returns a boolean if a field has been set.
func (o *RanNasRelCause) HasVar5gSmCause() bool {
	if o != nil && !isNil(o.Var5gSmCause) {
		return true
	}

	return false
}

// SetVar5gSmCause gets a reference to the given int32 and assigns it to the Var5gSmCause field.
func (o *RanNasRelCause) SetVar5gSmCause(v int32) {
	o.Var5gSmCause = &v
}

// GetEpsCause returns the EpsCause field value if set, zero value otherwise.
func (o *RanNasRelCause) GetEpsCause() string {
	if o == nil || isNil(o.EpsCause) {
		var ret string
		return ret
	}
	return *o.EpsCause
}

// GetEpsCauseOk returns a tuple with the EpsCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RanNasRelCause) GetEpsCauseOk() (*string, bool) {
	if o == nil || isNil(o.EpsCause) {
		return nil, false
	}
	return o.EpsCause, true
}

// HasEpsCause returns a boolean if a field has been set.
func (o *RanNasRelCause) HasEpsCause() bool {
	if o != nil && !isNil(o.EpsCause) {
		return true
	}

	return false
}

// SetEpsCause gets a reference to the given string and assigns it to the EpsCause field.
func (o *RanNasRelCause) SetEpsCause(v string) {
	o.EpsCause = &v
}

func (o RanNasRelCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RanNasRelCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NgApCause) {
		toSerialize["ngApCause"] = o.NgApCause
	}
	if !isNil(o.Var5gMmCause) {
		toSerialize["5gMmCause"] = o.Var5gMmCause
	}
	if !isNil(o.Var5gSmCause) {
		toSerialize["5gSmCause"] = o.Var5gSmCause
	}
	if !isNil(o.EpsCause) {
		toSerialize["epsCause"] = o.EpsCause
	}
	return toSerialize, nil
}

type NullableRanNasRelCause struct {
	value *RanNasRelCause
	isSet bool
}

func (v NullableRanNasRelCause) Get() *RanNasRelCause {
	return v.value
}

func (v *NullableRanNasRelCause) Set(val *RanNasRelCause) {
	v.value = val
	v.isSet = true
}

func (v NullableRanNasRelCause) IsSet() bool {
	return v.isSet
}

func (v *NullableRanNasRelCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRanNasRelCause(val *RanNasRelCause) *NullableRanNasRelCause {
	return &NullableRanNasRelCause{value: val, isSet: true}
}

func (v NullableRanNasRelCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRanNasRelCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


