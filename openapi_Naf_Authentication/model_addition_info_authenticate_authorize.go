/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_Authentication

import (
	"encoding/json"
)

// checks if the AdditionInfoAuthenticateAuthorize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionInfoAuthenticateAuthorize{}

// AdditionInfoAuthenticateAuthorize Indicates additional information during authentication failure
type AdditionInfoAuthenticateAuthorize struct {
	// Indicates to release the UAV resources during authentication failure, when set to \"true\". Default is set to \"false\". 
	UasResRelInd *bool `json:"uasResRelInd,omitempty"`
}

// NewAdditionInfoAuthenticateAuthorize instantiates a new AdditionInfoAuthenticateAuthorize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionInfoAuthenticateAuthorize() *AdditionInfoAuthenticateAuthorize {
	this := AdditionInfoAuthenticateAuthorize{}
	return &this
}

// NewAdditionInfoAuthenticateAuthorizeWithDefaults instantiates a new AdditionInfoAuthenticateAuthorize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionInfoAuthenticateAuthorizeWithDefaults() *AdditionInfoAuthenticateAuthorize {
	this := AdditionInfoAuthenticateAuthorize{}
	return &this
}

// GetUasResRelInd returns the UasResRelInd field value if set, zero value otherwise.
func (o *AdditionInfoAuthenticateAuthorize) GetUasResRelInd() bool {
	if o == nil || IsNil(o.UasResRelInd) {
		var ret bool
		return ret
	}
	return *o.UasResRelInd
}

// GetUasResRelIndOk returns a tuple with the UasResRelInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionInfoAuthenticateAuthorize) GetUasResRelIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UasResRelInd) {
		return nil, false
	}
	return o.UasResRelInd, true
}

// HasUasResRelInd returns a boolean if a field has been set.
func (o *AdditionInfoAuthenticateAuthorize) HasUasResRelInd() bool {
	if o != nil && !IsNil(o.UasResRelInd) {
		return true
	}

	return false
}

// SetUasResRelInd gets a reference to the given bool and assigns it to the UasResRelInd field.
func (o *AdditionInfoAuthenticateAuthorize) SetUasResRelInd(v bool) {
	o.UasResRelInd = &v
}

func (o AdditionInfoAuthenticateAuthorize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionInfoAuthenticateAuthorize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UasResRelInd) {
		toSerialize["uasResRelInd"] = o.UasResRelInd
	}
	return toSerialize, nil
}

type NullableAdditionInfoAuthenticateAuthorize struct {
	value *AdditionInfoAuthenticateAuthorize
	isSet bool
}

func (v NullableAdditionInfoAuthenticateAuthorize) Get() *AdditionInfoAuthenticateAuthorize {
	return v.value
}

func (v *NullableAdditionInfoAuthenticateAuthorize) Set(val *AdditionInfoAuthenticateAuthorize) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionInfoAuthenticateAuthorize) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionInfoAuthenticateAuthorize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionInfoAuthenticateAuthorize(val *AdditionInfoAuthenticateAuthorize) *NullableAdditionInfoAuthenticateAuthorize {
	return &NullableAdditionInfoAuthenticateAuthorize{value: val, isSet: true}
}

func (v NullableAdditionInfoAuthenticateAuthorize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionInfoAuthenticateAuthorize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


