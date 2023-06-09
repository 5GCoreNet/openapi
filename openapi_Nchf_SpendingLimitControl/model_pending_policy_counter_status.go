/*
Nchf_SpendingLimitControl

Nchf Spending Limit Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_SpendingLimitControl

import (
	"encoding/json"
	"time"
)

// checks if the PendingPolicyCounterStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingPolicyCounterStatus{}

// PendingPolicyCounterStatus Represents the data structure presenting the pending policy counter status.
type PendingPolicyCounterStatus struct {
	// Identifies the policy counter status applicable for a specific policy counter identified by the policyCounterId. The values (e.g. valid, invalid or any other status) are not specified. The interpretation and actions related to the defined values are out of scope of 3GPP.
	PolicyCounterStatus string `json:"policyCounterStatus"`
	// string with format 'date-time' as defined in OpenAPI.
	ActivationTime time.Time `json:"activationTime"`
}

// NewPendingPolicyCounterStatus instantiates a new PendingPolicyCounterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingPolicyCounterStatus(policyCounterStatus string, activationTime time.Time) *PendingPolicyCounterStatus {
	this := PendingPolicyCounterStatus{}
	this.PolicyCounterStatus = policyCounterStatus
	this.ActivationTime = activationTime
	return &this
}

// NewPendingPolicyCounterStatusWithDefaults instantiates a new PendingPolicyCounterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingPolicyCounterStatusWithDefaults() *PendingPolicyCounterStatus {
	this := PendingPolicyCounterStatus{}
	return &this
}

// GetPolicyCounterStatus returns the PolicyCounterStatus field value
func (o *PendingPolicyCounterStatus) GetPolicyCounterStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyCounterStatus
}

// GetPolicyCounterStatusOk returns a tuple with the PolicyCounterStatus field value
// and a boolean to check if the value has been set.
func (o *PendingPolicyCounterStatus) GetPolicyCounterStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyCounterStatus, true
}

// SetPolicyCounterStatus sets field value
func (o *PendingPolicyCounterStatus) SetPolicyCounterStatus(v string) {
	o.PolicyCounterStatus = v
}

// GetActivationTime returns the ActivationTime field value
func (o *PendingPolicyCounterStatus) GetActivationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActivationTime
}

// GetActivationTimeOk returns a tuple with the ActivationTime field value
// and a boolean to check if the value has been set.
func (o *PendingPolicyCounterStatus) GetActivationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationTime, true
}

// SetActivationTime sets field value
func (o *PendingPolicyCounterStatus) SetActivationTime(v time.Time) {
	o.ActivationTime = v
}

func (o PendingPolicyCounterStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingPolicyCounterStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policyCounterStatus"] = o.PolicyCounterStatus
	toSerialize["activationTime"] = o.ActivationTime
	return toSerialize, nil
}

type NullablePendingPolicyCounterStatus struct {
	value *PendingPolicyCounterStatus
	isSet bool
}

func (v NullablePendingPolicyCounterStatus) Get() *PendingPolicyCounterStatus {
	return v.value
}

func (v *NullablePendingPolicyCounterStatus) Set(val *PendingPolicyCounterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingPolicyCounterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingPolicyCounterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingPolicyCounterStatus(val *PendingPolicyCounterStatus) *NullablePendingPolicyCounterStatus {
	return &NullablePendingPolicyCounterStatus{value: val, isSet: true}
}

func (v NullablePendingPolicyCounterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingPolicyCounterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
