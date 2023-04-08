/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
)

// checks if the SelectedC2CommModeNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectedC2CommModeNotif{}

// SelectedC2CommModeNotif Represents information on the C2 Communication Mode selected by a UAS (i.e. pair of UAV and UAV-C). 
type SelectedC2CommModeNotif struct {
	UasId UasId `json:"uasId"`
	SelPrimaryC2CommMode C2CommMode `json:"selPrimaryC2CommMode"`
	SelSecondaryC2CommMode *C2CommMode `json:"selSecondaryC2CommMode,omitempty"`
}

// NewSelectedC2CommModeNotif instantiates a new SelectedC2CommModeNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedC2CommModeNotif(uasId UasId, selPrimaryC2CommMode C2CommMode) *SelectedC2CommModeNotif {
	this := SelectedC2CommModeNotif{}
	this.UasId = uasId
	this.SelPrimaryC2CommMode = selPrimaryC2CommMode
	return &this
}

// NewSelectedC2CommModeNotifWithDefaults instantiates a new SelectedC2CommModeNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedC2CommModeNotifWithDefaults() *SelectedC2CommModeNotif {
	this := SelectedC2CommModeNotif{}
	return &this
}

// GetUasId returns the UasId field value
func (o *SelectedC2CommModeNotif) GetUasId() UasId {
	if o == nil {
		var ret UasId
		return ret
	}

	return o.UasId
}

// GetUasIdOk returns a tuple with the UasId field value
// and a boolean to check if the value has been set.
func (o *SelectedC2CommModeNotif) GetUasIdOk() (*UasId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UasId, true
}

// SetUasId sets field value
func (o *SelectedC2CommModeNotif) SetUasId(v UasId) {
	o.UasId = v
}

// GetSelPrimaryC2CommMode returns the SelPrimaryC2CommMode field value
func (o *SelectedC2CommModeNotif) GetSelPrimaryC2CommMode() C2CommMode {
	if o == nil {
		var ret C2CommMode
		return ret
	}

	return o.SelPrimaryC2CommMode
}

// GetSelPrimaryC2CommModeOk returns a tuple with the SelPrimaryC2CommMode field value
// and a boolean to check if the value has been set.
func (o *SelectedC2CommModeNotif) GetSelPrimaryC2CommModeOk() (*C2CommMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelPrimaryC2CommMode, true
}

// SetSelPrimaryC2CommMode sets field value
func (o *SelectedC2CommModeNotif) SetSelPrimaryC2CommMode(v C2CommMode) {
	o.SelPrimaryC2CommMode = v
}

// GetSelSecondaryC2CommMode returns the SelSecondaryC2CommMode field value if set, zero value otherwise.
func (o *SelectedC2CommModeNotif) GetSelSecondaryC2CommMode() C2CommMode {
	if o == nil || isNil(o.SelSecondaryC2CommMode) {
		var ret C2CommMode
		return ret
	}
	return *o.SelSecondaryC2CommMode
}

// GetSelSecondaryC2CommModeOk returns a tuple with the SelSecondaryC2CommMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedC2CommModeNotif) GetSelSecondaryC2CommModeOk() (*C2CommMode, bool) {
	if o == nil || isNil(o.SelSecondaryC2CommMode) {
		return nil, false
	}
	return o.SelSecondaryC2CommMode, true
}

// HasSelSecondaryC2CommMode returns a boolean if a field has been set.
func (o *SelectedC2CommModeNotif) HasSelSecondaryC2CommMode() bool {
	if o != nil && !isNil(o.SelSecondaryC2CommMode) {
		return true
	}

	return false
}

// SetSelSecondaryC2CommMode gets a reference to the given C2CommMode and assigns it to the SelSecondaryC2CommMode field.
func (o *SelectedC2CommModeNotif) SetSelSecondaryC2CommMode(v C2CommMode) {
	o.SelSecondaryC2CommMode = &v
}

func (o SelectedC2CommModeNotif) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectedC2CommModeNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uasId"] = o.UasId
	toSerialize["selPrimaryC2CommMode"] = o.SelPrimaryC2CommMode
	if !isNil(o.SelSecondaryC2CommMode) {
		toSerialize["selSecondaryC2CommMode"] = o.SelSecondaryC2CommMode
	}
	return toSerialize, nil
}

type NullableSelectedC2CommModeNotif struct {
	value *SelectedC2CommModeNotif
	isSet bool
}

func (v NullableSelectedC2CommModeNotif) Get() *SelectedC2CommModeNotif {
	return v.value
}

func (v *NullableSelectedC2CommModeNotif) Set(val *SelectedC2CommModeNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedC2CommModeNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedC2CommModeNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedC2CommModeNotif(val *SelectedC2CommModeNotif) *NullableSelectedC2CommModeNotif {
	return &NullableSelectedC2CommModeNotif{value: val, isSet: true}
}

func (v NullableSelectedC2CommModeNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedC2CommModeNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


