/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the HsmfUpdatedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HsmfUpdatedData{}

// HsmfUpdatedData Data within Update Response from H-SMF, or from SMF to I-SMF
type HsmfUpdatedData struct {
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	N4Info       *N4Information   `json:"n4Info,omitempty"`
	N4InfoExt1   *N4Information   `json:"n4InfoExt1,omitempty"`
	N4InfoExt2   *N4Information   `json:"n4InfoExt2,omitempty"`
	DnaiList     []string         `json:"dnaiList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures               *string                        `json:"supportedFeatures,omitempty"`
	RoamingChargingProfile          *RoamingChargingProfile        `json:"roamingChargingProfile,omitempty"`
	HomeProvidedChargingId          *string                        `json:"homeProvidedChargingId,omitempty"`
	UpSecurity                      *UpSecurity                    `json:"upSecurity,omitempty"`
	MaxIntegrityProtectedDataRateUl *MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateUl,omitempty"`
	MaxIntegrityProtectedDataRateDl *MaxIntegrityProtectedDataRate `json:"maxIntegrityProtectedDataRateDl,omitempty"`
	Ipv6MultiHomingInd              *bool                          `json:"ipv6MultiHomingInd,omitempty"`
	QosFlowsSetupList               []QosFlowSetupItem             `json:"qosFlowsSetupList,omitempty"`
	SessionAmbr                     *Ambr                          `json:"sessionAmbr,omitempty"`
	EpsPdnCnxInfo                   *EpsPdnCnxInfo                 `json:"epsPdnCnxInfo,omitempty"`
	EpsBearerInfo                   []EpsBearerInfo                `json:"epsBearerInfo,omitempty"`
	// Procedure Transaction Identifier
	Pti *int32 `json:"pti,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	InterPlmnApiRoot *string `json:"interPlmnApiRoot,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	IntraPlmnApiRoot *string `json:"intraPlmnApiRoot,omitempty"`
}

// NewHsmfUpdatedData instantiates a new HsmfUpdatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHsmfUpdatedData() *HsmfUpdatedData {
	this := HsmfUpdatedData{}
	var ipv6MultiHomingInd bool = false
	this.Ipv6MultiHomingInd = &ipv6MultiHomingInd
	return &this
}

// NewHsmfUpdatedDataWithDefaults instantiates a new HsmfUpdatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHsmfUpdatedDataWithDefaults() *HsmfUpdatedData {
	this := HsmfUpdatedData{}
	var ipv6MultiHomingInd bool = false
	this.Ipv6MultiHomingInd = &ipv6MultiHomingInd
	return &this
}

// GetN1SmInfoToUe returns the N1SmInfoToUe field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetN1SmInfoToUe() RefToBinaryData {
	if o == nil || IsNil(o.N1SmInfoToUe) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N1SmInfoToUe
}

// GetN1SmInfoToUeOk returns a tuple with the N1SmInfoToUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetN1SmInfoToUeOk() (*RefToBinaryData, bool) {
	if o == nil || IsNil(o.N1SmInfoToUe) {
		return nil, false
	}
	return o.N1SmInfoToUe, true
}

// HasN1SmInfoToUe returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasN1SmInfoToUe() bool {
	if o != nil && !IsNil(o.N1SmInfoToUe) {
		return true
	}

	return false
}

// SetN1SmInfoToUe gets a reference to the given RefToBinaryData and assigns it to the N1SmInfoToUe field.
func (o *HsmfUpdatedData) SetN1SmInfoToUe(v RefToBinaryData) {
	o.N1SmInfoToUe = &v
}

// GetN4Info returns the N4Info field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetN4Info() N4Information {
	if o == nil || IsNil(o.N4Info) {
		var ret N4Information
		return ret
	}
	return *o.N4Info
}

// GetN4InfoOk returns a tuple with the N4Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetN4InfoOk() (*N4Information, bool) {
	if o == nil || IsNil(o.N4Info) {
		return nil, false
	}
	return o.N4Info, true
}

// HasN4Info returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasN4Info() bool {
	if o != nil && !IsNil(o.N4Info) {
		return true
	}

	return false
}

// SetN4Info gets a reference to the given N4Information and assigns it to the N4Info field.
func (o *HsmfUpdatedData) SetN4Info(v N4Information) {
	o.N4Info = &v
}

// GetN4InfoExt1 returns the N4InfoExt1 field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetN4InfoExt1() N4Information {
	if o == nil || IsNil(o.N4InfoExt1) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt1
}

// GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetN4InfoExt1Ok() (*N4Information, bool) {
	if o == nil || IsNil(o.N4InfoExt1) {
		return nil, false
	}
	return o.N4InfoExt1, true
}

// HasN4InfoExt1 returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasN4InfoExt1() bool {
	if o != nil && !IsNil(o.N4InfoExt1) {
		return true
	}

	return false
}

// SetN4InfoExt1 gets a reference to the given N4Information and assigns it to the N4InfoExt1 field.
func (o *HsmfUpdatedData) SetN4InfoExt1(v N4Information) {
	o.N4InfoExt1 = &v
}

// GetN4InfoExt2 returns the N4InfoExt2 field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetN4InfoExt2() N4Information {
	if o == nil || IsNil(o.N4InfoExt2) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt2
}

// GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetN4InfoExt2Ok() (*N4Information, bool) {
	if o == nil || IsNil(o.N4InfoExt2) {
		return nil, false
	}
	return o.N4InfoExt2, true
}

// HasN4InfoExt2 returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasN4InfoExt2() bool {
	if o != nil && !IsNil(o.N4InfoExt2) {
		return true
	}

	return false
}

// SetN4InfoExt2 gets a reference to the given N4Information and assigns it to the N4InfoExt2 field.
func (o *HsmfUpdatedData) SetN4InfoExt2(v N4Information) {
	o.N4InfoExt2 = &v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetDnaiList() []string {
	if o == nil || IsNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetDnaiListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnaiList) {
		return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasDnaiList() bool {
	if o != nil && !IsNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *HsmfUpdatedData) SetDnaiList(v []string) {
	o.DnaiList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *HsmfUpdatedData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetRoamingChargingProfile returns the RoamingChargingProfile field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetRoamingChargingProfile() RoamingChargingProfile {
	if o == nil || IsNil(o.RoamingChargingProfile) {
		var ret RoamingChargingProfile
		return ret
	}
	return *o.RoamingChargingProfile
}

// GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool) {
	if o == nil || IsNil(o.RoamingChargingProfile) {
		return nil, false
	}
	return o.RoamingChargingProfile, true
}

// HasRoamingChargingProfile returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasRoamingChargingProfile() bool {
	if o != nil && !IsNil(o.RoamingChargingProfile) {
		return true
	}

	return false
}

// SetRoamingChargingProfile gets a reference to the given RoamingChargingProfile and assigns it to the RoamingChargingProfile field.
func (o *HsmfUpdatedData) SetRoamingChargingProfile(v RoamingChargingProfile) {
	o.RoamingChargingProfile = &v
}

// GetHomeProvidedChargingId returns the HomeProvidedChargingId field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetHomeProvidedChargingId() string {
	if o == nil || IsNil(o.HomeProvidedChargingId) {
		var ret string
		return ret
	}
	return *o.HomeProvidedChargingId
}

// GetHomeProvidedChargingIdOk returns a tuple with the HomeProvidedChargingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetHomeProvidedChargingIdOk() (*string, bool) {
	if o == nil || IsNil(o.HomeProvidedChargingId) {
		return nil, false
	}
	return o.HomeProvidedChargingId, true
}

// HasHomeProvidedChargingId returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasHomeProvidedChargingId() bool {
	if o != nil && !IsNil(o.HomeProvidedChargingId) {
		return true
	}

	return false
}

// SetHomeProvidedChargingId gets a reference to the given string and assigns it to the HomeProvidedChargingId field.
func (o *HsmfUpdatedData) SetHomeProvidedChargingId(v string) {
	o.HomeProvidedChargingId = &v
}

// GetUpSecurity returns the UpSecurity field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetUpSecurity() UpSecurity {
	if o == nil || IsNil(o.UpSecurity) {
		var ret UpSecurity
		return ret
	}
	return *o.UpSecurity
}

// GetUpSecurityOk returns a tuple with the UpSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetUpSecurityOk() (*UpSecurity, bool) {
	if o == nil || IsNil(o.UpSecurity) {
		return nil, false
	}
	return o.UpSecurity, true
}

// HasUpSecurity returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasUpSecurity() bool {
	if o != nil && !IsNil(o.UpSecurity) {
		return true
	}

	return false
}

// SetUpSecurity gets a reference to the given UpSecurity and assigns it to the UpSecurity field.
func (o *HsmfUpdatedData) SetUpSecurity(v UpSecurity) {
	o.UpSecurity = &v
}

// GetMaxIntegrityProtectedDataRateUl returns the MaxIntegrityProtectedDataRateUl field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateUl() MaxIntegrityProtectedDataRate {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateUl) {
		var ret MaxIntegrityProtectedDataRate
		return ret
	}
	return *o.MaxIntegrityProtectedDataRateUl
}

// GetMaxIntegrityProtectedDataRateUlOk returns a tuple with the MaxIntegrityProtectedDataRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateUlOk() (*MaxIntegrityProtectedDataRate, bool) {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateUl) {
		return nil, false
	}
	return o.MaxIntegrityProtectedDataRateUl, true
}

// HasMaxIntegrityProtectedDataRateUl returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasMaxIntegrityProtectedDataRateUl() bool {
	if o != nil && !IsNil(o.MaxIntegrityProtectedDataRateUl) {
		return true
	}

	return false
}

// SetMaxIntegrityProtectedDataRateUl gets a reference to the given MaxIntegrityProtectedDataRate and assigns it to the MaxIntegrityProtectedDataRateUl field.
func (o *HsmfUpdatedData) SetMaxIntegrityProtectedDataRateUl(v MaxIntegrityProtectedDataRate) {
	o.MaxIntegrityProtectedDataRateUl = &v
}

// GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateDl) {
		var ret MaxIntegrityProtectedDataRate
		return ret
	}
	return *o.MaxIntegrityProtectedDataRateDl
}

// GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool) {
	if o == nil || IsNil(o.MaxIntegrityProtectedDataRateDl) {
		return nil, false
	}
	return o.MaxIntegrityProtectedDataRateDl, true
}

// HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasMaxIntegrityProtectedDataRateDl() bool {
	if o != nil && !IsNil(o.MaxIntegrityProtectedDataRateDl) {
		return true
	}

	return false
}

// SetMaxIntegrityProtectedDataRateDl gets a reference to the given MaxIntegrityProtectedDataRate and assigns it to the MaxIntegrityProtectedDataRateDl field.
func (o *HsmfUpdatedData) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate) {
	o.MaxIntegrityProtectedDataRateDl = &v
}

// GetIpv6MultiHomingInd returns the Ipv6MultiHomingInd field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetIpv6MultiHomingInd() bool {
	if o == nil || IsNil(o.Ipv6MultiHomingInd) {
		var ret bool
		return ret
	}
	return *o.Ipv6MultiHomingInd
}

// GetIpv6MultiHomingIndOk returns a tuple with the Ipv6MultiHomingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetIpv6MultiHomingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.Ipv6MultiHomingInd) {
		return nil, false
	}
	return o.Ipv6MultiHomingInd, true
}

// HasIpv6MultiHomingInd returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasIpv6MultiHomingInd() bool {
	if o != nil && !IsNil(o.Ipv6MultiHomingInd) {
		return true
	}

	return false
}

// SetIpv6MultiHomingInd gets a reference to the given bool and assigns it to the Ipv6MultiHomingInd field.
func (o *HsmfUpdatedData) SetIpv6MultiHomingInd(v bool) {
	o.Ipv6MultiHomingInd = &v
}

// GetQosFlowsSetupList returns the QosFlowsSetupList field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetQosFlowsSetupList() []QosFlowSetupItem {
	if o == nil || IsNil(o.QosFlowsSetupList) {
		var ret []QosFlowSetupItem
		return ret
	}
	return o.QosFlowsSetupList
}

// GetQosFlowsSetupListOk returns a tuple with the QosFlowsSetupList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetQosFlowsSetupListOk() ([]QosFlowSetupItem, bool) {
	if o == nil || IsNil(o.QosFlowsSetupList) {
		return nil, false
	}
	return o.QosFlowsSetupList, true
}

// HasQosFlowsSetupList returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasQosFlowsSetupList() bool {
	if o != nil && !IsNil(o.QosFlowsSetupList) {
		return true
	}

	return false
}

// SetQosFlowsSetupList gets a reference to the given []QosFlowSetupItem and assigns it to the QosFlowsSetupList field.
func (o *HsmfUpdatedData) SetQosFlowsSetupList(v []QosFlowSetupItem) {
	o.QosFlowsSetupList = v
}

// GetSessionAmbr returns the SessionAmbr field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetSessionAmbr() Ambr {
	if o == nil || IsNil(o.SessionAmbr) {
		var ret Ambr
		return ret
	}
	return *o.SessionAmbr
}

// GetSessionAmbrOk returns a tuple with the SessionAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetSessionAmbrOk() (*Ambr, bool) {
	if o == nil || IsNil(o.SessionAmbr) {
		return nil, false
	}
	return o.SessionAmbr, true
}

// HasSessionAmbr returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasSessionAmbr() bool {
	if o != nil && !IsNil(o.SessionAmbr) {
		return true
	}

	return false
}

// SetSessionAmbr gets a reference to the given Ambr and assigns it to the SessionAmbr field.
func (o *HsmfUpdatedData) SetSessionAmbr(v Ambr) {
	o.SessionAmbr = &v
}

// GetEpsPdnCnxInfo returns the EpsPdnCnxInfo field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetEpsPdnCnxInfo() EpsPdnCnxInfo {
	if o == nil || IsNil(o.EpsPdnCnxInfo) {
		var ret EpsPdnCnxInfo
		return ret
	}
	return *o.EpsPdnCnxInfo
}

// GetEpsPdnCnxInfoOk returns a tuple with the EpsPdnCnxInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetEpsPdnCnxInfoOk() (*EpsPdnCnxInfo, bool) {
	if o == nil || IsNil(o.EpsPdnCnxInfo) {
		return nil, false
	}
	return o.EpsPdnCnxInfo, true
}

// HasEpsPdnCnxInfo returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasEpsPdnCnxInfo() bool {
	if o != nil && !IsNil(o.EpsPdnCnxInfo) {
		return true
	}

	return false
}

// SetEpsPdnCnxInfo gets a reference to the given EpsPdnCnxInfo and assigns it to the EpsPdnCnxInfo field.
func (o *HsmfUpdatedData) SetEpsPdnCnxInfo(v EpsPdnCnxInfo) {
	o.EpsPdnCnxInfo = &v
}

// GetEpsBearerInfo returns the EpsBearerInfo field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetEpsBearerInfo() []EpsBearerInfo {
	if o == nil || IsNil(o.EpsBearerInfo) {
		var ret []EpsBearerInfo
		return ret
	}
	return o.EpsBearerInfo
}

// GetEpsBearerInfoOk returns a tuple with the EpsBearerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetEpsBearerInfoOk() ([]EpsBearerInfo, bool) {
	if o == nil || IsNil(o.EpsBearerInfo) {
		return nil, false
	}
	return o.EpsBearerInfo, true
}

// HasEpsBearerInfo returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasEpsBearerInfo() bool {
	if o != nil && !IsNil(o.EpsBearerInfo) {
		return true
	}

	return false
}

// SetEpsBearerInfo gets a reference to the given []EpsBearerInfo and assigns it to the EpsBearerInfo field.
func (o *HsmfUpdatedData) SetEpsBearerInfo(v []EpsBearerInfo) {
	o.EpsBearerInfo = v
}

// GetPti returns the Pti field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetPti() int32 {
	if o == nil || IsNil(o.Pti) {
		var ret int32
		return ret
	}
	return *o.Pti
}

// GetPtiOk returns a tuple with the Pti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetPtiOk() (*int32, bool) {
	if o == nil || IsNil(o.Pti) {
		return nil, false
	}
	return o.Pti, true
}

// HasPti returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasPti() bool {
	if o != nil && !IsNil(o.Pti) {
		return true
	}

	return false
}

// SetPti gets a reference to the given int32 and assigns it to the Pti field.
func (o *HsmfUpdatedData) SetPti(v int32) {
	o.Pti = &v
}

// GetInterPlmnApiRoot returns the InterPlmnApiRoot field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetInterPlmnApiRoot() string {
	if o == nil || IsNil(o.InterPlmnApiRoot) {
		var ret string
		return ret
	}
	return *o.InterPlmnApiRoot
}

// GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetInterPlmnApiRootOk() (*string, bool) {
	if o == nil || IsNil(o.InterPlmnApiRoot) {
		return nil, false
	}
	return o.InterPlmnApiRoot, true
}

// HasInterPlmnApiRoot returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasInterPlmnApiRoot() bool {
	if o != nil && !IsNil(o.InterPlmnApiRoot) {
		return true
	}

	return false
}

// SetInterPlmnApiRoot gets a reference to the given string and assigns it to the InterPlmnApiRoot field.
func (o *HsmfUpdatedData) SetInterPlmnApiRoot(v string) {
	o.InterPlmnApiRoot = &v
}

// GetIntraPlmnApiRoot returns the IntraPlmnApiRoot field value if set, zero value otherwise.
func (o *HsmfUpdatedData) GetIntraPlmnApiRoot() string {
	if o == nil || IsNil(o.IntraPlmnApiRoot) {
		var ret string
		return ret
	}
	return *o.IntraPlmnApiRoot
}

// GetIntraPlmnApiRootOk returns a tuple with the IntraPlmnApiRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HsmfUpdatedData) GetIntraPlmnApiRootOk() (*string, bool) {
	if o == nil || IsNil(o.IntraPlmnApiRoot) {
		return nil, false
	}
	return o.IntraPlmnApiRoot, true
}

// HasIntraPlmnApiRoot returns a boolean if a field has been set.
func (o *HsmfUpdatedData) HasIntraPlmnApiRoot() bool {
	if o != nil && !IsNil(o.IntraPlmnApiRoot) {
		return true
	}

	return false
}

// SetIntraPlmnApiRoot gets a reference to the given string and assigns it to the IntraPlmnApiRoot field.
func (o *HsmfUpdatedData) SetIntraPlmnApiRoot(v string) {
	o.IntraPlmnApiRoot = &v
}

func (o HsmfUpdatedData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HsmfUpdatedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.N1SmInfoToUe) {
		toSerialize["n1SmInfoToUe"] = o.N1SmInfoToUe
	}
	if !IsNil(o.N4Info) {
		toSerialize["n4Info"] = o.N4Info
	}
	if !IsNil(o.N4InfoExt1) {
		toSerialize["n4InfoExt1"] = o.N4InfoExt1
	}
	if !IsNil(o.N4InfoExt2) {
		toSerialize["n4InfoExt2"] = o.N4InfoExt2
	}
	if !IsNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.RoamingChargingProfile) {
		toSerialize["roamingChargingProfile"] = o.RoamingChargingProfile
	}
	if !IsNil(o.HomeProvidedChargingId) {
		toSerialize["homeProvidedChargingId"] = o.HomeProvidedChargingId
	}
	if !IsNil(o.UpSecurity) {
		toSerialize["upSecurity"] = o.UpSecurity
	}
	if !IsNil(o.MaxIntegrityProtectedDataRateUl) {
		toSerialize["maxIntegrityProtectedDataRateUl"] = o.MaxIntegrityProtectedDataRateUl
	}
	if !IsNil(o.MaxIntegrityProtectedDataRateDl) {
		toSerialize["maxIntegrityProtectedDataRateDl"] = o.MaxIntegrityProtectedDataRateDl
	}
	if !IsNil(o.Ipv6MultiHomingInd) {
		toSerialize["ipv6MultiHomingInd"] = o.Ipv6MultiHomingInd
	}
	if !IsNil(o.QosFlowsSetupList) {
		toSerialize["qosFlowsSetupList"] = o.QosFlowsSetupList
	}
	if !IsNil(o.SessionAmbr) {
		toSerialize["sessionAmbr"] = o.SessionAmbr
	}
	if !IsNil(o.EpsPdnCnxInfo) {
		toSerialize["epsPdnCnxInfo"] = o.EpsPdnCnxInfo
	}
	if !IsNil(o.EpsBearerInfo) {
		toSerialize["epsBearerInfo"] = o.EpsBearerInfo
	}
	if !IsNil(o.Pti) {
		toSerialize["pti"] = o.Pti
	}
	if !IsNil(o.InterPlmnApiRoot) {
		toSerialize["interPlmnApiRoot"] = o.InterPlmnApiRoot
	}
	if !IsNil(o.IntraPlmnApiRoot) {
		toSerialize["intraPlmnApiRoot"] = o.IntraPlmnApiRoot
	}
	return toSerialize, nil
}

type NullableHsmfUpdatedData struct {
	value *HsmfUpdatedData
	isSet bool
}

func (v NullableHsmfUpdatedData) Get() *HsmfUpdatedData {
	return v.value
}

func (v *NullableHsmfUpdatedData) Set(val *HsmfUpdatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableHsmfUpdatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableHsmfUpdatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHsmfUpdatedData(val *HsmfUpdatedData) *NullableHsmfUpdatedData {
	return &NullableHsmfUpdatedData{value: val, isSet: true}
}

func (v NullableHsmfUpdatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHsmfUpdatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
