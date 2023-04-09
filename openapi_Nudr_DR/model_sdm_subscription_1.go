/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"time"
)

// checks if the SdmSubscription1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdmSubscription1{}

// SdmSubscription1 struct for SdmSubscription1
type SdmSubscription1 struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfInstanceId        string `json:"nfInstanceId"`
	ImplicitUnsubscribe *bool  `json:"implicitUnsubscribe,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference     string       `json:"callbackReference"`
	AmfServiceName        *ServiceName `json:"amfServiceName,omitempty"`
	MonitoredResourceUris []string     `json:"monitoredResourceUris"`
	SingleNssai           *Snssai      `json:"singleNssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn             *string          `json:"dnn,omitempty"`
	SubscriptionId  *string          `json:"subscriptionId,omitempty"`
	PlmnId          *PlmnId1         `json:"plmnId,omitempty"`
	ImmediateReport *bool            `json:"immediateReport,omitempty"`
	Report          *ImmediateReport `json:"report,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures     *string                       `json:"supportedFeatures,omitempty"`
	ContextInfo           *ContextInfo                  `json:"contextInfo,omitempty"`
	NfChangeFilter        *bool                         `json:"nfChangeFilter,omitempty"`
	UniqueSubscription    *bool                         `json:"uniqueSubscription,omitempty"`
	ResetIds              []string                      `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter *UeContextInSmfDataSubFilter1 `json:"ueConSmfDataSubFilter,omitempty"`
	AdjacentPlmns         []PlmnId1                     `json:"adjacentPlmns,omitempty"`
	DisasterRoamingInd    *bool                         `json:"disasterRoamingInd,omitempty"`
}

// NewSdmSubscription1 instantiates a new SdmSubscription1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdmSubscription1(nfInstanceId string, callbackReference string, monitoredResourceUris []string) *SdmSubscription1 {
	this := SdmSubscription1{}
	this.NfInstanceId = nfInstanceId
	this.CallbackReference = callbackReference
	this.MonitoredResourceUris = monitoredResourceUris
	var immediateReport bool = false
	this.ImmediateReport = &immediateReport
	var nfChangeFilter bool = false
	this.NfChangeFilter = &nfChangeFilter
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	return &this
}

// NewSdmSubscription1WithDefaults instantiates a new SdmSubscription1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdmSubscription1WithDefaults() *SdmSubscription1 {
	this := SdmSubscription1{}
	var immediateReport bool = false
	this.ImmediateReport = &immediateReport
	var nfChangeFilter bool = false
	this.NfChangeFilter = &nfChangeFilter
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	return &this
}

// GetNfInstanceId returns the NfInstanceId field value
func (o *SdmSubscription1) GetNfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetNfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfInstanceId, true
}

// SetNfInstanceId sets field value
func (o *SdmSubscription1) SetNfInstanceId(v string) {
	o.NfInstanceId = v
}

// GetImplicitUnsubscribe returns the ImplicitUnsubscribe field value if set, zero value otherwise.
func (o *SdmSubscription1) GetImplicitUnsubscribe() bool {
	if o == nil || IsNil(o.ImplicitUnsubscribe) {
		var ret bool
		return ret
	}
	return *o.ImplicitUnsubscribe
}

// GetImplicitUnsubscribeOk returns a tuple with the ImplicitUnsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetImplicitUnsubscribeOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitUnsubscribe) {
		return nil, false
	}
	return o.ImplicitUnsubscribe, true
}

// HasImplicitUnsubscribe returns a boolean if a field has been set.
func (o *SdmSubscription1) HasImplicitUnsubscribe() bool {
	if o != nil && !IsNil(o.ImplicitUnsubscribe) {
		return true
	}

	return false
}

// SetImplicitUnsubscribe gets a reference to the given bool and assigns it to the ImplicitUnsubscribe field.
func (o *SdmSubscription1) SetImplicitUnsubscribe(v bool) {
	o.ImplicitUnsubscribe = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *SdmSubscription1) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *SdmSubscription1) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *SdmSubscription1) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetCallbackReference returns the CallbackReference field value
func (o *SdmSubscription1) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *SdmSubscription1) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetAmfServiceName returns the AmfServiceName field value if set, zero value otherwise.
func (o *SdmSubscription1) GetAmfServiceName() ServiceName {
	if o == nil || IsNil(o.AmfServiceName) {
		var ret ServiceName
		return ret
	}
	return *o.AmfServiceName
}

// GetAmfServiceNameOk returns a tuple with the AmfServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetAmfServiceNameOk() (*ServiceName, bool) {
	if o == nil || IsNil(o.AmfServiceName) {
		return nil, false
	}
	return o.AmfServiceName, true
}

// HasAmfServiceName returns a boolean if a field has been set.
func (o *SdmSubscription1) HasAmfServiceName() bool {
	if o != nil && !IsNil(o.AmfServiceName) {
		return true
	}

	return false
}

// SetAmfServiceName gets a reference to the given ServiceName and assigns it to the AmfServiceName field.
func (o *SdmSubscription1) SetAmfServiceName(v ServiceName) {
	o.AmfServiceName = &v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value
func (o *SdmSubscription1) GetMonitoredResourceUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoredResourceUris, true
}

// SetMonitoredResourceUris sets field value
func (o *SdmSubscription1) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *SdmSubscription1) GetSingleNssai() Snssai {
	if o == nil || IsNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SingleNssai) {
		return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *SdmSubscription1) HasSingleNssai() bool {
	if o != nil && !IsNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *SdmSubscription1) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SdmSubscription1) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SdmSubscription1) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SdmSubscription1) SetDnn(v string) {
	o.Dnn = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SdmSubscription1) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SdmSubscription1) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SdmSubscription1) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *SdmSubscription1) GetPlmnId() PlmnId1 {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId1
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetPlmnIdOk() (*PlmnId1, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *SdmSubscription1) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId1 and assigns it to the PlmnId field.
func (o *SdmSubscription1) SetPlmnId(v PlmnId1) {
	o.PlmnId = &v
}

// GetImmediateReport returns the ImmediateReport field value if set, zero value otherwise.
func (o *SdmSubscription1) GetImmediateReport() bool {
	if o == nil || IsNil(o.ImmediateReport) {
		var ret bool
		return ret
	}
	return *o.ImmediateReport
}

// GetImmediateReportOk returns a tuple with the ImmediateReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetImmediateReportOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmediateReport) {
		return nil, false
	}
	return o.ImmediateReport, true
}

// HasImmediateReport returns a boolean if a field has been set.
func (o *SdmSubscription1) HasImmediateReport() bool {
	if o != nil && !IsNil(o.ImmediateReport) {
		return true
	}

	return false
}

// SetImmediateReport gets a reference to the given bool and assigns it to the ImmediateReport field.
func (o *SdmSubscription1) SetImmediateReport(v bool) {
	o.ImmediateReport = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *SdmSubscription1) GetReport() ImmediateReport {
	if o == nil || IsNil(o.Report) {
		var ret ImmediateReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetReportOk() (*ImmediateReport, bool) {
	if o == nil || IsNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *SdmSubscription1) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given ImmediateReport and assigns it to the Report field.
func (o *SdmSubscription1) SetReport(v ImmediateReport) {
	o.Report = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SdmSubscription1) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SdmSubscription1) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SdmSubscription1) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *SdmSubscription1) GetContextInfo() ContextInfo {
	if o == nil || IsNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || IsNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *SdmSubscription1) HasContextInfo() bool {
	if o != nil && !IsNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *SdmSubscription1) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetNfChangeFilter returns the NfChangeFilter field value if set, zero value otherwise.
func (o *SdmSubscription1) GetNfChangeFilter() bool {
	if o == nil || IsNil(o.NfChangeFilter) {
		var ret bool
		return ret
	}
	return *o.NfChangeFilter
}

// GetNfChangeFilterOk returns a tuple with the NfChangeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetNfChangeFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.NfChangeFilter) {
		return nil, false
	}
	return o.NfChangeFilter, true
}

// HasNfChangeFilter returns a boolean if a field has been set.
func (o *SdmSubscription1) HasNfChangeFilter() bool {
	if o != nil && !IsNil(o.NfChangeFilter) {
		return true
	}

	return false
}

// SetNfChangeFilter gets a reference to the given bool and assigns it to the NfChangeFilter field.
func (o *SdmSubscription1) SetNfChangeFilter(v bool) {
	o.NfChangeFilter = &v
}

// GetUniqueSubscription returns the UniqueSubscription field value if set, zero value otherwise.
func (o *SdmSubscription1) GetUniqueSubscription() bool {
	if o == nil || IsNil(o.UniqueSubscription) {
		var ret bool
		return ret
	}
	return *o.UniqueSubscription
}

// GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetUniqueSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueSubscription) {
		return nil, false
	}
	return o.UniqueSubscription, true
}

// HasUniqueSubscription returns a boolean if a field has been set.
func (o *SdmSubscription1) HasUniqueSubscription() bool {
	if o != nil && !IsNil(o.UniqueSubscription) {
		return true
	}

	return false
}

// SetUniqueSubscription gets a reference to the given bool and assigns it to the UniqueSubscription field.
func (o *SdmSubscription1) SetUniqueSubscription(v bool) {
	o.UniqueSubscription = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *SdmSubscription1) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *SdmSubscription1) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *SdmSubscription1) SetResetIds(v []string) {
	o.ResetIds = v
}

// GetUeConSmfDataSubFilter returns the UeConSmfDataSubFilter field value if set, zero value otherwise.
func (o *SdmSubscription1) GetUeConSmfDataSubFilter() UeContextInSmfDataSubFilter1 {
	if o == nil || IsNil(o.UeConSmfDataSubFilter) {
		var ret UeContextInSmfDataSubFilter1
		return ret
	}
	return *o.UeConSmfDataSubFilter
}

// GetUeConSmfDataSubFilterOk returns a tuple with the UeConSmfDataSubFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetUeConSmfDataSubFilterOk() (*UeContextInSmfDataSubFilter1, bool) {
	if o == nil || IsNil(o.UeConSmfDataSubFilter) {
		return nil, false
	}
	return o.UeConSmfDataSubFilter, true
}

// HasUeConSmfDataSubFilter returns a boolean if a field has been set.
func (o *SdmSubscription1) HasUeConSmfDataSubFilter() bool {
	if o != nil && !IsNil(o.UeConSmfDataSubFilter) {
		return true
	}

	return false
}

// SetUeConSmfDataSubFilter gets a reference to the given UeContextInSmfDataSubFilter1 and assigns it to the UeConSmfDataSubFilter field.
func (o *SdmSubscription1) SetUeConSmfDataSubFilter(v UeContextInSmfDataSubFilter1) {
	o.UeConSmfDataSubFilter = &v
}

// GetAdjacentPlmns returns the AdjacentPlmns field value if set, zero value otherwise.
func (o *SdmSubscription1) GetAdjacentPlmns() []PlmnId1 {
	if o == nil || IsNil(o.AdjacentPlmns) {
		var ret []PlmnId1
		return ret
	}
	return o.AdjacentPlmns
}

// GetAdjacentPlmnsOk returns a tuple with the AdjacentPlmns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetAdjacentPlmnsOk() ([]PlmnId1, bool) {
	if o == nil || IsNil(o.AdjacentPlmns) {
		return nil, false
	}
	return o.AdjacentPlmns, true
}

// HasAdjacentPlmns returns a boolean if a field has been set.
func (o *SdmSubscription1) HasAdjacentPlmns() bool {
	if o != nil && !IsNil(o.AdjacentPlmns) {
		return true
	}

	return false
}

// SetAdjacentPlmns gets a reference to the given []PlmnId1 and assigns it to the AdjacentPlmns field.
func (o *SdmSubscription1) SetAdjacentPlmns(v []PlmnId1) {
	o.AdjacentPlmns = v
}

// GetDisasterRoamingInd returns the DisasterRoamingInd field value if set, zero value otherwise.
func (o *SdmSubscription1) GetDisasterRoamingInd() bool {
	if o == nil || IsNil(o.DisasterRoamingInd) {
		var ret bool
		return ret
	}
	return *o.DisasterRoamingInd
}

// GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubscription1) GetDisasterRoamingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DisasterRoamingInd) {
		return nil, false
	}
	return o.DisasterRoamingInd, true
}

// HasDisasterRoamingInd returns a boolean if a field has been set.
func (o *SdmSubscription1) HasDisasterRoamingInd() bool {
	if o != nil && !IsNil(o.DisasterRoamingInd) {
		return true
	}

	return false
}

// SetDisasterRoamingInd gets a reference to the given bool and assigns it to the DisasterRoamingInd field.
func (o *SdmSubscription1) SetDisasterRoamingInd(v bool) {
	o.DisasterRoamingInd = &v
}

func (o SdmSubscription1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdmSubscription1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfInstanceId"] = o.NfInstanceId
	if !IsNil(o.ImplicitUnsubscribe) {
		toSerialize["implicitUnsubscribe"] = o.ImplicitUnsubscribe
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	toSerialize["callbackReference"] = o.CallbackReference
	if !IsNil(o.AmfServiceName) {
		toSerialize["amfServiceName"] = o.AmfServiceName
	}
	toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	if !IsNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.ImmediateReport) {
		toSerialize["immediateReport"] = o.ImmediateReport
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !IsNil(o.NfChangeFilter) {
		toSerialize["nfChangeFilter"] = o.NfChangeFilter
	}
	if !IsNil(o.UniqueSubscription) {
		toSerialize["uniqueSubscription"] = o.UniqueSubscription
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	if !IsNil(o.UeConSmfDataSubFilter) {
		toSerialize["ueConSmfDataSubFilter"] = o.UeConSmfDataSubFilter
	}
	if !IsNil(o.AdjacentPlmns) {
		toSerialize["adjacentPlmns"] = o.AdjacentPlmns
	}
	if !IsNil(o.DisasterRoamingInd) {
		toSerialize["disasterRoamingInd"] = o.DisasterRoamingInd
	}
	return toSerialize, nil
}

type NullableSdmSubscription1 struct {
	value *SdmSubscription1
	isSet bool
}

func (v NullableSdmSubscription1) Get() *SdmSubscription1 {
	return v.value
}

func (v *NullableSdmSubscription1) Set(val *SdmSubscription1) {
	v.value = val
	v.isSet = true
}

func (v NullableSdmSubscription1) IsSet() bool {
	return v.isSet
}

func (v *NullableSdmSubscription1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdmSubscription1(val *SdmSubscription1) *NullableSdmSubscription1 {
	return &NullableSdmSubscription1{value: val, isSet: true}
}

func (v NullableSdmSubscription1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdmSubscription1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
