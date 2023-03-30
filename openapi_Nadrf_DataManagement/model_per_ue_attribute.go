/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the PerUeAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerUeAttribute{}

// PerUeAttribute UE application data collected per UE.
type PerUeAttribute struct {
	UeDest *LocationArea5G `json:"ueDest,omitempty"`
	Route *string `json:"route,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgSpeed *string `json:"avgSpeed,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeOfArrival *time.Time `json:"timeOfArrival,omitempty"`
}

// NewPerUeAttribute instantiates a new PerUeAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerUeAttribute() *PerUeAttribute {
	this := PerUeAttribute{}
	return &this
}

// NewPerUeAttributeWithDefaults instantiates a new PerUeAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerUeAttributeWithDefaults() *PerUeAttribute {
	this := PerUeAttribute{}
	return &this
}

// GetUeDest returns the UeDest field value if set, zero value otherwise.
func (o *PerUeAttribute) GetUeDest() LocationArea5G {
	if o == nil || IsNil(o.UeDest) {
		var ret LocationArea5G
		return ret
	}
	return *o.UeDest
}

// GetUeDestOk returns a tuple with the UeDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerUeAttribute) GetUeDestOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.UeDest) {
		return nil, false
	}
	return o.UeDest, true
}

// HasUeDest returns a boolean if a field has been set.
func (o *PerUeAttribute) HasUeDest() bool {
	if o != nil && !IsNil(o.UeDest) {
		return true
	}

	return false
}

// SetUeDest gets a reference to the given LocationArea5G and assigns it to the UeDest field.
func (o *PerUeAttribute) SetUeDest(v LocationArea5G) {
	o.UeDest = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *PerUeAttribute) GetRoute() string {
	if o == nil || IsNil(o.Route) {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerUeAttribute) GetRouteOk() (*string, bool) {
	if o == nil || IsNil(o.Route) {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *PerUeAttribute) HasRoute() bool {
	if o != nil && !IsNil(o.Route) {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *PerUeAttribute) SetRoute(v string) {
	o.Route = &v
}

// GetAvgSpeed returns the AvgSpeed field value if set, zero value otherwise.
func (o *PerUeAttribute) GetAvgSpeed() string {
	if o == nil || IsNil(o.AvgSpeed) {
		var ret string
		return ret
	}
	return *o.AvgSpeed
}

// GetAvgSpeedOk returns a tuple with the AvgSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerUeAttribute) GetAvgSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.AvgSpeed) {
		return nil, false
	}
	return o.AvgSpeed, true
}

// HasAvgSpeed returns a boolean if a field has been set.
func (o *PerUeAttribute) HasAvgSpeed() bool {
	if o != nil && !IsNil(o.AvgSpeed) {
		return true
	}

	return false
}

// SetAvgSpeed gets a reference to the given string and assigns it to the AvgSpeed field.
func (o *PerUeAttribute) SetAvgSpeed(v string) {
	o.AvgSpeed = &v
}

// GetTimeOfArrival returns the TimeOfArrival field value if set, zero value otherwise.
func (o *PerUeAttribute) GetTimeOfArrival() time.Time {
	if o == nil || IsNil(o.TimeOfArrival) {
		var ret time.Time
		return ret
	}
	return *o.TimeOfArrival
}

// GetTimeOfArrivalOk returns a tuple with the TimeOfArrival field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerUeAttribute) GetTimeOfArrivalOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeOfArrival) {
		return nil, false
	}
	return o.TimeOfArrival, true
}

// HasTimeOfArrival returns a boolean if a field has been set.
func (o *PerUeAttribute) HasTimeOfArrival() bool {
	if o != nil && !IsNil(o.TimeOfArrival) {
		return true
	}

	return false
}

// SetTimeOfArrival gets a reference to the given time.Time and assigns it to the TimeOfArrival field.
func (o *PerUeAttribute) SetTimeOfArrival(v time.Time) {
	o.TimeOfArrival = &v
}

func (o PerUeAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerUeAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeDest) {
		toSerialize["ueDest"] = o.UeDest
	}
	if !IsNil(o.Route) {
		toSerialize["route"] = o.Route
	}
	if !IsNil(o.AvgSpeed) {
		toSerialize["avgSpeed"] = o.AvgSpeed
	}
	if !IsNil(o.TimeOfArrival) {
		toSerialize["timeOfArrival"] = o.TimeOfArrival
	}
	return toSerialize, nil
}

type NullablePerUeAttribute struct {
	value *PerUeAttribute
	isSet bool
}

func (v NullablePerUeAttribute) Get() *PerUeAttribute {
	return v.value
}

func (v *NullablePerUeAttribute) Set(val *PerUeAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullablePerUeAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullablePerUeAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerUeAttribute(val *PerUeAttribute) *NullablePerUeAttribute {
	return &NullablePerUeAttribute{value: val, isSet: true}
}

func (v NullablePerUeAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerUeAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


