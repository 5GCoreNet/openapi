/*
Nudm_MT

UDM MT Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_MT

import (
	"encoding/json"
)

// checks if the Model5GSrvccInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GSrvccInfo{}

// Model5GSrvccInfo Represents 5G SRVCC information for a UE.
type Model5GSrvccInfo struct {
	Ue5GSrvccCapability bool `json:"ue5GSrvccCapability"`
	// String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003.
	StnSr *string `json:"stnSr,omitempty"`
	// String representing the C-MSISDN as defined in clause 18.7 of 3GPP TS 23.003.
	CMsisdn *string `json:"cMsisdn,omitempty"`
}

// NewModel5GSrvccInfo instantiates a new Model5GSrvccInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GSrvccInfo(ue5GSrvccCapability bool) *Model5GSrvccInfo {
	this := Model5GSrvccInfo{}
	this.Ue5GSrvccCapability = ue5GSrvccCapability
	return &this
}

// NewModel5GSrvccInfoWithDefaults instantiates a new Model5GSrvccInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GSrvccInfoWithDefaults() *Model5GSrvccInfo {
	this := Model5GSrvccInfo{}
	return &this
}

// GetUe5GSrvccCapability returns the Ue5GSrvccCapability field value
func (o *Model5GSrvccInfo) GetUe5GSrvccCapability() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ue5GSrvccCapability
}

// GetUe5GSrvccCapabilityOk returns a tuple with the Ue5GSrvccCapability field value
// and a boolean to check if the value has been set.
func (o *Model5GSrvccInfo) GetUe5GSrvccCapabilityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ue5GSrvccCapability, true
}

// SetUe5GSrvccCapability sets field value
func (o *Model5GSrvccInfo) SetUe5GSrvccCapability(v bool) {
	o.Ue5GSrvccCapability = v
}

// GetStnSr returns the StnSr field value if set, zero value otherwise.
func (o *Model5GSrvccInfo) GetStnSr() string {
	if o == nil || IsNil(o.StnSr) {
		var ret string
		return ret
	}
	return *o.StnSr
}

// GetStnSrOk returns a tuple with the StnSr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GSrvccInfo) GetStnSrOk() (*string, bool) {
	if o == nil || IsNil(o.StnSr) {
		return nil, false
	}
	return o.StnSr, true
}

// HasStnSr returns a boolean if a field has been set.
func (o *Model5GSrvccInfo) HasStnSr() bool {
	if o != nil && !IsNil(o.StnSr) {
		return true
	}

	return false
}

// SetStnSr gets a reference to the given string and assigns it to the StnSr field.
func (o *Model5GSrvccInfo) SetStnSr(v string) {
	o.StnSr = &v
}

// GetCMsisdn returns the CMsisdn field value if set, zero value otherwise.
func (o *Model5GSrvccInfo) GetCMsisdn() string {
	if o == nil || IsNil(o.CMsisdn) {
		var ret string
		return ret
	}
	return *o.CMsisdn
}

// GetCMsisdnOk returns a tuple with the CMsisdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GSrvccInfo) GetCMsisdnOk() (*string, bool) {
	if o == nil || IsNil(o.CMsisdn) {
		return nil, false
	}
	return o.CMsisdn, true
}

// HasCMsisdn returns a boolean if a field has been set.
func (o *Model5GSrvccInfo) HasCMsisdn() bool {
	if o != nil && !IsNil(o.CMsisdn) {
		return true
	}

	return false
}

// SetCMsisdn gets a reference to the given string and assigns it to the CMsisdn field.
func (o *Model5GSrvccInfo) SetCMsisdn(v string) {
	o.CMsisdn = &v
}

func (o Model5GSrvccInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GSrvccInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ue5GSrvccCapability"] = o.Ue5GSrvccCapability
	if !IsNil(o.StnSr) {
		toSerialize["stnSr"] = o.StnSr
	}
	if !IsNil(o.CMsisdn) {
		toSerialize["cMsisdn"] = o.CMsisdn
	}
	return toSerialize, nil
}

type NullableModel5GSrvccInfo struct {
	value *Model5GSrvccInfo
	isSet bool
}

func (v NullableModel5GSrvccInfo) Get() *Model5GSrvccInfo {
	return v.value
}

func (v *NullableModel5GSrvccInfo) Set(val *Model5GSrvccInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GSrvccInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GSrvccInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GSrvccInfo(val *Model5GSrvccInfo) *NullableModel5GSrvccInfo {
	return &NullableModel5GSrvccInfo{value: val, isSet: true}
}

func (v NullableModel5GSrvccInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GSrvccInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


