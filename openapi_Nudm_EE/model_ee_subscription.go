/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
)

// checks if the EeSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EeSubscription{}

// EeSubscription struct for EeSubscription
type EeSubscription struct {
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// A map (list of key-value pairs where ReferenceId serves as key) of MonitoringConfigurations
	MonitoringConfigurations map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ReportingOptions         *ReportingOptions                  `json:"reportingOptions,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string      `json:"supportedFeatures,omitempty"`
	SubscriptionId    *string      `json:"subscriptionId,omitempty"`
	ContextInfo       *ContextInfo `json:"contextInfo,omitempty"`
	EpcAppliedInd     *bool        `json:"epcAppliedInd,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamHost *string `json:"scefDiamHost,omitempty"`
	// Fully Qualified Domain Name
	ScefDiamRealm       *string `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SecondCallbackRef *string `json:"secondCallbackRef,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi            *string  `json:"gpsi,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	DataRestorationCallbackUri *string `json:"dataRestorationCallbackUri,omitempty"`
}

// NewEeSubscription instantiates a new EeSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEeSubscription(callbackReference string, monitoringConfigurations map[string]MonitoringConfiguration) *EeSubscription {
	this := EeSubscription{}
	this.CallbackReference = callbackReference
	this.MonitoringConfigurations = monitoringConfigurations
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// NewEeSubscriptionWithDefaults instantiates a new EeSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEeSubscriptionWithDefaults() *EeSubscription {
	this := EeSubscription{}
	var epcAppliedInd bool = false
	this.EpcAppliedInd = &epcAppliedInd
	return &this
}

// GetCallbackReference returns the CallbackReference field value
func (o *EeSubscription) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *EeSubscription) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetMonitoringConfigurations returns the MonitoringConfigurations field value
func (o *EeSubscription) GetMonitoringConfigurations() map[string]MonitoringConfiguration {
	if o == nil {
		var ret map[string]MonitoringConfiguration
		return ret
	}

	return o.MonitoringConfigurations
}

// GetMonitoringConfigurationsOk returns a tuple with the MonitoringConfigurations field value
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetMonitoringConfigurationsOk() (*map[string]MonitoringConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringConfigurations, true
}

// SetMonitoringConfigurations sets field value
func (o *EeSubscription) SetMonitoringConfigurations(v map[string]MonitoringConfiguration) {
	o.MonitoringConfigurations = v
}

// GetReportingOptions returns the ReportingOptions field value if set, zero value otherwise.
func (o *EeSubscription) GetReportingOptions() ReportingOptions {
	if o == nil || IsNil(o.ReportingOptions) {
		var ret ReportingOptions
		return ret
	}
	return *o.ReportingOptions
}

// GetReportingOptionsOk returns a tuple with the ReportingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetReportingOptionsOk() (*ReportingOptions, bool) {
	if o == nil || IsNil(o.ReportingOptions) {
		return nil, false
	}
	return o.ReportingOptions, true
}

// HasReportingOptions returns a boolean if a field has been set.
func (o *EeSubscription) HasReportingOptions() bool {
	if o != nil && !IsNil(o.ReportingOptions) {
		return true
	}

	return false
}

// SetReportingOptions gets a reference to the given ReportingOptions and assigns it to the ReportingOptions field.
func (o *EeSubscription) SetReportingOptions(v ReportingOptions) {
	o.ReportingOptions = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EeSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EeSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EeSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *EeSubscription) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *EeSubscription) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *EeSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *EeSubscription) GetContextInfo() ContextInfo {
	if o == nil || IsNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || IsNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *EeSubscription) HasContextInfo() bool {
	if o != nil && !IsNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *EeSubscription) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

// GetEpcAppliedInd returns the EpcAppliedInd field value if set, zero value otherwise.
func (o *EeSubscription) GetEpcAppliedInd() bool {
	if o == nil || IsNil(o.EpcAppliedInd) {
		var ret bool
		return ret
	}
	return *o.EpcAppliedInd
}

// GetEpcAppliedIndOk returns a tuple with the EpcAppliedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetEpcAppliedIndOk() (*bool, bool) {
	if o == nil || IsNil(o.EpcAppliedInd) {
		return nil, false
	}
	return o.EpcAppliedInd, true
}

// HasEpcAppliedInd returns a boolean if a field has been set.
func (o *EeSubscription) HasEpcAppliedInd() bool {
	if o != nil && !IsNil(o.EpcAppliedInd) {
		return true
	}

	return false
}

// SetEpcAppliedInd gets a reference to the given bool and assigns it to the EpcAppliedInd field.
func (o *EeSubscription) SetEpcAppliedInd(v bool) {
	o.EpcAppliedInd = &v
}

// GetScefDiamHost returns the ScefDiamHost field value if set, zero value otherwise.
func (o *EeSubscription) GetScefDiamHost() string {
	if o == nil || IsNil(o.ScefDiamHost) {
		var ret string
		return ret
	}
	return *o.ScefDiamHost
}

// GetScefDiamHostOk returns a tuple with the ScefDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetScefDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.ScefDiamHost) {
		return nil, false
	}
	return o.ScefDiamHost, true
}

// HasScefDiamHost returns a boolean if a field has been set.
func (o *EeSubscription) HasScefDiamHost() bool {
	if o != nil && !IsNil(o.ScefDiamHost) {
		return true
	}

	return false
}

// SetScefDiamHost gets a reference to the given string and assigns it to the ScefDiamHost field.
func (o *EeSubscription) SetScefDiamHost(v string) {
	o.ScefDiamHost = &v
}

// GetScefDiamRealm returns the ScefDiamRealm field value if set, zero value otherwise.
func (o *EeSubscription) GetScefDiamRealm() string {
	if o == nil || IsNil(o.ScefDiamRealm) {
		var ret string
		return ret
	}
	return *o.ScefDiamRealm
}

// GetScefDiamRealmOk returns a tuple with the ScefDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetScefDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.ScefDiamRealm) {
		return nil, false
	}
	return o.ScefDiamRealm, true
}

// HasScefDiamRealm returns a boolean if a field has been set.
func (o *EeSubscription) HasScefDiamRealm() bool {
	if o != nil && !IsNil(o.ScefDiamRealm) {
		return true
	}

	return false
}

// SetScefDiamRealm gets a reference to the given string and assigns it to the ScefDiamRealm field.
func (o *EeSubscription) SetScefDiamRealm(v string) {
	o.ScefDiamRealm = &v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *EeSubscription) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *EeSubscription) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *EeSubscription) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetSecondCallbackRef returns the SecondCallbackRef field value if set, zero value otherwise.
func (o *EeSubscription) GetSecondCallbackRef() string {
	if o == nil || IsNil(o.SecondCallbackRef) {
		var ret string
		return ret
	}
	return *o.SecondCallbackRef
}

// GetSecondCallbackRefOk returns a tuple with the SecondCallbackRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetSecondCallbackRefOk() (*string, bool) {
	if o == nil || IsNil(o.SecondCallbackRef) {
		return nil, false
	}
	return o.SecondCallbackRef, true
}

// HasSecondCallbackRef returns a boolean if a field has been set.
func (o *EeSubscription) HasSecondCallbackRef() bool {
	if o != nil && !IsNil(o.SecondCallbackRef) {
		return true
	}

	return false
}

// SetSecondCallbackRef gets a reference to the given string and assigns it to the SecondCallbackRef field.
func (o *EeSubscription) SetSecondCallbackRef(v string) {
	o.SecondCallbackRef = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EeSubscription) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EeSubscription) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EeSubscription) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExcludeGpsiList returns the ExcludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscription) GetExcludeGpsiList() []string {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		var ret []string
		return ret
	}
	return o.ExcludeGpsiList
}

// GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetExcludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		return nil, false
	}
	return o.ExcludeGpsiList, true
}

// HasExcludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscription) HasExcludeGpsiList() bool {
	if o != nil && !IsNil(o.ExcludeGpsiList) {
		return true
	}

	return false
}

// SetExcludeGpsiList gets a reference to the given []string and assigns it to the ExcludeGpsiList field.
func (o *EeSubscription) SetExcludeGpsiList(v []string) {
	o.ExcludeGpsiList = v
}

// GetIncludeGpsiList returns the IncludeGpsiList field value if set, zero value otherwise.
func (o *EeSubscription) GetIncludeGpsiList() []string {
	if o == nil || IsNil(o.IncludeGpsiList) {
		var ret []string
		return ret
	}
	return o.IncludeGpsiList
}

// GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetIncludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeGpsiList) {
		return nil, false
	}
	return o.IncludeGpsiList, true
}

// HasIncludeGpsiList returns a boolean if a field has been set.
func (o *EeSubscription) HasIncludeGpsiList() bool {
	if o != nil && !IsNil(o.IncludeGpsiList) {
		return true
	}

	return false
}

// SetIncludeGpsiList gets a reference to the given []string and assigns it to the IncludeGpsiList field.
func (o *EeSubscription) SetIncludeGpsiList(v []string) {
	o.IncludeGpsiList = v
}

// GetDataRestorationCallbackUri returns the DataRestorationCallbackUri field value if set, zero value otherwise.
func (o *EeSubscription) GetDataRestorationCallbackUri() string {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		var ret string
		return ret
	}
	return *o.DataRestorationCallbackUri
}

// GetDataRestorationCallbackUriOk returns a tuple with the DataRestorationCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EeSubscription) GetDataRestorationCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.DataRestorationCallbackUri) {
		return nil, false
	}
	return o.DataRestorationCallbackUri, true
}

// HasDataRestorationCallbackUri returns a boolean if a field has been set.
func (o *EeSubscription) HasDataRestorationCallbackUri() bool {
	if o != nil && !IsNil(o.DataRestorationCallbackUri) {
		return true
	}

	return false
}

// SetDataRestorationCallbackUri gets a reference to the given string and assigns it to the DataRestorationCallbackUri field.
func (o *EeSubscription) SetDataRestorationCallbackUri(v string) {
	o.DataRestorationCallbackUri = &v
}

func (o EeSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EeSubscription) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableEeSubscription struct {
	value *EeSubscription
	isSet bool
}

func (v NullableEeSubscription) Get() *EeSubscription {
	return v.value
}

func (v *NullableEeSubscription) Set(val *EeSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEeSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEeSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEeSubscription(val *EeSubscription) *NullableEeSubscription {
	return &NullableEeSubscription{value: val, isSet: true}
}

func (v NullableEeSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEeSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
