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

// checks if the EPN32SingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN32SingleAllOf{}

// EPN32SingleAllOf struct for EPN32SingleAllOf
type EPN32SingleAllOf struct {
	Attributes *EPRPAttr `json:"attributes,omitempty"`
}

// NewEPN32SingleAllOf instantiates a new EPN32SingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN32SingleAllOf() *EPN32SingleAllOf {
	this := EPN32SingleAllOf{}
	return &this
}

// NewEPN32SingleAllOfWithDefaults instantiates a new EPN32SingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN32SingleAllOfWithDefaults() *EPN32SingleAllOf {
	this := EPN32SingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN32SingleAllOf) GetAttributes() EPRPAttr {
	if o == nil || isNil(o.Attributes) {
		var ret EPRPAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN32SingleAllOf) GetAttributesOk() (*EPRPAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN32SingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPRPAttr and assigns it to the Attributes field.
func (o *EPN32SingleAllOf) SetAttributes(v EPRPAttr) {
	o.Attributes = &v
}

func (o EPN32SingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN32SingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPN32SingleAllOf struct {
	value *EPN32SingleAllOf
	isSet bool
}

func (v NullableEPN32SingleAllOf) Get() *EPN32SingleAllOf {
	return v.value
}

func (v *NullableEPN32SingleAllOf) Set(val *EPN32SingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN32SingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN32SingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN32SingleAllOf(val *EPN32SingleAllOf) *NullableEPN32SingleAllOf {
	return &NullableEPN32SingleAllOf{value: val, isSet: true}
}

func (v NullableEPN32SingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN32SingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


