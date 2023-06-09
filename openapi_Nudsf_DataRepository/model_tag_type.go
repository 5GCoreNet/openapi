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

// checks if the TagType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagType{}

// TagType Defines the Tag Type
type TagType struct {
	TagName  string  `json:"tagName"`
	KeyType  KeyType `json:"keyType"`
	Sort     *bool   `json:"sort,omitempty"`
	Presence *bool   `json:"presence,omitempty"`
}

// NewTagType instantiates a new TagType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagType(tagName string, keyType KeyType) *TagType {
	this := TagType{}
	this.TagName = tagName
	this.KeyType = keyType
	var sort bool = false
	this.Sort = &sort
	return &this
}

// NewTagTypeWithDefaults instantiates a new TagType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagTypeWithDefaults() *TagType {
	this := TagType{}
	var sort bool = false
	this.Sort = &sort
	return &this
}

// GetTagName returns the TagName field value
func (o *TagType) GetTagName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *TagType) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *TagType) SetTagName(v string) {
	o.TagName = v
}

// GetKeyType returns the KeyType field value
func (o *TagType) GetKeyType() KeyType {
	if o == nil {
		var ret KeyType
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *TagType) GetKeyTypeOk() (*KeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *TagType) SetKeyType(v KeyType) {
	o.KeyType = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TagType) GetSort() bool {
	if o == nil || IsNil(o.Sort) {
		var ret bool
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagType) GetSortOk() (*bool, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TagType) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given bool and assigns it to the Sort field.
func (o *TagType) SetSort(v bool) {
	o.Sort = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *TagType) GetPresence() bool {
	if o == nil || IsNil(o.Presence) {
		var ret bool
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagType) GetPresenceOk() (*bool, bool) {
	if o == nil || IsNil(o.Presence) {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *TagType) HasPresence() bool {
	if o != nil && !IsNil(o.Presence) {
		return true
	}

	return false
}

// SetPresence gets a reference to the given bool and assigns it to the Presence field.
func (o *TagType) SetPresence(v bool) {
	o.Presence = &v
}

func (o TagType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tagName"] = o.TagName
	toSerialize["keyType"] = o.KeyType
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Presence) {
		toSerialize["presence"] = o.Presence
	}
	return toSerialize, nil
}

type NullableTagType struct {
	value *TagType
	isSet bool
}

func (v NullableTagType) Get() *TagType {
	return v.value
}

func (v *NullableTagType) Set(val *TagType) {
	v.value = val
	v.isSet = true
}

func (v NullableTagType) IsSet() bool {
	return v.isSet
}

func (v *NullableTagType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagType(val *TagType) *NullableTagType {
	return &NullableTagType{value: val, isSet: true}
}

func (v NullableTagType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
