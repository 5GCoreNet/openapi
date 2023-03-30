/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the NefCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NefCond{}

// NefCond Subscription to a set of NF Instances (NEFs), identified by Event ID(s) provided by AF, S-NSSAI(s), AF Instance ID, Application Identifier, External Identifier, External Group Identifier, or domain name. 
type NefCond struct {
	ConditionType string `json:"conditionType"`
	AfEvents []AfEvent `json:"afEvents,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	PfdData *PfdData `json:"pfdData,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	ServedFqdnList []string `json:"servedFqdnList,omitempty"`
}

// NewNefCond instantiates a new NefCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefCond(conditionType string) *NefCond {
	this := NefCond{}
	this.ConditionType = conditionType
	return &this
}

// NewNefCondWithDefaults instantiates a new NefCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefCondWithDefaults() *NefCond {
	this := NefCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *NefCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *NefCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *NefCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetAfEvents returns the AfEvents field value if set, zero value otherwise.
func (o *NefCond) GetAfEvents() []AfEvent {
	if o == nil || IsNil(o.AfEvents) {
		var ret []AfEvent
		return ret
	}
	return o.AfEvents
}

// GetAfEventsOk returns a tuple with the AfEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetAfEventsOk() ([]AfEvent, bool) {
	if o == nil || IsNil(o.AfEvents) {
		return nil, false
	}
	return o.AfEvents, true
}

// HasAfEvents returns a boolean if a field has been set.
func (o *NefCond) HasAfEvents() bool {
	if o != nil && !IsNil(o.AfEvents) {
		return true
	}

	return false
}

// SetAfEvents gets a reference to the given []AfEvent and assigns it to the AfEvents field.
func (o *NefCond) SetAfEvents(v []AfEvent) {
	o.AfEvents = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NefCond) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NefCond) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NefCond) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetPfdData returns the PfdData field value if set, zero value otherwise.
func (o *NefCond) GetPfdData() PfdData {
	if o == nil || IsNil(o.PfdData) {
		var ret PfdData
		return ret
	}
	return *o.PfdData
}

// GetPfdDataOk returns a tuple with the PfdData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetPfdDataOk() (*PfdData, bool) {
	if o == nil || IsNil(o.PfdData) {
		return nil, false
	}
	return o.PfdData, true
}

// HasPfdData returns a boolean if a field has been set.
func (o *NefCond) HasPfdData() bool {
	if o != nil && !IsNil(o.PfdData) {
		return true
	}

	return false
}

// SetPfdData gets a reference to the given PfdData and assigns it to the PfdData field.
func (o *NefCond) SetPfdData(v PfdData) {
	o.PfdData = &v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *NefCond) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *NefCond) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *NefCond) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *NefCond) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *NefCond) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *NefCond) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetServedFqdnList returns the ServedFqdnList field value if set, zero value otherwise.
func (o *NefCond) GetServedFqdnList() []string {
	if o == nil || IsNil(o.ServedFqdnList) {
		var ret []string
		return ret
	}
	return o.ServedFqdnList
}

// GetServedFqdnListOk returns a tuple with the ServedFqdnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefCond) GetServedFqdnListOk() ([]string, bool) {
	if o == nil || IsNil(o.ServedFqdnList) {
		return nil, false
	}
	return o.ServedFqdnList, true
}

// HasServedFqdnList returns a boolean if a field has been set.
func (o *NefCond) HasServedFqdnList() bool {
	if o != nil && !IsNil(o.ServedFqdnList) {
		return true
	}

	return false
}

// SetServedFqdnList gets a reference to the given []string and assigns it to the ServedFqdnList field.
func (o *NefCond) SetServedFqdnList(v []string) {
	o.ServedFqdnList = v
}

func (o NefCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NefCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionType"] = o.ConditionType
	if !IsNil(o.AfEvents) {
		toSerialize["afEvents"] = o.AfEvents
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !IsNil(o.PfdData) {
		toSerialize["pfdData"] = o.PfdData
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.ExternalGroupIdentifiersRanges) {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if !IsNil(o.ServedFqdnList) {
		toSerialize["servedFqdnList"] = o.ServedFqdnList
	}
	return toSerialize, nil
}

type NullableNefCond struct {
	value *NefCond
	isSet bool
}

func (v NullableNefCond) Get() *NefCond {
	return v.value
}

func (v *NullableNefCond) Set(val *NefCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNefCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNefCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefCond(val *NefCond) *NullableNefCond {
	return &NullableNefCond{value: val, isSet: true}
}

func (v NullableNefCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


