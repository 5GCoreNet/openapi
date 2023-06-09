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

// checks if the UpfFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfFunctionSingleAllOfAttributes{}

// UpfFunctionSingleAllOfAttributes struct for UpfFunctionSingleAllOfAttributes
type UpfFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	PLMNInfoList          []PlmnInfo             `json:"pLMNInfoList,omitempty"`
	NRTACList             []int32                `json:"nRTACList,omitempty"`
	CNSIIdList            []string               `json:"cNSIIdList,omitempty"`
	SmfServingArea        *string                `json:"smfServingArea,omitempty"`
	InterfaceUpfInfoList  []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty"`
	IwkEpsInd             *bool                  `json:"iwkEpsInd,omitempty"`
	PduSessionTypes       *string                `json:"pduSessionTypes,omitempty"`
	AtsssCapability       *AtsssCapability       `json:"atsssCapability,omitempty"`
	UeIpAddrInd           *bool                  `json:"ueIpAddrInd,omitempty"`
	TaiList               []Tai                  `json:"taiList,omitempty"`
	TaiRangeList          []TaiRange             `json:"taiRangeList,omitempty"`
	WAgfInfo              *IpInterface           `json:"wAgfInfo,omitempty"`
	TngfInfo              *IpInterface           `json:"tngfInfo,omitempty"`
	TwifInfo              *IpInterface           `json:"twifInfo,omitempty"`
	Priority              *int32                 `json:"priority,omitempty"`
	RedundantGtpu         *bool                  `json:"redundantGtpu,omitempty"`
	Ipups                 *bool                  `json:"ipups,omitempty"`
	DataForwarding        *bool                  `json:"dataForwarding,omitempty"`
	SupportedPfcpFeatures *string                `json:"supportedPfcpFeatures,omitempty"`
	ManagedNFProfile      *ManagedNFProfile      `json:"managedNFProfile,omitempty"`
	SupportedBMOList      []string               `json:"supportedBMOList,omitempty"`
}

// NewUpfFunctionSingleAllOfAttributes instantiates a new UpfFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfFunctionSingleAllOfAttributes() *UpfFunctionSingleAllOfAttributes {
	this := UpfFunctionSingleAllOfAttributes{}
	return &this
}

// NewUpfFunctionSingleAllOfAttributesWithDefaults instantiates a new UpfFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfFunctionSingleAllOfAttributesWithDefaults() *UpfFunctionSingleAllOfAttributes {
	this := UpfFunctionSingleAllOfAttributes{}
	return &this
}

// GetPLMNInfoList returns the PLMNInfoList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetPLMNInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PLMNInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PLMNInfoList
}

// GetPLMNInfoListOk returns a tuple with the PLMNInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetPLMNInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PLMNInfoList) {
		return nil, false
	}
	return o.PLMNInfoList, true
}

// HasPLMNInfoList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasPLMNInfoList() bool {
	if o != nil && !IsNil(o.PLMNInfoList) {
		return true
	}

	return false
}

// SetPLMNInfoList gets a reference to the given []PlmnInfo and assigns it to the PLMNInfoList field.
func (o *UpfFunctionSingleAllOfAttributes) SetPLMNInfoList(v []PlmnInfo) {
	o.PLMNInfoList = v
}

// GetNRTACList returns the NRTACList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetNRTACList() []int32 {
	if o == nil || IsNil(o.NRTACList) {
		var ret []int32
		return ret
	}
	return o.NRTACList
}

// GetNRTACListOk returns a tuple with the NRTACList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetNRTACListOk() ([]int32, bool) {
	if o == nil || IsNil(o.NRTACList) {
		return nil, false
	}
	return o.NRTACList, true
}

// HasNRTACList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasNRTACList() bool {
	if o != nil && !IsNil(o.NRTACList) {
		return true
	}

	return false
}

// SetNRTACList gets a reference to the given []int32 and assigns it to the NRTACList field.
func (o *UpfFunctionSingleAllOfAttributes) SetNRTACList(v []int32) {
	o.NRTACList = v
}

// GetCNSIIdList returns the CNSIIdList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetCNSIIdList() []string {
	if o == nil || IsNil(o.CNSIIdList) {
		var ret []string
		return ret
	}
	return o.CNSIIdList
}

// GetCNSIIdListOk returns a tuple with the CNSIIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetCNSIIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.CNSIIdList) {
		return nil, false
	}
	return o.CNSIIdList, true
}

// HasCNSIIdList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasCNSIIdList() bool {
	if o != nil && !IsNil(o.CNSIIdList) {
		return true
	}

	return false
}

// SetCNSIIdList gets a reference to the given []string and assigns it to the CNSIIdList field.
func (o *UpfFunctionSingleAllOfAttributes) SetCNSIIdList(v []string) {
	o.CNSIIdList = v
}

// GetSmfServingArea returns the SmfServingArea field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetSmfServingArea() string {
	if o == nil || IsNil(o.SmfServingArea) {
		var ret string
		return ret
	}
	return *o.SmfServingArea
}

// GetSmfServingAreaOk returns a tuple with the SmfServingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetSmfServingAreaOk() (*string, bool) {
	if o == nil || IsNil(o.SmfServingArea) {
		return nil, false
	}
	return o.SmfServingArea, true
}

// HasSmfServingArea returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasSmfServingArea() bool {
	if o != nil && !IsNil(o.SmfServingArea) {
		return true
	}

	return false
}

// SetSmfServingArea gets a reference to the given string and assigns it to the SmfServingArea field.
func (o *UpfFunctionSingleAllOfAttributes) SetSmfServingArea(v string) {
	o.SmfServingArea = &v
}

// GetInterfaceUpfInfoList returns the InterfaceUpfInfoList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetInterfaceUpfInfoList() []InterfaceUpfInfoItem {
	if o == nil || IsNil(o.InterfaceUpfInfoList) {
		var ret []InterfaceUpfInfoItem
		return ret
	}
	return o.InterfaceUpfInfoList
}

// GetInterfaceUpfInfoListOk returns a tuple with the InterfaceUpfInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetInterfaceUpfInfoListOk() ([]InterfaceUpfInfoItem, bool) {
	if o == nil || IsNil(o.InterfaceUpfInfoList) {
		return nil, false
	}
	return o.InterfaceUpfInfoList, true
}

// HasInterfaceUpfInfoList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasInterfaceUpfInfoList() bool {
	if o != nil && !IsNil(o.InterfaceUpfInfoList) {
		return true
	}

	return false
}

// SetInterfaceUpfInfoList gets a reference to the given []InterfaceUpfInfoItem and assigns it to the InterfaceUpfInfoList field.
func (o *UpfFunctionSingleAllOfAttributes) SetInterfaceUpfInfoList(v []InterfaceUpfInfoItem) {
	o.InterfaceUpfInfoList = v
}

// GetIwkEpsInd returns the IwkEpsInd field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetIwkEpsInd() bool {
	if o == nil || IsNil(o.IwkEpsInd) {
		var ret bool
		return ret
	}
	return *o.IwkEpsInd
}

// GetIwkEpsIndOk returns a tuple with the IwkEpsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetIwkEpsIndOk() (*bool, bool) {
	if o == nil || IsNil(o.IwkEpsInd) {
		return nil, false
	}
	return o.IwkEpsInd, true
}

// HasIwkEpsInd returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasIwkEpsInd() bool {
	if o != nil && !IsNil(o.IwkEpsInd) {
		return true
	}

	return false
}

// SetIwkEpsInd gets a reference to the given bool and assigns it to the IwkEpsInd field.
func (o *UpfFunctionSingleAllOfAttributes) SetIwkEpsInd(v bool) {
	o.IwkEpsInd = &v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetPduSessionTypes() string {
	if o == nil || IsNil(o.PduSessionTypes) {
		var ret string
		return ret
	}
	return *o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetPduSessionTypesOk() (*string, bool) {
	if o == nil || IsNil(o.PduSessionTypes) {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasPduSessionTypes() bool {
	if o != nil && !IsNil(o.PduSessionTypes) {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given string and assigns it to the PduSessionTypes field.
func (o *UpfFunctionSingleAllOfAttributes) SetPduSessionTypes(v string) {
	o.PduSessionTypes = &v
}

// GetAtsssCapability returns the AtsssCapability field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetAtsssCapability() AtsssCapability {
	if o == nil || IsNil(o.AtsssCapability) {
		var ret AtsssCapability
		return ret
	}
	return *o.AtsssCapability
}

// GetAtsssCapabilityOk returns a tuple with the AtsssCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetAtsssCapabilityOk() (*AtsssCapability, bool) {
	if o == nil || IsNil(o.AtsssCapability) {
		return nil, false
	}
	return o.AtsssCapability, true
}

// HasAtsssCapability returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasAtsssCapability() bool {
	if o != nil && !IsNil(o.AtsssCapability) {
		return true
	}

	return false
}

// SetAtsssCapability gets a reference to the given AtsssCapability and assigns it to the AtsssCapability field.
func (o *UpfFunctionSingleAllOfAttributes) SetAtsssCapability(v AtsssCapability) {
	o.AtsssCapability = &v
}

// GetUeIpAddrInd returns the UeIpAddrInd field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetUeIpAddrInd() bool {
	if o == nil || IsNil(o.UeIpAddrInd) {
		var ret bool
		return ret
	}
	return *o.UeIpAddrInd
}

// GetUeIpAddrIndOk returns a tuple with the UeIpAddrInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetUeIpAddrIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UeIpAddrInd) {
		return nil, false
	}
	return o.UeIpAddrInd, true
}

// HasUeIpAddrInd returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasUeIpAddrInd() bool {
	if o != nil && !IsNil(o.UeIpAddrInd) {
		return true
	}

	return false
}

// SetUeIpAddrInd gets a reference to the given bool and assigns it to the UeIpAddrInd field.
func (o *UpfFunctionSingleAllOfAttributes) SetUeIpAddrInd(v bool) {
	o.UeIpAddrInd = &v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *UpfFunctionSingleAllOfAttributes) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *UpfFunctionSingleAllOfAttributes) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetWAgfInfo returns the WAgfInfo field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetWAgfInfo() IpInterface {
	if o == nil || IsNil(o.WAgfInfo) {
		var ret IpInterface
		return ret
	}
	return *o.WAgfInfo
}

// GetWAgfInfoOk returns a tuple with the WAgfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetWAgfInfoOk() (*IpInterface, bool) {
	if o == nil || IsNil(o.WAgfInfo) {
		return nil, false
	}
	return o.WAgfInfo, true
}

// HasWAgfInfo returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasWAgfInfo() bool {
	if o != nil && !IsNil(o.WAgfInfo) {
		return true
	}

	return false
}

// SetWAgfInfo gets a reference to the given IpInterface and assigns it to the WAgfInfo field.
func (o *UpfFunctionSingleAllOfAttributes) SetWAgfInfo(v IpInterface) {
	o.WAgfInfo = &v
}

// GetTngfInfo returns the TngfInfo field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetTngfInfo() IpInterface {
	if o == nil || IsNil(o.TngfInfo) {
		var ret IpInterface
		return ret
	}
	return *o.TngfInfo
}

// GetTngfInfoOk returns a tuple with the TngfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetTngfInfoOk() (*IpInterface, bool) {
	if o == nil || IsNil(o.TngfInfo) {
		return nil, false
	}
	return o.TngfInfo, true
}

// HasTngfInfo returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasTngfInfo() bool {
	if o != nil && !IsNil(o.TngfInfo) {
		return true
	}

	return false
}

// SetTngfInfo gets a reference to the given IpInterface and assigns it to the TngfInfo field.
func (o *UpfFunctionSingleAllOfAttributes) SetTngfInfo(v IpInterface) {
	o.TngfInfo = &v
}

// GetTwifInfo returns the TwifInfo field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetTwifInfo() IpInterface {
	if o == nil || IsNil(o.TwifInfo) {
		var ret IpInterface
		return ret
	}
	return *o.TwifInfo
}

// GetTwifInfoOk returns a tuple with the TwifInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetTwifInfoOk() (*IpInterface, bool) {
	if o == nil || IsNil(o.TwifInfo) {
		return nil, false
	}
	return o.TwifInfo, true
}

// HasTwifInfo returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasTwifInfo() bool {
	if o != nil && !IsNil(o.TwifInfo) {
		return true
	}

	return false
}

// SetTwifInfo gets a reference to the given IpInterface and assigns it to the TwifInfo field.
func (o *UpfFunctionSingleAllOfAttributes) SetTwifInfo(v IpInterface) {
	o.TwifInfo = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *UpfFunctionSingleAllOfAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetRedundantGtpu returns the RedundantGtpu field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetRedundantGtpu() bool {
	if o == nil || IsNil(o.RedundantGtpu) {
		var ret bool
		return ret
	}
	return *o.RedundantGtpu
}

// GetRedundantGtpuOk returns a tuple with the RedundantGtpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetRedundantGtpuOk() (*bool, bool) {
	if o == nil || IsNil(o.RedundantGtpu) {
		return nil, false
	}
	return o.RedundantGtpu, true
}

// HasRedundantGtpu returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasRedundantGtpu() bool {
	if o != nil && !IsNil(o.RedundantGtpu) {
		return true
	}

	return false
}

// SetRedundantGtpu gets a reference to the given bool and assigns it to the RedundantGtpu field.
func (o *UpfFunctionSingleAllOfAttributes) SetRedundantGtpu(v bool) {
	o.RedundantGtpu = &v
}

// GetIpups returns the Ipups field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetIpups() bool {
	if o == nil || IsNil(o.Ipups) {
		var ret bool
		return ret
	}
	return *o.Ipups
}

// GetIpupsOk returns a tuple with the Ipups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetIpupsOk() (*bool, bool) {
	if o == nil || IsNil(o.Ipups) {
		return nil, false
	}
	return o.Ipups, true
}

// HasIpups returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasIpups() bool {
	if o != nil && !IsNil(o.Ipups) {
		return true
	}

	return false
}

// SetIpups gets a reference to the given bool and assigns it to the Ipups field.
func (o *UpfFunctionSingleAllOfAttributes) SetIpups(v bool) {
	o.Ipups = &v
}

// GetDataForwarding returns the DataForwarding field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetDataForwarding() bool {
	if o == nil || IsNil(o.DataForwarding) {
		var ret bool
		return ret
	}
	return *o.DataForwarding
}

// GetDataForwardingOk returns a tuple with the DataForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetDataForwardingOk() (*bool, bool) {
	if o == nil || IsNil(o.DataForwarding) {
		return nil, false
	}
	return o.DataForwarding, true
}

// HasDataForwarding returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasDataForwarding() bool {
	if o != nil && !IsNil(o.DataForwarding) {
		return true
	}

	return false
}

// SetDataForwarding gets a reference to the given bool and assigns it to the DataForwarding field.
func (o *UpfFunctionSingleAllOfAttributes) SetDataForwarding(v bool) {
	o.DataForwarding = &v
}

// GetSupportedPfcpFeatures returns the SupportedPfcpFeatures field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetSupportedPfcpFeatures() string {
	if o == nil || IsNil(o.SupportedPfcpFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedPfcpFeatures
}

// GetSupportedPfcpFeaturesOk returns a tuple with the SupportedPfcpFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetSupportedPfcpFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedPfcpFeatures) {
		return nil, false
	}
	return o.SupportedPfcpFeatures, true
}

// HasSupportedPfcpFeatures returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasSupportedPfcpFeatures() bool {
	if o != nil && !IsNil(o.SupportedPfcpFeatures) {
		return true
	}

	return false
}

// SetSupportedPfcpFeatures gets a reference to the given string and assigns it to the SupportedPfcpFeatures field.
func (o *UpfFunctionSingleAllOfAttributes) SetSupportedPfcpFeatures(v string) {
	o.SupportedPfcpFeatures = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *UpfFunctionSingleAllOfAttributes) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetSupportedBMOList returns the SupportedBMOList field value if set, zero value otherwise.
func (o *UpfFunctionSingleAllOfAttributes) GetSupportedBMOList() []string {
	if o == nil || IsNil(o.SupportedBMOList) {
		var ret []string
		return ret
	}
	return o.SupportedBMOList
}

// GetSupportedBMOListOk returns a tuple with the SupportedBMOList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfFunctionSingleAllOfAttributes) GetSupportedBMOListOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedBMOList) {
		return nil, false
	}
	return o.SupportedBMOList, true
}

// HasSupportedBMOList returns a boolean if a field has been set.
func (o *UpfFunctionSingleAllOfAttributes) HasSupportedBMOList() bool {
	if o != nil && !IsNil(o.SupportedBMOList) {
		return true
	}

	return false
}

// SetSupportedBMOList gets a reference to the given []string and assigns it to the SupportedBMOList field.
func (o *UpfFunctionSingleAllOfAttributes) SetSupportedBMOList(v []string) {
	o.SupportedBMOList = v
}

func (o UpfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.PLMNInfoList) {
		toSerialize["pLMNInfoList"] = o.PLMNInfoList
	}
	if !IsNil(o.NRTACList) {
		toSerialize["nRTACList"] = o.NRTACList
	}
	if !IsNil(o.CNSIIdList) {
		toSerialize["cNSIIdList"] = o.CNSIIdList
	}
	if !IsNil(o.SmfServingArea) {
		toSerialize["smfServingArea"] = o.SmfServingArea
	}
	if !IsNil(o.InterfaceUpfInfoList) {
		toSerialize["interfaceUpfInfoList"] = o.InterfaceUpfInfoList
	}
	if !IsNil(o.IwkEpsInd) {
		toSerialize["iwkEpsInd"] = o.IwkEpsInd
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
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.SupportedBMOList) {
		toSerialize["supportedBMOList"] = o.SupportedBMOList
	}
	return toSerialize, nil
}

type NullableUpfFunctionSingleAllOfAttributes struct {
	value *UpfFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableUpfFunctionSingleAllOfAttributes) Get() *UpfFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableUpfFunctionSingleAllOfAttributes) Set(val *UpfFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfFunctionSingleAllOfAttributes(val *UpfFunctionSingleAllOfAttributes) *NullableUpfFunctionSingleAllOfAttributes {
	return &NullableUpfFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableUpfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
