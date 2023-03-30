/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the Model5MbsAuthorizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5MbsAuthorizationInfo{}

// Model5MbsAuthorizationInfo struct for Model5MbsAuthorizationInfo
type Model5MbsAuthorizationInfo struct {
	Var5mbsSessionIds []MbsSessionId `json:"5mbsSessionIds,omitempty"`
}

// NewModel5MbsAuthorizationInfo instantiates a new Model5MbsAuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5MbsAuthorizationInfo() *Model5MbsAuthorizationInfo {
	this := Model5MbsAuthorizationInfo{}
	return &this
}

// NewModel5MbsAuthorizationInfoWithDefaults instantiates a new Model5MbsAuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5MbsAuthorizationInfoWithDefaults() *Model5MbsAuthorizationInfo {
	this := Model5MbsAuthorizationInfo{}
	return &this
}

// GetVar5mbsSessionIds returns the Var5mbsSessionIds field value if set, zero value otherwise.
func (o *Model5MbsAuthorizationInfo) GetVar5mbsSessionIds() []MbsSessionId {
	if o == nil || IsNil(o.Var5mbsSessionIds) {
		var ret []MbsSessionId
		return ret
	}
	return o.Var5mbsSessionIds
}

// GetVar5mbsSessionIdsOk returns a tuple with the Var5mbsSessionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5MbsAuthorizationInfo) GetVar5mbsSessionIdsOk() ([]MbsSessionId, bool) {
	if o == nil || IsNil(o.Var5mbsSessionIds) {
		return nil, false
	}
	return o.Var5mbsSessionIds, true
}

// HasVar5mbsSessionIds returns a boolean if a field has been set.
func (o *Model5MbsAuthorizationInfo) HasVar5mbsSessionIds() bool {
	if o != nil && !IsNil(o.Var5mbsSessionIds) {
		return true
	}

	return false
}

// SetVar5mbsSessionIds gets a reference to the given []MbsSessionId and assigns it to the Var5mbsSessionIds field.
func (o *Model5MbsAuthorizationInfo) SetVar5mbsSessionIds(v []MbsSessionId) {
	o.Var5mbsSessionIds = v
}

func (o Model5MbsAuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5MbsAuthorizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var5mbsSessionIds) {
		toSerialize["5mbsSessionIds"] = o.Var5mbsSessionIds
	}
	return toSerialize, nil
}

type NullableModel5MbsAuthorizationInfo struct {
	value *Model5MbsAuthorizationInfo
	isSet bool
}

func (v NullableModel5MbsAuthorizationInfo) Get() *Model5MbsAuthorizationInfo {
	return v.value
}

func (v *NullableModel5MbsAuthorizationInfo) Set(val *Model5MbsAuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5MbsAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5MbsAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5MbsAuthorizationInfo(val *Model5MbsAuthorizationInfo) *NullableModel5MbsAuthorizationInfo {
	return &NullableModel5MbsAuthorizationInfo{value: val, isSet: true}
}

func (v NullableModel5MbsAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5MbsAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


