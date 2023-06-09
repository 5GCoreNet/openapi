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

// checks if the EventFilter1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFilter1{}

// EventFilter1 Represents event filter information for an event.
type EventFilter1 struct {
	Gpsis         []string                    `json:"gpsis,omitempty"`
	Supis         []string                    `json:"supis,omitempty"`
	ExterGroupIds []string                    `json:"exterGroupIds,omitempty"`
	InterGroupIds []string                    `json:"interGroupIds,omitempty"`
	AnyUeInd      *bool                       `json:"anyUeInd,omitempty"`
	AppIds        []string                    `json:"appIds,omitempty"`
	LocArea       *LocationArea5G             `json:"locArea,omitempty"`
	CollAttrs     []CollectiveBehaviourFilter `json:"collAttrs,omitempty"`
}

// NewEventFilter1 instantiates a new EventFilter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter1() *EventFilter1 {
	this := EventFilter1{}
	return &this
}

// NewEventFilter1WithDefaults instantiates a new EventFilter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilter1WithDefaults() *EventFilter1 {
	this := EventFilter1{}
	return &this
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *EventFilter1) GetGpsis() []string {
	if o == nil || IsNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetGpsisOk() ([]string, bool) {
	if o == nil || IsNil(o.Gpsis) {
		return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *EventFilter1) HasGpsis() bool {
	if o != nil && !IsNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *EventFilter1) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *EventFilter1) GetSupis() []string {
	if o == nil || IsNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.Supis) {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *EventFilter1) HasSupis() bool {
	if o != nil && !IsNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *EventFilter1) SetSupis(v []string) {
	o.Supis = v
}

// GetExterGroupIds returns the ExterGroupIds field value if set, zero value otherwise.
func (o *EventFilter1) GetExterGroupIds() []string {
	if o == nil || IsNil(o.ExterGroupIds) {
		var ret []string
		return ret
	}
	return o.ExterGroupIds
}

// GetExterGroupIdsOk returns a tuple with the ExterGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetExterGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExterGroupIds) {
		return nil, false
	}
	return o.ExterGroupIds, true
}

// HasExterGroupIds returns a boolean if a field has been set.
func (o *EventFilter1) HasExterGroupIds() bool {
	if o != nil && !IsNil(o.ExterGroupIds) {
		return true
	}

	return false
}

// SetExterGroupIds gets a reference to the given []string and assigns it to the ExterGroupIds field.
func (o *EventFilter1) SetExterGroupIds(v []string) {
	o.ExterGroupIds = v
}

// GetInterGroupIds returns the InterGroupIds field value if set, zero value otherwise.
func (o *EventFilter1) GetInterGroupIds() []string {
	if o == nil || IsNil(o.InterGroupIds) {
		var ret []string
		return ret
	}
	return o.InterGroupIds
}

// GetInterGroupIdsOk returns a tuple with the InterGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetInterGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InterGroupIds) {
		return nil, false
	}
	return o.InterGroupIds, true
}

// HasInterGroupIds returns a boolean if a field has been set.
func (o *EventFilter1) HasInterGroupIds() bool {
	if o != nil && !IsNil(o.InterGroupIds) {
		return true
	}

	return false
}

// SetInterGroupIds gets a reference to the given []string and assigns it to the InterGroupIds field.
func (o *EventFilter1) SetInterGroupIds(v []string) {
	o.InterGroupIds = v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *EventFilter1) GetAnyUeInd() bool {
	if o == nil || IsNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AnyUeInd) {
		return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *EventFilter1) HasAnyUeInd() bool {
	if o != nil && !IsNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *EventFilter1) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventFilter1) GetAppIds() []string {
	if o == nil || IsNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventFilter1) HasAppIds() bool {
	if o != nil && !IsNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventFilter1) SetAppIds(v []string) {
	o.AppIds = v
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *EventFilter1) GetLocArea() LocationArea5G {
	if o == nil || IsNil(o.LocArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.LocArea) {
		return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *EventFilter1) HasLocArea() bool {
	if o != nil && !IsNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given LocationArea5G and assigns it to the LocArea field.
func (o *EventFilter1) SetLocArea(v LocationArea5G) {
	o.LocArea = &v
}

// GetCollAttrs returns the CollAttrs field value if set, zero value otherwise.
func (o *EventFilter1) GetCollAttrs() []CollectiveBehaviourFilter {
	if o == nil || IsNil(o.CollAttrs) {
		var ret []CollectiveBehaviourFilter
		return ret
	}
	return o.CollAttrs
}

// GetCollAttrsOk returns a tuple with the CollAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter1) GetCollAttrsOk() ([]CollectiveBehaviourFilter, bool) {
	if o == nil || IsNil(o.CollAttrs) {
		return nil, false
	}
	return o.CollAttrs, true
}

// HasCollAttrs returns a boolean if a field has been set.
func (o *EventFilter1) HasCollAttrs() bool {
	if o != nil && !IsNil(o.CollAttrs) {
		return true
	}

	return false
}

// SetCollAttrs gets a reference to the given []CollectiveBehaviourFilter and assigns it to the CollAttrs field.
func (o *EventFilter1) SetCollAttrs(v []CollectiveBehaviourFilter) {
	o.CollAttrs = v
}

func (o EventFilter1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFilter1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !IsNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !IsNil(o.ExterGroupIds) {
		toSerialize["exterGroupIds"] = o.ExterGroupIds
	}
	if !IsNil(o.InterGroupIds) {
		toSerialize["interGroupIds"] = o.InterGroupIds
	}
	if !IsNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !IsNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !IsNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !IsNil(o.CollAttrs) {
		toSerialize["collAttrs"] = o.CollAttrs
	}
	return toSerialize, nil
}

type NullableEventFilter1 struct {
	value *EventFilter1
	isSet bool
}

func (v NullableEventFilter1) Get() *EventFilter1 {
	return v.value
}

func (v *NullableEventFilter1) Set(val *EventFilter1) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter1) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter1(val *EventFilter1) *NullableEventFilter1 {
	return &NullableEventFilter1{value: val, isSet: true}
}

func (v NullableEventFilter1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilter1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
