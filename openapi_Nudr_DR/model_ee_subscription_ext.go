/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the EeSubscriptionExt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EeSubscriptionExt{}

// EeSubscriptionExt struct for EeSubscriptionExt
type EeSubscriptionExt struct {
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations map[string]MonitoringConfiguration1 `json:"monitoringConfigurations"`
	ReportingOptions *ReportingOptions1 `json:"reportingOptions,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
	EpcAppliedInd *bool `json:"epcAppliedInd,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamHost *string `json:"scefDiamHost,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamRealm *string `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SecondCallbackRef *string `json:"secondCallbackRef,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri *string `json:"dataRestorationCallbackUri,omitempty"`
	AmfSubscriptionInfoList []AmfSubscriptionInfo `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo *SmfSubscriptionInfo `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
}

// NewEeSubscriptionExt instantiates a new EeSubscriptionExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEeSubscriptionExt(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration1) *EeSubscriptionExt {
	this := EeSubscriptionExt{}
	this.CallbackReference = callbackReference
	this.MonitoringConfigurations = monitoringConfigurations
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// NewEeSubscriptionExtWithDefaults instantiates a new EeSubscriptionExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEeSubscriptionExtWithDefaults() *EeSubscriptionExt {
	this := EeSubscriptionExt{}
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// GetCallbackReference returns the CallbackReference field value
func (o *EeSubscriptionExt) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *EeSubscriptionExt) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetMonitoringConfigurations returns the MonitoringConfigurations field value
func (o *EeSubscriptionExt) GetMonitoringConfigurations() map[string]MonitoringConfiguration1 {
	if o == nil {
		var ret map[string]MonitoringConfiguration1
		return ret
	}

	return o.MonitoringConfigurations
}

// GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field value
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringConfigurations, true
}

// SetMonitoringConfigurations sets field value
func (o *EeSubscriptionExt) SetMonitoringConfigurations(v map[string]MonitoringConfiguration1) {
	o.MonitoringConfigurations = v
}

// GetReportingOptions returns the ReportingOptions field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetReportingOptions() ReportingOptions1 {
	if o == nil || IsNil(o.ReportingOptions) {
		var ret ReportingOptions1
		return ret
	}
	return *o.ReportingOptions
}

// GetReportingOptionsOk returns a tuple with the ReportingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetReportingOptionsOk() (*ReportingOptions1, bool) {
	if o == nil || IsNil(o.ReportingOptions) {
		return nil, false
	}
	return o.ReportingOptions, true
}

// HasReportingOptions returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasReportingOptions() bool {
	if o != nil && !IsNil(o.ReportingOptions) {
		return true
	}

	return false
}

// SetReportingOptions gets a reference to the given ReportingOptions1 and assigns it to the ReportingOptions field.
func (o *EeSubscriptionExt) SetReportingOptions(v ReportingOptions1) {
	o.ReportingOptions = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EeSubscriptionExt) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *EeSubscriptionExt) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetContextInfo() ContextInfo {
	if o == nil || IsNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || IsNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasContextInfo() bool {
	if o != nil && !IsNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *EeSubscriptionExt) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetEpcAppliedInd returns the EpcAppliedInd field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetEpcAppliedInd() bool {
	if o == nil || IsNil(o.EpcAppliedInd) {
		var ret bool
		return ret
	}
	return *o.EpcAppliedInd
}

// GetEpcAppliedIndOk returns a tuple with the EpcAppliedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetEpcAppliedIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EpcAppliedInd) {
		return nil, false
	}
	return o.EpcAppliedInd, true
}

// HasEpcAppliedInd returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasEpcAppliedInd() bool {
	if o != nil && !IsNil(o.EpcAppliedInd) {
		return true
	}

	return false
}

// SetEpcAppliedInd gets a reference to the given bool and assigns it to the EpcAppliedInd field.
func (o *EeSubscriptionExt) SetEpcAppliedInd(v bool) {
	o.EpcAppliedInd = &v
}

// GetScefDiamHost returns the ScefDiamHost field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetScefDiamHost() string {
	if o == nil || IsNil(o.ScefDiamHost) {
		var ret string
		return ret
	}
	return *o.ScefDiamHost
}

// GetScefDiamHostOk returns a tuple with the ScefDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetScefDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.ScefDiamHost) {
		return nil, false
	}
	return o.ScefDiamHost, true
}

// HasScefDiamHost returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasScefDiamHost() bool {
	if o != nil && !IsNil(o.ScefDiamHost) {
		return true
	}

	return false
}

// SetScefDiamHost gets a reference to the given string and assigns it to the ScefDiamHost field.
func (o *EeSubscriptionExt) SetScefDiamHost(v string) {
	o.ScefDiamHost = &v
}

// GetScefDiamRealm returns the ScefDiamRealm field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetScefDiamRealm() string {
	if o == nil || IsNil(o.ScefDiamRealm) {
		var ret string
		return ret
	}
	return *o.ScefDiamRealm
}

// GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetScefDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.ScefDiamRealm) {
		return nil, false
	}
	return o.ScefDiamRealm, true
}

// HasScefDiamRealm returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasScefDiamRealm() bool {
	if o != nil && !IsNil(o.ScefDiamRealm) {
		return true
	}

	return false
}

// SetScefDiamRealm gets a reference to the given string and assigns it to the ScefDiamRealm field.
func (o *EeSubscriptionExt) SetScefDiamRealm(v string) {
	o.ScefDiamRealm = &v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *EeSubscriptionExt) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetSecondCallbackRef returns the SecondCallbackRef field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetSecondCallbackRef() string {
	if o == nil || IsNil(o.SecondCallbackRef) {
		var ret string
		return ret
	}
	return *o.SecondCallbackRef
}

// GetSecondCallbackRefOk returns a tuple with the SecondCallbackRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetSecondCallbackRefOk() (*string, bool) {
	if o == nil || IsNil(o.SecondCallbackRef) {
		return nil, false
	}
	return o.SecondCallbackRef, true
}

// HasSecondCallbackRef returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasSecondCallbackRef() bool {
	if o != nil && !IsNil(o.SecondCallbackRef) {
		return true
	}

	return false
}

// SetSecondCallbackRef gets a reference to the given string and assigns it to the SecondCallbackRef field.
func (o *EeSubscriptionExt) SetSecondCallbackRef(v string) {
	o.SecondCallbackRef = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EeSubscriptionExt) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExcludeGpsiList returns the ExcludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetExcludeGpsiList() []string {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		var ret []string
		return ret
	}
	return o.ExcludeGpsiList
}

// GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetExcludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		return nil, false
	}
	return o.ExcludeGpsiList, true
}

// HasExcludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasExcludeGpsiList() bool {
	if o != nil && !IsNil(o.ExcludeGpsiList) {
		return true
	}

	return false
}

// SetExcludeGpsiList gets a reference to the given []string and assigns it to the ExcludeGpsiList field.
func (o *EeSubscriptionExt) SetExcludeGpsiList(v []string) {
	o.ExcludeGpsiList = v
}

// GetIncludeGpsiList returns the IncludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetIncludeGpsiList() []string {
	if o == nil || IsNil(o.IncludeGpsiList) {
		var ret []string
		return ret
	}
	return o.IncludeGpsiList
}

// GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetIncludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeGpsiList) {
		return nil, false
	}
	return o.IncludeGpsiList, true
}

// HasIncludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasIncludeGpsiList() bool {
	if o != nil && !IsNil(o.IncludeGpsiList) {
		return true
	}

	return false
}

// SetIncludeGpsiList gets a reference to the given []string and assigns it to the IncludeGpsiList field.
func (o *EeSubscriptionExt) SetIncludeGpsiList(v []string) {
	o.IncludeGpsiList = v
}

// GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetDataRestorationCallbackUri() string {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.DataRestorationCallbackUri
}

// GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetDataRestorationCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		return nil, false
	}
	return o.DataRestorationCallbackUri, true
}

// HasDataRestorationCallbackUri returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasDataRestorationCallbackUri() bool {
	if o != nil && !IsNil(o.DataRestorationCallbackUri) {
		return true
	}

	return false
}

// SetDataRestorationCallbackUri gets a reference to the given string and assigns it to the DataRestorationCallbackUri field.
func (o *EeSubscriptionExt) SetDataRestorationCallbackUri(v string) {
	o.DataRestorationCallbackUri = &v
}

// GetAmfSubscriptionInfoList returns the AmfSubscriptionInfoList field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetAmfSubscriptionInfoList() []AmfSubscriptionInfo {
	if o == nil || IsNil(o.AmfSubscriptionInfoList) {
		var ret []AmfSubscriptionInfo
		return ret
	}
	return o.AmfSubscriptionInfoList
}

// GetAmfSubscriptionInfoListOk returns a tuple with the AmfSubscriptionInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetAmfSubscriptionInfoListOk() ([]AmfSubscriptionInfo, bool) {
	if o == nil || IsNil(o.AmfSubscriptionInfoList) {
		return nil, false
	}
	return o.AmfSubscriptionInfoList, true
}

// HasAmfSubscriptionInfoList returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasAmfSubscriptionInfoList() bool {
	if o != nil && !IsNil(o.AmfSubscriptionInfoList) {
		return true
	}

	return false
}

// SetAmfSubscriptionInfoList gets a reference to the given []AmfSubscriptionInfo and assigns it to the AmfSubscriptionInfoList field.
func (o *EeSubscriptionExt) SetAmfSubscriptionInfoList(v []AmfSubscriptionInfo) {
	o.AmfSubscriptionInfoList = v
}

// GetSmfSubscriptionInfo returns the SmfSubscriptionInfo field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetSmfSubscriptionInfo() SmfSubscriptionInfo {
	if o == nil || IsNil(o.SmfSubscriptionInfo) {
		var ret SmfSubscriptionInfo
		return ret
	}
	return *o.SmfSubscriptionInfo
}

// GetSmfSubscriptionInfoOk returns a tuple with the SmfSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetSmfSubscriptionInfoOk() (*SmfSubscriptionInfo, bool) {
	if o == nil || IsNil(o.SmfSubscriptionInfo) {
		return nil, false
	}
	return o.SmfSubscriptionInfo, true
}

// HasSmfSubscriptionInfo returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasSmfSubscriptionInfo() bool {
	if o != nil && !IsNil(o.SmfSubscriptionInfo) {
		return true
	}

	return false
}

// SetSmfSubscriptionInfo gets a reference to the given SmfSubscriptionInfo and assigns it to the SmfSubscriptionInfo field.
func (o *EeSubscriptionExt) SetSmfSubscriptionInfo(v SmfSubscriptionInfo) {
	o.SmfSubscriptionInfo = &v
}

// GetHssSubscriptionInfo returns the HssSubscriptionInfo field value if set, zero value otherwise.
func (o *EeSubscriptionExt) GetHssSubscriptionInfo() HssSubscriptionInfo {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		var ret HssSubscriptionInfo
		return ret
	}
	return *o.HssSubscriptionInfo
}

// GetHssSubscriptionInfoOk returns a tuple with the HssSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscriptionExt) GetHssSubscriptionInfoOk() (*HssSubscriptionInfo, bool) {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		return nil, false
	}
	return o.HssSubscriptionInfo, true
}

// HasHssSubscriptionInfo returns a boolean if a field has been set.
func (o *EeSubscriptionExt) HasHssSubscriptionInfo() bool {
	if o != nil && !IsNil(o.HssSubscriptionInfo) {
		return true
	}

	return false
}

// SetHssSubscriptionInfo gets a reference to the given HssSubscriptionInfo and assigns it to the HssSubscriptionInfo field.
func (o *EeSubscriptionExt) SetHssSubscriptionInfo(v HssSubscriptionInfo) {
	o.HssSubscriptionInfo = &v
}

func (o EeSubscriptionExt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EeSubscriptionExt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["callbackReference"] = o.CallbackReference
	toSerialize["monitoringConfigurations"] = o.MonitoringConfigurations
	if !IsNil(o.ReportingOptions) {
		toSerialize["reportingOptions"] = o.ReportingOptions
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	if !IsNil(o.EpcAppliedInd) {
		toSerialize["epcAppliedInd"] = o.EpcAppliedInd
	}
	if !IsNil(o.ScefDiamHost) {
		toSerialize["scefDiamHost"] = o.ScefDiamHost
	}
	if !IsNil(o.ScefDiamRealm) {
		toSerialize["scefDiamRealm"] = o.ScefDiamRealm
	}
	if !IsNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	if !IsNil(o.SecondCallbackRef) {
		toSerialize["secondCallbackRef"] = o.SecondCallbackRef
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.ExcludeGpsiList) {
		toSerialize["excludeGpsiList"] = o.ExcludeGpsiList
	}
	if !IsNil(o.IncludeGpsiList) {
		toSerialize["includeGpsiList"] = o.IncludeGpsiList
	}
	if !IsNil(o.DataRestorationCallbackUri) {
		toSerialize["dataRestorationCallbackUri"] = o.DataRestorationCallbackUri
	}
	if !IsNil(o.AmfSubscriptionInfoList) {
		toSerialize["amfSubscriptionInfoList"] = o.AmfSubscriptionInfoList
	}
	if !IsNil(o.SmfSubscriptionInfo) {
		toSerialize["smfSubscriptionInfo"] = o.SmfSubscriptionInfo
	}
	if !IsNil(o.HssSubscriptionInfo) {
		toSerialize["hssSubscriptionInfo"] = o.HssSubscriptionInfo
	}
	return toSerialize, nil
}

type NullableEeSubscriptionExt struct {
	value *EeSubscriptionExt
	isSet bool
}

func (v NullableEeSubscriptionExt) Get() *EeSubscriptionExt {
	return v.value
}

func (v *NullableEeSubscriptionExt) Set(val *EeSubscriptionExt) {
	v.value = val
	v.isSet = true
}

func (v NullableEeSubscriptionExt) IsSet() bool {
	return v.isSet
}

func (v *NullableEeSubscriptionExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEeSubscriptionExt(val *EeSubscriptionExt) *NullableEeSubscriptionExt {
	return &NullableEeSubscriptionExt{value: val, isSet: true}
}

func (v NullableEeSubscriptionExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEeSubscriptionExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

