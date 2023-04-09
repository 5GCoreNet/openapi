/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the DnnMbSmfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnnMbSmfInfoItem{}

// DnnMbSmfInfoItem Parameters supported by an MB-SMF for a given DNN
type DnnMbSmfInfoItem struct {
	Dnn DnnSmfInfoItemDnn `json:"dnn"`
}

// NewDnnMbSmfInfoItem instantiates a new DnnMbSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnMbSmfInfoItem(dnn DnnSmfInfoItemDnn) *DnnMbSmfInfoItem {
	this := DnnMbSmfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnMbSmfInfoItemWithDefaults instantiates a new DnnMbSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnMbSmfInfoItemWithDefaults() *DnnMbSmfInfoItem {
	this := DnnMbSmfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnMbSmfInfoItem) GetDnn() DnnSmfInfoItemDnn {
	if o == nil {
		var ret DnnSmfInfoItemDnn
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnMbSmfInfoItem) GetDnnOk() (*DnnSmfInfoItemDnn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnMbSmfInfoItem) SetDnn(v DnnSmfInfoItemDnn) {
	o.Dnn = v
}

func (o DnnMbSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnnMbSmfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	return toSerialize, nil
}

type NullableDnnMbSmfInfoItem struct {
	value *DnnMbSmfInfoItem
	isSet bool
}

func (v NullableDnnMbSmfInfoItem) Get() *DnnMbSmfInfoItem {
	return v.value
}

func (v *NullableDnnMbSmfInfoItem) Set(val *DnnMbSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnMbSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnMbSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnMbSmfInfoItem(val *DnnMbSmfInfoItem) *NullableDnnMbSmfInfoItem {
	return &NullableDnnMbSmfInfoItem{value: val, isSet: true}
}

func (v NullableDnnMbSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnMbSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
