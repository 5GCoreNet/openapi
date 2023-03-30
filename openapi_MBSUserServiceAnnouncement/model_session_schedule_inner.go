/*
MBS User Service Announcement Element units’ definition

MBS User Service Announcement Element units. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserServiceAnnouncement

import (
	"encoding/json"
	"time"
)

// checks if the SessionScheduleInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionScheduleInner{}

// SessionScheduleInner struct for SessionScheduleInner
type SessionScheduleInner struct {
	// string with format 'date-time' as defined in OpenAPI.
	Start time.Time `json:"start"`
	// string with format 'date-time' as defined in OpenAPI.
	Stop time.Time `json:"stop"`
	ReoccurencePattern *string `json:"reoccurencePattern,omitempty"`
	NumberOfTimes *int32 `json:"numberOfTimes,omitempty"`
	ReoccurenceStopTime *string `json:"reoccurenceStopTime,omitempty"`
	Index *int32 `json:"index,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	FDTInstanceURI *string `json:"FDTInstanceURI,omitempty"`
}

// NewSessionScheduleInner instantiates a new SessionScheduleInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionScheduleInner(start time.Time, stop time.Time) *SessionScheduleInner {
	this := SessionScheduleInner{}
	this.Start = start
	this.Stop = stop
	return &this
}

// NewSessionScheduleInnerWithDefaults instantiates a new SessionScheduleInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionScheduleInnerWithDefaults() *SessionScheduleInner {
	this := SessionScheduleInner{}
	return &this
}

// GetStart returns the Start field value
func (o *SessionScheduleInner) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *SessionScheduleInner) SetStart(v time.Time) {
	o.Start = v
}

// GetStop returns the Stop field value
func (o *SessionScheduleInner) GetStop() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Stop
}

// GetStopOk returns a tuple with the Stop field value
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetStopOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stop, true
}

// SetStop sets field value
func (o *SessionScheduleInner) SetStop(v time.Time) {
	o.Stop = v
}

// GetReoccurencePattern returns the ReoccurencePattern field value if set, zero value otherwise.
func (o *SessionScheduleInner) GetReoccurencePattern() string {
	if o == nil || IsNil(o.ReoccurencePattern) {
		var ret string
		return ret
	}
	return *o.ReoccurencePattern
}

// GetReoccurencePatternOk returns a tuple with the ReoccurencePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetReoccurencePatternOk() (*string, bool) {
	if o == nil || IsNil(o.ReoccurencePattern) {
		return nil, false
	}
	return o.ReoccurencePattern, true
}

// HasReoccurencePattern returns a boolean if a field has been set.
func (o *SessionScheduleInner) HasReoccurencePattern() bool {
	if o != nil && !IsNil(o.ReoccurencePattern) {
		return true
	}

	return false
}

// SetReoccurencePattern gets a reference to the given string and assigns it to the ReoccurencePattern field.
func (o *SessionScheduleInner) SetReoccurencePattern(v string) {
	o.ReoccurencePattern = &v
}

// GetNumberOfTimes returns the NumberOfTimes field value if set, zero value otherwise.
func (o *SessionScheduleInner) GetNumberOfTimes() int32 {
	if o == nil || IsNil(o.NumberOfTimes) {
		var ret int32
		return ret
	}
	return *o.NumberOfTimes
}

// GetNumberOfTimesOk returns a tuple with the NumberOfTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetNumberOfTimesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfTimes) {
		return nil, false
	}
	return o.NumberOfTimes, true
}

// HasNumberOfTimes returns a boolean if a field has been set.
func (o *SessionScheduleInner) HasNumberOfTimes() bool {
	if o != nil && !IsNil(o.NumberOfTimes) {
		return true
	}

	return false
}

// SetNumberOfTimes gets a reference to the given int32 and assigns it to the NumberOfTimes field.
func (o *SessionScheduleInner) SetNumberOfTimes(v int32) {
	o.NumberOfTimes = &v
}

// GetReoccurenceStopTime returns the ReoccurenceStopTime field value if set, zero value otherwise.
func (o *SessionScheduleInner) GetReoccurenceStopTime() string {
	if o == nil || IsNil(o.ReoccurenceStopTime) {
		var ret string
		return ret
	}
	return *o.ReoccurenceStopTime
}

// GetReoccurenceStopTimeOk returns a tuple with the ReoccurenceStopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetReoccurenceStopTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ReoccurenceStopTime) {
		return nil, false
	}
	return o.ReoccurenceStopTime, true
}

// HasReoccurenceStopTime returns a boolean if a field has been set.
func (o *SessionScheduleInner) HasReoccurenceStopTime() bool {
	if o != nil && !IsNil(o.ReoccurenceStopTime) {
		return true
	}

	return false
}

// SetReoccurenceStopTime gets a reference to the given string and assigns it to the ReoccurenceStopTime field.
func (o *SessionScheduleInner) SetReoccurenceStopTime(v string) {
	o.ReoccurenceStopTime = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SessionScheduleInner) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SessionScheduleInner) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *SessionScheduleInner) SetIndex(v int32) {
	o.Index = &v
}

// GetFDTInstanceURI returns the FDTInstanceURI field value if set, zero value otherwise.
func (o *SessionScheduleInner) GetFDTInstanceURI() string {
	if o == nil || IsNil(o.FDTInstanceURI) {
		var ret string
		return ret
	}
	return *o.FDTInstanceURI
}

// GetFDTInstanceURIOk returns a tuple with the FDTInstanceURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleInner) GetFDTInstanceURIOk() (*string, bool) {
	if o == nil || IsNil(o.FDTInstanceURI) {
		return nil, false
	}
	return o.FDTInstanceURI, true
}

// HasFDTInstanceURI returns a boolean if a field has been set.
func (o *SessionScheduleInner) HasFDTInstanceURI() bool {
	if o != nil && !IsNil(o.FDTInstanceURI) {
		return true
	}

	return false
}

// SetFDTInstanceURI gets a reference to the given string and assigns it to the FDTInstanceURI field.
func (o *SessionScheduleInner) SetFDTInstanceURI(v string) {
	o.FDTInstanceURI = &v
}

func (o SessionScheduleInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionScheduleInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	toSerialize["stop"] = o.Stop
	if !IsNil(o.ReoccurencePattern) {
		toSerialize["reoccurencePattern"] = o.ReoccurencePattern
	}
	if !IsNil(o.NumberOfTimes) {
		toSerialize["numberOfTimes"] = o.NumberOfTimes
	}
	if !IsNil(o.ReoccurenceStopTime) {
		toSerialize["reoccurenceStopTime"] = o.ReoccurenceStopTime
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.FDTInstanceURI) {
		toSerialize["FDTInstanceURI"] = o.FDTInstanceURI
	}
	return toSerialize, nil
}

type NullableSessionScheduleInner struct {
	value *SessionScheduleInner
	isSet bool
}

func (v NullableSessionScheduleInner) Get() *SessionScheduleInner {
	return v.value
}

func (v *NullableSessionScheduleInner) Set(val *SessionScheduleInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionScheduleInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionScheduleInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionScheduleInner(val *SessionScheduleInner) *NullableSessionScheduleInner {
	return &NullableSessionScheduleInner{value: val, isSet: true}
}

func (v NullableSessionScheduleInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionScheduleInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

