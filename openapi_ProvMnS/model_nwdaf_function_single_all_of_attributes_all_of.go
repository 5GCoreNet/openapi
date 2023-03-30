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

// checks if the NwdafFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafFunctionSingleAllOfAttributesAllOf{}

// NwdafFunctionSingleAllOfAttributesAllOf struct for NwdafFunctionSingleAllOfAttributesAllOf
type NwdafFunctionSingleAllOfAttributesAllOf struct {
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	SBIFqdn *string `json:"sBIFqdn,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList []CommModel `json:"commModelList,omitempty"`
	NetworkSliceInfoList []NetworkSliceInfo `json:"networkSliceInfoList,omitempty"`
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	NwdafInfo *NwdafInfo `json:"nwdafInfo,omitempty"`
}

// NewNwdafFunctionSingleAllOfAttributesAllOf instantiates a new NwdafFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafFunctionSingleAllOfAttributesAllOf() *NwdafFunctionSingleAllOfAttributesAllOf {
	this := NwdafFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewNwdafFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new NwdafFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafFunctionSingleAllOfAttributesAllOfWithDefaults() *NwdafFunctionSingleAllOfAttributesAllOf {
	this := NwdafFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

// GetNetworkSliceInfoList returns the NetworkSliceInfoList field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetNetworkSliceInfoList() []NetworkSliceInfo {
	if o == nil || IsNil(o.NetworkSliceInfoList) {
		var ret []NetworkSliceInfo
		return ret
	}
	return o.NetworkSliceInfoList
}

// GetNetworkSliceInfoListOk returns a tuple with the NetworkSliceInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetNetworkSliceInfoListOk() ([]NetworkSliceInfo, bool) {
	if o == nil || IsNil(o.NetworkSliceInfoList) {
		return nil, false
	}
	return o.NetworkSliceInfoList, true
}

// HasNetworkSliceInfoList returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasNetworkSliceInfoList() bool {
	if o != nil && !IsNil(o.NetworkSliceInfoList) {
		return true
	}

	return false
}

// SetNetworkSliceInfoList gets a reference to the given []NetworkSliceInfo and assigns it to the NetworkSliceInfoList field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetNetworkSliceInfoList(v []NetworkSliceInfo) {
	o.NetworkSliceInfoList = v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetNwdafInfo returns the NwdafInfo field value if set, zero value otherwise.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetNwdafInfo() NwdafInfo {
	if o == nil || IsNil(o.NwdafInfo) {
		var ret NwdafInfo
		return ret
	}
	return *o.NwdafInfo
}

// GetNwdafInfoOk returns a tuple with the NwdafInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) GetNwdafInfoOk() (*NwdafInfo, bool) {
	if o == nil || IsNil(o.NwdafInfo) {
		return nil, false
	}
	return o.NwdafInfo, true
}

// HasNwdafInfo returns a boolean if a field has been set.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) HasNwdafInfo() bool {
	if o != nil && !IsNil(o.NwdafInfo) {
		return true
	}

	return false
}

// SetNwdafInfo gets a reference to the given NwdafInfo and assigns it to the NwdafInfo field.
func (o *NwdafFunctionSingleAllOfAttributesAllOf) SetNwdafInfo(v NwdafInfo) {
	o.NwdafInfo = &v
}

func (o NwdafFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	if !IsNil(o.NetworkSliceInfoList) {
		toSerialize["networkSliceInfoList"] = o.NetworkSliceInfoList
	}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.NwdafInfo) {
		toSerialize["nwdafInfo"] = o.NwdafInfo
	}
	return toSerialize, nil
}

type NullableNwdafFunctionSingleAllOfAttributesAllOf struct {
	value *NwdafFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableNwdafFunctionSingleAllOfAttributesAllOf) Get() *NwdafFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableNwdafFunctionSingleAllOfAttributesAllOf) Set(val *NwdafFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafFunctionSingleAllOfAttributesAllOf(val *NwdafFunctionSingleAllOfAttributesAllOf) *NullableNwdafFunctionSingleAllOfAttributesAllOf {
	return &NullableNwdafFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableNwdafFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


