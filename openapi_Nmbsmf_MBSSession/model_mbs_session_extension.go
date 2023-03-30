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

// checks if the MbsSessionExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSessionExtension{}

// MbsSessionExtension MB-SMF API specific extensions to the MbsSession common data type
type MbsSessionExtension struct {
	MbsSecurityContext *MbsSecurityContext `json:"mbsSecurityContext,omitempty"`
	ContactPcfInd *bool `json:"contactPcfInd,omitempty"`
}

// NewMbsSessionExtension instantiates a new MbsSessionExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionExtension() *MbsSessionExtension {
	this := MbsSessionExtension{}
	var contactPcfInd bool = false
	this.ContactPcfInd = &contactPcfInd
	return &this
}

// NewMbsSessionExtensionWithDefaults instantiates a new MbsSessionExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionExtensionWithDefaults() *MbsSessionExtension {
	this := MbsSessionExtension{}
	var contactPcfInd bool = false
	this.ContactPcfInd = &contactPcfInd
	return &this
}

// GetMbsSecurityContext returns the MbsSecurityContext field value if set, zero value otherwise.
func (o *MbsSessionExtension) GetMbsSecurityContext() MbsSecurityContext {
	if o == nil || IsNil(o.MbsSecurityContext) {
		var ret MbsSecurityContext
		return ret
	}
	return *o.MbsSecurityContext
}

// GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionExtension) GetMbsSecurityContextOk() (*MbsSecurityContext, bool) {
	if o == nil || IsNil(o.MbsSecurityContext) {
		return nil, false
	}
	return o.MbsSecurityContext, true
}

// HasMbsSecurityContext returns a boolean if a field has been set.
func (o *MbsSessionExtension) HasMbsSecurityContext() bool {
	if o != nil && !IsNil(o.MbsSecurityContext) {
		return true
	}

	return false
}

// SetMbsSecurityContext gets a reference to the given MbsSecurityContext and assigns it to the MbsSecurityContext field.
func (o *MbsSessionExtension) SetMbsSecurityContext(v MbsSecurityContext) {
	o.MbsSecurityContext = &v
}

// GetContactPcfInd returns the ContactPcfInd field value if set, zero value otherwise.
func (o *MbsSessionExtension) GetContactPcfInd() bool {
	if o == nil || IsNil(o.ContactPcfInd) {
		var ret bool
		return ret
	}
	return *o.ContactPcfInd
}

// GetContactPcfIndOk returns a tuple with the ContactPcfInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionExtension) GetContactPcfIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ContactPcfInd) {
		return nil, false
	}
	return o.ContactPcfInd, true
}

// HasContactPcfInd returns a boolean if a field has been set.
func (o *MbsSessionExtension) HasContactPcfInd() bool {
	if o != nil && !IsNil(o.ContactPcfInd) {
		return true
	}

	return false
}

// SetContactPcfInd gets a reference to the given bool and assigns it to the ContactPcfInd field.
func (o *MbsSessionExtension) SetContactPcfInd(v bool) {
	o.ContactPcfInd = &v
}

func (o MbsSessionExtension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSessionExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsSecurityContext) {
		toSerialize["mbsSecurityContext"] = o.MbsSecurityContext
	}
	if !IsNil(o.ContactPcfInd) {
		toSerialize["contactPcfInd"] = o.ContactPcfInd
	}
	return toSerialize, nil
}

type NullableMbsSessionExtension struct {
	value *MbsSessionExtension
	isSet bool
}

func (v NullableMbsSessionExtension) Get() *MbsSessionExtension {
	return v.value
}

func (v *NullableMbsSessionExtension) Set(val *MbsSessionExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionExtension(val *MbsSessionExtension) *NullableMbsSessionExtension {
	return &NullableMbsSessionExtension{value: val, isSet: true}
}

func (v NullableMbsSessionExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

