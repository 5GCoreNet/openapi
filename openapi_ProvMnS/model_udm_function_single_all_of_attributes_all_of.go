/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the UdmFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdmFunctionSingleAllOfAttributesAllOf{}

// UdmFunctionSingleAllOfAttributesAllOf struct for UdmFunctionSingleAllOfAttributesAllOf
type UdmFunctionSingleAllOfAttributesAllOf struct {
	PLMNInfoList []PlmnInfo `json:"pLMNInfoList,omitempty"`
	SBIFqdn *string `json:"sBIFqdn,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList []CommModel `json:"commModelList,omitempty"`
	ECSAddrConfigInfo []string `json:"eCSAddrConfigInfo,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	RoutingIndicators []string `json:"routingIndicators,omitempty"`
	InternalGroupIdentifiersRanges []InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty"`
	SuciInfos []SuciInfo `json:"suciInfos,omitempty"`
}

// NewUdmFunctionSingleAllOfAttributesAllOf instantiates a new UdmFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdmFunctionSingleAllOfAttributesAllOf() *UdmFunctionSingleAllOfAttributesAllOf {
	this := UdmFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewUdmFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new UdmFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdmFunctionSingleAllOfAttributesAllOfWithDefaults() *UdmFunctionSingleAllOfAttributesAllOf {
	this := UdmFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPLMNInfoList returns the PLMNInfoList field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetPLMNInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PLMNInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PLMNInfoList
}

// GetPLMNInfoListOk returns a tuple with the PLMNInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetPLMNInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PLMNInfoList) {
		return nil, false
	}
	return o.PLMNInfoList, true
}

// HasPLMNInfoList returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasPLMNInfoList() bool {
	if o != nil && !IsNil(o.PLMNInfoList) {
		return true
	}

	return false
}

// SetPLMNInfoList gets a reference to the given []PlmnInfo and assigns it to the PLMNInfoList field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetPLMNInfoList(v []PlmnInfo) {
	o.PLMNInfoList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

// GetECSAddrConfigInfo returns the ECSAddrConfigInfo field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetECSAddrConfigInfo() []string {
	if o == nil || IsNil(o.ECSAddrConfigInfo) {
		var ret []string
		return ret
	}
	return o.ECSAddrConfigInfo
}

// GetECSAddrConfigInfoOk returns a tuple with the ECSAddrConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetECSAddrConfigInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.ECSAddrConfigInfo) {
		return nil, false
	}
	return o.ECSAddrConfigInfo, true
}

// HasECSAddrConfigInfo returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasECSAddrConfigInfo() bool {
	if o != nil && !IsNil(o.ECSAddrConfigInfo) {
		return true
	}

	return false
}

// SetECSAddrConfigInfo gets a reference to the given []string and assigns it to the ECSAddrConfigInfo field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetECSAddrConfigInfo(v []string) {
	o.ECSAddrConfigInfo = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetRoutingIndicators() []string {
	if o == nil || IsNil(o.RoutingIndicators) {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoutingIndicators) {
		return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasRoutingIndicators() bool {
	if o != nil && !IsNil(o.RoutingIndicators) {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

// GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		var ret []InternalGroupIdRange
		return ret
	}
	return o.InternalGroupIdentifiersRanges
}

// GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetInternalGroupIdentifiersRangesOk() ([]InternalGroupIdRange, bool) {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.InternalGroupIdentifiersRanges, true
}

// HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasInternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.InternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetInternalGroupIdentifiersRanges gets a reference to the given []InternalGroupIdRange and assigns it to the InternalGroupIdentifiersRanges field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange) {
	o.InternalGroupIdentifiersRanges = v
}

// GetSuciInfos returns the SuciInfos field value if set, zero value otherwise.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSuciInfos() []SuciInfo {
	if o == nil || IsNil(o.SuciInfos) {
		var ret []SuciInfo
		return ret
	}
	return o.SuciInfos
}

// GetSuciInfosOk returns a tuple with the SuciInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) GetSuciInfosOk() ([]SuciInfo, bool) {
	if o == nil || IsNil(o.SuciInfos) {
		return nil, false
	}
	return o.SuciInfos, true
}

// HasSuciInfos returns a boolean if a field has been set.
func (o *UdmFunctionSingleAllOfAttributesAllOf) HasSuciInfos() bool {
	if o != nil && !IsNil(o.SuciInfos) {
		return true
	}

	return false
}

// SetSuciInfos gets a reference to the given []SuciInfo and assigns it to the SuciInfos field.
func (o *UdmFunctionSingleAllOfAttributesAllOf) SetSuciInfos(v []SuciInfo) {
	o.SuciInfos = v
}

func (o UdmFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdmFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PLMNInfoList) {
		toSerialize["pLMNInfoList"] = o.PLMNInfoList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	if !IsNil(o.ECSAddrConfigInfo) {
		toSerialize["eCSAddrConfigInfo"] = o.ECSAddrConfigInfo
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.ExternalGroupIdentifiersRanges) {
		toSerialize["externalGroupIdentifiersRanges"] = o.ExternalGroupIdentifiersRanges
	}
	if !IsNil(o.RoutingIndicators) {
		toSerialize["routingIndicators"] = o.RoutingIndicators
	}
	if !IsNil(o.InternalGroupIdentifiersRanges) {
		toSerialize["internalGroupIdentifiersRanges"] = o.InternalGroupIdentifiersRanges
	}
	if !IsNil(o.SuciInfos) {
		toSerialize["suciInfos"] = o.SuciInfos
	}
	return toSerialize, nil
}

type NullableUdmFunctionSingleAllOfAttributesAllOf struct {
	value *UdmFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableUdmFunctionSingleAllOfAttributesAllOf) Get() *UdmFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableUdmFunctionSingleAllOfAttributesAllOf) Set(val *UdmFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUdmFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUdmFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdmFunctionSingleAllOfAttributesAllOf(val *UdmFunctionSingleAllOfAttributesAllOf) *NullableUdmFunctionSingleAllOfAttributesAllOf {
	return &NullableUdmFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableUdmFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdmFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

