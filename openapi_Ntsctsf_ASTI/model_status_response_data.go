/*
Ntsctsf_ASTI Service API

TSCTSF  Access Stratum time distribution configuration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_ASTI

import (
	"encoding/json"
)

// checks if the StatusResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseData{}

// StatusResponseData Contains the parameters for the status of the access stratum time distribution for a list of UEs. 
type StatusResponseData struct {
	InactiveUes []string `json:"inactiveUes,omitempty"`
	InactiveGpsis []string `json:"inactiveGpsis,omitempty"`
	ActiveUes []ActiveUe `json:"activeUes,omitempty"`
}

// NewStatusResponseData instantiates a new StatusResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseData() *StatusResponseData {
	this := StatusResponseData{}
	return &this
}

// NewStatusResponseDataWithDefaults instantiates a new StatusResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseDataWithDefaults() *StatusResponseData {
	this := StatusResponseData{}
	return &this
}

// GetInactiveUes returns the InactiveUes field value if set, zero value otherwise.
func (o *StatusResponseData) GetInactiveUes() []string {
	if o == nil || isNil(o.InactiveUes) {
		var ret []string
		return ret
	}
	return o.InactiveUes
}

// GetInactiveUesOk returns a tuple with the InactiveUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponseData) GetInactiveUesOk() ([]string, bool) {
	if o == nil || isNil(o.InactiveUes) {
		return nil, false
	}
	return o.InactiveUes, true
}

// HasInactiveUes returns a boolean if a field has been set.
func (o *StatusResponseData) HasInactiveUes() bool {
	if o != nil && !isNil(o.InactiveUes) {
		return true
	}

	return false
}

// SetInactiveUes gets a reference to the given []string and assigns it to the InactiveUes field.
func (o *StatusResponseData) SetInactiveUes(v []string) {
	o.InactiveUes = v
}

// GetInactiveGpsis returns the InactiveGpsis field value if set, zero value otherwise.
func (o *StatusResponseData) GetInactiveGpsis() []string {
	if o == nil || isNil(o.InactiveGpsis) {
		var ret []string
		return ret
	}
	return o.InactiveGpsis
}

// GetInactiveGpsisOk returns a tuple with the InactiveGpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponseData) GetInactiveGpsisOk() ([]string, bool) {
	if o == nil || isNil(o.InactiveGpsis) {
		return nil, false
	}
	return o.InactiveGpsis, true
}

// HasInactiveGpsis returns a boolean if a field has been set.
func (o *StatusResponseData) HasInactiveGpsis() bool {
	if o != nil && !isNil(o.InactiveGpsis) {
		return true
	}

	return false
}

// SetInactiveGpsis gets a reference to the given []string and assigns it to the InactiveGpsis field.
func (o *StatusResponseData) SetInactiveGpsis(v []string) {
	o.InactiveGpsis = v
}

// GetActiveUes returns the ActiveUes field value if set, zero value otherwise.
func (o *StatusResponseData) GetActiveUes() []ActiveUe {
	if o == nil || isNil(o.ActiveUes) {
		var ret []ActiveUe
		return ret
	}
	return o.ActiveUes
}

// GetActiveUesOk returns a tuple with the ActiveUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusResponseData) GetActiveUesOk() ([]ActiveUe, bool) {
	if o == nil || isNil(o.ActiveUes) {
		return nil, false
	}
	return o.ActiveUes, true
}

// HasActiveUes returns a boolean if a field has been set.
func (o *StatusResponseData) HasActiveUes() bool {
	if o != nil && !isNil(o.ActiveUes) {
		return true
	}

	return false
}

// SetActiveUes gets a reference to the given []ActiveUe and assigns it to the ActiveUes field.
func (o *StatusResponseData) SetActiveUes(v []ActiveUe) {
	o.ActiveUes = v
}

func (o StatusResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InactiveUes) {
		toSerialize["inactiveUes"] = o.InactiveUes
	}
	if !isNil(o.InactiveGpsis) {
		toSerialize["inactiveGpsis"] = o.InactiveGpsis
	}
	if !isNil(o.ActiveUes) {
		toSerialize["activeUes"] = o.ActiveUes
	}
	return toSerialize, nil
}

type NullableStatusResponseData struct {
	value *StatusResponseData
	isSet bool
}

func (v NullableStatusResponseData) Get() *StatusResponseData {
	return v.value
}

func (v *NullableStatusResponseData) Set(val *StatusResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseData(val *StatusResponseData) *NullableStatusResponseData {
	return &NullableStatusResponseData{value: val, isSet: true}
}

func (v NullableStatusResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


