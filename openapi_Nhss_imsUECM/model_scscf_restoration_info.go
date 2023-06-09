/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUECM

import (
	"encoding/json"
	"time"
)

// checks if the ScscfRestorationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScscfRestorationInfo{}

// ScscfRestorationInfo S-CSCF restoration information
type ScscfRestorationInfo struct {
	// IMS Private Identity of the UE
	UserName        *string           `json:"userName,omitempty"`
	RestorationInfo []RestorationInfo `json:"restorationInfo,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegistrationTimeOut     *time.Time               `json:"registrationTimeOut,omitempty"`
	SipAuthenticationScheme *SipAuthenticationScheme `json:"sipAuthenticationScheme,omitempty"`
}

// NewScscfRestorationInfo instantiates a new ScscfRestorationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScscfRestorationInfo() *ScscfRestorationInfo {
	this := ScscfRestorationInfo{}
	return &this
}

// NewScscfRestorationInfoWithDefaults instantiates a new ScscfRestorationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScscfRestorationInfoWithDefaults() *ScscfRestorationInfo {
	this := ScscfRestorationInfo{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScscfRestorationInfo) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfo) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScscfRestorationInfo) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScscfRestorationInfo) SetUserName(v string) {
	o.UserName = &v
}

// GetRestorationInfo returns the RestorationInfo field value if set, zero value otherwise.
func (o *ScscfRestorationInfo) GetRestorationInfo() []RestorationInfo {
	if o == nil || IsNil(o.RestorationInfo) {
		var ret []RestorationInfo
		return ret
	}
	return o.RestorationInfo
}

// GetRestorationInfoOk returns a tuple with the RestorationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfo) GetRestorationInfoOk() ([]RestorationInfo, bool) {
	if o == nil || IsNil(o.RestorationInfo) {
		return nil, false
	}
	return o.RestorationInfo, true
}

// HasRestorationInfo returns a boolean if a field has been set.
func (o *ScscfRestorationInfo) HasRestorationInfo() bool {
	if o != nil && !IsNil(o.RestorationInfo) {
		return true
	}

	return false
}

// SetRestorationInfo gets a reference to the given []RestorationInfo and assigns it to the RestorationInfo field.
func (o *ScscfRestorationInfo) SetRestorationInfo(v []RestorationInfo) {
	o.RestorationInfo = v
}

// GetRegistrationTimeOut returns the RegistrationTimeOut field value if set, zero value otherwise.
func (o *ScscfRestorationInfo) GetRegistrationTimeOut() time.Time {
	if o == nil || IsNil(o.RegistrationTimeOut) {
		var ret time.Time
		return ret
	}
	return *o.RegistrationTimeOut
}

// GetRegistrationTimeOutOk returns a tuple with the RegistrationTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfo) GetRegistrationTimeOutOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RegistrationTimeOut) {
		return nil, false
	}
	return o.RegistrationTimeOut, true
}

// HasRegistrationTimeOut returns a boolean if a field has been set.
func (o *ScscfRestorationInfo) HasRegistrationTimeOut() bool {
	if o != nil && !IsNil(o.RegistrationTimeOut) {
		return true
	}

	return false
}

// SetRegistrationTimeOut gets a reference to the given time.Time and assigns it to the RegistrationTimeOut field.
func (o *ScscfRestorationInfo) SetRegistrationTimeOut(v time.Time) {
	o.RegistrationTimeOut = &v
}

// GetSipAuthenticationScheme returns the SipAuthenticationScheme field value if set, zero value otherwise.
func (o *ScscfRestorationInfo) GetSipAuthenticationScheme() SipAuthenticationScheme {
	if o == nil || IsNil(o.SipAuthenticationScheme) {
		var ret SipAuthenticationScheme
		return ret
	}
	return *o.SipAuthenticationScheme
}

// GetSipAuthenticationSchemeOk returns a tuple with the SipAuthenticationScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScscfRestorationInfo) GetSipAuthenticationSchemeOk() (*SipAuthenticationScheme, bool) {
	if o == nil || IsNil(o.SipAuthenticationScheme) {
		return nil, false
	}
	return o.SipAuthenticationScheme, true
}

// HasSipAuthenticationScheme returns a boolean if a field has been set.
func (o *ScscfRestorationInfo) HasSipAuthenticationScheme() bool {
	if o != nil && !IsNil(o.SipAuthenticationScheme) {
		return true
	}

	return false
}

// SetSipAuthenticationScheme gets a reference to the given SipAuthenticationScheme and assigns it to the SipAuthenticationScheme field.
func (o *ScscfRestorationInfo) SetSipAuthenticationScheme(v SipAuthenticationScheme) {
	o.SipAuthenticationScheme = &v
}

func (o ScscfRestorationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScscfRestorationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.RestorationInfo) {
		toSerialize["restorationInfo"] = o.RestorationInfo
	}
	if !IsNil(o.RegistrationTimeOut) {
		toSerialize["registrationTimeOut"] = o.RegistrationTimeOut
	}
	if !IsNil(o.SipAuthenticationScheme) {
		toSerialize["sipAuthenticationScheme"] = o.SipAuthenticationScheme
	}
	return toSerialize, nil
}

type NullableScscfRestorationInfo struct {
	value *ScscfRestorationInfo
	isSet bool
}

func (v NullableScscfRestorationInfo) Get() *ScscfRestorationInfo {
	return v.value
}

func (v *NullableScscfRestorationInfo) Set(val *ScscfRestorationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableScscfRestorationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableScscfRestorationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScscfRestorationInfo(val *ScscfRestorationInfo) *NullableScscfRestorationInfo {
	return &NullableScscfRestorationInfo{value: val, isSet: true}
}

func (v NullableScscfRestorationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScscfRestorationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
