/*
VAE_MessageDelivery

API for VAE Message Delivery Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_MessageDelivery

import (
	"encoding/json"
	"time"
)

// checks if the DownlinkMessageDeliveryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownlinkMessageDeliveryData{}

// DownlinkMessageDeliveryData Contains the downlink V2X message delivery data.
type DownlinkMessageDeliveryData struct {
	// Represents the identifier of the V2X UE.
	UeId *string `json:"ueId,omitempty"`
	// Represents the group ID for which a V2X message is addressed.
	GroupId *string `json:"groupId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Duration *time.Time `json:"duration,omitempty"`
	// Represents a geographical area identifier.
	GeoId *string `json:"geoId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	Payload string `json:"payload"`
}

// NewDownlinkMessageDeliveryData instantiates a new DownlinkMessageDeliveryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownlinkMessageDeliveryData(payload string) *DownlinkMessageDeliveryData {
	this := DownlinkMessageDeliveryData{}
	this.Payload = payload
	return &this
}

// NewDownlinkMessageDeliveryDataWithDefaults instantiates a new DownlinkMessageDeliveryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownlinkMessageDeliveryDataWithDefaults() *DownlinkMessageDeliveryData {
	this := DownlinkMessageDeliveryData{}
	return &this
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *DownlinkMessageDeliveryData) GetUeId() string {
	if o == nil || IsNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownlinkMessageDeliveryData) GetUeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *DownlinkMessageDeliveryData) HasUeId() bool {
	if o != nil && !IsNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *DownlinkMessageDeliveryData) SetUeId(v string) {
	o.UeId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *DownlinkMessageDeliveryData) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownlinkMessageDeliveryData) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *DownlinkMessageDeliveryData) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *DownlinkMessageDeliveryData) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DownlinkMessageDeliveryData) GetDuration() time.Time {
	if o == nil || IsNil(o.Duration) {
		var ret time.Time
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownlinkMessageDeliveryData) GetDurationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DownlinkMessageDeliveryData) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given time.Time and assigns it to the Duration field.
func (o *DownlinkMessageDeliveryData) SetDuration(v time.Time) {
	o.Duration = &v
}

// GetGeoId returns the GeoId field value if set, zero value otherwise.
func (o *DownlinkMessageDeliveryData) GetGeoId() string {
	if o == nil || IsNil(o.GeoId) {
		var ret string
		return ret
	}
	return *o.GeoId
}

// GetGeoIdOk returns a tuple with the GeoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownlinkMessageDeliveryData) GetGeoIdOk() (*string, bool) {
	if o == nil || IsNil(o.GeoId) {
		return nil, false
	}
	return o.GeoId, true
}

// HasGeoId returns a boolean if a field has been set.
func (o *DownlinkMessageDeliveryData) HasGeoId() bool {
	if o != nil && !IsNil(o.GeoId) {
		return true
	}

	return false
}

// SetGeoId gets a reference to the given string and assigns it to the GeoId field.
func (o *DownlinkMessageDeliveryData) SetGeoId(v string) {
	o.GeoId = &v
}

// GetPayload returns the Payload field value
func (o *DownlinkMessageDeliveryData) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *DownlinkMessageDeliveryData) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *DownlinkMessageDeliveryData) SetPayload(v string) {
	o.Payload = v
}

func (o DownlinkMessageDeliveryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownlinkMessageDeliveryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.GeoId) {
		toSerialize["geoId"] = o.GeoId
	}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullableDownlinkMessageDeliveryData struct {
	value *DownlinkMessageDeliveryData
	isSet bool
}

func (v NullableDownlinkMessageDeliveryData) Get() *DownlinkMessageDeliveryData {
	return v.value
}

func (v *NullableDownlinkMessageDeliveryData) Set(val *DownlinkMessageDeliveryData) {
	v.value = val
	v.isSet = true
}

func (v NullableDownlinkMessageDeliveryData) IsSet() bool {
	return v.isSet
}

func (v *NullableDownlinkMessageDeliveryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownlinkMessageDeliveryData(val *DownlinkMessageDeliveryData) *NullableDownlinkMessageDeliveryData {
	return &NullableDownlinkMessageDeliveryData{value: val, isSet: true}
}

func (v NullableDownlinkMessageDeliveryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownlinkMessageDeliveryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


