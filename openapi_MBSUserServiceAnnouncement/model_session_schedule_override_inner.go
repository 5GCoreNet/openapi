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

// checks if the SessionScheduleOverrideInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionScheduleOverrideInner{}

// SessionScheduleOverrideInner struct for SessionScheduleOverrideInner
type SessionScheduleOverrideInner struct {
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Stop      *time.Time `json:"stop,omitempty"`
	Index     *int32     `json:"index,omitempty"`
	Cancelled *bool      `json:"cancelled,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SessionDescriptionURI *string `json:"sessionDescriptionURI,omitempty"`
}

// NewSessionScheduleOverrideInner instantiates a new SessionScheduleOverrideInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionScheduleOverrideInner() *SessionScheduleOverrideInner {
	this := SessionScheduleOverrideInner{}
	return &this
}

// NewSessionScheduleOverrideInnerWithDefaults instantiates a new SessionScheduleOverrideInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionScheduleOverrideInnerWithDefaults() *SessionScheduleOverrideInner {
	this := SessionScheduleOverrideInner{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SessionScheduleOverrideInner) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleOverrideInner) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SessionScheduleOverrideInner) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *SessionScheduleOverrideInner) SetStart(v time.Time) {
	o.Start = &v
}

// GetStop returns the Stop field value if set, zero value otherwise.
func (o *SessionScheduleOverrideInner) GetStop() time.Time {
	if o == nil || IsNil(o.Stop) {
		var ret time.Time
		return ret
	}
	return *o.Stop
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleOverrideInner) GetStopOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Stop) {
		return nil, false
	}
	return o.Stop, true
}

// HasStop returns a boolean if a field has been set.
func (o *SessionScheduleOverrideInner) HasStop() bool {
	if o != nil && !IsNil(o.Stop) {
		return true
	}

	return false
}

// SetStop gets a reference to the given time.Time and assigns it to the Stop field.
func (o *SessionScheduleOverrideInner) SetStop(v time.Time) {
	o.Stop = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SessionScheduleOverrideInner) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleOverrideInner) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SessionScheduleOverrideInner) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *SessionScheduleOverrideInner) SetIndex(v int32) {
	o.Index = &v
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *SessionScheduleOverrideInner) GetCancelled() bool {
	if o == nil || IsNil(o.Cancelled) {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleOverrideInner) GetCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancelled) {
		return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *SessionScheduleOverrideInner) HasCancelled() bool {
	if o != nil && !IsNil(o.Cancelled) {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *SessionScheduleOverrideInner) SetCancelled(v bool) {
	o.Cancelled = &v
}

// GetSessionDescriptionURI returns the SessionDescriptionURI field value if set, zero value otherwise.
func (o *SessionScheduleOverrideInner) GetSessionDescriptionURI() string {
	if o == nil || IsNil(o.SessionDescriptionURI) {
		var ret string
		return ret
	}
	return *o.SessionDescriptionURI
}

// GetSessionDescriptionURIOk returns a tuple with the SessionDescriptionURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionScheduleOverrideInner) GetSessionDescriptionURIOk() (*string, bool) {
	if o == nil || IsNil(o.SessionDescriptionURI) {
		return nil, false
	}
	return o.SessionDescriptionURI, true
}

// HasSessionDescriptionURI returns a boolean if a field has been set.
func (o *SessionScheduleOverrideInner) HasSessionDescriptionURI() bool {
	if o != nil && !IsNil(o.SessionDescriptionURI) {
		return true
	}

	return false
}

// SetSessionDescriptionURI gets a reference to the given string and assigns it to the SessionDescriptionURI field.
func (o *SessionScheduleOverrideInner) SetSessionDescriptionURI(v string) {
	o.SessionDescriptionURI = &v
}

func (o SessionScheduleOverrideInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionScheduleOverrideInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Stop) {
		toSerialize["stop"] = o.Stop
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Cancelled) {
		toSerialize["cancelled"] = o.Cancelled
	}
	if !IsNil(o.SessionDescriptionURI) {
		toSerialize["sessionDescriptionURI"] = o.SessionDescriptionURI
	}
	return toSerialize, nil
}

type NullableSessionScheduleOverrideInner struct {
	value *SessionScheduleOverrideInner
	isSet bool
}

func (v NullableSessionScheduleOverrideInner) Get() *SessionScheduleOverrideInner {
	return v.value
}

func (v *NullableSessionScheduleOverrideInner) Set(val *SessionScheduleOverrideInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionScheduleOverrideInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionScheduleOverrideInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionScheduleOverrideInner(val *SessionScheduleOverrideInner) *NullableSessionScheduleOverrideInner {
	return &NullableSessionScheduleOverrideInner{value: val, isSet: true}
}

func (v NullableSessionScheduleOverrideInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionScheduleOverrideInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
