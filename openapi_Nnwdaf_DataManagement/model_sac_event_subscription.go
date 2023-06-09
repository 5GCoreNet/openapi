/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the SACEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SACEventSubscription{}

// SACEventSubscription Request data to create the event subscription
type SACEventSubscription struct {
	Event SACEvent `json:"event"`
	// String providing an URI formatted according to RFC 3986.
	EventNotifyUri string `json:"eventNotifyUri"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	NfId                string  `json:"nfId"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	MaxReports          *int32  `json:"maxReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewSACEventSubscription instantiates a new SACEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACEventSubscription(event SACEvent, eventNotifyUri string, nfId string) *SACEventSubscription {
	this := SACEventSubscription{}
	this.Event = event
	this.EventNotifyUri = eventNotifyUri
	this.NfId = nfId
	return &this
}

// NewSACEventSubscriptionWithDefaults instantiates a new SACEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACEventSubscriptionWithDefaults() *SACEventSubscription {
	this := SACEventSubscription{}
	return &this
}

// GetEvent returns the Event field value
func (o *SACEventSubscription) GetEvent() SACEvent {
	if o == nil {
		var ret SACEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetEventOk() (*SACEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *SACEventSubscription) SetEvent(v SACEvent) {
	o.Event = v
}

// GetEventNotifyUri returns the EventNotifyUri field value
func (o *SACEventSubscription) GetEventNotifyUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventNotifyUri
}

// GetEventNotifyUriOk returns a tuple with the EventNotifyUri field value
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetEventNotifyUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventNotifyUri, true
}

// SetEventNotifyUri sets field value
func (o *SACEventSubscription) SetEventNotifyUri(v string) {
	o.EventNotifyUri = v
}

// GetNfId returns the NfId field value
func (o *SACEventSubscription) GetNfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetNfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfId, true
}

// SetNfId sets field value
func (o *SACEventSubscription) SetNfId(v string) {
	o.NfId = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *SACEventSubscription) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *SACEventSubscription) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *SACEventSubscription) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetMaxReports returns the MaxReports field value if set, zero value otherwise.
func (o *SACEventSubscription) GetMaxReports() int32 {
	if o == nil || IsNil(o.MaxReports) {
		var ret int32
		return ret
	}
	return *o.MaxReports
}

// GetMaxReportsOk returns a tuple with the MaxReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetMaxReportsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReports) {
		return nil, false
	}
	return o.MaxReports, true
}

// HasMaxReports returns a boolean if a field has been set.
func (o *SACEventSubscription) HasMaxReports() bool {
	if o != nil && !IsNil(o.MaxReports) {
		return true
	}

	return false
}

// SetMaxReports gets a reference to the given int32 and assigns it to the MaxReports field.
func (o *SACEventSubscription) SetMaxReports(v int32) {
	o.MaxReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *SACEventSubscription) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *SACEventSubscription) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *SACEventSubscription) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SACEventSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SACEventSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SACEventSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SACEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SACEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["eventNotifyUri"] = o.EventNotifyUri
	toSerialize["nfId"] = o.NfId
	if !IsNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	if !IsNil(o.MaxReports) {
		toSerialize["maxReports"] = o.MaxReports
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableSACEventSubscription struct {
	value *SACEventSubscription
	isSet bool
}

func (v NullableSACEventSubscription) Get() *SACEventSubscription {
	return v.value
}

func (v *NullableSACEventSubscription) Set(val *SACEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventSubscription(val *SACEventSubscription) *NullableSACEventSubscription {
	return &NullableSACEventSubscription{value: val, isSet: true}
}

func (v NullableSACEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
