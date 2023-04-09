/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the ApplicationVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationVolume{}

// ApplicationVolume Application data volume per Application Id.
type ApplicationVolume struct {
	// String providing an application identifier.
	AppId string `json:"appId"`
	// Unsigned integer identifying a volume in units of bytes.
	AppVolume int64 `json:"appVolume"`
}

// NewApplicationVolume instantiates a new ApplicationVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationVolume(appId string, appVolume int64) *ApplicationVolume {
	this := ApplicationVolume{}
	this.AppId = appId
	this.AppVolume = appVolume
	return &this
}

// NewApplicationVolumeWithDefaults instantiates a new ApplicationVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationVolumeWithDefaults() *ApplicationVolume {
	this := ApplicationVolume{}
	return &this
}

// GetAppId returns the AppId field value
func (o *ApplicationVolume) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *ApplicationVolume) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *ApplicationVolume) SetAppId(v string) {
	o.AppId = v
}

// GetAppVolume returns the AppVolume field value
func (o *ApplicationVolume) GetAppVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AppVolume
}

// GetAppVolumeOk returns a tuple with the AppVolume field value
// and a boolean to check if the value has been set.
func (o *ApplicationVolume) GetAppVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppVolume, true
}

// SetAppVolume sets field value
func (o *ApplicationVolume) SetAppVolume(v int64) {
	o.AppVolume = v
}

func (o ApplicationVolume) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	toSerialize["appVolume"] = o.AppVolume
	return toSerialize, nil
}

type NullableApplicationVolume struct {
	value *ApplicationVolume
	isSet bool
}

func (v NullableApplicationVolume) Get() *ApplicationVolume {
	return v.value
}

func (v *NullableApplicationVolume) Set(val *ApplicationVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationVolume(val *ApplicationVolume) *NullableApplicationVolume {
	return &NullableApplicationVolume{value: val, isSet: true}
}

func (v NullableApplicationVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
