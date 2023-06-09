/*
Nnef_PFDmanagement Service API

Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_PFDmanagement

import (
	"encoding/json"
)

// checks if the PfdSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PfdSubscription{}

// PfdSubscription Represents a PFD subscription.
type PfdSubscription struct {
	ApplicationIds []string `json:"applicationIds,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures"`
}

// NewPfdSubscription instantiates a new PfdSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfdSubscription(notifyUri string, supportedFeatures string) *PfdSubscription {
	this := PfdSubscription{}
	this.NotifyUri = notifyUri
	this.SupportedFeatures = supportedFeatures
	return &this
}

// NewPfdSubscriptionWithDefaults instantiates a new PfdSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdSubscriptionWithDefaults() *PfdSubscription {
	this := PfdSubscription{}
	return &this
}

// GetApplicationIds returns the ApplicationIds field value if set, zero value otherwise.
func (o *PfdSubscription) GetApplicationIds() []string {
	if o == nil || IsNil(o.ApplicationIds) {
		var ret []string
		return ret
	}
	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdSubscription) GetApplicationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationIds) {
		return nil, false
	}
	return o.ApplicationIds, true
}

// HasApplicationIds returns a boolean if a field has been set.
func (o *PfdSubscription) HasApplicationIds() bool {
	if o != nil && !IsNil(o.ApplicationIds) {
		return true
	}

	return false
}

// SetApplicationIds gets a reference to the given []string and assigns it to the ApplicationIds field.
func (o *PfdSubscription) SetApplicationIds(v []string) {
	o.ApplicationIds = v
}

// GetNotifyUri returns the NotifyUri field value
func (o *PfdSubscription) GetNotifyUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifyUri
}

// GetNotifyUriOk returns a tuple with the NotifyUri field value
// and a boolean to check if the value has been set.
func (o *PfdSubscription) GetNotifyUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyUri, true
}

// SetNotifyUri sets field value
func (o *PfdSubscription) SetNotifyUri(v string) {
	o.NotifyUri = v
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *PfdSubscription) GetSupportedFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *PfdSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *PfdSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = v
}

func (o PfdSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PfdSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationIds) {
		toSerialize["applicationIds"] = o.ApplicationIds
	}
	toSerialize["notifyUri"] = o.NotifyUri
	toSerialize["supportedFeatures"] = o.SupportedFeatures
	return toSerialize, nil
}

type NullablePfdSubscription struct {
	value *PfdSubscription
	isSet bool
}

func (v NullablePfdSubscription) Get() *PfdSubscription {
	return v.value
}

func (v *NullablePfdSubscription) Set(val *PfdSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdSubscription(val *PfdSubscription) *NullablePfdSubscription {
	return &NullablePfdSubscription{value: val, isSet: true}
}

func (v NullablePfdSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
