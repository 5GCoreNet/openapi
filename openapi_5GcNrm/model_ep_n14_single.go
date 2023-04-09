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

// checks if the EPN14Single type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPN14Single{}

// EPN14Single struct for EPN14Single
type EPN14Single struct {
	Top
	Attributes *EPN2SingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPN14Single instantiates a new EPN14Single object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPN14Single(id NullableString) *EPN14Single {
	this := EPN14Single{}
	this.Id = id
	return &this
}

// NewEPN14SingleWithDefaults instantiates a new EPN14Single object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPN14SingleWithDefaults() *EPN14Single {
	this := EPN14Single{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPN14Single) GetAttributes() EPN2SingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPN2SingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPN14Single) GetAttributesOk() (*EPN2SingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPN14Single) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPN2SingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPN14Single) SetAttributes(v EPN2SingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPN14Single) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPN14Single) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableEPN14Single struct {
	value *EPN14Single
	isSet bool
}

func (v NullableEPN14Single) Get() *EPN14Single {
	return v.value
}

func (v *NullableEPN14Single) Set(val *EPN14Single) {
	v.value = val
	v.isSet = true
}

func (v NullableEPN14Single) IsSet() bool {
	return v.isSet
}

func (v *NullableEPN14Single) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPN14Single(val *EPN14Single) *NullableEPN14Single {
	return &NullableEPN14Single{value: val, isSet: true}
}

func (v NullableEPN14Single) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPN14Single) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
