/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the ChfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChfInfo{}

// ChfInfo Information of a CHF NF Instance
type ChfInfo struct {
	SupiRangeList []SupiRange `json:"supiRangeList,omitempty"`
	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty"`
	PlmnRangeList []PlmnRange `json:"plmnRangeList,omitempty"`
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PrimaryChfInstance *string `json:"primaryChfInstance,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SecondaryChfInstance *string `json:"secondaryChfInstance,omitempty"`
}

// NewChfInfo instantiates a new ChfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChfInfo() *ChfInfo {
	this := ChfInfo{}
	return &this
}

// NewChfInfoWithDefaults instantiates a new ChfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChfInfoWithDefaults() *ChfInfo {
	this := ChfInfo{}
	return &this
}

// GetSupiRangeList returns the SupiRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetSupiRangeList() []SupiRange {
	if o == nil || isNil(o.SupiRangeList) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRangeList
}

// GetSupiRangeListOk returns a tuple with the SupiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetSupiRangeListOk() ([]SupiRange, bool) {
	if o == nil || isNil(o.SupiRangeList) {
		return nil, false
	}
	return o.SupiRangeList, true
}

// HasSupiRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasSupiRangeList() bool {
	if o != nil && !isNil(o.SupiRangeList) {
		return true
	}

	return false
}

// SetSupiRangeList gets a reference to the given []SupiRange and assigns it to the SupiRangeList field.
func (o *ChfInfo) SetSupiRangeList(v []SupiRange) {
	o.SupiRangeList = v
}

// GetGpsiRangeList returns the GpsiRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetGpsiRangeList() []IdentityRange {
	if o == nil || isNil(o.GpsiRangeList) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRangeList
}

// GetGpsiRangeListOk returns a tuple with the GpsiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetGpsiRangeListOk() ([]IdentityRange, bool) {
	if o == nil || isNil(o.GpsiRangeList) {
		return nil, false
	}
	return o.GpsiRangeList, true
}

// HasGpsiRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasGpsiRangeList() bool {
	if o != nil && !isNil(o.GpsiRangeList) {
		return true
	}

	return false
}

// SetGpsiRangeList gets a reference to the given []IdentityRange and assigns it to the GpsiRangeList field.
func (o *ChfInfo) SetGpsiRangeList(v []IdentityRange) {
	o.GpsiRangeList = v
}

// GetPlmnRangeList returns the PlmnRangeList field value if set, zero value otherwise.
func (o *ChfInfo) GetPlmnRangeList() []PlmnRange {
	if o == nil || isNil(o.PlmnRangeList) {
		var ret []PlmnRange
		return ret
	}
	return o.PlmnRangeList
}

// GetPlmnRangeListOk returns a tuple with the PlmnRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetPlmnRangeListOk() ([]PlmnRange, bool) {
	if o == nil || isNil(o.PlmnRangeList) {
		return nil, false
	}
	return o.PlmnRangeList, true
}

// HasPlmnRangeList returns a boolean if a field has been set.
func (o *ChfInfo) HasPlmnRangeList() bool {
	if o != nil && !isNil(o.PlmnRangeList) {
		return true
	}

	return false
}

// SetPlmnRangeList gets a reference to the given []PlmnRange and assigns it to the PlmnRangeList field.
func (o *ChfInfo) SetPlmnRangeList(v []PlmnRange) {
	o.PlmnRangeList = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ChfInfo) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ChfInfo) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ChfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPrimaryChfInstance returns the PrimaryChfInstance field value if set, zero value otherwise.
func (o *ChfInfo) GetPrimaryChfInstance() string {
	if o == nil || isNil(o.PrimaryChfInstance) {
		var ret string
		return ret
	}
	return *o.PrimaryChfInstance
}

// GetPrimaryChfInstanceOk returns a tuple with the PrimaryChfInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetPrimaryChfInstanceOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryChfInstance) {
		return nil, false
	}
	return o.PrimaryChfInstance, true
}

// HasPrimaryChfInstance returns a boolean if a field has been set.
func (o *ChfInfo) HasPrimaryChfInstance() bool {
	if o != nil && !isNil(o.PrimaryChfInstance) {
		return true
	}

	return false
}

// SetPrimaryChfInstance gets a reference to the given string and assigns it to the PrimaryChfInstance field.
func (o *ChfInfo) SetPrimaryChfInstance(v string) {
	o.PrimaryChfInstance = &v
}

// GetSecondaryChfInstance returns the SecondaryChfInstance field value if set, zero value otherwise.
func (o *ChfInfo) GetSecondaryChfInstance() string {
	if o == nil || isNil(o.SecondaryChfInstance) {
		var ret string
		return ret
	}
	return *o.SecondaryChfInstance
}

// GetSecondaryChfInstanceOk returns a tuple with the SecondaryChfInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChfInfo) GetSecondaryChfInstanceOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryChfInstance) {
		return nil, false
	}
	return o.SecondaryChfInstance, true
}

// HasSecondaryChfInstance returns a boolean if a field has been set.
func (o *ChfInfo) HasSecondaryChfInstance() bool {
	if o != nil && !isNil(o.SecondaryChfInstance) {
		return true
	}

	return false
}

// SetSecondaryChfInstance gets a reference to the given string and assigns it to the SecondaryChfInstance field.
func (o *ChfInfo) SetSecondaryChfInstance(v string) {
	o.SecondaryChfInstance = &v
}

func (o ChfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupiRangeList) {
		toSerialize["supiRangeList"] = o.SupiRangeList
	}
	if !isNil(o.GpsiRangeList) {
		toSerialize["gpsiRangeList"] = o.GpsiRangeList
	}
	if !isNil(o.PlmnRangeList) {
		toSerialize["plmnRangeList"] = o.PlmnRangeList
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.PrimaryChfInstance) {
		toSerialize["primaryChfInstance"] = o.PrimaryChfInstance
	}
	if !isNil(o.SecondaryChfInstance) {
		toSerialize["secondaryChfInstance"] = o.SecondaryChfInstance
	}
	return toSerialize, nil
}

type NullableChfInfo struct {
	value *ChfInfo
	isSet bool
}

func (v NullableChfInfo) Get() *ChfInfo {
	return v.value
}

func (v *NullableChfInfo) Set(val *ChfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChfInfo(val *ChfInfo) *NullableChfInfo {
	return &NullableChfInfo{value: val, isSet: true}
}

func (v NullableChfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


