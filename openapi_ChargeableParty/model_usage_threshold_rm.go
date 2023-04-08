/*
3gpp-chargeable-party

API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ChargeableParty

import (
	"encoding/json"
)

// checks if the UsageThresholdRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageThresholdRm{}

// UsageThresholdRm Represents the same as the UsageThreshold data type but with the nullable:true property.
type UsageThresholdRm struct {
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	Duration NullableInt32 `json:"duration,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	TotalVolume NullableInt64 `json:"totalVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	DownlinkVolume NullableInt64 `json:"downlinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	UplinkVolume NullableInt64 `json:"uplinkVolume,omitempty"`
}

// NewUsageThresholdRm instantiates a new UsageThresholdRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageThresholdRm() *UsageThresholdRm {
	this := UsageThresholdRm{}
	return &this
}

// NewUsageThresholdRmWithDefaults instantiates a new UsageThresholdRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageThresholdRmWithDefaults() *UsageThresholdRm {
	this := UsageThresholdRm{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageThresholdRm) GetDuration() int32 {
	if o == nil || isNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageThresholdRm) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *UsageThresholdRm) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *UsageThresholdRm) SetDuration(v int32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *UsageThresholdRm) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *UsageThresholdRm) UnsetDuration() {
	o.Duration.Unset()
}

// GetTotalVolume returns the TotalVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageThresholdRm) GetTotalVolume() int64 {
	if o == nil || isNil(o.TotalVolume.Get()) {
		var ret int64
		return ret
	}
	return *o.TotalVolume.Get()
}

// GetTotalVolumeOk returns a tuple with the TotalVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageThresholdRm) GetTotalVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalVolume.Get(), o.TotalVolume.IsSet()
}

// HasTotalVolume returns a boolean if a field has been set.
func (o *UsageThresholdRm) HasTotalVolume() bool {
	if o != nil && o.TotalVolume.IsSet() {
		return true
	}

	return false
}

// SetTotalVolume gets a reference to the given NullableInt64 and assigns it to the TotalVolume field.
func (o *UsageThresholdRm) SetTotalVolume(v int64) {
	o.TotalVolume.Set(&v)
}
// SetTotalVolumeNil sets the value for TotalVolume to be an explicit nil
func (o *UsageThresholdRm) SetTotalVolumeNil() {
	o.TotalVolume.Set(nil)
}

// UnsetTotalVolume ensures that no value is present for TotalVolume, not even an explicit nil
func (o *UsageThresholdRm) UnsetTotalVolume() {
	o.TotalVolume.Unset()
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageThresholdRm) GetDownlinkVolume() int64 {
	if o == nil || isNil(o.DownlinkVolume.Get()) {
		var ret int64
		return ret
	}
	return *o.DownlinkVolume.Get()
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageThresholdRm) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownlinkVolume.Get(), o.DownlinkVolume.IsSet()
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *UsageThresholdRm) HasDownlinkVolume() bool {
	if o != nil && o.DownlinkVolume.IsSet() {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given NullableInt64 and assigns it to the DownlinkVolume field.
func (o *UsageThresholdRm) SetDownlinkVolume(v int64) {
	o.DownlinkVolume.Set(&v)
}
// SetDownlinkVolumeNil sets the value for DownlinkVolume to be an explicit nil
func (o *UsageThresholdRm) SetDownlinkVolumeNil() {
	o.DownlinkVolume.Set(nil)
}

// UnsetDownlinkVolume ensures that no value is present for DownlinkVolume, not even an explicit nil
func (o *UsageThresholdRm) UnsetDownlinkVolume() {
	o.DownlinkVolume.Unset()
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageThresholdRm) GetUplinkVolume() int64 {
	if o == nil || isNil(o.UplinkVolume.Get()) {
		var ret int64
		return ret
	}
	return *o.UplinkVolume.Get()
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageThresholdRm) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UplinkVolume.Get(), o.UplinkVolume.IsSet()
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *UsageThresholdRm) HasUplinkVolume() bool {
	if o != nil && o.UplinkVolume.IsSet() {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given NullableInt64 and assigns it to the UplinkVolume field.
func (o *UsageThresholdRm) SetUplinkVolume(v int64) {
	o.UplinkVolume.Set(&v)
}
// SetUplinkVolumeNil sets the value for UplinkVolume to be an explicit nil
func (o *UsageThresholdRm) SetUplinkVolumeNil() {
	o.UplinkVolume.Set(nil)
}

// UnsetUplinkVolume ensures that no value is present for UplinkVolume, not even an explicit nil
func (o *UsageThresholdRm) UnsetUplinkVolume() {
	o.UplinkVolume.Unset()
}

func (o UsageThresholdRm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageThresholdRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.TotalVolume.IsSet() {
		toSerialize["totalVolume"] = o.TotalVolume.Get()
	}
	if o.DownlinkVolume.IsSet() {
		toSerialize["downlinkVolume"] = o.DownlinkVolume.Get()
	}
	if o.UplinkVolume.IsSet() {
		toSerialize["uplinkVolume"] = o.UplinkVolume.Get()
	}
	return toSerialize, nil
}

type NullableUsageThresholdRm struct {
	value *UsageThresholdRm
	isSet bool
}

func (v NullableUsageThresholdRm) Get() *UsageThresholdRm {
	return v.value
}

func (v *NullableUsageThresholdRm) Set(val *UsageThresholdRm) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageThresholdRm) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageThresholdRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageThresholdRm(val *UsageThresholdRm) *NullableUsageThresholdRm {
	return &NullableUsageThresholdRm{value: val, isSet: true}
}

func (v NullableUsageThresholdRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageThresholdRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


