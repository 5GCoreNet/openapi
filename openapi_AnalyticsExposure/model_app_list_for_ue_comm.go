/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"time"
)

// checks if the AppListForUeComm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppListForUeComm{}

// AppListForUeComm Represents the analytics of the application list used by UE.
type AppListForUeComm struct {
	// String providing an application identifier.
	AppId string `json:"appId"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	// indicating a time in seconds.
	AppDur *int32 `json:"appDur,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	OccurRatio      *int32           `json:"occurRatio,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty"`
}

// NewAppListForUeComm instantiates a new AppListForUeComm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppListForUeComm(appId string) *AppListForUeComm {
	this := AppListForUeComm{}
	this.AppId = appId
	return &this
}

// NewAppListForUeCommWithDefaults instantiates a new AppListForUeComm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppListForUeCommWithDefaults() *AppListForUeComm {
	this := AppListForUeComm{}
	return &this
}

// GetAppId returns the AppId field value
func (o *AppListForUeComm) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AppListForUeComm) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AppListForUeComm) SetAppId(v string) {
	o.AppId = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AppListForUeComm) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppListForUeComm) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AppListForUeComm) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AppListForUeComm) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetAppDur returns the AppDur field value if set, zero value otherwise.
func (o *AppListForUeComm) GetAppDur() int32 {
	if o == nil || IsNil(o.AppDur) {
		var ret int32
		return ret
	}
	return *o.AppDur
}

// GetAppDurOk returns a tuple with the AppDur field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppListForUeComm) GetAppDurOk() (*int32, bool) {
	if o == nil || IsNil(o.AppDur) {
		return nil, false
	}
	return o.AppDur, true
}

// HasAppDur returns a boolean if a field has been set.
func (o *AppListForUeComm) HasAppDur() bool {
	if o != nil && !IsNil(o.AppDur) {
		return true
	}

	return false
}

// SetAppDur gets a reference to the given int32 and assigns it to the AppDur field.
func (o *AppListForUeComm) SetAppDur(v int32) {
	o.AppDur = &v
}

// GetOccurRatio returns the OccurRatio field value if set, zero value otherwise.
func (o *AppListForUeComm) GetOccurRatio() int32 {
	if o == nil || IsNil(o.OccurRatio) {
		var ret int32
		return ret
	}
	return *o.OccurRatio
}

// GetOccurRatioOk returns a tuple with the OccurRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppListForUeComm) GetOccurRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.OccurRatio) {
		return nil, false
	}
	return o.OccurRatio, true
}

// HasOccurRatio returns a boolean if a field has been set.
func (o *AppListForUeComm) HasOccurRatio() bool {
	if o != nil && !IsNil(o.OccurRatio) {
		return true
	}

	return false
}

// SetOccurRatio gets a reference to the given int32 and assigns it to the OccurRatio field.
func (o *AppListForUeComm) SetOccurRatio(v int32) {
	o.OccurRatio = &v
}

// GetSpatialValidity returns the SpatialValidity field value if set, zero value otherwise.
func (o *AppListForUeComm) GetSpatialValidity() NetworkAreaInfo {
	if o == nil || IsNil(o.SpatialValidity) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.SpatialValidity
}

// GetSpatialValidityOk returns a tuple with the SpatialValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppListForUeComm) GetSpatialValidityOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.SpatialValidity) {
		return nil, false
	}
	return o.SpatialValidity, true
}

// HasSpatialValidity returns a boolean if a field has been set.
func (o *AppListForUeComm) HasSpatialValidity() bool {
	if o != nil && !IsNil(o.SpatialValidity) {
		return true
	}

	return false
}

// SetSpatialValidity gets a reference to the given NetworkAreaInfo and assigns it to the SpatialValidity field.
func (o *AppListForUeComm) SetSpatialValidity(v NetworkAreaInfo) {
	o.SpatialValidity = &v
}

func (o AppListForUeComm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppListForUeComm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.AppDur) {
		toSerialize["appDur"] = o.AppDur
	}
	if !IsNil(o.OccurRatio) {
		toSerialize["occurRatio"] = o.OccurRatio
	}
	if !IsNil(o.SpatialValidity) {
		toSerialize["spatialValidity"] = o.SpatialValidity
	}
	return toSerialize, nil
}

type NullableAppListForUeComm struct {
	value *AppListForUeComm
	isSet bool
}

func (v NullableAppListForUeComm) Get() *AppListForUeComm {
	return v.value
}

func (v *NullableAppListForUeComm) Set(val *AppListForUeComm) {
	v.value = val
	v.isSet = true
}

func (v NullableAppListForUeComm) IsSet() bool {
	return v.isSet
}

func (v *NullableAppListForUeComm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppListForUeComm(val *AppListForUeComm) *NullableAppListForUeComm {
	return &NullableAppListForUeComm{value: val, isSet: true}
}

func (v NullableAppListForUeComm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppListForUeComm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
