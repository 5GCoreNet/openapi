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

// checks if the NssfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NssfFunctionSingleAllOf{}

// NssfFunctionSingleAllOf struct for NssfFunctionSingleAllOf
type NssfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewNssfFunctionSingleAllOf instantiates a new NssfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssfFunctionSingleAllOf() *NssfFunctionSingleAllOf {
	this := NssfFunctionSingleAllOf{}
	return &this
}

// NewNssfFunctionSingleAllOfWithDefaults instantiates a new NssfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssfFunctionSingleAllOfWithDefaults() *NssfFunctionSingleAllOf {
	this := NssfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NssfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NssfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *NssfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o NssfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NssfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableNssfFunctionSingleAllOf struct {
	value *NssfFunctionSingleAllOf
	isSet bool
}

func (v NullableNssfFunctionSingleAllOf) Get() *NssfFunctionSingleAllOf {
	return v.value
}

func (v *NullableNssfFunctionSingleAllOf) Set(val *NssfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfFunctionSingleAllOf(val *NssfFunctionSingleAllOf) *NullableNssfFunctionSingleAllOf {
	return &NullableNssfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableNssfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


