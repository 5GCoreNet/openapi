/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the CsgInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsgInformation{}

// CsgInformation Information about a Closed Subscriber Group (CSG)
type CsgInformation struct {
	CsgId      string  `json:"csgId"`
	AccessMode *string `json:"accessMode,omitempty"`
	CMi        *bool   `json:"cMi,omitempty"`
}

// NewCsgInformation instantiates a new CsgInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsgInformation(csgId string) *CsgInformation {
	this := CsgInformation{}
	this.CsgId = csgId
	return &this
}

// NewCsgInformationWithDefaults instantiates a new CsgInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsgInformationWithDefaults() *CsgInformation {
	this := CsgInformation{}
	return &this
}

// GetCsgId returns the CsgId field value
func (o *CsgInformation) GetCsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CsgId
}

// GetCsgIdOk returns a tuple with the CsgId field value
// and a boolean to check if the value has been set.
func (o *CsgInformation) GetCsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsgId, true
}

// SetCsgId sets field value
func (o *CsgInformation) SetCsgId(v string) {
	o.CsgId = v
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *CsgInformation) GetAccessMode() string {
	if o == nil || IsNil(o.AccessMode) {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsgInformation) GetAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessMode) {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *CsgInformation) HasAccessMode() bool {
	if o != nil && !IsNil(o.AccessMode) {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *CsgInformation) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetCMi returns the CMi field value if set, zero value otherwise.
func (o *CsgInformation) GetCMi() bool {
	if o == nil || IsNil(o.CMi) {
		var ret bool
		return ret
	}
	return *o.CMi
}

// GetCMiOk returns a tuple with the CMi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsgInformation) GetCMiOk() (*bool, bool) {
	if o == nil || IsNil(o.CMi) {
		return nil, false
	}
	return o.CMi, true
}

// HasCMi returns a boolean if a field has been set.
func (o *CsgInformation) HasCMi() bool {
	if o != nil && !IsNil(o.CMi) {
		return true
	}

	return false
}

// SetCMi gets a reference to the given bool and assigns it to the CMi field.
func (o *CsgInformation) SetCMi(v bool) {
	o.CMi = &v
}

func (o CsgInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsgInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["csgId"] = o.CsgId
	if !IsNil(o.AccessMode) {
		toSerialize["accessMode"] = o.AccessMode
	}
	if !IsNil(o.CMi) {
		toSerialize["cMi"] = o.CMi
	}
	return toSerialize, nil
}

type NullableCsgInformation struct {
	value *CsgInformation
	isSet bool
}

func (v NullableCsgInformation) Get() *CsgInformation {
	return v.value
}

func (v *NullableCsgInformation) Set(val *CsgInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableCsgInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableCsgInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsgInformation(val *CsgInformation) *NullableCsgInformation {
	return &NullableCsgInformation{value: val, isSet: true}
}

func (v NullableCsgInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsgInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
