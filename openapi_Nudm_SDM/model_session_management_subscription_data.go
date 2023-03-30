/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the SessionManagementSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionManagementSubscriptionData{}

// SessionManagementSubscriptionData struct for SessionManagementSubscriptionData
type SessionManagementSubscriptionData struct {
	SingleNssai Snssai `json:"singleNssai"`
	// A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations
	DnnConfigurations *map[string]DnnConfiguration `json:"dnnConfigurations,omitempty"`
	InternalGroupIds []string `json:"internalGroupIds,omitempty"`
	// A map(list of key-value pairs) where GroupId serves as key of SharedDataId
	SharedVnGroupDataIds *map[string]string `json:"sharedVnGroupDataIds,omitempty"`
	SharedDnnConfigurationsId *string `json:"sharedDnnConfigurationsId,omitempty"`
	OdbPacketServices *OdbPacketServices `json:"odbPacketServices,omitempty"`
	TraceData NullableTraceData `json:"traceData,omitempty"`
	SharedTraceDataId *string `json:"sharedTraceDataId,omitempty"`
	// A map(list of key-value pairs) where Dnn serves as key of ExpectedUeBehaviourData
	ExpectedUeBehavioursList *map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`
	// A map(list of key-value pairs) where Dnn serves as key of SuggestedPacketNumDl
	SuggestedPacketNumDlList *map[string]SuggestedPacketNumDl `json:"suggestedPacketNumDlList,omitempty"`
	Var3gppChargingCharacteristics *string `json:"3gppChargingCharacteristics,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewSessionManagementSubscriptionData instantiates a new SessionManagementSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionManagementSubscriptionData(singleNssai Snssai) *SessionManagementSubscriptionData {
	this := SessionManagementSubscriptionData{}
	this.SingleNssai = singleNssai
	return &this
}

// NewSessionManagementSubscriptionDataWithDefaults instantiates a new SessionManagementSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionManagementSubscriptionDataWithDefaults() *SessionManagementSubscriptionData {
	this := SessionManagementSubscriptionData{}
	return &this
}

// GetSingleNssai returns the SingleNssai field value
func (o *SessionManagementSubscriptionData) GetSingleNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleNssai, true
}

// SetSingleNssai sets field value
func (o *SessionManagementSubscriptionData) SetSingleNssai(v Snssai) {
	o.SingleNssai = v
}

// GetDnnConfigurations returns the DnnConfigurations field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetDnnConfigurations() map[string]DnnConfiguration {
	if o == nil || IsNil(o.DnnConfigurations) {
		var ret map[string]DnnConfiguration
		return ret
	}
	return *o.DnnConfigurations
}

// GetDnnConfigurationsOk returns a tuple with the DnnConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetDnnConfigurationsOk() (*map[string]DnnConfiguration, bool) {
	if o == nil || IsNil(o.DnnConfigurations) {
		return nil, false
	}
	return o.DnnConfigurations, true
}

// HasDnnConfigurations returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasDnnConfigurations() bool {
	if o != nil && !IsNil(o.DnnConfigurations) {
		return true
	}

	return false
}

// SetDnnConfigurations gets a reference to the given map[string]DnnConfiguration and assigns it to the DnnConfigurations field.
func (o *SessionManagementSubscriptionData) SetDnnConfigurations(v map[string]DnnConfiguration) {
	o.DnnConfigurations = &v
}

// GetInternalGroupIds returns the InternalGroupIds field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetInternalGroupIds() []string {
	if o == nil || IsNil(o.InternalGroupIds) {
		var ret []string
		return ret
	}
	return o.InternalGroupIds
}

// GetInternalGroupIdsOk returns a tuple with the InternalGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetInternalGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalGroupIds) {
		return nil, false
	}
	return o.InternalGroupIds, true
}

// HasInternalGroupIds returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasInternalGroupIds() bool {
	if o != nil && !IsNil(o.InternalGroupIds) {
		return true
	}

	return false
}

// SetInternalGroupIds gets a reference to the given []string and assigns it to the InternalGroupIds field.
func (o *SessionManagementSubscriptionData) SetInternalGroupIds(v []string) {
	o.InternalGroupIds = v
}

// GetSharedVnGroupDataIds returns the SharedVnGroupDataIds field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetSharedVnGroupDataIds() map[string]string {
	if o == nil || IsNil(o.SharedVnGroupDataIds) {
		var ret map[string]string
		return ret
	}
	return *o.SharedVnGroupDataIds
}

// GetSharedVnGroupDataIdsOk returns a tuple with the SharedVnGroupDataIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSharedVnGroupDataIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.SharedVnGroupDataIds) {
		return nil, false
	}
	return o.SharedVnGroupDataIds, true
}

// HasSharedVnGroupDataIds returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasSharedVnGroupDataIds() bool {
	if o != nil && !IsNil(o.SharedVnGroupDataIds) {
		return true
	}

	return false
}

// SetSharedVnGroupDataIds gets a reference to the given map[string]string and assigns it to the SharedVnGroupDataIds field.
func (o *SessionManagementSubscriptionData) SetSharedVnGroupDataIds(v map[string]string) {
	o.SharedVnGroupDataIds = &v
}

// GetSharedDnnConfigurationsId returns the SharedDnnConfigurationsId field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetSharedDnnConfigurationsId() string {
	if o == nil || IsNil(o.SharedDnnConfigurationsId) {
		var ret string
		return ret
	}
	return *o.SharedDnnConfigurationsId
}

// GetSharedDnnConfigurationsIdOk returns a tuple with the SharedDnnConfigurationsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSharedDnnConfigurationsIdOk() (*string, bool) {
	if o == nil || IsNil(o.SharedDnnConfigurationsId) {
		return nil, false
	}
	return o.SharedDnnConfigurationsId, true
}

// HasSharedDnnConfigurationsId returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasSharedDnnConfigurationsId() bool {
	if o != nil && !IsNil(o.SharedDnnConfigurationsId) {
		return true
	}

	return false
}

// SetSharedDnnConfigurationsId gets a reference to the given string and assigns it to the SharedDnnConfigurationsId field.
func (o *SessionManagementSubscriptionData) SetSharedDnnConfigurationsId(v string) {
	o.SharedDnnConfigurationsId = &v
}

// GetOdbPacketServices returns the OdbPacketServices field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetOdbPacketServices() OdbPacketServices {
	if o == nil || IsNil(o.OdbPacketServices) {
		var ret OdbPacketServices
		return ret
	}
	return *o.OdbPacketServices
}

// GetOdbPacketServicesOk returns a tuple with the OdbPacketServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetOdbPacketServicesOk() (*OdbPacketServices, bool) {
	if o == nil || IsNil(o.OdbPacketServices) {
		return nil, false
	}
	return o.OdbPacketServices, true
}

// HasOdbPacketServices returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasOdbPacketServices() bool {
	if o != nil && !IsNil(o.OdbPacketServices) {
		return true
	}

	return false
}

// SetOdbPacketServices gets a reference to the given OdbPacketServices and assigns it to the OdbPacketServices field.
func (o *SessionManagementSubscriptionData) SetOdbPacketServices(v OdbPacketServices) {
	o.OdbPacketServices = &v
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionManagementSubscriptionData) GetTraceData() TraceData {
	if o == nil || IsNil(o.TraceData.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionManagementSubscriptionData) GetTraceDataOk() (*TraceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData and assigns it to the TraceData field.
func (o *SessionManagementSubscriptionData) SetTraceData(v TraceData) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *SessionManagementSubscriptionData) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *SessionManagementSubscriptionData) UnsetTraceData() {
	o.TraceData.Unset()
}

// GetSharedTraceDataId returns the SharedTraceDataId field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetSharedTraceDataId() string {
	if o == nil || IsNil(o.SharedTraceDataId) {
		var ret string
		return ret
	}
	return *o.SharedTraceDataId
}

// GetSharedTraceDataIdOk returns a tuple with the SharedTraceDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSharedTraceDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.SharedTraceDataId) {
		return nil, false
	}
	return o.SharedTraceDataId, true
}

// HasSharedTraceDataId returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasSharedTraceDataId() bool {
	if o != nil && !IsNil(o.SharedTraceDataId) {
		return true
	}

	return false
}

// SetSharedTraceDataId gets a reference to the given string and assigns it to the SharedTraceDataId field.
func (o *SessionManagementSubscriptionData) SetSharedTraceDataId(v string) {
	o.SharedTraceDataId = &v
}

// GetExpectedUeBehavioursList returns the ExpectedUeBehavioursList field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetExpectedUeBehavioursList() map[string]ExpectedUeBehaviourData {
	if o == nil || IsNil(o.ExpectedUeBehavioursList) {
		var ret map[string]ExpectedUeBehaviourData
		return ret
	}
	return *o.ExpectedUeBehavioursList
}

// GetExpectedUeBehavioursListOk returns a tuple with the ExpectedUeBehavioursList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetExpectedUeBehavioursListOk() (*map[string]ExpectedUeBehaviourData, bool) {
	if o == nil || IsNil(o.ExpectedUeBehavioursList) {
		return nil, false
	}
	return o.ExpectedUeBehavioursList, true
}

// HasExpectedUeBehavioursList returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasExpectedUeBehavioursList() bool {
	if o != nil && !IsNil(o.ExpectedUeBehavioursList) {
		return true
	}

	return false
}

// SetExpectedUeBehavioursList gets a reference to the given map[string]ExpectedUeBehaviourData and assigns it to the ExpectedUeBehavioursList field.
func (o *SessionManagementSubscriptionData) SetExpectedUeBehavioursList(v map[string]ExpectedUeBehaviourData) {
	o.ExpectedUeBehavioursList = &v
}

// GetSuggestedPacketNumDlList returns the SuggestedPacketNumDlList field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetSuggestedPacketNumDlList() map[string]SuggestedPacketNumDl {
	if o == nil || IsNil(o.SuggestedPacketNumDlList) {
		var ret map[string]SuggestedPacketNumDl
		return ret
	}
	return *o.SuggestedPacketNumDlList
}

// GetSuggestedPacketNumDlListOk returns a tuple with the SuggestedPacketNumDlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSuggestedPacketNumDlListOk() (*map[string]SuggestedPacketNumDl, bool) {
	if o == nil || IsNil(o.SuggestedPacketNumDlList) {
		return nil, false
	}
	return o.SuggestedPacketNumDlList, true
}

// HasSuggestedPacketNumDlList returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasSuggestedPacketNumDlList() bool {
	if o != nil && !IsNil(o.SuggestedPacketNumDlList) {
		return true
	}

	return false
}

// SetSuggestedPacketNumDlList gets a reference to the given map[string]SuggestedPacketNumDl and assigns it to the SuggestedPacketNumDlList field.
func (o *SessionManagementSubscriptionData) SetSuggestedPacketNumDlList(v map[string]SuggestedPacketNumDl) {
	o.SuggestedPacketNumDlList = &v
}

// GetVar3gppChargingCharacteristics returns the Var3gppChargingCharacteristics field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetVar3gppChargingCharacteristics() string {
	if o == nil || IsNil(o.Var3gppChargingCharacteristics) {
		var ret string
		return ret
	}
	return *o.Var3gppChargingCharacteristics
}

// GetVar3gppChargingCharacteristicsOk returns a tuple with the Var3gppChargingCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetVar3gppChargingCharacteristicsOk() (*string, bool) {
	if o == nil || IsNil(o.Var3gppChargingCharacteristics) {
		return nil, false
	}
	return o.Var3gppChargingCharacteristics, true
}

// HasVar3gppChargingCharacteristics returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasVar3gppChargingCharacteristics() bool {
	if o != nil && !IsNil(o.Var3gppChargingCharacteristics) {
		return true
	}

	return false
}

// SetVar3gppChargingCharacteristics gets a reference to the given string and assigns it to the Var3gppChargingCharacteristics field.
func (o *SessionManagementSubscriptionData) SetVar3gppChargingCharacteristics(v string) {
	o.Var3gppChargingCharacteristics = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SessionManagementSubscriptionData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionManagementSubscriptionData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SessionManagementSubscriptionData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SessionManagementSubscriptionData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SessionManagementSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionManagementSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["singleNssai"] = o.SingleNssai
	if !IsNil(o.DnnConfigurations) {
		toSerialize["dnnConfigurations"] = o.DnnConfigurations
	}
	if !IsNil(o.InternalGroupIds) {
		toSerialize["internalGroupIds"] = o.InternalGroupIds
	}
	if !IsNil(o.SharedVnGroupDataIds) {
		toSerialize["sharedVnGroupDataIds"] = o.SharedVnGroupDataIds
	}
	if !IsNil(o.SharedDnnConfigurationsId) {
		toSerialize["sharedDnnConfigurationsId"] = o.SharedDnnConfigurationsId
	}
	if !IsNil(o.OdbPacketServices) {
		toSerialize["odbPacketServices"] = o.OdbPacketServices
	}
	if o.TraceData.IsSet() {
		toSerialize["traceData"] = o.TraceData.Get()
	}
	if !IsNil(o.SharedTraceDataId) {
		toSerialize["sharedTraceDataId"] = o.SharedTraceDataId
	}
	if !IsNil(o.ExpectedUeBehavioursList) {
		toSerialize["expectedUeBehavioursList"] = o.ExpectedUeBehavioursList
	}
	if !IsNil(o.SuggestedPacketNumDlList) {
		toSerialize["suggestedPacketNumDlList"] = o.SuggestedPacketNumDlList
	}
	if !IsNil(o.Var3gppChargingCharacteristics) {
		toSerialize["3gppChargingCharacteristics"] = o.Var3gppChargingCharacteristics
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableSessionManagementSubscriptionData struct {
	value *SessionManagementSubscriptionData
	isSet bool
}

func (v NullableSessionManagementSubscriptionData) Get() *SessionManagementSubscriptionData {
	return v.value
}

func (v *NullableSessionManagementSubscriptionData) Set(val *SessionManagementSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionManagementSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionManagementSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionManagementSubscriptionData(val *SessionManagementSubscriptionData) *NullableSessionManagementSubscriptionData {
	return &NullableSessionManagementSubscriptionData{value: val, isSet: true}
}

func (v NullableSessionManagementSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionManagementSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


