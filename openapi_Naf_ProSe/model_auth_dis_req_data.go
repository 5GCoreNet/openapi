/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
)

// checks if the AuthDisReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthDisReqData{}

// AuthDisReqData Represents Data used to request the authorization for a UE of a 5G ProSe Direct  Discovery request.
type AuthDisReqData struct {
	AuthRequestType AuthRequestType `json:"authRequestType"`
	ProseAppId      []string        `json:"proseAppId,omitempty"`
	// contains the allowed number of suffixes.
	AllowedSuffixNum *int32 `json:"allowedSuffixNum,omitempty"`
	// Contains the Application Level Container.
	AppLevelContainer *string `json:"appLevelContainer,omitempty"`
	// Contains the RPAUID.
	Rpauid *string `json:"rpauid,omitempty"`
	// Contains the RPAUID.
	TargetRpauid *string `json:"targetRpauid,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	AuthUpdateCallbackUri *string `json:"authUpdateCallbackUri,omitempty"`
}

// NewAuthDisReqData instantiates a new AuthDisReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthDisReqData(authRequestType AuthRequestType) *AuthDisReqData {
	this := AuthDisReqData{}
	this.AuthRequestType = authRequestType
	return &this
}

// NewAuthDisReqDataWithDefaults instantiates a new AuthDisReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDisReqDataWithDefaults() *AuthDisReqData {
	this := AuthDisReqData{}
	return &this
}

// GetAuthRequestType returns the AuthRequestType field value
func (o *AuthDisReqData) GetAuthRequestType() AuthRequestType {
	if o == nil {
		var ret AuthRequestType
		return ret
	}

	return o.AuthRequestType
}

// GetAuthRequestTypeOk returns a tuple with the AuthRequestType field value
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetAuthRequestTypeOk() (*AuthRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthRequestType, true
}

// SetAuthRequestType sets field value
func (o *AuthDisReqData) SetAuthRequestType(v AuthRequestType) {
	o.AuthRequestType = v
}

// GetProseAppId returns the ProseAppId field value if set, zero value otherwise.
func (o *AuthDisReqData) GetProseAppId() []string {
	if o == nil || IsNil(o.ProseAppId) {
		var ret []string
		return ret
	}
	return o.ProseAppId
}

// GetProseAppIdOk returns a tuple with the ProseAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetProseAppIdOk() ([]string, bool) {
	if o == nil || IsNil(o.ProseAppId) {
		return nil, false
	}
	return o.ProseAppId, true
}

// HasProseAppId returns a boolean if a field has been set.
func (o *AuthDisReqData) HasProseAppId() bool {
	if o != nil && !IsNil(o.ProseAppId) {
		return true
	}

	return false
}

// SetProseAppId gets a reference to the given []string and assigns it to the ProseAppId field.
func (o *AuthDisReqData) SetProseAppId(v []string) {
	o.ProseAppId = v
}

// GetAllowedSuffixNum returns the AllowedSuffixNum field value if set, zero value otherwise.
func (o *AuthDisReqData) GetAllowedSuffixNum() int32 {
	if o == nil || IsNil(o.AllowedSuffixNum) {
		var ret int32
		return ret
	}
	return *o.AllowedSuffixNum
}

// GetAllowedSuffixNumOk returns a tuple with the AllowedSuffixNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetAllowedSuffixNumOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedSuffixNum) {
		return nil, false
	}
	return o.AllowedSuffixNum, true
}

// HasAllowedSuffixNum returns a boolean if a field has been set.
func (o *AuthDisReqData) HasAllowedSuffixNum() bool {
	if o != nil && !IsNil(o.AllowedSuffixNum) {
		return true
	}

	return false
}

// SetAllowedSuffixNum gets a reference to the given int32 and assigns it to the AllowedSuffixNum field.
func (o *AuthDisReqData) SetAllowedSuffixNum(v int32) {
	o.AllowedSuffixNum = &v
}

// GetAppLevelContainer returns the AppLevelContainer field value if set, zero value otherwise.
func (o *AuthDisReqData) GetAppLevelContainer() string {
	if o == nil || IsNil(o.AppLevelContainer) {
		var ret string
		return ret
	}
	return *o.AppLevelContainer
}

// GetAppLevelContainerOk returns a tuple with the AppLevelContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetAppLevelContainerOk() (*string, bool) {
	if o == nil || IsNil(o.AppLevelContainer) {
		return nil, false
	}
	return o.AppLevelContainer, true
}

// HasAppLevelContainer returns a boolean if a field has been set.
func (o *AuthDisReqData) HasAppLevelContainer() bool {
	if o != nil && !IsNil(o.AppLevelContainer) {
		return true
	}

	return false
}

// SetAppLevelContainer gets a reference to the given string and assigns it to the AppLevelContainer field.
func (o *AuthDisReqData) SetAppLevelContainer(v string) {
	o.AppLevelContainer = &v
}

// GetRpauid returns the Rpauid field value if set, zero value otherwise.
func (o *AuthDisReqData) GetRpauid() string {
	if o == nil || IsNil(o.Rpauid) {
		var ret string
		return ret
	}
	return *o.Rpauid
}

// GetRpauidOk returns a tuple with the Rpauid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetRpauidOk() (*string, bool) {
	if o == nil || IsNil(o.Rpauid) {
		return nil, false
	}
	return o.Rpauid, true
}

// HasRpauid returns a boolean if a field has been set.
func (o *AuthDisReqData) HasRpauid() bool {
	if o != nil && !IsNil(o.Rpauid) {
		return true
	}

	return false
}

// SetRpauid gets a reference to the given string and assigns it to the Rpauid field.
func (o *AuthDisReqData) SetRpauid(v string) {
	o.Rpauid = &v
}

// GetTargetRpauid returns the TargetRpauid field value if set, zero value otherwise.
func (o *AuthDisReqData) GetTargetRpauid() string {
	if o == nil || IsNil(o.TargetRpauid) {
		var ret string
		return ret
	}
	return *o.TargetRpauid
}

// GetTargetRpauidOk returns a tuple with the TargetRpauid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetTargetRpauidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetRpauid) {
		return nil, false
	}
	return o.TargetRpauid, true
}

// HasTargetRpauid returns a boolean if a field has been set.
func (o *AuthDisReqData) HasTargetRpauid() bool {
	if o != nil && !IsNil(o.TargetRpauid) {
		return true
	}

	return false
}

// SetTargetRpauid gets a reference to the given string and assigns it to the TargetRpauid field.
func (o *AuthDisReqData) SetTargetRpauid(v string) {
	o.TargetRpauid = &v
}

// GetAuthUpdateCallbackUri returns the AuthUpdateCallbackUri field value if set, zero value otherwise.
func (o *AuthDisReqData) GetAuthUpdateCallbackUri() string {
	if o == nil || IsNil(o.AuthUpdateCallbackUri) {
		var ret string
		return ret
	}
	return *o.AuthUpdateCallbackUri
}

// GetAuthUpdateCallbackUriOk returns a tuple with the AuthUpdateCallbackUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDisReqData) GetAuthUpdateCallbackUriOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUpdateCallbackUri) {
		return nil, false
	}
	return o.AuthUpdateCallbackUri, true
}

// HasAuthUpdateCallbackUri returns a boolean if a field has been set.
func (o *AuthDisReqData) HasAuthUpdateCallbackUri() bool {
	if o != nil && !IsNil(o.AuthUpdateCallbackUri) {
		return true
	}

	return false
}

// SetAuthUpdateCallbackUri gets a reference to the given string and assigns it to the AuthUpdateCallbackUri field.
func (o *AuthDisReqData) SetAuthUpdateCallbackUri(v string) {
	o.AuthUpdateCallbackUri = &v
}

func (o AuthDisReqData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthDisReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authRequestType"] = o.AuthRequestType
	if !IsNil(o.ProseAppId) {
		toSerialize["proseAppId"] = o.ProseAppId
	}
	if !IsNil(o.AllowedSuffixNum) {
		toSerialize["allowedSuffixNum"] = o.AllowedSuffixNum
	}
	if !IsNil(o.AppLevelContainer) {
		toSerialize["appLevelContainer"] = o.AppLevelContainer
	}
	if !IsNil(o.Rpauid) {
		toSerialize["rpauid"] = o.Rpauid
	}
	if !IsNil(o.TargetRpauid) {
		toSerialize["targetRpauid"] = o.TargetRpauid
	}
	if !IsNil(o.AuthUpdateCallbackUri) {
		toSerialize["authUpdateCallbackUri"] = o.AuthUpdateCallbackUri
	}
	return toSerialize, nil
}

type NullableAuthDisReqData struct {
	value *AuthDisReqData
	isSet bool
}

func (v NullableAuthDisReqData) Get() *AuthDisReqData {
	return v.value
}

func (v *NullableAuthDisReqData) Set(val *AuthDisReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthDisReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthDisReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthDisReqData(val *AuthDisReqData) *NullableAuthDisReqData {
	return &NullableAuthDisReqData{value: val, isSet: true}
}

func (v NullableAuthDisReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthDisReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
