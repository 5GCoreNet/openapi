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

// checks if the Configurable5QISetSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Configurable5QISetSingle{}

// Configurable5QISetSingle struct for Configurable5QISetSingle
type Configurable5QISetSingle struct {
	Top
	Attributes *Configurable5QISetSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewConfigurable5QISetSingle instantiates a new Configurable5QISetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurable5QISetSingle(id NullableString) *Configurable5QISetSingle {
	this := Configurable5QISetSingle{}
	this.Id = id
	return &this
}

// NewConfigurable5QISetSingleWithDefaults instantiates a new Configurable5QISetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurable5QISetSingleWithDefaults() *Configurable5QISetSingle {
	this := Configurable5QISetSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Configurable5QISetSingle) GetAttributes() Configurable5QISetSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret Configurable5QISetSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingle) GetAttributesOk() (*Configurable5QISetSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Configurable5QISetSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Configurable5QISetSingleAllOfAttributes and assigns it to the Attributes field.
func (o *Configurable5QISetSingle) SetAttributes(v Configurable5QISetSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o Configurable5QISetSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Configurable5QISetSingle) ToMap() (map[string]interface{}, error) {
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

type NullableConfigurable5QISetSingle struct {
	value *Configurable5QISetSingle
	isSet bool
}

func (v NullableConfigurable5QISetSingle) Get() *Configurable5QISetSingle {
	return v.value
}

func (v *NullableConfigurable5QISetSingle) Set(val *Configurable5QISetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurable5QISetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurable5QISetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurable5QISetSingle(val *Configurable5QISetSingle) *NullableConfigurable5QISetSingle {
	return &NullableConfigurable5QISetSingle{value: val, isSet: true}
}

func (v NullableConfigurable5QISetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurable5QISetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
