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

// checks if the UdsfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdsfFunctionSingleAllOf{}

// UdsfFunctionSingleAllOf struct for UdsfFunctionSingleAllOf
type UdsfFunctionSingleAllOf struct {
	Attributes *UdsfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewUdsfFunctionSingleAllOf instantiates a new UdsfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdsfFunctionSingleAllOf() *UdsfFunctionSingleAllOf {
	this := UdsfFunctionSingleAllOf{}
	return &this
}

// NewUdsfFunctionSingleAllOfWithDefaults instantiates a new UdsfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdsfFunctionSingleAllOfWithDefaults() *UdsfFunctionSingleAllOf {
	this := UdsfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UdsfFunctionSingleAllOf) GetAttributes() UdsfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret UdsfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfFunctionSingleAllOf) GetAttributesOk() (*UdsfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UdsfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UdsfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *UdsfFunctionSingleAllOf) SetAttributes(v UdsfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o UdsfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdsfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableUdsfFunctionSingleAllOf struct {
	value *UdsfFunctionSingleAllOf
	isSet bool
}

func (v NullableUdsfFunctionSingleAllOf) Get() *UdsfFunctionSingleAllOf {
	return v.value
}

func (v *NullableUdsfFunctionSingleAllOf) Set(val *UdsfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUdsfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUdsfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdsfFunctionSingleAllOf(val *UdsfFunctionSingleAllOf) *NullableUdsfFunctionSingleAllOf {
	return &NullableUdsfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableUdsfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdsfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
