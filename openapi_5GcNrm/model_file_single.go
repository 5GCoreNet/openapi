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

// checks if the FileSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSingle{}

// FileSingle struct for FileSingle
type FileSingle struct {
	Top
	Attributes *FileSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFileSingle instantiates a new FileSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSingle(id NullableString) *FileSingle {
	this := FileSingle{}
	this.Id = id
	return &this
}

// NewFileSingleWithDefaults instantiates a new FileSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSingleWithDefaults() *FileSingle {
	this := FileSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FileSingle) GetAttributes() FileSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret FileSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingle) GetAttributesOk() (*FileSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FileSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FileSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FileSingle) SetAttributes(v FileSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FileSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSingle) ToMap() (map[string]interface{}, error) {
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

type NullableFileSingle struct {
	value *FileSingle
	isSet bool
}

func (v NullableFileSingle) Get() *FileSingle {
	return v.value
}

func (v *NullableFileSingle) Set(val *FileSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSingle(val *FileSingle) *NullableFileSingle {
	return &NullableFileSingle{value: val, isSet: true}
}

func (v NullableFileSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
