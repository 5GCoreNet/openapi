/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_Authentication

import (
	"encoding/json"
)

// checks if the UAVAuthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UAVAuthResponse{}

// UAVAuthResponse UAV auth response data
type UAVAuthResponse struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi           string  `json:"gpsi"`
	ServiceLevelId *string `json:"serviceLevelId,omitempty"`
	// Deprecated
	AuthMsg       *UAVAuthInfoAuthMsg `json:"authMsg,omitempty"`
	AuthContainer []AuthContainer     `json:"authContainer,omitempty"`
	// Deprecated
	AuthResult   *UAVAuthResponseAuthResult `json:"authResult,omitempty"`
	NotifyCorrId *string                    `json:"notifyCorrId,omitempty"`
}

// NewUAVAuthResponse instantiates a new UAVAuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAVAuthResponse(gpsi string) *UAVAuthResponse {
	this := UAVAuthResponse{}
	this.Gpsi = gpsi
	return &this
}

// NewUAVAuthResponseWithDefaults instantiates a new UAVAuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUAVAuthResponseWithDefaults() *UAVAuthResponse {
	this := UAVAuthResponse{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *UAVAuthResponse) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetGpsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *UAVAuthResponse) SetGpsi(v string) {
	o.Gpsi = v
}

// GetServiceLevelId returns the ServiceLevelId field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetServiceLevelId() string {
	if o == nil || IsNil(o.ServiceLevelId) {
		var ret string
		return ret
	}
	return *o.ServiceLevelId
}

// GetServiceLevelIdOk returns a tuple with the ServiceLevelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetServiceLevelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLevelId) {
		return nil, false
	}
	return o.ServiceLevelId, true
}

// HasServiceLevelId returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasServiceLevelId() bool {
	if o != nil && !IsNil(o.ServiceLevelId) {
		return true
	}

	return false
}

// SetServiceLevelId gets a reference to the given string and assigns it to the ServiceLevelId field.
func (o *UAVAuthResponse) SetServiceLevelId(v string) {
	o.ServiceLevelId = &v
}

// GetAuthMsg returns the AuthMsg field value if set, zero value otherwise.
// Deprecated
func (o *UAVAuthResponse) GetAuthMsg() UAVAuthInfoAuthMsg {
	if o == nil || IsNil(o.AuthMsg) {
		var ret UAVAuthInfoAuthMsg
		return ret
	}
	return *o.AuthMsg
}

// GetAuthMsgOk returns a tuple with the AuthMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UAVAuthResponse) GetAuthMsgOk() (*UAVAuthInfoAuthMsg, bool) {
	if o == nil || IsNil(o.AuthMsg) {
		return nil, false
	}
	return o.AuthMsg, true
}

// HasAuthMsg returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasAuthMsg() bool {
	if o != nil && !IsNil(o.AuthMsg) {
		return true
	}

	return false
}

// SetAuthMsg gets a reference to the given UAVAuthInfoAuthMsg and assigns it to the AuthMsg field.
// Deprecated
func (o *UAVAuthResponse) SetAuthMsg(v UAVAuthInfoAuthMsg) {
	o.AuthMsg = &v
}

// GetAuthContainer returns the AuthContainer field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetAuthContainer() []AuthContainer {
	if o == nil || IsNil(o.AuthContainer) {
		var ret []AuthContainer
		return ret
	}
	return o.AuthContainer
}

// GetAuthContainerOk returns a tuple with the AuthContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetAuthContainerOk() ([]AuthContainer, bool) {
	if o == nil || IsNil(o.AuthContainer) {
		return nil, false
	}
	return o.AuthContainer, true
}

// HasAuthContainer returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasAuthContainer() bool {
	if o != nil && !IsNil(o.AuthContainer) {
		return true
	}

	return false
}

// SetAuthContainer gets a reference to the given []AuthContainer and assigns it to the AuthContainer field.
func (o *UAVAuthResponse) SetAuthContainer(v []AuthContainer) {
	o.AuthContainer = v
}

// GetAuthResult returns the AuthResult field value if set, zero value otherwise.
// Deprecated
func (o *UAVAuthResponse) GetAuthResult() UAVAuthResponseAuthResult {
	if o == nil || IsNil(o.AuthResult) {
		var ret UAVAuthResponseAuthResult
		return ret
	}
	return *o.AuthResult
}

// GetAuthResultOk returns a tuple with the AuthResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UAVAuthResponse) GetAuthResultOk() (*UAVAuthResponseAuthResult, bool) {
	if o == nil || IsNil(o.AuthResult) {
		return nil, false
	}
	return o.AuthResult, true
}

// HasAuthResult returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasAuthResult() bool {
	if o != nil && !IsNil(o.AuthResult) {
		return true
	}

	return false
}

// SetAuthResult gets a reference to the given UAVAuthResponseAuthResult and assigns it to the AuthResult field.
// Deprecated
func (o *UAVAuthResponse) SetAuthResult(v UAVAuthResponseAuthResult) {
	o.AuthResult = &v
}

// GetNotifyCorrId returns the NotifyCorrId field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetNotifyCorrId() string {
	if o == nil || IsNil(o.NotifyCorrId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrId
}

// GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetNotifyCorrIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrId) {
		return nil, false
	}
	return o.NotifyCorrId, true
}

// HasNotifyCorrId returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasNotifyCorrId() bool {
	if o != nil && !IsNil(o.NotifyCorrId) {
		return true
	}

	return false
}

// SetNotifyCorrId gets a reference to the given string and assigns it to the NotifyCorrId field.
func (o *UAVAuthResponse) SetNotifyCorrId(v string) {
	o.NotifyCorrId = &v
}

func (o UAVAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UAVAuthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gpsi"] = o.Gpsi
	if !IsNil(o.ServiceLevelId) {
		toSerialize["serviceLevelId"] = o.ServiceLevelId
	}
	if !IsNil(o.AuthMsg) {
		toSerialize["authMsg"] = o.AuthMsg
	}
	if !IsNil(o.AuthContainer) {
		toSerialize["authContainer"] = o.AuthContainer
	}
	if !IsNil(o.AuthResult) {
		toSerialize["authResult"] = o.AuthResult
	}
	if !IsNil(o.NotifyCorrId) {
		toSerialize["notifyCorrId"] = o.NotifyCorrId
	}
	return toSerialize, nil
}

type NullableUAVAuthResponse struct {
	value *UAVAuthResponse
	isSet bool
}

func (v NullableUAVAuthResponse) Get() *UAVAuthResponse {
	return v.value
}

func (v *NullableUAVAuthResponse) Set(val *UAVAuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUAVAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUAVAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUAVAuthResponse(val *UAVAuthResponse) *NullableUAVAuthResponse {
	return &NullableUAVAuthResponse{value: val, isSet: true}
}

func (v NullableUAVAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUAVAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
