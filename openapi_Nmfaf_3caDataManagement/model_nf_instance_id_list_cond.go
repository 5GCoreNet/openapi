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

// checks if the NfInstanceIdListCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfInstanceIdListCond{}

// NfInstanceIdListCond Subscription to a list of NF Instances
type NfInstanceIdListCond struct {
	NfInstanceIdList []string `json:"nfInstanceIdList"`
}

// NewNfInstanceIdListCond instantiates a new NfInstanceIdListCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfInstanceIdListCond(nfInstanceIdList []string) *NfInstanceIdListCond {
	this := NfInstanceIdListCond{}
	this.NfInstanceIdList = nfInstanceIdList
	return &this
}

// NewNfInstanceIdListCondWithDefaults instantiates a new NfInstanceIdListCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfInstanceIdListCondWithDefaults() *NfInstanceIdListCond {
	this := NfInstanceIdListCond{}
	return &this
}

// GetNfInstanceIdList returns the NfInstanceIdList field value
func (o *NfInstanceIdListCond) GetNfInstanceIdList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NfInstanceIdList
}

// GetNfInstanceIdListOk returns a tuple with the NfInstanceIdList field value
// and a boolean to check if the value has been set.
func (o *NfInstanceIdListCond) GetNfInstanceIdListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfInstanceIdList, true
}

// SetNfInstanceIdList sets field value
func (o *NfInstanceIdListCond) SetNfInstanceIdList(v []string) {
	o.NfInstanceIdList = v
}

func (o NfInstanceIdListCond) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfInstanceIdListCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfInstanceIdList"] = o.NfInstanceIdList
	return toSerialize, nil
}

type NullableNfInstanceIdListCond struct {
	value *NfInstanceIdListCond
	isSet bool
}

func (v NullableNfInstanceIdListCond) Get() *NfInstanceIdListCond {
	return v.value
}

func (v *NullableNfInstanceIdListCond) Set(val *NfInstanceIdListCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfInstanceIdListCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfInstanceIdListCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfInstanceIdListCond(val *NfInstanceIdListCond) *NullableNfInstanceIdListCond {
	return &NullableNfInstanceIdListCond{value: val, isSet: true}
}

func (v NullableNfInstanceIdListCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfInstanceIdListCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
