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

// checks if the ExternalGnbDuFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGnbDuFunctionSingleAllOf{}

// ExternalGnbDuFunctionSingleAllOf struct for ExternalGnbDuFunctionSingleAllOf
type ExternalGnbDuFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewExternalGnbDuFunctionSingleAllOf instantiates a new ExternalGnbDuFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbDuFunctionSingleAllOf() *ExternalGnbDuFunctionSingleAllOf {
	this := ExternalGnbDuFunctionSingleAllOf{}
	return &this
}

// NewExternalGnbDuFunctionSingleAllOfWithDefaults instantiates a new ExternalGnbDuFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbDuFunctionSingleAllOfWithDefaults() *ExternalGnbDuFunctionSingleAllOf {
	this := ExternalGnbDuFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalGnbDuFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbDuFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalGnbDuFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalGnbDuFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o ExternalGnbDuFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGnbDuFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalGnbDuFunctionSingleAllOf struct {
	value *ExternalGnbDuFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalGnbDuFunctionSingleAllOf) Get() *ExternalGnbDuFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalGnbDuFunctionSingleAllOf) Set(val *ExternalGnbDuFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbDuFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbDuFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbDuFunctionSingleAllOf(val *ExternalGnbDuFunctionSingleAllOf) *NullableExternalGnbDuFunctionSingleAllOf {
	return &NullableExternalGnbDuFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalGnbDuFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbDuFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


