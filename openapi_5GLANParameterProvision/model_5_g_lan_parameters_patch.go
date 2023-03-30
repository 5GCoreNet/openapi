/*
3gpp-5glan-pp

API for 5G LAN Parameter Provision.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GLANParameterProvision

import (
	"encoding/json"
)

// checks if the Model5GLanParametersPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GLanParametersPatch{}

// Model5GLanParametersPatch Represents 5G LAN service related parameters that need to be modified.
type Model5GLanParametersPatch struct {
	// Contains the list of 5G VN Group members, each member is identified by GPSI. Any string value can be used as a key of the map. 
	Gpsis *map[string]string `json:"gpsis,omitempty"`
	// Describes the operation systems and the corresponding applications for each operation systems. The key of map is osId. 
	AppDesps *map[string]AppDescriptorRm `json:"appDesps,omitempty"`
}

// NewModel5GLanParametersPatch instantiates a new Model5GLanParametersPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GLanParametersPatch() *Model5GLanParametersPatch {
	this := Model5GLanParametersPatch{}
	return &this
}

// NewModel5GLanParametersPatchWithDefaults instantiates a new Model5GLanParametersPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GLanParametersPatchWithDefaults() *Model5GLanParametersPatch {
	this := Model5GLanParametersPatch{}
	return &this
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *Model5GLanParametersPatch) GetGpsis() map[string]string {
	if o == nil || IsNil(o.Gpsis) {
		var ret map[string]string
		return ret
	}
	return *o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParametersPatch) GetGpsisOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Gpsis) {
		return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *Model5GLanParametersPatch) HasGpsis() bool {
	if o != nil && !IsNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given map[string]string and assigns it to the Gpsis field.
func (o *Model5GLanParametersPatch) SetGpsis(v map[string]string) {
	o.Gpsis = &v
}

// GetAppDesps returns the AppDesps field value if set, zero value otherwise.
func (o *Model5GLanParametersPatch) GetAppDesps() map[string]AppDescriptorRm {
	if o == nil || IsNil(o.AppDesps) {
		var ret map[string]AppDescriptorRm
		return ret
	}
	return *o.AppDesps
}

// GetAppDespsOk returns a tuple with the AppDesps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParametersPatch) GetAppDespsOk() (*map[string]AppDescriptorRm, bool) {
	if o == nil || IsNil(o.AppDesps) {
		return nil, false
	}
	return o.AppDesps, true
}

// HasAppDesps returns a boolean if a field has been set.
func (o *Model5GLanParametersPatch) HasAppDesps() bool {
	if o != nil && !IsNil(o.AppDesps) {
		return true
	}

	return false
}

// SetAppDesps gets a reference to the given map[string]AppDescriptorRm and assigns it to the AppDesps field.
func (o *Model5GLanParametersPatch) SetAppDesps(v map[string]AppDescriptorRm) {
	o.AppDesps = &v
}

func (o Model5GLanParametersPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GLanParametersPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !IsNil(o.AppDesps) {
		toSerialize["appDesps"] = o.AppDesps
	}
	return toSerialize, nil
}

type NullableModel5GLanParametersPatch struct {
	value *Model5GLanParametersPatch
	isSet bool
}

func (v NullableModel5GLanParametersPatch) Get() *Model5GLanParametersPatch {
	return v.value
}

func (v *NullableModel5GLanParametersPatch) Set(val *Model5GLanParametersPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GLanParametersPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GLanParametersPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GLanParametersPatch(val *Model5GLanParametersPatch) *NullableModel5GLanParametersPatch {
	return &NullableModel5GLanParametersPatch{value: val, isSet: true}
}

func (v NullableModel5GLanParametersPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GLanParametersPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


