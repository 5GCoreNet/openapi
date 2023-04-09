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

// checks if the EPNLSSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPNLSSingle{}

// EPNLSSingle struct for EPNLSSingle
type EPNLSSingle struct {
	Top
	Attributes *EPN2SingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPNLSSingle instantiates a new EPNLSSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPNLSSingle(id NullableString) *EPNLSSingle {
	this := EPNLSSingle{}
	this.Id = id
	return &this
}

// NewEPNLSSingleWithDefaults instantiates a new EPNLSSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPNLSSingleWithDefaults() *EPNLSSingle {
	this := EPNLSSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPNLSSingle) GetAttributes() EPN2SingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPN2SingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPNLSSingle) GetAttributesOk() (*EPN2SingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPNLSSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPN2SingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPNLSSingle) SetAttributes(v EPN2SingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPNLSSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPNLSSingle) ToMap() (map[string]interface{}, error) {
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

type NullableEPNLSSingle struct {
	value *EPNLSSingle
	isSet bool
}

func (v NullableEPNLSSingle) Get() *EPNLSSingle {
	return v.value
}

func (v *NullableEPNLSSingle) Set(val *EPNLSSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableEPNLSSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableEPNLSSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPNLSSingle(val *EPNLSSingle) *NullableEPNLSSingle {
	return &NullableEPNLSSingle{value: val, isSet: true}
}

func (v NullableEPNLSSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPNLSSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
