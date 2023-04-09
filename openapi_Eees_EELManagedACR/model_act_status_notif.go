/*
EES EEL Managed ACR Service

EES EEL Managed ACR Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EELManagedACR

import (
	"encoding/json"
)

// checks if the ACTStatusNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACTStatusNotif{}

// ACTStatusNotif Represents an ACT status notification.
type ACTStatusNotif struct {
	// Subscription identifier.
	SubscriptionId string    `json:"subscriptionId"`
	ActStatus      ACTResult `json:"actStatus"`
}

// NewACTStatusNotif instantiates a new ACTStatusNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACTStatusNotif(subscriptionId string, actStatus ACTResult) *ACTStatusNotif {
	this := ACTStatusNotif{}
	this.SubscriptionId = subscriptionId
	this.ActStatus = actStatus
	return &this
}

// NewACTStatusNotifWithDefaults instantiates a new ACTStatusNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACTStatusNotifWithDefaults() *ACTStatusNotif {
	this := ACTStatusNotif{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ACTStatusNotif) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ACTStatusNotif) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ACTStatusNotif) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetActStatus returns the ActStatus field value
func (o *ACTStatusNotif) GetActStatus() ACTResult {
	if o == nil {
		var ret ACTResult
		return ret
	}

	return o.ActStatus
}

// GetActStatusOk returns a tuple with the ActStatus field value
// and a boolean to check if the value has been set.
func (o *ACTStatusNotif) GetActStatusOk() (*ACTResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActStatus, true
}

// SetActStatus sets field value
func (o *ACTStatusNotif) SetActStatus(v ACTResult) {
	o.ActStatus = v
}

func (o ACTStatusNotif) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACTStatusNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["actStatus"] = o.ActStatus
	return toSerialize, nil
}

type NullableACTStatusNotif struct {
	value *ACTStatusNotif
	isSet bool
}

func (v NullableACTStatusNotif) Get() *ACTStatusNotif {
	return v.value
}

func (v *NullableACTStatusNotif) Set(val *ACTStatusNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableACTStatusNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableACTStatusNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACTStatusNotif(val *ACTStatusNotif) *NullableACTStatusNotif {
	return &NullableACTStatusNotif{value: val, isSet: true}
}

func (v NullableACTStatusNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACTStatusNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
