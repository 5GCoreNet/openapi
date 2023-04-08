/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"time"
)

// checks if the PerformanceDataInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceDataInfo{}

// PerformanceDataInfo Contains Performance Data Analytics related information collection.
type PerformanceDataInfo struct {
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
	UserLoc *UserLocation `json:"userLoc,omitempty"`
	AppLocs []string `json:"appLocs,omitempty"`
	AsAddr *AddrFqdn `json:"asAddr,omitempty"`
	PerfData PerformanceData `json:"perfData"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}

// NewPerformanceDataInfo instantiates a new PerformanceDataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceDataInfo(perfData PerformanceData, timeStamp time.Time) *PerformanceDataInfo {
	this := PerformanceDataInfo{}
	this.PerfData = perfData
	this.TimeStamp = timeStamp
	return &this
}

// NewPerformanceDataInfoWithDefaults instantiates a new PerformanceDataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceDataInfoWithDefaults() *PerformanceDataInfo {
	this := PerformanceDataInfo{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *PerformanceDataInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetUeIpAddr returns the UeIpAddr field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetUeIpAddr() IpAddr {
	if o == nil || isNil(o.UeIpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.UeIpAddr
}

// GetUeIpAddrOk returns a tuple with the UeIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetUeIpAddrOk() (*IpAddr, bool) {
	if o == nil || isNil(o.UeIpAddr) {
		return nil, false
	}
	return o.UeIpAddr, true
}

// HasUeIpAddr returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasUeIpAddr() bool {
	if o != nil && !isNil(o.UeIpAddr) {
		return true
	}

	return false
}

// SetUeIpAddr gets a reference to the given IpAddr and assigns it to the UeIpAddr field.
func (o *PerformanceDataInfo) SetUeIpAddr(v IpAddr) {
	o.UeIpAddr = &v
}

// GetIpTrafficFilter returns the IpTrafficFilter field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetIpTrafficFilter() FlowInfo {
	if o == nil || isNil(o.IpTrafficFilter) {
		var ret FlowInfo
		return ret
	}
	return *o.IpTrafficFilter
}

// GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetIpTrafficFilterOk() (*FlowInfo, bool) {
	if o == nil || isNil(o.IpTrafficFilter) {
		return nil, false
	}
	return o.IpTrafficFilter, true
}

// HasIpTrafficFilter returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasIpTrafficFilter() bool {
	if o != nil && !isNil(o.IpTrafficFilter) {
		return true
	}

	return false
}

// SetIpTrafficFilter gets a reference to the given FlowInfo and assigns it to the IpTrafficFilter field.
func (o *PerformanceDataInfo) SetIpTrafficFilter(v FlowInfo) {
	o.IpTrafficFilter = &v
}

// GetUserLoc returns the UserLoc field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetUserLoc() UserLocation {
	if o == nil || isNil(o.UserLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UserLoc
}

// GetUserLocOk returns a tuple with the UserLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetUserLocOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UserLoc) {
		return nil, false
	}
	return o.UserLoc, true
}

// HasUserLoc returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasUserLoc() bool {
	if o != nil && !isNil(o.UserLoc) {
		return true
	}

	return false
}

// SetUserLoc gets a reference to the given UserLocation and assigns it to the UserLoc field.
func (o *PerformanceDataInfo) SetUserLoc(v UserLocation) {
	o.UserLoc = &v
}

// GetAppLocs returns the AppLocs field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetAppLocs() []string {
	if o == nil || isNil(o.AppLocs) {
		var ret []string
		return ret
	}
	return o.AppLocs
}

// GetAppLocsOk returns a tuple with the AppLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetAppLocsOk() ([]string, bool) {
	if o == nil || isNil(o.AppLocs) {
		return nil, false
	}
	return o.AppLocs, true
}

// HasAppLocs returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasAppLocs() bool {
	if o != nil && !isNil(o.AppLocs) {
		return true
	}

	return false
}

// SetAppLocs gets a reference to the given []string and assigns it to the AppLocs field.
func (o *PerformanceDataInfo) SetAppLocs(v []string) {
	o.AppLocs = v
}

// GetAsAddr returns the AsAddr field value if set, zero value otherwise.
func (o *PerformanceDataInfo) GetAsAddr() AddrFqdn {
	if o == nil || isNil(o.AsAddr) {
		var ret AddrFqdn
		return ret
	}
	return *o.AsAddr
}

// GetAsAddrOk returns a tuple with the AsAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetAsAddrOk() (*AddrFqdn, bool) {
	if o == nil || isNil(o.AsAddr) {
		return nil, false
	}
	return o.AsAddr, true
}

// HasAsAddr returns a boolean if a field has been set.
func (o *PerformanceDataInfo) HasAsAddr() bool {
	if o != nil && !isNil(o.AsAddr) {
		return true
	}

	return false
}

// SetAsAddr gets a reference to the given AddrFqdn and assigns it to the AsAddr field.
func (o *PerformanceDataInfo) SetAsAddr(v AddrFqdn) {
	o.AsAddr = &v
}

// GetPerfData returns the PerfData field value
func (o *PerformanceDataInfo) GetPerfData() PerformanceData {
	if o == nil {
		var ret PerformanceData
		return ret
	}

	return o.PerfData
}

// GetPerfDataOk returns a tuple with the PerfData field value
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetPerfDataOk() (*PerformanceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerfData, true
}

// SetPerfData sets field value
func (o *PerformanceDataInfo) SetPerfData(v PerformanceData) {
	o.PerfData = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *PerformanceDataInfo) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *PerformanceDataInfo) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *PerformanceDataInfo) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

func (o PerformanceDataInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceDataInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.UeIpAddr) {
		toSerialize["ueIpAddr"] = o.UeIpAddr
	}
	if !isNil(o.IpTrafficFilter) {
		toSerialize["ipTrafficFilter"] = o.IpTrafficFilter
	}
	if !isNil(o.UserLoc) {
		toSerialize["userLoc"] = o.UserLoc
	}
	if !isNil(o.AppLocs) {
		toSerialize["appLocs"] = o.AppLocs
	}
	if !isNil(o.AsAddr) {
		toSerialize["asAddr"] = o.AsAddr
	}
	toSerialize["perfData"] = o.PerfData
	toSerialize["timeStamp"] = o.TimeStamp
	return toSerialize, nil
}

type NullablePerformanceDataInfo struct {
	value *PerformanceDataInfo
	isSet bool
}

func (v NullablePerformanceDataInfo) Get() *PerformanceDataInfo {
	return v.value
}

func (v *NullablePerformanceDataInfo) Set(val *PerformanceDataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceDataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceDataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceDataInfo(val *PerformanceDataInfo) *NullablePerformanceDataInfo {
	return &NullablePerformanceDataInfo{value: val, isSet: true}
}

func (v NullablePerformanceDataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceDataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


