/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// checks if the ConfiguredSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfiguredSnssai{}

// ConfiguredSnssai Contains the configured S-NSSAI(s) authorized by the NSSF in the serving PLMN and optional mapped home S-NSSAI
type ConfiguredSnssai struct {
	ConfiguredSnssai Snssai `json:"configuredSnssai"`
	MappedHomeSnssai *Snssai `json:"mappedHomeSnssai,omitempty"`
}

// NewConfiguredSnssai instantiates a new ConfiguredSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredSnssai(configuredSnssai Snssai) *ConfiguredSnssai {
	this := ConfiguredSnssai{}
	this.ConfiguredSnssai = configuredSnssai
	return &this
}

// NewConfiguredSnssaiWithDefaults instantiates a new ConfiguredSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredSnssaiWithDefaults() *ConfiguredSnssai {
	this := ConfiguredSnssai{}
	return &this
}

// GetConfiguredSnssai returns the ConfiguredSnssai field value
func (o *ConfiguredSnssai) GetConfiguredSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.ConfiguredSnssai
}

// GetConfiguredSnssaiOk returns a tuple with the ConfiguredSnssai field value
// and a boolean to check if the value has been set.
func (o *ConfiguredSnssai) GetConfiguredSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfiguredSnssai, true
}

// SetConfiguredSnssai sets field value
func (o *ConfiguredSnssai) SetConfiguredSnssai(v Snssai) {
	o.ConfiguredSnssai = v
}

// GetMappedHomeSnssai returns the MappedHomeSnssai field value if set, zero value otherwise.
func (o *ConfiguredSnssai) GetMappedHomeSnssai() Snssai {
	if o == nil || isNil(o.MappedHomeSnssai) {
		var ret Snssai
		return ret
	}
	return *o.MappedHomeSnssai
}

// GetMappedHomeSnssaiOk returns a tuple with the MappedHomeSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredSnssai) GetMappedHomeSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.MappedHomeSnssai) {
		return nil, false
	}
	return o.MappedHomeSnssai, true
}

// HasMappedHomeSnssai returns a boolean if a field has been set.
func (o *ConfiguredSnssai) HasMappedHomeSnssai() bool {
	if o != nil && !isNil(o.MappedHomeSnssai) {
		return true
	}

	return false
}

// SetMappedHomeSnssai gets a reference to the given Snssai and assigns it to the MappedHomeSnssai field.
func (o *ConfiguredSnssai) SetMappedHomeSnssai(v Snssai) {
	o.MappedHomeSnssai = &v
}

func (o ConfiguredSnssai) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfiguredSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuredSnssai"] = o.ConfiguredSnssai
	if !isNil(o.MappedHomeSnssai) {
		toSerialize["mappedHomeSnssai"] = o.MappedHomeSnssai
	}
	return toSerialize, nil
}

type NullableConfiguredSnssai struct {
	value *ConfiguredSnssai
	isSet bool
}

func (v NullableConfiguredSnssai) Get() *ConfiguredSnssai {
	return v.value
}

func (v *NullableConfiguredSnssai) Set(val *ConfiguredSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguredSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguredSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguredSnssai(val *ConfiguredSnssai) *NullableConfiguredSnssai {
	return &NullableConfiguredSnssai{value: val, isSet: true}
}

func (v NullableConfiguredSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguredSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


