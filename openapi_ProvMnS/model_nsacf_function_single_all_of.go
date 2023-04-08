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

// checks if the NsacfFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsacfFunctionSingleAllOf{}

// NsacfFunctionSingleAllOf struct for NsacfFunctionSingleAllOf
type NsacfFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewNsacfFunctionSingleAllOf instantiates a new NsacfFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfFunctionSingleAllOf() *NsacfFunctionSingleAllOf {
	this := NsacfFunctionSingleAllOf{}
	return &this
}

// NewNsacfFunctionSingleAllOfWithDefaults instantiates a new NsacfFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfFunctionSingleAllOfWithDefaults() *NsacfFunctionSingleAllOf {
	this := NsacfFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NsacfFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NsacfFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *NsacfFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o NsacfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsacfFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableNsacfFunctionSingleAllOf struct {
	value *NsacfFunctionSingleAllOf
	isSet bool
}

func (v NullableNsacfFunctionSingleAllOf) Get() *NsacfFunctionSingleAllOf {
	return v.value
}

func (v *NullableNsacfFunctionSingleAllOf) Set(val *NsacfFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfFunctionSingleAllOf(val *NsacfFunctionSingleAllOf) *NullableNsacfFunctionSingleAllOf {
	return &NullableNsacfFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableNsacfFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


