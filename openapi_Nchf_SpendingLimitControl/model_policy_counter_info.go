/*
Nchf_SpendingLimitControl

Nchf Spending Limit Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_SpendingLimitControl

import (
	"encoding/json"
)

// checks if the PolicyCounterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyCounterInfo{}

// PolicyCounterInfo Represents the data structure presenting the policy counter status.
type PolicyCounterInfo struct {
	// Identifies a policy counter.
	PolicyCounterId string `json:"policyCounterId"`
	// Identifies the policy counter status applicable for a specific policy counter identified by the policyCounterId. The values (e.g. valid, invalid or any other status) are not specified. The interpretation and actions related to the defined values are out of scope of 3GPP.
	CurrentStatus string `json:"currentStatus"`
	// Provides the pending policy counter status.
	PenPolCounterStatuses []PendingPolicyCounterStatus `json:"penPolCounterStatuses,omitempty"`
}

// NewPolicyCounterInfo instantiates a new PolicyCounterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyCounterInfo(policyCounterId string, currentStatus string) *PolicyCounterInfo {
	this := PolicyCounterInfo{}
	this.PolicyCounterId = policyCounterId
	this.CurrentStatus = currentStatus
	return &this
}

// NewPolicyCounterInfoWithDefaults instantiates a new PolicyCounterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyCounterInfoWithDefaults() *PolicyCounterInfo {
	this := PolicyCounterInfo{}
	return &this
}

// GetPolicyCounterId returns the PolicyCounterId field value
func (o *PolicyCounterInfo) GetPolicyCounterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyCounterId
}

// GetPolicyCounterIdOk returns a tuple with the PolicyCounterId field value
// and a boolean to check if the value has been set.
func (o *PolicyCounterInfo) GetPolicyCounterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyCounterId, true
}

// SetPolicyCounterId sets field value
func (o *PolicyCounterInfo) SetPolicyCounterId(v string) {
	o.PolicyCounterId = v
}

// GetCurrentStatus returns the CurrentStatus field value
func (o *PolicyCounterInfo) GetCurrentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value
// and a boolean to check if the value has been set.
func (o *PolicyCounterInfo) GetCurrentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStatus, true
}

// SetCurrentStatus sets field value
func (o *PolicyCounterInfo) SetCurrentStatus(v string) {
	o.CurrentStatus = v
}

// GetPenPolCounterStatuses returns the PenPolCounterStatuses field value if set, zero value otherwise.
func (o *PolicyCounterInfo) GetPenPolCounterStatuses() []PendingPolicyCounterStatus {
	if o == nil || IsNil(o.PenPolCounterStatuses) {
		var ret []PendingPolicyCounterStatus
		return ret
	}
	return o.PenPolCounterStatuses
}

// GetPenPolCounterStatusesOk returns a tuple with the PenPolCounterStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyCounterInfo) GetPenPolCounterStatusesOk() ([]PendingPolicyCounterStatus, bool) {
	if o == nil || IsNil(o.PenPolCounterStatuses) {
		return nil, false
	}
	return o.PenPolCounterStatuses, true
}

// HasPenPolCounterStatuses returns a boolean if a field has been set.
func (o *PolicyCounterInfo) HasPenPolCounterStatuses() bool {
	if o != nil && !IsNil(o.PenPolCounterStatuses) {
		return true
	}

	return false
}

// SetPenPolCounterStatuses gets a reference to the given []PendingPolicyCounterStatus and assigns it to the PenPolCounterStatuses field.
func (o *PolicyCounterInfo) SetPenPolCounterStatuses(v []PendingPolicyCounterStatus) {
	o.PenPolCounterStatuses = v
}

func (o PolicyCounterInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyCounterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policyCounterId"] = o.PolicyCounterId
	toSerialize["currentStatus"] = o.CurrentStatus
	if !IsNil(o.PenPolCounterStatuses) {
		toSerialize["penPolCounterStatuses"] = o.PenPolCounterStatuses
	}
	return toSerialize, nil
}

type NullablePolicyCounterInfo struct {
	value *PolicyCounterInfo
	isSet bool
}

func (v NullablePolicyCounterInfo) Get() *PolicyCounterInfo {
	return v.value
}

func (v *NullablePolicyCounterInfo) Set(val *PolicyCounterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyCounterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyCounterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyCounterInfo(val *PolicyCounterInfo) *NullablePolicyCounterInfo {
	return &NullablePolicyCounterInfo{value: val, isSet: true}
}

func (v NullablePolicyCounterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyCounterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
