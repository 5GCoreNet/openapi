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

// checks if the SubscriptionDataSubscriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDataSubscriptions{}

// SubscriptionDataSubscriptions A subscription to notifications.
type SubscriptionDataSubscriptions struct {
	// String represents the SUPI or GPSI
	UeId *string `json:"ueId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// String providing an URI formatted according to RFC 3986.
	OriginalCallbackReference *string `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	SdmSubscription *SdmSubscription1 `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	UniqueSubscription *bool `json:"uniqueSubscription,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	ImmediateReport *bool `json:"immediateReport,omitempty"`
	Report *ImmediateReport1 `json:"report,omitempty"`
}

// NewSubscriptionDataSubscriptions instantiates a new SubscriptionDataSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDataSubscriptions(callbackReference string, monitoredResourceUris []string) *SubscriptionDataSubscriptions {
	this := SubscriptionDataSubscriptions{}
	this.CallbackReference = callbackReference
	this.MonitoredResourceUris = monitoredResourceUris
	var immediateReport bool = false
	this.ImmediateReport = &immediateReport
	return &this
}

// NewSubscriptionDataSubscriptionsWithDefaults instantiates a new SubscriptionDataSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataSubscriptionsWithDefaults() *SubscriptionDataSubscriptions {
	this := SubscriptionDataSubscriptions{}
	var immediateReport bool = false
	this.ImmediateReport = &immediateReport
	return &this
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetUeId() string {
	if o == nil || IsNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetUeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasUeId() bool {
	if o != nil && !IsNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *SubscriptionDataSubscriptions) SetUeId(v string) {
	o.UeId = &v
}

// GetCallbackReference returns the CallbackReference field value
func (o *SubscriptionDataSubscriptions) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *SubscriptionDataSubscriptions) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetOriginalCallbackReference returns the OriginalCallbackReference field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReference() string {
	if o == nil || IsNil(o.OriginalCallbackReference) {
		var ret string
		return ret
	}
	return *o.OriginalCallbackReference
}

// GetOriginalCallbackReferenceOk returns a tuple with the OriginalCallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetOriginalCallbackReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalCallbackReference) {
		return nil, false
	}
	return o.OriginalCallbackReference, true
}

// HasOriginalCallbackReference returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasOriginalCallbackReference() bool {
	if o != nil && !IsNil(o.OriginalCallbackReference) {
		return true
	}

	return false
}

// SetOriginalCallbackReference gets a reference to the given string and assigns it to the OriginalCallbackReference field.
func (o *SubscriptionDataSubscriptions) SetOriginalCallbackReference(v string) {
	o.OriginalCallbackReference = &v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value
func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoredResourceUris, true
}

// SetMonitoredResourceUris sets field value
func (o *SubscriptionDataSubscriptions) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *SubscriptionDataSubscriptions) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSdmSubscription returns the SdmSubscription field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSdmSubscription() SdmSubscription1 {
	if o == nil || IsNil(o.SdmSubscription) {
		var ret SdmSubscription1
		return ret
	}
	return *o.SdmSubscription
}

// GetSdmSubscriptionOk returns a tuple with the SdmSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSdmSubscriptionOk() (*SdmSubscription1, bool) {
	if o == nil || IsNil(o.SdmSubscription) {
		return nil, false
	}
	return o.SdmSubscription, true
}

// HasSdmSubscription returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSdmSubscription() bool {
	if o != nil && !IsNil(o.SdmSubscription) {
		return true
	}

	return false
}

// SetSdmSubscription gets a reference to the given SdmSubscription1 and assigns it to the SdmSubscription field.
func (o *SubscriptionDataSubscriptions) SetSdmSubscription(v SdmSubscription1) {
	o.SdmSubscription = &v
}

// GetHssSubscriptionInfo returns the HssSubscriptionInfo field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetHssSubscriptionInfo() HssSubscriptionInfo {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		var ret HssSubscriptionInfo
		return ret
	}
	return *o.HssSubscriptionInfo
}

// GetHssSubscriptionInfoOk returns a tuple with the HssSubscriptionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetHssSubscriptionInfoOk() (*HssSubscriptionInfo, bool) {
	if o == nil || IsNil(o.HssSubscriptionInfo) {
		return nil, false
	}
	return o.HssSubscriptionInfo, true
}

// HasHssSubscriptionInfo returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasHssSubscriptionInfo() bool {
	if o != nil && !IsNil(o.HssSubscriptionInfo) {
		return true
	}

	return false
}

// SetHssSubscriptionInfo gets a reference to the given HssSubscriptionInfo and assigns it to the HssSubscriptionInfo field.
func (o *SubscriptionDataSubscriptions) SetHssSubscriptionInfo(v HssSubscriptionInfo) {
	o.HssSubscriptionInfo = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SubscriptionDataSubscriptions) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUniqueSubscription returns the UniqueSubscription field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetUniqueSubscription() bool {
	if o == nil || IsNil(o.UniqueSubscription) {
		var ret bool
		return ret
	}
	return *o.UniqueSubscription
}

// GetUniqueSubscriptionOk returns a tuple with the UniqueSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetUniqueSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueSubscription) {
		return nil, false
	}
	return o.UniqueSubscription, true
}

// HasUniqueSubscription returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasUniqueSubscription() bool {
	if o != nil && !IsNil(o.UniqueSubscription) {
		return true
	}

	return false
}

// SetUniqueSubscription gets a reference to the given bool and assigns it to the UniqueSubscription field.
func (o *SubscriptionDataSubscriptions) SetUniqueSubscription(v bool) {
	o.UniqueSubscription = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SubscriptionDataSubscriptions) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetImmediateReport returns the ImmediateReport field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetImmediateReport() bool {
	if o == nil || IsNil(o.ImmediateReport) {
		var ret bool
		return ret
	}
	return *o.ImmediateReport
}

// GetImmediateReportOk returns a tuple with the ImmediateReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetImmediateReportOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmediateReport) {
		return nil, false
	}
	return o.ImmediateReport, true
}

// HasImmediateReport returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasImmediateReport() bool {
	if o != nil && !IsNil(o.ImmediateReport) {
		return true
	}

	return false
}

// SetImmediateReport gets a reference to the given bool and assigns it to the ImmediateReport field.
func (o *SubscriptionDataSubscriptions) SetImmediateReport(v bool) {
	o.ImmediateReport = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *SubscriptionDataSubscriptions) GetReport() ImmediateReport1 {
	if o == nil || IsNil(o.Report) {
		var ret ImmediateReport1
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSubscriptions) GetReportOk() (*ImmediateReport1, bool) {
	if o == nil || IsNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *SubscriptionDataSubscriptions) HasReport() bool {
	if o != nil && !IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given ImmediateReport1 and assigns it to the Report field.
func (o *SubscriptionDataSubscriptions) SetReport(v ImmediateReport1) {
	o.Report = &v
}

func (o SubscriptionDataSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDataSubscriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	toSerialize["callbackReference"] = o.CallbackReference
	if !IsNil(o.OriginalCallbackReference) {
		toSerialize["originalCallbackReference"] = o.OriginalCallbackReference
	}
	toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.SdmSubscription) {
		toSerialize["sdmSubscription"] = o.SdmSubscription
	}
	if !IsNil(o.HssSubscriptionInfo) {
		toSerialize["hssSubscriptionInfo"] = o.HssSubscriptionInfo
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UniqueSubscription) {
		toSerialize["uniqueSubscription"] = o.UniqueSubscription
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ImmediateReport) {
		toSerialize["immediateReport"] = o.ImmediateReport
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	return toSerialize, nil
}

type NullableSubscriptionDataSubscriptions struct {
	value *SubscriptionDataSubscriptions
	isSet bool
}

func (v NullableSubscriptionDataSubscriptions) Get() *SubscriptionDataSubscriptions {
	return v.value
}

func (v *NullableSubscriptionDataSubscriptions) Set(val *SubscriptionDataSubscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDataSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDataSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDataSubscriptions(val *SubscriptionDataSubscriptions) *NullableSubscriptionDataSubscriptions {
	return &NullableSubscriptionDataSubscriptions{value: val, isSet: true}
}

func (v NullableSubscriptionDataSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDataSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


