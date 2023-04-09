/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the UeN1N2InfoSubscriptionCreatedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeN1N2InfoSubscriptionCreatedData{}

// UeN1N2InfoSubscriptionCreatedData Data for the created subscription for UE specific N1 and/or N2 information notification
type UeN1N2InfoSubscriptionCreatedData struct {
	N1n2NotifySubscriptionId string `json:"n1n2NotifySubscriptionId"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewUeN1N2InfoSubscriptionCreatedData instantiates a new UeN1N2InfoSubscriptionCreatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeN1N2InfoSubscriptionCreatedData(n1n2NotifySubscriptionId string) *UeN1N2InfoSubscriptionCreatedData {
	this := UeN1N2InfoSubscriptionCreatedData{}
	this.N1n2NotifySubscriptionId = n1n2NotifySubscriptionId
	return &this
}

// NewUeN1N2InfoSubscriptionCreatedDataWithDefaults instantiates a new UeN1N2InfoSubscriptionCreatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeN1N2InfoSubscriptionCreatedDataWithDefaults() *UeN1N2InfoSubscriptionCreatedData {
	this := UeN1N2InfoSubscriptionCreatedData{}
	return &this
}

// GetN1n2NotifySubscriptionId returns the N1n2NotifySubscriptionId field value
func (o *UeN1N2InfoSubscriptionCreatedData) GetN1n2NotifySubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N1n2NotifySubscriptionId
}

// GetN1n2NotifySubscriptionIdOk returns a tuple with the N1n2NotifySubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UeN1N2InfoSubscriptionCreatedData) GetN1n2NotifySubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N1n2NotifySubscriptionId, true
}

// SetN1n2NotifySubscriptionId sets field value
func (o *UeN1N2InfoSubscriptionCreatedData) SetN1n2NotifySubscriptionId(v string) {
	o.N1n2NotifySubscriptionId = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *UeN1N2InfoSubscriptionCreatedData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeN1N2InfoSubscriptionCreatedData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *UeN1N2InfoSubscriptionCreatedData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *UeN1N2InfoSubscriptionCreatedData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o UeN1N2InfoSubscriptionCreatedData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeN1N2InfoSubscriptionCreatedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n1n2NotifySubscriptionId"] = o.N1n2NotifySubscriptionId
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableUeN1N2InfoSubscriptionCreatedData struct {
	value *UeN1N2InfoSubscriptionCreatedData
	isSet bool
}

func (v NullableUeN1N2InfoSubscriptionCreatedData) Get() *UeN1N2InfoSubscriptionCreatedData {
	return v.value
}

func (v *NullableUeN1N2InfoSubscriptionCreatedData) Set(val *UeN1N2InfoSubscriptionCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeN1N2InfoSubscriptionCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeN1N2InfoSubscriptionCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeN1N2InfoSubscriptionCreatedData(val *UeN1N2InfoSubscriptionCreatedData) *NullableUeN1N2InfoSubscriptionCreatedData {
	return &NullableUeN1N2InfoSubscriptionCreatedData{value: val, isSet: true}
}

func (v NullableUeN1N2InfoSubscriptionCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeN1N2InfoSubscriptionCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
