/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the PlmnTargetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnTargetType{}

// PlmnTargetType The PLMN for which sessions shall be selected in the Trace Session in case of management based activation when several PLMNs are supported in the RAN (this means that shared cells and not shared cells are allowed for the specified PLMN. Note that the PLMN Target might differ from the PLMN specified in the Trace Reference, as that specifies the PLMN that is containing the management system requesting the Trace Session from the NE. See 3GPP TS 32.422 clause 5.9b for additional details.
type PlmnTargetType struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

// NewPlmnTargetType instantiates a new PlmnTargetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnTargetType(mcc string, mnc string) *PlmnTargetType {
	this := PlmnTargetType{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnTargetTypeWithDefaults instantiates a new PlmnTargetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnTargetTypeWithDefaults() *PlmnTargetType {
	this := PlmnTargetType{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnTargetType) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnTargetType) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnTargetType) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnTargetType) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnTargetType) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnTargetType) SetMnc(v string) {
	o.Mnc = v
}

func (o PlmnTargetType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnTargetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	return toSerialize, nil
}

type NullablePlmnTargetType struct {
	value *PlmnTargetType
	isSet bool
}

func (v NullablePlmnTargetType) Get() *PlmnTargetType {
	return v.value
}

func (v *NullablePlmnTargetType) Set(val *PlmnTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnTargetType(val *PlmnTargetType) *NullablePlmnTargetType {
	return &NullablePlmnTargetType{value: val, isSet: true}
}

func (v NullablePlmnTargetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

