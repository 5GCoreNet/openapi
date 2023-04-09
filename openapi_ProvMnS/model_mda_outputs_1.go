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

// checks if the MDAOutputs1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAOutputs1{}

// MDAOutputs1 struct for MDAOutputs1
type MDAOutputs1 struct {
	MDAType       *string          `json:"mDAType,omitempty"`
	MdaOutputList []MDAOutputEntry `json:"mdaOutputList,omitempty"`
	MDARequestRef *string          `json:"mDARequestRef,omitempty"`
}

// NewMDAOutputs1 instantiates a new MDAOutputs1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAOutputs1() *MDAOutputs1 {
	this := MDAOutputs1{}
	return &this
}

// NewMDAOutputs1WithDefaults instantiates a new MDAOutputs1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAOutputs1WithDefaults() *MDAOutputs1 {
	this := MDAOutputs1{}
	return &this
}

// GetMDAType returns the MDAType field value if set, zero value otherwise.
func (o *MDAOutputs1) GetMDAType() string {
	if o == nil || IsNil(o.MDAType) {
		var ret string
		return ret
	}
	return *o.MDAType
}

// GetMDATypeOk returns a tuple with the MDAType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs1) GetMDATypeOk() (*string, bool) {
	if o == nil || IsNil(o.MDAType) {
		return nil, false
	}
	return o.MDAType, true
}

// HasMDAType returns a boolean if a field has been set.
func (o *MDAOutputs1) HasMDAType() bool {
	if o != nil && !IsNil(o.MDAType) {
		return true
	}

	return false
}

// SetMDAType gets a reference to the given string and assigns it to the MDAType field.
func (o *MDAOutputs1) SetMDAType(v string) {
	o.MDAType = &v
}

// GetMdaOutputList returns the MdaOutputList field value if set, zero value otherwise.
func (o *MDAOutputs1) GetMdaOutputList() []MDAOutputEntry {
	if o == nil || IsNil(o.MdaOutputList) {
		var ret []MDAOutputEntry
		return ret
	}
	return o.MdaOutputList
}

// GetMdaOutputListOk returns a tuple with the MdaOutputList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs1) GetMdaOutputListOk() ([]MDAOutputEntry, bool) {
	if o == nil || IsNil(o.MdaOutputList) {
		return nil, false
	}
	return o.MdaOutputList, true
}

// HasMdaOutputList returns a boolean if a field has been set.
func (o *MDAOutputs1) HasMdaOutputList() bool {
	if o != nil && !IsNil(o.MdaOutputList) {
		return true
	}

	return false
}

// SetMdaOutputList gets a reference to the given []MDAOutputEntry and assigns it to the MdaOutputList field.
func (o *MDAOutputs1) SetMdaOutputList(v []MDAOutputEntry) {
	o.MdaOutputList = v
}

// GetMDARequestRef returns the MDARequestRef field value if set, zero value otherwise.
func (o *MDAOutputs1) GetMDARequestRef() string {
	if o == nil || IsNil(o.MDARequestRef) {
		var ret string
		return ret
	}
	return *o.MDARequestRef
}

// GetMDARequestRefOk returns a tuple with the MDARequestRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs1) GetMDARequestRefOk() (*string, bool) {
	if o == nil || IsNil(o.MDARequestRef) {
		return nil, false
	}
	return o.MDARequestRef, true
}

// HasMDARequestRef returns a boolean if a field has been set.
func (o *MDAOutputs1) HasMDARequestRef() bool {
	if o != nil && !IsNil(o.MDARequestRef) {
		return true
	}

	return false
}

// SetMDARequestRef gets a reference to the given string and assigns it to the MDARequestRef field.
func (o *MDAOutputs1) SetMDARequestRef(v string) {
	o.MDARequestRef = &v
}

func (o MDAOutputs1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAOutputs1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAType) {
		toSerialize["mDAType"] = o.MDAType
	}
	if !IsNil(o.MdaOutputList) {
		toSerialize["mdaOutputList"] = o.MdaOutputList
	}
	if !IsNil(o.MDARequestRef) {
		toSerialize["mDARequestRef"] = o.MDARequestRef
	}
	return toSerialize, nil
}

type NullableMDAOutputs1 struct {
	value *MDAOutputs1
	isSet bool
}

func (v NullableMDAOutputs1) Get() *MDAOutputs1 {
	return v.value
}

func (v *NullableMDAOutputs1) Set(val *MDAOutputs1) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAOutputs1) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAOutputs1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAOutputs1(val *MDAOutputs1) *NullableMDAOutputs1 {
	return &NullableMDAOutputs1{value: val, isSet: true}
}

func (v NullableMDAOutputs1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAOutputs1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
