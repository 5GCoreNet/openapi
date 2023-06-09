/*
UAE Server C2 Operation Mode Management Service

UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_C2OperationModeManagement

import (
	"encoding/json"
)

// checks if the C2CommModeSwitchNotif type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C2CommModeSwitchNotif{}

// C2CommModeSwitchNotif Represents information on the targeted C2 Communication Mode switching for a UAS (i.e. pair of UAV and UAV-C).
type C2CommModeSwitchNotif struct {
	// string providing an URI formatted according to IETF RFC 3986.
	UaeServerId          string              `json:"uaeServerId"`
	UasId                UasId               `json:"uasId"`
	C2CommModeSwitchType C2CommModeSwitching `json:"c2CommModeSwitchType"`
	SwitchingCause       *C2SwitchingCause   `json:"switchingCause,omitempty"`
}

// NewC2CommModeSwitchNotif instantiates a new C2CommModeSwitchNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC2CommModeSwitchNotif(uaeServerId string, uasId UasId, c2CommModeSwitchType C2CommModeSwitching) *C2CommModeSwitchNotif {
	this := C2CommModeSwitchNotif{}
	this.UaeServerId = uaeServerId
	this.UasId = uasId
	this.C2CommModeSwitchType = c2CommModeSwitchType
	return &this
}

// NewC2CommModeSwitchNotifWithDefaults instantiates a new C2CommModeSwitchNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC2CommModeSwitchNotifWithDefaults() *C2CommModeSwitchNotif {
	this := C2CommModeSwitchNotif{}
	return &this
}

// GetUaeServerId returns the UaeServerId field value
func (o *C2CommModeSwitchNotif) GetUaeServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UaeServerId
}

// GetUaeServerIdOk returns a tuple with the UaeServerId field value
// and a boolean to check if the value has been set.
func (o *C2CommModeSwitchNotif) GetUaeServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UaeServerId, true
}

// SetUaeServerId sets field value
func (o *C2CommModeSwitchNotif) SetUaeServerId(v string) {
	o.UaeServerId = v
}

// GetUasId returns the UasId field value
func (o *C2CommModeSwitchNotif) GetUasId() UasId {
	if o == nil {
		var ret UasId
		return ret
	}

	return o.UasId
}

// GetUasIdOk returns a tuple with the UasId field value
// and a boolean to check if the value has been set.
func (o *C2CommModeSwitchNotif) GetUasIdOk() (*UasId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UasId, true
}

// SetUasId sets field value
func (o *C2CommModeSwitchNotif) SetUasId(v UasId) {
	o.UasId = v
}

// GetC2CommModeSwitchType returns the C2CommModeSwitchType field value
func (o *C2CommModeSwitchNotif) GetC2CommModeSwitchType() C2CommModeSwitching {
	if o == nil {
		var ret C2CommModeSwitching
		return ret
	}

	return o.C2CommModeSwitchType
}

// GetC2CommModeSwitchTypeOk returns a tuple with the C2CommModeSwitchType field value
// and a boolean to check if the value has been set.
func (o *C2CommModeSwitchNotif) GetC2CommModeSwitchTypeOk() (*C2CommModeSwitching, bool) {
	if o == nil {
		return nil, false
	}
	return &o.C2CommModeSwitchType, true
}

// SetC2CommModeSwitchType sets field value
func (o *C2CommModeSwitchNotif) SetC2CommModeSwitchType(v C2CommModeSwitching) {
	o.C2CommModeSwitchType = v
}

// GetSwitchingCause returns the SwitchingCause field value if set, zero value otherwise.
func (o *C2CommModeSwitchNotif) GetSwitchingCause() C2SwitchingCause {
	if o == nil || IsNil(o.SwitchingCause) {
		var ret C2SwitchingCause
		return ret
	}
	return *o.SwitchingCause
}

// GetSwitchingCauseOk returns a tuple with the SwitchingCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C2CommModeSwitchNotif) GetSwitchingCauseOk() (*C2SwitchingCause, bool) {
	if o == nil || IsNil(o.SwitchingCause) {
		return nil, false
	}
	return o.SwitchingCause, true
}

// HasSwitchingCause returns a boolean if a field has been set.
func (o *C2CommModeSwitchNotif) HasSwitchingCause() bool {
	if o != nil && !IsNil(o.SwitchingCause) {
		return true
	}

	return false
}

// SetSwitchingCause gets a reference to the given C2SwitchingCause and assigns it to the SwitchingCause field.
func (o *C2CommModeSwitchNotif) SetSwitchingCause(v C2SwitchingCause) {
	o.SwitchingCause = &v
}

func (o C2CommModeSwitchNotif) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C2CommModeSwitchNotif) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uaeServerId"] = o.UaeServerId
	toSerialize["uasId"] = o.UasId
	toSerialize["c2CommModeSwitchType"] = o.C2CommModeSwitchType
	if !IsNil(o.SwitchingCause) {
		toSerialize["switchingCause"] = o.SwitchingCause
	}
	return toSerialize, nil
}

type NullableC2CommModeSwitchNotif struct {
	value *C2CommModeSwitchNotif
	isSet bool
}

func (v NullableC2CommModeSwitchNotif) Get() *C2CommModeSwitchNotif {
	return v.value
}

func (v *NullableC2CommModeSwitchNotif) Set(val *C2CommModeSwitchNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableC2CommModeSwitchNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableC2CommModeSwitchNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC2CommModeSwitchNotif(val *C2CommModeSwitchNotif) *NullableC2CommModeSwitchNotif {
	return &NullableC2CommModeSwitchNotif{value: val, isSet: true}
}

func (v NullableC2CommModeSwitchNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC2CommModeSwitchNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
