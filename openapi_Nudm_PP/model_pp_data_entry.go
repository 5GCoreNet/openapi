/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
	"time"
)

// checks if the PpDataEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PpDataEntry{}

// PpDataEntry struct for PpDataEntry
type PpDataEntry struct {
	CommunicationCharacteristics NullableCommunicationCharacteristicsAF `json:"communicationCharacteristics,omitempty"`
	ReferenceId                  *int32                                 `json:"referenceId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures            *string                   `json:"supportedFeatures,omitempty"`
	EcsAddrConfigInfo            NullableEcsAddrConfigInfo `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo       `json:"additionalEcsAddrConfigInfos,omitempty"`
	EcRestriction                NullableEcRestriction     `json:"ecRestriction,omitempty"`
}

// NewPpDataEntry instantiates a new PpDataEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpDataEntry() *PpDataEntry {
	this := PpDataEntry{}
	return &this
}

// NewPpDataEntryWithDefaults instantiates a new PpDataEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpDataEntryWithDefaults() *PpDataEntry {
	this := PpDataEntry{}
	return &this
}

// GetCommunicationCharacteristics returns the CommunicationCharacteristics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PpDataEntry) GetCommunicationCharacteristics() CommunicationCharacteristicsAF {
	if o == nil || IsNil(o.CommunicationCharacteristics.Get()) {
		var ret CommunicationCharacteristicsAF
		return ret
	}
	return *o.CommunicationCharacteristics.Get()
}

// GetCommunicationCharacteristicsOk returns a tuple with the CommunicationCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PpDataEntry) GetCommunicationCharacteristicsOk() (*CommunicationCharacteristicsAF, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommunicationCharacteristics.Get(), o.CommunicationCharacteristics.IsSet()
}

// HasCommunicationCharacteristics returns a boolean if a field has been set.
func (o *PpDataEntry) HasCommunicationCharacteristics() bool {
	if o != nil && o.CommunicationCharacteristics.IsSet() {
		return true
	}

	return false
}

// SetCommunicationCharacteristics gets a reference to the given NullableCommunicationCharacteristicsAF and assigns it to the CommunicationCharacteristics field.
func (o *PpDataEntry) SetCommunicationCharacteristics(v CommunicationCharacteristicsAF) {
	o.CommunicationCharacteristics.Set(&v)
}

// SetCommunicationCharacteristicsNil sets the value for CommunicationCharacteristics to be an explicit nil
func (o *PpDataEntry) SetCommunicationCharacteristicsNil() {
	o.CommunicationCharacteristics.Set(nil)
}

// UnsetCommunicationCharacteristics ensures that no value is present for CommunicationCharacteristics, not even an explicit nil
func (o *PpDataEntry) UnsetCommunicationCharacteristics() {
	o.CommunicationCharacteristics.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PpDataEntry) GetReferenceId() int32 {
	if o == nil || IsNil(o.ReferenceId) {
		var ret int32
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntry) GetReferenceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PpDataEntry) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.
func (o *PpDataEntry) SetReferenceId(v int32) {
	o.ReferenceId = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *PpDataEntry) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntry) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *PpDataEntry) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *PpDataEntry) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *PpDataEntry) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntry) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *PpDataEntry) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *PpDataEntry) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PpDataEntry) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntry) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PpDataEntry) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *PpDataEntry) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetEcsAddrConfigInfo returns the EcsAddrConfigInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PpDataEntry) GetEcsAddrConfigInfo() EcsAddrConfigInfo {
	if o == nil || IsNil(o.EcsAddrConfigInfo.Get()) {
		var ret EcsAddrConfigInfo
		return ret
	}
	return *o.EcsAddrConfigInfo.Get()
}

// GetEcsAddrConfigInfoOk returns a tuple with the EcsAddrConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PpDataEntry) GetEcsAddrConfigInfoOk() (*EcsAddrConfigInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.EcsAddrConfigInfo.Get(), o.EcsAddrConfigInfo.IsSet()
}

// HasEcsAddrConfigInfo returns a boolean if a field has been set.
func (o *PpDataEntry) HasEcsAddrConfigInfo() bool {
	if o != nil && o.EcsAddrConfigInfo.IsSet() {
		return true
	}

	return false
}

// SetEcsAddrConfigInfo gets a reference to the given NullableEcsAddrConfigInfo and assigns it to the EcsAddrConfigInfo field.
func (o *PpDataEntry) SetEcsAddrConfigInfo(v EcsAddrConfigInfo) {
	o.EcsAddrConfigInfo.Set(&v)
}

// SetEcsAddrConfigInfoNil sets the value for EcsAddrConfigInfo to be an explicit nil
func (o *PpDataEntry) SetEcsAddrConfigInfoNil() {
	o.EcsAddrConfigInfo.Set(nil)
}

// UnsetEcsAddrConfigInfo ensures that no value is present for EcsAddrConfigInfo, not even an explicit nil
func (o *PpDataEntry) UnsetEcsAddrConfigInfo() {
	o.EcsAddrConfigInfo.Unset()
}

// GetAdditionalEcsAddrConfigInfos returns the AdditionalEcsAddrConfigInfos field value if set, zero value otherwise.
func (o *PpDataEntry) GetAdditionalEcsAddrConfigInfos() []EcsAddrConfigInfo {
	if o == nil || IsNil(o.AdditionalEcsAddrConfigInfos) {
		var ret []EcsAddrConfigInfo
		return ret
	}
	return o.AdditionalEcsAddrConfigInfos
}

// GetAdditionalEcsAddrConfigInfosOk returns a tuple with the AdditionalEcsAddrConfigInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpDataEntry) GetAdditionalEcsAddrConfigInfosOk() ([]EcsAddrConfigInfo, bool) {
	if o == nil || IsNil(o.AdditionalEcsAddrConfigInfos) {
		return nil, false
	}
	return o.AdditionalEcsAddrConfigInfos, true
}

// HasAdditionalEcsAddrConfigInfos returns a boolean if a field has been set.
func (o *PpDataEntry) HasAdditionalEcsAddrConfigInfos() bool {
	if o != nil && !IsNil(o.AdditionalEcsAddrConfigInfos) {
		return true
	}

	return false
}

// SetAdditionalEcsAddrConfigInfos gets a reference to the given []EcsAddrConfigInfo and assigns it to the AdditionalEcsAddrConfigInfos field.
func (o *PpDataEntry) SetAdditionalEcsAddrConfigInfos(v []EcsAddrConfigInfo) {
	o.AdditionalEcsAddrConfigInfos = v
}

// GetEcRestriction returns the EcRestriction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PpDataEntry) GetEcRestriction() EcRestriction {
	if o == nil || IsNil(o.EcRestriction.Get()) {
		var ret EcRestriction
		return ret
	}
	return *o.EcRestriction.Get()
}

// GetEcRestrictionOk returns a tuple with the EcRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PpDataEntry) GetEcRestrictionOk() (*EcRestriction, bool) {
	if o == nil {
		return nil, false
	}
	return o.EcRestriction.Get(), o.EcRestriction.IsSet()
}

// HasEcRestriction returns a boolean if a field has been set.
func (o *PpDataEntry) HasEcRestriction() bool {
	if o != nil && o.EcRestriction.IsSet() {
		return true
	}

	return false
}

// SetEcRestriction gets a reference to the given NullableEcRestriction and assigns it to the EcRestriction field.
func (o *PpDataEntry) SetEcRestriction(v EcRestriction) {
	o.EcRestriction.Set(&v)
}

// SetEcRestrictionNil sets the value for EcRestriction to be an explicit nil
func (o *PpDataEntry) SetEcRestrictionNil() {
	o.EcRestriction.Set(nil)
}

// UnsetEcRestriction ensures that no value is present for EcRestriction, not even an explicit nil
func (o *PpDataEntry) UnsetEcRestriction() {
	o.EcRestriction.Unset()
}

func (o PpDataEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PpDataEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CommunicationCharacteristics.IsSet() {
		toSerialize["communicationCharacteristics"] = o.CommunicationCharacteristics.Get()
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if o.EcsAddrConfigInfo.IsSet() {
		toSerialize["ecsAddrConfigInfo"] = o.EcsAddrConfigInfo.Get()
	}
	if !IsNil(o.AdditionalEcsAddrConfigInfos) {
		toSerialize["additionalEcsAddrConfigInfos"] = o.AdditionalEcsAddrConfigInfos
	}
	if o.EcRestriction.IsSet() {
		toSerialize["ecRestriction"] = o.EcRestriction.Get()
	}
	return toSerialize, nil
}

type NullablePpDataEntry struct {
	value *PpDataEntry
	isSet bool
}

func (v NullablePpDataEntry) Get() *PpDataEntry {
	return v.value
}

func (v *NullablePpDataEntry) Set(val *PpDataEntry) {
	v.value = val
	v.isSet = true
}

func (v NullablePpDataEntry) IsSet() bool {
	return v.isSet
}

func (v *NullablePpDataEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpDataEntry(val *PpDataEntry) *NullablePpDataEntry {
	return &NullablePpDataEntry{value: val, isSet: true}
}

func (v NullablePpDataEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpDataEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
