/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"time"
)

// checks if the ExpectedUeBehaviourData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpectedUeBehaviourData{}

// ExpectedUeBehaviourData struct for ExpectedUeBehaviourData
type ExpectedUeBehaviourData struct {
	StationaryIndication *StationaryIndication `json:"stationaryIndication,omitempty"`
	// indicating a time in seconds.
	CommunicationDurationTime *int32 `json:"communicationDurationTime,omitempty"`
	// indicating a time in seconds.
	PeriodicTime *int32 `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime1 `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType *ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`
	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts []LocationArea `json:"expectedUmts,omitempty"`
	TrafficProfile *TrafficProfile `json:"trafficProfile,omitempty"`
	BatteryIndication *BatteryIndication `json:"batteryIndication,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
}

// NewExpectedUeBehaviourData instantiates a new ExpectedUeBehaviourData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpectedUeBehaviourData() *ExpectedUeBehaviourData {
	this := ExpectedUeBehaviourData{}
	return &this
}

// NewExpectedUeBehaviourDataWithDefaults instantiates a new ExpectedUeBehaviourData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpectedUeBehaviourDataWithDefaults() *ExpectedUeBehaviourData {
	this := ExpectedUeBehaviourData{}
	return &this
}

// GetStationaryIndication returns the StationaryIndication field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetStationaryIndication() StationaryIndication {
	if o == nil || IsNil(o.StationaryIndication) {
		var ret StationaryIndication
		return ret
	}
	return *o.StationaryIndication
}

// GetStationaryIndicationOk returns a tuple with the StationaryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetStationaryIndicationOk() (*StationaryIndication, bool) {
	if o == nil || IsNil(o.StationaryIndication) {
		return nil, false
	}
	return o.StationaryIndication, true
}

// HasStationaryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasStationaryIndication() bool {
	if o != nil && !IsNil(o.StationaryIndication) {
		return true
	}

	return false
}

// SetStationaryIndication gets a reference to the given StationaryIndication and assigns it to the StationaryIndication field.
func (o *ExpectedUeBehaviourData) SetStationaryIndication(v StationaryIndication) {
	o.StationaryIndication = &v
}

// GetCommunicationDurationTime returns the CommunicationDurationTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetCommunicationDurationTime() int32 {
	if o == nil || IsNil(o.CommunicationDurationTime) {
		var ret int32
		return ret
	}
	return *o.CommunicationDurationTime
}

// GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetCommunicationDurationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CommunicationDurationTime) {
		return nil, false
	}
	return o.CommunicationDurationTime, true
}

// HasCommunicationDurationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasCommunicationDurationTime() bool {
	if o != nil && !IsNil(o.CommunicationDurationTime) {
		return true
	}

	return false
}

// SetCommunicationDurationTime gets a reference to the given int32 and assigns it to the CommunicationDurationTime field.
func (o *ExpectedUeBehaviourData) SetCommunicationDurationTime(v int32) {
	o.CommunicationDurationTime = &v
}

// GetPeriodicTime returns the PeriodicTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetPeriodicTime() int32 {
	if o == nil || IsNil(o.PeriodicTime) {
		var ret int32
		return ret
	}
	return *o.PeriodicTime
}

// GetPeriodicTimeOk returns a tuple with the PeriodicTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetPeriodicTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodicTime) {
		return nil, false
	}
	return o.PeriodicTime, true
}

// HasPeriodicTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasPeriodicTime() bool {
	if o != nil && !IsNil(o.PeriodicTime) {
		return true
	}

	return false
}

// SetPeriodicTime gets a reference to the given int32 and assigns it to the PeriodicTime field.
func (o *ExpectedUeBehaviourData) SetPeriodicTime(v int32) {
	o.PeriodicTime = &v
}

// GetScheduledCommunicationTime returns the ScheduledCommunicationTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTime() ScheduledCommunicationTime1 {
	if o == nil || IsNil(o.ScheduledCommunicationTime) {
		var ret ScheduledCommunicationTime1
		return ret
	}
	return *o.ScheduledCommunicationTime
}

// GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime1, bool) {
	if o == nil || IsNil(o.ScheduledCommunicationTime) {
		return nil, false
	}
	return o.ScheduledCommunicationTime, true
}

// HasScheduledCommunicationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasScheduledCommunicationTime() bool {
	if o != nil && !IsNil(o.ScheduledCommunicationTime) {
		return true
	}

	return false
}

// SetScheduledCommunicationTime gets a reference to the given ScheduledCommunicationTime1 and assigns it to the ScheduledCommunicationTime field.
func (o *ExpectedUeBehaviourData) SetScheduledCommunicationTime(v ScheduledCommunicationTime1) {
	o.ScheduledCommunicationTime = &v
}

// GetScheduledCommunicationType returns the ScheduledCommunicationType field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetScheduledCommunicationType() ScheduledCommunicationType {
	if o == nil || IsNil(o.ScheduledCommunicationType) {
		var ret ScheduledCommunicationType
		return ret
	}
	return *o.ScheduledCommunicationType
}

// GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool) {
	if o == nil || IsNil(o.ScheduledCommunicationType) {
		return nil, false
	}
	return o.ScheduledCommunicationType, true
}

// HasScheduledCommunicationType returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasScheduledCommunicationType() bool {
	if o != nil && !IsNil(o.ScheduledCommunicationType) {
		return true
	}

	return false
}

// SetScheduledCommunicationType gets a reference to the given ScheduledCommunicationType and assigns it to the ScheduledCommunicationType field.
func (o *ExpectedUeBehaviourData) SetScheduledCommunicationType(v ScheduledCommunicationType) {
	o.ScheduledCommunicationType = &v
}

// GetExpectedUmts returns the ExpectedUmts field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetExpectedUmts() []LocationArea {
	if o == nil || IsNil(o.ExpectedUmts) {
		var ret []LocationArea
		return ret
	}
	return o.ExpectedUmts
}

// GetExpectedUmtsOk returns a tuple with the ExpectedUmts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetExpectedUmtsOk() ([]LocationArea, bool) {
	if o == nil || IsNil(o.ExpectedUmts) {
		return nil, false
	}
	return o.ExpectedUmts, true
}

// HasExpectedUmts returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasExpectedUmts() bool {
	if o != nil && !IsNil(o.ExpectedUmts) {
		return true
	}

	return false
}

// SetExpectedUmts gets a reference to the given []LocationArea and assigns it to the ExpectedUmts field.
func (o *ExpectedUeBehaviourData) SetExpectedUmts(v []LocationArea) {
	o.ExpectedUmts = v
}

// GetTrafficProfile returns the TrafficProfile field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetTrafficProfile() TrafficProfile {
	if o == nil || IsNil(o.TrafficProfile) {
		var ret TrafficProfile
		return ret
	}
	return *o.TrafficProfile
}

// GetTrafficProfileOk returns a tuple with the TrafficProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetTrafficProfileOk() (*TrafficProfile, bool) {
	if o == nil || IsNil(o.TrafficProfile) {
		return nil, false
	}
	return o.TrafficProfile, true
}

// HasTrafficProfile returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasTrafficProfile() bool {
	if o != nil && !IsNil(o.TrafficProfile) {
		return true
	}

	return false
}

// SetTrafficProfile gets a reference to the given TrafficProfile and assigns it to the TrafficProfile field.
func (o *ExpectedUeBehaviourData) SetTrafficProfile(v TrafficProfile) {
	o.TrafficProfile = &v
}

// GetBatteryIndication returns the BatteryIndication field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetBatteryIndication() BatteryIndication {
	if o == nil || IsNil(o.BatteryIndication) {
		var ret BatteryIndication
		return ret
	}
	return *o.BatteryIndication
}

// GetBatteryIndicationOk returns a tuple with the BatteryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetBatteryIndicationOk() (*BatteryIndication, bool) {
	if o == nil || IsNil(o.BatteryIndication) {
		return nil, false
	}
	return o.BatteryIndication, true
}

// HasBatteryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasBatteryIndication() bool {
	if o != nil && !IsNil(o.BatteryIndication) {
		return true
	}

	return false
}

// SetBatteryIndication gets a reference to the given BatteryIndication and assigns it to the BatteryIndication field.
func (o *ExpectedUeBehaviourData) SetBatteryIndication(v BatteryIndication) {
	o.BatteryIndication = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviourData) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviourData) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviourData) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *ExpectedUeBehaviourData) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

func (o ExpectedUeBehaviourData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpectedUeBehaviourData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StationaryIndication) {
		toSerialize["stationaryIndication"] = o.StationaryIndication
	}
	if !IsNil(o.CommunicationDurationTime) {
		toSerialize["communicationDurationTime"] = o.CommunicationDurationTime
	}
	if !IsNil(o.PeriodicTime) {
		toSerialize["periodicTime"] = o.PeriodicTime
	}
	if !IsNil(o.ScheduledCommunicationTime) {
		toSerialize["scheduledCommunicationTime"] = o.ScheduledCommunicationTime
	}
	if !IsNil(o.ScheduledCommunicationType) {
		toSerialize["scheduledCommunicationType"] = o.ScheduledCommunicationType
	}
	if !IsNil(o.ExpectedUmts) {
		toSerialize["expectedUmts"] = o.ExpectedUmts
	}
	if !IsNil(o.TrafficProfile) {
		toSerialize["trafficProfile"] = o.TrafficProfile
	}
	if !IsNil(o.BatteryIndication) {
		toSerialize["batteryIndication"] = o.BatteryIndication
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	return toSerialize, nil
}

type NullableExpectedUeBehaviourData struct {
	value *ExpectedUeBehaviourData
	isSet bool
}

func (v NullableExpectedUeBehaviourData) Get() *ExpectedUeBehaviourData {
	return v.value
}

func (v *NullableExpectedUeBehaviourData) Set(val *ExpectedUeBehaviourData) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectedUeBehaviourData) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectedUeBehaviourData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectedUeBehaviourData(val *ExpectedUeBehaviourData) *NullableExpectedUeBehaviourData {
	return &NullableExpectedUeBehaviourData{value: val, isSet: true}
}

func (v NullableExpectedUeBehaviourData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectedUeBehaviourData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


