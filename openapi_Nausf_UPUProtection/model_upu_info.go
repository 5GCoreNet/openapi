/*
Nausf_UPUProtection Service

AUSF UPU Protection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UPUProtection

import (
	"encoding/json"
)

// checks if the UpuInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpuInfo{}

// UpuInfo Contains the UE parameters update Information.
type UpuInfo struct {
	UpuDataList []UpuData `json:"upuDataList"`
	// Contains the \"UPU Header\" IE as specified in clause 9.11.3.53A of 3GPP TS 24.501  (octet 4), encoded as 2 hexadecimal characters.
	UpuHeader *string `json:"upuHeader,omitempty"`
	// Contains the indication of whether the acknowledgement from UE is needed.
	UpuAckInd bool `json:"upuAckInd"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UpuTransparentInfo *string `json:"upuTransparentInfo,omitempty"`
}

// NewUpuInfo instantiates a new UpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpuInfo(upuDataList []UpuData, upuAckInd bool) *UpuInfo {
	this := UpuInfo{}
	this.UpuDataList = upuDataList
	this.UpuAckInd = upuAckInd
	return &this
}

// NewUpuInfoWithDefaults instantiates a new UpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpuInfoWithDefaults() *UpuInfo {
	this := UpuInfo{}
	return &this
}

// GetUpuDataList returns the UpuDataList field value
func (o *UpuInfo) GetUpuDataList() []UpuData {
	if o == nil {
		var ret []UpuData
		return ret
	}

	return o.UpuDataList
}

// GetUpuDataListOk returns a tuple with the UpuDataList field value
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuDataListOk() ([]UpuData, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpuDataList, true
}

// SetUpuDataList sets field value
func (o *UpuInfo) SetUpuDataList(v []UpuData) {
	o.UpuDataList = v
}

// GetUpuHeader returns the UpuHeader field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuHeader() string {
	if o == nil || IsNil(o.UpuHeader) {
		var ret string
		return ret
	}
	return *o.UpuHeader
}

// GetUpuHeaderOk returns a tuple with the UpuHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.UpuHeader) {
		return nil, false
	}
	return o.UpuHeader, true
}

// HasUpuHeader returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuHeader() bool {
	if o != nil && !IsNil(o.UpuHeader) {
		return true
	}

	return false
}

// SetUpuHeader gets a reference to the given string and assigns it to the UpuHeader field.
func (o *UpuInfo) SetUpuHeader(v string) {
	o.UpuHeader = &v
}

// GetUpuAckInd returns the UpuAckInd field value
func (o *UpuInfo) GetUpuAckInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpuAckInd
}

// GetUpuAckIndOk returns a tuple with the UpuAckInd field value
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuAckIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpuAckInd, true
}

// SetUpuAckInd sets field value
func (o *UpuInfo) SetUpuAckInd(v bool) {
	o.UpuAckInd = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *UpuInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *UpuInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *UpuInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetUpuTransparentInfo returns the UpuTransparentInfo field value if set, zero value otherwise.
func (o *UpuInfo) GetUpuTransparentInfo() string {
	if o == nil || IsNil(o.UpuTransparentInfo) {
		var ret string
		return ret
	}
	return *o.UpuTransparentInfo
}

// GetUpuTransparentInfoOk returns a tuple with the UpuTransparentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuInfo) GetUpuTransparentInfoOk() (*string, bool) {
	if o == nil || IsNil(o.UpuTransparentInfo) {
		return nil, false
	}
	return o.UpuTransparentInfo, true
}

// HasUpuTransparentInfo returns a boolean if a field has been set.
func (o *UpuInfo) HasUpuTransparentInfo() bool {
	if o != nil && !IsNil(o.UpuTransparentInfo) {
		return true
	}

	return false
}

// SetUpuTransparentInfo gets a reference to the given string and assigns it to the UpuTransparentInfo field.
func (o *UpuInfo) SetUpuTransparentInfo(v string) {
	o.UpuTransparentInfo = &v
}

func (o UpuInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpuInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upuDataList"] = o.UpuDataList
	if !IsNil(o.UpuHeader) {
		toSerialize["upuHeader"] = o.UpuHeader
	}
	toSerialize["upuAckInd"] = o.UpuAckInd
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.UpuTransparentInfo) {
		toSerialize["upuTransparentInfo"] = o.UpuTransparentInfo
	}
	return toSerialize, nil
}

type NullableUpuInfo struct {
	value *UpuInfo
	isSet bool
}

func (v NullableUpuInfo) Get() *UpuInfo {
	return v.value
}

func (v *NullableUpuInfo) Set(val *UpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpuInfo(val *UpuInfo) *NullableUpuInfo {
	return &NullableUpuInfo{value: val, isSet: true}
}

func (v NullableUpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
