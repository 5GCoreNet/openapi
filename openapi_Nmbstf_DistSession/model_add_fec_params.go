/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
)

// checks if the AddFecParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFecParams{}

// AddFecParams Represents additional scheme-specific parameters for AL-FEC configuration.
type AddFecParams struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}

// NewAddFecParams instantiates a new AddFecParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFecParams(paramName string, paramValue string) *AddFecParams {
	this := AddFecParams{}
	this.ParamName = paramName
	this.ParamValue = paramValue
	return &this
}

// NewAddFecParamsWithDefaults instantiates a new AddFecParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFecParamsWithDefaults() *AddFecParams {
	this := AddFecParams{}
	return &this
}

// GetParamName returns the ParamName field value
func (o *AddFecParams) GetParamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParamName
}

// GetParamNameOk returns a tuple with the ParamName field value
// and a boolean to check if the value has been set.
func (o *AddFecParams) GetParamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParamName, true
}

// SetParamName sets field value
func (o *AddFecParams) SetParamName(v string) {
	o.ParamName = v
}

// GetParamValue returns the ParamValue field value
func (o *AddFecParams) GetParamValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParamValue
}

// GetParamValueOk returns a tuple with the ParamValue field value
// and a boolean to check if the value has been set.
func (o *AddFecParams) GetParamValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParamValue, true
}

// SetParamValue sets field value
func (o *AddFecParams) SetParamValue(v string) {
	o.ParamValue = v
}

func (o AddFecParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFecParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paramName"] = o.ParamName
	toSerialize["paramValue"] = o.ParamValue
	return toSerialize, nil
}

type NullableAddFecParams struct {
	value *AddFecParams
	isSet bool
}

func (v NullableAddFecParams) Get() *AddFecParams {
	return v.value
}

func (v *NullableAddFecParams) Set(val *AddFecParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFecParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFecParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFecParams(val *AddFecParams) *NullableAddFecParams {
	return &NullableAddFecParams{value: val, isSet: true}
}

func (v NullableAddFecParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFecParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
