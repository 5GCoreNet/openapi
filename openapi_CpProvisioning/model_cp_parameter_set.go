/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CpProvisioning

import (
	"encoding/json"
	"time"
)

// checks if the CpParameterSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpParameterSet{}

// CpParameterSet Represents an offered communication pattern parameter set.
type CpParameterSet struct {
	// SCS/AS-chosen correlator provided by the SCS/AS in the request to create a resource fo CP parameter set(s).
	SetId string `json:"setId"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ValidityTime                   *time.Time              `json:"validityTime,omitempty"`
	PeriodicCommunicationIndicator *CommunicationIndicator `json:"periodicCommunicationIndicator,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	CommunicationDurationTime *int32 `json:"communicationDurationTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	PeriodicTime               *int32                      `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType *ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`
	StationaryIndication       *StationaryIndication       `json:"stationaryIndication,omitempty"`
	BatteryInds                []BatteryIndication         `json:"batteryInds,omitempty"`
	TrafficProfile             *TrafficProfile             `json:"trafficProfile,omitempty"`
	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts []UmtLocationArea5G `json:"expectedUmts,omitempty"`
	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays shall be indicated with the next higher numbers. 7 shall indicate Sunday.
	ExpectedUmtDays *int32 `json:"expectedUmtDays,omitempty"`
}

// NewCpParameterSet instantiates a new CpParameterSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpParameterSet(setId string) *CpParameterSet {
	this := CpParameterSet{}
	this.SetId = setId
	return &this
}

// NewCpParameterSetWithDefaults instantiates a new CpParameterSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpParameterSetWithDefaults() *CpParameterSet {
	this := CpParameterSet{}
	return &this
}

// GetSetId returns the SetId field value
func (o *CpParameterSet) GetSetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetId
}

// GetSetIdOk returns a tuple with the SetId field value
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetSetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetId, true
}

// SetSetId sets field value
func (o *CpParameterSet) SetSetId(v string) {
	o.SetId = v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CpParameterSet) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CpParameterSet) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *CpParameterSet) SetSelf(v string) {
	o.Self = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *CpParameterSet) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *CpParameterSet) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *CpParameterSet) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetPeriodicCommunicationIndicator returns the PeriodicCommunicationIndicator field value if set, zero value otherwise.
func (o *CpParameterSet) GetPeriodicCommunicationIndicator() CommunicationIndicator {
	if o == nil || IsNil(o.PeriodicCommunicationIndicator) {
		var ret CommunicationIndicator
		return ret
	}
	return *o.PeriodicCommunicationIndicator
}

// GetPeriodicCommunicationIndicatorOk returns a tuple with the PeriodicCommunicationIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetPeriodicCommunicationIndicatorOk() (*CommunicationIndicator, bool) {
	if o == nil || IsNil(o.PeriodicCommunicationIndicator) {
		return nil, false
	}
	return o.PeriodicCommunicationIndicator, true
}

// HasPeriodicCommunicationIndicator returns a boolean if a field has been set.
func (o *CpParameterSet) HasPeriodicCommunicationIndicator() bool {
	if o != nil && !IsNil(o.PeriodicCommunicationIndicator) {
		return true
	}

	return false
}

// SetPeriodicCommunicationIndicator gets a reference to the given CommunicationIndicator and assigns it to the PeriodicCommunicationIndicator field.
func (o *CpParameterSet) SetPeriodicCommunicationIndicator(v CommunicationIndicator) {
	o.PeriodicCommunicationIndicator = &v
}

// GetCommunicationDurationTime returns the CommunicationDurationTime field value if set, zero value otherwise.
func (o *CpParameterSet) GetCommunicationDurationTime() int32 {
	if o == nil || IsNil(o.CommunicationDurationTime) {
		var ret int32
		return ret
	}
	return *o.CommunicationDurationTime
}

// GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetCommunicationDurationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CommunicationDurationTime) {
		return nil, false
	}
	return o.CommunicationDurationTime, true
}

// HasCommunicationDurationTime returns a boolean if a field has been set.
func (o *CpParameterSet) HasCommunicationDurationTime() bool {
	if o != nil && !IsNil(o.CommunicationDurationTime) {
		return true
	}

	return false
}

// SetCommunicationDurationTime gets a reference to the given int32 and assigns it to the CommunicationDurationTime field.
func (o *CpParameterSet) SetCommunicationDurationTime(v int32) {
	o.CommunicationDurationTime = &v
}

// GetPeriodicTime returns the PeriodicTime field value if set, zero value otherwise.
func (o *CpParameterSet) GetPeriodicTime() int32 {
	if o == nil || IsNil(o.PeriodicTime) {
		var ret int32
		return ret
	}
	return *o.PeriodicTime
}

// GetPeriodicTimeOk returns a tuple with the PeriodicTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetPeriodicTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodicTime) {
		return nil, false
	}
	return o.PeriodicTime, true
}

// HasPeriodicTime returns a boolean if a field has been set.
func (o *CpParameterSet) HasPeriodicTime() bool {
	if o != nil && !IsNil(o.PeriodicTime) {
		return true
	}

	return false
}

// SetPeriodicTime gets a reference to the given int32 and assigns it to the PeriodicTime field.
func (o *CpParameterSet) SetPeriodicTime(v int32) {
	o.PeriodicTime = &v
}

// GetScheduledCommunicationTime returns the ScheduledCommunicationTime field value if set, zero value otherwise.
func (o *CpParameterSet) GetScheduledCommunicationTime() ScheduledCommunicationTime {
	if o == nil || IsNil(o.ScheduledCommunicationTime) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.ScheduledCommunicationTime
}

// GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime, bool) {
	if o == nil || IsNil(o.ScheduledCommunicationTime) {
		return nil, false
	}
	return o.ScheduledCommunicationTime, true
}

// HasScheduledCommunicationTime returns a boolean if a field has been set.
func (o *CpParameterSet) HasScheduledCommunicationTime() bool {
	if o != nil && !IsNil(o.ScheduledCommunicationTime) {
		return true
	}

	return false
}

// SetScheduledCommunicationTime gets a reference to the given ScheduledCommunicationTime and assigns it to the ScheduledCommunicationTime field.
func (o *CpParameterSet) SetScheduledCommunicationTime(v ScheduledCommunicationTime) {
	o.ScheduledCommunicationTime = &v
}

// GetScheduledCommunicationType returns the ScheduledCommunicationType field value if set, zero value otherwise.
func (o *CpParameterSet) GetScheduledCommunicationType() ScheduledCommunicationType {
	if o == nil || IsNil(o.ScheduledCommunicationType) {
		var ret ScheduledCommunicationType
		return ret
	}
	return *o.ScheduledCommunicationType
}

// GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool) {
	if o == nil || IsNil(o.ScheduledCommunicationType) {
		return nil, false
	}
	return o.ScheduledCommunicationType, true
}

// HasScheduledCommunicationType returns a boolean if a field has been set.
func (o *CpParameterSet) HasScheduledCommunicationType() bool {
	if o != nil && !IsNil(o.ScheduledCommunicationType) {
		return true
	}

	return false
}

// SetScheduledCommunicationType gets a reference to the given ScheduledCommunicationType and assigns it to the ScheduledCommunicationType field.
func (o *CpParameterSet) SetScheduledCommunicationType(v ScheduledCommunicationType) {
	o.ScheduledCommunicationType = &v
}

// GetStationaryIndication returns the StationaryIndication field value if set, zero value otherwise.
func (o *CpParameterSet) GetStationaryIndication() StationaryIndication {
	if o == nil || IsNil(o.StationaryIndication) {
		var ret StationaryIndication
		return ret
	}
	return *o.StationaryIndication
}

// GetStationaryIndicationOk returns a tuple with the StationaryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetStationaryIndicationOk() (*StationaryIndication, bool) {
	if o == nil || IsNil(o.StationaryIndication) {
		return nil, false
	}
	return o.StationaryIndication, true
}

// HasStationaryIndication returns a boolean if a field has been set.
func (o *CpParameterSet) HasStationaryIndication() bool {
	if o != nil && !IsNil(o.StationaryIndication) {
		return true
	}

	return false
}

// SetStationaryIndication gets a reference to the given StationaryIndication and assigns it to the StationaryIndication field.
func (o *CpParameterSet) SetStationaryIndication(v StationaryIndication) {
	o.StationaryIndication = &v
}

// GetBatteryInds returns the BatteryInds field value if set, zero value otherwise.
func (o *CpParameterSet) GetBatteryInds() []BatteryIndication {
	if o == nil || IsNil(o.BatteryInds) {
		var ret []BatteryIndication
		return ret
	}
	return o.BatteryInds
}

// GetBatteryIndsOk returns a tuple with the BatteryInds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetBatteryIndsOk() ([]BatteryIndication, bool) {
	if o == nil || IsNil(o.BatteryInds) {
		return nil, false
	}
	return o.BatteryInds, true
}

// HasBatteryInds returns a boolean if a field has been set.
func (o *CpParameterSet) HasBatteryInds() bool {
	if o != nil && !IsNil(o.BatteryInds) {
		return true
	}

	return false
}

// SetBatteryInds gets a reference to the given []BatteryIndication and assigns it to the BatteryInds field.
func (o *CpParameterSet) SetBatteryInds(v []BatteryIndication) {
	o.BatteryInds = v
}

// GetTrafficProfile returns the TrafficProfile field value if set, zero value otherwise.
func (o *CpParameterSet) GetTrafficProfile() TrafficProfile {
	if o == nil || IsNil(o.TrafficProfile) {
		var ret TrafficProfile
		return ret
	}
	return *o.TrafficProfile
}

// GetTrafficProfileOk returns a tuple with the TrafficProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetTrafficProfileOk() (*TrafficProfile, bool) {
	if o == nil || IsNil(o.TrafficProfile) {
		return nil, false
	}
	return o.TrafficProfile, true
}

// HasTrafficProfile returns a boolean if a field has been set.
func (o *CpParameterSet) HasTrafficProfile() bool {
	if o != nil && !IsNil(o.TrafficProfile) {
		return true
	}

	return false
}

// SetTrafficProfile gets a reference to the given TrafficProfile and assigns it to the TrafficProfile field.
func (o *CpParameterSet) SetTrafficProfile(v TrafficProfile) {
	o.TrafficProfile = &v
}

// GetExpectedUmts returns the ExpectedUmts field value if set, zero value otherwise.
func (o *CpParameterSet) GetExpectedUmts() []UmtLocationArea5G {
	if o == nil || IsNil(o.ExpectedUmts) {
		var ret []UmtLocationArea5G
		return ret
	}
	return o.ExpectedUmts
}

// GetExpectedUmtsOk returns a tuple with the ExpectedUmts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetExpectedUmtsOk() ([]UmtLocationArea5G, bool) {
	if o == nil || IsNil(o.ExpectedUmts) {
		return nil, false
	}
	return o.ExpectedUmts, true
}

// HasExpectedUmts returns a boolean if a field has been set.
func (o *CpParameterSet) HasExpectedUmts() bool {
	if o != nil && !IsNil(o.ExpectedUmts) {
		return true
	}

	return false
}

// SetExpectedUmts gets a reference to the given []UmtLocationArea5G and assigns it to the ExpectedUmts field.
func (o *CpParameterSet) SetExpectedUmts(v []UmtLocationArea5G) {
	o.ExpectedUmts = v
}

// GetExpectedUmtDays returns the ExpectedUmtDays field value if set, zero value otherwise.
func (o *CpParameterSet) GetExpectedUmtDays() int32 {
	if o == nil || IsNil(o.ExpectedUmtDays) {
		var ret int32
		return ret
	}
	return *o.ExpectedUmtDays
}

// GetExpectedUmtDaysOk returns a tuple with the ExpectedUmtDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpParameterSet) GetExpectedUmtDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpectedUmtDays) {
		return nil, false
	}
	return o.ExpectedUmtDays, true
}

// HasExpectedUmtDays returns a boolean if a field has been set.
func (o *CpParameterSet) HasExpectedUmtDays() bool {
	if o != nil && !IsNil(o.ExpectedUmtDays) {
		return true
	}

	return false
}

// SetExpectedUmtDays gets a reference to the given int32 and assigns it to the ExpectedUmtDays field.
func (o *CpParameterSet) SetExpectedUmtDays(v int32) {
	o.ExpectedUmtDays = &v
}

func (o CpParameterSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpParameterSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["setId"] = o.SetId
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.PeriodicCommunicationIndicator) {
		toSerialize["periodicCommunicationIndicator"] = o.PeriodicCommunicationIndicator
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
	if !IsNil(o.StationaryIndication) {
		toSerialize["stationaryIndication"] = o.StationaryIndication
	}
	if !IsNil(o.BatteryInds) {
		toSerialize["batteryInds"] = o.BatteryInds
	}
	if !IsNil(o.TrafficProfile) {
		toSerialize["trafficProfile"] = o.TrafficProfile
	}
	if !IsNil(o.ExpectedUmts) {
		toSerialize["expectedUmts"] = o.ExpectedUmts
	}
	if !IsNil(o.ExpectedUmtDays) {
		toSerialize["expectedUmtDays"] = o.ExpectedUmtDays
	}
	return toSerialize, nil
}

type NullableCpParameterSet struct {
	value *CpParameterSet
	isSet bool
}

func (v NullableCpParameterSet) Get() *CpParameterSet {
	return v.value
}

func (v *NullableCpParameterSet) Set(val *CpParameterSet) {
	v.value = val
	v.isSet = true
}

func (v NullableCpParameterSet) IsSet() bool {
	return v.isSet
}

func (v *NullableCpParameterSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpParameterSet(val *CpParameterSet) *NullableCpParameterSet {
	return &NullableCpParameterSet{value: val, isSet: true}
}

func (v NullableCpParameterSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpParameterSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
