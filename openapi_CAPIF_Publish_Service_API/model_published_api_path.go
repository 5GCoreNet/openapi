/*
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Publish_Service_API

import (
	"encoding/json"
)

// checks if the PublishedApiPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishedApiPath{}

// PublishedApiPath Represents the published API path within the same CAPIF provider domain.
type PublishedApiPath struct {
	// A list of CCF identifiers where the service API is already published.
	CcfIds []string `json:"ccfIds,omitempty"`
}

// NewPublishedApiPath instantiates a new PublishedApiPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishedApiPath() *PublishedApiPath {
	this := PublishedApiPath{}
	return &this
}

// NewPublishedApiPathWithDefaults instantiates a new PublishedApiPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishedApiPathWithDefaults() *PublishedApiPath {
	this := PublishedApiPath{}
	return &this
}

// GetCcfIds returns the CcfIds field value if set, zero value otherwise.
func (o *PublishedApiPath) GetCcfIds() []string {
	if o == nil || IsNil(o.CcfIds) {
		var ret []string
		return ret
	}
	return o.CcfIds
}

// GetCcfIdsOk returns a tuple with the CcfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishedApiPath) GetCcfIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CcfIds) {
		return nil, false
	}
	return o.CcfIds, true
}

// HasCcfIds returns a boolean if a field has been set.
func (o *PublishedApiPath) HasCcfIds() bool {
	if o != nil && !IsNil(o.CcfIds) {
		return true
	}

	return false
}

// SetCcfIds gets a reference to the given []string and assigns it to the CcfIds field.
func (o *PublishedApiPath) SetCcfIds(v []string) {
	o.CcfIds = v
}

func (o PublishedApiPath) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishedApiPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CcfIds) {
		toSerialize["ccfIds"] = o.CcfIds
	}
	return toSerialize, nil
}

type NullablePublishedApiPath struct {
	value *PublishedApiPath
	isSet bool
}

func (v NullablePublishedApiPath) Get() *PublishedApiPath {
	return v.value
}

func (v *NullablePublishedApiPath) Set(val *PublishedApiPath) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishedApiPath) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishedApiPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishedApiPath(val *PublishedApiPath) *NullablePublishedApiPath {
	return &NullablePublishedApiPath{value: val, isSet: true}
}

func (v NullablePublishedApiPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishedApiPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
