/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
)

// checks if the MetaSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaSchema{}

// MetaSchema Defines the Meta Schema
type MetaSchema struct {
	// Represents the Identifier of a Meta schema.
	SchemaId string    `json:"schemaId"`
	MetaTags []TagType `json:"metaTags"`
}

// NewMetaSchema instantiates a new MetaSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaSchema(schemaId string, metaTags []TagType) *MetaSchema {
	this := MetaSchema{}
	this.SchemaId = schemaId
	this.MetaTags = metaTags
	return &this
}

// NewMetaSchemaWithDefaults instantiates a new MetaSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaSchemaWithDefaults() *MetaSchema {
	this := MetaSchema{}
	return &this
}

// GetSchemaId returns the SchemaId field value
func (o *MetaSchema) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *MetaSchema) GetSchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *MetaSchema) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetMetaTags returns the MetaTags field value
func (o *MetaSchema) GetMetaTags() []TagType {
	if o == nil {
		var ret []TagType
		return ret
	}

	return o.MetaTags
}

// GetMetaTagsOk returns a tuple with the MetaTags field value
// and a boolean to check if the value has been set.
func (o *MetaSchema) GetMetaTagsOk() ([]TagType, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetaTags, true
}

// SetMetaTags sets field value
func (o *MetaSchema) SetMetaTags(v []TagType) {
	o.MetaTags = v
}

func (o MetaSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemaId"] = o.SchemaId
	toSerialize["metaTags"] = o.MetaTags
	return toSerialize, nil
}

type NullableMetaSchema struct {
	value *MetaSchema
	isSet bool
}

func (v NullableMetaSchema) Get() *MetaSchema {
	return v.value
}

func (v *NullableMetaSchema) Set(val *MetaSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaSchema(val *MetaSchema) *NullableMetaSchema {
	return &NullableMetaSchema{value: val, isSet: true}
}

func (v NullableMetaSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
