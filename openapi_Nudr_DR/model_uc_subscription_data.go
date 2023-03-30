/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the UcSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UcSubscriptionData{}

// UcSubscriptionData Contains the User Consent Subscription Data.
type UcSubscriptionData struct {
	// A map(list of key-value pairs) where user consent purpose serves as key of user consent
	UserConsentPerPurposeList *map[string]UserConsent `json:"userConsentPerPurposeList,omitempty"`
}

// NewUcSubscriptionData instantiates a new UcSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcSubscriptionData() *UcSubscriptionData {
	this := UcSubscriptionData{}
	return &this
}

// NewUcSubscriptionDataWithDefaults instantiates a new UcSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcSubscriptionDataWithDefaults() *UcSubscriptionData {
	this := UcSubscriptionData{}
	return &this
}

// GetUserConsentPerPurposeList returns the UserConsentPerPurposeList field value if set, zero value otherwise.
func (o *UcSubscriptionData) GetUserConsentPerPurposeList() map[string]UserConsent {
	if o == nil || IsNil(o.UserConsentPerPurposeList) {
		var ret map[string]UserConsent
		return ret
	}
	return *o.UserConsentPerPurposeList
}

// GetUserConsentPerPurposeListOk returns a tuple with the UserConsentPerPurposeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcSubscriptionData) GetUserConsentPerPurposeListOk() (*map[string]UserConsent, bool) {
	if o == nil || IsNil(o.UserConsentPerPurposeList) {
		return nil, false
	}
	return o.UserConsentPerPurposeList, true
}

// HasUserConsentPerPurposeList returns a boolean if a field has been set.
func (o *UcSubscriptionData) HasUserConsentPerPurposeList() bool {
	if o != nil && !IsNil(o.UserConsentPerPurposeList) {
		return true
	}

	return false
}

// SetUserConsentPerPurposeList gets a reference to the given map[string]UserConsent and assigns it to the UserConsentPerPurposeList field.
func (o *UcSubscriptionData) SetUserConsentPerPurposeList(v map[string]UserConsent) {
	o.UserConsentPerPurposeList = &v
}

func (o UcSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UcSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserConsentPerPurposeList) {
		toSerialize["userConsentPerPurposeList"] = o.UserConsentPerPurposeList
	}
	return toSerialize, nil
}

type NullableUcSubscriptionData struct {
	value *UcSubscriptionData
	isSet bool
}

func (v NullableUcSubscriptionData) Get() *UcSubscriptionData {
	return v.value
}

func (v *NullableUcSubscriptionData) Set(val *UcSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableUcSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableUcSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcSubscriptionData(val *UcSubscriptionData) *NullableUcSubscriptionData {
	return &NullableUcSubscriptionData{value: val, isSet: true}
}

func (v NullableUcSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


