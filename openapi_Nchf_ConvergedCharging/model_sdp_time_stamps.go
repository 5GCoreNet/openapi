/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// checks if the SDPTimeStamps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDPTimeStamps{}

// SDPTimeStamps struct for SDPTimeStamps
type SDPTimeStamps struct {
	// string with format 'date-time' as defined in OpenAPI.
	SDPOfferTimestamp *time.Time `json:"sDPOfferTimestamp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SDPAnswerTimestamp *time.Time `json:"sDPAnswerTimestamp,omitempty"`
}

// NewSDPTimeStamps instantiates a new SDPTimeStamps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDPTimeStamps() *SDPTimeStamps {
	this := SDPTimeStamps{}
	return &this
}

// NewSDPTimeStampsWithDefaults instantiates a new SDPTimeStamps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDPTimeStampsWithDefaults() *SDPTimeStamps {
	this := SDPTimeStamps{}
	return &this
}

// GetSDPOfferTimestamp returns the SDPOfferTimestamp field value if set, zero value otherwise.
func (o *SDPTimeStamps) GetSDPOfferTimestamp() time.Time {
	if o == nil || IsNil(o.SDPOfferTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SDPOfferTimestamp
}

// GetSDPOfferTimestampOk returns a tuple with the SDPOfferTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPTimeStamps) GetSDPOfferTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SDPOfferTimestamp) {
		return nil, false
	}
	return o.SDPOfferTimestamp, true
}

// HasSDPOfferTimestamp returns a boolean if a field has been set.
func (o *SDPTimeStamps) HasSDPOfferTimestamp() bool {
	if o != nil && !IsNil(o.SDPOfferTimestamp) {
		return true
	}

	return false
}

// SetSDPOfferTimestamp gets a reference to the given time.Time and assigns it to the SDPOfferTimestamp field.
func (o *SDPTimeStamps) SetSDPOfferTimestamp(v time.Time) {
	o.SDPOfferTimestamp = &v
}

// GetSDPAnswerTimestamp returns the SDPAnswerTimestamp field value if set, zero value otherwise.
func (o *SDPTimeStamps) GetSDPAnswerTimestamp() time.Time {
	if o == nil || IsNil(o.SDPAnswerTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.SDPAnswerTimestamp
}

// GetSDPAnswerTimestampOk returns a tuple with the SDPAnswerTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPTimeStamps) GetSDPAnswerTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SDPAnswerTimestamp) {
		return nil, false
	}
	return o.SDPAnswerTimestamp, true
}

// HasSDPAnswerTimestamp returns a boolean if a field has been set.
func (o *SDPTimeStamps) HasSDPAnswerTimestamp() bool {
	if o != nil && !IsNil(o.SDPAnswerTimestamp) {
		return true
	}

	return false
}

// SetSDPAnswerTimestamp gets a reference to the given time.Time and assigns it to the SDPAnswerTimestamp field.
func (o *SDPTimeStamps) SetSDPAnswerTimestamp(v time.Time) {
	o.SDPAnswerTimestamp = &v
}

func (o SDPTimeStamps) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SDPTimeStamps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SDPOfferTimestamp) {
		toSerialize["sDPOfferTimestamp"] = o.SDPOfferTimestamp
	}
	if !IsNil(o.SDPAnswerTimestamp) {
		toSerialize["sDPAnswerTimestamp"] = o.SDPAnswerTimestamp
	}
	return toSerialize, nil
}

type NullableSDPTimeStamps struct {
	value *SDPTimeStamps
	isSet bool
}

func (v NullableSDPTimeStamps) Get() *SDPTimeStamps {
	return v.value
}

func (v *NullableSDPTimeStamps) Set(val *SDPTimeStamps) {
	v.value = val
	v.isSet = true
}

func (v NullableSDPTimeStamps) IsSet() bool {
	return v.isSet
}

func (v *NullableSDPTimeStamps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDPTimeStamps(val *SDPTimeStamps) *NullableSDPTimeStamps {
	return &NullableSDPTimeStamps{value: val, isSet: true}
}

func (v NullableSDPTimeStamps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSDPTimeStamps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
