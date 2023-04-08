/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the EPN3SingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN3SingleAllOf{}

// EPN3SingleAllOf struct for EPN3SingleAllOf
type EPN3SingleAllOf struct {
	Attributes *EPRPAttr `json:"attributes,omitempty"`
}

// NewEPN3SingleAllOf instantiates a new EPN3SingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN3SingleAllOf() *EPN3SingleAllOf {
	this := EPN3SingleAllOf{}
	return &this
}

// NewEPN3SingleAllOfWithDefaults instantiates a new EPN3SingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN3SingleAllOfWithDefaults() *EPN3SingleAllOf {
	this := EPN3SingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN3SingleAllOf) GetAttributes() EPRPAttr {
	if o == nil || isNil(o.Attributes) {
		var ret EPRPAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN3SingleAllOf) GetAttributesOk() (*EPRPAttr, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN3SingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPRPAttr and assigns it to the Attributes field.
func (o *EPN3SingleAllOf) SetAttributes(v EPRPAttr) {
	o.Attributes = &v
}

func (o EPN3SingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN3SingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPN3SingleAllOf struct {
	value *EPN3SingleAllOf
	isSet bool
}

func (v NullableEPN3SingleAllOf) Get() *EPN3SingleAllOf {
	return v.value
}

func (v *NullableEPN3SingleAllOf) Set(val *EPN3SingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN3SingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN3SingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN3SingleAllOf(val *EPN3SingleAllOf) *NullableEPN3SingleAllOf {
	return &NullableEPN3SingleAllOf{value: val, isSet: true}
}

func (v NullableEPN3SingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN3SingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


