/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
	"time"
)

// checks if the EDNConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EDNConfigInfo{}

// EDNConfigInfo Represents the EDN information.
type EDNConfigInfo struct {
	EdnConInfo EDNConInfo `json:"ednConInfo"`
	// Contains the list of EESs of the EDN.
	Eess []EESInfo `json:"eess"`
	// string with format \"date-time\" as defined in OpenAPI.
	LifeTime *time.Time `json:"lifeTime,omitempty"`
}

// NewEDNConfigInfo instantiates a new EDNConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEDNConfigInfo(ednConInfo EDNConInfo, eess []EESInfo) *EDNConfigInfo {
	this := EDNConfigInfo{}
	this.EdnConInfo = ednConInfo
	this.Eess = eess
	return &this
}

// NewEDNConfigInfoWithDefaults instantiates a new EDNConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEDNConfigInfoWithDefaults() *EDNConfigInfo {
	this := EDNConfigInfo{}
	return &this
}

// GetEdnConInfo returns the EdnConInfo field value
func (o *EDNConfigInfo) GetEdnConInfo() EDNConInfo {
	if o == nil {
		var ret EDNConInfo
		return ret
	}

	return o.EdnConInfo
}

// GetEdnConInfoOk returns a tuple with the EdnConInfo field value
// and a boolean to check if the value has been set.
func (o *EDNConfigInfo) GetEdnConInfoOk() (*EDNConInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdnConInfo, true
}

// SetEdnConInfo sets field value
func (o *EDNConfigInfo) SetEdnConInfo(v EDNConInfo) {
	o.EdnConInfo = v
}

// GetEess returns the Eess field value
func (o *EDNConfigInfo) GetEess() []EESInfo {
	if o == nil {
		var ret []EESInfo
		return ret
	}

	return o.Eess
}

// GetEessOk returns a tuple with the Eess field value
// and a boolean to check if the value has been set.
func (o *EDNConfigInfo) GetEessOk() ([]EESInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Eess, true
}

// SetEess sets field value
func (o *EDNConfigInfo) SetEess(v []EESInfo) {
	o.Eess = v
}

// GetLifeTime returns the LifeTime field value if set, zero value otherwise.
func (o *EDNConfigInfo) GetLifeTime() time.Time {
	if o == nil || IsNil(o.LifeTime) {
		var ret time.Time
		return ret
	}
	return *o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EDNConfigInfo) GetLifeTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LifeTime) {
		return nil, false
	}
	return o.LifeTime, true
}

// HasLifeTime returns a boolean if a field has been set.
func (o *EDNConfigInfo) HasLifeTime() bool {
	if o != nil && !IsNil(o.LifeTime) {
		return true
	}

	return false
}

// SetLifeTime gets a reference to the given time.Time and assigns it to the LifeTime field.
func (o *EDNConfigInfo) SetLifeTime(v time.Time) {
	o.LifeTime = &v
}

func (o EDNConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EDNConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ednConInfo"] = o.EdnConInfo
	toSerialize["eess"] = o.Eess
	if !IsNil(o.LifeTime) {
		toSerialize["lifeTime"] = o.LifeTime
	}
	return toSerialize, nil
}

type NullableEDNConfigInfo struct {
	value *EDNConfigInfo
	isSet bool
}

func (v NullableEDNConfigInfo) Get() *EDNConfigInfo {
	return v.value
}

func (v *NullableEDNConfigInfo) Set(val *EDNConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEDNConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEDNConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEDNConfigInfo(val *EDNConfigInfo) *NullableEDNConfigInfo {
	return &NullableEDNConfigInfo{value: val, isSet: true}
}

func (v NullableEDNConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEDNConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
