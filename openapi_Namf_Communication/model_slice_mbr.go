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

// checks if the SliceMbr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceMbr{}

// SliceMbr MBR related to slice
type SliceMbr struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	Uplink string `json:"uplink"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	Downlink string `json:"downlink"`
}

// NewSliceMbr instantiates a new SliceMbr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceMbr(uplink string, downlink string) *SliceMbr {
	this := SliceMbr{}
	this.Uplink = uplink
	this.Downlink = downlink
	return &this
}

// NewSliceMbrWithDefaults instantiates a new SliceMbr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceMbrWithDefaults() *SliceMbr {
	this := SliceMbr{}
	return &this
}

// GetUplink returns the Uplink field value
func (o *SliceMbr) GetUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value
// and a boolean to check if the value has been set.
func (o *SliceMbr) GetUplinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uplink, true
}

// SetUplink sets field value
func (o *SliceMbr) SetUplink(v string) {
	o.Uplink = v
}

// GetDownlink returns the Downlink field value
func (o *SliceMbr) GetDownlink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Downlink
}

// GetDownlinkOk returns a tuple with the Downlink field value
// and a boolean to check if the value has been set.
func (o *SliceMbr) GetDownlinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downlink, true
}

// SetDownlink sets field value
func (o *SliceMbr) SetDownlink(v string) {
	o.Downlink = v
}

func (o SliceMbr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceMbr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uplink"] = o.Uplink
	toSerialize["downlink"] = o.Downlink
	return toSerialize, nil
}

type NullableSliceMbr struct {
	value *SliceMbr
	isSet bool
}

func (v NullableSliceMbr) Get() *SliceMbr {
	return v.value
}

func (v *NullableSliceMbr) Set(val *SliceMbr) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceMbr) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceMbr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceMbr(val *SliceMbr) *NullableSliceMbr {
	return &NullableSliceMbr{value: val, isSet: true}
}

func (v NullableSliceMbr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceMbr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


