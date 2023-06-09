/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the AnalyticsEventFilterSubsc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsEventFilterSubsc{}

// AnalyticsEventFilterSubsc Represents an analytics event filter.
type AnalyticsEventFilterSubsc struct {
	NwPerfReqs []NetworkPerfRequirement `json:"nwPerfReqs,omitempty"`
	LocArea    *LocationArea5G          `json:"locArea,omitempty"`
	AppIds     []string                 `json:"appIds,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn              *string                    `json:"dnn,omitempty"`
	Dnais            []string                   `json:"dnais,omitempty"`
	ExcepRequs       []Exception                `json:"excepRequs,omitempty"`
	ExptAnaType      *ExpectedAnalyticsType     `json:"exptAnaType,omitempty"`
	ExptUeBehav      *ExpectedUeBehaviourData   `json:"exptUeBehav,omitempty"`
	MatchingDir      *MatchingDirection         `json:"matchingDir,omitempty"`
	ReptThlds        []ThresholdLevel           `json:"reptThlds,omitempty"`
	Snssai           *Snssai                    `json:"snssai,omitempty"`
	NsiIdInfos       []NsiIdInfo                `json:"nsiIdInfos,omitempty"`
	QosReq           *QosRequirement            `json:"qosReq,omitempty"`
	QosFlowRetThds   []RetainabilityThreshold   `json:"qosFlowRetThds,omitempty"`
	RanUeThrouThds   []string                   `json:"ranUeThrouThds,omitempty"`
	DisperReqs       []DispersionRequirement    `json:"disperReqs,omitempty"`
	ListOfAnaSubsets []AnalyticsSubset          `json:"listOfAnaSubsets,omitempty"`
	DnPerfReqs       []DnPerformanceReq         `json:"dnPerfReqs,omitempty"`
	BwRequs          []BwRequirement            `json:"bwRequs,omitempty"`
	RatFreqs         []RatFreqInformation       `json:"ratFreqs,omitempty"`
	AppServerAddrs   []AddrFqdn                 `json:"appServerAddrs,omitempty"`
	ExtraReportReq   *EventReportingRequirement `json:"extraReportReq,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTopAppUl *int32 `json:"maxNumOfTopAppUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTopAppDl *int32           `json:"maxNumOfTopAppDl,omitempty"`
	VisitedLocAreas  []LocationArea5G `json:"visitedLocAreas,omitempty"`
}

// NewAnalyticsEventFilterSubsc instantiates a new AnalyticsEventFilterSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsEventFilterSubsc() *AnalyticsEventFilterSubsc {
	this := AnalyticsEventFilterSubsc{}
	return &this
}

// NewAnalyticsEventFilterSubscWithDefaults instantiates a new AnalyticsEventFilterSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsEventFilterSubscWithDefaults() *AnalyticsEventFilterSubsc {
	this := AnalyticsEventFilterSubsc{}
	return &this
}

// GetNwPerfReqs returns the NwPerfReqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetNwPerfReqs() []NetworkPerfRequirement {
	if o == nil || IsNil(o.NwPerfReqs) {
		var ret []NetworkPerfRequirement
		return ret
	}
	return o.NwPerfReqs
}

// GetNwPerfReqsOk returns a tuple with the NwPerfReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetNwPerfReqsOk() ([]NetworkPerfRequirement, bool) {
	if o == nil || IsNil(o.NwPerfReqs) {
		return nil, false
	}
	return o.NwPerfReqs, true
}

// HasNwPerfReqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasNwPerfReqs() bool {
	if o != nil && !IsNil(o.NwPerfReqs) {
		return true
	}

	return false
}

// SetNwPerfReqs gets a reference to the given []NetworkPerfRequirement and assigns it to the NwPerfReqs field.
func (o *AnalyticsEventFilterSubsc) SetNwPerfReqs(v []NetworkPerfRequirement) {
	o.NwPerfReqs = v
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetLocArea() LocationArea5G {
	if o == nil || IsNil(o.LocArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.LocArea) {
		return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasLocArea() bool {
	if o != nil && !IsNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given LocationArea5G and assigns it to the LocArea field.
func (o *AnalyticsEventFilterSubsc) SetLocArea(v LocationArea5G) {
	o.LocArea = &v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetAppIds() []string {
	if o == nil || IsNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasAppIds() bool {
	if o != nil && !IsNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *AnalyticsEventFilterSubsc) SetAppIds(v []string) {
	o.AppIds = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *AnalyticsEventFilterSubsc) SetDnn(v string) {
	o.Dnn = &v
}

// GetDnais returns the Dnais field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetDnais() []string {
	if o == nil || IsNil(o.Dnais) {
		var ret []string
		return ret
	}
	return o.Dnais
}

// GetDnaisOk returns a tuple with the Dnais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetDnaisOk() ([]string, bool) {
	if o == nil || IsNil(o.Dnais) {
		return nil, false
	}
	return o.Dnais, true
}

// HasDnais returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasDnais() bool {
	if o != nil && !IsNil(o.Dnais) {
		return true
	}

	return false
}

// SetDnais gets a reference to the given []string and assigns it to the Dnais field.
func (o *AnalyticsEventFilterSubsc) SetDnais(v []string) {
	o.Dnais = v
}

// GetExcepRequs returns the ExcepRequs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetExcepRequs() []Exception {
	if o == nil || IsNil(o.ExcepRequs) {
		var ret []Exception
		return ret
	}
	return o.ExcepRequs
}

// GetExcepRequsOk returns a tuple with the ExcepRequs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetExcepRequsOk() ([]Exception, bool) {
	if o == nil || IsNil(o.ExcepRequs) {
		return nil, false
	}
	return o.ExcepRequs, true
}

// HasExcepRequs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasExcepRequs() bool {
	if o != nil && !IsNil(o.ExcepRequs) {
		return true
	}

	return false
}

// SetExcepRequs gets a reference to the given []Exception and assigns it to the ExcepRequs field.
func (o *AnalyticsEventFilterSubsc) SetExcepRequs(v []Exception) {
	o.ExcepRequs = v
}

// GetExptAnaType returns the ExptAnaType field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetExptAnaType() ExpectedAnalyticsType {
	if o == nil || IsNil(o.ExptAnaType) {
		var ret ExpectedAnalyticsType
		return ret
	}
	return *o.ExptAnaType
}

// GetExptAnaTypeOk returns a tuple with the ExptAnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool) {
	if o == nil || IsNil(o.ExptAnaType) {
		return nil, false
	}
	return o.ExptAnaType, true
}

// HasExptAnaType returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasExptAnaType() bool {
	if o != nil && !IsNil(o.ExptAnaType) {
		return true
	}

	return false
}

// SetExptAnaType gets a reference to the given ExpectedAnalyticsType and assigns it to the ExptAnaType field.
func (o *AnalyticsEventFilterSubsc) SetExptAnaType(v ExpectedAnalyticsType) {
	o.ExptAnaType = &v
}

// GetExptUeBehav returns the ExptUeBehav field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetExptUeBehav() ExpectedUeBehaviourData {
	if o == nil || IsNil(o.ExptUeBehav) {
		var ret ExpectedUeBehaviourData
		return ret
	}
	return *o.ExptUeBehav
}

// GetExptUeBehavOk returns a tuple with the ExptUeBehav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool) {
	if o == nil || IsNil(o.ExptUeBehav) {
		return nil, false
	}
	return o.ExptUeBehav, true
}

// HasExptUeBehav returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasExptUeBehav() bool {
	if o != nil && !IsNil(o.ExptUeBehav) {
		return true
	}

	return false
}

// SetExptUeBehav gets a reference to the given ExpectedUeBehaviourData and assigns it to the ExptUeBehav field.
func (o *AnalyticsEventFilterSubsc) SetExptUeBehav(v ExpectedUeBehaviourData) {
	o.ExptUeBehav = &v
}

// GetMatchingDir returns the MatchingDir field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetMatchingDir() MatchingDirection {
	if o == nil || IsNil(o.MatchingDir) {
		var ret MatchingDirection
		return ret
	}
	return *o.MatchingDir
}

// GetMatchingDirOk returns a tuple with the MatchingDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetMatchingDirOk() (*MatchingDirection, bool) {
	if o == nil || IsNil(o.MatchingDir) {
		return nil, false
	}
	return o.MatchingDir, true
}

// HasMatchingDir returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasMatchingDir() bool {
	if o != nil && !IsNil(o.MatchingDir) {
		return true
	}

	return false
}

// SetMatchingDir gets a reference to the given MatchingDirection and assigns it to the MatchingDir field.
func (o *AnalyticsEventFilterSubsc) SetMatchingDir(v MatchingDirection) {
	o.MatchingDir = &v
}

// GetReptThlds returns the ReptThlds field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetReptThlds() []ThresholdLevel {
	if o == nil || IsNil(o.ReptThlds) {
		var ret []ThresholdLevel
		return ret
	}
	return o.ReptThlds
}

// GetReptThldsOk returns a tuple with the ReptThlds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetReptThldsOk() ([]ThresholdLevel, bool) {
	if o == nil || IsNil(o.ReptThlds) {
		return nil, false
	}
	return o.ReptThlds, true
}

// HasReptThlds returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasReptThlds() bool {
	if o != nil && !IsNil(o.ReptThlds) {
		return true
	}

	return false
}

// SetReptThlds gets a reference to the given []ThresholdLevel and assigns it to the ReptThlds field.
func (o *AnalyticsEventFilterSubsc) SetReptThlds(v []ThresholdLevel) {
	o.ReptThlds = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *AnalyticsEventFilterSubsc) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetNsiIdInfos returns the NsiIdInfos field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetNsiIdInfos() []NsiIdInfo {
	if o == nil || IsNil(o.NsiIdInfos) {
		var ret []NsiIdInfo
		return ret
	}
	return o.NsiIdInfos
}

// GetNsiIdInfosOk returns a tuple with the NsiIdInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetNsiIdInfosOk() ([]NsiIdInfo, bool) {
	if o == nil || IsNil(o.NsiIdInfos) {
		return nil, false
	}
	return o.NsiIdInfos, true
}

// HasNsiIdInfos returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasNsiIdInfos() bool {
	if o != nil && !IsNil(o.NsiIdInfos) {
		return true
	}

	return false
}

// SetNsiIdInfos gets a reference to the given []NsiIdInfo and assigns it to the NsiIdInfos field.
func (o *AnalyticsEventFilterSubsc) SetNsiIdInfos(v []NsiIdInfo) {
	o.NsiIdInfos = v
}

// GetQosReq returns the QosReq field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetQosReq() QosRequirement {
	if o == nil || IsNil(o.QosReq) {
		var ret QosRequirement
		return ret
	}
	return *o.QosReq
}

// GetQosReqOk returns a tuple with the QosReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetQosReqOk() (*QosRequirement, bool) {
	if o == nil || IsNil(o.QosReq) {
		return nil, false
	}
	return o.QosReq, true
}

// HasQosReq returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasQosReq() bool {
	if o != nil && !IsNil(o.QosReq) {
		return true
	}

	return false
}

// SetQosReq gets a reference to the given QosRequirement and assigns it to the QosReq field.
func (o *AnalyticsEventFilterSubsc) SetQosReq(v QosRequirement) {
	o.QosReq = &v
}

// GetQosFlowRetThds returns the QosFlowRetThds field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetQosFlowRetThds() []RetainabilityThreshold {
	if o == nil || IsNil(o.QosFlowRetThds) {
		var ret []RetainabilityThreshold
		return ret
	}
	return o.QosFlowRetThds
}

// GetQosFlowRetThdsOk returns a tuple with the QosFlowRetThds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetQosFlowRetThdsOk() ([]RetainabilityThreshold, bool) {
	if o == nil || IsNil(o.QosFlowRetThds) {
		return nil, false
	}
	return o.QosFlowRetThds, true
}

// HasQosFlowRetThds returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasQosFlowRetThds() bool {
	if o != nil && !IsNil(o.QosFlowRetThds) {
		return true
	}

	return false
}

// SetQosFlowRetThds gets a reference to the given []RetainabilityThreshold and assigns it to the QosFlowRetThds field.
func (o *AnalyticsEventFilterSubsc) SetQosFlowRetThds(v []RetainabilityThreshold) {
	o.QosFlowRetThds = v
}

// GetRanUeThrouThds returns the RanUeThrouThds field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetRanUeThrouThds() []string {
	if o == nil || IsNil(o.RanUeThrouThds) {
		var ret []string
		return ret
	}
	return o.RanUeThrouThds
}

// GetRanUeThrouThdsOk returns a tuple with the RanUeThrouThds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetRanUeThrouThdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RanUeThrouThds) {
		return nil, false
	}
	return o.RanUeThrouThds, true
}

// HasRanUeThrouThds returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasRanUeThrouThds() bool {
	if o != nil && !IsNil(o.RanUeThrouThds) {
		return true
	}

	return false
}

// SetRanUeThrouThds gets a reference to the given []string and assigns it to the RanUeThrouThds field.
func (o *AnalyticsEventFilterSubsc) SetRanUeThrouThds(v []string) {
	o.RanUeThrouThds = v
}

// GetDisperReqs returns the DisperReqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetDisperReqs() []DispersionRequirement {
	if o == nil || IsNil(o.DisperReqs) {
		var ret []DispersionRequirement
		return ret
	}
	return o.DisperReqs
}

// GetDisperReqsOk returns a tuple with the DisperReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetDisperReqsOk() ([]DispersionRequirement, bool) {
	if o == nil || IsNil(o.DisperReqs) {
		return nil, false
	}
	return o.DisperReqs, true
}

// HasDisperReqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasDisperReqs() bool {
	if o != nil && !IsNil(o.DisperReqs) {
		return true
	}

	return false
}

// SetDisperReqs gets a reference to the given []DispersionRequirement and assigns it to the DisperReqs field.
func (o *AnalyticsEventFilterSubsc) SetDisperReqs(v []DispersionRequirement) {
	o.DisperReqs = v
}

// GetListOfAnaSubsets returns the ListOfAnaSubsets field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetListOfAnaSubsets() []AnalyticsSubset {
	if o == nil || IsNil(o.ListOfAnaSubsets) {
		var ret []AnalyticsSubset
		return ret
	}
	return o.ListOfAnaSubsets
}

// GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetListOfAnaSubsetsOk() ([]AnalyticsSubset, bool) {
	if o == nil || IsNil(o.ListOfAnaSubsets) {
		return nil, false
	}
	return o.ListOfAnaSubsets, true
}

// HasListOfAnaSubsets returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasListOfAnaSubsets() bool {
	if o != nil && !IsNil(o.ListOfAnaSubsets) {
		return true
	}

	return false
}

// SetListOfAnaSubsets gets a reference to the given []AnalyticsSubset and assigns it to the ListOfAnaSubsets field.
func (o *AnalyticsEventFilterSubsc) SetListOfAnaSubsets(v []AnalyticsSubset) {
	o.ListOfAnaSubsets = v
}

// GetDnPerfReqs returns the DnPerfReqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetDnPerfReqs() []DnPerformanceReq {
	if o == nil || IsNil(o.DnPerfReqs) {
		var ret []DnPerformanceReq
		return ret
	}
	return o.DnPerfReqs
}

// GetDnPerfReqsOk returns a tuple with the DnPerfReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetDnPerfReqsOk() ([]DnPerformanceReq, bool) {
	if o == nil || IsNil(o.DnPerfReqs) {
		return nil, false
	}
	return o.DnPerfReqs, true
}

// HasDnPerfReqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasDnPerfReqs() bool {
	if o != nil && !IsNil(o.DnPerfReqs) {
		return true
	}

	return false
}

// SetDnPerfReqs gets a reference to the given []DnPerformanceReq and assigns it to the DnPerfReqs field.
func (o *AnalyticsEventFilterSubsc) SetDnPerfReqs(v []DnPerformanceReq) {
	o.DnPerfReqs = v
}

// GetBwRequs returns the BwRequs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetBwRequs() []BwRequirement {
	if o == nil || IsNil(o.BwRequs) {
		var ret []BwRequirement
		return ret
	}
	return o.BwRequs
}

// GetBwRequsOk returns a tuple with the BwRequs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetBwRequsOk() ([]BwRequirement, bool) {
	if o == nil || IsNil(o.BwRequs) {
		return nil, false
	}
	return o.BwRequs, true
}

// HasBwRequs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasBwRequs() bool {
	if o != nil && !IsNil(o.BwRequs) {
		return true
	}

	return false
}

// SetBwRequs gets a reference to the given []BwRequirement and assigns it to the BwRequs field.
func (o *AnalyticsEventFilterSubsc) SetBwRequs(v []BwRequirement) {
	o.BwRequs = v
}

// GetRatFreqs returns the RatFreqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetRatFreqs() []RatFreqInformation {
	if o == nil || IsNil(o.RatFreqs) {
		var ret []RatFreqInformation
		return ret
	}
	return o.RatFreqs
}

// GetRatFreqsOk returns a tuple with the RatFreqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetRatFreqsOk() ([]RatFreqInformation, bool) {
	if o == nil || IsNil(o.RatFreqs) {
		return nil, false
	}
	return o.RatFreqs, true
}

// HasRatFreqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasRatFreqs() bool {
	if o != nil && !IsNil(o.RatFreqs) {
		return true
	}

	return false
}

// SetRatFreqs gets a reference to the given []RatFreqInformation and assigns it to the RatFreqs field.
func (o *AnalyticsEventFilterSubsc) SetRatFreqs(v []RatFreqInformation) {
	o.RatFreqs = v
}

// GetAppServerAddrs returns the AppServerAddrs field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetAppServerAddrs() []AddrFqdn {
	if o == nil || IsNil(o.AppServerAddrs) {
		var ret []AddrFqdn
		return ret
	}
	return o.AppServerAddrs
}

// GetAppServerAddrsOk returns a tuple with the AppServerAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetAppServerAddrsOk() ([]AddrFqdn, bool) {
	if o == nil || IsNil(o.AppServerAddrs) {
		return nil, false
	}
	return o.AppServerAddrs, true
}

// HasAppServerAddrs returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasAppServerAddrs() bool {
	if o != nil && !IsNil(o.AppServerAddrs) {
		return true
	}

	return false
}

// SetAppServerAddrs gets a reference to the given []AddrFqdn and assigns it to the AppServerAddrs field.
func (o *AnalyticsEventFilterSubsc) SetAppServerAddrs(v []AddrFqdn) {
	o.AppServerAddrs = v
}

// GetExtraReportReq returns the ExtraReportReq field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetExtraReportReq() EventReportingRequirement {
	if o == nil || IsNil(o.ExtraReportReq) {
		var ret EventReportingRequirement
		return ret
	}
	return *o.ExtraReportReq
}

// GetExtraReportReqOk returns a tuple with the ExtraReportReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetExtraReportReqOk() (*EventReportingRequirement, bool) {
	if o == nil || IsNil(o.ExtraReportReq) {
		return nil, false
	}
	return o.ExtraReportReq, true
}

// HasExtraReportReq returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasExtraReportReq() bool {
	if o != nil && !IsNil(o.ExtraReportReq) {
		return true
	}

	return false
}

// SetExtraReportReq gets a reference to the given EventReportingRequirement and assigns it to the ExtraReportReq field.
func (o *AnalyticsEventFilterSubsc) SetExtraReportReq(v EventReportingRequirement) {
	o.ExtraReportReq = &v
}

// GetMaxNumOfTopAppUl returns the MaxNumOfTopAppUl field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppUl() int32 {
	if o == nil || IsNil(o.MaxNumOfTopAppUl) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTopAppUl
}

// GetMaxNumOfTopAppUlOk returns a tuple with the MaxNumOfTopAppUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppUlOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfTopAppUl) {
		return nil, false
	}
	return o.MaxNumOfTopAppUl, true
}

// HasMaxNumOfTopAppUl returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasMaxNumOfTopAppUl() bool {
	if o != nil && !IsNil(o.MaxNumOfTopAppUl) {
		return true
	}

	return false
}

// SetMaxNumOfTopAppUl gets a reference to the given int32 and assigns it to the MaxNumOfTopAppUl field.
func (o *AnalyticsEventFilterSubsc) SetMaxNumOfTopAppUl(v int32) {
	o.MaxNumOfTopAppUl = &v
}

// GetMaxNumOfTopAppDl returns the MaxNumOfTopAppDl field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppDl() int32 {
	if o == nil || IsNil(o.MaxNumOfTopAppDl) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTopAppDl
}

// GetMaxNumOfTopAppDlOk returns a tuple with the MaxNumOfTopAppDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetMaxNumOfTopAppDlOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfTopAppDl) {
		return nil, false
	}
	return o.MaxNumOfTopAppDl, true
}

// HasMaxNumOfTopAppDl returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasMaxNumOfTopAppDl() bool {
	if o != nil && !IsNil(o.MaxNumOfTopAppDl) {
		return true
	}

	return false
}

// SetMaxNumOfTopAppDl gets a reference to the given int32 and assigns it to the MaxNumOfTopAppDl field.
func (o *AnalyticsEventFilterSubsc) SetMaxNumOfTopAppDl(v int32) {
	o.MaxNumOfTopAppDl = &v
}

// GetVisitedLocAreas returns the VisitedLocAreas field value if set, zero value otherwise.
func (o *AnalyticsEventFilterSubsc) GetVisitedLocAreas() []LocationArea5G {
	if o == nil || IsNil(o.VisitedLocAreas) {
		var ret []LocationArea5G
		return ret
	}
	return o.VisitedLocAreas
}

// GetVisitedLocAreasOk returns a tuple with the VisitedLocAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilterSubsc) GetVisitedLocAreasOk() ([]LocationArea5G, bool) {
	if o == nil || IsNil(o.VisitedLocAreas) {
		return nil, false
	}
	return o.VisitedLocAreas, true
}

// HasVisitedLocAreas returns a boolean if a field has been set.
func (o *AnalyticsEventFilterSubsc) HasVisitedLocAreas() bool {
	if o != nil && !IsNil(o.VisitedLocAreas) {
		return true
	}

	return false
}

// SetVisitedLocAreas gets a reference to the given []LocationArea5G and assigns it to the VisitedLocAreas field.
func (o *AnalyticsEventFilterSubsc) SetVisitedLocAreas(v []LocationArea5G) {
	o.VisitedLocAreas = v
}

func (o AnalyticsEventFilterSubsc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsEventFilterSubsc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NwPerfReqs) {
		toSerialize["nwPerfReqs"] = o.NwPerfReqs
	}
	if !IsNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !IsNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Dnais) {
		toSerialize["dnais"] = o.Dnais
	}
	if !IsNil(o.ExcepRequs) {
		toSerialize["excepRequs"] = o.ExcepRequs
	}
	if !IsNil(o.ExptAnaType) {
		toSerialize["exptAnaType"] = o.ExptAnaType
	}
	if !IsNil(o.ExptUeBehav) {
		toSerialize["exptUeBehav"] = o.ExptUeBehav
	}
	if !IsNil(o.MatchingDir) {
		toSerialize["matchingDir"] = o.MatchingDir
	}
	if !IsNil(o.ReptThlds) {
		toSerialize["reptThlds"] = o.ReptThlds
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.NsiIdInfos) {
		toSerialize["nsiIdInfos"] = o.NsiIdInfos
	}
	if !IsNil(o.QosReq) {
		toSerialize["qosReq"] = o.QosReq
	}
	if !IsNil(o.QosFlowRetThds) {
		toSerialize["qosFlowRetThds"] = o.QosFlowRetThds
	}
	if !IsNil(o.RanUeThrouThds) {
		toSerialize["ranUeThrouThds"] = o.RanUeThrouThds
	}
	if !IsNil(o.DisperReqs) {
		toSerialize["disperReqs"] = o.DisperReqs
	}
	if !IsNil(o.ListOfAnaSubsets) {
		toSerialize["listOfAnaSubsets"] = o.ListOfAnaSubsets
	}
	if !IsNil(o.DnPerfReqs) {
		toSerialize["dnPerfReqs"] = o.DnPerfReqs
	}
	if !IsNil(o.BwRequs) {
		toSerialize["bwRequs"] = o.BwRequs
	}
	if !IsNil(o.RatFreqs) {
		toSerialize["ratFreqs"] = o.RatFreqs
	}
	if !IsNil(o.AppServerAddrs) {
		toSerialize["appServerAddrs"] = o.AppServerAddrs
	}
	if !IsNil(o.ExtraReportReq) {
		toSerialize["extraReportReq"] = o.ExtraReportReq
	}
	if !IsNil(o.MaxNumOfTopAppUl) {
		toSerialize["maxNumOfTopAppUl"] = o.MaxNumOfTopAppUl
	}
	if !IsNil(o.MaxNumOfTopAppDl) {
		toSerialize["maxNumOfTopAppDl"] = o.MaxNumOfTopAppDl
	}
	if !IsNil(o.VisitedLocAreas) {
		toSerialize["visitedLocAreas"] = o.VisitedLocAreas
	}
	return toSerialize, nil
}

type NullableAnalyticsEventFilterSubsc struct {
	value *AnalyticsEventFilterSubsc
	isSet bool
}

func (v NullableAnalyticsEventFilterSubsc) Get() *AnalyticsEventFilterSubsc {
	return v.value
}

func (v *NullableAnalyticsEventFilterSubsc) Set(val *AnalyticsEventFilterSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsEventFilterSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsEventFilterSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsEventFilterSubsc(val *AnalyticsEventFilterSubsc) *NullableAnalyticsEventFilterSubsc {
	return &NullableAnalyticsEventFilterSubsc{value: val, isSet: true}
}

func (v NullableAnalyticsEventFilterSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsEventFilterSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
