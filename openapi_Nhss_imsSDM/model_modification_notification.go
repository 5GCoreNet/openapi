/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the ModificationNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModificationNotification{}

// ModificationNotification struct for ModificationNotification
type ModificationNotification struct {
	NotifyItems []NotifyItem `json:"notifyItems"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewModificationNotification instantiates a new ModificationNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModificationNotification(notifyItems []NotifyItem) *ModificationNotification {
	this := ModificationNotification{}
	this.NotifyItems = notifyItems
	return &this
}

// NewModificationNotificationWithDefaults instantiates a new ModificationNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationNotificationWithDefaults() *ModificationNotification {
	this := ModificationNotification{}
	return &this
}

// GetNotifyItems returns the NotifyItems field value
func (o *ModificationNotification) GetNotifyItems() []NotifyItem {
	if o == nil {
		var ret []NotifyItem
		return ret
	}

	return o.NotifyItems
}

// GetNotifyItemsOk returns a tuple with the NotifyItems field value
// and a boolean to check if the value has been set.
func (o *ModificationNotification) GetNotifyItemsOk() ([]NotifyItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyItems, true
}

// SetNotifyItems sets field value
func (o *ModificationNotification) SetNotifyItems(v []NotifyItem) {
	o.NotifyItems = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *ModificationNotification) GetSubscriptionId() string {
	if o == nil || isNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModificationNotification) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *ModificationNotification) HasSubscriptionId() bool {
	if o != nil && !isNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *ModificationNotification) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o ModificationNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModificationNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifyItems"] = o.NotifyItems
	if !isNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableModificationNotification struct {
	value *ModificationNotification
	isSet bool
}

func (v NullableModificationNotification) Get() *ModificationNotification {
	return v.value
}

func (v *NullableModificationNotification) Set(val *ModificationNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableModificationNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableModificationNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModificationNotification(val *ModificationNotification) *NullableModificationNotification {
	return &NullableModificationNotification{value: val, isSet: true}
}

func (v NullableModificationNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModificationNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


