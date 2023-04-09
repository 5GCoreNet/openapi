/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the NotifyMoiChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyMoiChanges{}

// NotifyMoiChanges struct for NotifyMoiChanges
type NotifyMoiChanges struct {
	NotificationHeader
	MoiChanges []MoiChange `json:"moiChanges"`
}

// NewNotifyMoiChanges instantiates a new NotifyMoiChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyMoiChanges(moiChanges []MoiChange, href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string) *NotifyMoiChanges {
	this := NotifyMoiChanges{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.MoiChanges = moiChanges
	return &this
}

// NewNotifyMoiChangesWithDefaults instantiates a new NotifyMoiChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyMoiChangesWithDefaults() *NotifyMoiChanges {
	this := NotifyMoiChanges{}
	return &this
}

// GetMoiChanges returns the MoiChanges field value
func (o *NotifyMoiChanges) GetMoiChanges() []MoiChange {
	if o == nil {
		var ret []MoiChange
		return ret
	}

	return o.MoiChanges
}

// GetMoiChangesOk returns a tuple with the MoiChanges field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiChanges) GetMoiChangesOk() ([]MoiChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoiChanges, true
}

// SetMoiChanges sets field value
func (o *NotifyMoiChanges) SetMoiChanges(v []MoiChange) {
	o.MoiChanges = v
}

func (o NotifyMoiChanges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyMoiChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationHeader, errNotificationHeader := json.Marshal(o.NotificationHeader)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	errNotificationHeader = json.Unmarshal([]byte(serializedNotificationHeader), &toSerialize)
	if errNotificationHeader != nil {
		return map[string]interface{}{}, errNotificationHeader
	}
	toSerialize["moiChanges"] = o.MoiChanges
	return toSerialize, nil
}

type NullableNotifyMoiChanges struct {
	value *NotifyMoiChanges
	isSet bool
}

func (v NullableNotifyMoiChanges) Get() *NotifyMoiChanges {
	return v.value
}

func (v *NullableNotifyMoiChanges) Set(val *NotifyMoiChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyMoiChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyMoiChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyMoiChanges(val *NotifyMoiChanges) *NullableNotifyMoiChanges {
	return &NullableNotifyMoiChanges{value: val, isSet: true}
}

func (v NullableNotifyMoiChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyMoiChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
