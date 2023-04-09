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

// checks if the GroupIdentifiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupIdentifiers{}

// GroupIdentifiers External or Internal Group Identifier with a list of group members
type GroupIdentifiers struct {
	ExtGroupId *string `json:"extGroupId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	IntGroupId   *string  `json:"intGroupId,omitempty"`
	UeIdList     []UeId   `json:"ueIdList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
}

// NewGroupIdentifiers instantiates a new GroupIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupIdentifiers() *GroupIdentifiers {
	this := GroupIdentifiers{}
	return &this
}

// NewGroupIdentifiersWithDefaults instantiates a new GroupIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupIdentifiersWithDefaults() *GroupIdentifiers {
	this := GroupIdentifiers{}
	return &this
}

// GetExtGroupId returns the ExtGroupId field value if set, zero value otherwise.
func (o *GroupIdentifiers) GetExtGroupId() string {
	if o == nil || IsNil(o.ExtGroupId) {
		var ret string
		return ret
	}
	return *o.ExtGroupId
}

// GetExtGroupIdOk returns a tuple with the ExtGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupIdentifiers) GetExtGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExtGroupId) {
		return nil, false
	}
	return o.ExtGroupId, true
}

// HasExtGroupId returns a boolean if a field has been set.
func (o *GroupIdentifiers) HasExtGroupId() bool {
	if o != nil && !IsNil(o.ExtGroupId) {
		return true
	}

	return false
}

// SetExtGroupId gets a reference to the given string and assigns it to the ExtGroupId field.
func (o *GroupIdentifiers) SetExtGroupId(v string) {
	o.ExtGroupId = &v
}

// GetIntGroupId returns the IntGroupId field value if set, zero value otherwise.
func (o *GroupIdentifiers) GetIntGroupId() string {
	if o == nil || IsNil(o.IntGroupId) {
		var ret string
		return ret
	}
	return *o.IntGroupId
}

// GetIntGroupIdOk returns a tuple with the IntGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupIdentifiers) GetIntGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntGroupId) {
		return nil, false
	}
	return o.IntGroupId, true
}

// HasIntGroupId returns a boolean if a field has been set.
func (o *GroupIdentifiers) HasIntGroupId() bool {
	if o != nil && !IsNil(o.IntGroupId) {
		return true
	}

	return false
}

// SetIntGroupId gets a reference to the given string and assigns it to the IntGroupId field.
func (o *GroupIdentifiers) SetIntGroupId(v string) {
	o.IntGroupId = &v
}

// GetUeIdList returns the UeIdList field value if set, zero value otherwise.
func (o *GroupIdentifiers) GetUeIdList() []UeId {
	if o == nil || IsNil(o.UeIdList) {
		var ret []UeId
		return ret
	}
	return o.UeIdList
}

// GetUeIdListOk returns a tuple with the UeIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupIdentifiers) GetUeIdListOk() ([]UeId, bool) {
	if o == nil || IsNil(o.UeIdList) {
		return nil, false
	}
	return o.UeIdList, true
}

// HasUeIdList returns a boolean if a field has been set.
func (o *GroupIdentifiers) HasUeIdList() bool {
	if o != nil && !IsNil(o.UeIdList) {
		return true
	}

	return false
}

// SetUeIdList gets a reference to the given []UeId and assigns it to the UeIdList field.
func (o *GroupIdentifiers) SetUeIdList(v []UeId) {
	o.UeIdList = v
}

// GetAllowedAfIds returns the AllowedAfIds field value if set, zero value otherwise.
func (o *GroupIdentifiers) GetAllowedAfIds() []string {
	if o == nil || IsNil(o.AllowedAfIds) {
		var ret []string
		return ret
	}
	return o.AllowedAfIds
}

// GetAllowedAfIdsOk returns a tuple with the AllowedAfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupIdentifiers) GetAllowedAfIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedAfIds) {
		return nil, false
	}
	return o.AllowedAfIds, true
}

// HasAllowedAfIds returns a boolean if a field has been set.
func (o *GroupIdentifiers) HasAllowedAfIds() bool {
	if o != nil && !IsNil(o.AllowedAfIds) {
		return true
	}

	return false
}

// SetAllowedAfIds gets a reference to the given []string and assigns it to the AllowedAfIds field.
func (o *GroupIdentifiers) SetAllowedAfIds(v []string) {
	o.AllowedAfIds = v
}

func (o GroupIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupIdentifiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtGroupId) {
		toSerialize["extGroupId"] = o.ExtGroupId
	}
	if !IsNil(o.IntGroupId) {
		toSerialize["intGroupId"] = o.IntGroupId
	}
	if !IsNil(o.UeIdList) {
		toSerialize["ueIdList"] = o.UeIdList
	}
	if !IsNil(o.AllowedAfIds) {
		toSerialize["allowedAfIds"] = o.AllowedAfIds
	}
	return toSerialize, nil
}

type NullableGroupIdentifiers struct {
	value *GroupIdentifiers
	isSet bool
}

func (v NullableGroupIdentifiers) Get() *GroupIdentifiers {
	return v.value
}

func (v *NullableGroupIdentifiers) Set(val *GroupIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupIdentifiers(val *GroupIdentifiers) *NullableGroupIdentifiers {
	return &NullableGroupIdentifiers{value: val, isSet: true}
}

func (v NullableGroupIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
