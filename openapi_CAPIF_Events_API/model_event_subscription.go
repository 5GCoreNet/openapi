/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Events_API

import (
	"encoding/json"
)

// checks if the EventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSubscription{}

// EventSubscription Represents an individual CAPIF Event Subscription resource.
type EventSubscription struct {
	// Subscribed events
	Events []CAPIFEvent `json:"events"`
	// Subscribed event filters
	EventFilters []CAPIFEventFilter `json:"eventFilters,omitempty"`
	EventReq *ReportingInformation `json:"eventReq,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination string `json:"notificationDestination"`
	// Set to true by Subscriber to request the CAPIF core function to send a test notification as defined in in clause 7.6. Set to false or omitted otherwise. 
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewEventSubscription instantiates a new EventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscription(events []CAPIFEvent, notificationDestination string) *EventSubscription {
	this := EventSubscription{}
	this.Events = events
	this.NotificationDestination = notificationDestination
	return &this
}

// NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscriptionWithDefaults() *EventSubscription {
	this := EventSubscription{}
	return &this
}

// GetEvents returns the Events field value
func (o *EventSubscription) GetEvents() []CAPIFEvent {
	if o == nil {
		var ret []CAPIFEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEventsOk() ([]CAPIFEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventSubscription) SetEvents(v []CAPIFEvent) {
	o.Events = v
}

// GetEventFilters returns the EventFilters field value if set, zero value otherwise.
func (o *EventSubscription) GetEventFilters() []CAPIFEventFilter {
	if o == nil || isNil(o.EventFilters) {
		var ret []CAPIFEventFilter
		return ret
	}
	return o.EventFilters
}

// GetEventFiltersOk returns a tuple with the EventFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEventFiltersOk() ([]CAPIFEventFilter, bool) {
	if o == nil || isNil(o.EventFilters) {
		return nil, false
	}
	return o.EventFilters, true
}

// HasEventFilters returns a boolean if a field has been set.
func (o *EventSubscription) HasEventFilters() bool {
	if o != nil && !isNil(o.EventFilters) {
		return true
	}

	return false
}

// SetEventFilters gets a reference to the given []CAPIFEventFilter and assigns it to the EventFilters field.
func (o *EventSubscription) SetEventFilters(v []CAPIFEventFilter) {
	o.EventFilters = v
}

// GetEventReq returns the EventReq field value if set, zero value otherwise.
func (o *EventSubscription) GetEventReq() ReportingInformation {
	if o == nil || isNil(o.EventReq) {
		var ret ReportingInformation
		return ret
	}
	return *o.EventReq
}

// GetEventReqOk returns a tuple with the EventReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEventReqOk() (*ReportingInformation, bool) {
	if o == nil || isNil(o.EventReq) {
		return nil, false
	}
	return o.EventReq, true
}

// HasEventReq returns a boolean if a field has been set.
func (o *EventSubscription) HasEventReq() bool {
	if o != nil && !isNil(o.EventReq) {
		return true
	}

	return false
}

// SetEventReq gets a reference to the given ReportingInformation and assigns it to the EventReq field.
func (o *EventSubscription) SetEventReq(v ReportingInformation) {
	o.EventReq = &v
}

// GetNotificationDestination returns the NotificationDestination field value
func (o *EventSubscription) GetNotificationDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetNotificationDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDestination, true
}

// SetNotificationDestination sets field value
func (o *EventSubscription) SetNotificationDestination(v string) {
	o.NotificationDestination = v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *EventSubscription) GetRequestTestNotification() bool {
	if o == nil || isNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *EventSubscription) HasRequestTestNotification() bool {
	if o != nil && !isNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *EventSubscription) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *EventSubscription) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || isNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || isNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *EventSubscription) HasWebsockNotifConfig() bool {
	if o != nil && !isNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *EventSubscription) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EventSubscription) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EventSubscription) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EventSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o EventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !isNil(o.EventFilters) {
		toSerialize["eventFilters"] = o.EventFilters
	}
	if !isNil(o.EventReq) {
		toSerialize["eventReq"] = o.EventReq
	}
	toSerialize["notificationDestination"] = o.NotificationDestination
	if !isNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !isNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableEventSubscription struct {
	value *EventSubscription
	isSet bool
}

func (v NullableEventSubscription) Get() *EventSubscription {
	return v.value
}

func (v *NullableEventSubscription) Set(val *EventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscription(val *EventSubscription) *NullableEventSubscription {
	return &NullableEventSubscription{value: val, isSet: true}
}

func (v NullableEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


