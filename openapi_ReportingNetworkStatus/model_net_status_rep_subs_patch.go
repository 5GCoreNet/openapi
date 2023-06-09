/*
3gpp-network-status-reporting

API for reporting network status.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ReportingNetworkStatus

import (
	"encoding/json"
	"time"
)

// checks if the NetStatusRepSubsPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetStatusRepSubsPatch{}

// NetStatusRepSubsPatch Represents the parameters to request the modification of network status reporting subscription.
type NetStatusRepSubsPatch struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string       `json:"notificationDestination,omitempty"`
	LocationArea            *LocationArea `json:"locationArea,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI with \"nullable=true\" property.
	TimeDuration    NullableTime     `json:"timeDuration,omitempty"`
	ThresholdValues []int32          `json:"thresholdValues,omitempty"`
	ThresholdTypes  []CongestionType `json:"thresholdTypes,omitempty"`
}

// NewNetStatusRepSubsPatch instantiates a new NetStatusRepSubsPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetStatusRepSubsPatch() *NetStatusRepSubsPatch {
	this := NetStatusRepSubsPatch{}
	return &this
}

// NewNetStatusRepSubsPatchWithDefaults instantiates a new NetStatusRepSubsPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetStatusRepSubsPatchWithDefaults() *NetStatusRepSubsPatch {
	this := NetStatusRepSubsPatch{}
	return &this
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *NetStatusRepSubsPatch) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetStatusRepSubsPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *NetStatusRepSubsPatch) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *NetStatusRepSubsPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetLocationArea returns the LocationArea field value if set, zero value otherwise.
func (o *NetStatusRepSubsPatch) GetLocationArea() LocationArea {
	if o == nil || IsNil(o.LocationArea) {
		var ret LocationArea
		return ret
	}
	return *o.LocationArea
}

// GetLocationAreaOk returns a tuple with the LocationArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetStatusRepSubsPatch) GetLocationAreaOk() (*LocationArea, bool) {
	if o == nil || IsNil(o.LocationArea) {
		return nil, false
	}
	return o.LocationArea, true
}

// HasLocationArea returns a boolean if a field has been set.
func (o *NetStatusRepSubsPatch) HasLocationArea() bool {
	if o != nil && !IsNil(o.LocationArea) {
		return true
	}

	return false
}

// SetLocationArea gets a reference to the given LocationArea and assigns it to the LocationArea field.
func (o *NetStatusRepSubsPatch) SetLocationArea(v LocationArea) {
	o.LocationArea = &v
}

// GetTimeDuration returns the TimeDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetStatusRepSubsPatch) GetTimeDuration() time.Time {
	if o == nil || IsNil(o.TimeDuration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TimeDuration.Get()
}

// GetTimeDurationOk returns a tuple with the TimeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetStatusRepSubsPatch) GetTimeDurationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeDuration.Get(), o.TimeDuration.IsSet()
}

// HasTimeDuration returns a boolean if a field has been set.
func (o *NetStatusRepSubsPatch) HasTimeDuration() bool {
	if o != nil && o.TimeDuration.IsSet() {
		return true
	}

	return false
}

// SetTimeDuration gets a reference to the given NullableTime and assigns it to the TimeDuration field.
func (o *NetStatusRepSubsPatch) SetTimeDuration(v time.Time) {
	o.TimeDuration.Set(&v)
}

// SetTimeDurationNil sets the value for TimeDuration to be an explicit nil
func (o *NetStatusRepSubsPatch) SetTimeDurationNil() {
	o.TimeDuration.Set(nil)
}

// UnsetTimeDuration ensures that no value is present for TimeDuration, not even an explicit nil
func (o *NetStatusRepSubsPatch) UnsetTimeDuration() {
	o.TimeDuration.Unset()
}

// GetThresholdValues returns the ThresholdValues field value if set, zero value otherwise.
func (o *NetStatusRepSubsPatch) GetThresholdValues() []int32 {
	if o == nil || IsNil(o.ThresholdValues) {
		var ret []int32
		return ret
	}
	return o.ThresholdValues
}

// GetThresholdValuesOk returns a tuple with the ThresholdValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetStatusRepSubsPatch) GetThresholdValuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ThresholdValues) {
		return nil, false
	}
	return o.ThresholdValues, true
}

// HasThresholdValues returns a boolean if a field has been set.
func (o *NetStatusRepSubsPatch) HasThresholdValues() bool {
	if o != nil && !IsNil(o.ThresholdValues) {
		return true
	}

	return false
}

// SetThresholdValues gets a reference to the given []int32 and assigns it to the ThresholdValues field.
func (o *NetStatusRepSubsPatch) SetThresholdValues(v []int32) {
	o.ThresholdValues = v
}

// GetThresholdTypes returns the ThresholdTypes field value if set, zero value otherwise.
func (o *NetStatusRepSubsPatch) GetThresholdTypes() []CongestionType {
	if o == nil || IsNil(o.ThresholdTypes) {
		var ret []CongestionType
		return ret
	}
	return o.ThresholdTypes
}

// GetThresholdTypesOk returns a tuple with the ThresholdTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetStatusRepSubsPatch) GetThresholdTypesOk() ([]CongestionType, bool) {
	if o == nil || IsNil(o.ThresholdTypes) {
		return nil, false
	}
	return o.ThresholdTypes, true
}

// HasThresholdTypes returns a boolean if a field has been set.
func (o *NetStatusRepSubsPatch) HasThresholdTypes() bool {
	if o != nil && !IsNil(o.ThresholdTypes) {
		return true
	}

	return false
}

// SetThresholdTypes gets a reference to the given []CongestionType and assigns it to the ThresholdTypes field.
func (o *NetStatusRepSubsPatch) SetThresholdTypes(v []CongestionType) {
	o.ThresholdTypes = v
}

func (o NetStatusRepSubsPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetStatusRepSubsPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !IsNil(o.LocationArea) {
		toSerialize["locationArea"] = o.LocationArea
	}
	if o.TimeDuration.IsSet() {
		toSerialize["timeDuration"] = o.TimeDuration.Get()
	}
	if !IsNil(o.ThresholdValues) {
		toSerialize["thresholdValues"] = o.ThresholdValues
	}
	if !IsNil(o.ThresholdTypes) {
		toSerialize["thresholdTypes"] = o.ThresholdTypes
	}
	return toSerialize, nil
}

type NullableNetStatusRepSubsPatch struct {
	value *NetStatusRepSubsPatch
	isSet bool
}

func (v NullableNetStatusRepSubsPatch) Get() *NetStatusRepSubsPatch {
	return v.value
}

func (v *NullableNetStatusRepSubsPatch) Set(val *NetStatusRepSubsPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableNetStatusRepSubsPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableNetStatusRepSubsPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetStatusRepSubsPatch(val *NetStatusRepSubsPatch) *NullableNetStatusRepSubsPatch {
	return &NullableNetStatusRepSubsPatch{value: val, isSet: true}
}

func (v NullableNetStatusRepSubsPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetStatusRepSubsPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
