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

// checks if the MDAReportAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MDAReportAllOf{}

// MDAReportAllOf struct for MDAReportAllOf
type MDAReportAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewMDAReportAllOf instantiates a new MDAReportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAReportAllOf() *MDAReportAllOf {
	this := MDAReportAllOf{}
	return &this
}

// NewMDAReportAllOfWithDefaults instantiates a new MDAReportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAReportAllOfWithDefaults() *MDAReportAllOf {
	this := MDAReportAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MDAReportAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAReportAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MDAReportAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *MDAReportAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o MDAReportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MDAReportAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableMDAReportAllOf struct {
	value *MDAReportAllOf
	isSet bool
}

func (v NullableMDAReportAllOf) Get() *MDAReportAllOf {
	return v.value
}

func (v *NullableMDAReportAllOf) Set(val *MDAReportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAReportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAReportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAReportAllOf(val *MDAReportAllOf) *NullableMDAReportAllOf {
	return &NullableMDAReportAllOf{value: val, isSet: true}
}

func (v NullableMDAReportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAReportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


