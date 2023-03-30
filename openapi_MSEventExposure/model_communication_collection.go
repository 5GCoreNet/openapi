/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
	"time"
)

// checks if the CommunicationCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationCollection{}

// CommunicationCollection Contains communication information.
type CommunicationCollection struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTime time.Time `json:"endTime"`
	// Unsigned integer identifying a volume in units of bytes.
	UlVol int64 `json:"ulVol"`
	// Unsigned integer identifying a volume in units of bytes.
	DlVol int64 `json:"dlVol"`
}

// NewCommunicationCollection instantiates a new CommunicationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationCollection(startTime time.Time, endTime time.Time, ulVol int64, dlVol int64) *CommunicationCollection {
	this := CommunicationCollection{}
	this.StartTime = startTime
	this.EndTime = endTime
	this.UlVol = ulVol
	this.DlVol = dlVol
	return &this
}

// NewCommunicationCollectionWithDefaults instantiates a new CommunicationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationCollectionWithDefaults() *CommunicationCollection {
	this := CommunicationCollection{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *CommunicationCollection) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *CommunicationCollection) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *CommunicationCollection) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *CommunicationCollection) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *CommunicationCollection) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *CommunicationCollection) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetUlVol returns the UlVol field value
func (o *CommunicationCollection) GetUlVol() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UlVol
}

// GetUlVolOk returns a tuple with the UlVol field value
// and a boolean to check if the value has been set.
func (o *CommunicationCollection) GetUlVolOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UlVol, true
}

// SetUlVol sets field value
func (o *CommunicationCollection) SetUlVol(v int64) {
	o.UlVol = v
}

// GetDlVol returns the DlVol field value
func (o *CommunicationCollection) GetDlVol() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DlVol
}

// GetDlVolOk returns a tuple with the DlVol field value
// and a boolean to check if the value has been set.
func (o *CommunicationCollection) GetDlVolOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DlVol, true
}

// SetDlVol sets field value
func (o *CommunicationCollection) SetDlVol(v int64) {
	o.DlVol = v
}

func (o CommunicationCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["ulVol"] = o.UlVol
	toSerialize["dlVol"] = o.DlVol
	return toSerialize, nil
}

type NullableCommunicationCollection struct {
	value *CommunicationCollection
	isSet bool
}

func (v NullableCommunicationCollection) Get() *CommunicationCollection {
	return v.value
}

func (v *NullableCommunicationCollection) Set(val *CommunicationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationCollection(val *CommunicationCollection) *NullableCommunicationCollection {
	return &NullableCommunicationCollection{value: val, isSet: true}
}

func (v NullableCommunicationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


