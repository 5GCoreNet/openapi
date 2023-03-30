/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the ECSFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSFunctionSingleAllOfAttributes{}

// ECSFunctionSingleAllOfAttributes struct for ECSFunctionSingleAllOfAttributes
type ECSFunctionSingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	VnfParametersList []VnfParameter `json:"vnfParametersList,omitempty"`
	PeeParametersList []PeeParameter `json:"peeParametersList,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
	ECSAddress *string `json:"eCSAddress,omitempty"`
	ProviderIdentifier *string `json:"providerIdentifier,omitempty"`
	EdgeDataNetworkRef []string `json:"edgeDataNetworkRef,omitempty"`
	EESFuncitonRef []string `json:"eESFuncitonRef,omitempty"`
	SoftwareImageInfo *SoftwareImageInfo `json:"softwareImageInfo,omitempty"`
	TrackingAreaIdList []Tai `json:"trackingAreaIdList,omitempty"`
	GeographicalLocation *GeoLoc `json:"geographicalLocation,omitempty"`
	Mcc *string `json:"mcc,omitempty"`
}

// NewECSFunctionSingleAllOfAttributes instantiates a new ECSFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSFunctionSingleAllOfAttributes() *ECSFunctionSingleAllOfAttributes {
	this := ECSFunctionSingleAllOfAttributes{}
	return &this
}

// NewECSFunctionSingleAllOfAttributesWithDefaults instantiates a new ECSFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSFunctionSingleAllOfAttributesWithDefaults() *ECSFunctionSingleAllOfAttributes {
	this := ECSFunctionSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ECSFunctionSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetVnfParametersList returns the VnfParametersList field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetVnfParametersList() []VnfParameter {
	if o == nil || IsNil(o.VnfParametersList) {
		var ret []VnfParameter
		return ret
	}
	return o.VnfParametersList
}

// GetVnfParametersListOk returns a tuple with the VnfParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetVnfParametersListOk() ([]VnfParameter, bool) {
	if o == nil || IsNil(o.VnfParametersList) {
		return nil, false
	}
	return o.VnfParametersList, true
}

// HasVnfParametersList returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasVnfParametersList() bool {
	if o != nil && !IsNil(o.VnfParametersList) {
		return true
	}

	return false
}

// SetVnfParametersList gets a reference to the given []VnfParameter and assigns it to the VnfParametersList field.
func (o *ECSFunctionSingleAllOfAttributes) SetVnfParametersList(v []VnfParameter) {
	o.VnfParametersList = v
}

// GetPeeParametersList returns the PeeParametersList field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetPeeParametersList() []PeeParameter {
	if o == nil || IsNil(o.PeeParametersList) {
		var ret []PeeParameter
		return ret
	}
	return o.PeeParametersList
}

// GetPeeParametersListOk returns a tuple with the PeeParametersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetPeeParametersListOk() ([]PeeParameter, bool) {
	if o == nil || IsNil(o.PeeParametersList) {
		return nil, false
	}
	return o.PeeParametersList, true
}

// HasPeeParametersList returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasPeeParametersList() bool {
	if o != nil && !IsNil(o.PeeParametersList) {
		return true
	}

	return false
}

// SetPeeParametersList gets a reference to the given []PeeParameter and assigns it to the PeeParametersList field.
func (o *ECSFunctionSingleAllOfAttributes) SetPeeParametersList(v []PeeParameter) {
	o.PeeParametersList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ECSFunctionSingleAllOfAttributes) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ECSFunctionSingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ECSFunctionSingleAllOfAttributes) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

// GetECSAddress returns the ECSAddress field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetECSAddress() string {
	if o == nil || IsNil(o.ECSAddress) {
		var ret string
		return ret
	}
	return *o.ECSAddress
}

// GetECSAddressOk returns a tuple with the ECSAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetECSAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ECSAddress) {
		return nil, false
	}
	return o.ECSAddress, true
}

// HasECSAddress returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasECSAddress() bool {
	if o != nil && !IsNil(o.ECSAddress) {
		return true
	}

	return false
}

// SetECSAddress gets a reference to the given string and assigns it to the ECSAddress field.
func (o *ECSFunctionSingleAllOfAttributes) SetECSAddress(v string) {
	o.ECSAddress = &v
}

// GetProviderIdentifier returns the ProviderIdentifier field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetProviderIdentifier() string {
	if o == nil || IsNil(o.ProviderIdentifier) {
		var ret string
		return ret
	}
	return *o.ProviderIdentifier
}

// GetProviderIdentifierOk returns a tuple with the ProviderIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetProviderIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderIdentifier) {
		return nil, false
	}
	return o.ProviderIdentifier, true
}

// HasProviderIdentifier returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasProviderIdentifier() bool {
	if o != nil && !IsNil(o.ProviderIdentifier) {
		return true
	}

	return false
}

// SetProviderIdentifier gets a reference to the given string and assigns it to the ProviderIdentifier field.
func (o *ECSFunctionSingleAllOfAttributes) SetProviderIdentifier(v string) {
	o.ProviderIdentifier = &v
}

// GetEdgeDataNetworkRef returns the EdgeDataNetworkRef field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetEdgeDataNetworkRef() []string {
	if o == nil || IsNil(o.EdgeDataNetworkRef) {
		var ret []string
		return ret
	}
	return o.EdgeDataNetworkRef
}

// GetEdgeDataNetworkRefOk returns a tuple with the EdgeDataNetworkRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetEdgeDataNetworkRefOk() ([]string, bool) {
	if o == nil || IsNil(o.EdgeDataNetworkRef) {
		return nil, false
	}
	return o.EdgeDataNetworkRef, true
}

// HasEdgeDataNetworkRef returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasEdgeDataNetworkRef() bool {
	if o != nil && !IsNil(o.EdgeDataNetworkRef) {
		return true
	}

	return false
}

// SetEdgeDataNetworkRef gets a reference to the given []string and assigns it to the EdgeDataNetworkRef field.
func (o *ECSFunctionSingleAllOfAttributes) SetEdgeDataNetworkRef(v []string) {
	o.EdgeDataNetworkRef = v
}

// GetEESFuncitonRef returns the EESFuncitonRef field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetEESFuncitonRef() []string {
	if o == nil || IsNil(o.EESFuncitonRef) {
		var ret []string
		return ret
	}
	return o.EESFuncitonRef
}

// GetEESFuncitonRefOk returns a tuple with the EESFuncitonRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetEESFuncitonRefOk() ([]string, bool) {
	if o == nil || IsNil(o.EESFuncitonRef) {
		return nil, false
	}
	return o.EESFuncitonRef, true
}

// HasEESFuncitonRef returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasEESFuncitonRef() bool {
	if o != nil && !IsNil(o.EESFuncitonRef) {
		return true
	}

	return false
}

// SetEESFuncitonRef gets a reference to the given []string and assigns it to the EESFuncitonRef field.
func (o *ECSFunctionSingleAllOfAttributes) SetEESFuncitonRef(v []string) {
	o.EESFuncitonRef = v
}

// GetSoftwareImageInfo returns the SoftwareImageInfo field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetSoftwareImageInfo() SoftwareImageInfo {
	if o == nil || IsNil(o.SoftwareImageInfo) {
		var ret SoftwareImageInfo
		return ret
	}
	return *o.SoftwareImageInfo
}

// GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool) {
	if o == nil || IsNil(o.SoftwareImageInfo) {
		return nil, false
	}
	return o.SoftwareImageInfo, true
}

// HasSoftwareImageInfo returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasSoftwareImageInfo() bool {
	if o != nil && !IsNil(o.SoftwareImageInfo) {
		return true
	}

	return false
}

// SetSoftwareImageInfo gets a reference to the given SoftwareImageInfo and assigns it to the SoftwareImageInfo field.
func (o *ECSFunctionSingleAllOfAttributes) SetSoftwareImageInfo(v SoftwareImageInfo) {
	o.SoftwareImageInfo = &v
}

// GetTrackingAreaIdList returns the TrackingAreaIdList field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetTrackingAreaIdList() []Tai {
	if o == nil || IsNil(o.TrackingAreaIdList) {
		var ret []Tai
		return ret
	}
	return o.TrackingAreaIdList
}

// GetTrackingAreaIdListOk returns a tuple with the TrackingAreaIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetTrackingAreaIdListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TrackingAreaIdList) {
		return nil, false
	}
	return o.TrackingAreaIdList, true
}

// HasTrackingAreaIdList returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasTrackingAreaIdList() bool {
	if o != nil && !IsNil(o.TrackingAreaIdList) {
		return true
	}

	return false
}

// SetTrackingAreaIdList gets a reference to the given []Tai and assigns it to the TrackingAreaIdList field.
func (o *ECSFunctionSingleAllOfAttributes) SetTrackingAreaIdList(v []Tai) {
	o.TrackingAreaIdList = v
}

// GetGeographicalLocation returns the GeographicalLocation field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetGeographicalLocation() GeoLoc {
	if o == nil || IsNil(o.GeographicalLocation) {
		var ret GeoLoc
		return ret
	}
	return *o.GeographicalLocation
}

// GetGeographicalLocationOk returns a tuple with the GeographicalLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetGeographicalLocationOk() (*GeoLoc, bool) {
	if o == nil || IsNil(o.GeographicalLocation) {
		return nil, false
	}
	return o.GeographicalLocation, true
}

// HasGeographicalLocation returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasGeographicalLocation() bool {
	if o != nil && !IsNil(o.GeographicalLocation) {
		return true
	}

	return false
}

// SetGeographicalLocation gets a reference to the given GeoLoc and assigns it to the GeographicalLocation field.
func (o *ECSFunctionSingleAllOfAttributes) SetGeographicalLocation(v GeoLoc) {
	o.GeographicalLocation = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributes) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributes) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributes) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *ECSFunctionSingleAllOfAttributes) SetMcc(v string) {
	o.Mcc = &v
}

func (o ECSFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.VnfParametersList) {
		toSerialize["vnfParametersList"] = o.VnfParametersList
	}
	if !IsNil(o.PeeParametersList) {
		toSerialize["peeParametersList"] = o.PeeParametersList
	}
	if !IsNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !IsNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !IsNil(o.SupportedTraceMetrics) {
		toSerialize["supportedTraceMetrics"] = o.SupportedTraceMetrics
	}
	if !IsNil(o.ECSAddress) {
		toSerialize["eCSAddress"] = o.ECSAddress
	}
	if !IsNil(o.ProviderIdentifier) {
		toSerialize["providerIdentifier"] = o.ProviderIdentifier
	}
	if !IsNil(o.EdgeDataNetworkRef) {
		toSerialize["edgeDataNetworkRef"] = o.EdgeDataNetworkRef
	}
	if !IsNil(o.EESFuncitonRef) {
		toSerialize["eESFuncitonRef"] = o.EESFuncitonRef
	}
	if !IsNil(o.SoftwareImageInfo) {
		toSerialize["softwareImageInfo"] = o.SoftwareImageInfo
	}
	if !IsNil(o.TrackingAreaIdList) {
		toSerialize["trackingAreaIdList"] = o.TrackingAreaIdList
	}
	if !IsNil(o.GeographicalLocation) {
		toSerialize["geographicalLocation"] = o.GeographicalLocation
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	return toSerialize, nil
}

type NullableECSFunctionSingleAllOfAttributes struct {
	value *ECSFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableECSFunctionSingleAllOfAttributes) Get() *ECSFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableECSFunctionSingleAllOfAttributes) Set(val *ECSFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableECSFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableECSFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSFunctionSingleAllOfAttributes(val *ECSFunctionSingleAllOfAttributes) *NullableECSFunctionSingleAllOfAttributes {
	return &NullableECSFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableECSFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

