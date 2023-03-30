/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the PcfSelectionAssistanceInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfSelectionAssistanceInfo1{}

// PcfSelectionAssistanceInfo1 struct for PcfSelectionAssistanceInfo1
type PcfSelectionAssistanceInfo1 struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	SingleNssai Snssai `json:"singleNssai"`
}

// NewPcfSelectionAssistanceInfo1 instantiates a new PcfSelectionAssistanceInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfSelectionAssistanceInfo1(dnn string, singleNssai Snssai) *PcfSelectionAssistanceInfo1 {
	this := PcfSelectionAssistanceInfo1{}
	this.Dnn = dnn
	this.SingleNssai = singleNssai
	return &this
}

// NewPcfSelectionAssistanceInfo1WithDefaults instantiates a new PcfSelectionAssistanceInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfSelectionAssistanceInfo1WithDefaults() *PcfSelectionAssistanceInfo1 {
	this := PcfSelectionAssistanceInfo1{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *PcfSelectionAssistanceInfo1) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PcfSelectionAssistanceInfo1) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PcfSelectionAssistanceInfo1) SetDnn(v string) {
	o.Dnn = v
}

// GetSingleNssai returns the SingleNssai field value
func (o *PcfSelectionAssistanceInfo1) GetSingleNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value
// and a boolean to check if the value has been set.
func (o *PcfSelectionAssistanceInfo1) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleNssai, true
}

// SetSingleNssai sets field value
func (o *PcfSelectionAssistanceInfo1) SetSingleNssai(v Snssai) {
	o.SingleNssai = v
}

func (o PcfSelectionAssistanceInfo1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfSelectionAssistanceInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	toSerialize["singleNssai"] = o.SingleNssai
	return toSerialize, nil
}

type NullablePcfSelectionAssistanceInfo1 struct {
	value *PcfSelectionAssistanceInfo1
	isSet bool
}

func (v NullablePcfSelectionAssistanceInfo1) Get() *PcfSelectionAssistanceInfo1 {
	return v.value
}

func (v *NullablePcfSelectionAssistanceInfo1) Set(val *PcfSelectionAssistanceInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfSelectionAssistanceInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfSelectionAssistanceInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfSelectionAssistanceInfo1(val *PcfSelectionAssistanceInfo1) *NullablePcfSelectionAssistanceInfo1 {
	return &NullablePcfSelectionAssistanceInfo1{value: val, isSet: true}
}

func (v NullablePcfSelectionAssistanceInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfSelectionAssistanceInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


