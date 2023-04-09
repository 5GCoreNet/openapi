/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the UdrInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UdrInfo{}

// UdrInfo Information of an UDR NF Instance
type UdrInfo struct {
	// Identifier of a group of NFs.
	GroupId                        *string             `json:"groupId,omitempty"`
	SupiRanges                     []SupiRange         `json:"supiRanges,omitempty"`
	GpsiRanges                     []IdentityRange     `json:"gpsiRanges,omitempty"`
	ExternalGroupIdentifiersRanges []IdentityRange     `json:"externalGroupIdentifiersRanges,omitempty"`
	SupportedDataSets              []DataSetId         `json:"supportedDataSets,omitempty"`
	SharedDataIdRanges             []SharedDataIdRange `json:"sharedDataIdRanges,omitempty"`
}

// NewUdrInfo instantiates a new UdrInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdrInfo() *UdrInfo {
	this := UdrInfo{}
	return &this
}

// NewUdrInfoWithDefaults instantiates a new UdrInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdrInfoWithDefaults() *UdrInfo {
	this := UdrInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdrInfo) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdrInfo) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdrInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdrInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *UdrInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetExternalGroupIdentifiersRanges() []IdentityRange {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.ExternalGroupIdentifiersRanges
}

// GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetExternalGroupIdentifiersRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.ExternalGroupIdentifiersRanges) {
		return nil, false
	}
	return o.ExternalGroupIdentifiersRanges, true
}

// HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasExternalGroupIdentifiersRanges() bool {
	if o != nil && !IsNil(o.ExternalGroupIdentifiersRanges) {
		return true
	}

	return false
}

// SetExternalGroupIdentifiersRanges gets a reference to the given []IdentityRange and assigns it to the ExternalGroupIdentifiersRanges field.
func (o *UdrInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange) {
	o.ExternalGroupIdentifiersRanges = v
}

// GetSupportedDataSets returns the SupportedDataSets field value if set, zero value otherwise.
func (o *UdrInfo) GetSupportedDataSets() []DataSetId {
	if o == nil || IsNil(o.SupportedDataSets) {
		var ret []DataSetId
		return ret
	}
	return o.SupportedDataSets
}

// GetSupportedDataSetsOk returns a tuple with the SupportedDataSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetSupportedDataSetsOk() ([]DataSetId, bool) {
	if o == nil || IsNil(o.SupportedDataSets) {
		return nil, false
	}
	return o.SupportedDataSets, true
}

// HasSupportedDataSets returns a boolean if a field has been set.
func (o *UdrInfo) HasSupportedDataSets() bool {
	if o != nil && !IsNil(o.SupportedDataSets) {
		return true
	}

	return false
}

// SetSupportedDataSets gets a reference to the given []DataSetId and assigns it to the SupportedDataSets field.
func (o *UdrInfo) SetSupportedDataSets(v []DataSetId) {
	o.SupportedDataSets = v
}

// GetSharedDataIdRanges returns the SharedDataIdRanges field value if set, zero value otherwise.
func (o *UdrInfo) GetSharedDataIdRanges() []SharedDataIdRange {
	if o == nil || IsNil(o.SharedDataIdRanges) {
		var ret []SharedDataIdRange
		return ret
	}
	return o.SharedDataIdRanges
}

// GetSharedDataIdRangesOk returns a tuple with the SharedDataIdRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdrInfo) GetSharedDataIdRangesOk() ([]SharedDataIdRange, bool) {
	if o == nil || IsNil(o.SharedDataIdRanges) {
		return nil, false
	}
	return o.SharedDataIdRanges, true
}

// HasSharedDataIdRanges returns a boolean if a field has been set.
func (o *UdrInfo) HasSharedDataIdRanges() bool {
	if o != nil && !IsNil(o.SharedDataIdRanges) {
		return true
	}

	return false
}

// SetSharedDataIdRanges gets a reference to the given []SharedDataIdRange and assigns it to the SharedDataIdRanges field.
func (o *UdrInfo) SetSharedDataIdRanges(v []SharedDataIdRange) {
	o.SharedDataIdRanges = v
}

func (o UdrInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UdrInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SupportedDataSets) {
		toSerialize["supportedDataSets"] = o.SupportedDataSets
	}
	if !IsNil(o.SharedDataIdRanges) {
		toSerialize["sharedDataIdRanges"] = o.SharedDataIdRanges
	}
	return toSerialize, nil
}

type NullableUdrInfo struct {
	value *UdrInfo
	isSet bool
}

func (v NullableUdrInfo) Get() *UdrInfo {
	return v.value
}

func (v *NullableUdrInfo) Set(val *UdrInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUdrInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUdrInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdrInfo(val *UdrInfo) *NullableUdrInfo {
	return &NullableUdrInfo{value: val, isSet: true}
}

func (v NullableUdrInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdrInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
