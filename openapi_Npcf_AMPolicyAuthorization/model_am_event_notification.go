/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyAuthorization

import (
	"encoding/json"
)

// checks if the AmEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmEventNotification{}

// AmEventNotification Describes the notification of a subscription.
type AmEventNotification struct {
	Event      AmEvent                  `json:"event"`
	AppliedCov *ServiceAreaCoverageInfo `json:"appliedCov,omitempty"`
	PduidInfo  *PduidInformation        `json:"pduidInfo,omitempty"`
}

// NewAmEventNotification instantiates a new AmEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmEventNotification(event AmEvent) *AmEventNotification {
	this := AmEventNotification{}
	this.Event = event
	return &this
}

// NewAmEventNotificationWithDefaults instantiates a new AmEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmEventNotificationWithDefaults() *AmEventNotification {
	this := AmEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *AmEventNotification) GetEvent() AmEvent {
	if o == nil {
		var ret AmEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AmEventNotification) GetEventOk() (*AmEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AmEventNotification) SetEvent(v AmEvent) {
	o.Event = v
}

// GetAppliedCov returns the AppliedCov field value if set, zero value otherwise.
func (o *AmEventNotification) GetAppliedCov() ServiceAreaCoverageInfo {
	if o == nil || IsNil(o.AppliedCov) {
		var ret ServiceAreaCoverageInfo
		return ret
	}
	return *o.AppliedCov
}

// GetAppliedCovOk returns a tuple with the AppliedCov field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventNotification) GetAppliedCovOk() (*ServiceAreaCoverageInfo, bool) {
	if o == nil || IsNil(o.AppliedCov) {
		return nil, false
	}
	return o.AppliedCov, true
}

// HasAppliedCov returns a boolean if a field has been set.
func (o *AmEventNotification) HasAppliedCov() bool {
	if o != nil && !IsNil(o.AppliedCov) {
		return true
	}

	return false
}

// SetAppliedCov gets a reference to the given ServiceAreaCoverageInfo and assigns it to the AppliedCov field.
func (o *AmEventNotification) SetAppliedCov(v ServiceAreaCoverageInfo) {
	o.AppliedCov = &v
}

// GetPduidInfo returns the PduidInfo field value if set, zero value otherwise.
func (o *AmEventNotification) GetPduidInfo() PduidInformation {
	if o == nil || IsNil(o.PduidInfo) {
		var ret PduidInformation
		return ret
	}
	return *o.PduidInfo
}

// GetPduidInfoOk returns a tuple with the PduidInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventNotification) GetPduidInfoOk() (*PduidInformation, bool) {
	if o == nil || IsNil(o.PduidInfo) {
		return nil, false
	}
	return o.PduidInfo, true
}

// HasPduidInfo returns a boolean if a field has been set.
func (o *AmEventNotification) HasPduidInfo() bool {
	if o != nil && !IsNil(o.PduidInfo) {
		return true
	}

	return false
}

// SetPduidInfo gets a reference to the given PduidInformation and assigns it to the PduidInfo field.
func (o *AmEventNotification) SetPduidInfo(v PduidInformation) {
	o.PduidInfo = &v
}

func (o AmEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.AppliedCov) {
		toSerialize["appliedCov"] = o.AppliedCov
	}
	if !IsNil(o.PduidInfo) {
		toSerialize["pduidInfo"] = o.PduidInfo
	}
	return toSerialize, nil
}

type NullableAmEventNotification struct {
	value *AmEventNotification
	isSet bool
}

func (v NullableAmEventNotification) Get() *AmEventNotification {
	return v.value
}

func (v *NullableAmEventNotification) Set(val *AmEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEventNotification(val *AmEventNotification) *NullableAmEventNotification {
	return &NullableAmEventNotification{value: val, isSet: true}
}

func (v NullableAmEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
