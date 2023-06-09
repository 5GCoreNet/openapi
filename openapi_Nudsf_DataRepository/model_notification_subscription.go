/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
	"time"
)

// checks if the NotificationSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSubscription{}

// NotificationSubscription Definition of a notification subscription
type NotificationSubscription struct {
	ClientId ClientId `json:"clientId"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
	// String providing an URI formatted according to RFC 3986.
	ExpiryCallbackReference *string `json:"expiryCallbackReference,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ExpiryNotification *int32              `json:"expiryNotification,omitempty"`
	SubFilter          *SubscriptionFilter `json:"subFilter,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewNotificationSubscription instantiates a new NotificationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSubscription(clientId ClientId, callbackReference string) *NotificationSubscription {
	this := NotificationSubscription{}
	this.ClientId = clientId
	this.CallbackReference = callbackReference
	return &this
}

// NewNotificationSubscriptionWithDefaults instantiates a new NotificationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSubscriptionWithDefaults() *NotificationSubscription {
	this := NotificationSubscription{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *NotificationSubscription) GetClientId() ClientId {
	if o == nil {
		var ret ClientId
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetClientIdOk() (*ClientId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *NotificationSubscription) SetClientId(v ClientId) {
	o.ClientId = v
}

// GetCallbackReference returns the CallbackReference field value
func (o *NotificationSubscription) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *NotificationSubscription) SetCallbackReference(v string) {
	o.CallbackReference = v
}

// GetExpiryCallbackReference returns the ExpiryCallbackReference field value if set, zero value otherwise.
func (o *NotificationSubscription) GetExpiryCallbackReference() string {
	if o == nil || IsNil(o.ExpiryCallbackReference) {
		var ret string
		return ret
	}
	return *o.ExpiryCallbackReference
}

// GetExpiryCallbackReferenceOk returns a tuple with the ExpiryCallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetExpiryCallbackReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryCallbackReference) {
		return nil, false
	}
	return o.ExpiryCallbackReference, true
}

// HasExpiryCallbackReference returns a boolean if a field has been set.
func (o *NotificationSubscription) HasExpiryCallbackReference() bool {
	if o != nil && !IsNil(o.ExpiryCallbackReference) {
		return true
	}

	return false
}

// SetExpiryCallbackReference gets a reference to the given string and assigns it to the ExpiryCallbackReference field.
func (o *NotificationSubscription) SetExpiryCallbackReference(v string) {
	o.ExpiryCallbackReference = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *NotificationSubscription) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *NotificationSubscription) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *NotificationSubscription) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetExpiryNotification returns the ExpiryNotification field value if set, zero value otherwise.
func (o *NotificationSubscription) GetExpiryNotification() int32 {
	if o == nil || IsNil(o.ExpiryNotification) {
		var ret int32
		return ret
	}
	return *o.ExpiryNotification
}

// GetExpiryNotificationOk returns a tuple with the ExpiryNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetExpiryNotificationOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryNotification) {
		return nil, false
	}
	return o.ExpiryNotification, true
}

// HasExpiryNotification returns a boolean if a field has been set.
func (o *NotificationSubscription) HasExpiryNotification() bool {
	if o != nil && !IsNil(o.ExpiryNotification) {
		return true
	}

	return false
}

// SetExpiryNotification gets a reference to the given int32 and assigns it to the ExpiryNotification field.
func (o *NotificationSubscription) SetExpiryNotification(v int32) {
	o.ExpiryNotification = &v
}

// GetSubFilter returns the SubFilter field value if set, zero value otherwise.
func (o *NotificationSubscription) GetSubFilter() SubscriptionFilter {
	if o == nil || IsNil(o.SubFilter) {
		var ret SubscriptionFilter
		return ret
	}
	return *o.SubFilter
}

// GetSubFilterOk returns a tuple with the SubFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetSubFilterOk() (*SubscriptionFilter, bool) {
	if o == nil || IsNil(o.SubFilter) {
		return nil, false
	}
	return o.SubFilter, true
}

// HasSubFilter returns a boolean if a field has been set.
func (o *NotificationSubscription) HasSubFilter() bool {
	if o != nil && !IsNil(o.SubFilter) {
		return true
	}

	return false
}

// SetSubFilter gets a reference to the given SubscriptionFilter and assigns it to the SubFilter field.
func (o *NotificationSubscription) SetSubFilter(v SubscriptionFilter) {
	o.SubFilter = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NotificationSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NotificationSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NotificationSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o NotificationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["callbackReference"] = o.CallbackReference
	if !IsNil(o.ExpiryCallbackReference) {
		toSerialize["expiryCallbackReference"] = o.ExpiryCallbackReference
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.ExpiryNotification) {
		toSerialize["expiryNotification"] = o.ExpiryNotification
	}
	if !IsNil(o.SubFilter) {
		toSerialize["subFilter"] = o.SubFilter
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableNotificationSubscription struct {
	value *NotificationSubscription
	isSet bool
}

func (v NullableNotificationSubscription) Get() *NotificationSubscription {
	return v.value
}

func (v *NullableNotificationSubscription) Set(val *NotificationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSubscription(val *NotificationSubscription) *NullableNotificationSubscription {
	return &NullableNotificationSubscription{value: val, isSet: true}
}

func (v NullableNotificationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
