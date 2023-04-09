/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the NotificationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationData{}

// NotificationData Data sent in notifications from NRF to subscribed NF Instances
type NotificationData struct {
	Event NotificationEventType `json:"event"`
	// String providing an URI formatted according to RFC 3986.
	NfInstanceUri       string               `json:"nfInstanceUri"`
	NfProfile           *NFProfile           `json:"nfProfile,omitempty"`
	ProfileChanges      []ChangeItem         `json:"profileChanges,omitempty"`
	ConditionEvent      *ConditionEventType  `json:"conditionEvent,omitempty"`
	SubscriptionContext *SubscriptionContext `json:"subscriptionContext,omitempty"`
	CompleteNfProfile   *NFProfile           `json:"completeNfProfile,omitempty"`
}

// NewNotificationData instantiates a new NotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationData(event NotificationEventType, nfInstanceUri string) *NotificationData {
	this := NotificationData{}
	return &this
}

// NewNotificationDataWithDefaults instantiates a new NotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDataWithDefaults() *NotificationData {
	this := NotificationData{}
	return &this
}

// GetEvent returns the Event field value
func (o *NotificationData) GetEvent() NotificationEventType {
	if o == nil {
		var ret NotificationEventType
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NotificationData) GetEventOk() (*NotificationEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NotificationData) SetEvent(v NotificationEventType) {
	o.Event = v
}

// GetNfInstanceUri returns the NfInstanceUri field value
func (o *NotificationData) GetNfInstanceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceUri
}

// GetNfInstanceUriOk returns a tuple with the NfInstanceUri field value
// and a boolean to check if the value has been set.
func (o *NotificationData) GetNfInstanceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceUri, true
}

// SetNfInstanceUri sets field value
func (o *NotificationData) SetNfInstanceUri(v string) {
	o.NfInstanceUri = v
}

// GetNfProfile returns the NfProfile field value if set, zero value otherwise.
func (o *NotificationData) GetNfProfile() NFProfile {
	if o == nil || IsNil(o.NfProfile) {
		var ret NFProfile
		return ret
	}
	return *o.NfProfile
}

// GetNfProfileOk returns a tuple with the NfProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationData) GetNfProfileOk() (*NFProfile, bool) {
	if o == nil || IsNil(o.NfProfile) {
		return nil, false
	}
	return o.NfProfile, true
}

// HasNfProfile returns a boolean if a field has been set.
func (o *NotificationData) HasNfProfile() bool {
	if o != nil && !IsNil(o.NfProfile) {
		return true
	}

	return false
}

// SetNfProfile gets a reference to the given NFProfile and assigns it to the NfProfile field.
func (o *NotificationData) SetNfProfile(v NFProfile) {
	o.NfProfile = &v
}

// GetProfileChanges returns the ProfileChanges field value if set, zero value otherwise.
func (o *NotificationData) GetProfileChanges() []ChangeItem {
	if o == nil || IsNil(o.ProfileChanges) {
		var ret []ChangeItem
		return ret
	}
	return o.ProfileChanges
}

// GetProfileChangesOk returns a tuple with the ProfileChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationData) GetProfileChangesOk() ([]ChangeItem, bool) {
	if o == nil || IsNil(o.ProfileChanges) {
		return nil, false
	}
	return o.ProfileChanges, true
}

// HasProfileChanges returns a boolean if a field has been set.
func (o *NotificationData) HasProfileChanges() bool {
	if o != nil && !IsNil(o.ProfileChanges) {
		return true
	}

	return false
}

// SetProfileChanges gets a reference to the given []ChangeItem and assigns it to the ProfileChanges field.
func (o *NotificationData) SetProfileChanges(v []ChangeItem) {
	o.ProfileChanges = v
}

// GetConditionEvent returns the ConditionEvent field value if set, zero value otherwise.
func (o *NotificationData) GetConditionEvent() ConditionEventType {
	if o == nil || IsNil(o.ConditionEvent) {
		var ret ConditionEventType
		return ret
	}
	return *o.ConditionEvent
}

// GetConditionEventOk returns a tuple with the ConditionEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationData) GetConditionEventOk() (*ConditionEventType, bool) {
	if o == nil || IsNil(o.ConditionEvent) {
		return nil, false
	}
	return o.ConditionEvent, true
}

// HasConditionEvent returns a boolean if a field has been set.
func (o *NotificationData) HasConditionEvent() bool {
	if o != nil && !IsNil(o.ConditionEvent) {
		return true
	}

	return false
}

// SetConditionEvent gets a reference to the given ConditionEventType and assigns it to the ConditionEvent field.
func (o *NotificationData) SetConditionEvent(v ConditionEventType) {
	o.ConditionEvent = &v
}

// GetSubscriptionContext returns the SubscriptionContext field value if set, zero value otherwise.
func (o *NotificationData) GetSubscriptionContext() SubscriptionContext {
	if o == nil || IsNil(o.SubscriptionContext) {
		var ret SubscriptionContext
		return ret
	}
	return *o.SubscriptionContext
}

// GetSubscriptionContextOk returns a tuple with the SubscriptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationData) GetSubscriptionContextOk() (*SubscriptionContext, bool) {
	if o == nil || IsNil(o.SubscriptionContext) {
		return nil, false
	}
	return o.SubscriptionContext, true
}

// HasSubscriptionContext returns a boolean if a field has been set.
func (o *NotificationData) HasSubscriptionContext() bool {
	if o != nil && !IsNil(o.SubscriptionContext) {
		return true
	}

	return false
}

// SetSubscriptionContext gets a reference to the given SubscriptionContext and assigns it to the SubscriptionContext field.
func (o *NotificationData) SetSubscriptionContext(v SubscriptionContext) {
	o.SubscriptionContext = &v
}

// GetCompleteNfProfile returns the CompleteNfProfile field value if set, zero value otherwise.
func (o *NotificationData) GetCompleteNfProfile() NFProfile {
	if o == nil || IsNil(o.CompleteNfProfile) {
		var ret NFProfile
		return ret
	}
	return *o.CompleteNfProfile
}

// GetCompleteNfProfileOk returns a tuple with the CompleteNfProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationData) GetCompleteNfProfileOk() (*NFProfile, bool) {
	if o == nil || IsNil(o.CompleteNfProfile) {
		return nil, false
	}
	return o.CompleteNfProfile, true
}

// HasCompleteNfProfile returns a boolean if a field has been set.
func (o *NotificationData) HasCompleteNfProfile() bool {
	if o != nil && !IsNil(o.CompleteNfProfile) {
		return true
	}

	return false
}

// SetCompleteNfProfile gets a reference to the given NFProfile and assigns it to the CompleteNfProfile field.
func (o *NotificationData) SetCompleteNfProfile(v NFProfile) {
	o.CompleteNfProfile = &v
}

func (o NotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["nfInstanceUri"] = o.NfInstanceUri
	if !IsNil(o.NfProfile) {
		toSerialize["nfProfile"] = o.NfProfile
	}
	if !IsNil(o.ProfileChanges) {
		toSerialize["profileChanges"] = o.ProfileChanges
	}
	if !IsNil(o.ConditionEvent) {
		toSerialize["conditionEvent"] = o.ConditionEvent
	}
	if !IsNil(o.SubscriptionContext) {
		toSerialize["subscriptionContext"] = o.SubscriptionContext
	}
	if !IsNil(o.CompleteNfProfile) {
		toSerialize["completeNfProfile"] = o.CompleteNfProfile
	}
	return toSerialize, nil
}

type NullableNotificationData struct {
	value *NotificationData
	isSet bool
}

func (v NullableNotificationData) Get() *NotificationData {
	return v.value
}

func (v *NullableNotificationData) Set(val *NotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationData(val *NotificationData) *NullableNotificationData {
	return &NullableNotificationData{value: val, isSet: true}
}

func (v NullableNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
