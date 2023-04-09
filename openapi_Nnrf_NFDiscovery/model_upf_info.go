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

// checks if the UpfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfInfo{}

// UpfInfo Information of an UPF NF Instance
type UpfInfo struct {
	SNssaiUpfInfoList     []SnssaiUpfInfoItem    `json:"sNssaiUpfInfoList"`
	SmfServingArea        []string               `json:"smfServingArea,omitempty"`
	InterfaceUpfInfoList  []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty"`
	IwkEpsInd             *bool                  `json:"iwkEpsInd,omitempty"`
	SxaInd                *bool                  `json:"sxaInd,omitempty"`
	PduSessionTypes       []PduSessionType       `json:"pduSessionTypes,omitempty"`
	AtsssCapability       *AtsssCapability       `json:"atsssCapability,omitempty"`
	UeIpAddrInd           *bool                  `json:"ueIpAddrInd,omitempty"`
	TaiList               []Tai                  `json:"taiList,omitempty"`
	TaiRangeList          []TaiRange             `json:"taiRangeList,omitempty"`
	WAgfInfo              *WAgfInfo              `json:"wAgfInfo,omitempty"`
	TngfInfo              *TngfInfo              `json:"tngfInfo,omitempty"`
	TwifInfo              *TwifInfo              `json:"twifInfo,omitempty"`
	Priority              *int32                 `json:"priority,omitempty"`
	RedundantGtpu         *bool                  `json:"redundantGtpu,omitempty"`
	Ipups                 *bool                  `json:"ipups,omitempty"`
	DataForwarding        *bool                  `json:"dataForwarding,omitempty"`
	SupportedPfcpFeatures *string                `json:"supportedPfcpFeatures,omitempty"`
}

// NewUpfInfo instantiates a new UpfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfInfo(sNssaiUpfInfoList []SnssaiUpfInfoItem) *UpfInfo {
	this := UpfInfo{}
	this.SNssaiUpfInfoList = sNssaiUpfInfoList
	var iwkEpsInd bool = false
	this.IwkEpsInd = &iwkEpsInd
	var ueIpAddrInd bool = false
	this.UeIpAddrInd = &ueIpAddrInd
	var redundantGtpu bool = false
	this.RedundantGtpu = &redundantGtpu
	var ipups bool = false
	this.Ipups = &ipups
	var dataForwarding bool = false
	this.DataForwarding = &dataForwarding
	return &this
}

// NewUpfInfoWithDefaults instantiates a new UpfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfInfoWithDefaults() *UpfInfo {
	this := UpfInfo{}
	var iwkEpsInd bool = false
	this.IwkEpsInd = &iwkEpsInd
	var ueIpAddrInd bool = false
	this.UeIpAddrInd = &ueIpAddrInd
	var redundantGtpu bool = false
	this.RedundantGtpu = &redundantGtpu
	var ipups bool = false
	this.Ipups = &ipups
	var dataForwarding bool = false
	this.DataForwarding = &dataForwarding
	return &this
}

// GetSNssaiUpfInfoList returns the SNssaiUpfInfoList field value
func (o *UpfInfo) GetSNssaiUpfInfoList() []SnssaiUpfInfoItem {
	if o == nil {
		var ret []SnssaiUpfInfoItem
		return ret
	}

	return o.SNssaiUpfInfoList
}

// GetSNssaiUpfInfoListOk returns a tuple with the SNssaiUpfInfoList field value
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSNssaiUpfInfoListOk() ([]SnssaiUpfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.SNssaiUpfInfoList, true
}

// SetSNssaiUpfInfoList sets field value
func (o *UpfInfo) SetSNssaiUpfInfoList(v []SnssaiUpfInfoItem) {
	o.SNssaiUpfInfoList = v
}

// GetSmfServingArea returns the SmfServingArea field value if set, zero value otherwise.
func (o *UpfInfo) GetSmfServingArea() []string {
	if o == nil || IsNil(o.SmfServingArea) {
		var ret []string
		return ret
	}
	return o.SmfServingArea
}

// GetSmfServingAreaOk returns a tuple with the SmfServingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSmfServingAreaOk() ([]string, bool) {
	if o == nil || IsNil(o.SmfServingArea) {
		return nil, false
	}
	return o.SmfServingArea, true
}

// HasSmfServingArea returns a boolean if a field has been set.
func (o *UpfInfo) HasSmfServingArea() bool {
	if o != nil && !IsNil(o.SmfServingArea) {
		return true
	}

	return false
}

// SetSmfServingArea gets a reference to the given []string and assigns it to the SmfServingArea field.
func (o *UpfInfo) SetSmfServingArea(v []string) {
	o.SmfServingArea = v
}

// GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field value if set, zero value otherwise.
func (o *UpfInfo) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem {
	if o == nil || IsNil(o.InterfaceUpfInfoList) {
		var ret []InterfaceUpfInfoItem
		return ret
	}
	return o.InterfaceUpfInfoList
}

// GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetInterfaceUpfInfoListOk() ([]InterfaceUpfInfoItem, bool) {
	if o == nil || IsNil(o.InterfaceUpfInfoList) {
		return nil, false
	}
	return o.InterfaceUpfInfoList, true
}

// HasInterfaceUpfInfoList returns a boolean if a field has been set.
func (o *UpfInfo) HasInterfaceUpfInfoList() bool {
	if o != nil && !IsNil(o.InterfaceUpfInfoList) {
		return true
	}

	return false
}

// SetInterfaceUpfInfoList gets a reference to the given []InterfaceUpfInfoItem and assigns it to the InterfaceUpfInfoList field.
func (o *UpfInfo) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem) {
	o.InterfaceUpfInfoList = v
}

// GetIwkEpsInd returns the IwkEpsInd field value if set, zero value otherwise.
func (o *UpfInfo) GetIwkEpsInd() bool {
	if o == nil || IsNil(o.IwkEpsInd) {
		var ret bool
		return ret
	}
	return *o.IwkEpsInd
}

// GetIwkEpsIndOk returns a tuple with the IwkEpsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetIwkEpsIndOk() (*bool, bool) {
	if o == nil || IsNil(o.IwkEpsInd) {
		return nil, false
	}
	return o.IwkEpsInd, true
}

// HasIwkEpsInd returns a boolean if a field has been set.
func (o *UpfInfo) HasIwkEpsInd() bool {
	if o != nil && !IsNil(o.IwkEpsInd) {
		return true
	}

	return false
}

// SetIwkEpsInd gets a reference to the given bool and assigns it to the IwkEpsInd field.
func (o *UpfInfo) SetIwkEpsInd(v bool) {
	o.IwkEpsInd = &v
}

// GetSxaInd returns the SxaInd field value if set, zero value otherwise.
func (o *UpfInfo) GetSxaInd() bool {
	if o == nil || IsNil(o.SxaInd) {
		var ret bool
		return ret
	}
	return *o.SxaInd
}

// GetSxaIndOk returns a tuple with the SxaInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSxaIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SxaInd) {
		return nil, false
	}
	return o.SxaInd, true
}

// HasSxaInd returns a boolean if a field has been set.
func (o *UpfInfo) HasSxaInd() bool {
	if o != nil && !IsNil(o.SxaInd) {
		return true
	}

	return false
}

// SetSxaInd gets a reference to the given bool and assigns it to the SxaInd field.
func (o *UpfInfo) SetSxaInd(v bool) {
	o.SxaInd = &v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *UpfInfo) GetPduSessionTypes() []PduSessionType {
	if o == nil || IsNil(o.PduSessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetPduSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || IsNil(o.PduSessionTypes) {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *UpfInfo) HasPduSessionTypes() bool {
	if o != nil && !IsNil(o.PduSessionTypes) {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given []PduSessionType and assigns it to the PduSessionTypes field.
func (o *UpfInfo) SetPduSessionTypes(v []PduSessionType) {
	o.PduSessionTypes = v
}

// GetAtsssCapability returns the AtsssCapability field value if set, zero value otherwise.
func (o *UpfInfo) GetAtsssCapability() AtsssCapability {
	if o == nil || IsNil(o.AtsssCapability) {
		var ret AtsssCapability
		return ret
	}
	return *o.AtsssCapability
}

// GetAtsssCapabilityOk returns a tuple with the AtsssCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetAtsssCapabilityOk() (*AtsssCapability, bool) {
	if o == nil || IsNil(o.AtsssCapability) {
		return nil, false
	}
	return o.AtsssCapability, true
}

// HasAtsssCapability returns a boolean if a field has been set.
func (o *UpfInfo) HasAtsssCapability() bool {
	if o != nil && !IsNil(o.AtsssCapability) {
		return true
	}

	return false
}

// SetAtsssCapability gets a reference to the given AtsssCapability and assigns it to the AtsssCapability field.
func (o *UpfInfo) SetAtsssCapability(v AtsssCapability) {
	o.AtsssCapability = &v
}

// GetUeIpAddrInd returns the UeIpAddrInd field value if set, zero value otherwise.
func (o *UpfInfo) GetUeIpAddrInd() bool {
	if o == nil || IsNil(o.UeIpAddrInd) {
		var ret bool
		return ret
	}
	return *o.UeIpAddrInd
}

// GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetUeIpAddrIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UeIpAddrInd) {
		return nil, false
	}
	return o.UeIpAddrInd, true
}

// HasUeIpAddrInd returns a boolean if a field has been set.
func (o *UpfInfo) HasUeIpAddrInd() bool {
	if o != nil && !IsNil(o.UeIpAddrInd) {
		return true
	}

	return false
}

// SetUeIpAddrInd gets a reference to the given bool and assigns it to the UeIpAddrInd field.
func (o *UpfInfo) SetUeIpAddrInd(v bool) {
	o.UeIpAddrInd = &v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *UpfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *UpfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *UpfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *UpfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *UpfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *UpfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetWAgfInfo returns the WAgfInfo field value if set, zero value otherwise.
func (o *UpfInfo) GetWAgfInfo() WAgfInfo {
	if o == nil || IsNil(o.WAgfInfo) {
		var ret WAgfInfo
		return ret
	}
	return *o.WAgfInfo
}

// GetWAgfInfoOk returns a tuple with the WAgfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetWAgfInfoOk() (*WAgfInfo, bool) {
	if o == nil || IsNil(o.WAgfInfo) {
		return nil, false
	}
	return o.WAgfInfo, true
}

// HasWAgfInfo returns a boolean if a field has been set.
func (o *UpfInfo) HasWAgfInfo() bool {
	if o != nil && !IsNil(o.WAgfInfo) {
		return true
	}

	return false
}

// SetWAgfInfo gets a reference to the given WAgfInfo and assigns it to the WAgfInfo field.
func (o *UpfInfo) SetWAgfInfo(v WAgfInfo) {
	o.WAgfInfo = &v
}

// GetTngfInfo returns the TngfInfo field value if set, zero value otherwise.
func (o *UpfInfo) GetTngfInfo() TngfInfo {
	if o == nil || IsNil(o.TngfInfo) {
		var ret TngfInfo
		return ret
	}
	return *o.TngfInfo
}

// GetTngfInfoOk returns a tuple with the TngfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetTngfInfoOk() (*TngfInfo, bool) {
	if o == nil || IsNil(o.TngfInfo) {
		return nil, false
	}
	return o.TngfInfo, true
}

// HasTngfInfo returns a boolean if a field has been set.
func (o *UpfInfo) HasTngfInfo() bool {
	if o != nil && !IsNil(o.TngfInfo) {
		return true
	}

	return false
}

// SetTngfInfo gets a reference to the given TngfInfo and assigns it to the TngfInfo field.
func (o *UpfInfo) SetTngfInfo(v TngfInfo) {
	o.TngfInfo = &v
}

// GetTwifInfo returns the TwifInfo field value if set, zero value otherwise.
func (o *UpfInfo) GetTwifInfo() TwifInfo {
	if o == nil || IsNil(o.TwifInfo) {
		var ret TwifInfo
		return ret
	}
	return *o.TwifInfo
}

// GetTwifInfoOk returns a tuple with the TwifInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetTwifInfoOk() (*TwifInfo, bool) {
	if o == nil || IsNil(o.TwifInfo) {
		return nil, false
	}
	return o.TwifInfo, true
}

// HasTwifInfo returns a boolean if a field has been set.
func (o *UpfInfo) HasTwifInfo() bool {
	if o != nil && !IsNil(o.TwifInfo) {
		return true
	}

	return false
}

// SetTwifInfo gets a reference to the given TwifInfo and assigns it to the TwifInfo field.
func (o *UpfInfo) SetTwifInfo(v TwifInfo) {
	o.TwifInfo = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *UpfInfo) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *UpfInfo) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *UpfInfo) SetPriority(v int32) {
	o.Priority = &v
}

// GetRedundantGtpu returns the RedundantGtpu field value if set, zero value otherwise.
func (o *UpfInfo) GetRedundantGtpu() bool {
	if o == nil || IsNil(o.RedundantGtpu) {
		var ret bool
		return ret
	}
	return *o.RedundantGtpu
}

// GetRedundantGtpuOk returns a tuple with the RedundantGtpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetRedundantGtpuOk() (*bool, bool) {
	if o == nil || IsNil(o.RedundantGtpu) {
		return nil, false
	}
	return o.RedundantGtpu, true
}

// HasRedundantGtpu returns a boolean if a field has been set.
func (o *UpfInfo) HasRedundantGtpu() bool {
	if o != nil && !IsNil(o.RedundantGtpu) {
		return true
	}

	return false
}

// SetRedundantGtpu gets a reference to the given bool and assigns it to the RedundantGtpu field.
func (o *UpfInfo) SetRedundantGtpu(v bool) {
	o.RedundantGtpu = &v
}

// GetIpups returns the Ipups field value if set, zero value otherwise.
func (o *UpfInfo) GetIpups() bool {
	if o == nil || IsNil(o.Ipups) {
		var ret bool
		return ret
	}
	return *o.Ipups
}

// GetIpupsOk returns a tuple with the Ipups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetIpupsOk() (*bool, bool) {
	if o == nil || IsNil(o.Ipups) {
		return nil, false
	}
	return o.Ipups, true
}

// HasIpups returns a boolean if a field has been set.
func (o *UpfInfo) HasIpups() bool {
	if o != nil && !IsNil(o.Ipups) {
		return true
	}

	return false
}

// SetIpups gets a reference to the given bool and assigns it to the Ipups field.
func (o *UpfInfo) SetIpups(v bool) {
	o.Ipups = &v
}

// GetDataForwarding returns the DataForwarding field value if set, zero value otherwise.
func (o *UpfInfo) GetDataForwarding() bool {
	if o == nil || IsNil(o.DataForwarding) {
		var ret bool
		return ret
	}
	return *o.DataForwarding
}

// GetDataForwardingOk returns a tuple with the DataForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetDataForwardingOk() (*bool, bool) {
	if o == nil || IsNil(o.DataForwarding) {
		return nil, false
	}
	return o.DataForwarding, true
}

// HasDataForwarding returns a boolean if a field has been set.
func (o *UpfInfo) HasDataForwarding() bool {
	if o != nil && !IsNil(o.DataForwarding) {
		return true
	}

	return false
}

// SetDataForwarding gets a reference to the given bool and assigns it to the DataForwarding field.
func (o *UpfInfo) SetDataForwarding(v bool) {
	o.DataForwarding = &v
}

// GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field value if set, zero value otherwise.
func (o *UpfInfo) GetSupportedPfcpFeatures() string {
	if o == nil || IsNil(o.SupportedPfcpFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedPfcpFeatures
}

// GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfInfo) GetSupportedPfcpFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedPfcpFeatures) {
		return nil, false
	}
	return o.SupportedPfcpFeatures, true
}

// HasSupportedPfcpFeatures returns a boolean if a field has been set.
func (o *UpfInfo) HasSupportedPfcpFeatures() bool {
	if o != nil && !IsNil(o.SupportedPfcpFeatures) {
		return true
	}

	return false
}

// SetSupportedPfcpFeatures gets a reference to the given string and assigns it to the SupportedPfcpFeatures field.
func (o *UpfInfo) SetSupportedPfcpFeatures(v string) {
	o.SupportedPfcpFeatures = &v
}

func (o UpfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssaiUpfInfoList"] = o.SNssaiUpfInfoList
	if !IsNil(o.SmfServingArea) {
		toSerialize["smfServingArea"] = o.SmfServingArea
	}
	if !IsNil(o.InterfaceUpfInfoList) {
		toSerialize["interfaceUpfInfoList"] = o.InterfaceUpfInfoList
	}
	if !IsNil(o.IwkEpsInd) {
		toSerialize["iwkEpsInd"] = o.IwkEpsInd
	}
	if !IsNil(o.SxaInd) {
		toSerialize["sxaInd"] = o.SxaInd
	}
	if !IsNil(o.PduSessionTypes) {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if !IsNil(o.AtsssCapability) {
		toSerialize["atsssCapability"] = o.AtsssCapability
	}
	if !IsNil(o.UeIpAddrInd) {
		toSerialize["ueIpAddrInd"] = o.UeIpAddrInd
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !IsNil(o.WAgfInfo) {
		toSerialize["wAgfInfo"] = o.WAgfInfo
	}
	if !IsNil(o.TngfInfo) {
		toSerialize["tngfInfo"] = o.TngfInfo
	}
	if !IsNil(o.TwifInfo) {
		toSerialize["twifInfo"] = o.TwifInfo
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.RedundantGtpu) {
		toSerialize["redundantGtpu"] = o.RedundantGtpu
	}
	if !IsNil(o.Ipups) {
		toSerialize["ipups"] = o.Ipups
	}
	if !IsNil(o.DataForwarding) {
		toSerialize["dataForwarding"] = o.DataForwarding
	}
	if !IsNil(o.SupportedPfcpFeatures) {
		toSerialize["supportedPfcpFeatures"] = o.SupportedPfcpFeatures
	}
	return toSerialize, nil
}

type NullableUpfInfo struct {
	value *UpfInfo
	isSet bool
}

func (v NullableUpfInfo) Get() *UpfInfo {
	return v.value
}

func (v *NullableUpfInfo) Set(val *UpfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfInfo(val *UpfInfo) *NullableUpfInfo {
	return &NullableUpfInfo{value: val, isSet: true}
}

func (v NullableUpfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
