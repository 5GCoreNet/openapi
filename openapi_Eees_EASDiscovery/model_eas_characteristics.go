/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the EasCharacteristics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasCharacteristics{}

// EasCharacteristics Represents the EAS chararcteristics.
type EasCharacteristics struct {
	// EAS application identifier.
	EasId *string `json:"easId,omitempty"`
	// EAS provider identifier.
	EasProvId *string `json:"easProvId,omitempty"`
	// EAS type.
	EasType  *string         `json:"easType,omitempty"`
	EasSched *TimeWindow     `json:"easSched,omitempty"`
	SvcArea  *LocationArea5G `json:"svcArea,omitempty"`
	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EasSvcContinuity []ACRScenario `json:"easSvcContinuity,omitempty"`
	// Service permissions level.
	SvcPermLevel *string `json:"svcPermLevel,omitempty"`
	// Service features.
	SvcFeats []string `json:"svcFeats,omitempty"`
}

// NewEasCharacteristics instantiates a new EasCharacteristics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasCharacteristics() *EasCharacteristics {
	this := EasCharacteristics{}
	return &this
}

// NewEasCharacteristicsWithDefaults instantiates a new EasCharacteristics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasCharacteristicsWithDefaults() *EasCharacteristics {
	this := EasCharacteristics{}
	return &this
}

// GetEasId returns the EasId field value if set, zero value otherwise.
func (o *EasCharacteristics) GetEasId() string {
	if o == nil || IsNil(o.EasId) {
		var ret string
		return ret
	}
	return *o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetEasIdOk() (*string, bool) {
	if o == nil || IsNil(o.EasId) {
		return nil, false
	}
	return o.EasId, true
}

// HasEasId returns a boolean if a field has been set.
func (o *EasCharacteristics) HasEasId() bool {
	if o != nil && !IsNil(o.EasId) {
		return true
	}

	return false
}

// SetEasId gets a reference to the given string and assigns it to the EasId field.
func (o *EasCharacteristics) SetEasId(v string) {
	o.EasId = &v
}

// GetEasProvId returns the EasProvId field value if set, zero value otherwise.
func (o *EasCharacteristics) GetEasProvId() string {
	if o == nil || IsNil(o.EasProvId) {
		var ret string
		return ret
	}
	return *o.EasProvId
}

// GetEasProvIdOk returns a tuple with the EasProvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetEasProvIdOk() (*string, bool) {
	if o == nil || IsNil(o.EasProvId) {
		return nil, false
	}
	return o.EasProvId, true
}

// HasEasProvId returns a boolean if a field has been set.
func (o *EasCharacteristics) HasEasProvId() bool {
	if o != nil && !IsNil(o.EasProvId) {
		return true
	}

	return false
}

// SetEasProvId gets a reference to the given string and assigns it to the EasProvId field.
func (o *EasCharacteristics) SetEasProvId(v string) {
	o.EasProvId = &v
}

// GetEasType returns the EasType field value if set, zero value otherwise.
func (o *EasCharacteristics) GetEasType() string {
	if o == nil || IsNil(o.EasType) {
		var ret string
		return ret
	}
	return *o.EasType
}

// GetEasTypeOk returns a tuple with the EasType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetEasTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EasType) {
		return nil, false
	}
	return o.EasType, true
}

// HasEasType returns a boolean if a field has been set.
func (o *EasCharacteristics) HasEasType() bool {
	if o != nil && !IsNil(o.EasType) {
		return true
	}

	return false
}

// SetEasType gets a reference to the given string and assigns it to the EasType field.
func (o *EasCharacteristics) SetEasType(v string) {
	o.EasType = &v
}

// GetEasSched returns the EasSched field value if set, zero value otherwise.
func (o *EasCharacteristics) GetEasSched() TimeWindow {
	if o == nil || IsNil(o.EasSched) {
		var ret TimeWindow
		return ret
	}
	return *o.EasSched
}

// GetEasSchedOk returns a tuple with the EasSched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetEasSchedOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.EasSched) {
		return nil, false
	}
	return o.EasSched, true
}

// HasEasSched returns a boolean if a field has been set.
func (o *EasCharacteristics) HasEasSched() bool {
	if o != nil && !IsNil(o.EasSched) {
		return true
	}

	return false
}

// SetEasSched gets a reference to the given TimeWindow and assigns it to the EasSched field.
func (o *EasCharacteristics) SetEasSched(v TimeWindow) {
	o.EasSched = &v
}

// GetSvcArea returns the SvcArea field value if set, zero value otherwise.
func (o *EasCharacteristics) GetSvcArea() LocationArea5G {
	if o == nil || IsNil(o.SvcArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.SvcArea
}

// GetSvcAreaOk returns a tuple with the SvcArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetSvcAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.SvcArea) {
		return nil, false
	}
	return o.SvcArea, true
}

// HasSvcArea returns a boolean if a field has been set.
func (o *EasCharacteristics) HasSvcArea() bool {
	if o != nil && !IsNil(o.SvcArea) {
		return true
	}

	return false
}

// SetSvcArea gets a reference to the given LocationArea5G and assigns it to the SvcArea field.
func (o *EasCharacteristics) SetSvcArea(v LocationArea5G) {
	o.SvcArea = &v
}

// GetEasSvcContinuity returns the EasSvcContinuity field value if set, zero value otherwise.
func (o *EasCharacteristics) GetEasSvcContinuity() []ACRScenario {
	if o == nil || IsNil(o.EasSvcContinuity) {
		var ret []ACRScenario
		return ret
	}
	return o.EasSvcContinuity
}

// GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetEasSvcContinuityOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.EasSvcContinuity) {
		return nil, false
	}
	return o.EasSvcContinuity, true
}

// HasEasSvcContinuity returns a boolean if a field has been set.
func (o *EasCharacteristics) HasEasSvcContinuity() bool {
	if o != nil && !IsNil(o.EasSvcContinuity) {
		return true
	}

	return false
}

// SetEasSvcContinuity gets a reference to the given []ACRScenario and assigns it to the EasSvcContinuity field.
func (o *EasCharacteristics) SetEasSvcContinuity(v []ACRScenario) {
	o.EasSvcContinuity = v
}

// GetSvcPermLevel returns the SvcPermLevel field value if set, zero value otherwise.
func (o *EasCharacteristics) GetSvcPermLevel() string {
	if o == nil || IsNil(o.SvcPermLevel) {
		var ret string
		return ret
	}
	return *o.SvcPermLevel
}

// GetSvcPermLevelOk returns a tuple with the SvcPermLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetSvcPermLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SvcPermLevel) {
		return nil, false
	}
	return o.SvcPermLevel, true
}

// HasSvcPermLevel returns a boolean if a field has been set.
func (o *EasCharacteristics) HasSvcPermLevel() bool {
	if o != nil && !IsNil(o.SvcPermLevel) {
		return true
	}

	return false
}

// SetSvcPermLevel gets a reference to the given string and assigns it to the SvcPermLevel field.
func (o *EasCharacteristics) SetSvcPermLevel(v string) {
	o.SvcPermLevel = &v
}

// GetSvcFeats returns the SvcFeats field value if set, zero value otherwise.
func (o *EasCharacteristics) GetSvcFeats() []string {
	if o == nil || IsNil(o.SvcFeats) {
		var ret []string
		return ret
	}
	return o.SvcFeats
}

// GetSvcFeatsOk returns a tuple with the SvcFeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasCharacteristics) GetSvcFeatsOk() ([]string, bool) {
	if o == nil || IsNil(o.SvcFeats) {
		return nil, false
	}
	return o.SvcFeats, true
}

// HasSvcFeats returns a boolean if a field has been set.
func (o *EasCharacteristics) HasSvcFeats() bool {
	if o != nil && !IsNil(o.SvcFeats) {
		return true
	}

	return false
}

// SetSvcFeats gets a reference to the given []string and assigns it to the SvcFeats field.
func (o *EasCharacteristics) SetSvcFeats(v []string) {
	o.SvcFeats = v
}

func (o EasCharacteristics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasCharacteristics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EasId) {
		toSerialize["easId"] = o.EasId
	}
	if !IsNil(o.EasProvId) {
		toSerialize["easProvId"] = o.EasProvId
	}
	if !IsNil(o.EasType) {
		toSerialize["easType"] = o.EasType
	}
	if !IsNil(o.EasSched) {
		toSerialize["easSched"] = o.EasSched
	}
	if !IsNil(o.SvcArea) {
		toSerialize["svcArea"] = o.SvcArea
	}
	if !IsNil(o.EasSvcContinuity) {
		toSerialize["easSvcContinuity"] = o.EasSvcContinuity
	}
	if !IsNil(o.SvcPermLevel) {
		toSerialize["svcPermLevel"] = o.SvcPermLevel
	}
	if !IsNil(o.SvcFeats) {
		toSerialize["svcFeats"] = o.SvcFeats
	}
	return toSerialize, nil
}

type NullableEasCharacteristics struct {
	value *EasCharacteristics
	isSet bool
}

func (v NullableEasCharacteristics) Get() *EasCharacteristics {
	return v.value
}

func (v *NullableEasCharacteristics) Set(val *EasCharacteristics) {
	v.value = val
	v.isSet = true
}

func (v NullableEasCharacteristics) IsSet() bool {
	return v.isSet
}

func (v *NullableEasCharacteristics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasCharacteristics(val *EasCharacteristics) *NullableEasCharacteristics {
	return &NullableEasCharacteristics{value: val, isSet: true}
}

func (v NullableEasCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasCharacteristics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
