/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the FilesSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSingle{}

// FilesSingle struct for FilesSingle
type FilesSingle struct {
	Top
	Attributes *FilesSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFilesSingle instantiates a new FilesSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSingle(id NullableString) *FilesSingle {
	this := FilesSingle{}
	this.Id = id
	return &this
}

// NewFilesSingleWithDefaults instantiates a new FilesSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSingleWithDefaults() *FilesSingle {
	this := FilesSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FilesSingle) GetAttributes() FilesSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret FilesSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingle) GetAttributesOk() (*FilesSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FilesSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FilesSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FilesSingle) SetAttributes(v FilesSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FilesSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSingle) ToMap() (map[string]interface{}, error) {
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

type NullableFilesSingle struct {
	value *FilesSingle
	isSet bool
}

func (v NullableFilesSingle) Get() *FilesSingle {
	return v.value
}

func (v *NullableFilesSingle) Set(val *FilesSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSingle(val *FilesSingle) *NullableFilesSingle {
	return &NullableFilesSingle{value: val, isSet: true}
}

func (v NullableFilesSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
