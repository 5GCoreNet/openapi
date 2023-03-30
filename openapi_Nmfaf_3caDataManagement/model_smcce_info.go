/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the SmcceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmcceInfo{}

// SmcceInfo Represents the Session Management congestion control experience information.
type SmcceInfo struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	SmcceUeList SmcceUeList `json:"smcceUeList"`
}

// NewSmcceInfo instantiates a new SmcceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmcceInfo(smcceUeList SmcceUeList) *SmcceInfo {
	this := SmcceInfo{}
	this.SmcceUeList = smcceUeList
	return &this
}

// NewSmcceInfoWithDefaults instantiates a new SmcceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmcceInfoWithDefaults() *SmcceInfo {
	this := SmcceInfo{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SmcceInfo) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmcceInfo) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SmcceInfo) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SmcceInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *SmcceInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmcceInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *SmcceInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *SmcceInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetSmcceUeList returns the SmcceUeList field value
func (o *SmcceInfo) GetSmcceUeList() SmcceUeList {
	if o == nil {
		var ret SmcceUeList
		return ret
	}

	return o.SmcceUeList
}

// GetSmcceUeListOk returns a tuple with the SmcceUeList field value
// and a boolean to check if the value has been set.
func (o *SmcceInfo) GetSmcceUeListOk() (*SmcceUeList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmcceUeList, true
}

// SetSmcceUeList sets field value
func (o *SmcceInfo) SetSmcceUeList(v SmcceUeList) {
	o.SmcceUeList = v
}

func (o SmcceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmcceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["smcceUeList"] = o.SmcceUeList
	return toSerialize, nil
}

type NullableSmcceInfo struct {
	value *SmcceInfo
	isSet bool
}

func (v NullableSmcceInfo) Get() *SmcceInfo {
	return v.value
}

func (v *NullableSmcceInfo) Set(val *SmcceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmcceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmcceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmcceInfo(val *SmcceInfo) *NullableSmcceInfo {
	return &NullableSmcceInfo{value: val, isSet: true}
}

func (v NullableSmcceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmcceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


