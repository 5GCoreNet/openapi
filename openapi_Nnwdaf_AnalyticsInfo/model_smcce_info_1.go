/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the SmcceInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmcceInfo1{}

// SmcceInfo1 Represents the Session Management congestion control experience information.
type SmcceInfo1 struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	SmcceUeList SmcceUeList1 `json:"smcceUeList"`
}

// NewSmcceInfo1 instantiates a new SmcceInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmcceInfo1(smcceUeList SmcceUeList1) *SmcceInfo1 {
	this := SmcceInfo1{}
	this.SmcceUeList = smcceUeList
	return &this
}

// NewSmcceInfo1WithDefaults instantiates a new SmcceInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmcceInfo1WithDefaults() *SmcceInfo1 {
	this := SmcceInfo1{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SmcceInfo1) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmcceInfo1) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SmcceInfo1) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SmcceInfo1) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *SmcceInfo1) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmcceInfo1) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *SmcceInfo1) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *SmcceInfo1) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetSmcceUeList returns the SmcceUeList field value
func (o *SmcceInfo1) GetSmcceUeList() SmcceUeList1 {
	if o == nil {
		var ret SmcceUeList1
		return ret
	}

	return o.SmcceUeList
}

// GetSmcceUeListOk returns a tuple with the SmcceUeList field value
// and a boolean to check if the value has been set.
func (o *SmcceInfo1) GetSmcceUeListOk() (*SmcceUeList1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmcceUeList, true
}

// SetSmcceUeList sets field value
func (o *SmcceInfo1) SetSmcceUeList(v SmcceUeList1) {
	o.SmcceUeList = v
}

func (o SmcceInfo1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmcceInfo1) ToMap() (map[string]interface{}, error) {
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

type NullableSmcceInfo1 struct {
	value *SmcceInfo1
	isSet bool
}

func (v NullableSmcceInfo1) Get() *SmcceInfo1 {
	return v.value
}

func (v *NullableSmcceInfo1) Set(val *SmcceInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableSmcceInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableSmcceInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmcceInfo1(val *SmcceInfo1) *NullableSmcceInfo1 {
	return &NullableSmcceInfo1{value: val, isSet: true}
}

func (v NullableSmcceInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmcceInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


