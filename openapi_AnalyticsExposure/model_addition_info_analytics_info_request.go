/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the AdditionInfoAnalyticsInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionInfoAnalyticsInfoRequest{}

// AdditionInfoAnalyticsInfoRequest Indicates additional information why the analytics request is rejected.
type AdditionInfoAnalyticsInfoRequest struct {
	// indicating a time in seconds.
	RvWaitTime *int32 `json:"rvWaitTime,omitempty"`
}

// NewAdditionInfoAnalyticsInfoRequest instantiates a new AdditionInfoAnalyticsInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionInfoAnalyticsInfoRequest() *AdditionInfoAnalyticsInfoRequest {
	this := AdditionInfoAnalyticsInfoRequest{}
	return &this
}

// NewAdditionInfoAnalyticsInfoRequestWithDefaults instantiates a new AdditionInfoAnalyticsInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionInfoAnalyticsInfoRequestWithDefaults() *AdditionInfoAnalyticsInfoRequest {
	this := AdditionInfoAnalyticsInfoRequest{}
	return &this
}

// GetRvWaitTime returns the RvWaitTime field value if set, zero value otherwise.
func (o *AdditionInfoAnalyticsInfoRequest) GetRvWaitTime() int32 {
	if o == nil || IsNil(o.RvWaitTime) {
		var ret int32
		return ret
	}
	return *o.RvWaitTime
}

// GetRvWaitTimeOk returns a tuple with the RvWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionInfoAnalyticsInfoRequest) GetRvWaitTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.RvWaitTime) {
		return nil, false
	}
	return o.RvWaitTime, true
}

// HasRvWaitTime returns a boolean if a field has been set.
func (o *AdditionInfoAnalyticsInfoRequest) HasRvWaitTime() bool {
	if o != nil && !IsNil(o.RvWaitTime) {
		return true
	}

	return false
}

// SetRvWaitTime gets a reference to the given int32 and assigns it to the RvWaitTime field.
func (o *AdditionInfoAnalyticsInfoRequest) SetRvWaitTime(v int32) {
	o.RvWaitTime = &v
}

func (o AdditionInfoAnalyticsInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionInfoAnalyticsInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RvWaitTime) {
		toSerialize["rvWaitTime"] = o.RvWaitTime
	}
	return toSerialize, nil
}

type NullableAdditionInfoAnalyticsInfoRequest struct {
	value *AdditionInfoAnalyticsInfoRequest
	isSet bool
}

func (v NullableAdditionInfoAnalyticsInfoRequest) Get() *AdditionInfoAnalyticsInfoRequest {
	return v.value
}

func (v *NullableAdditionInfoAnalyticsInfoRequest) Set(val *AdditionInfoAnalyticsInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionInfoAnalyticsInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionInfoAnalyticsInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionInfoAnalyticsInfoRequest(val *AdditionInfoAnalyticsInfoRequest) *NullableAdditionInfoAnalyticsInfoRequest {
	return &NullableAdditionInfoAnalyticsInfoRequest{value: val, isSet: true}
}

func (v NullableAdditionInfoAnalyticsInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionInfoAnalyticsInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
