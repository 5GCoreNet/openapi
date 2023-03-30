/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"encoding/json"
	"time"
)

// checks if the TimeSyncExposureSubsc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSyncExposureSubsc{}

// TimeSyncExposureSubsc Contains requested parameters for the subscription to the notification of time synchronization capability. 
type TimeSyncExposureSubsc struct {
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExterGroupId *string `json:"exterGroupId,omitempty"`
	// Contains a list of UE for which the time synchronization capabilities is requested. 
	Gpsis []string `json:"gpsis,omitempty"`
	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"false\" is used, if not present. 
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	// Identifies a service on behalf of which the AF is issuing the request.
	AfServiceId *string `json:"afServiceId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Notification Correlation ID assigned by the NF service consumer.
	SubsNotifId string `json:"subsNotifId"`
	// String providing an URI formatted according to RFC 3986.
	SubsNotifUri string `json:"subsNotifUri"`
	// Subscribed events
	SubscribedEvents []SubscribedEvent `json:"subscribedEvents,omitempty"`
	// Contains the filter conditions to match for notifying the event(s) of time synchronization capabilities for a list of UE(s). 
	EventFilters []EventFilter `json:"eventFilters,omitempty"`
	NotifMethod *NotificationMethod `json:"notifMethod,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr *int32 `json:"maxReportNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise. 
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewTimeSyncExposureSubsc instantiates a new TimeSyncExposureSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSyncExposureSubsc(subsNotifId string, subsNotifUri string) *TimeSyncExposureSubsc {
	this := TimeSyncExposureSubsc{}
	this.SubsNotifId = subsNotifId
	this.SubsNotifUri = subsNotifUri
	return &this
}

// NewTimeSyncExposureSubscWithDefaults instantiates a new TimeSyncExposureSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSyncExposureSubscWithDefaults() *TimeSyncExposureSubsc {
	this := TimeSyncExposureSubsc{}
	return &this
}

// GetExterGroupId returns the ExterGroupId field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetExterGroupId() string {
	if o == nil || IsNil(o.ExterGroupId) {
		var ret string
		return ret
	}
	return *o.ExterGroupId
}

// GetExterGroupIdOk returns a tuple with the ExterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetExterGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExterGroupId) {
		return nil, false
	}
	return o.ExterGroupId, true
}

// HasExterGroupId returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasExterGroupId() bool {
	if o != nil && !IsNil(o.ExterGroupId) {
		return true
	}

	return false
}

// SetExterGroupId gets a reference to the given string and assigns it to the ExterGroupId field.
func (o *TimeSyncExposureSubsc) SetExterGroupId(v string) {
	o.ExterGroupId = &v
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetGpsis() []string {
	if o == nil || IsNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetGpsisOk() ([]string, bool) {
	if o == nil || IsNil(o.Gpsis) {
		return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasGpsis() bool {
	if o != nil && !IsNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *TimeSyncExposureSubsc) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetAnyUeInd() bool {
	if o == nil || IsNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AnyUeInd) {
		return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasAnyUeInd() bool {
	if o != nil && !IsNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *TimeSyncExposureSubsc) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetAfServiceId returns the AfServiceId field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetAfServiceId() string {
	if o == nil || IsNil(o.AfServiceId) {
		var ret string
		return ret
	}
	return *o.AfServiceId
}

// GetAfServiceIdOk returns a tuple with the AfServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetAfServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfServiceId) {
		return nil, false
	}
	return o.AfServiceId, true
}

// HasAfServiceId returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasAfServiceId() bool {
	if o != nil && !IsNil(o.AfServiceId) {
		return true
	}

	return false
}

// SetAfServiceId gets a reference to the given string and assigns it to the AfServiceId field.
func (o *TimeSyncExposureSubsc) SetAfServiceId(v string) {
	o.AfServiceId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *TimeSyncExposureSubsc) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *TimeSyncExposureSubsc) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetSubsNotifId returns the SubsNotifId field value
func (o *TimeSyncExposureSubsc) GetSubsNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubsNotifId
}

// GetSubsNotifIdOk returns a tuple with the SubsNotifId field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetSubsNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubsNotifId, true
}

// SetSubsNotifId sets field value
func (o *TimeSyncExposureSubsc) SetSubsNotifId(v string) {
	o.SubsNotifId = v
}

// GetSubsNotifUri returns the SubsNotifUri field value
func (o *TimeSyncExposureSubsc) GetSubsNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubsNotifUri
}

// GetSubsNotifUriOk returns a tuple with the SubsNotifUri field value
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetSubsNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubsNotifUri, true
}

// SetSubsNotifUri sets field value
func (o *TimeSyncExposureSubsc) SetSubsNotifUri(v string) {
	o.SubsNotifUri = v
}

// GetSubscribedEvents returns the SubscribedEvents field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetSubscribedEvents() []SubscribedEvent {
	if o == nil || IsNil(o.SubscribedEvents) {
		var ret []SubscribedEvent
		return ret
	}
	return o.SubscribedEvents
}

// GetSubscribedEventsOk returns a tuple with the SubscribedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetSubscribedEventsOk() ([]SubscribedEvent, bool) {
	if o == nil || IsNil(o.SubscribedEvents) {
		return nil, false
	}
	return o.SubscribedEvents, true
}

// HasSubscribedEvents returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasSubscribedEvents() bool {
	if o != nil && !IsNil(o.SubscribedEvents) {
		return true
	}

	return false
}

// SetSubscribedEvents gets a reference to the given []SubscribedEvent and assigns it to the SubscribedEvents field.
func (o *TimeSyncExposureSubsc) SetSubscribedEvents(v []SubscribedEvent) {
	o.SubscribedEvents = v
}

// GetEventFilters returns the EventFilters field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetEventFilters() []EventFilter {
	if o == nil || IsNil(o.EventFilters) {
		var ret []EventFilter
		return ret
	}
	return o.EventFilters
}

// GetEventFiltersOk returns a tuple with the EventFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetEventFiltersOk() ([]EventFilter, bool) {
	if o == nil || IsNil(o.EventFilters) {
		return nil, false
	}
	return o.EventFilters, true
}

// HasEventFilters returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasEventFilters() bool {
	if o != nil && !IsNil(o.EventFilters) {
		return true
	}

	return false
}

// SetEventFilters gets a reference to the given []EventFilter and assigns it to the EventFilters field.
func (o *TimeSyncExposureSubsc) SetEventFilters(v []EventFilter) {
	o.EventFilters = v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetNotifMethod() NotificationMethod {
	if o == nil || IsNil(o.NotifMethod) {
		var ret NotificationMethod
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetNotifMethodOk() (*NotificationMethod, bool) {
	if o == nil || IsNil(o.NotifMethod) {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasNotifMethod() bool {
	if o != nil && !IsNil(o.NotifMethod) {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given NotificationMethod and assigns it to the NotifMethod field.
func (o *TimeSyncExposureSubsc) SetNotifMethod(v NotificationMethod) {
	o.NotifMethod = &v
}

// GetMaxReportNbr returns the MaxReportNbr field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetMaxReportNbr() int32 {
	if o == nil || IsNil(o.MaxReportNbr) {
		var ret int32
		return ret
	}
	return *o.MaxReportNbr
}

// GetMaxReportNbrOk returns a tuple with the MaxReportNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetMaxReportNbrOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxReportNbr) {
		return nil, false
	}
	return o.MaxReportNbr, true
}

// HasMaxReportNbr returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasMaxReportNbr() bool {
	if o != nil && !IsNil(o.MaxReportNbr) {
		return true
	}

	return false
}

// SetMaxReportNbr gets a reference to the given int32 and assigns it to the MaxReportNbr field.
func (o *TimeSyncExposureSubsc) SetMaxReportNbr(v int32) {
	o.MaxReportNbr = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *TimeSyncExposureSubsc) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetRepPeriod() int32 {
	if o == nil || IsNil(o.RepPeriod) {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetRepPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RepPeriod) {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasRepPeriod() bool {
	if o != nil && !IsNil(o.RepPeriod) {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *TimeSyncExposureSubsc) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetRequestTestNotification() bool {
	if o == nil || IsNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasRequestTestNotification() bool {
	if o != nil && !IsNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *TimeSyncExposureSubsc) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasWebsockNotifConfig() bool {
	if o != nil && !IsNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *TimeSyncExposureSubsc) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *TimeSyncExposureSubsc) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSyncExposureSubsc) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *TimeSyncExposureSubsc) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *TimeSyncExposureSubsc) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o TimeSyncExposureSubsc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSyncExposureSubsc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExterGroupId) {
		toSerialize["exterGroupId"] = o.ExterGroupId
	}
	if !IsNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !IsNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !IsNil(o.AfServiceId) {
		toSerialize["afServiceId"] = o.AfServiceId
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["subsNotifId"] = o.SubsNotifId
	toSerialize["subsNotifUri"] = o.SubsNotifUri
	if !IsNil(o.SubscribedEvents) {
		toSerialize["subscribedEvents"] = o.SubscribedEvents
	}
	if !IsNil(o.EventFilters) {
		toSerialize["eventFilters"] = o.EventFilters
	}
	if !IsNil(o.NotifMethod) {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if !IsNil(o.MaxReportNbr) {
		toSerialize["maxReportNbr"] = o.MaxReportNbr
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.RepPeriod) {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if !IsNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !IsNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableTimeSyncExposureSubsc struct {
	value *TimeSyncExposureSubsc
	isSet bool
}

func (v NullableTimeSyncExposureSubsc) Get() *TimeSyncExposureSubsc {
	return v.value
}

func (v *NullableTimeSyncExposureSubsc) Set(val *TimeSyncExposureSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSyncExposureSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSyncExposureSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSyncExposureSubsc(val *TimeSyncExposureSubsc) *NullableTimeSyncExposureSubsc {
	return &NullableTimeSyncExposureSubsc{value: val, isSet: true}
}

func (v NullableTimeSyncExposureSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSyncExposureSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


