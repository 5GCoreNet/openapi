/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"time"
)

// checks if the AuthorizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationInfo{}

// AuthorizationInfo Represents NIDD authorization information.
type AuthorizationInfo struct {
	Snssai Snssai `json:"snssai"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation"`
	// String providing an URI formatted according to RFC 3986.
	AuthUpdateCallbackUri string `json:"authUpdateCallbackUri"`
	AfId *string `json:"afId,omitempty"`
	// Identity of the NEF
	NefId *string `json:"nefId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	ContextInfo *ContextInfo `json:"contextInfo,omitempty"`
}

// NewAuthorizationInfo instantiates a new AuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationInfo(snssai Snssai, dnn string, mtcProviderInformation string, authUpdateCallbackUri string) *AuthorizationInfo {
	this := AuthorizationInfo{}
	this.Snssai = snssai
	this.Dnn = dnn
	this.MtcProviderInformation = mtcProviderInformation
	this.AuthUpdateCallbackUri = authUpdateCallbackUri
	return &this
}

// NewAuthorizationInfoWithDefaults instantiates a new AuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationInfoWithDefaults() *AuthorizationInfo {
	this := AuthorizationInfo{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *AuthorizationInfo) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *AuthorizationInfo) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetDnn returns the Dnn field value
func (o *AuthorizationInfo) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *AuthorizationInfo) SetDnn(v string) {
	o.Dnn = v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value
func (o *AuthorizationInfo) GetMtcProviderInformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MtcProviderInformation, true
}

// SetMtcProviderInformation sets field value
func (o *AuthorizationInfo) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = v
}

// GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field value
func (o *AuthorizationInfo) GetAuthUpdateCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthUpdateCallbackUri
}

// GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field value
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetAuthUpdateCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthUpdateCallbackUri, true
}

// SetAuthUpdateCallbackUri sets field value
func (o *AuthorizationInfo) SetAuthUpdateCallbackUri(v string) {
	o.AuthUpdateCallbackUri = v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *AuthorizationInfo) GetAfId() string {
	if o == nil || isNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetAfIdOk() (*string, bool) {
	if o == nil || isNil(o.AfId) {
		return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *AuthorizationInfo) HasAfId() bool {
	if o != nil && !isNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *AuthorizationInfo) SetAfId(v string) {
	o.AfId = &v
}

// GetNefId returns the NefId field value if set, zero value otherwise.
func (o *AuthorizationInfo) GetNefId() string {
	if o == nil || isNil(o.NefId) {
		var ret string
		return ret
	}
	return *o.NefId
}

// GetNefIdOk returns a tuple with the NefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetNefIdOk() (*string, bool) {
	if o == nil || isNil(o.NefId) {
		return nil, false
	}
	return o.NefId, true
}

// HasNefId returns a boolean if a field has been set.
func (o *AuthorizationInfo) HasNefId() bool {
	if o != nil && !isNil(o.NefId) {
		return true
	}

	return false
}

// SetNefId gets a reference to the given string and assigns it to the NefId field.
func (o *AuthorizationInfo) SetNefId(v string) {
	o.NefId = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *AuthorizationInfo) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *AuthorizationInfo) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *AuthorizationInfo) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetContextInfo returns the ContextInfo field value if set, zero value otherwise.
func (o *AuthorizationInfo) GetContextInfo() ContextInfo {
	if o == nil || isNil(o.ContextInfo) {
		var ret ContextInfo
		return ret
	}
	return *o.ContextInfo
}

// GetContextInfoOk returns a tuple with the ContextInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationInfo) GetContextInfoOk() (*ContextInfo, bool) {
	if o == nil || isNil(o.ContextInfo) {
		return nil, false
	}
	return o.ContextInfo, true
}

// HasContextInfo returns a boolean if a field has been set.
func (o *AuthorizationInfo) HasContextInfo() bool {
	if o != nil && !isNil(o.ContextInfo) {
		return true
	}

	return false
}

// SetContextInfo gets a reference to the given ContextInfo and assigns it to the ContextInfo field.
func (o *AuthorizationInfo) SetContextInfo(v ContextInfo) {
	o.ContextInfo = &v
}

func (o AuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snssai"] = o.Snssai
	toSerialize["dnn"] = o.Dnn
	toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	toSerialize["authUpdateCallbackUri"] = o.AuthUpdateCallbackUri
	if !isNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.NefId) {
		toSerialize["nefId"] = o.NefId
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !isNil(o.ContextInfo) {
		toSerialize["contextInfo"] = o.ContextInfo
	}
	return toSerialize, nil
}

type NullableAuthorizationInfo struct {
	value *AuthorizationInfo
	isSet bool
}

func (v NullableAuthorizationInfo) Get() *AuthorizationInfo {
	return v.value
}

func (v *NullableAuthorizationInfo) Set(val *AuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationInfo(val *AuthorizationInfo) *NullableAuthorizationInfo {
	return &NullableAuthorizationInfo{value: val, isSet: true}
}

func (v NullableAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


