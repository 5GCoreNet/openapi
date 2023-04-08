/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the FilesSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSingleAllOf{}

// FilesSingleAllOf struct for FilesSingleAllOf
type FilesSingleAllOf struct {
	Attributes *FilesSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewFilesSingleAllOf instantiates a new FilesSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSingleAllOf() *FilesSingleAllOf {
	this := FilesSingleAllOf{}
	return &this
}

// NewFilesSingleAllOfWithDefaults instantiates a new FilesSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSingleAllOfWithDefaults() *FilesSingleAllOf {
	this := FilesSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FilesSingleAllOf) GetAttributes() FilesSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret FilesSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingleAllOf) GetAttributesOk() (*FilesSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FilesSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FilesSingleAllOfAttributes and assigns it to the Attributes field.
func (o *FilesSingleAllOf) SetAttributes(v FilesSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o FilesSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableFilesSingleAllOf struct {
	value *FilesSingleAllOf
	isSet bool
}

func (v NullableFilesSingleAllOf) Get() *FilesSingleAllOf {
	return v.value
}

func (v *NullableFilesSingleAllOf) Set(val *FilesSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSingleAllOf(val *FilesSingleAllOf) *NullableFilesSingleAllOf {
	return &NullableFilesSingleAllOf{value: val, isSet: true}
}

func (v NullableFilesSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


