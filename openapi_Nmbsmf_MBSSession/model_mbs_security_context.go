/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the MbsSecurityContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSecurityContext{}

// MbsSecurityContext struct for MbsSecurityContext
type MbsSecurityContext struct {
	// A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbsSecurityContext
	KeyList map[string]MbsKeyInfo `json:"keyList"`
}

// NewMbsSecurityContext instantiates a new MbsSecurityContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSecurityContext(keyList map[string]MbsKeyInfo) *MbsSecurityContext {
	this := MbsSecurityContext{}
	this.KeyList = keyList
	return &this
}

// NewMbsSecurityContextWithDefaults instantiates a new MbsSecurityContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSecurityContextWithDefaults() *MbsSecurityContext {
	this := MbsSecurityContext{}
	return &this
}

// GetKeyList returns the KeyList field value
func (o *MbsSecurityContext) GetKeyList() map[string]MbsKeyInfo {
	if o == nil {
		var ret map[string]MbsKeyInfo
		return ret
	}

	return o.KeyList
}

// GetKeyListOk returns a tuple with the KeyList field value
// and a boolean to check if the value has been set.
func (o *MbsSecurityContext) GetKeyListOk() (*map[string]MbsKeyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyList, true
}

// SetKeyList sets field value
func (o *MbsSecurityContext) SetKeyList(v map[string]MbsKeyInfo) {
	o.KeyList = v
}

func (o MbsSecurityContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSecurityContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keyList"] = o.KeyList
	return toSerialize, nil
}

type NullableMbsSecurityContext struct {
	value *MbsSecurityContext
	isSet bool
}

func (v NullableMbsSecurityContext) Get() *MbsSecurityContext {
	return v.value
}

func (v *NullableMbsSecurityContext) Set(val *MbsSecurityContext) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSecurityContext) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSecurityContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSecurityContext(val *MbsSecurityContext) *NullableMbsSecurityContext {
	return &NullableMbsSecurityContext{value: val, isSet: true}
}

func (v NullableMbsSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSecurityContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


