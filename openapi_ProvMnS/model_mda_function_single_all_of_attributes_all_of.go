/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the MDAFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAFunctionSingleAllOfAttributesAllOf{}

// MDAFunctionSingleAllOfAttributesAllOf struct for MDAFunctionSingleAllOfAttributesAllOf
type MDAFunctionSingleAllOfAttributesAllOf struct {
	SupportedMDACapabilities []string `json:"supportedMDACapabilities,omitempty"`
}

// NewMDAFunctionSingleAllOfAttributesAllOf instantiates a new MDAFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAFunctionSingleAllOfAttributesAllOf() *MDAFunctionSingleAllOfAttributesAllOf {
	this := MDAFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewMDAFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new MDAFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAFunctionSingleAllOfAttributesAllOfWithDefaults() *MDAFunctionSingleAllOfAttributesAllOf {
	this := MDAFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetSupportedMDACapabilities returns the SupportedMDACapabilities field value if set, zero value otherwise.
func (o *MDAFunctionSingleAllOfAttributesAllOf) GetSupportedMDACapabilities() []string {
	if o == nil || IsNil(o.SupportedMDACapabilities) {
		var ret []string
		return ret
	}
	return o.SupportedMDACapabilities
}

// GetSupportedMDACapabilitiesOk returns a tuple with the SupportedMDACapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAFunctionSingleAllOfAttributesAllOf) GetSupportedMDACapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedMDACapabilities) {
		return nil, false
	}
	return o.SupportedMDACapabilities, true
}

// HasSupportedMDACapabilities returns a boolean if a field has been set.
func (o *MDAFunctionSingleAllOfAttributesAllOf) HasSupportedMDACapabilities() bool {
	if o != nil && !IsNil(o.SupportedMDACapabilities) {
		return true
	}

	return false
}

// SetSupportedMDACapabilities gets a reference to the given []string and assigns it to the SupportedMDACapabilities field.
func (o *MDAFunctionSingleAllOfAttributesAllOf) SetSupportedMDACapabilities(v []string) {
	o.SupportedMDACapabilities = v
}

func (o MDAFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedMDACapabilities) {
		toSerialize["supportedMDACapabilities"] = o.SupportedMDACapabilities
	}
	return toSerialize, nil
}

type NullableMDAFunctionSingleAllOfAttributesAllOf struct {
	value *MDAFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableMDAFunctionSingleAllOfAttributesAllOf) Get() *MDAFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableMDAFunctionSingleAllOfAttributesAllOf) Set(val *MDAFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAFunctionSingleAllOfAttributesAllOf(val *MDAFunctionSingleAllOfAttributesAllOf) *NullableMDAFunctionSingleAllOfAttributesAllOf {
	return &NullableMDAFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableMDAFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


