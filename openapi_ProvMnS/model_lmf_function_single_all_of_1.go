/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the LmfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LmfFunctionSingleAllOf1{}

// LmfFunctionSingleAllOf1 struct for LmfFunctionSingleAllOf1
type LmfFunctionSingleAllOf1 struct {
	EP_NLS []EPNLSSingle `json:"EP_NLS,omitempty"`
}

// NewLmfFunctionSingleAllOf1 instantiates a new LmfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLmfFunctionSingleAllOf1() *LmfFunctionSingleAllOf1 {
	this := LmfFunctionSingleAllOf1{}
	return &this
}

// NewLmfFunctionSingleAllOf1WithDefaults instantiates a new LmfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLmfFunctionSingleAllOf1WithDefaults() *LmfFunctionSingleAllOf1 {
	this := LmfFunctionSingleAllOf1{}
	return &this
}

// GetEP_NLS returns the EP_NLS field value if set, zero value otherwise.
func (o *LmfFunctionSingleAllOf1) GetEP_NLS() []EPNLSSingle {
	if o == nil || IsNil(o.EP_NLS) {
		var ret []EPNLSSingle
		return ret
	}
	return o.EP_NLS
}

// GetEP_NLSOk returns a tuple with the EP_NLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LmfFunctionSingleAllOf1) GetEP_NLSOk() ([]EPNLSSingle, bool) {
	if o == nil || IsNil(o.EP_NLS) {
		return nil, false
	}
	return o.EP_NLS, true
}

// HasEP_NLS returns a boolean if a field has been set.
func (o *LmfFunctionSingleAllOf1) HasEP_NLS() bool {
	if o != nil && !IsNil(o.EP_NLS) {
		return true
	}

	return false
}

// SetEP_NLS gets a reference to the given []EPNLSSingle and assigns it to the EP_NLS field.
func (o *LmfFunctionSingleAllOf1) SetEP_NLS(v []EPNLSSingle) {
	o.EP_NLS = v
}

func (o LmfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LmfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EP_NLS) {
		toSerialize["EP_NLS"] = o.EP_NLS
	}
	return toSerialize, nil
}

type NullableLmfFunctionSingleAllOf1 struct {
	value *LmfFunctionSingleAllOf1
	isSet bool
}

func (v NullableLmfFunctionSingleAllOf1) Get() *LmfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableLmfFunctionSingleAllOf1) Set(val *LmfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableLmfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableLmfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLmfFunctionSingleAllOf1(val *LmfFunctionSingleAllOf1) *NullableLmfFunctionSingleAllOf1 {
	return &NullableLmfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableLmfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLmfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


