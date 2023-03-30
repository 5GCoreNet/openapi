/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
	"time"
)

// checks if the PolicyDataSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyDataSubscription{}

// PolicyDataSubscription Identifies a subscription to policy data change notification.
type PolicyDataSubscription struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri"`
	NotifId *string `json:"notifId,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris"`
	MonResItems []ResourceItem `json:"monResItems,omitempty"`
	ExcludedResItems []ResourceItem `json:"excludedResItems,omitempty"`
	// If provided and set to true, it indicates that existing entries that match this subscription shall be immediately reported in the response. 
	ImmRep *bool `json:"immRep,omitempty"`
	// Immediate report with existing UDR entries.
	ImmReports []PolicyDataChangeNotification `json:"immReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewPolicyDataSubscription instantiates a new PolicyDataSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDataSubscription(notificationUri string, monitoredResourceUris []string) *PolicyDataSubscription {
	this := PolicyDataSubscription{}
	this.NotificationUri = notificationUri
	this.MonitoredResourceUris = monitoredResourceUris
	return &this
}

// NewPolicyDataSubscriptionWithDefaults instantiates a new PolicyDataSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDataSubscriptionWithDefaults() *PolicyDataSubscription {
	this := PolicyDataSubscription{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value
func (o *PolicyDataSubscription) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *PolicyDataSubscription) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetNotifId returns the NotifId field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetNotifId() string {
	if o == nil || IsNil(o.NotifId) {
		var ret string
		return ret
	}
	return *o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetNotifIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifId) {
		return nil, false
	}
	return o.NotifId, true
}

// HasNotifId returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasNotifId() bool {
	if o != nil && !IsNil(o.NotifId) {
		return true
	}

	return false
}

// SetNotifId gets a reference to the given string and assigns it to the NotifId field.
func (o *PolicyDataSubscription) SetNotifId(v string) {
	o.NotifId = &v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value
func (o *PolicyDataSubscription) GetMonitoredResourceUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoredResourceUris, true
}

// SetMonitoredResourceUris sets field value
func (o *PolicyDataSubscription) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetMonResItems returns the MonResItems field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetMonResItems() []ResourceItem {
	if o == nil || IsNil(o.MonResItems) {
		var ret []ResourceItem
		return ret
	}
	return o.MonResItems
}

// GetMonResItemsOk returns a tuple with the MonResItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetMonResItemsOk() ([]ResourceItem, bool) {
	if o == nil || IsNil(o.MonResItems) {
		return nil, false
	}
	return o.MonResItems, true
}

// HasMonResItems returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasMonResItems() bool {
	if o != nil && !IsNil(o.MonResItems) {
		return true
	}

	return false
}

// SetMonResItems gets a reference to the given []ResourceItem and assigns it to the MonResItems field.
func (o *PolicyDataSubscription) SetMonResItems(v []ResourceItem) {
	o.MonResItems = v
}

// GetExcludedResItems returns the ExcludedResItems field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetExcludedResItems() []ResourceItem {
	if o == nil || IsNil(o.ExcludedResItems) {
		var ret []ResourceItem
		return ret
	}
	return o.ExcludedResItems
}

// GetExcludedResItemsOk returns a tuple with the ExcludedResItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetExcludedResItemsOk() ([]ResourceItem, bool) {
	if o == nil || IsNil(o.ExcludedResItems) {
		return nil, false
	}
	return o.ExcludedResItems, true
}

// HasExcludedResItems returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasExcludedResItems() bool {
	if o != nil && !IsNil(o.ExcludedResItems) {
		return true
	}

	return false
}

// SetExcludedResItems gets a reference to the given []ResourceItem and assigns it to the ExcludedResItems field.
func (o *PolicyDataSubscription) SetExcludedResItems(v []ResourceItem) {
	o.ExcludedResItems = v
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetImmRep() bool {
	if o == nil || IsNil(o.ImmRep) {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetImmRepOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmRep) {
		return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasImmRep() bool {
	if o != nil && !IsNil(o.ImmRep) {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *PolicyDataSubscription) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetImmReports returns the ImmReports field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetImmReports() []PolicyDataChangeNotification {
	if o == nil || IsNil(o.ImmReports) {
		var ret []PolicyDataChangeNotification
		return ret
	}
	return o.ImmReports
}

// GetImmReportsOk returns a tuple with the ImmReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetImmReportsOk() ([]PolicyDataChangeNotification, bool) {
	if o == nil || IsNil(o.ImmReports) {
		return nil, false
	}
	return o.ImmReports, true
}

// HasImmReports returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasImmReports() bool {
	if o != nil && !IsNil(o.ImmReports) {
		return true
	}

	return false
}

// SetImmReports gets a reference to the given []PolicyDataChangeNotification and assigns it to the ImmReports field.
func (o *PolicyDataSubscription) SetImmReports(v []PolicyDataChangeNotification) {
	o.ImmReports = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *PolicyDataSubscription) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *PolicyDataSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *PolicyDataSubscription) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataSubscription) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *PolicyDataSubscription) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *PolicyDataSubscription) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o PolicyDataSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyDataSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notificationUri"] = o.NotificationUri
	if !IsNil(o.NotifId) {
		toSerialize["notifId"] = o.NotifId
	}
	toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	if !IsNil(o.MonResItems) {
		toSerialize["monResItems"] = o.MonResItems
	}
	if !IsNil(o.ExcludedResItems) {
		toSerialize["excludedResItems"] = o.ExcludedResItems
	}
	if !IsNil(o.ImmRep) {
		toSerialize["immRep"] = o.ImmRep
	}
	if !IsNil(o.ImmReports) {
		toSerialize["immReports"] = o.ImmReports
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return toSerialize, nil
}

type NullablePolicyDataSubscription struct {
	value *PolicyDataSubscription
	isSet bool
}

func (v NullablePolicyDataSubscription) Get() *PolicyDataSubscription {
	return v.value
}

func (v *NullablePolicyDataSubscription) Set(val *PolicyDataSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDataSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDataSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDataSubscription(val *PolicyDataSubscription) *NullablePolicyDataSubscription {
	return &NullablePolicyDataSubscription{value: val, isSet: true}
}

func (v NullablePolicyDataSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDataSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


