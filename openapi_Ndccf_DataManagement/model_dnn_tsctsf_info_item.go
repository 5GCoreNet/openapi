/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the DnnTsctsfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnnTsctsfInfoItem{}

// DnnTsctsfInfoItem Parameters supported by an TSCTSF for a given DNN
type DnnTsctsfInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
}

// NewDnnTsctsfInfoItem instantiates a new DnnTsctsfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnTsctsfInfoItem(dnn DnnSmfInfoItemDnn) *DnnTsctsfInfoItem {
	this := DnnTsctsfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnTsctsfInfoItemWithDefaults instantiates a new DnnTsctsfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnTsctsfInfoItemWithDefaults() *DnnTsctsfInfoItem {
	this := DnnTsctsfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnTsctsfInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnTsctsfInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnTsctsfInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

func (o DnnTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnnTsctsfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	return toSerialize, nil
}

type NullableDnnTsctsfInfoItem struct {
	value *DnnTsctsfInfoItem
	isSet bool
}

func (v NullableDnnTsctsfInfoItem) Get() *DnnTsctsfInfoItem {
	return v.value
}

func (v *NullableDnnTsctsfInfoItem) Set(val *DnnTsctsfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnTsctsfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnTsctsfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnTsctsfInfoItem(val *DnnTsctsfInfoItem) *NullableDnnTsctsfInfoItem {
	return &NullableDnnTsctsfInfoItem{value: val, isSet: true}
}

func (v NullableDnnTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnTsctsfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


