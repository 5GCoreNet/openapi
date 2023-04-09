/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
	"time"
)

// checks if the EasDiscoverySubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasDiscoverySubscription{}

// EasDiscoverySubscription Represents an Individual EAS Discovery Subscription resource.
type EasDiscoverySubscription struct {
	// Represents a unique identifier of the EEC.
	EecId string `json:"eecId"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	UeId               *string               `json:"ueId,omitempty"`
	EasEventType       EASDiscEventIDs       `json:"easEventType"`
	EasDiscoveryFilter *EasDiscoveryFilter   `json:"easDiscoveryFilter,omitempty"`
	EasDynInfoFilter   *EasDynamicInfoFilter `json:"easDynInfoFilter,omitempty"`
	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EasSvcContinuity []ACRScenario `json:"easSvcContinuity,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// Set to true by Subscriber to request the ECS to send a test notification. Set to false or omitted otherwise.
	RequestTestNotification *bool               `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig      *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewEasDiscoverySubscription instantiates a new EasDiscoverySubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDiscoverySubscription(eecId string, easEventType EASDiscEventIDs) *EasDiscoverySubscription {
	this := EasDiscoverySubscription{}
	this.EecId = eecId
	this.EasEventType = easEventType
	return &this
}

// NewEasDiscoverySubscriptionWithDefaults instantiates a new EasDiscoverySubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDiscoverySubscriptionWithDefaults() *EasDiscoverySubscription {
	this := EasDiscoverySubscription{}
	return &this
}

// GetEecId returns the EecId field value
func (o *EasDiscoverySubscription) GetEecId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EecId
}

// GetEecIdOk returns a tuple with the EecId field value
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetEecIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EecId, true
}

// SetEecId sets field value
func (o *EasDiscoverySubscription) SetEecId(v string) {
	o.EecId = v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetUeId() string {
	if o == nil || IsNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetUeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasUeId() bool {
	if o != nil && !IsNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *EasDiscoverySubscription) SetUeId(v string) {
	o.UeId = &v
}

// GetEasEventType returns the EasEventType field value
func (o *EasDiscoverySubscription) GetEasEventType() EASDiscEventIDs {
	if o == nil {
		var ret EASDiscEventIDs
		return ret
	}

	return o.EasEventType
}

// GetEasEventTypeOk returns a tuple with the EasEventType field value
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetEasEventTypeOk() (*EASDiscEventIDs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasEventType, true
}

// SetEasEventType sets field value
func (o *EasDiscoverySubscription) SetEasEventType(v EASDiscEventIDs) {
	o.EasEventType = v
}

// GetEasDiscoveryFilter returns the EasDiscoveryFilter field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetEasDiscoveryFilter() EasDiscoveryFilter {
	if o == nil || IsNil(o.EasDiscoveryFilter) {
		var ret EasDiscoveryFilter
		return ret
	}
	return *o.EasDiscoveryFilter
}

// GetEasDiscoveryFilterOk returns a tuple with the EasDiscoveryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetEasDiscoveryFilterOk() (*EasDiscoveryFilter, bool) {
	if o == nil || IsNil(o.EasDiscoveryFilter) {
		return nil, false
	}
	return o.EasDiscoveryFilter, true
}

// HasEasDiscoveryFilter returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasEasDiscoveryFilter() bool {
	if o != nil && !IsNil(o.EasDiscoveryFilter) {
		return true
	}

	return false
}

// SetEasDiscoveryFilter gets a reference to the given EasDiscoveryFilter and assigns it to the EasDiscoveryFilter field.
func (o *EasDiscoverySubscription) SetEasDiscoveryFilter(v EasDiscoveryFilter) {
	o.EasDiscoveryFilter = &v
}

// GetEasDynInfoFilter returns the EasDynInfoFilter field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetEasDynInfoFilter() EasDynamicInfoFilter {
	if o == nil || IsNil(o.EasDynInfoFilter) {
		var ret EasDynamicInfoFilter
		return ret
	}
	return *o.EasDynInfoFilter
}

// GetEasDynInfoFilterOk returns a tuple with the EasDynInfoFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetEasDynInfoFilterOk() (*EasDynamicInfoFilter, bool) {
	if o == nil || IsNil(o.EasDynInfoFilter) {
		return nil, false
	}
	return o.EasDynInfoFilter, true
}

// HasEasDynInfoFilter returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasEasDynInfoFilter() bool {
	if o != nil && !IsNil(o.EasDynInfoFilter) {
		return true
	}

	return false
}

// SetEasDynInfoFilter gets a reference to the given EasDynamicInfoFilter and assigns it to the EasDynInfoFilter field.
func (o *EasDiscoverySubscription) SetEasDynInfoFilter(v EasDynamicInfoFilter) {
	o.EasDynInfoFilter = &v
}

// GetEasSvcContinuity returns the EasSvcContinuity field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetEasSvcContinuity() []ACRScenario {
	if o == nil || IsNil(o.EasSvcContinuity) {
		var ret []ACRScenario
		return ret
	}
	return o.EasSvcContinuity
}

// GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetEasSvcContinuityOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.EasSvcContinuity) {
		return nil, false
	}
	return o.EasSvcContinuity, true
}

// HasEasSvcContinuity returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasEasSvcContinuity() bool {
	if o != nil && !IsNil(o.EasSvcContinuity) {
		return true
	}

	return false
}

// SetEasSvcContinuity gets a reference to the given []ACRScenario and assigns it to the EasSvcContinuity field.
func (o *EasDiscoverySubscription) SetEasSvcContinuity(v []ACRScenario) {
	o.EasSvcContinuity = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetExpTime() time.Time {
	if o == nil || IsNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasExpTime() bool {
	if o != nil && !IsNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *EasDiscoverySubscription) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *EasDiscoverySubscription) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetRequestTestNotification() bool {
	if o == nil || IsNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasRequestTestNotification() bool {
	if o != nil && !IsNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *EasDiscoverySubscription) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasWebsockNotifConfig() bool {
	if o != nil && !IsNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *EasDiscoverySubscription) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *EasDiscoverySubscription) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EasDiscoverySubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *EasDiscoverySubscription) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *EasDiscoverySubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o EasDiscoverySubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasDiscoverySubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eecId"] = o.EecId
	if !IsNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	toSerialize["easEventType"] = o.EasEventType
	if !IsNil(o.EasDiscoveryFilter) {
		toSerialize["easDiscoveryFilter"] = o.EasDiscoveryFilter
	}
	if !IsNil(o.EasDynInfoFilter) {
		toSerialize["easDynInfoFilter"] = o.EasDynInfoFilter
	}
	if !IsNil(o.EasSvcContinuity) {
		toSerialize["easSvcContinuity"] = o.EasSvcContinuity
	}
	if !IsNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
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

type NullableEasDiscoverySubscription struct {
	value *EasDiscoverySubscription
	isSet bool
}

func (v NullableEasDiscoverySubscription) Get() *EasDiscoverySubscription {
	return v.value
}

func (v *NullableEasDiscoverySubscription) Set(val *EasDiscoverySubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDiscoverySubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDiscoverySubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDiscoverySubscription(val *EasDiscoverySubscription) *NullableEasDiscoverySubscription {
	return &NullableEasDiscoverySubscription{value: val, isSet: true}
}

func (v NullableEasDiscoverySubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDiscoverySubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
