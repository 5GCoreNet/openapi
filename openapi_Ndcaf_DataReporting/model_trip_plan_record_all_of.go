/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
	"time"
)

// checks if the TripPlanRecordAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TripPlanRecordAllOf{}

// TripPlanRecordAllOf struct for TripPlanRecordAllOf
type TripPlanRecordAllOf struct {
	StartingPoint LocationData   `json:"startingPoint"`
	Waypoints     []LocationData `json:"waypoints,omitempty"`
	Destination   LocationData   `json:"destination"`
	// Indicates value of horizontal speed.
	EstimatedAverageSpeed *float32 `json:"estimatedAverageSpeed,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EstimatedArrivalTime *time.Time `json:"estimatedArrivalTime,omitempty"`
}

// NewTripPlanRecordAllOf instantiates a new TripPlanRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTripPlanRecordAllOf(startingPoint LocationData, destination LocationData) *TripPlanRecordAllOf {
	this := TripPlanRecordAllOf{}
	this.StartingPoint = startingPoint
	this.Destination = destination
	return &this
}

// NewTripPlanRecordAllOfWithDefaults instantiates a new TripPlanRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTripPlanRecordAllOfWithDefaults() *TripPlanRecordAllOf {
	this := TripPlanRecordAllOf{}
	return &this
}

// GetStartingPoint returns the StartingPoint field value
func (o *TripPlanRecordAllOf) GetStartingPoint() LocationData {
	if o == nil {
		var ret LocationData
		return ret
	}

	return o.StartingPoint
}

// GetStartingPointOk returns a tuple with the StartingPoint field value
// and a boolean to check if the value has been set.
func (o *TripPlanRecordAllOf) GetStartingPointOk() (*LocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingPoint, true
}

// SetStartingPoint sets field value
func (o *TripPlanRecordAllOf) SetStartingPoint(v LocationData) {
	o.StartingPoint = v
}

// GetWaypoints returns the Waypoints field value if set, zero value otherwise.
func (o *TripPlanRecordAllOf) GetWaypoints() []LocationData {
	if o == nil || IsNil(o.Waypoints) {
		var ret []LocationData
		return ret
	}
	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripPlanRecordAllOf) GetWaypointsOk() ([]LocationData, bool) {
	if o == nil || IsNil(o.Waypoints) {
		return nil, false
	}
	return o.Waypoints, true
}

// HasWaypoints returns a boolean if a field has been set.
func (o *TripPlanRecordAllOf) HasWaypoints() bool {
	if o != nil && !IsNil(o.Waypoints) {
		return true
	}

	return false
}

// SetWaypoints gets a reference to the given []LocationData and assigns it to the Waypoints field.
func (o *TripPlanRecordAllOf) SetWaypoints(v []LocationData) {
	o.Waypoints = v
}

// GetDestination returns the Destination field value
func (o *TripPlanRecordAllOf) GetDestination() LocationData {
	if o == nil {
		var ret LocationData
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *TripPlanRecordAllOf) GetDestinationOk() (*LocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *TripPlanRecordAllOf) SetDestination(v LocationData) {
	o.Destination = v
}

// GetEstimatedAverageSpeed returns the EstimatedAverageSpeed field value if set, zero value otherwise.
func (o *TripPlanRecordAllOf) GetEstimatedAverageSpeed() float32 {
	if o == nil || IsNil(o.EstimatedAverageSpeed) {
		var ret float32
		return ret
	}
	return *o.EstimatedAverageSpeed
}

// GetEstimatedAverageSpeedOk returns a tuple with the EstimatedAverageSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripPlanRecordAllOf) GetEstimatedAverageSpeedOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedAverageSpeed) {
		return nil, false
	}
	return o.EstimatedAverageSpeed, true
}

// HasEstimatedAverageSpeed returns a boolean if a field has been set.
func (o *TripPlanRecordAllOf) HasEstimatedAverageSpeed() bool {
	if o != nil && !IsNil(o.EstimatedAverageSpeed) {
		return true
	}

	return false
}

// SetEstimatedAverageSpeed gets a reference to the given float32 and assigns it to the EstimatedAverageSpeed field.
func (o *TripPlanRecordAllOf) SetEstimatedAverageSpeed(v float32) {
	o.EstimatedAverageSpeed = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *TripPlanRecordAllOf) GetEstimatedArrivalTime() time.Time {
	if o == nil || IsNil(o.EstimatedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TripPlanRecordAllOf) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *TripPlanRecordAllOf) HasEstimatedArrivalTime() bool {
	if o != nil && !IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given time.Time and assigns it to the EstimatedArrivalTime field.
func (o *TripPlanRecordAllOf) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = &v
}

func (o TripPlanRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TripPlanRecordAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startingPoint"] = o.StartingPoint
	if !IsNil(o.Waypoints) {
		toSerialize["waypoints"] = o.Waypoints
	}
	toSerialize["destination"] = o.Destination
	if !IsNil(o.EstimatedAverageSpeed) {
		toSerialize["estimatedAverageSpeed"] = o.EstimatedAverageSpeed
	}
	if !IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	}
	return toSerialize, nil
}

type NullableTripPlanRecordAllOf struct {
	value *TripPlanRecordAllOf
	isSet bool
}

func (v NullableTripPlanRecordAllOf) Get() *TripPlanRecordAllOf {
	return v.value
}

func (v *NullableTripPlanRecordAllOf) Set(val *TripPlanRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTripPlanRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTripPlanRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTripPlanRecordAllOf(val *TripPlanRecordAllOf) *NullableTripPlanRecordAllOf {
	return &NullableTripPlanRecordAllOf{value: val, isSet: true}
}

func (v NullableTripPlanRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTripPlanRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
