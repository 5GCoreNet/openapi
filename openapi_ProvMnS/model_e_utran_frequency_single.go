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

// checks if the EUtranFrequencySingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EUtranFrequencySingle{}

// EUtranFrequencySingle struct for EUtranFrequencySingle
type EUtranFrequencySingle struct {
	Top
	Attributes *EUtranFrequencySingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEUtranFrequencySingle instantiates a new EUtranFrequencySingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEUtranFrequencySingle(id NullableString) *EUtranFrequencySingle {
	this := EUtranFrequencySingle{}
	this.Id = id
	return &this
}

// NewEUtranFrequencySingleWithDefaults instantiates a new EUtranFrequencySingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEUtranFrequencySingleWithDefaults() *EUtranFrequencySingle {
	this := EUtranFrequencySingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EUtranFrequencySingle) GetAttributes() EUtranFrequencySingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret EUtranFrequencySingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFrequencySingle) GetAttributesOk() (*EUtranFrequencySingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EUtranFrequencySingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EUtranFrequencySingleAllOfAttributes and assigns it to the Attributes field.
func (o *EUtranFrequencySingle) SetAttributes(v EUtranFrequencySingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EUtranFrequencySingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EUtranFrequencySingle) ToMap() (map[string]interface{}, error) {
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

type NullableEUtranFrequencySingle struct {
	value *EUtranFrequencySingle
	isSet bool
}

func (v NullableEUtranFrequencySingle) Get() *EUtranFrequencySingle {
	return v.value
}

func (v *NullableEUtranFrequencySingle) Set(val *EUtranFrequencySingle) {
	v.value = val
	v.isSet = true
}

func (v NullableEUtranFrequencySingle) IsSet() bool {
	return v.isSet
}

func (v *NullableEUtranFrequencySingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEUtranFrequencySingle(val *EUtranFrequencySingle) *NullableEUtranFrequencySingle {
	return &NullableEUtranFrequencySingle{value: val, isSet: true}
}

func (v NullableEUtranFrequencySingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEUtranFrequencySingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
