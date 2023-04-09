/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the NnwdafEventsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NnwdafEventsSubscription{}

// NnwdafEventsSubscription Represents an Individual NWDAF Event Subscription resource.
type NnwdafEventsSubscription struct {
	// Subscribed events
	EventSubscriptions []EventSubscription   `json:"eventSubscriptions"`
	EvtReq             *ReportingInformation `json:"evtReq,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotificationURI *string `json:"notificationURI,omitempty"`
	// Notification correlation identifier.
	NotifCorrId *string `json:"notifCorrId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures  *string                `json:"supportedFeatures,omitempty"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}

// NewNnwdafEventsSubscription instantiates a new NnwdafEventsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNnwdafEventsSubscription(eventSubscriptions []EventSubscription) *NnwdafEventsSubscription {
	this := NnwdafEventsSubscription{}
	this.EventSubscriptions = eventSubscriptions
	return &this
}

// NewNnwdafEventsSubscriptionWithDefaults instantiates a new NnwdafEventsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNnwdafEventsSubscriptionWithDefaults() *NnwdafEventsSubscription {
	this := NnwdafEventsSubscription{}
	return &this
}

// GetEventSubscriptions returns the EventSubscriptions field value
func (o *NnwdafEventsSubscription) GetEventSubscriptions() []EventSubscription {
	if o == nil {
		var ret []EventSubscription
		return ret
	}

	return o.EventSubscriptions
}

// GetEventSubscriptionsOk returns a tuple with the EventSubscriptions field value
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetEventSubscriptionsOk() ([]EventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubscriptions, true
}

// SetEventSubscriptions sets field value
func (o *NnwdafEventsSubscription) SetEventSubscriptions(v []EventSubscription) {
	o.EventSubscriptions = v
}

// GetEvtReq returns the EvtReq field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetEvtReq() ReportingInformation {
	if o == nil || IsNil(o.EvtReq) {
		var ret ReportingInformation
		return ret
	}
	return *o.EvtReq
}

// GetEvtReqOk returns a tuple with the EvtReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetEvtReqOk() (*ReportingInformation, bool) {
	if o == nil || IsNil(o.EvtReq) {
		return nil, false
	}
	return o.EvtReq, true
}

// HasEvtReq returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasEvtReq() bool {
	if o != nil && !IsNil(o.EvtReq) {
		return true
	}

	return false
}

// SetEvtReq gets a reference to the given ReportingInformation and assigns it to the EvtReq field.
func (o *NnwdafEventsSubscription) SetEvtReq(v ReportingInformation) {
	o.EvtReq = &v
}

// GetNotificationURI returns the NotificationURI field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetNotificationURI() string {
	if o == nil || IsNil(o.NotificationURI) {
		var ret string
		return ret
	}
	return *o.NotificationURI
}

// GetNotificationURIOk returns a tuple with the NotificationURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetNotificationURIOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationURI) {
		return nil, false
	}
	return o.NotificationURI, true
}

// HasNotificationURI returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasNotificationURI() bool {
	if o != nil && !IsNil(o.NotificationURI) {
		return true
	}

	return false
}

// SetNotificationURI gets a reference to the given string and assigns it to the NotificationURI field.
func (o *NnwdafEventsSubscription) SetNotificationURI(v string) {
	o.NotificationURI = &v
}

// GetNotifCorrId returns the NotifCorrId field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetNotifCorrId() string {
	if o == nil || IsNil(o.NotifCorrId) {
		var ret string
		return ret
	}
	return *o.NotifCorrId
}

// GetNotifCorrIdOk returns a tuple with the NotifCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetNotifCorrIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifCorrId) {
		return nil, false
	}
	return o.NotifCorrId, true
}

// HasNotifCorrId returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasNotifCorrId() bool {
	if o != nil && !IsNil(o.NotifCorrId) {
		return true
	}

	return false
}

// SetNotifCorrId gets a reference to the given string and assigns it to the NotifCorrId field.
func (o *NnwdafEventsSubscription) SetNotifCorrId(v string) {
	o.NotifCorrId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NnwdafEventsSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetEventNotifications returns the EventNotifications field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetEventNotifications() []EventNotification {
	if o == nil || IsNil(o.EventNotifications) {
		var ret []EventNotification
		return ret
	}
	return o.EventNotifications
}

// GetEventNotificationsOk returns a tuple with the EventNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetEventNotificationsOk() ([]EventNotification, bool) {
	if o == nil || IsNil(o.EventNotifications) {
		return nil, false
	}
	return o.EventNotifications, true
}

// HasEventNotifications returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasEventNotifications() bool {
	if o != nil && !IsNil(o.EventNotifications) {
		return true
	}

	return false
}

// SetEventNotifications gets a reference to the given []EventNotification and assigns it to the EventNotifications field.
func (o *NnwdafEventsSubscription) SetEventNotifications(v []EventNotification) {
	o.EventNotifications = v
}

// GetFailEventReports returns the FailEventReports field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetFailEventReports() []FailureEventInfo {
	if o == nil || IsNil(o.FailEventReports) {
		var ret []FailureEventInfo
		return ret
	}
	return o.FailEventReports
}

// GetFailEventReportsOk returns a tuple with the FailEventReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetFailEventReportsOk() ([]FailureEventInfo, bool) {
	if o == nil || IsNil(o.FailEventReports) {
		return nil, false
	}
	return o.FailEventReports, true
}

// HasFailEventReports returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasFailEventReports() bool {
	if o != nil && !IsNil(o.FailEventReports) {
		return true
	}

	return false
}

// SetFailEventReports gets a reference to the given []FailureEventInfo and assigns it to the FailEventReports field.
func (o *NnwdafEventsSubscription) SetFailEventReports(v []FailureEventInfo) {
	o.FailEventReports = v
}

// GetPrevSub returns the PrevSub field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetPrevSub() PrevSubInfo {
	if o == nil || IsNil(o.PrevSub) {
		var ret PrevSubInfo
		return ret
	}
	return *o.PrevSub
}

// GetPrevSubOk returns a tuple with the PrevSub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetPrevSubOk() (*PrevSubInfo, bool) {
	if o == nil || IsNil(o.PrevSub) {
		return nil, false
	}
	return o.PrevSub, true
}

// HasPrevSub returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasPrevSub() bool {
	if o != nil && !IsNil(o.PrevSub) {
		return true
	}

	return false
}

// SetPrevSub gets a reference to the given PrevSubInfo and assigns it to the PrevSub field.
func (o *NnwdafEventsSubscription) SetPrevSub(v PrevSubInfo) {
	o.PrevSub = &v
}

// GetConsNfInfo returns the ConsNfInfo field value if set, zero value otherwise.
func (o *NnwdafEventsSubscription) GetConsNfInfo() ConsumerNfInformation {
	if o == nil || IsNil(o.ConsNfInfo) {
		var ret ConsumerNfInformation
		return ret
	}
	return *o.ConsNfInfo
}

// GetConsNfInfoOk returns a tuple with the ConsNfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NnwdafEventsSubscription) GetConsNfInfoOk() (*ConsumerNfInformation, bool) {
	if o == nil || IsNil(o.ConsNfInfo) {
		return nil, false
	}
	return o.ConsNfInfo, true
}

// HasConsNfInfo returns a boolean if a field has been set.
func (o *NnwdafEventsSubscription) HasConsNfInfo() bool {
	if o != nil && !IsNil(o.ConsNfInfo) {
		return true
	}

	return false
}

// SetConsNfInfo gets a reference to the given ConsumerNfInformation and assigns it to the ConsNfInfo field.
func (o *NnwdafEventsSubscription) SetConsNfInfo(v ConsumerNfInformation) {
	o.ConsNfInfo = &v
}

func (o NnwdafEventsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NnwdafEventsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventSubscriptions"] = o.EventSubscriptions
	if !IsNil(o.EvtReq) {
		toSerialize["evtReq"] = o.EvtReq
	}
	if !IsNil(o.NotificationURI) {
		toSerialize["notificationURI"] = o.NotificationURI
	}
	if !IsNil(o.NotifCorrId) {
		toSerialize["notifCorrId"] = o.NotifCorrId
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.EventNotifications) {
		toSerialize["eventNotifications"] = o.EventNotifications
	}
	if !IsNil(o.FailEventReports) {
		toSerialize["failEventReports"] = o.FailEventReports
	}
	if !IsNil(o.PrevSub) {
		toSerialize["prevSub"] = o.PrevSub
	}
	if !IsNil(o.ConsNfInfo) {
		toSerialize["consNfInfo"] = o.ConsNfInfo
	}
	return toSerialize, nil
}

type NullableNnwdafEventsSubscription struct {
	value *NnwdafEventsSubscription
	isSet bool
}

func (v NullableNnwdafEventsSubscription) Get() *NnwdafEventsSubscription {
	return v.value
}

func (v *NullableNnwdafEventsSubscription) Set(val *NnwdafEventsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNnwdafEventsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNnwdafEventsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNnwdafEventsSubscription(val *NnwdafEventsSubscription) *NullableNnwdafEventsSubscription {
	return &NullableNnwdafEventsSubscription{value: val, isSet: true}
}

func (v NullableNnwdafEventsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNnwdafEventsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
