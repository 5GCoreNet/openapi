/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the DnPerf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnPerf{}

// DnPerf Represents DN performance for the application.
type DnPerf struct {
	AppServerInsAddr *AddrFqdn `json:"appServerInsAddr,omitempty"`
	UpfInfo *UpfInformation `json:"upfInfo,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai *string `json:"dnai,omitempty"`
	PerfData PerfData `json:"perfData"`
	SpatialValidCon *NetworkAreaInfo `json:"spatialValidCon,omitempty"`
	TemporalValidCon *TimeWindow `json:"temporalValidCon,omitempty"`
}

// NewDnPerf instantiates a new DnPerf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnPerf(perfData PerfData) *DnPerf {
	this := DnPerf{}
	this.PerfData = perfData
	return &this
}

// NewDnPerfWithDefaults instantiates a new DnPerf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnPerfWithDefaults() *DnPerf {
	this := DnPerf{}
	return &this
}

// GetAppServerInsAddr returns the AppServerInsAddr field value if set, zero value otherwise.
func (o *DnPerf) GetAppServerInsAddr() AddrFqdn {
	if o == nil || isNil(o.AppServerInsAddr) {
		var ret AddrFqdn
		return ret
	}
	return *o.AppServerInsAddr
}

// GetAppServerInsAddrOk returns a tuple with the AppServerInsAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerf) GetAppServerInsAddrOk() (*AddrFqdn, bool) {
	if o == nil || isNil(o.AppServerInsAddr) {
		return nil, false
	}
	return o.AppServerInsAddr, true
}

// HasAppServerInsAddr returns a boolean if a field has been set.
func (o *DnPerf) HasAppServerInsAddr() bool {
	if o != nil && !isNil(o.AppServerInsAddr) {
		return true
	}

	return false
}

// SetAppServerInsAddr gets a reference to the given AddrFqdn and assigns it to the AppServerInsAddr field.
func (o *DnPerf) SetAppServerInsAddr(v AddrFqdn) {
	o.AppServerInsAddr = &v
}

// GetUpfInfo returns the UpfInfo field value if set, zero value otherwise.
func (o *DnPerf) GetUpfInfo() UpfInformation {
	if o == nil || isNil(o.UpfInfo) {
		var ret UpfInformation
		return ret
	}
	return *o.UpfInfo
}

// GetUpfInfoOk returns a tuple with the UpfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerf) GetUpfInfoOk() (*UpfInformation, bool) {
	if o == nil || isNil(o.UpfInfo) {
		return nil, false
	}
	return o.UpfInfo, true
}

// HasUpfInfo returns a boolean if a field has been set.
func (o *DnPerf) HasUpfInfo() bool {
	if o != nil && !isNil(o.UpfInfo) {
		return true
	}

	return false
}

// SetUpfInfo gets a reference to the given UpfInformation and assigns it to the UpfInfo field.
func (o *DnPerf) SetUpfInfo(v UpfInformation) {
	o.UpfInfo = &v
}

// GetDnai returns the Dnai field value if set, zero value otherwise.
func (o *DnPerf) GetDnai() string {
	if o == nil || isNil(o.Dnai) {
		var ret string
		return ret
	}
	return *o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerf) GetDnaiOk() (*string, bool) {
	if o == nil || isNil(o.Dnai) {
		return nil, false
	}
	return o.Dnai, true
}

// HasDnai returns a boolean if a field has been set.
func (o *DnPerf) HasDnai() bool {
	if o != nil && !isNil(o.Dnai) {
		return true
	}

	return false
}

// SetDnai gets a reference to the given string and assigns it to the Dnai field.
func (o *DnPerf) SetDnai(v string) {
	o.Dnai = &v
}

// GetPerfData returns the PerfData field value
func (o *DnPerf) GetPerfData() PerfData {
	if o == nil {
		var ret PerfData
		return ret
	}

	return o.PerfData
}

// GetPerfDataOk returns a tuple with the PerfData field value
// and a boolean to check if the value has been set.
func (o *DnPerf) GetPerfDataOk() (*PerfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerfData, true
}

// SetPerfData sets field value
func (o *DnPerf) SetPerfData(v PerfData) {
	o.PerfData = v
}

// GetSpatialValidCon returns the SpatialValidCon field value if set, zero value otherwise.
func (o *DnPerf) GetSpatialValidCon() NetworkAreaInfo {
	if o == nil || isNil(o.SpatialValidCon) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.SpatialValidCon
}

// GetSpatialValidConOk returns a tuple with the SpatialValidCon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerf) GetSpatialValidConOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.SpatialValidCon) {
		return nil, false
	}
	return o.SpatialValidCon, true
}

// HasSpatialValidCon returns a boolean if a field has been set.
func (o *DnPerf) HasSpatialValidCon() bool {
	if o != nil && !isNil(o.SpatialValidCon) {
		return true
	}

	return false
}

// SetSpatialValidCon gets a reference to the given NetworkAreaInfo and assigns it to the SpatialValidCon field.
func (o *DnPerf) SetSpatialValidCon(v NetworkAreaInfo) {
	o.SpatialValidCon = &v
}

// GetTemporalValidCon returns the TemporalValidCon field value if set, zero value otherwise.
func (o *DnPerf) GetTemporalValidCon() TimeWindow {
	if o == nil || isNil(o.TemporalValidCon) {
		var ret TimeWindow
		return ret
	}
	return *o.TemporalValidCon
}

// GetTemporalValidConOk returns a tuple with the TemporalValidCon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnPerf) GetTemporalValidConOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.TemporalValidCon) {
		return nil, false
	}
	return o.TemporalValidCon, true
}

// HasTemporalValidCon returns a boolean if a field has been set.
func (o *DnPerf) HasTemporalValidCon() bool {
	if o != nil && !isNil(o.TemporalValidCon) {
		return true
	}

	return false
}

// SetTemporalValidCon gets a reference to the given TimeWindow and assigns it to the TemporalValidCon field.
func (o *DnPerf) SetTemporalValidCon(v TimeWindow) {
	o.TemporalValidCon = &v
}

func (o DnPerf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnPerf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppServerInsAddr) {
		toSerialize["appServerInsAddr"] = o.AppServerInsAddr
	}
	if !isNil(o.UpfInfo) {
		toSerialize["upfInfo"] = o.UpfInfo
	}
	if !isNil(o.Dnai) {
		toSerialize["dnai"] = o.Dnai
	}
	toSerialize["perfData"] = o.PerfData
	if !isNil(o.SpatialValidCon) {
		toSerialize["spatialValidCon"] = o.SpatialValidCon
	}
	if !isNil(o.TemporalValidCon) {
		toSerialize["temporalValidCon"] = o.TemporalValidCon
	}
	return toSerialize, nil
}

type NullableDnPerf struct {
	value *DnPerf
	isSet bool
}

func (v NullableDnPerf) Get() *DnPerf {
	return v.value
}

func (v *NullableDnPerf) Set(val *DnPerf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnPerf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnPerf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnPerf(val *DnPerf) *NullableDnPerf {
	return &NullableDnPerf{value: val, isSet: true}
}

func (v NullableDnPerf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnPerf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


