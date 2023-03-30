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

// checks if the DDNMFFunctionSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DDNMFFunctionSingleAllOf{}

// DDNMFFunctionSingleAllOf struct for DDNMFFunctionSingleAllOf
type DDNMFFunctionSingleAllOf struct {
	Attributes *DDNMFFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDDNMFFunctionSingleAllOf instantiates a new DDNMFFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDDNMFFunctionSingleAllOf() *DDNMFFunctionSingleAllOf {
	this := DDNMFFunctionSingleAllOf{}
	return &this
}

// NewDDNMFFunctionSingleAllOfWithDefaults instantiates a new DDNMFFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDDNMFFunctionSingleAllOfWithDefaults() *DDNMFFunctionSingleAllOf {
	this := DDNMFFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DDNMFFunctionSingleAllOf) GetAttributes() DDNMFFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret DDNMFFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNMFFunctionSingleAllOf) GetAttributesOk() (*DDNMFFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DDNMFFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DDNMFFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *DDNMFFunctionSingleAllOf) SetAttributes(v DDNMFFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o DDNMFFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DDNMFFunctionSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDDNMFFunctionSingleAllOf struct {
	value *DDNMFFunctionSingleAllOf
	isSet bool
}

func (v NullableDDNMFFunctionSingleAllOf) Get() *DDNMFFunctionSingleAllOf {
	return v.value
}

func (v *NullableDDNMFFunctionSingleAllOf) Set(val *DDNMFFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDDNMFFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDDNMFFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDDNMFFunctionSingleAllOf(val *DDNMFFunctionSingleAllOf) *NullableDDNMFFunctionSingleAllOf {
	return &NullableDDNMFFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableDDNMFFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDDNMFFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


