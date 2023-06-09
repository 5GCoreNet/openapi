/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// checks if the MDAFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAFunctionSingleAllOfAttributes{}

// MDAFunctionSingleAllOfAttributes struct for MDAFunctionSingleAllOfAttributes
type MDAFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	SupportedMDACapabilities []string `json:"supportedMDACapabilities,omitempty"`
}

// NewMDAFunctionSingleAllOfAttributes instantiates a new MDAFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAFunctionSingleAllOfAttributes() *MDAFunctionSingleAllOfAttributes {
	this := MDAFunctionSingleAllOfAttributes{}
	return &this
}

// NewMDAFunctionSingleAllOfAttributesWithDefaults instantiates a new MDAFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAFunctionSingleAllOfAttributesWithDefaults() *MDAFunctionSingleAllOfAttributes {
	this := MDAFunctionSingleAllOfAttributes{}
	return &this
}

// GetSupportedMDACapabilities returns the SupportedMDACapabilities field value if set, zero value otherwise.
func (o *MDAFunctionSingleAllOfAttributes) GetSupportedMDACapabilities() []string {
	if o == nil || IsNil(o.SupportedMDACapabilities) {
		var ret []string
		return ret
	}
	return o.SupportedMDACapabilities
}

// GetSupportedMDACapabilitiesOk returns a tuple with the SupportedMDACapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAFunctionSingleAllOfAttributes) GetSupportedMDACapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedMDACapabilities) {
		return nil, false
	}
	return o.SupportedMDACapabilities, true
}

// HasSupportedMDACapabilities returns a boolean if a field has been set.
func (o *MDAFunctionSingleAllOfAttributes) HasSupportedMDACapabilities() bool {
	if o != nil && !IsNil(o.SupportedMDACapabilities) {
		return true
	}

	return false
}

// SetSupportedMDACapabilities gets a reference to the given []string and assigns it to the SupportedMDACapabilities field.
func (o *MDAFunctionSingleAllOfAttributes) SetSupportedMDACapabilities(v []string) {
	o.SupportedMDACapabilities = v
}

func (o MDAFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.SupportedMDACapabilities) {
		toSerialize["supportedMDACapabilities"] = o.SupportedMDACapabilities
	}
	return toSerialize, nil
}

type NullableMDAFunctionSingleAllOfAttributes struct {
	value *MDAFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableMDAFunctionSingleAllOfAttributes) Get() *MDAFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableMDAFunctionSingleAllOfAttributes) Set(val *MDAFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAFunctionSingleAllOfAttributes(val *MDAFunctionSingleAllOfAttributes) *NullableMDAFunctionSingleAllOfAttributes {
	return &NullableMDAFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableMDAFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
