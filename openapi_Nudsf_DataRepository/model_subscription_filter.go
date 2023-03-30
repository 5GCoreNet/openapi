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

// checks if the SubscriptionFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionFilter{}

// SubscriptionFilter A subscription filter
type SubscriptionFilter struct {
	// list of resources applicable to the subscription
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`
	// list of resources applicable to the subscription
	Operations []RecordOperation `json:"operations,omitempty"`
}

// NewSubscriptionFilter instantiates a new SubscriptionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionFilter() *SubscriptionFilter {
	this := SubscriptionFilter{}
	return &this
}

// NewSubscriptionFilterWithDefaults instantiates a new SubscriptionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionFilterWithDefaults() *SubscriptionFilter {
	this := SubscriptionFilter{}
	return &this
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value if set, zero value otherwise.
func (o *SubscriptionFilter) GetMonitoredResourceUris() []string {
	if o == nil || IsNil(o.MonitoredResourceUris) {
		var ret []string
		return ret
	}
	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilter) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.MonitoredResourceUris) {
		return nil, false
	}
	return o.MonitoredResourceUris, true
}

// HasMonitoredResourceUris returns a boolean if a field has been set.
func (o *SubscriptionFilter) HasMonitoredResourceUris() bool {
	if o != nil && !IsNil(o.MonitoredResourceUris) {
		return true
	}

	return false
}

// SetMonitoredResourceUris gets a reference to the given []string and assigns it to the MonitoredResourceUris field.
func (o *SubscriptionFilter) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *SubscriptionFilter) GetOperations() []RecordOperation {
	if o == nil || IsNil(o.Operations) {
		var ret []RecordOperation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilter) GetOperationsOk() ([]RecordOperation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *SubscriptionFilter) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []RecordOperation and assigns it to the Operations field.
func (o *SubscriptionFilter) SetOperations(v []RecordOperation) {
	o.Operations = v
}

func (o SubscriptionFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MonitoredResourceUris) {
		toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	return toSerialize, nil
}

type NullableSubscriptionFilter struct {
	value *SubscriptionFilter
	isSet bool
}

func (v NullableSubscriptionFilter) Get() *SubscriptionFilter {
	return v.value
}

func (v *NullableSubscriptionFilter) Set(val *SubscriptionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionFilter(val *SubscriptionFilter) *NullableSubscriptionFilter {
	return &NullableSubscriptionFilter{value: val, isSet: true}
}

func (v NullableSubscriptionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


