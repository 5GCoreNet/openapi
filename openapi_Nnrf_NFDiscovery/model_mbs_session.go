/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the MbsSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSession{}

// MbsSession MBS Session currently served by an MB-SMF
type MbsSession struct {
	MbsSessionId MbsSessionId `json:"mbsSessionId"`
	// A map (list of key-value pairs) where the key identifies an areaSessionId
	MbsAreaSessions *map[string]MbsServiceAreaInfo `json:"mbsAreaSessions,omitempty"`
}

// NewMbsSession instantiates a new MbsSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSession(mbsSessionId MbsSessionId) *MbsSession {
	this := MbsSession{}
	this.MbsSessionId = mbsSessionId
	return &this
}

// NewMbsSessionWithDefaults instantiates a new MbsSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionWithDefaults() *MbsSession {
	this := MbsSession{}
	return &this
}

// GetMbsSessionId returns the MbsSessionId field value
func (o *MbsSession) GetMbsSessionId() MbsSessionId {
	if o == nil {
		var ret MbsSessionId
		return ret
	}

	return o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value
// and a boolean to check if the value has been set.
func (o *MbsSession) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsSessionId, true
}

// SetMbsSessionId sets field value
func (o *MbsSession) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = v
}

// GetMbsAreaSessions returns the MbsAreaSessions field value if set, zero value otherwise.
func (o *MbsSession) GetMbsAreaSessions() map[string]MbsServiceAreaInfo {
	if o == nil || IsNil(o.MbsAreaSessions) {
		var ret map[string]MbsServiceAreaInfo
		return ret
	}
	return *o.MbsAreaSessions
}

// GetMbsAreaSessionsOk returns a tuple with the MbsAreaSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSession) GetMbsAreaSessionsOk() (*map[string]MbsServiceAreaInfo, bool) {
	if o == nil || IsNil(o.MbsAreaSessions) {
		return nil, false
	}
	return o.MbsAreaSessions, true
}

// HasMbsAreaSessions returns a boolean if a field has been set.
func (o *MbsSession) HasMbsAreaSessions() bool {
	if o != nil && !IsNil(o.MbsAreaSessions) {
		return true
	}

	return false
}

// SetMbsAreaSessions gets a reference to the given map[string]MbsServiceAreaInfo and assigns it to the MbsAreaSessions field.
func (o *MbsSession) SetMbsAreaSessions(v map[string]MbsServiceAreaInfo) {
	o.MbsAreaSessions = &v
}

func (o MbsSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsSessionId"] = o.MbsSessionId
	if !IsNil(o.MbsAreaSessions) {
		toSerialize["mbsAreaSessions"] = o.MbsAreaSessions
	}
	return toSerialize, nil
}

type NullableMbsSession struct {
	value *MbsSession
	isSet bool
}

func (v NullableMbsSession) Get() *MbsSession {
	return v.value
}

func (v *NullableMbsSession) Set(val *MbsSession) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSession) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSession(val *MbsSession) *NullableMbsSession {
	return &NullableMbsSession{value: val, isSet: true}
}

func (v NullableMbsSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
