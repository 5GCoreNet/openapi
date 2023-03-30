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

// checks if the RecordNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordNotification{}

// RecordNotification Definition of a notification on a record
type RecordNotification struct {
	Descriptor NotificationDescription `json:"descriptor"`
	Meta RecordMeta `json:"meta"`
	// list of opaque Block's in this Record
	Blocks []interface{} `json:"blocks,omitempty"`
}

// NewRecordNotification instantiates a new RecordNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordNotification(descriptor NotificationDescription, meta RecordMeta) *RecordNotification {
	this := RecordNotification{}
	this.Descriptor = descriptor
	this.Meta = meta
	return &this
}

// NewRecordNotificationWithDefaults instantiates a new RecordNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordNotificationWithDefaults() *RecordNotification {
	this := RecordNotification{}
	return &this
}

// GetDescriptor returns the Descriptor field value
func (o *RecordNotification) GetDescriptor() NotificationDescription {
	if o == nil {
		var ret NotificationDescription
		return ret
	}

	return o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value
// and a boolean to check if the value has been set.
func (o *RecordNotification) GetDescriptorOk() (*NotificationDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descriptor, true
}

// SetDescriptor sets field value
func (o *RecordNotification) SetDescriptor(v NotificationDescription) {
	o.Descriptor = v
}

// GetMeta returns the Meta field value
func (o *RecordNotification) GetMeta() RecordMeta {
	if o == nil {
		var ret RecordMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *RecordNotification) GetMetaOk() (*RecordMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *RecordNotification) SetMeta(v RecordMeta) {
	o.Meta = v
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *RecordNotification) GetBlocks() []interface{} {
	if o == nil || IsNil(o.Blocks) {
		var ret []interface{}
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordNotification) GetBlocksOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *RecordNotification) HasBlocks() bool {
	if o != nil && !IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []interface{} and assigns it to the Blocks field.
func (o *RecordNotification) SetBlocks(v []interface{}) {
	o.Blocks = v
}

func (o RecordNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptor"] = o.Descriptor
	toSerialize["meta"] = o.Meta
	if !IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	return toSerialize, nil
}

type NullableRecordNotification struct {
	value *RecordNotification
	isSet bool
}

func (v NullableRecordNotification) Get() *RecordNotification {
	return v.value
}

func (v *NullableRecordNotification) Set(val *RecordNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordNotification(val *RecordNotification) *NullableRecordNotification {
	return &NullableRecordNotification{value: val, isSet: true}
}

func (v NullableRecordNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


