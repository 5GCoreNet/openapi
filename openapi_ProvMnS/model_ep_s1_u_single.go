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

// checks if the EPS1USingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPS1USingle{}

// EPS1USingle struct for EPS1USingle
type EPS1USingle struct {
	Top
	Attributes *EPF1CSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEPS1USingle instantiates a new EPS1USingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPS1USingle(id NullableString) *EPS1USingle {
	this := EPS1USingle{}
	this.Id = id
	return &this
}

// NewEPS1USingleWithDefaults instantiates a new EPS1USingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPS1USingleWithDefaults() *EPS1USingle {
	this := EPS1USingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EPS1USingle) GetAttributes() EPF1CSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EPF1CSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPS1USingle) GetAttributesOk() (*EPF1CSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EPS1USingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EPF1CSingleAllOfAttributes and assigns it to the Attributes field.
func (o *EPS1USingle) SetAttributes(v EPF1CSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EPS1USingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPS1USingle) ToMap() (map[string]interface{}, error) {
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

type NullableEPS1USingle struct {
	value *EPS1USingle
	isSet bool
}

func (v NullableEPS1USingle) Get() *EPS1USingle {
	return v.value
}

func (v *NullableEPS1USingle) Set(val *EPS1USingle) {
	v.value = val
	v.isSet = true
}

func (v NullableEPS1USingle) IsSet() bool {
	return v.isSet
}

func (v *NullableEPS1USingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPS1USingle(val *EPS1USingle) *NullableEPS1USingle {
	return &NullableEPS1USingle{value: val, isSet: true}
}

func (v NullableEPS1USingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPS1USingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
