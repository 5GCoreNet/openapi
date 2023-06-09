/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the MonLocAreaInterestFltr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonLocAreaInterestFltr{}

// MonLocAreaInterestFltr Filter information indicate the area of interest and triggering events.
type MonLocAreaInterestFltr struct {
	LocInfoCri LocationInfoCriteria `json:"locInfoCri"`
	// Triggering events when to send information.
	TrigEvnts []MonLocTriggerEvent `json:"trigEvnts,omitempty"`
}

// NewMonLocAreaInterestFltr instantiates a new MonLocAreaInterestFltr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonLocAreaInterestFltr(locInfoCri LocationInfoCriteria) *MonLocAreaInterestFltr {
	this := MonLocAreaInterestFltr{}
	this.LocInfoCri = locInfoCri
	return &this
}

// NewMonLocAreaInterestFltrWithDefaults instantiates a new MonLocAreaInterestFltr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonLocAreaInterestFltrWithDefaults() *MonLocAreaInterestFltr {
	this := MonLocAreaInterestFltr{}
	return &this
}

// GetLocInfoCri returns the LocInfoCri field value
func (o *MonLocAreaInterestFltr) GetLocInfoCri() LocationInfoCriteria {
	if o == nil {
		var ret LocationInfoCriteria
		return ret
	}

	return o.LocInfoCri
}

// GetLocInfoCriOk returns a tuple with the LocInfoCri field value
// and a boolean to check if the value has been set.
func (o *MonLocAreaInterestFltr) GetLocInfoCriOk() (*LocationInfoCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocInfoCri, true
}

// SetLocInfoCri sets field value
func (o *MonLocAreaInterestFltr) SetLocInfoCri(v LocationInfoCriteria) {
	o.LocInfoCri = v
}

// GetTrigEvnts returns the TrigEvnts field value if set, zero value otherwise.
func (o *MonLocAreaInterestFltr) GetTrigEvnts() []MonLocTriggerEvent {
	if o == nil || IsNil(o.TrigEvnts) {
		var ret []MonLocTriggerEvent
		return ret
	}
	return o.TrigEvnts
}

// GetTrigEvntsOk returns a tuple with the TrigEvnts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonLocAreaInterestFltr) GetTrigEvntsOk() ([]MonLocTriggerEvent, bool) {
	if o == nil || IsNil(o.TrigEvnts) {
		return nil, false
	}
	return o.TrigEvnts, true
}

// HasTrigEvnts returns a boolean if a field has been set.
func (o *MonLocAreaInterestFltr) HasTrigEvnts() bool {
	if o != nil && !IsNil(o.TrigEvnts) {
		return true
	}

	return false
}

// SetTrigEvnts gets a reference to the given []MonLocTriggerEvent and assigns it to the TrigEvnts field.
func (o *MonLocAreaInterestFltr) SetTrigEvnts(v []MonLocTriggerEvent) {
	o.TrigEvnts = v
}

func (o MonLocAreaInterestFltr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonLocAreaInterestFltr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locInfoCri"] = o.LocInfoCri
	if !IsNil(o.TrigEvnts) {
		toSerialize["trigEvnts"] = o.TrigEvnts
	}
	return toSerialize, nil
}

type NullableMonLocAreaInterestFltr struct {
	value *MonLocAreaInterestFltr
	isSet bool
}

func (v NullableMonLocAreaInterestFltr) Get() *MonLocAreaInterestFltr {
	return v.value
}

func (v *NullableMonLocAreaInterestFltr) Set(val *MonLocAreaInterestFltr) {
	v.value = val
	v.isSet = true
}

func (v NullableMonLocAreaInterestFltr) IsSet() bool {
	return v.isSet
}

func (v *NullableMonLocAreaInterestFltr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonLocAreaInterestFltr(val *MonLocAreaInterestFltr) *NullableMonLocAreaInterestFltr {
	return &NullableMonLocAreaInterestFltr{value: val, isSet: true}
}

func (v NullableMonLocAreaInterestFltr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonLocAreaInterestFltr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
