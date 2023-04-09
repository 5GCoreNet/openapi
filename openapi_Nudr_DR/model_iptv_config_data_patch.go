/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the IptvConfigDataPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IptvConfigDataPatch{}

// IptvConfigDataPatch Represents the parameters to request the modification of an IPTV Configuration resource.
type IptvConfigDataPatch struct {
	// Identifies a list of multicast address access control information. Any string value can be used as a key of the map.
	MultiAccCtrls *map[string]MulticastAccessControl `json:"multiAccCtrls,omitempty"`
}

// NewIptvConfigDataPatch instantiates a new IptvConfigDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIptvConfigDataPatch() *IptvConfigDataPatch {
	this := IptvConfigDataPatch{}
	return &this
}

// NewIptvConfigDataPatchWithDefaults instantiates a new IptvConfigDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIptvConfigDataPatchWithDefaults() *IptvConfigDataPatch {
	this := IptvConfigDataPatch{}
	return &this
}

// GetMultiAccCtrls returns the MultiAccCtrls field value if set, zero value otherwise.
func (o *IptvConfigDataPatch) GetMultiAccCtrls() map[string]MulticastAccessControl {
	if o == nil || IsNil(o.MultiAccCtrls) {
		var ret map[string]MulticastAccessControl
		return ret
	}
	return *o.MultiAccCtrls
}

// GetMultiAccCtrlsOk returns a tuple with the MultiAccCtrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IptvConfigDataPatch) GetMultiAccCtrlsOk() (*map[string]MulticastAccessControl, bool) {
	if o == nil || IsNil(o.MultiAccCtrls) {
		return nil, false
	}
	return o.MultiAccCtrls, true
}

// HasMultiAccCtrls returns a boolean if a field has been set.
func (o *IptvConfigDataPatch) HasMultiAccCtrls() bool {
	if o != nil && !IsNil(o.MultiAccCtrls) {
		return true
	}

	return false
}

// SetMultiAccCtrls gets a reference to the given map[string]MulticastAccessControl and assigns it to the MultiAccCtrls field.
func (o *IptvConfigDataPatch) SetMultiAccCtrls(v map[string]MulticastAccessControl) {
	o.MultiAccCtrls = &v
}

func (o IptvConfigDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IptvConfigDataPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MultiAccCtrls) {
		toSerialize["multiAccCtrls"] = o.MultiAccCtrls
	}
	return toSerialize, nil
}

type NullableIptvConfigDataPatch struct {
	value *IptvConfigDataPatch
	isSet bool
}

func (v NullableIptvConfigDataPatch) Get() *IptvConfigDataPatch {
	return v.value
}

func (v *NullableIptvConfigDataPatch) Set(val *IptvConfigDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableIptvConfigDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableIptvConfigDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIptvConfigDataPatch(val *IptvConfigDataPatch) *NullableIptvConfigDataPatch {
	return &NullableIptvConfigDataPatch{value: val, isSet: true}
}

func (v NullableIptvConfigDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIptvConfigDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
