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

// checks if the VarSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VarSnssai{}

// VarSnssai When Snssai needs to be converted to string (e.g. when used in maps as key), the string shall be composed of one to three digits \"sst\" optionally followed by \"-\" and 6 hexadecimal digits \"sd\". 
type VarSnssai struct {
	// Unsigned integer, within the range 0 to 255, representing the Slice/Service Type.  It indicates the expected Network Slice behaviour in terms of features and services. Values 0 to 127 correspond to the standardized SST range. Values 128 to 255 correspond  to the Operator-specific range. See clause 28.4.2 of 3GPP TS 23.003. Standardized values are defined in clause 5.15.2.2 of 3GPP TS 23.501.  
	Sst int32 `json:"sst"`
	// 3-octet string, representing the Slice Differentiator, in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the SD shall appear first in the string, and the character representing the 4 least significant bit of the SD shall appear last in the string. This is an optional parameter that complements the Slice/Service type(s) to allow to  differentiate amongst multiple Network Slices of the same Slice/Service type. This IE shall be absent if no SD value is associated with the SST. 
	Sd *string `json:"sd,omitempty"`
}

// NewVarSnssai instantiates a new VarSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVarSnssai(sst int32) *VarSnssai {
	this := VarSnssai{}
	this.Sst = sst
	return &this
}

// NewVarSnssaiWithDefaults instantiates a new VarSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVarSnssaiWithDefaults() *VarSnssai {
	this := VarSnssai{}
	return &this
}

// GetSst returns the Sst field value
func (o *VarSnssai) GetSst() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sst
}

// GetSstOk returns a tuple with the Sst field value
// and a boolean to check if the value has been set.
func (o *VarSnssai) GetSstOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sst, true
}

// SetSst sets field value
func (o *VarSnssai) SetSst(v int32) {
	o.Sst = v
}

// GetSd returns the Sd field value if set, zero value otherwise.
func (o *VarSnssai) GetSd() string {
	if o == nil || isNil(o.Sd) {
		var ret string
		return ret
	}
	return *o.Sd
}

// GetSdOk returns a tuple with the Sd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VarSnssai) GetSdOk() (*string, bool) {
	if o == nil || isNil(o.Sd) {
		return nil, false
	}
	return o.Sd, true
}

// HasSd returns a boolean if a field has been set.
func (o *VarSnssai) HasSd() bool {
	if o != nil && !isNil(o.Sd) {
		return true
	}

	return false
}

// SetSd gets a reference to the given string and assigns it to the Sd field.
func (o *VarSnssai) SetSd(v string) {
	o.Sd = &v
}

func (o VarSnssai) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VarSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sst"] = o.Sst
	if !isNil(o.Sd) {
		toSerialize["sd"] = o.Sd
	}
	return toSerialize, nil
}

type NullableVarSnssai struct {
	value *VarSnssai
	isSet bool
}

func (v NullableVarSnssai) Get() *VarSnssai {
	return v.value
}

func (v *NullableVarSnssai) Set(val *VarSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableVarSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableVarSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVarSnssai(val *VarSnssai) *NullableVarSnssai {
	return &NullableVarSnssai{value: val, isSet: true}
}

func (v NullableVarSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVarSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


