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

// checks if the ApplicationDataSubs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationDataSubs{}

// ApplicationDataSubs Identifies a subscription to application data change notification.
type ApplicationDataSubs struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri"`
	DataFilters []DataFilter `json:"dataFilters,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Immediate reporting indication.
	ImmRep *bool `json:"immRep,omitempty"`
	// The AM Influence Data entries stored in the UDR that match a subscription.
	AmInfluEntries []AmInfluData `json:"amInfluEntries,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewApplicationDataSubs instantiates a new ApplicationDataSubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDataSubs(notificationUri string) *ApplicationDataSubs {
	this := ApplicationDataSubs{}
	this.NotificationUri = notificationUri
	return &this
}

// NewApplicationDataSubsWithDefaults instantiates a new ApplicationDataSubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDataSubsWithDefaults() *ApplicationDataSubs {
	this := ApplicationDataSubs{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value
func (o *ApplicationDataSubs) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *ApplicationDataSubs) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetDataFilters returns the DataFilters field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetDataFilters() []DataFilter {
	if o == nil || IsNil(o.DataFilters) {
		var ret []DataFilter
		return ret
	}
	return o.DataFilters
}

// GetDataFiltersOk returns a tuple with the DataFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetDataFiltersOk() ([]DataFilter, bool) {
	if o == nil || IsNil(o.DataFilters) {
		return nil, false
	}
	return o.DataFilters, true
}

// HasDataFilters returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasDataFilters() bool {
	if o != nil && !IsNil(o.DataFilters) {
		return true
	}

	return false
}

// SetDataFilters gets a reference to the given []DataFilter and assigns it to the DataFilters field.
func (o *ApplicationDataSubs) SetDataFilters(v []DataFilter) {
	o.DataFilters = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *ApplicationDataSubs) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetImmRep() bool {
	if o == nil || IsNil(o.ImmRep) {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetImmRepOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmRep) {
		return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasImmRep() bool {
	if o != nil && !IsNil(o.ImmRep) {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *ApplicationDataSubs) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetAmInfluEntries returns the AmInfluEntries field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetAmInfluEntries() []AmInfluData {
	if o == nil || IsNil(o.AmInfluEntries) {
		var ret []AmInfluData
		return ret
	}
	return o.AmInfluEntries
}

// GetAmInfluEntriesOk returns a tuple with the AmInfluEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetAmInfluEntriesOk() ([]AmInfluData, bool) {
	if o == nil || IsNil(o.AmInfluEntries) {
		return nil, false
	}
	return o.AmInfluEntries, true
}

// HasAmInfluEntries returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasAmInfluEntries() bool {
	if o != nil && !IsNil(o.AmInfluEntries) {
		return true
	}

	return false
}

// SetAmInfluEntries gets a reference to the given []AmInfluData and assigns it to the AmInfluEntries field.
func (o *ApplicationDataSubs) SetAmInfluEntries(v []AmInfluData) {
	o.AmInfluEntries = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ApplicationDataSubs) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *ApplicationDataSubs) GetResetIds() []string {
	if o == nil || IsNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationDataSubs) GetResetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResetIds) {
		return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *ApplicationDataSubs) HasResetIds() bool {
	if o != nil && !IsNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *ApplicationDataSubs) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o ApplicationDataSubs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationDataSubs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notificationUri"] = o.NotificationUri
	if !IsNil(o.DataFilters) {
		toSerialize["dataFilters"] = o.DataFilters
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.ImmRep) {
		toSerialize["immRep"] = o.ImmRep
	}
	if !IsNil(o.AmInfluEntries) {
		toSerialize["amInfluEntries"] = o.AmInfluEntries
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return toSerialize, nil
}

type NullableApplicationDataSubs struct {
	value *ApplicationDataSubs
	isSet bool
}

func (v NullableApplicationDataSubs) Get() *ApplicationDataSubs {
	return v.value
}

func (v *NullableApplicationDataSubs) Set(val *ApplicationDataSubs) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDataSubs) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDataSubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDataSubs(val *ApplicationDataSubs) *NullableApplicationDataSubs {
	return &NullableApplicationDataSubs{value: val, isSet: true}
}

func (v NullableApplicationDataSubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDataSubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


