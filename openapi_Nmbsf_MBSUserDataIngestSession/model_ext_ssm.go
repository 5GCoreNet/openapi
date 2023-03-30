/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// checks if the ExtSsm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtSsm{}

// ExtSsm SSM and Port Number
type ExtSsm struct {
	Ssm Ssm `json:"ssm"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}

// NewExtSsm instantiates a new ExtSsm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtSsm(ssm Ssm, portNumber int32) *ExtSsm {
	this := ExtSsm{}
	this.Ssm = ssm
	this.PortNumber = portNumber
	return &this
}

// NewExtSsmWithDefaults instantiates a new ExtSsm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtSsmWithDefaults() *ExtSsm {
	this := ExtSsm{}
	return &this
}

// GetSsm returns the Ssm field value
func (o *ExtSsm) GetSsm() Ssm {
	if o == nil {
		var ret Ssm
		return ret
	}

	return o.Ssm
}

// GetSsmOk returns a tuple with the Ssm field value
// and a boolean to check if the value has been set.
func (o *ExtSsm) GetSsmOk() (*Ssm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssm, true
}

// SetSsm sets field value
func (o *ExtSsm) SetSsm(v Ssm) {
	o.Ssm = v
}

// GetPortNumber returns the PortNumber field value
func (o *ExtSsm) GetPortNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value
// and a boolean to check if the value has been set.
func (o *ExtSsm) GetPortNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortNumber, true
}

// SetPortNumber sets field value
func (o *ExtSsm) SetPortNumber(v int32) {
	o.PortNumber = v
}

func (o ExtSsm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtSsm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ssm"] = o.Ssm
	toSerialize["portNumber"] = o.PortNumber
	return toSerialize, nil
}

type NullableExtSsm struct {
	value *ExtSsm
	isSet bool
}

func (v NullableExtSsm) Get() *ExtSsm {
	return v.value
}

func (v *NullableExtSsm) Set(val *ExtSsm) {
	v.value = val
	v.isSet = true
}

func (v NullableExtSsm) IsSet() bool {
	return v.isSet
}

func (v *NullableExtSsm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtSsm(val *ExtSsm) *NullableExtSsm {
	return &NullableExtSsm{value: val, isSet: true}
}

func (v NullableExtSsm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtSsm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


