/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
	"time"
)

// checks if the PerformanceDataCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceDataCollection{}

// PerformanceDataCollection Contains Performance Data Analytics related information collection.
type PerformanceDataCollection struct {
	// String providing an application identifier.
	AppId           *string         `json:"appId,omitempty"`
	UeIpAddr        *IpAddr         `json:"ueIpAddr,omitempty"`
	IpTrafficFilter *FlowInfo       `json:"ipTrafficFilter,omitempty"`
	UeLoc           *LocationArea5G `json:"ueLoc,omitempty"`
	AppLocs         []string        `json:"appLocs,omitempty"`
	AsAddr          *AddrFqdn       `json:"asAddr,omitempty"`
	PerfData        PerformanceData `json:"perfData"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}

// NewPerformanceDataCollection instantiates a new PerformanceDataCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceDataCollection(perfData PerformanceData, timeStamp time.Time) *PerformanceDataCollection {
	this := PerformanceDataCollection{}
	this.PerfData = perfData
	this.TimeStamp = timeStamp
	return &this
}

// NewPerformanceDataCollectionWithDefaults instantiates a new PerformanceDataCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceDataCollectionWithDefaults() *PerformanceDataCollection {
	this := PerformanceDataCollection{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *PerformanceDataCollection) SetAppId(v string) {
	o.AppId = &v
}

// GetUeIpAddr returns the UeIpAddr field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetUeIpAddr() IpAddr {
	if o == nil || IsNil(o.UeIpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.UeIpAddr
}

// GetUeIpAddrOk returns a tuple with the UeIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetUeIpAddrOk() (*IpAddr, bool) {
	if o == nil || IsNil(o.UeIpAddr) {
		return nil, false
	}
	return o.UeIpAddr, true
}

// HasUeIpAddr returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasUeIpAddr() bool {
	if o != nil && !IsNil(o.UeIpAddr) {
		return true
	}

	return false
}

// SetUeIpAddr gets a reference to the given IpAddr and assigns it to the UeIpAddr field.
func (o *PerformanceDataCollection) SetUeIpAddr(v IpAddr) {
	o.UeIpAddr = &v
}

// GetIpTrafficFilter returns the IpTrafficFilter field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetIpTrafficFilter() FlowInfo {
	if o == nil || IsNil(o.IpTrafficFilter) {
		var ret FlowInfo
		return ret
	}
	return *o.IpTrafficFilter
}

// GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetIpTrafficFilterOk() (*FlowInfo, bool) {
	if o == nil || IsNil(o.IpTrafficFilter) {
		return nil, false
	}
	return o.IpTrafficFilter, true
}

// HasIpTrafficFilter returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasIpTrafficFilter() bool {
	if o != nil && !IsNil(o.IpTrafficFilter) {
		return true
	}

	return false
}

// SetIpTrafficFilter gets a reference to the given FlowInfo and assigns it to the IpTrafficFilter field.
func (o *PerformanceDataCollection) SetIpTrafficFilter(v FlowInfo) {
	o.IpTrafficFilter = &v
}

// GetUeLoc returns the UeLoc field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetUeLoc() LocationArea5G {
	if o == nil || IsNil(o.UeLoc) {
		var ret LocationArea5G
		return ret
	}
	return *o.UeLoc
}

// GetUeLocOk returns a tuple with the UeLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetUeLocOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.UeLoc) {
		return nil, false
	}
	return o.UeLoc, true
}

// HasUeLoc returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasUeLoc() bool {
	if o != nil && !IsNil(o.UeLoc) {
		return true
	}

	return false
}

// SetUeLoc gets a reference to the given LocationArea5G and assigns it to the UeLoc field.
func (o *PerformanceDataCollection) SetUeLoc(v LocationArea5G) {
	o.UeLoc = &v
}

// GetAppLocs returns the AppLocs field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetAppLocs() []string {
	if o == nil || IsNil(o.AppLocs) {
		var ret []string
		return ret
	}
	return o.AppLocs
}

// GetAppLocsOk returns a tuple with the AppLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetAppLocsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppLocs) {
		return nil, false
	}
	return o.AppLocs, true
}

// HasAppLocs returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasAppLocs() bool {
	if o != nil && !IsNil(o.AppLocs) {
		return true
	}

	return false
}

// SetAppLocs gets a reference to the given []string and assigns it to the AppLocs field.
func (o *PerformanceDataCollection) SetAppLocs(v []string) {
	o.AppLocs = v
}

// GetAsAddr returns the AsAddr field value if set, zero value otherwise.
func (o *PerformanceDataCollection) GetAsAddr() AddrFqdn {
	if o == nil || IsNil(o.AsAddr) {
		var ret AddrFqdn
		return ret
	}
	return *o.AsAddr
}

// GetAsAddrOk returns a tuple with the AsAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetAsAddrOk() (*AddrFqdn, bool) {
	if o == nil || IsNil(o.AsAddr) {
		return nil, false
	}
	return o.AsAddr, true
}

// HasAsAddr returns a boolean if a field has been set.
func (o *PerformanceDataCollection) HasAsAddr() bool {
	if o != nil && !IsNil(o.AsAddr) {
		return true
	}

	return false
}

// SetAsAddr gets a reference to the given AddrFqdn and assigns it to the AsAddr field.
func (o *PerformanceDataCollection) SetAsAddr(v AddrFqdn) {
	o.AsAddr = &v
}

// GetPerfData returns the PerfData field value
func (o *PerformanceDataCollection) GetPerfData() PerformanceData {
	if o == nil {
		var ret PerformanceData
		return ret
	}

	return o.PerfData
}

// GetPerfDataOk returns a tuple with the PerfData field value
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetPerfDataOk() (*PerformanceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerfData, true
}

// SetPerfData sets field value
func (o *PerformanceDataCollection) SetPerfData(v PerformanceData) {
	o.PerfData = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *PerformanceDataCollection) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *PerformanceDataCollection) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *PerformanceDataCollection) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

func (o PerformanceDataCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceDataCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.UeIpAddr) {
		toSerialize["ueIpAddr"] = o.UeIpAddr
	}
	if !IsNil(o.IpTrafficFilter) {
		toSerialize["ipTrafficFilter"] = o.IpTrafficFilter
	}
	if !IsNil(o.UeLoc) {
		toSerialize["ueLoc"] = o.UeLoc
	}
	if !IsNil(o.AppLocs) {
		toSerialize["appLocs"] = o.AppLocs
	}
	if !IsNil(o.AsAddr) {
		toSerialize["asAddr"] = o.AsAddr
	}
	toSerialize["perfData"] = o.PerfData
	toSerialize["timeStamp"] = o.TimeStamp
	return toSerialize, nil
}

type NullablePerformanceDataCollection struct {
	value *PerformanceDataCollection
	isSet bool
}

func (v NullablePerformanceDataCollection) Get() *PerformanceDataCollection {
	return v.value
}

func (v *NullablePerformanceDataCollection) Set(val *PerformanceDataCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceDataCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceDataCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceDataCollection(val *PerformanceDataCollection) *NullablePerformanceDataCollection {
	return &NullablePerformanceDataCollection{value: val, isSet: true}
}

func (v NullablePerformanceDataCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceDataCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
