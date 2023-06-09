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

// checks if the AnalyticsEventFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsEventFilter{}

// AnalyticsEventFilter Represents analytics event filter information.
type AnalyticsEventFilter struct {
	LocArea *LocationArea5G `json:"locArea,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn              *string                  `json:"dnn,omitempty"`
	Dnais            []string                 `json:"dnais,omitempty"`
	NwPerfTypes      []NetworkPerfType        `json:"nwPerfTypes,omitempty"`
	AppIds           []string                 `json:"appIds,omitempty"`
	ExcepIds         []ExceptionId            `json:"excepIds,omitempty"`
	ExptAnaType      *ExpectedAnalyticsType   `json:"exptAnaType,omitempty"`
	ExptUeBehav      *ExpectedUeBehaviourData `json:"exptUeBehav,omitempty"`
	Snssai           *Snssai                  `json:"snssai,omitempty"`
	NsiIdInfos       []NsiIdInfo              `json:"nsiIdInfos,omitempty"`
	QosReq           *QosRequirement          `json:"qosReq,omitempty"`
	ListOfAnaSubsets []AnalyticsSubset        `json:"listOfAnaSubsets,omitempty"`
	DnPerfReqs       []DnPerformanceReq       `json:"dnPerfReqs,omitempty"`
	BwRequs          []BwRequirement          `json:"bwRequs,omitempty"`
	RatFreqs         []RatFreqInformation     `json:"ratFreqs,omitempty"`
	AppServerAddrs   []AddrFqdn               `json:"appServerAddrs,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTopAppUl *int32 `json:"maxNumOfTopAppUl,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumOfTopAppDl *int32           `json:"maxNumOfTopAppDl,omitempty"`
	VisitedLocAreas  []LocationArea5G `json:"visitedLocAreas,omitempty"`
}

// NewAnalyticsEventFilter instantiates a new AnalyticsEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsEventFilter() *AnalyticsEventFilter {
	this := AnalyticsEventFilter{}
	return &this
}

// NewAnalyticsEventFilterWithDefaults instantiates a new AnalyticsEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsEventFilterWithDefaults() *AnalyticsEventFilter {
	this := AnalyticsEventFilter{}
	return &this
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetLocArea() LocationArea5G {
	if o == nil || IsNil(o.LocArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.LocArea) {
		return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasLocArea() bool {
	if o != nil && !IsNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given LocationArea5G and assigns it to the LocArea field.
func (o *AnalyticsEventFilter) SetLocArea(v LocationArea5G) {
	o.LocArea = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *AnalyticsEventFilter) SetDnn(v string) {
	o.Dnn = &v
}

// GetDnais returns the Dnais field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetDnais() []string {
	if o == nil || IsNil(o.Dnais) {
		var ret []string
		return ret
	}
	return o.Dnais
}

// GetDnaisOk returns a tuple with the Dnais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetDnaisOk() ([]string, bool) {
	if o == nil || IsNil(o.Dnais) {
		return nil, false
	}
	return o.Dnais, true
}

// HasDnais returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasDnais() bool {
	if o != nil && !IsNil(o.Dnais) {
		return true
	}

	return false
}

// SetDnais gets a reference to the given []string and assigns it to the Dnais field.
func (o *AnalyticsEventFilter) SetDnais(v []string) {
	o.Dnais = v
}

// GetNwPerfTypes returns the NwPerfTypes field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetNwPerfTypes() []NetworkPerfType {
	if o == nil || IsNil(o.NwPerfTypes) {
		var ret []NetworkPerfType
		return ret
	}
	return o.NwPerfTypes
}

// GetNwPerfTypesOk returns a tuple with the NwPerfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetNwPerfTypesOk() ([]NetworkPerfType, bool) {
	if o == nil || IsNil(o.NwPerfTypes) {
		return nil, false
	}
	return o.NwPerfTypes, true
}

// HasNwPerfTypes returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasNwPerfTypes() bool {
	if o != nil && !IsNil(o.NwPerfTypes) {
		return true
	}

	return false
}

// SetNwPerfTypes gets a reference to the given []NetworkPerfType and assigns it to the NwPerfTypes field.
func (o *AnalyticsEventFilter) SetNwPerfTypes(v []NetworkPerfType) {
	o.NwPerfTypes = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetAppIds() []string {
	if o == nil || IsNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasAppIds() bool {
	if o != nil && !IsNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *AnalyticsEventFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetExcepIds returns the ExcepIds field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetExcepIds() []ExceptionId {
	if o == nil || IsNil(o.ExcepIds) {
		var ret []ExceptionId
		return ret
	}
	return o.ExcepIds
}

// GetExcepIdsOk returns a tuple with the ExcepIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetExcepIdsOk() ([]ExceptionId, bool) {
	if o == nil || IsNil(o.ExcepIds) {
		return nil, false
	}
	return o.ExcepIds, true
}

// HasExcepIds returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasExcepIds() bool {
	if o != nil && !IsNil(o.ExcepIds) {
		return true
	}

	return false
}

// SetExcepIds gets a reference to the given []ExceptionId and assigns it to the ExcepIds field.
func (o *AnalyticsEventFilter) SetExcepIds(v []ExceptionId) {
	o.ExcepIds = v
}

// GetExptAnaType returns the ExptAnaType field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetExptAnaType() ExpectedAnalyticsType {
	if o == nil || IsNil(o.ExptAnaType) {
		var ret ExpectedAnalyticsType
		return ret
	}
	return *o.ExptAnaType
}

// GetExptAnaTypeOk returns a tuple with the ExptAnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool) {
	if o == nil || IsNil(o.ExptAnaType) {
		return nil, false
	}
	return o.ExptAnaType, true
}

// HasExptAnaType returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasExptAnaType() bool {
	if o != nil && !IsNil(o.ExptAnaType) {
		return true
	}

	return false
}

// SetExptAnaType gets a reference to the given ExpectedAnalyticsType and assigns it to the ExptAnaType field.
func (o *AnalyticsEventFilter) SetExptAnaType(v ExpectedAnalyticsType) {
	o.ExptAnaType = &v
}

// GetExptUeBehav returns the ExptUeBehav field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetExptUeBehav() ExpectedUeBehaviourData {
	if o == nil || IsNil(o.ExptUeBehav) {
		var ret ExpectedUeBehaviourData
		return ret
	}
	return *o.ExptUeBehav
}

// GetExptUeBehavOk returns a tuple with the ExptUeBehav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool) {
	if o == nil || IsNil(o.ExptUeBehav) {
		return nil, false
	}
	return o.ExptUeBehav, true
}

// HasExptUeBehav returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasExptUeBehav() bool {
	if o != nil && !IsNil(o.ExptUeBehav) {
		return true
	}

	return false
}

// SetExptUeBehav gets a reference to the given ExpectedUeBehaviourData and assigns it to the ExptUeBehav field.
func (o *AnalyticsEventFilter) SetExptUeBehav(v ExpectedUeBehaviourData) {
	o.ExptUeBehav = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *AnalyticsEventFilter) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetNsiIdInfos returns the NsiIdInfos field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetNsiIdInfos() []NsiIdInfo {
	if o == nil || IsNil(o.NsiIdInfos) {
		var ret []NsiIdInfo
		return ret
	}
	return o.NsiIdInfos
}

// GetNsiIdInfosOk returns a tuple with the NsiIdInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetNsiIdInfosOk() ([]NsiIdInfo, bool) {
	if o == nil || IsNil(o.NsiIdInfos) {
		return nil, false
	}
	return o.NsiIdInfos, true
}

// HasNsiIdInfos returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasNsiIdInfos() bool {
	if o != nil && !IsNil(o.NsiIdInfos) {
		return true
	}

	return false
}

// SetNsiIdInfos gets a reference to the given []NsiIdInfo and assigns it to the NsiIdInfos field.
func (o *AnalyticsEventFilter) SetNsiIdInfos(v []NsiIdInfo) {
	o.NsiIdInfos = v
}

// GetQosReq returns the QosReq field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetQosReq() QosRequirement {
	if o == nil || IsNil(o.QosReq) {
		var ret QosRequirement
		return ret
	}
	return *o.QosReq
}

// GetQosReqOk returns a tuple with the QosReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetQosReqOk() (*QosRequirement, bool) {
	if o == nil || IsNil(o.QosReq) {
		return nil, false
	}
	return o.QosReq, true
}

// HasQosReq returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasQosReq() bool {
	if o != nil && !IsNil(o.QosReq) {
		return true
	}

	return false
}

// SetQosReq gets a reference to the given QosRequirement and assigns it to the QosReq field.
func (o *AnalyticsEventFilter) SetQosReq(v QosRequirement) {
	o.QosReq = &v
}

// GetListOfAnaSubsets returns the ListOfAnaSubsets field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetListOfAnaSubsets() []AnalyticsSubset {
	if o == nil || IsNil(o.ListOfAnaSubsets) {
		var ret []AnalyticsSubset
		return ret
	}
	return o.ListOfAnaSubsets
}

// GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetListOfAnaSubsetsOk() ([]AnalyticsSubset, bool) {
	if o == nil || IsNil(o.ListOfAnaSubsets) {
		return nil, false
	}
	return o.ListOfAnaSubsets, true
}

// HasListOfAnaSubsets returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasListOfAnaSubsets() bool {
	if o != nil && !IsNil(o.ListOfAnaSubsets) {
		return true
	}

	return false
}

// SetListOfAnaSubsets gets a reference to the given []AnalyticsSubset and assigns it to the ListOfAnaSubsets field.
func (o *AnalyticsEventFilter) SetListOfAnaSubsets(v []AnalyticsSubset) {
	o.ListOfAnaSubsets = v
}

// GetDnPerfReqs returns the DnPerfReqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetDnPerfReqs() []DnPerformanceReq {
	if o == nil || IsNil(o.DnPerfReqs) {
		var ret []DnPerformanceReq
		return ret
	}
	return o.DnPerfReqs
}

// GetDnPerfReqsOk returns a tuple with the DnPerfReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetDnPerfReqsOk() ([]DnPerformanceReq, bool) {
	if o == nil || IsNil(o.DnPerfReqs) {
		return nil, false
	}
	return o.DnPerfReqs, true
}

// HasDnPerfReqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasDnPerfReqs() bool {
	if o != nil && !IsNil(o.DnPerfReqs) {
		return true
	}

	return false
}

// SetDnPerfReqs gets a reference to the given []DnPerformanceReq and assigns it to the DnPerfReqs field.
func (o *AnalyticsEventFilter) SetDnPerfReqs(v []DnPerformanceReq) {
	o.DnPerfReqs = v
}

// GetBwRequs returns the BwRequs field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetBwRequs() []BwRequirement {
	if o == nil || IsNil(o.BwRequs) {
		var ret []BwRequirement
		return ret
	}
	return o.BwRequs
}

// GetBwRequsOk returns a tuple with the BwRequs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetBwRequsOk() ([]BwRequirement, bool) {
	if o == nil || IsNil(o.BwRequs) {
		return nil, false
	}
	return o.BwRequs, true
}

// HasBwRequs returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasBwRequs() bool {
	if o != nil && !IsNil(o.BwRequs) {
		return true
	}

	return false
}

// SetBwRequs gets a reference to the given []BwRequirement and assigns it to the BwRequs field.
func (o *AnalyticsEventFilter) SetBwRequs(v []BwRequirement) {
	o.BwRequs = v
}

// GetRatFreqs returns the RatFreqs field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetRatFreqs() []RatFreqInformation {
	if o == nil || IsNil(o.RatFreqs) {
		var ret []RatFreqInformation
		return ret
	}
	return o.RatFreqs
}

// GetRatFreqsOk returns a tuple with the RatFreqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetRatFreqsOk() ([]RatFreqInformation, bool) {
	if o == nil || IsNil(o.RatFreqs) {
		return nil, false
	}
	return o.RatFreqs, true
}

// HasRatFreqs returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasRatFreqs() bool {
	if o != nil && !IsNil(o.RatFreqs) {
		return true
	}

	return false
}

// SetRatFreqs gets a reference to the given []RatFreqInformation and assigns it to the RatFreqs field.
func (o *AnalyticsEventFilter) SetRatFreqs(v []RatFreqInformation) {
	o.RatFreqs = v
}

// GetAppServerAddrs returns the AppServerAddrs field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetAppServerAddrs() []AddrFqdn {
	if o == nil || IsNil(o.AppServerAddrs) {
		var ret []AddrFqdn
		return ret
	}
	return o.AppServerAddrs
}

// GetAppServerAddrsOk returns a tuple with the AppServerAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetAppServerAddrsOk() ([]AddrFqdn, bool) {
	if o == nil || IsNil(o.AppServerAddrs) {
		return nil, false
	}
	return o.AppServerAddrs, true
}

// HasAppServerAddrs returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasAppServerAddrs() bool {
	if o != nil && !IsNil(o.AppServerAddrs) {
		return true
	}

	return false
}

// SetAppServerAddrs gets a reference to the given []AddrFqdn and assigns it to the AppServerAddrs field.
func (o *AnalyticsEventFilter) SetAppServerAddrs(v []AddrFqdn) {
	o.AppServerAddrs = v
}

// GetMaxNumOfTopAppUl returns the MaxNumOfTopAppUl field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetMaxNumOfTopAppUl() int32 {
	if o == nil || IsNil(o.MaxNumOfTopAppUl) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTopAppUl
}

// GetMaxNumOfTopAppUlOk returns a tuple with the MaxNumOfTopAppUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetMaxNumOfTopAppUlOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfTopAppUl) {
		return nil, false
	}
	return o.MaxNumOfTopAppUl, true
}

// HasMaxNumOfTopAppUl returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasMaxNumOfTopAppUl() bool {
	if o != nil && !IsNil(o.MaxNumOfTopAppUl) {
		return true
	}

	return false
}

// SetMaxNumOfTopAppUl gets a reference to the given int32 and assigns it to the MaxNumOfTopAppUl field.
func (o *AnalyticsEventFilter) SetMaxNumOfTopAppUl(v int32) {
	o.MaxNumOfTopAppUl = &v
}

// GetMaxNumOfTopAppDl returns the MaxNumOfTopAppDl field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetMaxNumOfTopAppDl() int32 {
	if o == nil || IsNil(o.MaxNumOfTopAppDl) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfTopAppDl
}

// GetMaxNumOfTopAppDlOk returns a tuple with the MaxNumOfTopAppDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetMaxNumOfTopAppDlOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfTopAppDl) {
		return nil, false
	}
	return o.MaxNumOfTopAppDl, true
}

// HasMaxNumOfTopAppDl returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasMaxNumOfTopAppDl() bool {
	if o != nil && !IsNil(o.MaxNumOfTopAppDl) {
		return true
	}

	return false
}

// SetMaxNumOfTopAppDl gets a reference to the given int32 and assigns it to the MaxNumOfTopAppDl field.
func (o *AnalyticsEventFilter) SetMaxNumOfTopAppDl(v int32) {
	o.MaxNumOfTopAppDl = &v
}

// GetVisitedLocAreas returns the VisitedLocAreas field value if set, zero value otherwise.
func (o *AnalyticsEventFilter) GetVisitedLocAreas() []LocationArea5G {
	if o == nil || IsNil(o.VisitedLocAreas) {
		var ret []LocationArea5G
		return ret
	}
	return o.VisitedLocAreas
}

// GetVisitedLocAreasOk returns a tuple with the VisitedLocAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventFilter) GetVisitedLocAreasOk() ([]LocationArea5G, bool) {
	if o == nil || IsNil(o.VisitedLocAreas) {
		return nil, false
	}
	return o.VisitedLocAreas, true
}

// HasVisitedLocAreas returns a boolean if a field has been set.
func (o *AnalyticsEventFilter) HasVisitedLocAreas() bool {
	if o != nil && !IsNil(o.VisitedLocAreas) {
		return true
	}

	return false
}

// SetVisitedLocAreas gets a reference to the given []LocationArea5G and assigns it to the VisitedLocAreas field.
func (o *AnalyticsEventFilter) SetVisitedLocAreas(v []LocationArea5G) {
	o.VisitedLocAreas = v
}

func (o AnalyticsEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsEventFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Dnais) {
		toSerialize["dnais"] = o.Dnais
	}
	if !IsNil(o.NwPerfTypes) {
		toSerialize["nwPerfTypes"] = o.NwPerfTypes
	}
	if !IsNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !IsNil(o.ExcepIds) {
		toSerialize["excepIds"] = o.ExcepIds
	}
	if !IsNil(o.ExptAnaType) {
		toSerialize["exptAnaType"] = o.ExptAnaType
	}
	if !IsNil(o.ExptUeBehav) {
		toSerialize["exptUeBehav"] = o.ExptUeBehav
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

type NullableAnalyticsEventFilter struct {
	value *AnalyticsEventFilter
	isSet bool
}

func (v NullableAnalyticsEventFilter) Get() *AnalyticsEventFilter {
	return v.value
}

func (v *NullableAnalyticsEventFilter) Set(val *AnalyticsEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsEventFilter(val *AnalyticsEventFilter) *NullableAnalyticsEventFilter {
	return &NullableAnalyticsEventFilter{value: val, isSet: true}
}

func (v NullableAnalyticsEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
