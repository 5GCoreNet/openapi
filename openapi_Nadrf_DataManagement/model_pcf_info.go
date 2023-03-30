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

// checks if the PcfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfInfo{}

// PcfInfo Information of a PCF NF Instance
type PcfInfo struct {
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	DnnList []string `json:"dnnList,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
	// Fully Qualified Domain Name
	RxDiamHost *string `json:"rxDiamHost,omitempty"`
	// Fully Qualified Domain Name
	RxDiamRealm *string `json:"rxDiamRealm,omitempty"`
	V2xSupportInd *bool `json:"v2xSupportInd,omitempty"`
	ProseSupportInd *bool `json:"proseSupportInd,omitempty"`
	ProseCapability *ProSeCapability `json:"proseCapability,omitempty"`
	V2xCapability *V2xCapability `json:"v2xCapability,omitempty"`
}

// NewPcfInfo instantiates a new PcfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfInfo() *PcfInfo {
	this := PcfInfo{}
	var v2xSupportInd bool = false
	this.V2xSupportInd = &v2xSupportInd
	var proseSupportInd bool = false
	this.ProseSupportInd = &proseSupportInd
	return &this
}

// NewPcfInfoWithDefaults instantiates a new PcfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfInfoWithDefaults() *PcfInfo {
	this := PcfInfo{}
	var v2xSupportInd bool = false
	this.V2xSupportInd = &v2xSupportInd
	var proseSupportInd bool = false
	this.ProseSupportInd = &proseSupportInd
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PcfInfo) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PcfInfo) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PcfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *PcfInfo) GetDnnList() []string {
	if o == nil || IsNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetDnnListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnnList) {
		return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *PcfInfo) HasDnnList() bool {
	if o != nil && !IsNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *PcfInfo) SetDnnList(v []string) {
	o.DnnList = v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *PcfInfo) GetSupiRanges() []SupiRange {
	if o == nil || IsNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || IsNil(o.SupiRanges) {
		return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *PcfInfo) HasSupiRanges() bool {
	if o != nil && !IsNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *PcfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetGpsiRanges returns the GpsiRanges field value if set, zero value otherwise.
func (o *PcfInfo) GetGpsiRanges() []IdentityRange {
	if o == nil || IsNil(o.GpsiRanges) {
		var ret []IdentityRange
		return ret
	}
	return o.GpsiRanges
}

// GetGpsiRangesOk returns a tuple with the GpsiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetGpsiRangesOk() ([]IdentityRange, bool) {
	if o == nil || IsNil(o.GpsiRanges) {
		return nil, false
	}
	return o.GpsiRanges, true
}

// HasGpsiRanges returns a boolean if a field has been set.
func (o *PcfInfo) HasGpsiRanges() bool {
	if o != nil && !IsNil(o.GpsiRanges) {
		return true
	}

	return false
}

// SetGpsiRanges gets a reference to the given []IdentityRange and assigns it to the GpsiRanges field.
func (o *PcfInfo) SetGpsiRanges(v []IdentityRange) {
	o.GpsiRanges = v
}

// GetRxDiamHost returns the RxDiamHost field value if set, zero value otherwise.
func (o *PcfInfo) GetRxDiamHost() string {
	if o == nil || IsNil(o.RxDiamHost) {
		var ret string
		return ret
	}
	return *o.RxDiamHost
}

// GetRxDiamHostOk returns a tuple with the RxDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetRxDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamHost) {
		return nil, false
	}
	return o.RxDiamHost, true
}

// HasRxDiamHost returns a boolean if a field has been set.
func (o *PcfInfo) HasRxDiamHost() bool {
	if o != nil && !IsNil(o.RxDiamHost) {
		return true
	}

	return false
}

// SetRxDiamHost gets a reference to the given string and assigns it to the RxDiamHost field.
func (o *PcfInfo) SetRxDiamHost(v string) {
	o.RxDiamHost = &v
}

// GetRxDiamRealm returns the RxDiamRealm field value if set, zero value otherwise.
func (o *PcfInfo) GetRxDiamRealm() string {
	if o == nil || IsNil(o.RxDiamRealm) {
		var ret string
		return ret
	}
	return *o.RxDiamRealm
}

// GetRxDiamRealmOk returns a tuple with the RxDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetRxDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.RxDiamRealm) {
		return nil, false
	}
	return o.RxDiamRealm, true
}

// HasRxDiamRealm returns a boolean if a field has been set.
func (o *PcfInfo) HasRxDiamRealm() bool {
	if o != nil && !IsNil(o.RxDiamRealm) {
		return true
	}

	return false
}

// SetRxDiamRealm gets a reference to the given string and assigns it to the RxDiamRealm field.
func (o *PcfInfo) SetRxDiamRealm(v string) {
	o.RxDiamRealm = &v
}

// GetV2xSupportInd returns the V2xSupportInd field value if set, zero value otherwise.
func (o *PcfInfo) GetV2xSupportInd() bool {
	if o == nil || IsNil(o.V2xSupportInd) {
		var ret bool
		return ret
	}
	return *o.V2xSupportInd
}

// GetV2xSupportIndOk returns a tuple with the V2xSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetV2xSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.V2xSupportInd) {
		return nil, false
	}
	return o.V2xSupportInd, true
}

// HasV2xSupportInd returns a boolean if a field has been set.
func (o *PcfInfo) HasV2xSupportInd() bool {
	if o != nil && !IsNil(o.V2xSupportInd) {
		return true
	}

	return false
}

// SetV2xSupportInd gets a reference to the given bool and assigns it to the V2xSupportInd field.
func (o *PcfInfo) SetV2xSupportInd(v bool) {
	o.V2xSupportInd = &v
}

// GetProseSupportInd returns the ProseSupportInd field value if set, zero value otherwise.
func (o *PcfInfo) GetProseSupportInd() bool {
	if o == nil || IsNil(o.ProseSupportInd) {
		var ret bool
		return ret
	}
	return *o.ProseSupportInd
}

// GetProseSupportIndOk returns a tuple with the ProseSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetProseSupportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ProseSupportInd) {
		return nil, false
	}
	return o.ProseSupportInd, true
}

// HasProseSupportInd returns a boolean if a field has been set.
func (o *PcfInfo) HasProseSupportInd() bool {
	if o != nil && !IsNil(o.ProseSupportInd) {
		return true
	}

	return false
}

// SetProseSupportInd gets a reference to the given bool and assigns it to the ProseSupportInd field.
func (o *PcfInfo) SetProseSupportInd(v bool) {
	o.ProseSupportInd = &v
}

// GetProseCapability returns the ProseCapability field value if set, zero value otherwise.
func (o *PcfInfo) GetProseCapability() ProSeCapability {
	if o == nil || IsNil(o.ProseCapability) {
		var ret ProSeCapability
		return ret
	}
	return *o.ProseCapability
}

// GetProseCapabilityOk returns a tuple with the ProseCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetProseCapabilityOk() (*ProSeCapability, bool) {
	if o == nil || IsNil(o.ProseCapability) {
		return nil, false
	}
	return o.ProseCapability, true
}

// HasProseCapability returns a boolean if a field has been set.
func (o *PcfInfo) HasProseCapability() bool {
	if o != nil && !IsNil(o.ProseCapability) {
		return true
	}

	return false
}

// SetProseCapability gets a reference to the given ProSeCapability and assigns it to the ProseCapability field.
func (o *PcfInfo) SetProseCapability(v ProSeCapability) {
	o.ProseCapability = &v
}

// GetV2xCapability returns the V2xCapability field value if set, zero value otherwise.
func (o *PcfInfo) GetV2xCapability() V2xCapability {
	if o == nil || IsNil(o.V2xCapability) {
		var ret V2xCapability
		return ret
	}
	return *o.V2xCapability
}

// GetV2xCapabilityOk returns a tuple with the V2xCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfInfo) GetV2xCapabilityOk() (*V2xCapability, bool) {
	if o == nil || IsNil(o.V2xCapability) {
		return nil, false
	}
	return o.V2xCapability, true
}

// HasV2xCapability returns a boolean if a field has been set.
func (o *PcfInfo) HasV2xCapability() bool {
	if o != nil && !IsNil(o.V2xCapability) {
		return true
	}

	return false
}

// SetV2xCapability gets a reference to the given V2xCapability and assigns it to the V2xCapability field.
func (o *PcfInfo) SetV2xCapability(v V2xCapability) {
	o.V2xCapability = &v
}

func (o PcfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !IsNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !IsNil(o.GpsiRanges) {
		toSerialize["gpsiRanges"] = o.GpsiRanges
	}
	if !IsNil(o.RxDiamHost) {
		toSerialize["rxDiamHost"] = o.RxDiamHost
	}
	if !IsNil(o.RxDiamRealm) {
		toSerialize["rxDiamRealm"] = o.RxDiamRealm
	}
	if !IsNil(o.V2xSupportInd) {
		toSerialize["v2xSupportInd"] = o.V2xSupportInd
	}
	if !IsNil(o.ProseSupportInd) {
		toSerialize["proseSupportInd"] = o.ProseSupportInd
	}
	if !IsNil(o.ProseCapability) {
		toSerialize["proseCapability"] = o.ProseCapability
	}
	if !IsNil(o.V2xCapability) {
		toSerialize["v2xCapability"] = o.V2xCapability
	}
	return toSerialize, nil
}

type NullablePcfInfo struct {
	value *PcfInfo
	isSet bool
}

func (v NullablePcfInfo) Get() *PcfInfo {
	return v.value
}

func (v *NullablePcfInfo) Set(val *PcfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfInfo(val *PcfInfo) *NullablePcfInfo {
	return &NullablePcfInfo{value: val, isSet: true}
}

func (v NullablePcfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

