/*
ECS EES Registration_API

API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_EESRegistration

import (
	"encoding/json"
)

// checks if the EESProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EESProfile{}

// EESProfile Represents the EES profile information.
type EESProfile struct {
	// Identifier of the EES.
	EesId string   `json:"eesId"`
	EndPt EndPoint `json:"endPt"`
	// Application identifiers of EASs that are registered with EES.
	EasIds []string `json:"easIds,omitempty"`
	// Identifier of the ECSP that provides the EES provider.
	ProvId  *string      `json:"provId,omitempty"`
	SvcArea *ServiceArea `json:"svcArea,omitempty"`
	// List of DNAI(s) associated with the EES.
	AppLocs []string `json:"appLocs,omitempty"`
	// The ACR scenarios supported by the EES for service continuity.
	SvcContSupp []ACRScenario `json:"svcContSupp,omitempty"`
	// Set to true if the EEC is required to register to the EES to use edge service. Set to false if the EEC is not required to register to use edge services.
	EecRegConf bool `json:"eecRegConf"`
}

// NewEESProfile instantiates a new EESProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEESProfile(eesId string, endPt EndPoint, eecRegConf bool) *EESProfile {
	this := EESProfile{}
	this.EesId = eesId
	this.EndPt = endPt
	this.EecRegConf = eecRegConf
	return &this
}

// NewEESProfileWithDefaults instantiates a new EESProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEESProfileWithDefaults() *EESProfile {
	this := EESProfile{}
	return &this
}

// GetEesId returns the EesId field value
func (o *EESProfile) GetEesId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EesId
}

// GetEesIdOk returns a tuple with the EesId field value
// and a boolean to check if the value has been set.
func (o *EESProfile) GetEesIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EesId, true
}

// SetEesId sets field value
func (o *EESProfile) SetEesId(v string) {
	o.EesId = v
}

// GetEndPt returns the EndPt field value
func (o *EESProfile) GetEndPt() EndPoint {
	if o == nil {
		var ret EndPoint
		return ret
	}

	return o.EndPt
}

// GetEndPtOk returns a tuple with the EndPt field value
// and a boolean to check if the value has been set.
func (o *EESProfile) GetEndPtOk() (*EndPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPt, true
}

// SetEndPt sets field value
func (o *EESProfile) SetEndPt(v EndPoint) {
	o.EndPt = v
}

// GetEasIds returns the EasIds field value if set, zero value otherwise.
func (o *EESProfile) GetEasIds() []string {
	if o == nil || IsNil(o.EasIds) {
		var ret []string
		return ret
	}
	return o.EasIds
}

// GetEasIdsOk returns a tuple with the EasIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESProfile) GetEasIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EasIds) {
		return nil, false
	}
	return o.EasIds, true
}

// HasEasIds returns a boolean if a field has been set.
func (o *EESProfile) HasEasIds() bool {
	if o != nil && !IsNil(o.EasIds) {
		return true
	}

	return false
}

// SetEasIds gets a reference to the given []string and assigns it to the EasIds field.
func (o *EESProfile) SetEasIds(v []string) {
	o.EasIds = v
}

// GetProvId returns the ProvId field value if set, zero value otherwise.
func (o *EESProfile) GetProvId() string {
	if o == nil || IsNil(o.ProvId) {
		var ret string
		return ret
	}
	return *o.ProvId
}

// GetProvIdOk returns a tuple with the ProvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESProfile) GetProvIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProvId) {
		return nil, false
	}
	return o.ProvId, true
}

// HasProvId returns a boolean if a field has been set.
func (o *EESProfile) HasProvId() bool {
	if o != nil && !IsNil(o.ProvId) {
		return true
	}

	return false
}

// SetProvId gets a reference to the given string and assigns it to the ProvId field.
func (o *EESProfile) SetProvId(v string) {
	o.ProvId = &v
}

// GetSvcArea returns the SvcArea field value if set, zero value otherwise.
func (o *EESProfile) GetSvcArea() ServiceArea {
	if o == nil || IsNil(o.SvcArea) {
		var ret ServiceArea
		return ret
	}
	return *o.SvcArea
}

// GetSvcAreaOk returns a tuple with the SvcArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESProfile) GetSvcAreaOk() (*ServiceArea, bool) {
	if o == nil || IsNil(o.SvcArea) {
		return nil, false
	}
	return o.SvcArea, true
}

// HasSvcArea returns a boolean if a field has been set.
func (o *EESProfile) HasSvcArea() bool {
	if o != nil && !IsNil(o.SvcArea) {
		return true
	}

	return false
}

// SetSvcArea gets a reference to the given ServiceArea and assigns it to the SvcArea field.
func (o *EESProfile) SetSvcArea(v ServiceArea) {
	o.SvcArea = &v
}

// GetAppLocs returns the AppLocs field value if set, zero value otherwise.
func (o *EESProfile) GetAppLocs() []string {
	if o == nil || IsNil(o.AppLocs) {
		var ret []string
		return ret
	}
	return o.AppLocs
}

// GetAppLocsOk returns a tuple with the AppLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESProfile) GetAppLocsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppLocs) {
		return nil, false
	}
	return o.AppLocs, true
}

// HasAppLocs returns a boolean if a field has been set.
func (o *EESProfile) HasAppLocs() bool {
	if o != nil && !IsNil(o.AppLocs) {
		return true
	}

	return false
}

// SetAppLocs gets a reference to the given []string and assigns it to the AppLocs field.
func (o *EESProfile) SetAppLocs(v []string) {
	o.AppLocs = v
}

// GetSvcContSupp returns the SvcContSupp field value if set, zero value otherwise.
func (o *EESProfile) GetSvcContSupp() []ACRScenario {
	if o == nil || IsNil(o.SvcContSupp) {
		var ret []ACRScenario
		return ret
	}
	return o.SvcContSupp
}

// GetSvcContSuppOk returns a tuple with the SvcContSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EESProfile) GetSvcContSuppOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.SvcContSupp) {
		return nil, false
	}
	return o.SvcContSupp, true
}

// HasSvcContSupp returns a boolean if a field has been set.
func (o *EESProfile) HasSvcContSupp() bool {
	if o != nil && !IsNil(o.SvcContSupp) {
		return true
	}

	return false
}

// SetSvcContSupp gets a reference to the given []ACRScenario and assigns it to the SvcContSupp field.
func (o *EESProfile) SetSvcContSupp(v []ACRScenario) {
	o.SvcContSupp = v
}

// GetEecRegConf returns the EecRegConf field value
func (o *EESProfile) GetEecRegConf() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EecRegConf
}

// GetEecRegConfOk returns a tuple with the EecRegConf field value
// and a boolean to check if the value has been set.
func (o *EESProfile) GetEecRegConfOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EecRegConf, true
}

// SetEecRegConf sets field value
func (o *EESProfile) SetEecRegConf(v bool) {
	o.EecRegConf = v
}

func (o EESProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EESProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eesId"] = o.EesId
	toSerialize["endPt"] = o.EndPt
	if !IsNil(o.EasIds) {
		toSerialize["easIds"] = o.EasIds
	}
	if !IsNil(o.ProvId) {
		toSerialize["provId"] = o.ProvId
	}
	if !IsNil(o.SvcArea) {
		toSerialize["svcArea"] = o.SvcArea
	}
	if !IsNil(o.AppLocs) {
		toSerialize["appLocs"] = o.AppLocs
	}
	if !IsNil(o.SvcContSupp) {
		toSerialize["svcContSupp"] = o.SvcContSupp
	}
	toSerialize["eecRegConf"] = o.EecRegConf
	return toSerialize, nil
}

type NullableEESProfile struct {
	value *EESProfile
	isSet bool
}

func (v NullableEESProfile) Get() *EESProfile {
	return v.value
}

func (v *NullableEESProfile) Set(val *EESProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEESProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEESProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEESProfile(val *EESProfile) *NullableEESProfile {
	return &NullableEESProfile{value: val, isSet: true}
}

func (v NullableEESProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEESProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
