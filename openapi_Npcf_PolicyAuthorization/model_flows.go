/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the Flows type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Flows{}

// Flows Identifies the flows.
type Flows struct {
	ContVers []int32 `json:"contVers,omitempty"`
	FNums    []int32 `json:"fNums,omitempty"`
	MedCompN int32   `json:"medCompN"`
}

// NewFlows instantiates a new Flows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlows(medCompN int32) *Flows {
	this := Flows{}
	this.MedCompN = medCompN
	return &this
}

// NewFlowsWithDefaults instantiates a new Flows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowsWithDefaults() *Flows {
	this := Flows{}
	return &this
}

// GetContVers returns the ContVers field value if set, zero value otherwise.
func (o *Flows) GetContVers() []int32 {
	if o == nil || IsNil(o.ContVers) {
		var ret []int32
		return ret
	}
	return o.ContVers
}

// GetContVersOk returns a tuple with the ContVers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flows) GetContVersOk() ([]int32, bool) {
	if o == nil || IsNil(o.ContVers) {
		return nil, false
	}
	return o.ContVers, true
}

// HasContVers returns a boolean if a field has been set.
func (o *Flows) HasContVers() bool {
	if o != nil && !IsNil(o.ContVers) {
		return true
	}

	return false
}

// SetContVers gets a reference to the given []int32 and assigns it to the ContVers field.
func (o *Flows) SetContVers(v []int32) {
	o.ContVers = v
}

// GetFNums returns the FNums field value if set, zero value otherwise.
func (o *Flows) GetFNums() []int32 {
	if o == nil || IsNil(o.FNums) {
		var ret []int32
		return ret
	}
	return o.FNums
}

// GetFNumsOk returns a tuple with the FNums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Flows) GetFNumsOk() ([]int32, bool) {
	if o == nil || IsNil(o.FNums) {
		return nil, false
	}
	return o.FNums, true
}

// HasFNums returns a boolean if a field has been set.
func (o *Flows) HasFNums() bool {
	if o != nil && !IsNil(o.FNums) {
		return true
	}

	return false
}

// SetFNums gets a reference to the given []int32 and assigns it to the FNums field.
func (o *Flows) SetFNums(v []int32) {
	o.FNums = v
}

// GetMedCompN returns the MedCompN field value
func (o *Flows) GetMedCompN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MedCompN
}

// GetMedCompNOk returns a tuple with the MedCompN field value
// and a boolean to check if the value has been set.
func (o *Flows) GetMedCompNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MedCompN, true
}

// SetMedCompN sets field value
func (o *Flows) SetMedCompN(v int32) {
	o.MedCompN = v
}

func (o Flows) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Flows) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContVers) {
		toSerialize["contVers"] = o.ContVers
	}
	if !IsNil(o.FNums) {
		toSerialize["fNums"] = o.FNums
	}
	toSerialize["medCompN"] = o.MedCompN
	return toSerialize, nil
}

type NullableFlows struct {
	value *Flows
	isSet bool
}

func (v NullableFlows) Get() *Flows {
	return v.value
}

func (v *NullableFlows) Set(val *Flows) {
	v.value = val
	v.isSet = true
}

func (v NullableFlows) IsSet() bool {
	return v.isSet
}

func (v *NullableFlows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlows(val *Flows) *NullableFlows {
	return &NullableFlows{value: val, isSet: true}
}

func (v NullableFlows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
