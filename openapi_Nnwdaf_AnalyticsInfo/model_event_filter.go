/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the EventFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFilter{}

// EventFilter Represents the event filters used to identify the requested analytics.
type EventFilter struct {
	// FALSE represents not applicable for all slices. TRUE represents applicable for all slices. 
	AnySlice *bool `json:"anySlice,omitempty"`
	// Identification(s) of network slice.
	Snssais []Snssai `json:"snssais,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	Dnns []string `json:"dnns,omitempty"`
	Dnais []string `json:"dnais,omitempty"`
	// Identification(s) of LADN DNN to indicate the LADN service area as the AOI.
	LadnDnns []string `json:"ladnDnns,omitempty"`
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	VisitedAreas []NetworkAreaInfo `json:"visitedAreas,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppUlNbr *int32 `json:"maxTopAppUlNbr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxTopAppDlNbr *int32 `json:"maxTopAppDlNbr,omitempty"`
	NfInstanceIds []string `json:"nfInstanceIds,omitempty"`
	NfSetIds []string `json:"nfSetIds,omitempty"`
	NfTypes []NFType `json:"nfTypes,omitempty"`
	NsiIdInfos []NsiIdInfo `json:"nsiIdInfos,omitempty"`
	QosRequ *QosRequirement `json:"qosRequ,omitempty"`
	NwPerfTypes []NetworkPerfType `json:"nwPerfTypes,omitempty"`
	BwRequs []BwRequirement `json:"bwRequs,omitempty"`
	ExcepIds []ExceptionId `json:"excepIds,omitempty"`
	ExptAnaType *ExpectedAnalyticsType `json:"exptAnaType,omitempty"`
	ExptUeBehav *ExpectedUeBehaviourData `json:"exptUeBehav,omitempty"`
	RatFreqs []RatFreqInformation `json:"ratFreqs,omitempty"`
	DisperReqs []DispersionRequirement `json:"disperReqs,omitempty"`
	RedTransReqs []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	WlanReqs []WlanPerformanceReq `json:"wlanReqs,omitempty"`
	ListOfAnaSubsets []AnalyticsSubset `json:"listOfAnaSubsets,omitempty"`
	UpfInfo *UpfInformation `json:"upfInfo,omitempty"`
	AppServerAddrs []AddrFqdn `json:"appServerAddrs,omitempty"`
	DnPerfReqs []DnPerformanceReq `json:"dnPerfReqs,omitempty"`
}

// NewEventFilter instantiates a new EventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter() *EventFilter {
	this := EventFilter{}
	return &this
}

// NewEventFilterWithDefaults instantiates a new EventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilterWithDefaults() *EventFilter {
	this := EventFilter{}
	return &this
}

// GetAnySlice returns the AnySlice field value if set, zero value otherwise.
func (o *EventFilter) GetAnySlice() bool {
	if o == nil || isNil(o.AnySlice) {
		var ret bool
		return ret
	}
	return *o.AnySlice
}

// GetAnySliceOk returns a tuple with the AnySlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAnySliceOk() (*bool, bool) {
	if o == nil || isNil(o.AnySlice) {
		return nil, false
	}
	return o.AnySlice, true
}

// HasAnySlice returns a boolean if a field has been set.
func (o *EventFilter) HasAnySlice() bool {
	if o != nil && !isNil(o.AnySlice) {
		return true
	}

	return false
}

// SetAnySlice gets a reference to the given bool and assigns it to the AnySlice field.
func (o *EventFilter) SetAnySlice(v bool) {
	o.AnySlice = &v
}

// GetSnssais returns the Snssais field value if set, zero value otherwise.
func (o *EventFilter) GetSnssais() []Snssai {
	if o == nil || isNil(o.Snssais) {
		var ret []Snssai
		return ret
	}
	return o.Snssais
}

// GetSnssaisOk returns a tuple with the Snssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetSnssaisOk() ([]Snssai, bool) {
	if o == nil || isNil(o.Snssais) {
		return nil, false
	}
	return o.Snssais, true
}

// HasSnssais returns a boolean if a field has been set.
func (o *EventFilter) HasSnssais() bool {
	if o != nil && !isNil(o.Snssais) {
		return true
	}

	return false
}

// SetSnssais gets a reference to the given []Snssai and assigns it to the Snssais field.
func (o *EventFilter) SetSnssais(v []Snssai) {
	o.Snssais = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventFilter) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventFilter) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetDnns returns the Dnns field value if set, zero value otherwise.
func (o *EventFilter) GetDnns() []string {
	if o == nil || isNil(o.Dnns) {
		var ret []string
		return ret
	}
	return o.Dnns
}

// GetDnnsOk returns a tuple with the Dnns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDnnsOk() ([]string, bool) {
	if o == nil || isNil(o.Dnns) {
		return nil, false
	}
	return o.Dnns, true
}

// HasDnns returns a boolean if a field has been set.
func (o *EventFilter) HasDnns() bool {
	if o != nil && !isNil(o.Dnns) {
		return true
	}

	return false
}

// SetDnns gets a reference to the given []string and assigns it to the Dnns field.
func (o *EventFilter) SetDnns(v []string) {
	o.Dnns = v
}

// GetDnais returns the Dnais field value if set, zero value otherwise.
func (o *EventFilter) GetDnais() []string {
	if o == nil || isNil(o.Dnais) {
		var ret []string
		return ret
	}
	return o.Dnais
}

// GetDnaisOk returns a tuple with the Dnais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDnaisOk() ([]string, bool) {
	if o == nil || isNil(o.Dnais) {
		return nil, false
	}
	return o.Dnais, true
}

// HasDnais returns a boolean if a field has been set.
func (o *EventFilter) HasDnais() bool {
	if o != nil && !isNil(o.Dnais) {
		return true
	}

	return false
}

// SetDnais gets a reference to the given []string and assigns it to the Dnais field.
func (o *EventFilter) SetDnais(v []string) {
	o.Dnais = v
}

// GetLadnDnns returns the LadnDnns field value if set, zero value otherwise.
func (o *EventFilter) GetLadnDnns() []string {
	if o == nil || isNil(o.LadnDnns) {
		var ret []string
		return ret
	}
	return o.LadnDnns
}

// GetLadnDnnsOk returns a tuple with the LadnDnns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetLadnDnnsOk() ([]string, bool) {
	if o == nil || isNil(o.LadnDnns) {
		return nil, false
	}
	return o.LadnDnns, true
}

// HasLadnDnns returns a boolean if a field has been set.
func (o *EventFilter) HasLadnDnns() bool {
	if o != nil && !isNil(o.LadnDnns) {
		return true
	}

	return false
}

// SetLadnDnns gets a reference to the given []string and assigns it to the LadnDnns field.
func (o *EventFilter) SetLadnDnns(v []string) {
	o.LadnDnns = v
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *EventFilter) GetNetworkArea() NetworkAreaInfo {
	if o == nil || isNil(o.NetworkArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.NetworkArea) {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *EventFilter) HasNetworkArea() bool {
	if o != nil && !isNil(o.NetworkArea) {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *EventFilter) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetVisitedAreas returns the VisitedAreas field value if set, zero value otherwise.
func (o *EventFilter) GetVisitedAreas() []NetworkAreaInfo {
	if o == nil || isNil(o.VisitedAreas) {
		var ret []NetworkAreaInfo
		return ret
	}
	return o.VisitedAreas
}

// GetVisitedAreasOk returns a tuple with the VisitedAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetVisitedAreasOk() ([]NetworkAreaInfo, bool) {
	if o == nil || isNil(o.VisitedAreas) {
		return nil, false
	}
	return o.VisitedAreas, true
}

// HasVisitedAreas returns a boolean if a field has been set.
func (o *EventFilter) HasVisitedAreas() bool {
	if o != nil && !isNil(o.VisitedAreas) {
		return true
	}

	return false
}

// SetVisitedAreas gets a reference to the given []NetworkAreaInfo and assigns it to the VisitedAreas field.
func (o *EventFilter) SetVisitedAreas(v []NetworkAreaInfo) {
	o.VisitedAreas = v
}

// GetMaxTopAppUlNbr returns the MaxTopAppUlNbr field value if set, zero value otherwise.
func (o *EventFilter) GetMaxTopAppUlNbr() int32 {
	if o == nil || isNil(o.MaxTopAppUlNbr) {
		var ret int32
		return ret
	}
	return *o.MaxTopAppUlNbr
}

// GetMaxTopAppUlNbrOk returns a tuple with the MaxTopAppUlNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetMaxTopAppUlNbrOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTopAppUlNbr) {
		return nil, false
	}
	return o.MaxTopAppUlNbr, true
}

// HasMaxTopAppUlNbr returns a boolean if a field has been set.
func (o *EventFilter) HasMaxTopAppUlNbr() bool {
	if o != nil && !isNil(o.MaxTopAppUlNbr) {
		return true
	}

	return false
}

// SetMaxTopAppUlNbr gets a reference to the given int32 and assigns it to the MaxTopAppUlNbr field.
func (o *EventFilter) SetMaxTopAppUlNbr(v int32) {
	o.MaxTopAppUlNbr = &v
}

// GetMaxTopAppDlNbr returns the MaxTopAppDlNbr field value if set, zero value otherwise.
func (o *EventFilter) GetMaxTopAppDlNbr() int32 {
	if o == nil || isNil(o.MaxTopAppDlNbr) {
		var ret int32
		return ret
	}
	return *o.MaxTopAppDlNbr
}

// GetMaxTopAppDlNbrOk returns a tuple with the MaxTopAppDlNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetMaxTopAppDlNbrOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTopAppDlNbr) {
		return nil, false
	}
	return o.MaxTopAppDlNbr, true
}

// HasMaxTopAppDlNbr returns a boolean if a field has been set.
func (o *EventFilter) HasMaxTopAppDlNbr() bool {
	if o != nil && !isNil(o.MaxTopAppDlNbr) {
		return true
	}

	return false
}

// SetMaxTopAppDlNbr gets a reference to the given int32 and assigns it to the MaxTopAppDlNbr field.
func (o *EventFilter) SetMaxTopAppDlNbr(v int32) {
	o.MaxTopAppDlNbr = &v
}

// GetNfInstanceIds returns the NfInstanceIds field value if set, zero value otherwise.
func (o *EventFilter) GetNfInstanceIds() []string {
	if o == nil || isNil(o.NfInstanceIds) {
		var ret []string
		return ret
	}
	return o.NfInstanceIds
}

// GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfInstanceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.NfInstanceIds) {
		return nil, false
	}
	return o.NfInstanceIds, true
}

// HasNfInstanceIds returns a boolean if a field has been set.
func (o *EventFilter) HasNfInstanceIds() bool {
	if o != nil && !isNil(o.NfInstanceIds) {
		return true
	}

	return false
}

// SetNfInstanceIds gets a reference to the given []string and assigns it to the NfInstanceIds field.
func (o *EventFilter) SetNfInstanceIds(v []string) {
	o.NfInstanceIds = v
}

// GetNfSetIds returns the NfSetIds field value if set, zero value otherwise.
func (o *EventFilter) GetNfSetIds() []string {
	if o == nil || isNil(o.NfSetIds) {
		var ret []string
		return ret
	}
	return o.NfSetIds
}

// GetNfSetIdsOk returns a tuple with the NfSetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfSetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.NfSetIds) {
		return nil, false
	}
	return o.NfSetIds, true
}

// HasNfSetIds returns a boolean if a field has been set.
func (o *EventFilter) HasNfSetIds() bool {
	if o != nil && !isNil(o.NfSetIds) {
		return true
	}

	return false
}

// SetNfSetIds gets a reference to the given []string and assigns it to the NfSetIds field.
func (o *EventFilter) SetNfSetIds(v []string) {
	o.NfSetIds = v
}

// GetNfTypes returns the NfTypes field value if set, zero value otherwise.
func (o *EventFilter) GetNfTypes() []NFType {
	if o == nil || isNil(o.NfTypes) {
		var ret []NFType
		return ret
	}
	return o.NfTypes
}

// GetNfTypesOk returns a tuple with the NfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfTypesOk() ([]NFType, bool) {
	if o == nil || isNil(o.NfTypes) {
		return nil, false
	}
	return o.NfTypes, true
}

// HasNfTypes returns a boolean if a field has been set.
func (o *EventFilter) HasNfTypes() bool {
	if o != nil && !isNil(o.NfTypes) {
		return true
	}

	return false
}

// SetNfTypes gets a reference to the given []NFType and assigns it to the NfTypes field.
func (o *EventFilter) SetNfTypes(v []NFType) {
	o.NfTypes = v
}

// GetNsiIdInfos returns the NsiIdInfos field value if set, zero value otherwise.
func (o *EventFilter) GetNsiIdInfos() []NsiIdInfo {
	if o == nil || isNil(o.NsiIdInfos) {
		var ret []NsiIdInfo
		return ret
	}
	return o.NsiIdInfos
}

// GetNsiIdInfosOk returns a tuple with the NsiIdInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNsiIdInfosOk() ([]NsiIdInfo, bool) {
	if o == nil || isNil(o.NsiIdInfos) {
		return nil, false
	}
	return o.NsiIdInfos, true
}

// HasNsiIdInfos returns a boolean if a field has been set.
func (o *EventFilter) HasNsiIdInfos() bool {
	if o != nil && !isNil(o.NsiIdInfos) {
		return true
	}

	return false
}

// SetNsiIdInfos gets a reference to the given []NsiIdInfo and assigns it to the NsiIdInfos field.
func (o *EventFilter) SetNsiIdInfos(v []NsiIdInfo) {
	o.NsiIdInfos = v
}

// GetQosRequ returns the QosRequ field value if set, zero value otherwise.
func (o *EventFilter) GetQosRequ() QosRequirement {
	if o == nil || isNil(o.QosRequ) {
		var ret QosRequirement
		return ret
	}
	return *o.QosRequ
}

// GetQosRequOk returns a tuple with the QosRequ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetQosRequOk() (*QosRequirement, bool) {
	if o == nil || isNil(o.QosRequ) {
		return nil, false
	}
	return o.QosRequ, true
}

// HasQosRequ returns a boolean if a field has been set.
func (o *EventFilter) HasQosRequ() bool {
	if o != nil && !isNil(o.QosRequ) {
		return true
	}

	return false
}

// SetQosRequ gets a reference to the given QosRequirement and assigns it to the QosRequ field.
func (o *EventFilter) SetQosRequ(v QosRequirement) {
	o.QosRequ = &v
}

// GetNwPerfTypes returns the NwPerfTypes field value if set, zero value otherwise.
func (o *EventFilter) GetNwPerfTypes() []NetworkPerfType {
	if o == nil || isNil(o.NwPerfTypes) {
		var ret []NetworkPerfType
		return ret
	}
	return o.NwPerfTypes
}

// GetNwPerfTypesOk returns a tuple with the NwPerfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNwPerfTypesOk() ([]NetworkPerfType, bool) {
	if o == nil || isNil(o.NwPerfTypes) {
		return nil, false
	}
	return o.NwPerfTypes, true
}

// HasNwPerfTypes returns a boolean if a field has been set.
func (o *EventFilter) HasNwPerfTypes() bool {
	if o != nil && !isNil(o.NwPerfTypes) {
		return true
	}

	return false
}

// SetNwPerfTypes gets a reference to the given []NetworkPerfType and assigns it to the NwPerfTypes field.
func (o *EventFilter) SetNwPerfTypes(v []NetworkPerfType) {
	o.NwPerfTypes = v
}

// GetBwRequs returns the BwRequs field value if set, zero value otherwise.
func (o *EventFilter) GetBwRequs() []BwRequirement {
	if o == nil || isNil(o.BwRequs) {
		var ret []BwRequirement
		return ret
	}
	return o.BwRequs
}

// GetBwRequsOk returns a tuple with the BwRequs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetBwRequsOk() ([]BwRequirement, bool) {
	if o == nil || isNil(o.BwRequs) {
		return nil, false
	}
	return o.BwRequs, true
}

// HasBwRequs returns a boolean if a field has been set.
func (o *EventFilter) HasBwRequs() bool {
	if o != nil && !isNil(o.BwRequs) {
		return true
	}

	return false
}

// SetBwRequs gets a reference to the given []BwRequirement and assigns it to the BwRequs field.
func (o *EventFilter) SetBwRequs(v []BwRequirement) {
	o.BwRequs = v
}

// GetExcepIds returns the ExcepIds field value if set, zero value otherwise.
func (o *EventFilter) GetExcepIds() []ExceptionId {
	if o == nil || isNil(o.ExcepIds) {
		var ret []ExceptionId
		return ret
	}
	return o.ExcepIds
}

// GetExcepIdsOk returns a tuple with the ExcepIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExcepIdsOk() ([]ExceptionId, bool) {
	if o == nil || isNil(o.ExcepIds) {
		return nil, false
	}
	return o.ExcepIds, true
}

// HasExcepIds returns a boolean if a field has been set.
func (o *EventFilter) HasExcepIds() bool {
	if o != nil && !isNil(o.ExcepIds) {
		return true
	}

	return false
}

// SetExcepIds gets a reference to the given []ExceptionId and assigns it to the ExcepIds field.
func (o *EventFilter) SetExcepIds(v []ExceptionId) {
	o.ExcepIds = v
}

// GetExptAnaType returns the ExptAnaType field value if set, zero value otherwise.
func (o *EventFilter) GetExptAnaType() ExpectedAnalyticsType {
	if o == nil || isNil(o.ExptAnaType) {
		var ret ExpectedAnalyticsType
		return ret
	}
	return *o.ExptAnaType
}

// GetExptAnaTypeOk returns a tuple with the ExptAnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool) {
	if o == nil || isNil(o.ExptAnaType) {
		return nil, false
	}
	return o.ExptAnaType, true
}

// HasExptAnaType returns a boolean if a field has been set.
func (o *EventFilter) HasExptAnaType() bool {
	if o != nil && !isNil(o.ExptAnaType) {
		return true
	}

	return false
}

// SetExptAnaType gets a reference to the given ExpectedAnalyticsType and assigns it to the ExptAnaType field.
func (o *EventFilter) SetExptAnaType(v ExpectedAnalyticsType) {
	o.ExptAnaType = &v
}

// GetExptUeBehav returns the ExptUeBehav field value if set, zero value otherwise.
func (o *EventFilter) GetExptUeBehav() ExpectedUeBehaviourData {
	if o == nil || isNil(o.ExptUeBehav) {
		var ret ExpectedUeBehaviourData
		return ret
	}
	return *o.ExptUeBehav
}

// GetExptUeBehavOk returns a tuple with the ExptUeBehav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool) {
	if o == nil || isNil(o.ExptUeBehav) {
		return nil, false
	}
	return o.ExptUeBehav, true
}

// HasExptUeBehav returns a boolean if a field has been set.
func (o *EventFilter) HasExptUeBehav() bool {
	if o != nil && !isNil(o.ExptUeBehav) {
		return true
	}

	return false
}

// SetExptUeBehav gets a reference to the given ExpectedUeBehaviourData and assigns it to the ExptUeBehav field.
func (o *EventFilter) SetExptUeBehav(v ExpectedUeBehaviourData) {
	o.ExptUeBehav = &v
}

// GetRatFreqs returns the RatFreqs field value if set, zero value otherwise.
func (o *EventFilter) GetRatFreqs() []RatFreqInformation {
	if o == nil || isNil(o.RatFreqs) {
		var ret []RatFreqInformation
		return ret
	}
	return o.RatFreqs
}

// GetRatFreqsOk returns a tuple with the RatFreqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetRatFreqsOk() ([]RatFreqInformation, bool) {
	if o == nil || isNil(o.RatFreqs) {
		return nil, false
	}
	return o.RatFreqs, true
}

// HasRatFreqs returns a boolean if a field has been set.
func (o *EventFilter) HasRatFreqs() bool {
	if o != nil && !isNil(o.RatFreqs) {
		return true
	}

	return false
}

// SetRatFreqs gets a reference to the given []RatFreqInformation and assigns it to the RatFreqs field.
func (o *EventFilter) SetRatFreqs(v []RatFreqInformation) {
	o.RatFreqs = v
}

// GetDisperReqs returns the DisperReqs field value if set, zero value otherwise.
func (o *EventFilter) GetDisperReqs() []DispersionRequirement {
	if o == nil || isNil(o.DisperReqs) {
		var ret []DispersionRequirement
		return ret
	}
	return o.DisperReqs
}

// GetDisperReqsOk returns a tuple with the DisperReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDisperReqsOk() ([]DispersionRequirement, bool) {
	if o == nil || isNil(o.DisperReqs) {
		return nil, false
	}
	return o.DisperReqs, true
}

// HasDisperReqs returns a boolean if a field has been set.
func (o *EventFilter) HasDisperReqs() bool {
	if o != nil && !isNil(o.DisperReqs) {
		return true
	}

	return false
}

// SetDisperReqs gets a reference to the given []DispersionRequirement and assigns it to the DisperReqs field.
func (o *EventFilter) SetDisperReqs(v []DispersionRequirement) {
	o.DisperReqs = v
}

// GetRedTransReqs returns the RedTransReqs field value if set, zero value otherwise.
func (o *EventFilter) GetRedTransReqs() []RedundantTransmissionExpReq {
	if o == nil || isNil(o.RedTransReqs) {
		var ret []RedundantTransmissionExpReq
		return ret
	}
	return o.RedTransReqs
}

// GetRedTransReqsOk returns a tuple with the RedTransReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetRedTransReqsOk() ([]RedundantTransmissionExpReq, bool) {
	if o == nil || isNil(o.RedTransReqs) {
		return nil, false
	}
	return o.RedTransReqs, true
}

// HasRedTransReqs returns a boolean if a field has been set.
func (o *EventFilter) HasRedTransReqs() bool {
	if o != nil && !isNil(o.RedTransReqs) {
		return true
	}

	return false
}

// SetRedTransReqs gets a reference to the given []RedundantTransmissionExpReq and assigns it to the RedTransReqs field.
func (o *EventFilter) SetRedTransReqs(v []RedundantTransmissionExpReq) {
	o.RedTransReqs = v
}

// GetWlanReqs returns the WlanReqs field value if set, zero value otherwise.
func (o *EventFilter) GetWlanReqs() []WlanPerformanceReq {
	if o == nil || isNil(o.WlanReqs) {
		var ret []WlanPerformanceReq
		return ret
	}
	return o.WlanReqs
}

// GetWlanReqsOk returns a tuple with the WlanReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetWlanReqsOk() ([]WlanPerformanceReq, bool) {
	if o == nil || isNil(o.WlanReqs) {
		return nil, false
	}
	return o.WlanReqs, true
}

// HasWlanReqs returns a boolean if a field has been set.
func (o *EventFilter) HasWlanReqs() bool {
	if o != nil && !isNil(o.WlanReqs) {
		return true
	}

	return false
}

// SetWlanReqs gets a reference to the given []WlanPerformanceReq and assigns it to the WlanReqs field.
func (o *EventFilter) SetWlanReqs(v []WlanPerformanceReq) {
	o.WlanReqs = v
}

// GetListOfAnaSubsets returns the ListOfAnaSubsets field value if set, zero value otherwise.
func (o *EventFilter) GetListOfAnaSubsets() []AnalyticsSubset {
	if o == nil || isNil(o.ListOfAnaSubsets) {
		var ret []AnalyticsSubset
		return ret
	}
	return o.ListOfAnaSubsets
}

// GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetListOfAnaSubsetsOk() ([]AnalyticsSubset, bool) {
	if o == nil || isNil(o.ListOfAnaSubsets) {
		return nil, false
	}
	return o.ListOfAnaSubsets, true
}

// HasListOfAnaSubsets returns a boolean if a field has been set.
func (o *EventFilter) HasListOfAnaSubsets() bool {
	if o != nil && !isNil(o.ListOfAnaSubsets) {
		return true
	}

	return false
}

// SetListOfAnaSubsets gets a reference to the given []AnalyticsSubset and assigns it to the ListOfAnaSubsets field.
func (o *EventFilter) SetListOfAnaSubsets(v []AnalyticsSubset) {
	o.ListOfAnaSubsets = v
}

// GetUpfInfo returns the UpfInfo field value if set, zero value otherwise.
func (o *EventFilter) GetUpfInfo() UpfInformation {
	if o == nil || isNil(o.UpfInfo) {
		var ret UpfInformation
		return ret
	}
	return *o.UpfInfo
}

// GetUpfInfoOk returns a tuple with the UpfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetUpfInfoOk() (*UpfInformation, bool) {
	if o == nil || isNil(o.UpfInfo) {
		return nil, false
	}
	return o.UpfInfo, true
}

// HasUpfInfo returns a boolean if a field has been set.
func (o *EventFilter) HasUpfInfo() bool {
	if o != nil && !isNil(o.UpfInfo) {
		return true
	}

	return false
}

// SetUpfInfo gets a reference to the given UpfInformation and assigns it to the UpfInfo field.
func (o *EventFilter) SetUpfInfo(v UpfInformation) {
	o.UpfInfo = &v
}

// GetAppServerAddrs returns the AppServerAddrs field value if set, zero value otherwise.
func (o *EventFilter) GetAppServerAddrs() []AddrFqdn {
	if o == nil || isNil(o.AppServerAddrs) {
		var ret []AddrFqdn
		return ret
	}
	return o.AppServerAddrs
}

// GetAppServerAddrsOk returns a tuple with the AppServerAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAppServerAddrsOk() ([]AddrFqdn, bool) {
	if o == nil || isNil(o.AppServerAddrs) {
		return nil, false
	}
	return o.AppServerAddrs, true
}

// HasAppServerAddrs returns a boolean if a field has been set.
func (o *EventFilter) HasAppServerAddrs() bool {
	if o != nil && !isNil(o.AppServerAddrs) {
		return true
	}

	return false
}

// SetAppServerAddrs gets a reference to the given []AddrFqdn and assigns it to the AppServerAddrs field.
func (o *EventFilter) SetAppServerAddrs(v []AddrFqdn) {
	o.AppServerAddrs = v
}

// GetDnPerfReqs returns the DnPerfReqs field value if set, zero value otherwise.
func (o *EventFilter) GetDnPerfReqs() []DnPerformanceReq {
	if o == nil || isNil(o.DnPerfReqs) {
		var ret []DnPerformanceReq
		return ret
	}
	return o.DnPerfReqs
}

// GetDnPerfReqsOk returns a tuple with the DnPerfReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDnPerfReqsOk() ([]DnPerformanceReq, bool) {
	if o == nil || isNil(o.DnPerfReqs) {
		return nil, false
	}
	return o.DnPerfReqs, true
}

// HasDnPerfReqs returns a boolean if a field has been set.
func (o *EventFilter) HasDnPerfReqs() bool {
	if o != nil && !isNil(o.DnPerfReqs) {
		return true
	}

	return false
}

// SetDnPerfReqs gets a reference to the given []DnPerformanceReq and assigns it to the DnPerfReqs field.
func (o *EventFilter) SetDnPerfReqs(v []DnPerformanceReq) {
	o.DnPerfReqs = v
}

func (o EventFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnySlice) {
		toSerialize["anySlice"] = o.AnySlice
	}
	if !isNil(o.Snssais) {
		toSerialize["snssais"] = o.Snssais
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.Dnns) {
		toSerialize["dnns"] = o.Dnns
	}
	if !isNil(o.Dnais) {
		toSerialize["dnais"] = o.Dnais
	}
	if !isNil(o.LadnDnns) {
		toSerialize["ladnDnns"] = o.LadnDnns
	}
	if !isNil(o.NetworkArea) {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if !isNil(o.VisitedAreas) {
		toSerialize["visitedAreas"] = o.VisitedAreas
	}
	if !isNil(o.MaxTopAppUlNbr) {
		toSerialize["maxTopAppUlNbr"] = o.MaxTopAppUlNbr
	}
	if !isNil(o.MaxTopAppDlNbr) {
		toSerialize["maxTopAppDlNbr"] = o.MaxTopAppDlNbr
	}
	if !isNil(o.NfInstanceIds) {
		toSerialize["nfInstanceIds"] = o.NfInstanceIds
	}
	if !isNil(o.NfSetIds) {
		toSerialize["nfSetIds"] = o.NfSetIds
	}
	if !isNil(o.NfTypes) {
		toSerialize["nfTypes"] = o.NfTypes
	}
	if !isNil(o.NsiIdInfos) {
		toSerialize["nsiIdInfos"] = o.NsiIdInfos
	}
	if !isNil(o.QosRequ) {
		toSerialize["qosRequ"] = o.QosRequ
	}
	if !isNil(o.NwPerfTypes) {
		toSerialize["nwPerfTypes"] = o.NwPerfTypes
	}
	if !isNil(o.BwRequs) {
		toSerialize["bwRequs"] = o.BwRequs
	}
	if !isNil(o.ExcepIds) {
		toSerialize["excepIds"] = o.ExcepIds
	}
	if !isNil(o.ExptAnaType) {
		toSerialize["exptAnaType"] = o.ExptAnaType
	}
	if !isNil(o.ExptUeBehav) {
		toSerialize["exptUeBehav"] = o.ExptUeBehav
	}
	if !isNil(o.RatFreqs) {
		toSerialize["ratFreqs"] = o.RatFreqs
	}
	if !isNil(o.DisperReqs) {
		toSerialize["disperReqs"] = o.DisperReqs
	}
	if !isNil(o.RedTransReqs) {
		toSerialize["redTransReqs"] = o.RedTransReqs
	}
	if !isNil(o.WlanReqs) {
		toSerialize["wlanReqs"] = o.WlanReqs
	}
	if !isNil(o.ListOfAnaSubsets) {
		toSerialize["listOfAnaSubsets"] = o.ListOfAnaSubsets
	}
	if !isNil(o.UpfInfo) {
		toSerialize["upfInfo"] = o.UpfInfo
	}
	if !isNil(o.AppServerAddrs) {
		toSerialize["appServerAddrs"] = o.AppServerAddrs
	}
	if !isNil(o.DnPerfReqs) {
		toSerialize["dnPerfReqs"] = o.DnPerfReqs
	}
	return toSerialize, nil
}

type NullableEventFilter struct {
	value *EventFilter
	isSet bool
}

func (v NullableEventFilter) Get() *EventFilter {
	return v.value
}

func (v *NullableEventFilter) Set(val *EventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter(val *EventFilter) *NullableEventFilter {
	return &NullableEventFilter{value: val, isSet: true}
}

func (v NullableEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


