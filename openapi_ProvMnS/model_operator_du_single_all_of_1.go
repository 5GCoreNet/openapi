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

// checks if the OperatorDuSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorDuSingleAllOf1{}

// OperatorDuSingleAllOf1 struct for OperatorDuSingleAllOf1
type OperatorDuSingleAllOf1 struct {
	EPF1C *EPF1CSingle `json:"EP_F1C,omitempty"`
	EPF1U []EPF1USingle `json:"EP_F1U,omitempty"`
}

// NewOperatorDuSingleAllOf1 instantiates a new OperatorDuSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorDuSingleAllOf1() *OperatorDuSingleAllOf1 {
	this := OperatorDuSingleAllOf1{}
	return &this
}

// NewOperatorDuSingleAllOf1WithDefaults instantiates a new OperatorDuSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorDuSingleAllOf1WithDefaults() *OperatorDuSingleAllOf1 {
	this := OperatorDuSingleAllOf1{}
	return &this
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *OperatorDuSingleAllOf1) GetEPF1C() EPF1CSingle {
	if o == nil || isNil(o.EPF1C) {
		var ret EPF1CSingle
		return ret
	}
	return *o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingleAllOf1) GetEPF1COk() (*EPF1CSingle, bool) {
	if o == nil || isNil(o.EPF1C) {
		return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *OperatorDuSingleAllOf1) HasEPF1C() bool {
	if o != nil && !isNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given EPF1CSingle and assigns it to the EPF1C field.
func (o *OperatorDuSingleAllOf1) SetEPF1C(v EPF1CSingle) {
	o.EPF1C = &v
}

// GetEPF1U returns the EPF1U field value if set, zero value otherwise.
func (o *OperatorDuSingleAllOf1) GetEPF1U() []EPF1USingle {
	if o == nil || isNil(o.EPF1U) {
		var ret []EPF1USingle
		return ret
	}
	return o.EPF1U
}

// GetEPF1UOk returns a tuple with the EPF1U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDuSingleAllOf1) GetEPF1UOk() ([]EPF1USingle, bool) {
	if o == nil || isNil(o.EPF1U) {
		return nil, false
	}
	return o.EPF1U, true
}

// HasEPF1U returns a boolean if a field has been set.
func (o *OperatorDuSingleAllOf1) HasEPF1U() bool {
	if o != nil && !isNil(o.EPF1U) {
		return true
	}

	return false
}

// SetEPF1U gets a reference to the given []EPF1USingle and assigns it to the EPF1U field.
func (o *OperatorDuSingleAllOf1) SetEPF1U(v []EPF1USingle) {
	o.EPF1U = v
}

func (o OperatorDuSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorDuSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	if !isNil(o.EPF1U) {
		toSerialize["EP_F1U"] = o.EPF1U
	}
	return toSerialize, nil
}

type NullableOperatorDuSingleAllOf1 struct {
	value *OperatorDuSingleAllOf1
	isSet bool
}

func (v NullableOperatorDuSingleAllOf1) Get() *OperatorDuSingleAllOf1 {
	return v.value
}

func (v *NullableOperatorDuSingleAllOf1) Set(val *OperatorDuSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorDuSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorDuSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorDuSingleAllOf1(val *OperatorDuSingleAllOf1) *NullableOperatorDuSingleAllOf1 {
	return &NullableOperatorDuSingleAllOf1{value: val, isSet: true}
}

func (v NullableOperatorDuSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorDuSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


