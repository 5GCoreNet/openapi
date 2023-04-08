/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// checks if the MbsServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsServiceInfo{}

// MbsServiceInfo Represent MBS Service Information.
type MbsServiceInfo struct {
	MbsMediaComps map[string]MbsMediaCompRm `json:"mbsMediaComps"`
	MbsSdfResPrio *ReservPriority `json:"mbsSdfResPrio,omitempty"`
	// Contains an AF application identifier.
	AfAppId *string `json:"afAppId,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MbsSessionAmbr *string `json:"mbsSessionAmbr,omitempty"`
}

// NewMbsServiceInfo instantiates a new MbsServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsServiceInfo(mbsMediaComps map[string]MbsMediaCompRm) *MbsServiceInfo {
	this := MbsServiceInfo{}
	this.MbsMediaComps = mbsMediaComps
	return &this
}

// NewMbsServiceInfoWithDefaults instantiates a new MbsServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsServiceInfoWithDefaults() *MbsServiceInfo {
	this := MbsServiceInfo{}
	return &this
}

// GetMbsMediaComps returns the MbsMediaComps field value
func (o *MbsServiceInfo) GetMbsMediaComps() map[string]MbsMediaCompRm {
	if o == nil {
		var ret map[string]MbsMediaCompRm
		return ret
	}

	return o.MbsMediaComps
}

// GetMbsMediaCompsOk returns a tuple with the MbsMediaComps field value
// and a boolean to check if the value has been set.
func (o *MbsServiceInfo) GetMbsMediaCompsOk() (*map[string]MbsMediaCompRm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsMediaComps, true
}

// SetMbsMediaComps sets field value
func (o *MbsServiceInfo) SetMbsMediaComps(v map[string]MbsMediaCompRm) {
	o.MbsMediaComps = v
}

// GetMbsSdfResPrio returns the MbsSdfResPrio field value if set, zero value otherwise.
func (o *MbsServiceInfo) GetMbsSdfResPrio() ReservPriority {
	if o == nil || isNil(o.MbsSdfResPrio) {
		var ret ReservPriority
		return ret
	}
	return *o.MbsSdfResPrio
}

// GetMbsSdfResPrioOk returns a tuple with the MbsSdfResPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsServiceInfo) GetMbsSdfResPrioOk() (*ReservPriority, bool) {
	if o == nil || isNil(o.MbsSdfResPrio) {
		return nil, false
	}
	return o.MbsSdfResPrio, true
}

// HasMbsSdfResPrio returns a boolean if a field has been set.
func (o *MbsServiceInfo) HasMbsSdfResPrio() bool {
	if o != nil && !isNil(o.MbsSdfResPrio) {
		return true
	}

	return false
}

// SetMbsSdfResPrio gets a reference to the given ReservPriority and assigns it to the MbsSdfResPrio field.
func (o *MbsServiceInfo) SetMbsSdfResPrio(v ReservPriority) {
	o.MbsSdfResPrio = &v
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *MbsServiceInfo) GetAfAppId() string {
	if o == nil || isNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsServiceInfo) GetAfAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *MbsServiceInfo) HasAfAppId() bool {
	if o != nil && !isNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *MbsServiceInfo) SetAfAppId(v string) {
	o.AfAppId = &v
}

// GetMbsSessionAmbr returns the MbsSessionAmbr field value if set, zero value otherwise.
func (o *MbsServiceInfo) GetMbsSessionAmbr() string {
	if o == nil || isNil(o.MbsSessionAmbr) {
		var ret string
		return ret
	}
	return *o.MbsSessionAmbr
}

// GetMbsSessionAmbrOk returns a tuple with the MbsSessionAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsServiceInfo) GetMbsSessionAmbrOk() (*string, bool) {
	if o == nil || isNil(o.MbsSessionAmbr) {
		return nil, false
	}
	return o.MbsSessionAmbr, true
}

// HasMbsSessionAmbr returns a boolean if a field has been set.
func (o *MbsServiceInfo) HasMbsSessionAmbr() bool {
	if o != nil && !isNil(o.MbsSessionAmbr) {
		return true
	}

	return false
}

// SetMbsSessionAmbr gets a reference to the given string and assigns it to the MbsSessionAmbr field.
func (o *MbsServiceInfo) SetMbsSessionAmbr(v string) {
	o.MbsSessionAmbr = &v
}

func (o MbsServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsMediaComps"] = o.MbsMediaComps
	if !isNil(o.MbsSdfResPrio) {
		toSerialize["mbsSdfResPrio"] = o.MbsSdfResPrio
	}
	if !isNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	if !isNil(o.MbsSessionAmbr) {
		toSerialize["mbsSessionAmbr"] = o.MbsSessionAmbr
	}
	return toSerialize, nil
}

type NullableMbsServiceInfo struct {
	value *MbsServiceInfo
	isSet bool
}

func (v NullableMbsServiceInfo) Get() *MbsServiceInfo {
	return v.value
}

func (v *NullableMbsServiceInfo) Set(val *MbsServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceInfo(val *MbsServiceInfo) *NullableMbsServiceInfo {
	return &NullableMbsServiceInfo{value: val, isSet: true}
}

func (v NullableMbsServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


