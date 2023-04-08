/*
3gpp-bdt

API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ResourceManagementOfBdt

import (
	"encoding/json"
)

// checks if the ExNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExNotification{}

// ExNotification Represents a Background Data Transfer notification.
type ExNotification struct {
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
	LocationArea5G *LocationArea5G `json:"locationArea5G,omitempty"`
	TimeWindow *TimeWindow `json:"timeWindow,omitempty"`
	// This IE indicates a list of the candidate transfer policies from which the AF may select a new transfer policy due to network performance degradation.
	CandPolicies []TransferPolicy `json:"candPolicies,omitempty"`
}

// NewExNotification instantiates a new ExNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExNotification(bdtRefId string) *ExNotification {
	this := ExNotification{}
	this.BdtRefId = bdtRefId
	return &this
}

// NewExNotificationWithDefaults instantiates a new ExNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExNotificationWithDefaults() *ExNotification {
	this := ExNotification{}
	return &this
}

// GetBdtRefId returns the BdtRefId field value
func (o *ExNotification) GetBdtRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value
// and a boolean to check if the value has been set.
func (o *ExNotification) GetBdtRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BdtRefId, true
}

// SetBdtRefId sets field value
func (o *ExNotification) SetBdtRefId(v string) {
	o.BdtRefId = v
}

// GetLocationArea5G returns the LocationArea5G field value if set, zero value otherwise.
func (o *ExNotification) GetLocationArea5G() LocationArea5G {
	if o == nil || isNil(o.LocationArea5G) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocationArea5G
}

// GetLocationArea5GOk returns a tuple with the LocationArea5G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExNotification) GetLocationArea5GOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.LocationArea5G) {
		return nil, false
	}
	return o.LocationArea5G, true
}

// HasLocationArea5G returns a boolean if a field has been set.
func (o *ExNotification) HasLocationArea5G() bool {
	if o != nil && !isNil(o.LocationArea5G) {
		return true
	}

	return false
}

// SetLocationArea5G gets a reference to the given LocationArea5G and assigns it to the LocationArea5G field.
func (o *ExNotification) SetLocationArea5G(v LocationArea5G) {
	o.LocationArea5G = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *ExNotification) GetTimeWindow() TimeWindow {
	if o == nil || isNil(o.TimeWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExNotification) GetTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *ExNotification) HasTimeWindow() bool {
	if o != nil && !isNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given TimeWindow and assigns it to the TimeWindow field.
func (o *ExNotification) SetTimeWindow(v TimeWindow) {
	o.TimeWindow = &v
}

// GetCandPolicies returns the CandPolicies field value if set, zero value otherwise.
func (o *ExNotification) GetCandPolicies() []TransferPolicy {
	if o == nil || isNil(o.CandPolicies) {
		var ret []TransferPolicy
		return ret
	}
	return o.CandPolicies
}

// GetCandPoliciesOk returns a tuple with the CandPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExNotification) GetCandPoliciesOk() ([]TransferPolicy, bool) {
	if o == nil || isNil(o.CandPolicies) {
		return nil, false
	}
	return o.CandPolicies, true
}

// HasCandPolicies returns a boolean if a field has been set.
func (o *ExNotification) HasCandPolicies() bool {
	if o != nil && !isNil(o.CandPolicies) {
		return true
	}

	return false
}

// SetCandPolicies gets a reference to the given []TransferPolicy and assigns it to the CandPolicies field.
func (o *ExNotification) SetCandPolicies(v []TransferPolicy) {
	o.CandPolicies = v
}

func (o ExNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bdtRefId"] = o.BdtRefId
	if !isNil(o.LocationArea5G) {
		toSerialize["locationArea5G"] = o.LocationArea5G
	}
	if !isNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	if !isNil(o.CandPolicies) {
		toSerialize["candPolicies"] = o.CandPolicies
	}
	return toSerialize, nil
}

type NullableExNotification struct {
	value *ExNotification
	isSet bool
}

func (v NullableExNotification) Get() *ExNotification {
	return v.value
}

func (v *NullableExNotification) Set(val *ExNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableExNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableExNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExNotification(val *ExNotification) *NullableExNotification {
	return &NullableExNotification{value: val, isSet: true}
}

func (v NullableExNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


