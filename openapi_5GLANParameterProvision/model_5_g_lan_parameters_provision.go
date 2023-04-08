/*
3gpp-5glan-pp

API for 5G LAN Parameter Provision.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GLANParameterProvision

import (
	"encoding/json"
)

// checks if the Model5GLanParametersProvision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GLanParametersProvision{}

// Model5GLanParametersProvision Represents an individual 5G LAN parameters provision subscription resource.
type Model5GLanParametersProvision struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	Var5gLanParams Model5GLanParameters `json:"5gLanParams"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewModel5GLanParametersProvision instantiates a new Model5GLanParametersProvision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GLanParametersProvision(var5gLanParams Model5GLanParameters, suppFeat string) *Model5GLanParametersProvision {
	this := Model5GLanParametersProvision{}
	this.Var5gLanParams = var5gLanParams
	this.SuppFeat = suppFeat
	return &this
}

// NewModel5GLanParametersProvisionWithDefaults instantiates a new Model5GLanParametersProvision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GLanParametersProvisionWithDefaults() *Model5GLanParametersProvision {
	this := Model5GLanParametersProvision{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Model5GLanParametersProvision) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParametersProvision) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Model5GLanParametersProvision) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Model5GLanParametersProvision) SetSelf(v string) {
	o.Self = &v
}

// GetVar5gLanParams returns the Var5gLanParams field value
func (o *Model5GLanParametersProvision) GetVar5gLanParams() Model5GLanParameters {
	if o == nil {
		var ret Model5GLanParameters
		return ret
	}

	return o.Var5gLanParams
}

// GetVar5gLanParamsOk returns a tuple with the Var5gLanParams field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParametersProvision) GetVar5gLanParamsOk() (*Model5GLanParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5gLanParams, true
}

// SetVar5gLanParams sets field value
func (o *Model5GLanParametersProvision) SetVar5gLanParams(v Model5GLanParameters) {
	o.Var5gLanParams = v
}

// GetSuppFeat returns the SuppFeat field value
func (o *Model5GLanParametersProvision) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParametersProvision) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *Model5GLanParametersProvision) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o Model5GLanParametersProvision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GLanParametersProvision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	toSerialize["5gLanParams"] = o.Var5gLanParams
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullableModel5GLanParametersProvision struct {
	value *Model5GLanParametersProvision
	isSet bool
}

func (v NullableModel5GLanParametersProvision) Get() *Model5GLanParametersProvision {
	return v.value
}

func (v *NullableModel5GLanParametersProvision) Set(val *Model5GLanParametersProvision) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GLanParametersProvision) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GLanParametersProvision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GLanParametersProvision(val *Model5GLanParametersProvision) *NullableModel5GLanParametersProvision {
	return &NullableModel5GLanParametersProvision{value: val, isSet: true}
}

func (v NullableModel5GLanParametersProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GLanParametersProvision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


