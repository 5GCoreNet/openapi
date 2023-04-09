/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the NrSectorCarrierSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrSectorCarrierSingleAllOf1{}

// NrSectorCarrierSingleAllOf1 struct for NrSectorCarrierSingleAllOf1
type NrSectorCarrierSingleAllOf1 struct {
	CommonBeamformingFunction *CommonBeamformingFunctionSingle `json:"CommonBeamformingFunction,omitempty"`
}

// NewNrSectorCarrierSingleAllOf1 instantiates a new NrSectorCarrierSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrSectorCarrierSingleAllOf1() *NrSectorCarrierSingleAllOf1 {
	this := NrSectorCarrierSingleAllOf1{}
	return &this
}

// NewNrSectorCarrierSingleAllOf1WithDefaults instantiates a new NrSectorCarrierSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrSectorCarrierSingleAllOf1WithDefaults() *NrSectorCarrierSingleAllOf1 {
	this := NrSectorCarrierSingleAllOf1{}
	return &this
}

// GetCommonBeamformingFunction returns the CommonBeamformingFunction field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOf1) GetCommonBeamformingFunction() CommonBeamformingFunctionSingle {
	if o == nil || IsNil(o.CommonBeamformingFunction) {
		var ret CommonBeamformingFunctionSingle
		return ret
	}
	return *o.CommonBeamformingFunction
}

// GetCommonBeamformingFunctionOk returns a tuple with the CommonBeamformingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOf1) GetCommonBeamformingFunctionOk() (*CommonBeamformingFunctionSingle, bool) {
	if o == nil || IsNil(o.CommonBeamformingFunction) {
		return nil, false
	}
	return o.CommonBeamformingFunction, true
}

// HasCommonBeamformingFunction returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOf1) HasCommonBeamformingFunction() bool {
	if o != nil && !IsNil(o.CommonBeamformingFunction) {
		return true
	}

	return false
}

// SetCommonBeamformingFunction gets a reference to the given CommonBeamformingFunctionSingle and assigns it to the CommonBeamformingFunction field.
func (o *NrSectorCarrierSingleAllOf1) SetCommonBeamformingFunction(v CommonBeamformingFunctionSingle) {
	o.CommonBeamformingFunction = &v
}

func (o NrSectorCarrierSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrSectorCarrierSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonBeamformingFunction) {
		toSerialize["CommonBeamformingFunction"] = o.CommonBeamformingFunction
	}
	return toSerialize, nil
}

type NullableNrSectorCarrierSingleAllOf1 struct {
	value *NrSectorCarrierSingleAllOf1
	isSet bool
}

func (v NullableNrSectorCarrierSingleAllOf1) Get() *NrSectorCarrierSingleAllOf1 {
	return v.value
}

func (v *NullableNrSectorCarrierSingleAllOf1) Set(val *NrSectorCarrierSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNrSectorCarrierSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNrSectorCarrierSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrSectorCarrierSingleAllOf1(val *NrSectorCarrierSingleAllOf1) *NullableNrSectorCarrierSingleAllOf1 {
	return &NullableNrSectorCarrierSingleAllOf1{value: val, isSet: true}
}

func (v NullableNrSectorCarrierSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrSectorCarrierSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
