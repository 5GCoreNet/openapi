/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_Authentication

import (
	"encoding/json"
)

// checks if the UAVAuthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UAVAuthResponse{}

// UAVAuthResponse UAV auth response data
type UAVAuthResponse struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi          *string         `json:"gpsi,omitempty"`
	AuthContainer []AuthContainer `json:"authContainer,omitempty"`
	// Deprecated
	AuthMsg *string `json:"authMsg,omitempty"`
	// Deprecated
	AuthResult     *UAVAuthResponseAuthResult `json:"authResult,omitempty"`
	ServiceLevelId *string                    `json:"serviceLevelId,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	AuthSessAmbr  *string `json:"authSessAmbr,omitempty"`
	AuthProfIndex *string `json:"authProfIndex,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewUAVAuthResponse instantiates a new UAVAuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAVAuthResponse() *UAVAuthResponse {
	this := UAVAuthResponse{}
	return &this
}

// NewUAVAuthResponseWithDefaults instantiates a new UAVAuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUAVAuthResponseWithDefaults() *UAVAuthResponse {
	this := UAVAuthResponse{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UAVAuthResponse) SetGpsi(v string) {
	o.Gpsi = &v
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

// GetAuthMsg returns the AuthMsg field value if set, zero value otherwise.
// Deprecated
func (o *UAVAuthResponse) GetAuthMsg() string {
	if o == nil || IsNil(o.AuthMsg) {
		var ret string
		return ret
	}
	return *o.AuthMsg
}

// GetAuthMsgOk returns a tuple with the AuthMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UAVAuthResponse) GetAuthMsgOk() (*string, bool) {
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

// SetAuthMsg gets a reference to the given string and assigns it to the AuthMsg field.
// Deprecated
func (o *UAVAuthResponse) SetAuthMsg(v string) {
	o.AuthMsg = &v
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

// GetAuthSessAmbr returns the AuthSessAmbr field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetAuthSessAmbr() string {
	if o == nil || IsNil(o.AuthSessAmbr) {
		var ret string
		return ret
	}
	return *o.AuthSessAmbr
}

// GetAuthSessAmbrOk returns a tuple with the AuthSessAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetAuthSessAmbrOk() (*string, bool) {
	if o == nil || IsNil(o.AuthSessAmbr) {
		return nil, false
	}
	return o.AuthSessAmbr, true
}

// HasAuthSessAmbr returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasAuthSessAmbr() bool {
	if o != nil && !IsNil(o.AuthSessAmbr) {
		return true
	}

	return false
}

// SetAuthSessAmbr gets a reference to the given string and assigns it to the AuthSessAmbr field.
func (o *UAVAuthResponse) SetAuthSessAmbr(v string) {
	o.AuthSessAmbr = &v
}

// GetAuthProfIndex returns the AuthProfIndex field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetAuthProfIndex() string {
	if o == nil || IsNil(o.AuthProfIndex) {
		var ret string
		return ret
	}
	return *o.AuthProfIndex
}

// GetAuthProfIndexOk returns a tuple with the AuthProfIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetAuthProfIndexOk() (*string, bool) {
	if o == nil || IsNil(o.AuthProfIndex) {
		return nil, false
	}
	return o.AuthProfIndex, true
}

// HasAuthProfIndex returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasAuthProfIndex() bool {
	if o != nil && !IsNil(o.AuthProfIndex) {
		return true
	}

	return false
}

// SetAuthProfIndex gets a reference to the given string and assigns it to the AuthProfIndex field.
func (o *UAVAuthResponse) SetAuthProfIndex(v string) {
	o.AuthProfIndex = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *UAVAuthResponse) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthResponse) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *UAVAuthResponse) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *UAVAuthResponse) SetSuppFeat(v string) {
	o.SuppFeat = &v
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
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.AuthContainer) {
		toSerialize["authContainer"] = o.AuthContainer
	}
	if !IsNil(o.AuthMsg) {
		toSerialize["authMsg"] = o.AuthMsg
	}
	if !IsNil(o.AuthResult) {
		toSerialize["authResult"] = o.AuthResult
	}
	if !IsNil(o.ServiceLevelId) {
		toSerialize["serviceLevelId"] = o.ServiceLevelId
	}
	if !IsNil(o.AuthSessAmbr) {
		toSerialize["authSessAmbr"] = o.AuthSessAmbr
	}
	if !IsNil(o.AuthProfIndex) {
		toSerialize["authProfIndex"] = o.AuthProfIndex
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
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
