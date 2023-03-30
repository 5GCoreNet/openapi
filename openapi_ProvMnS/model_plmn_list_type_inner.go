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

// checks if the PlmnListTypeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnListTypeInner{}

// PlmnListTypeInner struct for PlmnListTypeInner
type PlmnListTypeInner struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

// NewPlmnListTypeInner instantiates a new PlmnListTypeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnListTypeInner(mcc string, mnc string) *PlmnListTypeInner {
	this := PlmnListTypeInner{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnListTypeInnerWithDefaults instantiates a new PlmnListTypeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnListTypeInnerWithDefaults() *PlmnListTypeInner {
	this := PlmnListTypeInner{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnListTypeInner) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnListTypeInner) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnListTypeInner) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnListTypeInner) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnListTypeInner) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnListTypeInner) SetMnc(v string) {
	o.Mnc = v
}

func (o PlmnListTypeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnListTypeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	return toSerialize, nil
}

type NullablePlmnListTypeInner struct {
	value *PlmnListTypeInner
	isSet bool
}

func (v NullablePlmnListTypeInner) Get() *PlmnListTypeInner {
	return v.value
}

func (v *NullablePlmnListTypeInner) Set(val *PlmnListTypeInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnListTypeInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnListTypeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnListTypeInner(val *PlmnListTypeInner) *NullablePlmnListTypeInner {
	return &NullablePlmnListTypeInner{value: val, isSet: true}
}

func (v NullablePlmnListTypeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnListTypeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


