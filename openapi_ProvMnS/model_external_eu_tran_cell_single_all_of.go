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

// checks if the ExternalEUTranCellSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEUTranCellSingleAllOf{}

// ExternalEUTranCellSingleAllOf struct for ExternalEUTranCellSingleAllOf
type ExternalEUTranCellSingleAllOf struct {
	Attributes *ExternalEUTranCellSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewExternalEUTranCellSingleAllOf instantiates a new ExternalEUTranCellSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEUTranCellSingleAllOf() *ExternalEUTranCellSingleAllOf {
	this := ExternalEUTranCellSingleAllOf{}
	return &this
}

// NewExternalEUTranCellSingleAllOfWithDefaults instantiates a new ExternalEUTranCellSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEUTranCellSingleAllOfWithDefaults() *ExternalEUTranCellSingleAllOf {
	this := ExternalEUTranCellSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalEUTranCellSingleAllOf) GetAttributes() ExternalEUTranCellSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ExternalEUTranCellSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEUTranCellSingleAllOf) GetAttributesOk() (*ExternalEUTranCellSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalEUTranCellSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ExternalEUTranCellSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ExternalEUTranCellSingleAllOf) SetAttributes(v ExternalEUTranCellSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o ExternalEUTranCellSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEUTranCellSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableExternalEUTranCellSingleAllOf struct {
	value *ExternalEUTranCellSingleAllOf
	isSet bool
}

func (v NullableExternalEUTranCellSingleAllOf) Get() *ExternalEUTranCellSingleAllOf {
	return v.value
}

func (v *NullableExternalEUTranCellSingleAllOf) Set(val *ExternalEUTranCellSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEUTranCellSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEUTranCellSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEUTranCellSingleAllOf(val *ExternalEUTranCellSingleAllOf) *NullableExternalEUTranCellSingleAllOf {
	return &NullableExternalEUTranCellSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalEUTranCellSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEUTranCellSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
