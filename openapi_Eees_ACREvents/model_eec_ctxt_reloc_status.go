/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
)

// checks if the EecCtxtRelocStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EecCtxtRelocStatus{}

// EecCtxtRelocStatus Indicates the registration id and expiry time of the registration.
type EecCtxtRelocStatus struct {
	ImplReg *ImplicitRegDetails `json:"implReg,omitempty"`
}

// NewEecCtxtRelocStatus instantiates a new EecCtxtRelocStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEecCtxtRelocStatus() *EecCtxtRelocStatus {
	this := EecCtxtRelocStatus{}
	return &this
}

// NewEecCtxtRelocStatusWithDefaults instantiates a new EecCtxtRelocStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEecCtxtRelocStatusWithDefaults() *EecCtxtRelocStatus {
	this := EecCtxtRelocStatus{}
	return &this
}

// GetImplReg returns the ImplReg field value if set, zero value otherwise.
func (o *EecCtxtRelocStatus) GetImplReg() ImplicitRegDetails {
	if o == nil || IsNil(o.ImplReg) {
		var ret ImplicitRegDetails
		return ret
	}
	return *o.ImplReg
}

// GetImplRegOk returns a tuple with the ImplReg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EecCtxtRelocStatus) GetImplRegOk() (*ImplicitRegDetails, bool) {
	if o == nil || IsNil(o.ImplReg) {
		return nil, false
	}
	return o.ImplReg, true
}

// HasImplReg returns a boolean if a field has been set.
func (o *EecCtxtRelocStatus) HasImplReg() bool {
	if o != nil && !IsNil(o.ImplReg) {
		return true
	}

	return false
}

// SetImplReg gets a reference to the given ImplicitRegDetails and assigns it to the ImplReg field.
func (o *EecCtxtRelocStatus) SetImplReg(v ImplicitRegDetails) {
	o.ImplReg = &v
}

func (o EecCtxtRelocStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EecCtxtRelocStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImplReg) {
		toSerialize["implReg"] = o.ImplReg
	}
	return toSerialize, nil
}

type NullableEecCtxtRelocStatus struct {
	value *EecCtxtRelocStatus
	isSet bool
}

func (v NullableEecCtxtRelocStatus) Get() *EecCtxtRelocStatus {
	return v.value
}

func (v *NullableEecCtxtRelocStatus) Set(val *EecCtxtRelocStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEecCtxtRelocStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEecCtxtRelocStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEecCtxtRelocStatus(val *EecCtxtRelocStatus) *NullableEecCtxtRelocStatus {
	return &NullableEecCtxtRelocStatus{value: val, isSet: true}
}

func (v NullableEecCtxtRelocStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEecCtxtRelocStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

