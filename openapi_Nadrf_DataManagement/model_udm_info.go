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

// checks if the UdmInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdmInfo{}

// UdmInfo Information of an UDM NF Instance
type UdmInfo struct {
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`
	RoutingIndicators []string `json:"routingIndicators,omitempty"`
	InternalGroupIdentifiersRanges []InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty"`
	SuciInfos []SuciInfo `json:"suciInfos,omitempty"`
}

// NewUdmInfo instantiates a new UdmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdmInfo() *UdmInfo {
	this := UdmInfo{}
	return &this
}

// NewUdmInfoWithDefaults instantiates a new UdmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdmInfoWithDefaults() *UdmInfo {
	this := UdmInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdmInfo) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdmInfo) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdmInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdmInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *UdmInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *UdmInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetRoutingIndicators returns the RoutingIndicators field value if set, zero value otherwise.
func (o *UdmInfo) GetRoutingIndicators() []string {
	if o == nil || IsNil(o.RoutingIndicators) {
		var ret []string
		return ret
	}
	return o.RoutingIndicators
}

// GetRoutingIndicatorsOk returns a tuple with the RoutingIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetRoutingIndicatorsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoutingIndicators) {
		return nil, false
	}
	return o.RoutingIndicators, true
}

// HasRoutingIndicators returns a boolean if a field has been set.
func (o *UdmInfo) HasRoutingIndicators() bool {
	if o != nil && !IsNil(o.RoutingIndicators) {
		return true
	}

	return false
}

// SetRoutingIndicators gets a reference to the given []string and assigns it to the RoutingIndicators field.
func (o *UdmInfo) SetRoutingIndicators(v []string) {
	o.RoutingIndicators = v
}

// GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdmInfo) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		var ret []InternalGroupIdRange
		return ret
	}
	return o.InternalGroupIdentifiersRanges
}

// GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetInternalGroupIdentifiersRangesOk() ([]InternalGroupIdRange, bool) {
	if o == nil || IsNil(o.InternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.InternalGroupIdentifiersRanges, true
}

// HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdmInfo) HasInternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.InternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetInternalGroupIdentifiersRanges gets a reference to the given []InternalGroupIdRange and assigns it to the InternalGroupIdentifiersRanges field.
func (o *UdmInfo) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange) {
	o.InternalGroupIdentifiersRanges = v
}

// GetSuciInfos returns the SuciInfos field value if set, zero value otherwise.
func (o *UdmInfo) GetSuciInfos() []SuciInfo {
	if o == nil || IsNil(o.SuciInfos) {
		var ret []SuciInfo
		return ret
	}
	return o.SuciInfos
}

// GetSuciInfosOk returns a tuple with the SuciInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdmInfo) GetSuciInfosOk() ([]SuciInfo, bool) {
	if o == nil || IsNil(o.SuciInfos) {
		return nil, false
	}
	return o.SuciInfos, true
}

// HasSuciInfos returns a boolean if a field has been set.
func (o *UdmInfo) HasSuciInfos() bool {
	if o != nil && !IsNil(o.SuciInfos) {
		return true
	}

	return false
}

// SetSuciInfos gets a reference to the given []SuciInfo and assigns it to the SuciInfos field.
func (o *UdmInfo) SetSuciInfos(v []SuciInfo) {
	o.SuciInfos = v
}

func (o UdmInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdmInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableUdmInfo struct {
	value *UdmInfo
	isSet bool
}

func (v NullableUdmInfo) Get() *UdmInfo {
	return v.value
}

func (v *NullableUdmInfo) Set(val *UdmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUdmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUdmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdmInfo(val *UdmInfo) *NullableUdmInfo {
	return &NullableUdmInfo{value: val, isSet: true}
}

func (v NullableUdmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


