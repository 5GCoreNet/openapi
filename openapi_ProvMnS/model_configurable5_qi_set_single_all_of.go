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

// checks if the Configurable5QISetSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Configurable5QISetSingleAllOf{}

// Configurable5QISetSingleAllOf struct for Configurable5QISetSingleAllOf
type Configurable5QISetSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewConfigurable5QISetSingleAllOf instantiates a new Configurable5QISetSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurable5QISetSingleAllOf() *Configurable5QISetSingleAllOf {
	this := Configurable5QISetSingleAllOf{}
	return &this
}

// NewConfigurable5QISetSingleAllOfWithDefaults instantiates a new Configurable5QISetSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurable5QISetSingleAllOfWithDefaults() *Configurable5QISetSingleAllOf {
	this := Configurable5QISetSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Configurable5QISetSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Configurable5QISetSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Configurable5QISetSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *Configurable5QISetSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o Configurable5QISetSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Configurable5QISetSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableConfigurable5QISetSingleAllOf struct {
	value *Configurable5QISetSingleAllOf
	isSet bool
}

func (v NullableConfigurable5QISetSingleAllOf) Get() *Configurable5QISetSingleAllOf {
	return v.value
}

func (v *NullableConfigurable5QISetSingleAllOf) Set(val *Configurable5QISetSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurable5QISetSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurable5QISetSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurable5QISetSingleAllOf(val *Configurable5QISetSingleAllOf) *NullableConfigurable5QISetSingleAllOf {
	return &NullableConfigurable5QISetSingleAllOf{value: val, isSet: true}
}

func (v NullableConfigurable5QISetSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurable5QISetSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


