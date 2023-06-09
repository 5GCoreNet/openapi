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

// checks if the SNssaiSmfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SNssaiSmfInfoItem{}

// SNssaiSmfInfoItem struct for SNssaiSmfInfoItem
type SNssaiSmfInfoItem struct {
	SNSSAI         *Snssai          `json:"sNSSAI,omitempty"`
	DnnSmfInfoList []DnnSmfInfoItem `json:"dnnSmfInfoList,omitempty"`
}

// NewSNssaiSmfInfoItem instantiates a new SNssaiSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSNssaiSmfInfoItem() *SNssaiSmfInfoItem {
	this := SNssaiSmfInfoItem{}
	return &this
}

// NewSNssaiSmfInfoItemWithDefaults instantiates a new SNssaiSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSNssaiSmfInfoItemWithDefaults() *SNssaiSmfInfoItem {
	this := SNssaiSmfInfoItem{}
	return &this
}

// GetSNSSAI returns the SNSSAI field value if set, zero value otherwise.
func (o *SNssaiSmfInfoItem) GetSNSSAI() Snssai {
	if o == nil || IsNil(o.SNSSAI) {
		var ret Snssai
		return ret
	}
	return *o.SNSSAI
}

// GetSNSSAIOk returns a tuple with the SNSSAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNssaiSmfInfoItem) GetSNSSAIOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SNSSAI) {
		return nil, false
	}
	return o.SNSSAI, true
}

// HasSNSSAI returns a boolean if a field has been set.
func (o *SNssaiSmfInfoItem) HasSNSSAI() bool {
	if o != nil && !IsNil(o.SNSSAI) {
		return true
	}

	return false
}

// SetSNSSAI gets a reference to the given Snssai and assigns it to the SNSSAI field.
func (o *SNssaiSmfInfoItem) SetSNSSAI(v Snssai) {
	o.SNSSAI = &v
}

// GetDnnSmfInfoList returns the DnnSmfInfoList field value if set, zero value otherwise.
func (o *SNssaiSmfInfoItem) GetDnnSmfInfoList() []DnnSmfInfoItem {
	if o == nil || IsNil(o.DnnSmfInfoList) {
		var ret []DnnSmfInfoItem
		return ret
	}
	return o.DnnSmfInfoList
}

// GetDnnSmfInfoListOk returns a tuple with the DnnSmfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNssaiSmfInfoItem) GetDnnSmfInfoListOk() ([]DnnSmfInfoItem, bool) {
	if o == nil || IsNil(o.DnnSmfInfoList) {
		return nil, false
	}
	return o.DnnSmfInfoList, true
}

// HasDnnSmfInfoList returns a boolean if a field has been set.
func (o *SNssaiSmfInfoItem) HasDnnSmfInfoList() bool {
	if o != nil && !IsNil(o.DnnSmfInfoList) {
		return true
	}

	return false
}

// SetDnnSmfInfoList gets a reference to the given []DnnSmfInfoItem and assigns it to the DnnSmfInfoList field.
func (o *SNssaiSmfInfoItem) SetDnnSmfInfoList(v []DnnSmfInfoItem) {
	o.DnnSmfInfoList = v
}

func (o SNssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SNssaiSmfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SNSSAI) {
		toSerialize["sNSSAI"] = o.SNSSAI
	}
	if !IsNil(o.DnnSmfInfoList) {
		toSerialize["dnnSmfInfoList"] = o.DnnSmfInfoList
	}
	return toSerialize, nil
}

type NullableSNssaiSmfInfoItem struct {
	value *SNssaiSmfInfoItem
	isSet bool
}

func (v NullableSNssaiSmfInfoItem) Get() *SNssaiSmfInfoItem {
	return v.value
}

func (v *NullableSNssaiSmfInfoItem) Set(val *SNssaiSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSNssaiSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSNssaiSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSNssaiSmfInfoItem(val *SNssaiSmfInfoItem) *NullableSNssaiSmfInfoItem {
	return &NullableSNssaiSmfInfoItem{value: val, isSet: true}
}

func (v NullableSNssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSNssaiSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
