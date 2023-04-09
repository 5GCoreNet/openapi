/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the PduSessionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PduSessionInfo{}

// PduSessionInfo indicates the DNN and S-NSSAI combination of a PDU session.
type PduSessionInfo struct {
	Snssai Snssai `json:"snssai"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`
}

// NewPduSessionInfo instantiates a new PduSessionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionInfo(snssai Snssai, dnn string) *PduSessionInfo {
	this := PduSessionInfo{}
	this.Snssai = snssai
	this.Dnn = dnn
	return &this
}

// NewPduSessionInfoWithDefaults instantiates a new PduSessionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionInfoWithDefaults() *PduSessionInfo {
	this := PduSessionInfo{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *PduSessionInfo) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *PduSessionInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *PduSessionInfo) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetDnn returns the Dnn field value
func (o *PduSessionInfo) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PduSessionInfo) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PduSessionInfo) SetDnn(v string) {
	o.Dnn = v
}

func (o PduSessionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PduSessionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snssai"] = o.Snssai
	toSerialize["dnn"] = o.Dnn
	return toSerialize, nil
}

type NullablePduSessionInfo struct {
	value *PduSessionInfo
	isSet bool
}

func (v NullablePduSessionInfo) Get() *PduSessionInfo {
	return v.value
}

func (v *NullablePduSessionInfo) Set(val *PduSessionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionInfo(val *PduSessionInfo) *NullablePduSessionInfo {
	return &NullablePduSessionInfo{value: val, isSet: true}
}

func (v NullablePduSessionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
