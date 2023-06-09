/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CpProvisioning

import (
	"encoding/json"
)

// checks if the UmtLocationArea5GAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmtLocationArea5GAllOf{}

// UmtLocationArea5GAllOf struct for UmtLocationArea5GAllOf
type UmtLocationArea5GAllOf struct {
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	UmtTime *string `json:"umtTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	UmtDuration *int32 `json:"umtDuration,omitempty"`
}

// NewUmtLocationArea5GAllOf instantiates a new UmtLocationArea5GAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmtLocationArea5GAllOf() *UmtLocationArea5GAllOf {
	this := UmtLocationArea5GAllOf{}
	return &this
}

// NewUmtLocationArea5GAllOfWithDefaults instantiates a new UmtLocationArea5GAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmtLocationArea5GAllOfWithDefaults() *UmtLocationArea5GAllOf {
	this := UmtLocationArea5GAllOf{}
	return &this
}

// GetUmtTime returns the UmtTime field value if set, zero value otherwise.
func (o *UmtLocationArea5GAllOf) GetUmtTime() string {
	if o == nil || IsNil(o.UmtTime) {
		var ret string
		return ret
	}
	return *o.UmtTime
}

// GetUmtTimeOk returns a tuple with the UmtTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmtLocationArea5GAllOf) GetUmtTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UmtTime) {
		return nil, false
	}
	return o.UmtTime, true
}

// HasUmtTime returns a boolean if a field has been set.
func (o *UmtLocationArea5GAllOf) HasUmtTime() bool {
	if o != nil && !IsNil(o.UmtTime) {
		return true
	}

	return false
}

// SetUmtTime gets a reference to the given string and assigns it to the UmtTime field.
func (o *UmtLocationArea5GAllOf) SetUmtTime(v string) {
	o.UmtTime = &v
}

// GetUmtDuration returns the UmtDuration field value if set, zero value otherwise.
func (o *UmtLocationArea5GAllOf) GetUmtDuration() int32 {
	if o == nil || IsNil(o.UmtDuration) {
		var ret int32
		return ret
	}
	return *o.UmtDuration
}

// GetUmtDurationOk returns a tuple with the UmtDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmtLocationArea5GAllOf) GetUmtDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.UmtDuration) {
		return nil, false
	}
	return o.UmtDuration, true
}

// HasUmtDuration returns a boolean if a field has been set.
func (o *UmtLocationArea5GAllOf) HasUmtDuration() bool {
	if o != nil && !IsNil(o.UmtDuration) {
		return true
	}

	return false
}

// SetUmtDuration gets a reference to the given int32 and assigns it to the UmtDuration field.
func (o *UmtLocationArea5GAllOf) SetUmtDuration(v int32) {
	o.UmtDuration = &v
}

func (o UmtLocationArea5GAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmtLocationArea5GAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UmtTime) {
		toSerialize["umtTime"] = o.UmtTime
	}
	if !IsNil(o.UmtDuration) {
		toSerialize["umtDuration"] = o.UmtDuration
	}
	return toSerialize, nil
}

type NullableUmtLocationArea5GAllOf struct {
	value *UmtLocationArea5GAllOf
	isSet bool
}

func (v NullableUmtLocationArea5GAllOf) Get() *UmtLocationArea5GAllOf {
	return v.value
}

func (v *NullableUmtLocationArea5GAllOf) Set(val *UmtLocationArea5GAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUmtLocationArea5GAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUmtLocationArea5GAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmtLocationArea5GAllOf(val *UmtLocationArea5GAllOf) *NullableUmtLocationArea5GAllOf {
	return &NullableUmtLocationArea5GAllOf{value: val, isSet: true}
}

func (v NullableUmtLocationArea5GAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmtLocationArea5GAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
