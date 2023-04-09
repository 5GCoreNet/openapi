/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// checks if the UeRequestedValueRep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeRequestedValueRep{}

// UeRequestedValueRep Contains the current applicable values corresponding to the policy control request triggers.
type UeRequestedValueRep struct {
	UserLoc *UserLocation `json:"userLoc,omitempty"`
	// Contains the UE presence statuses for tracking areas. The praId attribute within the PresenceInfo data type is the key of the map.
	PraStatuses  *map[string]PresenceInfo `json:"praStatuses,omitempty"`
	PlmnId       *PlmnIdNid               `json:"plmnId,omitempty"`
	ConnectState *CmState                 `json:"connectState,omitempty"`
}

// NewUeRequestedValueRep instantiates a new UeRequestedValueRep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeRequestedValueRep() *UeRequestedValueRep {
	this := UeRequestedValueRep{}
	return &this
}

// NewUeRequestedValueRepWithDefaults instantiates a new UeRequestedValueRep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeRequestedValueRepWithDefaults() *UeRequestedValueRep {
	this := UeRequestedValueRep{}
	return &this
}

// GetUserLoc returns the UserLoc field value if set, zero value otherwise.
func (o *UeRequestedValueRep) GetUserLoc() UserLocation {
	if o == nil || IsNil(o.UserLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UserLoc
}

// GetUserLocOk returns a tuple with the UserLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRequestedValueRep) GetUserLocOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLoc) {
		return nil, false
	}
	return o.UserLoc, true
}

// HasUserLoc returns a boolean if a field has been set.
func (o *UeRequestedValueRep) HasUserLoc() bool {
	if o != nil && !IsNil(o.UserLoc) {
		return true
	}

	return false
}

// SetUserLoc gets a reference to the given UserLocation and assigns it to the UserLoc field.
func (o *UeRequestedValueRep) SetUserLoc(v UserLocation) {
	o.UserLoc = &v
}

// GetPraStatuses returns the PraStatuses field value if set, zero value otherwise.
func (o *UeRequestedValueRep) GetPraStatuses() map[string]PresenceInfo {
	if o == nil || IsNil(o.PraStatuses) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PraStatuses
}

// GetPraStatusesOk returns a tuple with the PraStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRequestedValueRep) GetPraStatusesOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.PraStatuses) {
		return nil, false
	}
	return o.PraStatuses, true
}

// HasPraStatuses returns a boolean if a field has been set.
func (o *UeRequestedValueRep) HasPraStatuses() bool {
	if o != nil && !IsNil(o.PraStatuses) {
		return true
	}

	return false
}

// SetPraStatuses gets a reference to the given map[string]PresenceInfo and assigns it to the PraStatuses field.
func (o *UeRequestedValueRep) SetPraStatuses(v map[string]PresenceInfo) {
	o.PraStatuses = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *UeRequestedValueRep) GetPlmnId() PlmnIdNid {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnIdNid
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRequestedValueRep) GetPlmnIdOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *UeRequestedValueRep) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnIdNid and assigns it to the PlmnId field.
func (o *UeRequestedValueRep) SetPlmnId(v PlmnIdNid) {
	o.PlmnId = &v
}

// GetConnectState returns the ConnectState field value if set, zero value otherwise.
func (o *UeRequestedValueRep) GetConnectState() CmState {
	if o == nil || IsNil(o.ConnectState) {
		var ret CmState
		return ret
	}
	return *o.ConnectState
}

// GetConnectStateOk returns a tuple with the ConnectState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRequestedValueRep) GetConnectStateOk() (*CmState, bool) {
	if o == nil || IsNil(o.ConnectState) {
		return nil, false
	}
	return o.ConnectState, true
}

// HasConnectState returns a boolean if a field has been set.
func (o *UeRequestedValueRep) HasConnectState() bool {
	if o != nil && !IsNil(o.ConnectState) {
		return true
	}

	return false
}

// SetConnectState gets a reference to the given CmState and assigns it to the ConnectState field.
func (o *UeRequestedValueRep) SetConnectState(v CmState) {
	o.ConnectState = &v
}

func (o UeRequestedValueRep) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeRequestedValueRep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLoc) {
		toSerialize["userLoc"] = o.UserLoc
	}
	if !IsNil(o.PraStatuses) {
		toSerialize["praStatuses"] = o.PraStatuses
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.ConnectState) {
		toSerialize["connectState"] = o.ConnectState
	}
	return toSerialize, nil
}

type NullableUeRequestedValueRep struct {
	value *UeRequestedValueRep
	isSet bool
}

func (v NullableUeRequestedValueRep) Get() *UeRequestedValueRep {
	return v.value
}

func (v *NullableUeRequestedValueRep) Set(val *UeRequestedValueRep) {
	v.value = val
	v.isSet = true
}

func (v NullableUeRequestedValueRep) IsSet() bool {
	return v.isSet
}

func (v *NullableUeRequestedValueRep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeRequestedValueRep(val *UeRequestedValueRep) *NullableUeRequestedValueRep {
	return &NullableUeRequestedValueRep{value: val, isSet: true}
}

func (v NullableUeRequestedValueRep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeRequestedValueRep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
