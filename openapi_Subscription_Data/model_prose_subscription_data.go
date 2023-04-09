/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the ProseSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseSubscriptionData{}

// ProseSubscriptionData Contains the ProSe Subscription Data.
type ProseSubscriptionData struct {
	ProseServiceAuth *ProseServiceAuth `json:"proseServiceAuth,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	NrUePc5Ambr      *string            `json:"nrUePc5Ambr,omitempty"`
	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
}

// NewProseSubscriptionData instantiates a new ProseSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseSubscriptionData() *ProseSubscriptionData {
	this := ProseSubscriptionData{}
	return &this
}

// NewProseSubscriptionDataWithDefaults instantiates a new ProseSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseSubscriptionDataWithDefaults() *ProseSubscriptionData {
	this := ProseSubscriptionData{}
	return &this
}

// GetProseServiceAuth returns the ProseServiceAuth field value if set, zero value otherwise.
func (o *ProseSubscriptionData) GetProseServiceAuth() ProseServiceAuth {
	if o == nil || IsNil(o.ProseServiceAuth) {
		var ret ProseServiceAuth
		return ret
	}
	return *o.ProseServiceAuth
}

// GetProseServiceAuthOk returns a tuple with the ProseServiceAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseSubscriptionData) GetProseServiceAuthOk() (*ProseServiceAuth, bool) {
	if o == nil || IsNil(o.ProseServiceAuth) {
		return nil, false
	}
	return o.ProseServiceAuth, true
}

// HasProseServiceAuth returns a boolean if a field has been set.
func (o *ProseSubscriptionData) HasProseServiceAuth() bool {
	if o != nil && !IsNil(o.ProseServiceAuth) {
		return true
	}

	return false
}

// SetProseServiceAuth gets a reference to the given ProseServiceAuth and assigns it to the ProseServiceAuth field.
func (o *ProseSubscriptionData) SetProseServiceAuth(v ProseServiceAuth) {
	o.ProseServiceAuth = &v
}

// GetNrUePc5Ambr returns the NrUePc5Ambr field value if set, zero value otherwise.
func (o *ProseSubscriptionData) GetNrUePc5Ambr() string {
	if o == nil || IsNil(o.NrUePc5Ambr) {
		var ret string
		return ret
	}
	return *o.NrUePc5Ambr
}

// GetNrUePc5AmbrOk returns a tuple with the NrUePc5Ambr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseSubscriptionData) GetNrUePc5AmbrOk() (*string, bool) {
	if o == nil || IsNil(o.NrUePc5Ambr) {
		return nil, false
	}
	return o.NrUePc5Ambr, true
}

// HasNrUePc5Ambr returns a boolean if a field has been set.
func (o *ProseSubscriptionData) HasNrUePc5Ambr() bool {
	if o != nil && !IsNil(o.NrUePc5Ambr) {
		return true
	}

	return false
}

// SetNrUePc5Ambr gets a reference to the given string and assigns it to the NrUePc5Ambr field.
func (o *ProseSubscriptionData) SetNrUePc5Ambr(v string) {
	o.NrUePc5Ambr = &v
}

// GetProseAllowedPlmn returns the ProseAllowedPlmn field value if set, zero value otherwise.
func (o *ProseSubscriptionData) GetProseAllowedPlmn() []ProSeAllowedPlmn {
	if o == nil || IsNil(o.ProseAllowedPlmn) {
		var ret []ProSeAllowedPlmn
		return ret
	}
	return o.ProseAllowedPlmn
}

// GetProseAllowedPlmnOk returns a tuple with the ProseAllowedPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseSubscriptionData) GetProseAllowedPlmnOk() ([]ProSeAllowedPlmn, bool) {
	if o == nil || IsNil(o.ProseAllowedPlmn) {
		return nil, false
	}
	return o.ProseAllowedPlmn, true
}

// HasProseAllowedPlmn returns a boolean if a field has been set.
func (o *ProseSubscriptionData) HasProseAllowedPlmn() bool {
	if o != nil && !IsNil(o.ProseAllowedPlmn) {
		return true
	}

	return false
}

// SetProseAllowedPlmn gets a reference to the given []ProSeAllowedPlmn and assigns it to the ProseAllowedPlmn field.
func (o *ProseSubscriptionData) SetProseAllowedPlmn(v []ProSeAllowedPlmn) {
	o.ProseAllowedPlmn = v
}

func (o ProseSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProseServiceAuth) {
		toSerialize["proseServiceAuth"] = o.ProseServiceAuth
	}
	if !IsNil(o.NrUePc5Ambr) {
		toSerialize["nrUePc5Ambr"] = o.NrUePc5Ambr
	}
	if !IsNil(o.ProseAllowedPlmn) {
		toSerialize["proseAllowedPlmn"] = o.ProseAllowedPlmn
	}
	return toSerialize, nil
}

type NullableProseSubscriptionData struct {
	value *ProseSubscriptionData
	isSet bool
}

func (v NullableProseSubscriptionData) Get() *ProseSubscriptionData {
	return v.value
}

func (v *NullableProseSubscriptionData) Set(val *ProseSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableProseSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableProseSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseSubscriptionData(val *ProseSubscriptionData) *NullableProseSubscriptionData {
	return &NullableProseSubscriptionData{value: val, isSet: true}
}

func (v NullableProseSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
