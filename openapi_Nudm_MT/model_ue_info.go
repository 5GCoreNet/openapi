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

// checks if the UeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeInfo{}

// UeInfo Represents UE information.
type UeInfo struct {
	TadsInfo *UeContextInfo `json:"tadsInfo,omitempty"`
	UserState *Model5GsUserState `json:"userState,omitempty"`
	Var5gSrvccInfo *Model5GSrvccInfo `json:"5gSrvccInfo,omitempty"`
}

// NewUeInfo instantiates a new UeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeInfo() *UeInfo {
	this := UeInfo{}
	return &this
}

// NewUeInfoWithDefaults instantiates a new UeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeInfoWithDefaults() *UeInfo {
	this := UeInfo{}
	return &this
}

// GetTadsInfo returns the TadsInfo field value if set, zero value otherwise.
func (o *UeInfo) GetTadsInfo() UeContextInfo {
	if o == nil || IsNil(o.TadsInfo) {
		var ret UeContextInfo
		return ret
	}
	return *o.TadsInfo
}

// GetTadsInfoOk returns a tuple with the TadsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInfo) GetTadsInfoOk() (*UeContextInfo, bool) {
	if o == nil || IsNil(o.TadsInfo) {
		return nil, false
	}
	return o.TadsInfo, true
}

// HasTadsInfo returns a boolean if a field has been set.
func (o *UeInfo) HasTadsInfo() bool {
	if o != nil && !IsNil(o.TadsInfo) {
		return true
	}

	return false
}

// SetTadsInfo gets a reference to the given UeContextInfo and assigns it to the TadsInfo field.
func (o *UeInfo) SetTadsInfo(v UeContextInfo) {
	o.TadsInfo = &v
}

// GetUserState returns the UserState field value if set, zero value otherwise.
func (o *UeInfo) GetUserState() Model5GsUserState {
	if o == nil || IsNil(o.UserState) {
		var ret Model5GsUserState
		return ret
	}
	return *o.UserState
}

// GetUserStateOk returns a tuple with the UserState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInfo) GetUserStateOk() (*Model5GsUserState, bool) {
	if o == nil || IsNil(o.UserState) {
		return nil, false
	}
	return o.UserState, true
}

// HasUserState returns a boolean if a field has been set.
func (o *UeInfo) HasUserState() bool {
	if o != nil && !IsNil(o.UserState) {
		return true
	}

	return false
}

// SetUserState gets a reference to the given Model5GsUserState and assigns it to the UserState field.
func (o *UeInfo) SetUserState(v Model5GsUserState) {
	o.UserState = &v
}

// GetVar5gSrvccInfo returns the Var5gSrvccInfo field value if set, zero value otherwise.
func (o *UeInfo) GetVar5gSrvccInfo() Model5GSrvccInfo {
	if o == nil || IsNil(o.Var5gSrvccInfo) {
		var ret Model5GSrvccInfo
		return ret
	}
	return *o.Var5gSrvccInfo
}

// GetVar5gSrvccInfoOk returns a tuple with the Var5gSrvccInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInfo) GetVar5gSrvccInfoOk() (*Model5GSrvccInfo, bool) {
	if o == nil || IsNil(o.Var5gSrvccInfo) {
		return nil, false
	}
	return o.Var5gSrvccInfo, true
}

// HasVar5gSrvccInfo returns a boolean if a field has been set.
func (o *UeInfo) HasVar5gSrvccInfo() bool {
	if o != nil && !IsNil(o.Var5gSrvccInfo) {
		return true
	}

	return false
}

// SetVar5gSrvccInfo gets a reference to the given Model5GSrvccInfo and assigns it to the Var5gSrvccInfo field.
func (o *UeInfo) SetVar5gSrvccInfo(v Model5GSrvccInfo) {
	o.Var5gSrvccInfo = &v
}

func (o UeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TadsInfo) {
		toSerialize["tadsInfo"] = o.TadsInfo
	}
	if !IsNil(o.UserState) {
		toSerialize["userState"] = o.UserState
	}
	if !IsNil(o.Var5gSrvccInfo) {
		toSerialize["5gSrvccInfo"] = o.Var5gSrvccInfo
	}
	return toSerialize, nil
}

type NullableUeInfo struct {
	value *UeInfo
	isSet bool
}

func (v NullableUeInfo) Get() *UeInfo {
	return v.value
}

func (v *NullableUeInfo) Set(val *UeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeInfo(val *UeInfo) *NullableUeInfo {
	return &NullableUeInfo{value: val, isSet: true}
}

func (v NullableUeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


