/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the UpuData2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpuData2{}

// UpuData2 Contains UE parameters update data set (e.g., the updated Routing ID Data or the Default configured NSSAI).
type UpuData2 struct {
	// Contains a secure packet.
	SecPacket        *string  `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	// Represents a routing indicator.
	RoutingId *string `json:"routingId,omitempty"`
}

// NewUpuData2 instantiates a new UpuData2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpuData2() *UpuData2 {
	this := UpuData2{}
	return &this
}

// NewUpuData2WithDefaults instantiates a new UpuData2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpuData2WithDefaults() *UpuData2 {
	this := UpuData2{}
	return &this
}

// GetSecPacket returns the SecPacket field value if set, zero value otherwise.
func (o *UpuData2) GetSecPacket() string {
	if o == nil || IsNil(o.SecPacket) {
		var ret string
		return ret
	}
	return *o.SecPacket
}

// GetSecPacketOk returns a tuple with the SecPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData2) GetSecPacketOk() (*string, bool) {
	if o == nil || IsNil(o.SecPacket) {
		return nil, false
	}
	return o.SecPacket, true
}

// HasSecPacket returns a boolean if a field has been set.
func (o *UpuData2) HasSecPacket() bool {
	if o != nil && !IsNil(o.SecPacket) {
		return true
	}

	return false
}

// SetSecPacket gets a reference to the given string and assigns it to the SecPacket field.
func (o *UpuData2) SetSecPacket(v string) {
	o.SecPacket = &v
}

// GetDefaultConfNssai returns the DefaultConfNssai field value if set, zero value otherwise.
func (o *UpuData2) GetDefaultConfNssai() []Snssai {
	if o == nil || IsNil(o.DefaultConfNssai) {
		var ret []Snssai
		return ret
	}
	return o.DefaultConfNssai
}

// GetDefaultConfNssaiOk returns a tuple with the DefaultConfNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData2) GetDefaultConfNssaiOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.DefaultConfNssai) {
		return nil, false
	}
	return o.DefaultConfNssai, true
}

// HasDefaultConfNssai returns a boolean if a field has been set.
func (o *UpuData2) HasDefaultConfNssai() bool {
	if o != nil && !IsNil(o.DefaultConfNssai) {
		return true
	}

	return false
}

// SetDefaultConfNssai gets a reference to the given []Snssai and assigns it to the DefaultConfNssai field.
func (o *UpuData2) SetDefaultConfNssai(v []Snssai) {
	o.DefaultConfNssai = v
}

// GetRoutingId returns the RoutingId field value if set, zero value otherwise.
func (o *UpuData2) GetRoutingId() string {
	if o == nil || IsNil(o.RoutingId) {
		var ret string
		return ret
	}
	return *o.RoutingId
}

// GetRoutingIdOk returns a tuple with the RoutingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData2) GetRoutingIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingId) {
		return nil, false
	}
	return o.RoutingId, true
}

// HasRoutingId returns a boolean if a field has been set.
func (o *UpuData2) HasRoutingId() bool {
	if o != nil && !IsNil(o.RoutingId) {
		return true
	}

	return false
}

// SetRoutingId gets a reference to the given string and assigns it to the RoutingId field.
func (o *UpuData2) SetRoutingId(v string) {
	o.RoutingId = &v
}

func (o UpuData2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpuData2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecPacket) {
		toSerialize["secPacket"] = o.SecPacket
	}
	if !IsNil(o.DefaultConfNssai) {
		toSerialize["defaultConfNssai"] = o.DefaultConfNssai
	}
	if !IsNil(o.RoutingId) {
		toSerialize["routingId"] = o.RoutingId
	}
	return toSerialize, nil
}

type NullableUpuData2 struct {
	value *UpuData2
	isSet bool
}

func (v NullableUpuData2) Get() *UpuData2 {
	return v.value
}

func (v *NullableUpuData2) Set(val *UpuData2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpuData2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpuData2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpuData2(val *UpuData2) *NullableUpuData2 {
	return &NullableUpuData2{value: val, isSet: true}
}

func (v NullableUpuData2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpuData2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
