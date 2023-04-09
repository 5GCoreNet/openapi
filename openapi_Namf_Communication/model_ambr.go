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

// checks if the Ambr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ambr{}

// Ambr Contains the maximum aggregated uplink and downlink bit rates.
type Ambr struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	Uplink string `json:"uplink"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	Downlink string `json:"downlink"`
}

// NewAmbr instantiates a new Ambr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmbr(uplink string, downlink string) *Ambr {
	this := Ambr{}
	this.Uplink = uplink
	this.Downlink = downlink
	return &this
}

// NewAmbrWithDefaults instantiates a new Ambr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmbrWithDefaults() *Ambr {
	this := Ambr{}
	return &this
}

// GetUplink returns the Uplink field value
func (o *Ambr) GetUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value
// and a boolean to check if the value has been set.
func (o *Ambr) GetUplinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uplink, true
}

// SetUplink sets field value
func (o *Ambr) SetUplink(v string) {
	o.Uplink = v
}

// GetDownlink returns the Downlink field value
func (o *Ambr) GetDownlink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Downlink
}

// GetDownlinkOk returns a tuple with the Downlink field value
// and a boolean to check if the value has been set.
func (o *Ambr) GetDownlinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downlink, true
}

// SetDownlink sets field value
func (o *Ambr) SetDownlink(v string) {
	o.Downlink = v
}

func (o Ambr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ambr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uplink"] = o.Uplink
	toSerialize["downlink"] = o.Downlink
	return toSerialize, nil
}

type NullableAmbr struct {
	value *Ambr
	isSet bool
}

func (v NullableAmbr) Get() *Ambr {
	return v.value
}

func (v *NullableAmbr) Set(val *Ambr) {
	v.value = val
	v.isSet = true
}

func (v NullableAmbr) IsSet() bool {
	return v.isSet
}

func (v *NullableAmbr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmbr(val *Ambr) *NullableAmbr {
	return &NullableAmbr{value: val, isSet: true}
}

func (v NullableAmbr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmbr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
