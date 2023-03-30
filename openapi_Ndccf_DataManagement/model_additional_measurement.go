/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the AdditionalMeasurement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalMeasurement{}

// AdditionalMeasurement Represents additional measurement information.
type AdditionalMeasurement struct {
	UnexpLoc *NetworkAreaInfo `json:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription `json:"unexpFlowTeps,omitempty"`
	UnexpWakes []time.Time `json:"unexpWakes,omitempty"`
	DdosAttack *AddressList `json:"ddosAttack,omitempty"`
	WrgDest *AddressList `json:"wrgDest,omitempty"`
	Circums []CircumstanceDescription `json:"circums,omitempty"`
}

// NewAdditionalMeasurement instantiates a new AdditionalMeasurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalMeasurement() *AdditionalMeasurement {
	this := AdditionalMeasurement{}
	return &this
}

// NewAdditionalMeasurementWithDefaults instantiates a new AdditionalMeasurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalMeasurementWithDefaults() *AdditionalMeasurement {
	this := AdditionalMeasurement{}
	return &this
}

// GetUnexpLoc returns the UnexpLoc field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpLoc() NetworkAreaInfo {
	if o == nil || IsNil(o.UnexpLoc) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.UnexpLoc
}

// GetUnexpLocOk returns a tuple with the UnexpLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpLocOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.UnexpLoc) {
		return nil, false
	}
	return o.UnexpLoc, true
}

// HasUnexpLoc returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpLoc() bool {
	if o != nil && !IsNil(o.UnexpLoc) {
		return true
	}

	return false
}

// SetUnexpLoc gets a reference to the given NetworkAreaInfo and assigns it to the UnexpLoc field.
func (o *AdditionalMeasurement) SetUnexpLoc(v NetworkAreaInfo) {
	o.UnexpLoc = &v
}

// GetUnexpFlowTeps returns the UnexpFlowTeps field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpFlowTeps() []IpEthFlowDescription {
	if o == nil || IsNil(o.UnexpFlowTeps) {
		var ret []IpEthFlowDescription
		return ret
	}
	return o.UnexpFlowTeps
}

// GetUnexpFlowTepsOk returns a tuple with the UnexpFlowTeps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpFlowTepsOk() ([]IpEthFlowDescription, bool) {
	if o == nil || IsNil(o.UnexpFlowTeps) {
		return nil, false
	}
	return o.UnexpFlowTeps, true
}

// HasUnexpFlowTeps returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpFlowTeps() bool {
	if o != nil && !IsNil(o.UnexpFlowTeps) {
		return true
	}

	return false
}

// SetUnexpFlowTeps gets a reference to the given []IpEthFlowDescription and assigns it to the UnexpFlowTeps field.
func (o *AdditionalMeasurement) SetUnexpFlowTeps(v []IpEthFlowDescription) {
	o.UnexpFlowTeps = v
}

// GetUnexpWakes returns the UnexpWakes field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetUnexpWakes() []time.Time {
	if o == nil || IsNil(o.UnexpWakes) {
		var ret []time.Time
		return ret
	}
	return o.UnexpWakes
}

// GetUnexpWakesOk returns a tuple with the UnexpWakes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetUnexpWakesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.UnexpWakes) {
		return nil, false
	}
	return o.UnexpWakes, true
}

// HasUnexpWakes returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasUnexpWakes() bool {
	if o != nil && !IsNil(o.UnexpWakes) {
		return true
	}

	return false
}

// SetUnexpWakes gets a reference to the given []time.Time and assigns it to the UnexpWakes field.
func (o *AdditionalMeasurement) SetUnexpWakes(v []time.Time) {
	o.UnexpWakes = v
}

// GetDdosAttack returns the DdosAttack field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetDdosAttack() AddressList {
	if o == nil || IsNil(o.DdosAttack) {
		var ret AddressList
		return ret
	}
	return *o.DdosAttack
}

// GetDdosAttackOk returns a tuple with the DdosAttack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetDdosAttackOk() (*AddressList, bool) {
	if o == nil || IsNil(o.DdosAttack) {
		return nil, false
	}
	return o.DdosAttack, true
}

// HasDdosAttack returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasDdosAttack() bool {
	if o != nil && !IsNil(o.DdosAttack) {
		return true
	}

	return false
}

// SetDdosAttack gets a reference to the given AddressList and assigns it to the DdosAttack field.
func (o *AdditionalMeasurement) SetDdosAttack(v AddressList) {
	o.DdosAttack = &v
}

// GetWrgDest returns the WrgDest field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetWrgDest() AddressList {
	if o == nil || IsNil(o.WrgDest) {
		var ret AddressList
		return ret
	}
	return *o.WrgDest
}

// GetWrgDestOk returns a tuple with the WrgDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetWrgDestOk() (*AddressList, bool) {
	if o == nil || IsNil(o.WrgDest) {
		return nil, false
	}
	return o.WrgDest, true
}

// HasWrgDest returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasWrgDest() bool {
	if o != nil && !IsNil(o.WrgDest) {
		return true
	}

	return false
}

// SetWrgDest gets a reference to the given AddressList and assigns it to the WrgDest field.
func (o *AdditionalMeasurement) SetWrgDest(v AddressList) {
	o.WrgDest = &v
}

// GetCircums returns the Circums field value if set, zero value otherwise.
func (o *AdditionalMeasurement) GetCircums() []CircumstanceDescription {
	if o == nil || IsNil(o.Circums) {
		var ret []CircumstanceDescription
		return ret
	}
	return o.Circums
}

// GetCircumsOk returns a tuple with the Circums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalMeasurement) GetCircumsOk() ([]CircumstanceDescription, bool) {
	if o == nil || IsNil(o.Circums) {
		return nil, false
	}
	return o.Circums, true
}

// HasCircums returns a boolean if a field has been set.
func (o *AdditionalMeasurement) HasCircums() bool {
	if o != nil && !IsNil(o.Circums) {
		return true
	}

	return false
}

// SetCircums gets a reference to the given []CircumstanceDescription and assigns it to the Circums field.
func (o *AdditionalMeasurement) SetCircums(v []CircumstanceDescription) {
	o.Circums = v
}

func (o AdditionalMeasurement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalMeasurement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnexpLoc) {
		toSerialize["unexpLoc"] = o.UnexpLoc
	}
	if !IsNil(o.UnexpFlowTeps) {
		toSerialize["unexpFlowTeps"] = o.UnexpFlowTeps
	}
	if !IsNil(o.UnexpWakes) {
		toSerialize["unexpWakes"] = o.UnexpWakes
	}
	if !IsNil(o.DdosAttack) {
		toSerialize["ddosAttack"] = o.DdosAttack
	}
	if !IsNil(o.WrgDest) {
		toSerialize["wrgDest"] = o.WrgDest
	}
	if !IsNil(o.Circums) {
		toSerialize["circums"] = o.Circums
	}
	return toSerialize, nil
}

type NullableAdditionalMeasurement struct {
	value *AdditionalMeasurement
	isSet bool
}

func (v NullableAdditionalMeasurement) Get() *AdditionalMeasurement {
	return v.value
}

func (v *NullableAdditionalMeasurement) Set(val *AdditionalMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalMeasurement(val *AdditionalMeasurement) *NullableAdditionalMeasurement {
	return &NullableAdditionalMeasurement{value: val, isSet: true}
}

func (v NullableAdditionalMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


