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

// checks if the NsiLoadLevelInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsiLoadLevelInfo{}

// NsiLoadLevelInfo Represents the network slice and optionally the associated network slice instance and the  load level information. 
type NsiLoadLevelInfo struct {
	// Load level information of the network slice and the optionally associated network slice  instance. 
	LoadLevelInformation int32 `json:"loadLevelInformation"`
	Snssai Snssai `json:"snssai"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId *string `json:"nsiId,omitempty"`
	ResUsage *ResourceUsage `json:"resUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumOfExceedLoadLevelThr *int32 `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd *bool `json:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	TimePeriod *TimeWindow `json:"timePeriod,omitempty"`
	// Each element indicates the time elapsed between times each threshold is met or exceeded or crossed. The start time and end time are the exact time stamps of the resource usage threshold is reached or exceeded. May be present if the \"listOfAnaSubsets\" attribute is  provided and the maximum number of instances shall not exceed the value provided in the  \"numOfExceedLoadLevelThr\" attribute. 
	ResUsgThrCrossTimePeriod []TimeWindow `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes *NumberAverage `json:"numOfUes,omitempty"`
	NumOfPduSess *NumberAverage `json:"numOfPduSess,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewNsiLoadLevelInfo instantiates a new NsiLoadLevelInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsiLoadLevelInfo(loadLevelInformation int32, snssai Snssai) *NsiLoadLevelInfo {
	this := NsiLoadLevelInfo{}
	this.LoadLevelInformation = loadLevelInformation
	this.Snssai = snssai
	return &this
}

// NewNsiLoadLevelInfoWithDefaults instantiates a new NsiLoadLevelInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsiLoadLevelInfoWithDefaults() *NsiLoadLevelInfo {
	this := NsiLoadLevelInfo{}
	return &this
}

// GetLoadLevelInformation returns the LoadLevelInformation field value
func (o *NsiLoadLevelInfo) GetLoadLevelInformation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoadLevelInformation
}

// GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field value
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetLoadLevelInformationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadLevelInformation, true
}

// SetLoadLevelInformation sets field value
func (o *NsiLoadLevelInfo) SetLoadLevelInformation(v int32) {
	o.LoadLevelInformation = v
}

// GetSnssai returns the Snssai field value
func (o *NsiLoadLevelInfo) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *NsiLoadLevelInfo) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNsiId() string {
	if o == nil || IsNil(o.NsiId) {
		var ret string
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNsiIdOk() (*string, bool) {
	if o == nil || IsNil(o.NsiId) {
		return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNsiId() bool {
	if o != nil && !IsNil(o.NsiId) {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given string and assigns it to the NsiId field.
func (o *NsiLoadLevelInfo) SetNsiId(v string) {
	o.NsiId = &v
}

// GetResUsage returns the ResUsage field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetResUsage() ResourceUsage {
	if o == nil || IsNil(o.ResUsage) {
		var ret ResourceUsage
		return ret
	}
	return *o.ResUsage
}

// GetResUsageOk returns a tuple with the ResUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetResUsageOk() (*ResourceUsage, bool) {
	if o == nil || IsNil(o.ResUsage) {
		return nil, false
	}
	return o.ResUsage, true
}

// HasResUsage returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasResUsage() bool {
	if o != nil && !IsNil(o.ResUsage) {
		return true
	}

	return false
}

// SetResUsage gets a reference to the given ResourceUsage and assigns it to the ResUsage field.
func (o *NsiLoadLevelInfo) SetResUsage(v ResourceUsage) {
	o.ResUsage = &v
}

// GetNumOfExceedLoadLevelThr returns the NumOfExceedLoadLevelThr field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNumOfExceedLoadLevelThr() int32 {
	if o == nil || IsNil(o.NumOfExceedLoadLevelThr) {
		var ret int32
		return ret
	}
	return *o.NumOfExceedLoadLevelThr
}

// GetNumOfExceedLoadLevelThrOk returns a tuple with the NumOfExceedLoadLevelThr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNumOfExceedLoadLevelThrOk() (*int32, bool) {
	if o == nil || IsNil(o.NumOfExceedLoadLevelThr) {
		return nil, false
	}
	return o.NumOfExceedLoadLevelThr, true
}

// HasNumOfExceedLoadLevelThr returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNumOfExceedLoadLevelThr() bool {
	if o != nil && !IsNil(o.NumOfExceedLoadLevelThr) {
		return true
	}

	return false
}

// SetNumOfExceedLoadLevelThr gets a reference to the given int32 and assigns it to the NumOfExceedLoadLevelThr field.
func (o *NsiLoadLevelInfo) SetNumOfExceedLoadLevelThr(v int32) {
	o.NumOfExceedLoadLevelThr = &v
}

// GetExceedLoadLevelThrInd returns the ExceedLoadLevelThrInd field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetExceedLoadLevelThrInd() bool {
	if o == nil || IsNil(o.ExceedLoadLevelThrInd) {
		var ret bool
		return ret
	}
	return *o.ExceedLoadLevelThrInd
}

// GetExceedLoadLevelThrIndOk returns a tuple with the ExceedLoadLevelThrInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetExceedLoadLevelThrIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceedLoadLevelThrInd) {
		return nil, false
	}
	return o.ExceedLoadLevelThrInd, true
}

// HasExceedLoadLevelThrInd returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasExceedLoadLevelThrInd() bool {
	if o != nil && !IsNil(o.ExceedLoadLevelThrInd) {
		return true
	}

	return false
}

// SetExceedLoadLevelThrInd gets a reference to the given bool and assigns it to the ExceedLoadLevelThrInd field.
func (o *NsiLoadLevelInfo) SetExceedLoadLevelThrInd(v bool) {
	o.ExceedLoadLevelThrInd = &v
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil || IsNil(o.NetworkArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.NetworkArea) {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNetworkArea() bool {
	if o != nil && !IsNil(o.NetworkArea) {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *NsiLoadLevelInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetTimePeriod() TimeWindow {
	if o == nil || IsNil(o.TimePeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetTimePeriodOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.TimePeriod) {
		return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasTimePeriod() bool {
	if o != nil && !IsNil(o.TimePeriod) {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given TimeWindow and assigns it to the TimePeriod field.
func (o *NsiLoadLevelInfo) SetTimePeriod(v TimeWindow) {
	o.TimePeriod = &v
}

// GetResUsgThrCrossTimePeriod returns the ResUsgThrCrossTimePeriod field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetResUsgThrCrossTimePeriod() []TimeWindow {
	if o == nil || IsNil(o.ResUsgThrCrossTimePeriod) {
		var ret []TimeWindow
		return ret
	}
	return o.ResUsgThrCrossTimePeriod
}

// GetResUsgThrCrossTimePeriodOk returns a tuple with the ResUsgThrCrossTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetResUsgThrCrossTimePeriodOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.ResUsgThrCrossTimePeriod) {
		return nil, false
	}
	return o.ResUsgThrCrossTimePeriod, true
}

// HasResUsgThrCrossTimePeriod returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasResUsgThrCrossTimePeriod() bool {
	if o != nil && !IsNil(o.ResUsgThrCrossTimePeriod) {
		return true
	}

	return false
}

// SetResUsgThrCrossTimePeriod gets a reference to the given []TimeWindow and assigns it to the ResUsgThrCrossTimePeriod field.
func (o *NsiLoadLevelInfo) SetResUsgThrCrossTimePeriod(v []TimeWindow) {
	o.ResUsgThrCrossTimePeriod = v
}

// GetNumOfUes returns the NumOfUes field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNumOfUes() NumberAverage {
	if o == nil || IsNil(o.NumOfUes) {
		var ret NumberAverage
		return ret
	}
	return *o.NumOfUes
}

// GetNumOfUesOk returns a tuple with the NumOfUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNumOfUesOk() (*NumberAverage, bool) {
	if o == nil || IsNil(o.NumOfUes) {
		return nil, false
	}
	return o.NumOfUes, true
}

// HasNumOfUes returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNumOfUes() bool {
	if o != nil && !IsNil(o.NumOfUes) {
		return true
	}

	return false
}

// SetNumOfUes gets a reference to the given NumberAverage and assigns it to the NumOfUes field.
func (o *NsiLoadLevelInfo) SetNumOfUes(v NumberAverage) {
	o.NumOfUes = &v
}

// GetNumOfPduSess returns the NumOfPduSess field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetNumOfPduSess() NumberAverage {
	if o == nil || IsNil(o.NumOfPduSess) {
		var ret NumberAverage
		return ret
	}
	return *o.NumOfPduSess
}

// GetNumOfPduSessOk returns a tuple with the NumOfPduSess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetNumOfPduSessOk() (*NumberAverage, bool) {
	if o == nil || IsNil(o.NumOfPduSess) {
		return nil, false
	}
	return o.NumOfPduSess, true
}

// HasNumOfPduSess returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasNumOfPduSess() bool {
	if o != nil && !IsNil(o.NumOfPduSess) {
		return true
	}

	return false
}

// SetNumOfPduSess gets a reference to the given NumberAverage and assigns it to the NumOfPduSess field.
func (o *NsiLoadLevelInfo) SetNumOfPduSess(v NumberAverage) {
	o.NumOfPduSess = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *NsiLoadLevelInfo) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsiLoadLevelInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *NsiLoadLevelInfo) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *NsiLoadLevelInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o NsiLoadLevelInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsiLoadLevelInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loadLevelInformation"] = o.LoadLevelInformation
	toSerialize["snssai"] = o.Snssai
	if !IsNil(o.NsiId) {
		toSerialize["nsiId"] = o.NsiId
	}
	if !IsNil(o.ResUsage) {
		toSerialize["resUsage"] = o.ResUsage
	}
	if !IsNil(o.NumOfExceedLoadLevelThr) {
		toSerialize["numOfExceedLoadLevelThr"] = o.NumOfExceedLoadLevelThr
	}
	if !IsNil(o.ExceedLoadLevelThrInd) {
		toSerialize["exceedLoadLevelThrInd"] = o.ExceedLoadLevelThrInd
	}
	if !IsNil(o.NetworkArea) {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if !IsNil(o.TimePeriod) {
		toSerialize["timePeriod"] = o.TimePeriod
	}
	if !IsNil(o.ResUsgThrCrossTimePeriod) {
		toSerialize["resUsgThrCrossTimePeriod"] = o.ResUsgThrCrossTimePeriod
	}
	if !IsNil(o.NumOfUes) {
		toSerialize["numOfUes"] = o.NumOfUes
	}
	if !IsNil(o.NumOfPduSess) {
		toSerialize["numOfPduSess"] = o.NumOfPduSess
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	return toSerialize, nil
}

type NullableNsiLoadLevelInfo struct {
	value *NsiLoadLevelInfo
	isSet bool
}

func (v NullableNsiLoadLevelInfo) Get() *NsiLoadLevelInfo {
	return v.value
}

func (v *NullableNsiLoadLevelInfo) Set(val *NsiLoadLevelInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsiLoadLevelInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsiLoadLevelInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsiLoadLevelInfo(val *NsiLoadLevelInfo) *NullableNsiLoadLevelInfo {
	return &NullableNsiLoadLevelInfo{value: val, isSet: true}
}

func (v NullableNsiLoadLevelInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsiLoadLevelInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


