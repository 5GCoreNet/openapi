/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the ExternalGnbCuCpFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbCuCpFunctionSingleAllOf{}

// ExternalGnbCuCpFunctionSingleAllOf struct for ExternalGnbCuCpFunctionSingleAllOf
type ExternalGnbCuCpFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewExternalGnbCuCpFunctionSingleAllOf instantiates a new ExternalGnbCuCpFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbCuCpFunctionSingleAllOf() *ExternalGnbCuCpFunctionSingleAllOf {
	this := ExternalGnbCuCpFunctionSingleAllOf{}
	return &this
}

// NewExternalGnbCuCpFunctionSingleAllOfWithDefaults instantiates a new ExternalGnbCuCpFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbCuCpFunctionSingleAllOfWithDefaults() *ExternalGnbCuCpFunctionSingleAllOf {
	this := ExternalGnbCuCpFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalGnbCuCpFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o ExternalGnbCuCpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbCuCpFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalGnbCuCpFunctionSingleAllOf struct {
	value *ExternalGnbCuCpFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalGnbCuCpFunctionSingleAllOf) Get() *ExternalGnbCuCpFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOf) Set(val *ExternalGnbCuCpFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbCuCpFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbCuCpFunctionSingleAllOf(val *ExternalGnbCuCpFunctionSingleAllOf) *NullableExternalGnbCuCpFunctionSingleAllOf {
	return &NullableExternalGnbCuCpFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalGnbCuCpFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbCuCpFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


