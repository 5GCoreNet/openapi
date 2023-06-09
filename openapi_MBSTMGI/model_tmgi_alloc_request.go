/*
3gpp-mbs-tmgi

API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSTMGI

import (
	"encoding/json"
)

// checks if the TmgiAllocRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TmgiAllocRequest{}

// TmgiAllocRequest Represents the full set of parameters to initiate an MBS TMGI(s) allocation request or the refresh of the expiry time of already allocated TMGI(s).
type TmgiAllocRequest struct {
	AfId       string       `json:"afId"`
	TmgiParams TmgiAllocate `json:"tmgiParams"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationUri         *string             `json:"notificationUri,omitempty"`
	RequestTestNotification *bool               `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig      *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewTmgiAllocRequest instantiates a new TmgiAllocRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiAllocRequest(afId string, tmgiParams TmgiAllocate) *TmgiAllocRequest {
	this := TmgiAllocRequest{}
	this.AfId = afId
	this.TmgiParams = tmgiParams
	return &this
}

// NewTmgiAllocRequestWithDefaults instantiates a new TmgiAllocRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiAllocRequestWithDefaults() *TmgiAllocRequest {
	this := TmgiAllocRequest{}
	return &this
}

// GetAfId returns the AfId field value
func (o *TmgiAllocRequest) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetAfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *TmgiAllocRequest) SetAfId(v string) {
	o.AfId = v
}

// GetTmgiParams returns the TmgiParams field value
func (o *TmgiAllocRequest) GetTmgiParams() TmgiAllocate {
	if o == nil {
		var ret TmgiAllocate
		return ret
	}

	return o.TmgiParams
}

// GetTmgiParamsOk returns a tuple with the TmgiParams field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetTmgiParamsOk() (*TmgiAllocate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TmgiParams, true
}

// SetTmgiParams sets field value
func (o *TmgiAllocRequest) SetTmgiParams(v TmgiAllocate) {
	o.TmgiParams = v
}

// GetNotificationUri returns the NotificationUri field value if set, zero value otherwise.
func (o *TmgiAllocRequest) GetNotificationUri() string {
	if o == nil || IsNil(o.NotificationUri) {
		var ret string
		return ret
	}
	return *o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUri) {
		return nil, false
	}
	return o.NotificationUri, true
}

// HasNotificationUri returns a boolean if a field has been set.
func (o *TmgiAllocRequest) HasNotificationUri() bool {
	if o != nil && !IsNil(o.NotificationUri) {
		return true
	}

	return false
}

// SetNotificationUri gets a reference to the given string and assigns it to the NotificationUri field.
func (o *TmgiAllocRequest) SetNotificationUri(v string) {
	o.NotificationUri = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *TmgiAllocRequest) GetRequestTestNotification() bool {
	if o == nil || IsNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *TmgiAllocRequest) HasRequestTestNotification() bool {
	if o != nil && !IsNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *TmgiAllocRequest) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *TmgiAllocRequest) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *TmgiAllocRequest) HasWebsockNotifConfig() bool {
	if o != nil && !IsNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *TmgiAllocRequest) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *TmgiAllocRequest) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *TmgiAllocRequest) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *TmgiAllocRequest) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o TmgiAllocRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TmgiAllocRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afId"] = o.AfId
	toSerialize["tmgiParams"] = o.TmgiParams
	if !IsNil(o.NotificationUri) {
		toSerialize["notificationUri"] = o.NotificationUri
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

type NullableTmgiAllocRequest struct {
	value *TmgiAllocRequest
	isSet bool
}

func (v NullableTmgiAllocRequest) Get() *TmgiAllocRequest {
	return v.value
}

func (v *NullableTmgiAllocRequest) Set(val *TmgiAllocRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiAllocRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiAllocRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiAllocRequest(val *TmgiAllocRequest) *NullableTmgiAllocRequest {
	return &NullableTmgiAllocRequest{value: val, isSet: true}
}

func (v NullableTmgiAllocRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiAllocRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
