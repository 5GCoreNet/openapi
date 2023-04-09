/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the ImsProfileData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImsProfileData{}

// ImsProfileData User's IMS profile data
type ImsProfileData struct {
	ImsServiceProfiles       []ImsServiceProfile           `json:"imsServiceProfiles"`
	ChargingInfo             *ChargingInfo                 `json:"chargingInfo,omitempty"`
	ServiceLevelTraceInfo    *ServiceLevelTraceInformation `json:"serviceLevelTraceInfo,omitempty"`
	ServicePriorityLevelList []string                      `json:"servicePriorityLevelList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures    *string `json:"supportedFeatures,omitempty"`
	MaxAllowedSimulReg   *int32  `json:"maxAllowedSimulReg,omitempty"`
	ServicePriorityLevel *int32  `json:"servicePriorityLevel,omitempty"`
}

// NewImsProfileData instantiates a new ImsProfileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImsProfileData(imsServiceProfiles []ImsServiceProfile) *ImsProfileData {
	this := ImsProfileData{}
	this.ImsServiceProfiles = imsServiceProfiles
	return &this
}

// NewImsProfileDataWithDefaults instantiates a new ImsProfileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImsProfileDataWithDefaults() *ImsProfileData {
	this := ImsProfileData{}
	return &this
}

// GetImsServiceProfiles returns the ImsServiceProfiles field value
func (o *ImsProfileData) GetImsServiceProfiles() []ImsServiceProfile {
	if o == nil {
		var ret []ImsServiceProfile
		return ret
	}

	return o.ImsServiceProfiles
}

// GetImsServiceProfilesOk returns a tuple with the ImsServiceProfiles field value
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetImsServiceProfilesOk() ([]ImsServiceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImsServiceProfiles, true
}

// SetImsServiceProfiles sets field value
func (o *ImsProfileData) SetImsServiceProfiles(v []ImsServiceProfile) {
	o.ImsServiceProfiles = v
}

// GetChargingInfo returns the ChargingInfo field value if set, zero value otherwise.
func (o *ImsProfileData) GetChargingInfo() ChargingInfo {
	if o == nil || IsNil(o.ChargingInfo) {
		var ret ChargingInfo
		return ret
	}
	return *o.ChargingInfo
}

// GetChargingInfoOk returns a tuple with the ChargingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetChargingInfoOk() (*ChargingInfo, bool) {
	if o == nil || IsNil(o.ChargingInfo) {
		return nil, false
	}
	return o.ChargingInfo, true
}

// HasChargingInfo returns a boolean if a field has been set.
func (o *ImsProfileData) HasChargingInfo() bool {
	if o != nil && !IsNil(o.ChargingInfo) {
		return true
	}

	return false
}

// SetChargingInfo gets a reference to the given ChargingInfo and assigns it to the ChargingInfo field.
func (o *ImsProfileData) SetChargingInfo(v ChargingInfo) {
	o.ChargingInfo = &v
}

// GetServiceLevelTraceInfo returns the ServiceLevelTraceInfo field value if set, zero value otherwise.
func (o *ImsProfileData) GetServiceLevelTraceInfo() ServiceLevelTraceInformation {
	if o == nil || IsNil(o.ServiceLevelTraceInfo) {
		var ret ServiceLevelTraceInformation
		return ret
	}
	return *o.ServiceLevelTraceInfo
}

// GetServiceLevelTraceInfoOk returns a tuple with the ServiceLevelTraceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetServiceLevelTraceInfoOk() (*ServiceLevelTraceInformation, bool) {
	if o == nil || IsNil(o.ServiceLevelTraceInfo) {
		return nil, false
	}
	return o.ServiceLevelTraceInfo, true
}

// HasServiceLevelTraceInfo returns a boolean if a field has been set.
func (o *ImsProfileData) HasServiceLevelTraceInfo() bool {
	if o != nil && !IsNil(o.ServiceLevelTraceInfo) {
		return true
	}

	return false
}

// SetServiceLevelTraceInfo gets a reference to the given ServiceLevelTraceInformation and assigns it to the ServiceLevelTraceInfo field.
func (o *ImsProfileData) SetServiceLevelTraceInfo(v ServiceLevelTraceInformation) {
	o.ServiceLevelTraceInfo = &v
}

// GetServicePriorityLevelList returns the ServicePriorityLevelList field value if set, zero value otherwise.
func (o *ImsProfileData) GetServicePriorityLevelList() []string {
	if o == nil || IsNil(o.ServicePriorityLevelList) {
		var ret []string
		return ret
	}
	return o.ServicePriorityLevelList
}

// GetServicePriorityLevelListOk returns a tuple with the ServicePriorityLevelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetServicePriorityLevelListOk() ([]string, bool) {
	if o == nil || IsNil(o.ServicePriorityLevelList) {
		return nil, false
	}
	return o.ServicePriorityLevelList, true
}

// HasServicePriorityLevelList returns a boolean if a field has been set.
func (o *ImsProfileData) HasServicePriorityLevelList() bool {
	if o != nil && !IsNil(o.ServicePriorityLevelList) {
		return true
	}

	return false
}

// SetServicePriorityLevelList gets a reference to the given []string and assigns it to the ServicePriorityLevelList field.
func (o *ImsProfileData) SetServicePriorityLevelList(v []string) {
	o.ServicePriorityLevelList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ImsProfileData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ImsProfileData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ImsProfileData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetMaxAllowedSimulReg returns the MaxAllowedSimulReg field value if set, zero value otherwise.
func (o *ImsProfileData) GetMaxAllowedSimulReg() int32 {
	if o == nil || IsNil(o.MaxAllowedSimulReg) {
		var ret int32
		return ret
	}
	return *o.MaxAllowedSimulReg
}

// GetMaxAllowedSimulRegOk returns a tuple with the MaxAllowedSimulReg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetMaxAllowedSimulRegOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAllowedSimulReg) {
		return nil, false
	}
	return o.MaxAllowedSimulReg, true
}

// HasMaxAllowedSimulReg returns a boolean if a field has been set.
func (o *ImsProfileData) HasMaxAllowedSimulReg() bool {
	if o != nil && !IsNil(o.MaxAllowedSimulReg) {
		return true
	}

	return false
}

// SetMaxAllowedSimulReg gets a reference to the given int32 and assigns it to the MaxAllowedSimulReg field.
func (o *ImsProfileData) SetMaxAllowedSimulReg(v int32) {
	o.MaxAllowedSimulReg = &v
}

// GetServicePriorityLevel returns the ServicePriorityLevel field value if set, zero value otherwise.
func (o *ImsProfileData) GetServicePriorityLevel() int32 {
	if o == nil || IsNil(o.ServicePriorityLevel) {
		var ret int32
		return ret
	}
	return *o.ServicePriorityLevel
}

// GetServicePriorityLevelOk returns a tuple with the ServicePriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImsProfileData) GetServicePriorityLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ServicePriorityLevel) {
		return nil, false
	}
	return o.ServicePriorityLevel, true
}

// HasServicePriorityLevel returns a boolean if a field has been set.
func (o *ImsProfileData) HasServicePriorityLevel() bool {
	if o != nil && !IsNil(o.ServicePriorityLevel) {
		return true
	}

	return false
}

// SetServicePriorityLevel gets a reference to the given int32 and assigns it to the ServicePriorityLevel field.
func (o *ImsProfileData) SetServicePriorityLevel(v int32) {
	o.ServicePriorityLevel = &v
}

func (o ImsProfileData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImsProfileData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imsServiceProfiles"] = o.ImsServiceProfiles
	if !IsNil(o.ChargingInfo) {
		toSerialize["chargingInfo"] = o.ChargingInfo
	}
	if !IsNil(o.ServiceLevelTraceInfo) {
		toSerialize["serviceLevelTraceInfo"] = o.ServiceLevelTraceInfo
	}
	if !IsNil(o.ServicePriorityLevelList) {
		toSerialize["servicePriorityLevelList"] = o.ServicePriorityLevelList
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.MaxAllowedSimulReg) {
		toSerialize["maxAllowedSimulReg"] = o.MaxAllowedSimulReg
	}
	if !IsNil(o.ServicePriorityLevel) {
		toSerialize["servicePriorityLevel"] = o.ServicePriorityLevel
	}
	return toSerialize, nil
}

type NullableImsProfileData struct {
	value *ImsProfileData
	isSet bool
}

func (v NullableImsProfileData) Get() *ImsProfileData {
	return v.value
}

func (v *NullableImsProfileData) Set(val *ImsProfileData) {
	v.value = val
	v.isSet = true
}

func (v NullableImsProfileData) IsSet() bool {
	return v.isSet
}

func (v *NullableImsProfileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsProfileData(val *ImsProfileData) *NullableImsProfileData {
	return &NullableImsProfileData{value: val, isSet: true}
}

func (v NullableImsProfileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsProfileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
