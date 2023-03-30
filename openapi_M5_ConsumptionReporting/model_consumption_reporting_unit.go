/*
M5_ConsumptionReporting

5GMS AF M5 Consumption Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ConsumptionReporting

import (
	"encoding/json"
	"time"
)

// checks if the ConsumptionReportingUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumptionReportingUnit{}

// ConsumptionReportingUnit A Consumption Reporting Unit.
type ConsumptionReportingUnit struct {
	MediaConsumed string `json:"mediaConsumed"`
	MediaEndpointAddress *EndpointAddress `json:"mediaEndpointAddress,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`
	// indicating a time in seconds.
	Duration int32 `json:"duration"`
	Locations []TypedLocation `json:"locations,omitempty"`
}

// NewConsumptionReportingUnit instantiates a new ConsumptionReportingUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumptionReportingUnit(mediaConsumed string, startTime time.Time, duration int32) *ConsumptionReportingUnit {
	this := ConsumptionReportingUnit{}
	this.MediaConsumed = mediaConsumed
	this.StartTime = startTime
	this.Duration = duration
	return &this
}

// NewConsumptionReportingUnitWithDefaults instantiates a new ConsumptionReportingUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumptionReportingUnitWithDefaults() *ConsumptionReportingUnit {
	this := ConsumptionReportingUnit{}
	return &this
}

// GetMediaConsumed returns the MediaConsumed field value
func (o *ConsumptionReportingUnit) GetMediaConsumed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaConsumed
}

// GetMediaConsumedOk returns a tuple with the MediaConsumed field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReportingUnit) GetMediaConsumedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaConsumed, true
}

// SetMediaConsumed sets field value
func (o *ConsumptionReportingUnit) SetMediaConsumed(v string) {
	o.MediaConsumed = v
}

// GetMediaEndpointAddress returns the MediaEndpointAddress field value if set, zero value otherwise.
func (o *ConsumptionReportingUnit) GetMediaEndpointAddress() EndpointAddress {
	if o == nil || IsNil(o.MediaEndpointAddress) {
		var ret EndpointAddress
		return ret
	}
	return *o.MediaEndpointAddress
}

// GetMediaEndpointAddressOk returns a tuple with the MediaEndpointAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionReportingUnit) GetMediaEndpointAddressOk() (*EndpointAddress, bool) {
	if o == nil || IsNil(o.MediaEndpointAddress) {
		return nil, false
	}
	return o.MediaEndpointAddress, true
}

// HasMediaEndpointAddress returns a boolean if a field has been set.
func (o *ConsumptionReportingUnit) HasMediaEndpointAddress() bool {
	if o != nil && !IsNil(o.MediaEndpointAddress) {
		return true
	}

	return false
}

// SetMediaEndpointAddress gets a reference to the given EndpointAddress and assigns it to the MediaEndpointAddress field.
func (o *ConsumptionReportingUnit) SetMediaEndpointAddress(v EndpointAddress) {
	o.MediaEndpointAddress = &v
}

// GetStartTime returns the StartTime field value
func (o *ConsumptionReportingUnit) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReportingUnit) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ConsumptionReportingUnit) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetDuration returns the Duration field value
func (o *ConsumptionReportingUnit) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ConsumptionReportingUnit) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ConsumptionReportingUnit) SetDuration(v int32) {
	o.Duration = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ConsumptionReportingUnit) GetLocations() []TypedLocation {
	if o == nil || IsNil(o.Locations) {
		var ret []TypedLocation
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumptionReportingUnit) GetLocationsOk() ([]TypedLocation, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ConsumptionReportingUnit) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []TypedLocation and assigns it to the Locations field.
func (o *ConsumptionReportingUnit) SetLocations(v []TypedLocation) {
	o.Locations = v
}

func (o ConsumptionReportingUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumptionReportingUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mediaConsumed"] = o.MediaConsumed
	if !IsNil(o.MediaEndpointAddress) {
		toSerialize["mediaEndpointAddress"] = o.MediaEndpointAddress
	}
	toSerialize["startTime"] = o.StartTime
	toSerialize["duration"] = o.Duration
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	return toSerialize, nil
}

type NullableConsumptionReportingUnit struct {
	value *ConsumptionReportingUnit
	isSet bool
}

func (v NullableConsumptionReportingUnit) Get() *ConsumptionReportingUnit {
	return v.value
}

func (v *NullableConsumptionReportingUnit) Set(val *ConsumptionReportingUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumptionReportingUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumptionReportingUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumptionReportingUnit(val *ConsumptionReportingUnit) *NullableConsumptionReportingUnit {
	return &NullableConsumptionReportingUnit{value: val, isSet: true}
}

func (v NullableConsumptionReportingUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumptionReportingUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


