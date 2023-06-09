/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_UERCM

import (
	"encoding/json"
	"time"
)

// checks if the CreatedSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedSubscription{}

// CreatedSubscription Data for a creating a subscription response
type CreatedSubscription struct {
	DicEntryId int32 `json:"dicEntryId"`
	// string with format 'date-time' as defined in OpenAPI.
	ConfirmedExpires *time.Time `json:"confirmedExpires,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCreatedSubscription instantiates a new CreatedSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedSubscription(dicEntryId int32) *CreatedSubscription {
	this := CreatedSubscription{}
	this.DicEntryId = dicEntryId
	return &this
}

// NewCreatedSubscriptionWithDefaults instantiates a new CreatedSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedSubscriptionWithDefaults() *CreatedSubscription {
	this := CreatedSubscription{}
	return &this
}

// GetDicEntryId returns the DicEntryId field value
func (o *CreatedSubscription) GetDicEntryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DicEntryId
}

// GetDicEntryIdOk returns a tuple with the DicEntryId field value
// and a boolean to check if the value has been set.
func (o *CreatedSubscription) GetDicEntryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DicEntryId, true
}

// SetDicEntryId sets field value
func (o *CreatedSubscription) SetDicEntryId(v int32) {
	o.DicEntryId = v
}

// GetConfirmedExpires returns the ConfirmedExpires field value if set, zero value otherwise.
func (o *CreatedSubscription) GetConfirmedExpires() time.Time {
	if o == nil || IsNil(o.ConfirmedExpires) {
		var ret time.Time
		return ret
	}
	return *o.ConfirmedExpires
}

// GetConfirmedExpiresOk returns a tuple with the ConfirmedExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedSubscription) GetConfirmedExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConfirmedExpires) {
		return nil, false
	}
	return o.ConfirmedExpires, true
}

// HasConfirmedExpires returns a boolean if a field has been set.
func (o *CreatedSubscription) HasConfirmedExpires() bool {
	if o != nil && !IsNil(o.ConfirmedExpires) {
		return true
	}

	return false
}

// SetConfirmedExpires gets a reference to the given time.Time and assigns it to the ConfirmedExpires field.
func (o *CreatedSubscription) SetConfirmedExpires(v time.Time) {
	o.ConfirmedExpires = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CreatedSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CreatedSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CreatedSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CreatedSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dicEntryId"] = o.DicEntryId
	if !IsNil(o.ConfirmedExpires) {
		toSerialize["confirmedExpires"] = o.ConfirmedExpires
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableCreatedSubscription struct {
	value *CreatedSubscription
	isSet bool
}

func (v NullableCreatedSubscription) Get() *CreatedSubscription {
	return v.value
}

func (v *NullableCreatedSubscription) Set(val *CreatedSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedSubscription(val *CreatedSubscription) *NullableCreatedSubscription {
	return &NullableCreatedSubscription{value: val, isSet: true}
}

func (v NullableCreatedSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
