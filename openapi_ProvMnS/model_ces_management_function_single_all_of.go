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

// checks if the CESManagementFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CESManagementFunctionSingleAllOf{}

// CESManagementFunctionSingleAllOf struct for CESManagementFunctionSingleAllOf
type CESManagementFunctionSingleAllOf struct {
	Attributes *CESManagementFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewCESManagementFunctionSingleAllOf instantiates a new CESManagementFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCESManagementFunctionSingleAllOf() *CESManagementFunctionSingleAllOf {
	this := CESManagementFunctionSingleAllOf{}
	return &this
}

// NewCESManagementFunctionSingleAllOfWithDefaults instantiates a new CESManagementFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCESManagementFunctionSingleAllOfWithDefaults() *CESManagementFunctionSingleAllOf {
	this := CESManagementFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOf) GetAttributes() CESManagementFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CESManagementFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOf) GetAttributesOk() (*CESManagementFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CESManagementFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *CESManagementFunctionSingleAllOf) SetAttributes(v CESManagementFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o CESManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CESManagementFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableCESManagementFunctionSingleAllOf struct {
	value *CESManagementFunctionSingleAllOf
	isSet bool
}

func (v NullableCESManagementFunctionSingleAllOf) Get() *CESManagementFunctionSingleAllOf {
	return v.value
}

func (v *NullableCESManagementFunctionSingleAllOf) Set(val *CESManagementFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCESManagementFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCESManagementFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCESManagementFunctionSingleAllOf(val *CESManagementFunctionSingleAllOf) *NullableCESManagementFunctionSingleAllOf {
	return &NullableCESManagementFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableCESManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCESManagementFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


