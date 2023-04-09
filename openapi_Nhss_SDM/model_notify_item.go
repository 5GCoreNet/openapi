/*
Nhss_SDM

HSS Subscriber Data Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_SDM

import (
	"encoding/json"
)

// checks if the NotifyItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyItem{}

// NotifyItem Indicates changes on a resource.
type NotifyItem struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceId string       `json:"resourceId"`
	Changes    []ChangeItem `json:"changes"`
}

// NewNotifyItem instantiates a new NotifyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyItem(resourceId string, changes []ChangeItem) *NotifyItem {
	this := NotifyItem{}
	this.ResourceId = resourceId
	this.Changes = changes
	return &this
}

// NewNotifyItemWithDefaults instantiates a new NotifyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyItemWithDefaults() *NotifyItem {
	this := NotifyItem{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *NotifyItem) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *NotifyItem) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *NotifyItem) SetResourceId(v string) {
	o.ResourceId = v
}

// GetChanges returns the Changes field value
func (o *NotifyItem) GetChanges() []ChangeItem {
	if o == nil {
		var ret []ChangeItem
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *NotifyItem) GetChangesOk() ([]ChangeItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *NotifyItem) SetChanges(v []ChangeItem) {
	o.Changes = v
}

func (o NotifyItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["changes"] = o.Changes
	return toSerialize, nil
}

type NullableNotifyItem struct {
	value *NotifyItem
	isSet bool
}

func (v NullableNotifyItem) Get() *NotifyItem {
	return v.value
}

func (v *NullableNotifyItem) Set(val *NotifyItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyItem(val *NotifyItem) *NullableNotifyItem {
	return &NullableNotifyItem{value: val, isSet: true}
}

func (v NullableNotifyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
