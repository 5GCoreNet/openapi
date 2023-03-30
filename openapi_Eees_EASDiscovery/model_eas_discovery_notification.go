/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the EasDiscoveryNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDiscoveryNotification{}

// EasDiscoveryNotification Notification of EAS discovery information.
type EasDiscoveryNotification struct {
	// Identifier of the individual service provisioning subscription for which the service provisioning notification is delivered.
	SubId string `json:"subId"`
	EventType EASDiscEventIDs `json:"eventType"`
	// List of EAS discovery information.
	DiscoveredEas []DiscoveredEas `json:"discoveredEas"`
}

// NewEasDiscoveryNotification instantiates a new EasDiscoveryNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDiscoveryNotification(subId string, eventType EASDiscEventIDs, discoveredEas []DiscoveredEas) *EasDiscoveryNotification {
	this := EasDiscoveryNotification{}
	this.SubId = subId
	this.EventType = eventType
	this.DiscoveredEas = discoveredEas
	return &this
}

// NewEasDiscoveryNotificationWithDefaults instantiates a new EasDiscoveryNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDiscoveryNotificationWithDefaults() *EasDiscoveryNotification {
	this := EasDiscoveryNotification{}
	return &this
}

// GetSubId returns the SubId field value
func (o *EasDiscoveryNotification) GetSubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value
// and a boolean to check if the value has been set.
func (o *EasDiscoveryNotification) GetSubIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubId, true
}

// SetSubId sets field value
func (o *EasDiscoveryNotification) SetSubId(v string) {
	o.SubId = v
}

// GetEventType returns the EventType field value
func (o *EasDiscoveryNotification) GetEventType() EASDiscEventIDs {
	if o == nil {
		var ret EASDiscEventIDs
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EasDiscoveryNotification) GetEventTypeOk() (*EASDiscEventIDs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EasDiscoveryNotification) SetEventType(v EASDiscEventIDs) {
	o.EventType = v
}

// GetDiscoveredEas returns the DiscoveredEas field value
func (o *EasDiscoveryNotification) GetDiscoveredEas() []DiscoveredEas {
	if o == nil {
		var ret []DiscoveredEas
		return ret
	}

	return o.DiscoveredEas
}

// GetDiscoveredEasOk returns a tuple with the DiscoveredEas field value
// and a boolean to check if the value has been set.
func (o *EasDiscoveryNotification) GetDiscoveredEasOk() ([]DiscoveredEas, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoveredEas, true
}

// SetDiscoveredEas sets field value
func (o *EasDiscoveryNotification) SetDiscoveredEas(v []DiscoveredEas) {
	o.DiscoveredEas = v
}

func (o EasDiscoveryNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDiscoveryNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subId"] = o.SubId
	toSerialize["eventType"] = o.EventType
	toSerialize["discoveredEas"] = o.DiscoveredEas
	return toSerialize, nil
}

type NullableEasDiscoveryNotification struct {
	value *EasDiscoveryNotification
	isSet bool
}

func (v NullableEasDiscoveryNotification) Get() *EasDiscoveryNotification {
	return v.value
}

func (v *NullableEasDiscoveryNotification) Set(val *EasDiscoveryNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDiscoveryNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDiscoveryNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDiscoveryNotification(val *EasDiscoveryNotification) *NullableEasDiscoveryNotification {
	return &NullableEasDiscoveryNotification{value: val, isSet: true}
}

func (v NullableEasDiscoveryNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDiscoveryNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

