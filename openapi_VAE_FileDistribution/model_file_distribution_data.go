/*
VAE_FileDistribution

API for VAE File Distribution Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_FileDistribution

import (
	"encoding/json"
)

// checks if the FileDistributionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDistributionData{}

// FileDistributionData Represents an individual File Distribution resource for a V2X group ID.
type FileDistributionData struct {
	// Represents the group ID for which a V2X message is addressed.
	GroupId *string `json:"groupId,omitempty"`
	FileLists []FileList `json:"fileLists"`
	ServiceClass *string `json:"serviceClass,omitempty"`
	GeoArea GeographicArea `json:"geoArea"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBitrate string `json:"maxBitrate"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxDelay int32 `json:"maxDelay"`
	LocalMbmsInfo *LocalMbmsInfo `json:"localMbmsInfo,omitempty"`
	LocalMbmsActInd *bool `json:"localMbmsActInd,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewFileDistributionData instantiates a new FileDistributionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDistributionData(fileLists []FileList, geoArea GeographicArea, maxBitrate string, maxDelay int32) *FileDistributionData {
	this := FileDistributionData{}
	this.FileLists = fileLists
	this.GeoArea = geoArea
	this.MaxBitrate = maxBitrate
	this.MaxDelay = maxDelay
	return &this
}

// NewFileDistributionDataWithDefaults instantiates a new FileDistributionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDistributionDataWithDefaults() *FileDistributionData {
	this := FileDistributionData{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *FileDistributionData) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *FileDistributionData) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *FileDistributionData) SetGroupId(v string) {
	o.GroupId = &v
}

// GetFileLists returns the FileLists field value
func (o *FileDistributionData) GetFileLists() []FileList {
	if o == nil {
		var ret []FileList
		return ret
	}

	return o.FileLists
}

// GetFileListsOk returns a tuple with the FileLists field value
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetFileListsOk() ([]FileList, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileLists, true
}

// SetFileLists sets field value
func (o *FileDistributionData) SetFileLists(v []FileList) {
	o.FileLists = v
}

// GetServiceClass returns the ServiceClass field value if set, zero value otherwise.
func (o *FileDistributionData) GetServiceClass() string {
	if o == nil || isNil(o.ServiceClass) {
		var ret string
		return ret
	}
	return *o.ServiceClass
}

// GetServiceClassOk returns a tuple with the ServiceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetServiceClassOk() (*string, bool) {
	if o == nil || isNil(o.ServiceClass) {
		return nil, false
	}
	return o.ServiceClass, true
}

// HasServiceClass returns a boolean if a field has been set.
func (o *FileDistributionData) HasServiceClass() bool {
	if o != nil && !isNil(o.ServiceClass) {
		return true
	}

	return false
}

// SetServiceClass gets a reference to the given string and assigns it to the ServiceClass field.
func (o *FileDistributionData) SetServiceClass(v string) {
	o.ServiceClass = &v
}

// GetGeoArea returns the GeoArea field value
func (o *FileDistributionData) GetGeoArea() GeographicArea {
	if o == nil {
		var ret GeographicArea
		return ret
	}

	return o.GeoArea
}

// GetGeoAreaOk returns a tuple with the GeoArea field value
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetGeoAreaOk() (*GeographicArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeoArea, true
}

// SetGeoArea sets field value
func (o *FileDistributionData) SetGeoArea(v GeographicArea) {
	o.GeoArea = v
}

// GetMaxBitrate returns the MaxBitrate field value
func (o *FileDistributionData) GetMaxBitrate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxBitrate
}

// GetMaxBitrateOk returns a tuple with the MaxBitrate field value
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetMaxBitrateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBitrate, true
}

// SetMaxBitrate sets field value
func (o *FileDistributionData) SetMaxBitrate(v string) {
	o.MaxBitrate = v
}

// GetMaxDelay returns the MaxDelay field value
func (o *FileDistributionData) GetMaxDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxDelay
}

// GetMaxDelayOk returns a tuple with the MaxDelay field value
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetMaxDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxDelay, true
}

// SetMaxDelay sets field value
func (o *FileDistributionData) SetMaxDelay(v int32) {
	o.MaxDelay = v
}

// GetLocalMbmsInfo returns the LocalMbmsInfo field value if set, zero value otherwise.
func (o *FileDistributionData) GetLocalMbmsInfo() LocalMbmsInfo {
	if o == nil || isNil(o.LocalMbmsInfo) {
		var ret LocalMbmsInfo
		return ret
	}
	return *o.LocalMbmsInfo
}

// GetLocalMbmsInfoOk returns a tuple with the LocalMbmsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetLocalMbmsInfoOk() (*LocalMbmsInfo, bool) {
	if o == nil || isNil(o.LocalMbmsInfo) {
		return nil, false
	}
	return o.LocalMbmsInfo, true
}

// HasLocalMbmsInfo returns a boolean if a field has been set.
func (o *FileDistributionData) HasLocalMbmsInfo() bool {
	if o != nil && !isNil(o.LocalMbmsInfo) {
		return true
	}

	return false
}

// SetLocalMbmsInfo gets a reference to the given LocalMbmsInfo and assigns it to the LocalMbmsInfo field.
func (o *FileDistributionData) SetLocalMbmsInfo(v LocalMbmsInfo) {
	o.LocalMbmsInfo = &v
}

// GetLocalMbmsActInd returns the LocalMbmsActInd field value if set, zero value otherwise.
func (o *FileDistributionData) GetLocalMbmsActInd() bool {
	if o == nil || isNil(o.LocalMbmsActInd) {
		var ret bool
		return ret
	}
	return *o.LocalMbmsActInd
}

// GetLocalMbmsActIndOk returns a tuple with the LocalMbmsActInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetLocalMbmsActIndOk() (*bool, bool) {
	if o == nil || isNil(o.LocalMbmsActInd) {
		return nil, false
	}
	return o.LocalMbmsActInd, true
}

// HasLocalMbmsActInd returns a boolean if a field has been set.
func (o *FileDistributionData) HasLocalMbmsActInd() bool {
	if o != nil && !isNil(o.LocalMbmsActInd) {
		return true
	}

	return false
}

// SetLocalMbmsActInd gets a reference to the given bool and assigns it to the LocalMbmsActInd field.
func (o *FileDistributionData) SetLocalMbmsActInd(v bool) {
	o.LocalMbmsActInd = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *FileDistributionData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDistributionData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *FileDistributionData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *FileDistributionData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o FileDistributionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDistributionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	toSerialize["fileLists"] = o.FileLists
	if !isNil(o.ServiceClass) {
		toSerialize["serviceClass"] = o.ServiceClass
	}
	toSerialize["geoArea"] = o.GeoArea
	toSerialize["maxBitrate"] = o.MaxBitrate
	toSerialize["maxDelay"] = o.MaxDelay
	if !isNil(o.LocalMbmsInfo) {
		toSerialize["localMbmsInfo"] = o.LocalMbmsInfo
	}
	if !isNil(o.LocalMbmsActInd) {
		toSerialize["localMbmsActInd"] = o.LocalMbmsActInd
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableFileDistributionData struct {
	value *FileDistributionData
	isSet bool
}

func (v NullableFileDistributionData) Get() *FileDistributionData {
	return v.value
}

func (v *NullableFileDistributionData) Set(val *FileDistributionData) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDistributionData) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDistributionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDistributionData(val *FileDistributionData) *NullableFileDistributionData {
	return &NullableFileDistributionData{value: val, isSet: true}
}

func (v NullableFileDistributionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDistributionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


