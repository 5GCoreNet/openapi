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

// checks if the ECSFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSFunctionSingleAllOfAttributesAllOf{}

// ECSFunctionSingleAllOfAttributesAllOf struct for ECSFunctionSingleAllOfAttributesAllOf
type ECSFunctionSingleAllOfAttributesAllOf struct {
	ECSAddress           *string            `json:"eCSAddress,omitempty"`
	ProviderIdentifier   *string            `json:"providerIdentifier,omitempty"`
	EdgeDataNetworkRef   []string           `json:"edgeDataNetworkRef,omitempty"`
	EESFuncitonRef       []string           `json:"eESFuncitonRef,omitempty"`
	SoftwareImageInfo    *SoftwareImageInfo `json:"softwareImageInfo,omitempty"`
	TrackingAreaIdList   []Tai              `json:"trackingAreaIdList,omitempty"`
	GeographicalLocation *GeoLoc            `json:"geographicalLocation,omitempty"`
	Mcc                  *string            `json:"mcc,omitempty"`
}

// NewECSFunctionSingleAllOfAttributesAllOf instantiates a new ECSFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSFunctionSingleAllOfAttributesAllOf() *ECSFunctionSingleAllOfAttributesAllOf {
	this := ECSFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewECSFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new ECSFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSFunctionSingleAllOfAttributesAllOfWithDefaults() *ECSFunctionSingleAllOfAttributesAllOf {
	this := ECSFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetECSAddress returns the ECSAddress field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetECSAddress() string {
	if o == nil || IsNil(o.ECSAddress) {
		var ret string
		return ret
	}
	return *o.ECSAddress
}

// GetECSAddressOk returns a tuple with the ECSAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetECSAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ECSAddress) {
		return nil, false
	}
	return o.ECSAddress, true
}

// HasECSAddress returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasECSAddress() bool {
	if o != nil && !IsNil(o.ECSAddress) {
		return true
	}

	return false
}

// SetECSAddress gets a reference to the given string and assigns it to the ECSAddress field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetECSAddress(v string) {
	o.ECSAddress = &v
}

// GetProviderIdentifier returns the ProviderIdentifier field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetProviderIdentifier() string {
	if o == nil || IsNil(o.ProviderIdentifier) {
		var ret string
		return ret
	}
	return *o.ProviderIdentifier
}

// GetProviderIdentifierOk returns a tuple with the ProviderIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetProviderIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderIdentifier) {
		return nil, false
	}
	return o.ProviderIdentifier, true
}

// HasProviderIdentifier returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasProviderIdentifier() bool {
	if o != nil && !IsNil(o.ProviderIdentifier) {
		return true
	}

	return false
}

// SetProviderIdentifier gets a reference to the given string and assigns it to the ProviderIdentifier field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetProviderIdentifier(v string) {
	o.ProviderIdentifier = &v
}

// GetEdgeDataNetworkRef returns the EdgeDataNetworkRef field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetEdgeDataNetworkRef() []string {
	if o == nil || IsNil(o.EdgeDataNetworkRef) {
		var ret []string
		return ret
	}
	return o.EdgeDataNetworkRef
}

// GetEdgeDataNetworkRefOk returns a tuple with the EdgeDataNetworkRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetEdgeDataNetworkRefOk() ([]string, bool) {
	if o == nil || IsNil(o.EdgeDataNetworkRef) {
		return nil, false
	}
	return o.EdgeDataNetworkRef, true
}

// HasEdgeDataNetworkRef returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasEdgeDataNetworkRef() bool {
	if o != nil && !IsNil(o.EdgeDataNetworkRef) {
		return true
	}

	return false
}

// SetEdgeDataNetworkRef gets a reference to the given []string and assigns it to the EdgeDataNetworkRef field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetEdgeDataNetworkRef(v []string) {
	o.EdgeDataNetworkRef = v
}

// GetEESFuncitonRef returns the EESFuncitonRef field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetEESFuncitonRef() []string {
	if o == nil || IsNil(o.EESFuncitonRef) {
		var ret []string
		return ret
	}
	return o.EESFuncitonRef
}

// GetEESFuncitonRefOk returns a tuple with the EESFuncitonRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetEESFuncitonRefOk() ([]string, bool) {
	if o == nil || IsNil(o.EESFuncitonRef) {
		return nil, false
	}
	return o.EESFuncitonRef, true
}

// HasEESFuncitonRef returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasEESFuncitonRef() bool {
	if o != nil && !IsNil(o.EESFuncitonRef) {
		return true
	}

	return false
}

// SetEESFuncitonRef gets a reference to the given []string and assigns it to the EESFuncitonRef field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetEESFuncitonRef(v []string) {
	o.EESFuncitonRef = v
}

// GetSoftwareImageInfo returns the SoftwareImageInfo field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetSoftwareImageInfo() SoftwareImageInfo {
	if o == nil || IsNil(o.SoftwareImageInfo) {
		var ret SoftwareImageInfo
		return ret
	}
	return *o.SoftwareImageInfo
}

// GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool) {
	if o == nil || IsNil(o.SoftwareImageInfo) {
		return nil, false
	}
	return o.SoftwareImageInfo, true
}

// HasSoftwareImageInfo returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasSoftwareImageInfo() bool {
	if o != nil && !IsNil(o.SoftwareImageInfo) {
		return true
	}

	return false
}

// SetSoftwareImageInfo gets a reference to the given SoftwareImageInfo and assigns it to the SoftwareImageInfo field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetSoftwareImageInfo(v SoftwareImageInfo) {
	o.SoftwareImageInfo = &v
}

// GetTrackingAreaIdList returns the TrackingAreaIdList field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetTrackingAreaIdList() []Tai {
	if o == nil || IsNil(o.TrackingAreaIdList) {
		var ret []Tai
		return ret
	}
	return o.TrackingAreaIdList
}

// GetTrackingAreaIdListOk returns a tuple with the TrackingAreaIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetTrackingAreaIdListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TrackingAreaIdList) {
		return nil, false
	}
	return o.TrackingAreaIdList, true
}

// HasTrackingAreaIdList returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasTrackingAreaIdList() bool {
	if o != nil && !IsNil(o.TrackingAreaIdList) {
		return true
	}

	return false
}

// SetTrackingAreaIdList gets a reference to the given []Tai and assigns it to the TrackingAreaIdList field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetTrackingAreaIdList(v []Tai) {
	o.TrackingAreaIdList = v
}

// GetGeographicalLocation returns the GeographicalLocation field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetGeographicalLocation() GeoLoc {
	if o == nil || IsNil(o.GeographicalLocation) {
		var ret GeoLoc
		return ret
	}
	return *o.GeographicalLocation
}

// GetGeographicalLocationOk returns a tuple with the GeographicalLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetGeographicalLocationOk() (*GeoLoc, bool) {
	if o == nil || IsNil(o.GeographicalLocation) {
		return nil, false
	}
	return o.GeographicalLocation, true
}

// HasGeographicalLocation returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasGeographicalLocation() bool {
	if o != nil && !IsNil(o.GeographicalLocation) {
		return true
	}

	return false
}

// SetGeographicalLocation gets a reference to the given GeoLoc and assigns it to the GeographicalLocation field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetGeographicalLocation(v GeoLoc) {
	o.GeographicalLocation = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *ECSFunctionSingleAllOfAttributesAllOf) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *ECSFunctionSingleAllOfAttributesAllOf) SetMcc(v string) {
	o.Mcc = &v
}

func (o ECSFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableECSFunctionSingleAllOfAttributesAllOf struct {
	value *ECSFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableECSFunctionSingleAllOfAttributesAllOf) Get() *ECSFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableECSFunctionSingleAllOfAttributesAllOf) Set(val *ECSFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableECSFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableECSFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSFunctionSingleAllOfAttributesAllOf(val *ECSFunctionSingleAllOfAttributesAllOf) *NullableECSFunctionSingleAllOfAttributesAllOf {
	return &NullableECSFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableECSFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
