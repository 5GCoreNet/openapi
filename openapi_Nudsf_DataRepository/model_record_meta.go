/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
	"time"
)

// checks if the RecordMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordMeta{}

// RecordMeta Meta data of a Record
type RecordMeta struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ttl *time.Time `json:"ttl,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference *string `json:"callbackReference,omitempty"`
	// A dictionary of {\"tagName\": [ \"tagValue\", ...] }. A tag name can be used to retrieve a Record. The tagValue are unique.
	Tags *map[string][]string `json:"tags,omitempty"`
}

// NewRecordMeta instantiates a new RecordMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordMeta() *RecordMeta {
	this := RecordMeta{}
	return &this
}

// NewRecordMetaWithDefaults instantiates a new RecordMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordMetaWithDefaults() *RecordMeta {
	this := RecordMeta{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RecordMeta) GetTtl() time.Time {
	if o == nil || isNil(o.Ttl) {
		var ret time.Time
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMeta) GetTtlOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RecordMeta) HasTtl() bool {
	if o != nil && !isNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given time.Time and assigns it to the Ttl field.
func (o *RecordMeta) SetTtl(v time.Time) {
	o.Ttl = &v
}

// GetCallbackReference returns the CallbackReference field value if set, zero value otherwise.
func (o *RecordMeta) GetCallbackReference() string {
	if o == nil || isNil(o.CallbackReference) {
		var ret string
		return ret
	}
	return *o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMeta) GetCallbackReferenceOk() (*string, bool) {
	if o == nil || isNil(o.CallbackReference) {
		return nil, false
	}
	return o.CallbackReference, true
}

// HasCallbackReference returns a boolean if a field has been set.
func (o *RecordMeta) HasCallbackReference() bool {
	if o != nil && !isNil(o.CallbackReference) {
		return true
	}

	return false
}

// SetCallbackReference gets a reference to the given string and assigns it to the CallbackReference field.
func (o *RecordMeta) SetCallbackReference(v string) {
	o.CallbackReference = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RecordMeta) GetTags() map[string][]string {
	if o == nil || isNil(o.Tags) {
		var ret map[string][]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordMeta) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RecordMeta) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *RecordMeta) SetTags(v map[string][]string) {
	o.Tags = &v
}

func (o RecordMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !isNil(o.CallbackReference) {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableRecordMeta struct {
	value *RecordMeta
	isSet bool
}

func (v NullableRecordMeta) Get() *RecordMeta {
	return v.value
}

func (v *NullableRecordMeta) Set(val *RecordMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordMeta(val *RecordMeta) *NullableRecordMeta {
	return &NullableRecordMeta{value: val, isSet: true}
}

func (v NullableRecordMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


