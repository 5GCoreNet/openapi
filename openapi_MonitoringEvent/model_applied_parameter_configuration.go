/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the AppliedParameterConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppliedParameterConfiguration{}

// AppliedParameterConfiguration Represents the parameter configuration applied in the network.
type AppliedParameterConfiguration struct {
	// Each element uniquely identifies a user.
	ExternalIds []string `json:"externalIds,omitempty"`
	// Each element identifies the MS internal PSTN/ISDN number allocated for a UE.
	Msisdns []string `json:"msisdns,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	MaximumLatency *int32 `json:"maximumLatency,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	MaximumResponseTime *int32 `json:"maximumResponseTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	MaximumDetectionTime *int32 `json:"maximumDetectionTime,omitempty"`
}

// NewAppliedParameterConfiguration instantiates a new AppliedParameterConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppliedParameterConfiguration() *AppliedParameterConfiguration {
	this := AppliedParameterConfiguration{}
	return &this
}

// NewAppliedParameterConfigurationWithDefaults instantiates a new AppliedParameterConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppliedParameterConfigurationWithDefaults() *AppliedParameterConfiguration {
	this := AppliedParameterConfiguration{}
	return &this
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *AppliedParameterConfiguration) GetExternalIds() []string {
	if o == nil || isNil(o.ExternalIds) {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedParameterConfiguration) GetExternalIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *AppliedParameterConfiguration) HasExternalIds() bool {
	if o != nil && !isNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *AppliedParameterConfiguration) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetMsisdns returns the Msisdns field value if set, zero value otherwise.
func (o *AppliedParameterConfiguration) GetMsisdns() []string {
	if o == nil || isNil(o.Msisdns) {
		var ret []string
		return ret
	}
	return o.Msisdns
}

// GetMsisdnsOk returns a tuple with the Msisdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedParameterConfiguration) GetMsisdnsOk() ([]string, bool) {
	if o == nil || isNil(o.Msisdns) {
		return nil, false
	}
	return o.Msisdns, true
}

// HasMsisdns returns a boolean if a field has been set.
func (o *AppliedParameterConfiguration) HasMsisdns() bool {
	if o != nil && !isNil(o.Msisdns) {
		return true
	}

	return false
}

// SetMsisdns gets a reference to the given []string and assigns it to the Msisdns field.
func (o *AppliedParameterConfiguration) SetMsisdns(v []string) {
	o.Msisdns = v
}

// GetMaximumLatency returns the MaximumLatency field value if set, zero value otherwise.
func (o *AppliedParameterConfiguration) GetMaximumLatency() int32 {
	if o == nil || isNil(o.MaximumLatency) {
		var ret int32
		return ret
	}
	return *o.MaximumLatency
}

// GetMaximumLatencyOk returns a tuple with the MaximumLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedParameterConfiguration) GetMaximumLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumLatency) {
		return nil, false
	}
	return o.MaximumLatency, true
}

// HasMaximumLatency returns a boolean if a field has been set.
func (o *AppliedParameterConfiguration) HasMaximumLatency() bool {
	if o != nil && !isNil(o.MaximumLatency) {
		return true
	}

	return false
}

// SetMaximumLatency gets a reference to the given int32 and assigns it to the MaximumLatency field.
func (o *AppliedParameterConfiguration) SetMaximumLatency(v int32) {
	o.MaximumLatency = &v
}

// GetMaximumResponseTime returns the MaximumResponseTime field value if set, zero value otherwise.
func (o *AppliedParameterConfiguration) GetMaximumResponseTime() int32 {
	if o == nil || isNil(o.MaximumResponseTime) {
		var ret int32
		return ret
	}
	return *o.MaximumResponseTime
}

// GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedParameterConfiguration) GetMaximumResponseTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumResponseTime) {
		return nil, false
	}
	return o.MaximumResponseTime, true
}

// HasMaximumResponseTime returns a boolean if a field has been set.
func (o *AppliedParameterConfiguration) HasMaximumResponseTime() bool {
	if o != nil && !isNil(o.MaximumResponseTime) {
		return true
	}

	return false
}

// SetMaximumResponseTime gets a reference to the given int32 and assigns it to the MaximumResponseTime field.
func (o *AppliedParameterConfiguration) SetMaximumResponseTime(v int32) {
	o.MaximumResponseTime = &v
}

// GetMaximumDetectionTime returns the MaximumDetectionTime field value if set, zero value otherwise.
func (o *AppliedParameterConfiguration) GetMaximumDetectionTime() int32 {
	if o == nil || isNil(o.MaximumDetectionTime) {
		var ret int32
		return ret
	}
	return *o.MaximumDetectionTime
}

// GetMaximumDetectionTimeOk returns a tuple with the MaximumDetectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppliedParameterConfiguration) GetMaximumDetectionTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumDetectionTime) {
		return nil, false
	}
	return o.MaximumDetectionTime, true
}

// HasMaximumDetectionTime returns a boolean if a field has been set.
func (o *AppliedParameterConfiguration) HasMaximumDetectionTime() bool {
	if o != nil && !isNil(o.MaximumDetectionTime) {
		return true
	}

	return false
}

// SetMaximumDetectionTime gets a reference to the given int32 and assigns it to the MaximumDetectionTime field.
func (o *AppliedParameterConfiguration) SetMaximumDetectionTime(v int32) {
	o.MaximumDetectionTime = &v
}

func (o AppliedParameterConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppliedParameterConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !isNil(o.Msisdns) {
		toSerialize["msisdns"] = o.Msisdns
	}
	if !isNil(o.MaximumLatency) {
		toSerialize["maximumLatency"] = o.MaximumLatency
	}
	if !isNil(o.MaximumResponseTime) {
		toSerialize["maximumResponseTime"] = o.MaximumResponseTime
	}
	if !isNil(o.MaximumDetectionTime) {
		toSerialize["maximumDetectionTime"] = o.MaximumDetectionTime
	}
	return toSerialize, nil
}

type NullableAppliedParameterConfiguration struct {
	value *AppliedParameterConfiguration
	isSet bool
}

func (v NullableAppliedParameterConfiguration) Get() *AppliedParameterConfiguration {
	return v.value
}

func (v *NullableAppliedParameterConfiguration) Set(val *AppliedParameterConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedParameterConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedParameterConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedParameterConfiguration(val *AppliedParameterConfiguration) *NullableAppliedParameterConfiguration {
	return &NullableAppliedParameterConfiguration{value: val, isSet: true}
}

func (v NullableAppliedParameterConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedParameterConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


